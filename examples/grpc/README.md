# GRPC Plugins

Loading Go and Python plugins over GRPC.

## Getting Started

Building and running the Go GRPC plugin:

```
go build main.go && ./main
```

<!-- ```
$ export KV_PLUGIN="python plugin-python/plugin.py"
``` -->

## Generating Proto Files

For Go:

```bash
# In create-fullstack/core
protoc -I=proto --go_out=. --go-grpc_out=. proto/aug.proto
```

<!-- For Python:

```bash
python -m grpc_tools.protoc -I ./proto/ --python_out=./plugin-python/ --grpc_python_out=./plugin-python/ ./proto/kv.proto
``` -->

## Key Points

The `Plugins` map sent to:

**`main.go`**

```go
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: cfsplugin.AugmentPluginHandshake,
		Plugins:         cfsplugin.AugmentorManager.Plugins(),
		Cmd:             exec.Command(filepath.Join("build", meta.Id)),
		Logger:          logger,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
	})
```

and

**`plugin-go/main.go`**

```go
// In
// err := cfsplugin.ServeAugmentor(augmentor, logger)
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: AugmentPluginHandshake,
		Plugins:         AugmentorManager.Plugins(),
		Logger:          hcLogger,
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
```

**DO NOT** need to be the same reference to a map.

Instead, they must contain the same keys. The only difference is that the host process' `cfsplugin.AugmentorPluginManager.Plugins()` contains an empty implementation of the plugin, while the plugin provides a plugin with the implemented struct.
