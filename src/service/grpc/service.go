package grpc

import (
	pb "portService/grpc"
)

func NewGrpcService(data ServiceData) *grpcService {
	return &grpcService{
		data: data,
	}
}

//service is the implementation of PortInfoServer
type grpcService struct {
	data ServiceData
	pb.UnimplementedPortInfoServer
}

//GetPortInfo streams information about open ports on the current system
func (s *grpcService) GetPortInfo(req *pb.PortInfoRequest, stream pb.PortInfo_GetPortInfoServer) error {

	data, err := s.data.Get()
	if err != nil {
		return err
	}

	for _, portInfo := range data {
		if err := stream.Send(portInfo); err != nil {
			return err
		}
	}

	return nil
}
