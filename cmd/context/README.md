# `CmdContext`

Based off:

- https://github.com/cli/cli/blob/trunk/pkg/cmd/factory/default.go
- https://github.com/cli/cli/blob/791c7db632fc49e6b98f6b6c1df3d2b4d1a33675/pkg/cmdutil/factory.go#L19

```go
type Factory struct {
	Prompter  prompter.Prompter
    ...
	Remotes    func() (context.Remotes, error)
	Config     func() (config.Config, error)
	Branch     func() (string, error)
	...
	ExtensionManager extensions.ExtensionManager
	ExecutableName   string
}

// Executable is the path to the currently invoked binary
func (f *Factory) Executable() string {
	if !strings.ContainsRune(f.ExecutableName, os.PathSeparator) {
		f.ExecutableName = executable(f.ExecutableName)
	}
	return f.ExecutableName
}
```
