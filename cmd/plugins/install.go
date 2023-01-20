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
		Short:   "Install a plugin from URL(s) or file URI(s).",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("must specify at least one plugin url or file uri as an argument")
			}

			logger := hclog.New(&hclog.LoggerOptions{
				Name:   "plugin",
				Output: os.Stdout,
				Level:  hclog.Debug,
			})

			installer, err := cfsplugin.NewAugmentorPluginInstaller(cmdCtx.GlobalPluginsDir, logger)
			if err != nil {
				return fmt.Errorf("failed to initialize plugin installer: %s", err)
			}

			// Install all plugins from multiple args
			for _, pluginUri := range args {
				cmdCtx.CliUi.Warnf("Installing plugin from [%s]...\n", pluginUri)
				err = installer.GetPlugin(pluginUri)
				if err != nil {
					return fmt.Errorf("failed to install plugin %s: %s", pluginUri, err)
				}

				cmdCtx.CliUi.Successf("Successfully installed plugin from [%s]\n\n", pluginUri)
			}

			cmdCtx.CliUi.Successf("Successfully installed %d plugins\n", len(args))
			return nil
		},
	}

	return listCmd
}
