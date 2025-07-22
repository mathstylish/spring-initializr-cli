// Package prompt provides interactive terminal prompts for selecting project configuration options,
// such as build tools and other Spring Initializr metadata values.
package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	util "github.com/mathstylish/initzr/internal/util"
	model "github.com/mathstylish/initzr/metadata/model"
)

func SelectBuildTool(tools model.SelectableValue) (model.Option, error) {
	allowedIDs := []string{"maven-project", "gradle-project-kotlin", "gradle-project"}
	filtered := filterBuildToolOptions(tools.Values, allowedIDs)

	defaultIndex := util.FindIndexByOptionID(filtered, "maven-project")

	options := make([]string, len(filtered))
	labelToOption := make(map[string]model.Option, len(filtered))
	for i, opt := range filtered {
		options[i] = opt.Name
		labelToOption[opt.Name] = opt
	}

	var selected string
	prompt := &survey.Select{
		Message:  "🛠️ Project",
		Options:  options,
		PageSize: 10,
		Default:  options[defaultIndex],
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		return model.Option{}, fmt.Errorf("prompt failed: %w", err)
	}

	return labelToOption[selected], nil
}

func filterBuildToolOptions(options []model.Option, allowedIDs []string) []model.Option {
	allowedSet := make(map[string]struct{}, len(allowedIDs))
	for _, id := range allowedIDs {
		allowedSet[id] = struct{}{}
	}

	var filtered []model.Option
	for _, opt := range options {
		if _, ok := allowedSet[opt.ID]; ok {
			filtered = append(filtered, opt)
		}
	}

	return filtered
}
