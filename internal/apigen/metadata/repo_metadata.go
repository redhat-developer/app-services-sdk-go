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
	ModuleName   string `json:"package_name"`
	OpenApiFile  string `json:"openapi_file"`
	Disabled     bool   `json:"disabled,omitempty"`
}

func (s *SdkEntry) OutputPath() string {
	return fmt.Sprintf("%v/api%v/client", s.APIGroup, s.APIVersion)
}

func (s *SdkEntry) PackageName() string {
	return fmt.Sprintf("%vclient", s.APIGroup)
}

type RepoMetadata map[string]SdkEntry

type ClientPayload struct {
	ID          string `json:"id,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
}

func GetPackageName(sdkEntry *SdkEntry) string {
	return fmt.Sprintf("github.com/redhat-developer/app-services-sdk-go/%v/api%v", sdkEntry.APIGroup, sdkEntry.APIVersion)
}

func GetClientMetadata(pathToMetadata string) (*RepoMetadata, error) {
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
		return nil, err
	}

	return &repoMetadata, nil
}

func GetSdkEntry(clientID string, pathToMetadata string) (*SdkEntry, error) {
	repoMetadata, err := GetClientMetadata(pathToMetadata)
	if err != nil {
		return nil, err
	}
	repoMetadataV := *repoMetadata
	clientConfig, ok := repoMetadataV[clientID]

	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	clientConfig.OpenApiFile = path.Join(root, clientConfig.OpenApiFile)

	if !ok {
		return nil, fmt.Errorf("no SDK entry found for %v", clientID)
	}
	return &clientConfig, nil
}
