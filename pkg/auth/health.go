package auth

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"

	"google.golang.org/grpc"
)

// Blueprint struct for db service and auth provider.
type HealthRPC struct {
	grpc_health_v1.UnimplementedHealthServer
}

// Check applies the logic to determine which variant to be applied
// to the given request.
func (rpc *HealthRPC) Check(ctx context.Context, r *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

// Register wires up the RPC to the server
func (rpc *HealthRPC) Register(server *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(server, rpc)
}

// NewRPC returns a new initialized RPC object.
func NewHealthRPC() RPCInterface {
	return &HealthRPC{}
}
