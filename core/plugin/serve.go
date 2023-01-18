package plugin

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/jchen42703/create-fullstack/core/aug"
)

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
	})

	return nil
}
