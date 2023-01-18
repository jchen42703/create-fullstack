# Basic RPC Plugin Example

## Getting Started

```bash
go build main.go && ./main
```

## Key Points

The `Plugins` map sent to:

**`main.go`**

```go
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: cfsplugin.AugmentPluginHandshake,
		Plugins:         cfsplugin.AugmentorPluginManager.RawPlugins(),
		Cmd:             exec.Command("./plugin/aug"),
		Logger:          logger,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC,
		},
	})
```

and

**`plugin/aug.go`**

```go
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: cfsplugin.AugmentPluginHandshake,
		Plugins:         cfsplugin.AugmentorPluginManager.RawPlugins(),
		Logger:          logger,
	})
```

**DO NOT** need to be the same reference to a map.

Instead, they must contain the same keys. The only difference is that the host process' `cfsplugin.AugmentorPluginManager.RawPlugins()` contains an empty implementation of the plugin, while the plugin provides a plugin with the implemented struct.
