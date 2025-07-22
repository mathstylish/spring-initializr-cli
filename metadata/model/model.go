// Package model defines the structures that represent the metadata returned by the Spring Initializr API.
// This metadata includes information about available Spring Boot versions, build tools, packaging options,
// programming languages, Java versions, and dependencies that can be used to configure a new Spring project.
package model

type Option struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DependencyGroup struct {
	Name   string   `json:"name"`
	Values []Option `json:"values"`
}

type Dependencies struct {
	Type   string            `json:"type"`
	Values []DependencyGroup `json:"values"`
}

type TextValue struct {
	Type    string `json:"type"`
	Default string `json:"default"`
}

type SelectableValue struct {
	Default string   `json:"default"`
	Values  []Option `json:"values"`
}

type InitializrMetadata struct {
	Dependencies Dependencies    `json:"dependencies"`
	Type         SelectableValue `json:"type"`
	Packaging    SelectableValue `json:"packaging"`
	JavaVersion  SelectableValue `json:"javaVersion"`
	Language     SelectableValue `json:"language"`
	BootVersion  SelectableValue `json:"bootVersion"`
	GroupID      TextValue       `json:"groupId"`
	ArtifactID   TextValue       `json:"artifactId"`
	Version      TextValue       `json:"version"`
	Name         TextValue       `json:"name"`
	Description  TextValue       `json:"description"`
	PackageName  TextValue       `json:"packageName"`
}
