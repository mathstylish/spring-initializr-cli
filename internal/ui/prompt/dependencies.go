package prompt

import (
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mathstylish/initzr/metadata/model"
)

const finishLabel = "✅ Finish selection"

func SelectDependencies(options []model.Option) ([]string, error) {
	var selectedIDs []string
	availableOptions := make([]model.Option, len(options))
	copy(availableOptions, options)

	for {
		printSelectedDependencies(selectedIDs, options)

		choices, labelToID := buildChoices(availableOptions, selectedIDs)
		if len(choices) == 1 {
			break
		}

		selectedLabel, err := askUserToSelectDependency(choices)
		if err != nil {
			return nil, err
		}

		if selectedLabel == finishLabel {
			break
		}

		selectedIDs = append(selectedIDs, labelToID[selectedLabel])
	}

	return selectedIDs, nil
}

func FlattenDependencies(groups []model.DependencyGroup) []model.Option {
	var flat []model.Option
	for _, group := range groups {
		flat = append(flat, group.Values...)
	}
	return flat
}

func printSelectedDependencies(selectedIDs []string, allOptions []model.Option) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("📦 Selected dependencies:")
	if len(selectedIDs) == 0 {
		fmt.Println("  (none)")
	} else {
		for _, id := range selectedIDs {
			if name := findDependencyNameByID(allOptions, id); name != "" {
				fmt.Printf("    %s\n", name)
			}
		}
	}
	fmt.Println()
}

func buildChoices(options []model.Option, selectedIDs []string) ([]string, map[string]string) {
	choices := []string{finishLabel}
	labelMap := map[string]string{finishLabel: ""}

	for _, opt := range options {
		if !slices.Contains(selectedIDs, opt.ID) {
			choices = append(choices, opt.Name)
			labelMap[opt.Name] = opt.ID
		}
	}
	return choices, labelMap
}

func askUserToSelectDependency(choices []string) (string, error) {
	var chosen string
	prompt := &survey.Select{
		Message:  "➕ Add another dependency (or select Finish to continue)",
		Options:  choices,
		PageSize: 10,
	}
	if err := survey.AskOne(prompt, &chosen); err != nil {
		return "", err
	}
	return chosen, nil
}

func findDependencyNameByID(options []model.Option, id string) string {
	for _, opt := range options {
		if opt.ID == id {
			return opt.Name
		}
	}
	return ""
}
