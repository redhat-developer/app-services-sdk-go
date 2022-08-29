package apiv1

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

/**
* Create client with default auth and client for offline token
 */
func BuildAuthenticatedHTTPClient(offlineToken string) *http.Client {
	return BuildAuthenticatedHTTPClientCustom(offlineToken, DefaultClientID, DefaultAuthURL)
}

func BuildAuthenticatedHTTPClientCustom(offlineToken string, clientID string, authURL string) *http.Client {
	ctx := context.Background()
	cfg := oauth2.Config{
		ClientID: clientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:   authURL,
			TokenURL:  fmt.Sprintf("%s/%s", authURL, "protocol/openid-connect/token"),
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	ts := cfg.TokenSource(ctx, &oauth2.Token{
		RefreshToken: offlineToken,
	})

	return oauth2.NewClient(ctx, ts)
}
