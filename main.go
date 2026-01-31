package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/docker/cagent/cmd/root"
	"github.com/docker/cagent/pkg/metrics"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	if metrics.Enabled() {
		go func() {
			addr := metrics.Address()
			slog.Info("Starting metrics server", "addr", addr)
			if err := metrics.StartServer(ctx, addr); err != nil {
				slog.Error("Metrics server error", "error", err)
			}
		}()
	}

	if err := root.Execute(ctx, os.Stdin, os.Stdout, os.Stderr, os.Args[1:]...); err != nil {
		cancel()
		os.Exit(1)
	} else {
		cancel()
		os.Exit(0)
	}
}
