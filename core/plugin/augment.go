package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"github.com/jchen42703/create-fullstack/core/aug"
)

// Handshake is a common handshake that is shared by plugin and host.
var AugmentPluginHandshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "AUG_PLUGIN",
	MagicCookieValue: "2f4b473f573429bca56f545f02c6bac47cebcc3130fbdaf5e1a14dda66349d93", // echo -n create-fullstack-aug-plugin | sha256sum
}

// RPC Client to get Augmentor function results from server
type AugmentorRpcClient struct {
	client *rpc.Client
}

func (g *AugmentorRpcClient) Id() string {
	var resp string
	err := g.client.Call("Plugin.Id", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
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

func (s *AugmentorRpcServer) Id(args interface{}, resp *string) error {
	*resp = s.Impl.Id()
	return nil
}

func (s *AugmentorRpcServer) Augment(args interface{}, resp *string) error {
	return s.Impl.Augment()
}

// This is necessary for actually serving the interface implementation.
type AugmentorPlugin struct {
	// Impl is the interface
	Impl aug.TemplateAugmentor
}

func (p *AugmentorPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &AugmentorRpcServer{Impl: p.Impl}, nil
}

func (p *AugmentorPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &AugmentorRpcClient{client: c}, nil
}

var AugmentorManager = &PluginManager[aug.TemplateAugmentor]{
	plugins: map[string]*CfsPlugin[aug.TemplateAugmentor]{},
}
