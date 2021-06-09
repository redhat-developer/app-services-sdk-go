package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/google/go-github/v35/github"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/metadata"
	"golang.org/x/oauth2"
)

var clientID string
var accessToken string
var owner string
var repo string
var repoMetadataPath string

func init() {
	flag.StringVar(&repoMetadataPath, "repo-metadata", "", "path to repo metadata JSON")
	flag.StringVar(&clientID, "client-id", "", "API client ID")
	flag.StringVar(&accessToken, "token", "", "Access token")
	flag.StringVar(&owner, "owner", "", "User who creates the pull request")
	flag.StringVar(&repo, "repo", "", "The repo in which to create the pull request")
}

func main() {
	flag.Parse()
	if repoMetadataPath == "" {
		log.Fatalln("Missing required flag: --repo-metadata")
	}
	if clientID == "" {
		log.Fatalln("Missing required flag: --client-id")
	}
	if accessToken == "" {
		log.Fatalln("Missing required flag: --token")
	}
	if owner == "" {
		log.Fatalln("Missing required flag: --owner")
	}
	if repo == "" {
		log.Fatalln("Missing required flag: --repo")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	err := exec.Command("git", "config", "user.email", owner+"@users.noreply.github.com").Run()
	if err != nil {
		log.Fatalln(err)
	}
	err = exec.Command("git", "config", "user.name", owner).Run()
	if err != nil {
		log.Fatalln(err)
	}

	commitTitle := fmt.Sprintf("chore(%v): re-generate API client", clientID)
	err = exec.Command("git", "add", ".").Run()
	if err != nil {
		log.Fatalln(err)
	}
	err = exec.Command("git", "commit", "-m", commitTitle).Run()
	if err != nil {
		log.Fatalln(err)
	}

	sdkEntry, err := metadata.GetSdkEntry(clientID, repoMetadataPath)

	pullReq := github.NewPullRequest{
		Title: &commitTitle,
		Base:  github.String("main"),
		Head:  github.String("generate-client/" + clientID),
		Body: github.String(fmt.Sprintf(`
			_Auto generated pull request_

			This pull request updates the %v API client
		`, metadata.GetPackageName(sdkEntry))),
	}

	_, _, err = client.PullRequests.Create(ctx, owner, repo, &pullReq)
	if err != nil {
		log.Fatalln(err)
	}
}
