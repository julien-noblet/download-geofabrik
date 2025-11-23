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
| afghanistan                                 | Asia                        | Afghanistan                    | kHPpSs  |
| africa                                      |                             | Africa                         | kHPps   |
| albania                                     | Europe                      | Albania                        | kHPpSs  |
| alberta                                     | Canada                      | Alberta                        | kHPpSs  |
| algeria                                     | Africa                      | Algeria                        | kHPpSs  |
| alps                                        | Europe                      | Alps                           | kHPps   |
| alsace                                      | France                      | Alsace                         | kHPpSs  |
| american-oceania                            | Australia and Oceania       | American Oceania               | kHPpSs  |
| andalucia                                   | Spain                       | Andalucía                      | kHPpSs  |
| andorra                                     | Europe                      | Andorra                        | kHPpSs  |
| angola                                      | Africa                      | Angola                         | kHPpSs  |
| anhui                                       | China                       | Anhui                          | kHPpSs  |
| antarctica                                  |                             | Antarctica                     | kHPpSs  |
| aquitaine                                   | France                      | Aquitaine                      | kHPpSs  |
| aragon                                      | Spain                       | Aragón                         | kHPpSs  |
| argentina                                   | South America               | Argentina                      | kHPpSs  |
| armenia                                     | Asia                        | Armenia                        | kHPpSs  |
| arnsberg-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Arnsberg      | kHPpSs  |
| asia                                        |                             | Asia                           | kHPps   |
| asturias                                    | Spain                       | Asturias                       | kHPpSs  |
| australia                                   | Australia and Oceania       | Australia                      | kHPpSs  |
| australia-oceania                           |                             | Australia and Oceania          | kHPps   |
| austria                                     | Europe                      | Austria                        | kHPpSs  |
| auvergne                                    | France                      | Auvergne                       | kHPpSs  |
| azerbaijan                                  | Asia                        | Azerbaijan                     | kHPpSs  |
| azores                                      | Europe                      | Azores                         | kHPpSs  |
| baden-wuerttemberg                          | Germany                     | Baden-Württemberg              | kHPpSs  |
| bahamas                                     | Central America             | Bahamas                        | kHPpSs  |
| bangladesh                                  | Asia                        | Bangladesh                     | kHPpSs  |
| basse-normandie                             | France                      | Basse-Normandie                | kHPpSs  |
| bayern                                      | Germany                     | Bayern                         | kHPpSs  |
| bedfordshire                                | England                     | Bedfordshire                   | kHPpSs  |
| beijing                                     | China                       | Beijing                        | kHPpSs  |
| belarus                                     | Europe                      | Belarus                        | kHPpSs  |
| belgium                                     | Europe                      | Belgium                        | kHPpSs  |
| belize                                      | Central America             | Belize                         | kHPpSs  |
| benin                                       | Africa                      | Benin                          | kHPpSs  |
| berkshire                                   | England                     | Berkshire                      | kHPpSs  |
| berlin                                      | Germany                     | Berlin                         | kHPpSs  |
| bermuda                                     | United Kingdom              | Bermuda                        | kHPpSs  |
| bhutan                                      | Asia                        | Bhutan                         | kHPpSs  |
| bolivia                                     | South America               | Bolivia                        | kHPpSs  |
| bosnia-herzegovina                          | Europe                      | Bosnia-Herzegovina             | kHPpSs  |
| botswana                                    | Africa                      | Botswana                       | kHPpSs  |
| bourgogne                                   | France                      | Bourgogne                      | kHPpSs  |
| brandenburg                                 | Germany                     | Brandenburg (mit Berlin)       | kHPpSs  |
| brazil                                      | South America               | Brazil                         | kHPps   |
| bremen                                      | Germany                     | Bremen                         | kHPpSs  |
| bretagne                                    | France                      | Bretagne                       | kHPpSs  |
| bristol                                     | England                     | Bristol                        | kHPpSs  |
| britain-and-ireland                         | Europe                      | Britain and Ireland            | kHPps   |
| british-columbia                            | Canada                      | British Columbia               | kHPpSs  |
| buckinghamshire                             | England                     | Buckinghamshire                | kHPpSs  |
| bulgaria                                    | Europe                      | Bulgaria                       | kHPpSs  |
| burkina-faso                                | Africa                      | Burkina Faso                   | kHPpSs  |
| burundi                                     | Africa                      | Burundi                        | kHPpSs  |
| cambodia                                    | Asia                        | Cambodia                       | kHPpSs  |
| cambridgeshire                              | England                     | Cambridgeshire                 | kHPpSs  |
| cameroon                                    | Africa                      | Cameroon                       | kHPpSs  |
| canada                                      | North America               | Canada                         | kHPps   |
| canary-islands                              | Africa                      | Canary Islands                 | kHPpSs  |
| cantabria                                   | Spain                       | Cantabria                      | kHPpSs  |
| cape-verde                                  | Africa                      | Cape Verde                     | kHPpSs  |
| castilla-la-mancha                          | Spain                       | Castilla-La Mancha             | kHPpSs  |
| castilla-y-leon                             | Spain                       | Castilla y León                | kHPpSs  |
| cataluna                                    | Spain                       | Cataluña                       | kHPpSs  |
| central-african-republic                    | Africa                      | Central African Republic       | kHPpSs  |
| central-america                             |                             | Central America                | kHPps   |
| central-fed-district                        | Russian Federation          | Central Federal District       | kHPpSs  |
| central-zone                                | India                       | Central Zone                   | kHPpSs  |
| centre                                      | France                      | Centre                         | kHPpSs  |
| centro                                      | Italy                       | Centro                         | kHPpSs  |
| centro-oeste                                | Brazil                      | Centro-Oeste                   | kHPpSs  |
| ceuta                                       | Spain                       | Ceuta                          | kHPpSs  |
| chad                                        | Africa                      | Chad                           | kHPpSs  |
| champagne-ardenne                           | France                      | Champagne Ardenne              | kHPpSs  |
| cheshire                                    | England                     | Cheshire                       | kHPpSs  |
| chile                                       | South America               | Chile                          | kHPpSs  |
| china                                       | Asia                        | China                          | kHPpSs  |
| chongqing                                   | China                       | Chongqing                      | kHPpSs  |
| chubu                                       | Japan                       | Chūbu region                   | kHPpSs  |
| chugoku                                     | Japan                       | Chūgoku region                 | kHPpSs  |
| colombia                                    | South America               | Colombia                       | kHPpSs  |
| comores                                     | Africa                      | Comores                        | kHPpSs  |
| congo-brazzaville                           | Africa                      | Congo (Republic/Brazzaville)   | kHPpSs  |
| congo-democratic-republic                   | Africa                      | Congo (Democratic              | kHPpSs  |
|                                             |                             | Republic/Kinshasa)             |         |
| cook-islands                                | Australia and Oceania       | Cook Islands                   | kHPpSs  |
| cornwall                                    | England                     | Cornwall                       | kHPpSs  |
| corse                                       | France                      | Corse                          | kHPpSs  |
| costa-rica                                  | Central America             | Costa Rica                     | kHPpSs  |
| crimean-fed-district                        | Russian Federation          | Crimean Federal District       | kHPpSs  |
| croatia                                     | Europe                      | Croatia                        | kHPpSs  |
| cuba                                        | Central America             | Cuba                           | kHPpSs  |
| cumbria                                     | England                     | Cumbria                        | kHPpSs  |
| cyprus                                      | Europe                      | Cyprus                         | kHPpSs  |
| czech-republic                              | Europe                      | Czech Republic                 | kHPpSs  |
| dach                                        | Europe                      | Germany, Austria, Switzerland  | kHPps   |
| denmark                                     | Europe                      | Denmark                        | kHPpSs  |
| derbyshire                                  | England                     | Derbyshire                     | kHPpSs  |
| detmold-regbez                              | Nordrhein-Westfalen         | Regierungsbezirk Detmold       | kHPpSs  |
| devon                                       | England                     | Devon                          | kHPpSs  |
| djibouti                                    | Africa                      | Djibouti                       | kHPpSs  |
| dolnoslaskie                                | Poland                      | Województwo dolnośląskie<br    | kHPpSs  |
|                                             |                             | />(Lower Silesian Voivodeship) |         |
| dorset                                      | England                     | Dorset                         | kHPpSs  |
| drenthe                                     | Netherlands                 | Drenthe                        | kHPpSs  |
| duesseldorf-regbez                          | Nordrhein-Westfalen         | Regierungsbezirk Düsseldorf    | kHPpSs  |
| durham                                      | England                     | Durham                         | kHPpSs  |
| east-sussex                                 | England                     | East Sussex                    | kHPpSs  |
| east-timor                                  | Asia                        | East Timor                     | kHPpSs  |
| east-yorkshire-with-hull                    | England                     | East Yorkshire with Hull       | kHPpSs  |
| eastern-zone                                | India                       | Eastern Zone                   | kHPpSs  |
| ecuador                                     | South America               | Ecuador                        | kHPpSs  |
| egypt                                       | Africa                      | Egypt                          | kHPpSs  |
| el-salvador                                 | Central America             | El Salvador                    | kHPpSs  |
| enfield                                     | Greater London              | Enfield                        | kHPpSs  |
| england                                     | United Kingdom              | England                        | kHPpSs  |
| equatorial-guinea                           | Africa                      | Equatorial Guinea              | kHPpSs  |
| eritrea                                     | Africa                      | Eritrea                        | kHPpSs  |
| essex                                       | England                     | Essex                          | kHPpSs  |
| estonia                                     | Europe                      | Estonia                        | kHPpSs  |
| ethiopia                                    | Africa                      | Ethiopia                       | kHPpSs  |
| europe                                      |                             | Europe                         | kHPps   |
| extremadura                                 | Spain                       | Extremadura                    | kHPpSs  |
| falklands                                   | United Kingdom              | Falkland Islands               | kHPpSs  |
| far-eastern-fed-district                    | Russian Federation          | Far Eastern Federal District   | kHPpSs  |
| faroe-islands                               | Europe                      | Faroe Islands                  | kHPpSs  |
| fiji                                        | Australia and Oceania       | Fiji                           | kHPpSs  |
| finland                                     | Europe                      | Finland                        | kHPpSs  |
| flevoland                                   | Netherlands                 | Flevoland                      | kHPpSs  |
| france                                      | Europe                      | France                         | kHPps   |
| franche-comte                               | France                      | Franche Comte                  | kHPpSs  |
| freiburg-regbez                             | Baden-Württemberg           | Regierungsbezirk Freiburg      | kHPpSs  |
| friesland                                   | Netherlands                 | Friesland                      | kHPpSs  |
| fujian                                      | China                       | Fujian                         | kHPpSs  |
| gabon                                       | Africa                      | Gabon                          | kHPpSs  |
| galicia                                     | Spain                       | Galicia                        | kHPpSs  |
| gansu                                       | China                       | Gansu                          | kHPpSs  |
| gcc-states                                  | Asia                        | GCC States                     | kHPpSs  |
| gelderland                                  | Netherlands                 | Gelderland                     | kHPpSs  |
| georgia                                     | Europe                      | Georgia                        | kHPpSs  |
| germany                                     | Europe                      | Germany                        | kHPps   |
| ghana                                       | Africa                      | Ghana                          | kHPpSs  |
| gloucestershire                             | England                     | Gloucestershire                | kHPpSs  |
| great-britain                               | Europe                      | Great Britain                  | kHPps   |
| greater-london                              | England                     | Greater London                 | kHPpSs  |
| greater-manchester                          | England                     | Greater Manchester             | kHPpSs  |
| greece                                      | Europe                      | Greece                         | kHPpSs  |
| greenland                                   | North America               | Greenland                      | kHPpSs  |
| groningen                                   | Netherlands                 | Groningen                      | kHPpSs  |
| guadeloupe                                  | France                      | Guadeloupe                     | kHPpSs  |
| guangdong                                   | China                       | Guangdong (with Hong Kong and  | kHPpSs  |
|                                             |                             | Macau)                         |         |
| guangxi                                     | China                       | Guangxi                        | kHPpSs  |
| guatemala                                   | Central America             | Guatemala                      | kHPpSs  |
| guernsey-jersey                             | Europe                      | Guernsey and Jersey            | kHPpSs  |
| guinea                                      | Africa                      | Guinea                         | kHPpSs  |
| guinea-bissau                               | Africa                      | Guinea-Bissau                  | kHPpSs  |
| guizhou                                     | China                       | Guizhou                        | kHPpSs  |
| guyana                                      | South America               | Guyana                         | kHPpSs  |
| guyane                                      | France                      | Guyane                         | kHPpSs  |
| hainan                                      | China                       | Hainan                         | kHPpSs  |
| haiti-and-domrep                            | Central America             | Haiti and Dominican Republic   | kHPpSs  |
| hamburg                                     | Germany                     | Hamburg                        | kHPpSs  |
| hampshire                                   | England                     | Hampshire                      | kHPpSs  |
| haute-normandie                             | France                      | Haute-Normandie                | kHPpSs  |
| hebei                                       | China                       | Hebei (with Beijing and        | kHPpSs  |
|                                             |                             | Tianjin)                       |         |
| heilongjiang                                | China                       | Heilongjiang                   | kHPpSs  |
| henan                                       | China                       | Henan                          | kHPpSs  |
| herefordshire                               | England                     | Herefordshire                  | kHPpSs  |
| hertfordshire                               | England                     | Hertfordshire                  | kHPpSs  |
| hessen                                      | Germany                     | Hessen                         | kHPpSs  |
| hokkaido                                    | Japan                       | Hokkaidō                       | kHPpSs  |
| honduras                                    | Central America             | Honduras                       | kHPpSs  |
| hong-kong                                   | China                       | Hong Kong                      | kHPpSs  |
| hubei                                       | China                       | Hubei                          | kHPpSs  |
| hunan                                       | China                       | Hunan                          | kHPpSs  |
| hungary                                     | Europe                      | Hungary                        | kHPpSs  |
| iceland                                     | Europe                      | Iceland                        | kHPpSs  |
| ile-de-clipperton                           | Australia and Oceania       | Île de Clipperton              | kHPpSs  |
| ile-de-france                               | France                      | Ile-de-France                  | kHPpSs  |
| india                                       | Asia                        | India                          | kHPpSs  |
| indonesia                                   | Asia                        | Indonesia (with East Timor)    | kHPpSs  |
| inner-mongolia                              | China                       | Inner Mongolia                 | kHPpSs  |
| iran                                        | Asia                        | Iran                           | kHPpSs  |
| iraq                                        | Asia                        | Iraq                           | kHPpSs  |
| ireland-and-northern-ireland                | Europe                      | Ireland and Northern Ireland   | kHPpSs  |
| islas-baleares                              | Spain                       | Islas Baleares                 | kHPpSs  |
| isle-of-man                                 | Europe                      | Isle of Man                    | kHPpSs  |
| isle-of-wight                               | England                     | Isle of Wight                  | kHPpSs  |
| isole                                       | Italy                       | Isole                          | kHPpSs  |
| israel-and-palestine                        | Asia                        | Israel and Palestine           | kHPpSs  |
| italy                                       | Europe                      | Italy                          | kHPps   |
| ivory-coast                                 | Africa                      | Ivory Coast                    | kHPpSs  |
| jamaica                                     | Central America             | Jamaica                        | kHPpSs  |
| japan                                       | Asia                        | Japan                          | kHPps   |
| java                                        | Indonesia (with East Timor) | Java                           | kHPpSs  |
| jiangsu                                     | China                       | Jiangsu                        | kHPpSs  |
| jiangxi                                     | China                       | Jiangxi                        | kHPpSs  |
| jilin                                       | China                       | Jilin                          | kHPpSs  |
| jordan                                      | Asia                        | Jordan                         | kHPpSs  |
| kalimantan                                  | Indonesia (with East Timor) | Kalimantan                     | kHPpSs  |
| kaliningrad                                 | Russian Federation          | Kaliningrad                    | kHPpSs  |
| kansai                                      | Japan                       | Kansai region (a.k.a. Kinki    | kHPpSs  |
|                                             |                             | region)                        |         |
| kanto                                       | Japan                       | Kantō region                   | kHPpSs  |
| karlsruhe-regbez                            | Baden-Württemberg           | Regierungsbezirk Karlsruhe     | kHPpSs  |
| kazakhstan                                  | Asia                        | Kazakhstan                     | kHPpSs  |
| kent                                        | England                     | Kent                           | kHPpSs  |
| kenya                                       | Africa                      | Kenya                          | kHPpSs  |
| kiribati                                    | Australia and Oceania       | Kiribati                       | kHPpSs  |
| koeln-regbez                                | Nordrhein-Westfalen         | Regierungsbezirk Köln          | kHPpSs  |
| kosovo                                      | Europe                      | Kosovo                         | kHPpSs  |
| kujawsko-pomorskie                          | Poland                      | Województwo                    | kHPpSs  |
|                                             |                             | kujawsko-pomorskie<br          |         |
|                                             |                             | />(Kuyavian-Pomeranian         |         |
|                                             |                             | Voivodeship)                   |         |
| kyrgyzstan                                  | Asia                        | Kyrgyzstan                     | kHPpSs  |
| kyushu                                      | Japan                       | Kyūshū                         | kHPpSs  |
| la-rioja                                    | Spain                       | La Rioja                       | kHPpSs  |
| lancashire                                  | England                     | Lancashire                     | kHPpSs  |
| languedoc-roussillon                        | France                      | Languedoc-Roussillon           | kHPpSs  |
| laos                                        | Asia                        | Laos                           | kHPpSs  |
| latvia                                      | Europe                      | Latvia                         | kHPpSs  |
| lebanon                                     | Asia                        | Lebanon                        | kHPpSs  |
| leicestershire                              | England                     | Leicestershire                 | kHPpSs  |
| lesotho                                     | Africa                      | Lesotho                        | kHPpSs  |
| liaoning                                    | China                       | Liaoning                       | kHPpSs  |
| liberia                                     | Africa                      | Liberia                        | kHPpSs  |
| libya                                       | Africa                      | Libya                          | kHPpSs  |
| liechtenstein                               | Europe                      | Liechtenstein                  | kHPpSs  |
| limburg                                     | Netherlands                 | Limburg                        | kHPpSs  |
| limousin                                    | France                      | Limousin                       | kHPpSs  |
| lincolnshire                                | England                     | Lincolnshire                   | kHPpSs  |
| lithuania                                   | Europe                      | Lithuania                      | kHPpSs  |
| lodzkie                                     | Poland                      | Województwo łódzkie<br />(Łódź | kHPpSs  |
|                                             |                             | Voivodeship)                   |         |
| lorraine                                    | France                      | Lorraine                       | kHPpSs  |
| lubelskie                                   | Poland                      | Województwo lubelskie<br       | kHPpSs  |
|                                             |                             | />(Lublin Voivodeship)         |         |
| lubuskie                                    | Poland                      | Województwo lubuskie<br        | kHPpSs  |
|                                             |                             | />(Lubusz Voivodeship)         |         |
| luxembourg                                  | Europe                      | Luxembourg                     | kHPpSs  |
| macau                                       | China                       | Macau                          | kHPpSs  |
| macedonia                                   | Europe                      | Macedonia                      | kHPpSs  |
| madagascar                                  | Africa                      | Madagascar                     | kHPpSs  |
| madrid                                      | Spain                       | Madrid                         | kHPpSs  |
| malawi                                      | Africa                      | Malawi                         | kHPpSs  |
| malaysia-singapore-brunei                   | Asia                        | Malaysia, Singapore, and       | kHPpSs  |
|                                             |                             | Brunei                         |         |
| maldives                                    | Asia                        | Maldives                       | kHPpSs  |
| mali                                        | Africa                      | Mali                           | kHPpSs  |
| malopolskie                                 | Poland                      | Województwo małopolskie<br     | kHPpSs  |
|                                             |                             | />(Lesser Poland Voivodeship)  |         |
| malta                                       | Europe                      | Malta                          | kHPpSs  |
| maluku                                      | Indonesia (with East Timor) | Maluku                         | kHPpSs  |
| manitoba                                    | Canada                      | Manitoba                       | kHPpSs  |
| marshall-islands                            | Australia and Oceania       | Marshall Islands               | kHPpSs  |
| martinique                                  | France                      | Martinique                     | kHPpSs  |
| mauritania                                  | Africa                      | Mauritania                     | kHPpSs  |
| mauritius                                   | Africa                      | Mauritius                      | kHPpSs  |
| mayotte                                     | France                      | Mayotte                        | kHPpSs  |
| mazowieckie                                 | Poland                      | Województwo mazowieckie<br     | kHPpSs  |
|                                             |                             | />(Mazovian Voivodeship)       |         |
| mecklenburg-vorpommern                      | Germany                     | Mecklenburg-Vorpommern         | kHPpSs  |
| melilla                                     | Spain                       | Melilla                        | kHPpSs  |
| merseyside                                  | England                     | Merseyside                     | kHPpSs  |
| mexico                                      | North America               | Mexico                         | kHPpSs  |
| micronesia                                  | Australia and Oceania       | Micronesia                     | kHPpSs  |
| midi-pyrenees                               | France                      | Midi-Pyrenees                  | kHPpSs  |
| mittelfranken                               | Bayern                      | Mittelfranken                  | kHPpSs  |
| moldova                                     | Europe                      | Moldova                        | kHPpSs  |
| monaco                                      | Europe                      | Monaco                         | kHPpSs  |
| mongolia                                    | Asia                        | Mongolia                       | kHPpSs  |
| montenegro                                  | Europe                      | Montenegro                     | kHPpSs  |
| morocco                                     | Africa                      | Morocco                        | kHPpSs  |
| mozambique                                  | Africa                      | Mozambique                     | kHPpSs  |
| muenster-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Münster       | kHPpSs  |
| murcia                                      | Spain                       | Murcia                         | kHPpSs  |
| myanmar                                     | Asia                        | Myanmar (a.k.a. Burma)         | kHPpSs  |
| namibia                                     | Africa                      | Namibia                        | kHPpSs  |
| nauru                                       | Australia and Oceania       | Nauru                          | kHPpSs  |
| navarra                                     | Spain                       | Navarra                        | kHPpSs  |
| nepal                                       | Asia                        | Nepal                          | kHPpSs  |
| netherlands                                 | Europe                      | Netherlands                    | kHPpSs  |
| new-brunswick                               | Canada                      | New Brunswick                  | kHPpSs  |
| new-caledonia                               | Australia and Oceania       | New Caledonia                  | kHPpSs  |
| new-zealand                                 | Australia and Oceania       | New Zealand                    | kHPpSs  |
| newfoundland-and-labrador                   | Canada                      | Newfoundland and Labrador      | kHPpSs  |
| nicaragua                                   | Central America             | Nicaragua                      | kHPpSs  |
| niederbayern                                | Bayern                      | Niederbayern                   | kHPpSs  |
| niedersachsen                               | Germany                     | Niedersachsen (mit Bremen)     | kHPpSs  |
| niger                                       | Africa                      | Niger                          | kHPpSs  |
| nigeria                                     | Africa                      | Nigeria                        | kHPpSs  |
| ningxia                                     | China                       | Ningxia                        | kHPpSs  |
| niue                                        | Australia and Oceania       | Niue                           | kHPpSs  |
| noord-brabant                               | Netherlands                 | Noord-Brabant                  | kHPpSs  |
| noord-holland                               | Netherlands                 | Noord-Holland                  | kHPpSs  |
| norcal                                      | us/california               | Northern California            | kHPpSs  |
| nord-est                                    | Italy                       | Nord-Est                       | kHPpSs  |
| nord-ovest                                  | Italy                       | Nord-Ovest                     | kHPpSs  |
| nord-pas-de-calais                          | France                      | Nord-Pas-de-Calais             | kHPpSs  |
| nordeste                                    | Brazil                      | Nordeste                       | kHPpSs  |
| nordrhein-westfalen                         | Germany                     | Nordrhein-Westfalen            | kHPpSs  |
| norfolk                                     | England                     | Norfolk                        | kHPpSs  |
| norte                                       | Brazil                      | Norte                          | kHPpSs  |
| north-america                               |                             | North America                  | kHPps   |
| north-caucasus-fed-district                 | Russian Federation          | North Caucasus Federal         | kHPpSs  |
|                                             |                             | District                       |         |
| north-eastern-zone                          | India                       | North-Eastern Zone             | kHPpSs  |
| north-korea                                 | Asia                        | North Korea                    | kHPpSs  |
| north-yorkshire                             | England                     | North Yorkshire                | kHPpSs  |
| northamptonshire                            | England                     | Northamptonshire               | kHPpSs  |
| northern-zone                               | India                       | Northern Zone                  | kHPpSs  |
| northumberland                              | England                     | Northumberland                 | kHPpSs  |
| northwest-territories                       | Canada                      | Northwest Territories          | kHPpSs  |
| northwestern-fed-district                   | Russian Federation          | Northwestern Federal District  | kHPpSs  |
| norway                                      | Europe                      | Norway                         | kHPpSs  |
| nottinghamshire                             | England                     | Nottinghamshire                | kHPpSs  |
| nova-scotia                                 | Canada                      | Nova Scotia                    | kHPpSs  |
| nunavut                                     | Canada                      | Nunavut                        | kHPpSs  |
| nusa-tenggara                               | Indonesia (with East Timor) | Nusa-Tenggara                  | kHPpSs  |
| oberbayern                                  | Bayern                      | Oberbayern                     | kHPpSs  |
| oberfranken                                 | Bayern                      | Oberfranken                    | kHPpSs  |
| oberpfalz                                   | Bayern                      | Oberpfalz                      | kHPpSs  |
| ontario                                     | Canada                      | Ontario                        | kHPpSs  |
| opolskie                                    | Poland                      | Województwo opolskie<br        | kHPpSs  |
|                                             |                             | />(Opole Voivodeship)          |         |
| overijssel                                  | Netherlands                 | Overijssel                     | kHPpSs  |
| oxfordshire                                 | England                     | Oxfordshire                    | kHPpSs  |
| pais-vasco                                  | Spain                       | País Vasco                     | kHPpSs  |
| pakistan                                    | Asia                        | Pakistan                       | kHPpSs  |
| palau                                       | Australia and Oceania       | Palau                          | kHPpSs  |
| panama                                      | Central America             | Panama                         | kHPpSs  |
| papua                                       | Indonesia (with East Timor) | Papua                          | kHPpSs  |
| papua-new-guinea                            | Australia and Oceania       | Papua New Guinea               | kHPpSs  |
| paraguay                                    | South America               | Paraguay                       | kHPpSs  |
| pays-de-la-loire                            | France                      | Pays de la Loire               | kHPpSs  |
| peru                                        | South America               | Peru                           | kHPpSs  |
| philippines                                 | Asia                        | Philippines                    | kHPpSs  |
| picardie                                    | France                      | Picardie                       | kHPpSs  |
| pitcairn-islands                            | Australia and Oceania       | Pitcairn Islands               | kHPpSs  |
| podkarpackie                                | Poland                      | Województwo podkarpackie<br    | kHPpSs  |
|                                             |                             | />(Subcarpathian Voivodeship)  |         |
| podlaskie                                   | Poland                      | Województwo podlaskie<br       | kHPpSs  |
|                                             |                             | />(Podlaskie Voivodeship)      |         |
| poitou-charentes                            | France                      | Poitou-Charentes               | kHPpSs  |
| poland                                      | Europe                      | Poland                         | kHPps   |
| polynesie-francaise                         | Australia and Oceania       | Polynésie française (French    | kHPpSs  |
|                                             |                             | Polynesia)                     |         |
| pomorskie                                   | Poland                      | Województwo pomorskie<br       | kHPpSs  |
|                                             |                             | />(Pomeranian Voivodeship)     |         |
| portugal                                    | Europe                      | Portugal                       | kHPpSs  |
| prince-edward-island                        | Canada                      | Prince Edward Island           | kHPpSs  |
| provence-alpes-cote-d-azur                  | France                      | Provence Alpes-Cote-d'Azur     | kHPpSs  |
| qinghai                                     | China                       | Qinghai                        | kHPpSs  |
| quebec                                      | Canada                      | Quebec                         | kHPpSs  |
| reunion                                     | France                      | Reunion                        | kHPpSs  |
| rheinland-pfalz                             | Germany                     | Rheinland-Pfalz                | kHPpSs  |
| rhone-alpes                                 | France                      | Rhone-Alpes                    | kHPpSs  |
| romania                                     | Europe                      | Romania                        | kHPpSs  |
| russia                                      |                             | Russian Federation             | kHPps   |
| rutland                                     | England                     | Rutland                        | kHPpSs  |
| rwanda                                      | Africa                      | Rwanda                         | kHPpSs  |
| saarland                                    | Germany                     | Saarland                       | kHPpSs  |
| sachsen                                     | Germany                     | Sachsen                        | kHPpSs  |
| sachsen-anhalt                              | Germany                     | Sachsen-Anhalt                 | kHPpSs  |
| saint-helena-ascension-and-tristan-da-cunha | Africa                      | Saint Helena, Ascension, and   | kHPpSs  |
|                                             |                             | Tristan da Cunha               |         |
| samoa                                       | Australia and Oceania       | Samoa                          | kHPpSs  |
| sao-tome-and-principe                       | Africa                      | Sao Tome and Principe          | kHPpSs  |
| saskatchewan                                | Canada                      | Saskatchewan                   | kHPpSs  |
| schleswig-holstein                          | Germany                     | Schleswig-Holstein             | kHPpSs  |
| schwaben                                    | Bayern                      | Schwaben                       | kHPpSs  |
| scotland                                    | United Kingdom              | Scotland                       | kHPpSs  |
| sea                                         | Asia                        | South-East Asia                | kHPps   |
| senegal-and-gambia                          | Africa                      | Senegal and Gambia             | kHPpSs  |
| serbia                                      | Europe                      | Serbia                         | kHPpSs  |
| seychelles                                  | Africa                      | Seychelles                     | kHPpSs  |
| shaanxi                                     | China                       | Shaanxi                        | kHPpSs  |
| shandong                                    | China                       | Shandong                       | kHPpSs  |
| shanghai                                    | China                       | Shanghai                       | kHPpSs  |
| shanxi                                      | China                       | Shanxi                         | kHPpSs  |
| shikoku                                     | Japan                       | Shikoku                        | kHPpSs  |
| shropshire                                  | England                     | Shropshire                     | kHPpSs  |
| siberian-fed-district                       | Russian Federation          | Siberian Federal District      | kHPpSs  |
| sichuan                                     | China                       | Sichuan                        | kHPpSs  |
| sierra-leone                                | Africa                      | Sierra Leone                   | kHPpSs  |
| slaskie                                     | Poland                      | Województwo śląskie<br         | kHPpSs  |
|                                             |                             | />(Silesian Voivodeship)       |         |
| slovakia                                    | Europe                      | Slovakia                       | kHPpSs  |
| slovenia                                    | Europe                      | Slovenia                       | kHPpSs  |
| socal                                       | us/california               | Southern California            | kHPpSs  |
| solomon-islands                             | Australia and Oceania       | Solomon Islands                | kHPpSs  |
| somalia                                     | Africa                      | Somalia                        | kHPpSs  |
| somerset                                    | England                     | Somerset                       | kHPpSs  |
| south-africa                                | Africa                      | South Africa                   | kHPpSs  |
| south-africa-and-lesotho                    | Africa                      | South Africa (includes         | kHPps   |
|                                             |                             | Lesotho)                       |         |
| south-america                               |                             | South America                  | kHPps   |
| south-fed-district                          | Russian Federation          | South Federal District         | kHPpSs  |
| south-korea                                 | Asia                        | South Korea                    | kHPpSs  |
| south-sudan                                 | Africa                      | South Sudan                    | kHPpSs  |
| south-yorkshire                             | England                     | South Yorkshire                | kHPpSs  |
| southern-zone                               | India                       | Southern Zone                  | kHPpSs  |
| spain                                       | Europe                      | Spain                          | kHPpSs  |
| sri-lanka                                   | Asia                        | Sri Lanka                      | kHPpSs  |
| staffordshire                               | England                     | Staffordshire                  | kHPpSs  |
| stuttgart-regbez                            | Baden-Württemberg           | Regierungsbezirk Stuttgart     | kHPpSs  |
| sud                                         | Italy                       | Sud                            | kHPpSs  |
| sudan                                       | Africa                      | Sudan                          | kHPpSs  |
| sudeste                                     | Brazil                      | Sudeste                        | kHPpSs  |
| suffolk                                     | England                     | Suffolk                        | kHPpSs  |
| sul                                         | Brazil                      | Sul                            | kHPpSs  |
| sulawesi                                    | Indonesia (with East Timor) | Sulawesi                       | kHPpSs  |
| sumatra                                     | Indonesia (with East Timor) | Sumatra                        | kHPpSs  |
| suriname                                    | South America               | Suriname                       | kHPpSs  |
| surrey                                      | England                     | Surrey                         | kHPpSs  |
| swaziland                                   | Africa                      | Swaziland                      | kHPpSs  |
| sweden                                      | Europe                      | Sweden                         | kHPpSs  |
| swietokrzyskie                              | Poland                      | Województwo świętokrzyskie<br  | kHPpSs  |
|                                             |                             | />(Świętokrzyskie Voivodeship) |         |
| switzerland                                 | Europe                      | Switzerland                    | kHPpSs  |
| syria                                       | Asia                        | Syria                          | kHPpSs  |
| taiwan                                      | Asia                        | Taiwan                         | kHPpSs  |
| tajikistan                                  | Asia                        | Tajikistan                     | kHPpSs  |
| tanzania                                    | Africa                      | Tanzania                       | kHPpSs  |
| thailand                                    | Asia                        | Thailand                       | kHPpSs  |
| thueringen                                  | Germany                     | Thüringen                      | kHPpSs  |
| tianjin                                     | China                       | Tianjin                        | kHPpSs  |
| tibet                                       | China                       | Tibet                          | kHPpSs  |
| togo                                        | Africa                      | Togo                           | kHPpSs  |
| tohoku                                      | Japan                       | Tōhoku region                  | kHPpSs  |
| tokelau                                     | Australia and Oceania       | Tokelau                        | kHPpSs  |
| tonga                                       | Australia and Oceania       | Tonga                          | kHPpSs  |
| tuebingen-regbez                            | Baden-Württemberg           | Regierungsbezirk Tübingen      | kHPpSs  |
| tunisia                                     | Africa                      | Tunisia                        | kHPpSs  |
| turkey                                      | Europe                      | Turkey                         | kHPpSs  |
| turkmenistan                                | Asia                        | Turkmenistan                   | kHPpSs  |
| tuvalu                                      | Australia and Oceania       | Tuvalu                         | kHPpSs  |
| tyne-and-wear                               | England                     | Tyne and Wear                  | kHPpSs  |
| uganda                                      | Africa                      | Uganda                         | kHPpSs  |
| ukraine                                     | Europe                      | Ukraine (with Crimea)          | kHPpSs  |
| united-kingdom                              | Europe                      | United Kingdom                 | kHPps   |
| unterfranken                                | Bayern                      | Unterfranken                   | kHPpSs  |
| ural-fed-district                           | Russian Federation          | Ural Federal District          | kHPpSs  |
| uruguay                                     | South America               | Uruguay                        | kHPpSs  |
| us                                          | North America               | United States of America       | kHPps   |
| us-midwest                                  | North America               | US Midwest                     | kHPps   |
| us-northeast                                | North America               | US Northeast                   | kHPps   |
| us-pacific                                  | North America               | US Pacific                     | kHPps   |
| us-south                                    | North America               | US South                       | kHPps   |
| us-west                                     | North America               | US West                        | kHPps   |
| us/alabama                                  | North America               | us/alabama                     | kHPpSs  |
| us/alaska                                   | North America               | us/alaska                      | kHPpSs  |
| us/arizona                                  | North America               | us/arizona                     | kHPpSs  |
| us/arkansas                                 | North America               | us/arkansas                    | kHPpSs  |
| us/california                               | North America               | us/california                  | kHPps   |
| us/colorado                                 | North America               | us/colorado                    | kHPpSs  |
| us/connecticut                              | North America               | us/connecticut                 | kHPpSs  |
| us/delaware                                 | North America               | us/delaware                    | kHPpSs  |
| us/district-of-columbia                     | North America               | us/district-of-columbia        | kHPpSs  |
| us/florida                                  | North America               | us/florida                     | kHPpSs  |
| us/georgia                                  | North America               | Georgia                        | kHPpSs  |
| us/hawaii                                   | North America               | us/hawaii                      | kHPpSs  |
| us/idaho                                    | North America               | us/idaho                       | kHPpSs  |
| us/illinois                                 | North America               | us/illinois                    | kHPpSs  |
| us/indiana                                  | North America               | us/indiana                     | kHPpSs  |
| us/iowa                                     | North America               | us/iowa                        | kHPpSs  |
| us/kansas                                   | North America               | us/kansas                      | kHPpSs  |
| us/kentucky                                 | North America               | us/kentucky                    | kHPpSs  |
| us/louisiana                                | North America               | us/louisiana                   | kHPpSs  |
| us/maine                                    | North America               | us/maine                       | kHPpSs  |
| us/maryland                                 | North America               | us/maryland                    | kHPpSs  |
| us/massachusetts                            | North America               | us/massachusetts               | kHPpSs  |
| us/michigan                                 | North America               | us/michigan                    | kHPpSs  |
| us/minnesota                                | North America               | us/minnesota                   | kHPpSs  |
| us/mississippi                              | North America               | us/mississippi                 | kHPpSs  |
| us/missouri                                 | North America               | us/missouri                    | kHPpSs  |
| us/montana                                  | North America               | us/montana                     | kHPpSs  |
| us/nebraska                                 | North America               | us/nebraska                    | kHPpSs  |
| us/nevada                                   | North America               | us/nevada                      | kHPpSs  |
| us/new-hampshire                            | North America               | us/new-hampshire               | kHPpSs  |
| us/new-jersey                               | North America               | us/new-jersey                  | kHPpSs  |
| us/new-mexico                               | North America               | us/new-mexico                  | kHPpSs  |
| us/new-york                                 | North America               | us/new-york                    | kHPpSs  |
| us/north-carolina                           | North America               | us/north-carolina              | kHPpSs  |
| us/north-dakota                             | North America               | us/north-dakota                | kHPpSs  |
| us/ohio                                     | North America               | us/ohio                        | kHPpSs  |
| us/oklahoma                                 | North America               | us/oklahoma                    | kHPpSs  |
| us/oregon                                   | North America               | us/oregon                      | kHPpSs  |
| us/pennsylvania                             | North America               | us/pennsylvania                | kHPpSs  |
| us/puerto-rico                              | North America               | us/puerto-rico                 | kHPpSs  |
| us/rhode-island                             | North America               | us/rhode-island                | kHPpSs  |
| us/south-carolina                           | North America               | us/south-carolina              | kHPpSs  |
| us/south-dakota                             | North America               | us/south-dakota                | kHPpSs  |
| us/tennessee                                | North America               | us/tennessee                   | kHPpSs  |
| us/texas                                    | North America               | us/texas                       | kHPpSs  |
| us/us-virgin-islands                        | North America               | us/us-virgin-islands           | kHPpSs  |
| us/utah                                     | North America               | us/utah                        | kHPpSs  |
| us/vermont                                  | North America               | us/vermont                     | kHPpSs  |
| us/virginia                                 | North America               | us/virginia                    | kHPpSs  |
| us/washington                               | North America               | us/washington                  | kHPpSs  |
| us/west-virginia                            | North America               | us/west-virginia               | kHPpSs  |
| us/wisconsin                                | North America               | us/wisconsin                   | kHPpSs  |
| us/wyoming                                  | North America               | us/wyoming                     | kHPpSs  |
| utrecht                                     | Netherlands                 | Utrecht                        | kHPpSs  |
| uzbekistan                                  | Asia                        | Uzbekistan                     | kHPpSs  |
| valencia                                    | Spain                       | Valencia                       | kHPpSs  |
| vanuatu                                     | Australia and Oceania       | Vanuatu                        | kHPpSs  |
| venezuela                                   | South America               | Venezuela                      | kHPpSs  |
| vietnam                                     | Asia                        | Vietnam                        | kHPpSs  |
| volga-fed-district                          | Russian Federation          | Volga Federal District         | kHPpSs  |
| wales                                       | United Kingdom              | Wales                          | kHPpSs  |
| wallis-et-futuna                            | Australia and Oceania       | Wallis et Futuna               | kHPpSs  |
| warminsko-mazurskie                         | Poland                      | Województwo                    | kHPpSs  |
|                                             |                             | warmińsko-mazurskie<br         |         |
|                                             |                             | />(Warmian-Masurian            |         |
|                                             |                             | Voivodeship)                   |         |
| warwickshire                                | England                     | Warwickshire                   | kHPpSs  |
| west-midlands                               | England                     | West Midlands                  | kHPpSs  |
| west-sussex                                 | England                     | West Sussex                    | kHPpSs  |
| west-yorkshire                              | England                     | West Yorkshire                 | kHPpSs  |
| western-zone                                | India                       | Western Zone                   | kHPpSs  |
| wielkopolskie                               | Poland                      | Województwo wielkopolskie<br   | kHPpSs  |
|                                             |                             | />(Greater Poland Voivodeship) |         |
| wiltshire                                   | England                     | Wiltshire                      | kHPpSs  |
| worcestershire                              | England                     | Worcestershire                 | kHPpSs  |
| xinjiang                                    | China                       | Xinjiang                       | kHPpSs  |
| yemen                                       | Asia                        | Yemen                          | kHPpSs  |
| yukon                                       | Canada                      | Yukon                          | kHPpSs  |
| yunnan                                      | China                       | Yunnan                         | kHPpSs  |
| zachodniopomorskie                          | Poland                      | Województwo                    | kHPpSs  |
|                                             |                             | zachodniopomorskie<br />(West  |         |
|                                             |                             | Pomeranian Voivodeship)        |         |
| zambia                                      | Africa                      | Zambia                         | kHPpSs  |
| zeeland                                     | Netherlands                 | Zeeland                        | kHPpSs  |
| zhejiang                                    | China                       | Zhejiang                       | kHPpSs  |
| zimbabwe                                    | Africa                      | Zimbabwe                       | kHPpSs  |
| zuid-holland                                | Netherlands                 | Zuid-Holland                   | kHPpSs  |
Total elements: 512

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
| egypt                                    | africa                           | egypt                                    | Pps     |
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
| europe                                   |                                  | europe                                   | Pps     |
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
| france                                   | europe                           | france                                   | Pps     |
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
| haiti                                    | central-america                  | haiti                                    | Pps     |
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
| lebanon                                  | asia                             | lebanon                                  | Pps     |
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
| madagascar                               | africa                           | madagascar                               | Pps     |
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
| morocco                                  | africa                           | morocco                                  | Pps     |
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
| negros_island_region                     | philippines                      | negros_island_region                     | Pps     |
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
| north-america                            |                                  | north-america                            | Pps     |
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

