package plugins

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/jchen42703/create-fullstack/cmd/context"
	cfsplugin "github.com/jchen42703/create-fullstack/core/plugin"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// Creates the command to list all installed plugins.
func NewUninstallCmd(cmdCtx *context.CmdContext) *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "uninstall",
		GroupID: "plugins",
		Short:   "Uninstall a plugin from plugin id(s).",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("must specify at least one plugin id to uninstall")
			}

			idsToUninstall := lo.FindUniques(args)

			logger := hclog.New(&hclog.LoggerOptions{
				Name:   "plugin",
				Output: os.Stdout,
				Level:  hclog.Debug,
			})

			installer, err := cfsplugin.NewAugmentorPluginInstaller(cmdCtx.GlobalPluginsDir, logger)
			if err != nil {
				return fmt.Errorf("failed to initialize plugin installer: %s", err)
			}

			for _, pluginId := range idsToUninstall {
				err := installer.Uninstall(pluginId)
				if err != nil {
					return fmt.Errorf("failed to uninstall %s: %s", pluginId, err)
				}

				cmdCtx.CliUi.Successf("Successfully uninstalled plugin [%s]\n", pluginId)
			}

			cmdCtx.CliUi.Successf("\nSuccessfully uninstalled %d plugins\n", len(idsToUninstall))

			return nil
		},
	}

	return listCmd
}
