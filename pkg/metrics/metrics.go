package metrics

import (
	"context"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultMetricsAddr = ":9090"
)

var (
	registerOnce sync.Once

	agentRunsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cagent_agent_runs_total",
			Help: "Total number of agent runs started.",
		},
		[]string{"agent"},
	)

	agentRunDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cagent_agent_run_duration_seconds",
			Help:    "Duration of agent runs in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"agent"},
	)

	taskDelegationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cagent_task_delegations_total",
			Help: "Total number of task delegations between agents.",
		},
		[]string{"from_agent", "to_agent"},
	)

	taskDelegationDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cagent_task_delegation_duration_seconds",
			Help:    "Duration of task delegations in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"from_agent", "to_agent"},
	)

	mcpToolInvocationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cagent_mcp_tool_invocations_total",
			Help: "Total number of MCP tool invocations.",
		},
		[]string{"tool", "agent"},
	)

	mcpToolInvocationDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cagent_mcp_tool_invocation_duration_seconds",
			Help:    "Duration of MCP tool invocations in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"tool", "agent"},
	)
)

func Register() {
	registerOnce.Do(func() {
		prometheus.MustRegister(
			agentRunsTotal,
			agentRunDuration,
			taskDelegationsTotal,
			taskDelegationDuration,
			mcpToolInvocationsTotal,
			mcpToolInvocationDuration,
		)
	})
}

func Handler() http.Handler {
	Register()
	return promhttp.Handler()
}

func Enabled() bool {
	value := strings.TrimSpace(strings.ToLower(os.Getenv("ENABLE_METRICS")))
	return value != "" && value != "0" && value != "false"
}

func Address() string {
	if addr := strings.TrimSpace(os.Getenv("METRICS_ADDR")); addr != "" {
		return addr
	}
	return defaultMetricsAddr
}

func StartServer(ctx context.Context, addr string) error {
	Register()

	server := &http.Server{
		Addr:    addr,
		Handler: promhttp.Handler(),
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	case err := <-errCh:
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	}
}

func RecordAgentRun(agent string, duration time.Duration) {
	Register()
	agentRunsTotal.WithLabelValues(agent).Inc()
	agentRunDuration.WithLabelValues(agent).Observe(duration.Seconds())
}

func RecordTaskDelegation(fromAgent, toAgent string, duration time.Duration) {
	Register()
	taskDelegationsTotal.WithLabelValues(fromAgent, toAgent).Inc()
	taskDelegationDuration.WithLabelValues(fromAgent, toAgent).Observe(duration.Seconds())
}

func RecordMCPToolInvocation(toolName, agent string, duration time.Duration) {
	Register()
	mcpToolInvocationsTotal.WithLabelValues(toolName, agent).Inc()
	mcpToolInvocationDuration.WithLabelValues(toolName, agent).Observe(duration.Seconds())
}
