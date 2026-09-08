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

Run project tests:

```bash
# Fast check for modified packages only
go test ./internal/cli/... ./internal/provider/osmtw/...

# Full test suite
go test ./...
```

Run linter (strictly enforced rules: `varnamelen`, `goconst`, `wsl_v5`, `gci`):

```bash
golangci-lint run
```

Build the CLI binary:

```bash
go build ./cmd/download-geofabrik
```

The CLI is typically executed as:

```bash
./download-geofabrik --help
./download-geofabrik download element
./download-geofabrik download element -n  # dry-run / no-download (always use for testing!)
./download-geofabrik mcp
```

## Testing and linting expectations

- Tests are Go unit tests using `testing` and `testify/assert`.
- Prefer table-driven tests and keep assertions specific to the real behavior.
- Providers code should target 100% branch/statement coverage, or as close as reasonably possible.
- **Linting conventions to respect on first edit**:
  - Variable naming (`varnamelen`): avoid single-letter loop variables like `f`, `k`, `v` in non-trivial scopes; use descriptive names (e.g., `format`, `elementID`).
  - String constants (`goconst`): define a `const` for any string literal repeated 3 or more times (e.g. paths, prefixes).
  - Whitespace (`wsl`): leave blank lines before declarations (`var`), `return`, control flow statements (`for`, `if`, `switch`), and assignments unless tightly coupled.
  - Formatting (`gci`, `gofmt`): avoid extraneous blank lines at end of files or blocks.
- Run `golangci-lint run` early right after modifying code, before benchmarks.
- **Benchmarks**:
  - Run benchmarks only if a `.go` file changed, just before commit.
  - **Target only the modified package(s)** (e.g., `go test -bench=. -benchmem ./internal/cli`) rather than `./...` to avoid long execution times, massive log outputs, and excessive token usage.

## Agent workflow & Token optimization guidelines

To minimize back-and-forth and context token consumption:

1. **Investigating download / provider issues**:
   - Always check `<service>.yml` first to view the provider's defined formats and element list before diving into code.
   - Always reproduce with `-n` (`--nodownload`) to avoid triggering gigabyte-sized extract downloads.
   - Use targeted `grep_search` on the specific error message or symbol instead of browsing entire directories or files.
2. **File viewing**:
   - Target precise line ranges (`StartLine` / `EndLine`) when viewing files. Avoid dumping full 400+ line files when only a specific function is relevant.
3. **Write once, right once**:
   - When adding or updating formats in providers, remember to update:
     - `DefaultFormats()` in `internal/provider/<service>/scraper.go` (include `.md5` if supported by the provider)
     - `parseLink` / `FetchCatalog` in the scraper
     - `<service>.yml`
     - Provider tests in `scraper_test.go`
4. **Verification sequence**:
   - 1. Run targeted tests: `go test ./internal/<modified_pkg>`
   - 2. Run linter: `golangci-lint run` (fix any issues immediately)
   - 3. Run targeted benchmarks on modified packages only: `go test -bench=. -benchmem ./internal/<modified_pkg>`

## Files to inspect first

- `README.md` for user-facing behavior and supported services
- `internal/cli/root.go` for CLI registration and global flags
- `internal/cli/download.go` for download execution and format resolution
- `internal/provider/provider.go` for the provider contract
- `internal/mcp/tools.go` and `internal/mcp/resources.go` for tool/resource APIs

## Avoid

- Broad refactors without a clear need.
- Running full-suite benchmarks `go test -bench=. ./...` (consumes huge context and time; target modified packages).
- Adding new dependencies for simple logic.
- Modifying provider selection or config filenames without updating the corresponding tests and docs.
