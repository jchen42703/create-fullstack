package cliui_test

import (
	"testing"

	"github.com/jchen42703/create-fullstack/cmd/cliui"
)

func TestColorUi(t *testing.T) {
	cliUi := cliui.NewColorUi()
	// Colors don't work in go test
	cliUi.Bold("Should be bold text")
	cliUi.Error("Should be red text")
	cliUi.Errorf("Should be red text: %s %s", "formatted text", "formatted text 2\n")
	cliUi.Warn("Should be yellow text")
	cliUi.Warnf("Should be yellow text: %s %s", "formatted text", "formatted text 2\n")
	cliUi.Gray("Should be gray text")
	cliUi.Log("Should be regular text")
}
