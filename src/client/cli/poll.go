package cli

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	client "portService/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PollCmd struct {
	Interval    time.Duration `kong:"default='15s'"`
	Address     string        `kong:"default=''"`
	Port        int           `kong:"default='59001'"`
	PromAddress string        `kong:"default=''"`
	PromPort    int           `kong:"default='59002'"`
	ShowMetrics bool          `kong:"default='true'"`
}

func (cmd *PollCmd) Run(_ *Context) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cmd.Address, cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	c := client.NewPortInfoClient(conn)
	for {

		err := displayPortInfo(c)
		if err != nil {
			return err
		}

		if cmd.ShowMetrics {
			err = displayMetrics(cmd.PromAddress, cmd.PromPort)
			if err != nil {
				return err
			}
		}

		time.Sleep(cmd.Interval)
	}
}

func displayPortInfo(c client.PortInfoClient) error {

	formatPort := func(port uint32) string {
		if port == 0 {
			return "*"
		}

		return fmt.Sprintf("%d", port)
	}

	truncateText := func(s string, max int) string {
		if max > len(s) {
			return s
		}
		return fmt.Sprintf("%s...", s[:max])
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Printf("**************************************** %s: %s ****************************************\n\n", "Polling", time.Now().UTC())
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
			fmt.Printf("Destination.IpAddress=%s:%s, ", portInfo.Destination.IpAddress, formatPort(portInfo.Destination.Port))
			fmt.Printf("Command=%s", truncateText(portInfo.Command, 50))

			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println()

	return nil
}

func displayMetrics(address string, port int) error {

	fmt.Printf("**************************************** %s: %s ****************************************\n\n", "Metrics", time.Now().UTC())

	resp, err := http.Get(fmt.Sprintf("http://%s:%d/metrics", address, port))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
