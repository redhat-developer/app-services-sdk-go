## Authentication for Management SDKs

Management SDKs are used to created instances. 
For those instances we can use OAuth.

All Management SDK's are working with Red Hat Single Sign On (sso.redhat.com) server.
Authentication to services can be done using Keycloak.js library.

When authenticating please use following keycloak client configuration:
```
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

`http://localhost/sso-redhat-callback`, `http://127.0.0.1/sso-redhat-callback`

## Authentication for Instance SDK's

To authenticate to the Instance SDK's we need to create service account
Service account can be created using [RHOAS CLI](https://github.com/redhat-developer/app-services-cli/blob/main/docs/commands/rhoas_serviceaccount_create.adoc) or
by UI (https://cloud.redhat.com)

Once we have created service account we can use it to obtain token to the instances SDK

```bash 
curl --location --request POST 'https://identity.api.openshift.com/auth/realms/rhoas/protocol/openid-connect/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'client_id=YOURSERVICEACCOUNT' \
--data-urlencode 'client_secret=YOURSERVICEACCOUNT_SECRET' \
--data-urlencode 'scope=email' \
--data-urlencode 'grant_type=client_credentials'
```

 > NOTE: We are working on SDK for Authentication that will provide out of the box support for connections
 
