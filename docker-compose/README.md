# Docker Compose quickstart

This folder provides a generic `docker-compose.yml` that launches `cagent` alongside
common MCP services and observability tooling (Jaeger, Prometheus, Grafana).

## Prerequisites

- Docker Desktop or Docker Engine with Docker Compose.
- The `docker/cagent` image (pulled automatically by Docker Compose).
- Optional API keys for the model providers you plan to use.

## Configure environment variables

Create a `.env` file in this `docker-compose/` directory:

```bash
OPENAI_API_KEY=your-openai-key
ANTHROPIC_API_KEY=your-anthropic-key
GOOGLE_API_KEY=your-gemini-key
GITHUB_TOKEN=your-github-token

# Optional overrides
ENABLE_METRICS=true
METRICS_ADDR=:9090
MCP_DUCKDUCKGO_IMAGE=docker/mcp-server-duckduckgo:latest
MCP_GITHUB_IMAGE=docker/mcp-server-github-official:latest
```

> **Note**
> The MCP server images above are placeholders. If your environment uses a
> different registry or image tag, update the `.env` file accordingly.

## Start the stack

From this directory:

```bash
docker compose up -d
```

This starts:

- `cagent` (API server on port `8080`)
- MCP servers for DuckDuckGo and GitHub (ports `8081` and `8082`)
- Jaeger tracing (`16686` UI)
- Prometheus (`9090`)
- Grafana (`3000`)

## Run an example agent

The compose file mounts `../examples` into the container as `/examples`.
To run an agent using the running container:

```bash
docker compose exec cagent cagent run /examples/code.yaml "Summarize the repo"
```

If you want to update the agent file, edit it locally under `../examples` and rerun
`docker compose exec`.

## Access the UIs

- Jaeger: http://localhost:16686
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000 (default user/pass: `admin` / `admin`)

## Metrics and dashboards

The `ENABLE_METRICS` flag enables the `/metrics` endpoint in the `cagent` container.
Prometheus is preconfigured to scrape `cagent:9090`, and Grafana is provisioned
with a Prometheus data source pointing at `http://prometheus:9090`.
