# Usage Notes

This repository supports adding MCP servers to agents in two primary ways:

- `command`: execute a local or containerized MCP process directly.
- `ref`: reference a Docker MCP catalog entry such as `docker:duckduckgo`.

For examples, see:

- `examples/mcp-definitions.yaml`
- `examples/github_issue_manager.yaml`
- `examples/blog.yaml`

## Extending the toolset

In addition to existing built-in and catalog MCP integrations, the codebase now includes placeholder connector scaffolding under `pkg/tools/connectors/` for:

- Trello (`list_boards`)
- Slack / Teams messaging (`send_message`)
- PostgreSQL querying (`run_query`)

These scaffolds are intended as future integrations and include TODO guidance for authentication, API usage, and result handling.
