# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Model Context Protocol (MCP) test implementation in Go. The project demonstrates both MCP server and client capabilities using the official `github.com/modelcontextprotocol/go-sdk` package. It implements a simple greeting service to showcase the MCP protocol.

## Architecture

The codebase follows a CLI-based architecture using Cobra:

- **cmd/main.go**: Entry point with Cobra root command that registers server and client subcommands
- **cmd/mcp-server.go**: MCP server implementation that exposes a "greet" tool via stdio transport
- **cmd/mcp-client.go**: MCP client that connects to the server via command transport and calls the greet tool
- **model/data_model.go**: Input/Output data structures with JSON schema annotations for MCP tool parameters

The server uses stdio transport for communication, while the client spawns the server as a subprocess using command transport. Tool handlers are registered using the generic `mcp.AddTool` function with typed input/output models.

## Build and Run

Build the binary:
```bash
go build -o mcp ./cmd
```

Run the MCP server (typically used via client or MCP host):
```bash
./mcp server
```

Run the MCP client (spawns server and calls the greet tool):
```bash
./mcp client
```

## Common Development Commands

Ensure dependencies are in sync:
```bash
go mod tidy
```

Run tests:
```bash
go test ./...
```

Run tests for a specific package:
```bash
go test ./cmd
```

## Dependencies

- **github.com/modelcontextprotocol/go-sdk**: Official MCP protocol implementation for Go
- **github.com/spf13/cobra**: CLI framework used for command structure (server/client subcommands)

## Environment

The project uses `.env` for configuration (already in .gitignore).

## Coding Style

Use comments sparingly. Only comment complex code.
- The datatype model schema is defined in the `model/` folder. Reference it anytime you need to understand the structure of data objects.