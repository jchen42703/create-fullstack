# Plugins

## Use Cases

1. User wants to change an existing augmentation with their custom implementation.
2. User wants to add a completely new template augmentation.
   1. I.e. a new base GitHub Workflow generator
3. User wants to check what plugins they have installed.
4. User wants to install a new plugin.
5. User wants to download a new plugin.
6. User wants to upgrade a plugin version.

## Ideas

**Concept:**

1. Load all augmentation plugins.
2. For each plugin, check if it is overriding an existing augmentation or is a completely new one.
   1. If overriding existing augmentation, the pipelines should replace

**Necessary Changes:**

1. To make it easy to override existing augmentors, we need to make each augmentor has a unique id.
   1. Add an `Id() string` method to the `TemplateAugmentor` interface
   2. Store all loaded plugins in a `map[string]plugin.Plugin` where the ids are keys.
2. To make it easy to integrate a new template augmentation, we need to:
   1. Detect when a new augmentation is found.
      1.
   2. Figure out when to run it.
3. Need to create a versioned plugin system.
4. Need to create a separate directory for storing the installed plugin states (JSON).
   1. Need to store the installed plugins somewhere and their metadata.
5. Need example plugin

   1. Give example with GoReleaser
   2. Give example with just yaml config file:

   ```yaml
   plugin_name: husky-js-only
   version: 1.3
   ```

   3. Otherwise, defaults to commit hash as version

6. Need to create the commands to install, uninstall, update, and view plugins.

**Other Ideas:**

1. Planning Phase

   1. Checks that you have correct tools installed (i.e. `yarn`, `npm`, `pre-commit`, etc.)
   2. Installs those tools for you if not already installed.
   3. Generates a plan:

   ```
   Generate from: https://github.com/digota/digota
   ecommerce microservice
   - No preprocessing applied.

   Adding Pre-Commit with the following settings:
      ...

   Adding <CUSTOM_PLUGIN_ID>@<VERSION> with the following settings:
      ...
   ```

## Architecture [RPC-Only Version]

These are tbe foundation structs:

```go
// This needs to add `Id`
type TemplateAugmentor interface {
   Id() string
	Augment() error
}
```

```go
// This is necessary for actually serving the interface implementation.
type AugmentorPlugin struct {
	// Impl is the interface
	Impl aug.TemplateAugmentor

   // Plugin Management Parameters
   used bool
   execPath string // executable path
}

func (p *AugmentorPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &AugmentorRpcServer{Impl: p.Impl}, nil
}

func (*AugmentorPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &AugmentorRpcClient{client: c}, nil
}

type PluginMeta struct {
   Hanshake ...
   Version string
   InstallLink string // i.e. https://github.com/xxxx/example-plugin-repo
}
```

The `AugmentorRpcServer` is responsible for serving the implementation from a binary, while `AugmentorRpcClient` is what the `go-plugin.Plugin` uses to fetch the implementation from the RPC plugin server.

A simple example of what these may look like is:

```go
// RPC Client to get Augmentor function results from server
type AugmentorRpcClient struct {
	client *rpc.Client
}

func (g *AugmentorRpcClient) Augment() error {
	err := g.client.Call("Plugin.Augment", new(interface{}), nil)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		return err
	}

	return nil
}

// Here is the RPC server that AugmentorRpcClient talks to, conforming to
// the requirements of net/rpc
type AugmentorRpcServer struct {
	// This is the real implementation
	Impl aug.TemplateAugmentor
}

func (s *AugmentorRpcServer) Augment() error {
	return s.Impl.Augment()
}
```

To manage these plugins, we'll make a `PluginManager`:

```go
type AugmentorPluginManager struct {
   // Store generic plugin interfaces because we want to store multiple plugin types
   // But, might be better to just enforce AugmentorPlugin for now.
   // Maybe change this something that keeps track of the plugin metadata
   // - plugin version (from yaml config, go releaser, or commit hash)
   // - install link
   // - plugin struct
   plugins map[string]struct {
      Plugin plugin.Plugin
      Cfg PluginMeta
   }
}

// Needed for detecting new augmentations
func (m *AugmentorPluginManager) getUnusedPlugins() ([]plugin.Plugin, error) {
   ...
}

// Loads a plugin from an executable.
// Adds it to the PluginManager plugins map.
func (m *AugmentorPluginManager) LoadPlugin(execPath string) (plugin.Plugin, error) {
   ...
}

func (m *AugmentorPluginManager) LoadAllPlugins(pluginDir string) ([]plugin.Plugin, error) {
}

// Discovers new plugins
func (m *AugmentorPluginManager) Discover(execPath string) ([]plugin.Plugin, error) {
}

// Writes the current installed plugin state
// Need to be able to detect when state is corrupted (i.e. could not find state for installed executable)
// Plugin State:
// - version
// - install link
// - executable name
func (m *AugmentorPluginManager) WriteMetaToState() error {
}
```

**Misc:**

- We also need to store all of the handshake configs to something exportable like `HandshakeConfigs`
- Make separate plugin struct
  - `Load`
  - `String`

## Resources

- https://github.com/hashicorp/otto/blob/v0.2.0/command/plugin_manager.go
- https://github.com/hleong25/hashicorp-goplugin-separate-binary-example/blob/4254a153a1fb412ab0ef80c08fb958f887a492c7/src/lahenry.com/mainapp/pluginmgr/pluginmgr.go#L17
- https://github.com/hashicorp/go-plugin/issues/11
- https://github.com/hashicorp/go-plugin/blob/main/examples/negotiated/main.go
