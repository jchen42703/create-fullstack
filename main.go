package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/jchen42703/create-fullstack/cmd/context"
	"github.com/jchen42703/create-fullstack/cmd/root"
	"github.com/jchen42703/create-fullstack/internal/executable"
	"github.com/jchen42703/create-fullstack/internal/log"
	"github.com/spf13/cobra"
)

type exitCode int

const (
	exitOK     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
	exitAuth   exitCode = 4
)

func main() {
	code := runMain()
	os.Exit(int(code))
}

// Based on Github's gh CLI (which also uses cobra)
// https://github.com/cli/cli/blob/trunk/cmd/gh/main.go
func runMain() exitCode {
	// TODO: Check for CLI updates
	ioStreams := iostreams.System()
	cs := ioStreams.ColorScheme()

	// TODO: dynamically get the log file path for different OSes
	logFilePath := "./create-fullstack.log"
	// Initialize logger. Uses CFS_LOG_LVL env var to determine the log level.
	logger, err := log.CreateLogger(logFilePath)
	if err != nil {
		fmt.Fprint(os.Stderr, cs.Redf("Error initializing logger: %s", err.Error()))
		return exitError
	}

	// TODO: set the pager command for making viewing logs cleaner.
	currentTime := time.Now()
	cmdCtx := &context.CmdContext{
		Version:          "0.0.0-dev",
		BuildDate:        currentTime.Format("2006-01-02"),
		IoStreams:        ioStreams,
		ExecutableName:   executable.GetPath("create-fullstack"),
		Logger:           logger,
		GlobalPluginsDir: context.GetGlobalPluginsDir(runtime.GOOS),
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			// this sync error is safe to ignore, since stdout doesn't support syncing in Linux/OS X
			if !strings.HasSuffix(err.Error(), "sync /dev/stdout: invalid argument") {
				fmt.Fprint(os.Stderr, cs.Redf("Error cleaning up logger: %s", err.Error()))
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

	// 3. Overrides survey's default color
	// 4. Build rootCmd
	// 5. provide completions for aliases and extensions (includes plugin commands)
	// 6. Executes the rootCmd
	// 7. Checks if it errors out. Handles the error to provide a better UX.

	var SilentErr = errors.New("SilentErr")
	rootCmd := root.NewCmdRoot(cmdCtx)
	// Only prints usage for flag errors
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		fmt.Fprintln(os.Stderr, cs.Red(err.Error()))
		cmd.Println("\n" + cmd.UsageString())
		return SilentErr
	})

	err = rootCmd.Execute()
	if err != nil {
		if err != SilentErr {
			fmt.Fprint(os.Stderr, cs.Redf("Failed to execute command: %s\n", err.Error()))
		}

		return exitError
	}

	return exitOK
}
