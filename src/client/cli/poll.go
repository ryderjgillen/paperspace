package cli

import (
	"context"
	"fmt"
	"io"
	"time"

	client "portService/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PollCmd struct {
	Interval    time.Duration
	Address     string
	Port        int `kong:"required,default='50051'"`
	PromAddress string
	PromPort    int
}

func (cmd *PollCmd) Run(_ *Context) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cmd.Address, cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	c := client.NewPortInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetPortInfo(ctx, &client.PortInfoRequest{})
	if err != nil {
		return fmt.Errorf("ERROR: %s\n", err.Error())
	}

	for {
		portInfo, err := r.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		} else {

			fmt.Printf("Protocol=%s, ", portInfo.Protocol.String())
			fmt.Printf("Source.IpAddress=%s:%s, ", portInfo.Source.IpAddress, formatPort(portInfo.Source.Port))
			fmt.Printf("Destination.IpAddress=%s:%s", portInfo.Destination.IpAddress, formatPort(portInfo.Destination.Port))

			fmt.Println()
		}
	}

	return nil
}

func formatPort(port uint32) string {
	if port == 0 {
		return "*"
	}

	return fmt.Sprintf("%d", port)
}
