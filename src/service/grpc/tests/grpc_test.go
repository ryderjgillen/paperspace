package tests

import (
	"context"
	pb "portService/grpc"
	"portService/server/grpc"
	"testing"
	"time"
)

func Test_GrpcServer(t *testing.T) {

	//service := grpc.NewGrpcService(nil)
	data := []*pb.PortInfoResponse{
		{Command: "cmd1", Protocol: pb.Protocol_UDP, Source: &pb.IpPort{IpAddress: "source1", Port: 1}, Destination: &pb.IpPort{IpAddress: "dest1", Port: 2}},
		{Command: "cmd2", Protocol: pb.Protocol_TCP, Source: &pb.IpPort{IpAddress: "source2", Port: 3}, Destination: &pb.IpPort{IpAddress: "dest2", Port: 4}},
	}

	service := grpc.NewGrpcService(&serviceData{
		data: data,
	})

	server := grpc.NewGrpcServer(service, grpc.GrpcServerConfig{})
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()
	err := server.Run(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
}
