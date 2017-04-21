package launchpad

import (
	"fmt"
	"github.com/bergotorino/go-oauth/oauth"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const ACCESS_TOKEN_POOL_TIME = 10 * time.Second

type Launchpad struct {
	appName     string
	consumerKey string
	secrets     LaunchpadSecrets
	oauthClient oauth.Client
}

func LoginWith(s LaunchpadSecrets, applicationName string) (*Launchpad, error) {
	launchpad := newLaunchpadClient(s, applicationName)
	err := launchpad.doLogin()
	if err != nil {
		return nil, err
	}
	return launchpad, nil
}

func newLaunchpadClient(s LaunchpadSecrets, appName string) *Launchpad {

	consumerKey := "System-wide: Ubuntu"

	hostname, err := os.Hostname()
	if err == nil {
		consumerKey += fmt.Sprintf(" (%s)", hostname)
	}

	lp := Launchpad{appName: appName, secrets: s,
		oauthClient: oauth.Client{
			TemporaryCredentialRequestURI: "https://launchpad.net/+request-token",
			ResourceOwnerAuthorizationURI: "https://launchpad.net/+authorize-token",
			TokenRequestURI:               "https://launchpad.net/+access-token",
			Credentials: oauth.Credentials{
				Token:  consumerKey,
				Secret: "",
			},
			SignatureMethod: oauth.PLAINTEXT,
		}}

	return &lp
}

func (l *Launchpad) Get(resource string, form url.Values) (*http.Response, error) {
	return l.oauthClient.Get(nil, l.secrets.accessCredentials, resource, form)
}

func (l *Launchpad) doLogin() error {

	// Secrets already loaded, i.e. we have been already logged in
	// on this machine.
	if l.secrets.IsValid() {
		return nil
	}

	// Not logged in before, proceed with auth process

	tempCred, err := l.oauthClient.RequestTemporaryCredentials(nil, "", nil)
	if err != nil {
		return err
	}

	v := url.Values{}
	v.Set("allow_permission", "DESKTOP_INTEGRATION")
	authURL := l.oauthClient.AuthorizationURL(tempCred, v)

	fmt.Printf("Open this link:\n\n%s\n\n", authURL)
	fmt.Printf("to authorize this program to access Launchpad on your behalf.\n")
	fmt.Printf("Waiting to hear from Launchpad about your decision. . . .\n")

	// Here comes the hack! At this point the user has been asked
	// to authorize the application by visiting a URL. There is,
	// however, no chance for an application like this to know
	// when and if it happened. Luckily the Launchpad answers with
	// http errors to indicate what happened:
	// - 401: the user has not made the decision yet
	// - 403: the user has decided not to grant access
	// - anything else: error accessing the server
	// - no error: we are all good.
	for {
		time.Sleep(ACCESS_TOKEN_POOL_TIME)
		creds, _, err := l.oauthClient.RequestToken(nil, tempCred, "")
		if err != nil {
			// The hack is that we can learn the error code by
			// searching for it in the error string. It is in a
			// form: "OAuth server status ERRORCODE"
			if strings.Contains(err.Error(), "401") {
				continue
			}
			return err
		} else {
			l.secrets.accessCredentials = creds
			break
		}
	}

	l.secrets.Store()

	// Launchpad requires that certain Headers are set
	l.oauthClient.SetCustomHeader("accept", "application/json")
	l.oauthClient.SetCustomHeader("accept-encoding", "gzip, deflate")
	l.oauthClient.SetCustomHeader("Referer", "https://launchpad.net/")
	userAgent := fmt.Sprintf("application=\"%s\"; oauth_consumer=\"%s\"", l.appName, l.consumerKey)
	l.oauthClient.SetCustomHeader("user-agent", userAgent)

	return nil
}

func (l *Launchpad) People(name string) (*Person, error) {
	response, err := l.Get("https://api.launchpad.net/devel/~"+name, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data Person
	err = DecodeResponse(response, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (l *Launchpad) Snaps() *Snaps {
	return &Snaps{lp: l}
}

func (l *Launchpad) GitRepositories() *GitRepositories {
	return &GitRepositories{lp: l}
}
