package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"portService/server/grpc"
	"portService/server/prom"
	"syscall"
	"time"
)

type config struct {
	dataInterval time.Duration
	address      string
	port         uint16
	promAddress  string
	promPort     uint16
}

func parseConfig(args []string) (config, error) {

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	var (
		dataInterval = flags.Duration("data-interval", 15*time.Second, "port data refresh interval")
		address      = flags.String("address", "", "GRPC server listen address")
		port         = flags.Uint("port", 59001, "GRPC server listen port")
		promAddress  = flags.String("prom-address", "", "Prometheus metrics endpoint listen address")
		promPort     = flags.Uint("prom-port", 59002, "Prometheus metrics endpoint listen port")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return config{}, err
	}

	c := config{}
	c.dataInterval = *dataInterval
	c.address = *address
	c.port = uint16(*port)
	c.promAddress = *promAddress
	c.promPort = uint16(*promPort)

	return c, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT)

	config, err := parseConfig(os.Args)
	if err != nil {
		log.Fatalf("error reading config: %s", err.Error())
	}

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	errCh := make(chan error, 1)

	go func() {
		data := grpc.NewServiceData(ctx, config.dataInterval, grpc.GetPortInfo)
		grpcServer := grpc.NewGrpcServer(grpc.NewGrpcService(data), grpc.GrpcServerConfig{
			Address: config.address,
			Port:    config.port,
		})

		err := grpcServer.Run(ctx)
		if err != nil {
			errCh <- err
		}
	}()

	go func() {
		promServer := prom.NewPromServer(prom.PromServerConfig{
			Address: config.promAddress,
			Port:    config.promPort,
		})

		err := promServer.Run(ctx)
		if err != nil {
			errCh <- err
		}
	}()

	for {
		select {
		case sig := <-signalChan:
			switch sig {
			case syscall.SIGHUP:
				cancel()

				//cancel() unblocks errCh
				//waiting for err ensures server finishes handling ongoing requests
				log.Print((<-errCh).Error())

				//will retain same PID
				if err := syscall.Exec(os.Args[0], os.Args, os.Environ()); err != nil {
					log.Fatalf("error spawning process: %v", err)
				}
			case syscall.SIGINT:
				cancel()

				log.Fatal((<-errCh).Error())
			}
		case err = <-errCh:
			log.Fatal(err.Error())
		}
	}
}
