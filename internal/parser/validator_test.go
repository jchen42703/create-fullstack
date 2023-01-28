package parser_test

import (
	"testing"

	"github.com/jchen42703/create-fullstack/core/opts"
	"github.com/jchen42703/create-fullstack/core/ui"
	"github.com/jchen42703/create-fullstack/internal/parser"
)

func TestValidator(t *testing.T) {

	tests := map[string]parser.TemplateConfig{
		// validator tag file currently only checks if a file exists at the file path
		// there is an pr to update that to check if path is valid
		// for the two tests we just check that the file exists essentially checking that the file path
		// is valid
		"ValidConfig1": {
			OutputDirectoryPath: "C:\\Users\\19172\\create-fullstack\\internal\\parser\\parser.go",
			Base:                "asas",
			Language:            "typescript",
			AugmentOpts: &ui.UiAugmentOptions{
				AddScss:             &opts.PackageOptions{Version: "test"},
				AddTailwind:         &opts.PackageOptions{Version: "test"},
				AddStyledComponents: &opts.PackageOptions{Version: "test"},
				HuskyOpts:           &ui.HuskyOptions{Format: true, Lint: true},
				AddDockerfile:       true,
				AddCi:               "jenkins",
				GitOpts:             &ui.GitOptions{AddIssueTemplates: true, AddPRTemplates: true},
			},
		},
		"ValidConfig2": {
			OutputDirectoryPath: "C:\\Users\\19172\\create-fullstack\\internal\\parser\\validator.go",
			Base:                "asas",
			Language:            "typescript",
			AugmentOpts: &ui.UiAugmentOptions{
				AddScss:             &opts.PackageOptions{Version: "test2"},
				AddTailwind:         &opts.PackageOptions{Version: "test2"},
				AddStyledComponents: &opts.PackageOptions{Version: "test2"},
				HuskyOpts:           &ui.HuskyOptions{Format: true, Lint: true},
				AddDockerfile:       true,
				AddCi:               "jenkins",
				GitOpts:             &ui.GitOptions{AddIssueTemplates: false, AddPRTemplates: true},
			},
		},
	}

	for name, testEx := range tests {

		err := testEx.Validate()

		if err != nil {
			t.Fatalf("Example %s is not a valid config because %s", name, err)
		}

	}

}
