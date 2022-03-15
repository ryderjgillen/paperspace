package grpc

import (
	"context"
	"fmt"
	"net"

	pb "portService/grpc"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc"
)

type GrpcServerConfig struct {
	Address string
	Port    uint16
}

type grpcServer struct {
	config      GrpcServerConfig
	grpcServer  *grpc.Server
	grpcService *grpcService
}

func NewGrpcServer(grpcService *grpcService, config GrpcServerConfig) grpcServer {

	grpcServer := grpcServer{
		config:      config,
		grpcService: grpcService,
		grpcServer: grpc.NewServer(
			grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		),
	}

	pb.RegisterPortInfoServer(grpcServer.grpcServer, grpcServer.grpcService)
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer.grpcServer)

	return grpcServer
}

func (s *grpcServer) Run(ctx context.Context) error {

	go func() {
		select {
		case <-ctx.Done():
			s.grpcServer.GracefulStop()
		}
	}()

	var lc net.ListenConfig
	lis, err := lc.Listen(ctx, "tcp", fmt.Sprintf("%s:%d", s.config.Address, s.config.Port))

	if err != nil {
		return err
	} else if err = s.grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
