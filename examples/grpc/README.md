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
