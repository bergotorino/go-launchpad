package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/bergotorino/go-launchpad/launchpad"
)

const lpTeam string = "ubuntu-core-service"

type Package struct {
	Name    string
	Version string
	Arch    string
}

// map snap to lp team
var snaps2team = map[string]string{
	"core18": "ubuntu-core-service",
	"core":   "snappy-dev",
}

// map core to series
var core2series = map[string]string{
	"core":   "xenial",
	"core18": "bionic",
}

var (
	DefaultSnap     = "core18"
	DefaultRevision = 1076
	DefaultArch     = "armhf"
)

func main() {
	// Parse arguments.
	// Note that the architecture is not configurable at the moment
	var (
		snap     string
		revision int
	)
	flag.StringVar(&snap, "snap", DefaultSnap, "The snap to get source for")
	flag.IntVar(&revision, "revision", DefaultRevision, "The revision to get source for")
	flag.Parse()

	// Make sure the snap name is something that is supported
	_, ok := snaps2team[snap]
	if !ok {
		fmt.Printf("Invalid snap name.\nIt must be one of: core, core18\n")
		return
	}

	// Get and parse the manifest for selected snap / revision
	furl, err := getManifestFileUrl(snap, revision)
	if err != nil {
		log.Fatal("Failed to find the URL to the manifest.")
	}
	body, err := downloadFile(furl)
	if err != nil {
		log.Fatal("Failed to download manifest")
	}

	var packages []Package

	if snap == "core" {
		packages = parseCoreManifestFile(body)
	} else {
		packages = parseCore18ManifestFile(body)
	}

	// If the manifest file cannot be parsed be verbose
	if len(packages) == 0 {
		fmt.Printf("You should know that no packages were read from the manifest file: %s\n", furl)
		return
	}

	// Connect to Launchpad, get the archive
	sb := launchpad.SecretsFileBackend{File: "./launchpad.secrets.json"}
	lp := launchpad.NewClient(nil, "Example Client")
	err = lp.LoginWith(&sb)
	if err != nil {
		log.Fatal("Failed to get lp client")
	}
	ubuntu, err := lp.Distributions("ubuntu")
	if err != nil {
		log.Fatal("Failed to get ubuntu distro")
	}
	archive, err := ubuntu.GetPrimaryArchive()
	if err != nil {
		log.Fatal("Failed to get archive", err)
	}

	// At this stage the list of binary packages that were used to build the
	// snap is known. The next step is to find out what kind of sources
	// have been used to create those binary packages. It is not easy to
	// find out, i.e. lack of single API call, however, there is a way.
	//
	// For each binary pacakge it is possible to get its publishing history
	// entry. This entry contains the reference to the build that triggered
	// the publish. From the build it is possible to learn the source
	// package used. Note that the build has no information about the
	// version of the source package however this is going to correspond to
	// the binary package version and this information is already available
	// - is coming from the manifest.

	// This is to keep the list of sources that have been used to build the
	// snap associated with the url to the source archive.
	var sources map[string]string
	sources = make(map[string]string)

	// This is to map the source package name with the binary package. This
	// kind of relation is needed to know which version of the source
	// source package should be examined.
	var s2p map[string]Package
	s2p = make(map[string]Package)

	// Traverse the list of packages and build the map of src to bin
	for _, p := range packages {
		// in some cases the package name is name:arch. make sure
		// to capture just the name here
		name := strings.Split(p.Name, ":")[0]

		pubbin, err := archive.GetPublishedBinaries(launchpad.CreateSeriesArchUrl(core2series[snap], DefaultArch), name, p.Version)
		if err != nil {
			// TODO: in case of 503 save it on a list and
			// try again
			log.Println("Failed to get published binaries ", err)
			continue
		}

		// From the publishing history of a particular binary package
		// get the right one and process.
		for _, pb := range pubbin {

			// Filter out Deleted and Superseded uploads
			if pb.Status != "Published" {
				continue
			}

			// Just pick the right version
			if pb.BinaryPackageVersion != p.Version {
				continue
			}

			// At this point we have the right entry. Lets pick
			// the build that created it.
			build, _ := lp.NewBuild(pb.BuildLink)

			// Map the source with binary
			_, ok := s2p[build.SourcePackageName]
			if !ok {
				// put everything here, mark as null so that it
				// is easy to spot missing information at the
				// end.
				sources[build.SourcePackageName] = "NULL"

				// also save the mapping
				s2p[build.SourcePackageName] = p
			}
		}
	}

	// Now for every source package get its publishing history entry
	// key is the source name, value is the Package
	for k, v := range s2p {

		res, err := archive.GetPublishedSources(launchpad.CreateSeriesUrl(core2series[snap]), k)
		if err != nil {
			log.Fatal("Failed to get published binaries", err)
		}

		for _, r := range res {
			// Filter out Deleted and Superseded uploads
			if r.Status != "Published" {
				continue
			}

			// Just pick the right version
			if r.SourcePackageVersion != v.Version {
				continue
			}

			files, err := r.SourceFileUrls()
			if err != nil {
				log.Println("Unable to get files: ", err)
				continue
			}
			for _, f := range files {
				if strings.Contains(f, ".dsc") {
					continue
				}
				sources[k] = f
			}

		}

	}

	// Print links to sources
	for k, v := range sources {
		fmt.Printf("%s %s %s\n", k, s2p[k].Version, v)
	}
}

