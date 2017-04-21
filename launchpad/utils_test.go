package launchpad_test

import (
	. "gopkg.in/check.v1"

	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/bergotorino/go-launchpad/launchpad"
)

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TestDecodeResponseStatusCodeBad(c *C) {
	body := []byte("one two three")
	bReader := bytes.NewReader(body)
	request, _ := http.NewRequest("GET", "http://foo/bar", bReader)
	response := http.Response{StatusCode: 201, Request: request, Body: ioutil.NopCloser(bReader)}

	err := launchpad.DecodeResponse(&response, nil)
	c.Assert(err, Not(IsNil))
}

func (s *UtilsSuite) TestDecodeResponseStatusCodeGood(c *C) {
	body := []byte("{\"name\":\"Hello\",\"status\":\"good\"}")
	bReader := bytes.NewReader(body)
	request, _ := http.NewRequest("GET", "http://foo/bar", bReader)
	response := http.Response{StatusCode: 200, Request: request, Body: ioutil.NopCloser(bReader)}

	data := struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}{}

	err := launchpad.DecodeResponse(&response, &data)
	c.Assert(err, IsNil)
	c.Assert(data.Name, Equals, "Hello")
	c.Assert(data.Status, Equals, "good")
}

func (s *UtilsSuite) TestDecodeResponseStatusCodeGoodGzipBad(c *C) {
	body := []byte("{\"name\":\"Hello\",\"status\":\"good\"}")
	bReader := bytes.NewReader(body)
	request, _ := http.NewRequest("GET", "http://foo/bar", bReader)
	response := http.Response{StatusCode: 200, Request: request, Body: ioutil.NopCloser(bReader)}

	response.Header = make(map[string][]string)

	response.Header.Set("Content-Encoding", "gzip")

	data := struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}{}

	err := launchpad.DecodeResponse(&response, &data)
	c.Assert(err, Not(IsNil))
}

/*
func (s *UtilsSuite) TestDecodeResponseStatusCodeGoodGzipGood(c *C) {
	body := []byte("{\"name\":\"Hello\",\"status\":\"good\"}")

	var gzBody bytes.Buffer
	gz := gzip.NewWriter(&gzBody)
	_, err := gz.Write(body)
	if err != nil {
		panic(err)
	}
	gz.Close()

	gBdy := gzBody.Bytes()
	bReader := bytes.NewReader(gBdy)

	request, _ := http.NewRequest("GET", "http://foo/bar", nil)
	response := http.Response{StatusCode: 200, Request: request, Body: ioutil.NopCloser(bReader)}
	response.Header = make(map[string][]string)
	response.Header.Set("Content-Encoding", "gzip")

	data := struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}{}
	err = launchpad.DecodeResponse(&response, &data)
	c.Assert(err, IsNil)
	c.Assert(data.Name, Equals, "Hello")
	c.Assert(data.Status, Equals, "good")
}
*/
