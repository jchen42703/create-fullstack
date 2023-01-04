package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jchen42703/create-fullstack/internal/didyoumean"
	"github.com/jchen42703/create-fullstack/internal/log"
	"github.com/mitchellh/cli"
)

func runMain() int {
	Ui := &cli.BasicUi{
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
		Reader:      os.Stdin,
	}

	colorUi := &cli.ColoredUi{
		OutputColor: cli.UiColorNone,
		InfoColor:   cli.UiColorGreen,
		ErrorColor:  cli.UiColorRed,
		WarnColor:   cli.UiColorYellow,
		Ui:          Ui,
	}

	logger, err := log.CreateLogger("./create-fullstack.log")
	if err != nil {
		colorUi.Error(fmt.Sprintf("Error initializing logger: %s", err.Error()))
		return 1
	}

	// These error out for some reason
	defer func() {
		if err := logger.Sync(); err != nil {
			// this sync error is safe to ignore, since stdout doesn't support syncing in Linux/OS X
			if !strings.HasSuffix(err.Error(), "sync /dev/stdout: invalid argument") {
				colorUi.Error(fmt.Sprintf("Error cleaning up logger: %s", err.Error()))
			}
		}
	}()

	// // Used to return the file in CreateLogger so we could close it here, but idk why that also
	// // threw an error that it was already closed
	// defer func() {
	// 	if err := f.Close(); err != nil {
	// 		colorUi.Error(fmt.Sprintf("Error closing log file: %s", err.Error()))
	// 	}
	// }()

	binName := filepath.Base(os.Args[0])
	args := os.Args[1:]
	allCommands := InitCommands(colorUi)
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

			colorUi.Error(fmt.Sprintf("create-fullstack has no command named %q.%s\n\nTo see all of create-fullstack's top-level commands, run:\n  create-fullstack -help\n\n", cmd, suggestion))
			return 1
		}
	}

	exitCode, err := cliRunner.Run()
	if err != nil {
		colorUi.Error(fmt.Sprintf("Error executing CLI: %s", err.Error()))
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
