package auth

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	envoyAuthV2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoyAuthV3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
)

// ##################### V3 Auth GRPC #################### //

type RPCv2 struct {
	envoyAuthV2.UnimplementedAuthorizationServer
}

func (rpc *RPCv2) Check(ctx context.Context, r *envoyAuthV2.CheckRequest) (*envoyAuthV2.CheckResponse, error) {
	if attr := r.GetAttributes(); attr != nil {
		log.Debugf("request attributes.ContextExtension: %v", attr.GetContextExtensions())
		log.Debugf("request attributes.Destination: %v", attr.GetDestination())
		log.Debugf("request attributes.MetadataContext: %v", attr.GetMetadataContext())
		log.Debugf("request attributes.Request: %v", attr.GetRequest())
		log.Debugf("request attributes.Source: %v", attr.GetSource())
	}

	return &envoyAuthV2.CheckResponse{
		HttpResponse: &envoyAuthV2.CheckResponse_OkResponse{
			OkResponse: &envoyAuthV2.OkHttpResponse{},
		},
		Status: &status.Status{
			Code:    int32(codes.OK),
			Message: "authorized",
		},
	}, nil
}

// Register wires up the RPC to the server
func (rpc *RPCv2) Register(server *grpc.Server) {
	envoyAuthV2.RegisterAuthorizationServer(server, rpc)
}

// NewRPCv2 returns a new initialized RPC object.
func NewRPCv2() RPCInterface {
	return &RPCv2{}
}

// ##################### V3 Auth GRPC #################### //

type RPCv3 struct {
	envoyAuthV3.UnimplementedAuthorizationServer
}

func (rpc *RPCv3) Check(ctx context.Context, r *envoyAuthV3.CheckRequest) (*envoyAuthV3.CheckResponse, error) {
	if attr := r.GetAttributes(); attr != nil {
		log.Debugf("request attributes.ContextExtension: %v", attr.GetContextExtensions())
		log.Debugf("request attributes.Destination: %v", attr.GetDestination())
		log.Debugf("request attributes.MetadataContext: %v", attr.GetMetadataContext())
		log.Debugf("request attributes.Request: %v", attr.GetRequest())
		log.Debugf("request attributes.Source: %v", attr.GetSource())
	}

	return &envoyAuthV3.CheckResponse{
		HttpResponse: &envoyAuthV3.CheckResponse_OkResponse{
			OkResponse: &envoyAuthV3.OkHttpResponse{},
		},
		Status: &status.Status{
			Code:    int32(codes.OK),
			Message: "authorized",
		},
	}, nil
}

// Register wires up the RPC to the server
func (rpc *RPCv3) Register(server *grpc.Server) {
	envoyAuthV3.RegisterAuthorizationServer(server, rpc)
}

// NewRPCv3 returns a new initialized RPC object.
func NewRPCv3() RPCInterface {
	return &RPCv3{}
}
