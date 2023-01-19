package plugins

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/jchen42703/create-fullstack/cmd/context"
	cfsplugin "github.com/jchen42703/create-fullstack/core/plugin"
	"github.com/spf13/cobra"
)

// Creates the command to list all installed plugins.
func NewListCmd(cmdCtx *context.CmdContext) *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list",
		GroupID: "plugins",
		Short:   "List all installed plugins.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := hclog.New(&hclog.LoggerOptions{
				Name:   "plugin",
				Output: os.Stdout,
				Level:  hclog.Debug,
			})

			installer, err := cfsplugin.NewAugmentorPluginInstaller(cmdCtx.GlobalPluginsDir, logger)
			if err != nil {
				// idk what to do here lol
				fmt.Println("NewAugmentorPluginInstaller: ", err)
				return
			}

			fmt.Println("List all installed plugins")
			fmt.Println(installer.GetAllPlugins())
		},
	}

	return listCmd
}
