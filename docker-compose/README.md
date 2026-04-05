# Generic Docker Compose stack for cagent

This directory provides a batteries-included Docker Compose setup for running `cagent` with common MCP services and observability.

## What gets started

- `cagent` (`docker/cagent`) with example configs mounted from `../examples`.
- `duckduckgo` MCP server.
- `github-official` MCP server.
- Jaeger tracing UI and OTLP collector.
- Prometheus scraping `cagent` metrics.
- Grafana pre-provisioned with Prometheus as the default data source.

## Prerequisites

1. Docker Engine / Docker Desktop with `docker compose`.
2. Either:
   - the `docker/cagent` container image (used by this stack), or
   - a local `docker-agent` binary if you want to run commands outside Compose.
3. Optional provider credentials such as `OPENAI_API_KEY`, `ANTHROPIC_API_KEY`, and `GOOGLE_API_KEY`.

## Configure environment variables

Create a `.env` file in this `docker-compose/` directory:

```env
OPENAI_API_KEY=sk-...
ANTHROPIC_API_KEY=
GOOGLE_API_KEY=
GITHUB_TOKEN=ghp_...
```

Only set the keys you need.

## Start the full stack

From this directory:

```bash
docker compose -f docker-compose.yml up
```

This starts `cagent`, both MCP services, Jaeger, Prometheus, and Grafana.

## Run an example agent with this stack

The `cagent` service already mounts `../examples` at `/examples`.

To run an existing example manually:

```bash
docker compose -f docker-compose.yml run --rm cagent run --exec /examples/code.yaml "Summarize available tools"
```

## Observability endpoints

- Jaeger UI: <http://localhost:16686>
- Prometheus: <http://localhost:9090>
- Grafana: <http://localhost:3000> (default admin/admin)

`cagent` exposes Prometheus metrics on `/metrics` when `ENABLE_METRICS=true`.
