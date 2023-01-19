package plugins

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Creates the `plugins` command with the appropriate subcommands.
func NewCmd() *cobra.Command {
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

	pluginsCmd.AddCommand(NewListCmd())

	return pluginsCmd
}
