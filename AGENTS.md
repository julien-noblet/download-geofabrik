# AGENTS.md

## Project overview

This repository is a Go CLI for downloading OpenStreetMap extracts from multiple providers. The binary entrypoint lives in `cmd/download-geofabrik`, while the actual logic is organized under `internal/` and `pkg/`.

The most important runtime modules are:

- `internal/cli`: Cobra commands and CLI setup (`download`, `generate`, `list`, `mcp`, root flags)
- `internal/config`: default config/service naming and legacy config structures
- `internal/downloader`: HTTP download client, connection pooling, in-flight MD5 hashing, atomic writes
- `internal/generator`: config generation logic orchestrating providers
- `internal/mcp`: Model Context Protocol tools and resources served via stdio
- `internal/provider`: provider registry and provider-specific scrapers / API clients (10 providers)
- `internal/ui`: Table and JSON formatting for CLI output
- `pkg/catalog`: Thread-safe OSM catalog domain model, element/format definitions, YAML serialization
- `pkg/formats`: Shared format definitions, flag mappings, mini-format abbreviations

## Codebase Map & Routing Guide (Carte du Codebase)

Use this map to jump directly to the relevant file without scanning directories or wasting context tokens.

### 1. Package & File Index

| File / Package | Role & Key Symbols |
|---|---|
| `cmd/download-geofabrik/main.go` | Entrypoint binary. Injects `cli.Version`, calls `cli.Execute()`. |
| `internal/cli/root.go` | Root Cobra command, global flags (`--config/-c`, `--service/-s`, `--verbose`, `--quiet`), Viper bindings, `initConfig()`, `setupLogging()`. |
| `internal/cli/download.go` | `download [element]` command. Format flags (`-P`, `-H`, `-G`, `-B`, `-S`, `-p`, `-k`, `-g`, `-O`, etc.), format resolution (`selectDefaultFormat`, `resolveFormat`), execution (`runDownload`, `processDownload`). |
| `internal/cli/generate.go` | `generate` command (`-p/--progress`). Calls `generator.GenerateContext()`. |
| `internal/cli/list.go` | `list` command (`--markdown`, `--json`). Loads catalog, delegates to `internal/ui`. |
| `internal/cli/mcp.go` | `mcp` command. Starts stdio Model Context Protocol server. |
| `internal/config/config.go` | YAML schema (`Config`, `Options`), element searching (`FindElem`, `Exist` with hyphen/underscore fallback & dynamic Czech date `YYYY-MM-DD` resolution), URL resolution (`Elem2preURL`, `Elem2URL`), hashability check (`IsHashable`). |
| `internal/downloader/download.go` | `Downloader` struct, connection-pooled HTTP client (`DisableCompression: true` for pre-compressed OSM), atomic file download (`saveToFile` with `.tmp`), in-flight MD5 calculation via `io.MultiWriter`, progress bar via `pb/v3`, checksum control (`Checksum`, `verifyChecksum`). |
| `internal/downloader/hash.go` | Checksum verification: `ComputeMD5Hash`, `CheckFileHash`, `VerifyFileChecksum`. |
| `internal/downloader/pool.go` | `sync.Pool` 128KB byte buffers (`getBuffer`, `putBuffer`) for zero-alloc streaming I/O. |
| `internal/element/element.go` | Legacy element model (`Element`, `Formats`, `MapElement`). |
| `internal/generator/generator.go` | Catalog generator driver (`PerformGenerateContext`). Resolves provider via `provider.Get()`, fetches catalog, sorts formats, writes YAML. |
| `internal/generator/importer/geofabrik/geofabrik.go` | Legacy Geofabrik JSON index importer (`GetIndex`, `Convert`). |
| `internal/lists/lists.go` | Legacy table formatting (`ListAllRegions`, `CreateTable`, `GetSortedKeys`). |
| `internal/ui/printer.go` | CLI visual output: `PrintTable` (ASCII / Markdown via `tablewriter`) and `PrintJSON`. |
| `internal/mcp/server.go` | MCP server setup (`Server`, `NewServer`, `ServeStdio`) using `mark3labs/mcp-go`. |
| `internal/mcp/tools.go` | MCP tool handlers: `list_services`, `regenerate_catalog`, `list_elements`, `get_element`, `list_formats`, `download_element`. |
| `internal/mcp/resources.go` | MCP resource handlers: `geofabrik://services`, `geofabrik://formats`, `geofabrik://catalog/{service}`. |
| `internal/provider/provider.go` | `Provider` interface (`Name`, `Description`, `DefaultConfigFile`, `FetchCatalog`), thread-safe provider registry (`Register`, `Get`, `List`). |
| `internal/provider/defaults.go` | `RegisterDefaultProviders()`: registers all 10 built-in providers. |
| `pkg/catalog/catalog.go` | Core thread-safe domain model `Catalog` (`LoadFile`, `SaveFile`, `Find`, `MergeElement`, `SortedKeys`, `Elem2URL`). |
| `pkg/catalog/element.go` | Catalog element model (`Element`: `ID`, `File`, `Name`, `Parent`, `Formats`, `Meta`). |
| `pkg/catalog/format.go` | Format definitions & constants (`FormatOsmPbf`, `KeyOsmPbf`, etc.), `GetMiniFormats()`, `GetFormats()`. |
| `pkg/formats/formats.go` | Shared format constants, CLI flag keys, abbreviation map (`miniFormatMap`), `GetFormats()`. |

