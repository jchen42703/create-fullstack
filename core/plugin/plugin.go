package plugin

import (
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

// General struct for storing abstracting different types of plugins.
// The generic type is the interface the plugin implements.
type CfsPlugin[T any] struct {
	Plugin       plugin.Plugin
	Meta         *PluginMeta
	Used         bool
	ExecPath     string // executable path
	PluginClient *plugin.Client
}

// Gets the plugin ID.
func (p *CfsPlugin[T]) String() string {
	return p.Meta.Id
}

func (p *CfsPlugin[T]) Load() (T, error) {
	// Target interface to load into
	// This is declared up here so that we can return the zero value of the interface
	// on errors.
	var implementedI T

	// For now, just use non-versioned plugins, so just this current plugin
	plugins := plugin.PluginSet{}
	plugins[p.String()] = p.Plugin

	// We're a host. Start by launching the plugin process.
	// TODO: add negotiated versioning
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: AugmentPluginHandshake,
		Plugins:         plugins,
		Cmd:             exec.Command(p.ExecPath),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})

	p.PluginClient = client

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return implementedI, fmt.Errorf("failed to connect to plugin rpc: %s", err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(p.String())
	if err != nil {
		return implementedI, fmt.Errorf("failed to dispense plugin interface: %s", err)
	}

	implementedI = raw.(T)
	return implementedI, nil
}

// Cleans up the plugin process. If this is not called, the plugin subprocess will still be alive even after
// the main process exits.
func (p *CfsPlugin[T]) Cleanup() {
	p.PluginClient.Kill()
}
