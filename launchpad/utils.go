package launchpad

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/bergotorino/go-oauth/oauth"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// DecodeResponse decodes the JSON response
func DecodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		return fmt.Errorf("get %s returned status %d", resp.Request.URL, resp.StatusCode)
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	json.Unmarshal(body, &data)

	return nil
}

// Make the consumer key
func MakeConsumerKey() string {
	consumerKey := "System-wide: Ubuntu"
	hostname, err := os.Hostname()
	if err == nil {
		consumerKey += fmt.Sprintf(" (%s)", hostname)
	}
	return consumerKey
}

// Returns default oauth Client for Launchpad.net access
func DefaultOauthClient() *oauth.Client {
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://launchpad.net/+request-token",
		ResourceOwnerAuthorizationURI: "https://launchpad.net/+authorize-token",
		TokenRequestURI:               "https://launchpad.net/+access-token",
		Credentials: oauth.Credentials{
			Token:  MakeConsumerKey(),
			Secret: "",
		},
		SignatureMethod: oauth.PLAINTEXT,
	}
}
