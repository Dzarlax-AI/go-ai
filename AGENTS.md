# Project guidance

## Commands

- `go test -race -count=1 ./...`
- `go vet ./...`
- `go mod verify`

## Architecture

Package `ai` defines provider-neutral generation contracts. `providers/openai` and `providers/gemini` implement them; `internal/transport` handles validation, schema normalization and HTTP. Preserve one-attempt behavior, safe errors, usage semantics, and schema handling. Tests and examples use synthetic HTTP responses and must not call real providers.

## Environment variables

The library does not automatically read environment variables. Callers explicitly supply credentials through `ai.Config.APIKey`; do not put real credentials in tests, examples, or repository files.
