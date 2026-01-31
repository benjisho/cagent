package slack

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/docker/cagent/pkg/tools"
)

// SendMessageInput describes the payload for sending a Slack message.
type SendMessageInput struct {
	ChannelID string `json:"channel_id" jsonschema:"Slack channel ID"`
	Text      string `json:"text" jsonschema:"Message text to send"`
}

// SendMessageOutput represents a minimal response payload.
type SendMessageOutput struct {
	MessageID string `json:"message_id"`
}

// NewServer returns an MCP server that exposes Slack scaffolding tools.
// TODO: Authenticate with Slack OAuth tokens and call chat.postMessage.
func NewServer() *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "cagent-slack",
		Version: "0.1.0",
	}, nil)

	tool := &mcp.Tool{
		Name:         "send_message",
		Description:  "Send a Slack message to a channel.",
		InputSchema:  tools.MustSchemaFor[SendMessageInput](),
		OutputSchema: tools.MustSchemaFor[SendMessageOutput](),
	}

	mcp.AddTool(server, tool, func(_ context.Context, _ *mcp.CallToolRequest, _ SendMessageInput) (*mcp.CallToolResult, SendMessageOutput, error) {
		return nil, SendMessageOutput{}, fmt.Errorf("slack connector not configured: TODO add authentication and API calls")
	})

	return server
}
