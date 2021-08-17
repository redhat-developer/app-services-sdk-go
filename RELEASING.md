# Releasing

## Determing which module to release

The Go client libraries have several modules. If a file needs to be releases, you must release the closest ancestor module.

To see all modules:

```bash
$ cat `find . -name go.mod` | grep module
module github.com/redhat-developer/app-services-sdk-go/internal/apigen
module github.com/redhat-developer/app-services-sdk-go/internal/examples
module github.com/redhat-developer/app-services-sdk-go/kafkainstance
module github.com/redhat-developer/app-services-sdk-go/kafkamgmt
module github.com/redhat-developer/app-services-sdk-go/registrymgmt
module github.com/redhat-developer/app-services-sdk-go/connectormgmt
module github.com/redhat-developer/app-services-sdk-go/registryinstance
module github.com/redhat-developer/app-services-sdk-go
```

The `github.com/redhat-developer/app-services-sdk-go` is the repository root module. Each other module is a submodule.

If you need to release a change in `registrymgmt/apiv1/api_client.go`, the closest ancestor module is `github.com/redhat-developer/app-services-sdk-go/registrymgmt`, so you would need to release a new version of the `github.com/redhat-developer/app-services-sdk-go/registrymgmt` submodule.

If you need to release a change in `auth/oauth2/endpoint.go`, the closest ancestor module is `github.com/redhat-developer/app-services-sdk-go`, so you would need to release a new version of the `github.com/redhat-developer/app-services-sdk-go` repository root module. Note: releasing `github.com/redhat-developer/app-services-sdk-go` has no impact on any of the submodules, and vice-versa. They are release entirely independently.

## How to release

### Manual Release (`github.com/redhat-developer/app-services-sdk-go`)

**Requirements**

1. Navigate to the repository root and switch to `main`.
1. Run `git tag -l | grep -v beta | grep -v alpha` to see all existing releases. The current latest tag `$CURRENT_VERSION` is the largest tag. It should look like `vX.Y.Z` (note: ignore all `LIB/vX.Y.Z tags` - these are the tags for a specific submodule, not the root module). We'll call the current version `$CURRENT_VERSION` and the next version `$NEXT_VERSION`.
2. On `main`, run `git log $CURRENT_VERSION...` to list all the changes since the last release. Note: You must manually visually parse out changes to submodules (the git log is going to show you things in submodules, which are not going to be part of your release).
3. Edit `CHANGELOG.md` to include the summary of changes from the previous step.
4. Commit the changes, push to your fork and create a PR titled: `chore: release $NEXT_VERSION`.
5. Wait for the PR to be reviewed and merged. Once it's merged, and without merging any other PRs in the meantime:
    a. Switch the main branch
    b. Pull the latest changes
    c. Tag the repo with the next version: `git tag $NEXT_VERSION`
    d. Push the tag to origin: `git push origin $NEXT_VERSION`
6. Update the [releases page](https://github.com/redhat-developer/app-services-sdk-go/releases) with the new release, copying the contents of `CHANGELOG.md`.

### Manual Releases (submodules)

These steps assume we're releasing `github.com/redhat-developer/app-services-sdk-go/kafkamgmt` - change this value to be whichever module you are releasing.

1. Navigate to the repository root and switch to `main`.
1. Run `git tag -l | grep kafkamgmt | grep -v beta | grep -v alpha` to see all existing releases. The current latest tag `$CURRENT_VERSION` is the largest tag. It should look like `kafkamgmt/vX.Y.Z` (note: ignore all `LIB/vX.Y.Z tags` - these are the tags for a specific submodule, not the root module). We'll call the current version `$CURRENT_VERSION` and the next version `$NEXT_VERSION`.
2. On `main`, run `git log $CURRENT_VERSION.. -- kafkamgmt/` to list all the changes since the last release.
3. Edit `kafkamgmt/CHANGELOG.md` to include the summary of changes from the previous step.
4. Commit the changes, push to your fork and create a PR titled: `chore: release $NEXT_VERSION`. Note: The version of a submodule should be prefixed with the module name, e.g: `kafkamgmt/v1.0.1`
5. Wait for the PR to be reviewed and merged. Once it's merged, and without merging any other PRs in the meantime:
    a. Switch the main branch
    b. Pull the latest changes
    c. Tag the repo with the next version: `git tag $NEXT_VERSION`
    d. Push the tag to origin: `git push origin $NEXT_VERSION`
6. Update the [releases page](https://github.com/redhat-developer/app-services-sdk-go/releases) with the new release, copying the contents of `kafkamgmt/CHANGELOG.md`.