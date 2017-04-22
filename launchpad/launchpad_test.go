package launchpad_test

import (
	. "gopkg.in/check.v1"

	"github.com/bergotorino/go-launchpad/launchpad"
	"github.com/bergotorino/go-oauth/oauth"
	"net/http"
	"net/url"
)

type TestOauthClient struct {
}

func (c *TestOauthClient) AuthorizationHeader(credentials *oauth.Credentials, method string, u *url.URL, params url.Values) string {
	return ""
}

func (c *TestOauthClient) SetAuthorizationHeader(header http.Header, credentials *oauth.Credentials, method string, u *url.URL, params url.Values) error {
	return nil
}

func (c *TestOauthClient) Get(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return nil, nil
}
func (c *TestOauthClient) Post(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return nil, nil
}
func (c *TestOauthClient) Delete(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return nil, nil
}

func (c *TestOauthClient) Put(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return nil, nil
}

func (c *TestOauthClient) RequestTemporaryCredentials(client *http.Client, callbackURL string, additionalParams url.Values) (*oauth.Credentials, error) {
	return nil, nil
}

func (c *TestOauthClient) RequestToken(client *http.Client, temporaryCredentials *oauth.Credentials, verifier string) (*oauth.Credentials, url.Values, error) {
	return nil, nil, nil
}

func (c *TestOauthClient) AuthorizationURL(temporaryCredentials *oauth.Credentials, additionalParams url.Values) string {
	return ""
}

func (c *TestOauthClient) SetCustomHeader(key string, value string) {
}

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
	lp, err := launchpad.LoginWith(
}
