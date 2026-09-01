# AGENTS.md

## Project overview

This repository is a Go CLI for downloading OpenStreetMap extracts from multiple providers. The binary entrypoint lives in `cmd/download-geofabrik`, while the actual logic is organized under `internal/` and `pkg/`.

The most important runtime modules are:

- `internal/cli`: Cobra commands and CLI setup (`download`, `generate`, `list`, `mcp`, root flags)
- `internal/config`: default config/service naming
- `internal/provider`: provider-specific scrapers and catalog APIs
- `internal/mcp`: Model Context Protocol tools/resources served by the CLI
- `internal/generator`: config generation logic
- `pkg/catalog` and `pkg/formats`: shared catalog/format definitions

## Working conventions

- Prefer small, focused Go changes that fit the existing package boundaries.
- Keep provider implementations consistent with the `Provider` interface defined in `internal/provider/provider.go`.
- Config files are YAML and provider-specific; existing filenames are usually something like `geofabrik.yml`, `bbbike.yml`, etc.
- Keep CLI behavior aligned with Cobra/Viper conventions already used in `internal/cli/root.go`.
- For MCP-related changes, preserve the schema and tool/resource design already defined under `internal/mcp`.

## Build and verification

Run the project tests from the repo root:

```bash
go test ./...
```

Build the CLI binary:

```bash
go build ./cmd/download-geofabrik
```

The CLI is typically executed as:

```bash
./download-geofabrik --help
./download-geofabrik download element
./download-geofabrik mcp
```

## Testing expectations

- Tests are Go unit tests using `testing` and `testify/assert`.
- Prefer table-driven tests and keep assertions specific to the real behavior.
- When changing provider scraping or config generation, update or add tests in the relevant provider package.
- Providers code should target 100% branch/statement coverage, or as close as reasonably possible. If a branch is genuinely impossible or non-sensical to test, document the reason and keep the remaining uncovered area minimal.
- When Go code is modified, run a benchmark before commit and compare the results against the previous state to detect regressions.
- The benchmark must be done at the end of the work, just before commit, and only if a `.go` file changed.

## Files to inspect first

- `README.md` for user-facing behavior and supported services
- `internal/cli/root.go` for CLI registration and global flags
- `internal/provider/provider.go` for the provider contract
- `internal/mcp/tools.go` and `internal/mcp/resources.go` for tool/resource APIs

## Avoid

- Broad refactors without a clear need.
- Adding new dependencies for simple logic.
- Modifying provider selection or config filenames without updating the corresponding tests and docs.