### 2. Providers Matrix

All providers implement `provider.Provider` (`internal/provider/provider.go`) and register in `internal/provider/defaults.go`:

| Service Key | Implementation File | Default Config | Type | Target Source / URL |
|---|---|---|---|---|
| `geofabrik` (default) | `internal/provider/geofabrik/api.go` | `geofabrik.yml` | JSON API | `https://download.geofabrik.de/index-v1-nogeom.json` |
| `openstreetmap.fr` | `internal/provider/openstreetmapfr/scraper.go` | `openstreetmap.fr.yml` | HTML Scraper | `https://download.openstreetmap.fr/extracts/` |
| `bbbike` | `internal/provider/bbbike/scraper.go` | `bbbike.yml` | HTML Scraper | `https://download.bbbike.org/osm/bbbike/` |
| `geo2day` | `internal/provider/geo2day/scraper.go` | `geo2day.yml` | HTML Scraper | `https://geo2day.com/` |
| `movisda` | `internal/provider/movisda/api.go` | `movisda.yml` | GeoJSON API | `https://osm.download.movisda.io/admin/Admin-latest.geojson` |
| `planet.osm.ch` | `internal/provider/osmch/scraper.go` | `planet.osm.ch.yml` | HTML Scraper | `https://planet.osm.ch/` |
| `osm.kewl.lu` | `internal/provider/osmkewllu/scraper.go` | `osm.kewl.lu.yml` | HTML Scraper | `https://osm.kewl.lu/luxembourg.osm/` |
| `osm.fit.vutbr.cz` | `internal/provider/osmfitvutbr/scraper.go` | `osm.fit.vutbr.cz.yml` | HTML Scraper | `https://osm.fit.vutbr.cz/extracts/` |
| `osmit-estratti` | `internal/provider/osmit/scraper.go` | `osmit-estratti.yml` | HTML Scraper | `https://osmit-estratti.wmcloud.org/` |
| `osm.kcwu.csie.org` | `internal/provider/osmtw/scraper.go` | `osm.kcwu.csie.org.yml` | HTML Scraper | `https://osm.kcwu.csie.org/download/tw-extract/` |

### 3. Task-Oriented Routing (Where to go directly)

