package plugin

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/jchen42703/create-fullstack/core/aug"
)

// Called by a Go plugin to serve an implementation of aug.TemplateAugmentor.
func ServeAugmentor(augmentor aug.TemplateAugmentor, hcLogger hclog.Logger) error {
	addedPlugin := &CfsPlugin[aug.TemplateAugmentor]{
		Plugin: &AugmentorPlugin{
			Impl: augmentor,
		},
	}

	err := AugmentorManager.AddPlugin(augmentor.Id(), addedPlugin)
	if err != nil {
		return err
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: AugmentPluginHandshake,
		Plugins:         AugmentorManager.Plugins(),
		Logger:          hcLogger,
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})

	return nil
}
