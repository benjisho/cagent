package postgres

import (
	"context"

	"github.com/docker/docker-agent/pkg/tools"
)

const ToolNameRunQuery = "run_query"

// Server is a placeholder MCP-style connector for PostgreSQL operations.
//
// TODO:
//   - Add DSN and cloud-secret based authentication.
//   - Enforce read-only mode by default with optional write gating.
//   - Implement query execution, pagination, and result serialization.
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

type runQueryArgs struct {
	Query string `json:"query" jsonschema:"SQL query to execute against PostgreSQL."`
}

func (s *Server) runQuery(_ context.Context, _ runQueryArgs) (*tools.ToolCallResult, error) {
	return tools.ResultError("TODO: implement PostgreSQL connection management and SQL execution"), nil
}

// Tools returns placeholder tools exposed by the PostgreSQL connector.
func (s *Server) Tools(context.Context) ([]tools.Tool, error) {
	return []tools.Tool{
		{
			Name:         ToolNameRunQuery,
			Category:     "mcp",
			Description:  "Run a SQL query on a PostgreSQL database and return results.",
			Parameters:   tools.MustSchemaFor[runQueryArgs](),
			OutputSchema: tools.MustSchemaFor[string](),
			Handler:      tools.NewHandler(s.runQuery),
			Annotations: tools.ToolAnnotations{
				ReadOnlyHint: false,
				Title:        "Postgres run query",
			},
		},
	}, nil
}
