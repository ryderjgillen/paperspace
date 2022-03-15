package tests

import (
	"context"
	"net/http"
	"portService/server/prom"
	"testing"
	"time"
)

func Taest_promServer_cancel(t *testing.T) {

	//given
	promServer := prom.NewPromServer("", 0)
	ctx, cancel := context.WithCancel(context.Background())

	//when
	cancel()
	err := promServer.Run(ctx)

	//then
	if err != nil && err != context.Canceled {
		t.Fatalf(err.Error())
	}
}

func Test_promServer_200(t *testing.T) {

	//given
	promServer := prom.NewPromServer("", 65535)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)

	//when
	go func() {
		errCh <- promServer.Run(ctx)
	}()

	//then
	time.Sleep(1 * time.Second)
	res, err := http.Get("http://0.0.0.0:65535/metrics")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected: %d, got: %d", res.StatusCode, http.StatusOK)
	}
}
