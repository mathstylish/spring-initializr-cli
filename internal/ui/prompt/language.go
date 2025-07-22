// Package prompt provides interactive terminal prompts for selecting project configuration options,
// such as build tools and other Spring Initializr metadata values.
package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	util "github.com/mathstylish/initzr/internal/util"
	model "github.com/mathstylish/initzr/metadata/model"
)

func SelectLanguage(languages model.SelectableValue) (model.Option, error) {
	defaultIndex := util.FindIndexByOptionID(languages.Values, languages.Default)

	options := make([]string, len(languages.Values))
	labelToOption := make(map[string]model.Option, len(languages.Values))
	for i, lang := range languages.Values {
		options[i] = lang.Name
		labelToOption[lang.Name] = lang
	}

	var selected string
	prompt := &survey.Select{
		Message:  "💻 Language",
		Options:  options,
		PageSize: 10,
		Default:  options[defaultIndex],
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		return model.Option{}, fmt.Errorf("prompt failed: %w", err)
	}

	return labelToOption[selected], nil
}
