# download-geofabrik

[![Join the chat at https://gitter.im/julien-noblet/download-geofabrik](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/julien-noblet/download-geofabrik?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
![Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/julien-noblet/a509e15ea4734ca3e8e98f32ab5369c0/raw/7344619caf8ac5bce291793711071a9636536fce/coverage.json)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjulien-noblet%2Fdownload-geofabrik.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjulien-noblet%2Fdownload-geofabrik?ref=badge_shield)

## Version 2
Warning! command line have changed from V1
see [Usage](#usage)

## Docker
```shell
docker run -it --rm -v $PWD:/data download-geofabrik:latest download element
```
where ```element``` is one of geofabrik's files.
## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjulien-noblet%2Fdownload-geofabrik.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjulien-noblet%2Fdownload-geofabrik?ref=badge_large)

## Usage
```shell
./download-geofabrik download element
```
where ```element``` is one of geofabrik's files.
```shell
./download-geofabrik --help-long

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

## Model Context Protocol (MCP) Mode

`download-geofabrik` includes a built-in **MCP server** allowing Large Language Models (LLMs) such as Claude Desktop, Antigravity, Cursor, and Cline to interact with OpenStreetMap catalogs.

```shell
./download-geofabrik mcp
```

### Configuration for MCP Clients

Add the following to your MCP client configuration (e.g., `claude_desktop_config.json` or `antigravity.json`):

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

- **`list_services`**: List all 10 supported OSM data providers (`geofabrik`, `bbbike`, `openstreetmap.fr`, `movisda`, `geo2day`, etc.) and their configuration status.
- **`regenerate_catalog`**: Trigger remote catalog scraping and update local configuration files.
- **`list_elements`**: Search and list regions, countries, and cities for any provider with available file formats (`osm.pbf`, `shp.zip`, `poly`, `mbtiles`, etc.), supporting search and pagination.
- **`get_element`**: Retrieve detailed metadata for a specific extract, including resolved download URLs and MD5 checksum availability.
- **`list_formats`**: List supported data formats and extensions across services.
- **`download_element`**: Download OSM extracts for specified formats with checksum verification and `dry_run` simulation support.

### Available MCP Resources

- `geofabrik://services`: JSON list of all available provider services.
- `geofabrik://formats`: JSON list of all supported file formats.
- `geofabrik://catalog/{service}`: Full catalog elements and formats for the given service.

## List of elements

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
