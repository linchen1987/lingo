# AGENTS.md - Lin CLI Development Guide

## Project Overview

Lin is a CLI tool collection. Built with Cobra for CLI framework and Bubble Tea for TUI.

## Build Commands

```bash
# Build the project
go build -o lin .

# Build with optimizations (for release)
go build -ldflags="-s -w" -o lin .

# Run without building
go run .

# Run specific command
go run . datetime 1772103423
```

## Lint and Format Commands

```bash
# Format code
go fmt ./...

# Run go vet (static analysis)
go vet ./...

# Run all checks (recommended before committing)
go fmt ./... && go vet ./... && go test ./...
```

## Project Structure

```
lin-cli-go/
├── main.go              # Entry point - calls cmd.Execute()
├── cmd/                 # CLI commands (Cobra)
│   ├── root.go          # Root command, registers subcommands
│   ├── datetime.go      # datetime command
│   ├── base58.go        # base58 command
│   ├── base64.go        # base64 command
│   └── tui.go           # TUI implementation (Bubble Tea)
├── internal/            # Private packages (not importable externally)
│   └── tools/           # Core tool implementations
│       ├── datetime.go  # Timestamp/datetime conversion
│       ├── base58.go    # Base58 encode/decode
│       └── base64.go    # Base64 encode/decode
├── go.mod               # Module definition
├── go.sum               # Dependency checksums
└── .gitignore           # Git ignore rules
```

## Code Style Guidelines

### Adding New Tools

1. Create command in `cmd/xxx.go`:

```go
var xxxCmd = &cobra.Command{...}

func init() {
	rootCmd.AddCommand(xxxCmd)
}
```

2. Register in TUI by adding to `toolsList` in `cmd/tui.go`:

```go
{name: "xxx", description: "Description", actions: []string{"encode", "decode"}},
```


