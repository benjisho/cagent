package messaging

import (
	"context"

	"github.com/docker/docker-agent/pkg/tools"
)

const ToolNameSendMessage = "send_message"

// Server is a placeholder MCP-style connector for Slack and Microsoft Teams.
//
// TODO:
//   - Add OAuth or bot token authentication for Slack and Teams.
//   - Map channels/conversations between providers.
//   - Implement provider-specific message formatting and attachments.
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

type sendMessageArgs struct {
	Provider string `json:"provider" jsonschema:"Messaging provider: slack or teams."`
	Channel  string `json:"channel" jsonschema:"Channel, room, or conversation identifier."`
	Message  string `json:"message" jsonschema:"Message text payload to send."`
}

func (s *Server) sendMessage(_ context.Context, args sendMessageArgs) (*tools.ToolCallResult, error) {
	_ = args
	return tools.ResultError("TODO: implement Slack/Teams authentication and outbound messaging APIs"), nil
}

// Tools returns placeholder tools exposed by the messaging connector.
func (s *Server) Tools(context.Context) ([]tools.Tool, error) {
	return []tools.Tool{
		{
			Name:         ToolNameSendMessage,
			Category:     "mcp",
			Description:  "Send a notification message to Slack or Microsoft Teams.",
			Parameters:   tools.MustSchemaFor[sendMessageArgs](),
			OutputSchema: tools.MustSchemaFor[string](),
			Handler:      tools.NewHandler(s.sendMessage),
			Annotations: tools.ToolAnnotations{
				ReadOnlyHint: false,
				Title:        "Messaging send message",
			},
		},
	}, nil
}