// Based on the snap name and the revision this function tries to find the
// URL for the manifest file. Returns the url as string or error.
func getManifestFileUrl(sname string, revision int) (string, error) {
	sb := launchpad.SecretsFileBackend{File: "./launchpad.secrets.json"}

	lp := launchpad.NewClient(nil, "Example Client")
	err := lp.LoginWith(&sb)
	if err != nil {
		return "", err
	}

	team, err := lp.People(snaps2team[sname])
	if err != nil {
		return "", err
	}

	snaps := lp.Snaps()

	snap, err := snaps.GetByName(sname, team.SelfLink)
	if err != nil {
		return "", err
	}

	builds, err := snap.CompletedBuilds()
	if err != nil {
		return "", err
	}

	for _, b := range builds {
		// filter out unwanted revisions
		if b.StoreUploadRevision != revision {
			continue
		}

		// get all files produced by the build
		files, err := b.GetFileUrls()
		if err != nil {
			return "", err
		}

		// for each file produced by the build find .manifest
		// there will be only one like this so it is safe to
		// return blindly
		for _, f := range files {
			if strings.Contains(f, "manifest") {
				return f, nil
			}
		}
	}

	return "", errors.New("Manifest file not found in the build artifacts")
}

// Download file from the internet using HTTP GET
// Return contents of the file as string
func downloadFile(url string) (string, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Parse snap build manifest file and return the list of
// packages.
//
// For some reason the format of manifest file differs between core and core18.
// The interesting lines in the core18 manifest file are formatted as follows:
// ii <package> <version> <arch> <comment>
// ir <package> <version> <arch> <comment>
// On the other hand core manifest is formatted:
// <package> <version>
func parseCore18ManifestFile(file string) []Package {
	// peel out only the right lines as well as split on whitespace
	pattern := regexp.MustCompile("^[ir]")
	split := regexp.MustCompile("[ \t]+")

	var retval []Package

	scanner := bufio.NewScanner(strings.NewReader(file))
	for scanner.Scan() {
		token := scanner.Text()
		if pattern.MatchString(token) {
			tmp := split.Split(token, -1)
			if len(tmp) < 4 {
				continue
			}
			retval = append(retval, Package{
				Name:    tmp[1],
				Version: tmp[2],
				Arch:    tmp[3],
			})
		}
	}

	return retval
}

// Parse snap build manifest file and return the list of
// packages.
//
// For some reason the format of manifest file differs between core and core18.
// The interesting lines in the core18 manifest file are formatted as follows:
// ii <package> <version> <arch> <comment>
// ir <package> <version> <arch> <comment>
// On the other hand core manifest is formatted:
// <package> <version>
func parseCoreManifestFile(file string) []Package {
	// split on whitespace
	split := regexp.MustCompile("[ \t]+")

	var retval []Package

	scanner := bufio.NewScanner(strings.NewReader(file))
	for scanner.Scan() {
		token := scanner.Text()
		tmp := split.Split(token, -1)
		retval = append(retval, Package{
			Name:    tmp[0],
			Version: tmp[1],
		})
	}

	return retval
}
