package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/google/go-github/v38/github"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/metadata"
	"golang.org/x/oauth2"
)

var clientID string
var accessToken string
var owner string
var author string
var repo string
var repoMetadataPath string

func init() {
	flag.StringVar(&repoMetadataPath, "repo-metadata", "", "path to repo metadata JSON")
	flag.StringVar(&clientID, "client-id", "", "API client ID")
	flag.StringVar(&accessToken, "token", "", "Access token")
	flag.StringVar(&author, "author", "", "Pull request author")
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

	// TODO: Use https://github.com/go-git/go-git
	err := exec.Command("git", "config", "user.email", author+"@users.noreply.github.com").Run()
	if err != nil {
		log.Fatalln(err)
	}
	err = exec.Command("git", "config", "user.name", author).Run()
	if err != nil {
		log.Fatalln(err)
	}

	headBranch := fmt.Sprintf("generate-client/%v", clientID)
	cmd := exec.Command("git", "switch", "-c", headBranch)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = exec.Command("git", "push", "-u", "origin", headBranch).Run()
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
	if err != nil {
		log.Fatalln(err)
	}

	fullPackageName := metadata.GetPackageName(sdkEntry)
	goPkgURL := fmt.Sprintf("https://pkg.go.dev/%v", fullPackageName)

	// TODO: Externalize pull request options via flags
	pullReq := github.NewPullRequest{
		Title:               github.String(commitTitle),
		Base:                github.String("main"),
		MaintainerCanModify: github.Bool(true),
		Head:                github.String(fmt.Sprintf("generate-client/%v", clientID)),
		Body:                github.String(fmt.Sprintf("🤖 _This is an auto-generated pull request_. \n\nGenerated a new client for [%v](%v) package.", fullPackageName, goPkgURL)),
	}

	_, resp, err := client.PullRequests.Create(ctx, owner, repo, &pullReq)

	fmt.Println(resp.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}
}
