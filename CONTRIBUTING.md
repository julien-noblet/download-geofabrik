# Contributing to download-geofabrik

Thank you for contributing to `download-geofabrik`! We welcome bug fixes, documentation improvements, new OSM catalog providers, and performance optimizations.

Please take a few minutes to read this guide before submitting changes.

---

## Table of Contents

- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Development Workflow](#development-workflow)
- [Coding Guidelines & Linting](#coding-guidelines--linting)
- [Running Tests & Benchmarks](#running-tests--benchmarks)
- [Adding a New Provider](#adding-a-new-provider)
- [Submitting Pull Requests](#submitting-pull-requests)

---

## Prerequisites

- **Go**: Version **1.27.0 or later** (ensure `CGO_ENABLED=0` capability).
- **Nix** (recommended): The repository provides a hermetic development shell via `flake.nix`.
- **direnv** (optional): For automatic environment loading (`direnv allow`).
- **golangci-lint**: Configured via `.golangci.yml`.
- **Git**: Configured for Conventional Commits.

---

## Quick Start

### Option A: Using Nix & direnv (Recommended)

The Nix Flake provisions Go 1.27, `golangci-lint`, `gopls`, `goreleaser`, `delve`, and pre-configured shell completion:

```bash
# Clone the repository
git clone https://github.com/julien-noblet/download-geofabrik.git
cd download-geofabrik

# If using direnv
direnv allow

# Or enter the devShell directly
nix develop
```

### Option B: Using Standard Go Toolchain

```bash
# Clone the repository
git clone https://github.com/julien-noblet/download-geofabrik.git
cd download-geofabrik

# Verify Go version
go version # must be >= 1.27.0

# Download dependencies
go mod download
```

### Build & Smoke Test

```bash
# Build binary
go build -o download-geofabrik ./cmd/download-geofabrik

# Inspect CLI help
./download-geofabrik --help

# Test a download in dry-run mode (ALWAYS use -n to avoid downloading massive files!)
./download-geofabrik download monaco -n
```

---

## Development Workflow

1. **Fork & Branch**: Create a feature or fix branch from `master` / `main`:
   ```bash
   git checkout -b feat/my-new-feature
   ```
2. **Make Changes**: Follow package boundaries and coding conventions.
3. **Verify Locally**:
   - Run tests: `go test ./...`
   - Run linter: `golangci-lint run`
   - Verify flake: `nix flake check`
4. **Commit**: Use [Conventional Commits](https://www.conventionalcommits.org/):
   ```text
   feat(provider): add support for osm-chile extracts
   fix(downloader): handle http 429 rate-limiting with exponential backoff
   test(catalog): add table-driven tests for MergeElement
   ```
5. **Push & Open PR**: Push to your fork and submit a Pull Request.

---

## Coding Guidelines & Linting

We enforce strict Go idioms and linter rules defined in `.golangci.yml`:

### Mandatory Conventions

- **Whitespace (`wsl_v5`)**: Leave empty lines before control flow (`if`, `for`, `switch`), declarations (`var`), and `return` statements unless tightly coupled.
- **Variable Naming (`varnamelen`)**: Avoid single-letter loop/variable names like `f`, `k`, `v` in non-trivial scopes. Use descriptive names like `format`, `elementID`, `providerName`.
- **String Constants (`goconst`)**: Any string literal repeated 3 or more times must be extracted to a `const`.
- **Formatting & Imports (`gci`, `gofumpt`)**:
  Import blocks must be grouped into three distinct sections:
  1. Standard library packages
  2. Third-party dependencies
  3. Local module packages (`github.com/julien-noblet/download-geofabrik/...`)
- **Error Wrapping**: Always wrap errors with `%w` when adding context: `fmt.Errorf("fetching catalog: %w", err)`. Check sentinel errors with `errors.Is` and types with `errors.As`.
- **Nolint Directives (`nolintlint`)**: Any `//nolint:...` suppression must include an explicit explanation comment explaining why it is necessary.

### Running the Linter

```bash
# Run golangci-lint
golangci-lint run

# Or via Nix
nix develop --command golangci-lint run
```

---

## Running Tests & Benchmarks

### Unit Tests

We use Go standard `testing` and `testify/assert` / `testify/require`. All tests must be table-driven where appropriate.

```bash
# Fast check for targeted packages
go test -v ./internal/cli/... ./internal/provider/osmtw/...

# Full test suite with race detection
go test -race ./...
```

### Fuzz Testing

Native Go fuzz tests (`testing.F`) are implemented across `internal/mcp`, `internal/provider/*`, and `pkg/catalog`.

```bash
# Seed corpus is verified automatically in standard test passes:
go test ./...

# Run active mutation fuzzing on a specific target:
go test -fuzz=FuzzProviderScraper -fuzztime=15s ./internal/provider/osmch
```

### Benchmarks

- **Target only modified packages** when benchmarking to avoid huge execution times and log dumps:
  ```bash
  go test -bench=. -benchmem ./internal/cli
  ```
- **Do not run `./...` benchmarks** across the whole repository.

---

## Adding a New Provider

All catalog providers implement the `Provider` interface (`internal/provider/provider.go`):

```go
type Provider interface {
    Name() string
    Description() string
    DefaultConfigFile() string
    FetchCatalog(ctx context.Context) (*catalog.Catalog, error)
}
```

Follow this step-by-step checklist to integrate a new provider:

### Step 1: Create the Provider Package
Create a new package under `internal/provider/<name>/` (e.g. `internal/provider/osmcl/`):
- `scraper.go` (or `api.go` if the provider has a JSON/GeoJSON endpoint).
- Implement `Name()`, `Description()`, `DefaultConfigFile()`, and `FetchCatalog(ctx)`.
- Define `DefaultFormats() catalog.FormatDefinitions` (ensure `.md5` format is mapped if supported).

### Step 2: Implement Scraper Tests
Create `scraper_test.go` in `internal/provider/<name>/`:
- Use `net/http/httptest` to mock remote server responses.
- Test successful parsing, network timeouts, invalid HTML/JSON, and empty inputs.
- Provider code should target 100% statement and branch test coverage.

### Step 3: Register in Defaults
Update `internal/provider/defaults.go`:
- In `RegisterDefaultProviders()`: Add `Register(<name>.NewProvider())`.
- In `AllDefaultFormats()`: Append `<name>.DefaultFormats()`.

### Step 4: Add Service Constant
In `internal/generator/generator.go`:
- Define `Service<ProviderName> = "<provider-key>"`.

### Step 5: Update CLI Help
In `internal/cli/root.go`:
- Append the new provider key to the `--service` flag help description.

### Step 6: Generate Catalog YAML
Run the catalog generator to scrape and generate the initial catalog file:
```bash
./download-geofabrik --service="<provider-key>" generate
```
This generates `<provider-key>.yml` at the repository root.

### Step 7: Update Automation & Build Files
- Update `.pre-commit`: Add the catalog generation command, catalog doc markdown command, and `git add <provider-key>.yml`.
- Update `.github/workflows/genyml.yml`: Add generation step and catalog doc markdown generation.
- Update `.goreleaser.yaml`: Add `./<provider-key>.yml` to `extra_files` under `dockers_v2`.

> [!IMPORTANT]
> **Never modify `README.md` directly!**
> `README.md` is automatically assembled by concatenating `.README.md1`, the CLI help output, and `.README.md2`. Run `.pre-commit` or follow the steps in `.github/workflows/genyml.yml` to regenerate it.

---

## Submitting Pull Requests

Before opening your pull request, verify the following checklist:

- [ ] Code builds without errors: `go build ./cmd/download-geofabrik`
- [ ] Targeted tests pass: `go test -race ./internal/...`
- [ ] Linters pass with zero warnings: `golangci-lint run`
- [ ] Flake checks pass: `nix flake check`
- [ ] Downloads tested in dry-run mode: `./download-geofabrik download <element> -n`
- [ ] Never edit `README.md` manually; regenerate via pipeline if CLI flags changed.
- [ ] Commit messages adhere to Conventional Commits.

Thank you for helping make `download-geofabrik` better!
