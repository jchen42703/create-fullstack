package parser_test

import (
	"testing"

	"github.com/jchen42703/create-fullstack/internal/parser"
)

func TestValidator(t *testing.T) {

	tests := []parser.TemplateConfig {

		{OutputDirectoryPath: "asdsad", 
		Base: "asas", 
		Language: "typescript",
		AugmentOpts: &ui.UiAugmentOptions{}}
	}
}
