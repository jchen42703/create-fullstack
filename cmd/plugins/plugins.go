package plugins

import (
	"fmt"
	"os"

	"github.com/jchen42703/create-fullstack/cmd/context"
	"github.com/jchen42703/create-fullstack/internal/directory"
	"github.com/spf13/cobra"
)

// Creates the `plugins` command with the appropriate subcommands.
func NewCmd(cmdCtx *context.CmdContext) *cobra.Command {
	err := os.MkdirAll(cmdCtx.GlobalPluginsDir, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		// Idk what to do here
		panic(fmt.Errorf("failed to create global plugins dir %s: %s", cmdCtx.GlobalPluginsDir, err))
	}

	pluginsCmd := &cobra.Command{
		Use:     "plugins",
		GroupID: "core",
		Short:   "Manage your plugins",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Error: must also specify an operation like: list, install, or uninstall")
		},
	}

	pluginsCmd.AddGroup(&cobra.Group{
		ID:    "plugins",
		Title: "Plugins Commands",
	})

	pluginsCmd.AddCommand(NewListCmd(cmdCtx))

	return pluginsCmd
}
