package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bergotorino/go-launchpad/launchpad"

	"text/tabwriter"
)

func main() {
	sb := launchpad.SecretsFileBackend{File: "./launchpad.secrets.json"}

	lp := launchpad.NewClient(nil, "Example Client")
	err := lp.LoginWith(&sb)
	if err != nil {
		log.Fatal("lp.Login: ", err)
		return
	}

	team, err := lp.People("snappy-hwe-team")
	if err != nil {
		log.Fatal("lp.People: ", err)
		return
	}

	snaps, err := lp.GitRepositories().GetRepositories(team.SelfLink)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	for _, s := range snaps {
		gitrepository, err := lp.GitRepositories().GetByPath(s.DisplayName[3:])
		if err != nil {
			log.Fatal("Failed to get git repository")
		}
		landingcandidates, err := gitrepository.LandingCandidates()
		if err != nil {
			log.Fatal("Failed to get landing targets")
		}
		for _, mp := range landingcandidates {
			fmt.Fprintf(w, " %s\t| %s\t| %s\t| %d\t| %s\n",
				s.Name,
				mp.RegistrantLink[33:], mp.WebLink,
				uint(time.Now().Sub(mp.DateCreated).Hours()/24+0.5),
				mp.QueueStatus)
		}
	}

	w.Flush()
}
