package launchpad

import (
	"encoding/json"
	"github.com/bergotorino/go-oauth/oauth"
	"io"
	"io/ioutil"
	"os"
)

type LaunchpadSecrets struct {
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
	var data = make([]byte, 512)
	n, err := r.Read(data)
	if err != nil {
		return err
	}

	var ac oauth.Credentials
	err = json.Unmarshal(data[:n], &ac)
	if err != nil {
		return err
	}
	l.accessCredentials = &ac

	return nil
}

type SecretsBackend interface {
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
	f, err := os.Open(s.File)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	n, err = f.Read(data)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (l *LaunchpadSecrets) IsValid() bool {

	if l.accessCredentials == nil {
		return false
	} else {
		return true
	}
}
