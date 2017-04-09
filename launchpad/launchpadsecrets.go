package launchpad

import (
	"encoding/json"
	"github.com/bergotorino/go-oauth/oauth"
	"io/ioutil"
)

type LaunchpadSecrets struct {
	file              string
	accessCredentials *oauth.Credentials
}

func NewLaunchpadSecrets(secretsFile string) (*LaunchpadSecrets, error) {
	launchpadSecrets := new(LaunchpadSecrets)

	launchpadSecrets.accessCredentials = nil
	launchpadSecrets.file = secretsFile

	data, err := ioutil.ReadFile(secretsFile)
	if err != nil {
		return launchpadSecrets, nil
	}

	var ac oauth.Credentials

	err = json.Unmarshal(data, &ac)
	if err != nil {
		return nil, err
	}

	launchpadSecrets.accessCredentials = &ac

	return launchpadSecrets, nil
}

func (l *LaunchpadSecrets) IsValid() bool {

	if l.accessCredentials == nil {
		return false
	} else {
		return true
	}
}

func (l *LaunchpadSecrets) Store() error {
	data, err := json.Marshal(l.accessCredentials)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(l.file, data, 0664)
	if err != nil {
		return err
	}

	return nil
}
