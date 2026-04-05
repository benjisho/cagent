package metrics

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	registerMetrics sync.Once

	agentRunsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cagent_agent_runs_total",
			Help: "Total number of agent runtime executions.",
		},
		[]string{"result"},
	)

	agentRunDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cagent_agent_run_duration_seconds",
			Help:    "Duration of agent runtime executions.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"result"},
	)

	taskDelegationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cagent_task_delegations_total",
			Help: "Total number of task delegations between agents.",
		},
		[]string{"from_agent", "to_agent", "result"},
	)

	toolInvocationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cagent_mcp_tool_invocations_total",
			Help: "Total number of MCP and runtime tool invocations.",
		},
		[]string{"tool", "agent", "result"},
	)

	toolInvocationDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cagent_mcp_tool_invocation_duration_seconds",
			Help:    "Duration of MCP and runtime tool invocations.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"tool", "agent", "result"},
	)
)

func registerMetricsCollectors() {
	registerMetrics.Do(func() {
		prometheus.MustRegister(
			agentRunsTotal,
			agentRunDuration,
			taskDelegationsTotal,
			toolInvocationsTotal,
			toolInvocationDuration,
		)
	})
}

func RecordAgentRun(duration time.Duration, err error) {
	registerMetricsCollectors()
	result := resultLabel(err)
	agentRunsTotal.WithLabelValues(result).Inc()
	agentRunDuration.WithLabelValues(result).Observe(duration.Seconds())
}

func RecordTaskDelegation(fromAgent, toAgent string, err error) {
	registerMetricsCollectors()
	taskDelegationsTotal.WithLabelValues(fromAgent, toAgent, resultLabel(err)).Inc()
}

func RecordToolInvocation(toolName, agentName string, duration time.Duration, err error) {
	registerMetricsCollectors()
	result := resultLabel(err)
	toolInvocationsTotal.WithLabelValues(toolName, agentName, result).Inc()
	toolInvocationDuration.WithLabelValues(toolName, agentName, result).Observe(duration.Seconds())
}

func Handler() http.Handler {
	registerMetricsCollectors()
	return promhttp.Handler()
}

func StartServer(addr string) error {
	if addr == "" {
		addr = ":9090"
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", Handler())

	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("metrics server failed: %w", err)
	}

	return nil
}

func resultLabel(err error) string {
	if err != nil {
		return "error"
	}
	return "success"
}
