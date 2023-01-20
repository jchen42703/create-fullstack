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
func NewInstallCmd(cmdCtx *context.CmdContext) *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "install",
		GroupID: "plugins",
		Short:   "Install a plugin from a URL or file URI.",
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

			pluginUri := args[0]
			cmdCtx.CliUi.Warnf("Installing plugin from [%s]...\n", pluginUri)
			err = installer.GetPlugin(pluginUri)
			if err != nil {
				return fmt.Errorf("failed to install plugin %s: %s", pluginUri, err)
			}

			cmdCtx.CliUi.Successf("Successfully installed plugin from [%s]\n", pluginUri)

			return nil
		},
	}

	return listCmd
}
