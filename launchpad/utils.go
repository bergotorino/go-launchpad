package launchpad

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
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
