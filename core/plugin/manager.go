package plugin

import (
	"fmt"

	"github.com/hashicorp/go-plugin"
)

// Manages the state of installed plugins and loads plugins.
// Store generic plugin struct because we might have different plugin types in the future.
// We only want each plugin manager to handle one type of interface.
type PluginManager[T any] struct {
	plugins map[string]*CfsPlugin[T]
}

// Called by host to populate the host's plugin map with placeholder plugins that will be completed
// by the plugins.
func (m *PluginManager[T]) InitPlugin(id string, rawPlugin plugin.Plugin) error {
	if m.plugins == nil {
		return fmt.Errorf("must initialize plugins attribute first")
	}

	m.plugins[id] = &CfsPlugin[T]{
		Plugin: rawPlugin,
	}

	return nil
}

// Adds a full plugin implementation to the manager. This should be called by the plugins to add a real
// implementation of a plugin instead of having a placeholder.
func (m *PluginManager[T]) AddPlugin(id string, pluginI *CfsPlugin[T]) error {
	if m.plugins == nil {
		return fmt.Errorf("must initialize plugins attribute first")
	}

	m.plugins[id] = pluginI
	return nil
}

// Gets all raw plugins. This is useful for specifying the plugins in plugin.ClientConfig and
// plugin.ServeConfig.
func (m *PluginManager[T]) Plugins() plugin.PluginSet {
	rawPlugins := make(plugin.PluginSet)
	for id, cfsPlugin := range m.plugins {
		rawPlugins[id] = cfsPlugin.Plugin
	}

	return rawPlugins
}

// Gets all unused plugins. This is needed for detecting new augmentations (a.k.a. plugins are that not used yet.)
func (m *PluginManager[T]) GetUnusedPlugins() []*CfsPlugin[T] {
	unusedPlugins := []*CfsPlugin[T]{}
	for _, plugin := range m.plugins {
		if plugin.Used {
			unusedPlugins = append(unusedPlugins, plugin)
		}
	}

	return unusedPlugins
}

// Loads all plugins in directory using the state config. Adds those plugins to the manager state.
func (m *PluginManager[T]) LoadAllPlugins(pluginDir string) ([]*CfsPlugin[T], error) {
	return nil, nil
}

// Writes the current installed plugin state
// Need to be able to detect when state is corrupted (i.e. could not find state for installed executable)
// Plugin State:
// - version
// - install link
// - executable name
func (m *PluginManager[T]) WriteMetaToState(meta *PluginMeta, stateCfgPath string) error {
	// If state config doesn't exist already, create it.
	// Otherwise, unmarshal it.
	return nil
}
