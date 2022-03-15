package tests

import (
	"context"
	"fmt"
	pb "portService/grpc"
	"portService/server/grpc"
	"testing"
	"time"
)

func compare(a *pb.PortInfoResponse, b *pb.PortInfoResponse) bool {
	if a.Command != b.Command {
		return false
	}

	if a.Protocol != b.Protocol {
		return false
	}

	if a.Source.IpAddress != b.Source.IpAddress {
		return false
	}

	if a.Source.Port != b.Source.Port {
		return false
	}

	return true
}

func Test_IntervalData(t *testing.T) {

	//given
	ctx := context.Background()
	expectedData := []*pb.PortInfoResponse{}
	expectedErr := (error)(nil)

	serviceData := grpc.NewServiceData(ctx, 1*time.Second, func() ([]*pb.PortInfoResponse, error) {
		return expectedData, expectedErr
	})

	//when
	data, err := serviceData.Get()

	//then
	if len(expectedData) != len(data) {
		t.Fatalf("expected: data.len->%d, got: data.len->%d", len(expectedData), len(data))
	}

	for idx := range data {
		if compare(expectedData[idx], data[idx]) {
			t.Fatalf("expected: data->%v, got: data->%s", expectedData[idx], data[idx])
		}
	}

	if expectedErr != err {
		t.Fatalf("expected: error->nil, got: error->%s", err.Error())
	}
}

func Test_IntervalData_Error(t *testing.T) {

	//given
	ctx := context.Background()
	expectedData := ([]*pb.PortInfoResponse)(nil)
	expectedErr := fmt.Errorf("test error")

	serviceData := grpc.NewServiceData(ctx, 1*time.Second, func() ([]*pb.PortInfoResponse, error) {
		return expectedData, expectedErr
	})

	//when
	data, err := serviceData.Get()

	//then
	if len(expectedData) != len(data) {
		t.Fatalf("expected: data.len->%d, got: data.len->%d", len(expectedData), len(data))
	}

	for idx := range data {
		if compare(expectedData[idx], data[idx]) {
			t.Fatalf("expected: data->%v, got: data->%s", expectedData[idx], data[idx])
		}
	}

	if expectedErr != err {
		t.Fatalf("expected: error->%s, got: error->%s", expectedErr.Error(), err.Error())
	}
}
