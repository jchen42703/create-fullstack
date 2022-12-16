package main

import (
	"github.com/jchen42703/create-fullstack/internal/createcommands"
	"github.com/jchen42703/create-fullstack/internal/getproviders"
	"github.com/jchen42703/create-fullstack/version"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available create-fullstack commands.
// This is necessary for autocomplete to work as of 0.0.1.
var Commands map[string]cli.CommandFactory

// PrimaryCommands is an ordered sequence of the top-level commands (not
// subcommands) that we emphasize at the top of our help output. This is
// ordered so that we can show them in the typical workflow order, rather
// than in alphabetical order. Anything not in this sequence or in the
// HiddenCommands set appears under "all other commands".
var PrimaryCommands []string

// HiddenCommands is a set of top-level commands (not subcommands) that are
// not advertised in the top-level help at all. This is typically because
// they are either just stubs that return an error message about something
// no longer being supported or backward-compatibility aliases for other
// commands.
//
// No commands in the PrimaryCommands sequence should also appear in the
// HiddenCommands set, because that would be rather silly.
var HiddenCommands map[string]struct{}

// Initializes the commands for the CLI.
// Inspired by: https://github.com/hashicorp/terraform/blob/main/commands.go
func InitCommands(Ui cli.Ui) map[string]cli.CommandFactory {
	baseCommand := createcommands.BaseCommand{
		Ui: Ui,
	}

	// The command list is included in the create-fullstack -help
	// output, which is in turn included in the docs at
	// website/docs/cli/commands/index.html.markdown; if you
	// add, remove or reclassify commands then consider updating
	// that to match.

	Commands = map[string]cli.CommandFactory{
		"version": func() (cli.Command, error) {
			return &createcommands.VersionCommand{
				BaseCommand:       baseCommand,
				Version:           version.Version,
				VersionPrerelease: version.Prerelease,
				Platform:          getproviders.CurrentPlatform,
				// TODO: add a proper version check function down the line
				CheckFunc: nil,
			}, nil
		},
	}

	PrimaryCommands = []string{}

	HiddenCommands = map[string]struct{}{}
	return Commands
}
