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
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := hclog.New(&hclog.LoggerOptions{
				Name:   "plugin",
				Output: os.Stdout,
				Level:  hclog.Debug,
			})

			installer, err := cfsplugin.NewAugmentorPluginInstaller(cmdCtx.GlobalPluginsDir, logger)
			if err != nil {
				return fmt.Errorf("failed to initialize plugin installer: %s", err)
			}

			cs := cmdCtx.IoStreams.ColorScheme()
			allPlugins, err := installer.GetAllPlugins()
			if err != nil {
				return fmt.Errorf("failed to view all installed plugins: %s", err)
			}

			if len(allPlugins) == 0 {
				fmt.Fprint(os.Stdout, cs.Yellowf("No installed plugins found in the plugin directory '%s'.\nPlease run `create-fullstack plugins install <YOUR_PLUGIN_URL>` to install plugins.\n", cmdCtx.GlobalPluginsDir))
				return nil
			}

			fmt.Fprint(os.Stdout, cs.Bold("All Installed Plugins:\n"))
			for _, plugin := range allPlugins {
				fmt.Fprint(os.Stdout, cs.Grayf("%s [%s]\n", plugin.Id, plugin.Version))
			}

			return nil
		},
	}

	return listCmd
}
