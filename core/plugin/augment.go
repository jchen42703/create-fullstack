package plugin

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/jchen42703/create-fullstack/core/aug"
	"github.com/jchen42703/create-fullstack/core/proto"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var AugmentPluginHandshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "AUG_PLUGIN",
	MagicCookieValue: "2f4b473f573429bca56f545f02c6bac47cebcc3130fbdaf5e1a14dda66349d93", // echo -n create-fullstack-aug-plugin | sha256sum
}

// RPC Client to get Augmentor function results from server
type AugmentorGrpcClient struct {
	client proto.TemplateAugmentorClient
}

func (g *AugmentorGrpcClient) Id() string {
	// return resp
	resp, err := g.client.Id(context.Background(), &proto.Empty{})
	if err != nil {
		// This is bad
		panic(err)
	}

	return resp.Id
}

func (g *AugmentorGrpcClient) Augment() error {
	_, err := g.client.Augment(context.Background(), &proto.Empty{})
	if err != nil {
		return err
	}

	return nil
}

// Here is the RPC server that AugmentorGrpcClient talks to, conforming to
// the requirements of net/rpc
type AugmentorGrpcServer struct {
	proto.UnimplementedTemplateAugmentorServer
	// This is the real implementation
	Impl aug.TemplateAugmentor
}

func (s *AugmentorGrpcServer) Id(ctx context.Context, req *proto.Empty) (*proto.IdResponse, error) {
	return &proto.IdResponse{
		Id: s.Impl.Id(),
	}, nil
}

func (s *AugmentorGrpcServer) Augment(ctx context.Context, req *proto.Empty) (*proto.Empty, error) {
	err := s.Impl.Augment()
	if err != nil {
		return &proto.Empty{}, err
	}

	return &proto.Empty{}, nil
}

// This is necessary for actually serving the interface implementation.
type AugmentorPlugin struct {
	plugin.Plugin
	// Impl is the interface
	Impl aug.TemplateAugmentor
}

func (p *AugmentorPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterTemplateAugmentorServer(s, &AugmentorGrpcServer{Impl: p.Impl})
	return nil
}

func (p *AugmentorPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &AugmentorGrpcClient{client: proto.NewTemplateAugmentorClient(c)}, nil
}

var AugmentorManager = &PluginManager[aug.TemplateAugmentor]{
	plugins: map[string]*CfsPlugin[aug.TemplateAugmentor]{},
}
