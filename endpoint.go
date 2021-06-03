package rhoas

import "golang.org/x/oauth2"

// Endpoint is the default OAuth endpoints for RHOAS
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/auth",
	TokenURL: "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token",
}
