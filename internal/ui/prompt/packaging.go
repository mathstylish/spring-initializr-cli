// Package prompt provides interactive terminal prompts for selecting project configuration options,
// such as build tools and other Spring Initializr metadata values.
package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	util "github.com/mathstylish/initzr/internal/util"
	model "github.com/mathstylish/initzr/metadata/model"
)

func SelectPackaging(packaging model.SelectableValue) (model.Option, error) {
	defaultIndex := util.FindIndexByOptionID(packaging.Values, packaging.Default)

	options := make([]string, len(packaging.Values))
	labelToOption := make(map[string]model.Option, len(packaging.Values))
	for i, v := range packaging.Values {
		options[i] = v.Name
		labelToOption[v.Name] = v
	}

	var selected string
	prompt := &survey.Select{
		Message:  "📦 Packaging",
		Options:  options,
		PageSize: 10,
		Default:  options[defaultIndex],
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		return model.Option{}, fmt.Errorf("prompt failed: %w", err)
	}

	return labelToOption[selected], nil
}
