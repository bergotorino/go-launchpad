package launchpad_test

import (
	"bytes"
	"errors"
	. "gopkg.in/check.v1"
	"net/url"

	"github.com/bergotorino/go-launchpad/launchpad"
	"github.com/bergotorino/go-launchpad/launchpad/testutil"
	"github.com/bergotorino/go-oauth/oauth"
)

// Handlers for OauthClientMock
func rtcGood() (*oauth.Credentials, error) {
	retval := oauth.Credentials{Token: "TemporaryToken", Secret: "TemporarySecret"}
	return &retval, nil
}

func rtcBad() (*oauth.Credentials, error) {
	return nil, errors.New("something really bad")
}

func rtGood() (*oauth.Credentials, url.Values, error) {
	retval := oauth.Credentials{Token: "Token", Secret: "Secret"}
	return &retval, nil, nil
}

func rtGood2() (*oauth.Credentials, url.Values, error) {
	retval := oauth.Credentials{Token: "Token2", Secret: "Secret2"}
	return &retval, nil, nil
}

func rtBad() (*oauth.Credentials, url.Values, error) {
	return nil, nil, errors.New("something really bad")
}

// Here it begins
type LaunchpadSuite struct {
	lp launchpad.Launchpad
}

var _ = Suite(&LaunchpadSuite{})

func (s *LaunchpadSuite) TestSnaps(c *C) {
	snaps := s.lp.Snaps()
	c.Assert(snaps, Not(IsNil))
}

func (s *LaunchpadSuite) TestGitRepositories(c *C) {
	gr := s.lp.GitRepositories()
	c.Assert(gr, Not(IsNil))
}

func (s *LaunchpadSuite) TestLoginWithSuccess(c *C) {
	var data = make([]byte, 512)
	sb := bytes.NewBuffer(data)
	ocm := testutil.OauthClientMock{}
	ocm.Fn_RequestTemporaryCredentials = rtcGood
	ocm.Fn_RequestToken = rtGood

	lp := launchpad.NewClient(&ocm, "Some client")
	err := lp.LoginWith(sb)
	c.Assert(err, IsNil)
	n := bytes.IndexByte(data, 0x0)
	c.Assert(string(data[:n]), DeepEquals, string(`{"Token":"Token","Secret":"Secret"}`))
}

// Verify you can login to launchpad again after a successful login
// attempt and the credentials are not overwritten
func (s *LaunchpadSuite) TestLoginWithSuccessTwice(c *C) {
	var data = make([]byte, 512)
	sb := bytes.NewBuffer(data)
	ocm := testutil.OauthClientMock{}
	ocm.Fn_RequestTemporaryCredentials = rtcGood
	ocm.Fn_RequestToken = rtGood

	lp := launchpad.NewClient(&ocm, "Some client")
	err := lp.LoginWith(sb)
	c.Assert(err, IsNil)
	n := bytes.IndexByte(data, 0x0)
	c.Assert(string(data[:n]), DeepEquals, string(`{"Token":"Token","Secret":"Secret"}`))

	ocm.Fn_RequestToken = rtGood2
	err = lp.LoginWith(sb)
	c.Assert(err, IsNil)
	n = bytes.IndexByte(data, 0x0)
	c.Assert(string(data[:n]), DeepEquals, string(`{"Token":"Token","Secret":"Secret"}`))
}

func (s *LaunchpadSuite) TestLoginWithFailureOne(c *C) {
	var data = make([]byte, 512)
	sb := bytes.NewBuffer(data)
	ocm := testutil.OauthClientMock{}
	ocm.Fn_RequestTemporaryCredentials = rtcBad
	ocm.Fn_RequestToken = rtGood

	lp := launchpad.NewClient(&ocm, "Some client")
	err := lp.LoginWith(sb)
	c.Assert(err, Not(IsNil))
}
func (s *LaunchpadSuite) TestLoginWithFailureTwo(c *C) {
	var data = make([]byte, 512)
	sb := bytes.NewBuffer(data)
	ocm := testutil.OauthClientMock{}
	ocm.Fn_RequestTemporaryCredentials = rtcGood
	ocm.Fn_RequestToken = rtBad

	lp := launchpad.NewClient(&ocm, "Some client")
	err := lp.LoginWith(sb)
	c.Assert(err, Not(IsNil))
}

func (s *LaunchpadSuite) TestNewClient(c *C) {
	lp := launchpad.NewClient(nil, "Some client")
	c.Assert(lp, Not(IsNil))
}
