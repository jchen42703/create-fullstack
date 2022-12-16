package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jchen42703/create-fullstack/internal/didyoumean"
	"github.com/mitchellh/cli"
)

func runMain() int {
	Ui := &cli.BasicUi{
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
		Reader:      os.Stdin,
	}
	binName := filepath.Base(os.Args[0])
	args := os.Args[1:]
	allCommands := InitCommands(Ui)
	cliRunner := &cli.CLI{
		Name:                  binName,
		Args:                  args,
		Commands:              allCommands,
		HelpFunc:              helpFunc,
		HelpWriter:            os.Stdout,
		Autocomplete:          true,
		AutocompleteInstall:   "install-autocomplete",
		AutocompleteUninstall: "uninstall-autocomplete",
	}

	// Before we continue we'll check whether the requested command is
	// actually known. If not, we might be able to suggest an alternative
	// if it seems like the user made a typo.
	// (This bypasses the built-in help handling in cli.CLI for the situation
	// where a command isn't found, because it's likely more helpful to
	// mention what specifically went wrong, rather than just printing out
	// a big block of usage information.)

	// Check if this is being run via shell auto-complete, which uses the
	// binary name as the first argument and won't be listed as a subcommand.
	autoComplete := os.Getenv("COMP_LINE") != ""

	if cmd := cliRunner.Subcommand(); cmd != "" && !autoComplete {
		// Due to the design of cli.CLI, this special error message only works
		// for typos of top-level commands. For a subcommand typo, like
		// "terraform state posh", cmd would be "state" here and thus would
		// be considered to exist, and it would print out its own usage message.
		if _, exists := Commands[cmd]; !exists {
			suggestions := make([]string, 0, len(Commands))
			for name := range Commands {
				suggestions = append(suggestions, name)
			}

			suggestion := didyoumean.NameSuggestion(cmd, suggestions)
			if suggestion != "" {
				suggestion = fmt.Sprintf(" Did you mean %q?", suggestion)
			}

			fmt.Fprintf(os.Stderr, "create-fullstack has no command named %q.%s\n\nTo see all of create-fullstack's top-level commands, run:\n  create-fullstack -help\n\n", cmd, suggestion)
			return 1
		}
	}

	exitCode, err := cliRunner.Run()
	if err != nil {
		Ui.Error(fmt.Sprintf("Error executing CLI: %s", err.Error()))
		return 1
	}

	// if we are exiting with a non-zero code, check if it was caused by any
	// plugins crashing
	// if exitCode != 0 {
	// 	for _, panicLog := range logging.PluginPanics() {
	// 		Ui.Error(panicLog)
	// 	}
	// }

	return exitCode
}

func main() {
	os.Exit(runMain())
}
