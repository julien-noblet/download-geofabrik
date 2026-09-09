# 🗺️ download-geofabrik

<p align="center">
  <strong>Fast, multi-provider OpenStreetMap data extract downloader & Model Context Protocol (MCP) server for LLMs.</strong>
</p>

<p align="center">
  <a href="https://github.com/julien-noblet/download-geofabrik/releases"><img src="https://img.shields.io/github/v/release/julien-noblet/download-geofabrik?style=flat-square" alt="GitHub release" /></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.27+-00ADD8?style=flat-square&logo=go" alt="Go Version" /></a>
  <a href="https://pkg.go.dev/github.com/julien-noblet/download-geofabrik"><img src="https://pkg.go.dev/badge/github.com/julien-noblet/download-geofabrik.svg?style=flat-square" alt="Go Reference" /></a>
  <a href="https://github.com/julien-noblet/download-geofabrik/actions/workflows/gotest.yml"><img src="https://img.shields.io/github/actions/workflow/status/julien-noblet/download-geofabrik/gotest.yml?branch=master&label=tests&style=flat-square" alt="Build Status" /></a>
  <a href="https://github.com/julien-noblet/download-geofabrik/actions/workflows/nix.yml"><img src="https://img.shields.io/github/actions/workflow/status/julien-noblet/download-geofabrik/nix.yml?branch=master&label=nix&style=flat-square" alt="Nix Status" /></a>
  <a href="https://gist.githubusercontent.com/julien-noblet/a509e15ea4734ca3e8e98f32ab5369c0/raw/7344619caf8ac5bce291793711071a9636536fce/coverage.json"><img src="https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/julien-noblet/a509e15ea4734ca3e8e98f32ab5369c0/raw/7344619caf8ac5bce291793711071a9636536fce/coverage.json&style=flat-square" alt="Go Coverage" /></a>
  <a href="https://goreportcard.com/report/github.com/julien-noblet/download-geofabrik"><img src="https://goreportcard.com/badge/github.com/julien-noblet/download-geofabrik?style=flat-square" alt="Go Report Card" /></a>
  <a href="https://github.com/julien-noblet/download-geofabrik/pkgs/container/download-geofabrik"><img src="https://img.shields.io/badge/GHCR-download--geofabrik-blue?style=flat-square&logo=github" alt="GitHub Container Registry" /></a>
  <a href="https://modelcontextprotocol.io"><img src="https://img.shields.io/badge/MCP-Enabled-purple?style=flat-square" alt="MCP Enabled" /></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MPL_2.0-blue.svg?style=flat-square" alt="License: MPL 2.0" /></a>
</p>

---

## Table of Contents

