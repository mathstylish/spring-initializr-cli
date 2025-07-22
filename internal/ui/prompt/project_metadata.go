package prompt

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

type ProjectInfo struct {
	GroupID     string
	ArtifactID  string
	Name        string
	Description string
	PackageName string
}

func AskProjectMetadata(defaults ProjectInfo) (ProjectInfo, error) {
	entered := defaults

	fields := []struct {
		label    string
		defaultV func() string
		validate func(val string) error
		assign   func(val string)
	}{
		{
			label:    "👥 Group",
			defaultV: func() string { return entered.GroupID },
			validate: func(val string) error {
				if match, _ := regexp.MatchString(`^[a-zA-Z.]+$`, val); !match {
					return fmt.Errorf("only letters and dots are allowed (e.g., org.example)")
				}
				return nil
			},
			assign: func(val string) {
				entered.GroupID = val
				entered.PackageName = val + "." + entered.ArtifactID
			},
		},
		{
			label:    "🏺 Artifact",
			defaultV: func() string { return entered.ArtifactID },
			validate: func(val string) error {
				if match, _ := regexp.MatchString(`^[a-zA-Z-]+$`, val); !match {
					return fmt.Errorf("only letters and hyphens are allowed (e.g., my-app, myapp)")
				}
				return nil
			},
			assign: func(val string) {
				entered.ArtifactID = val
				entered.Name = val
				entered.PackageName = fmt.Sprintf("%s.%s", entered.GroupID, val)

				fmt.Print("\033[1A")
				fmt.Print("\033[2K")
			},
		},
		{
			label:    "✍️ Name",
			defaultV: func() string { return entered.Name },
			validate: nil, // no specific validation, just trim spaces
			assign: func(val string) {
				entered.Name = strings.TrimRight(val, " ")

				fmt.Print("\033[1A")
				fmt.Print("\033[2K")
			},
		},
		{
			label:    "🗒️ Description",
			defaultV: func() string { return entered.Description },
			validate: nil,
			assign: func(val string) {
				entered.Description = strings.TrimRight(val, " ")

				fmt.Print("\033[1A")
				fmt.Print("\033[2K")
			},
		},
		{
			label:    "📦 Package name",
			defaultV: func() string { return entered.PackageName },
			validate: func(val string) error {
				if match, _ := regexp.MatchString(`^[a-zA-Z.]+$`, val); !match {
					return fmt.Errorf("only letters and dots are allowed (e.g., com.example.myapp)")
				}
				return nil
			},
			assign: func(val string) {
				entered.PackageName = val

				fmt.Print("\033[1A")
				fmt.Print("\033[2K")
			},
		},
	}

	for _, field := range fields {
		prompt := promptui.Prompt{
			Label:    field.label,
			Default:  field.defaultV(),
			Validate: field.validate,
		}

		fmt.Print("\033[?25l")
		result, err := prompt.Run()
		if err != nil {
			return ProjectInfo{}, fmt.Errorf("error reading %s: %w", field.label, err)
		}
		defer fmt.Print("\033[?25h")

		field.assign(result)
	}

	return entered, nil
}
