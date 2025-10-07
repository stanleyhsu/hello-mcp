package main

import (
	"context"
	"git.kldmp.com/Golang/hello-mcp/model"
	"github.com/spf13/cobra"

	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func SayHello(ctx context.Context, req *mcp.CallToolRequest, input model.Input) (*mcp.CallToolResult, model.Output, error) {
	return nil, model.Output{Greeting: "Hi " + input.Name}, nil
}

var (
	// McpServerCmd represents the server command
	McpServerCmd = &cobra.Command{
		Use: "server",
		Run: runServer,
	}
)

func runServer(cmd *cobra.Command, args []string) {
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "say hi"}, SayHello)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