- [Features](#features)
- [Installation](#installation)
  - [Pre-built Binaries](#pre-built-binaries)
  - [Go Install](#go-install)
  - [Docker](#docker)
  - [Nix](#nix)
- [Quick Start](#quick-start)
  - [Basic Download](#basic-download)
  - [Download from Another Provider](#download-from-another-provider)
  - [Download Specific Formats](#download-specific-formats)
  - [Custom Output Directory & Options](#custom-output-directory--options)
  - [Exploring Elements](#exploring-elements)
- [CLI Reference](#cli-reference)
- [Supported Formats](#supported-formats)
- [Model Context Protocol (MCP) Mode](#model-context-protocol-mcp-mode)
  - [Configuration for MCP Clients](#configuration-for-mcp-clients)
  - [Available MCP Tools](#available-mcp-tools)
  - [Available MCP Resources](#available-mcp-resources)
- [Catalog Maintenance](#catalog-maintenance)
- [Supported Providers & Catalogs](#supported-providers--catalogs)
- [Contributing](#contributing)
- [Data Attribution](#data-attribution)
- [License](#license)

---

## Features

- 🌍 **10 OpenStreetMap Providers**: Download extracts from Geofabrik, OpenStreetMap France, BBBike, Geo2day, Movisda, Planet OSM Switzerland, OSM Luxembourg, FIT VUT Brno (Czech Republic), OpenStreetMap Italia, and OSM Taiwan.
- 📦 **15+ Data Formats**: Download `.osm.pbf`, `.osh.pbf` (history), `.shp.zip` (Shapefiles), `.gpkg` (GeoPackage), `.mbtiles`, `.geojson`, `.poly` (Osmosis boundaries), `.kml`, `.map` (Mapsforge), `.obf` (OsmAnd), `.o5m`/`.o5m.zst`, Garmin maps, CSV, and more.
- 🔒 **Automated Integrity & Smart Caching**: Built-in MD5 checksum verification (enabled by default). Automatically skips re-downloading files if the local copy matches the remote hash.
- 🤖 **Native Model Context Protocol (MCP)**: Run as a local stdio MCP server for LLMs and AI agents (Claude Desktop, Antigravity, Cursor, Cline, etc.) to query catalogs, inspect metadata, and download extracts autonomously.
- ⚡ **Lightweight & Portable**: Single static Go binary with minimal dependencies, responsive progress bar, and quiet/verbose output flags.
- 🔄 **Remote Catalog Scraping**: Built-in `generate` command to scrape and update local YAML catalog definitions.

---

## Installation

### Pre-built Binaries

Download pre-compiled binaries for Linux, macOS, and Windows from the [GitHub Releases page](https://github.com/julien-noblet/download-geofabrik/releases).

### Go Install

Requires **Go 1.27 or later**:

```bash
go install github.com/julien-noblet/download-geofabrik/cmd/download-geofabrik@latest
```

### Docker

```bash
docker run -it --rm -v "$PWD:/data" -w /data ghcr.io/julien-noblet/download-geofabrik:latest -c /geofabrik.yml download [element]
```

### Nix

Run directly with Nix Flakes:

```bash
nix run github:julien-noblet/download-geofabrik -- download [element]
```

Or install permanently to your Nix profile:

```bash
nix profile install github:julien-noblet/download-geofabrik
```

Or enter a development shell:

```bash
nix develop github:julien-noblet/download-geofabrik
```

---

## Quick Start

### Basic Download

Download the default `osm.pbf` extract from Geofabrik:

```bash
download-geofabrik download ile-de-france
```

### Download from Another Provider

Use the `--service` flag to target alternative OSM data sources:

```bash
# Download from OpenStreetMap France
download-geofabrik --service openstreetmap.fr download rhone_alpes

# Download from BBBike
download-geofabrik --service bbbike download Berlin

# Download from OSM Taiwan (automatically uses provider's default format, o5m)
download-geofabrik --service osm.kcwu.csie.org download taiwan
```

### Download Specific Formats

Select one or more output formats simultaneously using shorthand flags:

```bash
# Download Shapefile (-S), Osmosis Poly boundary (-p), and GeoJSON (-g)
download-geofabrik download -S -p -g ile-de-france

# Download GeoPackage format (-K)
download-geofabrik download -K monaco
```

### Custom Output Directory & Options

```bash
# Download to a specific destination folder
download-geofabrik download --output-dir /tmp/osm-extracts ile-de-france

# Skip checksum verification
download-geofabrik download --check=false ile-de-france

# Dry run (test URLs without downloading files)
download-geofabrik download --nodownload ile-de-france
```

### Exploring Elements

```bash
# Interactive / readable table
download-geofabrik list

# Output as JSON for scripting
download-geofabrik --service bbbike list --json

# Output as Markdown table
download-geofabrik --service openstreetmap.fr list --markdown
```

---

## CLI Reference

```text
download-geofabrik is a CLI tool for downloading OpenStreetMap data and extracts from multiple providers.

Usage:
  download-geofabrik [flags]
  download-geofabrik [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  download    Download element
  generate    Generate configuration file
  help        Help about any command
  list        Show elements available
  mcp         Start Model Context Protocol (MCP) server for LLM integration

Flags:
  -c, --config string    config file (default is geofabrik.yml)
  -h, --help             help for download-geofabrik
      --quiet            Quiet mode
  -s, --service string   Service to use (geofabrik, geofabrik-parse, openstreetmap.fr, geo2day, bbbike, movisda, planet.osm.ch, osm.kewl.lu, osm.fit.vutbr.cz, osmit-estratti, osm.kcwu.csie.org) (default "geofabrik")
      --verbose          Verbose mode
  -v, --version          version for download-geofabrik

Use "download-geofabrik [command] --help" for more information about a command.
```

---

## Supported Formats

The following flags can be passed to the `download` command to select specific file formats:

| Flag | Format Key | Extension | Description | Default |
| :--- | :--- | :--- | :--- | :---: |
| `-P` | `osm.pbf` | `.osm.pbf` | OpenStreetMap Protocolbuffer Binary Format | ✅ Yes |
| | `pbf` | `.pbf` | Standard Protocolbuffer (used by planet.osm.ch) | |
| `-H` | `osh.pbf` | `.osh.pbf` | OpenStreetMap History Protocolbuffer | |
| `-S` | `shp.zip` | `.shp.zip` | ESRI Shapefiles in ZIP archive | |
| `-K` | `gpkg` | `.gpkg` | OGC GeoPackage vector dataset | |
| `-g` | `geojson` | `.geojson` | GeoJSON boundary / extract | |
| `-p` | `poly` | `.poly` | Osmosis Polygon boundary | |
| `-k` | `kml` | `.kml` | Keyhole Markup Language boundary | |
| `-M` | `mbtiles` | `.mbtiles` | Vector / Raster Mapbox MBTiles | |
| `-m` | `map` | `.map` | Mapsforge vector map format | |
| `-G` | `osm.gz` | `.osm.gz` | Gzip-compressed OSM XML | |
| `-B` | `osm.bz2` | `.osm.bz2` | Bzip2-compressed OSM XML | |
| `-5` | `o5m` | `.o5m` | Packed OSM data format | |
| `-Z` | `o5m.zst` | `.o5m.zst` | Zstandard-compressed o5m | |
| `-O` | `garmin.osm` | `.zip` | Garmin GPS map | |
| `-r` | `garmin.onroad` | `.zip` | Garmin On-Road map | |
| `-t` | `garmin.ontrail`| `.zip` | Garmin On-Trail map | |
| `-o` | `garmin.opentopo`| `.zip` | Garmin OpenTopo map | |
| | `obf` | `.obf` | OsmAnd Binary Format | |
| | `state` | `state.txt`| Replication state timestamp | |
| `-C` | `csv` | `.csv` | City / POI list in CSV format | |

---

## Model Context Protocol (MCP) Mode

`download-geofabrik` includes a built-in **MCP server** allowing Large Language Models (LLMs) such as Claude Desktop, Antigravity, Cursor, Cline, and Roo Code to interact with OpenStreetMap catalogs.

```bash
download-geofabrik mcp
```

### Configuration for MCP Clients

Add `download-geofabrik` to your MCP client configuration using the absolute path to the binary (e.g. `/usr/local/bin/download-geofabrik` or `/home/user/go/bin/download-geofabrik`):

#### Claude Desktop (`claude_desktop_config.json`)
- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`
- Linux: `~/.config/Claude/claude_desktop_config.json`

```json
{
  "mcpServers": {
    "download-geofabrik": {
      "command": "/path/to/download-geofabrik",
      "args": ["mcp"]
    }
  }
}
```

#### Cursor (`.cursor/mcp.json` or Settings > Features > MCP)

```json
{
  "mcpServers": {
    "download-geofabrik": {
      "command": "/path/to/download-geofabrik",
      "args": ["mcp"]
    }
  }
}
```

#### Antigravity / Gemini CLI (`antigravity.json`)

```json
{
  "mcpServers": {
    "download-geofabrik": {
      "command": "download-geofabrik",
      "args": ["mcp"]
    }
  }
}
```

*Note: If local `.yml` catalog files are not present, the MCP server automatically fetches the catalog on the fly from upstream providers.*

### Available MCP Tools

- `list_services`: List all 10 supported OSM data providers (`geofabrik`, `bbbike`, `openstreetmap.fr`, `movisda`, `geo2day`, etc.) and their configuration status.
- `regenerate_catalog`: Trigger remote catalog scraping and update local configuration files.
- `list_elements`: Search and list regions, countries, and cities for any provider with available file formats (`osm.pbf`, `shp.zip`, `poly`, `mbtiles`, etc.), supporting search and pagination.
- `get_element`: Retrieve detailed metadata for a specific extract, including resolved download URLs and MD5 checksum availability.
- `list_formats`: List supported data formats and extensions across services.
- `download_element`: Download OSM extracts for specified formats with checksum verification and `dry_run` simulation support.

### Available MCP Resources

- `geofabrik://services`: JSON list of all available provider services.
- `geofabrik://formats`: JSON list of all supported file formats.
- `geofabrik://catalog/{service}`: Full catalog elements and formats for the given service.

---

## Catalog Maintenance

The `generate` command queries upstream provider directories and regenerates the local YAML database files:

```bash
# Update Geofabrik catalog (geofabrik.yml)
download-geofabrik generate

# Update a specific provider catalog
download-geofabrik --service openstreetmap.fr generate
download-geofabrik --service bbbike generate
download-geofabrik --service movisda generate
```

---

## Supported Providers & Catalogs

`download-geofabrik` supports 10 OpenStreetMap data providers. Browse the full extract lists for each provider using the links below:

| Provider | Service Key | Description | Elements | Formats | Catalog Link |
| :--- | :--- | :--- | :---: | :--- | :---: |
| **Geofabrik** | `geofabrik` | Continents, countries, and regional extracts (default) | 555 | `osm.pbf`, `shp.zip`, `osm.bz2`, `poly`, `kml`, `mbtiles` | [View Catalog](docs/catalogs/geofabrik.md) |
| **OpenStreetMap France** | `openstreetmap.fr` | High-detail extracts for French regions, departments, and world | 1,204 | `osm.pbf`, `shp.zip`, `poly`, `gpkg`, `geojson` | [View Catalog](docs/catalogs/openstreetmap.fr.md) |
| **BBBike** | `bbbike` | Extracts for 200+ major metropolitan cities worldwide | 238 | `osm.pbf`, `shp.zip`, `osm.gz`, `gpkg`, `map`, `mbtiles`, Garmin | [View Catalog](docs/catalogs/bbbike.md) |
| **Geo2day** | `geo2day` | Custom regional and metropolitan extracts | 1,022 | `osm.pbf`, `poly`, `geojson` | [View Catalog](docs/catalogs/geo2day.md) |
| **Movisda** | `movisda` | Worldwide country and administrative subdivisions | 3,291 | `osm.pbf`, `poly`, `geojson` | [View Catalog](docs/catalogs/movisda.md) |
| **Planet OSM Switzerland** | `planet.osm.ch` | Switzerland national and cantonal extracts | 3 | `osm.pbf`, `pbf`, `poly`, `obf`, Garmin | [View Catalog](docs/catalogs/planet.osm.ch.md) |
| **OSM Luxembourg** | `osm.kewl.lu` | Daily extracts for Luxembourg | 1 | `osm.pbf`, `osm.bz2` | [View Catalog](docs/catalogs/osm.kewl.lu.md) |
| **FIT VUT Brno** | `osm.fit.vutbr.cz` | Czech Republic extracts by date | 2 | `osm.pbf`, `osm.bz2`, `poly` | [View Catalog](docs/catalogs/osm.fit.vutbr.cz.md) |
| **OpenStreetMap Italia** | `osmit-estratti` | Italy national, regional, and provincial extracts | 128 | `osm.pbf`, `gpkg`, `poly` | [View Catalog](docs/catalogs/osmit-estratti.md) |
| **OSM Taiwan** | `osm.kcwu.csie.org` | Taiwan extracts and change history | 1 | `o5m`, `o5m.zst` | [View Catalog](docs/catalogs/osm.kcwu.csie.org.md) |

---

## Contributing

Contributions are welcome! Please review [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines on:
- Setting up the hermetic development shell with Nix (`nix develop`) or Go 1.27+
- Running unit and fuzz tests: `go test ./...`
- Enforcing strict linter rules: `golangci-lint run`
- Running targeted benchmarks on modified packages
- Adding new OSM catalog providers step-by-step

---

## Data Attribution

OpenStreetMap data is licensed under the [Open Data Commons Open Database License](https://opendatacommons.org/licenses/odbl/) (ODbL) by the OpenStreetMap Foundation (OSMF).
- © [OpenStreetMap contributors](https://www.openstreetmap.org/copyright)
- Extract servers provided by Geofabrik GmbH, OpenStreetMap France, BBBike.org, Geo2day, Movisda, SOSM Switzerland, OSM Luxembourg, FIT VUT Brno, Wikimedia Italia, and the OSM Taiwan community.

---

## License

This project is licensed under the [Mozilla Public License 2.0 (MPL-2.0)](LICENSE).
