/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/jchen42703/create-fullstack/cmd/context"
	"github.com/spf13/cobra"
)

// Creates the generator command.
func NewCmd(cmdCtx *context.CmdContext) *cobra.Command {
	genCmd := &cobra.Command{
		Use:     "gen",
		GroupID: "core",
		Short:   "Generate project templates with either a UI or YAML config",
		Example: heredoc.Doc(`
		To generate your template interactively with a UI:

		$ create-fullstack gen

		Generate templates using a yaml config:

		$ create-fullstack gen -f "<PATH_TO_CONFIG>"`),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("gen called with context: %+v\n", cmdCtx)
		},
	}

	genCmd.Flags().String("configPath", "", "The path to the YAML config file [optional].")

	return genCmd
}
