package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bergotorino/go-launchpad/launchpad"

	"text/tabwriter"
)

var snaps = []string{
	"alsa-utils",
	"bluez",
	"build-scripts",
	"captive-redirect",
	"engineering-tests",
	"git-repo",
	"modem-manager",
	"network-manager",
	"pulseaudio",
	"se-test-tools",
	"tests-extras",
	"tpm",
	"tpm2",
	"udisks2",
	"upower",
	"wifi-ap",
	"wpa-supplicant",
}

func main() {

	secrets, err := launchpad.NewLaunchpadSecrets("./launchpad.secrets.json")
	if err != nil {
		log.Fatal(err)
	}

	lp, err := launchpad.LoginWith(*secrets, "System-wide: Ubuntu(annapurna)")
	if err != nil {
		log.Fatal("lp.Login: ", err)
		return
	}

	log.Println("Logged in to Launchpad")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	for _, s := range snaps {
		gitrepository, err := lp.GitRepositories().GetByPath("~snappy-hwe-team/snappy-hwe-snaps/+git/" + s)
		if err != nil {
			log.Fatal("Failed to get git repository")
		}
		landingcandidates, err := gitrepository.LandingCandidates()
		if err != nil {
			log.Fatal("Failed to get landing targets")
		}
		for _, mp := range landingcandidates {
			fmt.Fprintf(w, " %s\t| %s\t| %s\t| %d\t| %s\n",
				s,
				mp.RegistrantLink[33:], mp.WebLink,
				uint(time.Now().Sub(mp.DateCreated).Hours()/24+0.5),
				mp.QueueStatus)
		}
	}

	w.Flush()
}
