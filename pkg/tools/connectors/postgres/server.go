package postgres

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/docker/cagent/pkg/tools"
)

// RunQueryInput describes the payload for running a SQL query.
type RunQueryInput struct {
	Query string `json:"query" jsonschema:"SQL query to execute"`
}

// RunQueryOutput represents a placeholder query response.
type RunQueryOutput struct {
	Rows any `json:"rows"`
}

// NewServer returns an MCP server that exposes Postgres scaffolding tools.
// TODO: Connect to Postgres with a DSN and use database/sql to execute queries safely.
func NewServer() *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "cagent-postgres",
		Version: "0.1.0",
	}, nil)

	tool := &mcp.Tool{
		Name:         "run_query",
		Description:  "Run a SQL query against a Postgres database.",
		InputSchema:  tools.MustSchemaFor[RunQueryInput](),
		OutputSchema: tools.MustSchemaFor[RunQueryOutput](),
	}

	mcp.AddTool(server, tool, func(_ context.Context, _ *mcp.CallToolRequest, _ RunQueryInput) (*mcp.CallToolResult, RunQueryOutput, error) {
		return nil, RunQueryOutput{}, fmt.Errorf("postgres connector not configured: TODO add database connection and query execution")
	})

	return server
}
