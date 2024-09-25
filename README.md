# download-geofabrik

[![Join the chat at https://gitter.im/julien-noblet/download-geofabrik](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/julien-noblet/download-geofabrik?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/julien-noblet/download-geofabrik)](https://goreportcard.com/report/github.com/julien-noblet/download-geofabrik)
[![Coverage Status](https://coveralls.io/repos/github/julien-noblet/download-geofabrik/badge.svg?branch=master)](https://coveralls.io/github/julien-noblet/download-geofabrik?branch=master)
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

usage: download-geofabrik [<flags>] <command> [<args> ...]

A command-line tool for downloading OSM files.

Flags:
      --[no-]help            Show context-sensitive help (also try --help-long
                             and --help-man).
  -c, --config="./geofabrik.yml"  
                             Set Config file.
  -n, --[no-]nodownload      Do not download file (test only)
  -v, --[no-]verbose         Be verbose
  -q, --[no-]quiet           Be quiet
      --[no-]progress        Add a progress bar (implies quiet)
      --service="geofabrik"  Can switch to another service. You can use
                             "geofabrik", "openstreetmap.fr" "osmtoday" or
                             "bbbike". It automatically change config file if -c
                             is unused.
      --[no-]version         Show application version.

Commands:
help [<command>...]
    Show help.


download [<flags>] <element>
    Download element

        --[no-]check           Control with checksum (default) Use --no-check to
                               discard control
    -d, --output_directory=OUTPUT_DIRECTORY  
                               Set output directory, you can use also OUTPUT_DIR
                               env variable
    -B, --[no-]osm.bz2         Download osm.bz2 if available
    -G, --[no-]osm.gz          Download osm.gz if available
    -S, --[no-]shp.zip         Download shp.zip if available
    -P, --[no-]osm.pbf         Download osm.pbf (default)
    -H, --[no-]osh.pbf         Download osh.pbf
    -s, --[no-]state           Download state.txt file
    -p, --[no-]poly            Download poly file
    -k, --[no-]kml             Download kml file
    -g, --[no-]geojson         Download GeoJSON file
    -O, --[no-]garmin-osm      Download Garmin OSM file
    -m, --[no-]mapsforge       Download Mapsforge file
    -M, --[no-]mbtiles         Download MBTiles file
    -C, --[no-]csv             Download CSV file
    -r, --[no-]garminonroad    Download Garmin Onroad file
    -t, --[no-]garminontrail   Download Garmin Ontrail file
    -o, --[no-]garminopenTopo  Download Garmin OpenTopo file

generate
    Generate a new config file


list [<flags>]
    Show elements available

    --[no-]markdown  generate list in Markdown format


```

## List of elements
|                  SHORTNAME                  |            IS IN            |           LONG NAME            | FORMATS |
|---------------------------------------------|-----------------------------|--------------------------------|---------|
| afghanistan                                 | Asia                        | Afghanistan                    | kHBPpSs |
| africa                                      |                             | Africa                         | kHBPps  |
| albania                                     | Europe                      | Albania                        | kHBPpSs |
| alberta                                     | Canada                      | Alberta                        | kHBPpSs |
| algeria                                     | Africa                      | Algeria                        | kHBPpSs |
| alps                                        | Europe                      | Alps                           | kHBPps  |
| alsace                                      | France                      | Alsace                         | kHBPpSs |
| american-oceania                            | Australia and Oceania       | American Oceania               | kHBPpSs |
| andalucia                                   | Spain                       | Andalucía                      | kHBPpSs |
| andorra                                     | Europe                      | Andorra                        | kHBPpSs |
| angola                                      | Africa                      | Angola                         | kHBPpSs |
| anhui                                       | China                       | Anhui                          | kHBPpSs |
| antarctica                                  |                             | Antarctica                     | kHBPpSs |
| aquitaine                                   | France                      | Aquitaine                      | kHBPpSs |
| aragon                                      | Spain                       | Aragón                         | kHBPpSs |
| argentina                                   | South America               | Argentina                      | kHBPpSs |
| armenia                                     | Asia                        | Armenia                        | kHBPpSs |
| arnsberg-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Arnsberg      | kHBPpSs |
| asia                                        |                             | Asia                           | kHBPps  |
| asturias                                    | Spain                       | Asturias                       | kHBPpSs |
| australia                                   | Australia and Oceania       | Australia                      | kHBPpSs |
| australia-oceania                           |                             | Australia and Oceania          | kHBPps  |
| austria                                     | Europe                      | Austria                        | kHBPpSs |
| auvergne                                    | France                      | Auvergne                       | kHBPpSs |
| azerbaijan                                  | Asia                        | Azerbaijan                     | kHBPpSs |
| azores                                      | Europe                      | Azores                         | kHBPpSs |
| baden-wuerttemberg                          | Germany                     | Baden-Württemberg              | kHBPps  |
| bahamas                                     | Central America             | Bahamas                        | kHBPpSs |
| bangladesh                                  | Asia                        | Bangladesh                     | kHBPpSs |
| basse-normandie                             | France                      | Basse-Normandie                | kHBPpSs |
| bayern                                      | Germany                     | Bayern                         | kHBPps  |
| bedfordshire                                | England                     | Bedfordshire                   | kHBPpSs |
| beijing                                     | China                       | Beijing                        | kHBPpSs |
| belarus                                     | Europe                      | Belarus                        | kHBPpSs |
| belgium                                     | Europe                      | Belgium                        | kHBPpSs |
| belize                                      | Central America             | Belize                         | kHBPpSs |
| benin                                       | Africa                      | Benin                          | kHBPpSs |
| berkshire                                   | England                     | Berkshire                      | kHBPpSs |
| berlin                                      | Germany                     | Berlin                         | kHBPpSs |
| bhutan                                      | Asia                        | Bhutan                         | kHBPpSs |
| bolivia                                     | South America               | Bolivia                        | kHBPpSs |
| bosnia-herzegovina                          | Europe                      | Bosnia-Herzegovina             | kHBPpSs |
| botswana                                    | Africa                      | Botswana                       | kHBPpSs |
| bourgogne                                   | France                      | Bourgogne                      | kHBPpSs |
| brandenburg                                 | Germany                     | Brandenburg (mit Berlin)       | kHBPpSs |
| brazil                                      | South America               | Brazil                         | kHBPps  |
| bremen                                      | Germany                     | Bremen                         | kHBPpSs |
| bretagne                                    | France                      | Bretagne                       | kHBPpSs |
| bristol                                     | England                     | Bristol                        | kHBPpSs |
| britain-and-ireland                         | Europe                      | Britain and Ireland            | kHBPps  |
| british-columbia                            | Canada                      | British Columbia               | kHBPpSs |
| buckinghamshire                             | England                     | Buckinghamshire                | kHBPpSs |
| bulgaria                                    | Europe                      | Bulgaria                       | kHBPpSs |
| burkina-faso                                | Africa                      | Burkina Faso                   | kHBPpSs |
| burundi                                     | Africa                      | Burundi                        | kHBPpSs |
| cambodia                                    | Asia                        | Cambodia                       | kHBPpSs |
| cambridgeshire                              | England                     | Cambridgeshire                 | kHBPpSs |
| cameroon                                    | Africa                      | Cameroon                       | kHBPpSs |
| canada                                      | North America               | Canada                         | kHBPps  |
| canary-islands                              | Africa                      | Canary Islands                 | kHBPpSs |
| cantabria                                   | Spain                       | Cantabria                      | kHBPpSs |
| cape-verde                                  | Africa                      | Cape Verde                     | kHBPpSs |
| castilla-la-mancha                          | Spain                       | Castilla-La Mancha             | kHBPpSs |
| castilla-y-leon                             | Spain                       | Castilla y León                | kHBPpSs |
| cataluna                                    | Spain                       | Cataluña                       | kHBPpSs |
| central-african-republic                    | Africa                      | Central African Republic       | kHBPpSs |
| central-america                             |                             | Central America                | kHBPps  |
| central-fed-district                        | Russian Federation          | Central Federal District       | kHBPpSs |
| central-zone                                | India                       | Central Zone                   | kHBPpSs |
| centre                                      | France                      | Centre                         | kHBPpSs |
| centro                                      | Italy                       | Centro                         | kHBPpSs |
| centro-oeste                                | Brazil                      | Centro-Oeste                   | kHBPpSs |
| ceuta                                       | Spain                       | Ceuta                          | kHBPpSs |
| chad                                        | Africa                      | Chad                           | kHBPpSs |
| champagne-ardenne                           | France                      | Champagne Ardenne              | kHBPpSs |
| cheshire                                    | England                     | Cheshire                       | kHBPpSs |
| chile                                       | South America               | Chile                          | kHBPpSs |
| china                                       | Asia                        | China                          | kHBPps  |
| chongqing                                   | China                       | Chongqing                      | kHBPpSs |
| chubu                                       | Japan                       | Chūbu region                   | kHBPpSs |
| chugoku                                     | Japan                       | Chūgoku region                 | kHBPpSs |
| colombia                                    | South America               | Colombia                       | kHBPpSs |
| comores                                     | Africa                      | Comores                        | kHBPpSs |
| congo-brazzaville                           | Africa                      | Congo (Republic/Brazzaville)   | kHBPpSs |
| congo-democratic-republic                   | Africa                      | Congo (Democratic              | kHBPpSs |
|                                             |                             | Republic/Kinshasa)             |         |
| cook-islands                                | Australia and Oceania       | Cook Islands                   | kHBPpSs |
| cornwall                                    | England                     | Cornwall                       | kHBPpSs |
| corse                                       | France                      | Corse                          | kHBPpSs |
| costa-rica                                  | Central America             | Costa Rica                     | kHBPpSs |
| crimean-fed-district                        | Russian Federation          | Crimean Federal District       | kHBPpSs |
| croatia                                     | Europe                      | Croatia                        | kHBPpSs |
| cuba                                        | Central America             | Cuba                           | kHBPpSs |
| cumbria                                     | England                     | Cumbria                        | kHBPpSs |
| cyprus                                      | Europe                      | Cyprus                         | kHBPpSs |
| czech-republic                              | Europe                      | Czech Republic                 | kHBPpSs |
| dach                                        | Europe                      | Germany, Austria, Switzerland  | kHBPps  |
| denmark                                     | Europe                      | Denmark                        | kHBPpSs |
| derbyshire                                  | England                     | Derbyshire                     | kHBPpSs |
| detmold-regbez                              | Nordrhein-Westfalen         | Regierungsbezirk Detmold       | kHBPpSs |
| devon                                       | England                     | Devon                          | kHBPpSs |
| djibouti                                    | Africa                      | Djibouti                       | kHBPpSs |
| dolnoslaskie                                | Poland                      | Województwo dolnośląskie<br    | kHBPpSs |
|                                             |                             | />(Lower Silesian Voivodeship) |         |
| dorset                                      | England                     | Dorset                         | kHBPpSs |
| drenthe                                     | Netherlands                 | Drenthe                        | kHBPpSs |
| duesseldorf-regbez                          | Nordrhein-Westfalen         | Regierungsbezirk Düsseldorf    | kHBPpSs |
| durham                                      | England                     | Durham                         | kHBPpSs |
| east-sussex                                 | England                     | East Sussex                    | kHBPpSs |
| east-timor                                  | Asia                        | East Timor                     | kHBPpSs |
| east-yorkshire-with-hull                    | England                     | East Yorkshire with Hull       | kHBPpSs |
| eastern-zone                                | India                       | Eastern Zone                   | kHBPpSs |
| ecuador                                     | South America               | Ecuador                        | kHBPps  |
| egypt                                       | Africa                      | Egypt                          | kHBPpSs |
| el-salvador                                 | Central America             | El Salvador                    | kHBPpSs |
| enfield                                     | Greater London              | Enfield                        | kHBPpSs |
| england                                     | United Kingdom              | England                        | kHBPpSs |
| equatorial-guinea                           | Africa                      | Equatorial Guinea              | kHBPpSs |
| eritrea                                     | Africa                      | Eritrea                        | kHBPpSs |
| essex                                       | England                     | Essex                          | kHBPpSs |
| estonia                                     | Europe                      | Estonia                        | kHBPpSs |
| ethiopia                                    | Africa                      | Ethiopia                       | kHBPpSs |
| europe                                      |                             | Europe                         | kHBPps  |
| extremadura                                 | Spain                       | Extremadura                    | kHBPpSs |
| far-eastern-fed-district                    | Russian Federation          | Far Eastern Federal District   | kHBPpSs |
| faroe-islands                               | Europe                      | Faroe Islands                  | kHBPpSs |
| fiji                                        | Australia and Oceania       | Fiji                           | kHBPpSs |
| finland                                     | Europe                      | Finland                        | kHBPpSs |
| flevoland                                   | Netherlands                 | Flevoland                      | kHBPpSs |
| france                                      | Europe                      | France                         | kHBPps  |
| franche-comte                               | France                      | Franche Comte                  | kHBPpSs |
| freiburg-regbez                             | Baden-Württemberg           | Regierungsbezirk Freiburg      | kHBPpSs |
| friesland                                   | Netherlands                 | Friesland                      | kHBPpSs |
| fujian                                      | China                       | Fujian                         | kHBPpSs |
| gabon                                       | Africa                      | Gabon                          | kHBPpSs |
| galicia                                     | Spain                       | Galicia                        | kHBPpSs |
| gansu                                       | China                       | Gansu                          | kHBPpSs |
| gcc-states                                  | Asia                        | GCC States                     | kHBPpSs |
| gelderland                                  | Netherlands                 | Gelderland                     | kHBPpSs |
| georgia                                     | Europe                      | Georgia                        | kHBPpSs |
| germany                                     | Europe                      | Germany                        | kHBPps  |
| ghana                                       | Africa                      | Ghana                          | kHBPpSs |
| gloucestershire                             | England                     | Gloucestershire                | kHBPpSs |
| great-britain                               | Europe                      | Great Britain                  | kHBPps  |
| greater-london                              | England                     | Greater London                 | kHBPpSs |
| greater-manchester                          | England                     | Greater Manchester             | kHBPpSs |
| greece                                      | Europe                      | Greece                         | kHBPpSs |
| greenland                                   | North America               | Greenland                      | kHBPpSs |
| groningen                                   | Netherlands                 | Groningen                      | kHBPpSs |
| guadeloupe                                  | France                      | Guadeloupe                     | kHBPps  |
| guangdong                                   | China                       | Guangdong (with Hong Kong and  | kHBPpSs |
|                                             |                             | Macau)                         |         |
| guangxi                                     | China                       | Guangxi                        | kHBPpSs |
| guatemala                                   | Central America             | Guatemala                      | kHBPpSs |
| guernsey-jersey                             | Europe                      | Guernsey and Jersey            | kHBPpSs |
| guinea                                      | Africa                      | Guinea                         | kHBPpSs |
| guinea-bissau                               | Africa                      | Guinea-Bissau                  | kHBPpSs |
| guizhou                                     | China                       | Guizhou                        | kHBPpSs |
| guyana                                      | South America               | Guyana                         | kHBPpSs |
| guyane                                      | France                      | Guyane                         | kHBPpSs |
| hainan                                      | China                       | Hainan                         | kHBPpSs |
| haiti-and-domrep                            | Central America             | Haiti and Dominican Republic   | kHBPpSs |
| hamburg                                     | Germany                     | Hamburg                        | kHBPpSs |
| hampshire                                   | England                     | Hampshire                      | kHBPpSs |
| haute-normandie                             | France                      | Haute-Normandie                | kHBPpSs |
| hebei                                       | China                       | Hebei (with Beijing and        | kHBPpSs |
|                                             |                             | Tianjin)                       |         |
| heilongjiang                                | China                       | Heilongjiang                   | kHBPpSs |
| henan                                       | China                       | Henan                          | kHBPpSs |
| herefordshire                               | England                     | Herefordshire                  | kHBPpSs |
| hertfordshire                               | England                     | Hertfordshire                  | kHBPpSs |
| hessen                                      | Germany                     | Hessen                         | kHBPpSs |
| hokkaido                                    | Japan                       | Hokkaidō                       | kHBPpSs |
| honduras                                    | Central America             | Honduras                       | kHBPpSs |
| hong-kong                                   | China                       | Hong Kong                      | kHBPpSs |
| hubei                                       | China                       | Hubei                          | kHBPpSs |
| hunan                                       | China                       | Hunan                          | kHBPpSs |
| hungary                                     | Europe                      | Hungary                        | kHBPpSs |
| iceland                                     | Europe                      | Iceland                        | kHBPpSs |
| ile-de-clipperton                           | Australia and Oceania       | Île de Clipperton              | kHBPpSs |
| ile-de-france                               | France                      | Ile-de-France                  | kHBPpSs |
| india                                       | Asia                        | India                          | kHBPps  |
| indonesia                                   | Asia                        | Indonesia (with East Timor)    | kHBPps  |
| inner-mongolia                              | China                       | Inner Mongolia                 | kHBPpSs |
| iran                                        | Asia                        | Iran                           | kHBPpSs |
| iraq                                        | Asia                        | Iraq                           | kHBPpSs |
| ireland-and-northern-ireland                | Europe                      | Ireland and Northern Ireland   | kHBPpSs |
| islas-baleares                              | Spain                       | Islas Baleares                 | kHBPpSs |
| isle-of-man                                 | Europe                      | Isle of Man                    | kHBPpSs |
| isle-of-wight                               | England                     | Isle of Wight                  | kHBPpSs |
| isole                                       | Italy                       | Isole                          | kHBPpSs |
| israel-and-palestine                        | Asia                        | Israel and Palestine           | kHBPpSs |
| italy                                       | Europe                      | Italy                          | kHBPps  |
| ivory-coast                                 | Africa                      | Ivory Coast                    | kHBPpSs |
| jamaica                                     | Central America             | Jamaica                        | kHBPpSs |
| japan                                       | Asia                        | Japan                          | kHBPps  |
| java                                        | Indonesia (with East Timor) | Java                           | kHBPpSs |
| jiangsu                                     | China                       | Jiangsu                        | kHBPpSs |
| jiangxi                                     | China                       | Jiangxi                        | kHBPpSs |
| jilin                                       | China                       | Jilin                          | kHBPpSs |
| jordan                                      | Asia                        | Jordan                         | kHBPpSs |
| kalimantan                                  | Indonesia (with East Timor) | Kalimantan                     | kHBPpSs |
| kaliningrad                                 | Russian Federation          | Kaliningrad                    | kHBPpSs |
| kansai                                      | Japan                       | Kansai region (a.k.a. Kinki    | kHBPpSs |
|                                             |                             | region)                        |         |
| kanto                                       | Japan                       | Kantō region                   | kHBPpSs |
| karlsruhe-regbez                            | Baden-Württemberg           | Regierungsbezirk Karlsruhe     | kHBPpSs |
| kazakhstan                                  | Asia                        | Kazakhstan                     | kHBPpSs |
| kent                                        | England                     | Kent                           | kHBPpSs |
| kenya                                       | Africa                      | Kenya                          | kHBPpSs |
| kiribati                                    | Australia and Oceania       | Kiribati                       | kHBPpSs |
| koeln-regbez                                | Nordrhein-Westfalen         | Regierungsbezirk Köln          | kHBPpSs |
| kosovo                                      | Europe                      | Kosovo                         | kHBPpSs |
| kujawsko-pomorskie                          | Poland                      | Województwo                    | kHBPpSs |
|                                             |                             | kujawsko-pomorskie<br          |         |
|                                             |                             | />(Kuyavian-Pomeranian         |         |
|                                             |                             | Voivodeship)                   |         |
| kyrgyzstan                                  | Asia                        | Kyrgyzstan                     | kHBPpSs |
| kyushu                                      | Japan                       | Kyūshū                         | kHBPpSs |
| la-rioja                                    | Spain                       | La Rioja                       | kHBPpSs |
| lancashire                                  | England                     | Lancashire                     | kHBPpSs |
| languedoc-roussillon                        | France                      | Languedoc-Roussillon           | kHBPpSs |
| laos                                        | Asia                        | Laos                           | kHBPpSs |
| latvia                                      | Europe                      | Latvia                         | kHBPpSs |
| lebanon                                     | Asia                        | Lebanon                        | kHBPpSs |
| leicestershire                              | England                     | Leicestershire                 | kHBPpSs |
| lesotho                                     | Africa                      | Lesotho                        | kHBPpSs |
| liaoning                                    | China                       | Liaoning                       | kHBPpSs |
| liberia                                     | Africa                      | Liberia                        | kHBPpSs |
| libya                                       | Africa                      | Libya                          | kHBPpSs |
| liechtenstein                               | Europe                      | Liechtenstein                  | kHBPpSs |
| limburg                                     | Netherlands                 | Limburg                        | kHBPpSs |
| limousin                                    | France                      | Limousin                       | kHBPpSs |
| lincolnshire                                | England                     | Lincolnshire                   | kHBPpSs |
| lithuania                                   | Europe                      | Lithuania                      | kHBPpSs |
| lodzkie                                     | Poland                      | Województwo łódzkie<br />(Łódź | kHBPpSs |
|                                             |                             | Voivodeship)                   |         |
| lorraine                                    | France                      | Lorraine                       | kHBPpSs |
| lubelskie                                   | Poland                      | Województwo lubelskie<br       | kHBPpSs |
|                                             |                             | />(Lublin Voivodeship)         |         |
| lubuskie                                    | Poland                      | Województwo lubuskie<br        | kHBPpSs |
|                                             |                             | />(Lubusz Voivodeship)         |         |
| luxembourg                                  | Europe                      | Luxembourg                     | kHBPpSs |
| macau                                       | China                       | Macau                          | kHBPpSs |
| macedonia                                   | Europe                      | Macedonia                      | kHBPpSs |
| madagascar                                  | Africa                      | Madagascar                     | kHBPpSs |
| madrid                                      | Spain                       | Madrid                         | kHBPpSs |
| malawi                                      | Africa                      | Malawi                         | kHBPpSs |
| malaysia-singapore-brunei                   | Asia                        | Malaysia, Singapore, and       | kHBPpSs |
|                                             |                             | Brunei                         |         |
| maldives                                    | Asia                        | Maldives                       | kHBPpSs |
| mali                                        | Africa                      | Mali                           | kHBPpSs |
| malopolskie                                 | Poland                      | Województwo małopolskie<br     | kHBPpSs |
|                                             |                             | />(Lesser Poland Voivodeship)  |         |
| malta                                       | Europe                      | Malta                          | kHBPpSs |
| maluku                                      | Indonesia (with East Timor) | Maluku                         | kHBPpSs |
| manitoba                                    | Canada                      | Manitoba                       | kHBPpSs |
| marshall-islands                            | Australia and Oceania       | Marshall Islands               | kHBPpSs |
| martinique                                  | France                      | Martinique                     | kHBPps  |
| mauritania                                  | Africa                      | Mauritania                     | kHBPpSs |
| mauritius                                   | Africa                      | Mauritius                      | kHBPpSs |
| mayotte                                     | France                      | Mayotte                        | kHBPps  |
| mazowieckie                                 | Poland                      | Województwo mazowieckie<br     | kHBPpSs |
|                                             |                             | />(Mazovian Voivodeship)       |         |
| mecklenburg-vorpommern                      | Germany                     | Mecklenburg-Vorpommern         | kHBPpSs |
| melilla                                     | Spain                       | Melilla                        | kHBPpSs |
| merseyside                                  | England                     | Merseyside                     | kHBPpSs |
| mexico                                      | North America               | Mexico                         | kHBPpSs |
| micronesia                                  | Australia and Oceania       | Micronesia                     | kHBPpSs |
| midi-pyrenees                               | France                      | Midi-Pyrenees                  | kHBPpSs |
| mittelfranken                               | Bayern                      | Mittelfranken                  | kHBPpSs |
| moldova                                     | Europe                      | Moldova                        | kHBPpSs |
| monaco                                      | Europe                      | Monaco                         | kHBPpSs |
| mongolia                                    | Asia                        | Mongolia                       | kHBPpSs |
| montenegro                                  | Europe                      | Montenegro                     | kHBPpSs |
| morocco                                     | Africa                      | Morocco                        | kHBPpSs |
| mozambique                                  | Africa                      | Mozambique                     | kHBPpSs |
| muenster-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Münster       | kHBPpSs |
| murcia                                      | Spain                       | Murcia                         | kHBPpSs |
| myanmar                                     | Asia                        | Myanmar (a.k.a. Burma)         | kHBPpSs |
| namibia                                     | Africa                      | Namibia                        | kHBPpSs |
| nauru                                       | Australia and Oceania       | Nauru                          | kHBPpSs |
| navarra                                     | Spain                       | Navarra                        | kHBPpSs |
| nepal                                       | Asia                        | Nepal                          | kHBPpSs |
| netherlands                                 | Europe                      | Netherlands                    | kHBPps  |
| new-brunswick                               | Canada                      | New Brunswick                  | kHBPpSs |
| new-caledonia                               | Australia and Oceania       | New Caledonia                  | kHBPpSs |
| new-zealand                                 | Australia and Oceania       | New Zealand                    | kHBPpSs |
| newfoundland-and-labrador                   | Canada                      | Newfoundland and Labrador      | kHBPpSs |
| nicaragua                                   | Central America             | Nicaragua                      | kHBPps  |
| niederbayern                                | Bayern                      | Niederbayern                   | kHBPpSs |
| niedersachsen                               | Germany                     | Niedersachsen                  | kHBPpSs |
| niger                                       | Africa                      | Niger                          | kHBPpSs |
| nigeria                                     | Africa                      | Nigeria                        | kHBPpSs |
| ningxia                                     | China                       | Ningxia                        | kHBPpSs |
| niue                                        | Australia and Oceania       | Niue                           | kHBPpSs |
| noord-brabant                               | Netherlands                 | Noord-Brabant                  | kHBPpSs |
| noord-holland                               | Netherlands                 | Noord-Holland                  | kHBPpSs |
| norcal                                      | us/california               | Northern California            | kHBPpSs |
| nord-est                                    | Italy                       | Nord-Est                       | kHBPpSs |
| nord-ovest                                  | Italy                       | Nord-Ovest                     | kHBPpSs |
| nord-pas-de-calais                          | France                      | Nord-Pas-de-Calais             | kHBPpSs |
| nordeste                                    | Brazil                      | Nordeste                       | kHBPpSs |
| nordrhein-westfalen                         | Germany                     | Nordrhein-Westfalen            | kHBPps  |
| norfolk                                     | England                     | Norfolk                        | kHBPpSs |
| norte                                       | Brazil                      | Norte                          | kHBPpSs |
| north-america                               |                             | North America                  | kHBPps  |
| north-caucasus-fed-district                 | Russian Federation          | North Caucasus Federal         | kHBPpSs |
|                                             |                             | District                       |         |
| north-eastern-zone                          | India                       | North-Eastern Zone             | kHBPpSs |
| north-korea                                 | Asia                        | North Korea                    | kHBPpSs |
| north-yorkshire                             | England                     | North Yorkshire                | kHBPpSs |
| northamptonshire                            | England                     | Northamptonshire               | kHBPpSs |
| northern-zone                               | India                       | Northern Zone                  | kHBPpSs |
| northumberland                              | England                     | Northumberland                 | kHBPpSs |
| northwest-territories                       | Canada                      | Northwest Territories          | kHBPpSs |
| northwestern-fed-district                   | Russian Federation          | Northwestern Federal District  | kHBPpSs |
| norway                                      | Europe                      | Norway                         | kHBPpSs |
| nottinghamshire                             | England                     | Nottinghamshire                | kHBPpSs |
| nova-scotia                                 | Canada                      | Nova Scotia                    | kHBPpSs |
| nunavut                                     | Canada                      | Nunavut                        | kHBPpSs |
| nusa-tenggara                               | Indonesia (with East Timor) | Nusa-Tenggara                  | kHBPpSs |
| oberbayern                                  | Bayern                      | Oberbayern                     | kHBPpSs |
| oberfranken                                 | Bayern                      | Oberfranken                    | kHBPpSs |
| oberpfalz                                   | Bayern                      | Oberpfalz                      | kHBPpSs |
| ontario                                     | Canada                      | Ontario                        | kHBPpSs |
| opolskie                                    | Poland                      | Województwo opolskie<br        | kHBPpSs |
|                                             |                             | />(Opole Voivodeship)          |         |
| overijssel                                  | Netherlands                 | Overijssel                     | kHBPpSs |
| oxfordshire                                 | England                     | Oxfordshire                    | kHBPpSs |
| pais-vasco                                  | Spain                       | País Vasco                     | kHBPpSs |
| pakistan                                    | Asia                        | Pakistan                       | kHBPpSs |
| palau                                       | Australia and Oceania       | Palau                          | kHBPpSs |
| panama                                      | Central America             | Panama                         | kHBPpSs |
| papua                                       | Indonesia (with East Timor) | Papua                          | kHBPpSs |
| papua-new-guinea                            | Australia and Oceania       | Papua New Guinea               | kHBPpSs |
| paraguay                                    | South America               | Paraguay                       | kHBPpSs |
| pays-de-la-loire                            | France                      | Pays de la Loire               | kHBPpSs |
| peru                                        | South America               | Peru                           | kHBPpSs |
| philippines                                 | Asia                        | Philippines                    | kHBPpSs |
| picardie                                    | France                      | Picardie                       | kHBPpSs |
| pitcairn-islands                            | Australia and Oceania       | Pitcairn Islands               | kHBPpSs |
| podkarpackie                                | Poland                      | Województwo podkarpackie<br    | kHBPpSs |
|                                             |                             | />(Subcarpathian Voivodeship)  |         |
| podlaskie                                   | Poland                      | Województwo podlaskie<br       | kHBPpSs |
|                                             |                             | />(Podlaskie Voivodeship)      |         |
| poitou-charentes                            | France                      | Poitou-Charentes               | kHBPpSs |
| poland                                      | Europe                      | Poland                         | kHBPps  |
| polynesie-francaise                         | Australia and Oceania       | Polynésie française (French    | kHBPpSs |
|                                             |                             | Polynesia)                     |         |
| pomorskie                                   | Poland                      | Województwo pomorskie<br       | kHBPpSs |
|                                             |                             | />(Pomeranian Voivodeship)     |         |
| portugal                                    | Europe                      | Portugal                       | kHBPpSs |
| prince-edward-island                        | Canada                      | Prince Edward Island           | kHBPpSs |
| provence-alpes-cote-d-azur                  | France                      | Provence Alpes-Cote-d'Azur     | kHBPpSs |
| qinghai                                     | China                       | Qinghai                        | kHBPpSs |
| quebec                                      | Canada                      | Quebec                         | kHBPpSs |
| reunion                                     | France                      | Reunion                        | kHBPps  |
| rheinland-pfalz                             | Germany                     | Rheinland-Pfalz                | kHBPpSs |
| rhone-alpes                                 | France                      | Rhone-Alpes                    | kHBPpSs |
| romania                                     | Europe                      | Romania                        | kHBPpSs |
| russia                                      |                             | Russian Federation             | kHBPps  |
| rutland                                     | England                     | Rutland                        | kHBPpSs |
| rwanda                                      | Africa                      | Rwanda                         | kHBPpSs |
| saarland                                    | Germany                     | Saarland                       | kHBPpSs |
| sachsen                                     | Germany                     | Sachsen                        | kHBPpSs |
| sachsen-anhalt                              | Germany                     | Sachsen-Anhalt                 | kHBPpSs |
| saint-helena-ascension-and-tristan-da-cunha | Africa                      | Saint Helena, Ascension, and   | kHBPpSs |
|                                             |                             | Tristan da Cunha               |         |
| samoa                                       | Australia and Oceania       | Samoa                          | kHBPpSs |
| sao-tome-and-principe                       | Africa                      | Sao Tome and Principe          | kHBPpSs |
| saskatchewan                                | Canada                      | Saskatchewan                   | kHBPpSs |
| schleswig-holstein                          | Germany                     | Schleswig-Holstein             | kHBPpSs |
| schwaben                                    | Bayern                      | Schwaben                       | kHBPpSs |
| scotland                                    | United Kingdom              | Scotland                       | kHBPpSs |
| senegal-and-gambia                          | Africa                      | Senegal and Gambia             | kHBPpSs |
| serbia                                      | Europe                      | Serbia                         | kHBPpSs |
| seychelles                                  | Africa                      | Seychelles                     | kHBPpSs |
| shaanxi                                     | China                       | Shaanxi                        | kHBPpSs |
| shandong                                    | China                       | Shandong                       | kHBPpSs |
| shanghai                                    | China                       | Shanghai                       | kHBPpSs |
| shanxi                                      | China                       | Shanxi                         | kHBPpSs |
| shikoku                                     | Japan                       | Shikoku                        | kHBPpSs |
| shropshire                                  | England                     | Shropshire                     | kHBPpSs |
| siberian-fed-district                       | Russian Federation          | Siberian Federal District      | kHBPpSs |
| sichuan                                     | China                       | Sichuan                        | kHBPpSs |
| sierra-leone                                | Africa                      | Sierra Leone                   | kHBPpSs |
| slaskie                                     | Poland                      | Województwo śląskie<br         | kHBPpSs |
|                                             |                             | />(Silesian Voivodeship)       |         |
| slovakia                                    | Europe                      | Slovakia                       | kHBPpSs |
| slovenia                                    | Europe                      | Slovenia                       | kHBPpSs |
| socal                                       | us/california               | Southern California            | kHBPpSs |
| solomon-islands                             | Australia and Oceania       | Solomon Islands                | kHBPpSs |
| somalia                                     | Africa                      | Somalia                        | kHBPpSs |
| somerset                                    | England                     | Somerset                       | kHBPpSs |
| south-africa                                | Africa                      | South Africa                   | kHBPpSs |
| south-africa-and-lesotho                    | Africa                      | South Africa (includes         | kHBPps  |
|                                             |                             | Lesotho)                       |         |
| south-america                               |                             | South America                  | kHBPps  |
| south-fed-district                          | Russian Federation          | South Federal District         | kHBPpSs |
| south-korea                                 | Asia                        | South Korea                    | kHBPpSs |
| south-sudan                                 | Africa                      | South Sudan                    | kHBPpSs |
| south-yorkshire                             | England                     | South Yorkshire                | kHBPpSs |
| southern-zone                               | India                       | Southern Zone                  | kHBPpSs |
| spain                                       | Europe                      | Spain                          | kHBPps  |
| sri-lanka                                   | Asia                        | Sri Lanka                      | kHBPpSs |
| staffordshire                               | England                     | Staffordshire                  | kHBPpSs |
| stuttgart-regbez                            | Baden-Württemberg           | Regierungsbezirk Stuttgart     | kHBPpSs |
| sud                                         | Italy                       | Sud                            | kHBPpSs |
| sudan                                       | Africa                      | Sudan                          | kHBPpSs |
| sudeste                                     | Brazil                      | Sudeste                        | kHBPpSs |
| suffolk                                     | England                     | Suffolk                        | kHBPpSs |
| sul                                         | Brazil                      | Sul                            | kHBPpSs |
| sulawesi                                    | Indonesia (with East Timor) | Sulawesi                       | kHBPpSs |
| sumatra                                     | Indonesia (with East Timor) | Sumatra                        | kHBPpSs |
| suriname                                    | South America               | Suriname                       | kHBPpSs |
| surrey                                      | England                     | Surrey                         | kHBPpSs |
| swaziland                                   | Africa                      | Swaziland                      | kHBPpSs |
| sweden                                      | Europe                      | Sweden                         | kHBPpSs |
| swietokrzyskie                              | Poland                      | Województwo świętokrzyskie<br  | kHBPpSs |
|                                             |                             | />(Świętokrzyskie Voivodeship) |         |
| switzerland                                 | Europe                      | Switzerland                    | kHBPpSs |
| syria                                       | Asia                        | Syria                          | kHBPpSs |
| taiwan                                      | Asia                        | Taiwan                         | kHBPpSs |
| tajikistan                                  | Asia                        | Tajikistan                     | kHBPpSs |
| tanzania                                    | Africa                      | Tanzania                       | kHBPpSs |
| thailand                                    | Asia                        | Thailand                       | kHBPpSs |
| thueringen                                  | Germany                     | Thüringen                      | kHBPpSs |
| tianjin                                     | China                       | Tianjin                        | kHBPpSs |
| tibet                                       | China                       | Tibet                          | kHBPpSs |
| togo                                        | Africa                      | Togo                           | kHBPpSs |
| tohoku                                      | Japan                       | Tōhoku region                  | kHBPpSs |
| tokelau                                     | Australia and Oceania       | Tokelau                        | kHBPpSs |
| tonga                                       | Australia and Oceania       | Tonga                          | kHBPpSs |
| tuebingen-regbez                            | Baden-Württemberg           | Regierungsbezirk Tübingen      | kHBPpSs |
| tunisia                                     | Africa                      | Tunisia                        | kHBPpSs |
| turkey                                      | Europe                      | Turkey                         | kHBPpSs |
| turkmenistan                                | Asia                        | Turkmenistan                   | kHBPpSs |
| tuvalu                                      | Australia and Oceania       | Tuvalu                         | kHBPpSs |
| tyne-and-wear                               | England                     | Tyne and Wear                  | kHBPpSs |
| uganda                                      | Africa                      | Uganda                         | kHBPpSs |
| ukraine                                     | Europe                      | Ukraine (with Crimea)          | kHBPpSs |
| united-kingdom                              | Europe                      | United Kingdom                 | kHBPps  |
| unterfranken                                | Bayern                      | Unterfranken                   | kHBPpSs |
| ural-fed-district                           | Russian Federation          | Ural Federal District          | kHBPpSs |
| uruguay                                     | South America               | Uruguay                        | kHBPpSs |
| us                                          | North America               | United States of America       | kHBPps  |
| us-midwest                                  | North America               | US Midwest                     | kHBPps  |
| us-northeast                                | North America               | US Northeast                   | kHBPps  |
| us-pacific                                  | North America               | US Pacific                     | kHBPps  |
| us-south                                    | North America               | US South                       | kHBPps  |
| us-west                                     | North America               | US West                        | kHBPps  |
| us/alabama                                  | North America               | us/alabama                     | kHBPpSs |
| us/alaska                                   | North America               | us/alaska                      | kHBPpSs |
| us/arizona                                  | North America               | us/arizona                     | kHBPpSs |
| us/arkansas                                 | North America               | us/arkansas                    | kHBPpSs |
| us/california                               | North America               | us/california                  | kHBPps  |
| us/colorado                                 | North America               | us/colorado                    | kHBPpSs |
| us/connecticut                              | North America               | us/connecticut                 | kHBPpSs |
| us/delaware                                 | North America               | us/delaware                    | kHBPpSs |
| us/district-of-columbia                     | North America               | us/district-of-columbia        | kHBPpSs |
| us/florida                                  | North America               | us/florida                     | kHBPpSs |
| us/georgia                                  | North America               | Georgia                        | kHBPpSs |
| us/hawaii                                   | North America               | us/hawaii                      | kHBPpSs |
| us/idaho                                    | North America               | us/idaho                       | kHBPpSs |
| us/illinois                                 | North America               | us/illinois                    | kHBPpSs |
| us/indiana                                  | North America               | us/indiana                     | kHBPpSs |
| us/iowa                                     | North America               | us/iowa                        | kHBPpSs |
| us/kansas                                   | North America               | us/kansas                      | kHBPpSs |
| us/kentucky                                 | North America               | us/kentucky                    | kHBPpSs |
| us/louisiana                                | North America               | us/louisiana                   | kHBPpSs |
| us/maine                                    | North America               | us/maine                       | kHBPpSs |
| us/maryland                                 | North America               | us/maryland                    | kHBPpSs |
| us/massachusetts                            | North America               | us/massachusetts               | kHBPpSs |
| us/michigan                                 | North America               | us/michigan                    | kHBPpSs |
| us/minnesota                                | North America               | us/minnesota                   | kHBPpSs |
| us/mississippi                              | North America               | us/mississippi                 | kHBPpSs |
| us/missouri                                 | North America               | us/missouri                    | kHBPpSs |
| us/montana                                  | North America               | us/montana                     | kHBPpSs |
| us/nebraska                                 | North America               | us/nebraska                    | kHBPpSs |
| us/nevada                                   | North America               | us/nevada                      | kHBPpSs |
| us/new-hampshire                            | North America               | us/new-hampshire               | kHBPpSs |
| us/new-jersey                               | North America               | us/new-jersey                  | kHBPpSs |
| us/new-mexico                               | North America               | us/new-mexico                  | kHBPpSs |
| us/new-york                                 | North America               | us/new-york                    | kHBPpSs |
| us/north-carolina                           | North America               | us/north-carolina              | kHBPpSs |
| us/north-dakota                             | North America               | us/north-dakota                | kHBPpSs |
| us/ohio                                     | North America               | us/ohio                        | kHBPpSs |
| us/oklahoma                                 | North America               | us/oklahoma                    | kHBPpSs |
| us/oregon                                   | North America               | us/oregon                      | kHBPpSs |
| us/pennsylvania                             | North America               | us/pennsylvania                | kHBPpSs |
| us/puerto-rico                              | North America               | us/puerto-rico                 | kHBPpSs |
| us/rhode-island                             | North America               | us/rhode-island                | kHBPpSs |
| us/south-carolina                           | North America               | us/south-carolina              | kHBPpSs |
| us/south-dakota                             | North America               | us/south-dakota                | kHBPpSs |
| us/tennessee                                | North America               | us/tennessee                   | kHBPpSs |
| us/texas                                    | North America               | us/texas                       | kHBPpSs |
| us/us-virgin-islands                        | North America               | us/us-virgin-islands           | kHBPpSs |
| us/utah                                     | North America               | us/utah                        | kHBPpSs |
| us/vermont                                  | North America               | us/vermont                     | kHBPpSs |
| us/virginia                                 | North America               | us/virginia                    | kHBPpSs |
| us/washington                               | North America               | us/washington                  | kHBPpSs |
| us/west-virginia                            | North America               | us/west-virginia               | kHBPpSs |
| us/wisconsin                                | North America               | us/wisconsin                   | kHBPpSs |
| us/wyoming                                  | North America               | us/wyoming                     | kHBPpSs |
| utrecht                                     | Netherlands                 | Utrecht                        | kHBPpSs |
| uzbekistan                                  | Asia                        | Uzbekistan                     | kHBPpSs |
| valencia                                    | Spain                       | Valencia                       | kHBPpSs |
| vanuatu                                     | Australia and Oceania       | Vanuatu                        | kHBPpSs |
| venezuela                                   | South America               | Venezuela                      | kHBPpSs |
| vietnam                                     | Asia                        | Vietnam                        | kHBPpSs |
| volga-fed-district                          | Russian Federation          | Volga Federal District         | kHBPpSs |
| wales                                       | United Kingdom              | Wales                          | kHBPpSs |
| wallis-et-futuna                            | Australia and Oceania       | Wallis et Futuna               | kHBPpSs |
| warminsko-mazurskie                         | Poland                      | Województwo                    | kHBPpSs |
|                                             |                             | warmińsko-mazurskie<br         |         |
|                                             |                             | />(Warmian-Masurian            |         |
|                                             |                             | Voivodeship)                   |         |
| warwickshire                                | England                     | Warwickshire                   | kHBPpSs |
| west-midlands                               | England                     | West Midlands                  | kHBPpSs |
| west-sussex                                 | England                     | West Sussex                    | kHBPpSs |
| west-yorkshire                              | England                     | West Yorkshire                 | kHBPpSs |
| western-zone                                | India                       | Western Zone                   | kHBPpSs |
| wielkopolskie                               | Poland                      | Województwo wielkopolskie<br   | kHBPpSs |
|                                             |                             | />(Greater Poland Voivodeship) |         |
| wiltshire                                   | England                     | Wiltshire                      | kHBPpSs |
| worcestershire                              | England                     | Worcestershire                 | kHBPpSs |
| xinjiang                                    | China                       | Xinjiang                       | kHBPpSs |
| yemen                                       | Asia                        | Yemen                          | kHBPpSs |
| yukon                                       | Canada                      | Yukon                          | kHBPpSs |
| yunnan                                      | China                       | Yunnan                         | kHBPpSs |
| zachodniopomorskie                          | Poland                      | Województwo                    | kHBPpSs |
|                                             |                             | zachodniopomorskie<br />(West  |         |
|                                             |                             | Pomeranian Voivodeship)        |         |
| zambia                                      | Africa                      | Zambia                         | kHBPpSs |
| zeeland                                     | Netherlands                 | Zeeland                        | kHBPpSs |
| zhejiang                                    | China                       | Zhejiang                       | kHBPpSs |
| zimbabwe                                    | Africa                      | Zimbabwe                       | kHBPpSs |
| zuid-holland                                | Netherlands                 | Zuid-Holland                   | kHBPpSs |
Total elements: 509

## List of elements from openstreetmap.fr
|                SHORTNAME                 |              IS IN               |                LONG NAME                 | FORMATS |
|------------------------------------------|----------------------------------|------------------------------------------|---------|
| aargau                                   | switzerland                      | aargau                                   | Pps     |
| abitibi_temiscamingue                    | quebec                           | abitibi_temiscamingue                    | Pps     |
| abruzzo                                  | italy                            | abruzzo                                  | Pps     |
| aceh                                     | indonesia                        | aceh                                     | Pps     |
| acre                                     | north                            | acre                                     | Pps     |
| adygea_republic                          | southern_federal_district        | adygea_republic                          | Pps     |
| aegean                                   | turkey                           | aegean                                   | Pps     |
| afghanistan                              | asia                             | afghanistan                              | Pps     |
| africa                                   |                                  | africa                                   | Pps     |
| africa_france_taaf                       | africa                           | africa_france_taaf                       | Pps     |
| aguascalientes                           | mexico                           | aguascalientes                           | Pps     |
| ain                                      | rhone_alpes                      | ain                                      | Pps     |
| aisne                                    | picardie                         | aisne                                    | Pps     |
| akershus                                 | norway                           | akershus                                 | Pps     |
| alagoas                                  | northeast                        | alagoas                                  | Pps     |
| alameda                                  | california                       | alameda                                  | Pps     |
| aland                                    | finland                          | aland                                    | Pps     |
| alava                                    | euskadi                          | alava                                    | Pps     |
| albacete                                 | castilla_la_mancha               | albacete                                 | Pps     |
| alberta                                  | canada                           | alberta                                  | Pps     |
| algeria                                  | africa                           | algeria                                  | Pps     |
| alicante                                 | comunitat_valenciana             | alicante                                 | Pps     |
| allier                                   | auvergne                         | allier                                   | Pps     |
| almeria                                  | andalucia                        | almeria                                  | Pps     |
| alpes_de_haute_provence                  | provence_alpes_cote_d_azur       | alpes_de_haute_provence                  | Pps     |
| alpes_maritimes                          | provence_alpes_cote_d_azur       | alpes_maritimes                          | Pps     |
| alpine                                   | california                       | alpine                                   | Pps     |
| alsace                                   | france                           | alsace                                   | Pps     |
| altai_krai                               | siberian_federal_district        | altai_krai                               | Pps     |
| altai_republic                           | siberian_federal_district        | altai_republic                           | Pps     |
| amador                                   | california                       | amador                                   | Pps     |
| amapa                                    | north                            | amapa                                    | Pps     |
| amazonas                                 | north                            | amazonas                                 | Pps     |
| american_samoa                           | oceania                          | american_samoa                           | Pps     |
| amur_oblast                              | far_eastern_federal_district     | amur_oblast                              | Pps     |
| andalucia                                | spain                            | andalucia                                | Pps     |
| andaman_and_nicobar_islands              | india                            | andaman_and_nicobar_islands              | Pps     |
| andhra_pradesh                           | india                            | andhra_pradesh                           | Pps     |
| angola                                   | africa                           | angola                                   | Pps     |
| anguilla                                 | central-america                  | anguilla                                 | Pps     |
| anhui                                    | china                            | anhui                                    | Pps     |
| antigua_and_barbuda                      | central-america                  | antigua_and_barbuda                      | Pps     |
| antofagasta                              | chile                            | antofagasta                              | Pps     |
| antwerp                                  | flanders                         | antwerp                                  | Pps     |
| appenzell_ausserrhoden                   | switzerland                      | appenzell_ausserrhoden                   | Pps     |
| appenzell_innerrhoden                    | switzerland                      | appenzell_innerrhoden                    | Pps     |
| appomattox                               | virginia                         | appomattox                               | Pps     |
| aquitaine                                | france                           | aquitaine                                | Pps     |
| aragon                                   | spain                            | aragon                                   | Pps     |
| araucania                                | chile                            | araucania                                | Pps     |
| ardeche                                  | rhone_alpes                      | ardeche                                  | Pps     |
| ardennes                                 | champagne_ardenne                | ardennes                                 | Pps     |
| argentina                                | south-america                    | argentina                                | Pps     |
| argentina_la_rioja                       | argentina                        | argentina_la_rioja                       | Pps     |
| argentina_santa_cruz                     | argentina                        | argentina_santa_cruz                     | Pps     |
| arica_y_parinacota                       | chile                            | arica_y_parinacota                       | Pps     |
| ariege                                   | midi_pyrenees                    | ariege                                   | Pps     |
| arkhangelsk_oblast                       | northwestern_federal_district    | arkhangelsk_oblast                       | Pps     |
| armenia                                  | asia                             | armenia                                  | Pps     |
| arnsberg                                 | nordrhein_westfalen              | arnsberg                                 | Pps     |
| aruba                                    | central-america                  | aruba                                    | Pps     |
| arunachal_pradesh                        | india                            | arunachal_pradesh                        | Pps     |
| ashmore_and_cartier_islands              | australia                        | ashmore_and_cartier_islands              | Pps     |
| asia                                     |                                  | asia                                     | Pps     |
| assam                                    | india                            | assam                                    | Pps     |
| astrakhan_oblast                         | southern_federal_district        | astrakhan_oblast                         | Pps     |
| asturias                                 | spain                            | asturias                                 | Pps     |
| atacama                                  | chile                            | atacama                                  | Pps     |
| aube                                     | champagne_ardenne                | aube                                     | Pps     |
| aude                                     | languedoc_roussillon             | aude                                     | Pps     |
| aust-agder                               | norway                           | aust-agder                               | Pps     |
| australia                                | oceania                          | australia                                | Pps     |
| australian_capital_territory             | australia                        | australian_capital_territory             | Pps     |
| austria                                  | europe                           | austria                                  | Pps     |
| auvergne                                 | france                           | auvergne                                 | Pps     |
| aveiro                                   | portugal                         | aveiro                                   | Pps     |
| aveyron                                  | midi_pyrenees                    | aveyron                                  | Pps     |
| avila                                    | castilla_y_leon                  | avila                                    | Pps     |
| aysen                                    | chile                            | aysen                                    | Pps     |
| azores                                   | portugal                         | azores                                   | Pps     |
| badajoz                                  | extremadura                      | badajoz                                  | Pps     |
| bahamas                                  | central-america                  | bahamas                                  | Pps     |
| bahia                                    | northeast                        | bahia                                    | Pps     |
| bahrain                                  | asia                             | bahrain                                  | Pps     |
| baja_california                          | mexico                           | baja_california                          | Pps     |
| baja_california_sur                      | mexico                           | baja_california_sur                      | Pps     |
| bali                                     | indonesia                        | bali                                     | Pps     |
| bangka_belitung_islands                  | indonesia                        | bangka_belitung_islands                  | Pps     |
| bangsamoro                               | philippines                      | bangsamoro                               | Pps     |
| banskobystricky                          | slovakia                         | banskobystricky                          | Pps     |
| banten                                   | indonesia                        | banten                                   | Pps     |
| barbados                                 | central-america                  | barbados                                 | Pps     |
| barcelona                                | catalunya                        | barcelona                                | Pps     |
| bas_rhin                                 | alsace                           | bas_rhin                                 | Pps     |
| bas_saint_laurent                        | quebec                           | bas_saint_laurent                        | Pps     |
| basel_landschaft                         | switzerland                      | basel_landschaft                         | Pps     |
| basel_stadt                              | switzerland                      | basel_stadt                              | Pps     |
| bashkortostan_republic                   | volga_federal_district           | bashkortostan_republic                   | Pps     |
| basilicata                               | italy                            | basilicata                               | Pps     |
| basse_normandie                          | france                           | basse_normandie                          | Pps     |
| beijing                                  | china                            | beijing                                  | Pps     |
| beja                                     | portugal                         | beja                                     | Pps     |
| belgium                                  | europe                           | belgium                                  | Pps     |
| belgorod_oblast                          | central_federal_district         | belgorod_oblast                          | Pps     |
| bengkulu                                 | indonesia                        | bengkulu                                 | Pps     |
| benin                                    | africa                           | benin                                    | Pps     |
| bermuda                                  | north-america                    | bermuda                                  | Pps     |
| bern                                     | switzerland                      | bern                                     | Pps     |
| bhutan                                   | asia                             | bhutan                                   | Pps     |
| bicol_region                             | philippines                      | bicol_region                             | Pps     |
| bihar                                    | india                            | bihar                                    | Pps     |
| biobio                                   | chile                            | biobio                                   | Pps     |
| bir_tawil                                | africa                           | bir_tawil                                | Pps     |
| black_sea                                | turkey                           | black_sea                                | Pps     |
| blekinge                                 | sweden                           | blekinge                                 | Pps     |
| bouches_du_rhone                         | provence_alpes_cote_d_azur       | bouches_du_rhone                         | Pps     |
| bourgogne                                | france                           | bourgogne                                | Pps     |
| bouvet_island                            | africa                           | bouvet_island                            | Pps     |
| braga                                    | portugal                         | braga                                    | Pps     |
| braganca                                 | portugal                         | braganca                                 | Pps     |
| bratislavsky                             | slovakia                         | bratislavsky                             | Pps     |
| brazil                                   | south-america                    | brazil                                   | Pps     |
| brazil_central-west                      | brazil                           | brazil_central-west                      | Pps     |
| brazil_north                             | brazil                           | brazil_north                             | Pps     |
| brazil_northeast                         | brazil                           | brazil_northeast                         | Pps     |
| brazil_south                             | brazil                           | brazil_south                             | Pps     |
| brazil_southeast                         | brazil                           | brazil_southeast                         | Pps     |
| bretagne                                 | france                           | bretagne                                 | Pps     |
| british_columbia                         | canada                           | british_columbia                         | Pps     |
| british_indian_ocean_territory           | asia                             | british_indian_ocean_territory           | Pps     |
| british_virgin_islands                   | central-america                  | british_virgin_islands                   | Pps     |
| brunei                                   | asia                             | brunei                                   | Pps     |
| brussels_capital_region                  | belgium                          | brussels_capital_region                  | Pps     |
| bryansk_oblast                           | central_federal_district         | bryansk_oblast                           | Pps     |
| buenos_aires                             | argentina                        | buenos_aires                             | Pps     |
| buenos_aires_city                        | argentina                        | buenos_aires_city                        | Pps     |
| burgenland                               | austria                          | burgenland                               | Pps     |
| burgos                                   | castilla_y_leon                  | burgos                                   | Pps     |
| burkina_faso                             | africa                           | burkina_faso                             | Pps     |
| burundi                                  | africa                           | burundi                                  | Pps     |
| buryatia_republic                        | siberian_federal_district        | buryatia_republic                        | Pps     |
| buskerud                                 | norway                           | buskerud                                 | Pps     |
| butte                                    | california                       | butte                                    | Pps     |
| cabo_delgado                             | mozambique                       | cabo_delgado                             | Pps     |
| caceres                                  | extremadura                      | caceres                                  | Pps     |
| cadiz                                    | andalucia                        | cadiz                                    | Pps     |
| cagayan_valley                           | philippines                      | cagayan_valley                           | Pps     |
| calabarzon                               | philippines                      | calabarzon                               | Pps     |
| calabria                                 | italy                            | calabria                                 | Pps     |
| calaveras                                | california                       | calaveras                                | Pps     |
| california                               | us-west                          | california                               | Pps     |
| california_lake                          | california                       | california_lake                          | Pps     |
| california_santa_cruz                    | california                       | california_santa_cruz                    | Pps     |
| calvados                                 | basse_normandie                  | calvados                                 | Pps     |
| cambodia                                 | asia                             | cambodia                                 | Pps     |
| cameroon                                 | africa                           | cameroon                                 | Pps     |
| campania                                 | italy                            | campania                                 | Pps     |
| campeche                                 | mexico                           | campeche                                 | Pps     |
| canada                                   | north-america                    | canada                                   | Pps     |
| canarias                                 | spain                            | canarias                                 | Pps     |
| cantabria                                | spain                            | cantabria                                | Pps     |
| cantal                                   | auvergne                         | cantal                                   | Pps     |
| cape_verde                               | africa                           | cape_verde                               | Pps     |
| capital_district                         | new-york                         | capital_district                         | Pps     |
| capitale_nationale                       | quebec                           | capitale_nationale                       | Pps     |
| caraga                                   | philippines                      | caraga                                   | Pps     |
| caribbean                                | central-america                  | caribbean                                | Pps     |
| castellon                                | comunitat_valenciana             | castellon                                | Pps     |
| castelo_branco                           | portugal                         | castelo_branco                           | Pps     |
| castilla_la_mancha                       | spain                            | castilla_la_mancha                       | Pps     |
| castilla_y_leon                          | spain                            | castilla_y_leon                          | Pps     |
| catalunya                                | spain                            | catalunya                                | Pps     |
| catamarca                                | argentina                        | catamarca                                | Pps     |
| cayman_islands                           | central-america                  | cayman_islands                           | Pps     |
| ceara                                    | northeast                        | ceara                                    | Pps     |
| central-america                          |                                  | central-america                          | Pps     |
| central-west                             | brazil                           | central-west                             |         |
| central_african_republic                 | africa                           | central_african_republic                 | Pps     |
| central_anatolia                         | turkey                           | central_anatolia                         | Pps     |
| central_federal_district                 | russia                           | central_federal_district                 | Pps     |
| central_finland                          | finland                          | central_finland                          | Pps     |
| central_java                             | indonesia                        | central_java                             | Pps     |
| central_kalimantan                       | indonesia                        | central_kalimantan                       | Pps     |
| central_luzon                            | philippines                      | central_luzon                            | Pps     |
| central_new_york                         | new-york                         | central_new_york                         | Pps     |
| central_ontario                          | ontario                          | central_ontario                          | Pps     |
| central_ostrobothnia                     | finland                          | central_ostrobothnia                     | Pps     |
| central_papua                            | indonesia                        | central_papua                            | Pps     |
| central_sulawesi                         | indonesia                        | central_sulawesi                         | Pps     |
| central_visayas                          | philippines                      | central_visayas                          | Pps     |
| centre                                   | france                           | centre                                   | Pps     |
| centre_du_quebec                         | quebec                           | centre_du_quebec                         | Pps     |
| ceuta                                    | spain                            | ceuta                                    | Pps     |
| chaco                                    | argentina                        | chaco                                    | Pps     |
| chad                                     | africa                           | chad                                     | Pps     |
| champagne_ardenne                        | france                           | champagne_ardenne                        | Pps     |
| chandigarh                               | india                            | chandigarh                               | Pps     |
| charente                                 | poitou_charentes                 | charente                                 | Pps     |
| charente_maritime                        | poitou_charentes                 | charente_maritime                        | Pps     |
| chaudiere_appalaches                     | quebec                           | chaudiere_appalaches                     | Pps     |
| chechen_republic                         | north_caucasian_federal_district | chechen_republic                         | Pps     |
| chelyabinsk_oblast                       | ural_federal_district            | chelyabinsk_oblast                       | Pps     |
| cher                                     | centre                           | cher                                     | Pps     |
| cherkasy_oblast                          | ukraine                          | cherkasy_oblast                          | Pps     |
| chernihiv_oblast                         | ukraine                          | chernihiv_oblast                         | Pps     |
| chernivtsi_oblast                        | ukraine                          | chernivtsi_oblast                        | Pps     |
| chesapeake                               | virginia                         | chesapeake                               | Pps     |
| chhattisgarh                             | india                            | chhattisgarh                             | Pps     |
| chiapas                                  | mexico                           | chiapas                                  | Pps     |
| chihuahua                                | mexico                           | chihuahua                                | Pps     |
| chile                                    | south-america                    | chile                                    | Pps     |
| china                                    | asia                             | china                                    | Pps     |
| chongqing                                | china                            | chongqing                                | Pps     |
| christmas_island                         | australia                        | christmas_island                         | Pps     |
| chubu                                    | japan                            | chubu                                    | Pps     |
| chubut                                   | argentina                        | chubut                                   | Pps     |
| chugoku                                  | japan                            | chugoku                                  | Pps     |
| chukotka_autonomous_okrug                | far_eastern_federal_district     | chukotka_autonomous_okrug                | Pps     |
| chuvash_republic                         | volga_federal_district           | chuvash_republic                         | Pps     |
| ciudad_real                              | castilla_la_mancha               | ciudad_real                              | Pps     |
| clipperton                               | oceania                          | clipperton                               | Pps     |
| coahuila                                 | mexico                           | coahuila                                 | Pps     |
| coastal                                  | tanzania                         | coastal                                  | Pps     |
| cocos_islands                            | australia                        | cocos_islands                            | Pps     |
| coimbra                                  | portugal                         | coimbra                                  | Pps     |
| colima                                   | mexico                           | colima                                   | Pps     |
| colorado                                 | us-west                          | colorado                                 | Pps     |
| colorado_northeast                       | colorado                         | colorado_northeast                       | Pps     |
| colorado_northwest                       | colorado                         | colorado_northwest                       | Pps     |
| colorado_southeast                       | colorado                         | colorado_southeast                       | Pps     |
| colorado_southwest                       | colorado                         | colorado_southwest                       | Pps     |
| colusa                                   | california                       | colusa                                   | Pps     |
| comoros                                  | africa                           | comoros                                  | Pps     |
| comunidad_de_madrid                      | spain                            | comunidad_de_madrid                      | Pps     |
| comunidad_foral_de_navarra               | spain                            | comunidad_foral_de_navarra               | Pps     |
| comunitat_valenciana                     | spain                            | comunitat_valenciana                     | Pps     |
| congo_brazzaville                        | africa                           | congo_brazzaville                        | Pps     |
| congo_kinshasa                           | africa                           | congo_kinshasa                           | Pps     |
| contra_costa                             | california                       | contra_costa                             | Pps     |
| cook                                     | illinois                         | cook                                     | Pps     |
| cook_islands                             | oceania                          | cook_islands                             | Pps     |
| coquimbo                                 | chile                            | coquimbo                                 | Pps     |
| coral_sea_islands                        | australia                        | coral_sea_islands                        | Pps     |
| cordillera_administrative_region         | philippines                      | cordillera_administrative_region         | Pps     |
| cordoba                                  | argentina                        | cordoba                                  | Pps     |
| correze                                  | limousin                         | correze                                  | Pps     |
| corrientes                               | argentina                        | corrientes                               | Pps     |
| corse                                    | france                           | corse                                    | Pps     |
| corse_du_sud                             | corse                            | corse_du_sud                             | Pps     |
| costa_rica                               | central-america                  | costa_rica                               | Pps     |
| cote_d_or                                | bourgogne                        | cote_d_or                                | Pps     |
| cote_nord                                | quebec                           | cote_nord                                | Pps     |
| cotes_d_armor                            | bretagne                         | cotes_d_armor                            | Pps     |
| creuse                                   | limousin                         | creuse                                   | Pps     |
| crimea                                   | ukraine                          | crimea                                   | Pps     |
| crimea_republic                          | southern_federal_district        | crimea_republic                          | Pps     |
| cuenca                                   | castilla_la_mancha               | cuenca                                   | Pps     |
| culpeper                                 | virginia                         | culpeper                                 | Pps     |
| curacao                                  | central-america                  | curacao                                  | Pps     |
| czech_republic                           | europe                           | czech_republic                           | Pps     |
| dadra_and_nagar_haveli                   | india                            | dadra_and_nagar_haveli                   | Pps     |
| dadra_and_nagar_haveli_and_daman_and_diu | india                            | dadra_and_nagar_haveli_and_daman_and_diu | Pps     |
| dagestan_republic                        | north_caucasian_federal_district | dagestan_republic                        | Pps     |
| dalarna                                  | sweden                           | dalarna                                  | Pps     |
| daman_and_diu                            | india                            | daman_and_diu                            | Pps     |
| davao_region                             | philippines                      | davao_region                             | Pps     |
| del_norte                                | california                       | del_norte                                | Pps     |
| denver                                   | colorado                         | denver                                   | Pps     |
| detmold                                  | nordrhein_westfalen              | detmold                                  | Pps     |
| detroit_metro                            | michigan                         | detroit_metro                            | Pps     |
| deux_sevres                              | poitou_charentes                 | deux_sevres                              | Pps     |
| distrito-federal                         | central-west                     | distrito-federal                         | Pps     |
| djibouti                                 | africa                           | djibouti                                 | Pps     |
| dnipropetrovsk_oblast                    | ukraine                          | dnipropetrovsk_oblast                    | Pps     |
| dolnoslaskie                             | poland                           | dolnoslaskie                             | Pps     |
| dominica                                 | central-america                  | dominica                                 | Pps     |
| dominican_republic                       | central-america                  | dominican_republic                       | Pps     |
| donetsk_oblast                           | ukraine                          | donetsk_oblast                           | Pps     |
| dordogne                                 | aquitaine                        | dordogne                                 | Pps     |
| doubs                                    | franche_comte                    | doubs                                    | Pps     |
| drenthe                                  | netherlands                      | drenthe                                  | Pps     |
| drome                                    | rhone_alpes                      | drome                                    | Pps     |
| durango                                  | mexico                           | durango                                  | Pps     |
| dusseldorf                               | nordrhein_westfalen              | dusseldorf                               | Pps     |
| east_flanders                            | flanders                         | east_flanders                            | Pps     |
| east_java                                | indonesia                        | east_java                                | Pps     |
| east_kalimantan                          | indonesia                        | east_kalimantan                          | Pps     |
| east_midlands                            | england                          | east_midlands                            | Pps     |
| east_nusa_tenggara                       | indonesia                        | east_nusa_tenggara                       | Pps     |
| east_timor                               | asia                             | east_timor                               | Pps     |
| eastern_anatolia                         | turkey                           | eastern_anatolia                         | Pps     |
| eastern_cape                             | south_africa                     | eastern_cape                             | Pps     |
| eastern_ontario                          | ontario                          | eastern_ontario                          | Pps     |
| eastern_visayas                          | philippines                      | eastern_visayas                          | Pps     |
| el_dorado                                | california                       | el_dorado                                | Pps     |
| el_salvador                              | central-america                  | el_salvador                              | Pps     |
| emilia_romagna                           | italy                            | emilia_romagna                           | Pps     |
| england                                  | united_kingdom                   | england                                  | Pps     |
| england_east                             | england                          | england_east                             | Pps     |
| england_north_east                       | england                          | england_north_east                       | Pps     |
| england_north_west                       | england                          | england_north_west                       | Pps     |
| england_south_east                       | england                          | england_south_east                       | Pps     |
| england_south_west                       | england                          | england_south_west                       | Pps     |
| entre_rios                               | argentina                        | entre_rios                               | Pps     |
| equatorial_guinea                        | africa                           | equatorial_guinea                        | Pps     |
| eritrea                                  | africa                           | eritrea                                  | Pps     |
| espirito-santo                           | southeast                        | espirito-santo                           | Pps     |
| essonne                                  | ile_de_france                    | essonne                                  | Pps     |
| estrie                                   | quebec                           | estrie                                   | Pps     |
| ethiopia                                 | africa                           | ethiopia                                 | Pps     |
| eure                                     | haute_normandie                  | eure                                     | Pps     |
| eure_et_loir                             | centre                           | eure_et_loir                             | Pps     |
| europe                                   |                                  | europe                                   | p       |
| europe_france                            | europe                           | europe_france                            | Pps     |
| euskadi                                  | spain                            | euskadi                                  | Pps     |
| evora                                    | portugal                         | evora                                    | Pps     |
| extremadura                              | spain                            | extremadura                              | Pps     |
| fairfax                                  | virginia                         | fairfax                                  | Pps     |
| falkland                                 | south-america                    | falkland                                 | Pps     |
| far_eastern_federal_district             | russia                           | far_eastern_federal_district             | Pps     |
| far_west                                 | new_south_wales                  | far_west                                 | Pps     |
| faro                                     | portugal                         | faro                                     | Pps     |
| fiji                                     | merge                            | fiji                                     | Pps     |
| fiji_east                                | oceania                          | fiji_east                                | Pps     |
| fiji_west                                | oceania                          | fiji_west                                | Pps     |
| finger_lakes                             | new-york                         | finger_lakes                             | Pps     |
| finistere                                | bretagne                         | finistere                                | Pps     |
| finland                                  | europe                           | finland                                  | Pps     |
| finnmark                                 | norway                           | finnmark                                 | Pps     |
| flanders                                 | belgium                          | flanders                                 | Pps     |
| flemish_brabant                          | flanders                         | flemish_brabant                          | Pps     |
| flevoland                                | netherlands                      | flevoland                                | Pps     |
| florida                                  | us-south                         | florida                                  | Pps     |
| florida_east_central                     | florida                          | florida_east_central                     | Pps     |
| florida_northeast                        | florida                          | florida_northeast                        | Pps     |
| florida_northwest                        | florida                          | florida_northwest                        | Pps     |
| florida_southwest                        | florida                          | florida_southwest                        | Pps     |
| formosa                                  | argentina                        | formosa                                  | Pps     |
| france                                   | europe                           | france                                   |         |
| france_metro_dom_com_nc                  | merge                            | france_metro_dom_com_nc                  | Pps     |
| franche_comte                            | france                           | franche_comte                            | Pps     |
| franche_comte_jura                       | franche_comte                    | franche_comte_jura                       | Pps     |
| free_state                               | south_africa                     | free_state                               | Pps     |
| fresno                                   | california                       | fresno                                   | Pps     |
| fribourg                                 | switzerland                      | fribourg                                 | Pps     |
| friesland                                | netherlands                      | friesland                                | Pps     |
| friuli_venezia_giulia                    | italy                            | friuli_venezia_giulia                    | Pps     |
| fujian                                   | china                            | fujian                                   | Pps     |
| gabon                                    | africa                           | gabon                                    | Pps     |
| galicia                                  | spain                            | galicia                                  | Pps     |
| gambia                                   | africa                           | gambia                                   | Pps     |
| gansu                                    | china                            | gansu                                    | Pps     |
| gard                                     | languedoc_roussillon             | gard                                     | Pps     |
| gaspesie_iles_de_la_madeleine            | quebec                           | gaspesie_iles_de_la_madeleine            | Pps     |
| gatorland                                | florida                          | gatorland                                | Pps     |
| gauteng                                  | south_africa                     | gauteng                                  | Pps     |
| gavleborg                                | sweden                           | gavleborg                                | Pps     |
| gaza                                     | mozambique                       | gaza                                     | Pps     |
| gelderland                               | netherlands                      | gelderland                               | Pps     |
| geneva                                   | switzerland                      | geneva                                   | Pps     |
| georgia                                  | asia                             | georgia                                  | Pps     |
| georgia_northeast                        | georgia                          | georgia_northeast                        | Pps     |
| georgia_northwest                        | georgia                          | georgia_northwest                        | Pps     |
| georgia_southeast                        | georgia                          | georgia_southeast                        | Pps     |
| georgia_southwest                        | georgia                          | georgia_southwest                        | Pps     |
| germany                                  | europe                           | germany                                  | Pps     |
| gers                                     | midi_pyrenees                    | gers                                     | Pps     |
| ghana                                    | africa                           | ghana                                    | Pps     |
| gibraltar                                | europe                           | gibraltar                                | Pps     |
| girona                                   | catalunya                        | girona                                   | Pps     |
| gironde                                  | aquitaine                        | gironde                                  | Pps     |
| glarus                                   | switzerland                      | glarus                                   | Pps     |
| glenn                                    | california                       | glenn                                    | Pps     |
| goa                                      | india                            | goa                                      | Pps     |
| goias                                    | central-west                     | goias                                    | Pps     |
| gold_coast                               | florida                          | gold_coast                               | Pps     |
| golden_horseshoe                         | ontario                          | golden_horseshoe                         | Pps     |
| google40770f7a526b2e5f                   |                                  | google40770f7a526b2e5f                   |         |
| gorontalo                                | indonesia                        | gorontalo                                | Pps     |
| gotland                                  | sweden                           | gotland                                  | Pps     |
| granada                                  | andalucia                        | granada                                  | Pps     |
| greater_london                           | england                          | greater_london                           | Pps     |
| greater_metropolitan_newcastle           | new_south_wales                  | greater_metropolitan_newcastle           | Pps     |
| greater_metropolitan_sydney              | new_south_wales                  | greater_metropolitan_sydney              | Pps     |
| grenada                                  | central-america                  | grenada                                  | Pps     |
| grisons                                  | switzerland                      | grisons                                  | Pps     |
| groningen                                | netherlands                      | groningen                                | Pps     |
| guadalajara                              | castilla_la_mancha               | guadalajara                              | Pps     |
| guadeloupe                               | central-america                  | guadeloupe                               | Pps     |
| guam                                     | oceania                          | guam                                     | Pps     |
| guanajuato                               | mexico                           | guanajuato                               | Pps     |
| guangdong                                | china                            | guangdong                                | Pps     |
| guangxi                                  | china                            | guangxi                                  | Pps     |
| guarda                                   | portugal                         | guarda                                   | Pps     |
| guernesey                                | europe                           | guernesey                                | Pps     |
| guerrero                                 | mexico                           | guerrero                                 | Pps     |
| guinea                                   | africa                           | guinea                                   | Pps     |
| guipuzcoa                                | euskadi                          | guipuzcoa                                | Pps     |
| guizhou                                  | china                            | guizhou                                  | Pps     |
| gujarat                                  | india                            | gujarat                                  | Pps     |
| guyana                                   | south-america                    | guyana                                   | Pps     |
| guyane                                   | south-america                    | guyane                                   | Pps     |
| hainan                                   | china                            | hainan                                   | Pps     |
| halland                                  | sweden                           | halland                                  | Pps     |
| haryana                                  | india                            | haryana                                  | Pps     |
| haut_rhin                                | alsace                           | haut_rhin                                | Pps     |
| haute_corse                              | corse                            | haute_corse                              | Pps     |
| haute_garonne                            | midi_pyrenees                    | haute_garonne                            | Pps     |
| haute_loire                              | auvergne                         | haute_loire                              | Pps     |
| haute_marne                              | champagne_ardenne                | haute_marne                              | Pps     |
| haute_normandie                          | france                           | haute_normandie                          | Pps     |
| haute_saone                              | franche_comte                    | haute_saone                              | Pps     |
| haute_savoie                             | rhone_alpes                      | haute_savoie                             | Pps     |
| haute_vienne                             | limousin                         | haute_vienne                             | Pps     |
| hautes_alpes                             | provence_alpes_cote_d_azur       | hautes_alpes                             | Pps     |
| hautes_pyrenees                          | midi_pyrenees                    | hautes_pyrenees                          | Pps     |
| hauts_de_seine                           | ile_de_france                    | hauts_de_seine                           | Pps     |
| heard_island_and_mcdonald_slands         | australia                        | heard_island_and_mcdonald_slands         | Pps     |
| hebei                                    | china                            | hebei                                    | Pps     |
| hedmark                                  | norway                           | hedmark                                  | Pps     |
| heilongjiang                             | china                            | heilongjiang                             | Pps     |
| henan                                    | china                            | henan                                    | Pps     |
| herault                                  | languedoc_roussillon             | herault                                  | Pps     |
| hidalgo                                  | mexico                           | hidalgo                                  | Pps     |
| highland_papua                           | indonesia                        | highland_papua                           | Pps     |
| himachal_pradesh                         | india                            | himachal_pradesh                         | Pps     |
| hokkaido                                 | japan                            | hokkaido                                 | Pps     |
| honduras                                 | central-america                  | honduras                                 | Pps     |
| hong_kong                                | china                            | hong_kong                                | Pps     |
| hordaland                                | norway                           | hordaland                                | Pps     |
| hubei                                    | china                            | hubei                                    | Pps     |
| hudson_valley                            | new-york                         | hudson_valley                            | Pps     |
| huelva                                   | andalucia                        | huelva                                   | Pps     |
| huesca                                   | aragon                           | huesca                                   | Pps     |
| humboldt                                 | california                       | humboldt                                 | Pps     |
| hunan                                    | china                            | hunan                                    | Pps     |
| ile_de_france                            | france                           | ile_de_france                            | Pps     |
| ilemi                                    | africa                           | ilemi                                    | Pps     |
| illawarra                                | new_south_wales                  | illawarra                                | Pps     |
| ille_et_vilaine                          | bretagne                         | ille_et_vilaine                          | Pps     |
| illes_balears                            | spain                            | illes_balears                            | Pps     |
| illinois                                 | us-midwest                       | illinois                                 | Pps     |
| illinois_central                         | illinois                         | illinois_central                         | Pps     |
| illinois_east_central                    | illinois                         | illinois_east_central                    | Pps     |
| illinois_north                           | illinois                         | illinois_north                           | Pps     |
| illinois_northeast                       | illinois                         | illinois_northeast                       | Pps     |
| illinois_northwest                       | illinois                         | illinois_northwest                       | Pps     |
| illinois_southern                        | illinois                         | illinois_southern                        | Pps     |
| illinois_southwest                       | illinois                         | illinois_southwest                       | Pps     |
| ilocos_region                            | philippines                      | ilocos_region                            | Pps     |
| imperial                                 | california                       | imperial                                 | Pps     |
| india                                    | asia                             | india                                    | Pps     |
| indonesia                                | asia                             | indonesia                                | Pps     |
| indre                                    | centre                           | indre                                    | Pps     |
| indre_et_loire                           | centre                           | indre_et_loire                           | Pps     |
| ingushetia_republic                      | north_caucasian_federal_district | ingushetia_republic                      | Pps     |
| inhambane                                | mozambique                       | inhambane                                | Pps     |
| inner_mongolia                           | china                            | inner_mongolia                           | Pps     |
| inyo                                     | california                       | inyo                                     | Pps     |
| ionian_sea                               | seas                             | ionian_sea                               | Pps     |
| ireland                                  | europe                           | ireland                                  | Pps     |
| irkutsk_oblast                           | siberian_federal_district        | irkutsk_oblast                           | Pps     |
| isere                                    | rhone_alpes                      | isere                                    | Pps     |
| israel                                   | asia                             | israel                                   | Pps     |
| israel_and_palestine                     | asia                             | israel_and_palestine                     | Pps     |
| israel_west_bank                         | asia                             | israel_west_bank                         | Pps     |
| italy                                    | europe                           | italy                                    | Pps     |
| ivano-frankivsk_oblast                   | ukraine                          | ivano-frankivsk_oblast                   | Pps     |
| ivanovo_oblast                           | central_federal_district         | ivanovo_oblast                           | Pps     |
| ivory_coast                              | africa                           | ivory_coast                              | Pps     |
| jaen                                     | andalucia                        | jaen                                     | Pps     |
| jakarta                                  | indonesia                        | jakarta                                  | Pps     |
| jalisco                                  | mexico                           | jalisco                                  | Pps     |
| jamaica                                  | central-america                  | jamaica                                  | Pps     |
| jambi                                    | indonesia                        | jambi                                    | Pps     |
| jammu_and_kashmir                        | india                            | jammu_and_kashmir                        | Pps     |
| jamtland                                 | sweden                           | jamtland                                 | Pps     |
| jan_mayen                                | norway                           | jan_mayen                                | Pps     |
| japan                                    | asia                             | japan                                    | Pps     |
| jersey                                   | europe                           | jersey                                   | Pps     |
| jewish_autonomous_oblast                 | far_eastern_federal_district     | jewish_autonomous_oblast                 | Pps     |
| jharkhand                                | india                            | jharkhand                                | Pps     |
| jiangsu                                  | china                            | jiangsu                                  | Pps     |
| jiangxi                                  | china                            | jiangxi                                  | Pps     |
| jihocesky                                | czech_republic                   | jihocesky                                | Pps     |
| jihomoravsky                             | czech_republic                   | jihomoravsky                             | Pps     |
| jilin                                    | china                            | jilin                                    | Pps     |
| jonkoping                                | sweden                           | jonkoping                                | Pps     |
| jujuy                                    | argentina                        | jujuy                                    | Pps     |
| kabardino_balkar_republic                | north_caucasian_federal_district | kabardino_balkar_republic                | Pps     |
| kainuu                                   | finland                          | kainuu                                   | Pps     |
| kaliningrad_oblast                       | northwestern_federal_district    | kaliningrad_oblast                       | Pps     |
| kalmar                                   | sweden                           | kalmar                                   | Pps     |
| kalmykia_republic                        | southern_federal_district        | kalmykia_republic                        | Pps     |
| kaluga_oblast                            | central_federal_district         | kaluga_oblast                            | Pps     |
| kamchatka_krai                           | far_eastern_federal_district     | kamchatka_krai                           | Pps     |
| kansai                                   | japan                            | kansai                                   | Pps     |
| kanta_hame                               | finland                          | kanta_hame                               | Pps     |
| kanto                                    | japan                            | kanto                                    | Pps     |
| karachay_cherkess_republic               | north_caucasian_federal_district | karachay_cherkess_republic               | Pps     |
| karelia_republic                         | northwestern_federal_district    | karelia_republic                         | Pps     |
| karlovarsky                              | czech_republic                   | karlovarsky                              | Pps     |
| karnataka                                | india                            | karnataka                                | Pps     |
| karnten                                  | austria                          | karnten                                  | Pps     |
| kemerovo_oblast                          | siberian_federal_district        | kemerovo_oblast                          | Pps     |
| kenya                                    | africa                           | kenya                                    | Pps     |
| kerala                                   | india                            | kerala                                   | Pps     |
| kern                                     | california                       | kern                                     | Pps     |
| khabarovsk_krai                          | far_eastern_federal_district     | khabarovsk_krai                          | Pps     |
| khakassia_republic                       | siberian_federal_district        | khakassia_republic                       | Pps     |
| khanty_mansi_autonomous_okrug            | ural_federal_district            | khanty_mansi_autonomous_okrug            | Pps     |
| kharkiv_oblast                           | ukraine                          | kharkiv_oblast                           | Pps     |
| kherson_oblast                           | ukraine                          | kherson_oblast                           | Pps     |
| khmelnytskyi_oblast                      | ukraine                          | khmelnytskyi_oblast                      | Pps     |
| kiev                                     | ukraine                          | kiev                                     | Pps     |
| kiev_oblast                              | ukraine                          | kiev_oblast                              | Pps     |
| kings                                    | california                       | kings                                    | Pps     |
| kiribati                                 | merge                            | kiribati                                 | Pps     |
| kiribati_east                            | oceania                          | kiribati_east                            | Pps     |
| kiribati_west                            | oceania                          | kiribati_west                            | Pps     |
| kirov_oblast                             | volga_federal_district           | kirov_oblast                             | Pps     |
| kirovohrad_oblast                        | ukraine                          | kirovohrad_oblast                        | Pps     |
| koln                                     | nordrhein_westfalen              | koln                                     | Pps     |
| komi_republic                            | northwestern_federal_district    | komi_republic                            | Pps     |
| kosicky                                  | slovakia                         | kosicky                                  | Pps     |
| kostroma_oblast                          | central_federal_district         | kostroma_oblast                          | Pps     |
| kralovehradecky                          | czech_republic                   | kralovehradecky                          | Pps     |
| krasnodar_krai                           | southern_federal_district        | krasnodar_krai                           | Pps     |
| krasnoyarsk_krai                         | siberian_federal_district        | krasnoyarsk_krai                         | Pps     |
| kronoberg                                | sweden                           | kronoberg                                | Pps     |
| kujawsko_pomorskie                       | poland                           | kujawsko_pomorskie                       | Pps     |
| kurgan_oblast                            | ural_federal_district            | kurgan_oblast                            | Pps     |
| kursk_oblast                             | central_federal_district         | kursk_oblast                             | Pps     |
| kuwait                                   | asia                             | kuwait                                   | Pps     |
| kwazulu_natal                            | south_africa                     | kwazulu_natal                            | Pps     |
| kymenlaakso                              | finland                          | kymenlaakso                              | Pps     |
| kyushu                                   | japan                            | kyushu                                   | Pps     |
| la_coruna                                | galicia                          | la_coruna                                | Pps     |
| la_pampa                                 | argentina                        | la_pampa                                 | Pps     |
| ladakh                                   | india                            | ladakh                                   | Pps     |
| lakshadweep                              | india                            | lakshadweep                              | Pps     |
| lampung                                  | indonesia                        | lampung                                  | Pps     |
| lanaudiere                               | quebec                           | lanaudiere                               | Pps     |
| landes                                   | aquitaine                        | landes                                   | Pps     |
| languedoc_roussillon                     | france                           | languedoc_roussillon                     | Pps     |
| laos                                     | asia                             | laos                                     | Pps     |
| lapland                                  | finland                          | lapland                                  | Pps     |
| las_palmas                               | canarias                         | las_palmas                               | Pps     |
| lassen                                   | california                       | lassen                                   | Pps     |
| laurentides                              | quebec                           | laurentides                              | Pps     |
| laval                                    | quebec                           | laval                                    | Pps     |
| lazio                                    | italy                            | lazio                                    | Pps     |
| leiria                                   | portugal                         | leiria                                   | Pps     |
| leningrad_oblast                         | northwestern_federal_district    | leningrad_oblast                         | Pps     |
| leon                                     | castilla_y_leon                  | leon                                     | Pps     |
| lesotho                                  | africa                           | lesotho                                  | Pps     |
| liaoning                                 | china                            | liaoning                                 | Pps     |
| liberecky                                | czech_republic                   | liberecky                                | Pps     |
| liguria                                  | italy                            | liguria                                  | Pps     |
| limburg                                  | netherlands                      | limburg                                  | Pps     |
| limousin                                 | france                           | limousin                                 | Pps     |
| limpopo                                  | south_africa                     | limpopo                                  | Pps     |
| lipetsk_oblast                           | central_federal_district         | lipetsk_oblast                           | Pps     |
| lisbon                                   | portugal                         | lisbon                                   | Pps     |
| lleida                                   | catalunya                        | lleida                                   | Pps     |
| lodzkie                                  | poland                           | lodzkie                                  | Pps     |
| loir_et_cher                             | centre                           | loir_et_cher                             | Pps     |
| loire                                    | rhone_alpes                      | loire                                    | Pps     |
| loire_atlantique                         | pays_de_la_loire                 | loire_atlantique                         | Pps     |
| loiret                                   | centre                           | loiret                                   | Pps     |
| lombardia                                | italy                            | lombardia                                | Pps     |
| long_island                              | new-york                         | long_island                              | Pps     |
| lorraine                                 | france                           | lorraine                                 | Pps     |
| los_angeles                              | california                       | los_angeles                              | Pps     |
| los_lagos                                | chile                            | los_lagos                                | Pps     |
| los_rios                                 | chile                            | los_rios                                 | Pps     |
| lot                                      | midi_pyrenees                    | lot                                      | Pps     |
| lot_et_garonne                           | aquitaine                        | lot_et_garonne                           | Pps     |
| lozere                                   | languedoc_roussillon             | lozere                                   | Pps     |
| lubelskie                                | poland                           | lubelskie                                | Pps     |
| lubuskie                                 | poland                           | lubuskie                                 | Pps     |
| lucerne                                  | switzerland                      | lucerne                                  | Pps     |
| lugo                                     | galicia                          | lugo                                     | Pps     |
| luhansk_oblast                           | ukraine                          | luhansk_oblast                           | Pps     |
| luxembourg                               | europe                           | luxembourg                               | Pps     |
| lviv_oblast                              | ukraine                          | lviv_oblast                              | Pps     |
| macau                                    | china                            | macau                                    | Pps     |
| madeira                                  | portugal                         | madeira                                  | Pps     |
| madera                                   | california                       | madera                                   | Pps     |
| madhya_pradesh                           | india                            | madhya_pradesh                           | Pps     |
| magadan_oblast                           | far_eastern_federal_district     | magadan_oblast                           | Pps     |
| magallanes                               | chile                            | magallanes                               | Pps     |
| maharashtra                              | india                            | maharashtra                              | Pps     |
| maine_et_loire                           | pays_de_la_loire                 | maine_et_loire                           | Pps     |
| malaga                                   | andalucia                        | malaga                                   | Pps     |
| malawi                                   | africa                           | malawi                                   | Pps     |
| malaysia                                 | asia                             | malaysia                                 | Pps     |
| maldives                                 | asia                             | maldives                                 | Pps     |
| mali                                     | africa                           | mali                                     | Pps     |
| malopolskie                              | poland                           | malopolskie                              | Pps     |
| maluku                                   | indonesia                        | maluku                                   | Pps     |
| manche                                   | basse_normandie                  | manche                                   | Pps     |
| manica                                   | mozambique                       | manica                                   | Pps     |
| manipur                                  | india                            | manipur                                  | Pps     |
| manitoba                                 | canada                           | manitoba                                 | Pps     |
| maputo                                   | mozambique                       | maputo                                   | Pps     |
| maputo_city                              | mozambique                       | maputo_city                              | Pps     |
| maranhao                                 | northeast                        | maranhao                                 | Pps     |
| marche                                   | italy                            | marche                                   | Pps     |
| mari_el_republic                         | volga_federal_district           | mari_el_republic                         | Pps     |
| marin                                    | california                       | marin                                    | Pps     |
| mariposa                                 | california                       | mariposa                                 | Pps     |
| marmara                                  | turkey                           | marmara                                  | Pps     |
| marne                                    | champagne_ardenne                | marne                                    | Pps     |
| marshall-islands                         | oceania                          | marshall-islands                         | Pps     |
| marshall_islands                         | oceania                          | marshall_islands                         | Pps     |
| martinique                               | central-america                  | martinique                               | Pps     |
| mato-grosso                              | central-west                     | mato-grosso                              | Pps     |
| mato-grosso-do-sul                       | central-west                     | mato-grosso-do-sul                       | Pps     |
| maule                                    | chile                            | maule                                    | Pps     |
| mauricie                                 | quebec                           | mauricie                                 | Pps     |
| mauritania                               | africa                           | mauritania                               | Pps     |
| mauritius                                | africa                           | mauritius                                | Pps     |
| mayenne                                  | pays_de_la_loire                 | mayenne                                  | Pps     |
| mayotte                                  | africa                           | mayotte                                  | Pps     |
| mazowieckie                              | poland                           | mazowieckie                              | Pps     |
| mediterranean                            | turkey                           | mediterranean                            | Pps     |
| meghalaya                                | india                            | meghalaya                                | Pps     |
| melilla                                  | spain                            | melilla                                  | Pps     |
| mendocino                                | california                       | mendocino                                | Pps     |
| mendoza                                  | argentina                        | mendoza                                  | Pps     |
| merced                                   | california                       | merced                                   | Pps     |
| merge                                    |                                  | merge                                    |         |
| merge_france_taaf                        | merge                            | merge_france_taaf                        | Pps     |
| metro_manila                             | philippines                      | metro_manila                             | Pps     |
| meurthe_et_moselle                       | lorraine                         | meurthe_et_moselle                       | Pps     |
| meuse                                    | lorraine                         | meuse                                    | Pps     |
| mexico                                   | north-america                    | mexico                                   | Pps     |
| mexico_city                              | mexico                           | mexico_city                              | Pps     |
| michigan                                 | us-midwest                       | michigan                                 | Pps     |
| michigan_central                         | michigan                         | michigan_central                         | Pps     |
| michigan_southeast                       | michigan                         | michigan_southeast                       | Pps     |
| michigan_southwest                       | michigan                         | michigan_southwest                       | Pps     |
| michigan_west                            | michigan                         | michigan_west                            | Pps     |
| michoacan                                | mexico                           | michoacan                                | Pps     |
| micronesia                               | oceania                          | micronesia                               | Pps     |
| mid_north_coast                          | new_south_wales                  | mid_north_coast                          | Pps     |
| midi_pyrenees                            | france                           | midi_pyrenees                            | Pps     |
| mimaropa                                 | philippines                      | mimaropa                                 | Pps     |
| minas-gerais                             | southeast                        | minas-gerais                             | Pps     |
| misiones                                 | argentina                        | misiones                                 | Pps     |
| mizoram                                  | india                            | mizoram                                  | Pps     |
| modoc                                    | california                       | modoc                                    | Pps     |
| moere_og_romsdal                         | norway                           | moere_og_romsdal                         | Pps     |
| mohawk_valley                            | new-york                         | mohawk_valley                            | Pps     |
| molise                                   | italy                            | molise                                   | Pps     |
| monaco                                   | europe                           | monaco                                   | Pps     |
| mono                                     | california                       | mono                                     | Pps     |
| monteregie                               | quebec                           | monteregie                               | Pps     |
| monterey                                 | california                       | monterey                                 | Pps     |
| montreal                                 | quebec                           | montreal                                 | Pps     |
| montserrat                               | central-america                  | montserrat                               | Pps     |
| moravskoslezsky                          | czech_republic                   | moravskoslezsky                          | Pps     |
| morbihan                                 | bretagne                         | morbihan                                 | Pps     |
| mordovia_republic                        | volga_federal_district           | mordovia_republic                        | Pps     |
| morelos                                  | mexico                           | morelos                                  | Pps     |
| moscow                                   | central_federal_district         | moscow                                   | Pps     |
| moscow_oblast                            | central_federal_district         | moscow_oblast                            | Pps     |
| moselle                                  | lorraine                         | moselle                                  | Pps     |
| mozambique                               | africa                           | mozambique                               | Pps     |
| mpumalanga                               | south_africa                     | mpumalanga                               | Pps     |
| munster                                  | nordrhein_westfalen              | munster                                  | Pps     |
| murmansk_oblast                          | northwestern_federal_district    | murmansk_oblast                          | Pps     |
| murray                                   | new_south_wales                  | murray                                   | Pps     |
| myanmar                                  | asia                             | myanmar                                  | Pps     |
| mykolaiv_oblast                          | ukraine                          | mykolaiv_oblast                          | Pps     |
| nagaland                                 | india                            | nagaland                                 | Pps     |
| namibia                                  | africa                           | namibia                                  | Pps     |
| nampula                                  | mozambique                       | nampula                                  | Pps     |
| napa                                     | california                       | napa                                     | Pps     |
| national_capital_territory_of_delhi      | india                            | national_capital_territory_of_delhi      | Pps     |
| nauru                                    | oceania                          | nauru                                    | Pps     |
| nayarit                                  | mexico                           | nayarit                                  | Pps     |
| nenets_autonomous_okrug                  | northwestern_federal_district    | nenets_autonomous_okrug                  | Pps     |
| netherlands                              | europe                           | netherlands                              | Pps     |
| neuchatel                                | switzerland                      | neuchatel                                | Pps     |
| neuquen                                  | argentina                        | neuquen                                  | Pps     |
| nevada                                   | california                       | nevada                                   | Pps     |
| new-york                                 | us-northeast                     | new-york                                 | Pps     |
| new_brunswick                            | canada                           | new_brunswick                            | Pps     |
| new_caledonia                            | oceania                          | new_caledonia                            | Pps     |
| new_south_wales                          | australia                        | new_south_wales                          | Pps     |
| new_south_wales_central_west             | new_south_wales                  | new_south_wales_central_west             | Pps     |
| new_south_wales_north_western            | new_south_wales                  | new_south_wales_north_western            | Pps     |
| new_south_wales_northern                 | new_south_wales                  | new_south_wales_northern                 | Pps     |
| new_york_city                            | new-york                         | new_york_city                            | Pps     |
| newfoundland_and_labrador                | canada                           | newfoundland_and_labrador                | Pps     |
| niassa                                   | mozambique                       | niassa                                   | Pps     |
| nicaragua                                | central-america                  | nicaragua                                | Pps     |
| nidwalden                                | switzerland                      | nidwalden                                | Pps     |
| niederosterreich                         | austria                          | niederosterreich                         | Pps     |
| nievre                                   | bourgogne                        | nievre                                   | Pps     |
| niger                                    | africa                           | niger                                    | Pps     |
| nigeria                                  | africa                           | nigeria                                  | Pps     |
| nigeria_north_central                    | nigeria                          | nigeria_north_central                    | Pps     |
| nigeria_north_east                       | nigeria                          | nigeria_north_east                       | Pps     |
| nigeria_north_west                       | nigeria                          | nigeria_north_west                       | Pps     |
| nigeria_south_east                       | nigeria                          | nigeria_south_east                       | Pps     |
| nigeria_south_south                      | nigeria                          | nigeria_south_south                      | Pps     |
| nigeria_south_west                       | nigeria                          | nigeria_south_west                       | Pps     |
| ningxia                                  | china                            | ningxia                                  | Pps     |
| nitriansky                               | slovakia                         | nitriansky                               | Pps     |
| niue                                     | oceania                          | niue                                     | Pps     |
| nizhny_novgorod_oblast                   | volga_federal_district           | nizhny_novgorod_oblast                   | Pps     |
| noord_brabant                            | netherlands                      | noord_brabant                            | Pps     |
| noord_holland                            | netherlands                      | noord_holland                            | Pps     |
| nord                                     | nord_pas_de_calais               | nord                                     | Pps     |
| nord_du_quebec                           | quebec                           | nord_du_quebec                           | Pps     |
| nord_pas_de_calais                       | france                           | nord_pas_de_calais                       | Pps     |
| nordland                                 | norway                           | nordland                                 | Pps     |
| nordrhein_westfalen                      | germany                          | nordrhein_westfalen                      | Pps     |
| norfolk_island                           | australia                        | norfolk_island                           | Pps     |
| norrbotten                               | sweden                           | norrbotten                               | Pps     |
| north                                    | brazil                           | north                                    |         |
| north-america                            |                                  | north-america                            | p       |
| north-carolina                           | us-south                         | north-carolina                           | Pps     |
| north-carolina_north_central             | north-carolina                   | north-carolina_north_central             | Pps     |
| north-carolina_northeast                 | north-carolina                   | north-carolina_northeast                 | Pps     |
| north-carolina_northwest                 | north-carolina                   | north-carolina_northwest                 | Pps     |
| north-carolina_south_central             | north-carolina                   | north-carolina_south_central             | Pps     |
| north-carolina_southeast                 | north-carolina                   | north-carolina_southeast                 | Pps     |
| north-carolina_western                   | north-carolina                   | north-carolina_western                   | Pps     |
| north_caucasian_federal_district         | russia                           | north_caucasian_federal_district         | Pps     |
| north_country                            | new-york                         | north_country                            | Pps     |
| north_kalimantan                         | indonesia                        | north_kalimantan                         | Pps     |
| north_karelia                            | finland                          | north_karelia                            | Pps     |
| north_maluku                             | indonesia                        | north_maluku                             | Pps     |
| north_metro                              | georgia                          | north_metro                              | Pps     |
| north_ossetia_alania_republic            | north_caucasian_federal_district | north_ossetia_alania_republic            | Pps     |
| north_ostrobothnia                       | finland                          | north_ostrobothnia                       | Pps     |
| north_savo                               | finland                          | north_savo                               | Pps     |
| north_sea                                | seas                             | north_sea                                | Pps     |
| north_sulawesi                           | indonesia                        | north_sulawesi                           | Pps     |
| north_sumatra                            | indonesia                        | north_sumatra                            | Pps     |
| northeast                                | brazil                           | northeast                                |         |
| northeastern_ontario                     | ontario                          | northeastern_ontario                     | Pps     |
| northern_cape                            | south_africa                     | northern_cape                            | Pps     |
| northern_ireland                         | united_kingdom                   | northern_ireland                         | Pps     |
| northern_lower                           | michigan                         | northern_lower                           | Pps     |
| northern_mariana_islands                 | oceania                          | northern_mariana_islands                 | Pps     |
| northern_mindanao                        | philippines                      | northern_mindanao                        | Pps     |
| northern_territory                       | australia                        | northern_territory                       | Pps     |
| northwest_territories                    | canada                           | northwest_territories                    | Pps     |
| northwestern_federal_district            | russia                           | northwestern_federal_district            | Pps     |
| northwestern_ontario                     | ontario                          | northwestern_ontario                     | Pps     |
| norway                                   | europe                           | norway                                   | Pps     |
| nova_scotia                              | canada                           | nova_scotia                              | Pps     |
| novgorod_oblast                          | northwestern_federal_district    | novgorod_oblast                          | Pps     |
| novosibirsk_oblast                       | siberian_federal_district        | novosibirsk_oblast                       | Pps     |
| nuble                                    | chile                            | nuble                                    | Pps     |
| nuevo_leon                               | mexico                           | nuevo_leon                               | Pps     |
| nunavut                                  | canada                           | nunavut                                  | Pps     |
| o_higgins                                | chile                            | o_higgins                                | Pps     |
| oaxaca                                   | mexico                           | oaxaca                                   | Pps     |
| oberosterreich                           | austria                          | oberosterreich                           | Pps     |
| obwalden                                 | switzerland                      | obwalden                                 | Pps     |
| oceania                                  |                                  | oceania                                  | Pps     |
| oceania_france_taaf                      | oceania                          | oceania_france_taaf                      | Pps     |
| odessa_oblast                            | ukraine                          | odessa_oblast                            | Pps     |
| odisha                                   | india                            | odisha                                   | Pps     |
| oestfold                                 | norway                           | oestfold                                 | Pps     |
| oise                                     | picardie                         | oise                                     | Pps     |
| olomoucky                                | czech_republic                   | olomoucky                                | Pps     |
| oman                                     | asia                             | oman                                     | Pps     |
| omsk_oblast                              | siberian_federal_district        | omsk_oblast                              | Pps     |
| ontario                                  | canada                           | ontario                                  | Pps     |
| opolskie                                 | poland                           | opolskie                                 | Pps     |
| oppland                                  | norway                           | oppland                                  | Pps     |
| orange                                   | california                       | orange                                   | Pps     |
| orebro                                   | sweden                           | orebro                                   | Pps     |
| orenburg_oblast                          | volga_federal_district           | orenburg_oblast                          | Pps     |
| orne                                     | basse_normandie                  | orne                                     | Pps     |
| oryol_oblast                             | central_federal_district         | oryol_oblast                             | Pps     |
| oslo                                     | norway                           | oslo                                     | Pps     |
| ostergotland                             | sweden                           | ostergotland                             | Pps     |
| ostrobothnia                             | finland                          | ostrobothnia                             | Pps     |
| ourense                                  | galicia                          | ourense                                  | Pps     |
| outaouais                                | quebec                           | outaouais                                | Pps     |
| overijssel                               | netherlands                      | overijssel                               | Pps     |
| paijat_hame                              | finland                          | paijat_hame                              | Pps     |
| palau                                    | oceania                          | palau                                    | Pps     |
| palencia                                 | castilla_y_leon                  | palencia                                 | Pps     |
| palestine                                | asia                             | palestine                                | Pps     |
| panama                                   | central-america                  | panama                                   | Pps     |
| panhandle                                | florida                          | panhandle                                | Pps     |
| papua                                    | indonesia                        | papua                                    | Pps     |
| papua_new_guinea                         | oceania                          | papua_new_guinea                         | Pps     |
| para                                     | north                            | para                                     | Pps     |
| paraguay                                 | south-america                    | paraguay                                 | Pps     |
| paraiba                                  | northeast                        | paraiba                                  | Pps     |
| parana                                   | south                            | parana                                   | Pps     |
| pardubicky                               | czech_republic                   | pardubicky                               | Pps     |
| paris                                    | ile_de_france                    | paris                                    | Pps     |
| pas_de_calais                            | nord_pas_de_calais               | pas_de_calais                            | Pps     |
| pays_de_la_loire                         | france                           | pays_de_la_loire                         | Pps     |
| penza_oblast                             | volga_federal_district           | penza_oblast                             | Pps     |
| perm_krai                                | volga_federal_district           | perm_krai                                | Pps     |
| pernambuco                               | northeast                        | pernambuco                               | Pps     |
| philippines                              | asia                             | philippines                              | Pps     |
| piaui                                    | northeast                        | piaui                                    | Pps     |
| picardie                                 | france                           | picardie                                 | Pps     |
| piedmont_triad                           | north-carolina                   | piedmont_triad                           | Pps     |
| piemonte                                 | italy                            | piemonte                                 | Pps     |
| pirkanmaa                                | finland                          | pirkanmaa                                | Pps     |
| pitcairn                                 | oceania                          | pitcairn                                 | Pps     |
| placer                                   | california                       | placer                                   | Pps     |
| plumas                                   | california                       | plumas                                   | Pps     |
| plzensky                                 | czech_republic                   | plzensky                                 | Pps     |
| podkarpackie                             | poland                           | podkarpackie                             | Pps     |
| podlaskie                                | poland                           | podlaskie                                | Pps     |
| poitou_charentes                         | france                           | poitou_charentes                         | Pps     |
| poland                                   | europe                           | poland                                   | Pps     |
| poltava_oblast                           | ukraine                          | poltava_oblast                           | Pps     |
| polygons-merge                           | merge                            | polygons-merge                           |         |
| polygons-merge_france_taaf               | polygons-merge                   | polygons-merge_france_taaf               | p       |
| polynesie                                | oceania                          | polynesie                                | Pps     |
| pomorskie                                | poland                           | pomorskie                                | Pps     |
| pontevedra                               | galicia                          | pontevedra                               | Pps     |
| portalegre                               | portugal                         | portalegre                               | Pps     |
| porto                                    | portugal                         | porto                                    | Pps     |
| portugal                                 | europe                           | portugal                                 | Pps     |
| praha                                    | czech_republic                   | praha                                    | Pps     |
| presovsky                                | slovakia                         | presovsky                                | Pps     |
| primorsky_krai                           | far_eastern_federal_district     | primorsky_krai                           | Pps     |
| prince_edward_island                     | canada                           | prince_edward_island                     | Pps     |
| provence_alpes_cote_d_azur               | france                           | provence_alpes_cote_d_azur               | Pps     |
| pskov_oblast                             | northwestern_federal_district    | pskov_oblast                             | Pps     |
| puducherry                               | india                            | puducherry                               | Pps     |
| puebla                                   | mexico                           | puebla                                   | Pps     |
| puerto_rico                              | central-america                  | puerto_rico                              | Pps     |
| puglia                                   | italy                            | puglia                                   | Pps     |
| punjab                                   | india                            | punjab                                   | Pps     |
| puy_de_dome                              | auvergne                         | puy_de_dome                              | Pps     |
| pyrenees_atlantiques                     | aquitaine                        | pyrenees_atlantiques                     | Pps     |
| pyrenees_orientales                      | languedoc_roussillon             | pyrenees_orientales                      | Pps     |
| qatar                                    | asia                             | qatar                                    | Pps     |
| qinghai                                  | china                            | qinghai                                  | Pps     |
| quebec                                   | canada                           | quebec                                   | Pps     |
| queensland                               | australia                        | queensland                               | Pps     |
| queretaro                                | mexico                           | queretaro                                | Pps     |
| quintana_roo                             | mexico                           | quintana_roo                             | Pps     |
| rajasthan                                | india                            | rajasthan                                | Pps     |
| region_de_murcia                         | spain                            | region_de_murcia                         | Pps     |
| reunion                                  | africa                           | reunion                                  | Pps     |
| rhone                                    | rhone_alpes                      | rhone                                    | Pps     |
| rhone_alpes                              | france                           | rhone_alpes                              | Pps     |
| riau                                     | indonesia                        | riau                                     | Pps     |
| riau_islands                             | indonesia                        | riau_islands                             | Pps     |
| richmond                                 | virginia                         | richmond                                 | Pps     |
| richmond_tweed                           | new_south_wales                  | richmond_tweed                           | Pps     |
| rio-de-janeiro                           | southeast                        | rio-de-janeiro                           | Pps     |
| rio-grande-do-norte                      | northeast                        | rio-grande-do-norte                      | Pps     |
| rio-grande-do-sul                        | south                            | rio-grande-do-sul                        | Pps     |
| rio_negro                                | argentina                        | rio_negro                                | Pps     |
| riverside                                | california                       | riverside                                | Pps     |
| rivne_oblast                             | ukraine                          | rivne_oblast                             | Pps     |
| robots                                   |                                  | robots                                   |         |
| rogaland                                 | norway                           | rogaland                                 | Pps     |
| rondonia                                 | north                            | rondonia                                 | Pps     |
| roraima                                  | north                            | roraima                                  | Pps     |
| rostov_oblast                            | southern_federal_district        | rostov_oblast                            | Pps     |
| russia                                   |                                  | russia                                   | Pps     |
| rwanda                                   | africa                           | rwanda                                   | Pps     |
| ryazan_oblast                            | central_federal_district         | ryazan_oblast                            | Pps     |
| sacramento                               | california                       | sacramento                               | Pps     |
| saguenay_lac_saint_jean                  | quebec                           | saguenay_lac_saint_jean                  | Pps     |
| saint_barthelemy                         | central-america                  | saint_barthelemy                         | Pps     |
| saint_gallen                             | switzerland                      | saint_gallen                             | Pps     |
| saint_helena_ascension_tristan_da_cunha  | africa                           | saint_helena_ascension_tristan_da_cunha  | Pps     |
| saint_kitts_and_nevis                    | central-america                  | saint_kitts_and_nevis                    | Pps     |
| saint_lucia                              | central-america                  | saint_lucia                              | Pps     |
| saint_martin                             | central-america                  | saint_martin                             | Pps     |
| saint_petersburg                         | northwestern_federal_district    | saint_petersburg                         | Pps     |
| saint_pierre_et_miquelon                 | north-america                    | saint_pierre_et_miquelon                 | Pps     |
| saint_vincent_and_the_grenadines         | central-america                  | saint_vincent_and_the_grenadines         | Pps     |
| sakha_republic                           | far_eastern_federal_district     | sakha_republic                           | Pps     |
| sakhalin_oblast                          | far_eastern_federal_district     | sakhalin_oblast                          | Pps     |
| salamanca                                | castilla_y_leon                  | salamanca                                | Pps     |
| salem                                    | virginia                         | salem                                    | Pps     |
| salta                                    | argentina                        | salta                                    | Pps     |
| salzburg                                 | austria                          | salzburg                                 | Pps     |
| samara_oblast                            | volga_federal_district           | samara_oblast                            | Pps     |
| samoa                                    | oceania                          | samoa                                    | Pps     |
| san_benito                               | california                       | san_benito                               | Pps     |
| san_bernardino                           | california                       | san_bernardino                           | Pps     |
| san_diego                                | california                       | san_diego                                | Pps     |
| san_francisco                            | california                       | san_francisco                            | Pps     |
| san_joaquin                              | california                       | san_joaquin                              | Pps     |
| san_juan                                 | argentina                        | san_juan                                 | Pps     |
| san_luis                                 | argentina                        | san_luis                                 | Pps     |
| san_luis_obispo                          | california                       | san_luis_obispo                          | Pps     |
| san_luis_potosi                          | mexico                           | san_luis_potosi                          | Pps     |
| san_marino                               | europe                           | san_marino                               | Pps     |
| san_mateo                                | california                       | san_mateo                                | Pps     |
| santa-catarina                           | south                            | santa-catarina                           | Pps     |
| santa_barbara                            | california                       | santa_barbara                            | Pps     |
| santa_clara                              | california                       | santa_clara                              | Pps     |
| santa_cruz_de_tenerife                   | canarias                         | santa_cruz_de_tenerife                   | Pps     |
| santa_fe                                 | argentina                        | santa_fe                                 | Pps     |
| santarem                                 | portugal                         | santarem                                 | Pps     |
| santiago                                 | chile                            | santiago                                 | Pps     |
| santiago_del_estero                      | argentina                        | santiago_del_estero                      | Pps     |
| sao-paulo                                | southeast                        | sao-paulo                                | Pps     |
| sao_tome_and_principe                    | africa                           | sao_tome_and_principe                    | Pps     |
| saone_et_loire                           | bourgogne                        | saone_et_loire                           | Pps     |
| saratov_oblast                           | volga_federal_district           | saratov_oblast                           | Pps     |
| sardegna                                 | italy                            | sardegna                                 | Pps     |
| sarthe                                   | pays_de_la_loire                 | sarthe                                   | Pps     |
| saskatchewan                             | canada                           | saskatchewan                             | Pps     |
| satakunta                                | finland                          | satakunta                                | Pps     |
| saudi_arabia                             | asia                             | saudi_arabia                             | Pps     |
| savoie                                   | rhone_alpes                      | savoie                                   | Pps     |
| schaffhausen                             | switzerland                      | schaffhausen                             | Pps     |
| schwyz                                   | switzerland                      | schwyz                                   | Pps     |
| seas                                     | europe                           | seas                                     |         |
| segovia                                  | castilla_y_leon                  | segovia                                  | Pps     |
| seine_et_marne                           | ile_de_france                    | seine_et_marne                           | Pps     |
| seine_maritime                           | haute_normandie                  | seine_maritime                           | Pps     |
| seine_saint_denis                        | ile_de_france                    | seine_saint_denis                        | Pps     |
| senegal                                  | africa                           | senegal                                  | Pps     |
| sergipe                                  | northeast                        | sergipe                                  | Pps     |
| setubal                                  | portugal                         | setubal                                  | Pps     |
| sevilla                                  | andalucia                        | sevilla                                  | Pps     |
| seychelles                               | africa                           | seychelles                               | Pps     |
| shaanxi                                  | china                            | shaanxi                                  | Pps     |
| shandong                                 | china                            | shandong                                 | Pps     |
| shanghai                                 | china                            | shanghai                                 | Pps     |
| shanxi                                   | china                            | shanxi                                   | Pps     |
| shasta                                   | california                       | shasta                                   | Pps     |
| shikoku                                  | japan                            | shikoku                                  | Pps     |
| siberian_federal_district                | russia                           | siberian_federal_district                | Pps     |
| sichuan                                  | china                            | sichuan                                  | Pps     |
| sicilia                                  | italy                            | sicilia                                  | Pps     |
| sierra                                   | california                       | sierra                                   | Pps     |
| sikkim                                   | india                            | sikkim                                   | Pps     |
| sinaloa                                  | mexico                           | sinaloa                                  | Pps     |
| singapore                                | asia                             | singapore                                | Pps     |
| sint_maarten                             | central-america                  | sint_maarten                             | Pps     |
| siskiyou                                 | california                       | siskiyou                                 | Pps     |
| skane                                    | sweden                           | skane                                    | Pps     |
| slaskie                                  | poland                           | slaskie                                  | Pps     |
| slovakia                                 | europe                           | slovakia                                 | Pps     |
| smolensk_oblast                          | central_federal_district         | smolensk_oblast                          | Pps     |
| soccsksargen                             | philippines                      | soccsksargen                             | Pps     |
| sodermanland                             | sweden                           | sodermanland                             | Pps     |
| sofala                                   | mozambique                       | sofala                                   | Pps     |
| sogn_og_fjordane                         | norway                           | sogn_og_fjordane                         | Pps     |
| solano                                   | california                       | solano                                   | Pps     |
| solomon_islands                          | oceania                          | solomon_islands                          | Pps     |
| solothurn                                | switzerland                      | solothurn                                | Pps     |
| somme                                    | picardie                         | somme                                    | Pps     |
| sonoma                                   | california                       | sonoma                                   | Pps     |
| sonora                                   | mexico                           | sonora                                   | Pps     |
| soria                                    | castilla_y_leon                  | soria                                    | Pps     |
| south                                    | brazil                           | south                                    |         |
| south-america                            |                                  | south-america                            | Pps     |
| south_africa                             | africa                           | south_africa                             | Pps     |
| south_africa_north_west                  | south_africa                     | south_africa_north_west                  | Pps     |
| south_australia                          | australia                        | south_australia                          | Pps     |
| south_east_region                        | new_south_wales                  | south_east_region                        | Pps     |
| south_georgia_and_south_sandwich         | south-america                    | south_georgia_and_south_sandwich         | Pps     |
| south_kalimantan                         | indonesia                        | south_kalimantan                         | Pps     |
| south_karelia                            | finland                          | south_karelia                            | Pps     |
| south_ostrobothnia                       | finland                          | south_ostrobothnia                       | Pps     |
| south_papua                              | indonesia                        | south_papua                              | Pps     |
| south_savo                               | finland                          | south_savo                               | Pps     |
| south_sudan                              | africa                           | south_sudan                              | Pps     |
| south_sulawesi                           | indonesia                        | south_sulawesi                           | Pps     |
| south_sumatra                            | indonesia                        | south_sumatra                            | Pps     |
| southeast                                | brazil                           | southeast                                |         |
| southeast_sulawesi                       | indonesia                        | southeast_sulawesi                       | Pps     |
| southeastern_anatolia                    | turkey                           | southeastern_anatolia                    | Pps     |
| southern_federal_district                | russia                           | southern_federal_district                | Pps     |
| southern_federal_district_sevastopol     | southern_federal_district        | southern_federal_district_sevastopol     | Pps     |
| southern_highlands                       | tanzania                         | southern_highlands                       | Pps     |
| southern_tier                            | new-york                         | southern_tier                            | Pps     |
| southwest_finland                        | finland                          | southwest_finland                        | Pps     |
| southwest_papua                          | indonesia                        | southwest_papua                          | Pps     |
| southwestern_ontario                     | ontario                          | southwestern_ontario                     | Pps     |
| spain                                    | europe                           | spain                                    | Pps     |
| spain_la_rioja                           | spain                            | spain_la_rioja                           | Pps     |
| stanislaus                               | california                       | stanislaus                               | Pps     |
| state_of_mexico                          | mexico                           | state_of_mexico                          | Pps     |
| stavropol_krai                           | north_caucasian_federal_district | stavropol_krai                           | Pps     |
| steiermark                               | austria                          | steiermark                               | Pps     |
| stockholm                                | sweden                           | stockholm                                | Pps     |
| stredocesky                              | czech_republic                   | stredocesky                              | Pps     |
| sudan                                    | africa                           | sudan                                    | Pps     |
| sumy_oblast                              | ukraine                          | sumy_oblast                              | Pps     |
| suncoast                                 | florida                          | suncoast                                 | Pps     |
| suriname                                 | south-america                    | suriname                                 | Pps     |
| sutter                                   | california                       | sutter                                   | Pps     |
| svalbard                                 | norway                           | svalbard                                 | Pps     |
| sverdlovsk_oblast                        | ural_federal_district            | sverdlovsk_oblast                        | Pps     |
| swaziland                                | africa                           | swaziland                                | Pps     |
| sweden                                   | europe                           | sweden                                   | Pps     |
| swietokrzyskie                           | poland                           | swietokrzyskie                           | Pps     |
| switzerland                              | europe                           | switzerland                              | Pps     |
| switzerland_jura                         | switzerland                      | switzerland_jura                         | Pps     |
| sydney_surrounds                         | new_south_wales                  | sydney_surrounds                         | Pps     |
| tabasco                                  | mexico                           | tabasco                                  | Pps     |
| tamaulipas                               | mexico                           | tamaulipas                               | Pps     |
| tambov_oblast                            | central_federal_district         | tambov_oblast                            | Pps     |
| tamil_nadu                               | india                            | tamil_nadu                               | Pps     |
| tanzania                                 | africa                           | tanzania                                 | Pps     |
| tanzania_central                         | tanzania                         | tanzania_central                         | Pps     |
| tanzania_lake                            | tanzania                         | tanzania_lake                            | Pps     |
| tanzania_northern                        | tanzania                         | tanzania_northern                        | Pps     |
| tanzania_western                         | tanzania                         | tanzania_western                         | Pps     |
| tarapaca                                 | chile                            | tarapaca                                 | Pps     |
| tarn                                     | midi_pyrenees                    | tarn                                     | Pps     |
| tarn_et_garonne                          | midi_pyrenees                    | tarn_et_garonne                          | Pps     |
| tarragona                                | catalunya                        | tarragona                                | Pps     |
| tasmania                                 | australia                        | tasmania                                 | Pps     |
| tatarstan_republic                       | volga_federal_district           | tatarstan_republic                       | Pps     |
| tehama                                   | california                       | tehama                                   | Pps     |
| telangana                                | india                            | telangana                                | Pps     |
| telemark                                 | norway                           | telemark                                 | Pps     |
| ternopil_oblast                          | ukraine                          | ternopil_oblast                          | Pps     |
| territoire_de_belfort                    | franche_comte                    | territoire_de_belfort                    | Pps     |
| teruel                                   | aragon                           | teruel                                   | Pps     |
| tete                                     | mozambique                       | tete                                     | Pps     |
| texas                                    | us-south                         | texas                                    | Pps     |
| texas_central                            | texas                            | texas_central                            | Pps     |
| texas_north                              | texas                            | texas_north                              | Pps     |
| texas_northwest                          | texas                            | texas_northwest                          | Pps     |
| texas_south                              | texas                            | texas_south                              | Pps     |
| texas_southeast                          | texas                            | texas_southeast                          | Pps     |
| texas_west                               | texas                            | texas_west                               | Pps     |
| the_riverina                             | new_south_wales                  | the_riverina                             | Pps     |
| thurgau                                  | switzerland                      | thurgau                                  | Pps     |
| tianjin                                  | china                            | tianjin                                  | Pps     |
| tibet                                    | china                            | tibet                                    | Pps     |
| ticino                                   | switzerland                      | ticino                                   | Pps     |
| tierra_del_fuego                         | argentina                        | tierra_del_fuego                         | Pps     |
| tirol                                    | austria                          | tirol                                    | Pps     |
| tlaxcala                                 | mexico                           | tlaxcala                                 | Pps     |
| tocantins                                | north                            | tocantins                                | Pps     |
| togo                                     | africa                           | togo                                     | Pps     |
| tohoku                                   | japan                            | tohoku                                   | Pps     |
| tokelau                                  | oceania                          | tokelau                                  | Pps     |
| toledo                                   | castilla_la_mancha               | toledo                                   | Pps     |
| tomsk_oblast                             | siberian_federal_district        | tomsk_oblast                             | Pps     |
| tonga                                    | oceania                          | tonga                                    | Pps     |
| toscana                                  | italy                            | toscana                                  | Pps     |
| treasure_coast                           | florida                          | treasure_coast                           | Pps     |
| trenciansky                              | slovakia                         | trenciansky                              | Pps     |
| trentino_alto_adige                      | italy                            | trentino_alto_adige                      | Pps     |
| trinidad_and_tobago                      | central-america                  | trinidad_and_tobago                      | Pps     |
| trinity                                  | california                       | trinity                                  | Pps     |
| tripura                                  | india                            | tripura                                  | Pps     |
| trnavsky                                 | slovakia                         | trnavsky                                 | Pps     |
| troendelag                               | norway                           | troendelag                               | Pps     |
| troms                                    | norway                           | troms                                    | Pps     |
| tucuman                                  | argentina                        | tucuman                                  | Pps     |
| tula_oblast                              | central_federal_district         | tula_oblast                              | Pps     |
| tulare                                   | california                       | tulare                                   | Pps     |
| tunisia                                  | africa                           | tunisia                                  | Pps     |
| tuolumne                                 | california                       | tuolumne                                 | Pps     |
| turkey                                   | europe                           | turkey                                   | Pps     |
| turks_and_caicos_islands                 | central-america                  | turks_and_caicos_islands                 | Pps     |
| tuva_republic                            | siberian_federal_district        | tuva_republic                            | Pps     |
| tuvalu                                   | oceania                          | tuvalu                                   | Pps     |
| tver_oblast                              | central_federal_district         | tver_oblast                              | Pps     |
| tyumen_oblast                            | ural_federal_district            | tyumen_oblast                            | Pps     |
| udmurt_republic                          | volga_federal_district           | udmurt_republic                          | Pps     |
| uganda                                   | africa                           | uganda                                   | Pps     |
| uganda_central                           | uganda                           | uganda_central                           | Pps     |
| uganda_eastern                           | uganda                           | uganda_eastern                           | Pps     |
| uganda_northern                          | uganda                           | uganda_northern                          | Pps     |
| uganda_western                           | uganda                           | uganda_western                           | Pps     |
| ukraine                                  | europe                           | ukraine                                  | Pps     |
| ukraine_sevastopol                       | ukraine                          | ukraine_sevastopol                       | Pps     |
| ulyanovsk_oblast                         | volga_federal_district           | ulyanovsk_oblast                         | Pps     |
| umbria                                   | italy                            | umbria                                   | Pps     |
| united_arab_emirates                     | asia                             | united_arab_emirates                     | Pps     |
| united_kingdom                           | europe                           | united_kingdom                           | Pps     |
| united_states_virgin_islands             | central-america                  | united_states_virgin_islands             | Pps     |
| upper_peninsula                          | michigan                         | upper_peninsula                          | Pps     |
| uppsala                                  | sweden                           | uppsala                                  | Pps     |
| ural_federal_district                    | russia                           | ural_federal_district                    | Pps     |
| uri                                      | switzerland                      | uri                                      | Pps     |
| us-midwest                               | north-america                    | us-midwest                               | Pps     |
| us-northeast                             | north-america                    | us-northeast                             | Pps     |
| us-south                                 | north-america                    | us-south                                 | Pps     |
| us-west                                  | north-america                    | us-west                                  | Pps     |
| usa_virgin_islands                       | central-america                  | usa_virgin_islands                       | Pps     |
| ustecky                                  | czech_republic                   | ustecky                                  | Pps     |
| utrecht                                  | netherlands                      | utrecht                                  | Pps     |
| uttar_pradesh                            | india                            | uttar_pradesh                            | Pps     |
| uttarakhand                              | india                            | uttarakhand                              | Pps     |
| uusimaa                                  | finland                          | uusimaa                                  | Pps     |
| val_d_oise                               | ile_de_france                    | val_d_oise                               | Pps     |
| val_de_marne                             | ile_de_france                    | val_de_marne                             | Pps     |
| valais                                   | switzerland                      | valais                                   | Pps     |
| valencia                                 | comunitat_valenciana             | valencia                                 | Pps     |
| valladolid                               | castilla_y_leon                  | valladolid                               | Pps     |
| valle_aosta                              | italy                            | valle_aosta                              | Pps     |
| valparaiso                               | chile                            | valparaiso                               | Pps     |
| vanuatu                                  | oceania                          | vanuatu                                  | Pps     |
| var                                      | provence_alpes_cote_d_azur       | var                                      | Pps     |
| varmland                                 | sweden                           | varmland                                 | Pps     |
| vasterbotten                             | sweden                           | vasterbotten                             | Pps     |
| vasternorrland                           | sweden                           | vasternorrland                           | Pps     |
| vastmanland                              | sweden                           | vastmanland                              | Pps     |
| vastra_gotaland                          | sweden                           | vastra_gotaland                          | Pps     |
| vatican_city                             | europe                           | vatican_city                             | Pps     |
| vaucluse                                 | provence_alpes_cote_d_azur       | vaucluse                                 | Pps     |
| vaud                                     | switzerland                      | vaud                                     | Pps     |
| vendee                                   | pays_de_la_loire                 | vendee                                   | Pps     |
| veneto                                   | italy                            | veneto                                   | Pps     |
| venezuela                                | south-america                    | venezuela                                | Pps     |
| ventura                                  | california                       | ventura                                  | Pps     |
| veracruz                                 | mexico                           | veracruz                                 | Pps     |
| vest-agder                               | norway                           | vest-agder                               | Pps     |
| vestfold                                 | norway                           | vestfold                                 | Pps     |
| viana_do_castelo                         | portugal                         | viana_do_castelo                         | Pps     |
| victoria                                 | australia                        | victoria                                 | Pps     |
| vienne                                   | poitou_charentes                 | vienne                                   | Pps     |
| vila_real                                | portugal                         | vila_real                                | Pps     |
| vinnytsia_oblast                         | ukraine                          | vinnytsia_oblast                         | Pps     |
| virginia                                 | us-south                         | virginia                                 | Pps     |
| viseu                                    | portugal                         | viseu                                    | Pps     |
| vizcaya                                  | euskadi                          | vizcaya                                  | Pps     |
| vladimir_oblast                          | central_federal_district         | vladimir_oblast                          | Pps     |
| volga_federal_district                   | russia                           | volga_federal_district                   | Pps     |
| volgograd_oblast                         | southern_federal_district        | volgograd_oblast                         | Pps     |
| vologda_oblast                           | northwestern_federal_district    | vologda_oblast                           | Pps     |
| volyn_oblast                             | ukraine                          | volyn_oblast                             | Pps     |
| vorarlberg                               | austria                          | vorarlberg                               | Pps     |
| voronezh_oblast                          | central_federal_district         | voronezh_oblast                          | Pps     |
| vosges                                   | lorraine                         | vosges                                   | Pps     |
| vysocina                                 | czech_republic                   | vysocina                                 | Pps     |
| wallis_et_futuna                         | oceania                          | wallis_et_futuna                         | Pps     |
| wallonia_french_community                | belgium                          | wallonia_french_community                | Pps     |
| wallonia_german_community                | belgium                          | wallonia_german_community                | Pps     |
| warminsko_mazurskie                      | poland                           | warminsko_mazurskie                      | Pps     |
| west_bengal                              | india                            | west_bengal                              | Pps     |
| west_flanders                            | flanders                         | west_flanders                            | Pps     |
| west_java                                | indonesia                        | west_java                                | Pps     |
| west_kalimantan                          | indonesia                        | west_kalimantan                          | Pps     |
| west_midlands                            | england                          | west_midlands                            | Pps     |
| west_nusa_tenggara                       | indonesia                        | west_nusa_tenggara                       | Pps     |
| west_papua                               | indonesia                        | west_papua                               | Pps     |
| west_sulawesi                            | indonesia                        | west_sulawesi                            | Pps     |
| west_sumatra                             | indonesia                        | west_sumatra                             | Pps     |
| western_australia                        | australia                        | western_australia                        | Pps     |
| western_cape                             | south_africa                     | western_cape                             | Pps     |
| western_new_york                         | new-york                         | western_new_york                         | Pps     |
| western_sahara                           | africa                           | western_sahara                           | Pps     |
| western_visayas                          | philippines                      | western_visayas                          | Pps     |
| wielkopolskie                            | poland                           | wielkopolskie                            | Pps     |
| wien                                     | austria                          | wien                                     | Pps     |
| wytheville                               | virginia                         | wytheville                               | Pps     |
| xinjiang                                 | china                            | xinjiang                                 | Pps     |
| yamalo_nenets_autonomous_okrug           | ural_federal_district            | yamalo_nenets_autonomous_okrug           | Pps     |
| yaroslavl_oblast                         | central_federal_district         | yaroslavl_oblast                         | Pps     |
| yogyakarta                               | indonesia                        | yogyakarta                               | Pps     |
| yolo                                     | california                       | yolo                                     | Pps     |
| yonne                                    | bourgogne                        | yonne                                    | Pps     |
| yorkshire_and_the_humber                 | england                          | yorkshire_and_the_humber                 | Pps     |
| yuba                                     | california                       | yuba                                     | Pps     |
| yucatan                                  | mexico                           | yucatan                                  | Pps     |
| yukon                                    | canada                           | yukon                                    | Pps     |
| yunnan                                   | china                            | yunnan                                   | Pps     |
| yvelines                                 | ile_de_france                    | yvelines                                 | Pps     |
| zabaykalsky_krai                         | siberian_federal_district        | zabaykalsky_krai                         | Pps     |
| zacatecas                                | mexico                           | zacatecas                                | Pps     |
| zachodniopomorskie                       | poland                           | zachodniopomorskie                       | Pps     |
| zakarpattia_oblast                       | ukraine                          | zakarpattia_oblast                       | Pps     |
| zambezia                                 | mozambique                       | zambezia                                 | Pps     |
| zambia                                   | africa                           | zambia                                   | Pps     |
| zamboanga_peninsula                      | philippines                      | zamboanga_peninsula                      | Pps     |
| zamora                                   | castilla_y_leon                  | zamora                                   | Pps     |
| zanzibar                                 | tanzania                         | zanzibar                                 | Pps     |
| zaporizhia_oblast                        | ukraine                          | zaporizhia_oblast                        | Pps     |
| zaragoza                                 | aragon                           | zaragoza                                 | Pps     |
| zeeland                                  | netherlands                      | zeeland                                  | Pps     |
| zhejiang                                 | china                            | zhejiang                                 | Pps     |
| zhytomyr_oblast                          | ukraine                          | zhytomyr_oblast                          | Pps     |
| zilinsky                                 | slovakia                         | zilinsky                                 | Pps     |
| zimbabwe                                 | africa                           | zimbabwe                                 | Pps     |
| zlinsky                                  | czech_republic                   | zlinsky                                  | Pps     |
| zug                                      | switzerland                      | zug                                      | Pps     |
| zuid_holland                             | netherlands                      | zuid_holland                             | Pps     |
| zurich                                   | switzerland                      | zurich                                   | Pps     |
Total elements: 1196

## List of elements from bbbike.org
|    SHORTNAME     | IS IN |    LONG NAME     |   FORMATS    |
|------------------|-------|------------------|--------------|
| Aachen           |       | Aachen           | CgMrtoOGmPpS |
| Aarhus           |       | Aarhus           | CgMrtoOGmPpS |
| Adelaide         |       | Adelaide         | CgMrtoOGmPpS |
| Albuquerque      |       | Albuquerque      | CgMrtoOGmPpS |
| Alexandria       |       | Alexandria       | CgMrtoOGmPpS |
| Amsterdam        |       | Amsterdam        | CgMrtoOGmPpS |
| Antwerpen        |       | Antwerpen        | CgMrtoOGmPpS |
| Arnhem           |       | Arnhem           | CgMrtoOGmPpS |
| Auckland         |       | Auckland         | CgMrtoOGmPpS |
| Augsburg         |       | Augsburg         | CgMrtoOGmPpS |
| Austin           |       | Austin           | CgMrtoOGmPpS |
| Baghdad          |       | Baghdad          | CgMrtoOGmPpS |
| Baku             |       | Baku             | CgMrtoOGmPpS |
| Balaton          |       | Balaton          | CgMrtoOGmPpS |
| Bamberg          |       | Bamberg          | CgMrtoOGmPpS |
| Bangkok          |       | Bangkok          | CgMrtoOGmPpS |
| Barcelona        |       | Barcelona        | CgMrtoOGmPpS |
| Basel            |       | Basel            | CgMrtoOGmPpS |
| Beijing          |       | Beijing          | CgMrtoOGmPpS |
| Beirut           |       | Beirut           | CgMrtoOGmPpS |
| Berkeley         |       | Berkeley         | CgMrtoOGmPpS |
| Berlin           |       | Berlin           | CgMrtoOGmPpS |
| Bern             |       | Bern             | CgMrtoOGmPpS |
| Bielefeld        |       | Bielefeld        | CgMrtoOGmPpS |
| Birmingham       |       | Birmingham       | CgMrtoOGmPpS |
| Bochum           |       | Bochum           | CgMrtoOGmPpS |
| Bogota           |       | Bogota           | CgMrtoOGmPpS |
| Bombay           |       | Bombay           | CgMrtoOGmPpS |
| Bonn             |       | Bonn             | CgMrtoOGmPpS |
| Bordeaux         |       | Bordeaux         | CgMrtoOGmPpS |
| Boulder          |       | Boulder          | CgMrtoOGmPpS |
| BrandenburgHavel |       | BrandenburgHavel | CgMrtoOGmPpS |
| Braunschweig     |       | Braunschweig     | CgMrtoOGmPpS |
| Bremen           |       | Bremen           | CgMrtoOGmPpS |
| Bremerhaven      |       | Bremerhaven      | CgMrtoOGmPpS |
| Brisbane         |       | Brisbane         | CgMrtoOGmPpS |
| Bristol          |       | Bristol          | CgMrtoOGmPpS |
| Brno             |       | Brno             | CgMrtoOGmPpS |
| Bruegge          |       | Bruegge          | CgMrtoOGmPpS |
| Bruessel         |       | Bruessel         | CgMrtoOGmPpS |
| Budapest         |       | Budapest         | CgMrtoOGmPpS |
| BuenosAires      |       | BuenosAires      | CgMrtoOGmPpS |
| Cairo            |       | Cairo            | CgMrtoOGmPpS |
| Calgary          |       | Calgary          | CgMrtoOGmPpS |
| Cambridge        |       | Cambridge        | CgMrtoOGmPpS |
| CambridgeMa      |       | CambridgeMa      | CgMrtoOGmPpS |
| Canberra         |       | Canberra         | CgMrtoOGmPpS |
| CapeTown         |       | CapeTown         | CgMrtoOGmPpS |
| Chemnitz         |       | Chemnitz         | CgMrtoOGmPpS |
| Chicago          |       | Chicago          | CgMrtoOGmPpS |
| ClermontFerrand  |       | ClermontFerrand  | CgMrtoOGmPpS |
| Colmar           |       | Colmar           | CgMrtoOGmPpS |
| Copenhagen       |       | Copenhagen       | CgMrtoOGmPpS |
| Cork             |       | Cork             | CgMrtoOGmPpS |
| Corsica          |       | Corsica          | CgMrtoOGmPpS |
| Corvallis        |       | Corvallis        | CgMrtoOGmPpS |
| Cottbus          |       | Cottbus          | CgMrtoOGmPpS |
| Cracow           |       | Cracow           | CgMrtoOGmPpS |
| CraterLake       |       | CraterLake       | CgMrtoOGmPpS |
| Curitiba         |       | Curitiba         | CgMrtoOGmPpS |
| Cusco            |       | Cusco            | CgMrtoOGmPpS |
| Dallas           |       | Dallas           | CgMrtoOGmPpS |
| Darmstadt        |       | Darmstadt        | CgMrtoOGmPpS |
| Davis            |       | Davis            | CgMrtoOGmPpS |
| DenHaag          |       | DenHaag          | CgMrtoOGmPpS |
| Denver           |       | Denver           | CgMrtoOGmPpS |
| Dessau           |       | Dessau           | CgMrtoOGmPpS |
| Dortmund         |       | Dortmund         | CgMrtoOGmPpS |
| Dresden          |       | Dresden          | CgMrtoOGmPpS |
| Dublin           |       | Dublin           | CgMrtoOGmPpS |
| Duesseldorf      |       | Duesseldorf      | CgMrtoOGmPpS |
| Duisburg         |       | Duisburg         | CgMrtoOGmPpS |
| Edinburgh        |       | Edinburgh        | CgMrtoOGmPpS |
| Eindhoven        |       | Eindhoven        | CgMrtoOGmPpS |
| Emden            |       | Emden            | CgMrtoOGmPpS |
| Erfurt           |       | Erfurt           | CgMrtoOGmPpS |
| Erlangen         |       | Erlangen         | CgMrtoOGmPpS |
| Eugene           |       | Eugene           | CgMrtoOGmPpS |
| Flensburg        |       | Flensburg        | CgMrtoOGmPpS |
| FortCollins      |       | FortCollins      | CgMrtoOGmPpS |
| Frankfurt        |       | Frankfurt        | CgMrtoOGmPpS |
| FrankfurtOder    |       | FrankfurtOder    | CgMrtoOGmPpS |
| Freiburg         |       | Freiburg         | CgMrtoOGmPpS |
| Gdansk           |       | Gdansk           | CgMrtoOGmPpS |
| Genf             |       | Genf             | CgMrtoOGmPpS |
| Gent             |       | Gent             | CgMrtoOGmPpS |
| Gera             |       | Gera             | CgMrtoOGmPpS |
| Glasgow          |       | Glasgow          | CgMrtoOGmPpS |
| Gliwice          |       | Gliwice          | CgMrtoOGmPpS |
| Goerlitz         |       | Goerlitz         | CgMrtoOGmPpS |
| Goeteborg        |       | Goeteborg        | CgMrtoOGmPpS |
| Goettingen       |       | Goettingen       | CgMrtoOGmPpS |
| Graz             |       | Graz             | CgMrtoOGmPpS |
| Groningen        |       | Groningen        | CgMrtoOGmPpS |
| Halifax          |       | Halifax          | CgMrtoOGmPpS |
| Halle            |       | Halle            | CgMrtoOGmPpS |
| Hamburg          |       | Hamburg          | CgMrtoOGmPpS |
| Hamm             |       | Hamm             | CgMrtoOGmPpS |
| Hannover         |       | Hannover         | CgMrtoOGmPpS |
| Heilbronn        |       | Heilbronn        | CgMrtoOGmPpS |
| Helsinki         |       | Helsinki         | CgMrtoOGmPpS |
| Hertogenbosch    |       | Hertogenbosch    | CgMrtoOGmPpS |
| Huntsville       |       | Huntsville       | CgMrtoOGmPpS |
| Innsbruck        |       | Innsbruck        | CgMrtoOGmPpS |
| Istanbul         |       | Istanbul         | CgMrtoOGmPpS |
| Jena             |       | Jena             | CgMrtoOGmPpS |
| Jerusalem        |       | Jerusalem        | CgMrtoOGmPpS |
| Johannesburg     |       | Johannesburg     | CgMrtoOGmPpS |
| Kaiserslautern   |       | Kaiserslautern   | CgMrtoOGmPpS |
| Karlsruhe        |       | Karlsruhe        | CgMrtoOGmPpS |
| Kassel           |       | Kassel           | CgMrtoOGmPpS |
| Katowice         |       | Katowice         | CgMrtoOGmPpS |
| Kaunas           |       | Kaunas           | CgMrtoOGmPpS |
| Kiel             |       | Kiel             | CgMrtoOGmPpS |
| Kiew             |       | Kiew             | CgMrtoOGmPpS |
| Koblenz          |       | Koblenz          | CgMrtoOGmPpS |
| Koeln            |       | Koeln            | CgMrtoOGmPpS |
| Konstanz         |       | Konstanz         | CgMrtoOGmPpS |
| LaPaz            |       | LaPaz            | CgMrtoOGmPpS |
| LaPlata          |       | LaPlata          | CgMrtoOGmPpS |
| LakeGarda        |       | LakeGarda        | CgMrtoOGmPpS |
| Lausanne         |       | Lausanne         | CgMrtoOGmPpS |
| Leeds            |       | Leeds            | CgMrtoOGmPpS |
| Leipzig          |       | Leipzig          | CgMrtoOGmPpS |
| Lima             |       | Lima             | CgMrtoOGmPpS |
| Linz             |       | Linz             | CgMrtoOGmPpS |
| Lisbon           |       | Lisbon           | CgMrtoOGmPpS |
| Liverpool        |       | Liverpool        | CgMrtoOGmPpS |
| Ljubljana        |       | Ljubljana        | CgMrtoOGmPpS |
| Lodz             |       | Lodz             | CgMrtoOGmPpS |
| London           |       | London           | CgMrtoOGmPpS |
| LosAngeles       |       | LosAngeles       | CgMrtoOGmPpS |
| Luebeck          |       | Luebeck          | CgMrtoOGmPpS |
| Luxemburg        |       | Luxemburg        | CgMrtoOGmPpS |
| Lyon             |       | Lyon             | CgMrtoOGmPpS |
| Maastricht       |       | Maastricht       | CgMrtoOGmPpS |
| Madison          |       | Madison          | CgMrtoOGmPpS |
| Madrid           |       | Madrid           | CgMrtoOGmPpS |
| Magdeburg        |       | Magdeburg        | CgMrtoOGmPpS |
| Mainz            |       | Mainz            | CgMrtoOGmPpS |
| Malmoe           |       | Malmoe           | CgMrtoOGmPpS |
| Manchester       |       | Manchester       | CgMrtoOGmPpS |
| Mannheim         |       | Mannheim         | CgMrtoOGmPpS |
| Marseille        |       | Marseille        | CgMrtoOGmPpS |
| Melbourne        |       | Melbourne        | CgMrtoOGmPpS |
| Memphis          |       | Memphis          | CgMrtoOGmPpS |
| MexicoCity       |       | MexicoCity       | CgMrtoOGmPpS |
| Miami            |       | Miami            | CgMrtoOGmPpS |
| Minsk            |       | Minsk            | CgMrtoOGmPpS |
| Moenchengladbach |       | Moenchengladbach | CgMrtoOGmPpS |
| Montevideo       |       | Montevideo       | CgMrtoOGmPpS |
| Montpellier      |       | Montpellier      | CgMrtoOGmPpS |
| Montreal         |       | Montreal         | CgMrtoOGmPpS |
| Moscow           |       | Moscow           | CgMrtoOGmPpS |
| Muenchen         |       | Muenchen         | CgMrtoOGmPpS |
| Muenster         |       | Muenster         | CgMrtoOGmPpS |
| NewDelhi         |       | NewDelhi         | CgMrtoOGmPpS |
| NewOrleans       |       | NewOrleans       | CgMrtoOGmPpS |
| NewYork          |       | NewYork          | CgMrtoOGmPpS |
| Nuernberg        |       | Nuernberg        | CgMrtoOGmPpS |
| Oldenburg        |       | Oldenburg        | CgMrtoOGmPpS |
| Oranienburg      |       | Oranienburg      | CgMrtoOGmPpS |
| Orlando          |       | Orlando          | CgMrtoOGmPpS |
| Oslo             |       | Oslo             | CgMrtoOGmPpS |
| Osnabrueck       |       | Osnabrueck       | CgMrtoOGmPpS |
| Ostrava          |       | Ostrava          | CgMrtoOGmPpS |
| Ottawa           |       | Ottawa           | CgMrtoOGmPpS |
| Paderborn        |       | Paderborn        | CgMrtoOGmPpS |
| Palma            |       | Palma            | CgMrtoOGmPpS |
| PaloAlto         |       | PaloAlto         | CgMrtoOGmPpS |
| Paris            |       | Paris            | CgMrtoOGmPpS |
| Perth            |       | Perth            | CgMrtoOGmPpS |
| Philadelphia     |       | Philadelphia     | CgMrtoOGmPpS |
| PhnomPenh        |       | PhnomPenh        | CgMrtoOGmPpS |
| Portland         |       | Portland         | CgMrtoOGmPpS |
| PortlandME       |       | PortlandME       | CgMrtoOGmPpS |
| Porto            |       | Porto            | CgMrtoOGmPpS |
| PortoAlegre      |       | PortoAlegre      | CgMrtoOGmPpS |
| Potsdam          |       | Potsdam          | CgMrtoOGmPpS |
| Poznan           |       | Poznan           | CgMrtoOGmPpS |
| Prag             |       | Prag             | CgMrtoOGmPpS |
| Providence       |       | Providence       | CgMrtoOGmPpS |
| Regensburg       |       | Regensburg       | CgMrtoOGmPpS |
| Riga             |       | Riga             | CgMrtoOGmPpS |
| RiodeJaneiro     |       | RiodeJaneiro     | CgMrtoOGmPpS |
| Rostock          |       | Rostock          | CgMrtoOGmPpS |
| Rotterdam        |       | Rotterdam        | CgMrtoOGmPpS |
| Ruegen           |       | Ruegen           | CgMrtoOGmPpS |
| Saarbruecken     |       | Saarbruecken     | CgMrtoOGmPpS |
| Sacramento       |       | Sacramento       | CgMrtoOGmPpS |
| Saigon           |       | Saigon           | CgMrtoOGmPpS |
| Salzburg         |       | Salzburg         | CgMrtoOGmPpS |
| SanFrancisco     |       | SanFrancisco     | CgMrtoOGmPpS |
| SanJose          |       | SanJose          | CgMrtoOGmPpS |
| SanktPetersburg  |       | SanktPetersburg  | CgMrtoOGmPpS |
| SantaBarbara     |       | SantaBarbara     | CgMrtoOGmPpS |
| SantaCruz        |       | SantaCruz        | CgMrtoOGmPpS |
| Santiago         |       | Santiago         | CgMrtoOGmPpS |
| Sarajewo         |       | Sarajewo         | CgMrtoOGmPpS |
| Schwerin         |       | Schwerin         | CgMrtoOGmPpS |
| Seattle          |       | Seattle          | CgMrtoOGmPpS |
| Seoul            |       | Seoul            | CgMrtoOGmPpS |
| Sheffield        |       | Sheffield        | CgMrtoOGmPpS |
| Singapore        |       | Singapore        | CgMrtoOGmPpS |
| Sofia            |       | Sofia            | CgMrtoOGmPpS |
| Stockholm        |       | Stockholm        | CgMrtoOGmPpS |
| Stockton         |       | Stockton         | CgMrtoOGmPpS |
| Strassburg       |       | Strassburg       | CgMrtoOGmPpS |
| Stuttgart        |       | Stuttgart        | CgMrtoOGmPpS |
| Sucre            |       | Sucre            | CgMrtoOGmPpS |
| Sydney           |       | Sydney           | CgMrtoOGmPpS |
| Szczecin         |       | Szczecin         | CgMrtoOGmPpS |
| Tallinn          |       | Tallinn          | CgMrtoOGmPpS |
| Tehran           |       | Tehran           | CgMrtoOGmPpS |
| Tilburg          |       | Tilburg          | CgMrtoOGmPpS |
| Tokyo            |       | Tokyo            | CgMrtoOGmPpS |
| Toronto          |       | Toronto          | CgMrtoOGmPpS |
| Toulouse         |       | Toulouse         | CgMrtoOGmPpS |
| Trondheim        |       | Trondheim        | CgMrtoOGmPpS |
| Tucson           |       | Tucson           | CgMrtoOGmPpS |
| Turin            |       | Turin            | CgMrtoOGmPpS |
| UlanBator        |       | UlanBator        | CgMrtoOGmPpS |
| Ulm              |       | Ulm              | CgMrtoOGmPpS |
| Usedom           |       | Usedom           | CgMrtoOGmPpS |
| Utrecht          |       | Utrecht          | CgMrtoOGmPpS |
| Vancouver        |       | Vancouver        | CgMrtoOGmPpS |
| Victoria         |       | Victoria         | CgMrtoOGmPpS |
| WarenMueritz     |       | WarenMueritz     | CgMrtoOGmPpS |
| Warsaw           |       | Warsaw           | CgMrtoOGmPpS |
| WashingtonDC     |       | WashingtonDC     | CgMrtoOGmPpS |
| Waterloo         |       | Waterloo         | CgMrtoOGmPpS |
| Wien             |       | Wien             | CgMrtoOGmPpS |
| Wroclaw          |       | Wroclaw          | CgMrtoOGmPpS |
| Wuerzburg        |       | Wuerzburg        | CgMrtoOGmPpS |
| Wuppertal        |       | Wuppertal        | CgMrtoOGmPpS |
| Zagreb           |       | Zagreb           | CgMrtoOGmPpS |
| Zuerich          |       | Zuerich          | CgMrtoOGmPpS |
Total elements: 237

## List of elements from osmtoday
|                SHORTNAME                 |             IS IN              |           LONG NAME            | FORMATS |
|------------------------------------------|--------------------------------|--------------------------------|---------|
|                                          |                                |                                | gPp     |
| aargau                                   | Switzerland                    | Aargau                         | gPp     |
| abitibi_temiscamingue                    | Quebec                         | Abitibi - Temiskaming          | gPp     |
| abruzzo                                  | Italy                          | Abruzzo                        | gPp     |
| aceh                                     | Indonesia                      | Aceh                           | gPp     |
| acre                                     | North                          | Acre                           | gPp     |
| adygea_republic                          | Southern Federal District      | Republic of Adygea             | gPp     |
| afghanistan                              | Asia                           | Afghanistan                    | gPp     |
| africa                                   |                                | Africa                         | gPp     |
| aguascalientes                           | Mexico                         | Aguascalientes                 | gPp     |
| akershus                                 | Norway                         | Akershus                       | gPp     |
| alabama                                  | United States of America       | Alabama                        | gPp     |
| alagoas                                  | Northeast                      | Alagoas                        | gPp     |
| aland                                    | Finland                        | Åland Islands                  | gPp     |
| alaska                                   | United States of America       | Alaska                         | gPp     |
| alava                                    | Euskadi                        | Alava                          | gPp     |
| albacete                                 | Castile-La Mancha              | Albacete                       | gPp     |
| alberta                                  | Canada                         | Alberta                        | gPp     |
| algeria                                  | Africa                         | Algeria                        | gPp     |
| alicante                                 | Valencia                       | Alicante                       | gPp     |
| almeria                                  | Andalusia                      | Almeria                        | gPp     |
| alsace                                   | France                         | Alsace                         | gPp     |
| altai_krai                               | Siberian Federal District      | Altai region                   | gPp     |
| altai_republic                           | Siberian Federal District      | Altai Republic                 | gPp     |
| amapa                                    | North                          | Amapa                          | gPp     |
| amazonas                                 | North                          | Amazons                        | gPp     |
| american-oceania                         | Australia Oceania              | America-Oceania                | gPp     |
| amur_oblast                              | Far Eastern Federal District   | Amur region                    | gPp     |
| andalucia                                | Spain                          | Andalusia                      | gPp     |
| andaman_and_nicobar_islands              | India                          | Andaman and Nicobar Islands    | gPp     |
| andhra_pradesh                           | India                          | Andhra Pradesh                 | gPp     |
| angola                                   | Africa                         | Angola                         | gPp     |
| anguilla                                 | Central America                | Anguilla                       | gPp     |
| anhui                                    | China                          | anhui                          | gPp     |
| antarctica                               |                                | Antarctica                     | gPp     |
| antigua_and_barbuda                      | Central America                | Antigua and Barbuda            | gPp     |
| antwerp                                  | Flanders                       | Antwerp                        | gPp     |
| appenzell_ausserrhoden                   | Switzerland                    | Appenzell Ausserrhoden         | gPp     |
| appenzell_innerrhoden                    | Switzerland                    | Appenzell-Innerrhoden          | gPp     |
| aquitaine                                | France                         | Aquitaine                      | gPp     |
| aragon                                   | Spain                          | Aragon                         | gPp     |
| argentina                                | South America                  | Argentina                      | gPp     |
| arizona                                  | United States of America       | Arizona                        | gPp     |
| arkansas                                 | United States of America       | Arkansas                       | gPp     |
| arkhangelsk_oblast                       | Northwestern Federal District  | Arkhangelsk region             | gPp     |
| armenia                                  | Asia                           | Armenia                        | gPp     |
| arnsberg                                 | North Rhine Westphalia         | Arnsberg                       | gPp     |
| aruba                                    | Central America                | Aruba                          | gPp     |
| arunachal_pradesh                        | India                          | Arunachal Pradesh              | gPp     |
| asia                                     |                                | Asia                           | gPp     |
| assam                                    | India                          | Assam                          | gPp     |
| astrakhan_oblast                         | Southern Federal District      | Astrakhan Region               | gPp     |
| asturias                                 | Spain                          | Asturias                       | gPp     |
| aust-agder                               | Norway                         | Aust-Agder                     | gPp     |
| australia                                | Australia Oceania              | Australia                      | gPp     |
| australia_oceania                        |                                | Australia Oceania              | gPp     |
| austria                                  | Europe                         | Austria                        | gPp     |
| auvergne                                 | France                         | Auvergne                       | gPp     |
| avila                                    | Castile and Leon               | Avila                          | gPp     |
| azerbaijan                               | Asia                           | Azerbaijan                     | gPp     |
| azores                                   | Portugal                       | Azores                         | gPp     |
| badajoz                                  | Extremadura                    | Badajoz                        | gPp     |
| baden_wuerttemberg                       | Germany                        | Baden Württemberg              | gPp     |
| bahamas                                  | Central America                | Bahamas                        | gPp     |
| bahia                                    | Northeast                      | Bahia                          | gPp     |
| bahrain                                  | Asia                           | Bahrain                        | gPp     |
| baja_california                          | Mexico                         | Baja California                | gPp     |
| baja_california_sur                      | Mexico                         | Baja California Sur            | gPp     |
| bali                                     | Indonesia                      | Bali                           | gPp     |
| bangka_belitung_islands                  | Indonesia                      | Bangka Belitung Islands        | gPp     |
| bangladesh                               | Asia                           | Bangladesh                     | gPp     |
| banskobystricky                          | Slovakia                       | Banskofast                     | gPp     |
| banten                                   | Indonesia                      | banten                         | gPp     |
| barbados                                 | Central America                | Barbados                       | gPp     |
| barcelona                                | Catalonia                      | Barcelona                      | gPp     |
| bas_saint_laurent                        | Quebec                         | Lower Saint Lawrence           | gPp     |
| basel_landschaft                         | Switzerland                    | Basel Land                     | gPp     |
| basel_stadt                              | Switzerland                    | Basel-Stadt                    | gPp     |
| bashkortostan_republic                   | Volga Federal District         | Republic of Bashkortostan      | gPp     |
| basilicata                               | Italy                          | Basilicata                     | gPp     |
| basse_normandie                          | France                         | Lower Normandy                 | gPp     |
| bayern                                   | Germany                        | Bavaria                        | gPp     |
| bedfordshire                             | England                        | Bedfordshire                   | gPp     |
| beijing                                  | China                          | Beijing                        | gPp     |
| belarus                                  | Europe                         | Belarus                        | gPp     |
| belgium                                  | Europe                         | Belgium                        | gPp     |
| belgorod_oblast                          | Central Federal District       | Belgorod region                | gPp     |
| belize                                   | Central America                | Belize                         | gPp     |
| bengkulu                                 | Indonesia                      | Bengkulu                       | gPp     |
| benin                                    | Africa                         | Benin                          | gPp     |
| berkshire                                | England                        | Berkshire                      | gPp     |
| berlin                                   | Germany                        | Berlin                         | gPp     |
| bermuda                                  | North America                  | Bermuda                        | gPp     |
| bern                                     | Switzerland                    | Berne                          | gPp     |
| bhutan                                   | Asia                           | Butane                         | gPp     |
| bihar                                    | India                          | Bihar                          | gPp     |
| blekinge                                 | Sweden                         | Blekinge                       | gPp     |
| bolivia                                  | South America                  | Bolivia                        | gPp     |
| botswana                                 | Africa                         | Botswana                       | gPp     |
| bourgogne                                | France                         | Burgundian                     | gPp     |
| bouvet_island                            | Africa                         | Bouvet Island                  | gPp     |
| brandenburg                              | Germany                        | Brandenburg                    | gPp     |
| bratislavsky                             | Slovakia                       | Bratislava                     | gPp     |
| brazil                                   | South America                  | Brazil                         | gPp     |
| bremen                                   | Germany                        | Bremen                         | gPp     |
| brestakaya                               | Belarus                        | Brest region                   | gPp     |
| bretagne                                 | France                         | Brittany                       | gPp     |
| bristol                                  | England                        | Bristol                        | gPp     |
| british_columbia                         | Canada                         | British Columbia               | gPp     |
| british_indian_ocean_territory           | Asia                           | British Indian Ocean Territory | gPp     |
| british_virgin_islands                   | Central America                | British Virgin Islands         | gPp     |
| brunei                                   | Asia                           | Brunei                         | gPp     |
| brussels_capital_region                  | Belgium                        | Brussels Capital Region        | gPp     |
| bryansk_oblast                           | Central Federal District       | Bryansk Region                 | gPp     |
| buckinghamshire                          | England                        | Buckinghamshire                | gPp     |
| buenos_aires                             | Argentina                      | Buenos Aires                   | gPp     |
| buenos_aires_city                        | Argentina                      | city of Buenos Aires           | gPp     |
| burgenland                               | Austria                        | Burgenland                     | gPp     |
| burgos                                   | Castile and Leon               | Burgos                         | gPp     |
| burkina_faso                             | Africa                         | Burkina Faso                   | gPp     |
| burundi                                  | Africa                         | Burundi                        | gPp     |
| buryatia_republic                        | Siberian Federal District      | The Republic of Buryatia       | gPp     |
| buskerud                                 | Norway                         | Buskerud                       | gPp     |
| caceres                                  | Extremadura                    | Caceres                        | gPp     |
| cadiz                                    | Andalusia                      | Cadiz                          | gPp     |
| calabria                                 | Italy                          | Calabria                       | gPp     |
| california                               | United States of America       | California                     | gPp     |
| cambodia                                 | Asia                           | Cambodia                       | gPp     |
| cambridgeshire                           | England                        | Cambridgeshire                 | gPp     |
| cameroon                                 | Africa                         | Cameroon                       | gPp     |
| campania                                 | Italy                          | Campaign                       | gPp     |
| campeche                                 | Mexico                         | campeche                       | gPp     |
| canada                                   | North America                  | Canada                         | gPp     |
| canary-islands                           | Africa                         | Canary Islands                 | gPp     |
| cantabria                                | Spain                          | Cantabria                      | gPp     |
| cape_verde                               | Africa                         | Cape Verde                     | gPp     |
| capitale_nationale                       | Quebec                         | Capital Nacional               | gPp     |
| caribbean                                | Central America                | caribbean                      | gPp     |
| castellon                                | Valencia                       | Castellón                      | gPp     |
| castilla_la_mancha                       | Spain                          | Castile-La Mancha              | gPp     |
| castilla_y_leon                          | Spain                          | Castile and Leon               | gPp     |
| catalunya                                | Spain                          | Catalonia                      | gPp     |
| catamarca                                | Argentina                      | Catamarca                      | gPp     |
| cayman_islands                           | Central America                | Cayman islands                 | gPp     |
| ceara                                    | Northeast                      | Ceara                          | gPp     |
| central-west                             | Brazil                         | Central Western                | gPp     |
| central-zone                             | India                          | Central Zone                   | gPp     |
| central_african_republic                 | Africa                         | Central African Republic       | gPp     |
| central_america                          |                                | Central America                | gPp     |
| central_federal_district                 | Russia                         | Central Federal District       | gPp     |
| central_finland                          | Finland                        | Central Finland                | gPp     |
| central_java                             | Indonesia                      | Central Java                   | gPp     |
| central_kalimantan                       | Indonesia                      | Central Kalimantan             | gPp     |
| central_ontario                          | Ontario                        | Central Ontario                | gPp     |
| central_ostrobothnia                     | Finland                        | Central Ostrobothnia           | gPp     |
| central_sulawesi                         | Indonesia                      | Central Sulawesi               | gPp     |
| centre                                   | France                         | Center                         | gPp     |
| centre_du_quebec                         | Quebec                         | Central Quebec                 | gPp     |
| chaco                                    | Argentina                      | Chaco                          | gPp     |
| chad                                     | Africa                         | Chad                           | gPp     |
| champagne_ardenne                        | France                         | Champagne - Ardennes           | gPp     |
| chandigarh                               | India                          | Chandigarh                     | gPp     |
| chaudiere_appalaches                     | Quebec                         | Chaudieres - Appalachians      | gPp     |
| chechen_republic                         | North Caucasian Federal        | Chechen Republic               | gPp     |
|                                          | District                       |                                |         |
| chelyabinsk_oblast                       | Ural federal district          | Chelyabinsk region             | gPp     |
| cherkasy_oblast                          | Ukraine                        | Cherkasy region                | gPp     |
| chernihiv_oblast                         | Ukraine                        | Chernihiv region               | gPp     |
| chernivtsi_oblast                        | Ukraine                        | Chernivtsi Region              | gPp     |
| cheshire                                 | England                        | Cheshire                       | gPp     |
| chhattisgarh                             | India                          | Chhattisgarh                   | gPp     |
| chiapas                                  | Mexico                         | Chiapas                        | gPp     |
| chihuahua                                | Mexico                         | Chihuahua                      | gPp     |
| chile                                    | South America                  | Chile                          | gPp     |
| china                                    | Asia                           | China                          | gPp     |
| chongqing                                | China                          | chongqing                      | gPp     |
| chubu                                    | Japan                          | Chubu                          | gPp     |
| chubut                                   | Argentina                      | Chubut                         | gPp     |
| chugoku                                  | Japan                          | Chugoku                        | gPp     |
| chukotka_autonomous_okrug                | Far Eastern Federal District   | Chukotka Autonomous Okrug      | gPp     |
| chuvash_republic                         | Volga Federal District         | Chuvash Republic               | gPp     |
| ciudad_real                              | Castile-La Mancha              | Ciudad Real                    | gPp     |
| coahuila                                 | Mexico                         | Coahuila                       | gPp     |
| colima                                   | Mexico                         | Colima                         | gPp     |
| colombia                                 | South America                  | Colombia                       | gPp     |
| colorado                                 | United States of America       | Colorado                       | gPp     |
| comoros                                  | Africa                         | Comoros                        | gPp     |
| comunidad_de_madrid                      | Spain                          | Madrid                         | gPp     |
| comunidad_foral_de_navarra               | Spain                          | Navarre                        | gPp     |
| comunitat_valenciana                     | Spain                          | Valencia                       | gPp     |
| congo-democratic-republic                | Africa                         | Democratic Republic of the     | gPp     |
|                                          |                                | Congo                          |         |
| congo_brazzaville                        | Africa                         | Brazzaville                    | gPp     |
| congo_kinshasa                           | Africa                         | Kinshasa                       | gPp     |
| connecticut                              | United States of America       | Connecticut                    | gPp     |
| cook-islands                             | Australia Oceania              | Cook Islands                   | gPp     |
| cordoba-andalucia                        | Andalusia                      | Cordoba                        | Pp      |
| cordoba-argentina                        | Argentina                      | Cordoba                        | Pp      |
| cornwall                                 | England                        | Cornwall                       | gPp     |
| corrientes                               | Argentina                      | Corrientes                     | gPp     |
| corse                                    | France                         | Kors                           | gPp     |
| costa_rica                               | Central America                | Costa Rica                     | gPp     |
| cote_nord                                | Quebec                         | Côte Nor                       | gPp     |
| crimea                                   | Ukraine                        | Crimea                         | gPp     |
| crimea_republic                          | Southern Federal District      | Republic of Crimea             | gPp     |
| cuba                                     | Central America                | Cuba                           | gPp     |
| cuenca                                   | Castile-La Mancha              | Cuenca                         | gPp     |
| cumbria                                  | England                        | Cumbria                        | gPp     |
| curacao                                  | Central America                | Curacao                        | gPp     |
| czech_republic                           | Europe                         | Czech Republic                 | gPp     |
| dadra_and_nagar_haveli                   | India                          | Dadra and Nagar Haveli         | gPp     |
| dadra_and_nagar_haveli_and_daman_and_diu | India                          | Dadra and Nagar Haveli and     | gPp     |
|                                          |                                | Daman and Diu                  |         |
| dagestan_republic                        | North Caucasian Federal        | The Republic of Dagestan       | gPp     |
|                                          | District                       |                                |         |
| dalarna                                  | Sweden                         | Dalarna                        | gPp     |
| daman_and_diu                            | India                          | Daman and Diu                  | gPp     |
| delaware                                 | United States of America       | Delaware                       | gPp     |
| derbyshire                               | England                        | Derbyshire                     | gPp     |
| detmold                                  | North Rhine Westphalia         | Detmold                        | gPp     |
| devon                                    | England                        | Devonian                       | gPp     |
| district-of-columbia                     | United States of America       | District of Colombia           | gPp     |
| distrito-federal                         | Central Western                | Distrito Federal               | gPp     |
| djibouti                                 | Africa                         | Djibouti                       | gPp     |
| dnipropetrovsk_oblast                    | Ukraine                        | Dnipropetrovsk Region          | gPp     |
| dolnoslaskie                             | Poland                         | Dolnoslaskie                   | gPp     |
| dominica                                 | Central America                | Dominica                       | gPp     |
| dominican_republic                       | Central America                | Dominican Republic             | gPp     |
| donetsk_oblast                           | Ukraine                        | Donetsk Oblast                 | gPp     |
| dorset                                   | England                        | Dorset                         | gPp     |
| drenthe                                  | Netherlands                    | Drenthe                        | gPp     |
| durango                                  | Mexico                         | Durango                        | gPp     |
| durham                                   | England                        | Durham                         | gPp     |
| dusseldorf                               | North Rhine Westphalia         | Dusseldorf                     | gPp     |
| east                                     | England                        | East                           | gPp     |
| east_flanders                            | Flanders                       | East Flemish                   | gPp     |
| east_java                                | Indonesia                      | East Java                      | gPp     |
| east_kalimantan                          | Indonesia                      | East Kalimantan                | gPp     |
| east_midlands                            | England                        | East Midlands                  | gPp     |
| east_nusa_tenggara                       | Indonesia                      | East Nusa Tenggara             | gPp     |
| east_sussex                              | England                        | East Sussex                    | gPp     |
| east_timor                               | Asia                           | East Timor                     | gPp     |
| east_yorkshire_with_hull                 | England                        | East Riding of Yorkshire       | gPp     |
| eastern-zone                             | India                          | Eastern Zone                   | gPp     |
| eastern_ontario                          | Ontario                        | Eastern Ontario                | gPp     |
| ecuador                                  | South America                  | Ecuador                        | gPp     |
| egypt                                    | Africa                         | Egypt                          | gPp     |
| el_salvador                              | Central America                | Salvador                       | gPp     |
| emilia_romagna                           | Italy                          | Emilia-Romagna                 | gPp     |
| england                                  | United Kingdom                 | England                        | gPp     |
| entre_rios                               | Argentina                      | Entre Rios                     | gPp     |
| equatorial_guinea                        | Africa                         | Equatorial Guinea              | gPp     |
| eritrea                                  | Africa                         | Eritrea                        | gPp     |
| espirito-santo                           | Southeast                      | Duhito Santo                   | gPp     |
| essex                                    | England                        | Essex                          | gPp     |
| estrie                                   | Quebec                         | Estri                          | gPp     |
| ethiopia                                 | Africa                         | Ethiopia                       | gPp     |
| europe                                   |                                | Europe                         | gPp     |
| euskadi                                  | Spain                          | Euskadi                        | gPp     |
| extremadura                              | Spain                          | Extremadura                    | gPp     |
| falkland                                 | South America                  | Falkland                       | gPp     |
| far_eastern_federal_district             | Russia                         | Far Eastern Federal District   | gPp     |
| fiji                                     | Australia Oceania              | Fiji                           | gPp     |
| finland                                  | Europe                         | Finland                        | gPp     |
| finnmark                                 | Norway                         | Finnmark                       | gPp     |
| flanders                                 | Belgium                        | Flanders                       | gPp     |
| flemish_brabant                          | Flanders                       | Flemish Brabant                | gPp     |
| flevoland                                | Netherlands                    | Flevoland                      | gPp     |
| florida                                  | United States of America       | Florida                        | gPp     |
| formosa                                  | Argentina                      | Formosa                        | gPp     |
| france                                   | Europe                         | France                         | gPp     |
| france_taaf                              | Africa                         | French Southern and Antarctic  | gPp     |
|                                          |                                | Territories                    |         |
| franche_comte                            | France                         | franche-comté                  | gPp     |
| fribourg                                 | Switzerland                    | Friborg                        | gPp     |
| friesland                                | Netherlands                    | friesland                      | gPp     |
| friuli_venezia_giulia                    | Italy                          | Friuli Venezia Giulia          | gPp     |
| fujian                                   | China                          | Fujian                         | gPp     |
| gabon                                    | Africa                         | Gabon                          | gPp     |
| galicia                                  | Spain                          | Galicia                        | gPp     |
| gambia                                   | Africa                         | Gambia                         | gPp     |
| gansu                                    | China                          | Gansu                          | gPp     |
| gaspesie_iles_de_la_madeleine            | Quebec                         | Gaspesie - Madeleine Islands   | gPp     |
| gavleborg                                | Sweden                         | Gävleborg                      | gPp     |
| gcc_states                               | Asia                           | Cooperation Council for the    | gPp     |
|                                          |                                | Arab States of the Gulf        |         |
| gelderland                               | Netherlands                    | Gelderland                     | gPp     |
| geneva                                   | Switzerland                    | Geneva                         | gPp     |
| georgia-asia                             | Asia                           | Georgia                        | Pp      |
| georgia-us                               | United States of America       | Georgia                        | Pp      |
| germany                                  | Europe                         | Germany                        | gPp     |
| ghana                                    | Africa                         | Ghana                          | gPp     |
| gibraltar                                | Europe                         | Gibraltar                      | gPp     |
| girona                                   | Catalonia                      | Girona                         | gPp     |
| glarus                                   | Switzerland                    | Glarus                         | gPp     |
| gloucestershire                          | England                        | Gloucestershire                | gPp     |
| goa                                      | India                          | Goa                            | gPp     |
| goias                                    | Central Western                | Goiás                          | gPp     |
| golden_horseshoe                         | Ontario                        | Golden Horseshoe               | gPp     |
| gorontalo                                | Indonesia                      | Gorontalo                      | gPp     |
| gotland                                  | Sweden                         | Gotland                        | gPp     |
| granada                                  | Andalusia                      | Granada                        | gPp     |
| greater_london                           | England                        | Greater London                 | gPp     |
| greater_manchester                       | England                        | Greater Manchester             | gPp     |
| greenland                                | North America                  | Greenland                      | gPp     |
| grenada                                  | Central America                | Grenada                        | gPp     |
| grisons                                  | Switzerland                    | Grisons                        | gPp     |
| grodnenskaya                             | Belarus                        | The Grodno region              | gPp     |
| groningen                                | Netherlands                    | Groningen                      | gPp     |
| guadalajara                              | Castile-La Mancha              | Guadalajara                    | gPp     |
| guadeloupe                               | Central America                | Guadeloupe                     | gPp     |
| guanajuato                               | Mexico                         | Guanajuato                     | gPp     |
| guangdong                                | China                          | Guangdong                      | gPp     |
| guangxi                                  | China                          | Guangxi                        | gPp     |
| guatemala                                | Central America                | Guatemala                      | gPp     |
| guernesey                                | Europe                         | guernsey                       | gPp     |
| guerrero                                 | Mexico                         | Guerrero                       | gPp     |
| guinea                                   | Africa                         | Guinea                         | gPp     |
| guipuzcoa                                | Euskadi                        | Gipuzkoa                       | gPp     |
| guizhou                                  | China                          | Guizhou                        | gPp     |
| gujarat                                  | India                          | Gujarat                        | gPp     |
| guyana                                   | South America                  | Guyana                         | gPp     |
| guyane                                   | South America                  | Guyan                          | gPp     |
| guyane-france                            | France                         | Guyan                          | Pp      |
| hainan                                   | China                          | Hainan                         | gPp     |
| haiti_and_domrep                         | Central America                | Republic of Haiti and          | gPp     |
|                                          |                                | Domenican Republic             |         |
| halland                                  | Sweden                         | Halland                        | gPp     |
| hamburg                                  | Germany                        | Hamburg                        | gPp     |
| hampshire                                | England                        | Hampshire                      | gPp     |
| haryana                                  | India                          | Haryana                        | gPp     |
| haute_normandie                          | France                         | Upper Normandy                 | gPp     |
| hawaii                                   | United States of America       | Hawaii                         | gPp     |
| hebei                                    | China                          | Hebei                          | gPp     |
| hedmark                                  | Norway                         | headmark                       | gPp     |
| heilongjiang                             | China                          | heilongjiang                   | gPp     |
| henan                                    | China                          | Henan                          | gPp     |
| herefordshire                            | England                        | Herefordshire                  | gPp     |
| hertfordshire                            | England                        | Hertfordshire                  | gPp     |
| hessen                                   | Germany                        | Hesse                          | gPp     |
| hidalgo                                  | Mexico                         | Hidalgo                        | gPp     |
| himachal_pradesh                         | India                          | Himachal Pradesh               | gPp     |
| hokkaido                                 | Japan                          | Hokkaido                       | gPp     |
| homelskay                                | Belarus                        | Gomel region                   | gPp     |
| honduras                                 | Central America                | Honduras                       | gPp     |
| hong_kong                                | China                          | Hong Kong                      | gPp     |
| hordaland                                | Norway                         | Hordaland                      | gPp     |
| hubei                                    | China                          | Hubei                          | gPp     |
| huelva                                   | Andalusia                      | Huelva                         | gPp     |
| huesca                                   | Aragon                         | Huesca                         | gPp     |
| hunan                                    | China                          | Hunan                          | gPp     |
| idaho                                    | United States of America       | Idaho                          | gPp     |
| ile-de-clipperton                        | Australia Oceania              | Ile De Clipperton              | gPp     |
| ile_de_france                            | France                         | Île de France                  | gPp     |
| illes_balears                            | Spain                          | Balearic Islands               | gPp     |
| illinois                                 | United States of America       | Illinois                       | gPp     |
| india                                    | Asia                           | India                          | gPp     |
| indiana                                  | United States of America       | Indiana                        | gPp     |
| indonesia                                | Asia                           | Indonesia                      | gPp     |
| ingushetia_republic                      | North Caucasian Federal        | Republic of Ingushetia         | gPp     |
|                                          | District                       |                                |         |
| inner_mongolia                           | China                          | Inner Mongolia                 | gPp     |
| iowa                                     | United States of America       | Iowa                           | gPp     |
| iran                                     | Asia                           | Iran                           | gPp     |
| iraq                                     | Asia                           | Iraq                           | gPp     |
| ireland                                  | Europe                         | Ireland                        | gPp     |
| irkutsk_oblast                           | Siberian Federal District      | Irkutsk region                 | gPp     |
| isle_of_wight                            | England                        | Isle of Wight                  | gPp     |
| israel                                   | Asia                           | Israel                         | gPp     |
| israel_and_palestine                     | Asia                           | Israel and Palestine           | gPp     |
| italy                                    | Europe                         | Italy                          | gPp     |
| ivano-frankivsk_oblast                   | Ukraine                        | Ivano-Frankivsk region         | gPp     |
| ivanovo_oblast                           | Central Federal District       | Ivanovo Region                 | gPp     |
| ivory_coast                              | Africa                         | Ivory Coast                    | gPp     |
| jaen                                     | Andalusia                      | Haen                           | gPp     |
| jakarta                                  | Indonesia                      | Jakarta                        | gPp     |
| jalisco                                  | Mexico                         | Jalisco                        | gPp     |
| jamaica                                  | Central America                | Jamaica                        | gPp     |
| jambi                                    | Indonesia                      | Jambi                          | gPp     |
| jammu_and_kashmir                        | India                          | Jammu and Kashmir              | gPp     |
| jamtland                                 | Sweden                         | gemtland                       | gPp     |
| jan_mayen                                | Norway                         | Jan Mayen                      | gPp     |
| japan                                    | Asia                           | Japan                          | gPp     |
| java                                     | Indonesia                      | Java                           | gPp     |
| jersey                                   | Europe                         | channel islands                | gPp     |
| jewish_autonomous_oblast                 | Far Eastern Federal District   | Jewish Autonomous Region       | gPp     |
| jharkhand                                | India                          | Jharkhand                      | gPp     |
| jiangsu                                  | China                          | Jiangsu                        | gPp     |
| jiangxi                                  | China                          | Jiangxi                        | gPp     |
| jihocesky                                | Czech Republic                 | Jigocheski                     | gPp     |
| jihomoravsky                             | Czech Republic                 | Jichomoravian                  | gPp     |
| jilin                                    | China                          | Jilin                          | gPp     |
| jonkoping                                | Sweden                         | Yonkoping                      | gPp     |
| jordan                                   | Asia                           | Jordan                         | gPp     |
| jujuy                                    | Argentina                      | Jujuy                          | gPp     |
| jura                                     | Switzerland                    | Yura                           | gPp     |
| kabardino_balkar_republic                | North Caucasian Federal        | Kabardino Balkar Republic      | gPp     |
|                                          | District                       |                                |         |
| kainuu                                   | Finland                        | Kainuu                         | gPp     |
| kalimantan                               | Indonesia                      | kalimantan                     | gPp     |
| kaliningrad_oblast                       | Northwestern Federal District  | Kaliningrad region             | gPp     |
| kalmar                                   | Sweden                         | Squid                          | gPp     |
| kalmykia_republic                        | Southern Federal District      | Republic of Kalmykia           | gPp     |
| kaluga_oblast                            | Central Federal District       | Kaluga region                  | gPp     |
| kamchatka_krai                           | Far Eastern Federal District   | Kamchatka                      | gPp     |
| kansai                                   | Japan                          | Kansai                         | gPp     |
| kansas                                   | United States of America       | Kansas                         | gPp     |
| kanta_hame                               | Finland                        | Kanta-Hame                     | gPp     |
| kanto                                    | Japan                          | Kanto                          | gPp     |
| karachay_cherkess_republic               | North Caucasian Federal        | Karachaev Circassian Republic  | gPp     |
|                                          | District                       |                                |         |
| karelia_republic                         | Northwestern Federal District  | Republic of Karelia            | gPp     |
| karlovarsky                              | Czech Republic                 | Karlovy Vary                   | gPp     |
| karnataka                                | India                          | Karnataka                      | gPp     |
| karnten                                  | Austria                        | Carnten                        | gPp     |
| kazakhstan                               | Asia                           | Kazakhstan                     | gPp     |
| kemerovo_oblast                          | Siberian Federal District      | Kemerovo Region                | gPp     |
| kent                                     | England                        | Kent                           | gPp     |
| kentucky                                 | United States of America       | Kentucky                       | gPp     |
| kenya                                    | Africa                         | Kenya                          | gPp     |
| kerala                                   | India                          | Kerala                         | gPp     |
| khabarovsk_krai                          | Far Eastern Federal District   | Khabarovsk region              | gPp     |
| khakassia_republic                       | Siberian Federal District      | The Republic of Khakassia      | gPp     |
| khanty_mansi_autonomous_okrug            | Ural federal district          | Khanty-Mansi Autonomous Okrug  | gPp     |
| kharkiv_oblast                           | Ukraine                        | Kharkov region                 | gPp     |
| kherson_oblast                           | Ukraine                        | Kherson region                 | gPp     |
| khmelnytskyi_oblast                      | Ukraine                        | Khmelnitsky region             | gPp     |
| kiev                                     | Ukraine                        | Kyiv                           | gPp     |
| kiev_oblast                              | Ukraine                        | Kyiv region                    | gPp     |
| kiribati                                 | Australia Oceania              | Kiribati                       | gPp     |
| kirov_oblast                             | Volga Federal District         | Kirov region                   | gPp     |
| kirovohrad_oblast                        | Ukraine                        | Kirovograd Region              | gPp     |
| koln                                     | North Rhine Westphalia         | Column                         | gPp     |
| komi_republic                            | Northwestern Federal District  | Komi Republic                  | gPp     |
| kosicky                                  | Slovakia                       | Kositsky                       | gPp     |
| kostroma_oblast                          | Central Federal District       | Kostroma region                | gPp     |
| kralovehradecky                          | Czech Republic                 | Kralovogradetsky               | gPp     |
| krasnodar_krai                           | Southern Federal District      | Krasnodar region               | gPp     |
| krasnoyarsk_krai                         | Siberian Federal District      | Krasnoyarsk Region             | gPp     |
| kronoberg                                | Sweden                         | Kronoberg                      | gPp     |
| kujawsko_pomorskie                       | Poland                         | Kuyavia-Pomeranian             | gPp     |
| kurgan_oblast                            | Ural federal district          | Kurgan region                  | gPp     |
| kursk_oblast                             | Central Federal District       | Kursk Region                   | gPp     |
| kuwait                                   | Asia                           | Kuwait                         | gPp     |
| kymenlaakso                              | Finland                        | Kymenlaakso                    | gPp     |
| kyrgyzstan                               | Asia                           | Kyrgyzstan                     | gPp     |
| kyushu                                   | Japan                          | Kyushu                         | gPp     |
| la_coruna                                | Galicia                        | La Coruna                      | gPp     |
| la_pampa                                 | Argentina                      | La Pampa                       | gPp     |
| la_rioja-argentina                       | Argentina                      | Rioja                          | Pp      |
| la_rioja-spain                           | Spain                          | Rioja                          | Pp      |
| lakshadweep                              | India                          | Lakshadweep                    | gPp     |
| lampung                                  | Indonesia                      | Lampung                        | gPp     |
| lanaudiere                               | Quebec                         | Lanaudière                     | gPp     |
| lancashire                               | England                        | Lancashire                     | gPp     |
| languedoc_roussillon                     | France                         | Languedoc Roussillon           | gPp     |
| laos                                     | Asia                           | Laos                           | gPp     |
| lapland                                  | Finland                        | Lapland                        | gPp     |
| laurentides                              | Quebec                         | Laurentides                    | gPp     |
| laval                                    | Quebec                         | Laval                          | gPp     |
| lazio                                    | Italy                          | Lazio                          | gPp     |
| lebanon                                  | Asia                           | Lebanon                        | gPp     |
| leicestershire                           | England                        | Leicestershire                 | gPp     |
| leningrad_oblast                         | Northwestern Federal District  | Leningrad region               | gPp     |
| leon                                     | Castile and Leon               | Leon                           | gPp     |
| lesotho                                  | Africa                         | Lesotho                        | gPp     |
| liaoning                                 | China                          | Liaoning                       | gPp     |
| liberecky                                | Czech Republic                 | Liberec                        | gPp     |
| liberia                                  | Africa                         | Liberia                        | gPp     |
| libya                                    | Africa                         | Libya                          | gPp     |
| liguria                                  | Italy                          | Liguria                        | gPp     |
| limburg-flanders                         | Flanders                       | Limburg                        | Pp      |
| limburg-netherlands                      | Netherlands                    | Limburg                        | Pp      |
| limousin                                 | France                         | Limousine                      | gPp     |
| lincolnshire                             | England                        | Lincolnshire                   | gPp     |
| lipetsk_oblast                           | Central Federal District       | Lipetsk Region                 | gPp     |
| lleida                                   | Catalonia                      | Lleida                         | gPp     |
| lodzkie                                  | Poland                         | Lodz                           | gPp     |
| lombardia                                | Italy                          | Lombardy                       | gPp     |
| lorraine                                 | France                         | Lorraine                       | gPp     |
| louisiana                                | United States of America       | Louisiana                      | gPp     |
| lubelskie                                | Poland                         | Lubelskie                      | gPp     |
| lubuskie                                 | Poland                         | Lubuski                        | gPp     |
| lucerne                                  | Switzerland                    | Alfalfa                        | gPp     |
| lugo                                     | Galicia                        | Lugo                           | gPp     |
| luhansk_oblast                           | Ukraine                        | Lugansk region                 | gPp     |
| luxembourg                               | Europe                         | Luxembourg                     | gPp     |
| lviv_oblast                              | Ukraine                        | Lviv region                    | gPp     |
| macau                                    | China                          | Macau                          | gPp     |
| madagascar                               | Africa                         | Madagascar                     | gPp     |
| madeira                                  | Portugal                       | Madeira                        | gPp     |
| madhya_pradesh                           | India                          | Madhya Pradesh                 | gPp     |
| magadan_oblast                           | Far Eastern Federal District   | Magadan Region                 | gPp     |
| maharashtra                              | India                          | Maharashtra                    | gPp     |
| maine                                    | United States of America       | Maine                          | gPp     |
| malaga                                   | Andalusia                      | Malaga                         | gPp     |
| malawi                                   | Africa                         | Malawi                         | gPp     |
| malaysia                                 | Asia                           | Malaysia                       | gPp     |
| maldives                                 | Asia                           | Maldives                       | gPp     |
| mali                                     | Africa                         | Mali                           | gPp     |
| malopolskie                              | Poland                         | Lesser Poland                  | gPp     |
| maluku                                   | Indonesia                      | Maluku                         | gPp     |
| manipur                                  | India                          | Manipur                        | gPp     |
| manitoba                                 | Canada                         | Manitoba                       | gPp     |
| maranhao                                 | Northeast                      | Maranhao                       | gPp     |
| marche                                   | Italy                          | Marche                         | gPp     |
| mari_el_republic                         | Volga Federal District         | Mari El Republic               | gPp     |
| marshall-islands                         | Australia Oceania              | Marshall Islands               | gPp     |
| martinique                               | Central America                | Martinique                     | gPp     |
| maryland                                 | United States of America       | Maryland                       | gPp     |
| massachusetts                            | United States of America       | Massachusetts                  | gPp     |
| mato-grosso                              | Central Western                | Mata Grosso                    | gPp     |
| mato-grosso-do-sul                       | Central Western                | Mato Grosso Do Sul             | gPp     |
| mauricie                                 | Quebec                         | Morisi                         | gPp     |
| mauritania                               | Africa                         | Mauritania                     | gPp     |
| mauritius                                | Africa                         | Mauritius                      | gPp     |
| mayotte                                  | Africa                         | Mayotte                        | gPp     |
| mazowieckie                              | Poland                         | Masovian                       | gPp     |
| mecklenburg_vorpommern                   | Germany                        | Mecklenburg Western Pomerania  | gPp     |
| meghalaya                                | India                          | Megalaya                       | gPp     |
| mendoza                                  | Argentina                      | Mendoza                        | gPp     |
| merseyside                               | England                        | Merseyside                     | gPp     |
| mexico                                   | North America                  | Mexico                         | gPp     |
| mexico_city                              | Mexico                         | mexico city                    | gPp     |
| michigan                                 | United States of America       | Michigan                       | gPp     |
| michoacan                                | Mexico                         | Michoacán                      | gPp     |
| micronesia                               | Australia Oceania              | micronesia                     | gPp     |
| midi_pyrenees                            | France                         | Midi Pyrenees                  | gPp     |
| minas-gerais                             | Southeast                      | Minas Gerais                   | gPp     |
| minnesota                                | United States of America       | Minnesota                      | gPp     |
| minskaya                                 | Belarus                        | Minsk Region                   | gPp     |
| misiones                                 | Argentina                      | missiones                      | gPp     |
| mississippi                              | United States of America       | Mississippi                    | gPp     |
| missouri                                 | United States of America       | Missouri                       | gPp     |
| mittelfranken                            | Bavaria                        | Mittelfranken                  | gPp     |
| mizoram                                  | India                          | Mizoram                        | gPp     |
| moere_og_romsdal                         | Norway                         | Møre og Romsdal                | gPp     |
| mogilevskaya                             | Belarus                        | Mogilev region                 | gPp     |
| molise                                   | Italy                          | plea                           | gPp     |
| monaco                                   | Europe                         | Monaco                         | gPp     |
| mongolia                                 | Asia                           | Mongolia                       | gPp     |
| montana                                  | United States of America       | Montana                        | gPp     |
| monteregie                               | Quebec                         | Monteregi                      | gPp     |
| montreal                                 | Quebec                         | Montreal                       | gPp     |
| montserrat                               | Central America                | Montserrat                     | gPp     |
| moravskoslezsky                          | Czech Republic                 | Moravian-Slase                 | gPp     |
| mordovia_republic                        | Volga Federal District         | The Republic of Mordovia       | gPp     |
| morelos                                  | Mexico                         | Morelos                        | gPp     |
| morocco                                  | Africa                         | Morocco                        | gPp     |
| moscow                                   | Central Federal District       | Moscow                         | gPp     |
| moscow_oblast                            | Central Federal District       | Moscow region                  | gPp     |
| mozambique                               | Africa                         | Mozambique                     | gPp     |
| munster                                  | North Rhine Westphalia         | Munster                        | gPp     |
| murmansk_oblast                          | Northwestern Federal District  | Murmansk region                | gPp     |
| myanmar                                  | Asia                           | Myanmar                        | gPp     |
| mykolaiv_oblast                          | Ukraine                        | Mykolaiv Region                | gPp     |
| nagaland                                 | India                          | Nagaland                       | gPp     |
| namibia                                  | Africa                         | Namibia                        | gPp     |
| national_capital_territory_of_delhi      | India                          | National Capital Territory     | gPp     |
|                                          |                                | Delhi                          |         |
| nauru                                    | Australia Oceania              | Nauru                          | gPp     |
| nayarit                                  | Mexico                         | Nayarite                       | gPp     |
| nebraska                                 | United States of America       | Nebraska                       | gPp     |
| nenets_autonomous_okrug                  | Northwestern Federal District  | Nenets Autonomous Okrug        | gPp     |
| nepal                                    | Asia                           | Nepal                          | gPp     |
| netherlands                              | Europe                         | Netherlands                    | gPp     |
| neuchatel                                | Switzerland                    | Neuchâtel                      | gPp     |
| neuquen                                  | Argentina                      | Neuquen                        | gPp     |
| nevada                                   | United States of America       | Nevada                         | gPp     |
| new-caledonia                            | Australia Oceania              | New Caledonia                  | gPp     |
| new-hampshire                            | United States of America       | New Hampshire                  | gPp     |
| new-jersey                               | United States of America       | New Jersey                     | gPp     |
| new-mexico                               | United States of America       | New Mexico                     | gPp     |
| new-york                                 | United States of America       | New York                       | gPp     |
| new-zealand                              | Australia Oceania              | New Zealand                    | gPp     |
| new_brunswick                            | Canada                         | New Brunswick                  | gPp     |
| newfoundland_and_labrador                | Canada                         | Newfoundland and Labrador      | gPp     |
| nicaragua                                | Central America                | Nicaragua                      | gPp     |
| nidwalden                                | Switzerland                    | Nidwalden                      | gPp     |
| niederbayern                             | Bavaria                        | Niederbayern                   | gPp     |
| niederosterreich                         | Austria                        | Niederosterreich               | gPp     |
| niedersachsen                            | Germany                        | Lower Saxony                   | gPp     |
| niger                                    | Africa                         | Niger                          | gPp     |
| nigeria                                  | Africa                         | Nigeria                        | gPp     |
| ningxia                                  | China                          | Ningxia                        | gPp     |
| nitriansky                               | Slovakia                       | Nitryansky                     | gPp     |
| niue                                     | Australia Oceania              | Niue                           | gPp     |
| nizhny_novgorod_oblast                   | Volga Federal District         | Nizhny Novgorod Region         | gPp     |
| noord_brabant                            | Netherlands                    | North Brabant                  | gPp     |
| noord_holland                            | Netherlands                    | North Holland                  | gPp     |
| norcal                                   | California                     | California North               | gPp     |
| nord_du_quebec                           | Quebec                         | North of Quebec                | gPp     |
| nord_pas_de_calais                       | France                         | Nord - Pas de Calais           | gPp     |
| nordland                                 | Norway                         | Nordland                       | gPp     |
| nordrhein_westfalen                      | Germany                        | North Rhine Westphalia         | gPp     |
| norfolk                                  | England                        | Norfolk                        | gPp     |
| norrbotten                               | Sweden                         | Norrbotten                     | gPp     |
| north                                    | Brazil                         | North                          | gPp     |
| north-carolina                           | United States of America       | North Carolina                 | gPp     |
| north-dakota                             | United States of America       | North Dakota                   | gPp     |
| north-eastern-zone                       | India                          | North East Zone                | gPp     |
| north_america                            |                                | North America                  | gPp     |
| north_caucasian_federal_district         | Russia                         | North Caucasian Federal        | gPp     |
|                                          |                                | District                       |         |
| north_east                               | England                        | Northeast                      | gPp     |
| north_kalimantan                         | Indonesia                      | North Kalimantan               | gPp     |
| north_karelia                            | Finland                        | North Karelia                  | gPp     |
| north_korea                              | Asia                           | North Korea                    | gPp     |
| north_maluku                             | Indonesia                      | North Maluku                   | gPp     |
| north_ossetia_alania_republic            | North Caucasian Federal        | Republic of North Ossetia      | gPp     |
|                                          | District                       | Alania                         |         |
| north_ostrobothnia                       | Finland                        | North Ostrobothnia             | gPp     |
| north_savo                               | Finland                        | North Savo                     | gPp     |
| north_sulawesi                           | Indonesia                      | North Sulawesi                 | gPp     |
| north_sumatra                            | Indonesia                      | North Sumatra                  | gPp     |
| north_west                               | England                        | Northwest                      | gPp     |
| north_yorkshire                          | England                        | North Yorkshire                | gPp     |
| northamptonshire                         | England                        | Northamptonshire               | gPp     |
| northeast                                | Brazil                         | Northeast                      | gPp     |
| northeastern_ontario                     | Ontario                        | Northeastern Ontario           | gPp     |
| northern-zone                            | India                          | North Zone                     | gPp     |
| northern_ireland                         | United Kingdom                 | Northern Ireland               | gPp     |
| northumberland                           | England                        | Northumberland                 | gPp     |
| northwest_territories                    | Canada                         | Northwest Territories          | gPp     |
| northwestern_federal_district            | Russia                         | Northwestern Federal District  | gPp     |
| northwestern_ontario                     | Ontario                        | Northwestern Ontario           | gPp     |
| norway                                   | Europe                         | Norway                         | gPp     |
| nottinghamshire                          | England                        | Nottinghamshire                | gPp     |
| nova_scotia                              | Canada                         | Nova Scotia                    | gPp     |
| novgorod_oblast                          | Northwestern Federal District  | Novgorod region                | gPp     |
| novosibirsk_oblast                       | Siberian Federal District      | Novosibirsk region             | gPp     |
| nuevo_leon                               | Mexico                         | Nuevo Leon                     | gPp     |
| nunavut                                  | Canada                         | Nunavut                        | gPp     |
| nusa_tenggara                            | Indonesia                      | West Nusa Tenggara             | gPp     |
| oaxaca                                   | Mexico                         | Oaxaca                         | gPp     |
| oberbayern                               | Bavaria                        | Oberbayern                     | gPp     |
| oberfranken                              | Bavaria                        | Oberfranken                    | gPp     |
| oberosterreich                           | Austria                        | Oberosterreich                 | gPp     |
| oberpfalz                                | Bavaria                        | Oberpfalz                      | gPp     |
| obwalden                                 | Switzerland                    | obwalden                       | gPp     |
| odessa_oblast                            | Ukraine                        | Odessa region                  | gPp     |
| odisha                                   | India                          | Odisha                         | gPp     |
| oestfold                                 | Norway                         | estfold                        | gPp     |
| ohio                                     | United States of America       | Ohio                           | gPp     |
| oklahoma                                 | United States of America       | Oklahoma                       | gPp     |
| olomoucky                                | Czech Republic                 | Olomouc                        | gPp     |
| oman                                     | Asia                           | Oman                           | gPp     |
| omsk_oblast                              | Siberian Federal District      | Omsk Region                    | gPp     |
| ontario                                  | Canada                         | Ontario                        | gPp     |
| opolskie                                 | Poland                         | Opole                          | gPp     |
| oppland                                  | Norway                         | Oppland                        | gPp     |
| orebro                                   | Sweden                         | Orebro                         | gPp     |
| oregon                                   | United States of America       | Oregon                         | gPp     |
| orenburg_oblast                          | Volga Federal District         | Orenburg region                | gPp     |
| oryol_oblast                             | Central Federal District       | Oryol Region                   | gPp     |
| oslo                                     | Norway                         | Oslo                           | gPp     |
| ostergotland                             | Sweden                         | Ostergotland                   | gPp     |
| ostrobothnia                             | Finland                        | Ostrobothnia                   | gPp     |
| ourense                                  | Galicia                        | Ourense                        | gPp     |
| outaouais                                | Quebec                         | Outaouais                      | gPp     |
| overijssel                               | Netherlands                    | Overijssel                     | gPp     |
| oxfordshire                              | England                        | Oxfordshire                    | gPp     |
| paijat_hame                              | Finland                        | Payat-Hame                     | gPp     |
| pakistan                                 | Asia                           | Pakistan                       | gPp     |
| palau                                    | Australia Oceania              | Palau                          | gPp     |
| palencia                                 | Castile and Leon               | Palencia                       | gPp     |
| palestine                                | Asia                           | Palestine                      | gPp     |
| panama                                   | Central America                | Panama                         | gPp     |
| papua                                    | Indonesia                      | Papua                          | gPp     |
| papua-new-guinea                         | Australia Oceania              | Papua New Guinea               | gPp     |
| para                                     | North                          | Paragraph                      | gPp     |
| paraguay                                 | South America                  | Paraguay                       | gPp     |
| paraiba                                  | Northeast                      | Paraiba                        | gPp     |
| parana                                   | South                          | Paraná                         | gPp     |
| pardubicky                               | Czech Republic                 | Pardubice                      | gPp     |
| pays_de_la_loire                         | France                         | Pays de la Loire               | gPp     |
| pennsylvania                             | United States of America       | Pennsylvania                   | gPp     |
| penza_oblast                             | Volga Federal District         | Penza region                   | gPp     |
| perm_krai                                | Volga Federal District         | Perm region                    | gPp     |
| pernambuco                               | Northeast                      | Pernambuco                     | gPp     |
| peru                                     | South America                  | Peru                           | gPp     |
| philippines                              | Asia                           | Philippines                    | gPp     |
| piaui                                    | Northeast                      | Piaui                          | gPp     |
| picardie                                 | France                         | Picardy                        | gPp     |
| piemonte                                 | Italy                          | Piedmont                       | gPp     |
| pirkanmaa                                | Finland                        | Pirkanmaa                      | gPp     |
| pitcairn-islands                         | Australia Oceania              | Pitcairn Islands               | gPp     |
| plzensky                                 | Czech Republic                 | Pilsensky                      | gPp     |
| podkarpackie                             | Poland                         | Subcarpathia                   | gPp     |
| podlaskie                                | Poland                         | Podlaskie Voivodeship          | gPp     |
| poitou_charentes                         | France                         | Poitou Charente                | gPp     |
| poland                                   | Europe                         | Poland                         | gPp     |
| poltava_oblast                           | Ukraine                        | Poltava region                 | gPp     |
| polynesie-francaise                      | Australia Oceania              | Polynesia-France               | gPp     |
| pomorskie                                | Poland                         | Pomeranian                     | gPp     |
| pontevedra                               | Galicia                        | Pontevedra                     | gPp     |
| portugal                                 | Europe                         | Portugal                       | gPp     |
| praha                                    | Czech Republic                 | Prague                         | gPp     |
| presovsky                                | Slovakia                       | Presovsky                      | gPp     |
| primorsky_krai                           | Far Eastern Federal District   | Primorsky Krai                 | gPp     |
| prince_edward_island                     | Canada                         | Prince Edward Island           | gPp     |
| provence_alpes_cote_d_azur               | France                         | Provence - Alpes - Cote d'Azur | gPp     |
| pskov_oblast                             | Northwestern Federal District  | Pskov region                   | gPp     |
| puducherry                               | India                          | Puducherry                     | gPp     |
| puebla                                   | Mexico                         | puebla                         | gPp     |
| puerto-rico                              | United States of America       | Puerto Rico                    | gPp     |
| puerto_rico                              | Central America                | Puerto Rico                    | gPp     |
| puglia                                   | Italy                          | Apulia                         | gPp     |
| punjab                                   | India                          | Punjab                         | gPp     |
| qatar                                    | Asia                           | Qatar                          | gPp     |
| qinghai                                  | China                          | Qinghai                        | gPp     |
| quebec                                   | Canada                         | Quebec                         | gPp     |
| queretaro                                | Mexico                         | Querétaro                      | gPp     |
| quintana_roo                             | Mexico                         | Quintana Roo                   | gPp     |
| rajasthan                                | India                          | Rajasthan                      | gPp     |
| region_de_murcia                         | Spain                          | Murcia                         | gPp     |
| reunion                                  | Africa                         | reunion                        | gPp     |
| rheinland_pfalz                          | Germany                        | Rhineland Palatinate           | gPp     |
| rhode-island                             | United States of America       | Rhode Island                   | gPp     |
| rhone_alpes                              | France                         | Rhone Alps                     | gPp     |
| riau                                     | Indonesia                      | Riau                           | gPp     |
| riau_islands                             | Indonesia                      | Kepulauan Riau                 | gPp     |
| rio-de-janeiro                           | Southeast                      | Rio de Janeiro                 | gPp     |
| rio-grande-do-norte                      | Northeast                      | Rio Grande do Norte            | gPp     |
| rio-grande-do-sul                        | South                          | Rio Grande do Sul              | gPp     |
| rio_negro                                | Argentina                      | Rio Negro                      | gPp     |
| rivne_oblast                             | Ukraine                        | Rivne Region                   | gPp     |
| rogaland                                 | Norway                         | Rogaland                       | gPp     |
| rondonia                                 | North                          | Rondonia                       | gPp     |
| roraima                                  | North                          | Roraima                        | gPp     |
| rostov_oblast                            | Southern Federal District      | Rostov region                  | gPp     |
| russia                                   |                                | Russia                         | gPp     |
| rutland                                  | England                        | rutland                        | gPp     |
| rwanda                                   | Africa                         | Rwanda                         | gPp     |
| ryazan_oblast                            | Central Federal District       | Ryazan Oblast                  | gPp     |
| saarland                                 | Germany                        | Saar                           | gPp     |
| sachsen                                  | Germany                        | Sachsen                        | gPp     |
| sachsen_anhalt                           | Germany                        | Saxony-Anhalt                  | gPp     |
| saguenay_lac_saint_jean                  | Quebec                         | Saguenay - Lake Saint-Jean     | gPp     |
| saint_barthelemy                         | Central America                | Saint Barthélemy               | gPp     |
| saint_gallen                             | Switzerland                    | St. Gallen                     | gPp     |
| saint_helena_ascension_tristan_da_cunha  | Africa                         | Saint Helena                   | gPp     |
| saint_kitts_and_nevis                    | Central America                | Saint Kitts and Nevis          | gPp     |
| saint_lucia                              | Central America                | Saint Lucia                    | gPp     |
| saint_martin                             | Central America                | Saint Martin                   | gPp     |
| saint_petersburg                         | Northwestern Federal District  | St. Petersburg                 | gPp     |
| saint_pierre_et_miquelon                 | North America                  | Saint Pierre and Miquelon      | gPp     |
| saint_vincent_and_the_grenadines         | Central America                | Saint Vincent and the          | gPp     |
|                                          |                                | Grenadines                     |         |
| sakha_republic                           | Far Eastern Federal District   | Republic of Yakutia            | gPp     |
| sakhalin_oblast                          | Far Eastern Federal District   | Sakhalin Region                | gPp     |
| salamanca                                | Castile and Leon               | Salamanca                      | gPp     |
| salta                                    | Argentina                      | Salta                          | gPp     |
| salzburg                                 | Austria                        | Salzburg                       | gPp     |
| samara_oblast                            | Volga Federal District         | Samara Region                  | gPp     |
| samoa                                    | Australia Oceania              | Samoa                          | gPp     |
| san_juan                                 | Argentina                      | San Juan                       | gPp     |
| san_luis                                 | Argentina                      | San Luis                       | gPp     |
| san_luis_potosi                          | Mexico                         | San Luis Potosi                | gPp     |
| san_marino                               | Europe                         | San Marino                     | gPp     |
| santa-catarina                           | South                          | Santa Catarina                 | gPp     |
| santa_cruz                               | Argentina                      | Santa Cruz                     | gPp     |
| santa_fe                                 | Argentina                      | Santa Fe                       | gPp     |
| santiago_del_estero                      | Argentina                      | Santiago Del Estero            | gPp     |
| sao-paulo                                | Southeast                      | Sao Paulo                      | gPp     |
| sao_tome_and_principe                    | Africa                         | Sao Tome and Principe          | gPp     |
| saratov_oblast                           | Volga Federal District         | Saratov region                 | gPp     |
| sardegna                                 | Italy                          | Sardinia                       | gPp     |
| saskatchewan                             | Canada                         | Saskatchewan                   | gPp     |
| satakunta                                | Finland                        | Satakunta                      | gPp     |
| saudi_arabia                             | Asia                           | Saudi Arabia                   | gPp     |
| schaffhausen                             | Switzerland                    | Schaffhausen                   | gPp     |
| schleswig_holstein                       | Germany                        | Schleswig Holstein             | gPp     |
| schwaben                                 | Bavaria                        | Schwaben                       | gPp     |
| schwyz                                   | Switzerland                    | Schwyz                         | gPp     |
| scotland                                 | United Kingdom                 | Scotland                       | gPp     |
| segovia                                  | Castile and Leon               | Segovia                        | gPp     |
| senegal                                  | Africa                         | Senegal                        | gPp     |
| sergipe                                  | Northeast                      | Sergeep                        | gPp     |
| sevastopol                               | Southern Federal District      | Sevastopol                     | gPp     |
| sevastopol-ukraine                       | Ukraine                        | Sevastopol                     | Pp      |
| sevilla                                  | Andalusia                      | Seville                        | gPp     |
| seychelles                               | Africa                         | Seychelles                     | gPp     |
| shaanxi                                  | China                          | Shaanxi                        | gPp     |
| shandong                                 | China                          | Shandong                       | gPp     |
| shanghai                                 | China                          | Shanghai                       | gPp     |
| shanxi                                   | China                          | Shanxi                         | gPp     |
| shikoku                                  | Japan                          | Shikoku                        | gPp     |
| shropshire                               | England                        | Shropshire                     | gPp     |
| siberian_federal_district                | Russia                         | Siberian Federal District      | gPp     |
| sichuan                                  | China                          | Sichuan                        | gPp     |
| sicilia                                  | Italy                          | Sicily                         | gPp     |
| sierra-leone                             | Africa                         | Sierra Leone                   | gPp     |
| sikkim                                   | India                          | Sikkim                         | gPp     |
| sinaloa                                  | Mexico                         | Sinaloa                        | gPp     |
| singapore                                | Asia                           | Singapore                      | gPp     |
| sint_maarten                             | Central America                | Sint Maarten                   | gPp     |
| skane                                    | Sweden                         | Skane                          | gPp     |
| slaskie                                  | Poland                         | Sweet                          | gPp     |
| slovakia                                 | Europe                         | Slovakia                       | gPp     |
| smolensk_oblast                          | Central Federal District       | Smolensk region                | gPp     |
| socal                                    | California                     | California South               | gPp     |
| sodermanland                             | Sweden                         | Södermanland                   | gPp     |
| sogn_og_fjordane                         | Norway                         | Sogn-og-Fyurane                | gPp     |
| solomon-islands                          | Australia Oceania              | Solomon islands                | gPp     |
| solothurn                                | Switzerland                    | Solothurn                      | gPp     |
| somalia                                  | Africa                         | Somalia                        | gPp     |
| somerset                                 | England                        | Somerset                       | gPp     |
| sonora                                   | Mexico                         | Sonora                         | gPp     |
| soria                                    | Castile and Leon               | Soria                          | gPp     |
| south                                    | Brazil                         | South                          | gPp     |
| south-africa-and-lesotho                 | Africa                         | South Africa and Lesotho       | gPp     |
| south-carolina                           | United States of America       | South Carolina                 | gPp     |
| south-dakota                             | United States of America       | North Dakota                   | gPp     |
| south_africa                             | Africa                         | South Africa                   | gPp     |
| south_america                            |                                | South America                  | gPp     |
| south_east                               | England                        | South East                     | gPp     |
| south_georgia_and_south_sandwich         | South America                  | South Georgia And South        | gPp     |
|                                          |                                | Sandwich                       |         |
| south_kalimantan                         | Indonesia                      | South Kalimantan               | gPp     |
| south_karelia                            | Finland                        | South Karelia                  | gPp     |
| south_korea                              | Asia                           | South Korea                    | gPp     |
| south_ostrobothnia                       | Finland                        | Southern Ostrobothnia          | gPp     |
| south_savo                               | Finland                        | South Savo                     | gPp     |
| south_sudan                              | Africa                         | South Sudan                    | gPp     |
| south_sulawesi                           | Indonesia                      | South Sulawesi                 | gPp     |
| south_sumatra                            | Indonesia                      | South Sumatra                  | gPp     |
| south_west                               | England                        | South West                     | gPp     |
| south_yorkshire                          | England                        | South Yorkshire                | gPp     |
| southeast                                | Brazil                         | Southeast                      | gPp     |
| southeast_sulawesi                       | Indonesia                      | Southeast Sulawesi             | gPp     |
| southern-zone                            | India                          | South Zone                     | gPp     |
| southern_federal_district                | Russia                         | Southern Federal District      | gPp     |
| southwest_finland                        | Finland                        | Southwest Finland              | gPp     |
| southwestern_ontario                     | Ontario                        | Southwestern Ontario           | gPp     |
| spain                                    | Europe                         | Spain                          | gPp     |
| sri_lanka                                | Asia                           | Sri Lanka                      | gPp     |
| staffordshire                            | England                        | Staffordshire                  | gPp     |
| state_of_mexico                          | Mexico                         | mexico city                    | gPp     |
| stavropol_krai                           | North Caucasian Federal        | Stavropol region               | gPp     |
|                                          | District                       |                                |         |
| steiermark                               | Austria                        | Stirmarka                      | gPp     |
| stockholm                                | Sweden                         | Stockholm                      | gPp     |
| stredocesky                              | Czech Republic                 | Stredochsky                    | gPp     |
| sudan                                    | Africa                         | Sudan                          | gPp     |
| suffolk                                  | England                        | Suffolk                        | gPp     |
| sulawesi                                 | Indonesia                      | Sulawesi                       | gPp     |
| sumatra                                  | Indonesia                      | Sumatra                        | gPp     |
| sumy_oblast                              | Ukraine                        | Sumy region                    | gPp     |
| suriname                                 | South America                  | Suriname                       | gPp     |
| surrey                                   | England                        | Surrey                         | gPp     |
| svalbard                                 | Norway                         | Svalbard                       | gPp     |
| sverdlovsk_oblast                        | Ural federal district          | Sverdlovsk region              | gPp     |
| swaziland                                | Africa                         | Swaziland                      | gPp     |
| sweden                                   | Europe                         | Sweden                         | gPp     |
| swietokrzyskie                           | Poland                         | Świętokrzyskie                 | gPp     |
| switzerland                              | Europe                         | Switzerland                    | gPp     |
| syria                                    | Asia                           | Syria                          | gPp     |
| tabasco                                  | Mexico                         | Tabasco                        | gPp     |
| taiwan                                   | Asia                           | Taiwan                         | gPp     |
| tajikistan                               | Asia                           | Tajikistan                     | gPp     |
| tamaulipas                               | Mexico                         | Tamaulipas                     | gPp     |
| tambov_oblast                            | Central Federal District       | Tambov Region                  | gPp     |
| tamil_nadu                               | India                          | Tamilnadu                      | gPp     |
| tanzania                                 | Africa                         | Tanzania                       | gPp     |
| tarragona                                | Catalonia                      | Tarragona                      | gPp     |
| tatarstan_republic                       | Volga Federal District         | Republic of Tatarstan          | gPp     |
| telangana                                | India                          | Telengana                      | gPp     |
| telemark                                 | Norway                         | Telemark                       | gPp     |
| tennessee                                | United States of America       | Tennessee                      | gPp     |
| ternopil_oblast                          | Ukraine                        | Ternopil Region                | gPp     |
| teruel                                   | Aragon                         | Teruel                         | gPp     |
| texas                                    | United States of America       | Texas                          | gPp     |
| thailand                                 | Asia                           | Thailand                       | gPp     |
| thueringen                               | Germany                        | Thuringian                     | gPp     |
| thurgau                                  | Switzerland                    | Thurgau                        | gPp     |
| tianjin                                  | China                          | Tianjin                        | gPp     |
| tibet                                    | China                          | Tibet                          | gPp     |
| ticino                                   | Switzerland                    | Ticino                         | gPp     |
| tierra_del_fuego                         | Argentina                      | Tierra del Fuego               | gPp     |
| tirol                                    | Austria                        | Tyrol                          | gPp     |
| tlaxcala                                 | Mexico                         | Tlaxcala                       | gPp     |
| tocantins                                | North                          | Tocantines                     | gPp     |
| togo                                     | Africa                         | Togo                           | gPp     |
| tohoku                                   | Japan                          | Tohoku                         | gPp     |
| tokelau                                  | Australia Oceania              | Tokelau                        | gPp     |
| toledo                                   | Castile-La Mancha              | Toledo                         | gPp     |
| tomsk_oblast                             | Siberian Federal District      | Tomsk Region                   | gPp     |
| tonga                                    | Australia Oceania              | Tonga                          | gPp     |
| toscana                                  | Italy                          | Tuscany                        | gPp     |
| trenciansky                              | Slovakia                       | Trenchansky                    | gPp     |
| trentino_alto_adige                      | Italy                          | Trentino Alto Adige            | gPp     |
| trinidad_and_tobago                      | Central America                | Trinidad and Tobago            | gPp     |
| tripura                                  | India                          | Tripura                        | gPp     |
| trnavsky                                 | Slovakia                       | Trnavsky                       | gPp     |
| troendelag                               | Norway                         | Trendelag                      | gPp     |
| troms                                    | Norway                         | Troms                          | gPp     |
| tucuman                                  | Argentina                      | Tucuman                        | gPp     |
| tula_oblast                              | Central Federal District       | Tula region                    | gPp     |
| tunisia                                  | Africa                         | Tunisia                        | gPp     |
| turkmenistan                             | Asia                           | Turkmenistan                   | gPp     |
| turks_and_caicos_islands                 | Central America                | Turks and Caicos Islands       | gPp     |
| tuva_republic                            | Siberian Federal District      | Tyva Republic                  | gPp     |
| tuvalu                                   | Australia Oceania              | Tuvalu                         | gPp     |
| tver_oblast                              | Central Federal District       | Tver region                    | gPp     |
| tyne_and_wear                            | England                        | Tyne And Wear                  | gPp     |
| tyumen_oblast                            | Ural federal district          | Tyumen region                  | gPp     |
| udmurt_republic                          | Volga Federal District         | Republic of Udmurtia           | gPp     |
| uganda                                   | Africa                         | Uganda                         | gPp     |
| ukraine                                  | Europe                         | Ukraine                        | gPp     |
| ulyanovsk_oblast                         | Volga Federal District         | Ulyanovsk region               | gPp     |
| umbria                                   | Italy                          | Umbria                         | gPp     |
| united_arab_emirates                     | Asia                           | United Arab Emirates           | gPp     |
| united_kingdom                           | Europe                         | United Kingdom                 | gPp     |
| united_states_virgin_islands             | Central America                | Virgin Islands                 | gPp     |
| unterfranken                             | Bavaria                        | uncontracted                   | gPp     |
| uppsala                                  | Sweden                         | Uppsala                        | gPp     |
| ural_federal_district                    | Russia                         | Ural federal district          | gPp     |
| uri                                      | Switzerland                    | Uri                            | gPp     |
| uruguay                                  | South America                  | Uruguay                        | gPp     |
| us                                       | North America                  | United States of America       | gPp     |
| us-midwest                               | North America                  | US-Midwest                     | gPp     |
| us-northeast                             | North America                  | USA-Northeast                  | gPp     |
| us-pacific                               | North America                  | USA-Pacific                    | gPp     |
| us-south                                 | North America                  | US South                       | gPp     |
| us-virgin-islands                        | United States of America       | US Virgin Islands              | gPp     |
| us-west                                  | North America                  | USA-West                       | gPp     |
| usa_virgin_islands                       | Central America                | Virgin Islands                 | gPp     |
| ustecky                                  | Czech Republic                 | Ustetsky                       | gPp     |
| utah                                     | United States of America       | Utah                           | gPp     |
| utrecht                                  | Netherlands                    | Utrecht                        | gPp     |
| uttar_pradesh                            | India                          | Uttar Pradesh                  | gPp     |
| uttarakhand                              | India                          | Uttarakhand                    | gPp     |
| uusimaa                                  | Finland                        | Uusimaa                        | gPp     |
| uzbekistan                               | Asia                           | Uzbekistan                     | gPp     |
| valais                                   | Switzerland                    | Valais                         | gPp     |
| valencia                                 | Valencia                       | Valencia                       | gPp     |
| valladolid                               | Castile and Leon               | Valladolid                     | gPp     |
| valle_aosta                              | Italy                          | Valle Aosta                    | gPp     |
| vanuatu                                  | Australia Oceania              | Vanuatu                        | gPp     |
| varmland                                 | Sweden                         | Värmland                       | gPp     |
| vasterbotten                             | Sweden                         | Vasterbotten                   | gPp     |
| vasternorrland                           | Sweden                         | Vast Severnaya Zemlya          | gPp     |
| vastmanland                              | Sweden                         | Vast Earth                     | gPp     |
| vastra_gotaland                          | Sweden                         | Västra Götaland                | gPp     |
| vatican_city                             | Europe                         | Vatican                        | gPp     |
| vaud                                     | Switzerland                    | In                             | gPp     |
| veneto                                   | Italy                          | Veneto                         | gPp     |
| venezuela                                | South America                  | Venezuela                      | gPp     |
| veracruz                                 | Mexico                         | Veracruz                       | gPp     |
| vermont                                  | United States of America       | Vermont                        | gPp     |
| vest-agder                               | Norway                         | Vest-Agder                     | gPp     |
| vestfold                                 | Norway                         | Vest                           | gPp     |
| vietnam                                  | Asia                           | Vietnam                        | gPp     |
| vinnytsia_oblast                         | Ukraine                        | Vinnytsia region               | gPp     |
| virginia                                 | United States of America       | Virginia                       | gPp     |
| vitebskaya                               | Belarus                        | Vitebsk region                 | gPp     |
| vizcaya                                  | Euskadi                        | Biscay                         | gPp     |
| vladimir_oblast                          | Central Federal District       | Vladimir region                | gPp     |
| volga_federal_district                   | Russia                         | Volga Federal District         | gPp     |
| volgograd_oblast                         | Southern Federal District      | Volgograd region               | gPp     |
| vologda_oblast                           | Northwestern Federal District  | Vologda Region                 | gPp     |
| volyn_oblast                             | Ukraine                        | Volyn Region                   | gPp     |
| vorarlberg                               | Austria                        | Vorarlberg                     | gPp     |
| voronezh_oblast                          | Central Federal District       | Voronezh region                | gPp     |
| vysocina                                 | Czech Republic                 | Vysochina                      | gPp     |
| wales                                    | United Kingdom                 | Wales                          | gPp     |
| wallis-et-futuna                         | Australia Oceania              | Wallis and Futuna              | gPp     |
| wallonia_french_community                | Belgium                        | French Community of Belgium    | gPp     |
| wallonia_german_community                | Belgium                        | wallonia                       | gPp     |
| warminsko_mazurskie                      | Poland                         | Warmian-Masurian               | gPp     |
| warwickshire                             | England                        | Warwickshire                   | gPp     |
| washington                               | United States of America       | Washington                     | gPp     |
| west-virginia                            | United States of America       | West Virginia                  | gPp     |
| west_bengal                              | India                          | West Bengal                    | gPp     |
| west_flanders                            | Flanders                       | West Flanders                  | gPp     |
| west_java                                | Indonesia                      | West Java                      | gPp     |
| west_kalimantan                          | Indonesia                      | West Kalimantan                | gPp     |
| west_midlands                            | England                        | Western Middle-earth           | gPp     |
| west_nusa_tenggara                       | Indonesia                      | West Nusa Tenggara             | gPp     |
| west_papua                               | Indonesia                      | West Papua                     | gPp     |
| west_sulawesi                            | Indonesia                      | West Sulawesi                  | gPp     |
| west_sumatra                             | Indonesia                      | West Sumatra                   | gPp     |
| west_sussex                              | England                        | West Sussex                    | gPp     |
| west_yorkshire                           | England                        | West Yorkshire                 | gPp     |
| western-zone                             | India                          | West Zone                      | gPp     |
| western_sahara                           | Africa                         | West Sahara                    | gPp     |
| wielkopolskie                            | Poland                         | Greater Poland                 | gPp     |
| wien                                     | Austria                        | Vein                           | gPp     |
| wiltshire                                | England                        | Wiltshire                      | gPp     |
| wisconsin                                | United States of America       | Wisconsin                      | gPp     |
| worcestershire                           | England                        | Worcestershire                 | gPp     |
| wyoming                                  | United States of America       | Wyoming                        | gPp     |
| xinjiang                                 | China                          | xinjiang                       | gPp     |
| yamalo_nenets_autonomous_okrug           | Ural federal district          | Yamalo Nenets Autonomous Okrug | gPp     |
| yaroslavl_oblast                         | Central Federal District       | Yaroslavskaya oblast           | gPp     |
| yemen                                    | Asia                           | Yemen                          | gPp     |
| yogyakarta                               | Indonesia                      | Yogyakarta                     | gPp     |
| yorkshire_and_the_humber                 | England                        | Yorkshire and the Humber       | gPp     |
| yucatan                                  | Mexico                         | Yucatan                        | gPp     |
| yukon                                    | Canada                         | Yukon                          | gPp     |
| yunnan                                   | China                          | Yunnan                         | gPp     |
| zabaykalsky_krai                         | Siberian Federal District      | Transbaikal region             | gPp     |
| zacatecas                                | Mexico                         | Zacatecas                      | gPp     |
| zachodniopomorskie                       | Poland                         | West Pomeranian                | gPp     |
| zakarpattia_oblast                       | Ukraine                        | Zakarpattia Oblast             | gPp     |
| zambia                                   | Africa                         | Zambia                         | gPp     |
| zamora                                   | Castile and Leon               | Zamora                         | gPp     |
| zaporizhia_oblast                        | Ukraine                        | Zaporozhye region              | gPp     |
| zaragoza                                 | Aragon                         | Zaragoza                       | gPp     |
| zeeland                                  | Netherlands                    | Zealand                        | gPp     |
| zhejiang                                 | China                          | Zhejiang                       | gPp     |
| zhytomyr_oblast                          | Ukraine                        | Zhytomyr Oblast                | gPp     |
| zilinsky                                 | Slovakia                       | Zhilinsky                      | gPp     |
| zimbabwe                                 | Africa                         | Zimbabwe                       | gPp     |
| zlinsky                                  | Czech Republic                 | Zlinsky                        | gPp     |
| zug                                      | Switzerland                    | Zug                            | gPp     |
| zuid_holland                             | Netherlands                    | South Holland                  | gPp     |
| zurich                                   | Switzerland                    | Zurich                         | gPp     |
Total elements: 1003

