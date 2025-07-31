# CRUSH.md - Moonwalk Project Guide

## Build/Test Commands
```bash
go build -o moonwalk .                    # Build binary
go run main.go                           # Run from source
go run main.go -tui                      # Run TUI mode from source
go test ./...                           # Run all tests
go test -v ./slide                      # Run single package tests
go mod tidy                             # Clean dependencies
go fmt ./...                            # Format code
go vet ./...                            # Static analysis
```

## Code Style Guidelines

### Imports
- Standard library first, then third-party, then local packages
- Use blank line separation between groups
- Avoid dot imports

### Naming Conventions
- Use camelCase for variables and functions
- Use PascalCase for exported types and functions
- Use ALL_CAPS for constants
- Package names should be lowercase, single word

### Types & Error Handling
- Always handle errors explicitly, don't ignore them
- Use structured types for complex data (FileInfo, Statistics)
- Prefer composition over inheritance
- Use interfaces for testability

### Formatting
- Use gofmt for consistent formatting
- Line length should be reasonable (~100 chars)
- Use meaningful variable names over comments