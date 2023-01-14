package plugin

import "github.com/hashicorp/go-plugin"

// Used to define plugin metadata. This struct will be written to a state config file whenever a plugin is installed.
type PluginMeta struct {
	Id          string                 `json:"pluginId"`
	Version     string                 `json:"version"`        // could do a semi-ver, could be git commit hash, etc.
	InstallLink string                 `json:"installLink"`    // i.e. https://github.com/xxxx/example-plugin-repo
	Executable  string                 `json:"executableName"` // Just the executable name, not the absolute path to it.
	Handshake   plugin.HandshakeConfig `json:"-"`              // should not write to JSON
}
