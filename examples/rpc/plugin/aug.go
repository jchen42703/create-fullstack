package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/jchen42703/create-fullstack/core/aug"
	cfsplugin "github.com/jchen42703/create-fullstack/core/plugin"
)

type ExampleAugmentor struct {
	logger hclog.Logger
}

func (a *ExampleAugmentor) Id() string {
	return "ExampleAugmentor"
}

func (a *ExampleAugmentor) Augment() error {
	a.logger.Debug("message from ExampleAugmentor.Augment")
	return nil
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	augmentor := &ExampleAugmentor{
		logger: logger,
	}

	execPath := "./plugin/aug"
	addedPlugin := &cfsplugin.CfsPlugin[aug.TemplateAugmentor]{
		ExecPath: execPath,
		Plugin: &cfsplugin.AugmentorPlugin{
			Impl: augmentor,
		},
	}

	logger.Debug("message from plugin", augmentor.Id())
	cfsplugin.AugmentorPluginManager.AddPlugin(augmentor.Id(), addedPlugin)
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: cfsplugin.AugmentPluginHandshake,
		Plugins:         cfsplugin.AugmentorPluginManager.RawPlugins(),
		Logger:          logger,
	})
}
