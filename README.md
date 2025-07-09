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
| niedersachsen                               | Germany                     | Niedersachsen (mit Bremen)     | kHBPpSs |
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
| aargau                                   | switzerland                      | aargau                                   | p       |
| abitibi_temiscamingue                    | quebec                           | abitibi_temiscamingue                    | Pps     |
| abruzzo                                  | italy                            | abruzzo                                  | p       |
| aceh                                     | indonesia                        | aceh                                     | Pps     |
| acre                                     | north                            | acre                                     | Pps     |
| adygea_republic                          | southern_federal_district        | adygea_republic                          | Pps     |
| aegean                                   | turkey                           | aegean                                   | p       |
| afghanistan                              | asia                             | afghanistan                              | Pps     |
| africa                                   |                                  | africa                                   | Pps     |
| africa_france_taaf                       | africa                           | africa_france_taaf                       | Pps     |
| aguascalientes                           | mexico                           | aguascalientes                           | p       |
| ain                                      | rhone_alpes                      | ain                                      | p       |
| aisne                                    | picardie                         | aisne                                    | p       |
| akershus                                 | norway                           | akershus                                 | p       |
| alagoas                                  | northeast                        | alagoas                                  | Pps     |
| alameda                                  | california                       | alameda                                  | p       |
| aland                                    | finland                          | aland                                    | p       |
| alava                                    | euskadi                          | alava                                    | p       |
| albacete                                 | castilla_la_mancha               | albacete                                 | p       |
| alberta                                  | canada                           | alberta                                  | Pps     |
| algeria                                  | africa                           | algeria                                  | Pps     |
| alicante                                 | comunitat_valenciana             | alicante                                 | p       |
| allier                                   | auvergne                         | allier                                   | p       |
| almeria                                  | andalucia                        | almeria                                  | p       |
| alpes_de_haute_provence                  | provence_alpes_cote_d_azur       | alpes_de_haute_provence                  | p       |
| alpes_maritimes                          | provence_alpes_cote_d_azur       | alpes_maritimes                          | p       |
| alpine                                   | california                       | alpine                                   | p       |
| alsace                                   | france                           | alsace                                   | p       |
| altai_krai                               | siberian_federal_district        | altai_krai                               | Pps     |
| altai_republic                           | siberian_federal_district        | altai_republic                           | Pps     |
| amador                                   | california                       | amador                                   | p       |
| amapa                                    | north                            | amapa                                    | Pps     |
| amazonas                                 | north                            | amazonas                                 | Pps     |
| american_samoa                           | oceania                          | american_samoa                           | p       |
| amur_oblast                              | far_eastern_federal_district     | amur_oblast                              | Pps     |
| andalucia                                | spain                            | andalucia                                | p       |
| andaman_and_nicobar_islands              | india                            | andaman_and_nicobar_islands              | p       |
| andhra_pradesh                           | india                            | andhra_pradesh                           | p       |
| angola                                   | africa                           | angola                                   | Pps     |
| anguilla                                 | central-america                  | anguilla                                 | p       |
| anhui                                    | china                            | anhui                                    | Pps     |
| antigua_and_barbuda                      | central-america                  | antigua_and_barbuda                      | p       |
| antofagasta                              | chile                            | antofagasta                              | Pps     |
| antwerp                                  | flanders                         | antwerp                                  | p       |
| appenzell_ausserrhoden                   | switzerland                      | appenzell_ausserrhoden                   | p       |
| appenzell_innerrhoden                    | switzerland                      | appenzell_innerrhoden                    | p       |
| appomattox                               | virginia                         | appomattox                               | Pps     |
| aquitaine                                | france                           | aquitaine                                | p       |
| aragon                                   | spain                            | aragon                                   | p       |
| araucania                                | chile                            | araucania                                | Pps     |
| ardeche                                  | rhone_alpes                      | ardeche                                  | p       |
| ardennes                                 | champagne_ardenne                | ardennes                                 | p       |
| argentina                                | south-america                    | argentina                                | Pps     |
| argentina_la_rioja                       | argentina                        | argentina_la_rioja                       | Pps     |
| argentina_santa_cruz                     | argentina                        | argentina_santa_cruz                     | Pps     |
| arica_y_parinacota                       | chile                            | arica_y_parinacota                       | Pps     |
| ariege                                   | midi_pyrenees                    | ariege                                   | p       |
| arkhangelsk_oblast                       | northwestern_federal_district    | arkhangelsk_oblast                       | Pps     |
| armenia                                  | asia                             | armenia                                  | Pps     |
| arnsberg                                 | nordrhein_westfalen              | arnsberg                                 | p       |
| aruba                                    | central-america                  | aruba                                    | p       |
| arunachal_pradesh                        | india                            | arunachal_pradesh                        | p       |
| ashmore_and_cartier_islands              | australia                        | ashmore_and_cartier_islands              | p       |
| asia                                     |                                  | asia                                     | Pps     |
| assam                                    | india                            | assam                                    | p       |
| astrakhan_oblast                         | southern_federal_district        | astrakhan_oblast                         | Pps     |
| asturias                                 | spain                            | asturias                                 | p       |
| atacama                                  | chile                            | atacama                                  | Pps     |
| aube                                     | champagne_ardenne                | aube                                     | p       |
| aude                                     | languedoc_roussillon             | aude                                     | p       |
| aust-agder                               | norway                           | aust-agder                               | p       |
| australia                                | oceania                          | australia                                | p       |
| australian_capital_territory             | australia                        | australian_capital_territory             | p       |
| austria                                  | europe                           | austria                                  | p       |
| auvergne                                 | france                           | auvergne                                 | p       |
| aveiro                                   | portugal                         | aveiro                                   | p       |
| aveyron                                  | midi_pyrenees                    | aveyron                                  | p       |
| avila                                    | castilla_y_leon                  | avila                                    | p       |
| aysen                                    | chile                            | aysen                                    | Pps     |
| azores                                   | portugal                         | azores                                   | p       |
| badajoz                                  | extremadura                      | badajoz                                  | p       |
| bahamas                                  | central-america                  | bahamas                                  | p       |
| bahia                                    | northeast                        | bahia                                    | Pps     |
| bahrain                                  | asia                             | bahrain                                  | Pps     |
| baja_california                          | mexico                           | baja_california                          | p       |
| baja_california_sur                      | mexico                           | baja_california_sur                      | p       |
| bali                                     | indonesia                        | bali                                     | Pps     |
| bangka_belitung_islands                  | indonesia                        | bangka_belitung_islands                  | Pps     |
| bangsamoro                               | philippines                      | bangsamoro                               | Pps     |
| banskobystricky                          | slovakia                         | banskobystricky                          | p       |
| banten                                   | indonesia                        | banten                                   | Pps     |
| barbados                                 | central-america                  | barbados                                 | p       |
| barcelona                                | catalunya                        | barcelona                                | p       |
| bas_rhin                                 | alsace                           | bas_rhin                                 | p       |
| bas_saint_laurent                        | quebec                           | bas_saint_laurent                        | Pps     |
| basel_landschaft                         | switzerland                      | basel_landschaft                         | p       |
| basel_stadt                              | switzerland                      | basel_stadt                              | p       |
| bashkortostan_republic                   | volga_federal_district           | bashkortostan_republic                   | Pps     |
| basilicata                               | italy                            | basilicata                               | p       |
| basse_normandie                          | france                           | basse_normandie                          | p       |
| beijing                                  | china                            | beijing                                  | Pps     |
| beja                                     | portugal                         | beja                                     | p       |
| belgium                                  | europe                           | belgium                                  | p       |
| belgorod_oblast                          | central_federal_district         | belgorod_oblast                          | Pps     |
| bengkulu                                 | indonesia                        | bengkulu                                 | Pps     |
| benin                                    | africa                           | benin                                    | Pps     |
| bermuda                                  | north-america                    | bermuda                                  | Pps     |
| bern                                     | switzerland                      | bern                                     | p       |
| bhutan                                   | asia                             | bhutan                                   | Pps     |
| bicol_region                             | philippines                      | bicol_region                             | Pps     |
| bihar                                    | india                            | bihar                                    | p       |
| biobio                                   | chile                            | biobio                                   | Pps     |
| bir_tawil                                | africa                           | bir_tawil                                | Pps     |
| black_sea                                | turkey                           | black_sea                                | p       |
| blekinge                                 | sweden                           | blekinge                                 | p       |
| bouches_du_rhone                         | provence_alpes_cote_d_azur       | bouches_du_rhone                         | p       |
| bourgogne                                | france                           | bourgogne                                | p       |
| bouvet_island                            | africa                           | bouvet_island                            | Pps     |
| braga                                    | portugal                         | braga                                    | p       |
| braganca                                 | portugal                         | braganca                                 | p       |
| bratislavsky                             | slovakia                         | bratislavsky                             | p       |
| brazil                                   | south-america                    | brazil                                   | Pps     |
| brazil_central-west                      | brazil                           | brazil_central-west                      | Pps     |
| brazil_north                             | brazil                           | brazil_north                             | Pps     |
| brazil_northeast                         | brazil                           | brazil_northeast                         | Pps     |
| brazil_south                             | brazil                           | brazil_south                             | Pps     |
| brazil_southeast                         | brazil                           | brazil_southeast                         | Pps     |
| bretagne                                 | france                           | bretagne                                 | p       |
| british_columbia                         | canada                           | british_columbia                         | Pps     |
| british_indian_ocean_territory           | asia                             | british_indian_ocean_territory           | Pps     |
| british_virgin_islands                   | central-america                  | british_virgin_islands                   | p       |
| brunei                                   | asia                             | brunei                                   | Pps     |
| brussels_capital_region                  | belgium                          | brussels_capital_region                  | p       |
| bryansk_oblast                           | central_federal_district         | bryansk_oblast                           | Pps     |
| buenos_aires                             | argentina                        | buenos_aires                             | Pps     |
| buenos_aires_city                        | argentina                        | buenos_aires_city                        | Pps     |
| burgenland                               | austria                          | burgenland                               | p       |
| burgos                                   | castilla_y_leon                  | burgos                                   | p       |
| burkina_faso                             | africa                           | burkina_faso                             | Pps     |
| burundi                                  | africa                           | burundi                                  | Pps     |
| buryatia_republic                        | siberian_federal_district        | buryatia_republic                        | Pps     |
| buskerud                                 | norway                           | buskerud                                 | p       |
| butte                                    | california                       | butte                                    | p       |
| cabo_delgado                             | mozambique                       | cabo_delgado                             | Pps     |
| caceres                                  | extremadura                      | caceres                                  | p       |
| cadiz                                    | andalucia                        | cadiz                                    | p       |
| cagayan_valley                           | philippines                      | cagayan_valley                           | Pps     |
| calabarzon                               | philippines                      | calabarzon                               | Pps     |
| calabria                                 | italy                            | calabria                                 | p       |
| calaveras                                | california                       | calaveras                                | p       |
| california                               | us-west                          | california                               | Pps     |
| california_lake                          | california                       | california_lake                          | p       |
| california_santa_cruz                    | california                       | california_santa_cruz                    | p       |
| calvados                                 | basse_normandie                  | calvados                                 | p       |
| cambodia                                 | asia                             | cambodia                                 | Pps     |
| cameroon                                 | africa                           | cameroon                                 | Pps     |
| campania                                 | italy                            | campania                                 | p       |
| campeche                                 | mexico                           | campeche                                 | p       |
| canada                                   | north-america                    | canada                                   | Pps     |
| canarias                                 | spain                            | canarias                                 | Pps     |
| cantabria                                | spain                            | cantabria                                | p       |
| cantal                                   | auvergne                         | cantal                                   | p       |
| cape_verde                               | africa                           | cape_verde                               | Pps     |
| capital_district                         | new-york                         | capital_district                         | Pps     |
| capitale_nationale                       | quebec                           | capitale_nationale                       | Pps     |
| caraga                                   | philippines                      | caraga                                   | Pps     |
| caribbean                                | central-america                  | caribbean                                | p       |
| castellon                                | comunitat_valenciana             | castellon                                | p       |
| castelo_branco                           | portugal                         | castelo_branco                           | p       |
| castilla_la_mancha                       | spain                            | castilla_la_mancha                       | p       |
| castilla_y_leon                          | spain                            | castilla_y_leon                          | p       |
| catalunya                                | spain                            | catalunya                                | p       |
| catamarca                                | argentina                        | catamarca                                | Pps     |
| cayman_islands                           | central-america                  | cayman_islands                           | p       |
| ceara                                    | northeast                        | ceara                                    | Pps     |
| central-america                          |                                  | central-america                          | Pps     |
| central-west                             | brazil                           | central-west                             |         |
| central_african_republic                 | africa                           | central_african_republic                 | Pps     |
| central_anatolia                         | turkey                           | central_anatolia                         | p       |
| central_federal_district                 | russia                           | central_federal_district                 | Pps     |
| central_finland                          | finland                          | central_finland                          | p       |
| central_java                             | indonesia                        | central_java                             | Pps     |
| central_kalimantan                       | indonesia                        | central_kalimantan                       | Pps     |
| central_luzon                            | philippines                      | central_luzon                            | Pps     |
| central_new_york                         | new-york                         | central_new_york                         | Pps     |
| central_ontario                          | ontario                          | central_ontario                          | Pps     |
| central_ostrobothnia                     | finland                          | central_ostrobothnia                     | p       |
| central_papua                            | indonesia                        | central_papua                            | Pps     |
| central_sulawesi                         | indonesia                        | central_sulawesi                         | Pps     |
| central_visayas                          | philippines                      | central_visayas                          | Pps     |
| centre                                   | france                           | centre                                   | p       |
| centre_du_quebec                         | quebec                           | centre_du_quebec                         | Pps     |
| ceuta                                    | spain                            | ceuta                                    | Pps     |
| chaco                                    | argentina                        | chaco                                    | Pps     |
| chad                                     | africa                           | chad                                     | Pps     |
| champagne_ardenne                        | france                           | champagne_ardenne                        | p       |
| chandigarh                               | india                            | chandigarh                               | p       |
| charente                                 | poitou_charentes                 | charente                                 | p       |
| charente_maritime                        | poitou_charentes                 | charente_maritime                        | p       |
| chaudiere_appalaches                     | quebec                           | chaudiere_appalaches                     | Pps     |
| chechen_republic                         | north_caucasian_federal_district | chechen_republic                         | Pps     |
| chelyabinsk_oblast                       | ural_federal_district            | chelyabinsk_oblast                       | Pps     |
| cher                                     | centre                           | cher                                     | p       |
| cherkasy_oblast                          | ukraine                          | cherkasy_oblast                          | p       |
| chernihiv_oblast                         | ukraine                          | chernihiv_oblast                         | p       |
| chernivtsi_oblast                        | ukraine                          | chernivtsi_oblast                        | p       |
| chesapeake                               | virginia                         | chesapeake                               | Pps     |
| chhattisgarh                             | india                            | chhattisgarh                             | p       |
| chiapas                                  | mexico                           | chiapas                                  | p       |
| chihuahua                                | mexico                           | chihuahua                                | p       |
| chile                                    | south-america                    | chile                                    | Pps     |
| china                                    | asia                             | china                                    | Pps     |
| chongqing                                | china                            | chongqing                                | Pps     |
| christmas_island                         | australia                        | christmas_island                         | p       |
| chubu                                    | japan                            | chubu                                    | Pps     |
| chubut                                   | argentina                        | chubut                                   | Pps     |
| chugoku                                  | japan                            | chugoku                                  | Pps     |
| chukotka_autonomous_okrug                | far_eastern_federal_district     | chukotka_autonomous_okrug                | Pps     |
| chuvash_republic                         | volga_federal_district           | chuvash_republic                         | Pps     |
| ciudad_real                              | castilla_la_mancha               | ciudad_real                              | p       |
| clipperton                               | oceania                          | clipperton                               | p       |
| coahuila                                 | mexico                           | coahuila                                 | p       |
| coastal                                  | tanzania                         | coastal                                  | Pps     |
| cocos_islands                            | australia                        | cocos_islands                            | p       |
| coimbra                                  | portugal                         | coimbra                                  | p       |
| colima                                   | mexico                           | colima                                   | p       |
| colorado                                 | us-west                          | colorado                                 | Pps     |
| colorado_northeast                       | colorado                         | colorado_northeast                       | Pps     |
| colorado_northwest                       | colorado                         | colorado_northwest                       | Pps     |
| colorado_southeast                       | colorado                         | colorado_southeast                       | Pps     |
| colorado_southwest                       | colorado                         | colorado_southwest                       | Pps     |
| colusa                                   | california                       | colusa                                   | p       |
| comoros                                  | africa                           | comoros                                  | Pps     |
| comunidad_de_madrid                      | spain                            | comunidad_de_madrid                      | p       |
| comunidad_foral_de_navarra               | spain                            | comunidad_foral_de_navarra               | p       |
| comunitat_valenciana                     | spain                            | comunitat_valenciana                     | p       |
| congo_brazzaville                        | africa                           | congo_brazzaville                        | Pps     |
| congo_kinshasa                           | africa                           | congo_kinshasa                           | Pps     |
| contra_costa                             | california                       | contra_costa                             | p       |
| cook                                     | illinois                         | cook                                     | Pps     |
| cook_islands                             | oceania                          | cook_islands                             | p       |
| coquimbo                                 | chile                            | coquimbo                                 | Pps     |
| coral_sea_islands                        | australia                        | coral_sea_islands                        | p       |
| cordillera_administrative_region         | philippines                      | cordillera_administrative_region         | Pps     |
| cordoba                                  | argentina                        | cordoba                                  | Pps     |
| correze                                  | limousin                         | correze                                  | p       |
| corrientes                               | argentina                        | corrientes                               | Pps     |
| corse                                    | france                           | corse                                    | p       |
| corse_du_sud                             | corse                            | corse_du_sud                             | p       |
| costa_rica                               | central-america                  | costa_rica                               | p       |
| cote_d_or                                | bourgogne                        | cote_d_or                                | p       |
| cote_nord                                | quebec                           | cote_nord                                | Pps     |
| cotes_d_armor                            | bretagne                         | cotes_d_armor                            | p       |
| creuse                                   | limousin                         | creuse                                   | p       |
| crimea                                   | ukraine                          | crimea                                   | p       |
| crimea_republic                          | southern_federal_district        | crimea_republic                          | Pps     |
| cuenca                                   | castilla_la_mancha               | cuenca                                   | p       |
| culpeper                                 | virginia                         | culpeper                                 | Pps     |
| curacao                                  | central-america                  | curacao                                  | p       |
| czech_republic                           | europe                           | czech_republic                           | p       |
| dadra_and_nagar_haveli                   | india                            | dadra_and_nagar_haveli                   | p       |
| dadra_and_nagar_haveli_and_daman_and_diu | india                            | dadra_and_nagar_haveli_and_daman_and_diu | p       |
| dagestan_republic                        | north_caucasian_federal_district | dagestan_republic                        | Pps     |
| dalarna                                  | sweden                           | dalarna                                  | p       |
| daman_and_diu                            | india                            | daman_and_diu                            | p       |
| davao_region                             | philippines                      | davao_region                             | Pps     |
| del_norte                                | california                       | del_norte                                | p       |
| denver                                   | colorado                         | denver                                   | Pps     |
| detmold                                  | nordrhein_westfalen              | detmold                                  | p       |
| detroit_metro                            | michigan                         | detroit_metro                            | Pps     |
| deux_sevres                              | poitou_charentes                 | deux_sevres                              | p       |
| distrito-federal                         | central-west                     | distrito-federal                         | Pps     |
| djibouti                                 | africa                           | djibouti                                 | Pps     |
| dnipropetrovsk_oblast                    | ukraine                          | dnipropetrovsk_oblast                    | p       |
| dolnoslaskie                             | poland                           | dolnoslaskie                             | p       |
| dominica                                 | central-america                  | dominica                                 | p       |
| dominican_republic                       | central-america                  | dominican_republic                       | p       |
| donetsk_oblast                           | ukraine                          | donetsk_oblast                           | p       |
| dordogne                                 | aquitaine                        | dordogne                                 | p       |
| doubs                                    | franche_comte                    | doubs                                    | p       |
| drenthe                                  | netherlands                      | drenthe                                  | p       |
| drome                                    | rhone_alpes                      | drome                                    | p       |
| durango                                  | mexico                           | durango                                  | p       |
| dusseldorf                               | nordrhein_westfalen              | dusseldorf                               | p       |
| east_flanders                            | flanders                         | east_flanders                            | p       |
| east_java                                | indonesia                        | east_java                                | Pps     |
| east_kalimantan                          | indonesia                        | east_kalimantan                          | Pps     |
| east_midlands                            | england                          | east_midlands                            | p       |
| east_nusa_tenggara                       | indonesia                        | east_nusa_tenggara                       | Pps     |
| east_timor                               | asia                             | east_timor                               | Pps     |
| eastern_anatolia                         | turkey                           | eastern_anatolia                         | p       |
| eastern_cape                             | south_africa                     | eastern_cape                             | Pps     |
| eastern_ontario                          | ontario                          | eastern_ontario                          | Pps     |
| eastern_visayas                          | philippines                      | eastern_visayas                          | Pps     |
| egypt                                    | africa                           | egypt                                    | Pps     |
| el_dorado                                | california                       | el_dorado                                | p       |
| el_salvador                              | central-america                  | el_salvador                              | p       |
| emilia_romagna                           | italy                            | emilia_romagna                           | p       |
| england                                  | united_kingdom                   | england                                  | p       |
| england_east                             | england                          | england_east                             | p       |
| england_north_east                       | england                          | england_north_east                       | p       |
| england_north_west                       | england                          | england_north_west                       | p       |
| england_south_east                       | england                          | england_south_east                       | p       |
| england_south_west                       | england                          | england_south_west                       | p       |
| entre_rios                               | argentina                        | entre_rios                               | Pps     |
| equatorial_guinea                        | africa                           | equatorial_guinea                        | Pps     |
| eritrea                                  | africa                           | eritrea                                  | Pps     |
| espirito-santo                           | southeast                        | espirito-santo                           | Pps     |
| essonne                                  | ile_de_france                    | essonne                                  | p       |
| estrie                                   | quebec                           | estrie                                   | Pps     |
| ethiopia                                 | africa                           | ethiopia                                 | Pps     |
| eure                                     | haute_normandie                  | eure                                     | p       |
| eure_et_loir                             | centre                           | eure_et_loir                             | p       |
| europe                                   |                                  | europe                                   | p       |
| euskadi                                  | spain                            | euskadi                                  | p       |
| evora                                    | portugal                         | evora                                    | p       |
| extremadura                              | spain                            | extremadura                              | p       |
| fairfax                                  | virginia                         | fairfax                                  | Pps     |
| falkland                                 | south-america                    | falkland                                 | Pps     |
| far_eastern_federal_district             | russia                           | far_eastern_federal_district             | Pps     |
| far_west                                 | new_south_wales                  | far_west                                 | p       |
| faro                                     | portugal                         | faro                                     | p       |
| fiji                                     | merge                            | fiji                                     | Pps     |
| fiji_east                                | oceania                          | fiji_east                                | p       |
| fiji_west                                | oceania                          | fiji_west                                | p       |
| finger_lakes                             | new-york                         | finger_lakes                             | Pps     |
| finistere                                | bretagne                         | finistere                                | p       |
| finland                                  | europe                           | finland                                  | p       |
| finnmark                                 | norway                           | finnmark                                 | p       |
| flanders                                 | belgium                          | flanders                                 | p       |
| flemish_brabant                          | flanders                         | flemish_brabant                          | p       |
| flevoland                                | netherlands                      | flevoland                                | p       |
| florida                                  | us-south                         | florida                                  | Pps     |
| florida_east_central                     | florida                          | florida_east_central                     | Pps     |
| florida_northeast                        | florida                          | florida_northeast                        | Pps     |
| florida_northwest                        | florida                          | florida_northwest                        | Pps     |
| florida_southwest                        | florida                          | florida_southwest                        | Pps     |
| formosa                                  | argentina                        | formosa                                  | Pps     |
| france                                   | europe                           | france                                   | p       |
| france_metro_dom_com_nc                  | merge                            | france_metro_dom_com_nc                  | Pps     |
| franche_comte                            | france                           | franche_comte                            | p       |
| franche_comte_jura                       | franche_comte                    | franche_comte_jura                       | p       |
| free_state                               | south_africa                     | free_state                               | Pps     |
| fresno                                   | california                       | fresno                                   | p       |
| fribourg                                 | switzerland                      | fribourg                                 | p       |
| friesland                                | netherlands                      | friesland                                | p       |
| friuli_venezia_giulia                    | italy                            | friuli_venezia_giulia                    | p       |
| fujian                                   | china                            | fujian                                   | Pps     |
| gabon                                    | africa                           | gabon                                    | Pps     |
| galicia                                  | spain                            | galicia                                  | p       |
| gambia                                   | africa                           | gambia                                   | Pps     |
| gansu                                    | china                            | gansu                                    | Pps     |
| gard                                     | languedoc_roussillon             | gard                                     | p       |
| gaspesie_iles_de_la_madeleine            | quebec                           | gaspesie_iles_de_la_madeleine            | Pps     |
| gatorland                                | florida                          | gatorland                                | Pps     |
| gauteng                                  | south_africa                     | gauteng                                  | Pps     |
| gavleborg                                | sweden                           | gavleborg                                | p       |
| gaza                                     | mozambique                       | gaza                                     | Pps     |
| gelderland                               | netherlands                      | gelderland                               | p       |
| geneva                                   | switzerland                      | geneva                                   | p       |
| georgia                                  | asia                             | georgia                                  | Pps     |
| georgia_northeast                        | georgia                          | georgia_northeast                        | Pps     |
| georgia_northwest                        | georgia                          | georgia_northwest                        | Pps     |
| georgia_southeast                        | georgia                          | georgia_southeast                        | Pps     |
| georgia_southwest                        | georgia                          | georgia_southwest                        | Pps     |
| germany                                  | europe                           | germany                                  | p       |
| gers                                     | midi_pyrenees                    | gers                                     | p       |
| ghana                                    | africa                           | ghana                                    | Pps     |
| gibraltar                                | europe                           | gibraltar                                | p       |
| girona                                   | catalunya                        | girona                                   | p       |
| gironde                                  | aquitaine                        | gironde                                  | p       |
| glarus                                   | switzerland                      | glarus                                   | p       |
| glenn                                    | california                       | glenn                                    | p       |
| goa                                      | india                            | goa                                      | p       |
| goias                                    | central-west                     | goias                                    | Pps     |
| gold_coast                               | florida                          | gold_coast                               | Pps     |
| golden_horseshoe                         | ontario                          | golden_horseshoe                         | Pps     |
| gorontalo                                | indonesia                        | gorontalo                                | Pps     |
| gotland                                  | sweden                           | gotland                                  | p       |
| granada                                  | andalucia                        | granada                                  | p       |
| greater_london                           | england                          | greater_london                           | p       |
| greater_metropolitan_newcastle           | new_south_wales                  | greater_metropolitan_newcastle           | p       |
| greater_metropolitan_sydney              | new_south_wales                  | greater_metropolitan_sydney              | p       |
| grenada                                  | central-america                  | grenada                                  | p       |
| grisons                                  | switzerland                      | grisons                                  | p       |
| groningen                                | netherlands                      | groningen                                | p       |
| guadalajara                              | castilla_la_mancha               | guadalajara                              | p       |
| guadeloupe                               | central-america                  | guadeloupe                               | p       |
| guam                                     | oceania                          | guam                                     | p       |
| guanajuato                               | mexico                           | guanajuato                               | p       |
| guangdong                                | china                            | guangdong                                | Pps     |
| guangxi                                  | china                            | guangxi                                  | Pps     |
| guarda                                   | portugal                         | guarda                                   | p       |
| guernesey                                | europe                           | guernesey                                | p       |
| guerrero                                 | mexico                           | guerrero                                 | p       |
| guinea                                   | africa                           | guinea                                   | Pps     |
| guipuzcoa                                | euskadi                          | guipuzcoa                                | p       |
| guizhou                                  | china                            | guizhou                                  | Pps     |
| gujarat                                  | india                            | gujarat                                  | p       |
| guyana                                   | south-america                    | guyana                                   | Pps     |
| guyane                                   | south-america                    | guyane                                   | Pps     |
| hainan                                   | china                            | hainan                                   | Pps     |
| haiti                                    | central-america                  | haiti                                    | p       |
| halland                                  | sweden                           | halland                                  | p       |
| haryana                                  | india                            | haryana                                  | p       |
| haut_rhin                                | alsace                           | haut_rhin                                | p       |
| haute_corse                              | corse                            | haute_corse                              | p       |
| haute_garonne                            | midi_pyrenees                    | haute_garonne                            | p       |
| haute_loire                              | auvergne                         | haute_loire                              | p       |
| haute_marne                              | champagne_ardenne                | haute_marne                              | p       |
| haute_normandie                          | france                           | haute_normandie                          | p       |
| haute_saone                              | franche_comte                    | haute_saone                              | p       |
| haute_savoie                             | rhone_alpes                      | haute_savoie                             | p       |
| haute_vienne                             | limousin                         | haute_vienne                             | p       |
| hautes_alpes                             | provence_alpes_cote_d_azur       | hautes_alpes                             | p       |
| hautes_pyrenees                          | midi_pyrenees                    | hautes_pyrenees                          | p       |
| hauts_de_seine                           | ile_de_france                    | hauts_de_seine                           | p       |
| heard_island_and_mcdonald_slands         | australia                        | heard_island_and_mcdonald_slands         | p       |
| hebei                                    | china                            | hebei                                    | Pps     |
| hedmark                                  | norway                           | hedmark                                  | p       |
| heilongjiang                             | china                            | heilongjiang                             | Pps     |
| henan                                    | china                            | henan                                    | Pps     |
| herault                                  | languedoc_roussillon             | herault                                  | p       |
| hidalgo                                  | mexico                           | hidalgo                                  | p       |
| highland_papua                           | indonesia                        | highland_papua                           | Pps     |
| himachal_pradesh                         | india                            | himachal_pradesh                         | p       |
| hokkaido                                 | japan                            | hokkaido                                 | Pps     |
| honduras                                 | central-america                  | honduras                                 | p       |
| hong_kong                                | china                            | hong_kong                                | Pps     |
| hordaland                                | norway                           | hordaland                                | p       |
| hubei                                    | china                            | hubei                                    | Pps     |
| hudson_valley                            | new-york                         | hudson_valley                            | Pps     |
| huelva                                   | andalucia                        | huelva                                   | p       |
| huesca                                   | aragon                           | huesca                                   | p       |
| humboldt                                 | california                       | humboldt                                 | p       |
| hunan                                    | china                            | hunan                                    | Pps     |
| ile_de_france                            | france                           | ile_de_france                            | p       |
| ilemi                                    | africa                           | ilemi                                    | Pps     |
| illawarra                                | new_south_wales                  | illawarra                                | p       |
| ille_et_vilaine                          | bretagne                         | ille_et_vilaine                          | p       |
| illes_balears                            | spain                            | illes_balears                            | p       |
| illinois                                 | us-midwest                       | illinois                                 | Pps     |
| illinois_central                         | illinois                         | illinois_central                         | Pps     |
| illinois_east_central                    | illinois                         | illinois_east_central                    | Pps     |
| illinois_north                           | illinois                         | illinois_north                           | Pps     |
| illinois_northeast                       | illinois                         | illinois_northeast                       | Pps     |
| illinois_northwest                       | illinois                         | illinois_northwest                       | Pps     |
| illinois_southern                        | illinois                         | illinois_southern                        | Pps     |
| illinois_southwest                       | illinois                         | illinois_southwest                       | Pps     |
| ilocos_region                            | philippines                      | ilocos_region                            | Pps     |
| imperial                                 | california                       | imperial                                 | p       |
| india                                    | asia                             | india                                    | Pps     |
| indonesia                                | asia                             | indonesia                                | Pps     |
| indre                                    | centre                           | indre                                    | p       |
| indre_et_loire                           | centre                           | indre_et_loire                           | p       |
| ingushetia_republic                      | north_caucasian_federal_district | ingushetia_republic                      | Pps     |
| inhambane                                | mozambique                       | inhambane                                | Pps     |
| inner_mongolia                           | china                            | inner_mongolia                           | Pps     |
| inyo                                     | california                       | inyo                                     | p       |
| ionian_sea                               | seas                             | ionian_sea                               | p       |
| ireland                                  | europe                           | ireland                                  | p       |
| irkutsk_oblast                           | siberian_federal_district        | irkutsk_oblast                           | Pps     |
| isere                                    | rhone_alpes                      | isere                                    | p       |
| israel                                   | asia                             | israel                                   | Pps     |
| israel_and_palestine                     | asia                             | israel_and_palestine                     | Pps     |
| israel_west_bank                         | asia                             | israel_west_bank                         | Pps     |
| italy                                    | europe                           | italy                                    | p       |
| ivano-frankivsk_oblast                   | ukraine                          | ivano-frankivsk_oblast                   | p       |
| ivanovo_oblast                           | central_federal_district         | ivanovo_oblast                           | Pps     |
| ivory_coast                              | africa                           | ivory_coast                              | Pps     |
| jaen                                     | andalucia                        | jaen                                     | p       |
| jakarta                                  | indonesia                        | jakarta                                  | Pps     |
| jalisco                                  | mexico                           | jalisco                                  | p       |
| jamaica                                  | central-america                  | jamaica                                  | p       |
| jambi                                    | indonesia                        | jambi                                    | Pps     |
| jammu_and_kashmir                        | india                            | jammu_and_kashmir                        | p       |
| jamtland                                 | sweden                           | jamtland                                 | p       |
| jan_mayen                                | norway                           | jan_mayen                                | p       |
| japan                                    | asia                             | japan                                    | Pps     |
| jersey                                   | europe                           | jersey                                   | p       |
| jewish_autonomous_oblast                 | far_eastern_federal_district     | jewish_autonomous_oblast                 | Pps     |
| jharkhand                                | india                            | jharkhand                                | p       |
| jiangsu                                  | china                            | jiangsu                                  | Pps     |
| jiangxi                                  | china                            | jiangxi                                  | Pps     |
| jihocesky                                | czech_republic                   | jihocesky                                | p       |
| jihomoravsky                             | czech_republic                   | jihomoravsky                             | p       |
| jilin                                    | china                            | jilin                                    | Pps     |
| jonkoping                                | sweden                           | jonkoping                                | p       |
| jujuy                                    | argentina                        | jujuy                                    | Pps     |
| kabardino_balkar_republic                | north_caucasian_federal_district | kabardino_balkar_republic                | Pps     |
| kainuu                                   | finland                          | kainuu                                   | p       |
| kaliningrad_oblast                       | northwestern_federal_district    | kaliningrad_oblast                       | Pps     |
| kalmar                                   | sweden                           | kalmar                                   | p       |
| kalmykia_republic                        | southern_federal_district        | kalmykia_republic                        | Pps     |
| kaluga_oblast                            | central_federal_district         | kaluga_oblast                            | Pps     |
| kamchatka_krai                           | far_eastern_federal_district     | kamchatka_krai                           | Pps     |
| kansai                                   | japan                            | kansai                                   | Pps     |
| kanta_hame                               | finland                          | kanta_hame                               | p       |
| kanto                                    | japan                            | kanto                                    | Pps     |
| karachay_cherkess_republic               | north_caucasian_federal_district | karachay_cherkess_republic               | Pps     |
| karelia_republic                         | northwestern_federal_district    | karelia_republic                         | Pps     |
| karlovarsky                              | czech_republic                   | karlovarsky                              | p       |
| karnataka                                | india                            | karnataka                                | p       |
| karnten                                  | austria                          | karnten                                  | p       |
| kemerovo_oblast                          | siberian_federal_district        | kemerovo_oblast                          | Pps     |
| kenya                                    | africa                           | kenya                                    | Pps     |
| kerala                                   | india                            | kerala                                   | p       |
| kern                                     | california                       | kern                                     | p       |
| khabarovsk_krai                          | far_eastern_federal_district     | khabarovsk_krai                          | Pps     |
| khakassia_republic                       | siberian_federal_district        | khakassia_republic                       | Pps     |
| khanty_mansi_autonomous_okrug            | ural_federal_district            | khanty_mansi_autonomous_okrug            | Pps     |
| kharkiv_oblast                           | ukraine                          | kharkiv_oblast                           | p       |
| kherson_oblast                           | ukraine                          | kherson_oblast                           | p       |
| khmelnytskyi_oblast                      | ukraine                          | khmelnytskyi_oblast                      | p       |
| kiev                                     | ukraine                          | kiev                                     | p       |
| kiev_oblast                              | ukraine                          | kiev_oblast                              | p       |
| kings                                    | california                       | kings                                    | p       |
| kiribati                                 | merge                            | kiribati                                 | Pps     |
| kiribati_east                            | oceania                          | kiribati_east                            | p       |
| kiribati_west                            | oceania                          | kiribati_west                            | p       |
| kirov_oblast                             | volga_federal_district           | kirov_oblast                             | Pps     |
| kirovohrad_oblast                        | ukraine                          | kirovohrad_oblast                        | p       |
| koln                                     | nordrhein_westfalen              | koln                                     | p       |
| komi_republic                            | northwestern_federal_district    | komi_republic                            | Pps     |
| kosicky                                  | slovakia                         | kosicky                                  | p       |
| kostroma_oblast                          | central_federal_district         | kostroma_oblast                          | Pps     |
| kralovehradecky                          | czech_republic                   | kralovehradecky                          | p       |
| krasnodar_krai                           | southern_federal_district        | krasnodar_krai                           | Pps     |
| krasnoyarsk_krai                         | siberian_federal_district        | krasnoyarsk_krai                         | Pps     |
| kronoberg                                | sweden                           | kronoberg                                | p       |
| kujawsko_pomorskie                       | poland                           | kujawsko_pomorskie                       | p       |
| kurgan_oblast                            | ural_federal_district            | kurgan_oblast                            | Pps     |
| kursk_oblast                             | central_federal_district         | kursk_oblast                             | Pps     |
| kuwait                                   | asia                             | kuwait                                   | Pps     |
| kwazulu_natal                            | south_africa                     | kwazulu_natal                            | Pps     |
| kymenlaakso                              | finland                          | kymenlaakso                              | p       |
| kyushu                                   | japan                            | kyushu                                   | Pps     |
| la_coruna                                | galicia                          | la_coruna                                | p       |
| la_pampa                                 | argentina                        | la_pampa                                 | Pps     |
| ladakh                                   | india                            | ladakh                                   | p       |
| lakshadweep                              | india                            | lakshadweep                              | p       |
| lampung                                  | indonesia                        | lampung                                  | Pps     |
| lanaudiere                               | quebec                           | lanaudiere                               | Pps     |
| landes                                   | aquitaine                        | landes                                   | p       |
| languedoc_roussillon                     | france                           | languedoc_roussillon                     | p       |
| laos                                     | asia                             | laos                                     | Pps     |
| lapland                                  | finland                          | lapland                                  | p       |
| las_palmas                               | canarias                         | las_palmas                               | Pps     |
| lassen                                   | california                       | lassen                                   | p       |
| laurentides                              | quebec                           | laurentides                              | Pps     |
| laval                                    | quebec                           | laval                                    | Pps     |
| lazio                                    | italy                            | lazio                                    | p       |
| lebanon                                  | asia                             | lebanon                                  | Pps     |
| leiria                                   | portugal                         | leiria                                   | p       |
| leningrad_oblast                         | northwestern_federal_district    | leningrad_oblast                         | Pps     |
| leon                                     | castilla_y_leon                  | leon                                     | p       |
| lesotho                                  | africa                           | lesotho                                  | Pps     |
| liaoning                                 | china                            | liaoning                                 | Pps     |
| liberecky                                | czech_republic                   | liberecky                                | p       |
| liguria                                  | italy                            | liguria                                  | p       |
| limburg                                  | netherlands                      | limburg                                  | p       |
| limousin                                 | france                           | limousin                                 | p       |
| limpopo                                  | south_africa                     | limpopo                                  | Pps     |
| lipetsk_oblast                           | central_federal_district         | lipetsk_oblast                           | Pps     |
| lisbon                                   | portugal                         | lisbon                                   | p       |
| lleida                                   | catalunya                        | lleida                                   | p       |
| lodzkie                                  | poland                           | lodzkie                                  | p       |
| loir_et_cher                             | centre                           | loir_et_cher                             | p       |
| loire                                    | rhone_alpes                      | loire                                    | p       |
| loire_atlantique                         | pays_de_la_loire                 | loire_atlantique                         | p       |
| loiret                                   | centre                           | loiret                                   | p       |
| lombardia                                | italy                            | lombardia                                | p       |
| long_island                              | new-york                         | long_island                              | Pps     |
| lorraine                                 | france                           | lorraine                                 | p       |
| los_angeles                              | california                       | los_angeles                              | p       |
| los_lagos                                | chile                            | los_lagos                                | Pps     |
| los_rios                                 | chile                            | los_rios                                 | Pps     |
| lot                                      | midi_pyrenees                    | lot                                      | p       |
| lot_et_garonne                           | aquitaine                        | lot_et_garonne                           | p       |
| lozere                                   | languedoc_roussillon             | lozere                                   | p       |
| lubelskie                                | poland                           | lubelskie                                | p       |
| lubuskie                                 | poland                           | lubuskie                                 | p       |
| lucerne                                  | switzerland                      | lucerne                                  | p       |
| lugo                                     | galicia                          | lugo                                     | p       |
| luhansk_oblast                           | ukraine                          | luhansk_oblast                           | p       |
| luxembourg                               | europe                           | luxembourg                               | p       |
| lviv_oblast                              | ukraine                          | lviv_oblast                              | p       |
| macau                                    | china                            | macau                                    | Pps     |
| madagascar                               | africa                           | madagascar                               | Pps     |
| madeira                                  | portugal                         | madeira                                  | p       |
| madera                                   | california                       | madera                                   | p       |
| madhya_pradesh                           | india                            | madhya_pradesh                           | p       |
| magadan_oblast                           | far_eastern_federal_district     | magadan_oblast                           | Pps     |
| magallanes                               | chile                            | magallanes                               | Pps     |
| maharashtra                              | india                            | maharashtra                              | p       |
| maine_et_loire                           | pays_de_la_loire                 | maine_et_loire                           | p       |
| malaga                                   | andalucia                        | malaga                                   | p       |
| malawi                                   | africa                           | malawi                                   | Pps     |
| malaysia                                 | asia                             | malaysia                                 | Pps     |
| maldives                                 | asia                             | maldives                                 | Pps     |
| mali                                     | africa                           | mali                                     | Pps     |
| malopolskie                              | poland                           | malopolskie                              | p       |
| maluku                                   | indonesia                        | maluku                                   | Pps     |
| manche                                   | basse_normandie                  | manche                                   | p       |
| manica                                   | mozambique                       | manica                                   | Pps     |
| manipur                                  | india                            | manipur                                  | p       |
| manitoba                                 | canada                           | manitoba                                 | Pps     |
| maputo                                   | mozambique                       | maputo                                   | Pps     |
| maputo_city                              | mozambique                       | maputo_city                              | Pps     |
| maranhao                                 | northeast                        | maranhao                                 | Pps     |
| marche                                   | italy                            | marche                                   | p       |
| mari_el_republic                         | volga_federal_district           | mari_el_republic                         | Pps     |
| marin                                    | california                       | marin                                    | p       |
| mariposa                                 | california                       | mariposa                                 | p       |
| marmara                                  | turkey                           | marmara                                  | p       |
| marne                                    | champagne_ardenne                | marne                                    | p       |
| marshall-islands                         | oceania                          | marshall-islands                         | p       |
| marshall_islands                         | oceania                          | marshall_islands                         | p       |
| martinique                               | central-america                  | martinique                               | p       |
| mato-grosso                              | central-west                     | mato-grosso                              | Pps     |
| mato-grosso-do-sul                       | central-west                     | mato-grosso-do-sul                       | Pps     |
| maule                                    | chile                            | maule                                    | Pps     |
| mauricie                                 | quebec                           | mauricie                                 | Pps     |
| mauritania                               | africa                           | mauritania                               | Pps     |
| mauritius                                | africa                           | mauritius                                | Pps     |
| mayenne                                  | pays_de_la_loire                 | mayenne                                  | p       |
| mayotte                                  | africa                           | mayotte                                  | Pps     |
| mazowieckie                              | poland                           | mazowieckie                              | p       |
| mediterranean                            | turkey                           | mediterranean                            | p       |
| meghalaya                                | india                            | meghalaya                                | p       |
| melilla                                  | spain                            | melilla                                  | Pps     |
| mendocino                                | california                       | mendocino                                | p       |
| mendoza                                  | argentina                        | mendoza                                  | Pps     |
| merced                                   | california                       | merced                                   | p       |
| merge                                    |                                  | merge                                    |         |
| merge_france_taaf                        | merge                            | merge_france_taaf                        | Pps     |
| metro_manila                             | philippines                      | metro_manila                             | Pps     |
| meurthe_et_moselle                       | lorraine                         | meurthe_et_moselle                       | p       |
| meuse                                    | lorraine                         | meuse                                    | p       |
| mexico                                   | north-america                    | mexico                                   | Pps     |
| mexico_city                              | mexico                           | mexico_city                              | p       |
| michigan                                 | us-midwest                       | michigan                                 | Pps     |
| michigan_central                         | michigan                         | michigan_central                         | Pps     |
| michigan_southeast                       | michigan                         | michigan_southeast                       | Pps     |
| michigan_southwest                       | michigan                         | michigan_southwest                       | Pps     |
| michigan_west                            | michigan                         | michigan_west                            | Pps     |
| michoacan                                | mexico                           | michoacan                                | p       |
| micronesia                               | oceania                          | micronesia                               | p       |
| mid_north_coast                          | new_south_wales                  | mid_north_coast                          | p       |
| midi_pyrenees                            | france                           | midi_pyrenees                            | p       |
| mimaropa                                 | philippines                      | mimaropa                                 | Pps     |
| minas-gerais                             | southeast                        | minas-gerais                             | Pps     |
| misiones                                 | argentina                        | misiones                                 | Pps     |
| mizoram                                  | india                            | mizoram                                  | p       |
| modoc                                    | california                       | modoc                                    | p       |
| moere_og_romsdal                         | norway                           | moere_og_romsdal                         | p       |
| mohawk_valley                            | new-york                         | mohawk_valley                            | Pps     |
| molise                                   | italy                            | molise                                   | p       |
| monaco                                   | europe                           | monaco                                   | p       |
| mono                                     | california                       | mono                                     | p       |
| monteregie                               | quebec                           | monteregie                               | Pps     |
| monterey                                 | california                       | monterey                                 | p       |
| montreal                                 | quebec                           | montreal                                 | Pps     |
| montserrat                               | central-america                  | montserrat                               | p       |
| moravskoslezsky                          | czech_republic                   | moravskoslezsky                          | p       |
| morbihan                                 | bretagne                         | morbihan                                 | p       |
| mordovia_republic                        | volga_federal_district           | mordovia_republic                        | Pps     |
| morelos                                  | mexico                           | morelos                                  | p       |
| morocco                                  | africa                           | morocco                                  | Pps     |
| moscow                                   | central_federal_district         | moscow                                   | Pps     |
| moscow_oblast                            | central_federal_district         | moscow_oblast                            | Pps     |
| moselle                                  | lorraine                         | moselle                                  | p       |
| mozambique                               | africa                           | mozambique                               | Pps     |
| mpumalanga                               | south_africa                     | mpumalanga                               | Pps     |
| munster                                  | nordrhein_westfalen              | munster                                  | p       |
| murmansk_oblast                          | northwestern_federal_district    | murmansk_oblast                          | Pps     |
| murray                                   | new_south_wales                  | murray                                   | p       |
| myanmar                                  | asia                             | myanmar                                  | Pps     |
| mykolaiv_oblast                          | ukraine                          | mykolaiv_oblast                          | p       |
| nagaland                                 | india                            | nagaland                                 | p       |
| namibia                                  | africa                           | namibia                                  | Pps     |
| nampula                                  | mozambique                       | nampula                                  | Pps     |
| napa                                     | california                       | napa                                     | p       |
| national_capital_territory_of_delhi      | india                            | national_capital_territory_of_delhi      | p       |
| nauru                                    | oceania                          | nauru                                    | p       |
| nayarit                                  | mexico                           | nayarit                                  | p       |
| negros_island_region                     | philippines                      | negros_island_region                     | Pps     |
| nenets_autonomous_okrug                  | northwestern_federal_district    | nenets_autonomous_okrug                  | Pps     |
| netherlands                              | europe                           | netherlands                              | p       |
| neuchatel                                | switzerland                      | neuchatel                                | p       |
| neuquen                                  | argentina                        | neuquen                                  | Pps     |
| nevada                                   | california                       | nevada                                   | p       |
| new-york                                 | us-northeast                     | new-york                                 | Pps     |
| new_brunswick                            | canada                           | new_brunswick                            | Pps     |
| new_caledonia                            | oceania                          | new_caledonia                            | p       |
| new_south_wales                          | australia                        | new_south_wales                          | p       |
| new_south_wales_central_west             | new_south_wales                  | new_south_wales_central_west             | p       |
| new_south_wales_north_western            | new_south_wales                  | new_south_wales_north_western            | p       |
| new_south_wales_northern                 | new_south_wales                  | new_south_wales_northern                 | p       |
| new_york_city                            | new-york                         | new_york_city                            | Pps     |
| newfoundland_and_labrador                | canada                           | newfoundland_and_labrador                | Pps     |
| niassa                                   | mozambique                       | niassa                                   | Pps     |
| nicaragua                                | central-america                  | nicaragua                                | p       |
| nidwalden                                | switzerland                      | nidwalden                                | p       |
| niederosterreich                         | austria                          | niederosterreich                         | p       |
| nievre                                   | bourgogne                        | nievre                                   | p       |
| niger                                    | africa                           | niger                                    | Pps     |
| nigeria                                  | africa                           | nigeria                                  | Pps     |
| nigeria_north_central                    | nigeria                          | nigeria_north_central                    | Pps     |
| nigeria_north_east                       | nigeria                          | nigeria_north_east                       | Pps     |
| nigeria_north_west                       | nigeria                          | nigeria_north_west                       | Pps     |
| nigeria_south_east                       | nigeria                          | nigeria_south_east                       | Pps     |
| nigeria_south_south                      | nigeria                          | nigeria_south_south                      | Pps     |
| nigeria_south_west                       | nigeria                          | nigeria_south_west                       | Pps     |
| ningxia                                  | china                            | ningxia                                  | Pps     |
| nitriansky                               | slovakia                         | nitriansky                               | p       |
| niue                                     | oceania                          | niue                                     | p       |
| nizhny_novgorod_oblast                   | volga_federal_district           | nizhny_novgorod_oblast                   | Pps     |
| noord_brabant                            | netherlands                      | noord_brabant                            | p       |
| noord_holland                            | netherlands                      | noord_holland                            | p       |
| nord                                     | nord_pas_de_calais               | nord                                     | p       |
| nord_du_quebec                           | quebec                           | nord_du_quebec                           | Pps     |
| nord_pas_de_calais                       | france                           | nord_pas_de_calais                       | p       |
| nordland                                 | norway                           | nordland                                 | p       |
| nordrhein_westfalen                      | germany                          | nordrhein_westfalen                      | p       |
| norfolk_island                           | australia                        | norfolk_island                           | p       |
| norrbotten                               | sweden                           | norrbotten                               | p       |
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
| north_karelia                            | finland                          | north_karelia                            | p       |
| north_maluku                             | indonesia                        | north_maluku                             | Pps     |
| north_metro                              | georgia                          | north_metro                              | Pps     |
| north_ossetia_alania_republic            | north_caucasian_federal_district | north_ossetia_alania_republic            | Pps     |
| north_ostrobothnia                       | finland                          | north_ostrobothnia                       | p       |
| north_savo                               | finland                          | north_savo                               | p       |
| north_sea                                | seas                             | north_sea                                | p       |
| north_sulawesi                           | indonesia                        | north_sulawesi                           | Pps     |
| north_sumatra                            | indonesia                        | north_sumatra                            | Pps     |
| northeast                                | brazil                           | northeast                                |         |
| northeastern_ontario                     | ontario                          | northeastern_ontario                     | Pps     |
| northern_cape                            | south_africa                     | northern_cape                            | Pps     |
| northern_ireland                         | united_kingdom                   | northern_ireland                         | p       |
| northern_lower                           | michigan                         | northern_lower                           | Pps     |
| northern_mariana_islands                 | oceania                          | northern_mariana_islands                 | p       |
| northern_mindanao                        | philippines                      | northern_mindanao                        | Pps     |
| northern_territory                       | australia                        | northern_territory                       | p       |
| northwest_territories                    | canada                           | northwest_territories                    | Pps     |
| northwestern_federal_district            | russia                           | northwestern_federal_district            | Pps     |
| northwestern_ontario                     | ontario                          | northwestern_ontario                     | Pps     |
| norway                                   | europe                           | norway                                   | p       |
| nova_scotia                              | canada                           | nova_scotia                              | Pps     |
| novgorod_oblast                          | northwestern_federal_district    | novgorod_oblast                          | Pps     |
| novosibirsk_oblast                       | siberian_federal_district        | novosibirsk_oblast                       | Pps     |
| nuble                                    | chile                            | nuble                                    | Pps     |
| nuevo_leon                               | mexico                           | nuevo_leon                               | p       |
| nunavut                                  | canada                           | nunavut                                  | Pps     |
| o_higgins                                | chile                            | o_higgins                                | Pps     |
| oaxaca                                   | mexico                           | oaxaca                                   | p       |
| oberosterreich                           | austria                          | oberosterreich                           | p       |
| obwalden                                 | switzerland                      | obwalden                                 | p       |
| oceania                                  |                                  | oceania                                  | Pps     |
| oceania_france_taaf                      | oceania                          | oceania_france_taaf                      | p       |
| odessa_oblast                            | ukraine                          | odessa_oblast                            | p       |
| odisha                                   | india                            | odisha                                   | p       |
| oestfold                                 | norway                           | oestfold                                 | p       |
| oise                                     | picardie                         | oise                                     | p       |
| olomoucky                                | czech_republic                   | olomoucky                                | p       |
| oman                                     | asia                             | oman                                     | Pps     |
| omsk_oblast                              | siberian_federal_district        | omsk_oblast                              | Pps     |
| ontario                                  | canada                           | ontario                                  | Pps     |
| opolskie                                 | poland                           | opolskie                                 | p       |
| oppland                                  | norway                           | oppland                                  | p       |
| orange                                   | california                       | orange                                   | p       |
| orebro                                   | sweden                           | orebro                                   | p       |
| orenburg_oblast                          | volga_federal_district           | orenburg_oblast                          | Pps     |
| orne                                     | basse_normandie                  | orne                                     | p       |
| oryol_oblast                             | central_federal_district         | oryol_oblast                             | Pps     |
| oslo                                     | norway                           | oslo                                     | p       |
| ostergotland                             | sweden                           | ostergotland                             | p       |
| ostrobothnia                             | finland                          | ostrobothnia                             | p       |
| ourense                                  | galicia                          | ourense                                  | p       |
| outaouais                                | quebec                           | outaouais                                | Pps     |
| overijssel                               | netherlands                      | overijssel                               | p       |
| paijat_hame                              | finland                          | paijat_hame                              | p       |
| palau                                    | oceania                          | palau                                    | p       |
| palencia                                 | castilla_y_leon                  | palencia                                 | p       |
| palestine                                | asia                             | palestine                                | Pps     |
| panama                                   | central-america                  | panama                                   | p       |
| panhandle                                | florida                          | panhandle                                | Pps     |
| papua                                    | indonesia                        | papua                                    | Pps     |
| papua_new_guinea                         | oceania                          | papua_new_guinea                         | p       |
| para                                     | north                            | para                                     | Pps     |
| paraguay                                 | south-america                    | paraguay                                 | Pps     |
| paraiba                                  | northeast                        | paraiba                                  | Pps     |
| parana                                   | south                            | parana                                   | Pps     |
| pardubicky                               | czech_republic                   | pardubicky                               | p       |
| paris                                    | ile_de_france                    | paris                                    | p       |
| pas_de_calais                            | nord_pas_de_calais               | pas_de_calais                            | p       |
| pays_de_la_loire                         | france                           | pays_de_la_loire                         | p       |
| penza_oblast                             | volga_federal_district           | penza_oblast                             | Pps     |
| perm_krai                                | volga_federal_district           | perm_krai                                | Pps     |
| pernambuco                               | northeast                        | pernambuco                               | Pps     |
| philippines                              | asia                             | philippines                              | Pps     |
| piaui                                    | northeast                        | piaui                                    | Pps     |
| picardie                                 | france                           | picardie                                 | p       |
| piedmont_triad                           | north-carolina                   | piedmont_triad                           | Pps     |
| piemonte                                 | italy                            | piemonte                                 | p       |
| pirkanmaa                                | finland                          | pirkanmaa                                | p       |
| pitcairn                                 | oceania                          | pitcairn                                 | p       |
| placer                                   | california                       | placer                                   | p       |
| plumas                                   | california                       | plumas                                   | p       |
| plzensky                                 | czech_republic                   | plzensky                                 | p       |
| podkarpackie                             | poland                           | podkarpackie                             | p       |
| podlaskie                                | poland                           | podlaskie                                | p       |
| poitou_charentes                         | france                           | poitou_charentes                         | p       |
| poland                                   | europe                           | poland                                   | p       |
| poltava_oblast                           | ukraine                          | poltava_oblast                           | p       |
| polygons-merge                           | merge                            | polygons-merge                           |         |
| polygons-merge_france_taaf               | polygons-merge                   | polygons-merge_france_taaf               | p       |
| polynesie                                | oceania                          | polynesie                                | p       |
| pomorskie                                | poland                           | pomorskie                                | p       |
| pontevedra                               | galicia                          | pontevedra                               | p       |
| portalegre                               | portugal                         | portalegre                               | p       |
| porto                                    | portugal                         | porto                                    | p       |
| portugal                                 | europe                           | portugal                                 | p       |
| praha                                    | czech_republic                   | praha                                    | p       |
| presovsky                                | slovakia                         | presovsky                                | p       |
| primorsky_krai                           | far_eastern_federal_district     | primorsky_krai                           | Pps     |
| prince_edward_island                     | canada                           | prince_edward_island                     | Pps     |
| provence_alpes_cote_d_azur               | france                           | provence_alpes_cote_d_azur               | p       |
| pskov_oblast                             | northwestern_federal_district    | pskov_oblast                             | Pps     |
| puducherry                               | india                            | puducherry                               | p       |
| puebla                                   | mexico                           | puebla                                   | p       |
| puerto_rico                              | central-america                  | puerto_rico                              | p       |
| puglia                                   | italy                            | puglia                                   | p       |
| punjab                                   | india                            | punjab                                   | p       |
| puy_de_dome                              | auvergne                         | puy_de_dome                              | p       |
| pyrenees_atlantiques                     | aquitaine                        | pyrenees_atlantiques                     | p       |
| pyrenees_orientales                      | languedoc_roussillon             | pyrenees_orientales                      | p       |
| qatar                                    | asia                             | qatar                                    | Pps     |
| qinghai                                  | china                            | qinghai                                  | Pps     |
| quebec                                   | canada                           | quebec                                   | Pps     |
| queensland                               | australia                        | queensland                               | p       |
| queretaro                                | mexico                           | queretaro                                | p       |
| quintana_roo                             | mexico                           | quintana_roo                             | p       |
| rajasthan                                | india                            | rajasthan                                | p       |
| region_de_murcia                         | spain                            | region_de_murcia                         | p       |
| reunion                                  | africa                           | reunion                                  | Pps     |
| rhone                                    | rhone_alpes                      | rhone                                    | p       |
| rhone_alpes                              | france                           | rhone_alpes                              | p       |
| riau                                     | indonesia                        | riau                                     | Pps     |
| riau_islands                             | indonesia                        | riau_islands                             | Pps     |
| richmond                                 | virginia                         | richmond                                 | Pps     |
| richmond_tweed                           | new_south_wales                  | richmond_tweed                           | p       |
| rio-de-janeiro                           | southeast                        | rio-de-janeiro                           | Pps     |
| rio-grande-do-norte                      | northeast                        | rio-grande-do-norte                      | Pps     |
| rio-grande-do-sul                        | south                            | rio-grande-do-sul                        | Pps     |
| rio_negro                                | argentina                        | rio_negro                                | Pps     |
| riverside                                | california                       | riverside                                | p       |
| rivne_oblast                             | ukraine                          | rivne_oblast                             | p       |
| rogaland                                 | norway                           | rogaland                                 | p       |
| rondonia                                 | north                            | rondonia                                 | Pps     |
| roraima                                  | north                            | roraima                                  | Pps     |
| rostov_oblast                            | southern_federal_district        | rostov_oblast                            | Pps     |
| russia                                   |                                  | russia                                   | Pps     |
| rwanda                                   | africa                           | rwanda                                   | Pps     |
| ryazan_oblast                            | central_federal_district         | ryazan_oblast                            | Pps     |
| sacramento                               | california                       | sacramento                               | p       |
| saguenay_lac_saint_jean                  | quebec                           | saguenay_lac_saint_jean                  | Pps     |
| saint_barthelemy                         | central-america                  | saint_barthelemy                         | p       |
| saint_gallen                             | switzerland                      | saint_gallen                             | p       |
| saint_helena_ascension_tristan_da_cunha  | africa                           | saint_helena_ascension_tristan_da_cunha  | Pps     |
| saint_kitts_and_nevis                    | central-america                  | saint_kitts_and_nevis                    | p       |
| saint_lucia                              | central-america                  | saint_lucia                              | p       |
| saint_martin                             | central-america                  | saint_martin                             | p       |
| saint_petersburg                         | northwestern_federal_district    | saint_petersburg                         | Pps     |
| saint_pierre_et_miquelon                 | north-america                    | saint_pierre_et_miquelon                 | Pps     |
| saint_vincent_and_the_grenadines         | central-america                  | saint_vincent_and_the_grenadines         | p       |
| sakha_republic                           | far_eastern_federal_district     | sakha_republic                           | Pps     |
| sakhalin_oblast                          | far_eastern_federal_district     | sakhalin_oblast                          | Pps     |
| salamanca                                | castilla_y_leon                  | salamanca                                | p       |
| salem                                    | virginia                         | salem                                    | Pps     |
| salta                                    | argentina                        | salta                                    | Pps     |
| salzburg                                 | austria                          | salzburg                                 | p       |
| samara_oblast                            | volga_federal_district           | samara_oblast                            | Pps     |
| samoa                                    | oceania                          | samoa                                    | p       |
| san_benito                               | california                       | san_benito                               | p       |
| san_bernardino                           | california                       | san_bernardino                           | p       |
| san_diego                                | california                       | san_diego                                | p       |
| san_francisco                            | california                       | san_francisco                            | p       |
| san_joaquin                              | california                       | san_joaquin                              | p       |
| san_juan                                 | argentina                        | san_juan                                 | Pps     |
| san_luis                                 | argentina                        | san_luis                                 | Pps     |
| san_luis_obispo                          | california                       | san_luis_obispo                          | p       |
| san_luis_potosi                          | mexico                           | san_luis_potosi                          | p       |
| san_marino                               | europe                           | san_marino                               | p       |
| san_mateo                                | california                       | san_mateo                                | p       |
| santa-catarina                           | south                            | santa-catarina                           | Pps     |
| santa_barbara                            | california                       | santa_barbara                            | p       |
| santa_clara                              | california                       | santa_clara                              | p       |
| santa_cruz_de_tenerife                   | canarias                         | santa_cruz_de_tenerife                   | Pps     |
| santa_fe                                 | argentina                        | santa_fe                                 | Pps     |
| santarem                                 | portugal                         | santarem                                 | p       |
| santiago                                 | chile                            | santiago                                 | Pps     |
| santiago_del_estero                      | argentina                        | santiago_del_estero                      | Pps     |
| sao-paulo                                | southeast                        | sao-paulo                                | Pps     |
| sao_tome_and_principe                    | africa                           | sao_tome_and_principe                    | Pps     |
| saone_et_loire                           | bourgogne                        | saone_et_loire                           | p       |
| saratov_oblast                           | volga_federal_district           | saratov_oblast                           | Pps     |
| sardegna                                 | italy                            | sardegna                                 | p       |
| sarthe                                   | pays_de_la_loire                 | sarthe                                   | p       |
| saskatchewan                             | canada                           | saskatchewan                             | Pps     |
| satakunta                                | finland                          | satakunta                                | p       |
| saudi_arabia                             | asia                             | saudi_arabia                             | Pps     |
| savoie                                   | rhone_alpes                      | savoie                                   | p       |
| schaffhausen                             | switzerland                      | schaffhausen                             | p       |
| schwyz                                   | switzerland                      | schwyz                                   | p       |
| seas                                     | europe                           | seas                                     |         |
| segovia                                  | castilla_y_leon                  | segovia                                  | p       |
| seine_et_marne                           | ile_de_france                    | seine_et_marne                           | p       |
| seine_maritime                           | haute_normandie                  | seine_maritime                           | p       |
| seine_saint_denis                        | ile_de_france                    | seine_saint_denis                        | p       |
| senegal                                  | africa                           | senegal                                  | Pps     |
| sergipe                                  | northeast                        | sergipe                                  | Pps     |
| setubal                                  | portugal                         | setubal                                  | p       |
| sevilla                                  | andalucia                        | sevilla                                  | p       |
| seychelles                               | africa                           | seychelles                               | Pps     |
| shaanxi                                  | china                            | shaanxi                                  | Pps     |
| shandong                                 | china                            | shandong                                 | Pps     |
| shanghai                                 | china                            | shanghai                                 | Pps     |
| shanxi                                   | china                            | shanxi                                   | Pps     |
| shasta                                   | california                       | shasta                                   | p       |
| shikoku                                  | japan                            | shikoku                                  | Pps     |
| siberian_federal_district                | russia                           | siberian_federal_district                | Pps     |
| sichuan                                  | china                            | sichuan                                  | Pps     |
| sicilia                                  | italy                            | sicilia                                  | p       |
| sierra                                   | california                       | sierra                                   | p       |
| sikkim                                   | india                            | sikkim                                   | p       |
| sinaloa                                  | mexico                           | sinaloa                                  | p       |
| singapore                                | asia                             | singapore                                | Pps     |
| sint_maarten                             | central-america                  | sint_maarten                             | p       |
| siskiyou                                 | california                       | siskiyou                                 | p       |
| skane                                    | sweden                           | skane                                    | p       |
| slaskie                                  | poland                           | slaskie                                  | p       |
| slovakia                                 | europe                           | slovakia                                 | p       |
| smolensk_oblast                          | central_federal_district         | smolensk_oblast                          | Pps     |
| soccsksargen                             | philippines                      | soccsksargen                             | Pps     |
| sodermanland                             | sweden                           | sodermanland                             | p       |
| sofala                                   | mozambique                       | sofala                                   | Pps     |
| sogn_og_fjordane                         | norway                           | sogn_og_fjordane                         | p       |
| solano                                   | california                       | solano                                   | p       |
| solomon_islands                          | oceania                          | solomon_islands                          | p       |
| solothurn                                | switzerland                      | solothurn                                | p       |
| somme                                    | picardie                         | somme                                    | p       |
| sonoma                                   | california                       | sonoma                                   | p       |
| sonora                                   | mexico                           | sonora                                   | p       |
| soria                                    | castilla_y_leon                  | soria                                    | p       |
| south                                    | brazil                           | south                                    |         |
| south-america                            |                                  | south-america                            | Pps     |
| south_africa                             | africa                           | south_africa                             | Pps     |
| south_africa_north_west                  | south_africa                     | south_africa_north_west                  | Pps     |
| south_australia                          | australia                        | south_australia                          | p       |
| south_east_region                        | new_south_wales                  | south_east_region                        | p       |
| south_georgia_and_south_sandwich         | south-america                    | south_georgia_and_south_sandwich         | Pps     |
| south_kalimantan                         | indonesia                        | south_kalimantan                         | Pps     |
| south_karelia                            | finland                          | south_karelia                            | p       |
| south_ostrobothnia                       | finland                          | south_ostrobothnia                       | p       |
| south_papua                              | indonesia                        | south_papua                              | Pps     |
| south_savo                               | finland                          | south_savo                               | p       |
| south_sudan                              | africa                           | south_sudan                              | Pps     |
| south_sulawesi                           | indonesia                        | south_sulawesi                           | Pps     |
| south_sumatra                            | indonesia                        | south_sumatra                            | Pps     |
| southeast                                | brazil                           | southeast                                |         |
| southeast_sulawesi                       | indonesia                        | southeast_sulawesi                       | Pps     |
| southeastern_anatolia                    | turkey                           | southeastern_anatolia                    | p       |
| southern_federal_district                | russia                           | southern_federal_district                | Pps     |
| southern_federal_district_sevastopol     | southern_federal_district        | southern_federal_district_sevastopol     | Pps     |
| southern_highlands                       | tanzania                         | southern_highlands                       | Pps     |
| southern_tier                            | new-york                         | southern_tier                            | Pps     |
| southwest_finland                        | finland                          | southwest_finland                        | p       |
| southwest_papua                          | indonesia                        | southwest_papua                          | Pps     |
| southwestern_ontario                     | ontario                          | southwestern_ontario                     | Pps     |
| spain                                    | europe                           | spain                                    | p       |
| spain_la_rioja                           | spain                            | spain_la_rioja                           | p       |
| stanislaus                               | california                       | stanislaus                               | p       |
| state_of_mexico                          | mexico                           | state_of_mexico                          | p       |
| stavropol_krai                           | north_caucasian_federal_district | stavropol_krai                           | Pps     |
| steiermark                               | austria                          | steiermark                               | p       |
| stockholm                                | sweden                           | stockholm                                | p       |
| stredocesky                              | czech_republic                   | stredocesky                              | p       |
| sudan                                    | africa                           | sudan                                    | Pps     |
| sumy_oblast                              | ukraine                          | sumy_oblast                              | p       |
| suncoast                                 | florida                          | suncoast                                 | Pps     |
| suriname                                 | south-america                    | suriname                                 | Pps     |
| sutter                                   | california                       | sutter                                   | p       |
| svalbard                                 | norway                           | svalbard                                 | p       |
| sverdlovsk_oblast                        | ural_federal_district            | sverdlovsk_oblast                        | Pps     |
| swaziland                                | africa                           | swaziland                                | Pps     |
| sweden                                   | europe                           | sweden                                   | p       |
| swietokrzyskie                           | poland                           | swietokrzyskie                           | p       |
| switzerland                              | europe                           | switzerland                              | p       |
| switzerland_jura                         | switzerland                      | switzerland_jura                         | p       |
| sydney_surrounds                         | new_south_wales                  | sydney_surrounds                         | p       |
| tabasco                                  | mexico                           | tabasco                                  | p       |
| tamaulipas                               | mexico                           | tamaulipas                               | p       |
| tambov_oblast                            | central_federal_district         | tambov_oblast                            | Pps     |
| tamil_nadu                               | india                            | tamil_nadu                               | p       |
| tanzania                                 | africa                           | tanzania                                 | Pps     |
| tanzania_central                         | tanzania                         | tanzania_central                         | Pps     |
| tanzania_lake                            | tanzania                         | tanzania_lake                            | Pps     |
| tanzania_northern                        | tanzania                         | tanzania_northern                        | Pps     |
| tanzania_western                         | tanzania                         | tanzania_western                         | Pps     |
| tarapaca                                 | chile                            | tarapaca                                 | Pps     |
| tarn                                     | midi_pyrenees                    | tarn                                     | p       |
| tarn_et_garonne                          | midi_pyrenees                    | tarn_et_garonne                          | p       |
| tarragona                                | catalunya                        | tarragona                                | p       |
| tasmania                                 | australia                        | tasmania                                 | p       |
| tatarstan_republic                       | volga_federal_district           | tatarstan_republic                       | Pps     |
| tehama                                   | california                       | tehama                                   | p       |
| telangana                                | india                            | telangana                                | p       |
| telemark                                 | norway                           | telemark                                 | p       |
| ternopil_oblast                          | ukraine                          | ternopil_oblast                          | p       |
| territoire_de_belfort                    | franche_comte                    | territoire_de_belfort                    | p       |
| teruel                                   | aragon                           | teruel                                   | p       |
| tete                                     | mozambique                       | tete                                     | Pps     |
| texas                                    | us-south                         | texas                                    | Pps     |
| texas_central                            | texas                            | texas_central                            | Pps     |
| texas_north                              | texas                            | texas_north                              | Pps     |
| texas_northwest                          | texas                            | texas_northwest                          | Pps     |
| texas_south                              | texas                            | texas_south                              | Pps     |
| texas_southeast                          | texas                            | texas_southeast                          | Pps     |
| texas_west                               | texas                            | texas_west                               | Pps     |
| the_riverina                             | new_south_wales                  | the_riverina                             | p       |
| thurgau                                  | switzerland                      | thurgau                                  | p       |
| tianjin                                  | china                            | tianjin                                  | Pps     |
| tibet                                    | china                            | tibet                                    | Pps     |
| ticino                                   | switzerland                      | ticino                                   | p       |
| tierra_del_fuego                         | argentina                        | tierra_del_fuego                         | Pps     |
| tirol                                    | austria                          | tirol                                    | p       |
| tlaxcala                                 | mexico                           | tlaxcala                                 | p       |
| tocantins                                | north                            | tocantins                                | Pps     |
| togo                                     | africa                           | togo                                     | Pps     |
| tohoku                                   | japan                            | tohoku                                   | Pps     |
| tokelau                                  | oceania                          | tokelau                                  | p       |
| toledo                                   | castilla_la_mancha               | toledo                                   | p       |
| tomsk_oblast                             | siberian_federal_district        | tomsk_oblast                             | Pps     |
| tonga                                    | oceania                          | tonga                                    | p       |
| toscana                                  | italy                            | toscana                                  | p       |
| treasure_coast                           | florida                          | treasure_coast                           | Pps     |
| trenciansky                              | slovakia                         | trenciansky                              | p       |
| trentino_alto_adige                      | italy                            | trentino_alto_adige                      | p       |
| trinidad_and_tobago                      | central-america                  | trinidad_and_tobago                      | p       |
| trinity                                  | california                       | trinity                                  | p       |
| tripura                                  | india                            | tripura                                  | p       |
| trnavsky                                 | slovakia                         | trnavsky                                 | p       |
| troendelag                               | norway                           | troendelag                               | p       |
| troms                                    | norway                           | troms                                    | p       |
| tucuman                                  | argentina                        | tucuman                                  | Pps     |
| tula_oblast                              | central_federal_district         | tula_oblast                              | Pps     |
| tulare                                   | california                       | tulare                                   | p       |
| tunisia                                  | africa                           | tunisia                                  | Pps     |
| tuolumne                                 | california                       | tuolumne                                 | p       |
| turkey                                   | europe                           | turkey                                   | p       |
| turks_and_caicos_islands                 | central-america                  | turks_and_caicos_islands                 | p       |
| tuva_republic                            | siberian_federal_district        | tuva_republic                            | Pps     |
| tuvalu                                   | oceania                          | tuvalu                                   | p       |
| tver_oblast                              | central_federal_district         | tver_oblast                              | Pps     |
| tyumen_oblast                            | ural_federal_district            | tyumen_oblast                            | Pps     |
| udmurt_republic                          | volga_federal_district           | udmurt_republic                          | Pps     |
| uganda                                   | africa                           | uganda                                   | Pps     |
| uganda_central                           | uganda                           | uganda_central                           | Pps     |
| uganda_eastern                           | uganda                           | uganda_eastern                           | Pps     |
| uganda_northern                          | uganda                           | uganda_northern                          | Pps     |
| uganda_western                           | uganda                           | uganda_western                           | Pps     |
| ukraine                                  | europe                           | ukraine                                  | p       |
| ukraine_sevastopol                       | ukraine                          | ukraine_sevastopol                       | p       |
| ulyanovsk_oblast                         | volga_federal_district           | ulyanovsk_oblast                         | Pps     |
| umbria                                   | italy                            | umbria                                   | p       |
| united_arab_emirates                     | asia                             | united_arab_emirates                     | Pps     |
| united_kingdom                           | europe                           | united_kingdom                           | p       |
| united_states_virgin_islands             | central-america                  | united_states_virgin_islands             | p       |
| upper_peninsula                          | michigan                         | upper_peninsula                          | Pps     |
| uppsala                                  | sweden                           | uppsala                                  | p       |
| ural_federal_district                    | russia                           | ural_federal_district                    | Pps     |
| uri                                      | switzerland                      | uri                                      | p       |
| us-midwest                               | north-america                    | us-midwest                               | Pps     |
| us-northeast                             | north-america                    | us-northeast                             | Pps     |
| us-south                                 | north-america                    | us-south                                 | Pps     |
| us-west                                  | north-america                    | us-west                                  | Pps     |
| usa_virgin_islands                       | central-america                  | usa_virgin_islands                       | p       |
| ustecky                                  | czech_republic                   | ustecky                                  | p       |
| utrecht                                  | netherlands                      | utrecht                                  | p       |
| uttar_pradesh                            | india                            | uttar_pradesh                            | p       |
| uttarakhand                              | india                            | uttarakhand                              | p       |
| uusimaa                                  | finland                          | uusimaa                                  | p       |
| val_d_oise                               | ile_de_france                    | val_d_oise                               | p       |
| val_de_marne                             | ile_de_france                    | val_de_marne                             | p       |
| valais                                   | switzerland                      | valais                                   | p       |
| valencia                                 | comunitat_valenciana             | valencia                                 | p       |
| valladolid                               | castilla_y_leon                  | valladolid                               | p       |
| valle_aosta                              | italy                            | valle_aosta                              | p       |
| valparaiso                               | chile                            | valparaiso                               | Pps     |
| vanuatu                                  | oceania                          | vanuatu                                  | p       |
| var                                      | provence_alpes_cote_d_azur       | var                                      | p       |
| varmland                                 | sweden                           | varmland                                 | p       |
| vasterbotten                             | sweden                           | vasterbotten                             | p       |
| vasternorrland                           | sweden                           | vasternorrland                           | p       |
| vastmanland                              | sweden                           | vastmanland                              | p       |
| vastra_gotaland                          | sweden                           | vastra_gotaland                          | p       |
| vatican_city                             | europe                           | vatican_city                             | p       |
| vaucluse                                 | provence_alpes_cote_d_azur       | vaucluse                                 | p       |
| vaud                                     | switzerland                      | vaud                                     | p       |
| vendee                                   | pays_de_la_loire                 | vendee                                   | p       |
| veneto                                   | italy                            | veneto                                   | p       |
| venezuela                                | south-america                    | venezuela                                | Pps     |
| ventura                                  | california                       | ventura                                  | p       |
| veracruz                                 | mexico                           | veracruz                                 | p       |
| vest-agder                               | norway                           | vest-agder                               | p       |
| vestfold                                 | norway                           | vestfold                                 | p       |
| viana_do_castelo                         | portugal                         | viana_do_castelo                         | p       |
| victoria                                 | australia                        | victoria                                 | p       |
| vienne                                   | poitou_charentes                 | vienne                                   | p       |
| vila_real                                | portugal                         | vila_real                                | p       |
| vinnytsia_oblast                         | ukraine                          | vinnytsia_oblast                         | p       |
| virginia                                 | us-south                         | virginia                                 | Pps     |
| viseu                                    | portugal                         | viseu                                    | p       |
| vizcaya                                  | euskadi                          | vizcaya                                  | p       |
| vladimir_oblast                          | central_federal_district         | vladimir_oblast                          | Pps     |
| volga_federal_district                   | russia                           | volga_federal_district                   | Pps     |
| volgograd_oblast                         | southern_federal_district        | volgograd_oblast                         | Pps     |
| vologda_oblast                           | northwestern_federal_district    | vologda_oblast                           | Pps     |
| volyn_oblast                             | ukraine                          | volyn_oblast                             | p       |
| vorarlberg                               | austria                          | vorarlberg                               | p       |
| voronezh_oblast                          | central_federal_district         | voronezh_oblast                          | Pps     |
| vosges                                   | lorraine                         | vosges                                   | p       |
| vysocina                                 | czech_republic                   | vysocina                                 | p       |
| wallis_et_futuna                         | oceania                          | wallis_et_futuna                         | p       |
| wallonia_french_community                | belgium                          | wallonia_french_community                | p       |
| wallonia_german_community                | belgium                          | wallonia_german_community                | p       |
| warminsko_mazurskie                      | poland                           | warminsko_mazurskie                      | p       |
| west_bengal                              | india                            | west_bengal                              | p       |
| west_flanders                            | flanders                         | west_flanders                            | p       |
| west_java                                | indonesia                        | west_java                                | Pps     |
| west_kalimantan                          | indonesia                        | west_kalimantan                          | Pps     |
| west_midlands                            | england                          | west_midlands                            | p       |
| west_nusa_tenggara                       | indonesia                        | west_nusa_tenggara                       | Pps     |
| west_papua                               | indonesia                        | west_papua                               | Pps     |
| west_sulawesi                            | indonesia                        | west_sulawesi                            | Pps     |
| west_sumatra                             | indonesia                        | west_sumatra                             | Pps     |
| western_australia                        | australia                        | western_australia                        | p       |
| western_cape                             | south_africa                     | western_cape                             | Pps     |
| western_new_york                         | new-york                         | western_new_york                         | Pps     |
| western_sahara                           | africa                           | western_sahara                           | Pps     |
| western_visayas                          | philippines                      | western_visayas                          | Pps     |
| wielkopolskie                            | poland                           | wielkopolskie                            | p       |
| wien                                     | austria                          | wien                                     | p       |
| wytheville                               | virginia                         | wytheville                               | Pps     |
| xinjiang                                 | china                            | xinjiang                                 | Pps     |
| yamalo_nenets_autonomous_okrug           | ural_federal_district            | yamalo_nenets_autonomous_okrug           | Pps     |
| yaroslavl_oblast                         | central_federal_district         | yaroslavl_oblast                         | Pps     |
| yogyakarta                               | indonesia                        | yogyakarta                               | Pps     |
| yolo                                     | california                       | yolo                                     | p       |
| yonne                                    | bourgogne                        | yonne                                    | p       |
| yorkshire_and_the_humber                 | england                          | yorkshire_and_the_humber                 | p       |
| yuba                                     | california                       | yuba                                     | p       |
| yucatan                                  | mexico                           | yucatan                                  | p       |
| yukon                                    | canada                           | yukon                                    | Pps     |
| yunnan                                   | china                            | yunnan                                   | Pps     |
| yvelines                                 | ile_de_france                    | yvelines                                 | p       |
| zabaykalsky_krai                         | siberian_federal_district        | zabaykalsky_krai                         | Pps     |
| zacatecas                                | mexico                           | zacatecas                                | p       |
| zachodniopomorskie                       | poland                           | zachodniopomorskie                       | p       |
| zakarpattia_oblast                       | ukraine                          | zakarpattia_oblast                       | p       |
| zambezia                                 | mozambique                       | zambezia                                 | Pps     |
| zambia                                   | africa                           | zambia                                   | Pps     |
| zamboanga_peninsula                      | philippines                      | zamboanga_peninsula                      | Pps     |
| zamora                                   | castilla_y_leon                  | zamora                                   | p       |
| zanzibar                                 | tanzania                         | zanzibar                                 | Pps     |
| zaporizhia_oblast                        | ukraine                          | zaporizhia_oblast                        | p       |
| zaragoza                                 | aragon                           | zaragoza                                 | p       |
| zeeland                                  | netherlands                      | zeeland                                  | p       |
| zhejiang                                 | china                            | zhejiang                                 | Pps     |
| zhytomyr_oblast                          | ukraine                          | zhytomyr_oblast                          | p       |
| zilinsky                                 | slovakia                         | zilinsky                                 | p       |
| zimbabwe                                 | africa                           | zimbabwe                                 | Pps     |
| zlinsky                                  | czech_republic                   | zlinsky                                  | p       |
| zug                                      | switzerland                      | zug                                      | p       |
| zuid_holland                             | netherlands                      | zuid_holland                             | p       |
| zurich                                   | switzerland                      | zurich                                   | p       |
Total elements: 1199

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
| SHORTNAME | IS IN | LONG NAME | FORMATS |
|-----------|-------|-----------|---------|
Total elements: 0

