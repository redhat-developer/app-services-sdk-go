package metadata

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
