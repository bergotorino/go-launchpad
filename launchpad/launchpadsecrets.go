package launchpad

import (
	"encoding/json"
	"github.com/bergotorino/go-oauth/oauth"
	"io"
	"io/ioutil"
)

type LaunchpadSecrets struct {
	file              string
	accessCredentials *oauth.Credentials
}

func (l *LaunchpadSecrets) Write(w io.Writer) error {
	data, err := json.Marshal(l.accessCredentials)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func (l *LaunchpadSecrets) Read(r io.Reader) error {
	var data []byte
	_, err := r.Read(data)
	if err != nil {
		return err
	}

	var ac oauth.Credentials
	err = json.Unmarshal(data, &ac)
	if err != nil {
		return err
	}
	l.accessCredentials = &ac

	return nil
}

type SecretsReaderWriter interface {
	Write(p []byte) (n int, err error)
	Read(p []byte) (n int, err error)
}

type SecretsFileBackend struct {
	File string
}

func (s *SecretsFileBackend) Write(data []byte) (n int, err error) {
	err = ioutil.WriteFile(s.File, data, 0664)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

func (s *SecretsFileBackend) Read(data []byte) (n int, err error) {
	data, err = ioutil.ReadFile(s.File)
	if err != nil {
		return 0, err
	}
	return len(data), nil
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
