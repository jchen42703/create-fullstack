package parser_test

import (
	"testing"

	"github.com/jchen42703/create-fullstack/core/opts"
	"github.com/jchen42703/create-fullstack/core/ui"
	"github.com/jchen42703/create-fullstack/internal/parser"
)

func TestValidator(t *testing.T) {

	tests := map[string]parser.TemplateConfig{

		"ValidDirectoryPath": {
			OutputDirectoryPath: "C:\\Users\\19172\\create-fullstack\\internal\\parser",
			Base:                "asas",
			Language:            "typescript",
			AugmentOpts:         &ui.UiAugmentOptions{},
		},
		"NoAugmentOpts": {
			OutputDirectoryPath: "asdsad",
			Base:                "asas",
			Language:            "typescript",
			AugmentOpts:         &ui.UiAugmentOptions{},
		},
		"ValidConfigMissingGitOpts": {
			OutputDirectoryPath: "C:\\Users\\19172\\create-fullstack\\internal\\parser",
			Base:                "asas",
			Language:            "typescript",
			AugmentOpts: &ui.UiAugmentOptions{
				AddScss:             &opts.PackageOptions{Version: "test"},
				AddTailwind:         &opts.PackageOptions{Version: "test"},
				AddStyledComponents: &opts.PackageOptions{Version: "test"},
				HuskyOpts:           &ui.HuskyOptions{Format: true, Lint: true},
				AddDockerfile:       true,
				AddCi:               "jenkins",
			},
		},
	}

	for name, testEx := range tests {

		if err := testEx.Validate(); err != nil {
			t.Fatalf("Example %s is not a valid config because %s", name, err)
		}

	}

}
