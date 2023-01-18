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
	// rawPlugins plugin.PluginSet
}

func (m *PluginManager[T]) AddPlugin(id string, pluginI *CfsPlugin[T]) error {
	if m.plugins == nil {
		return fmt.Errorf("must initialize plugins attribute first")
	}

	m.plugins[id] = pluginI
	return nil
}

// Called by host.
func (m *PluginManager[T]) InitializePlugin(id string, rawPlugin plugin.Plugin) error {
	if m.plugins == nil {
		return fmt.Errorf("must initialize plugins attribute first")
	}

	m.plugins[id] = &CfsPlugin[T]{
		Plugin: rawPlugin,
	}

	return nil
}

// Gets all raw plugins. This is useful for specifying the plugins in plugin.ClientConfig.
func (m *PluginManager[T]) RawPlugins() plugin.PluginSet {
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

// // Loads a plugin from an executable.
// // Adds it to the PluginManager plugins map.
// func (m *PluginManager[T]) LoadPlugin(execPath string) (*CfsPlugin[T], error) {
// 	loadedPlugin := &CfsPlugin[T]{
// 		ExecPath: execPath,
// 		Plugin:   &AugmentorPlugin{},
// 	}

// 	_, err := loadedPlugin.Load()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load plugin: %s", err)
// 	}

// 	return loadedPlugin, nil
// }

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
