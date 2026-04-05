package trello

import (
	"context"

	"github.com/docker/docker-agent/pkg/tools"
)

const ToolNameListBoards = "list_boards"

// Server is a placeholder MCP-style connector for Trello integration.
//
// TODO:
//   - Add Trello API key and token authentication.
//   - Resolve user/account scoped board access.
//   - Return structured board metadata from Trello's REST API.
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

type listBoardsArgs struct {
	WorkspaceID string `json:"workspace_id,omitempty" jsonschema:"Optional Trello workspace id to scope board listing."`
}

func (s *Server) listBoards(_ context.Context, _ listBoardsArgs) (*tools.ToolCallResult, error) {
	return tools.ResultError("TODO: implement Trello authentication and board listing API calls"), nil
}

// Tools returns placeholder tools exposed by the Trello connector.
func (s *Server) Tools(context.Context) ([]tools.Tool, error) {
	return []tools.Tool{
		{
			Name:         ToolNameListBoards,
			Category:     "mcp",
			Description:  "List Trello boards visible to the configured account.",
			Parameters:   tools.MustSchemaFor[listBoardsArgs](),
			OutputSchema: tools.MustSchemaFor[string](),
			Handler:      tools.NewHandler(s.listBoards),
			Annotations: tools.ToolAnnotations{
				ReadOnlyHint: true,
				Title:        "Trello list boards",
			},
		},
	}, nil
}
