package launchpad_test

import (
	"bytes"
	. "gopkg.in/check.v1"

	"github.com/bergotorino/go-launchpad/launchpad"
	"github.com/bergotorino/go-oauth/oauth"
	"net/http"
	"net/url"
)

// Mock oauth.Client
type OauthClientMock struct {
}

func (c *OauthClientMock) SetAuthorizationHeader(header http.Header, credentials *oauth.Credentials, method string, u *url.URL, params url.Values) error {
	return nil
}

func (c *OauthClientMock) Get(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return nil, nil
}
func (c *OauthClientMock) Post(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return nil, nil
}

func (c *OauthClientMock) RequestTemporaryCredentials(client *http.Client, callbackURL string, additionalParams url.Values) (*oauth.Credentials, error) {

	retval := oauth.Credentials{Token: "TemporaryToken", Secret: "TemporarySecret"}
	return &retval, nil
}

func (c *OauthClientMock) RequestToken(client *http.Client, temporaryCredentials *oauth.Credentials, verifier string) (*oauth.Credentials, url.Values, error) {
	retval := oauth.Credentials{Token: "Token", Secret: "Secret"}
	return &retval, nil, nil
}

func (c *OauthClientMock) AuthorizationURL(temporaryCredentials *oauth.Credentials, additionalParams url.Values) string {
	return "http://authorization.url"
}

func (c *OauthClientMock) SetCustomHeader(key string, value string) {
}

// Below methods are not needed
func (c *OauthClientMock) AuthorizationHeader(credentials *oauth.Credentials, method string, u *url.URL, params url.Values) string {
	return ""
}

func (c *OauthClientMock) SignForm(credentials *oauth.Credentials, method, urlStr string, form url.Values) error {
	return nil
}

func (c *OauthClientMock) SignParam(credentials *oauth.Credentials, method, urlStr string, params url.Values) {
	return
}

func (c *OauthClientMock) Delete(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return nil, nil
}

func (c *OauthClientMock) Put(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return nil, nil
}

func (c *OauthClientMock) RequestTokenXAuth(client *http.Client, temporaryCredentials *oauth.Credentials, user, password string) (*oauth.Credentials, url.Values, error) {
	return nil, nil, nil
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

func (s *LaunchpadSuite) TestLoginWith(c *C) {
	var data = make([]byte, 512)
	sb := bytes.NewBuffer(data)
	ocm := OauthClientMock{}

	lp := launchpad.NewClient(&ocm, "Some client")
	err := lp.LoginWith(sb)
	c.Assert(err, IsNil)
	n := bytes.IndexByte(data, 0x0)
	c.Assert(string(data[:n]), DeepEquals, string(`{"Token":"Token","Secret":"Secret"}`))
}

func (s *LaunchpadSuite) TestNewClient(c *C) {
	lp := launchpad.NewClient(nil, "Some client")
	c.Assert(lp, Not(IsNil))
}
