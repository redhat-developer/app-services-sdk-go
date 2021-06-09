package metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type SdkEntry struct {
	APIGroup     string `json:"apiGroup"`
	APIVersion   string `json:"apiVersion"`
	ReleaseLevel string `json:"release_level"`
}

type RepoMetadata map[string]SdkEntry

type ClientPayload struct {
	ID          string `json:"id,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
}

func GetPackageName(sdkEntry *SdkEntry) string {
	return fmt.Sprintf("github.com/redhat-developer/app-services-sdk-go/%v/api%v", sdkEntry.APIGroup, sdkEntry.APIVersion)
}

func GetSdkEntry(clientID string, pathToMetadata string) (*SdkEntry, error) {
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fullMetadataPath := path.Join(root, pathToMetadata)
	f, err := ioutil.ReadFile(fullMetadataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var repoMetadata RepoMetadata
	err = json.Unmarshal([]byte(f), &repoMetadata)
	if err != nil {
		log.Fatalln(err)
	}

	clientConfig, ok := repoMetadata[clientID]
	if !ok {
		return nil, fmt.Errorf("no SDK entry found for %v", clientID)
	}
	return &clientConfig, nil
}
