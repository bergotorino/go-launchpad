package testutil

import (
	"github.com/bergotorino/go-oauth/oauth"
	"net/http"
	"net/url"
)

// Mock oauth.Client
type OauthClientMock struct {
	Fn_RequestTemporaryCredentials func() (*oauth.Credentials, error)
	Fn_RequestToken                func() (*oauth.Credentials, url.Values, error)
	Fn_Get                         func(url string, form url.Values) (*http.Response, error)
}

func (c *OauthClientMock) SetAuthorizationHeader(header http.Header, credentials *oauth.Credentials, method string, u *url.URL, params url.Values) error {
	return nil
}

func (c *OauthClientMock) Get(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {

	if c.Fn_Get != nil {
		return c.Fn_Get(urlStr, form)
	}

	return nil, nil
}
func (c *OauthClientMock) Post(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return nil, nil
}

func (c *OauthClientMock) RequestTemporaryCredentials(client *http.Client, callbackURL string, additionalParams url.Values) (*oauth.Credentials, error) {

	if c.Fn_RequestTemporaryCredentials != nil {
		return c.Fn_RequestTemporaryCredentials()
	}

	return nil, nil
}

func (c *OauthClientMock) RequestToken(client *http.Client, temporaryCredentials *oauth.Credentials, verifier string) (*oauth.Credentials, url.Values, error) {

	if c.Fn_RequestToken != nil {
		return c.Fn_RequestToken()
	}

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
