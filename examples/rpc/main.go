package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/jchen42703/create-fullstack/core/aug"
	cfsplugin "github.com/jchen42703/create-fullstack/core/plugin"
)

// Runs the plugins
func main() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	cfsplugin.AugmentorManager.InitPlugin("ExampleAugmentor", &cfsplugin.AugmentorPlugin{})
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: cfsplugin.AugmentPluginHandshake,
		Plugins:         cfsplugin.AugmentorManager.Plugins(),
		Cmd:             exec.Command("./plugin/aug"),
		Logger:          logger,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC,
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

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	augmentor := raw.(aug.TemplateAugmentor)
	fmt.Println(augmentor.Id())
	fmt.Println(augmentor.Augment())
}
