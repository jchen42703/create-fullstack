// Based off:
// https://github.com/cli/cli/blob/791c7db632fc49e6b98f6b6c1df3d2b4d1a33675/pkg/cmd/root/root.go#L40
// This is the main entrypoint for the CLI commands.
package root

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/jchen42703/create-fullstack/cmd/augment"
	"github.com/jchen42703/create-fullstack/cmd/context"
	genCmd "github.com/jchen42703/create-fullstack/cmd/gen"
	"github.com/jchen42703/create-fullstack/cmd/plugins"
	versionCmd "github.com/jchen42703/create-fullstack/cmd/version"
	"github.com/spf13/cobra"
)

// Creates the main entrypoint for the Cobra command runner.
// - The root command manages the top level commands and command groups.
func NewCmdRoot(cmdCtx *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-fullstack <command> <subcommand> [flags]",
		Short: "Create Fullstack CLI",
		Long:  `Create new or augment existing project templates for consistent and quick fullstack application development.`,
		Example: heredoc.Doc(`
			$ create-fullstack gen
			$ create-fullstack gen -f "<PATH_TO_CONFIG>"
			$ create-fullstack augment ui tailwind
			$ create-fullstack plugins list`),
		Annotations: map[string]string{
			"versionInfo": versionCmd.Format(cmdCtx.Version, cmdCtx.BuildDate),
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")

	// // override Cobra's default behaviors unless an opt-out has been set
	// if os.Getenv("CFS_COBRA") == "" {
	// 	cmd.SilenceErrors = true
	// 	cmd.SilenceUsage = true

	// 	// this --version flag is checked in rootHelpFunc
	// 	cmd.Flags().Bool("version", false, "Show gh version")

	// 	cmd.SetHelpFunc(func(c *cobra.Command, args []string) {
	// 		rootHelpFunc(f, c, args)
	// 	})
	// 	cmd.SetUsageFunc(func(c *cobra.Command) error {
	// 		return rootUsageFunc(f.IOStreams.ErrOut, c)
	// 	})
	// 	cmd.SetFlagErrorFunc(rootFlagErrorFunc)
	// }

	// Misc Commands
	cmd.AddCommand(versionCmd.NewCmd(cmdCtx))

	// Core Commands
	cmd.AddGroup(&cobra.Group{
		ID:    "core",
		Title: "Core Commands",
	})

	cmd.AddCommand(augment.NewCmd())
	cmd.AddCommand(plugins.NewCmd(cmdCtx))
	cmd.AddCommand(genCmd.NewCmd(cmdCtx))

	return cmd
}
