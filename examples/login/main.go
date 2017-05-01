package main

import (
	"fmt"
	"github.com/bergotorino/go-launchpad/launchpad"
	"log"
)

func main() {
	// Create a place to save and load the credentials
	// The file will be created if it not exists.
	sb := launchpad.SecretsFileBackend{File: "./launchpad.secrets.json"}

	// Get a handle to the Launchpad client. All further requests
	// are proxied through it.
	lp := launchpad.NewClient(nil, "Example Client")

	// Loginto Launchpad using the previously created secrets backend
	err := lp.LoginWith(&sb)
	if err != nil {
		log.Fatal("lp.Login: ", err)
		return
	}

	fmt.Println("Logged in to Launchpad")
}
