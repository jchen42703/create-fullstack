package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/jchen42703/create-fullstack/core/aug"
	cfsplugin "github.com/jchen42703/create-fullstack/core/plugin"
)

func runExampleAugmentor(pluginId string, logger hclog.Logger, cmd *exec.Cmd) error {
	cfsplugin.AugmentorManager.InitPlugin(pluginId, &cfsplugin.AugmentorPlugin{})
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: cfsplugin.AugmentPluginHandshake,
		Plugins:         cfsplugin.AugmentorManager.Plugins(),
		Cmd:             cmd,
		Logger:          logger,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
	})

	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return fmt.Errorf("failed to create rpc client: %s", err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(pluginId)
	if err != nil {
		return fmt.Errorf("failed to dispense: %s", err)
	}

	logger.Debug("Dispensed", pluginId)

	// We should have a TemplateAugmentor now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	augmentor := raw.(aug.TemplateAugmentor)
	logger.Debug("Id", augmentor.Id())
	logger.Debug("Augment", augmentor.Augment())
	return nil
}

func runPlugins(parentPluginDir string, logger hclog.Logger) error {
	installer, err := cfsplugin.NewAugmentorPluginInstaller(parentPluginDir, logger)
	if err != nil {
		return fmt.Errorf("failed to create augmentor plugin installer: %s", err)
	}

	// Copying plugins to the parent plugin dir
	// TODO: example with github link.
	err = installer.GetPlugin("./plugin-go")
	if err != nil {
		return fmt.Errorf("failed to get plugin-go: %s", err)
	}

	err = installer.GetPlugin("./plugin-python")
	if err != nil {
		return fmt.Errorf("failed to get plugin-python: %s", err)
	}

	// Install all plugins.
	metas, err := installer.InstallAll()
	if err != nil {
		return fmt.Errorf("failed to install plugins: %s", err)
	}

	for _, meta := range metas {
		metaPluginDir := filepath.Join(parentPluginDir, meta.Id)
		runCmd, err := installer.GetRunCmd(filepath.Join(metaPluginDir, meta.Entrypoint), filepath.Join(metaPluginDir, meta.Id))
		if err != nil {
			return fmt.Errorf("failed to get run cmd: %s", err)
		}

		err = runExampleAugmentor(meta.Id, logger, runCmd)
		if err != nil {
			return fmt.Errorf("failed to run grpc plugin: %s", err)
		}
	}

	return nil
}

// Runs the plugins
func main() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	pluginInstallDir := "./plugins"
	defer func() {
		err := os.RemoveAll(pluginInstallDir)
		if err != nil {
			logger.Error("failed to cleanup plugins dir")
		}
	}()

	err := runPlugins(pluginInstallDir, logger)
	if err != nil {
		logger.Debug(fmt.Sprintf("failed to run plugin: %s", err))
	}
}
