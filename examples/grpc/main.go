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
	"github.com/jchen42703/create-fullstack/internal/log"
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

func installGoPlugin(logger hclog.Logger) (*cfsplugin.PluginMeta, error) {
	// Install plugin prior to running
	installerLogger, err := log.CreateLogger("./build/create-fullstack.log")
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %s", err)
	}

	installer := cfsplugin.AugmentorPluginInstaller{
		ParentOutputPluginDir: "./build",
		Writer:                logger.StandardWriter(&hclog.StandardLoggerOptions{}),
		Logger:                installerLogger,
	}

	meta, err := installer.Install("./plugin-go", "./build")
	if err != nil {
		return nil, fmt.Errorf("failed to install go grpc plugin: %s", err)
	}

	logger.Debug("meta", meta)
	return meta, nil
}

// Runs the plugins
func main() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// Go GRPC
	goMeta, err := installGoPlugin(logger)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = runExampleAugmentor(goMeta.Id, logger, exec.Command(filepath.Join("build", goMeta.Id)))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to run go grpc plugin: %s", err))
	}

	// TODO: Should read from python meta.
	pythonMetaId := "jchen42703-ExampleAugmentor-python"
	err = runExampleAugmentor(pythonMetaId, logger, exec.Command("python", "./plugin-python/plugin.py"))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to run python grpc plugin: %s", err))
	}
}
