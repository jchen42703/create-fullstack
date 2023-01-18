package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/jchen42703/create-fullstack/core/aug"
	cfsplugin "github.com/jchen42703/create-fullstack/core/plugin"
)

func runExampleAug1(logger hclog.Logger) {
	// Install plugin prior to running
	installer := cfsplugin.AugmentorPluginInstaller{
		ParentOutputPluginDir: "./build",
		Logger:                logger.StandardWriter(&hclog.StandardLoggerOptions{}),
	}

	meta, err := installer.Install("./plugin-go", "./build")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Debug("meta", meta)

	cfsplugin.AugmentorManager.InitPlugin("ExampleAugmentor", &cfsplugin.AugmentorPlugin{})
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: cfsplugin.AugmentPluginHandshake,
		Plugins:         cfsplugin.AugmentorManager.Plugins(),
		Cmd:             exec.Command(filepath.Join("build", meta.Id)),
		Logger:          logger,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
	})

	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		logger.Error(err.Error())
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("ExampleAugmentor")
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Debug("Dispensed augmentor")

	// We should have a TemplateAugmentor now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	augmentor := raw.(aug.TemplateAugmentor)
	logger.Debug("Id", augmentor.Id())
	logger.Debug("Augment", augmentor.Augment())
}

func runExampleAug2(logger hclog.Logger) {
	cfsplugin.AugmentorManager.InitPlugin("ExampleAugmentor2", &cfsplugin.AugmentorPlugin{})
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: cfsplugin.AugmentPluginHandshake,
		Plugins:         cfsplugin.AugmentorManager.Plugins(),
		Cmd:             exec.Command("python", "./plugin-python/plugin.py"),
		Logger:          logger,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
	})

	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		logger.Error(err.Error())
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("ExampleAugmentor2")
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Debug("Dispensed augmentor 2")

	// We should have a TemplateAugmentor now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	augmentor := raw.(aug.TemplateAugmentor)
	logger.Debug("Id", augmentor.Id())
	logger.Debug("Augment", augmentor.Augment())
}

// Runs the plugins
func main() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})
	runExampleAug1(logger)
	runExampleAug2(logger)
}
