package openapi

import (
	"net/url"
	"strings"
)

func GetFileName(downloadURL string) (string, error) {
	fileURL, err := url.Parse(downloadURL)
	if err != nil {
		return "", err
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	return fileName, nil
}
