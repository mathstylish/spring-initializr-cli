package main

import (
	"fmt"
	"log"

	"github.com/mathstylish/initzr/internal/ui/prompt"
	"github.com/mathstylish/initzr/metadata/client"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	metadata, err := client.FetchMetadata()
	if err != nil {
		return fmt.Errorf("fetch metadata: %w", err)
	}

	fmt.Printf("🗂️  Welcome to initzr. Your Spring project will be ready shortly! 🚀\n\n")

	buildTool, err := prompt.SelectBuildTool(metadata.Type)
	if err != nil {
		return fmt.Errorf("build tool selection: %w", err)
	}

	language, err := prompt.SelectLanguage(metadata.Language)
	if err != nil {
		return fmt.Errorf("language selection: %w", err)
	}

	bootVersion, err := prompt.SelectBootVersion(metadata.BootVersion)
	if err != nil {
		return fmt.Errorf("boot version selection: %w", err)
	}

	projectMetadata, err := prompt.AskProjectMetadata(prompt.ProjectInfo{
		GroupID:     metadata.GroupID.Default,
		ArtifactID:  metadata.ArtifactID.Default,
		Name:        metadata.Name.Default,
		Description: metadata.Description.Default,
		PackageName: metadata.PackageName.Default,
	})
	if err != nil {
		return fmt.Errorf("project metadata input: %w", err)
	}

	packaging, err := prompt.SelectPackaging(metadata.Packaging)
	if err != nil {
		return fmt.Errorf("packaging selection: %w", err)
	}

	javaVersion, err := prompt.SelectJavaVersion(metadata.JavaVersion)
	if err != nil {
		return fmt.Errorf("javaVersion selection: %w", err)
	}

	allDeps := prompt.FlattenDependencies(metadata.Dependencies.Values)
	selectedDeps, err := prompt.SelectDependencies(allDeps)
	if err != nil {
		log.Fatalf("dependencies selection: %v", err)
	}

	err = client.DownloadProject(client.DownloadRequest{
		Type:         buildTool.ID,
		Language:     language.ID,
		BootVersion:  bootVersion.Name,
		GroupID:      projectMetadata.GroupID,
		ArtifactID:   projectMetadata.ArtifactID,
		Name:         projectMetadata.Name,
		Description:  projectMetadata.Description,
		PackageName:  projectMetadata.PackageName,
		Packaging:    packaging.ID,
		JavaVersion:  javaVersion.ID,
		Dependencies: selectedDeps,
	})
	if err != nil {
		log.Fatalf("downloading project: %v", err)
	}

	fmt.Printf("\n🚀 Your Spring Boot project has been successfully downloaded\n")
	return nil
}
