# 🗺️ download-geofabrik

<p align="center">
  <strong>Fast, multi-provider OpenStreetMap data extract downloader & Model Context Protocol (MCP) server for LLMs.</strong>
</p>

<p align="center">
  <a href="https://github.com/julien-noblet/download-geofabrik/releases"><img src="https://img.shields.io/github/v/release/julien-noblet/download-geofabrik?style=flat-square" alt="GitHub release" /></a>
  <a href="https://github.com/julien-noblet/download-geofabrik/actions/workflows/gotest.yml"><img src="https://img.shields.io/github/actions/workflow/status/julien-noblet/download-geofabrik/gotest.yml?branch=master&label=tests&style=flat-square" alt="Build Status" /></a>
  <a href="https://gist.githubusercontent.com/julien-noblet/a509e15ea4734ca3e8e98f32ab5369c0/raw/7344619caf8ac5bce291793711071a9636536fce/coverage.json"><img src="https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/julien-noblet/a509e15ea4734ca3e8e98f32ab5369c0/raw/7344619caf8ac5bce291793711071a9636536fce/coverage.json&style=flat-square" alt="Go Coverage" /></a>
  <a href="https://goreportcard.com/report/github.com/julien-noblet/download-geofabrik"><img src="https://goreportcard.com/badge/github.com/julien-noblet/download-geofabrik?style=flat-square" alt="Go Report Card" /></a>
  <a href="https://hub.docker.com/r/juliennoblet/download-geofabrik"><img src="https://img.shields.io/docker/pulls/juliennoblet/download-geofabrik?style=flat-square" alt="Docker Pulls" /></a>
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
- [Supported Formats](#supported-formats)
- [Model Context Protocol (MCP) Mode](#model-context-protocol-mcp-mode)
  - [Configuration for MCP Clients](#configuration-for-mcp-clients)
  - [Available MCP Tools](#available-mcp-tools)
  - [Available MCP Resources](#available-mcp-resources)
- [Catalog Maintenance](#catalog-maintenance)
- [CLI Reference](#cli-reference)
- [Available Extracts by Provider](#available-extracts-by-provider)
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

```bash
go install github.com/julien-noblet/download-geofabrik/cmd/download-geofabrik@latest
```

### Docker

```bash
docker run -it --rm -v "$PWD:/data" -w /data juliennoblet/download-geofabrik:latest download [element]
```

### Nix

Run directly with Nix Flakes:

```bash
nix run github:julien-noblet/download-geofabrik -- download [element]
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
download-geofabrik --service openstreetmap.fr download rhone-alpes

# Download from BBBike
download-geofabrik --service bbbike download Berlin
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
download-geofabrik download --no-check ile-de-france

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

Add `download-geofabrik` to your MCP client configuration (e.g., `claude_desktop_config.json` or `antigravity.json`):

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

## License

This project is licensed under the [Mozilla Public License 2.0 (MPL-2.0)](LICENSE).

---

## Available Extracts by Provider

Below is the complete list of elements and available formats across all supported providers. Click any provider below to expand its catalog table.

<details>
<summary><strong>Geofabrik Extracts</strong> (default service)</summary>

| SHORT NAME                                  | IS IN                       | LONG NAME                                                             | FORMATS |
|---------------------------------------------|-----------------------------|-----------------------------------------------------------------------|---------|
| act                                         | Australia                   | Australian Capital Territory                                          | kHPpSs  |
| afghanistan                                 | Asia                        | Afghanistan                                                           | kHPpSs  |
| africa                                      |                             | Africa                                                                | kHPps   |
| albania                                     | Europe                      | Albania                                                               | kHPpSs  |
| alberta                                     | Canada                      | Alberta                                                               | kHPpSs  |
| algeria                                     | Africa                      | Algeria                                                               | kHPpSs  |
| alps                                        | Europe                      | Alps                                                                  | kHPps   |
| alsace                                      | France                      | Alsace                                                                | kHPpSs  |
| american-oceania                            | Australia and Oceania       | American Oceania                                                      | kHPpSs  |
| andalucia                                   | Spain                       | Andalucía                                                             | kHPpSs  |
| andorra                                     | Europe                      | Andorra                                                               | kHPpSs  |
| angola                                      | Africa                      | Angola                                                                | kHPpSs  |
| anhui                                       | China                       | Anhui                                                                 | kHPpSs  |
| antarctica                                  |                             | Antarctica                                                            | kHPpSs  |
| aquitaine                                   | France                      | Aquitaine                                                             | kHPpSs  |
| aragon                                      | Spain                       | Aragón                                                                | kHPpSs  |
| argentina                                   | South America               | Argentina                                                             | kHPpSs  |
| armenia                                     | Asia                        | Armenia                                                               | kHPpSs  |
| arnsberg-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Arnsberg                                             | kHPpSs  |
| ashmore-cartier                             | Australia                   | Ashmore and Cartier Islands                                           | kHPpSs  |
| asia                                        |                             | Asia                                                                  | kHPps   |
| asturias                                    | Spain                       | Asturias                                                              | kHPpSs  |
| australia                                   | Australia and Oceania       | Australia                                                             | kHPpSs  |
| australia-oceania                           |                             | Australia and Oceania                                                 | kHPps   |
| austria                                     | Europe                      | Austria                                                               | kHPpSs  |
| auvergne                                    | France                      | Auvergne                                                              | kHPpSs  |
| azerbaijan                                  | Asia                        | Azerbaijan                                                            | kHPpSs  |
| azores                                      | Europe                      | Azores                                                                | kHPpSs  |
| baden-wuerttemberg                          | Germany                     | Baden-Württemberg                                                     | kHPpSs  |
| bahamas                                     | Central America             | Bahamas                                                               | kHPpSs  |
| bangladesh                                  | Asia                        | Bangladesh                                                            | kHPpSs  |
| basse-normandie                             | France                      | Basse-Normandie                                                       | kHPpSs  |
| bayern                                      | Germany                     | Bayern                                                                | kHPpSs  |
| bedfordshire                                | England                     | Bedfordshire                                                          | kHPpSs  |
| beijing                                     | China                       | Beijing                                                               | kHPpSs  |
| belarus                                     | Europe                      | Belarus                                                               | kHPpSs  |
| belgium                                     | Europe                      | Belgium                                                               | kHPpSs  |
| belize                                      | Central America             | Belize                                                                | kHPpSs  |
| benin                                       | Africa                      | Benin                                                                 | kHPpSs  |
| berkshire                                   | England                     | Berkshire                                                             | kHPpSs  |
| berlin                                      | Germany                     | Berlin                                                                | kHPpSs  |
| bermuda                                     | United Kingdom              | Bermuda                                                               | kHPpSs  |
| bhutan                                      | Asia                        | Bhutan                                                                | kHPpSs  |
| bolivia                                     | South America               | Bolivia                                                               | kHPpSs  |
| bosnia-herzegovina                          | Europe                      | Bosnia-Herzegovina                                                    | kHPpSs  |
| botswana                                    | Africa                      | Botswana                                                              | kHPpSs  |
| bourgogne                                   | France                      | Bourgogne                                                             | kHPpSs  |
| brandenburg                                 | Germany                     | Brandenburg (mit Berlin)                                              | kHPpSs  |
| brazil                                      | South America               | Brazil                                                                | kHPps   |
| bremen                                      | Germany                     | Bremen                                                                | kHPpSs  |
| bretagne                                    | France                      | Bretagne                                                              | kHPpSs  |
| bristol                                     | England                     | Bristol                                                               | kHPpSs  |
| britain-and-ireland                         | Europe                      | Britain and Ireland                                                   | kHPps   |
| british-columbia                            | Canada                      | British Columbia                                                      | kHPpSs  |
| buckinghamshire                             | England                     | Buckinghamshire                                                       | kHPpSs  |
| bulgaria                                    | Europe                      | Bulgaria                                                              | kHPpSs  |
| burkina-faso                                | Africa                      | Burkina Faso                                                          | kHPpSs  |
| burundi                                     | Africa                      | Burundi                                                               | kHPpSs  |
| cambodia                                    | Asia                        | Cambodia                                                              | kHPpSs  |
| cambridgeshire                              | England                     | Cambridgeshire                                                        | kHPpSs  |
| cameroon                                    | Africa                      | Cameroon                                                              | kHPpSs  |
| canada                                      | North America               | Canada                                                                | kHPps   |
| canary-islands                              | Africa                      | Canary Islands                                                        | kHPpSs  |
| cantabria                                   | Spain                       | Cantabria                                                             | kHPpSs  |
| cape-verde                                  | Africa                      | Cape Verde                                                            | kHPpSs  |
| castilla-la-mancha                          | Spain                       | Castilla-La Mancha                                                    | kHPpSs  |
| castilla-y-leon                             | Spain                       | Castilla y León                                                       | kHPpSs  |
| cataluna                                    | Spain                       | Cataluña                                                              | kHPpSs  |
| central-african-republic                    | Africa                      | Central African Republic                                              | kHPpSs  |
| central-america                             |                             | Central America                                                       | kHPps   |
| central-fed-district                        | Russian Federation          | Central Federal District                                              | kHPpSs  |
| central-zone                                | India                       | Central Zone                                                          | kHPpSs  |
| centre                                      | France                      | Centre                                                                | kHPpSs  |
| centro                                      | Italy                       | Centro                                                                | kHPpSs  |
| centro-oeste                                | Brazil                      | Centro-Oeste                                                          | kHPpSs  |
| ceuta                                       | Spain                       | Ceuta                                                                 | kHPpSs  |
| chad                                        | Africa                      | Chad                                                                  | kHPpSs  |
| champagne-ardenne                           | France                      | Champagne Ardenne                                                     | kHPpSs  |
| cheshire                                    | England                     | Cheshire                                                              | kHPpSs  |
| chile                                       | South America               | Chile                                                                 | kHPpSs  |
| china                                       | Asia                        | China                                                                 | kHPpSs  |
| chongqing                                   | China                       | Chongqing                                                             | kHPpSs  |
| christmas-island                            | Australia                   | Christmas Island                                                      | kHPpSs  |
| chubu                                       | Japan                       | Chūbu region                                                          | kHPpSs  |
| chugoku                                     | Japan                       | Chūgoku region                                                        | kHPpSs  |
| cocos-islands                               | Australia                   | Cocos (Keeling) Islands                                               | kHPpSs  |
| colombia                                    | South America               | Colombia                                                              | kHPpSs  |
| comores                                     | Africa                      | Comores                                                               | kHPpSs  |
| congo-brazzaville                           | Africa                      | Congo (Republic/Brazzaville)                                          | kHPpSs  |
| congo-democratic-republic                   | Africa                      | Congo (Democratic Republic/Kinshasa)                                  | kHPpSs  |
| cook-islands                                | Australia and Oceania       | Cook Islands                                                          | kHPpSs  |
| coral-sea-islands                           | Australia                   | Coral Sea Islands                                                     | kHPpSs  |
| cornwall                                    | England                     | Cornwall                                                              | kHPpSs  |
| corse                                       | France                      | Corse                                                                 | kHPpSs  |
| costa-rica                                  | Central America             | Costa Rica                                                            | kHPpSs  |
| crimean-fed-district                        | Russian Federation          | Crimean Federal District                                              | kHPpSs  |
| croatia                                     | Europe                      | Croatia                                                               | kHPpSs  |
| cuba                                        | Central America             | Cuba                                                                  | kHPpSs  |
| cumbria                                     | England                     | Cumbria                                                               | kHPpSs  |
| cyprus                                      | Europe                      | Cyprus                                                                | kHPpSs  |
| czech-republic                              | Europe                      | Czech Republic                                                        | kHPpSs  |
| dach                                        | Europe                      | Germany, Austria, Switzerland                                         | kHPps   |
| denmark                                     | Europe                      | Denmark                                                               | kHPpSs  |
| derbyshire                                  | England                     | Derbyshire                                                            | kHPpSs  |
| detmold-regbez                              | Nordrhein-Westfalen         | Regierungsbezirk Detmold                                              | kHPpSs  |
| devon                                       | England                     | Devon                                                                 | kHPpSs  |
| djibouti                                    | Africa                      | Djibouti                                                              | kHPpSs  |
| dolnoslaskie                                | Poland                      | Województwo dolnośląskie<br />(Lower Silesian Voivodeship)            | kHPpSs  |
| dorset                                      | England                     | Dorset                                                                | kHPpSs  |
| drenthe                                     | Netherlands                 | Drenthe                                                               | kHPpSs  |
| duesseldorf-regbez                          | Nordrhein-Westfalen         | Regierungsbezirk Düsseldorf                                           | kHPpSs  |
| durham                                      | England                     | Durham                                                                | kHPpSs  |
| east-sussex                                 | England                     | East Sussex                                                           | kHPpSs  |
| east-timor                                  | Asia                        | East Timor                                                            | kHPpSs  |
| east-yorkshire-with-hull                    | England                     | East Yorkshire with Hull                                              | kHPpSs  |
| eastern-zone                                | India                       | Eastern Zone                                                          | kHPpSs  |
| ecuador                                     | South America               | Ecuador                                                               | kHPpSs  |
| egypt                                       | Africa                      | Egypt                                                                 | kHPpSs  |
| el-salvador                                 | Central America             | El Salvador                                                           | kHPpSs  |
| enfield                                     | Greater London              | Enfield                                                               | kHPpSs  |
| england                                     | United Kingdom              | England                                                               | kHPpSs  |
| equatorial-guinea                           | Africa                      | Equatorial Guinea                                                     | kHPpSs  |
| eritrea                                     | Africa                      | Eritrea                                                               | kHPpSs  |
| essex                                       | England                     | Essex                                                                 | kHPpSs  |
| estonia                                     | Europe                      | Estonia                                                               | kHPpSs  |
| ethiopia                                    | Africa                      | Ethiopia                                                              | kHPpSs  |
| europe                                      |                             | Europe                                                                | kHPps   |
| extremadura                                 | Spain                       | Extremadura                                                           | kHPpSs  |
| falklands                                   | United Kingdom              | Falkland Islands                                                      | kHPpSs  |
| far-eastern-fed-district                    | Russian Federation          | Far Eastern Federal District                                          | kHPpSs  |
| faroe-islands                               | Europe                      | Faroe Islands                                                         | kHPpSs  |
| fiji                                        | Australia and Oceania       | Fiji                                                                  | kHPpSs  |
| finland                                     | Europe                      | Finland                                                               | kHPpSs  |
| flevoland                                   | Netherlands                 | Flevoland                                                             | kHPpSs  |
| france                                      | Europe                      | France                                                                | kHPps   |
| franche-comte                               | France                      | Franche Comte                                                         | kHPpSs  |
| freiburg-regbez                             | Baden-Württemberg           | Regierungsbezirk Freiburg                                             | kHPpSs  |
| friesland                                   | Netherlands                 | Friesland                                                             | kHPpSs  |
| fujian                                      | China                       | Fujian                                                                | kHPpSs  |
| gabon                                       | Africa                      | Gabon                                                                 | kHPpSs  |
| galicia                                     | Spain                       | Galicia                                                               | kHPpSs  |
| gansu                                       | China                       | Gansu                                                                 | kHPpSs  |
| gcc-states                                  | Asia                        | GCC States                                                            | kHPpSs  |
| gelderland                                  | Netherlands                 | Gelderland                                                            | kHPpSs  |
| georgia                                     | Europe                      | Georgia                                                               | kHPpSs  |
| germany                                     | Europe                      | Germany                                                               | kHPps   |
| ghana                                       | Africa                      | Ghana                                                                 | kHPpSs  |
| gloucestershire                             | England                     | Gloucestershire                                                       | kHPpSs  |
| great-britain                               | Europe                      | Great Britain                                                         | kHPps   |
| greater-london                              | England                     | Greater London                                                        | kHPpSs  |
| greater-manchester                          | England                     | Greater Manchester                                                    | kHPpSs  |
| greece                                      | Europe                      | Greece                                                                | kHPpSs  |
| greenland                                   | North America               | Greenland                                                             | kHPpSs  |
| groningen                                   | Netherlands                 | Groningen                                                             | kHPpSs  |
| guadeloupe                                  | France                      | Guadeloupe                                                            | kHPpSs  |
| guangdong                                   | China                       | Guangdong (with Hong Kong and Macau)                                  | kHPpSs  |
| guangxi                                     | China                       | Guangxi                                                               | kHPpSs  |
| guatemala                                   | Central America             | Guatemala                                                             | kHPpSs  |
| guernsey-jersey                             | Europe                      | Guernsey and Jersey                                                   | kHPpSs  |
| guinea                                      | Africa                      | Guinea                                                                | kHPpSs  |
| guinea-bissau                               | Africa                      | Guinea-Bissau                                                         | kHPpSs  |
| guizhou                                     | China                       | Guizhou                                                               | kHPpSs  |
| guyana                                      | South America               | Guyana                                                                | kHPpSs  |
| guyane                                      | France                      | Guyane                                                                | kHPpSs  |
| hainan                                      | China                       | Hainan                                                                | kHPpSs  |
| haiti-and-domrep                            | Central America             | Haiti and Dominican Republic                                          | kHPpSs  |
| hamburg                                     | Germany                     | Hamburg                                                               | kHPpSs  |
| hampshire                                   | England                     | Hampshire                                                             | kHPpSs  |
| haute-normandie                             | France                      | Haute-Normandie                                                       | kHPpSs  |
| heard-mcdonald                              | Australia                   | Heard Island and McDonald Islands                                     | kHPpSs  |
| hebei                                       | China                       | Hebei (with Beijing and Tianjin)                                      | kHPpSs  |
| heilongjiang                                | China                       | Heilongjiang                                                          | kHPpSs  |
| henan                                       | China                       | Henan                                                                 | kHPpSs  |
| herefordshire                               | England                     | Herefordshire                                                         | kHPpSs  |
| hertfordshire                               | England                     | Hertfordshire                                                         | kHPpSs  |
| hessen                                      | Germany                     | Hessen                                                                | kHPpSs  |
| hokkaido                                    | Japan                       | Hokkaidō                                                              | kHPpSs  |
| honduras                                    | Central America             | Honduras                                                              | kHPpSs  |
| hong-kong                                   | China                       | Hong Kong                                                             | kHPpSs  |
| hubei                                       | China                       | Hubei                                                                 | kHPpSs  |
| hunan                                       | China                       | Hunan                                                                 | kHPpSs  |
| hungary                                     | Europe                      | Hungary                                                               | kHPpSs  |
| iceland                                     | Europe                      | Iceland                                                               | kHPpSs  |
| ile-de-clipperton                           | Australia and Oceania       | Île de Clipperton                                                     | kHPpSs  |
| ile-de-france                               | France                      | Ile-de-France                                                         | kHPpSs  |
| india                                       | Asia                        | India                                                                 | kHPpSs  |
| indonesia                                   | Asia                        | Indonesia (with East Timor)                                           | kHPpSs  |
| inner-mongolia                              | China                       | Inner Mongolia                                                        | kHPpSs  |
| interior-admreg                             | British Columbia            | Interior Administrative Region                                        | kHPpSs  |
| iran                                        | Asia                        | Iran                                                                  | kHPpSs  |
| iraq                                        | Asia                        | Iraq                                                                  | kHPpSs  |
| ireland-and-northern-ireland                | Europe                      | Ireland and Northern Ireland                                          | kHPpSs  |
| island-admreg                               | British Columbia            | Island Administrative Region                                          | kHPpSs  |
| islas-baleares                              | Spain                       | Islas Baleares                                                        | kHPpSs  |
| isle-of-man                                 | Europe                      | Isle of Man                                                           | kHPpSs  |
| isle-of-wight                               | England                     | Isle of Wight                                                         | kHPpSs  |
| isole                                       | Italy                       | Isole                                                                 | kHPpSs  |
| israel-and-palestine                        | Asia                        | Israel and Palestine                                                  | kHPpSs  |
| italy                                       | Europe                      | Italy                                                                 | kHPps   |
| ivory-coast                                 | Africa                      | Ivory Coast                                                           | kHPpSs  |
| jamaica                                     | Central America             | Jamaica                                                               | kHPpSs  |
| japan                                       | Asia                        | Japan                                                                 | kHPps   |
| java                                        | Indonesia (with East Timor) | Java                                                                  | kHPpSs  |
| jiangsu                                     | China                       | Jiangsu                                                               | kHPpSs  |
| jiangxi                                     | China                       | Jiangxi                                                               | kHPpSs  |
| jihocesky                                   | Czech Republic              | Jihočeský kraj                                                        | kHPpSs  |
| jihomoravsky                                | Czech Republic              | Jihomoravský kraj                                                     | kHPpSs  |
| jilin                                       | China                       | Jilin                                                                 | kHPpSs  |
| jordan                                      | Asia                        | Jordan                                                                | kHPpSs  |
| kalimantan                                  | Indonesia (with East Timor) | Kalimantan                                                            | kHPpSs  |
| kaliningrad                                 | Russian Federation          | Kaliningrad                                                           | kHPpSs  |
| kansai                                      | Japan                       | Kansai region (a.k.a. Kinki region)                                   | kHPpSs  |
| kanto                                       | Japan                       | Kantō region                                                          | kHPpSs  |
| karlovarsky                                 | Czech Republic              | Karlovarský kraj                                                      | kHPpSs  |
| karlsruhe-regbez                            | Baden-Württemberg           | Regierungsbezirk Karlsruhe                                            | kHPpSs  |
| kazakhstan                                  | Asia                        | Kazakhstan                                                            | kHPpSs  |
| kent                                        | England                     | Kent                                                                  | kHPpSs  |
| kenya                                       | Africa                      | Kenya                                                                 | kHPpSs  |
| kiribati                                    | Australia and Oceania       | Kiribati                                                              | kHPpSs  |
| kitikmeot                                   | Nunavut                     | Kitikmeot Region                                                      | kHPpSs  |
| kivalliq                                    | Nunavut                     | Kivalliq Region                                                       | kHPpSs  |
| koeln-regbez                                | Nordrhein-Westfalen         | Regierungsbezirk Köln                                                 | kHPpSs  |
| kootenay-admreg                             | British Columbia            | Kootenay Administrative Region                                        | kHPpSs  |
| kosovo                                      | Europe                      | Kosovo                                                                | kHPpSs  |
| kralovehradecky                             | Czech Republic              | Královéhradecký kraj                                                  | kHPpSs  |
| kujawsko-pomorskie                          | Poland                      | Województwo kujawsko-pomorskie<br />(Kuyavian-Pomeranian Voivodeship) | kHPpSs  |
| kyrgyzstan                                  | Asia                        | Kyrgyzstan                                                            | kHPpSs  |
| kyushu                                      | Japan                       | Kyūshū                                                                | kHPpSs  |
| la-rioja                                    | Spain                       | La Rioja                                                              | kHPpSs  |
| lancashire                                  | England                     | Lancashire                                                            | kHPpSs  |
| languedoc-roussillon                        | France                      | Languedoc-Roussillon                                                  | kHPpSs  |
| laos                                        | Asia                        | Laos                                                                  | kHPpSs  |
| latvia                                      | Europe                      | Latvia                                                                | kHPpSs  |
| lebanon                                     | Asia                        | Lebanon                                                               | kHPpSs  |
| leicestershire                              | England                     | Leicestershire                                                        | kHPpSs  |
| lesotho                                     | Africa                      | Lesotho                                                               | kHPpSs  |
| liaoning                                    | China                       | Liaoning                                                              | kHPpSs  |
| liberecky                                   | Czech Republic              | Liberecký kraj                                                        | kHPpSs  |
| liberia                                     | Africa                      | Liberia                                                               | kHPpSs  |
| libya                                       | Africa                      | Libya                                                                 | kHPpSs  |
| liechtenstein                               | Europe                      | Liechtenstein                                                         | kHPpSs  |
| limburg                                     | Netherlands                 | Limburg                                                               | kHPpSs  |
| limousin                                    | France                      | Limousin                                                              | kHPpSs  |
| lincolnshire                                | England                     | Lincolnshire                                                          | kHPpSs  |
| lithuania                                   | Europe                      | Lithuania                                                             | kHPpSs  |
| lodzkie                                     | Poland                      | Województwo łódzkie<br />(Łódź Voivodeship)                           | kHPpSs  |
| lorraine                                    | France                      | Lorraine                                                              | kHPpSs  |
| lubelskie                                   | Poland                      | Województwo lubelskie<br />(Lublin Voivodeship)                       | kHPpSs  |
| lubuskie                                    | Poland                      | Województwo lubuskie<br />(Lubusz Voivodeship)                        | kHPpSs  |
| luxembourg                                  | Europe                      | Luxembourg                                                            | kHPpSs  |
| macau                                       | China                       | Macau                                                                 | kHPpSs  |
| macedonia                                   | Europe                      | Macedonia                                                             | kHPpSs  |
| madagascar                                  | Africa                      | Madagascar                                                            | kHPpSs  |
| madrid                                      | Spain                       | Madrid                                                                | kHPpSs  |
| malawi                                      | Africa                      | Malawi                                                                | kHPpSs  |
| malaysia-singapore-brunei                   | Asia                        | Malaysia, Singapore, and Brunei                                       | kHPpSs  |
| maldives                                    | Asia                        | Maldives                                                              | kHPpSs  |
| mali                                        | Africa                      | Mali                                                                  | kHPpSs  |
| malopolskie                                 | Poland                      | Województwo małopolskie<br />(Lesser Poland Voivodeship)              | kHPpSs  |
| malta                                       | Europe                      | Malta                                                                 | kHPpSs  |
| maluku                                      | Indonesia (with East Timor) | Maluku                                                                | kHPpSs  |
| manitoba                                    | Canada                      | Manitoba                                                              | kHPpSs  |
| marshall-islands                            | Australia and Oceania       | Marshall Islands                                                      | kHPpSs  |
| martinique                                  | France                      | Martinique                                                            | kHPpSs  |
| mauritania                                  | Africa                      | Mauritania                                                            | kHPpSs  |
| mauritius                                   | Africa                      | Mauritius                                                             | kHPpSs  |
| mayotte                                     | France                      | Mayotte                                                               | kHPpSs  |
| mazowieckie                                 | Poland                      | Województwo mazowieckie<br />(Mazovian Voivodeship)                   | kHPpSs  |
| mecklenburg-vorpommern                      | Germany                     | Mecklenburg-Vorpommern                                                | kHPpSs  |
| melilla                                     | Spain                       | Melilla                                                               | kHPpSs  |
| merseyside                                  | England                     | Merseyside                                                            | kHPpSs  |
| mexico                                      | North America               | Mexico                                                                | kHPpSs  |
| micronesia                                  | Australia and Oceania       | Micronesia                                                            | kHPpSs  |
| midi-pyrenees                               | France                      | Midi-Pyrenees                                                         | kHPpSs  |
| mittelfranken                               | Bayern                      | Mittelfranken                                                         | kHPpSs  |
| moldova                                     | Europe                      | Moldova                                                               | kHPpSs  |
| monaco                                      | Europe                      | Monaco                                                                | kHPpSs  |
| mongolia                                    | Asia                        | Mongolia                                                              | kHPpSs  |
| montenegro                                  | Europe                      | Montenegro                                                            | kHPpSs  |
| moravskoslezky                              | Czech Republic              | Moravskoslezský kraj                                                  | kHPpSs  |
| morocco                                     | Africa                      | Morocco                                                               | kHPpSs  |
| mozambique                                  | Africa                      | Mozambique                                                            | kHPpSs  |
| muenster-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Münster                                              | kHPpSs  |
| murcia                                      | Spain                       | Murcia                                                                | kHPpSs  |
| myanmar                                     | Asia                        | Myanmar (a.k.a. Burma)                                                | kHPpSs  |
| namibia                                     | Africa                      | Namibia                                                               | kHPpSs  |
| nauru                                       | Australia and Oceania       | Nauru                                                                 | kHPpSs  |
| navarra                                     | Spain                       | Navarra                                                               | kHPpSs  |
| nepal                                       | Asia                        | Nepal                                                                 | kHPpSs  |
| netherlands                                 | Europe                      | Netherlands                                                           | kHPpSs  |
| new-brunswick                               | Canada                      | New Brunswick                                                         | kHPpSs  |
| new-caledonia                               | Australia and Oceania       | New Caledonia                                                         | kHPpSs  |
| new-south-wales                             | Australia                   | New South Wales (with ACT and JBT)                                    | kHPpSs  |
| new-zealand                                 | Australia and Oceania       | New Zealand                                                           | kHPpSs  |
| newfoundland-and-labrador                   | Canada                      | Newfoundland and Labrador                                             | kHPpSs  |
| nicaragua                                   | Central America             | Nicaragua                                                             | kHPpSs  |
| niederbayern                                | Bayern                      | Niederbayern                                                          | kHPpSs  |
| niedersachsen                               | Germany                     | Niedersachsen (mit Bremen)                                            | kHPpSs  |
| niger                                       | Africa                      | Niger                                                                 | kHPpSs  |
| nigeria                                     | Africa                      | Nigeria                                                               | kHPpSs  |
| ningxia                                     | China                       | Ningxia                                                               | kHPpSs  |
| niue                                        | Australia and Oceania       | Niue                                                                  | kHPpSs  |
| noord-brabant                               | Netherlands                 | Noord-Brabant                                                         | kHPpSs  |
| noord-holland                               | Netherlands                 | Noord-Holland                                                         | kHPpSs  |
| norcal                                      | us/california               | Northern California                                                   | kHPpSs  |
| nord-est                                    | Italy                       | Nord-Est                                                              | kHPpSs  |
| nord-norge                                  | Norway                      | Nord-Norge<br />(Northern Norway)                                     | kHPpSs  |
| nord-ovest                                  | Italy                       | Nord-Ovest                                                            | kHPpSs  |
| nord-pas-de-calais                          | France                      | Nord-Pas-de-Calais                                                    | kHPpSs  |
| nordeste                                    | Brazil                      | Nordeste                                                              | kHPpSs  |
| nordrhein-westfalen                         | Germany                     | Nordrhein-Westfalen                                                   | kHPpSs  |
| norfolk                                     | England                     | Norfolk                                                               | kHPpSs  |
| norfolk-island                              | Australia                   | Norfolk Island                                                        | kHPpSs  |
| norte                                       | Brazil                      | Norte                                                                 | kHPpSs  |
| north-admreg                                | British Columbia            | North Administrative Region                                           | kHPpSs  |
| north-america                               |                             | North America                                                         | kHPps   |
| north-caucasus-fed-district                 | Russian Federation          | North Caucasus Federal District                                       | kHPpSs  |
| north-eastern-zone                          | India                       | North-Eastern Zone                                                    | kHPpSs  |
| north-korea                                 | Asia                        | North Korea                                                           | kHPpSs  |
| north-yorkshire                             | England                     | North Yorkshire                                                       | kHPpSs  |
| northamptonshire                            | England                     | Northamptonshire                                                      | kHPpSs  |
| northern-territory                          | Australia                   | Northern Territory                                                    | kHPpSs  |
| northern-zone                               | India                       | Northern Zone                                                         | kHPpSs  |
| northumberland                              | England                     | Northumberland                                                        | kHPpSs  |
| northwest-territories                       | Canada                      | Northwest Territories                                                 | kHPpSs  |
| northwestern-fed-district                   | Russian Federation          | Northwestern Federal District                                         | kHPpSs  |
| norway                                      | Europe                      | Norway                                                                | kHPpSs  |
| nottinghamshire                             | England                     | Nottinghamshire                                                       | kHPpSs  |
| nova-scotia                                 | Canada                      | Nova Scotia                                                           | kHPpSs  |
| nunavut                                     | Canada                      | Nunavut                                                               | kHPpSs  |
| nusa-tenggara                               | Indonesia (with East Timor) | Nusa-Tenggara                                                         | kHPpSs  |
| oberbayern                                  | Bayern                      | Oberbayern                                                            | kHPpSs  |
| oberfranken                                 | Bayern                      | Oberfranken                                                           | kHPpSs  |
| oberpfalz                                   | Bayern                      | Oberpfalz                                                             | kHPpSs  |
| okanagan-admreg                             | British Columbia            | Okanagan Administrative Region                                        | kHPpSs  |
| olomoucky                                   | Czech Republic              | Olomoucký kraj                                                        | kHPpSs  |
| ontario                                     | Canada                      | Ontario                                                               | kHPpSs  |
| opolskie                                    | Poland                      | Województwo opolskie<br />(Opole Voivodeship)                         | kHPpSs  |
| ostlandet                                   | Norway                      | Østlandet<br />(Eastern Norway)                                       | kHPpSs  |
| overijssel                                  | Netherlands                 | Overijssel                                                            | kHPpSs  |
| oxfordshire                                 | England                     | Oxfordshire                                                           | kHPpSs  |
| pais-vasco                                  | Spain                       | País Vasco                                                            | kHPpSs  |
| pakistan                                    | Asia                        | Pakistan                                                              | kHPpSs  |
| palau                                       | Australia and Oceania       | Palau                                                                 | kHPpSs  |
| panama                                      | Central America             | Panama                                                                | kHPpSs  |
| papua                                       | Indonesia (with East Timor) | Papua                                                                 | kHPpSs  |
| papua-new-guinea                            | Australia and Oceania       | Papua New Guinea                                                      | kHPpSs  |
| paraguay                                    | South America               | Paraguay                                                              | kHPpSs  |
| pardubicky                                  | Czech Republic              | Pardubický kraj                                                       | kHPpSs  |
| pays-de-la-loire                            | France                      | Pays de la Loire                                                      | kHPpSs  |
| peru                                        | South America               | Peru                                                                  | kHPpSs  |
| philippines                                 | Asia                        | Philippines                                                           | kHPpSs  |
| picardie                                    | France                      | Picardie                                                              | kHPpSs  |
| pitcairn-islands                            | Australia and Oceania       | Pitcairn Islands                                                      | kHPpSs  |
| plzensky                                    | Czech Republic              | Plzeňský kraj                                                         | kHPpSs  |
| podkarpackie                                | Poland                      | Województwo podkarpackie<br />(Subcarpathian Voivodeship)             | kHPpSs  |
| podlaskie                                   | Poland                      | Województwo podlaskie<br />(Podlaskie Voivodeship)                    | kHPpSs  |
| poitou-charentes                            | France                      | Poitou-Charentes                                                      | kHPpSs  |
| poland                                      | Europe                      | Poland                                                                | kHPps   |
| polynesie-francaise                         | Australia and Oceania       | Polynésie française (French Polynesia)                                | kHPpSs  |
| pomorskie                                   | Poland                      | Województwo pomorskie<br />(Pomeranian Voivodeship)                   | kHPpSs  |
| portugal                                    | Europe                      | Portugal                                                              | kHPpSs  |
| praha                                       | Czech Republic              | Praha                                                                 | kHPpSs  |
| prince-edward-island                        | Canada                      | Prince Edward Island                                                  | kHPpSs  |
| provence-alpes-cote-d-azur                  | France                      | Provence Alpes-Cote-d'Azur                                            | kHPpSs  |
| qikiqtaaluk                                 | Nunavut                     | Qikiqtaaluk Region                                                    | kHPpSs  |
| qinghai                                     | China                       | Qinghai                                                               | kHPpSs  |
| quebec                                      | Canada                      | Quebec                                                                | kHPpSs  |
| queensland                                  | Australia                   | Queensland                                                            | kHPpSs  |
| reunion                                     | France                      | Reunion                                                               | kHPpSs  |
| rheinland-pfalz                             | Germany                     | Rheinland-Pfalz                                                       | kHPpSs  |
| rhone-alpes                                 | France                      | Rhone-Alpes                                                           | kHPpSs  |
| romania                                     | Europe                      | Romania                                                               | kHPpSs  |
| russia                                      |                             | Russian Federation                                                    | kHPps   |
| rutland                                     | England                     | Rutland                                                               | kHPpSs  |
| rwanda                                      | Africa                      | Rwanda                                                                | kHPpSs  |
| saarland                                    | Germany                     | Saarland                                                              | kHPpSs  |
| sachsen                                     | Germany                     | Sachsen                                                               | kHPpSs  |
| sachsen-anhalt                              | Germany                     | Sachsen-Anhalt                                                        | kHPpSs  |
| saint-helena-ascension-and-tristan-da-cunha | Africa                      | Saint Helena, Ascension, and Tristan da Cunha                         | kHPpSs  |
| samoa                                       | Australia and Oceania       | Samoa                                                                 | kHPpSs  |
| sao-tome-and-principe                       | Africa                      | Sao Tome and Principe                                                 | kHPpSs  |
| saskatchewan                                | Canada                      | Saskatchewan                                                          | kHPpSs  |
| schleswig-holstein                          | Germany                     | Schleswig-Holstein                                                    | kHPpSs  |
| schwaben                                    | Bayern                      | Schwaben                                                              | kHPpSs  |
| scotland                                    | United Kingdom              | Scotland                                                              | kHPpSs  |
| sea                                         | Asia                        | South-East Asia                                                       | kHPps   |
| senegal-and-gambia                          | Africa                      | Senegal and Gambia                                                    | kHPpSs  |
| serbia                                      | Europe                      | Serbia                                                                | kHPpSs  |
| seychelles                                  | Africa                      | Seychelles                                                            | kHPpSs  |
| shaanxi                                     | China                       | Shaanxi                                                               | kHPpSs  |
| shandong                                    | China                       | Shandong                                                              | kHPpSs  |
| shanghai                                    | China                       | Shanghai                                                              | kHPpSs  |
| shanxi                                      | China                       | Shanxi                                                                | kHPpSs  |
| shikoku                                     | Japan                       | Shikoku                                                               | kHPpSs  |
| shropshire                                  | England                     | Shropshire                                                            | kHPpSs  |
| siberian-fed-district                       | Russian Federation          | Siberian Federal District                                             | kHPpSs  |
| sichuan                                     | China                       | Sichuan                                                               | kHPpSs  |
| sierra-leone                                | Africa                      | Sierra Leone                                                          | kHPpSs  |
| slaskie                                     | Poland                      | Województwo śląskie<br />(Silesian Voivodeship)                       | kHPpSs  |
| slovakia                                    | Europe                      | Slovakia                                                              | kHPpSs  |
| slovenia                                    | Europe                      | Slovenia                                                              | kHPpSs  |
| socal                                       | us/california               | Southern California                                                   | kHPpSs  |
| solomon-islands                             | Australia and Oceania       | Solomon Islands                                                       | kHPpSs  |
| somalia                                     | Africa                      | Somalia                                                               | kHPpSs  |
| somerset                                    | England                     | Somerset                                                              | kHPpSs  |
| sorlandet                                   | Norway                      | Sørlandet<br />(Southern Norway)                                      | kHPpSs  |
| south-africa                                | Africa                      | South Africa                                                          | kHPpSs  |
| south-africa-and-lesotho                    | Africa                      | South Africa (includes Lesotho)                                       | kHPps   |
| south-america                               |                             | South America                                                         | kHPps   |
| south-australia                             | Australia                   | South Australia                                                       | kHPpSs  |
| south-fed-district                          | Russian Federation          | South Federal District                                                | kHPpSs  |
| south-korea                                 | Asia                        | South Korea                                                           | kHPpSs  |
| south-sudan                                 | Africa                      | South Sudan                                                           | kHPpSs  |
| south-yorkshire                             | England                     | South Yorkshire                                                       | kHPpSs  |
| southcoast-admreg                           | British Columbia            | South Coast Administrative Region                                     | kHPpSs  |
| southern-zone                               | India                       | Southern Zone                                                         | kHPpSs  |
| spain                                       | Europe                      | Spain                                                                 | kHPpSs  |
| sri-lanka                                   | Asia                        | Sri Lanka                                                             | kHPpSs  |
| staffordshire                               | England                     | Staffordshire                                                         | kHPpSs  |
| stredocesky                                 | Czech Republic              | Středočeský kraj (with Praha)                                         | kHPpSs  |
| stuttgart-regbez                            | Baden-Württemberg           | Regierungsbezirk Stuttgart                                            | kHPpSs  |
| sud                                         | Italy                       | Sud                                                                   | kHPpSs  |
| sudan                                       | Africa                      | Sudan                                                                 | kHPpSs  |
| sudeste                                     | Brazil                      | Sudeste                                                               | kHPpSs  |
| suffolk                                     | England                     | Suffolk                                                               | kHPpSs  |
| sul                                         | Brazil                      | Sul                                                                   | kHPpSs  |
| sulawesi                                    | Indonesia (with East Timor) | Sulawesi                                                              | kHPpSs  |
| sumatra                                     | Indonesia (with East Timor) | Sumatra                                                               | kHPpSs  |
| suriname                                    | South America               | Suriname                                                              | kHPpSs  |
| surrey                                      | England                     | Surrey                                                                | kHPpSs  |
| svalbard-janmayen                           | Norway                      | Svalbard and Jan Mayen                                                | kHPpSs  |
| swaziland                                   | Africa                      | Swaziland                                                             | kHPpSs  |
| sweden                                      | Europe                      | Sweden                                                                | kHPpSs  |
| swietokrzyskie                              | Poland                      | Województwo świętokrzyskie<br />(Świętokrzyskie Voivodeship)          | kHPpSs  |
| switzerland                                 | Europe                      | Switzerland                                                           | kHPpSs  |
| syria                                       | Asia                        | Syria                                                                 | kHPpSs  |
| taiwan                                      | Asia                        | Taiwan                                                                | kHPpSs  |
| tajikistan                                  | Asia                        | Tajikistan                                                            | kHPpSs  |
| tanzania                                    | Africa                      | Tanzania                                                              | kHPpSs  |
| tasmania                                    | Australia                   | Tasmania                                                              | kHPpSs  |
| thailand                                    | Asia                        | Thailand                                                              | kHPpSs  |
| thueringen                                  | Germany                     | Thüringen                                                             | kHPpSs  |
| tianjin                                     | China                       | Tianjin                                                               | kHPpSs  |
| tibet                                       | China                       | Tibet                                                                 | kHPpSs  |
| togo                                        | Africa                      | Togo                                                                  | kHPpSs  |
| tohoku                                      | Japan                       | Tōhoku region                                                         | kHPpSs  |
| tokelau                                     | Australia and Oceania       | Tokelau                                                               | kHPpSs  |
| tonga                                       | Australia and Oceania       | Tonga                                                                 | kHPpSs  |
| trondelag                                   | Norway                      | Trøndelag<br />(Central Norway)                                       | kHPpSs  |
| tuebingen-regbez                            | Baden-Württemberg           | Regierungsbezirk Tübingen                                             | kHPpSs  |
| tunisia                                     | Africa                      | Tunisia                                                               | kHPpSs  |
| turkey                                      | Europe                      | Turkey                                                                | kHPpSs  |
| turkmenistan                                | Asia                        | Turkmenistan                                                          | kHPpSs  |
| tuvalu                                      | Australia and Oceania       | Tuvalu                                                                | kHPpSs  |
| tyne-and-wear                               | England                     | Tyne and Wear                                                         | kHPpSs  |
| uganda                                      | Africa                      | Uganda                                                                | kHPpSs  |
| ukraine                                     | Europe                      | Ukraine (with Crimea)                                                 | kHPpSs  |
| united-kingdom                              | Europe                      | United Kingdom                                                        | kHPps   |
| unterfranken                                | Bayern                      | Unterfranken                                                          | kHPpSs  |
| ural-fed-district                           | Russian Federation          | Ural Federal District                                                 | kHPpSs  |
| uruguay                                     | South America               | Uruguay                                                               | kHPpSs  |
| us                                          | North America               | United States of America                                              | kHPps   |
| us-midwest                                  | North America               | US Midwest                                                            | kHPps   |
| us-northeast                                | North America               | US Northeast                                                          | kHPps   |
| us-pacific                                  | North America               | US Pacific                                                            | kHPps   |
| us-south                                    | North America               | US South                                                              | kHPps   |
| us-west                                     | North America               | US West                                                               | kHPps   |
| us/alabama                                  | North America               | us/alabama                                                            | kHPpSs  |
| us/alaska                                   | North America               | us/alaska                                                             | kHPpSs  |
| us/arizona                                  | North America               | us/arizona                                                            | kHPpSs  |
| us/arkansas                                 | North America               | us/arkansas                                                           | kHPpSs  |
| us/california                               | North America               | us/california                                                         | kHPps   |
| us/colorado                                 | North America               | us/colorado                                                           | kHPpSs  |
| us/connecticut                              | North America               | us/connecticut                                                        | kHPpSs  |
| us/delaware                                 | North America               | us/delaware                                                           | kHPpSs  |
| us/district-of-columbia                     | North America               | us/district-of-columbia                                               | kHPpSs  |
| us/florida                                  | North America               | us/florida                                                            | kHPpSs  |
| us/georgia                                  | North America               | Georgia                                                               | kHPpSs  |
| us/hawaii                                   | North America               | us/hawaii                                                             | kHPpSs  |
| us/idaho                                    | North America               | us/idaho                                                              | kHPpSs  |
| us/illinois                                 | North America               | us/illinois                                                           | kHPpSs  |
| us/indiana                                  | North America               | us/indiana                                                            | kHPpSs  |
| us/iowa                                     | North America               | us/iowa                                                               | kHPpSs  |
| us/kansas                                   | North America               | us/kansas                                                             | kHPpSs  |
| us/kentucky                                 | North America               | us/kentucky                                                           | kHPpSs  |
| us/louisiana                                | North America               | us/louisiana                                                          | kHPpSs  |
| us/maine                                    | North America               | us/maine                                                              | kHPpSs  |
| us/maryland                                 | North America               | us/maryland                                                           | kHPpSs  |
| us/massachusetts                            | North America               | us/massachusetts                                                      | kHPpSs  |
| us/michigan                                 | North America               | us/michigan                                                           | kHPpSs  |
| us/minnesota                                | North America               | us/minnesota                                                          | kHPpSs  |
| us/mississippi                              | North America               | us/mississippi                                                        | kHPpSs  |
| us/missouri                                 | North America               | us/missouri                                                           | kHPpSs  |
| us/montana                                  | North America               | us/montana                                                            | kHPpSs  |
| us/nebraska                                 | North America               | us/nebraska                                                           | kHPpSs  |
| us/nevada                                   | North America               | us/nevada                                                             | kHPpSs  |
| us/new-hampshire                            | North America               | us/new-hampshire                                                      | kHPpSs  |
| us/new-jersey                               | North America               | us/new-jersey                                                         | kHPpSs  |
| us/new-mexico                               | North America               | us/new-mexico                                                         | kHPpSs  |
| us/new-york                                 | North America               | us/new-york                                                           | kHPpSs  |
| us/north-carolina                           | North America               | us/north-carolina                                                     | kHPpSs  |
| us/north-dakota                             | North America               | us/north-dakota                                                       | kHPpSs  |
| us/ohio                                     | North America               | us/ohio                                                               | kHPpSs  |
| us/oklahoma                                 | North America               | us/oklahoma                                                           | kHPpSs  |
| us/oregon                                   | North America               | us/oregon                                                             | kHPpSs  |
| us/pennsylvania                             | North America               | us/pennsylvania                                                       | kHPpSs  |
| us/puerto-rico                              | North America               | us/puerto-rico                                                        | kHPpSs  |
| us/rhode-island                             | North America               | us/rhode-island                                                       | kHPpSs  |
| us/south-carolina                           | North America               | us/south-carolina                                                     | kHPpSs  |
| us/south-dakota                             | North America               | us/south-dakota                                                       | kHPpSs  |
| us/tennessee                                | North America               | us/tennessee                                                          | kHPpSs  |
| us/texas                                    | North America               | us/texas                                                              | kHPpSs  |
| us/us-virgin-islands                        | North America               | us/us-virgin-islands                                                  | kHPpSs  |
| us/utah                                     | North America               | us/utah                                                               | kHPpSs  |
| us/vermont                                  | North America               | us/vermont                                                            | kHPpSs  |
| us/virginia                                 | North America               | us/virginia                                                           | kHPpSs  |
| us/washington                               | North America               | us/washington                                                         | kHPpSs  |
| us/west-virginia                            | North America               | us/west-virginia                                                      | kHPpSs  |
| us/wisconsin                                | North America               | us/wisconsin                                                          | kHPpSs  |
| us/wyoming                                  | North America               | us/wyoming                                                            | kHPpSs  |
| ustecky                                     | Czech Republic              | Ústecký kraj                                                          | kHPpSs  |
| utrecht                                     | Netherlands                 | Utrecht                                                               | kHPpSs  |
| uzbekistan                                  | Asia                        | Uzbekistan                                                            | kHPpSs  |
| valencia                                    | Spain                       | Valencia                                                              | kHPpSs  |
| vanuatu                                     | Australia and Oceania       | Vanuatu                                                               | kHPpSs  |
| venezuela                                   | South America               | Venezuela                                                             | kHPpSs  |
| vestlandet                                  | Norway                      | Vestlandet<br />(Western Norway)                                      | kHPpSs  |
| victoria                                    | Australia                   | Victoria                                                              | kHPpSs  |
| vietnam                                     | Asia                        | Vietnam                                                               | kHPpSs  |
| volga-fed-district                          | Russian Federation          | Volga Federal District                                                | kHPpSs  |
| vysocina                                    | Czech Republic              | Kraj Vysočina                                                         | kHPpSs  |
| wales                                       | United Kingdom              | Wales                                                                 | kHPpSs  |
| wallis-et-futuna                            | Australia and Oceania       | Wallis et Futuna                                                      | kHPpSs  |
| warminsko-mazurskie                         | Poland                      | Województwo warmińsko-mazurskie<br />(Warmian-Masurian Voivodeship)   | kHPpSs  |
| warwickshire                                | England                     | Warwickshire                                                          | kHPpSs  |
| west-midlands                               | England                     | West Midlands                                                         | kHPpSs  |
| west-sussex                                 | England                     | West Sussex                                                           | kHPpSs  |
| west-yorkshire                              | England                     | West Yorkshire                                                        | kHPpSs  |
| western-australia                           | Australia                   | Western Australia                                                     | kHPpSs  |
| western-zone                                | India                       | Western Zone                                                          | kHPpSs  |
| wielkopolskie                               | Poland                      | Województwo wielkopolskie<br />(Greater Poland Voivodeship)           | kHPpSs  |
| wiltshire                                   | England                     | Wiltshire                                                             | kHPpSs  |
| worcestershire                              | England                     | Worcestershire                                                        | kHPpSs  |
| xinjiang                                    | China                       | Xinjiang                                                              | kHPpSs  |
| yemen                                       | Asia                        | Yemen                                                                 | kHPpSs  |
| yukon                                       | Canada                      | Yukon                                                                 | kHPpSs  |
| yunnan                                      | China                       | Yunnan                                                                | kHPpSs  |
| zachodniopomorskie                          | Poland                      | Województwo zachodniopomorskie<br />(West Pomeranian Voivodeship)     | kHPpSs  |
| zambia                                      | Africa                      | Zambia                                                                | kHPpSs  |
| zeeland                                     | Netherlands                 | Zeeland                                                               | kHPpSs  |
| zhejiang                                    | China                       | Zhejiang                                                              | kHPpSs  |
| zimbabwe                                    | Africa                      | Zimbabwe                                                              | kHPpSs  |
| zlinsky                                     | Czech Republic              | Zlínský kraj                                                          | kHPpSs  |
| zuid-holland                                | Netherlands                 | Zuid-Holland                                                          | kHPpSs  |

</details>

<details>
<summary><strong>OpenStreetMap France Extracts (openstreetmap.fr)</strong></summary>

| SHORT NAME                               | IS IN                            | LONG NAME                                | FORMATS |
|------------------------------------------|----------------------------------|------------------------------------------|---------|
| aargau                                   | switzerland                      | aargau                                   | Ps      |
| abitibi_temiscamingue                    | quebec                           | abitibi_temiscamingue                    | Ps      |
| abruzzo                                  | italy                            | abruzzo                                  | Ps      |
| aceh                                     | indonesia                        | aceh                                     | Ps      |
| acre                                     | north                            | acre                                     | Ps      |
| adygea_republic                          | southern_federal_district        | adygea_republic                          | Ps      |
| aegean                                   | turkey                           | aegean                                   | Ps      |
| afghanistan                              | asia                             | afghanistan                              | Ps      |
| africa                                   |                                  | africa                                   | Ps      |
| africa_france_taaf                       | africa                           | africa_france_taaf                       | Ps      |
| aguascalientes                           | mexico                           | aguascalientes                           | Ps      |
| ain                                      | rhone_alpes                      | ain                                      | Ps      |
| aisne                                    | picardie                         | aisne                                    | Ps      |
| akershus                                 | norway                           | akershus                                 | Ps      |
| alagoas                                  | northeast                        | alagoas                                  | Ps      |
| alameda                                  | california                       | alameda                                  | Ps      |
| aland                                    | finland                          | aland                                    | Ps      |
| alava                                    | euskadi                          | alava                                    | Ps      |
| albacete                                 | castilla_la_mancha               | albacete                                 | Ps      |
| alberta                                  | canada                           | alberta                                  | Ps      |
| algeria                                  | africa                           | algeria                                  | Ps      |
| alicante                                 | comunitat_valenciana             | alicante                                 | Ps      |
| allier                                   | auvergne                         | allier                                   | Ps      |
| almeria                                  | andalucia                        | almeria                                  | Ps      |
| alpes_de_haute_provence                  | provence_alpes_cote_d_azur       | alpes_de_haute_provence                  | Ps      |
| alpes_maritimes                          | provence_alpes_cote_d_azur       | alpes_maritimes                          | Ps      |
| alpine                                   | california                       | alpine                                   | Ps      |
| alsace                                   | france                           | alsace                                   | Ps      |
| altai_krai                               | siberian_federal_district        | altai_krai                               | Ps      |
| altai_republic                           | siberian_federal_district        | altai_republic                           | Ps      |
| amador                                   | california                       | amador                                   | Ps      |
| amapa                                    | north                            | amapa                                    | Ps      |
| amazonas                                 | north                            | amazonas                                 | Ps      |
| american_samoa                           | oceania                          | american_samoa                           | Ps      |
| amur_oblast                              | far_eastern_federal_district     | amur_oblast                              | Ps      |
| andalucia                                | spain                            | andalucia                                | Ps      |
| andaman_and_nicobar_islands              | india                            | andaman_and_nicobar_islands              | Ps      |
| andhra_pradesh                           | india                            | andhra_pradesh                           | Ps      |
| andorra                                  | europe                           | andorra                                  | Ps      |
| angola                                   | africa                           | angola                                   | Ps      |
| anguilla                                 | central-america                  | anguilla                                 | Ps      |
| anhui                                    | china                            | anhui                                    | Ps      |
| antigua_and_barbuda                      | central-america                  | antigua_and_barbuda                      | Ps      |
| antofagasta                              | chile                            | antofagasta                              | Ps      |
| antwerp                                  | flanders                         | antwerp                                  | Ps      |
| appenzell_ausserrhoden                   | switzerland                      | appenzell_ausserrhoden                   | Ps      |
| appenzell_innerrhoden                    | switzerland                      | appenzell_innerrhoden                    | Ps      |
| appomattox                               | virginia                         | appomattox                               | Ps      |
| aquitaine                                | france                           | aquitaine                                | Ps      |
| aragon                                   | spain                            | aragon                                   | Ps      |
| araucania                                | chile                            | araucania                                | Ps      |
| ardeche                                  | rhone_alpes                      | ardeche                                  | Ps      |
| ardennes                                 | champagne_ardenne                | ardennes                                 | Ps      |
| argentina                                | south-america                    | argentina                                | Ps      |
| argentina_la_rioja                       | argentina                        | argentina_la_rioja                       | Ps      |
| argentina_santa_cruz                     | argentina                        | argentina_santa_cruz                     | Ps      |
| arica_y_parinacota                       | chile                            | arica_y_parinacota                       | Ps      |
| ariege                                   | midi_pyrenees                    | ariege                                   | Ps      |
| arkhangelsk_oblast                       | northwestern_federal_district    | arkhangelsk_oblast                       | Ps      |
| armenia                                  | asia                             | armenia                                  | Ps      |
| arnsberg                                 | nordrhein_westfalen              | arnsberg                                 | Ps      |
| aruba                                    | central-america                  | aruba                                    | Ps      |
| arunachal_pradesh                        | india                            | arunachal_pradesh                        | Ps      |
| ashmore_and_cartier_islands              | australia                        | ashmore_and_cartier_islands              | Ps      |
| asia                                     |                                  | asia                                     | Ps      |
| assam                                    | india                            | assam                                    | Ps      |
| astrakhan_oblast                         | southern_federal_district        | astrakhan_oblast                         | Ps      |
| asturias                                 | spain                            | asturias                                 | Ps      |
| atacama                                  | chile                            | atacama                                  | Ps      |
| aube                                     | champagne_ardenne                | aube                                     | Ps      |
| aude                                     | languedoc_roussillon             | aude                                     | Ps      |
| aust-agder                               | norway                           | aust-agder                               | Ps      |
| australia                                | oceania                          | australia                                | Ps      |
| australian_capital_territory             | australia                        | australian_capital_territory             | Ps      |
| austria                                  | europe                           | austria                                  | Ps      |
| auvergne                                 | france                           | auvergne                                 | Ps      |
| aveiro                                   | portugal                         | aveiro                                   | Ps      |
| aveyron                                  | midi_pyrenees                    | aveyron                                  | Ps      |
| avila                                    | castilla_y_leon                  | avila                                    | Ps      |
| aysen                                    | chile                            | aysen                                    | Ps      |
| azores                                   | portugal                         | azores                                   | Ps      |
| badajoz                                  | extremadura                      | badajoz                                  | Ps      |
| bahamas                                  | central-america                  | bahamas                                  | Ps      |
| bahia                                    | northeast                        | bahia                                    | Ps      |
| bahrain                                  | asia                             | bahrain                                  | Ps      |
| baja_california                          | mexico                           | baja_california                          | Ps      |
| baja_california_sur                      | mexico                           | baja_california_sur                      | Ps      |
| bali                                     | indonesia                        | bali                                     | Ps      |
| bangka_belitung_islands                  | indonesia                        | bangka_belitung_islands                  | Ps      |
| bangsamoro                               | philippines                      | bangsamoro                               | Ps      |
| banskobystricky                          | slovakia                         | banskobystricky                          | Ps      |
| banten                                   | indonesia                        | banten                                   | Ps      |
| barbados                                 | central-america                  | barbados                                 | Ps      |
| barcelona                                | catalunya                        | barcelona                                | Ps      |
| bas_rhin                                 | alsace                           | bas_rhin                                 | Ps      |
| bas_saint_laurent                        | quebec                           | bas_saint_laurent                        | Ps      |
| basel_landschaft                         | switzerland                      | basel_landschaft                         | Ps      |
| basel_stadt                              | switzerland                      | basel_stadt                              | Ps      |
| bashkortostan_republic                   | volga_federal_district           | bashkortostan_republic                   | Ps      |
| basilicata                               | italy                            | basilicata                               | Ps      |
| basse_normandie                          | france                           | basse_normandie                          | Ps      |
| beijing                                  | china                            | beijing                                  | Ps      |
| beja                                     | portugal                         | beja                                     | Ps      |
| belgium                                  | europe                           | belgium                                  | Ps      |
| belgorod_oblast                          | central_federal_district         | belgorod_oblast                          | Ps      |
| bengkulu                                 | indonesia                        | bengkulu                                 | Ps      |
| benin                                    | africa                           | benin                                    | Ps      |
| bermuda                                  | north-america                    | bermuda                                  | Ps      |
| bern                                     | switzerland                      | bern                                     | Ps      |
| bhutan                                   | asia                             | bhutan                                   | Ps      |
| bicol_region                             | philippines                      | bicol_region                             | Ps      |
| bihar                                    | india                            | bihar                                    | Ps      |
| biobio                                   | chile                            | biobio                                   | Ps      |
| bir_tawil                                | africa                           | bir_tawil                                | Ps      |
| black_sea                                | turkey                           | black_sea                                | Ps      |
| blekinge                                 | sweden                           | blekinge                                 | Ps      |
| bouches_du_rhone                         | provence_alpes_cote_d_azur       | bouches_du_rhone                         | Ps      |
| bourgogne                                | france                           | bourgogne                                | Ps      |
| bouvet_island                            | africa                           | bouvet_island                            | Ps      |
| braga                                    | portugal                         | braga                                    | Ps      |
| braganca                                 | portugal                         | braganca                                 | Ps      |
| bratislavsky                             | slovakia                         | bratislavsky                             | Ps      |
| brazil                                   | south-america                    | brazil                                   | Ps      |
| brazil_central-west                      | brazil                           | brazil_central-west                      | Ps      |
| brazil_north                             | brazil                           | brazil_north                             | Ps      |
| brazil_northeast                         | brazil                           | brazil_northeast                         | Ps      |
| brazil_south                             | brazil                           | brazil_south                             | Ps      |
| brazil_southeast                         | brazil                           | brazil_southeast                         | Ps      |
| bretagne                                 | france                           | bretagne                                 | Ps      |
| british_columbia                         | canada                           | british_columbia                         | Ps      |
| british_indian_ocean_territory           | asia                             | british_indian_ocean_territory           | Ps      |
| british_virgin_islands                   | central-america                  | british_virgin_islands                   | Ps      |
| brunei                                   | asia                             | brunei                                   | Ps      |
| brussels_capital_region                  | belgium                          | brussels_capital_region                  | Ps      |
| bryansk_oblast                           | central_federal_district         | bryansk_oblast                           | Ps      |
| buenos_aires                             | argentina                        | buenos_aires                             | Ps      |
| buenos_aires_city                        | argentina                        | buenos_aires_city                        | Ps      |
| burgenland                               | austria                          | burgenland                               | Ps      |
| burgos                                   | castilla_y_leon                  | burgos                                   | Ps      |
| burkina_faso                             | africa                           | burkina_faso                             | Ps      |
| burundi                                  | africa                           | burundi                                  | Ps      |
| buryatia_republic                        | siberian_federal_district        | buryatia_republic                        | Ps      |
| buskerud                                 | norway                           | buskerud                                 | Ps      |
| butte                                    | california                       | butte                                    | Ps      |
| cabo_delgado                             | mozambique                       | cabo_delgado                             | Ps      |
| caceres                                  | extremadura                      | caceres                                  | Ps      |
| cadiz                                    | andalucia                        | cadiz                                    | Ps      |
| cagayan_valley                           | philippines                      | cagayan_valley                           | Ps      |
| calabarzon                               | philippines                      | calabarzon                               | Ps      |
| calabria                                 | italy                            | calabria                                 | Ps      |
| calaveras                                | california                       | calaveras                                | Ps      |
| california                               | us-west                          | california                               | Ps      |
| california_lake                          | california                       | california_lake                          | Ps      |
| california_santa_cruz                    | california                       | california_santa_cruz                    | Ps      |
| calvados                                 | basse_normandie                  | calvados                                 | Ps      |
| cambodia                                 | asia                             | cambodia                                 | Ps      |
| cameroon                                 | africa                           | cameroon                                 | Ps      |
| campania                                 | italy                            | campania                                 | Ps      |
| campeche                                 | mexico                           | campeche                                 | Ps      |
| canada                                   | north-america                    | canada                                   | Ps      |
| canarias                                 | spain                            | canarias                                 | Ps      |
| cantabria                                | spain                            | cantabria                                | Ps      |
| cantal                                   | auvergne                         | cantal                                   | Ps      |
| cape_verde                               | africa                           | cape_verde                               | Ps      |
| capital_district                         | new-york                         | capital_district                         | Ps      |
| capitale_nationale                       | quebec                           | capitale_nationale                       | Ps      |
| caraga                                   | philippines                      | caraga                                   | Ps      |
| caribbean                                | central-america                  | caribbean                                | Ps      |
| castellon                                | comunitat_valenciana             | castellon                                | Ps      |
| castelo_branco                           | portugal                         | castelo_branco                           | Ps      |
| castilla_la_mancha                       | spain                            | castilla_la_mancha                       | Ps      |
| castilla_y_leon                          | spain                            | castilla_y_leon                          | Ps      |
| catalunya                                | spain                            | catalunya                                | Ps      |
| catamarca                                | argentina                        | catamarca                                | Ps      |
| cayman_islands                           | central-america                  | cayman_islands                           | Ps      |
| ceara                                    | northeast                        | ceara                                    | Ps      |
| central-america                          |                                  | central-america                          | Ps      |
| central-west                             | brazil                           | central-west                             |         |
| central_african_republic                 | africa                           | central_african_republic                 | Ps      |
| central_anatolia                         | turkey                           | central_anatolia                         | Ps      |
| central_federal_district                 | russia                           | central_federal_district                 | Ps      |
| central_finland                          | finland                          | central_finland                          | Ps      |
| central_java                             | indonesia                        | central_java                             | Ps      |
| central_kalimantan                       | indonesia                        | central_kalimantan                       | Ps      |
| central_luzon                            | philippines                      | central_luzon                            | Ps      |
| central_new_york                         | new-york                         | central_new_york                         | Ps      |
| central_ontario                          | ontario                          | central_ontario                          | Ps      |
| central_ostrobothnia                     | finland                          | central_ostrobothnia                     | Ps      |
| central_papua                            | indonesia                        | central_papua                            | Ps      |
| central_sulawesi                         | indonesia                        | central_sulawesi                         | Ps      |
| central_visayas                          | philippines                      | central_visayas                          | Ps      |
| centre                                   | france                           | centre                                   | Ps      |
| centre_du_quebec                         | quebec                           | centre_du_quebec                         | Ps      |
| ceuta                                    | spain                            | ceuta                                    | Ps      |
| chaco                                    | argentina                        | chaco                                    | Ps      |
| chad                                     | africa                           | chad                                     | Ps      |
| champagne_ardenne                        | france                           | champagne_ardenne                        | Ps      |
| chandigarh                               | india                            | chandigarh                               | Ps      |
| charente                                 | poitou_charentes                 | charente                                 | Ps      |
| charente_maritime                        | poitou_charentes                 | charente_maritime                        | Ps      |
| chaudiere_appalaches                     | quebec                           | chaudiere_appalaches                     | Ps      |
| chechen_republic                         | north_caucasian_federal_district | chechen_republic                         | Ps      |
| chelyabinsk_oblast                       | ural_federal_district            | chelyabinsk_oblast                       | Ps      |
| cher                                     | centre                           | cher                                     | Ps      |
| cherkasy_oblast                          | ukraine                          | cherkasy_oblast                          | Ps      |
| chernihiv_oblast                         | ukraine                          | chernihiv_oblast                         | Ps      |
| chernivtsi_oblast                        | ukraine                          | chernivtsi_oblast                        | Ps      |
| chesapeake                               | virginia                         | chesapeake                               | Ps      |
| chhattisgarh                             | india                            | chhattisgarh                             | Ps      |
| chiapas                                  | mexico                           | chiapas                                  | Ps      |
| chihuahua                                | mexico                           | chihuahua                                | Ps      |
| chile                                    | south-america                    | chile                                    | Ps      |
| china                                    | asia                             | china                                    | Ps      |
| chongqing                                | china                            | chongqing                                | Ps      |
| christmas_island                         | australia                        | christmas_island                         | Ps      |
| chubu                                    | japan                            | chubu                                    | Ps      |
| chubut                                   | argentina                        | chubut                                   | Ps      |
| chugoku                                  | japan                            | chugoku                                  | Ps      |
| chukotka_autonomous_okrug                | far_eastern_federal_district     | chukotka_autonomous_okrug                | Ps      |
| chuvash_republic                         | volga_federal_district           | chuvash_republic                         | Ps      |
| ciudad_real                              | castilla_la_mancha               | ciudad_real                              | Ps      |
| clipperton                               | oceania                          | clipperton                               | Ps      |
| coahuila                                 | mexico                           | coahuila                                 | Ps      |
| coastal                                  | tanzania                         | coastal                                  | Ps      |
| cocos_islands                            | australia                        | cocos_islands                            | Ps      |
| coimbra                                  | portugal                         | coimbra                                  | Ps      |
| colima                                   | mexico                           | colima                                   | Ps      |
| colorado                                 | us-west                          | colorado                                 | Ps      |
| colorado_northeast                       | colorado                         | colorado_northeast                       | Ps      |
| colorado_northwest                       | colorado                         | colorado_northwest                       | Ps      |
| colorado_southeast                       | colorado                         | colorado_southeast                       | Ps      |
| colorado_southwest                       | colorado                         | colorado_southwest                       | Ps      |
| colusa                                   | california                       | colusa                                   | Ps      |
| comoros                                  | africa                           | comoros                                  | Ps      |
| comunidad_de_madrid                      | spain                            | comunidad_de_madrid                      | Ps      |
| comunidad_foral_de_navarra               | spain                            | comunidad_foral_de_navarra               | Ps      |
| comunitat_valenciana                     | spain                            | comunitat_valenciana                     | Ps      |
| congo_brazzaville                        | africa                           | congo_brazzaville                        | Ps      |
| congo_kinshasa                           | africa                           | congo_kinshasa                           | Ps      |
| contra_costa                             | california                       | contra_costa                             | Ps      |
| cook                                     | illinois                         | cook                                     | Ps      |
| cook_islands                             | oceania                          | cook_islands                             | Ps      |
| coquimbo                                 | chile                            | coquimbo                                 | Ps      |
| coral_sea_islands                        | australia                        | coral_sea_islands                        | Ps      |
| cordillera_administrative_region         | philippines                      | cordillera_administrative_region         | Ps      |
| cordoba                                  | argentina                        | cordoba                                  | Ps      |
| correze                                  | limousin                         | correze                                  | Ps      |
| corrientes                               | argentina                        | corrientes                               | Ps      |
| corse                                    | france                           | corse                                    | Ps      |
| corse_du_sud                             | corse                            | corse_du_sud                             | Ps      |
| costa_rica                               | central-america                  | costa_rica                               | Ps      |
| cote_d_or                                | bourgogne                        | cote_d_or                                | Ps      |
| cote_nord                                | quebec                           | cote_nord                                | Ps      |
| cotes_d_armor                            | bretagne                         | cotes_d_armor                            | Ps      |
| creuse                                   | limousin                         | creuse                                   | Ps      |
| crimea                                   | ukraine                          | crimea                                   | Ps      |
| crimea_republic                          | southern_federal_district        | crimea_republic                          | Ps      |
| cuenca                                   | castilla_la_mancha               | cuenca                                   | Ps      |
| culpeper                                 | virginia                         | culpeper                                 | Ps      |
| curacao                                  | central-america                  | curacao                                  | Ps      |
| czech_republic                           | europe                           | czech_republic                           | Ps      |
| dadra_and_nagar_haveli                   | india                            | dadra_and_nagar_haveli                   | Ps      |
| dadra_and_nagar_haveli_and_daman_and_diu | india                            | dadra_and_nagar_haveli_and_daman_and_diu | Ps      |
| dagestan_republic                        | north_caucasian_federal_district | dagestan_republic                        | Ps      |
| dalarna                                  | sweden                           | dalarna                                  | Ps      |
| daman_and_diu                            | india                            | daman_and_diu                            | Ps      |
| davao_region                             | philippines                      | davao_region                             | Ps      |
| del_norte                                | california                       | del_norte                                | Ps      |
| denmark                                  | europe                           | denmark                                  | Ps      |
| denver                                   | colorado                         | denver                                   | Ps      |
| detmold                                  | nordrhein_westfalen              | detmold                                  | Ps      |
| detroit_metro                            | michigan                         | detroit_metro                            | Ps      |
| deux_sevres                              | poitou_charentes                 | deux_sevres                              | Ps      |
| distrito-federal                         | central-west                     | distrito-federal                         | Ps      |
| djibouti                                 | africa                           | djibouti                                 | Ps      |
| dnipropetrovsk_oblast                    | ukraine                          | dnipropetrovsk_oblast                    | Ps      |
| dolnoslaskie                             | poland                           | dolnoslaskie                             | Ps      |
| dominica                                 | central-america                  | dominica                                 | Ps      |
| dominican_republic                       | central-america                  | dominican_republic                       | Ps      |
| donetsk_oblast                           | ukraine                          | donetsk_oblast                           | Ps      |
| dordogne                                 | aquitaine                        | dordogne                                 | Ps      |
| doubs                                    | franche_comte                    | doubs                                    | Ps      |
| drenthe                                  | netherlands                      | drenthe                                  | Ps      |
| drome                                    | rhone_alpes                      | drome                                    | Ps      |
| durango                                  | mexico                           | durango                                  | Ps      |
| dusseldorf                               | nordrhein_westfalen              | dusseldorf                               | Ps      |
| east_flanders                            | flanders                         | east_flanders                            | Ps      |
| east_java                                | indonesia                        | east_java                                | Ps      |
| east_kalimantan                          | indonesia                        | east_kalimantan                          | Ps      |
| east_midlands                            | england                          | east_midlands                            | Ps      |
| east_nusa_tenggara                       | indonesia                        | east_nusa_tenggara                       | Ps      |
| east_timor                               | asia                             | east_timor                               | Ps      |
| eastern_anatolia                         | turkey                           | eastern_anatolia                         | Ps      |
| eastern_cape                             | south_africa                     | eastern_cape                             | Ps      |
| eastern_ontario                          | ontario                          | eastern_ontario                          | Ps      |
| eastern_visayas                          | philippines                      | eastern_visayas                          | Ps      |
| egypt                                    | africa                           | egypt                                    | Ps      |
| el_dorado                                | california                       | el_dorado                                | Ps      |
| el_salvador                              | central-america                  | el_salvador                              | Ps      |
| emilia_romagna                           | italy                            | emilia_romagna                           | Ps      |
| england                                  | united_kingdom                   | england                                  | Ps      |
| england_east                             | england                          | england_east                             | Ps      |
| england_north_east                       | england                          | england_north_east                       | Ps      |
| england_north_west                       | england                          | england_north_west                       | Ps      |
| england_south_east                       | england                          | england_south_east                       | Ps      |
| england_south_west                       | england                          | england_south_west                       | Ps      |
| entre_rios                               | argentina                        | entre_rios                               | Ps      |
| equatorial_guinea                        | africa                           | equatorial_guinea                        | Ps      |
| eritrea                                  | africa                           | eritrea                                  | Ps      |
| espirito-santo                           | southeast                        | espirito-santo                           | Ps      |
| essonne                                  | ile_de_france                    | essonne                                  | Ps      |
| estrie                                   | quebec                           | estrie                                   | Ps      |
| ethiopia                                 | africa                           | ethiopia                                 | Ps      |
| eure                                     | haute_normandie                  | eure                                     | Ps      |
| eure_et_loir                             | centre                           | eure_et_loir                             | Ps      |
| europe                                   |                                  | europe                                   | Ps      |
| euskadi                                  | spain                            | euskadi                                  | Ps      |
| evora                                    | portugal                         | evora                                    | Ps      |
| extremadura                              | spain                            | extremadura                              | Ps      |
| fairfax                                  | virginia                         | fairfax                                  | Ps      |
| falkland                                 | south-america                    | falkland                                 | Ps      |
| far_eastern_federal_district             | russia                           | far_eastern_federal_district             | Ps      |
| far_west                                 | new_south_wales                  | far_west                                 | Ps      |
| faro                                     | portugal                         | faro                                     | Ps      |
| fiji                                     | merge                            | fiji                                     | Ps      |
| fiji_east                                | oceania                          | fiji_east                                | Ps      |
| fiji_west                                | oceania                          | fiji_west                                | Ps      |
| finger_lakes                             | new-york                         | finger_lakes                             | Ps      |
| finistere                                | bretagne                         | finistere                                | Ps      |
| finland                                  | europe                           | finland                                  | Ps      |
| finnmark                                 | norway                           | finnmark                                 | Ps      |
| flanders                                 | belgium                          | flanders                                 | Ps      |
| flemish_brabant                          | flanders                         | flemish_brabant                          | Ps      |
| flevoland                                | netherlands                      | flevoland                                | Ps      |
| florida                                  | us-south                         | florida                                  | Ps      |
| florida_east_central                     | florida                          | florida_east_central                     | Ps      |
| florida_northeast                        | florida                          | florida_northeast                        | Ps      |
| florida_northwest                        | florida                          | florida_northwest                        | Ps      |
| florida_southwest                        | florida                          | florida_southwest                        | Ps      |
| formosa                                  | argentina                        | formosa                                  | Ps      |
| france                                   | europe                           | france                                   | Ps      |
| france_metro_dom_com_nc                  | merge                            | france_metro_dom_com_nc                  | Ps      |
| franche_comte                            | france                           | franche_comte                            | Ps      |
| franche_comte_jura                       | franche_comte                    | franche_comte_jura                       | Ps      |
| free_state                               | south_africa                     | free_state                               | Ps      |
| fresno                                   | california                       | fresno                                   | Ps      |
| fribourg                                 | switzerland                      | fribourg                                 | Ps      |
| friesland                                | netherlands                      | friesland                                | Ps      |
| friuli_venezia_giulia                    | italy                            | friuli_venezia_giulia                    | Ps      |
| fujian                                   | china                            | fujian                                   | Ps      |
| gabon                                    | africa                           | gabon                                    | Ps      |
| galicia                                  | spain                            | galicia                                  | Ps      |
| gambia                                   | africa                           | gambia                                   | Ps      |
| gansu                                    | china                            | gansu                                    | Ps      |
| gard                                     | languedoc_roussillon             | gard                                     | Ps      |
| gaspesie_iles_de_la_madeleine            | quebec                           | gaspesie_iles_de_la_madeleine            | Ps      |
| gatorland                                | florida                          | gatorland                                | Ps      |
| gauteng                                  | south_africa                     | gauteng                                  | Ps      |
| gavleborg                                | sweden                           | gavleborg                                | Ps      |
| gaza                                     | mozambique                       | gaza                                     | Ps      |
| gelderland                               | netherlands                      | gelderland                               | Ps      |
| geneva                                   | switzerland                      | geneva                                   | Ps      |
| georgia                                  | asia                             | georgia                                  | Ps      |
| georgia_northeast                        | georgia                          | georgia_northeast                        | Ps      |
| georgia_northwest                        | georgia                          | georgia_northwest                        | Ps      |
| georgia_southeast                        | georgia                          | georgia_southeast                        | Ps      |
| georgia_southwest                        | georgia                          | georgia_southwest                        | Ps      |
| germany                                  | europe                           | germany                                  | Ps      |
| gers                                     | midi_pyrenees                    | gers                                     | Ps      |
| ghana                                    | africa                           | ghana                                    | Ps      |
| gibraltar                                | europe                           | gibraltar                                | Ps      |
| girona                                   | catalunya                        | girona                                   | Ps      |
| gironde                                  | aquitaine                        | gironde                                  | Ps      |
| glarus                                   | switzerland                      | glarus                                   | Ps      |
| glenn                                    | california                       | glenn                                    | Ps      |
| goa                                      | india                            | goa                                      | Ps      |
| goias                                    | central-west                     | goias                                    | Ps      |
| gold_coast                               | florida                          | gold_coast                               | Ps      |
| golden_horseshoe                         | ontario                          | golden_horseshoe                         | Ps      |
| gorontalo                                | indonesia                        | gorontalo                                | Ps      |
| gotland                                  | sweden                           | gotland                                  | Ps      |
| granada                                  | andalucia                        | granada                                  | Ps      |
| greater_london                           | england                          | greater_london                           | Ps      |
| greater_metropolitan_newcastle           | new_south_wales                  | greater_metropolitan_newcastle           | Ps      |
| greater_metropolitan_sydney              | new_south_wales                  | greater_metropolitan_sydney              | Ps      |
| grenada                                  | central-america                  | grenada                                  | Ps      |
| grisons                                  | switzerland                      | grisons                                  | Ps      |
| groningen                                | netherlands                      | groningen                                | Ps      |
| guadalajara                              | castilla_la_mancha               | guadalajara                              | Ps      |
| guadeloupe                               | central-america                  | guadeloupe                               | Ps      |
| guam                                     | oceania                          | guam                                     | Ps      |
| guanajuato                               | mexico                           | guanajuato                               | Ps      |
| guangdong                                | china                            | guangdong                                | Ps      |
| guangxi                                  | china                            | guangxi                                  | Ps      |
| guarda                                   | portugal                         | guarda                                   | Ps      |
| guernesey                                | europe                           | guernesey                                | Ps      |
| guerrero                                 | mexico                           | guerrero                                 | Ps      |
| guinea                                   | africa                           | guinea                                   | Ps      |
| guipuzcoa                                | euskadi                          | guipuzcoa                                | Ps      |
| guizhou                                  | china                            | guizhou                                  | Ps      |
| gujarat                                  | india                            | gujarat                                  | Ps      |
| guyana                                   | south-america                    | guyana                                   | Ps      |
| guyane                                   | south-america                    | guyane                                   | Ps      |
| hainan                                   | china                            | hainan                                   | Ps      |
| haiti                                    | central-america                  | haiti                                    | Ps      |
| halland                                  | sweden                           | halland                                  | Ps      |
| haryana                                  | india                            | haryana                                  | Ps      |
| haut_rhin                                | alsace                           | haut_rhin                                | Ps      |
| haute_corse                              | corse                            | haute_corse                              | Ps      |
| haute_garonne                            | midi_pyrenees                    | haute_garonne                            | Ps      |
| haute_loire                              | auvergne                         | haute_loire                              | Ps      |
| haute_marne                              | champagne_ardenne                | haute_marne                              | Ps      |
| haute_normandie                          | france                           | haute_normandie                          | Ps      |
| haute_saone                              | franche_comte                    | haute_saone                              | Ps      |
| haute_savoie                             | rhone_alpes                      | haute_savoie                             | Ps      |
| haute_vienne                             | limousin                         | haute_vienne                             | Ps      |
| hautes_alpes                             | provence_alpes_cote_d_azur       | hautes_alpes                             | Ps      |
| hautes_pyrenees                          | midi_pyrenees                    | hautes_pyrenees                          | Ps      |
| hauts_de_seine                           | ile_de_france                    | hauts_de_seine                           | Ps      |
| heard_island_and_mcdonald_slands         | australia                        | heard_island_and_mcdonald_slands         | Ps      |
| hebei                                    | china                            | hebei                                    | Ps      |
| hedmark                                  | norway                           | hedmark                                  | Ps      |
| heilongjiang                             | china                            | heilongjiang                             | Ps      |
| henan                                    | china                            | henan                                    | Ps      |
| herault                                  | languedoc_roussillon             | herault                                  | Ps      |
| hidalgo                                  | mexico                           | hidalgo                                  | Ps      |
| highland_papua                           | indonesia                        | highland_papua                           | Ps      |
| himachal_pradesh                         | india                            | himachal_pradesh                         | Ps      |
| hokkaido                                 | japan                            | hokkaido                                 | Ps      |
| honduras                                 | central-america                  | honduras                                 | Ps      |
| hong_kong                                | china                            | hong_kong                                | Ps      |
| hordaland                                | norway                           | hordaland                                | Ps      |
| hovedstaden                              | denmark                          | hovedstaden                              | Ps      |
| hubei                                    | china                            | hubei                                    | Ps      |
| hudson_valley                            | new-york                         | hudson_valley                            | Ps      |
| huelva                                   | andalucia                        | huelva                                   | Ps      |
| huesca                                   | aragon                           | huesca                                   | Ps      |
| humboldt                                 | california                       | humboldt                                 | Ps      |
| hunan                                    | china                            | hunan                                    | Ps      |
| ile_de_france                            | france                           | ile_de_france                            | Ps      |
| ilemi                                    | africa                           | ilemi                                    | Ps      |
| illawarra                                | new_south_wales                  | illawarra                                | Ps      |
| ille_et_vilaine                          | bretagne                         | ille_et_vilaine                          | Ps      |
| illes_balears                            | spain                            | illes_balears                            | Ps      |
| illinois                                 | us-midwest                       | illinois                                 | Ps      |
| illinois_central                         | illinois                         | illinois_central                         | Ps      |
| illinois_east_central                    | illinois                         | illinois_east_central                    | Ps      |
| illinois_north                           | illinois                         | illinois_north                           | Ps      |
| illinois_northeast                       | illinois                         | illinois_northeast                       | Ps      |
| illinois_northwest                       | illinois                         | illinois_northwest                       | Ps      |
| illinois_southern                        | illinois                         | illinois_southern                        | Ps      |
| illinois_southwest                       | illinois                         | illinois_southwest                       | Ps      |
| ilocos_region                            | philippines                      | ilocos_region                            | Ps      |
| imperial                                 | california                       | imperial                                 | Ps      |
| india                                    | asia                             | india                                    | Ps      |
| indonesia                                | asia                             | indonesia                                | Ps      |
| indre                                    | centre                           | indre                                    | Ps      |
| indre_et_loire                           | centre                           | indre_et_loire                           | Ps      |
| ingushetia_republic                      | north_caucasian_federal_district | ingushetia_republic                      | Ps      |
| inhambane                                | mozambique                       | inhambane                                | Ps      |
| inner_mongolia                           | china                            | inner_mongolia                           | Ps      |
| inyo                                     | california                       | inyo                                     | Ps      |
| ionian_sea                               | seas                             | ionian_sea                               | Ps      |
| ireland                                  | europe                           | ireland                                  | Ps      |
| irkutsk_oblast                           | siberian_federal_district        | irkutsk_oblast                           | Ps      |
| isere                                    | rhone_alpes                      | isere                                    | Ps      |
| israel                                   | asia                             | israel                                   | Ps      |
| israel_and_palestine                     | asia                             | israel_and_palestine                     | Ps      |
| israel_west_bank                         | asia                             | israel_west_bank                         | Ps      |
| italy                                    | europe                           | italy                                    | Ps      |
| ivano-frankivsk_oblast                   | ukraine                          | ivano-frankivsk_oblast                   | Ps      |
| ivanovo_oblast                           | central_federal_district         | ivanovo_oblast                           | Ps      |
| ivory_coast                              | africa                           | ivory_coast                              | Ps      |
| jaen                                     | andalucia                        | jaen                                     | Ps      |
| jakarta                                  | indonesia                        | jakarta                                  | Ps      |
| jalisco                                  | mexico                           | jalisco                                  | Ps      |
| jamaica                                  | central-america                  | jamaica                                  | Ps      |
| jambi                                    | indonesia                        | jambi                                    | Ps      |
| jammu_and_kashmir                        | india                            | jammu_and_kashmir                        | Ps      |
| jamtland                                 | sweden                           | jamtland                                 | Ps      |
| jan_mayen                                | norway                           | jan_mayen                                | Ps      |
| japan                                    | asia                             | japan                                    | Ps      |
| jersey                                   | europe                           | jersey                                   | Ps      |
| jewish_autonomous_oblast                 | far_eastern_federal_district     | jewish_autonomous_oblast                 | Ps      |
| jharkhand                                | india                            | jharkhand                                | Ps      |
| jiangsu                                  | china                            | jiangsu                                  | Ps      |
| jiangxi                                  | china                            | jiangxi                                  | Ps      |
| jihocesky                                | czech_republic                   | jihocesky                                | Ps      |
| jihomoravsky                             | czech_republic                   | jihomoravsky                             | Ps      |
| jilin                                    | china                            | jilin                                    | Ps      |
| jonkoping                                | sweden                           | jonkoping                                | Ps      |
| jujuy                                    | argentina                        | jujuy                                    | Ps      |
| kabardino_balkar_republic                | north_caucasian_federal_district | kabardino_balkar_republic                | Ps      |
| kainuu                                   | finland                          | kainuu                                   | Ps      |
| kaliningrad_oblast                       | northwestern_federal_district    | kaliningrad_oblast                       | Ps      |
| kalmar                                   | sweden                           | kalmar                                   | Ps      |
| kalmykia_republic                        | southern_federal_district        | kalmykia_republic                        | Ps      |
| kaluga_oblast                            | central_federal_district         | kaluga_oblast                            | Ps      |
| kamchatka_krai                           | far_eastern_federal_district     | kamchatka_krai                           | Ps      |
| kansai                                   | japan                            | kansai                                   | Ps      |
| kanta_hame                               | finland                          | kanta_hame                               | Ps      |
| kanto                                    | japan                            | kanto                                    | Ps      |
| karachay_cherkess_republic               | north_caucasian_federal_district | karachay_cherkess_republic               | Ps      |
| karelia_republic                         | northwestern_federal_district    | karelia_republic                         | Ps      |
| karlovarsky                              | czech_republic                   | karlovarsky                              | Ps      |
| karnataka                                | india                            | karnataka                                | Ps      |
| karnten                                  | austria                          | karnten                                  | Ps      |
| kemerovo_oblast                          | siberian_federal_district        | kemerovo_oblast                          | Ps      |
| kenya                                    | africa                           | kenya                                    | Ps      |
| kerala                                   | india                            | kerala                                   | Ps      |
| kern                                     | california                       | kern                                     | Ps      |
| khabarovsk_krai                          | far_eastern_federal_district     | khabarovsk_krai                          | Ps      |
| khakassia_republic                       | siberian_federal_district        | khakassia_republic                       | Ps      |
| khanty_mansi_autonomous_okrug            | ural_federal_district            | khanty_mansi_autonomous_okrug            | Ps      |
| kharkiv_oblast                           | ukraine                          | kharkiv_oblast                           | Ps      |
| kherson_oblast                           | ukraine                          | kherson_oblast                           | Ps      |
| khmelnytskyi_oblast                      | ukraine                          | khmelnytskyi_oblast                      | Ps      |
| kiev                                     | ukraine                          | kiev                                     | Ps      |
| kiev_oblast                              | ukraine                          | kiev_oblast                              | Ps      |
| kings                                    | california                       | kings                                    | Ps      |
| kiribati                                 | merge                            | kiribati                                 | Ps      |
| kiribati_east                            | oceania                          | kiribati_east                            | Ps      |
| kiribati_west                            | oceania                          | kiribati_west                            | Ps      |
| kirov_oblast                             | volga_federal_district           | kirov_oblast                             | Ps      |
| kirovohrad_oblast                        | ukraine                          | kirovohrad_oblast                        | Ps      |
| koln                                     | nordrhein_westfalen              | koln                                     | Ps      |
| komi_republic                            | northwestern_federal_district    | komi_republic                            | Ps      |
| kosicky                                  | slovakia                         | kosicky                                  | Ps      |
| kostroma_oblast                          | central_federal_district         | kostroma_oblast                          | Ps      |
| kralovehradecky                          | czech_republic                   | kralovehradecky                          | Ps      |
| krasnodar_krai                           | southern_federal_district        | krasnodar_krai                           | Ps      |
| krasnoyarsk_krai                         | siberian_federal_district        | krasnoyarsk_krai                         | Ps      |
| kronoberg                                | sweden                           | kronoberg                                | Ps      |
| kujawsko_pomorskie                       | poland                           | kujawsko_pomorskie                       | Ps      |
| kurgan_oblast                            | ural_federal_district            | kurgan_oblast                            | Ps      |
| kursk_oblast                             | central_federal_district         | kursk_oblast                             | Ps      |
| kuwait                                   | asia                             | kuwait                                   | Ps      |
| kwazulu_natal                            | south_africa                     | kwazulu_natal                            | Ps      |
| kymenlaakso                              | finland                          | kymenlaakso                              | Ps      |
| kyushu                                   | japan                            | kyushu                                   | Ps      |
| la_coruna                                | galicia                          | la_coruna                                | Ps      |
| la_pampa                                 | argentina                        | la_pampa                                 | Ps      |
| ladakh                                   | india                            | ladakh                                   | Ps      |
| lakshadweep                              | india                            | lakshadweep                              | Ps      |
| lampung                                  | indonesia                        | lampung                                  | Ps      |
| lanaudiere                               | quebec                           | lanaudiere                               | Ps      |
| landes                                   | aquitaine                        | landes                                   | Ps      |
| languedoc_roussillon                     | france                           | languedoc_roussillon                     | Ps      |
| laos                                     | asia                             | laos                                     | Ps      |
| lapland                                  | finland                          | lapland                                  | Ps      |
| las_palmas                               | canarias                         | las_palmas                               | Ps      |
| lassen                                   | california                       | lassen                                   | Ps      |
| laurentides                              | quebec                           | laurentides                              | Ps      |
| laval                                    | quebec                           | laval                                    | Ps      |
| lazio                                    | italy                            | lazio                                    | Ps      |
| lebanon                                  | asia                             | lebanon                                  | Ps      |
| leiria                                   | portugal                         | leiria                                   | Ps      |
| leningrad_oblast                         | northwestern_federal_district    | leningrad_oblast                         | Ps      |
| leon                                     | castilla_y_leon                  | leon                                     | Ps      |
| lesotho                                  | africa                           | lesotho                                  | Ps      |
| liaoning                                 | china                            | liaoning                                 | Ps      |
| liberecky                                | czech_republic                   | liberecky                                | Ps      |
| liguria                                  | italy                            | liguria                                  | Ps      |
| limburg                                  | netherlands                      | limburg                                  | Ps      |
| limousin                                 | france                           | limousin                                 | Ps      |
| limpopo                                  | south_africa                     | limpopo                                  | Ps      |
| lipetsk_oblast                           | central_federal_district         | lipetsk_oblast                           | Ps      |
| lisbon                                   | portugal                         | lisbon                                   | Ps      |
| lleida                                   | catalunya                        | lleida                                   | Ps      |
| lodzkie                                  | poland                           | lodzkie                                  | Ps      |
| loir_et_cher                             | centre                           | loir_et_cher                             | Ps      |
| loire                                    | rhone_alpes                      | loire                                    | Ps      |
| loire_atlantique                         | pays_de_la_loire                 | loire_atlantique                         | Ps      |
| loiret                                   | centre                           | loiret                                   | Ps      |
| lombardia                                | italy                            | lombardia                                | Ps      |
| long_island                              | new-york                         | long_island                              | Ps      |
| lorraine                                 | france                           | lorraine                                 | Ps      |
| los_angeles                              | california                       | los_angeles                              | Ps      |
| los_lagos                                | chile                            | los_lagos                                | Ps      |
| los_rios                                 | chile                            | los_rios                                 | Ps      |
| lot                                      | midi_pyrenees                    | lot                                      | Ps      |
| lot_et_garonne                           | aquitaine                        | lot_et_garonne                           | Ps      |
| lozere                                   | languedoc_roussillon             | lozere                                   | Ps      |
| lubelskie                                | poland                           | lubelskie                                | Ps      |
| lubuskie                                 | poland                           | lubuskie                                 | Ps      |
| lucerne                                  | switzerland                      | lucerne                                  | Ps      |
| lugo                                     | galicia                          | lugo                                     | Ps      |
| luhansk_oblast                           | ukraine                          | luhansk_oblast                           | Ps      |
| luxembourg                               | europe                           | luxembourg                               | Ps      |
| lviv_oblast                              | ukraine                          | lviv_oblast                              | Ps      |
| macau                                    | china                            | macau                                    | Ps      |
| madagascar                               | africa                           | madagascar                               | Ps      |
| madeira                                  | portugal                         | madeira                                  | Ps      |
| madera                                   | california                       | madera                                   | Ps      |
| madhya_pradesh                           | india                            | madhya_pradesh                           | Ps      |
| magadan_oblast                           | far_eastern_federal_district     | magadan_oblast                           | Ps      |
| magallanes                               | chile                            | magallanes                               | Ps      |
| maharashtra                              | india                            | maharashtra                              | Ps      |
| maine_et_loire                           | pays_de_la_loire                 | maine_et_loire                           | Ps      |
| malaga                                   | andalucia                        | malaga                                   | Ps      |
| malawi                                   | africa                           | malawi                                   | Ps      |
| malaysia                                 | asia                             | malaysia                                 | Ps      |
| maldives                                 | asia                             | maldives                                 | Ps      |
| mali                                     | africa                           | mali                                     | Ps      |
| malopolskie                              | poland                           | malopolskie                              | Ps      |
| maluku                                   | indonesia                        | maluku                                   | Ps      |
| manche                                   | basse_normandie                  | manche                                   | Ps      |
| manica                                   | mozambique                       | manica                                   | Ps      |
| manipur                                  | india                            | manipur                                  | Ps      |
| manitoba                                 | canada                           | manitoba                                 | Ps      |
| maputo                                   | mozambique                       | maputo                                   | Ps      |
| maputo_city                              | mozambique                       | maputo_city                              | Ps      |
| maranhao                                 | northeast                        | maranhao                                 | Ps      |
| marche                                   | italy                            | marche                                   | Ps      |
| mari_el_republic                         | volga_federal_district           | mari_el_republic                         | Ps      |
| marin                                    | california                       | marin                                    | Ps      |
| mariposa                                 | california                       | mariposa                                 | Ps      |
| marmara                                  | turkey                           | marmara                                  | Ps      |
| marne                                    | champagne_ardenne                | marne                                    | Ps      |
| marshall-islands                         | oceania                          | marshall-islands                         | Ps      |
| marshall_islands                         | oceania                          | marshall_islands                         | Ps      |
| martinique                               | central-america                  | martinique                               | Ps      |
| mato-grosso                              | central-west                     | mato-grosso                              | Ps      |
| mato-grosso-do-sul                       | central-west                     | mato-grosso-do-sul                       | Ps      |
| maule                                    | chile                            | maule                                    | Ps      |
| mauricie                                 | quebec                           | mauricie                                 | Ps      |
| mauritania                               | africa                           | mauritania                               | Ps      |
| mauritius                                | africa                           | mauritius                                | Ps      |
| mayenne                                  | pays_de_la_loire                 | mayenne                                  | Ps      |
| mayotte                                  | africa                           | mayotte                                  | Ps      |
| mazowieckie                              | poland                           | mazowieckie                              | Ps      |
| mediterranean                            | turkey                           | mediterranean                            | Ps      |
| meghalaya                                | india                            | meghalaya                                | Ps      |
| melilla                                  | spain                            | melilla                                  | Ps      |
| mendocino                                | california                       | mendocino                                | Ps      |
| mendoza                                  | argentina                        | mendoza                                  | Ps      |
| merced                                   | california                       | merced                                   | Ps      |
| merge                                    |                                  | merge                                    |         |
| merge_france_taaf                        | merge                            | merge_france_taaf                        | Ps      |
| metro_manila                             | philippines                      | metro_manila                             | Ps      |
| meurthe_et_moselle                       | lorraine                         | meurthe_et_moselle                       | Ps      |
| meuse                                    | lorraine                         | meuse                                    | Ps      |
| mexico                                   | north-america                    | mexico                                   | Ps      |
| mexico_city                              | mexico                           | mexico_city                              | Ps      |
| michigan                                 | us-midwest                       | michigan                                 | Ps      |
| michigan_central                         | michigan                         | michigan_central                         | Ps      |
| michigan_southeast                       | michigan                         | michigan_southeast                       | Ps      |
| michigan_southwest                       | michigan                         | michigan_southwest                       | Ps      |
| michigan_west                            | michigan                         | michigan_west                            | Ps      |
| michoacan                                | mexico                           | michoacan                                | Ps      |
| micronesia                               | oceania                          | micronesia                               | Ps      |
| mid_north_coast                          | new_south_wales                  | mid_north_coast                          | Ps      |
| midi_pyrenees                            | france                           | midi_pyrenees                            | Ps      |
| midtjylland                              | denmark                          | midtjylland                              | Ps      |
| mimaropa                                 | philippines                      | mimaropa                                 | Ps      |
| minas-gerais                             | southeast                        | minas-gerais                             | Ps      |
| misiones                                 | argentina                        | misiones                                 | Ps      |
| mizoram                                  | india                            | mizoram                                  | Ps      |
| modoc                                    | california                       | modoc                                    | Ps      |
| moere_og_romsdal                         | norway                           | moere_og_romsdal                         | Ps      |
| mohawk_valley                            | new-york                         | mohawk_valley                            | Ps      |
| molise                                   | italy                            | molise                                   | Ps      |
| monaco                                   | europe                           | monaco                                   | Ps      |
| mono                                     | california                       | mono                                     | Ps      |
| monteregie                               | quebec                           | monteregie                               | Ps      |
| monterey                                 | california                       | monterey                                 | Ps      |
| montreal                                 | quebec                           | montreal                                 | Ps      |
| montserrat                               | central-america                  | montserrat                               | Ps      |
| moravskoslezsky                          | czech_republic                   | moravskoslezsky                          | Ps      |
| morbihan                                 | bretagne                         | morbihan                                 | Ps      |
| mordovia_republic                        | volga_federal_district           | mordovia_republic                        | Ps      |
| morelos                                  | mexico                           | morelos                                  | Ps      |
| morocco                                  | africa                           | morocco                                  | Ps      |
| moscow                                   | central_federal_district         | moscow                                   | Ps      |
| moscow_oblast                            | central_federal_district         | moscow_oblast                            | Ps      |
| moselle                                  | lorraine                         | moselle                                  | Ps      |
| mozambique                               | africa                           | mozambique                               | Ps      |
| mpumalanga                               | south_africa                     | mpumalanga                               | Ps      |
| munster                                  | nordrhein_westfalen              | munster                                  | Ps      |
| murmansk_oblast                          | northwestern_federal_district    | murmansk_oblast                          | Ps      |
| murray                                   | new_south_wales                  | murray                                   | Ps      |
| myanmar                                  | asia                             | myanmar                                  | Ps      |
| mykolaiv_oblast                          | ukraine                          | mykolaiv_oblast                          | Ps      |
| nagaland                                 | india                            | nagaland                                 | Ps      |
| namibia                                  | africa                           | namibia                                  | Ps      |
| nampula                                  | mozambique                       | nampula                                  | Ps      |
| napa                                     | california                       | napa                                     | Ps      |
| national_capital_territory_of_delhi      | india                            | national_capital_territory_of_delhi      | Ps      |
| nauru                                    | oceania                          | nauru                                    | Ps      |
| nayarit                                  | mexico                           | nayarit                                  | Ps      |
| negros_island_region                     | philippines                      | negros_island_region                     | Ps      |
| nenets_autonomous_okrug                  | northwestern_federal_district    | nenets_autonomous_okrug                  | Ps      |
| netherlands                              | europe                           | netherlands                              | Ps      |
| neuchatel                                | switzerland                      | neuchatel                                | Ps      |
| neuquen                                  | argentina                        | neuquen                                  | Ps      |
| nevada                                   | california                       | nevada                                   | Ps      |
| new-york                                 | us-northeast                     | new-york                                 | Ps      |
| new_brunswick                            | canada                           | new_brunswick                            | Ps      |
| new_caledonia                            | oceania                          | new_caledonia                            | Ps      |
| new_south_wales                          | australia                        | new_south_wales                          | Ps      |
| new_south_wales_central_west             | new_south_wales                  | new_south_wales_central_west             | Ps      |
| new_south_wales_north_western            | new_south_wales                  | new_south_wales_north_western            | Ps      |
| new_south_wales_northern                 | new_south_wales                  | new_south_wales_northern                 | Ps      |
| new_york_city                            | new-york                         | new_york_city                            | Ps      |
| newfoundland_and_labrador                | canada                           | newfoundland_and_labrador                | Ps      |
| niassa                                   | mozambique                       | niassa                                   | Ps      |
| nicaragua                                | central-america                  | nicaragua                                | Ps      |
| nidwalden                                | switzerland                      | nidwalden                                | Ps      |
| niederosterreich                         | austria                          | niederosterreich                         | Ps      |
| nievre                                   | bourgogne                        | nievre                                   | Ps      |
| niger                                    | africa                           | niger                                    | Ps      |
| nigeria                                  | africa                           | nigeria                                  | Ps      |
| nigeria_north_central                    | nigeria                          | nigeria_north_central                    | Ps      |
| nigeria_north_east                       | nigeria                          | nigeria_north_east                       | Ps      |
| nigeria_north_west                       | nigeria                          | nigeria_north_west                       | Ps      |
| nigeria_south_east                       | nigeria                          | nigeria_south_east                       | Ps      |
| nigeria_south_south                      | nigeria                          | nigeria_south_south                      | Ps      |
| nigeria_south_west                       | nigeria                          | nigeria_south_west                       | Ps      |
| ningxia                                  | china                            | ningxia                                  | Ps      |
| nitriansky                               | slovakia                         | nitriansky                               | Ps      |
| niue                                     | oceania                          | niue                                     | Ps      |
| nizhny_novgorod_oblast                   | volga_federal_district           | nizhny_novgorod_oblast                   | Ps      |
| noord_brabant                            | netherlands                      | noord_brabant                            | Ps      |
| noord_holland                            | netherlands                      | noord_holland                            | Ps      |
| nord                                     | nord_pas_de_calais               | nord                                     | Ps      |
| nord_du_quebec                           | quebec                           | nord_du_quebec                           | Ps      |
| nord_pas_de_calais                       | france                           | nord_pas_de_calais                       | Ps      |
| nordjylland                              | denmark                          | nordjylland                              | Ps      |
| nordland                                 | norway                           | nordland                                 | Ps      |
| nordrhein_westfalen                      | germany                          | nordrhein_westfalen                      | Ps      |
| norfolk_island                           | australia                        | norfolk_island                           | Ps      |
| norrbotten                               | sweden                           | norrbotten                               | Ps      |
| north                                    | brazil                           | north                                    |         |
| north-america                            |                                  | north-america                            | Ps      |
| north-carolina                           | us-south                         | north-carolina                           | Ps      |
| north-carolina_north_central             | north-carolina                   | north-carolina_north_central             | Ps      |
| north-carolina_northeast                 | north-carolina                   | north-carolina_northeast                 | Ps      |
| north-carolina_northwest                 | north-carolina                   | north-carolina_northwest                 | Ps      |
| north-carolina_south_central             | north-carolina                   | north-carolina_south_central             | Ps      |
| north-carolina_southeast                 | north-carolina                   | north-carolina_southeast                 | Ps      |
| north-carolina_western                   | north-carolina                   | north-carolina_western                   | Ps      |
| north_caucasian_federal_district         | russia                           | north_caucasian_federal_district         | Ps      |
| north_country                            | new-york                         | north_country                            | Ps      |
| north_kalimantan                         | indonesia                        | north_kalimantan                         | Ps      |
| north_karelia                            | finland                          | north_karelia                            | Ps      |
| north_maluku                             | indonesia                        | north_maluku                             | Ps      |
| north_metro                              | georgia                          | north_metro                              | Ps      |
| north_ossetia_alania_republic            | north_caucasian_federal_district | north_ossetia_alania_republic            | Ps      |
| north_ostrobothnia                       | finland                          | north_ostrobothnia                       | Ps      |
| north_savo                               | finland                          | north_savo                               | Ps      |
| north_sea                                | seas                             | north_sea                                | Ps      |
| north_sulawesi                           | indonesia                        | north_sulawesi                           | Ps      |
| north_sumatra                            | indonesia                        | north_sumatra                            | Ps      |
| northeast                                | brazil                           | northeast                                |         |
| northeastern_ontario                     | ontario                          | northeastern_ontario                     | Ps      |
| northern_cape                            | south_africa                     | northern_cape                            | Ps      |
| northern_ireland                         | united_kingdom                   | northern_ireland                         | Ps      |
| northern_lower                           | michigan                         | northern_lower                           | Ps      |
| northern_mariana_islands                 | oceania                          | northern_mariana_islands                 | Ps      |
| northern_mindanao                        | philippines                      | northern_mindanao                        | Ps      |
| northern_territory                       | australia                        | northern_territory                       | Ps      |
| northwest_territories                    | canada                           | northwest_territories                    | Ps      |
| northwestern_federal_district            | russia                           | northwestern_federal_district            | Ps      |
| northwestern_ontario                     | ontario                          | northwestern_ontario                     | Ps      |
| norway                                   | europe                           | norway                                   | Ps      |
| nova_scotia                              | canada                           | nova_scotia                              | Ps      |
| novgorod_oblast                          | northwestern_federal_district    | novgorod_oblast                          | Ps      |
| novosibirsk_oblast                       | siberian_federal_district        | novosibirsk_oblast                       | Ps      |
| nuble                                    | chile                            | nuble                                    | Ps      |
| nuevo_leon                               | mexico                           | nuevo_leon                               | Ps      |
| nunavut                                  | canada                           | nunavut                                  | Ps      |
| o_higgins                                | chile                            | o_higgins                                | Ps      |
| oaxaca                                   | mexico                           | oaxaca                                   | Ps      |
| oberosterreich                           | austria                          | oberosterreich                           | Ps      |
| obwalden                                 | switzerland                      | obwalden                                 | Ps      |
| oceania                                  |                                  | oceania                                  | Ps      |
| oceania_france_taaf                      | oceania                          | oceania_france_taaf                      | Ps      |
| odessa_oblast                            | ukraine                          | odessa_oblast                            | Ps      |
| odisha                                   | india                            | odisha                                   | Ps      |
| oestfold                                 | norway                           | oestfold                                 | Ps      |
| oise                                     | picardie                         | oise                                     | Ps      |
| olomoucky                                | czech_republic                   | olomoucky                                | Ps      |
| oman                                     | asia                             | oman                                     | Ps      |
| omsk_oblast                              | siberian_federal_district        | omsk_oblast                              | Ps      |
| ontario                                  | canada                           | ontario                                  | Ps      |
| opolskie                                 | poland                           | opolskie                                 | Ps      |
| oppland                                  | norway                           | oppland                                  | Ps      |
| orange                                   | california                       | orange                                   | Ps      |
| orebro                                   | sweden                           | orebro                                   | Ps      |
| orenburg_oblast                          | volga_federal_district           | orenburg_oblast                          | Ps      |
| orne                                     | basse_normandie                  | orne                                     | Ps      |
| oryol_oblast                             | central_federal_district         | oryol_oblast                             | Ps      |
| oslo                                     | norway                           | oslo                                     | Ps      |
| ostergotland                             | sweden                           | ostergotland                             | Ps      |
| ostrobothnia                             | finland                          | ostrobothnia                             | Ps      |
| ourense                                  | galicia                          | ourense                                  | Ps      |
| outaouais                                | quebec                           | outaouais                                | Ps      |
| overijssel                               | netherlands                      | overijssel                               | Ps      |
| paijat_hame                              | finland                          | paijat_hame                              | Ps      |
| palau                                    | oceania                          | palau                                    | Ps      |
| palencia                                 | castilla_y_leon                  | palencia                                 | Ps      |
| palestine                                | asia                             | palestine                                | Ps      |
| panama                                   | central-america                  | panama                                   | Ps      |
| panhandle                                | florida                          | panhandle                                | Ps      |
| papua                                    | indonesia                        | papua                                    | Ps      |
| papua_new_guinea                         | oceania                          | papua_new_guinea                         | Ps      |
| para                                     | north                            | para                                     | Ps      |
| paraguay                                 | south-america                    | paraguay                                 | Ps      |
| paraiba                                  | northeast                        | paraiba                                  | Ps      |
| parana                                   | south                            | parana                                   | Ps      |
| pardubicky                               | czech_republic                   | pardubicky                               | Ps      |
| paris                                    | ile_de_france                    | paris                                    | Ps      |
| pas_de_calais                            | nord_pas_de_calais               | pas_de_calais                            | Ps      |
| pays_de_la_loire                         | france                           | pays_de_la_loire                         | Ps      |
| penza_oblast                             | volga_federal_district           | penza_oblast                             | Ps      |
| perm_krai                                | volga_federal_district           | perm_krai                                | Ps      |
| pernambuco                               | northeast                        | pernambuco                               | Ps      |
| philippines                              | asia                             | philippines                              | Ps      |
| piaui                                    | northeast                        | piaui                                    | Ps      |
| picardie                                 | france                           | picardie                                 | Ps      |
| piedmont_triad                           | north-carolina                   | piedmont_triad                           | Ps      |
| piemonte                                 | italy                            | piemonte                                 | Ps      |
| pirkanmaa                                | finland                          | pirkanmaa                                | Ps      |
| pitcairn                                 | oceania                          | pitcairn                                 | Ps      |
| placer                                   | california                       | placer                                   | Ps      |
| plumas                                   | california                       | plumas                                   | Ps      |
| plzensky                                 | czech_republic                   | plzensky                                 | Ps      |
| podkarpackie                             | poland                           | podkarpackie                             | Ps      |
| podlaskie                                | poland                           | podlaskie                                | Ps      |
| poitou_charentes                         | france                           | poitou_charentes                         | Ps      |
| poland                                   | europe                           | poland                                   | Ps      |
| poltava_oblast                           | ukraine                          | poltava_oblast                           | Ps      |
| polynesie                                | oceania                          | polynesie                                | Ps      |
| pomorskie                                | poland                           | pomorskie                                | Ps      |
| pontevedra                               | galicia                          | pontevedra                               | Ps      |
| portalegre                               | portugal                         | portalegre                               | Ps      |
| porto                                    | portugal                         | porto                                    | Ps      |
| portugal                                 | europe                           | portugal                                 | Ps      |
| praha                                    | czech_republic                   | praha                                    | Ps      |
| presovsky                                | slovakia                         | presovsky                                | Ps      |
| primorsky_krai                           | far_eastern_federal_district     | primorsky_krai                           | Ps      |
| prince_edward_island                     | canada                           | prince_edward_island                     | Ps      |
| provence_alpes_cote_d_azur               | france                           | provence_alpes_cote_d_azur               | Ps      |
| pskov_oblast                             | northwestern_federal_district    | pskov_oblast                             | Ps      |
| puducherry                               | india                            | puducherry                               | Ps      |
| puebla                                   | mexico                           | puebla                                   | Ps      |
| puerto_rico                              | central-america                  | puerto_rico                              | Ps      |
| puglia                                   | italy                            | puglia                                   | Ps      |
| punjab                                   | india                            | punjab                                   | Ps      |
| puy_de_dome                              | auvergne                         | puy_de_dome                              | Ps      |
| pyrenees_atlantiques                     | aquitaine                        | pyrenees_atlantiques                     | Ps      |
| pyrenees_orientales                      | languedoc_roussillon             | pyrenees_orientales                      | Ps      |
| qatar                                    | asia                             | qatar                                    | Ps      |
| qinghai                                  | china                            | qinghai                                  | Ps      |
| quebec                                   | canada                           | quebec                                   | Ps      |
| queensland                               | australia                        | queensland                               | Ps      |
| queretaro                                | mexico                           | queretaro                                | Ps      |
| quintana_roo                             | mexico                           | quintana_roo                             | Ps      |
| rajasthan                                | india                            | rajasthan                                | Ps      |
| region_de_murcia                         | spain                            | region_de_murcia                         | Ps      |
| reunion                                  | africa                           | reunion                                  | Ps      |
| rhone                                    | rhone_alpes                      | rhone                                    | Ps      |
| rhone_alpes                              | france                           | rhone_alpes                              | Ps      |
| riau                                     | indonesia                        | riau                                     | Ps      |
| riau_islands                             | indonesia                        | riau_islands                             | Ps      |
| richmond                                 | virginia                         | richmond                                 | Ps      |
| richmond_tweed                           | new_south_wales                  | richmond_tweed                           | Ps      |
| rio-de-janeiro                           | southeast                        | rio-de-janeiro                           | Ps      |
| rio-grande-do-norte                      | northeast                        | rio-grande-do-norte                      | Ps      |
| rio-grande-do-sul                        | south                            | rio-grande-do-sul                        | Ps      |
| rio_negro                                | argentina                        | rio_negro                                | Ps      |
| riverside                                | california                       | riverside                                | Ps      |
| rivne_oblast                             | ukraine                          | rivne_oblast                             | Ps      |
| rogaland                                 | norway                           | rogaland                                 | Ps      |
| rondonia                                 | north                            | rondonia                                 | Ps      |
| roraima                                  | north                            | roraima                                  | Ps      |
| rostov_oblast                            | southern_federal_district        | rostov_oblast                            | Ps      |
| russia                                   |                                  | russia                                   | Ps      |
| rwanda                                   | africa                           | rwanda                                   | Ps      |
| ryazan_oblast                            | central_federal_district         | ryazan_oblast                            | Ps      |
| sacramento                               | california                       | sacramento                               | Ps      |
| saguenay_lac_saint_jean                  | quebec                           | saguenay_lac_saint_jean                  | Ps      |
| saint_barthelemy                         | central-america                  | saint_barthelemy                         | Ps      |
| saint_gallen                             | switzerland                      | saint_gallen                             | Ps      |
| saint_helena_ascension_tristan_da_cunha  | africa                           | saint_helena_ascension_tristan_da_cunha  | Ps      |
| saint_kitts_and_nevis                    | central-america                  | saint_kitts_and_nevis                    | Ps      |
| saint_lucia                              | central-america                  | saint_lucia                              | Ps      |
| saint_martin                             | central-america                  | saint_martin                             | Ps      |
| saint_petersburg                         | northwestern_federal_district    | saint_petersburg                         | Ps      |
| saint_pierre_et_miquelon                 | north-america                    | saint_pierre_et_miquelon                 | Ps      |
| saint_vincent_and_the_grenadines         | central-america                  | saint_vincent_and_the_grenadines         | Ps      |
| sakha_republic                           | far_eastern_federal_district     | sakha_republic                           | Ps      |
| sakhalin_oblast                          | far_eastern_federal_district     | sakhalin_oblast                          | Ps      |
| salamanca                                | castilla_y_leon                  | salamanca                                | Ps      |
| salem                                    | virginia                         | salem                                    | Ps      |
| salta                                    | argentina                        | salta                                    | Ps      |
| salzburg                                 | austria                          | salzburg                                 | Ps      |
| samara_oblast                            | volga_federal_district           | samara_oblast                            | Ps      |
| samoa                                    | oceania                          | samoa                                    | Ps      |
| san_benito                               | california                       | san_benito                               | Ps      |
| san_bernardino                           | california                       | san_bernardino                           | Ps      |
| san_diego                                | california                       | san_diego                                | Ps      |
| san_francisco                            | california                       | san_francisco                            | Ps      |
| san_joaquin                              | california                       | san_joaquin                              | Ps      |
| san_juan                                 | argentina                        | san_juan                                 | Ps      |
| san_luis                                 | argentina                        | san_luis                                 | Ps      |
| san_luis_obispo                          | california                       | san_luis_obispo                          | Ps      |
| san_luis_potosi                          | mexico                           | san_luis_potosi                          | Ps      |
| san_marino                               | europe                           | san_marino                               | Ps      |
| san_mateo                                | california                       | san_mateo                                | Ps      |
| santa-catarina                           | south                            | santa-catarina                           | Ps      |
| santa_barbara                            | california                       | santa_barbara                            | Ps      |
| santa_clara                              | california                       | santa_clara                              | Ps      |
| santa_cruz_de_tenerife                   | canarias                         | santa_cruz_de_tenerife                   | Ps      |
| santa_fe                                 | argentina                        | santa_fe                                 | Ps      |
| santarem                                 | portugal                         | santarem                                 | Ps      |
| santiago                                 | chile                            | santiago                                 | Ps      |
| santiago_del_estero                      | argentina                        | santiago_del_estero                      | Ps      |
| sao-paulo                                | southeast                        | sao-paulo                                | Ps      |
| sao_tome_and_principe                    | africa                           | sao_tome_and_principe                    | Ps      |
| saone_et_loire                           | bourgogne                        | saone_et_loire                           | Ps      |
| saratov_oblast                           | volga_federal_district           | saratov_oblast                           | Ps      |
| sardegna                                 | italy                            | sardegna                                 | Ps      |
| sarthe                                   | pays_de_la_loire                 | sarthe                                   | Ps      |
| saskatchewan                             | canada                           | saskatchewan                             | Ps      |
| satakunta                                | finland                          | satakunta                                | Ps      |
| saudi_arabia                             | asia                             | saudi_arabia                             | Ps      |
| savoie                                   | rhone_alpes                      | savoie                                   | Ps      |
| schaffhausen                             | switzerland                      | schaffhausen                             | Ps      |
| schwyz                                   | switzerland                      | schwyz                                   | Ps      |
| seas                                     | europe                           | seas                                     |         |
| segovia                                  | castilla_y_leon                  | segovia                                  | Ps      |
| seine_et_marne                           | ile_de_france                    | seine_et_marne                           | Ps      |
| seine_maritime                           | haute_normandie                  | seine_maritime                           | Ps      |
| seine_saint_denis                        | ile_de_france                    | seine_saint_denis                        | Ps      |
| senegal                                  | africa                           | senegal                                  | Ps      |
| sergipe                                  | northeast                        | sergipe                                  | Ps      |
| setubal                                  | portugal                         | setubal                                  | Ps      |
| sevilla                                  | andalucia                        | sevilla                                  | Ps      |
| seychelles                               | africa                           | seychelles                               | Ps      |
| shaanxi                                  | china                            | shaanxi                                  | Ps      |
| shandong                                 | china                            | shandong                                 | Ps      |
| shanghai                                 | china                            | shanghai                                 | Ps      |
| shanxi                                   | china                            | shanxi                                   | Ps      |
| shasta                                   | california                       | shasta                                   | Ps      |
| shikoku                                  | japan                            | shikoku                                  | Ps      |
| siberian_federal_district                | russia                           | siberian_federal_district                | Ps      |
| sichuan                                  | china                            | sichuan                                  | Ps      |
| sicilia                                  | italy                            | sicilia                                  | Ps      |
| sierra                                   | california                       | sierra                                   | Ps      |
| sikkim                                   | india                            | sikkim                                   | Ps      |
| sinaloa                                  | mexico                           | sinaloa                                  | Ps      |
| singapore                                | asia                             | singapore                                | Ps      |
| sint_maarten                             | central-america                  | sint_maarten                             | Ps      |
| siskiyou                                 | california                       | siskiyou                                 | Ps      |
| sjaelland                                | denmark                          | sjaelland                                | Ps      |
| skane                                    | sweden                           | skane                                    | Ps      |
| slaskie                                  | poland                           | slaskie                                  | Ps      |
| slovakia                                 | europe                           | slovakia                                 | Ps      |
| smolensk_oblast                          | central_federal_district         | smolensk_oblast                          | Ps      |
| soccsksargen                             | philippines                      | soccsksargen                             | Ps      |
| sodermanland                             | sweden                           | sodermanland                             | Ps      |
| sofala                                   | mozambique                       | sofala                                   | Ps      |
| sogn_og_fjordane                         | norway                           | sogn_og_fjordane                         | Ps      |
| solano                                   | california                       | solano                                   | Ps      |
| solomon_islands                          | oceania                          | solomon_islands                          | Ps      |
| solothurn                                | switzerland                      | solothurn                                | Ps      |
| somme                                    | picardie                         | somme                                    | Ps      |
| sonoma                                   | california                       | sonoma                                   | Ps      |
| sonora                                   | mexico                           | sonora                                   | Ps      |
| soria                                    | castilla_y_leon                  | soria                                    | Ps      |
| south                                    | brazil                           | south                                    |         |
| south-america                            |                                  | south-america                            | Ps      |
| south_africa                             | africa                           | south_africa                             | Ps      |
| south_africa_north_west                  | south_africa                     | south_africa_north_west                  | Ps      |
| south_australia                          | australia                        | south_australia                          | Ps      |
| south_east_region                        | new_south_wales                  | south_east_region                        | Ps      |
| south_georgia_and_south_sandwich         | south-america                    | south_georgia_and_south_sandwich         | Ps      |
| south_kalimantan                         | indonesia                        | south_kalimantan                         | Ps      |
| south_karelia                            | finland                          | south_karelia                            | Ps      |
| south_ostrobothnia                       | finland                          | south_ostrobothnia                       | Ps      |
| south_papua                              | indonesia                        | south_papua                              | Ps      |
| south_savo                               | finland                          | south_savo                               | Ps      |
| south_sudan                              | africa                           | south_sudan                              | Ps      |
| south_sulawesi                           | indonesia                        | south_sulawesi                           | Ps      |
| south_sumatra                            | indonesia                        | south_sumatra                            | Ps      |
| southeast                                | brazil                           | southeast                                |         |
| southeast_sulawesi                       | indonesia                        | southeast_sulawesi                       | Ps      |
| southeastern_anatolia                    | turkey                           | southeastern_anatolia                    | Ps      |
| southern_federal_district                | russia                           | southern_federal_district                | Ps      |
| southern_federal_district_sevastopol     | southern_federal_district        | southern_federal_district_sevastopol     | Ps      |
| southern_highlands                       | tanzania                         | southern_highlands                       | Ps      |
| southern_tier                            | new-york                         | southern_tier                            | Ps      |
| southwest_finland                        | finland                          | southwest_finland                        | Ps      |
| southwest_papua                          | indonesia                        | southwest_papua                          | Ps      |
| southwestern_ontario                     | ontario                          | southwestern_ontario                     | Ps      |
| spain                                    | europe                           | spain                                    | Ps      |
| spain_la_rioja                           | spain                            | spain_la_rioja                           | Ps      |
| stanislaus                               | california                       | stanislaus                               | Ps      |
| state_of_mexico                          | mexico                           | state_of_mexico                          | Ps      |
| stavropol_krai                           | north_caucasian_federal_district | stavropol_krai                           | Ps      |
| steiermark                               | austria                          | steiermark                               | Ps      |
| stockholm                                | sweden                           | stockholm                                | Ps      |
| stredocesky                              | czech_republic                   | stredocesky                              | Ps      |
| sudan                                    | africa                           | sudan                                    | Ps      |
| sumy_oblast                              | ukraine                          | sumy_oblast                              | Ps      |
| suncoast                                 | florida                          | suncoast                                 | Ps      |
| suriname                                 | south-america                    | suriname                                 | Ps      |
| sutter                                   | california                       | sutter                                   | Ps      |
| svalbard                                 | norway                           | svalbard                                 | Ps      |
| sverdlovsk_oblast                        | ural_federal_district            | sverdlovsk_oblast                        | Ps      |
| swaziland                                | africa                           | swaziland                                | Ps      |
| sweden                                   | europe                           | sweden                                   | Ps      |
| swietokrzyskie                           | poland                           | swietokrzyskie                           | Ps      |
| switzerland                              | europe                           | switzerland                              | Ps      |
| switzerland_jura                         | switzerland                      | switzerland_jura                         | Ps      |
| syddanmark                               | denmark                          | syddanmark                               | Ps      |
| sydney_surrounds                         | new_south_wales                  | sydney_surrounds                         | Ps      |
| tabasco                                  | mexico                           | tabasco                                  | Ps      |
| tamaulipas                               | mexico                           | tamaulipas                               | Ps      |
| tambov_oblast                            | central_federal_district         | tambov_oblast                            | Ps      |
| tamil_nadu                               | india                            | tamil_nadu                               | Ps      |
| tanzania                                 | africa                           | tanzania                                 | Ps      |
| tanzania_central                         | tanzania                         | tanzania_central                         | Ps      |
| tanzania_lake                            | tanzania                         | tanzania_lake                            | Ps      |
| tanzania_northern                        | tanzania                         | tanzania_northern                        | Ps      |
| tanzania_western                         | tanzania                         | tanzania_western                         | Ps      |
| tarapaca                                 | chile                            | tarapaca                                 | Ps      |
| tarn                                     | midi_pyrenees                    | tarn                                     | Ps      |
| tarn_et_garonne                          | midi_pyrenees                    | tarn_et_garonne                          | Ps      |
| tarragona                                | catalunya                        | tarragona                                | Ps      |
| tasmania                                 | australia                        | tasmania                                 | Ps      |
| tatarstan_republic                       | volga_federal_district           | tatarstan_republic                       | Ps      |
| tehama                                   | california                       | tehama                                   | Ps      |
| telangana                                | india                            | telangana                                | Ps      |
| telemark                                 | norway                           | telemark                                 | Ps      |
| ternopil_oblast                          | ukraine                          | ternopil_oblast                          | Ps      |
| territoire_de_belfort                    | franche_comte                    | territoire_de_belfort                    | Ps      |
| teruel                                   | aragon                           | teruel                                   | Ps      |
| tete                                     | mozambique                       | tete                                     | Ps      |
| texas                                    | us-south                         | texas                                    | Ps      |
| texas_central                            | texas                            | texas_central                            | Ps      |
| texas_north                              | texas                            | texas_north                              | Ps      |
| texas_northwest                          | texas                            | texas_northwest                          | Ps      |
| texas_south                              | texas                            | texas_south                              | Ps      |
| texas_southeast                          | texas                            | texas_southeast                          | Ps      |
| texas_west                               | texas                            | texas_west                               | Ps      |
| the_riverina                             | new_south_wales                  | the_riverina                             | Ps      |
| thurgau                                  | switzerland                      | thurgau                                  | Ps      |
| tianjin                                  | china                            | tianjin                                  | Ps      |
| tibet                                    | china                            | tibet                                    | Ps      |
| ticino                                   | switzerland                      | ticino                                   | Ps      |
| tierra_del_fuego                         | argentina                        | tierra_del_fuego                         | Ps      |
| tirol                                    | austria                          | tirol                                    | Ps      |
| tlaxcala                                 | mexico                           | tlaxcala                                 | Ps      |
| tocantins                                | north                            | tocantins                                | Ps      |
| togo                                     | africa                           | togo                                     | Ps      |
| tohoku                                   | japan                            | tohoku                                   | Ps      |
| tokelau                                  | oceania                          | tokelau                                  | Ps      |
| toledo                                   | castilla_la_mancha               | toledo                                   | Ps      |
| tomsk_oblast                             | siberian_federal_district        | tomsk_oblast                             | Ps      |
| tonga                                    | oceania                          | tonga                                    | Ps      |
| toscana                                  | italy                            | toscana                                  | Ps      |
| treasure_coast                           | florida                          | treasure_coast                           | Ps      |
| trenciansky                              | slovakia                         | trenciansky                              | Ps      |
| trentino_alto_adige                      | italy                            | trentino_alto_adige                      | Ps      |
| trinidad_and_tobago                      | central-america                  | trinidad_and_tobago                      | Ps      |
| trinity                                  | california                       | trinity                                  | Ps      |
| tripura                                  | india                            | tripura                                  | Ps      |
| trnavsky                                 | slovakia                         | trnavsky                                 | Ps      |
| troendelag                               | norway                           | troendelag                               | Ps      |
| troms                                    | norway                           | troms                                    | Ps      |
| tucuman                                  | argentina                        | tucuman                                  | Ps      |
| tula_oblast                              | central_federal_district         | tula_oblast                              | Ps      |
| tulare                                   | california                       | tulare                                   | Ps      |
| tunisia                                  | africa                           | tunisia                                  | Ps      |
| tuolumne                                 | california                       | tuolumne                                 | Ps      |
| turkey                                   | europe                           | turkey                                   | Ps      |
| turks_and_caicos_islands                 | central-america                  | turks_and_caicos_islands                 | Ps      |
| tuva_republic                            | siberian_federal_district        | tuva_republic                            | Ps      |
| tuvalu                                   | oceania                          | tuvalu                                   | Ps      |
| tver_oblast                              | central_federal_district         | tver_oblast                              | Ps      |
| tyumen_oblast                            | ural_federal_district            | tyumen_oblast                            | Ps      |
| udmurt_republic                          | volga_federal_district           | udmurt_republic                          | Ps      |
| uganda                                   | africa                           | uganda                                   | Ps      |
| uganda_central                           | uganda                           | uganda_central                           | Ps      |
| uganda_eastern                           | uganda                           | uganda_eastern                           | Ps      |
| uganda_northern                          | uganda                           | uganda_northern                          | Ps      |
| uganda_western                           | uganda                           | uganda_western                           | Ps      |
| ukraine                                  | europe                           | ukraine                                  | Ps      |
| ukraine_sevastopol                       | ukraine                          | ukraine_sevastopol                       | Ps      |
| ulyanovsk_oblast                         | volga_federal_district           | ulyanovsk_oblast                         | Ps      |
| umbria                                   | italy                            | umbria                                   | Ps      |
| united_arab_emirates                     | asia                             | united_arab_emirates                     | Ps      |
| united_kingdom                           | europe                           | united_kingdom                           | Ps      |
| united_states_virgin_islands             | central-america                  | united_states_virgin_islands             | Ps      |
| upper_peninsula                          | michigan                         | upper_peninsula                          | Ps      |
| uppsala                                  | sweden                           | uppsala                                  | Ps      |
| ural_federal_district                    | russia                           | ural_federal_district                    | Ps      |
| uri                                      | switzerland                      | uri                                      | Ps      |
| us-midwest                               | north-america                    | us-midwest                               | Ps      |
| us-northeast                             | north-america                    | us-northeast                             | Ps      |
| us-south                                 | north-america                    | us-south                                 | Ps      |
| us-west                                  | north-america                    | us-west                                  | Ps      |
| usa_virgin_islands                       | central-america                  | usa_virgin_islands                       | Ps      |
| ustecky                                  | czech_republic                   | ustecky                                  | Ps      |
| utrecht                                  | netherlands                      | utrecht                                  | Ps      |
| uttar_pradesh                            | india                            | uttar_pradesh                            | Ps      |
| uttarakhand                              | india                            | uttarakhand                              | Ps      |
| uusimaa                                  | finland                          | uusimaa                                  | Ps      |
| val_d_oise                               | ile_de_france                    | val_d_oise                               | Ps      |
| val_de_marne                             | ile_de_france                    | val_de_marne                             | Ps      |
| valais                                   | switzerland                      | valais                                   | Ps      |
| valencia                                 | comunitat_valenciana             | valencia                                 | Ps      |
| valladolid                               | castilla_y_leon                  | valladolid                               | Ps      |
| valle_aosta                              | italy                            | valle_aosta                              | Ps      |
| valparaiso                               | chile                            | valparaiso                               | Ps      |
| vanuatu                                  | oceania                          | vanuatu                                  | Ps      |
| var                                      | provence_alpes_cote_d_azur       | var                                      | Ps      |
| varmland                                 | sweden                           | varmland                                 | Ps      |
| vasterbotten                             | sweden                           | vasterbotten                             | Ps      |
| vasternorrland                           | sweden                           | vasternorrland                           | Ps      |
| vastmanland                              | sweden                           | vastmanland                              | Ps      |
| vastra_gotaland                          | sweden                           | vastra_gotaland                          | Ps      |
| vatican_city                             | europe                           | vatican_city                             | Ps      |
| vaucluse                                 | provence_alpes_cote_d_azur       | vaucluse                                 | Ps      |
| vaud                                     | switzerland                      | vaud                                     | Ps      |
| vendee                                   | pays_de_la_loire                 | vendee                                   | Ps      |
| veneto                                   | italy                            | veneto                                   | Ps      |
| venezuela                                | south-america                    | venezuela                                | Ps      |
| ventura                                  | california                       | ventura                                  | Ps      |
| veracruz                                 | mexico                           | veracruz                                 | Ps      |
| vest-agder                               | norway                           | vest-agder                               | Ps      |
| vestfold                                 | norway                           | vestfold                                 | Ps      |
| viana_do_castelo                         | portugal                         | viana_do_castelo                         | Ps      |
| victoria                                 | australia                        | victoria                                 | Ps      |
| vienne                                   | poitou_charentes                 | vienne                                   | Ps      |
| vila_real                                | portugal                         | vila_real                                | Ps      |
| vinnytsia_oblast                         | ukraine                          | vinnytsia_oblast                         | Ps      |
| virginia                                 | us-south                         | virginia                                 | Ps      |
| viseu                                    | portugal                         | viseu                                    | Ps      |
| vizcaya                                  | euskadi                          | vizcaya                                  | Ps      |
| vladimir_oblast                          | central_federal_district         | vladimir_oblast                          | Ps      |
| volga_federal_district                   | russia                           | volga_federal_district                   | Ps      |
| volgograd_oblast                         | southern_federal_district        | volgograd_oblast                         | Ps      |
| vologda_oblast                           | northwestern_federal_district    | vologda_oblast                           | Ps      |
| volyn_oblast                             | ukraine                          | volyn_oblast                             | Ps      |
| vorarlberg                               | austria                          | vorarlberg                               | Ps      |
| voronezh_oblast                          | central_federal_district         | voronezh_oblast                          | Ps      |
| vosges                                   | lorraine                         | vosges                                   | Ps      |
| vysocina                                 | czech_republic                   | vysocina                                 | Ps      |
| wallis_et_futuna                         | oceania                          | wallis_et_futuna                         | Ps      |
| wallonia_french_community                | belgium                          | wallonia_french_community                | Ps      |
| wallonia_german_community                | belgium                          | wallonia_german_community                | Ps      |
| warminsko_mazurskie                      | poland                           | warminsko_mazurskie                      | Ps      |
| west_bengal                              | india                            | west_bengal                              | Ps      |
| west_flanders                            | flanders                         | west_flanders                            | Ps      |
| west_java                                | indonesia                        | west_java                                | Ps      |
| west_kalimantan                          | indonesia                        | west_kalimantan                          | Ps      |
| west_midlands                            | england                          | west_midlands                            | Ps      |
| west_nusa_tenggara                       | indonesia                        | west_nusa_tenggara                       | Ps      |
| west_papua                               | indonesia                        | west_papua                               | Ps      |
| west_sulawesi                            | indonesia                        | west_sulawesi                            | Ps      |
| west_sumatra                             | indonesia                        | west_sumatra                             | Ps      |
| western_australia                        | australia                        | western_australia                        | Ps      |
| western_cape                             | south_africa                     | western_cape                             | Ps      |
| western_new_york                         | new-york                         | western_new_york                         | Ps      |
| western_sahara                           | africa                           | western_sahara                           | Ps      |
| western_visayas                          | philippines                      | western_visayas                          | Ps      |
| wielkopolskie                            | poland                           | wielkopolskie                            | Ps      |
| wien                                     | austria                          | wien                                     | Ps      |
| wytheville                               | virginia                         | wytheville                               | Ps      |
| xinjiang                                 | china                            | xinjiang                                 | Ps      |
| yamalo_nenets_autonomous_okrug           | ural_federal_district            | yamalo_nenets_autonomous_okrug           | Ps      |
| yaroslavl_oblast                         | central_federal_district         | yaroslavl_oblast                         | Ps      |
| yogyakarta                               | indonesia                        | yogyakarta                               | Ps      |
| yolo                                     | california                       | yolo                                     | Ps      |
| yonne                                    | bourgogne                        | yonne                                    | Ps      |
| yorkshire_and_the_humber                 | england                          | yorkshire_and_the_humber                 | Ps      |
| yuba                                     | california                       | yuba                                     | Ps      |
| yucatan                                  | mexico                           | yucatan                                  | Ps      |
| yukon                                    | canada                           | yukon                                    | Ps      |
| yunnan                                   | china                            | yunnan                                   | Ps      |
| yvelines                                 | ile_de_france                    | yvelines                                 | Ps      |
| zabaykalsky_krai                         | siberian_federal_district        | zabaykalsky_krai                         | Ps      |
| zacatecas                                | mexico                           | zacatecas                                | Ps      |
| zachodniopomorskie                       | poland                           | zachodniopomorskie                       | Ps      |
| zakarpattia_oblast                       | ukraine                          | zakarpattia_oblast                       | Ps      |
| zambezia                                 | mozambique                       | zambezia                                 | Ps      |
| zambia                                   | africa                           | zambia                                   | Ps      |
| zamboanga_peninsula                      | philippines                      | zamboanga_peninsula                      | Ps      |
| zamora                                   | castilla_y_leon                  | zamora                                   | Ps      |
| zanzibar                                 | tanzania                         | zanzibar                                 | Ps      |
| zaporizhia_oblast                        | ukraine                          | zaporizhia_oblast                        | Ps      |
| zaragoza                                 | aragon                           | zaragoza                                 | Ps      |
| zeeland                                  | netherlands                      | zeeland                                  | Ps      |
| zhejiang                                 | china                            | zhejiang                                 | Ps      |
| zhytomyr_oblast                          | ukraine                          | zhytomyr_oblast                          | Ps      |
| zilinsky                                 | slovakia                         | zilinsky                                 | Ps      |
| zimbabwe                                 | africa                           | zimbabwe                                 | Ps      |
| zlinsky                                  | czech_republic                   | zlinsky                                  | Ps      |
| zug                                      | switzerland                      | zug                                      | Ps      |
| zuid_holland                             | netherlands                      | zuid_holland                             | Ps      |
| zurich                                   | switzerland                      | zurich                                   | Ps      |

</details>

<details>
<summary><strong>BBBike Extracts (bbbike.org)</strong></summary>

| SHORT NAME       | IS IN | LONG NAME        | FORMATS |
|------------------|-------|------------------|---------|
| Aachen           |       | Aachen           | gGPpS   |
| Aarhus           |       | Aarhus           | gGPpS   |
| Adelaide         |       | Adelaide         | gGPpS   |
| Albuquerque      |       | Albuquerque      | gGPpS   |
| Alexandria       |       | Alexandria       | gGPpS   |
| Amsterdam        |       | Amsterdam        | gGPpS   |
| Antwerpen        |       | Antwerpen        | gGPpS   |
| Arnhem           |       | Arnhem           | gGPpS   |
| Auckland         |       | Auckland         | gGPpS   |
| Augsburg         |       | Augsburg         | gGPpS   |
| Austin           |       | Austin           | gGPpS   |
| Baghdad          |       | Baghdad          | gGPpS   |
| Baku             |       | Baku             | gGPpS   |
| Balaton          |       | Balaton          | gGPpS   |
| Bamberg          |       | Bamberg          | gGPpS   |
| Bangkok          |       | Bangkok          | gGPpS   |
| Barcelona        |       | Barcelona        | gGPpS   |
| Basel            |       | Basel            | gGPpS   |
| Beijing          |       | Beijing          | gGPpS   |
| Beirut           |       | Beirut           | gGPpS   |
| Berkeley         |       | Berkeley         | gGPpS   |
| Berlin           |       | Berlin           | gGPpS   |
| Bern             |       | Bern             | gGPpS   |
| Bielefeld        |       | Bielefeld        | gGPpS   |
| Birmingham       |       | Birmingham       | gGPpS   |
| Bochum           |       | Bochum           | gGPpS   |
| Bogota           |       | Bogota           | gGPpS   |
| Bombay           |       | Bombay           | gGPpS   |
| Bonn             |       | Bonn             | gGPpS   |
| Bordeaux         |       | Bordeaux         | gGPpS   |
| Boulder          |       | Boulder          | gGPpS   |
| BrandenburgHavel |       | BrandenburgHavel | gGPpS   |
| Braunschweig     |       | Braunschweig     | gGPpS   |
| Bremen           |       | Bremen           | gGPpS   |
| Bremerhaven      |       | Bremerhaven      | gGPpS   |
| Brisbane         |       | Brisbane         | gGPpS   |
| Bristol          |       | Bristol          | gGPpS   |
| Brno             |       | Brno             | gGPpS   |
| Bruegge          |       | Bruegge          | gGPpS   |
| Bruessel         |       | Bruessel         | gGPpS   |
| Budapest         |       | Budapest         | gGPpS   |
| BuenosAires      |       | BuenosAires      | gGPpS   |
| Cairo            |       | Cairo            | gGPpS   |
| Calgary          |       | Calgary          | gGPpS   |
| Cambridge        |       | Cambridge        | gGPpS   |
| CambridgeMa      |       | CambridgeMa      | gGPpS   |
| Canberra         |       | Canberra         | gGPpS   |
| CapeTown         |       | CapeTown         | gGPpS   |
| Chemnitz         |       | Chemnitz         | gGPpS   |
| Chicago          |       | Chicago          | gGPpS   |
| ClermontFerrand  |       | ClermontFerrand  | gGPpS   |
| Colmar           |       | Colmar           | gGPpS   |
| Copenhagen       |       | Copenhagen       | gGPpS   |
| Cork             |       | Cork             | gGPpS   |
| Corsica          |       | Corsica          | gGPpS   |
| Corvallis        |       | Corvallis        | gGPpS   |
| Cottbus          |       | Cottbus          | gGPpS   |
| Cracow           |       | Cracow           | gGPpS   |
| CraterLake       |       | CraterLake       | gGPpS   |
| Curitiba         |       | Curitiba         | gGPpS   |
| Cusco            |       | Cusco            | gGPpS   |
| Dallas           |       | Dallas           | gGPpS   |
| Damaskus         |       | Damaskus         | gGPpS   |
| Darmstadt        |       | Darmstadt        | gGPpS   |
| Davis            |       | Davis            | gGPpS   |
| DenHaag          |       | DenHaag          | gGPpS   |
| Denver           |       | Denver           | gGPpS   |
| Dessau           |       | Dessau           | gGPpS   |
| Dortmund         |       | Dortmund         | gGPpS   |
| Dresden          |       | Dresden          | gGPpS   |
| Dublin           |       | Dublin           | gGPpS   |
| Duesseldorf      |       | Duesseldorf      | gGPpS   |
| Duisburg         |       | Duisburg         | gGPpS   |
| Edinburgh        |       | Edinburgh        | gGPpS   |
| Eindhoven        |       | Eindhoven        | gGPpS   |
| Emden            |       | Emden            | gGPpS   |
| Erfurt           |       | Erfurt           | gGPpS   |
| Erlangen         |       | Erlangen         | gGPpS   |
| Eugene           |       | Eugene           | gGPpS   |
| Flensburg        |       | Flensburg        | gGPpS   |
| FortCollins      |       | FortCollins      | gGPpS   |
| Frankfurt        |       | Frankfurt        | gGPpS   |
| FrankfurtOder    |       | FrankfurtOder    | gGPpS   |
| Freiburg         |       | Freiburg         | gGPpS   |
| Gdansk           |       | Gdansk           | gGPpS   |
| Genf             |       | Genf             | gGPpS   |
| Gent             |       | Gent             | gGPpS   |
| Gera             |       | Gera             | gGPpS   |
| Glasgow          |       | Glasgow          | gGPpS   |
| Gliwice          |       | Gliwice          | gGPpS   |
| Goerlitz         |       | Goerlitz         | gGPpS   |
| Goeteborg        |       | Goeteborg        | gGPpS   |
| Goettingen       |       | Goettingen       | gGPpS   |
| Graz             |       | Graz             | gGPpS   |
| Groningen        |       | Groningen        | gGPpS   |
| Halifax          |       | Halifax          | gGPpS   |
| Halle            |       | Halle            | gGPpS   |
| Hamburg          |       | Hamburg          | gGPpS   |
| Hamm             |       | Hamm             | gGPpS   |
| Hannover         |       | Hannover         | gGPpS   |
| Heilbronn        |       | Heilbronn        | gGPpS   |
| Helsinki         |       | Helsinki         | gGPpS   |
| Hertogenbosch    |       | Hertogenbosch    | gGPpS   |
| Huntsville       |       | Huntsville       | gGPpS   |
| Innsbruck        |       | Innsbruck        | gGPpS   |
| Istanbul         |       | Istanbul         | gGPpS   |
| Jena             |       | Jena             | gGPpS   |
| Jerusalem        |       | Jerusalem        | gGPpS   |
| Johannesburg     |       | Johannesburg     | gGPpS   |
| Kaiserslautern   |       | Kaiserslautern   | gGPpS   |
| Karlsruhe        |       | Karlsruhe        | gGPpS   |
| Kassel           |       | Kassel           | gGPpS   |
| Katowice         |       | Katowice         | gGPpS   |
| Kaunas           |       | Kaunas           | gGPpS   |
| Kiel             |       | Kiel             | gGPpS   |
| Kiew             |       | Kiew             | gGPpS   |
| Koblenz          |       | Koblenz          | gGPpS   |
| Koeln            |       | Koeln            | gGPpS   |
| Konstanz         |       | Konstanz         | gGPpS   |
| LaPaz            |       | LaPaz            | gGPpS   |
| LaPlata          |       | LaPlata          | gGPpS   |
| LakeGarda        |       | LakeGarda        | gGPpS   |
| Lausanne         |       | Lausanne         | gGPpS   |
| Leeds            |       | Leeds            | gGPpS   |
| Leipzig          |       | Leipzig          | gGPpS   |
| Lima             |       | Lima             | gGPpS   |
| Linz             |       | Linz             | gGPpS   |
| Lisbon           |       | Lisbon           | gGPpS   |
| Liverpool        |       | Liverpool        | gGPpS   |
| Ljubljana        |       | Ljubljana        | gGPpS   |
| Lodz             |       | Lodz             | gGPpS   |
| London           |       | London           | gGPpS   |
| LosAngeles       |       | LosAngeles       | gGPpS   |
| Luebeck          |       | Luebeck          | gGPpS   |
| Luxemburg        |       | Luxemburg        | gGPpS   |
| Lyon             |       | Lyon             | gGPpS   |
| Maastricht       |       | Maastricht       | gGPpS   |
| Madison          |       | Madison          | gGPpS   |
| Madrid           |       | Madrid           | gGPpS   |
| Magdeburg        |       | Magdeburg        | gGPpS   |
| Mainz            |       | Mainz            | gGPpS   |
| Malmoe           |       | Malmoe           | gGPpS   |
| Manchester       |       | Manchester       | gGPpS   |
| Mannheim         |       | Mannheim         | gGPpS   |
| Marseille        |       | Marseille        | gGPpS   |
| Melbourne        |       | Melbourne        | gGPpS   |
| Memphis          |       | Memphis          | gGPpS   |
| MexicoCity       |       | MexicoCity       | gGPpS   |
| Miami            |       | Miami            | gGPpS   |
| Minsk            |       | Minsk            | gGPpS   |
| Moenchengladbach |       | Moenchengladbach | gGPpS   |
| Montevideo       |       | Montevideo       | gGPpS   |
| Montpellier      |       | Montpellier      | gGPpS   |
| Montreal         |       | Montreal         | gGPpS   |
| Moscow           |       | Moscow           | gGPpS   |
| Muenchen         |       | Muenchen         | gGPpS   |
| Muenster         |       | Muenster         | gGPpS   |
| NewDelhi         |       | NewDelhi         | gGPpS   |
| NewOrleans       |       | NewOrleans       | gGPpS   |
| NewYork          |       | NewYork          | gGPpS   |
| Nuernberg        |       | Nuernberg        | gGPpS   |
| Oldenburg        |       | Oldenburg        | gGPpS   |
| Oranienburg      |       | Oranienburg      | gGPpS   |
| Orlando          |       | Orlando          | gGPpS   |
| Oslo             |       | Oslo             | gGPpS   |
| Osnabrueck       |       | Osnabrueck       | gGPpS   |
| Ostrava          |       | Ostrava          | gGPpS   |
| Ottawa           |       | Ottawa           | gGPpS   |
| Paderborn        |       | Paderborn        | gGPpS   |
| Palma            |       | Palma            | gGPpS   |
| PaloAlto         |       | PaloAlto         | gGPpS   |
| Paris            |       | Paris            | gGPpS   |
| Perth            |       | Perth            | gGPpS   |
| Philadelphia     |       | Philadelphia     | gGPpS   |
| PhnomPenh        |       | PhnomPenh        | gGPpS   |
| Portland         |       | Portland         | gGPpS   |
| PortlandME       |       | PortlandME       | gGPpS   |
| Porto            |       | Porto            | gGPpS   |
| PortoAlegre      |       | PortoAlegre      | gGPpS   |
| Potsdam          |       | Potsdam          | gGPpS   |
| Poznan           |       | Poznan           | gGPpS   |
| Prag             |       | Prag             | gGPpS   |
| Providence       |       | Providence       | gGPpS   |
| Regensburg       |       | Regensburg       | gGPpS   |
| Riga             |       | Riga             | gGPpS   |
| RiodeJaneiro     |       | RiodeJaneiro     | gGPpS   |
| Rostock          |       | Rostock          | gGPpS   |
| Rotterdam        |       | Rotterdam        | gGPpS   |
| Ruegen           |       | Ruegen           | gGPpS   |
| Saarbruecken     |       | Saarbruecken     | gGPpS   |
| Sacramento       |       | Sacramento       | gGPpS   |
| Saigon           |       | Saigon           | gGPpS   |
| Salzburg         |       | Salzburg         | gGPpS   |
| SanFrancisco     |       | SanFrancisco     | gGPpS   |
| SanJose          |       | SanJose          | gGPpS   |
| SanktPetersburg  |       | SanktPetersburg  | gGPpS   |
| SantaBarbara     |       | SantaBarbara     | gGPpS   |
| SantaCruz        |       | SantaCruz        | gGPpS   |
| Santiago         |       | Santiago         | gGPpS   |
| Sarajewo         |       | Sarajewo         | gGPpS   |
| Schwerin         |       | Schwerin         | gGPpS   |
| Seattle          |       | Seattle          | gGPpS   |
| Seoul            |       | Seoul            | gGPpS   |
| Sheffield        |       | Sheffield        | gGPpS   |
| Singapore        |       | Singapore        | gGPpS   |
| Sofia            |       | Sofia            | gGPpS   |
| Stockholm        |       | Stockholm        | gGPpS   |
| Stockton         |       | Stockton         | gGPpS   |
| Strassburg       |       | Strassburg       | gGPpS   |
| Stuttgart        |       | Stuttgart        | gGPpS   |
| Sucre            |       | Sucre            | gGPpS   |
| Sydney           |       | Sydney           | gGPpS   |
| Szczecin         |       | Szczecin         | gGPpS   |
| Tallinn          |       | Tallinn          | gGPpS   |
| Tehran           |       | Tehran           | gGPpS   |
| Tilburg          |       | Tilburg          | gGPpS   |
| Tokyo            |       | Tokyo            | gGPpS   |
| Toronto          |       | Toronto          | gGPpS   |
| Toulouse         |       | Toulouse         | gGPpS   |
| Trondheim        |       | Trondheim        | gGPpS   |
| Tucson           |       | Tucson           | gGPpS   |
| Turin            |       | Turin            | gGPpS   |
| UlanBator        |       | UlanBator        | gGPpS   |
| Ulm              |       | Ulm              | gGPpS   |
| Usedom           |       | Usedom           | gGPpS   |
| Utrecht          |       | Utrecht          | gGPpS   |
| Vancouver        |       | Vancouver        | gGPpS   |
| Victoria         |       | Victoria         | gGPpS   |
| WarenMueritz     |       | WarenMueritz     | gGPpS   |
| Warsaw           |       | Warsaw           | gGPpS   |
| WashingtonDC     |       | WashingtonDC     | gGPpS   |
| Waterloo         |       | Waterloo         | gGPpS   |
| Wien             |       | Wien             | gGPpS   |
| Wroclaw          |       | Wroclaw          | gGPpS   |
| Wuerzburg        |       | Wuerzburg        | gGPpS   |
| Wuppertal        |       | Wuppertal        | gGPpS   |
| Zagreb           |       | Zagreb           | gGPpS   |
| Zuerich          |       | Zuerich          | gGPpS   |

</details>

<details>
<summary><strong>Geo2day Extracts</strong></summary>

| SHORT NAME        | IS IN | LONG NAME                | FORMATS |
|-------------------|-------|--------------------------|---------|
| africa            |       | [africa.poly]            | Pp      |
| antarctica        |       | [antarctica.poly]        | Pp      |
| asia              |       | [asia.poly]              | Pp      |
| australia_oceania |       | [australia_oceania.poly] | Pp      |
| central_america   |       | [central_america.poly]   | Pp      |
| europe            |       | [europe.poly]            | Pp      |
| north_america     |       | [north_america.poly]     | Pp      |
| russia            |       | [russia.poly]            | Pp      |
| south_america     |       | [south_america.poly]     | Pp      |

</details>

<details>
<summary><strong>Movisda Extracts</strong></summary>

| SHORT NAME | IS IN                                        | LONG NAME                                                   | FORMATS |
|------------|----------------------------------------------|-------------------------------------------------------------|---------|
| ad         |                                              | Andorra                                                     | gPp     |
| ae         |                                              | United Arab Emirates                                        | gPp     |
| ae-aj      | United Arab Emirates                         | Ajman Emirate                                               | gPp     |
| ae-az      | United Arab Emirates                         | Abu Dhabi Emirate                                           | gPp     |
| ae-du      | United Arab Emirates                         | Dubai Emirate                                               | gPp     |
| ae-fu      | United Arab Emirates                         | Fujairah Emirate                                            | gPp     |
| ae-rk      | United Arab Emirates                         | Ras al-Khaimah Emirate                                      | gPp     |
| ae-sh      | United Arab Emirates                         | Sharjah Emirate                                             | gPp     |
| ae-uq      | United Arab Emirates                         | Umm al-Quwain Emirate                                       | gPp     |
| af         |                                              | Afghanistan                                                 | gPp     |
| af-bal     | Afghanistan                                  | Balkh Province                                              | gPp     |
| af-bam     | Afghanistan                                  | Bamyan Province                                             | gPp     |
| af-bdg     | Afghanistan                                  | Badghis Province                                            | gPp     |
| af-bds     | Afghanistan                                  | Badakhshan Province                                         | gPp     |
| af-bgl     | Afghanistan                                  | Baghlan Province                                            | gPp     |
| af-day     | Afghanistan                                  | Daykundi Province                                           | gPp     |
| af-fra     | Afghanistan                                  | Farah Province                                              | gPp     |
| af-fyb     | Afghanistan                                  | Faryab Province                                             | gPp     |
| af-gha     | Afghanistan                                  | Ghazni Province                                             | gPp     |
| af-gho     | Afghanistan                                  | Ghor Province                                               | gPp     |
| af-hel     | Afghanistan                                  | Helmand Province                                            | gPp     |
| af-her     | Afghanistan                                  | Herat Province                                              | gPp     |
| af-jow     | Afghanistan                                  | Jowzjan Province                                            | gPp     |
| af-kab     | Afghanistan                                  | Kabul Province                                              | gPp     |
| af-kan     | Afghanistan                                  | Kandahar Province                                           | gPp     |
| af-kap     | Afghanistan                                  | Kapisa Province                                             | gPp     |
| af-kdz     | Afghanistan                                  | Kunduz Province                                             | gPp     |
| af-kho     | Afghanistan                                  | Khost Province                                              | gPp     |
| af-knr     | Afghanistan                                  | Kunar Province                                              | gPp     |
| af-lag     | Afghanistan                                  | Laghman Province                                            | gPp     |
| af-log     | Afghanistan                                  | Logar Province                                              | gPp     |
| af-nan     | Afghanistan                                  | Nangarhar Province                                          | gPp     |
| af-nim     | Afghanistan                                  | Nimruz Province                                             | gPp     |
| af-nur     | Afghanistan                                  | Nuristan Province                                           | gPp     |
| af-pan     | Afghanistan                                  | Panjshir Province                                           | gPp     |
| af-par     | Afghanistan                                  | Parwan Province                                             | gPp     |
| af-pia     | Afghanistan                                  | Paktia Province                                             | gPp     |
| af-pka     | Afghanistan                                  | Paktika Province                                            | gPp     |
| af-sam     | Afghanistan                                  | Samangan Province                                           | gPp     |
| af-sar     | Afghanistan                                  | Sar-e Pol Province                                          | gPp     |
| af-tak     | Afghanistan                                  | Takhar Province                                             | gPp     |
| af-uru     | Afghanistan                                  | Uruzgan Province                                            | gPp     |
| af-war     | Afghanistan                                  | Maidan Wardak Province                                      | gPp     |
| af-zab     | Afghanistan                                  | Zabul Province                                              | gPp     |
| ag         |                                              | Antigua and Barbuda                                         | gPp     |
| ai         |                                              | Anguilla                                                    | gPp     |
| al         |                                              | Albania                                                     | gPp     |
| al-01      | Albania                                      | Berat County                                                | gPp     |
| al-02      | Albania                                      | Durrës County                                               | gPp     |
| al-03      | Albania                                      | Elbasan County                                              | gPp     |
| al-04      | Albania                                      | Fier County                                                 | gPp     |
| al-05      | Albania                                      | Gjirokastër County                                          | gPp     |
| al-06      | Albania                                      | Korçë County                                                | gPp     |
| al-07      | Albania                                      | Kukës County                                                | gPp     |
| al-08      | Albania                                      | Lezhë County                                                | gPp     |
| al-09      | Albania                                      | Dibër County                                                | gPp     |
| al-10      | Albania                                      | Shkodër County                                              | gPp     |
| al-11      | Albania                                      | Tirana County                                               | gPp     |
| al-12      | Albania                                      | Vlorë County                                                | gPp     |
| am         |                                              | Armenia                                                     | gPp     |
| am-ag      | Armenia                                      | Aragatsotn Province                                         | gPp     |
| am-ar      | Armenia                                      | Ararat Province                                             | gPp     |
| am-av      | Armenia                                      | Armavir Province                                            | gPp     |
| am-er      | Armenia                                      | Yerevan                                                     | gPp     |
| am-gr      | Armenia                                      | Gegharkunik Province                                        | gPp     |
| am-kt      | Armenia                                      | Kotayk Province                                             | gPp     |
| am-lo      | Armenia                                      | Lori Province                                               | gPp     |
| am-sh      | Armenia                                      | Shirak Province                                             | gPp     |
| am-su      | Armenia                                      | Syunik Province                                             | gPp     |
| am-tv      | Armenia                                      | Tavush Province                                             | gPp     |
| am-vd      | Armenia                                      | Vayots Dzor Province                                        | gPp     |
| ao         |                                              | Angola                                                      | gPp     |
| ao-bgo     | Angola                                       | Bengo Province                                              | gPp     |
| ao-bgu     | Angola                                       | Benguela Province                                           | gPp     |
| ao-bie     | Angola                                       | Bié Province                                                | gPp     |
| ao-cab     | Angola                                       | Cabinda Province                                            | gPp     |
| ao-ccu     | Angola                                       | Cuando Cubango Province                                     | gPp     |
| ao-cnn     | Angola                                       | Cunene Province                                             | gPp     |
| ao-cno     | Angola                                       | Cuanza Norte Province                                       | gPp     |
| ao-cus     | Angola                                       | Cuanza Sul Province                                         | gPp     |
| ao-hua     | Angola                                       | Huambo Province                                             | gPp     |
| ao-hui     | Angola                                       | Huíla Province                                              | gPp     |
| ao-lno     | Angola                                       | Lunda Norte Province                                        | gPp     |
| ao-lsu     | Angola                                       | Lunda Sul Province                                          | gPp     |
| ao-lua     | Angola                                       | Luanda Province                                             | gPp     |
| ao-mal     | Angola                                       | Malanje Province                                            | gPp     |
| ao-mox     | Angola                                       | Moxico Province                                             | gPp     |
| ao-nam     | Angola                                       | Namibe Province                                             | gPp     |
| ao-uig     | Angola                                       | Uíge Province                                               | gPp     |
| ao-zai     | Angola                                       | Zaire Province                                              | gPp     |
| aq         |                                              | Antarctica                                                  | gPp     |
| ar         |                                              | Argentina                                                   | gPp     |
| ar-a       | Argentina                                    | Salta                                                       | gPp     |
| ar-b       | Argentina                                    | Buenos Aires                                                | gPp     |
| ar-c       | Argentina                                    | Autonomous City of Buenos Aires                             | gPp     |
| ar-d       | Argentina                                    | San Luis                                                    | gPp     |
| ar-e       | Argentina                                    | Entre Ríos Province                                         | gPp     |
| ar-f       | Argentina                                    | La Rioja                                                    | gPp     |
| ar-g       | Argentina                                    | Santiago del Estero                                         | gPp     |
| ar-h       | Argentina                                    | Chaco                                                       | gPp     |
| ar-j       | Argentina                                    | San Juan                                                    | gPp     |
| ar-k       | Argentina                                    | Catamarca                                                   | gPp     |
| ar-l       | Argentina                                    | La Pampa                                                    | gPp     |
| ar-m       | Argentina                                    | Mendoza                                                     | gPp     |
| ar-n       | Argentina                                    | Misiones                                                    | gPp     |
| ar-p       | Argentina                                    | Formosa                                                     | gPp     |
| ar-q       | Argentina                                    | Neuquén                                                     | gPp     |
| ar-r       | Argentina                                    | Río Negro                                                   | gPp     |
| ar-s       | Argentina                                    | Santa Fe                                                    | gPp     |
| ar-t       | Argentina                                    | Tucumán                                                     | gPp     |
| ar-u       | Argentina                                    | Chubut                                                      | gPp     |
| ar-v       | Argentina                                    | Tierra del Fuego                                            | gPp     |
| ar-w       | Argentina                                    | Corrientes                                                  | gPp     |
| ar-x       | Argentina                                    | Córdoba                                                     | gPp     |
| ar-y       | Argentina                                    | Jujuy                                                       | gPp     |
| ar-z       | Argentina                                    | Santa Cruz                                                  | gPp     |
| at         |                                              | Austria                                                     | gPp     |
| at-1       | Austria                                      | Burgenland                                                  | gPp     |
| at-2       | Austria                                      | Carinthia                                                   | gPp     |
| at-3       | Austria                                      | Lower Austria                                               | gPp     |
| at-4       | Austria                                      | Upper Austria                                               | gPp     |
| at-5       | Austria                                      | Salzburg                                                    | gPp     |
| at-6       | Austria                                      | Styria                                                      | gPp     |
| at-7       | Austria                                      | Tyrol                                                       | gPp     |
| at-8       | Austria                                      | Vorarlberg                                                  | gPp     |
| at-9       | Austria                                      | Vienna                                                      | gPp     |
| au         |                                              | Australia                                                   | gPp     |
| au-act     | Australia                                    | Australian Capital Territory                                | gPp     |
| au-cc      | Australia                                    | Cocos (Keeling) Islands                                     | gPp     |
| au-cx      | Australia                                    | Christmas Island                                            | gPp     |
| au-hm      | Australia                                    | Heard Island and McDonald Islands                           | gPp     |
| au-nf      | Australia                                    | Norfolk Island                                              | gPp     |
| au-nsw     | Australia                                    | New South Wales                                             | gPp     |
| au-nt      | Australia                                    | Northern Territory                                          | gPp     |
| au-qld     | Australia                                    | Queensland                                                  | gPp     |
| au-sa      | Australia                                    | South Australia                                             | gPp     |
| au-tas     | Australia                                    | Tasmania                                                    | gPp     |
| au-vic     | Australia                                    | Victoria                                                    | gPp     |
| au-wa      | Australia                                    | Western Australia                                           | gPp     |
| az         |                                              | Azerbaijan                                                  | gPp     |
| az-abs     | Azerbaijan                                   | Absheron District                                           | gPp     |
| az-aga     | Azerbaijan                                   | Aghstafa District                                           | gPp     |
| az-agc     | Azerbaijan                                   | Aghjabadi District                                          | gPp     |
| az-agd     | Azerbaijan                                   | Agdere District                                             | gPp     |
| az-agm     | Azerbaijan                                   | Aghdam District                                             | gPp     |
| az-ags     | Azerbaijan                                   | Agdash District                                             | gPp     |
| az-agu     | Azerbaijan                                   | Agsu District                                               | gPp     |
| az-ast     | Azerbaijan                                   | Astara District                                             | gPp     |
| az-ba      | Azerbaijan                                   | Baku Zone                                                   | gPp     |
| az-bal     | Azerbaijan                                   | Balakan District                                            | gPp     |
| az-bar     | Azerbaijan                                   | Barda District                                              | gPp     |
| az-bey     | Azerbaijan                                   | Beylagan District                                           | gPp     |
| az-bil     | Azerbaijan                                   | Bilasuvar District                                          | gPp     |
| az-cab     | Azerbaijan                                   | Jabrayil District                                           | gPp     |
| az-cal     | Azerbaijan                                   | Jalilabad District                                          | gPp     |
| az-das     | Azerbaijan                                   | Dashkasan District                                          | gPp     |
| az-fuz     | Azerbaijan                                   | Fizuli District                                             | gPp     |
| az-ga      | Azerbaijan                                   | Ganja City                                                  | gPp     |
| az-gad     | Azerbaijan                                   | Gedebey District                                            | gPp     |
| az-gor     | Azerbaijan                                   | Goranboy District                                           | gPp     |
| az-goy     | Azerbaijan                                   | Goychay District                                            | gPp     |
| az-gyg     | Azerbaijan                                   | Goygol District                                             | gPp     |
| az-hac     | Azerbaijan                                   | Hajigabul District                                          | gPp     |
| az-imi     | Azerbaijan                                   | Imishli District                                            | gPp     |
| az-ism     | Azerbaijan                                   | Ismailli District                                           | gPp     |
| az-kal     | Azerbaijan                                   | Kalbajar District                                           | gPp     |
| az-kur     | Azerbaijan                                   | Kurdamir District                                           | gPp     |
| az-la      | Azerbaijan                                   | Lankaran                                                    | gPp     |
| az-lac     | Azerbaijan                                   | Lachin District                                             | gPp     |
| az-lan     | Azerbaijan                                   | Lankaran District                                           | gPp     |
| az-ler     | Azerbaijan                                   | Lerik District                                              | gPp     |
| az-mas     | Azerbaijan                                   | Masally District                                            | gPp     |
| az-mi      | Azerbaijan                                   | Mingachevir                                                 | gPp     |
| az-na      | Azerbaijan                                   | Naftalan                                                    | gPp     |
| az-nef     | Azerbaijan                                   | Neftchala District                                          | gPp     |
| az-nx      | Azerbaijan                                   | Nakhchivan Autonomous Republic                              | gPp     |
| az-ogu     | Azerbaijan                                   | Oghuz District                                              | gPp     |
| az-qab     | Azerbaijan                                   | Qabala District                                             | gPp     |
| az-qax     | Azerbaijan                                   | Qakh District                                               | gPp     |
| az-qaz     | Azerbaijan                                   | Qazakh District                                             | gPp     |
| az-qba     | Azerbaijan                                   | Quba District                                               | gPp     |
| az-qbi     | Azerbaijan                                   | Qubadli District                                            | gPp     |
| az-qob     | Azerbaijan                                   | Gobustan District                                           | gPp     |
| az-qus     | Azerbaijan                                   | Qusar District                                              | gPp     |
| az-sa      | Azerbaijan                                   | Shaki                                                       | gPp     |
| az-sab     | Azerbaijan                                   | Sabirabad District                                          | gPp     |
| az-sak     | Azerbaijan                                   | Sheki District                                              | gPp     |
| az-sal     | Azerbaijan                                   | Salyan District                                             | gPp     |
| az-sat     | Azerbaijan                                   | Saatly District                                             | gPp     |
| az-sbn     | Azerbaijan                                   | Shabran District                                            | gPp     |
| az-siy     | Azerbaijan                                   | Siazan District                                             | gPp     |
| az-skr     | Azerbaijan                                   | Shamkir District                                            | gPp     |
| az-sm      | Azerbaijan                                   | Sumqayit                                                    | gPp     |
| az-smi     | Azerbaijan                                   | Shamakhi District                                           | gPp     |
| az-smx     | Azerbaijan                                   | Samukh District                                             | gPp     |
| az-sr      | Azerbaijan                                   | Shirvan District                                            | gPp     |
| az-sus     | Azerbaijan                                   | Shusha District                                             | gPp     |
| az-tar     | Azerbaijan                                   | Tartar District                                             | gPp     |
| az-tov     | Azerbaijan                                   | Tovuz District                                              | gPp     |
| az-uca     | Azerbaijan                                   | Ujar District                                               | gPp     |
| az-xa      | Azerbaijan                                   | Khankendi                                                   | gPp     |
| az-xac     | Azerbaijan                                   | Khachmaz District                                           | gPp     |
| az-xci     | Azerbaijan                                   | Khojaly District                                            | gPp     |
| az-xiz     | Azerbaijan                                   | Khizi District                                              | gPp     |
| az-xvd     | Azerbaijan                                   | Khojavend District                                          | gPp     |
| az-yar     | Azerbaijan                                   | Yardymli District                                           | gPp     |
| az-ye      | Azerbaijan                                   | Yevlakh                                                     | gPp     |
| az-yev     | Azerbaijan                                   | Yevlakh District                                            | gPp     |
| az-zan     | Azerbaijan                                   | Zangilan District                                           | gPp     |
| az-zaq     | Azerbaijan                                   | Zaqatala District                                           | gPp     |
| az-zar     | Azerbaijan                                   | Zardab District                                             | gPp     |
| ba         |                                              | Bosnia and Herzegovina                                      | gPp     |
| ba-bih     | Bosnia and Herzegovina                       | Federation of Bosnia and Herzegovina                        | gPp     |
| ba-brc     | Bosnia and Herzegovina                       | Brčko District                                              | gPp     |
| ba-srp     | Bosnia and Herzegovina                       | Republika Srpska                                            | gPp     |
| bb         |                                              | Barbados                                                    | gPp     |
| bb-01      | Barbados                                     | Christ Church                                               | gPp     |
| bb-02      | Barbados                                     | Saint Andrew                                                | gPp     |
| bb-03      | Barbados                                     | Saint George                                                | gPp     |
| bb-04      | Barbados                                     | Saint James                                                 | gPp     |
| bb-05      | Barbados                                     | Saint John                                                  | gPp     |
| bb-06      | Barbados                                     | Saint Joseph                                                | gPp     |
| bb-07      | Barbados                                     | Saint Lucy                                                  | gPp     |
| bb-08      | Barbados                                     | Saint Michael                                               | gPp     |
| bb-09      | Barbados                                     | Saint Peter                                                 | gPp     |
| bb-10      | Barbados                                     | Saint Philip                                                | gPp     |
| bb-11      | Barbados                                     | Saint Thomas                                                | gPp     |
| bd         |                                              | Bangladesh                                                  | gPp     |
| bd-a       | Bangladesh                                   | Barishal Division                                           | gPp     |
| bd-b       | Bangladesh                                   | Chattogram Division                                         | gPp     |
| bd-c       | Bangladesh                                   | Dhaka Division                                              | gPp     |
| bd-d       | Bangladesh                                   | Khulna Division                                             | gPp     |
| bd-e       | Bangladesh                                   | Rajshahi Division                                           | gPp     |
| bd-f       | Bangladesh                                   | Rangpur Division                                            | gPp     |
| bd-g       | Bangladesh                                   | Sylhet Division                                             | gPp     |
| bd-h       | Bangladesh                                   | Mymensingh Division                                         | gPp     |
| be         |                                              | Belgium                                                     | gPp     |
| be-bru     | Belgium                                      | Brussels-Capital                                            | gPp     |
| be-vlg     | Belgium                                      | Flanders                                                    | gPp     |
| be-wal     | Belgium                                      | Wallonia                                                    | gPp     |
| bf         |                                              | Burkina Faso                                                | gPp     |
| bf-01      | Burkina Faso                                 | Boucle du Mouhoun                                           | gPp     |
| bf-02      | Burkina Faso                                 | Waterfalls                                                  | gPp     |
| bf-03      | Burkina Faso                                 | Centre                                                      | gPp     |
| bf-04      | Burkina Faso                                 | Central-East                                                | gPp     |
| bf-05      | Burkina Faso                                 | Central-North                                               | gPp     |
| bf-06      | Burkina Faso                                 | Central-West                                                | gPp     |
| bf-07      | Burkina Faso                                 | Central-South                                               | gPp     |
| bf-08      | Burkina Faso                                 | East                                                        | gPp     |
| bf-09      | Burkina Faso                                 | Upper-Basins                                                | gPp     |
| bf-10      | Burkina Faso                                 | North                                                       | gPp     |
| bf-11      | Burkina Faso                                 | Central-Plateau                                             | gPp     |
| bf-12      | Burkina Faso                                 | Sahel                                                       | gPp     |
| bf-13      | Burkina Faso                                 | Southwest                                                   | gPp     |
| bg         |                                              | Bulgaria                                                    | gPp     |
| bg-01      | Bulgaria                                     | Blagoevgrad                                                 | gPp     |
| bg-02      | Bulgaria                                     | Burgas                                                      | gPp     |
| bg-03      | Bulgaria                                     | Varna                                                       | gPp     |
| bg-04      | Bulgaria                                     | Veliko Tarnovo                                              | gPp     |
| bg-05      | Bulgaria                                     | Vidin                                                       | gPp     |
| bg-06      | Bulgaria                                     | Vratsa                                                      | gPp     |
| bg-07      | Bulgaria                                     | Gabrovo                                                     | gPp     |
| bg-08      | Bulgaria                                     | Dobrich                                                     | gPp     |
| bg-09      | Bulgaria                                     | Kardzhali                                                   | gPp     |
| bg-10      | Bulgaria                                     | Kyustendil                                                  | gPp     |
| bg-11      | Bulgaria                                     | Lovech                                                      | gPp     |
| bg-12      | Bulgaria                                     | Montana                                                     | gPp     |
| bg-13      | Bulgaria                                     | Pazardzhik                                                  | gPp     |
| bg-14      | Bulgaria                                     | Pernik                                                      | gPp     |
| bg-15      | Bulgaria                                     | Pleven                                                      | gPp     |
| bg-16      | Bulgaria                                     | Plovdiv                                                     | gPp     |
| bg-17      | Bulgaria                                     | Razgrad                                                     | gPp     |
| bg-18      | Bulgaria                                     | Ruse                                                        | gPp     |
| bg-19      | Bulgaria                                     | Silistra                                                    | gPp     |
| bg-20      | Bulgaria                                     | Sliven                                                      | gPp     |
| bg-21      | Bulgaria                                     | Smolyan                                                     | gPp     |
| bg-22      | Bulgaria                                     | Sofia-City                                                  | gPp     |
| bg-23      | Bulgaria                                     | Sofia                                                       | gPp     |
| bg-24      | Bulgaria                                     | Stara Zagora                                                | gPp     |
| bg-25      | Bulgaria                                     | Targovishte                                                 | gPp     |
| bg-26      | Bulgaria                                     | Haskovo                                                     | gPp     |
| bg-27      | Bulgaria                                     | Shumen                                                      | gPp     |
| bg-28      | Bulgaria                                     | Yambol                                                      | gPp     |
| bh         |                                              | Bahrain                                                     | gPp     |
| bh-13      | Bahrain                                      | Capital Governorate                                         | gPp     |
| bh-14      | Bahrain                                      | Southern Governorate                                        | gPp     |
| bh-15      | Bahrain                                      | Muharraq Governorate                                        | gPp     |
| bh-17      | Bahrain                                      | Northern Governorate                                        | gPp     |
| bi         |                                              | Burundi                                                     | gPp     |
| bi-bb      | Burundi                                      | Bubanza                                                     | gPp     |
| bi-bl      | Burundi                                      | Bujumbura Rural Province                                    | gPp     |
| bi-bm      | Burundi                                      | Bujumbura Mairie                                            | gPp     |
| bi-br      | Burundi                                      | Bururi                                                      | gPp     |
| bi-ca      | Burundi                                      | Cankuzo                                                     | gPp     |
| bi-ci      | Burundi                                      | Cibitoke                                                    | gPp     |
| bi-gi      | Burundi                                      | Gitega                                                      | gPp     |
| bi-ki      | Burundi                                      | Kirundo                                                     | gPp     |
| bi-kr      | Burundi                                      | Karuzi                                                      | gPp     |
| bi-ky      | Burundi                                      | Kayanza                                                     | gPp     |
| bi-ma      | Burundi                                      | Makamba                                                     | gPp     |
| bi-mu      | Burundi                                      | Muramvya                                                    | gPp     |
| bi-mw      | Burundi                                      | Mwaro                                                       | gPp     |
| bi-my      | Burundi                                      | Muyinga                                                     | gPp     |
| bi-ng      | Burundi                                      | Ngozi                                                       | gPp     |
| bi-rm      | Burundi                                      | Rumonge                                                     | gPp     |
| bi-rt      | Burundi                                      | Rutana                                                      | gPp     |
| bi-ry      | Burundi                                      | Ruyigi                                                      | gPp     |
| bj         |                                              | Benin                                                       | gPp     |
| bj-ak      | Benin                                        | Atakora Department                                          | gPp     |
| bj-al      | Benin                                        | Alibori Department                                          | gPp     |
| bj-aq      | Benin                                        | Atlantique Department                                       | gPp     |
| bj-bo      | Benin                                        | Borgou Department                                           | gPp     |
| bj-co      | Benin                                        | Collines Department                                         | gPp     |
| bj-do      | Benin                                        | Donga Department                                            | gPp     |
| bj-ko      | Benin                                        | Kouffo Department                                           | gPp     |
| bj-li      | Benin                                        | Littoral Department                                         | gPp     |
| bj-mo      | Benin                                        | Mono Department                                             | gPp     |
| bj-ou      | Benin                                        | Ouémé Department                                            | gPp     |
| bj-pl      | Benin                                        | Plateau Department                                          | gPp     |
| bj-zo      | Benin                                        | Zou Department                                              | gPp     |
| bm         |                                              | Bermuda                                                     | gPp     |
| bn         |                                              | Brunei                                                      | gPp     |
| bn-be      | Brunei                                       | Belait District                                             | gPp     |
| bn-bm      | Brunei                                       | Brunei-Muara District                                       | gPp     |
| bn-te      | Brunei                                       | Temburong District                                          | gPp     |
| bn-tu      | Brunei                                       | Tutong District                                             | gPp     |
| bo         |                                              | Bolivia                                                     | gPp     |
| bo-b       | Bolivia                                      | Beni                                                        | gPp     |
| bo-c       | Bolivia                                      | Cochabamba                                                  | gPp     |
| bo-h       | Bolivia                                      | Chuquisaca                                                  | gPp     |
| bo-l       | Bolivia                                      | La Paz                                                      | gPp     |
| bo-n       | Bolivia                                      | Pando                                                       | gPp     |
| bo-o       | Bolivia                                      | Oruro                                                       | gPp     |
| bo-p       | Bolivia                                      | Potosí                                                      | gPp     |
| bo-s       | Bolivia                                      | Santa Cruz                                                  | gPp     |
| bo-t       | Bolivia                                      | Tarija                                                      | gPp     |
| br         |                                              | Brazil                                                      | gPp     |
| br-ac      | Brazil                                       | Acre                                                        | gPp     |
| br-al      | Brazil                                       | Alagoas                                                     | gPp     |
| br-am      | Brazil                                       | Amazonas                                                    | gPp     |
| br-ap      | Brazil                                       | Amapá                                                       | gPp     |
| br-ba      | Brazil                                       | Bahia                                                       | gPp     |
| br-ce      | Brazil                                       | Ceará                                                       | gPp     |
| br-df      | Brazil                                       | Federal District                                            | gPp     |
| br-es      | Brazil                                       | Espírito Santo                                              | gPp     |
| br-go      | Brazil                                       | Goiás                                                       | gPp     |
| br-ma      | Brazil                                       | Maranhão                                                    | gPp     |
| br-mg      | Brazil                                       | Minas Gerais                                                | gPp     |
| br-ms      | Brazil                                       | Mato Grosso do Sul                                          | gPp     |
| br-mt      | Brazil                                       | Mato Grosso                                                 | gPp     |
| br-pa      | Brazil                                       | Pará                                                        | gPp     |
| br-pb      | Brazil                                       | Paraíba                                                     | gPp     |
| br-pe      | Brazil                                       | Pernambuco                                                  | gPp     |
| br-pi      | Brazil                                       | Piauí                                                       | gPp     |
| br-pr      | Brazil                                       | Paraná                                                      | gPp     |
| br-rj      | Brazil                                       | Rio de Janeiro                                              | gPp     |
| br-rn      | Brazil                                       | Rio Grande do Norte                                         | gPp     |
| br-ro      | Brazil                                       | Rondônia                                                    | gPp     |
| br-rr      | Brazil                                       | Roraima                                                     | gPp     |
| br-rs      | Brazil                                       | Rio Grande do Sul                                           | gPp     |
| br-sc      | Brazil                                       | Santa Catarina                                              | gPp     |
| br-se      | Brazil                                       | Sergipe                                                     | gPp     |
| br-sp      | Brazil                                       | São Paulo                                                   | gPp     |
| br-to      | Brazil                                       | Tocantins                                                   | gPp     |
| bs         |                                              | Bahamas                                                     | gPp     |
| bt         |                                              | Bhutan                                                      | gPp     |
| bt-11      | Bhutan                                       | Paro District                                               | gPp     |
| bt-12      | Bhutan                                       | Chukha                                                      | gPp     |
| bt-13      | Bhutan                                       | Haa District                                                | gPp     |
| bt-14      | Bhutan                                       | Samtse District                                             | gPp     |
| bt-15      | Bhutan                                       | Thimphu District                                            | gPp     |
| bt-21      | Bhutan                                       | Tsirang District                                            | gPp     |
| bt-22      | Bhutan                                       | Dagana District                                             | gPp     |
| bt-23      | Bhutan                                       | Punakha District                                            | gPp     |
| bt-24      | Bhutan                                       | Wangdue Phodrang District                                   | gPp     |
| bt-31      | Bhutan                                       | Sarpang District                                            | gPp     |
| bt-32      | Bhutan                                       | Trongsa District                                            | gPp     |
| bt-33      | Bhutan                                       | Bumthang District                                           | gPp     |
| bt-34      | Bhutan                                       | Zhemgang District                                           | gPp     |
| bt-41      | Bhutan                                       | Trashigang District                                         | gPp     |
| bt-42      | Bhutan                                       | Mongar District                                             | gPp     |
| bt-43      | Bhutan                                       | Pemagatshel District                                        | gPp     |
| bt-44      | Bhutan                                       | Lhuntse District                                            | gPp     |
| bt-45      | Bhutan                                       | Samdrup Jongkhar District                                   | gPp     |
| bt-ga      | Bhutan                                       | Gasa District                                               | gPp     |
| bt-ty      | Bhutan                                       | Trashiyangtse District                                      | gPp     |
| bw         |                                              | Botswana                                                    | gPp     |
| bw-ce      | Botswana                                     | Central District                                            | gPp     |
| bw-ch      | Botswana                                     | Chobe District                                              | gPp     |
| bw-fr      | Botswana                                     | Francistown                                                 | gPp     |
| bw-ga      | Botswana                                     | Gaborone                                                    | gPp     |
| bw-gh      | Botswana                                     | Ghanzi District                                             | gPp     |
| bw-kg      | Botswana                                     | Kgalagadi District                                          | gPp     |
| bw-kl      | Botswana                                     | Kgatleng District                                           | gPp     |
| bw-kw      | Botswana                                     | Kweneng District                                            | gPp     |
| bw-lo      | Botswana                                     | Lobatse                                                     | gPp     |
| bw-ne      | Botswana                                     | North-East District                                         | gPp     |
| bw-nw      | Botswana                                     | North-West District                                         | gPp     |
| bw-se      | Botswana                                     | South-East District                                         | gPp     |
| bw-so      | Botswana                                     | Southern District                                           | gPp     |
| bw-sp      | Botswana                                     | Selebi Phikwe                                               | gPp     |
| bw-st      | Botswana                                     | Sowa Town                                                   | gPp     |
| bx         |                                              | Bir Tawil                                                   | gPp     |
| by         |                                              | Belarus                                                     | gPp     |
| by-br      | Belarus                                      | Brest Region                                                | gPp     |
| by-hm      | Belarus                                      | Minsk                                                       | gPp     |
| by-ho      | Belarus                                      | Homyel Region                                               | gPp     |
| by-hr      | Belarus                                      | Hrodna Region                                               | gPp     |
| by-ma      | Belarus                                      | Mahilyow Region                                             | gPp     |
| by-mi      | Belarus                                      | Minsk Region                                                | gPp     |
| by-vi      | Belarus                                      | Vitsebsk Region                                             | gPp     |
| bz         |                                              | Belize                                                      | gPp     |
| bz-bz      | Belize                                       | Belize District                                             | gPp     |
| bz-cy      | Belize                                       | Cayo                                                        | gPp     |
| bz-czl     | Belize                                       | Corozal                                                     | gPp     |
| bz-ow      | Belize                                       | Orange Walk                                                 | gPp     |
| bz-sc      | Belize                                       | Stann Creek                                                 | gPp     |
| bz-tol     | Belize                                       | Toledo                                                      | gPp     |
| ca         |                                              | Canada                                                      | gPp     |
| ca-ab      | Canada                                       | Alberta                                                     | gPp     |
| ca-bc      | Canada                                       | British Columbia                                            | gPp     |
| ca-mb      | Canada                                       | Manitoba                                                    | gPp     |
| ca-nb      | Canada                                       | New Brunswick                                               | gPp     |
| ca-nl      | Canada                                       | Newfoundland and Labrador                                   | gPp     |
| ca-ns      | Canada                                       | Nova Scotia                                                 | gPp     |
| ca-nt      | Canada                                       | Northwest Territories                                       | gPp     |
| ca-nu      | Canada                                       | Nunavut                                                     | gPp     |
| ca-on      | Canada                                       | Ontario                                                     | gPp     |
| ca-pe      | Canada                                       | Prince Edward Island                                        | gPp     |
| ca-qc      | Canada                                       | Quebec                                                      | gPp     |
| ca-sk      | Canada                                       | Saskatchewan                                                | gPp     |
| ca-yt      | Canada                                       | Yukon                                                       | gPp     |
| cd         |                                              | Democratic Republic of the Congo                            | gPp     |
| cd-bc      | Democratic Republic of the Congo             | Kongo Central                                               | gPp     |
| cd-bu      | Democratic Republic of the Congo             | Lower Uele                                                  | gPp     |
| cd-eq      | Democratic Republic of the Congo             | Équateur                                                    | gPp     |
| cd-hk      | Democratic Republic of the Congo             | Haut-Katanga                                                | gPp     |
| cd-hl      | Democratic Republic of the Congo             | Haut-Lomami                                                 | gPp     |
| cd-hu      | Democratic Republic of the Congo             | Upper Uele                                                  | gPp     |
| cd-it      | Democratic Republic of the Congo             | Ituri                                                       | gPp     |
| cd-kc      | Democratic Republic of the Congo             | Central Kasai                                               | gPp     |
| cd-ke      | Democratic Republic of the Congo             | Kasaï-Oriental                                              | gPp     |
| cd-kg      | Democratic Republic of the Congo             | Kwango                                                      | gPp     |
| cd-kl      | Democratic Republic of the Congo             | Kwilu                                                       | gPp     |
| cd-kn      | Democratic Republic of the Congo             | Kinshasa                                                    | gPp     |
| cd-ks      | Democratic Republic of the Congo             | Kasai                                                       | gPp     |
| cd-lo      | Democratic Republic of the Congo             | Lomami                                                      | gPp     |
| cd-lu      | Democratic Republic of the Congo             | Lualaba                                                     | gPp     |
| cd-ma      | Democratic Republic of the Congo             | Maniema                                                     | gPp     |
| cd-mn      | Democratic Republic of the Congo             | Mai-Ndombe                                                  | gPp     |
| cd-mo      | Democratic Republic of the Congo             | Mongala                                                     | gPp     |
| cd-nk      | Democratic Republic of the Congo             | North Kivu                                                  | gPp     |
| cd-nu      | Democratic Republic of the Congo             | Nord-Ubangi                                                 | gPp     |
| cd-sa      | Democratic Republic of the Congo             | Sankuru                                                     | gPp     |
| cd-sk      | Democratic Republic of the Congo             | South Kivu                                                  | gPp     |
| cd-su      | Democratic Republic of the Congo             | Sud-Ubangi                                                  | gPp     |
| cd-ta      | Democratic Republic of the Congo             | Tanganyika                                                  | gPp     |
| cd-to      | Democratic Republic of the Congo             | Tshopo                                                      | gPp     |
| cd-tu      | Democratic Republic of the Congo             | Tshuapa                                                     | gPp     |
| cf         |                                              | Central African Republic                                    | gPp     |
| cf-ac      | Central African Republic                     | Ouham                                                       | gPp     |
| cf-bb      | Central African Republic                     | Bamingui-Bangoran                                           | gPp     |
| cf-bgf     | Central African Republic                     | Bangui                                                      | gPp     |
| cf-bk      | Central African Republic                     | Basse-Kotto                                                 | gPp     |
| cf-hk      | Central African Republic                     | Haute-Kotto                                                 | gPp     |
| cf-hm      | Central African Republic                     | Haut-Mbomou                                                 | gPp     |
| cf-hs      | Central African Republic                     | Mambéré-Kadéï                                               | gPp     |
| cf-kb      | Central African Republic                     | Nana-Grébizi                                                | gPp     |
| cf-kg      | Central African Republic                     | Kémo                                                        | gPp     |
| cf-lb      | Central African Republic                     | Lobaye                                                      | gPp     |
| cf-mb      | Central African Republic                     | Mbomou                                                      | gPp     |
| cf-mp      | Central African Republic                     | Ombella-M'Poko                                              | gPp     |
| cf-nm      | Central African Republic                     | Nana-Mambéré                                                | gPp     |
| cf-op      | Central African Republic                     | Ouham-Pendé                                                 | gPp     |
| cf-se      | Central African Republic                     | Sangha-Mbaéré                                               | gPp     |
| cf-uk      | Central African Republic                     | Ouaka                                                       | gPp     |
| cf-vk      | Central African Republic                     | Vakaga                                                      | gPp     |
| cg         |                                              | Congo-Brazzaville                                           | gPp     |
| cg-11      | Congo-Brazzaville                            | Bouenza Department                                          | gPp     |
| cg-12      | Congo-Brazzaville                            | Pool Department                                             | gPp     |
| cg-13      | Congo-Brazzaville                            | Sangha                                                      | gPp     |
| cg-14      | Congo-Brazzaville                            | Plateaux Department                                         | gPp     |
| cg-15      | Congo-Brazzaville                            | Cuvette-Ouest Department                                    | gPp     |
| cg-16      | Congo-Brazzaville                            | Pointe-Noire (département)                                  | gPp     |
| cg-2       | Congo-Brazzaville                            | Lékoumou Department                                         | gPp     |
| cg-5       | Congo-Brazzaville                            | Kouilou Department                                          | gPp     |
| cg-7       | Congo-Brazzaville                            | Likouala                                                    | gPp     |
| cg-8       | Congo-Brazzaville                            | Cuvette Department                                          | gPp     |
| cg-9       | Congo-Brazzaville                            | Niari Department                                            | gPp     |
| cg-bzv     | Congo-Brazzaville                            | Brazzaville (department)                                    | gPp     |
| ch         |                                              | Switzerland                                                 | gPp     |
| ch-ag      | Switzerland                                  | Aargau                                                      | gPp     |
| ch-ai      | Switzerland                                  | Appenzell Innerrhoden                                       | gPp     |
| ch-ar      | Switzerland                                  | Appenzell Ausserrhoden                                      | gPp     |
| ch-be      | Switzerland                                  | Bern                                                        | gPp     |
| ch-bl      | Switzerland                                  | Basel-Landschaft                                            | gPp     |
| ch-bs      | Switzerland                                  | Basel-City                                                  | gPp     |
| ch-fr      | Switzerland                                  | Fribourg                                                    | gPp     |
| ch-ge      | Switzerland                                  | Geneva                                                      | gPp     |
| ch-gl      | Switzerland                                  | Glarus                                                      | gPp     |
| ch-gr      | Switzerland                                  | Grisons                                                     | gPp     |
| ch-ju      | Switzerland                                  | Jura                                                        | gPp     |
| ch-lu      | Switzerland                                  | Lucerne                                                     | gPp     |
| ch-ne      | Switzerland                                  | Neuchâtel                                                   | gPp     |
| ch-nw      | Switzerland                                  | Nidwalden                                                   | gPp     |
| ch-ow      | Switzerland                                  | Obwalden                                                    | gPp     |
| ch-sg      | Switzerland                                  | St. Gallen                                                  | gPp     |
| ch-sh      | Switzerland                                  | Schaffhausen                                                | gPp     |
| ch-so      | Switzerland                                  | Solothurn                                                   | gPp     |
| ch-sz      | Switzerland                                  | Schwyz                                                      | gPp     |
| ch-tg      | Switzerland                                  | Thurgau                                                     | gPp     |
| ch-ti      | Switzerland                                  | Ticino                                                      | gPp     |
| ch-ur      | Switzerland                                  | Uri                                                         | gPp     |
| ch-vd      | Switzerland                                  | Vaud                                                        | gPp     |
| ch-vs      | Switzerland                                  | Valais/Wallis                                               | gPp     |
| ch-zg      | Switzerland                                  | Zug                                                         | gPp     |
| ch-zh      | Switzerland                                  | Zurich                                                      | gPp     |
| ci         |                                              | Ivory Coast                                                 | gPp     |
| ci-ab      | Ivory Coast                                  | Abidjan                                                     | gPp     |
| ci-bs      | Ivory Coast                                  | Bas-Sassandra                                               | gPp     |
| ci-cm      | Ivory Coast                                  | Comoé                                                       | gPp     |
| ci-dn      | Ivory Coast                                  | Denguélé District                                           | gPp     |
| ci-gd      | Ivory Coast                                  | Gôh-Djiboua                                                 | gPp     |
| ci-lc      | Ivory Coast                                  | Lacs                                                        | gPp     |
| ci-lg      | Ivory Coast                                  | Lagunes                                                     | gPp     |
| ci-mg      | Ivory Coast                                  | Montagnes                                                   | gPp     |
| ci-sm      | Ivory Coast                                  | Sassandra-Marahoué                                          | gPp     |
| ci-sv      | Ivory Coast                                  | Savanes                                                     | gPp     |
| ci-vb      | Ivory Coast                                  | Vallée du Bandama                                           | gPp     |
| ci-wr      | Ivory Coast                                  | Woroba                                                      | gPp     |
| ci-ym      | Ivory Coast                                  | Yamoussoukro                                                | gPp     |
| ci-zz      | Ivory Coast                                  | Zanzan                                                      | gPp     |
| ck         |                                              | Cook Islands                                                | gPp     |
| cl         |                                              | Chile                                                       | gPp     |
| cl-ai      | Chile                                        | Aysen del General Carlos Ibanez del Campo Region            | gPp     |
| cl-an      | Chile                                        | Antofagasta Region                                          | gPp     |
| cl-ap      | Chile                                        | Arica y Parinacota Region                                   | gPp     |
| cl-ar      | Chile                                        | Araucania Region                                            | gPp     |
| cl-at      | Chile                                        | Atacama Region                                              | gPp     |
| cl-bi      | Chile                                        | Biobio Region                                               | gPp     |
| cl-co      | Chile                                        | Coquimbo Region                                             | gPp     |
| cl-li      | Chile                                        | O'Higgins Region                                            | gPp     |
| cl-ll      | Chile                                        | Los Lagos Region                                            | gPp     |
| cl-lr      | Chile                                        | Los Ríos                                                    | gPp     |
| cl-ma      | Chile                                        | Magallanes and Chilean Antarctica Region                    | gPp     |
| cl-ml      | Chile                                        | Maule Region                                                | gPp     |
| cl-nb      | Chile                                        | Nuble Region                                                | gPp     |
| cl-rm      | Chile                                        | Santiago Metropolitan Region                                | gPp     |
| cl-ta      | Chile                                        | Tarapacа Region                                             | gPp     |
| cl-vs      | Chile                                        | Valparaiso Region                                           | gPp     |
| cm         |                                              | Cameroon                                                    | gPp     |
| cm-ad      | Cameroon                                     | Adamawa                                                     | gPp     |
| cm-ce      | Cameroon                                     | Centre                                                      | gPp     |
| cm-en      | Cameroon                                     | Far-North                                                   | gPp     |
| cm-es      | Cameroon                                     | East                                                        | gPp     |
| cm-lt      | Cameroon                                     | Littoral                                                    | gPp     |
| cm-no      | Cameroon                                     | North                                                       | gPp     |
| cm-nw      | Cameroon                                     | Northwest                                                   | gPp     |
| cm-ou      | Cameroon                                     | West                                                        | gPp     |
| cm-su      | Cameroon                                     | South                                                       | gPp     |
| cm-sw      | Cameroon                                     | Southwest                                                   | gPp     |
| cn         |                                              | China                                                       | gPp     |
| cn-ah      | China                                        | Anhui                                                       | gPp     |
| cn-bj      | China                                        | Beijing                                                     | gPp     |
| cn-cq      | China                                        | Chongqing                                                   | gPp     |
| cn-fj      | China                                        | Fujian                                                      | gPp     |
| cn-gd      | China                                        | Guangdong                                                   | gPp     |
| cn-gs      | China                                        | Gansu                                                       | gPp     |
| cn-gx      | China                                        | Guangxi                                                     | gPp     |
| cn-gz      | China                                        | Guizhou                                                     | gPp     |
| cn-ha      | China                                        | Henan                                                       | gPp     |
| cn-hb      | China                                        | Hubei                                                       | gPp     |
| cn-he      | China                                        | Hebei                                                       | gPp     |
| cn-hi      | China                                        | Hainan                                                      | gPp     |
| cn-hk      | China                                        | Hong Kong                                                   | gPp     |
| cn-hl      | China                                        | Heilongjiang                                                | gPp     |
| cn-hn      | China                                        | Hunan                                                       | gPp     |
| cn-jl      | China                                        | Jilin                                                       | gPp     |
| cn-js      | China                                        | Jiangsu                                                     | gPp     |
| cn-jx      | China                                        | Jiangxi                                                     | gPp     |
| cn-ln      | China                                        | Liaoning                                                    | gPp     |
| cn-mo      | China                                        | Macau                                                       | gPp     |
| cn-nm      | China                                        | Inner Mongolia                                              | gPp     |
| cn-nx      | China                                        | Ningxia                                                     | gPp     |
| cn-qh      | China                                        | Qinghai                                                     | gPp     |
| cn-sc      | China                                        | Sichuan                                                     | gPp     |
| cn-sd      | China                                        | Shandong                                                    | gPp     |
| cn-sh      | China                                        | Shanghai                                                    | gPp     |
| cn-sn      | China                                        | Shaanxi                                                     | gPp     |
| cn-sx      | China                                        | Shanxi                                                      | gPp     |
| cn-tj      | China                                        | Tianjin                                                     | gPp     |
| cn-xj      | China                                        | Xinjiang                                                    | gPp     |
| cn-xz      | China                                        | Xizang                                                      | gPp     |
| cn-yn      | China                                        | Yunnan                                                      | gPp     |
| cn-zj      | China                                        | Zhejiang                                                    | gPp     |
| co         |                                              | Colombia                                                    | gPp     |
| co-ama     | Colombia                                     | Amazonas                                                    | gPp     |
| co-ant     | Colombia                                     | Antioquia                                                   | gPp     |
| co-ara     | Colombia                                     | Arauca                                                      | gPp     |
| co-atl     | Colombia                                     | Atlántico                                                   | gPp     |
| co-bol     | Colombia                                     | Bolívar                                                     | gPp     |
| co-boy     | Colombia                                     | Boyacá                                                      | gPp     |
| co-cal     | Colombia                                     | Caldas                                                      | gPp     |
| co-caq     | Colombia                                     | Caquetá                                                     | gPp     |
| co-cas     | Colombia                                     | Casanare                                                    | gPp     |
| co-cau     | Colombia                                     | Cauca                                                       | gPp     |
| co-ces     | Colombia                                     | Cesar                                                       | gPp     |
| co-cho     | Colombia                                     | Chocó                                                       | gPp     |
| co-cor     | Colombia                                     | Córdoba                                                     | gPp     |
| co-cun     | Colombia                                     | Cundinamarca                                                | gPp     |
| co-dc      | Colombia                                     | Bogota, Capital District                                    | gPp     |
| co-gua     | Colombia                                     | Guainía                                                     | gPp     |
| co-guv     | Colombia                                     | Guaviare                                                    | gPp     |
| co-hui     | Colombia                                     | Huila                                                       | gPp     |
| co-lag     | Colombia                                     | La Guajira                                                  | gPp     |
| co-mag     | Colombia                                     | Magdalena                                                   | gPp     |
| co-met     | Colombia                                     | Meta                                                        | gPp     |
| co-nar     | Colombia                                     | Nariño                                                      | gPp     |
| co-nsa     | Colombia                                     | Norte de Santander                                          | gPp     |
| co-put     | Colombia                                     | Putumayo                                                    | gPp     |
| co-qui     | Colombia                                     | Quindío                                                     | gPp     |
| co-ris     | Colombia                                     | Risaralda                                                   | gPp     |
| co-san     | Colombia                                     | Santander                                                   | gPp     |
| co-sap     | Colombia                                     | Archipelago of San Andrés, Providencia and Santa Catalina   | gPp     |
| co-suc     | Colombia                                     | Sucre                                                       | gPp     |
| co-tol     | Colombia                                     | Tolima                                                      | gPp     |
| co-vac     | Colombia                                     | Valle del Cauca                                             | gPp     |
| co-vau     | Colombia                                     | Vaupés                                                      | gPp     |
| co-vid     | Colombia                                     | Vichada                                                     | gPp     |
| cr         |                                              | Costa Rica                                                  | gPp     |
| cr-a       | Costa Rica                                   | Alajuela Province                                           | gPp     |
| cr-c       | Costa Rica                                   | Cartago Province                                            | gPp     |
| cr-g       | Costa Rica                                   | Guanacaste                                                  | gPp     |
| cr-h       | Costa Rica                                   | Heredia Province                                            | gPp     |
| cr-l       | Costa Rica                                   | Limón Province                                              | gPp     |
| cr-p       | Costa Rica                                   | Puntarenas Province                                         | gPp     |
| cr-sj      | Costa Rica                                   | San Jose Province                                           | gPp     |
| cu         |                                              | Cuba                                                        | gPp     |
| cu-01      | Cuba                                         | Pinar del Rio                                               | gPp     |
| cu-03      | Cuba                                         | Havana                                                      | gPp     |
| cu-04      | Cuba                                         | Matanzas                                                    | gPp     |
| cu-05      | Cuba                                         | Villa Clara                                                 | gPp     |
| cu-06      | Cuba                                         | Cienfuegos                                                  | gPp     |
| cu-07      | Cuba                                         | Sancti Spiritus                                             | gPp     |
| cu-08      | Cuba                                         | Ciego de Avila                                              | gPp     |
| cu-09      | Cuba                                         | Camagüey                                                    | gPp     |
| cu-10      | Cuba                                         | Las Tunas                                                   | gPp     |
| cu-11      | Cuba                                         | Holguín                                                     | gPp     |
| cu-12      | Cuba                                         | Granma                                                      | gPp     |
| cu-13      | Cuba                                         | Santiago de Cuba                                            | gPp     |
| cu-14      | Cuba                                         | Guantánamo                                                  | gPp     |
| cu-15      | Cuba                                         | Artemisa                                                    | gPp     |
| cu-16      | Cuba                                         | Mayabeque                                                   | gPp     |
| cu-99      | Cuba                                         | Isle of Youth                                               | gPp     |
| cv         |                                              | Cape Verde                                                  | gPp     |
| cv-br      | Cape Verde                                   | Brava                                                       | gPp     |
| cv-bv      | Cape Verde                                   | Boa Vista                                                   | gPp     |
| cv-ca      | Cape Verde                                   | Santa Catarina                                              | gPp     |
| cv-cf      | Cape Verde                                   | Santa Catarina do Fogo                                      | gPp     |
| cv-cr      | Cape Verde                                   | Santa Cruz                                                  | gPp     |
| cv-ma      | Cape Verde                                   | Maio                                                        | gPp     |
| cv-mo      | Cape Verde                                   | Mosteiros                                                   | gPp     |
| cv-pa      | Cape Verde                                   | Paul                                                        | gPp     |
| cv-pn      | Cape Verde                                   | Porto Novo                                                  | gPp     |
| cv-pr      | Cape Verde                                   | Praia                                                       | gPp     |
| cv-rb      | Cape Verde                                   | Ribeira Brava                                               | gPp     |
| cv-rg      | Cape Verde                                   | Ribeira Grande                                              | gPp     |
| cv-rs      | Cape Verde                                   | Ribeira Grande de Santiago                                  | gPp     |
| cv-sd      | Cape Verde                                   | São Domingos                                                | gPp     |
| cv-sf      | Cape Verde                                   | São Filipe                                                  | gPp     |
| cv-sl      | Cape Verde                                   | Sal                                                         | gPp     |
| cv-sm      | Cape Verde                                   | São Miguel                                                  | gPp     |
| cv-so      | Cape Verde                                   | São Lourenço dos Órgãos                                     | gPp     |
| cv-ss      | Cape Verde                                   | São Salvador do Mundo                                       | gPp     |
| cv-sv      | Cape Verde                                   | São Vicente                                                 | gPp     |
| cv-ta      | Cape Verde                                   | Tarrafal                                                    | gPp     |
| cv-ts      | Cape Verde                                   | Tarrafal de São Nicolau                                     | gPp     |
| cy         |                                              | Cyprus                                                      | gPp     |
| cy-01      | Cyprus                                       | Nicosia District                                            | gPp     |
| cy-02      | Cyprus                                       | Limassol District                                           | gPp     |
| cy-03      | Cyprus                                       | Larnaca District                                            | gPp     |
| cy-04      | Cyprus                                       | Famagusta District                                          | gPp     |
| cy-05      | Cyprus                                       | Paphos District                                             | gPp     |
| cy-xx      | Cyprus                                       | Northern Cyprus                                             | gPp     |
| cz         |                                              | Czechia                                                     | gPp     |
| cz-10      | Czechia                                      | Prague                                                      | gPp     |
| cz-20      | Czechia                                      | Central Bohemian Region                                     | gPp     |
| cz-31      | Czechia                                      | South Bohemian Region                                       | gPp     |
| cz-32      | Czechia                                      | Plzeň Region                                                | gPp     |
| cz-41      | Czechia                                      | Karlovy Vary Region                                         | gPp     |
| cz-42      | Czechia                                      | Ústí nad Labem Region                                       | gPp     |
| cz-51      | Czechia                                      | Liberec Region                                              | gPp     |
| cz-52      | Czechia                                      | Hradec Králové Region                                       | gPp     |
| cz-53      | Czechia                                      | Pardubice Region                                            | gPp     |
| cz-63      | Czechia                                      | Vysočina Region                                             | gPp     |
| cz-64      | Czechia                                      | South Moravian Region                                       | gPp     |
| cz-71      | Czechia                                      | Olomouc Region                                              | gPp     |
| cz-72      | Czechia                                      | Zlín Region                                                 | gPp     |
| cz-80      | Czechia                                      | Moravian-Silesian Region                                    | gPp     |
| da         |                                              | Akrotiri and Dhekelia                                       | gPp     |
| de         |                                              | Germany                                                     | gPp     |
| de-bb      | Germany                                      | Brandenburg                                                 | gPp     |
| de-be      | Germany                                      | Berlin                                                      | gPp     |
| de-bw      | Germany                                      | Baden-Württemberg                                           | gPp     |
| de-by      | Germany                                      | Bavaria                                                     | gPp     |
| de-hb      | Germany                                      | Bremen                                                      | gPp     |
| de-he      | Germany                                      | Hesse                                                       | gPp     |
| de-hh      | Germany                                      | Hamburg                                                     | gPp     |
| de-mv      | Germany                                      | Mecklenburg-Vorpommern                                      | gPp     |
| de-ni      | Germany                                      | Lower Saxony                                                | gPp     |
| de-nw      | Germany                                      | North Rhine-Westphalia                                      | gPp     |
| de-rp      | Germany                                      | Rhineland-Palatinate                                        | gPp     |
| de-sh      | Germany                                      | Schleswig-Holstein                                          | gPp     |
| de-sl      | Germany                                      | Saarland                                                    | gPp     |
| de-sn      | Germany                                      | Saxony                                                      | gPp     |
| de-st      | Germany                                      | Saxony-Anhalt                                               | gPp     |
| de-th      | Germany                                      | Thuringia                                                   | gPp     |
| dj         |                                              | Djibouti                                                    | gPp     |
| dj-ar      | Djibouti                                     | Arta                                                        | gPp     |
| dj-as      | Djibouti                                     | Ali Sabieh                                                  | gPp     |
| dj-di      | Djibouti                                     | Dikhil                                                      | gPp     |
| dj-dj      | Djibouti                                     | Djibouti                                                    | gPp     |
| dj-ob      | Djibouti                                     | Obock                                                       | gPp     |
| dj-ta      | Djibouti                                     | Tadjourah                                                   | gPp     |
| dk         |                                              | Denmark                                                     | gPp     |
| dk-81      | Denmark                                      | North Denmark Region                                        | gPp     |
| dk-82      | Denmark                                      | Central Denmark Region                                      | gPp     |
| dk-83      | Denmark                                      | Region of Southern Denmark                                  | gPp     |
| dk-84      | Denmark                                      | Capital Region of Denmark                                   | gPp     |
| dk-85      | Denmark                                      | Region Zealand                                              | gPp     |
| dm         |                                              | Dominica                                                    | gPp     |
| dm-02      | Dominica                                     | Saint Andrew Parish                                         | gPp     |
| dm-03      | Dominica                                     | Saint David Parish                                          | gPp     |
| dm-04      | Dominica                                     | Saint George Parish                                         | gPp     |
| dm-05      | Dominica                                     | Saint John Parish                                           | gPp     |
| dm-06      | Dominica                                     | Saint Joseph Parish                                         | gPp     |
| dm-07      | Dominica                                     | Saint Luke Parish                                           | gPp     |
| dm-08      | Dominica                                     | Saint Mark Parish                                           | gPp     |
| dm-09      | Dominica                                     | Saint Patrick Parish                                        | gPp     |
| dm-10      | Dominica                                     | Saint Paul Parish                                           | gPp     |
| dm-11      | Dominica                                     | Saint Peter Parish                                          | gPp     |
| do         |                                              | Dominican Republic                                          | gPp     |
| dz         |                                              | Algeria                                                     | gPp     |
| dz-01      | Algeria                                      | Adrar                                                       | gPp     |
| dz-02      | Algeria                                      | Chlef                                                       | gPp     |
| dz-03      | Algeria                                      | Laghouat                                                    | gPp     |
| dz-04      | Algeria                                      | Oum El Bouaghi                                              | gPp     |
| dz-05      | Algeria                                      | Batna                                                       | gPp     |
| dz-06      | Algeria                                      | Bejaia                                                      | gPp     |
| dz-07      | Algeria                                      | Biskra                                                      | gPp     |
| dz-08      | Algeria                                      | Bechar                                                      | gPp     |
| dz-09      | Algeria                                      | Blida                                                       | gPp     |
| dz-10      | Algeria                                      | Bouira                                                      | gPp     |
| dz-11      | Algeria                                      | Tamanrasset                                                 | gPp     |
| dz-12      | Algeria                                      | Tébessa                                                     | gPp     |
| dz-13      | Algeria                                      | Tlemcen                                                     | gPp     |
| dz-14      | Algeria                                      | Tiaret                                                      | gPp     |
| dz-15      | Algeria                                      | Tizi Ouzou                                                  | gPp     |
| dz-16      | Algeria                                      | Algiers                                                     | gPp     |
| dz-17      | Algeria                                      | Djelfa                                                      | gPp     |
| dz-18      | Algeria                                      | Jijel                                                       | gPp     |
| dz-19      | Algeria                                      | Setif                                                       | gPp     |
| dz-20      | Algeria                                      | Saïda                                                       | gPp     |
| dz-21      | Algeria                                      | Skikda                                                      | gPp     |
| dz-22      | Algeria                                      | Sidi Bel Abbès                                              | gPp     |
| dz-23      | Algeria                                      | Annaba                                                      | gPp     |
| dz-24      | Algeria                                      | Guelma                                                      | gPp     |
| dz-25      | Algeria                                      | Constantine                                                 | gPp     |
| dz-26      | Algeria                                      | Médéa                                                       | gPp     |
| dz-27      | Algeria                                      | Mostaganem                                                  | gPp     |
| dz-28      | Algeria                                      | M'Sila                                                      | gPp     |
| dz-29      | Algeria                                      | Mascara                                                     | gPp     |
| dz-30      | Algeria                                      | Ouargla                                                     | gPp     |
| dz-31      | Algeria                                      | Oran                                                        | gPp     |
| dz-32      | Algeria                                      | El Bayadh                                                   | gPp     |
| dz-33      | Algeria                                      | Illizi                                                      | gPp     |
| dz-34      | Algeria                                      | Bordj Bou Arreridj                                          | gPp     |
| dz-35      | Algeria                                      | Boumerdes                                                   | gPp     |
| dz-36      | Algeria                                      | El Tarf                                                     | gPp     |
| dz-37      | Algeria                                      | Tinduf                                                      | gPp     |
| dz-38      | Algeria                                      | Tissemsilt                                                  | gPp     |
| dz-39      | Algeria                                      | El Oued                                                     | gPp     |
| dz-40      | Algeria                                      | Khenchela                                                   | gPp     |
| dz-41      | Algeria                                      | Souk Ahras                                                  | gPp     |
| dz-42      | Algeria                                      | Tipaza                                                      | gPp     |
| dz-43      | Algeria                                      | Mila                                                        | gPp     |
| dz-44      | Algeria                                      | Aïn Defla                                                   | gPp     |
| dz-45      | Algeria                                      | Naâma                                                       | gPp     |
| dz-46      | Algeria                                      | Aïn Témouchent                                              | gPp     |
| dz-47      | Algeria                                      | Ghardaia                                                    | gPp     |
| dz-48      | Algeria                                      | Relizane                                                    | gPp     |
| dz-49      | Algeria                                      | Timimoun Province                                           | gPp     |
| dz-50      | Algeria                                      | Bordj Badji Mokhtar                                         | gPp     |
| dz-51      | Algeria                                      | Ouled Djellal                                               | gPp     |
| dz-52      | Algeria                                      | Beni Abbes                                                  | gPp     |
| dz-53      | Algeria                                      | In Salah                                                    | gPp     |
| dz-54      | Algeria                                      | In Guezzam                                                  | gPp     |
| dz-55      | Algeria                                      | Touggourt                                                   | gPp     |
| dz-56      | Algeria                                      | Djanet                                                      | gPp     |
| dz-57      | Algeria                                      | El M'Ghair                                                  | gPp     |
| dz-58      | Algeria                                      | El Menia                                                    | gPp     |
| dz-59      | Algeria                                      | Aflou                                                       | gPp     |
| dz-60      | Algeria                                      | Barika                                                      | gPp     |
| dz-61      | Algeria                                      | El Kantara                                                  | gPp     |
| dz-62      | Algeria                                      | Bir El Ater                                                 | gPp     |
| dz-64      | Algeria                                      | Ksar Chellala                                               | gPp     |
| dz-65      | Algeria                                      | Aïn Oussara                                                 | gPp     |
| dz-66      | Algeria                                      | Messaad                                                     | gPp     |
| dz-67      | Algeria                                      | Ksar El Boukhari Province                                   | gPp     |
| dz-68      | Algeria                                      | Bou Saâda Province                                          | gPp     |
| dz-69      | Algeria                                      | El Abiodh Sidi Cheikh Province                              | gPp     |
| ec         |                                              | Ecuador                                                     | gPp     |
| ec-a       | Ecuador                                      | Azuay                                                       | gPp     |
| ec-b       | Ecuador                                      | Bolívar                                                     | gPp     |
| ec-c       | Ecuador                                      | Carchi                                                      | gPp     |
| ec-d       | Ecuador                                      | Orellana                                                    | gPp     |
| ec-e       | Ecuador                                      | Esmeraldas                                                  | gPp     |
| ec-f       | Ecuador                                      | Cañar                                                       | gPp     |
| ec-g       | Ecuador                                      | Guayas                                                      | gPp     |
| ec-h       | Ecuador                                      | Chimborazo                                                  | gPp     |
| ec-i       | Ecuador                                      | Imbabura                                                    | gPp     |
| ec-l       | Ecuador                                      | Loja                                                        | gPp     |
| ec-m       | Ecuador                                      | Manabí                                                      | gPp     |
| ec-n       | Ecuador                                      | Napo                                                        | gPp     |
| ec-o       | Ecuador                                      | El Oro                                                      | gPp     |
| ec-p       | Ecuador                                      | Pichincha                                                   | gPp     |
| ec-r       | Ecuador                                      | Los Ríos                                                    | gPp     |
| ec-s       | Ecuador                                      | Morona Santiago                                             | gPp     |
| ec-sd      | Ecuador                                      | Santo Domingo de los Tsáchilas                              | gPp     |
| ec-se      | Ecuador                                      | Santa Elena Province                                        | gPp     |
| ec-t       | Ecuador                                      | Tungurahua                                                  | gPp     |
| ec-u       | Ecuador                                      | Sucumbíos                                                   | gPp     |
| ec-w       | Ecuador                                      | Galápagos                                                   | gPp     |
| ec-x       | Ecuador                                      | Cotopaxi                                                    | gPp     |
| ec-y       | Ecuador                                      | Pastaza                                                     | gPp     |
| ec-z       | Ecuador                                      | Zamora Chinchipe                                            | gPp     |
| ee         |                                              | Estonia                                                     | gPp     |
| ee-37      | Estonia                                      | Harju County                                                | gPp     |
| ee-39      | Estonia                                      | Hiiu County                                                 | gPp     |
| ee-45      | Estonia                                      | Ida-Viru County                                             | gPp     |
| ee-50      | Estonia                                      | Jõgeva County                                               | gPp     |
| ee-52      | Estonia                                      | Järva County                                                | gPp     |
| ee-56      | Estonia                                      | Lääne County                                                | gPp     |
| ee-60      | Estonia                                      | Lääne-Viru County                                           | gPp     |
| ee-64      | Estonia                                      | Põlva County                                                | gPp     |
| ee-68      | Estonia                                      | Pärnu County                                                | gPp     |
| ee-71      | Estonia                                      | Rapla County                                                | gPp     |
| ee-74      | Estonia                                      | Saare County                                                | gPp     |
| ee-79      | Estonia                                      | Tartu County                                                | gPp     |
| ee-81      | Estonia                                      | Valga County                                                | gPp     |
| ee-84      | Estonia                                      | Viljandi County                                             | gPp     |
| ee-87      | Estonia                                      | Võru County                                                 | gPp     |
| eg         |                                              | Egypt                                                       | gPp     |
| eg-alx     | Egypt                                        | Alexandria                                                  | gPp     |
| eg-asn     | Egypt                                        | Aswan                                                       | gPp     |
| eg-ast     | Egypt                                        | Asyut                                                       | gPp     |
| eg-ba      | Egypt                                        | Red Sea                                                     | gPp     |
| eg-bh      | Egypt                                        | Lake                                                        | gPp     |
| eg-bns     | Egypt                                        | Bani Sweif                                                  | gPp     |
| eg-c       | Egypt                                        | Cairo                                                       | gPp     |
| eg-dk      | Egypt                                        | Ad Daqahliyya                                               | gPp     |
| eg-dt      | Egypt                                        | Damietta                                                    | gPp     |
| eg-fym     | Egypt                                        | Faiyum                                                      | gPp     |
| eg-gh      | Egypt                                        | Western                                                     | gPp     |
| eg-gz      | Egypt                                        | Aj Jiza                                                     | gPp     |
| eg-is      | Egypt                                        | Al Ismailiya                                                | gPp     |
| eg-js      | Egypt                                        | South Sinai                                                 | gPp     |
| eg-kb      | Egypt                                        | Al Qalyubiya                                                | gPp     |
| eg-kfs     | Egypt                                        | Kafr El Sheikh                                              | gPp     |
| eg-kn      | Egypt                                        | Qena                                                        | gPp     |
| eg-lx      | Egypt                                        | Luxor                                                       | gPp     |
| eg-mn      | Egypt                                        | Al Minya                                                    | gPp     |
| eg-mnf     | Egypt                                        | El Minufiyya                                                | gPp     |
| eg-mt      | Egypt                                        | Matruh                                                      | gPp     |
| eg-pts     | Egypt                                        | Port Said                                                   | gPp     |
| eg-shg     | Egypt                                        | Suhaj                                                       | gPp     |
| eg-shr     | Egypt                                        | Eastern                                                     | gPp     |
| eg-sin     | Egypt                                        | North Sinai                                                 | gPp     |
| eg-suz     | Egypt                                        | Suez                                                        | gPp     |
| eg-wad     | Egypt                                        | New Valley                                                  | gPp     |
| el         |                                              | Ilemi Triangle                                              | gPp     |
| er         |                                              | Eritrea                                                     | gPp     |
| er-an      | Eritrea                                      | Anseba                                                      | gPp     |
| er-dk      | Eritrea                                      | Southern Red Sea Region                                     | gPp     |
| er-du      | Eritrea                                      | Debub Region                                                | gPp     |
| er-gb      | Eritrea                                      | Gash-Barka                                                  | gPp     |
| er-ma      | Eritrea                                      | Maekel Region                                               | gPp     |
| er-sk      | Eritrea                                      | Northern Red Sea Region                                     | gPp     |
| es         |                                              | Spain                                                       | gPp     |
| es-an      | Spain                                        | Andalusia                                                   | gPp     |
| es-ar      | Spain                                        | Aragon                                                      | gPp     |
| es-as      | Spain                                        | Asturias                                                    | gPp     |
| es-cb      | Spain                                        | Cantabria                                                   | gPp     |
| es-ce      | Spain                                        | Ceuta                                                       | gPp     |
| es-cl      | Spain                                        | Castile and León                                            | gPp     |
| es-cm      | Spain                                        | Castile-La Mancha                                           | gPp     |
| es-cn      | Spain                                        | Canary Islands                                              | gPp     |
| es-ct      | Spain                                        | Catalonia                                                   | gPp     |
| es-ex      | Spain                                        | Extremadura                                                 | gPp     |
| es-ga      | Spain                                        | Galicia                                                     | gPp     |
| es-ib      | Spain                                        | Balearic Islands                                            | gPp     |
| es-mc      | Spain                                        | Region of Murcia                                            | gPp     |
| es-md      | Spain                                        | Community of Madrid                                         | gPp     |
| es-ml      | Spain                                        | Melilla                                                     | gPp     |
| es-nc      | Spain                                        | Navarre                                                     | gPp     |
| es-pv      | Spain                                        | Autonomous Community of the Basque Country                  | gPp     |
| es-ri      | Spain                                        | Rioja                                                       | gPp     |
| es-vc      | Spain                                        | Valencian Community                                         | gPp     |
| et         |                                              | Ethiopia                                                    | gPp     |
| et-aa      | Ethiopia                                     | Addis Ababa                                                 | gPp     |
| et-af      | Ethiopia                                     | Afar Region                                                 | gPp     |
| et-am      | Ethiopia                                     | Amhara Region                                               | gPp     |
| et-be      | Ethiopia                                     | Benishangul-Gumuz Region                                    | gPp     |
| et-ce      | Ethiopia                                     | Central Ethiopia Regional State                             | gPp     |
| et-dd      | Ethiopia                                     | Dire Dawa                                                   | gPp     |
| et-ga      | Ethiopia                                     | Gambela Region                                              | gPp     |
| et-ha      | Ethiopia                                     | Harar                                                       | gPp     |
| et-or      | Ethiopia                                     | Oromia Region                                               | gPp     |
| et-se      | Ethiopia                                     | South Ethiopia Regional State                               | gPp     |
| et-si      | Ethiopia                                     | Sidama                                                      | gPp     |
| et-so      | Ethiopia                                     | Somali Region                                               | gPp     |
| et-sw      | Ethiopia                                     | South West Ethiopia Peoples                                 | gPp     |
| et-ti      | Ethiopia                                     | Tigray                                                      | gPp     |
| ex         |                                              | Sahrawi Arab Democratic Republic                            | gPp     |
| fi         |                                              | Finland                                                     | gPp     |
| fi-01      | Finland                                      | Åland Islands                                               | gPp     |
| fi-02      | Finland                                      | South Karelia                                               | gPp     |
| fi-03      | Finland                                      | South Ostrobothnia                                          | gPp     |
| fi-04      | Finland                                      | South Savo                                                  | gPp     |
| fi-05      | Finland                                      | Kainuu                                                      | gPp     |
| fi-06      | Finland                                      | Kanta-Häme                                                  | gPp     |
| fi-07      | Finland                                      | Central Ostrobothnia                                        | gPp     |
| fi-08      | Finland                                      | Central Finland                                             | gPp     |
| fi-09      | Finland                                      | Kymenlaakso                                                 | gPp     |
| fi-10      | Finland                                      | Lapland                                                     | gPp     |
| fi-11      | Finland                                      | Pirkanmaa                                                   | gPp     |
| fi-12      | Finland                                      | Ostrobothnia                                                | gPp     |
| fi-13      | Finland                                      | North Karelia                                               | gPp     |
| fi-14      | Finland                                      | North Ostrobothnia                                          | gPp     |
| fi-15      | Finland                                      | North Savo                                                  | gPp     |
| fi-16      | Finland                                      | Päijät-Häme                                                 | gPp     |
| fi-17      | Finland                                      | Satakunta                                                   | gPp     |
| fi-18      | Finland                                      | Uusimaa                                                     | gPp     |
| fi-19      | Finland                                      | Southwest Finland                                           | gPp     |
| fj         |                                              | Fiji                                                        | gPp     |
| fj-c       | Fiji                                         | Central                                                     | gPp     |
| fj-e       | Fiji                                         | Eastern                                                     | gPp     |
| fj-n       | Fiji                                         | Northern                                                    | gPp     |
| fj-r       | Fiji                                         | Rotuma                                                      | gPp     |
| fj-w       | Fiji                                         | Western                                                     | gPp     |
| fk         |                                              | Falkland Islands                                            | gPp     |
| fm         |                                              | Federated States of Micronesia                              | gPp     |
| fm-ksa     | Federated States of Micronesia               | Kosrae                                                      | gPp     |
| fm-pni     | Federated States of Micronesia               | Pohnpei                                                     | gPp     |
| fm-trk     | Federated States of Micronesia               | Chuuk                                                       | gPp     |
| fm-yap     | Federated States of Micronesia               | Yap                                                         | gPp     |
| fo         |                                              | Faroe Islands                                               | gPp     |
| fr         |                                              | France                                                      | gPp     |
| fr-20r     | France                                       | Corsica                                                     | gPp     |
| fr-971     | France                                       | Guadeloupe                                                  | gPp     |
| fr-972     | France                                       | Martinique                                                  | gPp     |
| fr-973     | France                                       | French Guiana                                               | gPp     |
| fr-974     | France                                       | Réunion                                                     | gPp     |
| fr-976     | France                                       | Mayotte                                                     | gPp     |
| fr-ara     | France                                       | Auvergne-Rhône-Alpes                                        | gPp     |
| fr-bfc     | France                                       | Bourgogne – Franche-Comté                                   | gPp     |
| fr-bl      | France                                       | Saint Barthélemy                                            | gPp     |
| fr-bre     | France                                       | Brittany                                                    | gPp     |
| fr-cp      | France                                       | Clipperton Island                                           | gPp     |
| fr-cvl     | France                                       | Centre-Val de Loire                                         | gPp     |
| fr-ges     | France                                       | Grand Est                                                   | gPp     |
| fr-hdf     | France                                       | Hauts-de-France                                             | gPp     |
| fr-idf     | France                                       | Ile-de-France                                               | gPp     |
| fr-mf      | France                                       | Saint Martin                                                | gPp     |
| fr-naq     | France                                       | Nouvelle-Aquitaine                                          | gPp     |
| fr-nc      | France                                       | New Caledonia                                               | gPp     |
| fr-nor     | France                                       | Normandy                                                    | gPp     |
| fr-occ     | France                                       | Occitania                                                   | gPp     |
| fr-pac     | France                                       | Provence-Alpes-Côte d'Azur                                  | gPp     |
| fr-pdl     | France                                       | Pays de la Loire                                            | gPp     |
| fr-pf      | France                                       | French Polynesia                                            | gPp     |
| fr-pm      | France                                       | Saint Pierre and Miquelon                                   | gPp     |
| fr-tf      | France                                       | French Southern and Antarctic Lands                         | gPp     |
| fr-wf      | France                                       | Wallis and Futuna                                           | gPp     |
| ga         |                                              | Gabon                                                       | gPp     |
| ga-1       | Gabon                                        | Estuaire Province                                           | gPp     |
| ga-2       | Gabon                                        | Haut-Ogooué Province                                        | gPp     |
| ga-3       | Gabon                                        | Moyen-Ogooué Province                                       | gPp     |
| ga-4       | Gabon                                        | Ngounié Province                                            | gPp     |
| ga-5       | Gabon                                        | Nyanga Province                                             | gPp     |
| ga-6       | Gabon                                        | Ogooué-Ivindo                                               | gPp     |
| ga-7       | Gabon                                        | Ogooué-Lolo Province                                        | gPp     |
| ga-8       | Gabon                                        | Ogooué-Maritime Province                                    | gPp     |
| ga-9       | Gabon                                        | Woleu-Ntem                                                  | gPp     |
| gb         |                                              | United Kingdom                                              | gPp     |
| gb-eng     | United Kingdom                               | England                                                     | gPp     |
| gb-nir     | United Kingdom                               | Northern Ireland                                            | gPp     |
| gb-sct     | United Kingdom                               | Scotland                                                    | gPp     |
| gb-wls     | United Kingdom                               | Wales                                                       | gPp     |
| gd         |                                              | Grenada                                                     | gPp     |
| gd-01      | Grenada                                      | Saint Andrew                                                | gPp     |
| gd-02      | Grenada                                      | Saint David                                                 | gPp     |
| gd-03      | Grenada                                      | Saint George                                                | gPp     |
| gd-04      | Grenada                                      | Saint John                                                  | gPp     |
| gd-05      | Grenada                                      | Saint Mark                                                  | gPp     |
| gd-06      | Grenada                                      | Saint Patrick                                               | gPp     |
| gd-10      | Grenada                                      | Carriacou and Petite Martinique                             | gPp     |
| ge         |                                              | Georgia                                                     | gPp     |
| ge-ab      | Georgia                                      | Abkhazia                                                    | gPp     |
| ge-aj      | Georgia                                      | Autonomous Republic of Adjara                               | gPp     |
| ge-gu      | Georgia                                      | Guria                                                       | gPp     |
| ge-im      | Georgia                                      | Imereti                                                     | gPp     |
| ge-ka      | Georgia                                      | Kakheti                                                     | gPp     |
| ge-kk      | Georgia                                      | Lower Kartli                                                | gPp     |
| ge-mm      | Georgia                                      | Mtskheta-Mtianeti                                           | gPp     |
| ge-rl      | Georgia                                      | Racha-Lechkhumi and Kvemo Svaneti                           | gPp     |
| ge-sj      | Georgia                                      | Samtskhe-Javakheti                                          | gPp     |
| ge-sk      | Georgia                                      | Shida Kartli                                                | gPp     |
| ge-sz      | Georgia                                      | Samegrelo-Upper Svaneti                                     | gPp     |
| ge-tb      | Georgia                                      | Tbilisi                                                     | gPp     |
| ge-xx      | Georgia                                      | South Ossetia                                               | gPp     |
| gg         |                                              | Guernsey                                                    | gPp     |
| gh         |                                              | Ghana                                                       | gPp     |
| gh-aa      | Ghana                                        | Greater Accra Region                                        | gPp     |
| gh-af      | Ghana                                        | Ahafo Region                                                | gPp     |
| gh-ah      | Ghana                                        | Ashanti Region                                              | gPp     |
| gh-be      | Ghana                                        | Bono East Region                                            | gPp     |
| gh-bo      | Ghana                                        | Bono Region                                                 | gPp     |
| gh-cp      | Ghana                                        | Central Region                                              | gPp     |
| gh-ep      | Ghana                                        | Eastern Region                                              | gPp     |
| gh-ne      | Ghana                                        | North East Region                                           | gPp     |
| gh-np      | Ghana                                        | Northern Region                                             | gPp     |
| gh-ot      | Ghana                                        | Oti Region                                                  | gPp     |
| gh-sv      | Ghana                                        | Savannah Region                                             | gPp     |
| gh-tv      | Ghana                                        | Volta Region                                                | gPp     |
| gh-ue      | Ghana                                        | Upper East Region                                           | gPp     |
| gh-uw      | Ghana                                        | Upper West Region                                           | gPp     |
| gh-wn      | Ghana                                        | Western North Region                                        | gPp     |
| gh-wp      | Ghana                                        | Western Region                                              | gPp     |
| gi         |                                              | Gibraltar                                                   | gPp     |
| gl         |                                              | Greenland                                                   | gPp     |
| gm         |                                              | The Gambia                                                  | gPp     |
| gm-b       | The Gambia                                   | Banjul                                                      | gPp     |
| gm-l       | The Gambia                                   | Lower River Division                                        | gPp     |
| gm-m       | The Gambia                                   | Central River Division                                      | gPp     |
| gm-n       | The Gambia                                   | North Bank Division                                         | gPp     |
| gm-u       | The Gambia                                   | Upper River Division                                        | gPp     |
| gm-w       | The Gambia                                   | West Coast Division                                         | gPp     |
| gn         |                                              | Guinea                                                      | gPp     |
| gn-b       | Guinea                                       | Boké Region                                                 | gPp     |
| gn-c       | Guinea                                       | Conakry                                                     | gPp     |
| gn-d       | Guinea                                       | Kindia Region                                               | gPp     |
| gn-f       | Guinea                                       | Faranah Region                                              | gPp     |
| gn-k       | Guinea                                       | Kankan Region                                               | gPp     |
| gn-l       | Guinea                                       | Labé Region                                                 | gPp     |
| gn-m       | Guinea                                       | Mamou Region                                                | gPp     |
| gn-n       | Guinea                                       | Nzérékoré Region                                            | gPp     |
| gq         |                                              | Equatorial Guinea                                           | gPp     |
| gq-c       | Equatorial Guinea                            | Región Continental                                          | gPp     |
| gq-i       | Equatorial Guinea                            | Región Insular                                              | gPp     |
| gr         |                                              | Greece                                                      | gPp     |
| gr-69      | Greece                                       | Autonomous Monastic State of the Holy Mountain              | gPp     |
| gr-a       | Greece                                       | Eastern Macedonia and Thrace                                | gPp     |
| gr-b       | Greece                                       | Central Macedonia                                           | gPp     |
| gr-c       | Greece                                       | Western Macedonia                                           | gPp     |
| gr-d       | Greece                                       | Epirus                                                      | gPp     |
| gr-e       | Greece                                       | Thessaly                                                    | gPp     |
| gr-el3     | Greece                                       | Attica                                                      | gPp     |
| gr-f       | Greece                                       | Ioanian Islands                                             | gPp     |
| gr-g       | Greece                                       | Western Greece                                              | gPp     |
| gr-h       | Greece                                       | Central Greece                                              | gPp     |
| gr-j       | Greece                                       | Peloponnese Region                                          | gPp     |
| gr-k       | Greece                                       | Northern Aegean                                             | gPp     |
| gr-l       | Greece                                       | South Aegean                                                | gPp     |
| gr-m       | Greece                                       | Region of Crete                                             | gPp     |
| gs         |                                              | South Georgia and the South Sandwich Islands                | gPp     |
| gt         |                                              | Guatemala                                                   | gPp     |
| gt-01      | Guatemala                                    | Guatemala Department                                        | gPp     |
| gt-02      | Guatemala                                    | El Progreso                                                 | gPp     |
| gt-03      | Guatemala                                    | Sacatepéquez                                                | gPp     |
| gt-04      | Guatemala                                    | Chimaltenango                                               | gPp     |
| gt-05      | Guatemala                                    | Escuintla                                                   | gPp     |
| gt-06      | Guatemala                                    | Santa Rosa                                                  | gPp     |
| gt-07      | Guatemala                                    | Sololá                                                      | gPp     |
| gt-08      | Guatemala                                    | Totonicapán                                                 | gPp     |
| gt-09      | Guatemala                                    | Quetzaltenango                                              | gPp     |
| gt-10      | Guatemala                                    | Suchitepéquez                                               | gPp     |
| gt-11      | Guatemala                                    | Retalhuleu                                                  | gPp     |
| gt-12      | Guatemala                                    | San Marcos                                                  | gPp     |
| gt-13      | Guatemala                                    | Huehuetenango                                               | gPp     |
| gt-14      | Guatemala                                    | Quiché                                                      | gPp     |
| gt-15      | Guatemala                                    | Baja Verapaz                                                | gPp     |
| gt-16      | Guatemala                                    | Alta Verapaz                                                | gPp     |
| gt-17      | Guatemala                                    | Petén                                                       | gPp     |
| gt-18      | Guatemala                                    | Izabal                                                      | gPp     |
| gt-19      | Guatemala                                    | Zacapa                                                      | gPp     |
| gt-20      | Guatemala                                    | Chiquimula                                                  | gPp     |
| gt-21      | Guatemala                                    | Jalapa                                                      | gPp     |
| gt-22      | Guatemala                                    | Jutiapa                                                     | gPp     |
| gw         |                                              | Guinea-Bissau                                               | gPp     |
| gw-bs      | Guinea-Bissau                                | Bissau Autonomous Sector                                    | gPp     |
| gw-l       | Guinea-Bissau                                | East                                                        | gPp     |
| gw-n       | Guinea-Bissau                                | North                                                       | gPp     |
| gw-s       | Guinea-Bissau                                | South                                                       | gPp     |
| gy         |                                              | Guyana                                                      | gPp     |
| gy-ba      | Guyana                                       | Barima-Waini                                                | gPp     |
| gy-cu      | Guyana                                       | Cuyuni-Mazaruni                                             | gPp     |
| gy-de      | Guyana                                       | Demerara-Mahaica                                            | gPp     |
| gy-eb      | Guyana                                       | East Berbice-Corentyne                                      | gPp     |
| gy-es      | Guyana                                       | Essequibo Islands                                           | gPp     |
| gy-ma      | Guyana                                       | Mahaica-Berbice                                             | gPp     |
| gy-pm      | Guyana                                       | Pomeroon-Supenaam                                           | gPp     |
| gy-pt      | Guyana                                       | Potaro-Siparuni                                             | gPp     |
| gy-ud      | Guyana                                       | Upper Demerara-Berbice                                      | gPp     |
| gy-ut      | Guyana                                       | Upper Takutu-Upper Essequibo                                | gPp     |
| hn         |                                              | Honduras                                                    | gPp     |
| hn-at      | Honduras                                     | Atlántida                                                   | gPp     |
| hn-ch      | Honduras                                     | Choluteca                                                   | gPp     |
| hn-cl      | Honduras                                     | Colón                                                       | gPp     |
| hn-cm      | Honduras                                     | Comayagua                                                   | gPp     |
| hn-cp      | Honduras                                     | Copán                                                       | gPp     |
| hn-cr      | Honduras                                     | Cortés                                                      | gPp     |
| hn-ep      | Honduras                                     | El Paraíso                                                  | gPp     |
| hn-fm      | Honduras                                     | Francisco Morazán                                           | gPp     |
| hn-gd      | Honduras                                     | Gracias a Dios                                              | gPp     |
| hn-ib      | Honduras                                     | Bay Islands                                                 | gPp     |
| hn-in      | Honduras                                     | Intibucá                                                    | gPp     |
| hn-le      | Honduras                                     | Lempira                                                     | gPp     |
| hn-lp      | Honduras                                     | La Paz                                                      | gPp     |
| hn-oc      | Honduras                                     | Ocotepeque                                                  | gPp     |
| hn-ol      | Honduras                                     | Olancho                                                     | gPp     |
| hn-sb      | Honduras                                     | Santa Bárbara                                               | gPp     |
| hn-va      | Honduras                                     | Valle                                                       | gPp     |
| hn-yo      | Honduras                                     | Yoro                                                        | gPp     |
| hr         |                                              | Croatia                                                     | gPp     |
| hr-01      | Croatia                                      | Zagreb County                                               | gPp     |
| hr-02      | Croatia                                      | Krapina-Zagorje County                                      | gPp     |
| hr-03      | Croatia                                      | Sisak-Moslavina County                                      | gPp     |
| hr-04      | Croatia                                      | Karlovac County                                             | gPp     |
| hr-05      | Croatia                                      | Varaždin County                                             | gPp     |
| hr-06      | Croatia                                      | Koprivnica-Križevci County                                  | gPp     |
| hr-07      | Croatia                                      | Bjelovar-Bilogora County                                    | gPp     |
| hr-08      | Croatia                                      | Primorje-Gorski Kotar County                                | gPp     |
| hr-09      | Croatia                                      | Lika-Senj County                                            | gPp     |
| hr-10      | Croatia                                      | Virovitica-Podravina County                                 | gPp     |
| hr-11      | Croatia                                      | Požega-Slavonia County                                      | gPp     |
| hr-12      | Croatia                                      | Brod-Posavina County                                        | gPp     |
| hr-13      | Croatia                                      | Zadar County                                                | gPp     |
| hr-14      | Croatia                                      | Osijek-Baranja County                                       | gPp     |
| hr-15      | Croatia                                      | Šibenik-Knin County                                         | gPp     |
| hr-16      | Croatia                                      | Vukovar-Srijem County                                       | gPp     |
| hr-17      | Croatia                                      | Split-Dalmatia County                                       | gPp     |
| hr-18      | Croatia                                      | Istria County                                               | gPp     |
| hr-19      | Croatia                                      | Dubrovnik-Neretva County                                    | gPp     |
| hr-20      | Croatia                                      | Međimurje County                                            | gPp     |
| hr-21      | Croatia                                      | City of Zagreb                                              | gPp     |
| ht         |                                              | Haiti                                                       | gPp     |
| ht-ar      | Haiti                                        | Artibonite Department                                       | gPp     |
| ht-ce      | Haiti                                        | Centre Department                                           | gPp     |
| ht-ga      | Haiti                                        | Département de la Grande-Anse                               | gPp     |
| ht-nd      | Haiti                                        | Département du Nord                                         | gPp     |
| ht-ne      | Haiti                                        | Nord-Est Department                                         | gPp     |
| ht-ni      | Haiti                                        | Département des Nippes                                      | gPp     |
| ht-no      | Haiti                                        | Nord-Ouest Department                                       | gPp     |
| ht-ou      | Haiti                                        | Département de l'Ouest                                      | gPp     |
| ht-sd      | Haiti                                        | Sud Department                                              | gPp     |
| ht-se      | Haiti                                        | Département du Sud-Est                                      | gPp     |
| hu         |                                              | Hungary                                                     | gPp     |
| hu-hu10    | Hungary                                      | Central Hungary                                             | gPp     |
| hu-hu21    | Hungary                                      | Central Transdanubia                                        | gPp     |
| hu-hu22    | Hungary                                      | Western Transdanubia                                        | gPp     |
| hu-hu23    | Hungary                                      | Southern Transdanubia                                       | gPp     |
| hu-hu31    | Hungary                                      | North Hungary                                               | gPp     |
| hu-hu32    | Hungary                                      | North Great Plain                                           | gPp     |
| hu-hu33    | Hungary                                      | South Great Plain                                           | gPp     |
| id         |                                              | Indonesia                                                   | gPp     |
| id-ac      | Indonesia                                    | Aceh                                                        | gPp     |
| id-ba      | Indonesia                                    | Bali                                                        | gPp     |
| id-bb      | Indonesia                                    | Bangka-Belitung Islands                                     | gPp     |
| id-be      | Indonesia                                    | Bengkulu                                                    | gPp     |
| id-bt      | Indonesia                                    | Banten                                                      | gPp     |
| id-go      | Indonesia                                    | Gorontalo                                                   | gPp     |
| id-ja      | Indonesia                                    | Jambi                                                       | gPp     |
| id-jb      | Indonesia                                    | West Java                                                   | gPp     |
| id-ji      | Indonesia                                    | East Java                                                   | gPp     |
| id-jk      | Indonesia                                    | Special Capital Region of Jakarta                           | gPp     |
| id-jt      | Indonesia                                    | Central Java                                                | gPp     |
| id-kb      | Indonesia                                    | West Kalimantan                                             | gPp     |
| id-ki      | Indonesia                                    | East Kalimantan                                             | gPp     |
| id-kr      | Indonesia                                    | Riau Islands                                                | gPp     |
| id-ks      | Indonesia                                    | South Kalimantan                                            | gPp     |
| id-kt      | Indonesia                                    | Central Kalimantan                                          | gPp     |
| id-ku      | Indonesia                                    | North Kalimantan                                            | gPp     |
| id-la      | Indonesia                                    | Lampung                                                     | gPp     |
| id-ma      | Indonesia                                    | Maluku                                                      | gPp     |
| id-mu      | Indonesia                                    | North Maluku                                                | gPp     |
| id-nb      | Indonesia                                    | West Nusa Tenggara                                          | gPp     |
| id-nt      | Indonesia                                    | East Nusa Tenggara                                          | gPp     |
| id-pa      | Indonesia                                    | Papua                                                       | gPp     |
| id-pb      | Indonesia                                    | West Papua                                                  | gPp     |
| id-pd      | Indonesia                                    | Southwest Papua                                             | gPp     |
| id-pe      | Indonesia                                    | Highland Papua                                              | gPp     |
| id-ps      | Indonesia                                    | South Papua                                                 | gPp     |
| id-pt      | Indonesia                                    | Central Papua                                               | gPp     |
| id-ri      | Indonesia                                    | Riau                                                        | gPp     |
| id-sa      | Indonesia                                    | North Sulawesi                                              | gPp     |
| id-sb      | Indonesia                                    | West Sumatra                                                | gPp     |
| id-sg      | Indonesia                                    | Southeast Sulawesi                                          | gPp     |
| id-sn      | Indonesia                                    | South Sulawesi                                              | gPp     |
| id-sr      | Indonesia                                    | West Sulawesi                                               | gPp     |
| id-ss      | Indonesia                                    | South Sumatra                                               | gPp     |
| id-st      | Indonesia                                    | Central Sulawesi                                            | gPp     |
| id-su      | Indonesia                                    | North Sumatra                                               | gPp     |
| id-yo      | Indonesia                                    | Special Region of Yogyakarta                                | gPp     |
| ie         |                                              | Ireland                                                     | gPp     |
| ie-c       | Ireland                                      | Connacht                                                    | gPp     |
| ie-l       | Ireland                                      | Leinster                                                    | gPp     |
| ie-m       | Ireland                                      | Munster                                                     | gPp     |
| ie-u       | Ireland                                      | Ulster                                                      | gPp     |
| il         |                                              | Israel                                                      | gPp     |
| il-d       | Israel                                       | South District                                              | gPp     |
| il-ha      | Israel                                       | Haifa District                                              | gPp     |
| il-jm      | Israel                                       | Jerusalem District                                          | gPp     |
| il-m       | Israel                                       | Center District                                             | gPp     |
| il-ta      | Israel                                       | Tel-Aviv District                                           | gPp     |
| il-z       | Israel                                       | North District                                              | gPp     |
| im         |                                              | Isle of Man                                                 | gPp     |
| in         |                                              | India                                                       | gPp     |
| in-an      | India                                        | Andaman and Nicobar Islands                                 | gPp     |
| in-ap      | India                                        | Andhra Pradesh                                              | gPp     |
| in-ar      | India                                        | Arunachal Pradesh                                           | gPp     |
| in-as      | India                                        | Assam                                                       | gPp     |
| in-br      | India                                        | Bihar                                                       | gPp     |
| in-cg      | India                                        | Chhattisgarh                                                | gPp     |
| in-ch      | India                                        | Chandigarh                                                  | gPp     |
| in-dh      | India                                        | Dadra and Nagar Haveli and Daman and Diu                    | gPp     |
| in-dl      | India                                        | Delhi                                                       | gPp     |
| in-ga      | India                                        | Goa                                                         | gPp     |
| in-gj      | India                                        | Gujarat                                                     | gPp     |
| in-hp      | India                                        | Himachal Pradesh                                            | gPp     |
| in-hr      | India                                        | Haryana                                                     | gPp     |
| in-jh      | India                                        | Jharkhand                                                   | gPp     |
| in-jk      | India                                        | Jammu and Kashmir                                           | gPp     |
| in-ka      | India                                        | Karnataka                                                   | gPp     |
| in-kl      | India                                        | Kerala                                                      | gPp     |
| in-la      | India                                        | Ladakh                                                      | gPp     |
| in-ld      | India                                        | Lakshadweep                                                 | gPp     |
| in-mh      | India                                        | Maharashtra                                                 | gPp     |
| in-ml      | India                                        | Meghalaya                                                   | gPp     |
| in-mn      | India                                        | Manipur                                                     | gPp     |
| in-mp      | India                                        | Madhya Pradesh                                              | gPp     |
| in-mz      | India                                        | Mizoram                                                     | gPp     |
| in-nl      | India                                        | Nagaland                                                    | gPp     |
| in-od      | India                                        | Odisha                                                      | gPp     |
| in-pb      | India                                        | Punjab                                                      | gPp     |
| in-py      | India                                        | Puducherry                                                  | gPp     |
| in-rj      | India                                        | Rajasthan                                                   | gPp     |
| in-sk      | India                                        | Sikkim                                                      | gPp     |
| in-tn      | India                                        | Tamil Nadu                                                  | gPp     |
| in-tr      | India                                        | Tripura                                                     | gPp     |
| in-ts      | India                                        | Telangana                                                   | gPp     |
| in-uk      | India                                        | Uttarakhand                                                 | gPp     |
| in-up      | India                                        | Uttar Pradesh                                               | gPp     |
| in-wb      | India                                        | West Bengal                                                 | gPp     |
| io         |                                              | British Indian Ocean Territory                              | gPp     |
| iq         |                                              | Iraq                                                        | gPp     |
| iq-an      | Iraq                                         | Al-Anbar Governorate                                        | gPp     |
| iq-ba      | Iraq                                         | Basra Governorate                                           | gPp     |
| iq-bb      | Iraq                                         | Babil Governorate                                           | gPp     |
| iq-bg      | Iraq                                         | Baghdad Governorate                                         | gPp     |
| iq-di      | Iraq                                         | Diyala Governorate                                          | gPp     |
| iq-dq      | Iraq                                         | Dhi Qar Governorate                                         | gPp     |
| iq-ka      | Iraq                                         | Karbala Governorate                                         | gPp     |
| iq-ki      | Iraq                                         | Kirkuk Governorate                                          | gPp     |
| iq-kr      | Iraq                                         | Iraqi Kurdistan Region                                      | gPp     |
| iq-ma      | Iraq                                         | Maysan Governorate                                          | gPp     |
| iq-mu      | Iraq                                         | Muthanna Governorate                                        | gPp     |
| iq-na      | Iraq                                         | Al-Najaf Governorate                                        | gPp     |
| iq-ni      | Iraq                                         | Nineveh Governorate                                         | gPp     |
| iq-qa      | Iraq                                         | Al-Qadisiyah Governorate                                    | gPp     |
| iq-sd      | Iraq                                         | Saladin Governorate                                         | gPp     |
| iq-wa      | Iraq                                         | Wasit Governorate                                           | gPp     |
| ir         |                                              | Iran                                                        | gPp     |
| ir-00      | Iran                                         | Markazi Province                                            | gPp     |
| ir-01      | Iran                                         | Gilan Province                                              | gPp     |
| ir-02      | Iran                                         | Mazandaran Province                                         | gPp     |
| ir-03      | Iran                                         | East Azerbaijan Province                                    | gPp     |
| ir-04      | Iran                                         | West Azerbaijan Province                                    | gPp     |
| ir-05      | Iran                                         | Kermanshah Province                                         | gPp     |
| ir-06      | Iran                                         | Khuzestan Province                                          | gPp     |
| ir-07      | Iran                                         | Fars Province                                               | gPp     |
| ir-08      | Iran                                         | Kerman Province                                             | gPp     |
| ir-09      | Iran                                         | Razavi Khorasan                                             | gPp     |
| ir-10      | Iran                                         | Isfahan Province                                            | gPp     |
| ir-11      | Iran                                         | Sistan and Baluchestan Province                             | gPp     |
| ir-12      | Iran                                         | Kurdistan Province                                          | gPp     |
| ir-13      | Iran                                         | Hamadan Province                                            | gPp     |
| ir-14      | Iran                                         | Chaharmahal and Bakhtiyari Province                         | gPp     |
| ir-15      | Iran                                         | Lorestan Province                                           | gPp     |
| ir-16      | Iran                                         | Ilam Province                                               | gPp     |
| ir-17      | Iran                                         | Kohgiluye and Buyer Ahmad Province                          | gPp     |
| ir-18      | Iran                                         | Bushehr Province                                            | gPp     |
| ir-19      | Iran                                         | Zanjan Province                                             | gPp     |
| ir-20      | Iran                                         | Semnan Province                                             | gPp     |
| ir-21      | Iran                                         | Yazd Province                                               | gPp     |
| ir-22      | Iran                                         | Hormozgan Province                                          | gPp     |
| ir-23      | Iran                                         | Tehran Province                                             | gPp     |
| ir-24      | Iran                                         | Ardabil Province                                            | gPp     |
| ir-25      | Iran                                         | Qom Province                                                | gPp     |
| ir-26      | Iran                                         | Qazvin Province                                             | gPp     |
| ir-27      | Iran                                         | Golestan Province                                           | gPp     |
| ir-28      | Iran                                         | North Khorasan Province                                     | gPp     |
| ir-29      | Iran                                         | South Khorasan Province                                     | gPp     |
| ir-30      | Iran                                         | Alborz Province                                             | gPp     |
| is         |                                              | Iceland                                                     | gPp     |
| is-1       | Iceland                                      | Capital Region                                              | gPp     |
| is-2       | Iceland                                      | Southern Peninsula                                          | gPp     |
| is-3       | Iceland                                      | Western Region                                              | gPp     |
| is-4       | Iceland                                      | Westfjords Region                                           | gPp     |
| is-5       | Iceland                                      | Northwestern Region                                         | gPp     |
| is-6       | Iceland                                      | Northeastern Region                                         | gPp     |
| is-7       | Iceland                                      | Eastern Region                                              | gPp     |
| is-8       | Iceland                                      | Southern Region                                             | gPp     |
| it         |                                              | Italy                                                       | gPp     |
| it-21      | Italy                                        | Piedmont                                                    | gPp     |
| it-23      | Italy                                        | Aosta Valley                                                | gPp     |
| it-25      | Italy                                        | Lombardy                                                    | gPp     |
| it-32      | Italy                                        | Trentino – Alto Adige/Südtirol                              | gPp     |
| it-34      | Italy                                        | Veneto                                                      | gPp     |
| it-36      | Italy                                        | Friuli – Venezia Giulia                                     | gPp     |
| it-42      | Italy                                        | Liguria                                                     | gPp     |
| it-45      | Italy                                        | Emilia-Romagna                                              | gPp     |
| it-52      | Italy                                        | Tuscany                                                     | gPp     |
| it-55      | Italy                                        | Umbria                                                      | gPp     |
| it-57      | Italy                                        | Marche                                                      | gPp     |
| it-62      | Italy                                        | Lazio                                                       | gPp     |
| it-65      | Italy                                        | Abruzzo                                                     | gPp     |
| it-67      | Italy                                        | Molise                                                      | gPp     |
| it-72      | Italy                                        | Campania                                                    | gPp     |
| it-75      | Italy                                        | Apulia                                                      | gPp     |
| it-77      | Italy                                        | Basilicata                                                  | gPp     |
| it-78      | Italy                                        | Calabria                                                    | gPp     |
| it-82      | Italy                                        | Sicily                                                      | gPp     |
| it-88      | Italy                                        | Sardinia                                                    | gPp     |
| je         |                                              | Jersey                                                      | gPp     |
| jm         |                                              | Jamaica                                                     | gPp     |
| jm-01      | Jamaica                                      | Kingston                                                    | gPp     |
| jm-02      | Jamaica                                      | Saint Andrew                                                | gPp     |
| jm-03      | Jamaica                                      | Saint Thomas                                                | gPp     |
| jm-04      | Jamaica                                      | Portland                                                    | gPp     |
| jm-05      | Jamaica                                      | Saint Mary                                                  | gPp     |
| jm-06      | Jamaica                                      | Saint Ann                                                   | gPp     |
| jm-07      | Jamaica                                      | Trelawny                                                    | gPp     |
| jm-08      | Jamaica                                      | Saint James                                                 | gPp     |
| jm-09      | Jamaica                                      | Hanover                                                     | gPp     |
| jm-10      | Jamaica                                      | Westmoreland                                                | gPp     |
| jm-11      | Jamaica                                      | Saint Elizabeth                                             | gPp     |
| jm-12      | Jamaica                                      | Manchester                                                  | gPp     |
| jm-13      | Jamaica                                      | Clarendon                                                   | gPp     |
| jm-14      | Jamaica                                      | Saint Catherine                                             | gPp     |
| jo         |                                              | Jordan                                                      | gPp     |
| jo-aj      | Jordan                                       | Ajlun                                                       | gPp     |
| jo-am      | Jordan                                       | Amman                                                       | gPp     |
| jo-aq      | Jordan                                       | Aqaba                                                       | gPp     |
| jo-at      | Jordan                                       | Tafilah                                                     | gPp     |
| jo-az      | Jordan                                       | Zarqa                                                       | gPp     |
| jo-ba      | Jordan                                       | Balqa                                                       | gPp     |
| jo-ir      | Jordan                                       | Irbid                                                       | gPp     |
| jo-ja      | Jordan                                       | Jarash                                                      | gPp     |
| jo-ka      | Jordan                                       | Karak                                                       | gPp     |
| jo-ma      | Jordan                                       | Mafraq                                                      | gPp     |
| jo-md      | Jordan                                       | Madaba                                                      | gPp     |
| jo-mn      | Jordan                                       | Maan                                                        | gPp     |
| jp         |                                              | Japan                                                       | gPp     |
| jp-01      | Japan                                        | Hokkaido Prefecture                                         | gPp     |
| jp-02      | Japan                                        | Aomori Prefecture                                           | gPp     |
| jp-03      | Japan                                        | Iwate Prefecture                                            | gPp     |
| jp-04      | Japan                                        | Miyagi Prefecture                                           | gPp     |
| jp-05      | Japan                                        | Akita Prefecture                                            | gPp     |
| jp-06      | Japan                                        | Yamagata Prefecture                                         | gPp     |
| jp-07      | Japan                                        | Fukushima Prefecture                                        | gPp     |
| jp-08      | Japan                                        | Ibaraki Prefecture                                          | gPp     |
| jp-09      | Japan                                        | Tochigi Prefecture                                          | gPp     |
| jp-10      | Japan                                        | Gunma Prefecture                                            | gPp     |
| jp-11      | Japan                                        | Saitama Prefecture                                          | gPp     |
| jp-12      | Japan                                        | Chiba Prefecture                                            | gPp     |
| jp-13      | Japan                                        | Tokyo                                                       | gPp     |
| jp-14      | Japan                                        | Kanagawa Prefecture                                         | gPp     |
| jp-15      | Japan                                        | Niigata Prefecture                                          | gPp     |
| jp-16      | Japan                                        | Toyama Prefecture                                           | gPp     |
| jp-17      | Japan                                        | Ishikawa Prefecture                                         | gPp     |
| jp-18      | Japan                                        | Fukui Prefecture                                            | gPp     |
| jp-19      | Japan                                        | Yamanashi Prefecture                                        | gPp     |
| jp-20      | Japan                                        | Nagano Prefecture                                           | gPp     |
| jp-21      | Japan                                        | Gifu Prefecture                                             | gPp     |
| jp-22      | Japan                                        | Shizuoka Prefecture                                         | gPp     |
| jp-23      | Japan                                        | Aichi Prefecture                                            | gPp     |
| jp-24      | Japan                                        | Mie Prefecture                                              | gPp     |
| jp-25      | Japan                                        | Shiga Prefecture                                            | gPp     |
| jp-26      | Japan                                        | Kyoto Prefecture                                            | gPp     |
| jp-27      | Japan                                        | Osaka Prefecture                                            | gPp     |
| jp-28      | Japan                                        | Hyogo Prefecture                                            | gPp     |
| jp-29      | Japan                                        | Nara Prefecture                                             | gPp     |
| jp-30      | Japan                                        | Wakayama Prefecture                                         | gPp     |
| jp-31      | Japan                                        | Tottori Prefecture                                          | gPp     |
| jp-32      | Japan                                        | Shimane Prefecture                                          | gPp     |
| jp-33      | Japan                                        | Okayama Prefecture                                          | gPp     |
| jp-34      | Japan                                        | Hiroshima Prefecture                                        | gPp     |
| jp-35      | Japan                                        | Yamaguchi Prefecture                                        | gPp     |
| jp-36      | Japan                                        | Tokushima Prefecture                                        | gPp     |
| jp-37      | Japan                                        | Kagawa Prefecture                                           | gPp     |
| jp-38      | Japan                                        | Ehime Prefecture                                            | gPp     |
| jp-39      | Japan                                        | Kochi Prefecture                                            | gPp     |
| jp-40      | Japan                                        | Fukuoka Prefecture                                          | gPp     |
| jp-41      | Japan                                        | Saga Prefecture                                             | gPp     |
| jp-42      | Japan                                        | Nagasaki Prefecture                                         | gPp     |
| jp-43      | Japan                                        | Kumamoto Prefecture                                         | gPp     |
| jp-44      | Japan                                        | Oita Prefecture                                             | gPp     |
| jp-45      | Japan                                        | Miyazaki Prefecture                                         | gPp     |
| jp-46      | Japan                                        | Kagoshima Prefecture                                        | gPp     |
| jp-47      | Japan                                        | Okinawa Prefecture                                          | gPp     |
| js         |                                              | Judea and Samaria                                           | gPp     |
| ke         |                                              | Kenya                                                       | gPp     |
| ke-01      | Kenya                                        | Baringo                                                     | gPp     |
| ke-02      | Kenya                                        | Bomet                                                       | gPp     |
| ke-03      | Kenya                                        | Bungoma County                                              | gPp     |
| ke-04      | Kenya                                        | Busia County                                                | gPp     |
| ke-05      | Kenya                                        | Elgeyo-Marakwet County                                      | gPp     |
| ke-06      | Kenya                                        | Embu                                                        | gPp     |
| ke-07      | Kenya                                        | Garissa                                                     | gPp     |
| ke-08      | Kenya                                        | Homa Bay County                                             | gPp     |
| ke-09      | Kenya                                        | Isiolo                                                      | gPp     |
| ke-10      | Kenya                                        | Kajiado County                                              | gPp     |
| ke-11      | Kenya                                        | Kakamega County                                             | gPp     |
| ke-12      | Kenya                                        | Kericho County                                              | gPp     |
| ke-13      | Kenya                                        | Kiambu                                                      | gPp     |
| ke-14      | Kenya                                        | Kilifi County                                               | gPp     |
| ke-15      | Kenya                                        | Kirinyaga County                                            | gPp     |
| ke-16      | Kenya                                        | Kisii County                                                | gPp     |
| ke-17      | Kenya                                        | Kisumu County                                               | gPp     |
| ke-18      | Kenya                                        | Kitui County                                                | gPp     |
| ke-19      | Kenya                                        | Kwale                                                       | gPp     |
| ke-20      | Kenya                                        | Laikipia County                                             | gPp     |
| ke-21      | Kenya                                        | Lamu                                                        | gPp     |
| ke-22      | Kenya                                        | Machakos County                                             | gPp     |
| ke-23      | Kenya                                        | Makueni                                                     | gPp     |
| ke-24      | Kenya                                        | Mandera County                                              | gPp     |
| ke-25      | Kenya                                        | Marsabit County                                             | gPp     |
| ke-26      | Kenya                                        | Meru County                                                 | gPp     |
| ke-27      | Kenya                                        | Migori County                                               | gPp     |
| ke-28      | Kenya                                        | Mombasa County                                              | gPp     |
| ke-29      | Kenya                                        | Murang'a County                                             | gPp     |
| ke-30      | Kenya                                        | Nairobi County                                              | gPp     |
| ke-31      | Kenya                                        | Nakuru                                                      | gPp     |
| ke-32      | Kenya                                        | Nandi County                                                | gPp     |
| ke-33      | Kenya                                        | Narok                                                       | gPp     |
| ke-34      | Kenya                                        | Nyamira County                                              | gPp     |
| ke-35      | Kenya                                        | Nyandarua                                                   | gPp     |
| ke-36      | Kenya                                        | Nyeri                                                       | gPp     |
| ke-37      | Kenya                                        | Samburu                                                     | gPp     |
| ke-38      | Kenya                                        | Siaya County                                                | gPp     |
| ke-39      | Kenya                                        | Taita–Taveta                                                | gPp     |
| ke-40      | Kenya                                        | Tana River County                                           | gPp     |
| ke-41      | Kenya                                        | Tharaka-Nithi                                               | gPp     |
| ke-42      | Kenya                                        | Trans-Nzoia County                                          | gPp     |
| ke-43      | Kenya                                        | Turkana County                                              | gPp     |
| ke-44      | Kenya                                        | Uasin Gishu County                                          | gPp     |
| ke-45      | Kenya                                        | Vihiga County                                               | gPp     |
| ke-46      | Kenya                                        | Wajir County                                                | gPp     |
| ke-47      | Kenya                                        | West Pokot                                                  | gPp     |
| kg         |                                              | Kyrgyzstan                                                  | gPp     |
| kg-b       | Kyrgyzstan                                   | Batken Region                                               | gPp     |
| kg-c       | Kyrgyzstan                                   | Chuy Region                                                 | gPp     |
| kg-gb      | Kyrgyzstan                                   | Bishkek City                                                | gPp     |
| kg-go      | Kyrgyzstan                                   | Osh City                                                    | gPp     |
| kg-j       | Kyrgyzstan                                   | Jalal-Abad Region                                           | gPp     |
| kg-n       | Kyrgyzstan                                   | Naryn Region                                                | gPp     |
| kg-o       | Kyrgyzstan                                   | Osh Region                                                  | gPp     |
| kg-t       | Kyrgyzstan                                   | Talas Region                                                | gPp     |
| kg-y       | Kyrgyzstan                                   | Issyk-Kul Region                                            | gPp     |
| kh         |                                              | Cambodia                                                    | gPp     |
| kh-1       | Cambodia                                     | Bantey Meanchey                                             | gPp     |
| kh-10      | Cambodia                                     | Kratie                                                      | gPp     |
| kh-11      | Cambodia                                     | Mondulkiri                                                  | gPp     |
| kh-12      | Cambodia                                     | Phnom Penh                                                  | gPp     |
| kh-13      | Cambodia                                     | Preah Vihear                                                | gPp     |
| kh-14      | Cambodia                                     | Prey Veng                                                   | gPp     |
| kh-15      | Cambodia                                     | Pursat                                                      | gPp     |
| kh-16      | Cambodia                                     | Ratanakiri                                                  | gPp     |
| kh-17      | Cambodia                                     | Siem Reap                                                   | gPp     |
| kh-18      | Cambodia                                     | Khaet Preah Sihanouk                                        | gPp     |
| kh-19      | Cambodia                                     | Stung Treng                                                 | gPp     |
| kh-2       | Cambodia                                     | Battambang                                                  | gPp     |
| kh-20      | Cambodia                                     | Svay Rieng                                                  | gPp     |
| kh-21      | Cambodia                                     | Takeo                                                       | gPp     |
| kh-22      | Cambodia                                     | Oddar Meanchey                                              | gPp     |
| kh-23      | Cambodia                                     | Kep                                                         | gPp     |
| kh-24      | Cambodia                                     | Pailin                                                      | gPp     |
| kh-25      | Cambodia                                     | Tbong Khmum                                                 | gPp     |
| kh-3       | Cambodia                                     | Kampong Cham                                                | gPp     |
| kh-4       | Cambodia                                     | Kampong Chhnang                                             | gPp     |
| kh-5       | Cambodia                                     | Kampong Speu                                                | gPp     |
| kh-6       | Cambodia                                     | Kampong Thom                                                | gPp     |
| kh-7       | Cambodia                                     | Kampot                                                      | gPp     |
| kh-8       | Cambodia                                     | Kandal                                                      | gPp     |
| kh-9       | Cambodia                                     | Koh Kong                                                    | gPp     |
| ki         |                                              | Kiribati                                                    | gPp     |
| ki-g       | Kiribati                                     | Gilbert Islands                                             | gPp     |
| ki-l       | Kiribati                                     | Line Islands                                                | gPp     |
| ki-p       | Kiribati                                     | Phoenix Islands                                             | gPp     |
| km         |                                              | Comoros                                                     | gPp     |
| km-a       | Comoros                                      | Anjouan                                                     | gPp     |
| km-g       | Comoros                                      | Grande Comore                                               | gPp     |
| km-m       | Comoros                                      | Moheli                                                      | gPp     |
| kn         |                                              | Saint Kitts and Nevis                                       | gPp     |
| kn-k       | Saint Kitts and Nevis                        | Saint Kitts                                                 | gPp     |
| kn-n       | Saint Kitts and Nevis                        | Nevis                                                       | gPp     |
| kp         |                                              | North Korea                                                 | gPp     |
| kp-01      | North Korea                                  | P'yŏngyang                                                  | gPp     |
| kp-02      | North Korea                                  | South Pyongan                                               | gPp     |
| kp-03      | North Korea                                  | North Pyongan                                               | gPp     |
| kp-04      | North Korea                                  | Chagang                                                     | gPp     |
| kp-05      | North Korea                                  | South Hwanghae                                              | gPp     |
| kp-06      | North Korea                                  | North Hwanghae                                              | gPp     |
| kp-07      | North Korea                                  | Kangwon                                                     | gPp     |
| kp-08      | North Korea                                  | South Hamgyong                                              | gPp     |
| kp-09      | North Korea                                  | North Hamgyong                                              | gPp     |
| kp-10      | North Korea                                  | Ryanggang                                                   | gPp     |
| kp-13      | North Korea                                  | Rason                                                       | gPp     |
| kp-14      | North Korea                                  | Nampo                                                       | gPp     |
| kp-15      | North Korea                                  | Kaesong                                                     | gPp     |
| kr         |                                              | South Korea                                                 | gPp     |
| kr-11      | South Korea                                  | Seoul                                                       | gPp     |
| kr-26      | South Korea                                  | Busan                                                       | gPp     |
| kr-27      | South Korea                                  | Daegu                                                       | gPp     |
| kr-28      | South Korea                                  | Incheon                                                     | gPp     |
| kr-30      | South Korea                                  | Daejeon                                                     | gPp     |
| kr-31      | South Korea                                  | Ulsan                                                       | gPp     |
| kr-41      | South Korea                                  | Gyeonggi                                                    | gPp     |
| kr-42      | South Korea                                  | Gangwon State                                               | gPp     |
| kr-43      | South Korea                                  | North Chungcheong                                           | gPp     |
| kr-44      | South Korea                                  | South Chungcheong                                           | gPp     |
| kr-45      | South Korea                                  | Jeonbuk-do                                                  | gPp     |
| kr-47      | South Korea                                  | North Gyeongsang                                            | gPp     |
| kr-48      | South Korea                                  | South Gyeongsang                                            | gPp     |
| kr-49      | South Korea                                  | Jeju                                                        | gPp     |
| kr-50      | South Korea                                  | Sejong                                                      | gPp     |
| kr-xx      | South Korea                                  | Jeonnam-Gwangju Special Metropolitan City                   | gPp     |
| kw         |                                              | Kuwait                                                      | gPp     |
| kw-ah      | Kuwait                                       | Ahmadi Governorate                                          | gPp     |
| kw-fa      | Kuwait                                       | Farwaniya Governorate                                       | gPp     |
| kw-ha      | Kuwait                                       | Hawalli Governorate                                         | gPp     |
| kw-ja      | Kuwait                                       | Jahra Governorate                                           | gPp     |
| kw-ku      | Kuwait                                       | Capital Governorate                                         | gPp     |
| kw-mu      | Kuwait                                       | Mubarak al-Kabir Governorate                                | gPp     |
| ky         |                                              | Cayman Islands                                              | gPp     |
| kz         |                                              | Kazakhstan                                                  | gPp     |
| kz-10      | Kazakhstan                                   | Abay Region                                                 | gPp     |
| kz-11      | Kazakhstan                                   | Akmola Region                                               | gPp     |
| kz-15      | Kazakhstan                                   | Aqtöbe Region                                               | gPp     |
| kz-19      | Kazakhstan                                   | Almaty Region                                               | gPp     |
| kz-23      | Kazakhstan                                   | Atyrau Region                                               | gPp     |
| kz-27      | Kazakhstan                                   | West Kazakhstan Region                                      | gPp     |
| kz-31      | Kazakhstan                                   | Jambyl Region                                               | gPp     |
| kz-33      | Kazakhstan                                   | Jetisu Region                                               | gPp     |
| kz-35      | Kazakhstan                                   | Karaganda Region                                            | gPp     |
| kz-39      | Kazakhstan                                   | Kostanay Region                                             | gPp     |
| kz-43      | Kazakhstan                                   | Kyzylorda Region                                            | gPp     |
| kz-47      | Kazakhstan                                   | Mangystau Region                                            | gPp     |
| kz-55      | Kazakhstan                                   | Pavlodar Region                                             | gPp     |
| kz-59      | Kazakhstan                                   | North Kazakhstan Region                                     | gPp     |
| kz-61      | Kazakhstan                                   | Turkistan Region                                            | gPp     |
| kz-62      | Kazakhstan                                   | Ulytau Region                                               | gPp     |
| kz-63      | Kazakhstan                                   | East Kazakhstan Region                                      | gPp     |
| kz-71      | Kazakhstan                                   | Astana                                                      | gPp     |
| kz-75      | Kazakhstan                                   | Almaty                                                      | gPp     |
| kz-79      | Kazakhstan                                   | Shymkent                                                    | gPp     |
| la         |                                              | Laos                                                        | gPp     |
| la-at      | Laos                                         | Attapeu                                                     | gPp     |
| la-bk      | Laos                                         | Bokeo Province                                              | gPp     |
| la-bl      | Laos                                         | Bolikhamsai                                                 | gPp     |
| la-ch      | Laos                                         | Champasak Province                                          | gPp     |
| la-ho      | Laos                                         | Houaphanh                                                   | gPp     |
| la-kh      | Laos                                         | Khammouane                                                  | gPp     |
| la-lm      | Laos                                         | Luang Namtha                                                | gPp     |
| la-lp      | Laos                                         | Luang Prabang                                               | gPp     |
| la-ou      | Laos                                         | Oudomxay                                                    | gPp     |
| la-ph      | Laos                                         | Phongsaly                                                   | gPp     |
| la-sl      | Laos                                         | Salavan Province                                            | gPp     |
| la-sv      | Laos                                         | Savannakhet Province                                        | gPp     |
| la-vi      | Laos                                         | Vientiane Province                                          | gPp     |
| la-vt      | Laos                                         | Vientiane Prefecture                                        | gPp     |
| la-xa      | Laos                                         | Sainyabuli Province                                         | gPp     |
| la-xe      | Laos                                         | Sekong Province                                             | gPp     |
| la-xi      | Laos                                         | Xiangkhouang Province                                       | gPp     |
| la-xs      | Laos                                         | Xaisomboun Province                                         | gPp     |
| lb         |                                              | Lebanon                                                     | gPp     |
| lb-ak      | Lebanon                                      | Akkar Governorate                                           | gPp     |
| lb-as      | Lebanon                                      | North Governorate                                           | gPp     |
| lb-ba      | Lebanon                                      | Beirut Governorate                                          | gPp     |
| lb-bh      | Lebanon                                      | Baalbek-Hermel Governorate                                  | gPp     |
| lb-bi      | Lebanon                                      | Beqaa Governorate                                           | gPp     |
| lb-ja      | Lebanon                                      | South Governorate                                           | gPp     |
| lb-jl      | Lebanon                                      | Mount Lebanon Governorate                                   | gPp     |
| lb-na      | Lebanon                                      | Nabatieh Governorate                                        | gPp     |
| lb-xx      | Lebanon                                      | Keserwan-Jbeil Governorate                                  | gPp     |
| lc         |                                              | Saint Lucia                                                 | gPp     |
| lc-01      | Saint Lucia                                  | Anse La Raye                                                | gPp     |
| lc-02      | Saint Lucia                                  | Castries                                                    | gPp     |
| lc-03      | Saint Lucia                                  | Choiseul                                                    | gPp     |
| lc-05      | Saint Lucia                                  | Dennery                                                     | gPp     |
| lc-06      | Saint Lucia                                  | Gros Islet                                                  | gPp     |
| lc-07      | Saint Lucia                                  | Laborie                                                     | gPp     |
| lc-08      | Saint Lucia                                  | Micoud                                                      | gPp     |
| lc-10      | Saint Lucia                                  | Soufrière                                                   | gPp     |
| lc-11      | Saint Lucia                                  | Vieux Fort                                                  | gPp     |
| lc-12      | Saint Lucia                                  | Canaries                                                    | gPp     |
| li         |                                              | Liechtenstein                                               | gPp     |
| lk         |                                              | Sri Lanka                                                   | gPp     |
| lk-1       | Sri Lanka                                    | Western Province                                            | gPp     |
| lk-2       | Sri Lanka                                    | Central Province                                            | gPp     |
| lk-3       | Sri Lanka                                    | Southern Province                                           | gPp     |
| lk-4       | Sri Lanka                                    | Northern Province                                           | gPp     |
| lk-5       | Sri Lanka                                    | Eastern Province                                            | gPp     |
| lk-6       | Sri Lanka                                    | North Western Province                                      | gPp     |
| lk-7       | Sri Lanka                                    | North Central Province                                      | gPp     |
| lk-8       | Sri Lanka                                    | Uva Province                                                | gPp     |
| lk-9       | Sri Lanka                                    | Sabaragamuwa Province                                       | gPp     |
| lr         |                                              | Liberia                                                     | gPp     |
| lr-bg      | Liberia                                      | Bong County                                                 | gPp     |
| lr-bm      | Liberia                                      | Bomi County                                                 | gPp     |
| lr-cm      | Liberia                                      | Grand Cape Mount County                                     | gPp     |
| lr-gb      | Liberia                                      | Grand Bassa County                                          | gPp     |
| lr-gg      | Liberia                                      | Grand Gedeh County                                          | gPp     |
| lr-gk      | Liberia                                      | Grand Kru County                                            | gPp     |
| lr-gp      | Liberia                                      | Gbarpolu County                                             | gPp     |
| lr-lo      | Liberia                                      | Lofa County                                                 | gPp     |
| lr-mg      | Liberia                                      | Margibi County                                              | gPp     |
| lr-mo      | Liberia                                      | Montserrado County                                          | gPp     |
| lr-my      | Liberia                                      | Maryland County                                             | gPp     |
| lr-ni      | Liberia                                      | Nimba County                                                | gPp     |
| lr-rg      | Liberia                                      | River Gee County                                            | gPp     |
| lr-ri      | Liberia                                      | Rivercess County                                            | gPp     |
| lr-si      | Liberia                                      | Sinoe County                                                | gPp     |
| ls         |                                              | Lesotho                                                     | gPp     |
| ls-a       | Lesotho                                      | Maseru District                                             | gPp     |
| ls-b       | Lesotho                                      | Butha-Buthe District                                        | gPp     |
| ls-c       | Lesotho                                      | Leribe District                                             | gPp     |
| ls-d       | Lesotho                                      | Berea District                                              | gPp     |
| ls-e       | Lesotho                                      | Mafeteng District                                           | gPp     |
| ls-f       | Lesotho                                      | Mohale's Hoek District                                      | gPp     |
| ls-g       | Lesotho                                      | Quthing District                                            | gPp     |
| ls-h       | Lesotho                                      | Qacha's Nek District                                        | gPp     |
| ls-j       | Lesotho                                      | Mokhotlong District                                         | gPp     |
| ls-k       | Lesotho                                      | Thaba-Tseka District                                        | gPp     |
| lt         |                                              | Lithuania                                                   | gPp     |
| lt-al      | Lithuania                                    | Alytus County                                               | gPp     |
| lt-kl      | Lithuania                                    | Klaipėda County                                             | gPp     |
| lt-ku      | Lithuania                                    | Kaunas County                                               | gPp     |
| lt-mr      | Lithuania                                    | Marijampolė County                                          | gPp     |
| lt-pn      | Lithuania                                    | Panevėžys County                                            | gPp     |
| lt-sa      | Lithuania                                    | Šiauliai County                                             | gPp     |
| lt-ta      | Lithuania                                    | Tauragė County                                              | gPp     |
| lt-te      | Lithuania                                    | Telšiai County                                              | gPp     |
| lt-ut      | Lithuania                                    | Utena County                                                | gPp     |
| lt-vl      | Lithuania                                    | Vilnius County                                              | gPp     |
| lu         |                                              | Luxembourg                                                  | gPp     |
| lu-ca      | Luxembourg                                   | Canton Capellen                                             | gPp     |
| lu-cl      | Luxembourg                                   | Canton Clervaux                                             | gPp     |
| lu-di      | Luxembourg                                   | Canton Diekirch                                             | gPp     |
| lu-ec      | Luxembourg                                   | Canton Echternach                                           | gPp     |
| lu-es      | Luxembourg                                   | Canton Esch-sur-Alzette                                     | gPp     |
| lu-gr      | Luxembourg                                   | Canton Grevenmacher                                         | gPp     |
| lu-lu      | Luxembourg                                   | Canton Luxembourg                                           | gPp     |
| lu-me      | Luxembourg                                   | Canton Mersch                                               | gPp     |
| lu-rd      | Luxembourg                                   | Canton Redange                                              | gPp     |
| lu-rm      | Luxembourg                                   | Canton Remich                                               | gPp     |
| lu-vd      | Luxembourg                                   | Canton Vianden                                              | gPp     |
| lu-wi      | Luxembourg                                   | Canton Wiltz                                                | gPp     |
| lv         |                                              | Latvia                                                      | gPp     |
| lv-002     | Latvia                                       | Aizkraukle Municipality                                     | gPp     |
| lv-007     | Latvia                                       | Alūksnes novads                                             | gPp     |
| lv-011     | Latvia                                       | Ādažu novads                                                | gPp     |
| lv-015     | Latvia                                       | Balvu novads                                                | gPp     |
| lv-016     | Latvia                                       | Bauskas novads                                              | gPp     |
| lv-022     | Latvia                                       | Cēsu novads                                                 | gPp     |
| lv-026     | Latvia                                       | Dobeles novads                                              | gPp     |
| lv-033     | Latvia                                       | Gulbenes novads                                             | gPp     |
| lv-041     | Latvia                                       | Jelgavas novads                                             | gPp     |
| lv-042     | Latvia                                       | Jēkabpils novads                                            | gPp     |
| lv-047     | Latvia                                       | Krāslavas novads                                            | gPp     |
| lv-050     | Latvia                                       | Kuldīgas novads                                             | gPp     |
| lv-052     | Latvia                                       | Ķekavas novads                                              | gPp     |
| lv-054     | Latvia                                       | Limbažu novads                                              | gPp     |
| lv-056     | Latvia                                       | Līvānu novads                                               | gPp     |
| lv-058     | Latvia                                       | Ludzas novads                                               | gPp     |
| lv-059     | Latvia                                       | Madona Municipality                                         | gPp     |
| lv-062     | Latvia                                       | Mārupes novads                                              | gPp     |
| lv-067     | Latvia                                       | Ogres novads                                                | gPp     |
| lv-068     | Latvia                                       | Olaines novads                                              | gPp     |
| lv-073     | Latvia                                       | Preiļu novads                                               | gPp     |
| lv-077     | Latvia                                       | Rēzeknes novads                                             | gPp     |
| lv-080     | Latvia                                       | Ropažu novads                                               | gPp     |
| lv-087     | Latvia                                       | Salaspils novads                                            | gPp     |
| lv-088     | Latvia                                       | Saldus novads                                               | gPp     |
| lv-089     | Latvia                                       | Saulkrastu novads                                           | gPp     |
| lv-091     | Latvia                                       | Siguldas novads                                             | gPp     |
| lv-094     | Latvia                                       | Smiltenes novads                                            | gPp     |
| lv-097     | Latvia                                       | Talsi Municipality                                          | gPp     |
| lv-099     | Latvia                                       | Tukuma novads                                               | gPp     |
| lv-101     | Latvia                                       | Valkas novads                                               | gPp     |
| lv-106     | Latvia                                       | Ventspils novads                                            | gPp     |
| lv-111     | Latvia                                       | Augšdaugavas novads                                         | gPp     |
| lv-112     | Latvia                                       | South Kurzeme Municipality                                  | gPp     |
| lv-113     | Latvia                                       | Valmieras novads                                            | gPp     |
| lv-dgv     | Latvia                                       | Daugavpils                                                  | gPp     |
| lv-jel     | Latvia                                       | Jelgava                                                     | gPp     |
| lv-jur     | Latvia                                       | Jūrmala                                                     | gPp     |
| lv-lpx     | Latvia                                       | Liepāja                                                     | gPp     |
| lv-rez     | Latvia                                       | Rēzekne                                                     | gPp     |
| lv-rix     | Latvia                                       | Riga                                                        | gPp     |
| lv-ven     | Latvia                                       | Ventspils                                                   | gPp     |
| ly         |                                              | Libya                                                       | gPp     |
| ly-ba      | Libya                                        | Benghazi                                                    | gPp     |
| ly-bu      | Libya                                        | Butnan                                                      | gPp     |
| ly-dr      | Libya                                        | Derna                                                       | gPp     |
| ly-gt      | Libya                                        | Ghat                                                        | gPp     |
| ly-ja      | Libya                                        | Jabal al Akhdar                                             | gPp     |
| ly-jg      | Libya                                        | Nafusa Mountains                                            | gPp     |
| ly-ji      | Libya                                        | Jafara                                                      | gPp     |
| ly-ju      | Libya                                        | District of Al-Jufra                                        | gPp     |
| ly-kf      | Libya                                        | Kufra                                                       | gPp     |
| ly-mb      | Libya                                        | Murqub                                                      | gPp     |
| ly-mi      | Libya                                        | Bani Walid                                                  | gPp     |
| ly-mj      | Libya                                        | Marj                                                        | gPp     |
| ly-mq      | Libya                                        | Murzuq                                                      | gPp     |
| ly-nl      | Libya                                        | Nalut                                                       | gPp     |
| ly-nq      | Libya                                        | Nuqat al Khams                                              | gPp     |
| ly-sb      | Libya                                        | Sabha                                                       | gPp     |
| ly-sr      | Libya                                        | Sirte                                                       | gPp     |
| ly-su      | Libya                                        | Baladiyah Surman                                            | gPp     |
| ly-tb      | Libya                                        | Tripoli                                                     | gPp     |
| ly-wa      | Libya                                        | Al Wahat                                                    | gPp     |
| ly-wd      | Libya                                        | Wadi al Hayaa                                               | gPp     |
| ly-ws      | Libya                                        | Wadi al Shatii                                              | gPp     |
| ly-za      | Libya                                        | Az Zawiya District                                          | gPp     |
| ma         |                                              | Morocco                                                     | gPp     |
| ma-01      | Morocco                                      | Tangier-Tetouan-Al Hoceima                                  | gPp     |
| ma-02      | Morocco                                      | Oriental                                                    | gPp     |
| ma-03      | Morocco                                      | Fez-Meknes                                                  | gPp     |
| ma-04      | Morocco                                      | Rabat-Salé-Kénitra                                          | gPp     |
| ma-05      | Morocco                                      | Béni Mellal-Khénifra                                        | gPp     |
| ma-06      | Morocco                                      | Casablanca-Settat                                           | gPp     |
| ma-07      | Morocco                                      | Marrakech-Safi                                              | gPp     |
| ma-08      | Morocco                                      | Drâa-Tafilalet                                              | gPp     |
| ma-09      | Morocco                                      | Souss-Massa                                                 | gPp     |
| ma-10      | Morocco                                      | Guelmim-Oued Noun                                           | gPp     |
| ma-11      | Morocco                                      | Laâyoune-Sakia El Hamra                                     | gPp     |
| ma-12      | Morocco                                      | Dakhla-Oued Ed-Dahab                                        | gPp     |
| mc         |                                              | Monaco                                                      | gPp     |
| md         |                                              | Moldova                                                     | gPp     |
| md-an      | Moldova                                      | Anenii Noi District                                         | gPp     |
| md-ba      | Moldova                                      | Bălți Municipality                                          | gPp     |
| md-bd      | Moldova                                      | Bender Municipality                                         | gPp     |
| md-br      | Moldova                                      | Briceni District                                            | gPp     |
| md-bs      | Moldova                                      | Basarabeasca District                                       | gPp     |
| md-ca      | Moldova                                      | Cahul District                                              | gPp     |
| md-cl      | Moldova                                      | Călărași District                                           | gPp     |
| md-cm      | Moldova                                      | Cimișlia District                                           | gPp     |
| md-cr      | Moldova                                      | Criuleni District                                           | gPp     |
| md-cs      | Moldova                                      | Căușeni District                                            | gPp     |
| md-ct      | Moldova                                      | Cantemir District                                           | gPp     |
| md-cu      | Moldova                                      | Chișinău Municipality                                       | gPp     |
| md-do      | Moldova                                      | Dondușeni District                                          | gPp     |
| md-dr      | Moldova                                      | Drochia District                                            | gPp     |
| md-du      | Moldova                                      | Dubăsari District                                           | gPp     |
| md-ed      | Moldova                                      | Edineț District                                             | gPp     |
| md-fa      | Moldova                                      | Fălești District                                            | gPp     |
| md-fl      | Moldova                                      | Florești District                                           | gPp     |
| md-ga      | Moldova                                      | Gagauzia                                                    | gPp     |
| md-gl      | Moldova                                      | Glodeni District                                            | gPp     |
| md-hi      | Moldova                                      | Hîncești District                                           | gPp     |
| md-ia      | Moldova                                      | Ialoveni District                                           | gPp     |
| md-le      | Moldova                                      | Leova District                                              | gPp     |
| md-ni      | Moldova                                      | Nisporeni District                                          | gPp     |
| md-oc      | Moldova                                      | Ocnița District                                             | gPp     |
| md-or      | Moldova                                      | Orhei District                                              | gPp     |
| md-re      | Moldova                                      | Rezina District                                             | gPp     |
| md-ri      | Moldova                                      | Rîșcani District                                            | gPp     |
| md-sd      | Moldova                                      | Șoldănești District                                         | gPp     |
| md-si      | Moldova                                      | Sîngerei District                                           | gPp     |
| md-sn      | Moldova                                      | Administrative-territorial units of the left bank of Nistru | gPp     |
| md-so      | Moldova                                      | Soroca District                                             | gPp     |
| md-st      | Moldova                                      | Strășeni District                                           | gPp     |
| md-sv      | Moldova                                      | Ștefan Vodă District                                        | gPp     |
| md-ta      | Moldova                                      | Taraclia District                                           | gPp     |
| md-te      | Moldova                                      | Telenești District                                          | gPp     |
| md-un      | Moldova                                      | Ungheni District                                            | gPp     |
| me         |                                              | Montenegro                                                  | gPp     |
| me-01      | Montenegro                                   | Andrijevica Municipality                                    | gPp     |
| me-02      | Montenegro                                   | Bar Municipality                                            | gPp     |
| me-03      | Montenegro                                   | Berane Municipality                                         | gPp     |
| me-04      | Montenegro                                   | Bijelo Polje Municipality                                   | gPp     |
| me-05      | Montenegro                                   | Budva Municipality                                          | gPp     |
| me-06      | Montenegro                                   | Old Royal Capital Cetinje                                   | gPp     |
| me-07      | Montenegro                                   | Danilovgrad Municipality                                    | gPp     |
| me-08      | Montenegro                                   | Herceg Novi Municipality                                    | gPp     |
| me-09      | Montenegro                                   | Kolašin Municipality                                        | gPp     |
| me-10      | Montenegro                                   | Kotor Municipality                                          | gPp     |
| me-11      | Montenegro                                   | Mojkovac Municipality                                       | gPp     |
| me-12      | Montenegro                                   | Nikšić Municipality                                         | gPp     |
| me-13      | Montenegro                                   | Plav Municipality                                           | gPp     |
| me-14      | Montenegro                                   | Pljevlja Municipality                                       | gPp     |
| me-15      | Montenegro                                   | Plužine Municipality                                        | gPp     |
| me-16      | Montenegro                                   | Podgorica Capital City                                      | gPp     |
| me-17      | Montenegro                                   | Rožaje Municipality                                         | gPp     |
| me-18      | Montenegro                                   | Šavnik Municipality                                         | gPp     |
| me-19      | Montenegro                                   | Tivat Municipality                                          | gPp     |
| me-20      | Montenegro                                   | Ulcinj Municipality                                         | gPp     |
| me-21      | Montenegro                                   | Žabljak Municipality                                        | gPp     |
| me-22      | Montenegro                                   | Gusinje Municipality                                        | gPp     |
| me-23      | Montenegro                                   | Petnjica Municipality                                       | gPp     |
| me-24      | Montenegro                                   | Tuzi Municipality                                           | gPp     |
| me-25      | Montenegro                                   | Zeta Municipality                                           | gPp     |
| mg         |                                              | Madagascar                                                  | gPp     |
| mg-a       | Madagascar                                   | Toamasina Province                                          | gPp     |
| mg-d       | Madagascar                                   | Antsiranana Province                                        | gPp     |
| mg-f       | Madagascar                                   | Province de Fianarantsoa                                    | gPp     |
| mg-m       | Madagascar                                   | Province de Mahajanga                                       | gPp     |
| mg-t       | Madagascar                                   | Province d’Antananarivo                                     | gPp     |
| mg-u       | Madagascar                                   | Province de Toliara                                         | gPp     |
| mh         |                                              | Marshall Islands                                            | gPp     |
| mk         |                                              | North Macedonia                                             | gPp     |
| ml         |                                              | Mali                                                        | gPp     |
| ml-1       | Mali                                         | Kayes                                                       | gPp     |
| ml-10      | Mali                                         | Taoudénit Region                                            | gPp     |
| ml-2       | Mali                                         | Koulikoro                                                   | gPp     |
| ml-3       | Mali                                         | Sikasso Region                                              | gPp     |
| ml-4       | Mali                                         | Ségou Region                                                | gPp     |
| ml-5       | Mali                                         | Mopti                                                       | gPp     |
| ml-6       | Mali                                         | Timbuktu                                                    | gPp     |
| ml-7       | Mali                                         | Gao                                                         | gPp     |
| ml-8       | Mali                                         | Kidal                                                       | gPp     |
| ml-9       | Mali                                         | Ménaka                                                      | gPp     |
| ml-bko     | Mali                                         | Bamako                                                      | gPp     |
| mm         |                                              | Myanmar                                                     | gPp     |
| mm-01      | Myanmar                                      | Sagaing Region                                              | gPp     |
| mm-02      | Myanmar                                      | Bago Region                                                 | gPp     |
| mm-03      | Myanmar                                      | Magway                                                      | gPp     |
| mm-04      | Myanmar                                      | Mandalay                                                    | gPp     |
| mm-05      | Myanmar                                      | Tanintharyi Region                                          | gPp     |
| mm-06      | Myanmar                                      | Yangon                                                      | gPp     |
| mm-07      | Myanmar                                      | Ayeyarwady                                                  | gPp     |
| mm-11      | Myanmar                                      | Kachin State                                                | gPp     |
| mm-12      | Myanmar                                      | Kayah State                                                 | gPp     |
| mm-13      | Myanmar                                      | ကရင်ပြည်နယ်                                                    | gPp     |
| mm-14      | Myanmar                                      | Chin                                                        | gPp     |
| mm-15      | Myanmar                                      | Mon State                                                   | gPp     |
| mm-16      | Myanmar                                      | Rakhine                                                     | gPp     |
| mm-17      | Myanmar                                      | Shan State                                                  | gPp     |
| mm-18      | Myanmar                                      | Naypyitaw Union Territory                                   | gPp     |
| mn         |                                              | Mongolia                                                    | gPp     |
| mn-035     | Mongolia                                     | Orkhon                                                      | gPp     |
| mn-037     | Mongolia                                     | Darkhan-Uul                                                 | gPp     |
| mn-039     | Mongolia                                     | Khentii                                                     | gPp     |
| mn-041     | Mongolia                                     | Hovsgel                                                     | gPp     |
| mn-043     | Mongolia                                     | Khovd                                                       | gPp     |
| mn-046     | Mongolia                                     | Uvs                                                         | gPp     |
| mn-047     | Mongolia                                     | Töv                                                         | gPp     |
| mn-049     | Mongolia                                     | Selenge                                                     | gPp     |
| mn-051     | Mongolia                                     | Sükhbaatar                                                  | gPp     |
| mn-053     | Mongolia                                     | Ömnögovi                                                    | gPp     |
| mn-055     | Mongolia                                     | Övörkhangai                                                 | gPp     |
| mn-057     | Mongolia                                     | Zavkhan                                                     | gPp     |
| mn-059     | Mongolia                                     | Dundgovi                                                    | gPp     |
| mn-061     | Mongolia                                     | Dornod                                                      | gPp     |
| mn-063     | Mongolia                                     | Dornogovi                                                   | gPp     |
| mn-064     | Mongolia                                     | Govisumber                                                  | gPp     |
| mn-065     | Mongolia                                     | Govi-Altai                                                  | gPp     |
| mn-067     | Mongolia                                     | Bulgan                                                      | gPp     |
| mn-069     | Mongolia                                     | Bayankhongor                                                | gPp     |
| mn-071     | Mongolia                                     | Bayan-Ölgii                                                 | gPp     |
| mn-073     | Mongolia                                     | Arkhangai                                                   | gPp     |
| mn-1       | Mongolia                                     | Ulaanbaatar                                                 | gPp     |
| mr         |                                              | Mauritania                                                  | gPp     |
| mr-01      | Mauritania                                   | Hodh Ech Chargui                                            | gPp     |
| mr-02      | Mauritania                                   | Hodh El Gharbi                                              | gPp     |
| mr-03      | Mauritania                                   | Assaba                                                      | gPp     |
| mr-04      | Mauritania                                   | Gorgol                                                      | gPp     |
| mr-05      | Mauritania                                   | Brakna                                                      | gPp     |
| mr-06      | Mauritania                                   | Trarza                                                      | gPp     |
| mr-07      | Mauritania                                   | Adrar                                                       | gPp     |
| mr-08      | Mauritania                                   | Dakhlet Nouadhibou                                          | gPp     |
| mr-09      | Mauritania                                   | Tagant                                                      | gPp     |
| mr-10      | Mauritania                                   | Guidimaka                                                   | gPp     |
| mr-11      | Mauritania                                   | Tiris Zemmour                                               | gPp     |
| mr-12      | Mauritania                                   | Inchiri                                                     | gPp     |
| mr-13      | Mauritania                                   | Nouakchott-Ouest Region                                     | gPp     |
| mr-14      | Mauritania                                   | Nouakchott-Nord Region                                      | gPp     |
| mr-15      | Mauritania                                   | Nouakchott-Sud Region                                       | gPp     |
| ms         |                                              | Montserrat                                                  | gPp     |
| mt         |                                              | Malta                                                       | gPp     |
| mu         |                                              | Mauritius                                                   | gPp     |
| mv         |                                              | Maldives                                                    | gPp     |
| mv-00      | Maldives                                     | South Ari Atoll                                             | gPp     |
| mv-01      | Maldives                                     | Addu Atoll                                                  | gPp     |
| mv-02      | Maldives                                     | North Ari Atoll                                             | gPp     |
| mv-03      | Maldives                                     | Faadhippolhu                                                | gPp     |
| mv-04      | Maldives                                     | Felidhu Atoll                                               | gPp     |
| mv-05      | Maldives                                     | Hadhdhunmathi                                               | gPp     |
| mv-07      | Maldives                                     | North Thiladhunmathi                                        | gPp     |
| mv-08      | Maldives                                     | Kolhumadulu                                                 | gPp     |
| mv-12      | Maldives                                     | Mulaku Atoll                                                | gPp     |
| mv-13      | Maldives                                     | North Maalhosmadulu                                         | gPp     |
| mv-14      | Maldives                                     | North Nilandhe Atoll                                        | gPp     |
| mv-17      | Maldives                                     | South Nilandhe Atoll                                        | gPp     |
| mv-20      | Maldives                                     | South Maalhosmadulu                                         | gPp     |
| mv-23      | Maldives                                     | South Thiladhunmathi                                        | gPp     |
| mv-24      | Maldives                                     | North Miladhunmadulu                                        | gPp     |
| mv-25      | Maldives                                     | South Miladhunmadulu                                        | gPp     |
| mv-26      | Maldives                                     | Malé Atoll                                                  | gPp     |
| mv-27      | Maldives                                     | North Huvadhu Atoll                                         | gPp     |
| mv-28      | Maldives                                     | South Huvadhu Atoll                                         | gPp     |
| mv-29      | Maldives                                     | Fuvahmulah                                                  | gPp     |
| mv-mle     | Maldives                                     | Malé                                                        | gPp     |
| mw         |                                              | Malawi                                                      | gPp     |
| mw-c       | Malawi                                       | Central Region, Malawi                                      | gPp     |
| mw-n       | Malawi                                       | Northern Region, Malawi                                     | gPp     |
| mw-s       | Malawi                                       | Southern Region, Malawi                                     | gPp     |
| mx         |                                              | Mexico                                                      | gPp     |
| mx-agu     | Mexico                                       | Aguascalientes                                              | gPp     |
| mx-bcn     | Mexico                                       | Baja California                                             | gPp     |
| mx-bcs     | Mexico                                       | Baja California Sur                                         | gPp     |
| mx-cam     | Mexico                                       | Campeche                                                    | gPp     |
| mx-chh     | Mexico                                       | Chihuahua                                                   | gPp     |
| mx-chp     | Mexico                                       | Chiapas                                                     | gPp     |
| mx-cmx     | Mexico                                       | Mexico City                                                 | gPp     |
| mx-coa     | Mexico                                       | Coahuila                                                    | gPp     |
| mx-col     | Mexico                                       | Colima                                                      | gPp     |
| mx-dur     | Mexico                                       | Durango                                                     | gPp     |
| mx-gro     | Mexico                                       | Guerrero                                                    | gPp     |
| mx-gua     | Mexico                                       | Guanajuato                                                  | gPp     |
| mx-hid     | Mexico                                       | Hidalgo                                                     | gPp     |
| mx-jal     | Mexico                                       | Jalisco                                                     | gPp     |
| mx-mex     | Mexico                                       | State of Mexico                                             | gPp     |
| mx-mic     | Mexico                                       | Michoacán                                                   | gPp     |
| mx-mor     | Mexico                                       | Morelos                                                     | gPp     |
| mx-nay     | Mexico                                       | Nayarit                                                     | gPp     |
| mx-nle     | Mexico                                       | Nuevo León                                                  | gPp     |
| mx-oax     | Mexico                                       | Oaxaca                                                      | gPp     |
| mx-pue     | Mexico                                       | Puebla                                                      | gPp     |
| mx-que     | Mexico                                       | Querétaro                                                   | gPp     |
| mx-roo     | Mexico                                       | Quintana Roo                                                | gPp     |
| mx-sin     | Mexico                                       | Sinaloa                                                     | gPp     |
| mx-slp     | Mexico                                       | San Luis Potosí                                             | gPp     |
| mx-son     | Mexico                                       | Sonora                                                      | gPp     |
| mx-tab     | Mexico                                       | Tabasco                                                     | gPp     |
| mx-tam     | Mexico                                       | Tamaulipas                                                  | gPp     |
| mx-tla     | Mexico                                       | Tlaxcala                                                    | gPp     |
| mx-ver     | Mexico                                       | Veracruz                                                    | gPp     |
| mx-yuc     | Mexico                                       | Yucatán                                                     | gPp     |
| mx-zac     | Mexico                                       | Zacatecas                                                   | gPp     |
| my         |                                              | Malaysia                                                    | gPp     |
| my-01      | Malaysia                                     | Johor                                                       | gPp     |
| my-02      | Malaysia                                     | Kedah                                                       | gPp     |
| my-03      | Malaysia                                     | Kelantan                                                    | gPp     |
| my-04      | Malaysia                                     | Malacca                                                     | gPp     |
| my-05      | Malaysia                                     | Negeri Sembilan                                             | gPp     |
| my-06      | Malaysia                                     | Pahang                                                      | gPp     |
| my-07      | Malaysia                                     | Penang                                                      | gPp     |
| my-08      | Malaysia                                     | Perak                                                       | gPp     |
| my-09      | Malaysia                                     | Perlis                                                      | gPp     |
| my-10      | Malaysia                                     | Selangor                                                    | gPp     |
| my-11      | Malaysia                                     | Terengganu                                                  | gPp     |
| my-12      | Malaysia                                     | Sabah                                                       | gPp     |
| my-13      | Malaysia                                     | Sarawak                                                     | gPp     |
| my-14      | Malaysia                                     | Kuala Lumpur                                                | gPp     |
| my-15      | Malaysia                                     | Labuan                                                      | gPp     |
| my-16      | Malaysia                                     | Putrajaya                                                   | gPp     |
| mz         |                                              | Mozambique                                                  | gPp     |
| mz-a       | Mozambique                                   | Niassa Province                                             | gPp     |
| mz-b       | Mozambique                                   | Manica Province                                             | gPp     |
| mz-g       | Mozambique                                   | Gaza Province                                               | gPp     |
| mz-i       | Mozambique                                   | Inhambane Province                                          | gPp     |
| mz-l       | Mozambique                                   | Maputo Province                                             | gPp     |
| mz-mpm     | Mozambique                                   | Cidade de Maputo                                            | gPp     |
| mz-n       | Mozambique                                   | Nampula Province                                            | gPp     |
| mz-p       | Mozambique                                   | Cabo Delgado Province                                       | gPp     |
| mz-q       | Mozambique                                   | Zambezia Province                                           | gPp     |
| mz-s       | Mozambique                                   | Sofala Province                                             | gPp     |
| mz-t       | Mozambique                                   | Tete Province                                               | gPp     |
| na         |                                              | Namibia                                                     | gPp     |
| na-ca      | Namibia                                      | Zambezi                                                     | gPp     |
| na-er      | Namibia                                      | Erongo Region                                               | gPp     |
| na-ha      | Namibia                                      | Hardap                                                      | gPp     |
| na-ka      | Namibia                                      | Karas                                                       | gPp     |
| na-ke      | Namibia                                      | Kavango East                                                | gPp     |
| na-kh      | Namibia                                      | Khomas                                                      | gPp     |
| na-ku      | Namibia                                      | Kunene Region                                               | gPp     |
| na-kw      | Namibia                                      | Kavango West                                                | gPp     |
| na-od      | Namibia                                      | Otjozondjupa                                                | gPp     |
| na-oh      | Namibia                                      | Omaheke                                                     | gPp     |
| na-on      | Namibia                                      | Oshana                                                      | gPp     |
| na-os      | Namibia                                      | Omusati                                                     | gPp     |
| na-ot      | Namibia                                      | Oshikoto                                                    | gPp     |
| na-ow      | Namibia                                      | Ohangwena                                                   | gPp     |
| ne         |                                              | Niger                                                       | gPp     |
| ne-1       | Niger                                        | Agadez Region                                               | gPp     |
| ne-2       | Niger                                        | Diffa Region                                                | gPp     |
| ne-3       | Niger                                        | Dosso Region                                                | gPp     |
| ne-4       | Niger                                        | Maradi Region                                               | gPp     |
| ne-5       | Niger                                        | Tahoua Region                                               | gPp     |
| ne-6       | Niger                                        | Tillabéri Region                                            | gPp     |
| ne-7       | Niger                                        | Zinder Region                                               | gPp     |
| ne-8       | Niger                                        | Niamey                                                      | gPp     |
| ng         |                                              | Nigeria                                                     | gPp     |
| ng-ab      | Nigeria                                      | Abia State                                                  | gPp     |
| ng-ad      | Nigeria                                      | Adamawa State                                               | gPp     |
| ng-ak      | Nigeria                                      | Akwa Ibom State                                             | gPp     |
| ng-an      | Nigeria                                      | Anambra State                                               | gPp     |
| ng-ba      | Nigeria                                      | Bauchi State                                                | gPp     |
| ng-be      | Nigeria                                      | Benue State                                                 | gPp     |
| ng-bo      | Nigeria                                      | Borno State                                                 | gPp     |
| ng-by      | Nigeria                                      | Bayelsa State                                               | gPp     |
| ng-cr      | Nigeria                                      | Cross River State                                           | gPp     |
| ng-de      | Nigeria                                      | Delta State                                                 | gPp     |
| ng-eb      | Nigeria                                      | Ebonyi State                                                | gPp     |
| ng-ed      | Nigeria                                      | Edo State                                                   | gPp     |
| ng-ek      | Nigeria                                      | Ekiti State                                                 | gPp     |
| ng-en      | Nigeria                                      | Enugu State                                                 | gPp     |
| ng-fc      | Nigeria                                      | Federal Capital Territory                                   | gPp     |
| ng-go      | Nigeria                                      | Gombe State                                                 | gPp     |
| ng-im      | Nigeria                                      | Imo State                                                   | gPp     |
| ng-ji      | Nigeria                                      | Jigawa State                                                | gPp     |
| ng-kd      | Nigeria                                      | Kaduna State                                                | gPp     |
| ng-ke      | Nigeria                                      | Kebbi State                                                 | gPp     |
| ng-kn      | Nigeria                                      | Kano State                                                  | gPp     |
| ng-ko      | Nigeria                                      | Kogi State                                                  | gPp     |
| ng-kt      | Nigeria                                      | Katsina State                                               | gPp     |
| ng-kw      | Nigeria                                      | Kwara State                                                 | gPp     |
| ng-la      | Nigeria                                      | Lagos State                                                 | gPp     |
| ng-na      | Nigeria                                      | Nasarawa State                                              | gPp     |
| ng-ni      | Nigeria                                      | Niger State                                                 | gPp     |
| ng-og      | Nigeria                                      | Ogun State                                                  | gPp     |
| ng-on      | Nigeria                                      | Ondo State                                                  | gPp     |
| ng-os      | Nigeria                                      | Osun State                                                  | gPp     |
| ng-oy      | Nigeria                                      | Oyo State                                                   | gPp     |
| ng-pl      | Nigeria                                      | Plateau State                                               | gPp     |
| ng-ri      | Nigeria                                      | Rivers State                                                | gPp     |
| ng-so      | Nigeria                                      | Sokoto State                                                | gPp     |
| ng-ta      | Nigeria                                      | Taraba State                                                | gPp     |
| ng-yo      | Nigeria                                      | Yobe                                                        | gPp     |
| ng-za      | Nigeria                                      | Zamfara State                                               | gPp     |
| ni         |                                              | Nicaragua                                                   | gPp     |
| ni-an      | Nicaragua                                    | North Caribbean Coast                                       | gPp     |
| ni-as      | Nicaragua                                    | South Caribbean Coast                                       | gPp     |
| ni-bo      | Nicaragua                                    | Boaco                                                       | gPp     |
| ni-ca      | Nicaragua                                    | Carazo                                                      | gPp     |
| ni-ci      | Nicaragua                                    | Chinandega                                                  | gPp     |
| ni-co      | Nicaragua                                    | Chontales                                                   | gPp     |
| ni-es      | Nicaragua                                    | Estelí                                                      | gPp     |
| ni-gr      | Nicaragua                                    | Granada Department                                          | gPp     |
| ni-ji      | Nicaragua                                    | Jinotega                                                    | gPp     |
| ni-le      | Nicaragua                                    | León                                                        | gPp     |
| ni-md      | Nicaragua                                    | Madriz Department                                           | gPp     |
| ni-mn      | Nicaragua                                    | Managua                                                     | gPp     |
| ni-ms      | Nicaragua                                    | Masaya                                                      | gPp     |
| ni-mt      | Nicaragua                                    | Matagalpa                                                   | gPp     |
| ni-ns      | Nicaragua                                    | Nueva Segovia                                               | gPp     |
| ni-ri      | Nicaragua                                    | Rivas                                                       | gPp     |
| ni-sj      | Nicaragua                                    | Río San Juan                                                | gPp     |
| nl         |                                              | Netherlands                                                 | gPp     |
| nl-aw      | Netherlands                                  | Aruba                                                       | gPp     |
| nl-bq1     | Netherlands                                  | Bonaire                                                     | gPp     |
| nl-cw      | Netherlands                                  | Curacao                                                     | gPp     |
| nl-dr      | Netherlands                                  | Drenthe                                                     | gPp     |
| nl-fl      | Netherlands                                  | Flevoland                                                   | gPp     |
| nl-fr      | Netherlands                                  | Frisia                                                      | gPp     |
| nl-ge      | Netherlands                                  | Gelderland                                                  | gPp     |
| nl-gr      | Netherlands                                  | Groningen                                                   | gPp     |
| nl-li      | Netherlands                                  | Limburg                                                     | gPp     |
| nl-nb      | Netherlands                                  | North Brabant                                               | gPp     |
| nl-nh      | Netherlands                                  | North Holland                                               | gPp     |
| nl-ov      | Netherlands                                  | Overijssel                                                  | gPp     |
| nl-sx      | Netherlands                                  | Sint Maarten                                                | gPp     |
| nl-ut      | Netherlands                                  | Utrecht                                                     | gPp     |
| nl-ze      | Netherlands                                  | Zeeland                                                     | gPp     |
| nl-zh      | Netherlands                                  | South Holland                                               | gPp     |
| no         |                                              | Norway                                                      | gPp     |
| no-03      | Norway                                       | Oslo                                                        | gPp     |
| no-11      | Norway                                       | Rogaland                                                    | gPp     |
| no-15      | Norway                                       | Møre og Romsdal                                             | gPp     |
| no-18      | Norway                                       | Nordland                                                    | gPp     |
| no-21      | Norway                                       | Svalbard                                                    | gPp     |
| no-22      | Norway                                       | Jan Mayen                                                   | gPp     |
| no-31      | Norway                                       | Østfold                                                     | gPp     |
| no-32      | Norway                                       | Akershus                                                    | gPp     |
| no-33      | Norway                                       | Buskerud                                                    | gPp     |
| no-34      | Norway                                       | Innlandet                                                   | gPp     |
| no-39      | Norway                                       | Vestfold                                                    | gPp     |
| no-40      | Norway                                       | Telemark                                                    | gPp     |
| no-42      | Norway                                       | Agder                                                       | gPp     |
| no-46      | Norway                                       | Vestland                                                    | gPp     |
| no-50      | Norway                                       | Trøndelag                                                   | gPp     |
| no-55      | Norway                                       | Troms                                                       | gPp     |
| no-56      | Norway                                       | Finnmark                                                    | gPp     |
| no-bv      | Norway                                       | Bouvet Island                                               | gPp     |
| np         |                                              | Nepal                                                       | gPp     |
| np-p1      | Nepal                                        | Koshi Province                                              | gPp     |
| np-p2      | Nepal                                        | Madhesh Province                                            | gPp     |
| np-p3      | Nepal                                        | Bagamati Province                                           | gPp     |
| np-p4      | Nepal                                        | Gandaki Province                                            | gPp     |
| np-p5      | Nepal                                        | Lumbini Province                                            | gPp     |
| np-p6      | Nepal                                        | Karnali Province                                            | gPp     |
| np-p7      | Nepal                                        | Sudurpashchim Province                                      | gPp     |
| nr         |                                              | Nauru                                                       | gPp     |
| nr-01      | Nauru                                        | Aiwo District                                               | gPp     |
| nr-02      | Nauru                                        | Anabar District                                             | gPp     |
| nr-03      | Nauru                                        | Anetan District                                             | gPp     |
| nr-04      | Nauru                                        | Anibare District                                            | gPp     |
| nr-05      | Nauru                                        | Baitsi District                                             | gPp     |
| nr-06      | Nauru                                        | Boe District                                                | gPp     |
| nr-07      | Nauru                                        | Buada District                                              | gPp     |
| nr-08      | Nauru                                        | Denigomodu District                                         | gPp     |
| nr-09      | Nauru                                        | Ewa District                                                | gPp     |
| nr-10      | Nauru                                        | Ijuw District                                               | gPp     |
| nr-11      | Nauru                                        | Meneng District                                             | gPp     |
| nr-12      | Nauru                                        | Nibok District                                              | gPp     |
| nr-13      | Nauru                                        | Uaboe District                                              | gPp     |
| nr-14      | Nauru                                        | Yaren District                                              | gPp     |
| nu         |                                              | Niue                                                        | gPp     |
| nz         |                                              | New Zealand                                                 | gPp     |
| nz-auk     | New Zealand                                  | Auckland                                                    | gPp     |
| nz-bop     | New Zealand                                  | Bay of Plenty                                               | gPp     |
| nz-can     | New Zealand                                  | Canterbury                                                  | gPp     |
| nz-cit     | New Zealand                                  | Chatham Islands                                             | gPp     |
| nz-gis     | New Zealand                                  | Gisborne                                                    | gPp     |
| nz-hkb     | New Zealand                                  | Hawke's Bay                                                 | gPp     |
| nz-mbh     | New Zealand                                  | Marlborough                                                 | gPp     |
| nz-mwt     | New Zealand                                  | Manawatū-Whanganui                                          | gPp     |
| nz-nsn     | New Zealand                                  | Nelson                                                      | gPp     |
| nz-ntl     | New Zealand                                  | Northland                                                   | gPp     |
| nz-ota     | New Zealand                                  | Otago                                                       | gPp     |
| nz-stl     | New Zealand                                  | Southland                                                   | gPp     |
| nz-tas     | New Zealand                                  | Tasman                                                      | gPp     |
| nz-tki     | New Zealand                                  | Taranaki                                                    | gPp     |
| nz-wgn     | New Zealand                                  | Wellington                                                  | gPp     |
| nz-wko     | New Zealand                                  | Waikato                                                     | gPp     |
| nz-wtc     | New Zealand                                  | West Coast                                                  | gPp     |
| om         |                                              | Oman                                                        | gPp     |
| om-bj      | Oman                                         | Al Batinah South Governorate                                | gPp     |
| om-bs      | Oman                                         | Al Batinah North Governorate                                | gPp     |
| om-bu      | Oman                                         | Al Buraimi Governorate                                      | gPp     |
| om-da      | Oman                                         | Ad Dakhiliyah Governorate                                   | gPp     |
| om-ma      | Oman                                         | Muscat Governorate                                          | gPp     |
| om-mu      | Oman                                         | Musandam Governorate                                        | gPp     |
| om-sj      | Oman                                         | Ash Sharqiyah South Governorate                             | gPp     |
| om-ss      | Oman                                         | Ash Sharqiyah North Governorate                             | gPp     |
| om-wu      | Oman                                         | Al Wusta Governorate                                        | gPp     |
| om-za      | Oman                                         | Ad Dhahirah Governorate                                     | gPp     |
| om-zu      | Oman                                         | Dhofar Governorate                                          | gPp     |
| pa         |                                              | Panama                                                      | gPp     |
| pa-1       | Panama                                       | Bocas del Toro                                              | gPp     |
| pa-10      | Panama                                       | Panamá Oeste                                                | gPp     |
| pa-2       | Panama                                       | Coclé                                                       | gPp     |
| pa-3       | Panama                                       | Colón                                                       | gPp     |
| pa-4       | Panama                                       | Chiriquí                                                    | gPp     |
| pa-5       | Panama                                       | Darién                                                      | gPp     |
| pa-6       | Panama                                       | Herrera                                                     | gPp     |
| pa-7       | Panama                                       | Los Santos                                                  | gPp     |
| pa-8       | Panama                                       | Panamá Province                                             | gPp     |
| pa-9       | Panama                                       | Veraguas                                                    | gPp     |
| pa-em      | Panama                                       | Emberá-Wounaan                                              | gPp     |
| pa-ky      | Panama                                       | Guna Yala                                                   | gPp     |
| pa-nb      | Panama                                       | Ngäbe-Buglé                                                 | gPp     |
| pa-nt      | Panama                                       | Naso Tjër Di                                                | gPp     |
| pe         |                                              | Peru                                                        | gPp     |
| pe-ama     | Peru                                         | Amazonas                                                    | gPp     |
| pe-anc     | Peru                                         | Ancash                                                      | gPp     |
| pe-apu     | Peru                                         | Apurímac                                                    | gPp     |
| pe-are     | Peru                                         | Arequipa                                                    | gPp     |
| pe-aya     | Peru                                         | Ayacucho                                                    | gPp     |
| pe-caj     | Peru                                         | Cajamarca                                                   | gPp     |
| pe-cal     | Peru                                         | Callao                                                      | gPp     |
| pe-cus     | Peru                                         | Cusco                                                       | gPp     |
| pe-huc     | Peru                                         | Huánuco                                                     | gPp     |
| pe-huv     | Peru                                         | Huancavelica                                                | gPp     |
| pe-ica     | Peru                                         | Ica                                                         | gPp     |
| pe-jun     | Peru                                         | Junín                                                       | gPp     |
| pe-lal     | Peru                                         | La Libertad                                                 | gPp     |
| pe-lam     | Peru                                         | Lambayeque                                                  | gPp     |
| pe-lim     | Peru                                         | Lima                                                        | gPp     |
| pe-lor     | Peru                                         | Loreto                                                      | gPp     |
| pe-mdd     | Peru                                         | Madre de Dios                                               | gPp     |
| pe-moq     | Peru                                         | Moquegua                                                    | gPp     |
| pe-pas     | Peru                                         | Pasco                                                       | gPp     |
| pe-piu     | Peru                                         | Piura                                                       | gPp     |
| pe-pun     | Peru                                         | Puno                                                        | gPp     |
| pe-sam     | Peru                                         | San Martín                                                  | gPp     |
| pe-tac     | Peru                                         | Tacna                                                       | gPp     |
| pe-tum     | Peru                                         | Tumbes                                                      | gPp     |
| pe-uca     | Peru                                         | Ucayali                                                     | gPp     |
| pg         |                                              | Papua New Guinea                                            | gPp     |
| pg-cpk     | Papua New Guinea                             | Simbu                                                       | gPp     |
| pg-cpm     | Papua New Guinea                             | Central                                                     | gPp     |
| pg-ebr     | Papua New Guinea                             | East New Britain                                            | gPp     |
| pg-ehg     | Papua New Guinea                             | Eastern Highlands                                           | gPp     |
| pg-epw     | Papua New Guinea                             | Enga                                                        | gPp     |
| pg-esw     | Papua New Guinea                             | East Sepik                                                  | gPp     |
| pg-gpk     | Papua New Guinea                             | Gulf                                                        | gPp     |
| pg-hla     | Papua New Guinea                             | Hela                                                        | gPp     |
| pg-jwk     | Papua New Guinea                             | Jiwaka                                                      | gPp     |
| pg-mba     | Papua New Guinea                             | Milne Bay                                                   | gPp     |
| pg-mpl     | Papua New Guinea                             | Morobe                                                      | gPp     |
| pg-mpm     | Papua New Guinea                             | Madang                                                      | gPp     |
| pg-mrl     | Papua New Guinea                             | Manus                                                       | gPp     |
| pg-ncd     | Papua New Guinea                             | National Capital District                                   | gPp     |
| pg-nik     | Papua New Guinea                             | New Ireland                                                 | gPp     |
| pg-npp     | Papua New Guinea                             | Oro                                                         | gPp     |
| pg-nsb     | Papua New Guinea                             | Autonomous Region of Bougainville                           | gPp     |
| pg-san     | Papua New Guinea                             | Sandaun                                                     | gPp     |
| pg-shm     | Papua New Guinea                             | Southern Highlands                                          | gPp     |
| pg-wbk     | Papua New Guinea                             | West New Britain                                            | gPp     |
| pg-whm     | Papua New Guinea                             | Western Highlands                                           | gPp     |
| pg-wpd     | Papua New Guinea                             | Western                                                     | gPp     |
| ph         |                                              | Philippines                                                 | gPp     |
| ph-00      | Philippines                                  | Metro Manila                                                | gPp     |
| ph-01      | Philippines                                  | Ilocos Region                                               | gPp     |
| ph-02      | Philippines                                  | Cagayan Valley                                              | gPp     |
| ph-03      | Philippines                                  | Central Luzon                                               | gPp     |
| ph-05      | Philippines                                  | Bicol Region                                                | gPp     |
| ph-06      | Philippines                                  | Western Visayas                                             | gPp     |
| ph-07      | Philippines                                  | Central Visayas                                             | gPp     |
| ph-08      | Philippines                                  | Eastern Visayas                                             | gPp     |
| ph-09      | Philippines                                  | Zamboanga Peninsula                                         | gPp     |
| ph-10      | Philippines                                  | Northern Mindanao                                           | gPp     |
| ph-11      | Philippines                                  | Davao Region                                                | gPp     |
| ph-12      | Philippines                                  | Soccsksargen                                                | gPp     |
| ph-13      | Philippines                                  | Caraga                                                      | gPp     |
| ph-14      | Philippines                                  | Bangsamoro                                                  | gPp     |
| ph-15      | Philippines                                  | Cordillera Administrative Region                            | gPp     |
| ph-40      | Philippines                                  | Calabarzon                                                  | gPp     |
| ph-41      | Philippines                                  | Mimaropa                                                    | gPp     |
| ph-nec     | Philippines                                  | Negros Occidental                                           | gPp     |
| ph-ner     | Philippines                                  | Negros Oriental                                             | gPp     |
| ph-sig     | Philippines                                  | Siquijor                                                    | gPp     |
| pk         |                                              | Pakistan                                                    | gPp     |
| pk-ba      | Pakistan                                     | Balochistan                                                 | gPp     |
| pk-gb      | Pakistan                                     | Gilgit-Baltistan                                            | gPp     |
| pk-is      | Pakistan                                     | Islamabad Capital Territory                                 | gPp     |
| pk-jk      | Pakistan                                     | Azad Kashmir                                                | gPp     |
| pk-kp      | Pakistan                                     | Khyber Pakhtunkhwa                                          | gPp     |
| pk-pb      | Pakistan                                     | Punjab                                                      | gPp     |
| pk-sd      | Pakistan                                     | Sindh                                                       | gPp     |
| pl         |                                              | Poland                                                      | gPp     |
| pl-02      | Poland                                       | Lower Silesian Voivodeship                                  | gPp     |
| pl-04      | Poland                                       | Kuyavian-Pomeranian Voivodeship                             | gPp     |
| pl-06      | Poland                                       | Lublin Voivodeship                                          | gPp     |
| pl-08      | Poland                                       | Lubusz Voivodeship                                          | gPp     |
| pl-10      | Poland                                       | Łódź Voivodeship                                            | gPp     |
| pl-12      | Poland                                       | Lesser Poland Voivodeship                                   | gPp     |
| pl-14      | Poland                                       | Masovian Voivodeship                                        | gPp     |
| pl-16      | Poland                                       | Opole Voivodeship                                           | gPp     |
| pl-18      | Poland                                       | Subcarpathian Voivodeship                                   | gPp     |
| pl-20      | Poland                                       | Podlachia Voivodeship                                       | gPp     |
| pl-22      | Poland                                       | Pomeranian Voivodeship                                      | gPp     |
| pl-24      | Poland                                       | Silesian Voivodeship                                        | gPp     |
| pl-26      | Poland                                       | Holy Cross Voivodeship                                      | gPp     |
| pl-28      | Poland                                       | Warmian-Masurian Voivodeship                                | gPp     |
| pl-30      | Poland                                       | Greater Poland Voivodeship                                  | gPp     |
| pl-32      | Poland                                       | West Pomeranian Voivodeship                                 | gPp     |
| pn         |                                              | Pitcairn Islands                                            | gPp     |
| ps         |                                              | Palestinian Territories                                     | gPp     |
| ps-gzz     | Palestinian Territories                      | Gaza Strip                                                  | gPp     |
| ps-wbk     | Palestinian Territories                      | West Bank                                                   | gPp     |
| pt         |                                              | Portugal                                                    | gPp     |
| pt-01      | Portugal                                     | Aveiro                                                      | gPp     |
| pt-02      | Portugal                                     | Beja                                                        | gPp     |
| pt-03      | Portugal                                     | Braga                                                       | gPp     |
| pt-04      | Portugal                                     | Bragança                                                    | gPp     |
| pt-05      | Portugal                                     | Castelo Branco                                              | gPp     |
| pt-06      | Portugal                                     | Coimbra                                                     | gPp     |
| pt-07      | Portugal                                     | Évora                                                       | gPp     |
| pt-08      | Portugal                                     | Faro                                                        | gPp     |
| pt-09      | Portugal                                     | Guarda                                                      | gPp     |
| pt-10      | Portugal                                     | Leiria                                                      | gPp     |
| pt-11      | Portugal                                     | Lisbon                                                      | gPp     |
| pt-12      | Portugal                                     | Portalegre                                                  | gPp     |
| pt-13      | Portugal                                     | Porto                                                       | gPp     |
| pt-14      | Portugal                                     | Santarém                                                    | gPp     |
| pt-15      | Portugal                                     | Setúbal                                                     | gPp     |
| pt-16      | Portugal                                     | Viana do Castelo                                            | gPp     |
| pt-17      | Portugal                                     | Vila Real                                                   | gPp     |
| pt-18      | Portugal                                     | Viseu                                                       | gPp     |
| pt-20      | Portugal                                     | Azores                                                      | gPp     |
| pt-30      | Portugal                                     | Madeira                                                     | gPp     |
| pw         |                                              | Palau                                                       | gPp     |
| pw-002     | Palau                                        | Aimeliik                                                    | gPp     |
| pw-004     | Palau                                        | Airai                                                       | gPp     |
| pw-010     | Palau                                        | Angaur                                                      | gPp     |
| pw-050     | Palau                                        | Hatohobei                                                   | gPp     |
| pw-100     | Palau                                        | Kayangel                                                    | gPp     |
| pw-150     | Palau                                        | Koror                                                       | gPp     |
| pw-212     | Palau                                        | Melekeok                                                    | gPp     |
| pw-214     | Palau                                        | Ngaraard                                                    | gPp     |
| pw-218     | Palau                                        | Ngarchelong                                                 | gPp     |
| pw-222     | Palau                                        | Ngardmau                                                    | gPp     |
| pw-224     | Palau                                        | Ngatpang                                                    | gPp     |
| pw-226     | Palau                                        | Ngchesar                                                    | gPp     |
| pw-227     | Palau                                        | Ngeremlengui                                                | gPp     |
| pw-228     | Palau                                        | Ngiwal                                                      | gPp     |
| pw-350     | Palau                                        | Peleliu                                                     | gPp     |
| pw-370     | Palau                                        | Sonsorol                                                    | gPp     |
| py         |                                              | Paraguay                                                    | gPp     |
| py-1       | Paraguay                                     | Concepción                                                  | gPp     |
| py-10      | Paraguay                                     | Alto Paraná                                                 | gPp     |
| py-11      | Paraguay                                     | Central Department                                          | gPp     |
| py-12      | Paraguay                                     | Ñeembucú                                                    | gPp     |
| py-13      | Paraguay                                     | Amambay                                                     | gPp     |
| py-14      | Paraguay                                     | Canindeyú                                                   | gPp     |
| py-15      | Paraguay                                     | Presidente Hayes                                            | gPp     |
| py-16      | Paraguay                                     | Alto Paraguay                                               | gPp     |
| py-19      | Paraguay                                     | Boquerón                                                    | gPp     |
| py-2       | Paraguay                                     | San Pedro                                                   | gPp     |
| py-3       | Paraguay                                     | Cordillera Department                                       | gPp     |
| py-4       | Paraguay                                     | Guairá                                                      | gPp     |
| py-5       | Paraguay                                     | Caaguazú                                                    | gPp     |
| py-6       | Paraguay                                     | Caazapá                                                     | gPp     |
| py-7       | Paraguay                                     | Itapúa                                                      | gPp     |
| py-8       | Paraguay                                     | Misiones                                                    | gPp     |
| py-9       | Paraguay                                     | Department of Paraguari                                     | gPp     |
| qa         |                                              | Qatar                                                       | gPp     |
| qa-da      | Qatar                                        | Doha                                                        | gPp     |
| qa-kh      | Qatar                                        | Al Khor and Al Thakhira                                     | gPp     |
| qa-ms      | Qatar                                        | Ash Shamal                                                  | gPp     |
| qa-ra      | Qatar                                        | Al Rayyan                                                   | gPp     |
| qa-sh      | Qatar                                        | Al Shahaniya                                                | gPp     |
| qa-us      | Qatar                                        | Umm Salal                                                   | gPp     |
| qa-wa      | Qatar                                        | Al Wakrah                                                   | gPp     |
| qa-za      | Qatar                                        | Al-Daayen                                                   | gPp     |
| ro         |                                              | Romania                                                     | gPp     |
| ro-ab      | Romania                                      | Alba                                                        | gPp     |
| ro-ag      | Romania                                      | Argeș                                                       | gPp     |
| ro-ar      | Romania                                      | Arad                                                        | gPp     |
| ro-b       | Romania                                      | Bucharest                                                   | gPp     |
| ro-bc      | Romania                                      | Bacău                                                       | gPp     |
| ro-bh      | Romania                                      | Bihor                                                       | gPp     |
| ro-bn      | Romania                                      | Bistrița-Năsăud                                             | gPp     |
| ro-br      | Romania                                      | Brăila                                                      | gPp     |
| ro-bt      | Romania                                      | Botoșani                                                    | gPp     |
| ro-bv      | Romania                                      | Brașov                                                      | gPp     |
| ro-bz      | Romania                                      | Buzău                                                       | gPp     |
| ro-cj      | Romania                                      | Cluj                                                        | gPp     |
| ro-cl      | Romania                                      | Călărași                                                    | gPp     |
| ro-cs      | Romania                                      | Caraș-Severin                                               | gPp     |
| ro-ct      | Romania                                      | Constanța                                                   | gPp     |
| ro-cv      | Romania                                      | Covasna                                                     | gPp     |
| ro-db      | Romania                                      | Dâmbovița                                                   | gPp     |
| ro-dj      | Romania                                      | Dolj                                                        | gPp     |
| ro-gj      | Romania                                      | Gorj                                                        | gPp     |
| ro-gl      | Romania                                      | Galați                                                      | gPp     |
| ro-gr      | Romania                                      | Giurgiu                                                     | gPp     |
| ro-hd      | Romania                                      | Hunedoara                                                   | gPp     |
| ro-hr      | Romania                                      | Harghita                                                    | gPp     |
| ro-if      | Romania                                      | Ilfov                                                       | gPp     |
| ro-il      | Romania                                      | Ialomița                                                    | gPp     |
| ro-is      | Romania                                      | Iași                                                        | gPp     |
| ro-mh      | Romania                                      | Mehedinți                                                   | gPp     |
| ro-mm      | Romania                                      | Maramureș                                                   | gPp     |
| ro-ms      | Romania                                      | Mureș                                                       | gPp     |
| ro-nt      | Romania                                      | Neamț                                                       | gPp     |
| ro-ot      | Romania                                      | Olt                                                         | gPp     |
| ro-ph      | Romania                                      | Prahova                                                     | gPp     |
| ro-sb      | Romania                                      | Sibiu                                                       | gPp     |
| ro-sj      | Romania                                      | Sălaj                                                       | gPp     |
| ro-sm      | Romania                                      | Satu Mare                                                   | gPp     |
| ro-sv      | Romania                                      | Suceava                                                     | gPp     |
| ro-tl      | Romania                                      | Tulcea                                                      | gPp     |
| ro-tm      | Romania                                      | Timiș                                                       | gPp     |
| ro-tr      | Romania                                      | Teleorman                                                   | gPp     |
| ro-vl      | Romania                                      | Vâlcea                                                      | gPp     |
| ro-vn      | Romania                                      | Vrancea                                                     | gPp     |
| ro-vs      | Romania                                      | Vaslui                                                      | gPp     |
| rs         |                                              | Serbia                                                      | gPp     |
| rs-00      | Serbia                                       | City of Belgrade                                            | gPp     |
| rs-08      | Serbia                                       | Macva Administrative District                               | gPp     |
| rs-09      | Serbia                                       | Kolubara Administrative District                            | gPp     |
| rs-10      | Serbia                                       | Podunavlje Administrative District                          | gPp     |
| rs-11      | Serbia                                       | Branicevo Administrative District                           | gPp     |
| rs-12      | Serbia                                       | Sumadija Administrative District                            | gPp     |
| rs-13      | Serbia                                       | Pomoravlje Administrative District                          | gPp     |
| rs-14      | Serbia                                       | Bor Administrative District                                 | gPp     |
| rs-15      | Serbia                                       | Zajecar Administrative District                             | gPp     |
| rs-16      | Serbia                                       | Zlatibor Administrative District                            | gPp     |
| rs-17      | Serbia                                       | Moravica Administrative District                            | gPp     |
| rs-18      | Serbia                                       | Raska Administrative District                               | gPp     |
| rs-19      | Serbia                                       | Rasina Administrative District                              | gPp     |
| rs-20      | Serbia                                       | Nisava Administrative District                              | gPp     |
| rs-21      | Serbia                                       | Toplica Administrative District                             | gPp     |
| rs-22      | Serbia                                       | Pirot Administrative District                               | gPp     |
| rs-23      | Serbia                                       | Jablanica Administrative District                           | gPp     |
| rs-24      | Serbia                                       | Pcinja Administrative District                              | gPp     |
| rs-vo      | Serbia                                       | Vojvodina                                                   | gPp     |
| ru         |                                              | Russia                                                      | gPp     |
| ru-ad      | Russia                                       | Republic of Adygea                                          | gPp     |
| ru-al      | Russia                                       | Altai Republic                                              | gPp     |
| ru-alt     | Russia                                       | Altai Krai                                                  | gPp     |
| ru-amu     | Russia                                       | Amur Oblast                                                 | gPp     |
| ru-ark     | Russia                                       | Arkhangelsk Oblast                                          | gPp     |
| ru-ast     | Russia                                       | Astrakhan Oblast                                            | gPp     |
| ru-ba      | Russia                                       | Bashkortostan                                               | gPp     |
| ru-bel     | Russia                                       | Belgorod Oblast                                             | gPp     |
| ru-bry     | Russia                                       | Bryansk Oblast                                              | gPp     |
| ru-bu      | Russia                                       | Buryatia                                                    | gPp     |
| ru-ce      | Russia                                       | Chechnya                                                    | gPp     |
| ru-che     | Russia                                       | Chelyabinsk Oblast                                          | gPp     |
| ru-chu     | Russia                                       | Chukotka Autonomous Okrug                                   | gPp     |
| ru-cu      | Russia                                       | Chuvashia                                                   | gPp     |
| ru-da      | Russia                                       | Dagestan                                                    | gPp     |
| ru-in      | Russia                                       | Ingushetia                                                  | gPp     |
| ru-irk     | Russia                                       | Irkutsk Oblast                                              | gPp     |
| ru-iva     | Russia                                       | Ivanovo Oblast                                              | gPp     |
| ru-kam     | Russia                                       | Kamchatka Krai                                              | gPp     |
| ru-kb      | Russia                                       | Kabardino-Balkaria                                          | gPp     |
| ru-kc      | Russia                                       | Karachay-Cherkessia                                         | gPp     |
| ru-kda     | Russia                                       | Krasnodar Krai                                              | gPp     |
| ru-kem     | Russia                                       | Kemerovo Oblast–Kuzbass                                     | gPp     |
| ru-kgd     | Russia                                       | Kaliningrad                                                 | gPp     |
| ru-kgn     | Russia                                       | Kurgan Oblast                                               | gPp     |
| ru-kha     | Russia                                       | Khabarovsk Krai                                             | gPp     |
| ru-khm     | Russia                                       | Khanty-Mansiysk Autonomous Okrug – Ugra                     | gPp     |
| ru-kir     | Russia                                       | Kirov Oblast                                                | gPp     |
| ru-kk      | Russia                                       | Khakassia                                                   | gPp     |
| ru-kl      | Russia                                       | Republic of Kalmykia                                        | gPp     |
| ru-klu     | Russia                                       | Kaluga Oblast                                               | gPp     |
| ru-ko      | Russia                                       | Komi Republic                                               | gPp     |
| ru-kos     | Russia                                       | Kostroma Oblast                                             | gPp     |
| ru-kr      | Russia                                       | Karelia                                                     | gPp     |
| ru-krs     | Russia                                       | Kursk Oblast                                                | gPp     |
| ru-kya     | Russia                                       | Krasnoyarsk Krai                                            | gPp     |
| ru-len     | Russia                                       | Leningrad Oblast                                            | gPp     |
| ru-lip     | Russia                                       | Lipetsk Oblast                                              | gPp     |
| ru-mag     | Russia                                       | Magadan Oblast                                              | gPp     |
| ru-me      | Russia                                       | Mari El Republic                                            | gPp     |
| ru-mo      | Russia                                       | Republic of Mordovia                                        | gPp     |
| ru-mos     | Russia                                       | Moscow Oblast                                               | gPp     |
| ru-mow     | Russia                                       | Moscow                                                      | gPp     |
| ru-mur     | Russia                                       | Murmansk Oblast                                             | gPp     |
| ru-nen     | Russia                                       | Nenets Autonomous Okrug                                     | gPp     |
| ru-ngr     | Russia                                       | Novgorod Oblast                                             | gPp     |
| ru-niz     | Russia                                       | Nizhny Novgorod Oblast                                      | gPp     |
| ru-nvs     | Russia                                       | Novosibirsk Oblast                                          | gPp     |
| ru-oms     | Russia                                       | Omsk Oblast                                                 | gPp     |
| ru-ore     | Russia                                       | Orenburg Oblast                                             | gPp     |
| ru-orl     | Russia                                       | Oryol Oblast                                                | gPp     |
| ru-per     | Russia                                       | Perm Krai                                                   | gPp     |
| ru-pnz     | Russia                                       | Penza Oblast                                                | gPp     |
| ru-pri     | Russia                                       | Primorsky Krai                                              | gPp     |
| ru-psk     | Russia                                       | Pskov Oblast                                                | gPp     |
| ru-ros     | Russia                                       | Rostov Oblast                                               | gPp     |
| ru-rya     | Russia                                       | Ryazan Oblast                                               | gPp     |
| ru-sa      | Russia                                       | Sakha Republic                                              | gPp     |
| ru-sak     | Russia                                       | Sakhalin Oblast                                             | gPp     |
| ru-sam     | Russia                                       | Samara Oblast                                               | gPp     |
| ru-sar     | Russia                                       | Saratov Oblast                                              | gPp     |
| ru-se      | Russia                                       | Republic of North Ossetia – Alania                          | gPp     |
| ru-smo     | Russia                                       | Smolensk Oblast                                             | gPp     |
| ru-spe     | Russia                                       | Saint Petersburg                                            | gPp     |
| ru-sta     | Russia                                       | Stavropol Krai                                              | gPp     |
| ru-sve     | Russia                                       | Sverdlovsk Oblast                                           | gPp     |
| ru-ta      | Russia                                       | Tatarstan                                                   | gPp     |
| ru-tam     | Russia                                       | Tambov Oblast                                               | gPp     |
| ru-tom     | Russia                                       | Tomsk Oblast                                                | gPp     |
| ru-tul     | Russia                                       | Tula Oblast                                                 | gPp     |
| ru-tve     | Russia                                       | Tver Oblast                                                 | gPp     |
| ru-ty      | Russia                                       | Tuva Republic                                               | gPp     |
| ru-tyu     | Russia                                       | Tyumen Oblast                                               | gPp     |
| ru-ud      | Russia                                       | Udmurtia                                                    | gPp     |
| ru-uly     | Russia                                       | Ulyanovsk Oblast                                            | gPp     |
| ru-vgg     | Russia                                       | Volgograd Oblast                                            | gPp     |
| ru-vla     | Russia                                       | Vladimir Oblast                                             | gPp     |
| ru-vlg     | Russia                                       | Vologda Oblast                                              | gPp     |
| ru-vor     | Russia                                       | Voronezh Oblast                                             | gPp     |
| ru-yan     | Russia                                       | Yamalo-Nenets Autonomous Okrug                              | gPp     |
| ru-yar     | Russia                                       | Yaroslavl Oblast                                            | gPp     |
| ru-yev     | Russia                                       | Jewish Autonomous Oblast                                    | gPp     |
| ru-zab     | Russia                                       | Zabaykalsky Krai                                            | gPp     |
| rw         |                                              | Rwanda                                                      | gPp     |
| rw-01      | Rwanda                                       | City of Kigali                                              | gPp     |
| rw-02      | Rwanda                                       | Eastern Province                                            | gPp     |
| rw-03      | Rwanda                                       | Northern Province                                           | gPp     |
| rw-04      | Rwanda                                       | Western Province                                            | gPp     |
| rw-05      | Rwanda                                       | Southern Province                                           | gPp     |
| sa         |                                              | Saudi Arabia                                                | gPp     |
| sa-01      | Saudi Arabia                                 | Riyadh Region                                               | gPp     |
| sa-02      | Saudi Arabia                                 | Makkah Region                                               | gPp     |
| sa-03      | Saudi Arabia                                 | Medina Province                                             | gPp     |
| sa-04      | Saudi Arabia                                 | Eastern Province                                            | gPp     |
| sa-05      | Saudi Arabia                                 | Al-Qassim Province                                          | gPp     |
| sa-06      | Saudi Arabia                                 | Ḥa'il Province                                              | gPp     |
| sa-07      | Saudi Arabia                                 | Tabuk Province                                              | gPp     |
| sa-08      | Saudi Arabia                                 | Northern Borders Province                                   | gPp     |
| sa-09      | Saudi Arabia                                 | Jazan Province                                              | gPp     |
| sa-10      | Saudi Arabia                                 | Najran Region                                               | gPp     |
| sa-11      | Saudi Arabia                                 | Al-Bahah Province                                           | gPp     |
| sa-12      | Saudi Arabia                                 | Al Jawf Region                                              | gPp     |
| sa-14      | Saudi Arabia                                 | 'Asir Province                                              | gPp     |
| sb         |                                              | Solomon Islands                                             | gPp     |
| sb-ce      | Solomon Islands                              | Central Province                                            | gPp     |
| sb-ch      | Solomon Islands                              | Choiseul                                                    | gPp     |
| sb-ct      | Solomon Islands                              | Honiara                                                     | gPp     |
| sb-gu      | Solomon Islands                              | Guadalcanal                                                 | gPp     |
| sb-is      | Solomon Islands                              | Isabel                                                      | gPp     |
| sb-mk      | Solomon Islands                              | Makira-Ulawa                                                | gPp     |
| sb-ml      | Solomon Islands                              | Malaita                                                     | gPp     |
| sb-rb      | Solomon Islands                              | Rennell and Bellona                                         | gPp     |
| sb-te      | Solomon Islands                              | Temotu                                                      | gPp     |
| sb-we      | Solomon Islands                              | Western                                                     | gPp     |
| sc         |                                              | Seychelles                                                  | gPp     |
| sc-01      | Seychelles                                   | Anse Aux Pins                                               | gPp     |
| sc-02      | Seychelles                                   | Anse Boileau                                                | gPp     |
| sc-03      | Seychelles                                   | Anse Etoile                                                 | gPp     |
| sc-04      | Seychelles                                   | Au Cap                                                      | gPp     |
| sc-05      | Seychelles                                   | Anse Royale                                                 | gPp     |
| sc-06      | Seychelles                                   | Baie Lazare                                                 | gPp     |
| sc-07      | Seychelles                                   | Baie Sainte Anne Praslin                                    | gPp     |
| sc-08      | Seychelles                                   | Beau Vallon                                                 | gPp     |
| sc-09      | Seychelles                                   | Bel Air                                                     | gPp     |
| sc-10      | Seychelles                                   | Bel Ombre                                                   | gPp     |
| sc-11      | Seychelles                                   | Cascade                                                     | gPp     |
| sc-12      | Seychelles                                   | Glacis                                                      | gPp     |
| sc-13      | Seychelles                                   | Grand Anse Mahe                                             | gPp     |
| sc-14      | Seychelles                                   | Grand Anse Praslin                                          | gPp     |
| sc-15      | Seychelles                                   | La Digue                                                    | gPp     |
| sc-16      | Seychelles                                   | English River                                               | gPp     |
| sc-17      | Seychelles                                   | Mont Buxton                                                 | gPp     |
| sc-18      | Seychelles                                   | Mont Fleuri                                                 | gPp     |
| sc-19      | Seychelles                                   | Plaisance                                                   | gPp     |
| sc-20      | Seychelles                                   | Pointe La Rue                                               | gPp     |
| sc-21      | Seychelles                                   | Port Glaud                                                  | gPp     |
| sc-22      | Seychelles                                   | Saint Louis                                                 | gPp     |
| sc-23      | Seychelles                                   | Takamaka                                                    | gPp     |
| sc-24      | Seychelles                                   | Les Mamelles                                                | gPp     |
| sc-25      | Seychelles                                   | Roche Caiman                                                | gPp     |
| sc-26      | Seychelles                                   | Ile Perseverance 1                                          | gPp     |
| sc-27      | Seychelles                                   | Ile Perseverance 2                                          | gPp     |
| sd         |                                              | Sudan                                                       | gPp     |
| sd-dc      | Sudan                                        | Central Darfur                                              | gPp     |
| sd-de      | Sudan                                        | East Darfur                                                 | gPp     |
| sd-dn      | Sudan                                        | North Darfur                                                | gPp     |
| sd-ds      | Sudan                                        | South Darfur                                                | gPp     |
| sd-dw      | Sudan                                        | West Darfur                                                 | gPp     |
| sd-gd      | Sudan                                        | Al Qadarif State                                            | gPp     |
| sd-gk      | Sudan                                        | West Kordufan                                               | gPp     |
| sd-gz      | Sudan                                        | Al Jazirah                                                  | gPp     |
| sd-ka      | Sudan                                        | Kassala State                                               | gPp     |
| sd-kh      | Sudan                                        | Al Khartum State                                            | gPp     |
| sd-kn      | Sudan                                        | North Kordufan                                              | gPp     |
| sd-ks      | Sudan                                        | South Kordufan                                              | gPp     |
| sd-nb      | Sudan                                        | Blue Nile                                                   | gPp     |
| sd-no      | Sudan                                        | Northern State                                              | gPp     |
| sd-nr      | Sudan                                        | River Nile State                                            | gPp     |
| sd-nw      | Sudan                                        | White Nile                                                  | gPp     |
| sd-rs      | Sudan                                        | Red Sea State                                               | gPp     |
| sd-si      | Sudan                                        | Sennar State                                                | gPp     |
| se         |                                              | Sweden                                                      | gPp     |
| se-ab      | Sweden                                       | Stockholm County                                            | gPp     |
| se-ac      | Sweden                                       | Västerbotten County                                         | gPp     |
| se-bd      | Sweden                                       | Norrbotten County                                           | gPp     |
| se-c       | Sweden                                       | Uppsala County                                              | gPp     |
| se-d       | Sweden                                       | Södermanland County                                         | gPp     |
| se-e       | Sweden                                       | Östergötland County                                         | gPp     |
| se-f       | Sweden                                       | Jönköping County                                            | gPp     |
| se-g       | Sweden                                       | Kronoberg County                                            | gPp     |
| se-h       | Sweden                                       | Kalmar County                                               | gPp     |
| se-i       | Sweden                                       | Gotland County                                              | gPp     |
| se-k       | Sweden                                       | Blekinge County                                             | gPp     |
| se-m       | Sweden                                       | Skåne County                                                | gPp     |
| se-n       | Sweden                                       | Halland County                                              | gPp     |
| se-o       | Sweden                                       | Västra Götaland County                                      | gPp     |
| se-s       | Sweden                                       | Värmland County                                             | gPp     |
| se-t       | Sweden                                       | Örebro County                                               | gPp     |
| se-u       | Sweden                                       | Västmanland County                                          | gPp     |
| se-w       | Sweden                                       | Dalarna County                                              | gPp     |
| se-x       | Sweden                                       | Gävleborg County                                            | gPp     |
| se-y       | Sweden                                       | Västernorrland County                                       | gPp     |
| se-z       | Sweden                                       | Jämtland County                                             | gPp     |
| sg         |                                              | Singapore                                                   | gPp     |
| sh         |                                              | Saint Helena, Ascension and Tristan da Cunha                | gPp     |
| sh-ac      | Saint Helena, Ascension and Tristan da Cunha | Ascension Island                                            | gPp     |
| sh-hl      | Saint Helena, Ascension and Tristan da Cunha | Saint Helena                                                | gPp     |
| sh-ta      | Saint Helena, Ascension and Tristan da Cunha | Tristan da Cunha                                            | gPp     |
| si         |                                              | Slovenia                                                    | gPp     |
| sk         |                                              | Slovakia                                                    | gPp     |
| sk-bc      | Slovakia                                     | Region of Banská Bystrica                                   | gPp     |
| sk-bl      | Slovakia                                     | Region of Bratislava                                        | gPp     |
| sk-ki      | Slovakia                                     | Region of Košice                                            | gPp     |
| sk-ni      | Slovakia                                     | Region of Nitra                                             | gPp     |
| sk-pv      | Slovakia                                     | Region of Prešov                                            | gPp     |
| sk-ta      | Slovakia                                     | Region of Trnava                                            | gPp     |
| sk-tc      | Slovakia                                     | Region of Trenčín                                           | gPp     |
| sk-zi      | Slovakia                                     | Region of Žilina                                            | gPp     |
| sl         |                                              | Sierra Leone                                                | gPp     |
| sl-e       | Sierra Leone                                 | Eastern Province                                            | gPp     |
| sl-n       | Sierra Leone                                 | Northern Province, Sierra Leone                             | gPp     |
| sl-nw      | Sierra Leone                                 | North West Province, Sierra Leone                           | gPp     |
| sl-s       | Sierra Leone                                 | Southern Province, Sierra Leone                             | gPp     |
| sl-w       | Sierra Leone                                 | Western Area                                                | gPp     |
| sm         |                                              | San Marino                                                  | gPp     |
| sn         |                                              | Senegal                                                     | gPp     |
| sn-db      | Senegal                                      | Diourbel Region                                             | gPp     |
| sn-dk      | Senegal                                      | Dakar Region                                                | gPp     |
| sn-fk      | Senegal                                      | Fatick Region                                               | gPp     |
| sn-ka      | Senegal                                      | Kaffrine Region                                             | gPp     |
| sn-kd      | Senegal                                      | Kolda Region                                                | gPp     |
| sn-ke      | Senegal                                      | Kédougou Region                                             | gPp     |
| sn-kl      | Senegal                                      | Kaolack Region                                              | gPp     |
| sn-lg      | Senegal                                      | Louga Region                                                | gPp     |
| sn-mt      | Senegal                                      | Matam Region                                                | gPp     |
| sn-se      | Senegal                                      | Sédhiou Region                                              | gPp     |
| sn-sl      | Senegal                                      | Saint-Louis Region                                          | gPp     |
| sn-tc      | Senegal                                      | Tambacounda Region                                          | gPp     |
| sn-th      | Senegal                                      | Thiès Region                                                | gPp     |
| sn-zg      | Senegal                                      | Ziguinchor Region                                           | gPp     |
| so         |                                              | Somalia                                                     | gPp     |
| so-aw      | Somalia                                      | Awdal                                                       | gPp     |
| so-bk      | Somalia                                      | Bakool                                                      | gPp     |
| so-bn      | Somalia                                      | Banaadir                                                    | gPp     |
| so-br      | Somalia                                      | Bari                                                        | gPp     |
| so-by      | Somalia                                      | Bay                                                         | gPp     |
| so-ga      | Somalia                                      | Galgaduud                                                   | gPp     |
| so-ge      | Somalia                                      | Gedo                                                        | gPp     |
| so-hi      | Somalia                                      | Hiiraan                                                     | gPp     |
| so-jd      | Somalia                                      | Middle Juba                                                 | gPp     |
| so-jh      | Somalia                                      | Lower Juba                                                  | gPp     |
| so-mu      | Somalia                                      | Mudug                                                       | gPp     |
| so-nu      | Somalia                                      | Nugaal                                                      | gPp     |
| so-sa      | Somalia                                      | Sanaag                                                      | gPp     |
| so-sd      | Somalia                                      | Middle Shebelle                                             | gPp     |
| so-sh      | Somalia                                      | Lower Shabelle                                              | gPp     |
| so-so      | Somalia                                      | Sool                                                        | gPp     |
| so-to      | Somalia                                      | Togdheer                                                    | gPp     |
| so-wo      | Somalia                                      | Woqooyi Galbeed                                             | gPp     |
| so-xx      | Somalia                                      | Sahil                                                       | gPp     |
| sr         |                                              | Suriname                                                    | gPp     |
| sr-br      | Suriname                                     | Brokopondo                                                  | gPp     |
| sr-cm      | Suriname                                     | Commewijne                                                  | gPp     |
| sr-cr      | Suriname                                     | Coronie                                                     | gPp     |
| sr-ma      | Suriname                                     | Marowijne                                                   | gPp     |
| sr-ni      | Suriname                                     | Nickerie                                                    | gPp     |
| sr-pm      | Suriname                                     | Paramaribo                                                  | gPp     |
| sr-pr      | Suriname                                     | Para                                                        | gPp     |
| sr-sa      | Suriname                                     | Saramacca                                                   | gPp     |
| sr-si      | Suriname                                     | Sipaliwini                                                  | gPp     |
| sr-wa      | Suriname                                     | Wanica                                                      | gPp     |
| ss         |                                              | South Sudan                                                 | gPp     |
| st         |                                              | São Tomé and Príncipe                                       | gPp     |
| st-01      | São Tomé and Príncipe                        | Água Grande                                                 | gPp     |
| st-02      | São Tomé and Príncipe                        | Cantagalo                                                   | gPp     |
| st-03      | São Tomé and Príncipe                        | Caué                                                        | gPp     |
| st-04      | São Tomé and Príncipe                        | Lembá                                                       | gPp     |
| st-05      | São Tomé and Príncipe                        | Lobata                                                      | gPp     |
| st-06      | São Tomé and Príncipe                        | Mé-Zóchi                                                    | gPp     |
| st-p       | São Tomé and Príncipe                        | Autonomous Region of Príncipe                               | gPp     |
| sv         |                                              | El Salvador                                                 | gPp     |
| sv-ah      | El Salvador                                  | Ahuachapán                                                  | gPp     |
| sv-ca      | El Salvador                                  | Cabañas                                                     | gPp     |
| sv-ch      | El Salvador                                  | Chalatenango                                                | gPp     |
| sv-cu      | El Salvador                                  | Cuscatlán                                                   | gPp     |
| sv-li      | El Salvador                                  | La Libertad                                                 | gPp     |
| sv-mo      | El Salvador                                  | Morazán                                                     | gPp     |
| sv-pa      | El Salvador                                  | La Paz                                                      | gPp     |
| sv-sa      | El Salvador                                  | Santa Ana                                                   | gPp     |
| sv-sm      | El Salvador                                  | San Miguel                                                  | gPp     |
| sv-so      | El Salvador                                  | Sonsonate                                                   | gPp     |
| sv-ss      | El Salvador                                  | San Salvador                                                | gPp     |
| sv-sv      | El Salvador                                  | San Vicente                                                 | gPp     |
| sv-un      | El Salvador                                  | La Unión                                                    | gPp     |
| sv-us      | El Salvador                                  | Usulután                                                    | gPp     |
| sy         |                                              | Syria                                                       | gPp     |
| sy-di      | Syria                                        | Damascus Governorate                                        | gPp     |
| sy-dr      | Syria                                        | Dar'a Governorate                                           | gPp     |
| sy-dy      | Syria                                        | Deir ez-Zor Governorate                                     | gPp     |
| sy-ha      | Syria                                        | Al-Hasaka Governorate                                       | gPp     |
| sy-hi      | Syria                                        | Homs Governorate                                            | gPp     |
| sy-hl      | Syria                                        | Aleppo Governorate                                          | gPp     |
| sy-hm      | Syria                                        | Hama Governorate                                            | gPp     |
| sy-id      | Syria                                        | Idleb Governorate                                           | gPp     |
| sy-la      | Syria                                        | Latakia Governorate                                         | gPp     |
| sy-qu      | Syria                                        | Al-Quneitra Governorate                                     | gPp     |
| sy-ra      | Syria                                        | Ar-Raqqa Governorate                                        | gPp     |
| sy-rd      | Syria                                        | Rif Dimashq Governorate                                     | gPp     |
| sy-su      | Syria                                        | As-Suweida Governorate                                      | gPp     |
| sy-ta      | Syria                                        | Tartus Governorate                                          | gPp     |
| sz         |                                              | Eswatini                                                    | gPp     |
| sz-hh      | Eswatini                                     | Hhohho Region                                               | gPp     |
| sz-lu      | Eswatini                                     | Lubombo                                                     | gPp     |
| sz-ma      | Eswatini                                     | Manzini Region                                              | gPp     |
| sz-sh      | Eswatini                                     | Shiselweni                                                  | gPp     |
| tc         |                                              | Turks and Caicos Islands                                    | gPp     |
| td         |                                              | Chad                                                        | gPp     |
| td-ba      | Chad                                         | Batha                                                       | gPp     |
| td-bg      | Chad                                         | Bahr el Gazel                                               | gPp     |
| td-bo      | Chad                                         | Borkou                                                      | gPp     |
| td-cb      | Chad                                         | Chari-Baguirmi                                              | gPp     |
| td-ee      | Chad                                         | East Ennedi                                                 | gPp     |
| td-eo      | Chad                                         | Ennedi-Ouest                                                | gPp     |
| td-gr      | Chad                                         | Guéra                                                       | gPp     |
| td-hl      | Chad                                         | Hadjer-Lamis                                                | gPp     |
| td-ka      | Chad                                         | Kanem                                                       | gPp     |
| td-lc      | Chad                                         | Lac                                                         | gPp     |
| td-lo      | Chad                                         | Logone Occidental                                           | gPp     |
| td-lr      | Chad                                         | Logone Oriental                                             | gPp     |
| td-ma      | Chad                                         | Mandoul                                                     | gPp     |
| td-mc      | Chad                                         | Moyen-Chari                                                 | gPp     |
| td-me      | Chad                                         | Mayo-Kebbi Est                                              | gPp     |
| td-mo      | Chad                                         | Mayo-Kebbi Ouest                                            | gPp     |
| td-nd      | Chad                                         | N'Djamena                                                   | gPp     |
| td-od      | Chad                                         | Ouaddaï                                                     | gPp     |
| td-sa      | Chad                                         | Salamat                                                     | gPp     |
| td-si      | Chad                                         | Sila                                                        | gPp     |
| td-ta      | Chad                                         | Tandjilé                                                    | gPp     |
| td-ti      | Chad                                         | Tibesti                                                     | gPp     |
| td-wf      | Chad                                         | Wadi Fira                                                   | gPp     |
| tg         |                                              | Togo                                                        | gPp     |
| tg-c       | Togo                                         | Centrale Region                                             | gPp     |
| tg-k       | Togo                                         | Kara Region                                                 | gPp     |
| tg-m       | Togo                                         | Maritime Region                                             | gPp     |
| tg-p       | Togo                                         | Plateaux Region                                             | gPp     |
| tg-s       | Togo                                         | Savanes Region                                              | gPp     |
| th         |                                              | Thailand                                                    | gPp     |
| th-10      | Thailand                                     | Bangkok                                                     | gPp     |
| th-11      | Thailand                                     | Samut Prakan Province                                       | gPp     |
| th-12      | Thailand                                     | Nonthaburi Province                                         | gPp     |
| th-13      | Thailand                                     | Pathum Thani Province                                       | gPp     |
| th-14      | Thailand                                     | Phra Nakhon Si Ayutthaya Province                           | gPp     |
| th-15      | Thailand                                     | Ang Thong Province                                          | gPp     |
| th-16      | Thailand                                     | Lop Buri Province                                           | gPp     |
| th-17      | Thailand                                     | Sing Buri Province                                          | gPp     |
| th-18      | Thailand                                     | Chai Nat Province                                           | gPp     |
| th-19      | Thailand                                     | Saraburi Province                                           | gPp     |
| th-20      | Thailand                                     | Chon Buri Province                                          | gPp     |
| th-21      | Thailand                                     | Rayong Province                                             | gPp     |
| th-22      | Thailand                                     | Chanthaburi Province                                        | gPp     |
| th-23      | Thailand                                     | Trat Province                                               | gPp     |
| th-24      | Thailand                                     | Chachoengsao Province                                       | gPp     |
| th-25      | Thailand                                     | Prachin Buri Province                                       | gPp     |
| th-26      | Thailand                                     | Nakhon Nayok Province                                       | gPp     |
| th-27      | Thailand                                     | Sa Kaeo Province                                            | gPp     |
| th-30      | Thailand                                     | Nakhon Ratchasima Province                                  | gPp     |
| th-31      | Thailand                                     | Buri Ram Province                                           | gPp     |
| th-32      | Thailand                                     | Surin Province                                              | gPp     |
| th-33      | Thailand                                     | Si Sa Ket Province                                          | gPp     |
| th-34      | Thailand                                     | Ubon Ratchathani Province                                   | gPp     |
| th-35      | Thailand                                     | Yasothon Province                                           | gPp     |
| th-36      | Thailand                                     | Chaiyaphum Province                                         | gPp     |
| th-37      | Thailand                                     | Amnat Charoen Province                                      | gPp     |
| th-38      | Thailand                                     | Bueng Kan Province                                          | gPp     |
| th-39      | Thailand                                     | Nong Bua Lam Phu Province                                   | gPp     |
| th-40      | Thailand                                     | Khon Kaen Province                                          | gPp     |
| th-41      | Thailand                                     | Udon Thani Province                                         | gPp     |
| th-42      | Thailand                                     | Loei Province                                               | gPp     |
| th-43      | Thailand                                     | Nong Khai Province                                          | gPp     |
| th-44      | Thailand                                     | Maha Sarakham Province                                      | gPp     |
| th-45      | Thailand                                     | Roi Et Province                                             | gPp     |
| th-46      | Thailand                                     | Kalasin Province                                            | gPp     |
| th-47      | Thailand                                     | Sakon Nakhon Province                                       | gPp     |
| th-48      | Thailand                                     | Nakhon Phanom Province                                      | gPp     |
| th-49      | Thailand                                     | Mukdahan Province                                           | gPp     |
| th-50      | Thailand                                     | Chiang Mai Province                                         | gPp     |
| th-51      | Thailand                                     | Lamphun Province                                            | gPp     |
| th-52      | Thailand                                     | Lampang Province                                            | gPp     |
| th-53      | Thailand                                     | Uttaradit Province                                          | gPp     |
| th-54      | Thailand                                     | Phrae Province                                              | gPp     |
| th-55      | Thailand                                     | Nan Province                                                | gPp     |
| th-56      | Thailand                                     | Phayao Province                                             | gPp     |
| th-57      | Thailand                                     | Chiang Rai Province                                         | gPp     |
| th-58      | Thailand                                     | Mae Hong Son Province                                       | gPp     |
| th-60      | Thailand                                     | Nakhon Sawan Province                                       | gPp     |
| th-61      | Thailand                                     | Uthai Thani Province                                        | gPp     |
| th-62      | Thailand                                     | Kamphaeng Phet Province                                     | gPp     |
| th-63      | Thailand                                     | Tak Province                                                | gPp     |
| th-64      | Thailand                                     | Sukhothai Province                                          | gPp     |
| th-65      | Thailand                                     | Phitsanulok Province                                        | gPp     |
| th-66      | Thailand                                     | Phichit Province                                            | gPp     |
| th-67      | Thailand                                     | Phetchabun Province                                         | gPp     |
| th-70      | Thailand                                     | Ratchaburi Province                                         | gPp     |
| th-71      | Thailand                                     | Kanchanaburi Province                                       | gPp     |
| th-72      | Thailand                                     | Suphan Buri Province                                        | gPp     |
| th-73      | Thailand                                     | Nakhon Pathom Province                                      | gPp     |
| th-74      | Thailand                                     | Samut Sakhon Province                                       | gPp     |
| th-75      | Thailand                                     | Samut Songkhram Province                                    | gPp     |
| th-76      | Thailand                                     | Phetchaburi Province                                        | gPp     |
| th-77      | Thailand                                     | Prachuap Khiri Khan Province                                | gPp     |
| th-80      | Thailand                                     | Nakhon Si Thammarat Province                                | gPp     |
| th-81      | Thailand                                     | Krabi Province                                              | gPp     |
| th-82      | Thailand                                     | Phang-nga Province                                          | gPp     |
| th-83      | Thailand                                     | Phuket Province                                             | gPp     |
| th-84      | Thailand                                     | Surat Thani Province                                        | gPp     |
| th-85      | Thailand                                     | Ranong Province                                             | gPp     |
| th-86      | Thailand                                     | Chumphon Province                                           | gPp     |
| th-90      | Thailand                                     | Songkhla Province                                           | gPp     |
| th-91      | Thailand                                     | Satun Province                                              | gPp     |
| th-92      | Thailand                                     | Trang Province                                              | gPp     |
| th-93      | Thailand                                     | Phatthalung Province                                        | gPp     |
| th-94      | Thailand                                     | Pattani Province                                            | gPp     |
| th-95      | Thailand                                     | Yala Province                                               | gPp     |
| th-96      | Thailand                                     | Narathiwat Province                                         | gPp     |
| tj         |                                              | Tajikistan                                                  | gPp     |
| tj-du      | Tajikistan                                   | Dushanbe                                                    | gPp     |
| tj-gb      | Tajikistan                                   | Gorno-Badakhshan Autonomous Region                          | gPp     |
| tj-kt      | Tajikistan                                   | Khatlon Region                                              | gPp     |
| tj-ra      | Tajikistan                                   | Districts of Republican Subordination                       | gPp     |
| tj-su      | Tajikistan                                   | Sughd Region                                                | gPp     |
| tk         |                                              | Tokelau                                                     | gPp     |
| tl         |                                              | East Timor                                                  | gPp     |
| tl-al      | East Timor                                   | Aileu                                                       | gPp     |
| tl-an      | East Timor                                   | Ainaro                                                      | gPp     |
| tl-ba      | East Timor                                   | Baucau                                                      | gPp     |
| tl-bo      | East Timor                                   | Bobonaro                                                    | gPp     |
| tl-co      | East Timor                                   | Cova Lima                                                   | gPp     |
| tl-di      | East Timor                                   | Dili                                                        | gPp     |
| tl-er      | East Timor                                   | Ermera                                                      | gPp     |
| tl-la      | East Timor                                   | Lautém                                                      | gPp     |
| tl-li      | East Timor                                   | Liquiçá                                                     | gPp     |
| tl-mf      | East Timor                                   | Manufahi                                                    | gPp     |
| tl-mt      | East Timor                                   | Manatuto                                                    | gPp     |
| tl-oe      | East Timor                                   | Oecussi-Ambeno                                              | gPp     |
| tl-vi      | East Timor                                   | Viqueque                                                    | gPp     |
| tm         |                                              | Turkmenistan                                                | gPp     |
| tm-a       | Turkmenistan                                 | Ahal Region                                                 | gPp     |
| tm-b       | Turkmenistan                                 | Balkan Region                                               | gPp     |
| tm-d       | Turkmenistan                                 | Dashoguz Region                                             | gPp     |
| tm-l       | Turkmenistan                                 | Lebap Region                                                | gPp     |
| tm-m       | Turkmenistan                                 | Mary Region                                                 | gPp     |
| tm-s       | Turkmenistan                                 | Ashgabat                                                    | gPp     |
| tn         |                                              | Tunisia                                                     | gPp     |
| tn-11      | Tunisia                                      | Tunis                                                       | gPp     |
| tn-12      | Tunisia                                      | Ariana                                                      | gPp     |
| tn-13      | Tunisia                                      | Ben Arous                                                   | gPp     |
| tn-14      | Tunisia                                      | Manouba                                                     | gPp     |
| tn-21      | Tunisia                                      | Nabeul                                                      | gPp     |
| tn-22      | Tunisia                                      | Zaghouan Governorate                                        | gPp     |
| tn-23      | Tunisia                                      | Bizerte                                                     | gPp     |
| tn-31      | Tunisia                                      | Béja                                                        | gPp     |
| tn-32      | Tunisia                                      | Jendouba                                                    | gPp     |
| tn-33      | Tunisia                                      | Al Kaf                                                      | gPp     |
| tn-34      | Tunisia                                      | Siliana                                                     | gPp     |
| tn-41      | Tunisia                                      | Kairouan                                                    | gPp     |
| tn-42      | Tunisia                                      | Kasserine                                                   | gPp     |
| tn-43      | Tunisia                                      | Sidi Bouzid                                                 | gPp     |
| tn-51      | Tunisia                                      | Sousse                                                      | gPp     |
| tn-52      | Tunisia                                      | Monastir                                                    | gPp     |
| tn-53      | Tunisia                                      | Mahdia                                                      | gPp     |
| tn-61      | Tunisia                                      | Sfax                                                        | gPp     |
| tn-71      | Tunisia                                      | Gafsa                                                       | gPp     |
| tn-72      | Tunisia                                      | Tozeur                                                      | gPp     |
| tn-73      | Tunisia                                      | Kébili                                                      | gPp     |
| tn-81      | Tunisia                                      | Gabès                                                       | gPp     |
| tn-82      | Tunisia                                      | Médenine                                                    | gPp     |
| tn-83      | Tunisia                                      | Tataouine                                                   | gPp     |
| to         |                                              | Tonga                                                       | gPp     |
| to-01      | Tonga                                        | ʻEua                                                        | gPp     |
| to-02      | Tonga                                        | Haʻapai                                                     | gPp     |
| to-03      | Tonga                                        | Ongo Niua                                                   | gPp     |
| to-04      | Tonga                                        | Tongatapu                                                   | gPp     |
| to-05      | Tonga                                        | Vavaʻu                                                      | gPp     |
| tr         |                                              | Turkey                                                      | gPp     |
| tr-01      | Turkey                                       | Adana                                                       | gPp     |
| tr-02      | Turkey                                       | Adıyaman                                                    | gPp     |
| tr-03      | Turkey                                       | Afyonkarahisar                                              | gPp     |
| tr-04      | Turkey                                       | Ağrı                                                        | gPp     |
| tr-05      | Turkey                                       | Amasya                                                      | gPp     |
| tr-06      | Turkey                                       | Ankara                                                      | gPp     |
| tr-07      | Turkey                                       | Antalya                                                     | gPp     |
| tr-08      | Turkey                                       | Artvin                                                      | gPp     |
| tr-09      | Turkey                                       | Aydın                                                       | gPp     |
| tr-10      | Turkey                                       | Balıkesir                                                   | gPp     |
| tr-11      | Turkey                                       | Bilecik                                                     | gPp     |
| tr-12      | Turkey                                       | Bingöl                                                      | gPp     |
| tr-13      | Turkey                                       | Bitlis                                                      | gPp     |
| tr-14      | Turkey                                       | Bolu                                                        | gPp     |
| tr-15      | Turkey                                       | Burdur                                                      | gPp     |
| tr-16      | Turkey                                       | Bursa                                                       | gPp     |
| tr-17      | Turkey                                       | Canakkale                                                   | gPp     |
| tr-18      | Turkey                                       | Çankırı                                                     | gPp     |
| tr-19      | Turkey                                       | Çorum                                                       | gPp     |
| tr-20      | Turkey                                       | Denizli                                                     | gPp     |
| tr-21      | Turkey                                       | Diyarbakır                                                  | gPp     |
| tr-22      | Turkey                                       | Edirne                                                      | gPp     |
| tr-23      | Turkey                                       | Elazığ                                                      | gPp     |
| tr-24      | Turkey                                       | Erzincan                                                    | gPp     |
| tr-25      | Turkey                                       | Erzurum                                                     | gPp     |
| tr-26      | Turkey                                       | Eskişehir                                                   | gPp     |
| tr-27      | Turkey                                       | Gaziantep                                                   | gPp     |
| tr-28      | Turkey                                       | Giresun                                                     | gPp     |
| tr-29      | Turkey                                       | Gümüşhane                                                   | gPp     |
| tr-30      | Turkey                                       | Hakkâri                                                     | gPp     |
| tr-31      | Turkey                                       | Hatay                                                       | gPp     |
| tr-32      | Turkey                                       | Isparta                                                     | gPp     |
| tr-33      | Turkey                                       | Mersin                                                      | gPp     |
| tr-34      | Turkey                                       | Istanbul                                                    | gPp     |
| tr-35      | Turkey                                       | Izmir                                                       | gPp     |
| tr-36      | Turkey                                       | Kars                                                        | gPp     |
| tr-37      | Turkey                                       | Kastamonu                                                   | gPp     |
| tr-38      | Turkey                                       | Kayseri                                                     | gPp     |
| tr-39      | Turkey                                       | Kırklareli                                                  | gPp     |
| tr-40      | Turkey                                       | Kırşehir                                                    | gPp     |
| tr-41      | Turkey                                       | Kocaeli                                                     | gPp     |
| tr-42      | Turkey                                       | Konya                                                       | gPp     |
| tr-43      | Turkey                                       | Kütahya                                                     | gPp     |
| tr-44      | Turkey                                       | Malatya                                                     | gPp     |
| tr-45      | Turkey                                       | Manisa                                                      | gPp     |
| tr-46      | Turkey                                       | Kahramanmaraş                                               | gPp     |
| tr-47      | Turkey                                       | Mardin                                                      | gPp     |
| tr-48      | Turkey                                       | Muğla                                                       | gPp     |
| tr-49      | Turkey                                       | Muş                                                         | gPp     |
| tr-50      | Turkey                                       | Nevşehir                                                    | gPp     |
| tr-51      | Turkey                                       | Niğde                                                       | gPp     |
| tr-52      | Turkey                                       | Ordu                                                        | gPp     |
| tr-53      | Turkey                                       | Rize                                                        | gPp     |
| tr-54      | Turkey                                       | Sakarya                                                     | gPp     |
| tr-55      | Turkey                                       | Samsun                                                      | gPp     |
| tr-56      | Turkey                                       | Siirt                                                       | gPp     |
| tr-57      | Turkey                                       | Sinop                                                       | gPp     |
| tr-58      | Turkey                                       | Sivas                                                       | gPp     |
| tr-59      | Turkey                                       | Tekirdağ                                                    | gPp     |
| tr-60      | Turkey                                       | Tokat                                                       | gPp     |
| tr-61      | Turkey                                       | Trabzon                                                     | gPp     |
| tr-62      | Turkey                                       | Tunceli                                                     | gPp     |
| tr-63      | Turkey                                       | Şanlıurfa                                                   | gPp     |
| tr-64      | Turkey                                       | Uşak                                                        | gPp     |
| tr-65      | Turkey                                       | Van                                                         | gPp     |
| tr-66      | Turkey                                       | Yozgat                                                      | gPp     |
| tr-67      | Turkey                                       | Zonguldak                                                   | gPp     |
| tr-68      | Turkey                                       | Aksaray                                                     | gPp     |
| tr-69      | Turkey                                       | Bayburt                                                     | gPp     |
| tr-70      | Turkey                                       | Karaman                                                     | gPp     |
| tr-71      | Turkey                                       | Kırıkkale                                                   | gPp     |
| tr-72      | Turkey                                       | Batman                                                      | gPp     |
| tr-73      | Turkey                                       | Şırnak                                                      | gPp     |
| tr-74      | Turkey                                       | Bartın                                                      | gPp     |
| tr-75      | Turkey                                       | Ardahan                                                     | gPp     |
| tr-76      | Turkey                                       | Iğdır                                                       | gPp     |
| tr-77      | Turkey                                       | Yalova                                                      | gPp     |
| tr-78      | Turkey                                       | Karabük                                                     | gPp     |
| tr-79      | Turkey                                       | Kilis                                                       | gPp     |
| tr-80      | Turkey                                       | Osmaniye                                                    | gPp     |
| tr-81      | Turkey                                       | Düzce                                                       | gPp     |
| tt         |                                              | Trinidad and Tobago                                         | gPp     |
| tt-ari     | Trinidad and Tobago                          | Arima                                                       | gPp     |
| tt-cha     | Trinidad and Tobago                          | Chaguanas                                                   | gPp     |
| tt-ctt     | Trinidad and Tobago                          | Couva-Tabaquite-Talparo                                     | gPp     |
| tt-dmn     | Trinidad and Tobago                          | Diego Martin                                                | gPp     |
| tt-mrc     | Trinidad and Tobago                          | Mayaro-Rio Claro                                            | gPp     |
| tt-ped     | Trinidad and Tobago                          | Penal-Debe                                                  | gPp     |
| tt-pos     | Trinidad and Tobago                          | Port of Spain                                               | gPp     |
| tt-prt     | Trinidad and Tobago                          | Princes Town                                                | gPp     |
| tt-ptf     | Trinidad and Tobago                          | Point Fortin                                                | gPp     |
| tt-sfo     | Trinidad and Tobago                          | San Fernando                                                | gPp     |
| tt-sge     | Trinidad and Tobago                          | Sangre Grande                                               | gPp     |
| tt-sip     | Trinidad and Tobago                          | Siparia                                                     | gPp     |
| tt-sjl     | Trinidad and Tobago                          | San Juan-Laventille                                         | gPp     |
| tt-tob     | Trinidad and Tobago                          | Tobago                                                      | gPp     |
| tt-tup     | Trinidad and Tobago                          | Tunapuna-Piarco                                             | gPp     |
| tv         |                                              | Tuvalu                                                      | gPp     |
| tw         |                                              | Taiwan                                                      | gPp     |
| tw-cha     | Taiwan                                       | Changhua County                                             | gPp     |
| tw-cyi     | Taiwan                                       | Chiayi                                                      | gPp     |
| tw-cyq     | Taiwan                                       | Chiayi County                                               | gPp     |
| tw-hsq     | Taiwan                                       | Hsinchu County                                              | gPp     |
| tw-hsz     | Taiwan                                       | Hsinchu                                                     | gPp     |
| tw-hua     | Taiwan                                       | Hualien County                                              | gPp     |
| tw-ila     | Taiwan                                       | Yilan County                                                | gPp     |
| tw-kee     | Taiwan                                       | Keelung                                                     | gPp     |
| tw-khh     | Taiwan                                       | Kaohsiung                                                   | gPp     |
| tw-kin     | Taiwan                                       | Kinmen                                                      | gPp     |
| tw-lie     | Taiwan                                       | Lienchiang County                                           | gPp     |
| tw-mia     | Taiwan                                       | Miaoli County                                               | gPp     |
| tw-nan     | Taiwan                                       | Nantou County                                               | gPp     |
| tw-nwt     | Taiwan                                       | New Taipei                                                  | gPp     |
| tw-pen     | Taiwan                                       | Penghu                                                      | gPp     |
| tw-pif     | Taiwan                                       | Pingtung County                                             | gPp     |
| tw-tao     | Taiwan                                       | Taoyuan City                                                | gPp     |
| tw-tnn     | Taiwan                                       | Tainan                                                      | gPp     |
| tw-tpe     | Taiwan                                       | Taipei                                                      | gPp     |
| tw-ttt     | Taiwan                                       | Taitung County                                              | gPp     |
| tw-txg     | Taiwan                                       | Taichung                                                    | gPp     |
| tw-yun     | Taiwan                                       | Yunlin County                                               | gPp     |
| tz         |                                              | Tanzania                                                    | gPp     |
| tz-01      | Tanzania                                     | Arusha                                                      | gPp     |
| tz-02      | Tanzania                                     | Dar es Salaam                                               | gPp     |
| tz-03      | Tanzania                                     | Dodoma Region                                               | gPp     |
| tz-04      | Tanzania                                     | Iringa Region                                               | gPp     |
| tz-05      | Tanzania                                     | Kagera                                                      | gPp     |
| tz-06      | Tanzania                                     | North Pemba                                                 | gPp     |
| tz-07      | Tanzania                                     | Zanzibar North                                              | gPp     |
| tz-08      | Tanzania                                     | Kigoma Region                                               | gPp     |
| tz-09      | Tanzania                                     | Kilimanjaro                                                 | gPp     |
| tz-10      | Tanzania                                     | South Pemba                                                 | gPp     |
| tz-11      | Tanzania                                     | Zanzibar South & Central                                    | gPp     |
| tz-12      | Tanzania                                     | Lindi Region                                                | gPp     |
| tz-13      | Tanzania                                     | Mara Region                                                 | gPp     |
| tz-14      | Tanzania                                     | Mbeya Region                                                | gPp     |
| tz-15      | Tanzania                                     | Zanzibar Urban/West                                         | gPp     |
| tz-16      | Tanzania                                     | Morogoro Region                                             | gPp     |
| tz-17      | Tanzania                                     | Mtwara Region                                               | gPp     |
| tz-18      | Tanzania                                     | Mwanza Region                                               | gPp     |
| tz-19      | Tanzania                                     | Pwani Region                                                | gPp     |
| tz-20      | Tanzania                                     | Rukwa Region                                                | gPp     |
| tz-21      | Tanzania                                     | Ruvuma Region                                               | gPp     |
| tz-22      | Tanzania                                     | Shinyanga Region                                            | gPp     |
| tz-23      | Tanzania                                     | Singida Region                                              | gPp     |
| tz-24      | Tanzania                                     | Tabora Region                                               | gPp     |
| tz-25      | Tanzania                                     | Tanga Region                                                | gPp     |
| tz-26      | Tanzania                                     | Manyara Region                                              | gPp     |
| tz-27      | Tanzania                                     | Geita                                                       | gPp     |
| tz-28      | Tanzania                                     | Katavi Region                                               | gPp     |
| tz-29      | Tanzania                                     | Njombe Region                                               | gPp     |
| tz-30      | Tanzania                                     | Simiyu                                                      | gPp     |
| tz-31      | Tanzania                                     | Songwe Region                                               | gPp     |
| ua         |                                              | Ukraine                                                     | gPp     |
| ua-05      | Ukraine                                      | Vinnytsia Oblast                                            | gPp     |
| ua-07      | Ukraine                                      | Volyn Oblast                                                | gPp     |
| ua-09      | Ukraine                                      | Luhansk Oblast                                              | gPp     |
| ua-12      | Ukraine                                      | Dnipropetrovsk Oblast                                       | gPp     |
| ua-14      | Ukraine                                      | Donetsk Oblast                                              | gPp     |
| ua-18      | Ukraine                                      | Zhytomyr Oblast                                             | gPp     |
| ua-21      | Ukraine                                      | Zakarpattia Oblast                                          | gPp     |
| ua-23      | Ukraine                                      | Zaporizhia Oblast                                           | gPp     |
| ua-26      | Ukraine                                      | Ivano-Frankivsk Oblast                                      | gPp     |
| ua-30      | Ukraine                                      | Kyiv                                                        | gPp     |
| ua-32      | Ukraine                                      | Kyiv Oblast                                                 | gPp     |
| ua-35      | Ukraine                                      | Kirovohrad Oblast                                           | gPp     |
| ua-40      | Ukraine                                      | Sevastopol                                                  | gPp     |
| ua-43      | Ukraine                                      | Autonomous Republic of Crimea                               | gPp     |
| ua-46      | Ukraine                                      | Lviv Oblast                                                 | gPp     |
| ua-48      | Ukraine                                      | Mykolaiv Oblast                                             | gPp     |
| ua-51      | Ukraine                                      | Odesa Oblast                                                | gPp     |
| ua-53      | Ukraine                                      | Poltava Oblast                                              | gPp     |
| ua-56      | Ukraine                                      | Rivne Oblast                                                | gPp     |
| ua-59      | Ukraine                                      | Sumy Oblast                                                 | gPp     |
| ua-61      | Ukraine                                      | Ternopil Oblast                                             | gPp     |
| ua-63      | Ukraine                                      | Kharkiv Oblast                                              | gPp     |
| ua-65      | Ukraine                                      | Kherson Oblast                                              | gPp     |
| ua-68      | Ukraine                                      | Khmelnytskyi Oblast                                         | gPp     |
| ua-71      | Ukraine                                      | Cherkasy Oblast                                             | gPp     |
| ua-74      | Ukraine                                      | Chernihiv Oblast                                            | gPp     |
| ua-77      | Ukraine                                      | Chernivtsi Oblast                                           | gPp     |
| ug         |                                              | Uganda                                                      | gPp     |
| ug-c       | Uganda                                       | Central Region                                              | gPp     |
| ug-e       | Uganda                                       | Eastern Region                                              | gPp     |
| ug-n       | Uganda                                       | Northern Region                                             | gPp     |
| ug-w       | Uganda                                       | Western Region                                              | gPp     |
| us         |                                              | United States                                               | gPp     |
| us-ak      | United States                                | Alaska                                                      | gPp     |
| us-al      | United States                                | Alabama                                                     | gPp     |
| us-ar      | United States                                | Arkansas                                                    | gPp     |
| us-as      | United States                                | American Samoa                                              | gPp     |
| us-az      | United States                                | Arizona                                                     | gPp     |
| us-ca      | United States                                | California                                                  | gPp     |
| us-co      | United States                                | Colorado                                                    | gPp     |
| us-ct      | United States                                | Connecticut                                                 | gPp     |
| us-dc      | United States                                | District of Columbia                                        | gPp     |
| us-de      | United States                                | Delaware                                                    | gPp     |
| us-fl      | United States                                | Florida                                                     | gPp     |
| us-ga      | United States                                | Georgia                                                     | gPp     |
| us-gu      | United States                                | Guam                                                        | gPp     |
| us-hi      | United States                                | Hawaii                                                      | gPp     |
| us-ia      | United States                                | Iowa                                                        | gPp     |
| us-id      | United States                                | Idaho                                                       | gPp     |
| us-il      | United States                                | Illinois                                                    | gPp     |
| us-in      | United States                                | Indiana                                                     | gPp     |
| us-ks      | United States                                | Kansas                                                      | gPp     |
| us-ky      | United States                                | Kentucky                                                    | gPp     |
| us-la      | United States                                | Louisiana                                                   | gPp     |
| us-ma      | United States                                | Massachusetts                                               | gPp     |
| us-md      | United States                                | Maryland                                                    | gPp     |
| us-me      | United States                                | Maine                                                       | gPp     |
| us-mi      | United States                                | Michigan                                                    | gPp     |
| us-mn      | United States                                | Minnesota                                                   | gPp     |
| us-mo      | United States                                | Missouri                                                    | gPp     |
| us-mp      | United States                                | Northern Mariana Islands                                    | gPp     |
| us-ms      | United States                                | Mississippi                                                 | gPp     |
| us-mt      | United States                                | Montana                                                     | gPp     |
| us-nc      | United States                                | North Carolina                                              | gPp     |
| us-nd      | United States                                | North Dakota                                                | gPp     |
| us-ne      | United States                                | Nebraska                                                    | gPp     |
| us-nh      | United States                                | New Hampshire                                               | gPp     |
| us-nj      | United States                                | New Jersey                                                  | gPp     |
| us-nm      | United States                                | New Mexico                                                  | gPp     |
| us-nv      | United States                                | Nevada                                                      | gPp     |
| us-ny      | United States                                | New York                                                    | gPp     |
| us-oh      | United States                                | Ohio                                                        | gPp     |
| us-ok      | United States                                | Oklahoma                                                    | gPp     |
| us-or      | United States                                | Oregon                                                      | gPp     |
| us-pa      | United States                                | Pennsylvania                                                | gPp     |
| us-pr      | United States                                | Puerto Rico                                                 | gPp     |
| us-ri      | United States                                | Rhode Island                                                | gPp     |
| us-sc      | United States                                | South Carolina                                              | gPp     |
| us-sd      | United States                                | South Dakota                                                | gPp     |
| us-tn      | United States                                | Tennessee                                                   | gPp     |
| us-tx      | United States                                | Texas                                                       | gPp     |
| us-ut      | United States                                | Utah                                                        | gPp     |
| us-va      | United States                                | Virginia                                                    | gPp     |
| us-vi      | United States                                | United States Virgin Islands                                | gPp     |
| us-vt      | United States                                | Vermont                                                     | gPp     |
| us-wa      | United States                                | Washington                                                  | gPp     |
| us-wi      | United States                                | Wisconsin                                                   | gPp     |
| us-wv      | United States                                | West Virginia                                               | gPp     |
| us-wy      | United States                                | Wyoming                                                     | gPp     |
| uy         |                                              | Uruguay                                                     | gPp     |
| uy-ar      | Uruguay                                      | Artigas                                                     | gPp     |
| uy-ca      | Uruguay                                      | Canelones                                                   | gPp     |
| uy-cl      | Uruguay                                      | Cerro Largo                                                 | gPp     |
| uy-co      | Uruguay                                      | Colonia                                                     | gPp     |
| uy-du      | Uruguay                                      | Durazno                                                     | gPp     |
| uy-fd      | Uruguay                                      | Florida                                                     | gPp     |
| uy-fs      | Uruguay                                      | Flores                                                      | gPp     |
| uy-la      | Uruguay                                      | Lavalleja                                                   | gPp     |
| uy-ma      | Uruguay                                      | Maldonado                                                   | gPp     |
| uy-mo      | Uruguay                                      | Montevideo                                                  | gPp     |
| uy-pa      | Uruguay                                      | Paysandú                                                    | gPp     |
| uy-rn      | Uruguay                                      | Río Negro                                                   | gPp     |
| uy-ro      | Uruguay                                      | Rocha                                                       | gPp     |
| uy-rv      | Uruguay                                      | Rivera                                                      | gPp     |
| uy-sa      | Uruguay                                      | Salto                                                       | gPp     |
| uy-sj      | Uruguay                                      | San José                                                    | gPp     |
| uy-so      | Uruguay                                      | Soriano                                                     | gPp     |
| uy-ta      | Uruguay                                      | Tacuarembó                                                  | gPp     |
| uy-tt      | Uruguay                                      | Treinta y Tres                                              | gPp     |
| uz         |                                              | Uzbekistan                                                  | gPp     |
| uz-an      | Uzbekistan                                   | Andijan Region                                              | gPp     |
| uz-bu      | Uzbekistan                                   | Bukhara Region                                              | gPp     |
| uz-fa      | Uzbekistan                                   | Fergana Region                                              | gPp     |
| uz-ji      | Uzbekistan                                   | Jizzakh Region                                              | gPp     |
| uz-ng      | Uzbekistan                                   | Namangan Region                                             | gPp     |
| uz-nw      | Uzbekistan                                   | Navoiy Region                                               | gPp     |
| uz-qa      | Uzbekistan                                   | Qashqadaryo Region                                          | gPp     |
| uz-qr      | Uzbekistan                                   | Republic of Karakalpakstan                                  | gPp     |
| uz-sa      | Uzbekistan                                   | Samarqand Region                                            | gPp     |
| uz-si      | Uzbekistan                                   | Sirdaryo Region                                             | gPp     |
| uz-su      | Uzbekistan                                   | Surxondaryo Region                                          | gPp     |
| uz-tk      | Uzbekistan                                   | Tashkent                                                    | gPp     |
| uz-to      | Uzbekistan                                   | Tashkent Region                                             | gPp     |
| uz-xo      | Uzbekistan                                   | Xorazm Region                                               | gPp     |
| va         |                                              | Vatican City                                                | gPp     |
| vc         |                                              | Saint Vincent and the Grenadines                            | gPp     |
| vc-01      | Saint Vincent and the Grenadines             | Charlotte                                                   | gPp     |
| vc-02      | Saint Vincent and the Grenadines             | Saint Andrew                                                | gPp     |
| vc-03      | Saint Vincent and the Grenadines             | Saint David                                                 | gPp     |
| vc-04      | Saint Vincent and the Grenadines             | Saint George                                                | gPp     |
| vc-05      | Saint Vincent and the Grenadines             | Saint Patrick                                               | gPp     |
| vc-06      | Saint Vincent and the Grenadines             | Grenadines                                                  | gPp     |
| ve         |                                              | Venezuela                                                   | gPp     |
| ve-a       | Venezuela                                    | Capital District                                            | gPp     |
| ve-b       | Venezuela                                    | Anzoategui State                                            | gPp     |
| ve-c       | Venezuela                                    | Apure State                                                 | gPp     |
| ve-d       | Venezuela                                    | Aragua State                                                | gPp     |
| ve-e       | Venezuela                                    | Barinas State                                               | gPp     |
| ve-f       | Venezuela                                    | Bolivar State                                               | gPp     |
| ve-g       | Venezuela                                    | Carabobo State                                              | gPp     |
| ve-h       | Venezuela                                    | Cojedes State                                               | gPp     |
| ve-i       | Venezuela                                    | Falcon State                                                | gPp     |
| ve-j       | Venezuela                                    | Guarico State                                               | gPp     |
| ve-k       | Venezuela                                    | Lara State                                                  | gPp     |
| ve-l       | Venezuela                                    | Merida State                                                | gPp     |
| ve-m       | Venezuela                                    | Miranda State                                               | gPp     |
| ve-n       | Venezuela                                    | Monagas State                                               | gPp     |
| ve-o       | Venezuela                                    | Nueva Esparta State                                         | gPp     |
| ve-p       | Venezuela                                    | Portuguesa State                                            | gPp     |
| ve-r       | Venezuela                                    | Sucre State                                                 | gPp     |
| ve-s       | Venezuela                                    | Tachira State                                               | gPp     |
| ve-t       | Venezuela                                    | Trujillo State                                              | gPp     |
| ve-u       | Venezuela                                    | Yaracuy State                                               | gPp     |
| ve-v       | Venezuela                                    | Zulia State                                                 | gPp     |
| ve-w       | Venezuela                                    | Federal Dependencies                                        | gPp     |
| ve-x       | Venezuela                                    | Vargas State                                                | gPp     |
| ve-y       | Venezuela                                    | Delta Amacuro State                                         | gPp     |
| ve-z       | Venezuela                                    | Amazonas State                                              | gPp     |
| vg         |                                              | British Virgin Islands                                      | gPp     |
| vn         |                                              | Vietnam                                                     | gPp     |
| vn-01      | Vietnam                                      | Lai Châu Province                                           | gPp     |
| vn-02      | Vietnam                                      | Lào Cai Province                                            | gPp     |
| vn-04      | Vietnam                                      | Cao Bằng Province                                           | gPp     |
| vn-05      | Vietnam                                      | Sơn La Province                                             | gPp     |
| vn-07      | Vietnam                                      | Tuyên Quang Province                                        | gPp     |
| vn-09      | Vietnam                                      | Lạng Sơn Province                                           | gPp     |
| vn-13      | Vietnam                                      | Quang Ninh City                                             | gPp     |
| vn-18      | Vietnam                                      | Ninh Bình Province                                          | gPp     |
| vn-21      | Vietnam                                      | Thanh Hóa Province                                          | gPp     |
| vn-22      | Vietnam                                      | Nghệ An Province                                            | gPp     |
| vn-23      | Vietnam                                      | Hà Tĩnh Province                                            | gPp     |
| vn-25      | Vietnam                                      | Quảng Trị Province                                          | gPp     |
| vn-26      | Vietnam                                      | Huế                                                         | gPp     |
| vn-29      | Vietnam                                      | Quảng Ngãi Province                                         | gPp     |
| vn-30      | Vietnam                                      | Gia Lai Province                                            | gPp     |
| vn-33      | Vietnam                                      | Đắk Lắk Province                                            | gPp     |
| vn-34      | Vietnam                                      | Khánh Hòa Province                                          | gPp     |
| vn-35      | Vietnam                                      | Lâm Đồng Province                                           | gPp     |
| vn-37      | Vietnam                                      | Tây Ninh Province                                           | gPp     |
| vn-39      | Vietnam                                      | Đồng Nai                                                    | gPp     |
| vn-44      | Vietnam                                      | An Giang Province                                           | gPp     |
| vn-45      | Vietnam                                      | Đồng Tháp Province                                          | gPp     |
| vn-49      | Vietnam                                      | Vĩnh Long Province                                          | gPp     |
| vn-56      | Vietnam                                      | Bắc Ninh City                                               | gPp     |
| vn-59      | Vietnam                                      | Cà Mau Province                                             | gPp     |
| vn-66      | Vietnam                                      | Hưng Yên Province                                           | gPp     |
| vn-68      | Vietnam                                      | Phú Thọ Province                                            | gPp     |
| vn-69      | Vietnam                                      | Thái Nguyên Province                                        | gPp     |
| vn-71      | Vietnam                                      | Điện Biên Province                                          | gPp     |
| vn-ct      | Vietnam                                      | Cần Thơ                                                     | gPp     |
| vn-dn      | Vietnam                                      | Đà Nẵng                                                     | gPp     |
| vn-hn      | Vietnam                                      | Hà Nội                                                      | gPp     |
| vn-hp      | Vietnam                                      | Hải Phòng                                                   | gPp     |
| vn-sg      | Vietnam                                      | Ho Chi Minh City                                            | gPp     |
| vu         |                                              | Vanuatu                                                     | gPp     |
| vu-map     | Vanuatu                                      | Malampa                                                     | gPp     |
| vu-pam     | Vanuatu                                      | Penama                                                      | gPp     |
| vu-sam     | Vanuatu                                      | Sanma                                                       | gPp     |
| vu-see     | Vanuatu                                      | Shefa Province                                              | gPp     |
| vu-tae     | Vanuatu                                      | Tafea                                                       | gPp     |
| vu-tob     | Vanuatu                                      | Torba                                                       | gPp     |
| ws         |                                              | Samoa                                                       | gPp     |
| ws-aa      | Samoa                                        | Aʻana                                                       | gPp     |
| ws-al      | Samoa                                        | Aiga-i-le-Tai                                               | gPp     |
| ws-at      | Samoa                                        | Ātua                                                        | gPp     |
| ws-fa      | Samoa                                        | Faʻasaleleaga                                               | gPp     |
| ws-ge      | Samoa                                        | Gagaʻemauga                                                 | gPp     |
| ws-gi      | Samoa                                        | Gagaʻifomauga                                               | gPp     |
| ws-pa      | Samoa                                        | Palauli                                                     | gPp     |
| ws-sa      | Samoa                                        | Satupaʻitea                                                 | gPp     |
| ws-tu      | Samoa                                        | Tuamasaga                                                   | gPp     |
| ws-vf      | Samoa                                        | Vaʻa-o-Fonoti                                               | gPp     |
| ws-vs      | Samoa                                        | Vaisigano                                                   | gPp     |
| xk         |                                              | Kosovo                                                      | gPp     |
| ye         |                                              | Yemen                                                       | gPp     |
| ye-ab      | Yemen                                        | Abyan Governorate                                           | gPp     |
| ye-ad      | Yemen                                        | Aden Governorate                                            | gPp     |
| ye-am      | Yemen                                        | 'Amran Governorate                                          | gPp     |
| ye-ba      | Yemen                                        | Al Bayda' Governorate                                       | gPp     |
| ye-da      | Yemen                                        | Ad Dali' Governorate                                        | gPp     |
| ye-dh      | Yemen                                        | Dhamar Governorate                                          | gPp     |
| ye-hd      | Yemen                                        | Hadramaut Governorate                                       | gPp     |
| ye-hj      | Yemen                                        | Hajjah Governorate                                          | gPp     |
| ye-hu      | Yemen                                        | Al Hudaydah Governorate                                     | gPp     |
| ye-ib      | Yemen                                        | Ibb Governorate                                             | gPp     |
| ye-ja      | Yemen                                        | Al Jawf Governorate                                         | gPp     |
| ye-la      | Yemen                                        | Lahij Governorate                                           | gPp     |
| ye-ma      | Yemen                                        | Marib Governorate                                           | gPp     |
| ye-mr      | Yemen                                        | Al Mahrah Governorate                                       | gPp     |
| ye-mw      | Yemen                                        | Al Mahwit Governorate                                       | gPp     |
| ye-ra      | Yemen                                        | Raymah Governorate                                          | gPp     |
| ye-sa      | Yemen                                        | Amanat Al Asimah                                            | gPp     |
| ye-sd      | Yemen                                        | Sa'dah Governorate                                          | gPp     |
| ye-sh      | Yemen                                        | Shabwah Governorate                                         | gPp     |
| ye-sn      | Yemen                                        | Sana'a Governorate                                          | gPp     |
| ye-su      | Yemen                                        | Socotra Governorate                                         | gPp     |
| ye-ta      | Yemen                                        | Ta'izz Governorate                                          | gPp     |
| za         |                                              | South Africa                                                | gPp     |
| za-ec      | South Africa                                 | Eastern Cape                                                | gPp     |
| za-fs      | South Africa                                 | Free State                                                  | gPp     |
| za-gp      | South Africa                                 | Gauteng                                                     | gPp     |
| za-kzn     | South Africa                                 | KwaZulu-Natal                                               | gPp     |
| za-lp      | South Africa                                 | Limpopo                                                     | gPp     |
| za-mp      | South Africa                                 | Mpumalanga                                                  | gPp     |
| za-nc      | South Africa                                 | Northern Cape                                               | gPp     |
| za-nw      | South Africa                                 | North West                                                  | gPp     |
| za-wc      | South Africa                                 | Western Cape                                                | gPp     |
| zm         |                                              | Zambia                                                      | gPp     |
| zm-01      | Zambia                                       | Western Province                                            | gPp     |
| zm-02      | Zambia                                       | Central Province                                            | gPp     |
| zm-03      | Zambia                                       | Eastern Province                                            | gPp     |
| zm-04      | Zambia                                       | Luapula Province                                            | gPp     |
| zm-05      | Zambia                                       | Northern Province                                           | gPp     |
| zm-06      | Zambia                                       | North-Western Province                                      | gPp     |
| zm-07      | Zambia                                       | Southern Province                                           | gPp     |
| zm-08      | Zambia                                       | Copperbelt Province                                         | gPp     |
| zm-09      | Zambia                                       | Lusaka Province                                             | gPp     |
| zm-10      | Zambia                                       | Muchinga Province                                           | gPp     |
| zw         |                                              | Zimbabwe                                                    | gPp     |
| zw-bu      | Zimbabwe                                     | Bulawayo Province                                           | gPp     |
| zw-ha      | Zimbabwe                                     | Harare                                                      | gPp     |
| zw-ma      | Zimbabwe                                     | Manicaland Province                                         | gPp     |
| zw-mc      | Zimbabwe                                     | Mashonaland Central Province                                | gPp     |
| zw-me      | Zimbabwe                                     | Mashonaland East Province                                   | gPp     |
| zw-mi      | Zimbabwe                                     | Midlands Province                                           | gPp     |
| zw-mn      | Zimbabwe                                     | Matabeleland North Province                                 | gPp     |
| zw-ms      | Zimbabwe                                     | Matabeleland South Province                                 | gPp     |
| zw-mv      | Zimbabwe                                     | Masvingo Province                                           | gPp     |
| zw-mw      | Zimbabwe                                     | Mashonaland West Province                                   | gPp     |

</details>

<details>
<summary><strong>Planet OSM Switzerland Extracts (planet.osm.ch)</strong></summary>

| SHORT NAME         | IS IN | LONG NAME          | FORMATS |
|--------------------|-------|--------------------|---------|
| switzerland        |       | switzerland        | oP      |
| switzerland-exact  |       | switzerland-exact  | Pp      |
| switzerland-padded |       | switzerland-padded | Pp      |

</details>

<details>
<summary><strong>OSM Luxembourg Extracts (osm.kewl.lu)</strong></summary>

| SHORT NAME | IS IN | LONG NAME  | FORMATS |
|------------|-------|------------|---------|
| luxembourg |       | luxembourg | BP      |

</details>

<details>
<summary><strong>FIT VUT Brno Extracts (osm.fit.vutbr.cz)</strong></summary>

| SHORT NAME     | IS IN | LONG NAME      | FORMATS |
|----------------|-------|----------------|---------|
| czech_republic |       | Czech Republic | BPp     |

</details>

<details>
<summary><strong>OpenStreetMap Italia Extracts (osmit-estratti)</strong></summary>

| SHORT NAME            | IS IN | LONG NAME             | FORMATS |
|-----------------------|-------|-----------------------|---------|
| abruzzo               |       | Abruzzo               | KoP     |
| agrigento             |       | Agrigento             | KP      |
| alessandria           |       | Alessandria           | KP      |
| ancona                |       | Ancona                | KP      |
| aosta                 |       | Aosta                 | KP      |
| arezzo                |       | Arezzo                | KP      |
| ascoli-piceno         |       | Ascoli Piceno         | KP      |
| asti                  |       | Asti                  | KP      |
| avellino              |       | Avellino              | KP      |
| bari                  |       | Bari                  | KP      |
| barletta-andria-trani |       | Barletta-Andria-Trani | KP      |
| basilicata            |       | Basilicata            | KoP     |
| belluno               |       | Belluno               | KP      |
| benevento             |       | Benevento             | KP      |
| bergamo               |       | Bergamo               | KP      |
| biella                |       | Biella                | KP      |
| bologna               |       | Bologna               | KP      |
| bolzano               |       | Bolzano               | KP      |
| brescia               |       | Brescia               | KP      |
| brindisi              |       | Brindisi              | KP      |
| cagliari              |       | Cagliari              | KP      |
| calabria              |       | Calabria              | KoP     |
| caltanissetta         |       | Caltanissetta         | KP      |
| campania              |       | Campania              | KoP     |
| campobasso            |       | Campobasso            | KP      |
| caserta               |       | Caserta               | KP      |
| catania               |       | Catania               | KP      |
| catanzaro             |       | Catanzaro             | KP      |
| chieti                |       | Chieti                | KP      |
| como                  |       | Como                  | KP      |
| cosenza               |       | Cosenza               | KP      |
| cremona               |       | Cremona               | KP      |
| crotone               |       | Crotone               | KP      |
| cuneo                 |       | Cuneo                 | KP      |
| emilia-romagna        |       | Emilia-Romagna        | KoP     |
| enna                  |       | Enna                  | KP      |
| fermo                 |       | Fermo                 | KP      |
| ferrara               |       | Ferrara               | KP      |
| firenze               |       | Firenze               | KP      |
| foggia                |       | Foggia                | KP      |
| forli-cesena          |       | Forlì-Cesena          | KP      |
| friuli-venezia-giulia |       | Friuli-Venezia Giulia | KoP     |
| frosinone             |       | Frosinone             | KP      |
| genova                |       | Genova                | KP      |
| gorizia               |       | Gorizia               | KP      |
| grosseto              |       | Grosseto              | KP      |
| imperia               |       | Imperia               | KP      |
| isernia               |       | Isernia               | KP      |
| la-spezia             |       | La Spezia             | KP      |
| laquila               |       | L'Aquila              | KP      |
| latina                |       | Latina                | KP      |
| lazio                 |       | Lazio                 | KoP     |
| lecce                 |       | Lecce                 | KP      |
| lecco                 |       | Lecco                 | KP      |
| liguria               |       | Liguria               | KoP     |
| livorno               |       | Livorno               | KP      |
| lodi                  |       | Lodi                  | KP      |
| lombardia             |       | Lombardia             | KoP     |
| lucca                 |       | Lucca                 | KP      |
| macerata              |       | Macerata              | KP      |
| mantova               |       | Mantova               | KP      |
| marche                |       | Marche                | KoP     |
| massa-carrara         |       | Massa-Carrara         | KP      |
| matera                |       | Matera                | KP      |
| messina               |       | Messina               | KP      |
| milano                |       | Milano                | KP      |
| modena                |       | Modena                | KP      |
| molise                |       | Molise                | KoP     |
| monza-brianza         |       | Monza e Brianza       | KP      |
| napoli                |       | Napoli                | KP      |
| novara                |       | Novara                | KP      |
| nuoro                 |       | Nuoro                 | KP      |
| oristano              |       | Oristano              | KP      |
| padova                |       | Padova                | KP      |
| palermo               |       | Palermo               | KP      |
| parma                 |       | Parma                 | KP      |
| pavia                 |       | Pavia                 | KP      |
| perugia               |       | Perugia               | KP      |
| pesaro-urbino         |       | Pesaro e Urbino       | KP      |
| pescara               |       | Pescara               | KP      |
| piacenza              |       | Piacenza              | KP      |
| piemonte              |       | Piemonte              | KoP     |
| pisa                  |       | Pisa                  | KP      |
| pistoia               |       | Pistoia               | KP      |
| pordenone             |       | Pordenone             | KP      |
| potenza               |       | Potenza               | KP      |
| prato                 |       | Prato                 | KP      |
| puglia                |       | Puglia                | KoP     |
| ragusa                |       | Ragusa                | KP      |
| ravenna               |       | Ravenna               | KP      |
| reggio-calabria       |       | Reggio Calabria       | KP      |
| reggio-emilia         |       | Reggio Emilia         | KP      |
| rieti                 |       | Rieti                 | KP      |
| rimini                |       | Rimini                | KP      |
| roma                  |       | Roma                  | KP      |
| rovigo                |       | Rovigo                | KP      |
| salerno               |       | Salerno               | KP      |
| sardegna              |       | Sardegna              | KoP     |
| sassari               |       | Sassari               | KP      |
| savona                |       | Savona                | KP      |
| sicilia               |       | Sicilia               | KoP     |
| siena                 |       | Siena                 | KP      |
| siracusa              |       | Siracusa              | KP      |
| sondrio               |       | Sondrio               | KP      |
| sud-sardegna          |       | Sud Sardegna          | KP      |
| taranto               |       | Taranto               | KP      |
| teramo                |       | Teramo                | KP      |
| terni                 |       | Terni                 | KP      |
| torino                |       | Torino                | KP      |
| toscana               |       | Toscana               | KoP     |
| trapani               |       | Trapani               | KP      |
| trentino-alto-adige   |       | Trentino-Alto Adige   | KoP     |
| trento                |       | Trento                | KP      |
| treviso               |       | Treviso               | KP      |
| trieste               |       | Trieste               | KP      |
| udine                 |       | Udine                 | KP      |
| umbria                |       | Umbria                | KoP     |
| valle-daosta          |       | Valle d'Aosta         | KoP     |
| varese                |       | Varese                | KP      |
| veneto                |       | Veneto                | KoP     |
| venezia               |       | Venezia               | KP      |
| verbano-cusio-ossola  |       | Verbano-Cusio-Ossola  | KP      |
| vercelli              |       | Vercelli              | KP      |
| verona                |       | Verona                | KP      |
| vibo-valentia         |       | Vibo Valentia         | KP      |
| vicenza               |       | Vicenza               | KP      |
| viterbo               |       | Viterbo               | KP      |

</details>

<details>
<summary><strong>OSM Taiwan Extracts (osm.kcwu.csie.org)</strong></summary>

| SHORT NAME | IS IN | LONG NAME | FORMATS |
|------------|-------|-----------|---------|
| taiwan     |       | Taiwan    | 5Z      |

</details>