| Goal / Issue | Primary Files to Inspect / Modify | Secondary Files / Context |
|---|---|---|
| **Download failure / 404 / format error** | `internal/cli/download.go`<br>`internal/downloader/download.go` | `<service>.yml` (inspect format list)<br>`pkg/catalog/catalog.go` (`Elem2URL`) |
| **Checksum / MD5 verification issue** | `internal/downloader/hash.go`<br>`internal/downloader/download.go` (`Checksum`, `verifyChecksum`) | `<service>.yml` (check `.md5` format presence) |
| **Provider catalog scraping bug / outdated URLs** | `internal/provider/<service>/scraper.go` (or `api.go`) | `internal/provider/<service>/scraper_test.go`<br>`<service>.yml` |
| **Add or update a file format** | 1. `pkg/formats/formats.go`<br>2. `pkg/catalog/format.go`<br>3. `internal/cli/download.go`<br>4. `internal/provider/<service>/scraper.go` (`DefaultFormats()`) | `<service>.yml`<br>`scraper_test.go` |
| **Add a new provider** | 1. `internal/provider/<newpkg>/scraper.go`<br>2. `internal/provider/defaults.go`<br>3. `internal/generator/generator.go` (add constant)<br>4. `internal/cli/root.go` (service flag help) | `.github/workflows/genyml.yml`<br>`<newpkg>.yml` |
| **CLI command or flag modification** | `internal/cli/root.go`<br>`internal/cli/<command>.go` | `internal/cli/<command>_test.go` |
| **MCP tool or resource update** | `internal/mcp/tools.go`<br>`internal/mcp/resources.go` | `internal/mcp/server.go`<br>`internal/mcp/server_test.go` |
| **CLI listing or table display formatting** | `internal/ui/printer.go`<br>`internal/cli/list.go` | `internal/ui/printer_test.go` |

### 4. Data Flows & Execution Pipelines

- **Download pipeline**: `cmd/.../main.go` &rarr; `cli.Execute()` &rarr; `runDownload()` &rarr; `config.LoadConfig(cfgFile)` &rarr; `config.FindElem()` &rarr; `selectDefaultFormat()` / `resolveFormat()` &rarr; `downloaderInstance.DownloadFile()` (atomic `.tmp` + in-flight MD5) &rarr; `downloaderInstance.Checksum()`.
- **Generate pipeline**: `cmd/.../main.go` &rarr; `runGenerate()` &rarr; `generator.PerformGenerateContext()` &rarr; `provider.Get(service)` &rarr; `prov.FetchCatalog(ctx)` &rarr; sort formats &rarr; `cat.SaveFile(cfgFile)`.
- **List pipeline**: `cmd/.../main.go` &rarr; `runList()` &rarr; `catalog.LoadFile(cfgFile)` &rarr; `ui.PrintTable()` (ASCII/Markdown) or `ui.PrintJSON()`.
- **MCP server pipeline**: `cmd/.../main.go` &rarr; `runMCP()` &rarr; `mcp.NewServer(Version)` &rarr; register tools/resources &rarr; `srv.ServeStdio(ctx)`.

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
- **Fuzz testing**:
  - Native Go fuzz tests (`testing.F`) are implemented across `internal/mcp`, `internal/provider/*`, and `pkg/catalog` (`*_fuzz_test.go` and `*_fuzz_internal_test.go`).
  - Seed corpuses run automatically as part of normal test passes (`go test ./...`).
  - Active mutation fuzzing can be run on-demand: `go test -fuzz=FuzzName -fuzztime=10s ./internal/mcp`.
- **Test commenting conventions**:
  - Keep tests self-documenting via explicit subtest names (`t.Run("descriptive case name", ...)`) and descriptive table fields.
  - Avoid redundant comments that simply paraphrase the assertion or function call.
  - Preserve all `//nolint:...` directives and comments explaining non-obvious edge cases, algorithmic constraints, or magic numbers.
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
   - 1. Run linter: `golangci-lint run` (fix any issues immediately)
   - 1. Run targeted benchmarks on modified packages only: `go test -bench=. -benchmem ./internal/<modified_pkg>`

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
- Modify directly the README.md. You can find the way to generate README.md from .github/workflows/genyml.yml file. Adapt the workflow to generate the catalog files too.
