// Package prompt provides interactive terminal prompts for selecting project configuration options,
// such as build tools and other Spring Initializr metadata values.
package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mathstylish/initzr/internal/util"
	"github.com/mathstylish/initzr/metadata/model"
)

func SelectJavaVersion(javaVersion model.SelectableValue) (model.Option, error) {
	defaultIndex := util.FindIndexByOptionID(javaVersion.Values, javaVersion.Default)

	options := buildOptions(javaVersion.Values)

	prompt := &survey.Select{
		Message:  " Java",
		Options:  options,
		PageSize: 10,
		Default:  options[defaultIndex],
	}

	var selected string
	if err := survey.AskOne(prompt, &selected); err != nil {
		return model.Option{}, fmt.Errorf("prompt failed: %w", err)
	}

	return findOptionByName(javaVersion.Values, selected)
}

func buildOptions(values []model.Option) []string {
	options := make([]string, len(values))
	for i, v := range values {
		options[i] = v.Name
	}
	return options
}

func findOptionByName(options []model.Option, name string) (model.Option, error) {
	for _, opt := range options {
		if opt.Name == name {
			return opt, nil
		}
	}
	return model.Option{}, fmt.Errorf("option with name %q not found", name)
}
