package prom

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PromServerConfig struct {
	Address string
	Port    uint16
}

type promServer struct {
	httpServer *http.Server
}

func NewPromServer(config PromServerConfig) promServer {

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	service := promServer{
		httpServer: &http.Server{Addr: fmt.Sprintf("%s:%d", config.Address, config.Port), Handler: mux},
	}
	return service
}

func (s promServer) Run(ctx context.Context) error {

	errCh := make(chan error, 1)

	go func() {
		select {
		case <-ctx.Done():
			errCh <- s.httpServer.Shutdown(ctx)
		}
	}()

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			errCh <- err
		}
	}()

	return <-errCh
}
