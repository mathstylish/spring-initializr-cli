// Package client provides a function to fetch and parse Spring Initializr metadata
// from the official start.spring.io endpoint.
package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/mathstylish/initzr/internal/util"
	"github.com/mathstylish/initzr/metadata/model"
)

type DownloadRequest struct {
	Type         string
	Language     string
	BootVersion  string
	GroupID      string
	ArtifactID   string
	Name         string
	Description  string
	PackageName  string
	Packaging    string
	JavaVersion  string
	Dependencies []string
}

func FetchMetadata() (model.InitializrMetadata, error) {
	client := resty.New()

	response, err := client.R().Get("https://start.spring.io/metadata/client")
	if err != nil {
		return model.InitializrMetadata{}, err
	}

	var initializr model.InitializrMetadata

	if err := json.Unmarshal(response.Body(), &initializr); err != nil {
		return model.InitializrMetadata{}, err
	}

	return initializr, nil
}

func DownloadProject(request DownloadRequest) error {
	url := buildSpringURL(request)

	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download project zip: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status downloading zip: %s", response.Status)
	}

	err = util.SaveAndUnzip(request.ArtifactID, response.Body)
	if err != nil {
		return fmt.Errorf("save and unzip project: %w", err)
	}

	return nil
}

func buildSpringURL(metadata DownloadRequest) string {
	base := "https://start.spring.io/starter.zip"
	params := fmt.Sprintf(
		"type=%s&language=%s&bootVersion=%s&baseDir=%s&groupId=%s&artifactId=%s&name=%s&description=%s&packageName=%s&packaging=%s&javaVersion=%s&dependencies=%s",
		metadata.Type,
		metadata.Language,
		metadata.BootVersion,
		metadata.ArtifactID,
		metadata.GroupID,
		metadata.ArtifactID,
		metadata.Name,
		strings.ReplaceAll(metadata.Description, " ", "+"),
		metadata.PackageName,
		metadata.Packaging,
		metadata.JavaVersion,
		strings.Join(metadata.Dependencies, ","),
	)
	return base + "?" + params
}
