package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/docker/cli/cli"

	"github.com/docker/docker-agent/cmd/root"
	"github.com/docker/docker-agent/pkg/metrics"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	maybeStartMetricsServer()

	if err := root.Execute(ctx, os.Stdin, os.Stdout, os.Stderr, os.Args[1:]...); err != nil {
		cancel()
		if statusErr, ok := errors.AsType[cli.StatusError](err); ok {
			os.Exit(statusErr.StatusCode)
		}
		os.Exit(1)
	} else {
		cancel()
		os.Exit(0)
	}
}

func maybeStartMetricsServer() {
	if !isMetricsEnabled() {
		return
	}

	addr := strings.TrimSpace(os.Getenv("METRICS_ADDR"))
	if addr == "" {
		addr = ":9090"
	}

	go func() {
		slog.Info("Starting Prometheus metrics endpoint", "address", addr, "path", "/metrics")
		if err := metrics.StartServer(addr); err != nil {
			slog.Error("Prometheus metrics endpoint stopped", "error", err)
		}
	}()
}

func isMetricsEnabled() bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("ENABLE_METRICS"))) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}
