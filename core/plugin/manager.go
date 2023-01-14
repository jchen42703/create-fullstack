package plugin

import "github.com/hashicorp/go-plugin"

// Manages the state of installed plugins and loads plugins.
// Store generic plugin struct because we might have different plugin types in the future.
// We only want each plugin manager to handle one type of interface.
type PluginManager[T any] struct {
	plugins map[string]*CfsPlugin[T]
}

// Gets all raw plugins. This is useful for specifying the plugins in plugin.ClientConfig.
func (m *PluginManager[T]) Plugins() plugin.PluginSet {
	plugins := plugin.PluginSet{}
	for id, plugin := range m.plugins {
		plugins[id] = plugin.Plugin
	}
	return plugins
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

// Loads a plugin from an executable.
// Adds it to the PluginManager plugins map.
func (m *PluginManager[T]) LoadPlugin(execPath string) (*CfsPlugin[T], error) {
	return nil, nil
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
