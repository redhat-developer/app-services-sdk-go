# Authentication for Management SDKs

All Management SDK's are working with Red Hat Single Sign On (sso.redhat.com) server.
Authentication to services can be done using Keycloak.js library.

When authenticating please use following keycloak client configuration:

```json
{
  "realm": "redhat-external",
  "auth-server-url": "https://sso.redhat.com/auth/",
  "ssl-required": "all",
  "resource": "rhoas-cli-prod",
  "public-client": true,
  "confidential-port": 0
}
```

Valid redirect urls:

`http://localhost/sso-redhat-callback`
`http://127.0.0.1/sso-redhat-callback`
