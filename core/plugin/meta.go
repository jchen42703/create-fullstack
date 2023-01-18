package plugin

// Used to define plugin metadata. This struct will be written to a state config file whenever a plugin is installed.
type PluginMeta struct {
	Id          string `json:"pluginId"`    // GitHubUsername-NameOfPluginStruct; used to name the folder
	Version     string `json:"version"`     // could do a semi-ver, could be git commit hash, etc.
	InstallLink string `json:"installLink"` // i.e. https://github.com/xxxx/example-plugin-repo
}
