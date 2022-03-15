package tests

import (
	"context"
	"portService/server/grpc"
	"testing"

	pb "portService/grpc"

	g "google.golang.org/grpc"
)

type portInfo_GetPortInfoServer struct {
	g.ServerStream
	ctx            context.Context
	recvToServer   chan *pb.PortInfoRequest
	sentFromServer chan *pb.PortInfoResponse
}

func (s portInfo_GetPortInfoServer) Send(resp *pb.PortInfoResponse) error {
	s.sentFromServer <- resp

	return nil
}

type serviceData struct {
	data []*pb.PortInfoResponse
}

func (d serviceData) Get() []*pb.PortInfoResponse {
	return d.data
}

func aTest_GrpcService(t *testing.T) {

	//given
	data := []*pb.PortInfoResponse{
		{Command: "cmd1", Protocol: pb.Protocol_UDP, Source: &pb.IpPort{IpAddress: "source1", Port: 1}, Destination: &pb.IpPort{IpAddress: "dest1", Port: 2}},
		{Command: "cmd2", Protocol: pb.Protocol_TCP, Source: &pb.IpPort{IpAddress: "source2", Port: 3}, Destination: &pb.IpPort{IpAddress: "dest2", Port: 4}},
	}

	service := grpc.NewGrpcService(&serviceData{
		data: data,
	})

	stream := &portInfo_GetPortInfoServer{
		ctx:            context.Background(),
		recvToServer:   make(chan *pb.PortInfoRequest, 10),
		sentFromServer: make(chan *pb.PortInfoResponse, 10),
	}

	//when
	err := service.GetPortInfo(&pb.PortInfoRequest{}, stream)
	close(stream.sentFromServer)
	close(stream.recvToServer)

	//then
	if err != nil {
		t.Fatal(err.Error())
	}

	count := 0
	for msg := range stream.sentFromServer {
		if data[count] != msg {
			t.Fatal()
		}

		count++
	}

	if count != len(data) {
		t.Fatal()
	}
}
