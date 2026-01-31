package trello

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/docker/cagent/pkg/tools"
)

// ListBoardsInput describes the inputs for listing Trello boards.
type ListBoardsInput struct {
	OrganizationID string `json:"organization_id,omitempty" jsonschema:"ID of the Trello organization (optional)"`
}

// Board represents a minimal Trello board summary.
type Board struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewServer returns an MCP server that exposes Trello scaffolding tools.
// TODO: Wire Trello OAuth (API key + token) and call the Trello REST API.
func NewServer() *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "cagent-trello",
		Version: "0.1.0",
	}, nil)

	tool := &mcp.Tool{
		Name:         "list_boards",
		Description:  "List Trello boards for a user or organization.",
		InputSchema:  tools.MustSchemaFor[ListBoardsInput](),
		OutputSchema: tools.MustSchemaFor[[]Board](),
	}

	mcp.AddTool(server, tool, func(_ context.Context, _ *mcp.CallToolRequest, _ ListBoardsInput) (*mcp.CallToolResult, []Board, error) {
		return nil, nil, fmt.Errorf("trello connector not configured: TODO add authentication and API calls")
	})

	return server
}
