package launchpad

import (
	"fmt"
	"github.com/bergotorino/go-oauth/oauth"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const ACCESS_TOKEN_POOL_TIME = 10 * time.Second

type UserNotifier interface {
	Notify(message string)
}

type CliNotifier struct{}

func (c *CliNotifier) Notify(message string) {
	fmt.Printf(message)
}

type Launchpad struct {
	appName        string
	consumerKey    string
	secrets        LaunchpadSecrets
	secretsBackend SecretsBackend
	oauthClient    oauth.AbstractClient
	Notifier       UserNotifier
}

func NewClient(client oauth.AbstractClient, name string) *Launchpad {
	lp := Launchpad{appName: name, Notifier: &CliNotifier{}}

	if client == nil {
		lp.oauthClient = DefaultOauthClient()
	} else {
		lp.oauthClient = client
	}

	return &lp
}

func (l *Launchpad) LoginWith(sb SecretsBackend) error {
	l.secretsBackend = sb

	// Load secrets. If there is an error we do not fail but instead
	// proceed with authentication.
	err := l.secrets.Read(l.secretsBackend)
	if err == nil {
		if l.secrets.IsValid() {
			return nil
		}
	}

	// At this point we know that we have not been logged in before
	// Let's proceed with auth process

	tempCred, err := l.oauthClient.RequestTemporaryCredentials(nil, "", nil)
	if err != nil {
		return err
	}

	v := url.Values{}
	v.Set("allow_permission", "DESKTOP_INTEGRATION")
	authURL := l.oauthClient.AuthorizationURL(tempCred, v)

	msg := fmt.Sprintf("Open this link:\n\n%s\n\n%s\n%s\n",
		authURL, "to authorize this program to access Launchpad on your behalf",
		"Waiting to hear from Launchpad about your decision")
	l.Notifier.Notify(msg)

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

	l.secrets.Write(l.secretsBackend)

	// Launchpad requires that certain Headers are set
	l.oauthClient.SetCustomHeader("accept", "application/json")
	l.oauthClient.SetCustomHeader("accept-encoding", "gzip, deflate")
	l.oauthClient.SetCustomHeader("Referer", "https://launchpad.net/")
	userAgent := fmt.Sprintf("application=\"%s\"; oauth_consumer=\"%s\"", l.appName, l.consumerKey)
	l.oauthClient.SetCustomHeader("user-agent", userAgent)

	return nil
}

func (l *Launchpad) Get(resource string, form url.Values) (*http.Response, error) {
	return l.oauthClient.Get(nil, l.secrets.accessCredentials, resource, form)
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
