package grpcserver

import (
	"net"

	"google.golang.org/grpc"
)

// GRPCServer adapter is an interface for the gRPC server to make it easier to mock for testing.
//
//go:generate mockgen -destination=mocks/mock_grpc.go -package=mocks -source=grpcserver.go
type GRPCServer interface {
	GracefulStop()
	Serve(listener net.Listener) error
}

type RealGRPCServer struct {
	Server *grpc.Server
}

func (r *RealGRPCServer) GracefulStop() {
	r.Server.GracefulStop()
}

func (r *RealGRPCServer) Serve(listener net.Listener) error {
	return r.Server.Serve(listener)
}
