# Auth module

Module provides useful helper functions for authentication in the RHOAS SDKs.

## Example

```go
 offlineToken := os.Getenv("OFFLINE_TOKEN")
 httpClient := rhoasAuth.BuildAuthenticatedHTTPClient(offlineToken)

 ctx := context.Background()

 client := kafkamgmt.NewAPIClient(&kafkamgmt.Config{
  HTTPClient: httpClient,
 })

 _, _, err := client.DefaultApi.GetKafkas(ctx).Execute()
 if err != nil {
  panic(err)
 }
```
