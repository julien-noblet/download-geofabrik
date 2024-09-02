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
      --service="geofabrik"  Can switch to another service. You can use
                             "geofabrik", "openstreetmap.fr" "osmtoday" or
                             "bbbike". It automatically change config file if -c
                             is unused.
  -c, --config="./geofabrik.yml"  
                             Set Config file.
  -n, --[no-]nodownload      Do not download file (test only)
  -v, --[no-]verbose         Be verbose
  -q, --[no-]quiet           Be quiet
      --[no-]progress        Add a progress bar (implie quiet)
      --[no-]version         Show application version.

Commands:
help [<command>...]
    Show help.


list [<flags>]
    Show elements available

    --[no-]markdown  generate list in Markdown format

download [<flags>] <element>
    Download element

    -B, --[no-]osm.bz2  Download osm.bz2 if available
    -G, --[no-]osm.gz   Download osm.gz if available
    -S, --[no-]shp.zip  Download shp.zip if available
    -P, --[no-]osm.pbf  Download osm.pbf (default)
    -H, --[no-]osh.pbf  Download osh.pbf
    -s, --[no-]state    Download state.txt file
    -p, --[no-]poly     Download poly file
    -k, --[no-]kml      Download kml file
    -g, --[no-]geojson  Download geojson file
        --[no-]check    Control with checksum (default) Use --no-check to
                        discard control
    -d, --output_directory=OUTPUT_DIRECTORY  
                        Set output directory, you can use also OUTPUT_DIR env
                        variable

generate
    Generate a new config file



```

## List of elements
|                  SHORTNAME                  |            IS IN            |           LONG NAME            | FORMATS |
|---------------------------------------------|-----------------------------|--------------------------------|---------|
| afghanistan                                 | Asia                        | Afghanistan                    | sPBHpSk |
| africa                                      |                             | Africa                         | sPBHpk  |
| albania                                     | Europe                      | Albania                        | sPBHpSk |
| alberta                                     | Canada                      | Alberta                        | sPBHpSk |
| algeria                                     | Africa                      | Algeria                        | sPBHpSk |
| alps                                        | Europe                      | Alps                           | sPBHpk  |
| alsace                                      | France                      | Alsace                         | sPBHpSk |
| american-oceania                            | Australia and Oceania       | American Oceania               | sPBHpSk |
| andalucia                                   | Spain                       | Andalucía                      | sPBHpSk |
| andorra                                     | Europe                      | Andorra                        | sPBHpSk |
| angola                                      | Africa                      | Angola                         | sPBHpSk |
| anhui                                       | China                       | Anhui                          | sPBHpSk |
| antarctica                                  |                             | Antarctica                     | sPBHpSk |
| aquitaine                                   | France                      | Aquitaine                      | sPBHpSk |
| aragon                                      | Spain                       | Aragón                         | sPBHpSk |
| argentina                                   | South America               | Argentina                      | sPBHpSk |
| armenia                                     | Asia                        | Armenia                        | sPBHpSk |
| arnsberg-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Arnsberg      | sPBHpSk |
| asia                                        |                             | Asia                           | sPBHpk  |
| asturias                                    | Spain                       | Asturias                       | sPBHpSk |
| australia                                   | Australia and Oceania       | Australia                      | sPBHpSk |
| australia-oceania                           |                             | Australia and Oceania          | sPBHpk  |
| austria                                     | Europe                      | Austria                        | sPBHpSk |
| auvergne                                    | France                      | Auvergne                       | sPBHpSk |
| azerbaijan                                  | Asia                        | Azerbaijan                     | sPBHpSk |
| azores                                      | Europe                      | Azores                         | sPBHpSk |
| baden-wuerttemberg                          | Germany                     | Baden-Württemberg              | sPBHpk  |
| bahamas                                     | Central America             | Bahamas                        | sPBHpSk |
| bangladesh                                  | Asia                        | Bangladesh                     | sPBHpSk |
| basse-normandie                             | France                      | Basse-Normandie                | sPBHpSk |
| bayern                                      | Germany                     | Bayern                         | sPBHpk  |
| bedfordshire                                | England                     | Bedfordshire                   | sPBHpSk |
| beijing                                     | China                       | Beijing                        | sPBHpSk |
| belarus                                     | Europe                      | Belarus                        | sPBHpSk |
| belgium                                     | Europe                      | Belgium                        | sPBHpSk |
| belize                                      | Central America             | Belize                         | sPBHpSk |
| benin                                       | Africa                      | Benin                          | sPBHpSk |
| berkshire                                   | England                     | Berkshire                      | sPBHpSk |
| berlin                                      | Germany                     | Berlin                         | sPBHpSk |
| bhutan                                      | Asia                        | Bhutan                         | sPBHpSk |
| bolivia                                     | South America               | Bolivia                        | sPBHpSk |
| bosnia-herzegovina                          | Europe                      | Bosnia-Herzegovina             | sPBHpSk |
| botswana                                    | Africa                      | Botswana                       | sPBHpSk |
| bourgogne                                   | France                      | Bourgogne                      | sPBHpSk |
| brandenburg                                 | Germany                     | Brandenburg (mit Berlin)       | sPBHpSk |
| brazil                                      | South America               | Brazil                         | sPBHpk  |
| bremen                                      | Germany                     | Bremen                         | sPBHpSk |
| bretagne                                    | France                      | Bretagne                       | sPBHpSk |
| bristol                                     | England                     | Bristol                        | sPBHpSk |
| britain-and-ireland                         | Europe                      | Britain and Ireland            | sPBHpk  |
| british-columbia                            | Canada                      | British Columbia               | sPBHpSk |
| buckinghamshire                             | England                     | Buckinghamshire                | sPBHpSk |
| bulgaria                                    | Europe                      | Bulgaria                       | sPBHpSk |
| burkina-faso                                | Africa                      | Burkina Faso                   | sPBHpSk |
| burundi                                     | Africa                      | Burundi                        | sPBHpSk |
| cambodia                                    | Asia                        | Cambodia                       | sPBHpSk |
| cambridgeshire                              | England                     | Cambridgeshire                 | sPBHpSk |
| cameroon                                    | Africa                      | Cameroon                       | sPBHpSk |
| canada                                      | North America               | Canada                         | sPBHpk  |
| canary-islands                              | Africa                      | Canary Islands                 | sPBHpSk |
| cantabria                                   | Spain                       | Cantabria                      | sPBHpSk |
| cape-verde                                  | Africa                      | Cape Verde                     | sPBHpSk |
| castilla-la-mancha                          | Spain                       | Castilla-La Mancha             | sPBHpSk |
| castilla-y-leon                             | Spain                       | Castilla y León                | sPBHpSk |
| cataluna                                    | Spain                       | Cataluña                       | sPBHpSk |
| central-african-republic                    | Africa                      | Central African Republic       | sPBHpSk |
| central-america                             |                             | Central America                | sPBHpk  |
| central-fed-district                        | Russian Federation          | Central Federal District       | sPBHpSk |
| central-zone                                | India                       | Central Zone                   | sPBHpSk |
| centre                                      | France                      | Centre                         | sPBHpSk |
| centro                                      | Italy                       | Centro                         | sPBHpSk |
| centro-oeste                                | Brazil                      | Centro-Oeste                   | sPBHpSk |
| ceuta                                       | Spain                       | Ceuta                          | sPBHpSk |
| chad                                        | Africa                      | Chad                           | sPBHpSk |
| champagne-ardenne                           | France                      | Champagne Ardenne              | sPBHpSk |
| cheshire                                    | England                     | Cheshire                       | sPBHpSk |
| chile                                       | South America               | Chile                          | sPBHpSk |
| china                                       | Asia                        | China                          | sPBHpk  |
| chongqing                                   | China                       | Chongqing                      | sPBHpSk |
| chubu                                       | Japan                       | Chūbu region                   | sPBHpSk |
| chugoku                                     | Japan                       | Chūgoku region                 | sPBHpSk |
| colombia                                    | South America               | Colombia                       | sPBHpSk |
| comores                                     | Africa                      | Comores                        | sPBHpSk |
| congo-brazzaville                           | Africa                      | Congo (Republic/Brazzaville)   | sPBHpSk |
| congo-democratic-republic                   | Africa                      | Congo (Democratic              | sPBHpSk |
|                                             |                             | Republic/Kinshasa)             |         |
| cook-islands                                | Australia and Oceania       | Cook Islands                   | sPBHpSk |
| cornwall                                    | England                     | Cornwall                       | sPBHpSk |
| corse                                       | France                      | Corse                          | sPBHpSk |
| costa-rica                                  | Central America             | Costa Rica                     | sPBHpSk |
| crimean-fed-district                        | Russian Federation          | Crimean Federal District       | sPBHpSk |
| croatia                                     | Europe                      | Croatia                        | sPBHpSk |
| cuba                                        | Central America             | Cuba                           | sPBHpSk |
| cumbria                                     | England                     | Cumbria                        | sPBHpSk |
| cyprus                                      | Europe                      | Cyprus                         | sPBHpSk |
| czech-republic                              | Europe                      | Czech Republic                 | sPBHpSk |
| dach                                        | Europe                      | Germany, Austria, Switzerland  | sPBHpk  |
| denmark                                     | Europe                      | Denmark                        | sPBHpSk |
| derbyshire                                  | England                     | Derbyshire                     | sPBHpSk |
| detmold-regbez                              | Nordrhein-Westfalen         | Regierungsbezirk Detmold       | sPBHpSk |
| devon                                       | England                     | Devon                          | sPBHpSk |
| djibouti                                    | Africa                      | Djibouti                       | sPBHpSk |
| dolnoslaskie                                | Poland                      | Województwo dolnośląskie<br    | sPBHpSk |
|                                             |                             | />(Lower Silesian Voivodeship) |         |
| dorset                                      | England                     | Dorset                         | sPBHpSk |
| drenthe                                     | Netherlands                 | Drenthe                        | sPBHpSk |
| duesseldorf-regbez                          | Nordrhein-Westfalen         | Regierungsbezirk Düsseldorf    | sPBHpSk |
| durham                                      | England                     | Durham                         | sPBHpSk |
| east-sussex                                 | England                     | East Sussex                    | sPBHpSk |
| east-timor                                  | Asia                        | East Timor                     | sPBHpSk |
| east-yorkshire-with-hull                    | England                     | East Yorkshire with Hull       | sPBHpSk |
| eastern-zone                                | India                       | Eastern Zone                   | sPBHpSk |
| ecuador                                     | South America               | Ecuador                        | sPBHpk  |
| egypt                                       | Africa                      | Egypt                          | sPBHpSk |
| el-salvador                                 | Central America             | El Salvador                    | sPBHpSk |
| enfield                                     | Greater London              | Enfield                        | sPBHpSk |
| england                                     | United Kingdom              | England                        | sPBHpSk |
| equatorial-guinea                           | Africa                      | Equatorial Guinea              | sPBHpSk |
| eritrea                                     | Africa                      | Eritrea                        | sPBHpSk |
| essex                                       | England                     | Essex                          | sPBHpSk |
| estonia                                     | Europe                      | Estonia                        | sPBHpSk |
| ethiopia                                    | Africa                      | Ethiopia                       | sPBHpSk |
| europe                                      |                             | Europe                         | sPBHpk  |
| extremadura                                 | Spain                       | Extremadura                    | sPBHpSk |
| far-eastern-fed-district                    | Russian Federation          | Far Eastern Federal District   | sPBHpSk |
| faroe-islands                               | Europe                      | Faroe Islands                  | sPBHpSk |
| fiji                                        | Australia and Oceania       | Fiji                           | sPBHpSk |
| finland                                     | Europe                      | Finland                        | sPBHpSk |
| flevoland                                   | Netherlands                 | Flevoland                      | sPBHpSk |
| france                                      | Europe                      | France                         | sPBHpk  |
| franche-comte                               | France                      | Franche Comte                  | sPBHpSk |
| freiburg-regbez                             | Baden-Württemberg           | Regierungsbezirk Freiburg      | sPBHpSk |
| friesland                                   | Netherlands                 | Friesland                      | sPBHpSk |
| fujian                                      | China                       | Fujian                         | sPBHpSk |
| gabon                                       | Africa                      | Gabon                          | sPBHpSk |
| galicia                                     | Spain                       | Galicia                        | sPBHpSk |
| gansu                                       | China                       | Gansu                          | sPBHpSk |
| gcc-states                                  | Asia                        | GCC States                     | sPBHpSk |
| gelderland                                  | Netherlands                 | Gelderland                     | sPBHpSk |
| georgia                                     | Europe                      | Georgia                        | sPBHpSk |
| germany                                     | Europe                      | Germany                        | sPBHpk  |
| ghana                                       | Africa                      | Ghana                          | sPBHpSk |
| gloucestershire                             | England                     | Gloucestershire                | sPBHpSk |
| great-britain                               | Europe                      | Great Britain                  | sPBHpk  |
| greater-london                              | England                     | Greater London                 | sPBHpSk |
| greater-manchester                          | England                     | Greater Manchester             | sPBHpSk |
| greece                                      | Europe                      | Greece                         | sPBHpSk |
| greenland                                   | North America               | Greenland                      | sPBHpSk |
| groningen                                   | Netherlands                 | Groningen                      | sPBHpSk |
| guadeloupe                                  | France                      | Guadeloupe                     | sPBHpk  |
| guangdong                                   | China                       | Guangdong (with Hong Kong and  | sPBHpSk |
|                                             |                             | Macau)                         |         |
| guangxi                                     | China                       | Guangxi                        | sPBHpSk |
| guatemala                                   | Central America             | Guatemala                      | sPBHpSk |
| guernsey-jersey                             | Europe                      | Guernsey and Jersey            | sPBHpSk |
| guinea                                      | Africa                      | Guinea                         | sPBHpSk |
| guinea-bissau                               | Africa                      | Guinea-Bissau                  | sPBHpSk |
| guizhou                                     | China                       | Guizhou                        | sPBHpSk |
| guyana                                      | South America               | Guyana                         | sPBHpSk |
| guyane                                      | France                      | Guyane                         | sPBHpSk |
| hainan                                      | China                       | Hainan                         | sPBHpSk |
| haiti-and-domrep                            | Central America             | Haiti and Dominican Republic   | sPBHpSk |
| hamburg                                     | Germany                     | Hamburg                        | sPBHpSk |
| hampshire                                   | England                     | Hampshire                      | sPBHpSk |
| haute-normandie                             | France                      | Haute-Normandie                | sPBHpSk |
| hebei                                       | China                       | Hebei (with Beijing and        | sPBHpSk |
|                                             |                             | Tianjin)                       |         |
| heilongjiang                                | China                       | Heilongjiang                   | sPBHpSk |
| henan                                       | China                       | Henan                          | sPBHpSk |
| herefordshire                               | England                     | Herefordshire                  | sPBHpSk |
| hertfordshire                               | England                     | Hertfordshire                  | sPBHpSk |
| hessen                                      | Germany                     | Hessen                         | sPBHpSk |
| hokkaido                                    | Japan                       | Hokkaidō                       | sPBHpSk |
| honduras                                    | Central America             | Honduras                       | sPBHpSk |
| hong-kong                                   | China                       | Hong Kong                      | sPBHpSk |
| hubei                                       | China                       | Hubei                          | sPBHpSk |
| hunan                                       | China                       | Hunan                          | sPBHpSk |
| hungary                                     | Europe                      | Hungary                        | sPBHpSk |
| iceland                                     | Europe                      | Iceland                        | sPBHpSk |
| ile-de-clipperton                           | Australia and Oceania       | Île de Clipperton              | sPBHpSk |
| ile-de-france                               | France                      | Ile-de-France                  | sPBHpSk |
| india                                       | Asia                        | India                          | sPBHpk  |
| indonesia                                   | Asia                        | Indonesia (with East Timor)    | sPBHpk  |
| inner-mongolia                              | China                       | Inner Mongolia                 | sPBHpSk |
| iran                                        | Asia                        | Iran                           | sPBHpSk |
| iraq                                        | Asia                        | Iraq                           | sPBHpSk |
| ireland-and-northern-ireland                | Europe                      | Ireland and Northern Ireland   | sPBHpSk |
| islas-baleares                              | Spain                       | Islas Baleares                 | sPBHpSk |
| isle-of-man                                 | Europe                      | Isle of Man                    | sPBHpSk |
| isle-of-wight                               | England                     | Isle of Wight                  | sPBHpSk |
| isole                                       | Italy                       | Isole                          | sPBHpSk |
| israel-and-palestine                        | Asia                        | Israel and Palestine           | sPBHpSk |
| italy                                       | Europe                      | Italy                          | sPBHpk  |
| ivory-coast                                 | Africa                      | Ivory Coast                    | sPBHpSk |
| jamaica                                     | Central America             | Jamaica                        | sPBHpSk |
| japan                                       | Asia                        | Japan                          | sPBHpk  |
| java                                        | Indonesia (with East Timor) | Java                           | sPBHpSk |
| jiangsu                                     | China                       | Jiangsu                        | sPBHpSk |
| jiangxi                                     | China                       | Jiangxi                        | sPBHpSk |
| jilin                                       | China                       | Jilin                          | sPBHpSk |
| jordan                                      | Asia                        | Jordan                         | sPBHpSk |
| kalimantan                                  | Indonesia (with East Timor) | Kalimantan                     | sPBHpSk |
| kaliningrad                                 | Russian Federation          | Kaliningrad                    | sPBHpSk |
| kansai                                      | Japan                       | Kansai region (a.k.a. Kinki    | sPBHpSk |
|                                             |                             | region)                        |         |
| kanto                                       | Japan                       | Kantō region                   | sPBHpSk |
| karlsruhe-regbez                            | Baden-Württemberg           | Regierungsbezirk Karlsruhe     | sPBHpSk |
| kazakhstan                                  | Asia                        | Kazakhstan                     | sPBHpSk |
| kent                                        | England                     | Kent                           | sPBHpSk |
| kenya                                       | Africa                      | Kenya                          | sPBHpSk |
| kiribati                                    | Australia and Oceania       | Kiribati                       | sPBHpSk |
| koeln-regbez                                | Nordrhein-Westfalen         | Regierungsbezirk Köln          | sPBHpSk |
| kosovo                                      | Europe                      | Kosovo                         | sPBHpSk |
| kujawsko-pomorskie                          | Poland                      | Województwo                    | sPBHpSk |
|                                             |                             | kujawsko-pomorskie<br          |         |
|                                             |                             | />(Kuyavian-Pomeranian         |         |
|                                             |                             | Voivodeship)                   |         |
| kyrgyzstan                                  | Asia                        | Kyrgyzstan                     | sPBHpSk |
| kyushu                                      | Japan                       | Kyūshū                         | sPBHpSk |
| la-rioja                                    | Spain                       | La Rioja                       | sPBHpSk |
| lancashire                                  | England                     | Lancashire                     | sPBHpSk |
| languedoc-roussillon                        | France                      | Languedoc-Roussillon           | sPBHpSk |
| laos                                        | Asia                        | Laos                           | sPBHpSk |
| latvia                                      | Europe                      | Latvia                         | sPBHpSk |
| lebanon                                     | Asia                        | Lebanon                        | sPBHpSk |
| leicestershire                              | England                     | Leicestershire                 | sPBHpSk |
| lesotho                                     | Africa                      | Lesotho                        | sPBHpSk |
| liaoning                                    | China                       | Liaoning                       | sPBHpSk |
| liberia                                     | Africa                      | Liberia                        | sPBHpSk |
| libya                                       | Africa                      | Libya                          | sPBHpSk |
| liechtenstein                               | Europe                      | Liechtenstein                  | sPBHpSk |
| limburg                                     | Netherlands                 | Limburg                        | sPBHpSk |
| limousin                                    | France                      | Limousin                       | sPBHpSk |
| lincolnshire                                | England                     | Lincolnshire                   | sPBHpSk |
| lithuania                                   | Europe                      | Lithuania                      | sPBHpSk |
| lodzkie                                     | Poland                      | Województwo łódzkie<br />(Łódź | sPBHpSk |
|                                             |                             | Voivodeship)                   |         |
| lorraine                                    | France                      | Lorraine                       | sPBHpSk |
| lubelskie                                   | Poland                      | Województwo lubelskie<br       | sPBHpSk |
|                                             |                             | />(Lublin Voivodeship)         |         |
| lubuskie                                    | Poland                      | Województwo lubuskie<br        | sPBHpSk |
|                                             |                             | />(Lubusz Voivodeship)         |         |
| luxembourg                                  | Europe                      | Luxembourg                     | sPBHpSk |
| macau                                       | China                       | Macau                          | sPBHpSk |
| macedonia                                   | Europe                      | Macedonia                      | sPBHpSk |
| madagascar                                  | Africa                      | Madagascar                     | sPBHpSk |
| madrid                                      | Spain                       | Madrid                         | sPBHpSk |
| malawi                                      | Africa                      | Malawi                         | sPBHpSk |
| malaysia-singapore-brunei                   | Asia                        | Malaysia, Singapore, and       | sPBHpSk |
|                                             |                             | Brunei                         |         |
| maldives                                    | Asia                        | Maldives                       | sPBHpSk |
| mali                                        | Africa                      | Mali                           | sPBHpSk |
| malopolskie                                 | Poland                      | Województwo małopolskie<br     | sPBHpSk |
|                                             |                             | />(Lesser Poland Voivodeship)  |         |
| malta                                       | Europe                      | Malta                          | sPBHpSk |
| maluku                                      | Indonesia (with East Timor) | Maluku                         | sPBHpSk |
| manitoba                                    | Canada                      | Manitoba                       | sPBHpSk |
| marshall-islands                            | Australia and Oceania       | Marshall Islands               | sPBHpSk |
| martinique                                  | France                      | Martinique                     | sPBHpk  |
| mauritania                                  | Africa                      | Mauritania                     | sPBHpSk |
| mauritius                                   | Africa                      | Mauritius                      | sPBHpSk |
| mayotte                                     | France                      | Mayotte                        | sPBHpk  |
| mazowieckie                                 | Poland                      | Województwo mazowieckie<br     | sPBHpSk |
|                                             |                             | />(Mazovian Voivodeship)       |         |
| mecklenburg-vorpommern                      | Germany                     | Mecklenburg-Vorpommern         | sPBHpSk |
| melilla                                     | Spain                       | Melilla                        | sPBHpSk |
| merseyside                                  | England                     | Merseyside                     | sPBHpSk |
| mexico                                      | North America               | Mexico                         | sPBHpSk |
| micronesia                                  | Australia and Oceania       | Micronesia                     | sPBHpSk |
| midi-pyrenees                               | France                      | Midi-Pyrenees                  | sPBHpSk |
| mittelfranken                               | Bayern                      | Mittelfranken                  | sPBHpSk |
| moldova                                     | Europe                      | Moldova                        | sPBHpSk |
| monaco                                      | Europe                      | Monaco                         | sPBHpSk |
| mongolia                                    | Asia                        | Mongolia                       | sPBHpSk |
| montenegro                                  | Europe                      | Montenegro                     | sPBHpSk |
| morocco                                     | Africa                      | Morocco                        | sPBHpSk |
| mozambique                                  | Africa                      | Mozambique                     | sPBHpSk |
| muenster-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Münster       | sPBHpSk |
| murcia                                      | Spain                       | Murcia                         | sPBHpSk |
| myanmar                                     | Asia                        | Myanmar (a.k.a. Burma)         | sPBHpSk |
| namibia                                     | Africa                      | Namibia                        | sPBHpSk |
| nauru                                       | Australia and Oceania       | Nauru                          | sPBHpSk |
| navarra                                     | Spain                       | Navarra                        | sPBHpSk |
| nepal                                       | Asia                        | Nepal                          | sPBHpSk |
| netherlands                                 | Europe                      | Netherlands                    | sPBHpk  |
| new-brunswick                               | Canada                      | New Brunswick                  | sPBHpSk |
| new-caledonia                               | Australia and Oceania       | New Caledonia                  | sPBHpSk |
| new-zealand                                 | Australia and Oceania       | New Zealand                    | sPBHpSk |
| newfoundland-and-labrador                   | Canada                      | Newfoundland and Labrador      | sPBHpSk |
| nicaragua                                   | Central America             | Nicaragua                      | sPBHpk  |
| niederbayern                                | Bayern                      | Niederbayern                   | sPBHpSk |
| niedersachsen                               | Germany                     | Niedersachsen                  | sPBHpSk |
| niger                                       | Africa                      | Niger                          | sPBHpSk |
| nigeria                                     | Africa                      | Nigeria                        | sPBHpSk |
| ningxia                                     | China                       | Ningxia                        | sPBHpSk |
| niue                                        | Australia and Oceania       | Niue                           | sPBHpSk |
| noord-brabant                               | Netherlands                 | Noord-Brabant                  | sPBHpSk |
| noord-holland                               | Netherlands                 | Noord-Holland                  | sPBHpSk |
| norcal                                      | us/california               | Northern California            | sPBHpSk |
| nord-est                                    | Italy                       | Nord-Est                       | sPBHpSk |
| nord-ovest                                  | Italy                       | Nord-Ovest                     | sPBHpSk |
| nord-pas-de-calais                          | France                      | Nord-Pas-de-Calais             | sPBHpSk |
| nordeste                                    | Brazil                      | Nordeste                       | sPBHpSk |
| nordrhein-westfalen                         | Germany                     | Nordrhein-Westfalen            | sPBHpk  |
| norfolk                                     | England                     | Norfolk                        | sPBHpSk |
| norte                                       | Brazil                      | Norte                          | sPBHpSk |
| north-america                               |                             | North America                  | sPBHpk  |
| north-caucasus-fed-district                 | Russian Federation          | North Caucasus Federal         | sPBHpSk |
|                                             |                             | District                       |         |
| north-eastern-zone                          | India                       | North-Eastern Zone             | sPBHpSk |
| north-korea                                 | Asia                        | North Korea                    | sPBHpSk |
| north-yorkshire                             | England                     | North Yorkshire                | sPBHpSk |
| northamptonshire                            | England                     | Northamptonshire               | sPBHpSk |
| northern-zone                               | India                       | Northern Zone                  | sPBHpSk |
| northumberland                              | England                     | Northumberland                 | sPBHpSk |
| northwest-territories                       | Canada                      | Northwest Territories          | sPBHpSk |
| northwestern-fed-district                   | Russian Federation          | Northwestern Federal District  | sPBHpSk |
| norway                                      | Europe                      | Norway                         | sPBHpSk |
| nottinghamshire                             | England                     | Nottinghamshire                | sPBHpSk |
| nova-scotia                                 | Canada                      | Nova Scotia                    | sPBHpSk |
| nunavut                                     | Canada                      | Nunavut                        | sPBHpSk |
| nusa-tenggara                               | Indonesia (with East Timor) | Nusa-Tenggara                  | sPBHpSk |
| oberbayern                                  | Bayern                      | Oberbayern                     | sPBHpSk |
| oberfranken                                 | Bayern                      | Oberfranken                    | sPBHpSk |
| oberpfalz                                   | Bayern                      | Oberpfalz                      | sPBHpSk |
| ontario                                     | Canada                      | Ontario                        | sPBHpSk |
| opolskie                                    | Poland                      | Województwo opolskie<br        | sPBHpSk |
|                                             |                             | />(Opole Voivodeship)          |         |
| overijssel                                  | Netherlands                 | Overijssel                     | sPBHpSk |
| oxfordshire                                 | England                     | Oxfordshire                    | sPBHpSk |
| pais-vasco                                  | Spain                       | País Vasco                     | sPBHpSk |
| pakistan                                    | Asia                        | Pakistan                       | sPBHpSk |
| palau                                       | Australia and Oceania       | Palau                          | sPBHpSk |
| panama                                      | Central America             | Panama                         | sPBHpSk |
| papua                                       | Indonesia (with East Timor) | Papua                          | sPBHpSk |
| papua-new-guinea                            | Australia and Oceania       | Papua New Guinea               | sPBHpSk |
| paraguay                                    | South America               | Paraguay                       | sPBHpSk |
| pays-de-la-loire                            | France                      | Pays de la Loire               | sPBHpSk |
| peru                                        | South America               | Peru                           | sPBHpSk |
| philippines                                 | Asia                        | Philippines                    | sPBHpSk |
| picardie                                    | France                      | Picardie                       | sPBHpSk |
| pitcairn-islands                            | Australia and Oceania       | Pitcairn Islands               | sPBHpSk |
| podkarpackie                                | Poland                      | Województwo podkarpackie<br    | sPBHpSk |
|                                             |                             | />(Subcarpathian Voivodeship)  |         |
| podlaskie                                   | Poland                      | Województwo podlaskie<br       | sPBHpSk |
|                                             |                             | />(Podlaskie Voivodeship)      |         |
| poitou-charentes                            | France                      | Poitou-Charentes               | sPBHpSk |
| poland                                      | Europe                      | Poland                         | sPBHpk  |
| polynesie-francaise                         | Australia and Oceania       | Polynésie française (French    | sPBHpSk |
|                                             |                             | Polynesia)                     |         |
| pomorskie                                   | Poland                      | Województwo pomorskie<br       | sPBHpSk |
|                                             |                             | />(Pomeranian Voivodeship)     |         |
| portugal                                    | Europe                      | Portugal                       | sPBHpSk |
| prince-edward-island                        | Canada                      | Prince Edward Island           | sPBHpSk |
| provence-alpes-cote-d-azur                  | France                      | Provence Alpes-Cote-d'Azur     | sPBHpSk |
| qinghai                                     | China                       | Qinghai                        | sPBHpSk |
| quebec                                      | Canada                      | Quebec                         | sPBHpSk |
| reunion                                     | France                      | Reunion                        | sPBHpk  |
| rheinland-pfalz                             | Germany                     | Rheinland-Pfalz                | sPBHpSk |
| rhone-alpes                                 | France                      | Rhone-Alpes                    | sPBHpSk |
| romania                                     | Europe                      | Romania                        | sPBHpSk |
| russia                                      |                             | Russian Federation             | sPBHpk  |
| rutland                                     | England                     | Rutland                        | sPBHpSk |
| rwanda                                      | Africa                      | Rwanda                         | sPBHpSk |
| saarland                                    | Germany                     | Saarland                       | sPBHpSk |
| sachsen                                     | Germany                     | Sachsen                        | sPBHpSk |
| sachsen-anhalt                              | Germany                     | Sachsen-Anhalt                 | sPBHpSk |
| saint-helena-ascension-and-tristan-da-cunha | Africa                      | Saint Helena, Ascension, and   | sPBHpSk |
|                                             |                             | Tristan da Cunha               |         |
| samoa                                       | Australia and Oceania       | Samoa                          | sPBHpSk |
| sao-tome-and-principe                       | Africa                      | Sao Tome and Principe          | sPBHpSk |
| saskatchewan                                | Canada                      | Saskatchewan                   | sPBHpSk |
| schleswig-holstein                          | Germany                     | Schleswig-Holstein             | sPBHpSk |
| schwaben                                    | Bayern                      | Schwaben                       | sPBHpSk |
| scotland                                    | United Kingdom              | Scotland                       | sPBHpSk |
| senegal-and-gambia                          | Africa                      | Senegal and Gambia             | sPBHpSk |
| serbia                                      | Europe                      | Serbia                         | sPBHpSk |
| seychelles                                  | Africa                      | Seychelles                     | sPBHpSk |
| shaanxi                                     | China                       | Shaanxi                        | sPBHpSk |
| shandong                                    | China                       | Shandong                       | sPBHpSk |
| shanghai                                    | China                       | Shanghai                       | sPBHpSk |
| shanxi                                      | China                       | Shanxi                         | sPBHpSk |
| shikoku                                     | Japan                       | Shikoku                        | sPBHpSk |
| shropshire                                  | England                     | Shropshire                     | sPBHpSk |
| siberian-fed-district                       | Russian Federation          | Siberian Federal District      | sPBHpSk |
| sichuan                                     | China                       | Sichuan                        | sPBHpSk |
| sierra-leone                                | Africa                      | Sierra Leone                   | sPBHpSk |
| slaskie                                     | Poland                      | Województwo śląskie<br         | sPBHpSk |
|                                             |                             | />(Silesian Voivodeship)       |         |
| slovakia                                    | Europe                      | Slovakia                       | sPBHpSk |
| slovenia                                    | Europe                      | Slovenia                       | sPBHpSk |
| socal                                       | us/california               | Southern California            | sPBHpSk |
| solomon-islands                             | Australia and Oceania       | Solomon Islands                | sPBHpSk |
| somalia                                     | Africa                      | Somalia                        | sPBHpSk |
| somerset                                    | England                     | Somerset                       | sPBHpSk |
| south-africa                                | Africa                      | South Africa                   | sPBHpSk |
| south-africa-and-lesotho                    | Africa                      | South Africa (includes         | sPBHpk  |
|                                             |                             | Lesotho)                       |         |
| south-america                               |                             | South America                  | sPBHpk  |
| south-fed-district                          | Russian Federation          | South Federal District         | sPBHpSk |
| south-korea                                 | Asia                        | South Korea                    | sPBHpSk |
| south-sudan                                 | Africa                      | South Sudan                    | sPBHpSk |
| south-yorkshire                             | England                     | South Yorkshire                | sPBHpSk |
| southern-zone                               | India                       | Southern Zone                  | sPBHpSk |
| spain                                       | Europe                      | Spain                          | sPBHpk  |
| sri-lanka                                   | Asia                        | Sri Lanka                      | sPBHpSk |
| staffordshire                               | England                     | Staffordshire                  | sPBHpSk |
| stuttgart-regbez                            | Baden-Württemberg           | Regierungsbezirk Stuttgart     | sPBHpSk |
| sud                                         | Italy                       | Sud                            | sPBHpSk |
| sudan                                       | Africa                      | Sudan                          | sPBHpSk |
| sudeste                                     | Brazil                      | Sudeste                        | sPBHpSk |
| suffolk                                     | England                     | Suffolk                        | sPBHpSk |
| sul                                         | Brazil                      | Sul                            | sPBHpSk |
| sulawesi                                    | Indonesia (with East Timor) | Sulawesi                       | sPBHpSk |
| sumatra                                     | Indonesia (with East Timor) | Sumatra                        | sPBHpSk |
| suriname                                    | South America               | Suriname                       | sPBHpSk |
| surrey                                      | England                     | Surrey                         | sPBHpSk |
| swaziland                                   | Africa                      | Swaziland                      | sPBHpSk |
| sweden                                      | Europe                      | Sweden                         | sPBHpSk |
| swietokrzyskie                              | Poland                      | Województwo świętokrzyskie<br  | sPBHpSk |
|                                             |                             | />(Świętokrzyskie Voivodeship) |         |
| switzerland                                 | Europe                      | Switzerland                    | sPBHpSk |
| syria                                       | Asia                        | Syria                          | sPBHpSk |
| taiwan                                      | Asia                        | Taiwan                         | sPBHpSk |
| tajikistan                                  | Asia                        | Tajikistan                     | sPBHpSk |
| tanzania                                    | Africa                      | Tanzania                       | sPBHpSk |
| thailand                                    | Asia                        | Thailand                       | sPBHpSk |
| thueringen                                  | Germany                     | Thüringen                      | sPBHpSk |
| tianjin                                     | China                       | Tianjin                        | sPBHpSk |
| tibet                                       | China                       | Tibet                          | sPBHpSk |
| togo                                        | Africa                      | Togo                           | sPBHpSk |
| tohoku                                      | Japan                       | Tōhoku region                  | sPBHpSk |
| tokelau                                     | Australia and Oceania       | Tokelau                        | sPBHpSk |
| tonga                                       | Australia and Oceania       | Tonga                          | sPBHpSk |
| tuebingen-regbez                            | Baden-Württemberg           | Regierungsbezirk Tübingen      | sPBHpSk |
| tunisia                                     | Africa                      | Tunisia                        | sPBHpSk |
| turkey                                      | Europe                      | Turkey                         | sPBHpSk |
| turkmenistan                                | Asia                        | Turkmenistan                   | sPBHpSk |
| tuvalu                                      | Australia and Oceania       | Tuvalu                         | sPBHpSk |
| tyne-and-wear                               | England                     | Tyne and Wear                  | sPBHpSk |
| uganda                                      | Africa                      | Uganda                         | sPBHpSk |
| ukraine                                     | Europe                      | Ukraine (with Crimea)          | sPBHpSk |
| united-kingdom                              | Europe                      | United Kingdom                 | sPBHpk  |
| unterfranken                                | Bayern                      | Unterfranken                   | sPBHpSk |
| ural-fed-district                           | Russian Federation          | Ural Federal District          | sPBHpSk |
| uruguay                                     | South America               | Uruguay                        | sPBHpSk |
| us                                          | North America               | United States of America       | sPBHpk  |
| us-midwest                                  | North America               | US Midwest                     | sPBHpk  |
| us-northeast                                | North America               | US Northeast                   | sPBHpk  |
| us-pacific                                  | North America               | US Pacific                     | sPBHpk  |
| us-south                                    | North America               | US South                       | sPBHpk  |
| us-west                                     | North America               | US West                        | sPBHpk  |
| us/alabama                                  | North America               | us/alabama                     | sPBHpSk |
| us/alaska                                   | North America               | us/alaska                      | sPBHpSk |
| us/arizona                                  | North America               | us/arizona                     | sPBHpSk |
| us/arkansas                                 | North America               | us/arkansas                    | sPBHpSk |
| us/california                               | North America               | us/california                  | sPBHpk  |
| us/colorado                                 | North America               | us/colorado                    | sPBHpSk |
| us/connecticut                              | North America               | us/connecticut                 | sPBHpSk |
| us/delaware                                 | North America               | us/delaware                    | sPBHpSk |
| us/district-of-columbia                     | North America               | us/district-of-columbia        | sPBHpSk |
| us/florida                                  | North America               | us/florida                     | sPBHpSk |
| us/georgia                                  | North America               | Georgia                        | sPBHpSk |
| us/hawaii                                   | North America               | us/hawaii                      | sPBHpSk |
| us/idaho                                    | North America               | us/idaho                       | sPBHpSk |
| us/illinois                                 | North America               | us/illinois                    | sPBHpSk |
| us/indiana                                  | North America               | us/indiana                     | sPBHpSk |
| us/iowa                                     | North America               | us/iowa                        | sPBHpSk |
| us/kansas                                   | North America               | us/kansas                      | sPBHpSk |
| us/kentucky                                 | North America               | us/kentucky                    | sPBHpSk |
| us/louisiana                                | North America               | us/louisiana                   | sPBHpSk |
| us/maine                                    | North America               | us/maine                       | sPBHpSk |
| us/maryland                                 | North America               | us/maryland                    | sPBHpSk |
| us/massachusetts                            | North America               | us/massachusetts               | sPBHpSk |
| us/michigan                                 | North America               | us/michigan                    | sPBHpSk |
| us/minnesota                                | North America               | us/minnesota                   | sPBHpSk |
| us/mississippi                              | North America               | us/mississippi                 | sPBHpSk |
| us/missouri                                 | North America               | us/missouri                    | sPBHpSk |
| us/montana                                  | North America               | us/montana                     | sPBHpSk |
| us/nebraska                                 | North America               | us/nebraska                    | sPBHpSk |
| us/nevada                                   | North America               | us/nevada                      | sPBHpSk |
| us/new-hampshire                            | North America               | us/new-hampshire               | sPBHpSk |
| us/new-jersey                               | North America               | us/new-jersey                  | sPBHpSk |
| us/new-mexico                               | North America               | us/new-mexico                  | sPBHpSk |
| us/new-york                                 | North America               | us/new-york                    | sPBHpSk |
| us/north-carolina                           | North America               | us/north-carolina              | sPBHpSk |
| us/north-dakota                             | North America               | us/north-dakota                | sPBHpSk |
| us/ohio                                     | North America               | us/ohio                        | sPBHpSk |
| us/oklahoma                                 | North America               | us/oklahoma                    | sPBHpSk |
| us/oregon                                   | North America               | us/oregon                      | sPBHpSk |
| us/pennsylvania                             | North America               | us/pennsylvania                | sPBHpSk |
| us/puerto-rico                              | North America               | us/puerto-rico                 | sPBHpSk |
| us/rhode-island                             | North America               | us/rhode-island                | sPBHpSk |
| us/south-carolina                           | North America               | us/south-carolina              | sPBHpSk |
| us/south-dakota                             | North America               | us/south-dakota                | sPBHpSk |
| us/tennessee                                | North America               | us/tennessee                   | sPBHpSk |
| us/texas                                    | North America               | us/texas                       | sPBHpSk |
| us/us-virgin-islands                        | North America               | us/us-virgin-islands           | sPBHpSk |
| us/utah                                     | North America               | us/utah                        | sPBHpSk |
| us/vermont                                  | North America               | us/vermont                     | sPBHpSk |
| us/virginia                                 | North America               | us/virginia                    | sPBHpSk |
| us/washington                               | North America               | us/washington                  | sPBHpSk |
| us/west-virginia                            | North America               | us/west-virginia               | sPBHpSk |
| us/wisconsin                                | North America               | us/wisconsin                   | sPBHpSk |
| us/wyoming                                  | North America               | us/wyoming                     | sPBHpSk |
| utrecht                                     | Netherlands                 | Utrecht                        | sPBHpSk |
| uzbekistan                                  | Asia                        | Uzbekistan                     | sPBHpSk |
| valencia                                    | Spain                       | Valencia                       | sPBHpSk |
| vanuatu                                     | Australia and Oceania       | Vanuatu                        | sPBHpSk |
| venezuela                                   | South America               | Venezuela                      | sPBHpSk |
| vietnam                                     | Asia                        | Vietnam                        | sPBHpSk |
| volga-fed-district                          | Russian Federation          | Volga Federal District         | sPBHpSk |
| wales                                       | United Kingdom              | Wales                          | sPBHpSk |
| wallis-et-futuna                            | Australia and Oceania       | Wallis et Futuna               | sPBHpSk |
| warminsko-mazurskie                         | Poland                      | Województwo                    | sPBHpSk |
|                                             |                             | warmińsko-mazurskie<br         |         |
|                                             |                             | />(Warmian-Masurian            |         |
|                                             |                             | Voivodeship)                   |         |
| warwickshire                                | England                     | Warwickshire                   | sPBHpSk |
| west-midlands                               | England                     | West Midlands                  | sPBHpSk |
| west-sussex                                 | England                     | West Sussex                    | sPBHpSk |
| west-yorkshire                              | England                     | West Yorkshire                 | sPBHpSk |
| western-zone                                | India                       | Western Zone                   | sPBHpSk |
| wielkopolskie                               | Poland                      | Województwo wielkopolskie<br   | sPBHpSk |
|                                             |                             | />(Greater Poland Voivodeship) |         |
| wiltshire                                   | England                     | Wiltshire                      | sPBHpSk |
| worcestershire                              | England                     | Worcestershire                 | sPBHpSk |
| xinjiang                                    | China                       | Xinjiang                       | sPBHpSk |
| yemen                                       | Asia                        | Yemen                          | sPBHpSk |
| yukon                                       | Canada                      | Yukon                          | sPBHpSk |
| yunnan                                      | China                       | Yunnan                         | sPBHpSk |
| zachodniopomorskie                          | Poland                      | Województwo                    | sPBHpSk |
|                                             |                             | zachodniopomorskie<br />(West  |         |
|                                             |                             | Pomeranian Voivodeship)        |         |
| zambia                                      | Africa                      | Zambia                         | sPBHpSk |
| zeeland                                     | Netherlands                 | Zeeland                        | sPBHpSk |
| zhejiang                                    | China                       | Zhejiang                       | sPBHpSk |
| zimbabwe                                    | Africa                      | Zimbabwe                       | sPBHpSk |
| zuid-holland                                | Netherlands                 | Zuid-Holland                   | sPBHpSk |
Total elements: 509

## List of elements from openstreetmap.fr
|                SHORTNAME                 |              IS IN               |                LONG NAME                 | FORMATS |
|------------------------------------------|----------------------------------|------------------------------------------|---------|
| aargau                                   | switzerland                      | aargau                                   | sPp     |
| abitibi_temiscamingue                    | quebec                           | abitibi_temiscamingue                    | sPp     |
| abruzzo                                  | italy                            | abruzzo                                  | sPp     |
| aceh                                     | indonesia                        | aceh                                     | sPp     |
| acre                                     | north                            | acre                                     | sPp     |
| adygea_republic                          | southern_federal_district        | adygea_republic                          | sPp     |
| aegean                                   | turkey                           | aegean                                   | sPp     |
| afghanistan                              | asia                             | afghanistan                              | sPp     |
| africa                                   |                                  | africa                                   | sPp     |
| africa_france_taaf                       | africa                           | africa_france_taaf                       | sPp     |
| aguascalientes                           | mexico                           | aguascalientes                           | sPp     |
| ain                                      | rhone_alpes                      | ain                                      | sPp     |
| aisne                                    | picardie                         | aisne                                    | sPp     |
| akershus                                 | norway                           | akershus                                 | sPp     |
| alagoas                                  | northeast                        | alagoas                                  | sPp     |
| alameda                                  | california                       | alameda                                  | sPp     |
| aland                                    | finland                          | aland                                    | sPp     |
| alava                                    | euskadi                          | alava                                    | sPp     |
| albacete                                 | castilla_la_mancha               | albacete                                 | sPp     |
| alberta                                  | canada                           | alberta                                  | sPp     |
| algeria                                  | africa                           | algeria                                  | sPp     |
| alicante                                 | comunitat_valenciana             | alicante                                 | sPp     |
| allier                                   | auvergne                         | allier                                   | sPp     |
| almeria                                  | andalucia                        | almeria                                  | sPp     |
| alpes_de_haute_provence                  | provence_alpes_cote_d_azur       | alpes_de_haute_provence                  | sPp     |
| alpes_maritimes                          | provence_alpes_cote_d_azur       | alpes_maritimes                          | sPp     |
| alpine                                   | california                       | alpine                                   | sPp     |
| alsace                                   | france                           | alsace                                   | sPp     |
| altai_krai                               | siberian_federal_district        | altai_krai                               | sPp     |
| altai_republic                           | siberian_federal_district        | altai_republic                           | sPp     |
| amador                                   | california                       | amador                                   | sPp     |
| amapa                                    | north                            | amapa                                    | sPp     |
| amazonas                                 | north                            | amazonas                                 | sPp     |
| american_samoa                           | oceania                          | american_samoa                           | sPp     |
| amur_oblast                              | far_eastern_federal_district     | amur_oblast                              | sPp     |
| andalucia                                | spain                            | andalucia                                | sPp     |
| andaman_and_nicobar_islands              | india                            | andaman_and_nicobar_islands              | sPp     |
| andhra_pradesh                           | india                            | andhra_pradesh                           | sPp     |
| angola                                   | africa                           | angola                                   | sPp     |
| anguilla                                 | central-america                  | anguilla                                 | sPp     |
| anhui                                    | china                            | anhui                                    | sPp     |
| antigua_and_barbuda                      | central-america                  | antigua_and_barbuda                      | sPp     |
| antofagasta                              | chile                            | antofagasta                              | sPp     |
| antwerp                                  | flanders                         | antwerp                                  | sPp     |
| appenzell_ausserrhoden                   | switzerland                      | appenzell_ausserrhoden                   | sPp     |
| appenzell_innerrhoden                    | switzerland                      | appenzell_innerrhoden                    | sPp     |
| appomattox                               | virginia                         | appomattox                               | sPp     |
| aquitaine                                | france                           | aquitaine                                | sPp     |
| aragon                                   | spain                            | aragon                                   | sPp     |
| araucania                                | chile                            | araucania                                | sPp     |
| ardeche                                  | rhone_alpes                      | ardeche                                  | sPp     |
| ardennes                                 | champagne_ardenne                | ardennes                                 | sPp     |
| argentina                                | south-america                    | argentina                                | sPp     |
| arica_y_parinacota                       | chile                            | arica_y_parinacota                       | sPp     |
| ariege                                   | midi_pyrenees                    | ariege                                   | sPp     |
| arkhangelsk_oblast                       | northwestern_federal_district    | arkhangelsk_oblast                       | sPp     |
| armenia                                  | asia                             | armenia                                  | sPp     |
| arnsberg                                 | nordrhein_westfalen              | arnsberg                                 | sPp     |
| aruba                                    | central-america                  | aruba                                    | sPp     |
| arunachal_pradesh                        | india                            | arunachal_pradesh                        | sPp     |
| ashmore_and_cartier_islands              | australia                        | ashmore_and_cartier_islands              | sPp     |
| asia                                     |                                  | asia                                     | sPp     |
| assam                                    | india                            | assam                                    | sPp     |
| astrakhan_oblast                         | southern_federal_district        | astrakhan_oblast                         | sPp     |
| asturias                                 | spain                            | asturias                                 | sPp     |
| atacama                                  | chile                            | atacama                                  | sPp     |
| aube                                     | champagne_ardenne                | aube                                     | sPp     |
| aude                                     | languedoc_roussillon             | aude                                     | sPp     |
| aust-agder                               | norway                           | aust-agder                               | sPp     |
| australia                                | oceania                          | australia                                | sPp     |
| australian_capital_territory             | australia                        | australian_capital_territory             | sPp     |
| austria                                  | europe                           | austria                                  | sPp     |
| auvergne                                 | france                           | auvergne                                 | sPp     |
| aveiro                                   | portugal                         | aveiro                                   | sPp     |
| aveyron                                  | midi_pyrenees                    | aveyron                                  | sPp     |
| avila                                    | castilla_y_leon                  | avila                                    | sPp     |
| aysen                                    | chile                            | aysen                                    | sPp     |
| azores                                   | portugal                         | azores                                   | sPp     |
| badajoz                                  | extremadura                      | badajoz                                  | sPp     |
| bahamas                                  | central-america                  | bahamas                                  | sPp     |
| bahia                                    | northeast                        | bahia                                    | sPp     |
| bahrain                                  | asia                             | bahrain                                  | sPp     |
| baja_california                          | mexico                           | baja_california                          | sPp     |
| baja_california_sur                      | mexico                           | baja_california_sur                      | sPp     |
| bali                                     | indonesia                        | bali                                     | sPp     |
| bangka_belitung_islands                  | indonesia                        | bangka_belitung_islands                  | sPp     |
| bangsamoro                               | philippines                      | bangsamoro                               | sPp     |
| banskobystricky                          | slovakia                         | banskobystricky                          | sPp     |
| banten                                   | indonesia                        | banten                                   | sPp     |
| barbados                                 | central-america                  | barbados                                 | sPp     |
| barcelona                                | catalunya                        | barcelona                                | sPp     |
| bas_rhin                                 | alsace                           | bas_rhin                                 | sPp     |
| bas_saint_laurent                        | quebec                           | bas_saint_laurent                        | sPp     |
| basel_landschaft                         | switzerland                      | basel_landschaft                         | sPp     |
| basel_stadt                              | switzerland                      | basel_stadt                              | sPp     |
| bashkortostan_republic                   | volga_federal_district           | bashkortostan_republic                   | sPp     |
| basilicata                               | italy                            | basilicata                               | sPp     |
| basse_normandie                          | france                           | basse_normandie                          | sPp     |
| beijing                                  | china                            | beijing                                  | sPp     |
| beja                                     | portugal                         | beja                                     | sPp     |
| belgium                                  | europe                           | belgium                                  | sPp     |
| belgorod_oblast                          | central_federal_district         | belgorod_oblast                          | sPp     |
| bengkulu                                 | indonesia                        | bengkulu                                 | sPp     |
| benin                                    | africa                           | benin                                    | sPp     |
| bermuda                                  | north-america                    | bermuda                                  | sPp     |
| bern                                     | switzerland                      | bern                                     | sPp     |
| bhutan                                   | asia                             | bhutan                                   | sPp     |
| bicol_region                             | philippines                      | bicol_region                             | sPp     |
| bihar                                    | india                            | bihar                                    | sPp     |
| biobio                                   | chile                            | biobio                                   | sPp     |
| bir_tawil                                | africa                           | bir_tawil                                | sPp     |
| black_sea                                | turkey                           | black_sea                                | sPp     |
| blekinge                                 | sweden                           | blekinge                                 | sPp     |
| bouches_du_rhone                         | provence_alpes_cote_d_azur       | bouches_du_rhone                         | sPp     |
| bourgogne                                | france                           | bourgogne                                | sPp     |
| bouvet_island                            | africa                           | bouvet_island                            | sPp     |
| braga                                    | portugal                         | braga                                    | sPp     |
| braganca                                 | portugal                         | braganca                                 | sPp     |
| bratislavsky                             | slovakia                         | bratislavsky                             | sPp     |
| brazil                                   | south-america                    | brazil                                   | sPp     |
| brazil_central-west                      | brazil                           | brazil_central-west                      | sPp     |
| brazil_north                             | brazil                           | brazil_north                             | sPp     |
| brazil_northeast                         | brazil                           | brazil_northeast                         | sPp     |
| brazil_south                             | brazil                           | brazil_south                             | sPp     |
| brazil_southeast                         | brazil                           | brazil_southeast                         | sPp     |
| bretagne                                 | france                           | bretagne                                 | sPp     |
| british_columbia                         | canada                           | british_columbia                         | sPp     |
| british_indian_ocean_territory           | asia                             | british_indian_ocean_territory           | sPp     |
| british_virgin_islands                   | central-america                  | british_virgin_islands                   | sPp     |
| brunei                                   | asia                             | brunei                                   | sPp     |
| brussels_capital_region                  | belgium                          | brussels_capital_region                  | sPp     |
| bryansk_oblast                           | central_federal_district         | bryansk_oblast                           | sPp     |
| buenos_aires                             | argentina                        | buenos_aires                             | sPp     |
| buenos_aires_city                        | argentina                        | buenos_aires_city                        | sPp     |
| burgenland                               | austria                          | burgenland                               | sPp     |
| burgos                                   | castilla_y_leon                  | burgos                                   | sPp     |
| burkina_faso                             | africa                           | burkina_faso                             | sPp     |
| burundi                                  | africa                           | burundi                                  | sPp     |
| buryatia_republic                        | siberian_federal_district        | buryatia_republic                        | sPp     |
| buskerud                                 | norway                           | buskerud                                 | sPp     |
| butte                                    | california                       | butte                                    | sPp     |
| cabo_delgado                             | mozambique                       | cabo_delgado                             | sPp     |
| caceres                                  | extremadura                      | caceres                                  | sPp     |
| cadiz                                    | andalucia                        | cadiz                                    | sPp     |
| cagayan_valley                           | philippines                      | cagayan_valley                           | sPp     |
| calabarzon                               | philippines                      | calabarzon                               | sPp     |
| calabria                                 | italy                            | calabria                                 | sPp     |
| calaveras                                | california                       | calaveras                                | sPp     |
| california                               | us-west                          | california                               | sPp     |
| california_lake                          | california                       | california_lake                          | sPp     |
| calvados                                 | basse_normandie                  | calvados                                 | sPp     |
| cambodia                                 | asia                             | cambodia                                 | sPp     |
| cameroon                                 | africa                           | cameroon                                 | sPp     |
| campania                                 | italy                            | campania                                 | sPp     |
| campeche                                 | mexico                           | campeche                                 | sPp     |
| canada                                   | north-america                    | canada                                   | sPp     |
| canarias                                 | spain                            | canarias                                 | sPp     |
| cantabria                                | spain                            | cantabria                                | sPp     |
| cantal                                   | auvergne                         | cantal                                   | sPp     |
| cape_verde                               | africa                           | cape_verde                               | sPp     |
| capital_district                         | new-york                         | capital_district                         | sPp     |
| capitale_nationale                       | quebec                           | capitale_nationale                       | sPp     |
| caraga                                   | philippines                      | caraga                                   | sPp     |
| caribbean                                | central-america                  | caribbean                                | sPp     |
| castellon                                | comunitat_valenciana             | castellon                                | sPp     |
| castelo_branco                           | portugal                         | castelo_branco                           | sPp     |
| castilla_la_mancha                       | spain                            | castilla_la_mancha                       | sPp     |
| castilla_y_leon                          | spain                            | castilla_y_leon                          | sPp     |
| catalunya                                | spain                            | catalunya                                | sPp     |
| catamarca                                | argentina                        | catamarca                                | sPp     |
| cayman_islands                           | central-america                  | cayman_islands                           | sPp     |
| ceara                                    | northeast                        | ceara                                    | sPp     |
| central-america                          |                                  | central-america                          | sPp     |
| central-west                             | brazil                           | central-west                             |         |
| central_african_republic                 | africa                           | central_african_republic                 | sPp     |
| central_anatolia                         | turkey                           | central_anatolia                         | sPp     |
| central_federal_district                 | russia                           | central_federal_district                 | sPp     |
| central_finland                          | finland                          | central_finland                          | sPp     |
| central_java                             | indonesia                        | central_java                             | sPp     |
| central_kalimantan                       | indonesia                        | central_kalimantan                       | sPp     |
| central_luzon                            | philippines                      | central_luzon                            | sPp     |
| central_new_york                         | new-york                         | central_new_york                         | sPp     |
| central_ontario                          | ontario                          | central_ontario                          | sPp     |
| central_ostrobothnia                     | finland                          | central_ostrobothnia                     | sPp     |
| central_papua                            | indonesia                        | central_papua                            | sPp     |
| central_sulawesi                         | indonesia                        | central_sulawesi                         | sPp     |
| central_visayas                          | philippines                      | central_visayas                          | sPp     |
| centre                                   | france                           | centre                                   | sPp     |
| centre_du_quebec                         | quebec                           | centre_du_quebec                         | sPp     |
| ceuta                                    | spain                            | ceuta                                    | sPp     |
| chaco                                    | argentina                        | chaco                                    | sPp     |
| chad                                     | africa                           | chad                                     | sPp     |
| champagne_ardenne                        | france                           | champagne_ardenne                        | sPp     |
| chandigarh                               | india                            | chandigarh                               | sPp     |
| charente                                 | poitou_charentes                 | charente                                 | sPp     |
| charente_maritime                        | poitou_charentes                 | charente_maritime                        | sPp     |
| chaudiere_appalaches                     | quebec                           | chaudiere_appalaches                     | sPp     |
| chechen_republic                         | north_caucasian_federal_district | chechen_republic                         | sPp     |
| chelyabinsk_oblast                       | ural_federal_district            | chelyabinsk_oblast                       | sPp     |
| cher                                     | centre                           | cher                                     | sPp     |
| cherkasy_oblast                          | ukraine                          | cherkasy_oblast                          | sPp     |
| chernihiv_oblast                         | ukraine                          | chernihiv_oblast                         | sPp     |
| chernivtsi_oblast                        | ukraine                          | chernivtsi_oblast                        | sPp     |
| chesapeake                               | virginia                         | chesapeake                               | sPp     |
| chhattisgarh                             | india                            | chhattisgarh                             | sPp     |
| chiapas                                  | mexico                           | chiapas                                  | sPp     |
| chihuahua                                | mexico                           | chihuahua                                | sPp     |
| chile                                    | south-america                    | chile                                    | sPp     |
| china                                    | asia                             | china                                    | sPp     |
| chongqing                                | china                            | chongqing                                | sPp     |
| christmas_island                         | australia                        | christmas_island                         | sPp     |
| chubu                                    | japan                            | chubu                                    | sPp     |
| chubut                                   | argentina                        | chubut                                   | sPp     |
| chugoku                                  | japan                            | chugoku                                  | sPp     |
| chukotka_autonomous_okrug                | far_eastern_federal_district     | chukotka_autonomous_okrug                | sPp     |
| chuvash_republic                         | volga_federal_district           | chuvash_republic                         | sPp     |
| ciudad_real                              | castilla_la_mancha               | ciudad_real                              | sPp     |
| clipperton                               | oceania                          | clipperton                               | sPp     |
| coahuila                                 | mexico                           | coahuila                                 | sPp     |
| coastal                                  | tanzania                         | coastal                                  | sPp     |
| cocos_islands                            | australia                        | cocos_islands                            | sPp     |
| coimbra                                  | portugal                         | coimbra                                  | sPp     |
| colima                                   | mexico                           | colima                                   | sPp     |
| colorado                                 | us-west                          | colorado                                 | sPp     |
| colorado_northeast                       | colorado                         | colorado_northeast                       | sPp     |
| colorado_northwest                       | colorado                         | colorado_northwest                       | sPp     |
| colorado_southeast                       | colorado                         | colorado_southeast                       | sPp     |
| colorado_southwest                       | colorado                         | colorado_southwest                       | sPp     |
| colusa                                   | california                       | colusa                                   | sPp     |
| comoros                                  | africa                           | comoros                                  | sPp     |
| comunidad_de_madrid                      | spain                            | comunidad_de_madrid                      | sPp     |
| comunidad_foral_de_navarra               | spain                            | comunidad_foral_de_navarra               | sPp     |
| comunitat_valenciana                     | spain                            | comunitat_valenciana                     | sPp     |
| congo_brazzaville                        | africa                           | congo_brazzaville                        | sPp     |
| congo_kinshasa                           | africa                           | congo_kinshasa                           | sPp     |
| contra_costa                             | california                       | contra_costa                             | sPp     |
| cook                                     | illinois                         | cook                                     | sPp     |
| cook_islands                             | oceania                          | cook_islands                             | sPp     |
| coquimbo                                 | chile                            | coquimbo                                 | sPp     |
| coral_sea_islands                        | australia                        | coral_sea_islands                        | sPp     |
| cordillera_administrative_region         | philippines                      | cordillera_administrative_region         | sPp     |
| cordoba                                  | argentina                        | cordoba                                  | sPp     |
| correze                                  | limousin                         | correze                                  | sPp     |
| corrientes                               | argentina                        | corrientes                               | sPp     |
| corse                                    | france                           | corse                                    | sPp     |
| corse_du_sud                             | corse                            | corse_du_sud                             | sPp     |
| costa_rica                               | central-america                  | costa_rica                               | sPp     |
| cote_d_or                                | bourgogne                        | cote_d_or                                | sPp     |
| cote_nord                                | quebec                           | cote_nord                                | sPp     |
| cotes_d_armor                            | bretagne                         | cotes_d_armor                            | sPp     |
| creuse                                   | limousin                         | creuse                                   | sPp     |
| crimea                                   | ukraine                          | crimea                                   | sPp     |
| crimea_republic                          | southern_federal_district        | crimea_republic                          | sPp     |
| cuenca                                   | castilla_la_mancha               | cuenca                                   | sPp     |
| culpeper                                 | virginia                         | culpeper                                 | sPp     |
| curacao                                  | central-america                  | curacao                                  | sPp     |
| czech_republic                           | europe                           | czech_republic                           | sPp     |
| dadra_and_nagar_haveli                   | india                            | dadra_and_nagar_haveli                   | sPp     |
| dadra_and_nagar_haveli_and_daman_and_diu | india                            | dadra_and_nagar_haveli_and_daman_and_diu | sPp     |
| dagestan_republic                        | north_caucasian_federal_district | dagestan_republic                        | sPp     |
| dalarna                                  | sweden                           | dalarna                                  | sPp     |
| daman_and_diu                            | india                            | daman_and_diu                            | sPp     |
| davao_region                             | philippines                      | davao_region                             | sPp     |
| del_norte                                | california                       | del_norte                                | sPp     |
| denver                                   | colorado                         | denver                                   | sPp     |
| detmold                                  | nordrhein_westfalen              | detmold                                  | sPp     |
| detroit_metro                            | michigan                         | detroit_metro                            | sPp     |
| deux_sevres                              | poitou_charentes                 | deux_sevres                              | sPp     |
| distrito-federal                         | central-west                     | distrito-federal                         | sPp     |
| djibouti                                 | africa                           | djibouti                                 | sPp     |
| dnipropetrovsk_oblast                    | ukraine                          | dnipropetrovsk_oblast                    | sPp     |
| dolnoslaskie                             | poland                           | dolnoslaskie                             | sPp     |
| dominica                                 | central-america                  | dominica                                 | sPp     |
| dominican_republic                       | central-america                  | dominican_republic                       | sPp     |
| donetsk_oblast                           | ukraine                          | donetsk_oblast                           | sPp     |
| dordogne                                 | aquitaine                        | dordogne                                 | sPp     |
| doubs                                    | franche_comte                    | doubs                                    | sPp     |
| drenthe                                  | netherlands                      | drenthe                                  | sPp     |
| drome                                    | rhone_alpes                      | drome                                    | sPp     |
| durango                                  | mexico                           | durango                                  | sPp     |
| dusseldorf                               | nordrhein_westfalen              | dusseldorf                               | sPp     |
| east_flanders                            | flanders                         | east_flanders                            | sPp     |
| east_java                                | indonesia                        | east_java                                | sPp     |
| east_kalimantan                          | indonesia                        | east_kalimantan                          | sPp     |
| east_midlands                            | england                          | east_midlands                            | sPp     |
| east_nusa_tenggara                       | indonesia                        | east_nusa_tenggara                       | sPp     |
| east_timor                               | asia                             | east_timor                               | sPp     |
| eastern_anatolia                         | turkey                           | eastern_anatolia                         | sPp     |
| eastern_cape                             | south_africa                     | eastern_cape                             | sPp     |
| eastern_ontario                          | ontario                          | eastern_ontario                          | sPp     |
| eastern_visayas                          | philippines                      | eastern_visayas                          | sPp     |
| el_dorado                                | california                       | el_dorado                                | sPp     |
| el_salvador                              | central-america                  | el_salvador                              | sPp     |
| emilia_romagna                           | italy                            | emilia_romagna                           | sPp     |
| england                                  | united_kingdom                   | england                                  | sPp     |
| england_east                             | england                          | england_east                             | sPp     |
| england_north_east                       | england                          | england_north_east                       | sPp     |
| england_north_west                       | england                          | england_north_west                       | sPp     |
| england_south_east                       | england                          | england_south_east                       | sPp     |
| england_south_west                       | england                          | england_south_west                       | sPp     |
| entre_rios                               | argentina                        | entre_rios                               | sPp     |
| equatorial_guinea                        | africa                           | equatorial_guinea                        | sPp     |
| eritrea                                  | africa                           | eritrea                                  | sPp     |
| espirito-santo                           | southeast                        | espirito-santo                           | sPp     |
| essonne                                  | ile_de_france                    | essonne                                  | sPp     |
| estrie                                   | quebec                           | estrie                                   | sPp     |
| ethiopia                                 | africa                           | ethiopia                                 | sPp     |
| eure                                     | haute_normandie                  | eure                                     | sPp     |
| eure_et_loir                             | centre                           | eure_et_loir                             | sPp     |
| europe                                   |                                  | europe                                   | p       |
| europe_france                            | europe                           | europe_france                            | sPp     |
| euskadi                                  | spain                            | euskadi                                  | sPp     |
| evora                                    | portugal                         | evora                                    | sPp     |
| extremadura                              | spain                            | extremadura                              | sPp     |
| fairfax                                  | virginia                         | fairfax                                  | sPp     |
| falkland                                 | south-america                    | falkland                                 | sPp     |
| far_eastern_federal_district             | russia                           | far_eastern_federal_district             | sPp     |
| far_west                                 | new_south_wales                  | far_west                                 | sPp     |
| faro                                     | portugal                         | faro                                     | sPp     |
| fiji                                     | merge                            | fiji                                     | sPp     |
| fiji_east                                | oceania                          | fiji_east                                | sPp     |
| fiji_west                                | oceania                          | fiji_west                                | sPp     |
| finger_lakes                             | new-york                         | finger_lakes                             | sPp     |
| finistere                                | bretagne                         | finistere                                | sPp     |
| finland                                  | europe                           | finland                                  | sPp     |
| finnmark                                 | norway                           | finnmark                                 | sPp     |
| flanders                                 | belgium                          | flanders                                 | sPp     |
| flemish_brabant                          | flanders                         | flemish_brabant                          | sPp     |
| flevoland                                | netherlands                      | flevoland                                | sPp     |
| florida                                  | us-south                         | florida                                  | sPp     |
| florida_east_central                     | florida                          | florida_east_central                     | sPp     |
| florida_northeast                        | florida                          | florida_northeast                        | sPp     |
| florida_northwest                        | florida                          | florida_northwest                        | sPp     |
| florida_southwest                        | florida                          | florida_southwest                        | sPp     |
| formosa                                  | argentina                        | formosa                                  | sPp     |
| france                                   | europe                           | france                                   |         |
| france_metro_dom_com_nc                  | merge                            | france_metro_dom_com_nc                  | sPp     |
| franche_comte                            | france                           | franche_comte                            | sPp     |
| free_state                               | south_africa                     | free_state                               | sPp     |
| fresno                                   | california                       | fresno                                   | sPp     |
| fribourg                                 | switzerland                      | fribourg                                 | sPp     |
| friesland                                | netherlands                      | friesland                                | sPp     |
| friuli_venezia_giulia                    | italy                            | friuli_venezia_giulia                    | sPp     |
| fujian                                   | china                            | fujian                                   | sPp     |
| gabon                                    | africa                           | gabon                                    | sPp     |
| galicia                                  | spain                            | galicia                                  | sPp     |
| gambia                                   | africa                           | gambia                                   | sPp     |
| gansu                                    | china                            | gansu                                    | sPp     |
| gard                                     | languedoc_roussillon             | gard                                     | sPp     |
| gaspesie_iles_de_la_madeleine            | quebec                           | gaspesie_iles_de_la_madeleine            | sPp     |
| gatorland                                | florida                          | gatorland                                | sPp     |
| gauteng                                  | south_africa                     | gauteng                                  | sPp     |
| gavleborg                                | sweden                           | gavleborg                                | sPp     |
| gaza                                     | mozambique                       | gaza                                     | sPp     |
| gelderland                               | netherlands                      | gelderland                               | sPp     |
| geneva                                   | switzerland                      | geneva                                   | sPp     |
| georgia                                  | asia                             | georgia                                  | sPp     |
| georgia_northeast                        | georgia                          | georgia_northeast                        | sPp     |
| georgia_northwest                        | georgia                          | georgia_northwest                        | sPp     |
| georgia_southeast                        | georgia                          | georgia_southeast                        | sPp     |
| georgia_southwest                        | georgia                          | georgia_southwest                        | sPp     |
| germany                                  | europe                           | germany                                  | sPp     |
| gers                                     | midi_pyrenees                    | gers                                     | sPp     |
| ghana                                    | africa                           | ghana                                    | sPp     |
| gibraltar                                | europe                           | gibraltar                                | sPp     |
| girona                                   | catalunya                        | girona                                   | sPp     |
| gironde                                  | aquitaine                        | gironde                                  | sPp     |
| glarus                                   | switzerland                      | glarus                                   | sPp     |
| glenn                                    | california                       | glenn                                    | sPp     |
| goa                                      | india                            | goa                                      | sPp     |
| goias                                    | central-west                     | goias                                    | sPp     |
| gold_coast                               | florida                          | gold_coast                               | sPp     |
| golden_horseshoe                         | ontario                          | golden_horseshoe                         | sPp     |
| google40770f7a526b2e5f                   |                                  | google40770f7a526b2e5f                   |         |
| gorontalo                                | indonesia                        | gorontalo                                | sPp     |
| gotland                                  | sweden                           | gotland                                  | sPp     |
| granada                                  | andalucia                        | granada                                  | sPp     |
| greater_london                           | england                          | greater_london                           | sPp     |
| greater_metropolitan_newcastle           | new_south_wales                  | greater_metropolitan_newcastle           | sPp     |
| greater_metropolitan_sydney              | new_south_wales                  | greater_metropolitan_sydney              | sPp     |
| grenada                                  | central-america                  | grenada                                  | sPp     |
| grisons                                  | switzerland                      | grisons                                  | sPp     |
| groningen                                | netherlands                      | groningen                                | sPp     |
| guadalajara                              | castilla_la_mancha               | guadalajara                              | sPp     |
| guadeloupe                               | central-america                  | guadeloupe                               | sPp     |
| guam                                     | oceania                          | guam                                     | sPp     |
| guanajuato                               | mexico                           | guanajuato                               | sPp     |
| guangdong                                | china                            | guangdong                                | sPp     |
| guangxi                                  | china                            | guangxi                                  | sPp     |
| guarda                                   | portugal                         | guarda                                   | sPp     |
| guernesey                                | europe                           | guernesey                                | sPp     |
| guerrero                                 | mexico                           | guerrero                                 | sPp     |
| guinea                                   | africa                           | guinea                                   | sPp     |
| guipuzcoa                                | euskadi                          | guipuzcoa                                | sPp     |
| guizhou                                  | china                            | guizhou                                  | sPp     |
| gujarat                                  | india                            | gujarat                                  | sPp     |
| guyana                                   | south-america                    | guyana                                   | sPp     |
| guyane                                   | south-america                    | guyane                                   | sPp     |
| hainan                                   | china                            | hainan                                   | sPp     |
| halland                                  | sweden                           | halland                                  | sPp     |
| haryana                                  | india                            | haryana                                  | sPp     |
| haut_rhin                                | alsace                           | haut_rhin                                | sPp     |
| haute_corse                              | corse                            | haute_corse                              | sPp     |
| haute_garonne                            | midi_pyrenees                    | haute_garonne                            | sPp     |
| haute_loire                              | auvergne                         | haute_loire                              | sPp     |
| haute_marne                              | champagne_ardenne                | haute_marne                              | sPp     |
| haute_normandie                          | france                           | haute_normandie                          | sPp     |
| haute_saone                              | franche_comte                    | haute_saone                              | sPp     |
| haute_savoie                             | rhone_alpes                      | haute_savoie                             | sPp     |
| haute_vienne                             | limousin                         | haute_vienne                             | sPp     |
| hautes_alpes                             | provence_alpes_cote_d_azur       | hautes_alpes                             | sPp     |
| hautes_pyrenees                          | midi_pyrenees                    | hautes_pyrenees                          | sPp     |
| hauts_de_seine                           | ile_de_france                    | hauts_de_seine                           | sPp     |
| heard_island_and_mcdonald_slands         | australia                        | heard_island_and_mcdonald_slands         | sPp     |
| hebei                                    | china                            | hebei                                    | sPp     |
| hedmark                                  | norway                           | hedmark                                  | sPp     |
| heilongjiang                             | china                            | heilongjiang                             | sPp     |
| henan                                    | china                            | henan                                    | sPp     |
| herault                                  | languedoc_roussillon             | herault                                  | sPp     |
| hidalgo                                  | mexico                           | hidalgo                                  | sPp     |
| highland_papua                           | indonesia                        | highland_papua                           | sPp     |
| himachal_pradesh                         | india                            | himachal_pradesh                         | sPp     |
| hokkaido                                 | japan                            | hokkaido                                 | sPp     |
| honduras                                 | central-america                  | honduras                                 | sPp     |
| hong_kong                                | china                            | hong_kong                                | sPp     |
| hordaland                                | norway                           | hordaland                                | sPp     |
| hubei                                    | china                            | hubei                                    | sPp     |
| hudson_valley                            | new-york                         | hudson_valley                            | sPp     |
| huelva                                   | andalucia                        | huelva                                   | sPp     |
| huesca                                   | aragon                           | huesca                                   | sPp     |
| humboldt                                 | california                       | humboldt                                 | sPp     |
| hunan                                    | china                            | hunan                                    | sPp     |
| ile_de_france                            | france                           | ile_de_france                            | sPp     |
| ilemi                                    | africa                           | ilemi                                    | sPp     |
| illawarra                                | new_south_wales                  | illawarra                                | sPp     |
| ille_et_vilaine                          | bretagne                         | ille_et_vilaine                          | sPp     |
| illes_balears                            | spain                            | illes_balears                            | sPp     |
| illinois                                 | us-midwest                       | illinois                                 | sPp     |
| illinois_central                         | illinois                         | illinois_central                         | sPp     |
| illinois_east_central                    | illinois                         | illinois_east_central                    | sPp     |
| illinois_north                           | illinois                         | illinois_north                           | sPp     |
| illinois_northeast                       | illinois                         | illinois_northeast                       | sPp     |
| illinois_northwest                       | illinois                         | illinois_northwest                       | sPp     |
| illinois_southern                        | illinois                         | illinois_southern                        | sPp     |
| illinois_southwest                       | illinois                         | illinois_southwest                       | sPp     |
| ilocos_region                            | philippines                      | ilocos_region                            | sPp     |
| imperial                                 | california                       | imperial                                 | sPp     |
| india                                    | asia                             | india                                    | sPp     |
| indonesia                                | asia                             | indonesia                                | sPp     |
| indre                                    | centre                           | indre                                    | sPp     |
| indre_et_loire                           | centre                           | indre_et_loire                           | sPp     |
| ingushetia_republic                      | north_caucasian_federal_district | ingushetia_republic                      | sPp     |
| inhambane                                | mozambique                       | inhambane                                | sPp     |
| inner_mongolia                           | china                            | inner_mongolia                           | sPp     |
| inyo                                     | california                       | inyo                                     | sPp     |
| ionian_sea                               | seas                             | ionian_sea                               | sPp     |
| ireland                                  | europe                           | ireland                                  | sPp     |
| irkutsk_oblast                           | siberian_federal_district        | irkutsk_oblast                           | sPp     |
| isere                                    | rhone_alpes                      | isere                                    | sPp     |
| israel                                   | asia                             | israel                                   | sPp     |
| israel_and_palestine                     | asia                             | israel_and_palestine                     | sPp     |
| israel_west_bank                         | asia                             | israel_west_bank                         | sPp     |
| italy                                    | europe                           | italy                                    | sPp     |
| ivano-frankivsk_oblast                   | ukraine                          | ivano-frankivsk_oblast                   | sPp     |
| ivanovo_oblast                           | central_federal_district         | ivanovo_oblast                           | sPp     |
| ivory_coast                              | africa                           | ivory_coast                              | sPp     |
| jaen                                     | andalucia                        | jaen                                     | sPp     |
| jakarta                                  | indonesia                        | jakarta                                  | sPp     |
| jalisco                                  | mexico                           | jalisco                                  | sPp     |
| jamaica                                  | central-america                  | jamaica                                  | sPp     |
| jambi                                    | indonesia                        | jambi                                    | sPp     |
| jammu_and_kashmir                        | india                            | jammu_and_kashmir                        | sPp     |
| jamtland                                 | sweden                           | jamtland                                 | sPp     |
| jan_mayen                                | norway                           | jan_mayen                                | sPp     |
| japan                                    | asia                             | japan                                    | sPp     |
| jersey                                   | europe                           | jersey                                   | sPp     |
| jewish_autonomous_oblast                 | far_eastern_federal_district     | jewish_autonomous_oblast                 | sPp     |
| jharkhand                                | india                            | jharkhand                                | sPp     |
| jiangsu                                  | china                            | jiangsu                                  | sPp     |
| jiangxi                                  | china                            | jiangxi                                  | sPp     |
| jihocesky                                | czech_republic                   | jihocesky                                | sPp     |
| jihomoravsky                             | czech_republic                   | jihomoravsky                             | sPp     |
| jilin                                    | china                            | jilin                                    | sPp     |
| jonkoping                                | sweden                           | jonkoping                                | sPp     |
| jujuy                                    | argentina                        | jujuy                                    | sPp     |
| jura                                     | switzerland                      | jura                                     | sPp     |
| kabardino_balkar_republic                | north_caucasian_federal_district | kabardino_balkar_republic                | sPp     |
| kainuu                                   | finland                          | kainuu                                   | sPp     |
| kaliningrad_oblast                       | northwestern_federal_district    | kaliningrad_oblast                       | sPp     |
| kalmar                                   | sweden                           | kalmar                                   | sPp     |
| kalmykia_republic                        | southern_federal_district        | kalmykia_republic                        | sPp     |
| kaluga_oblast                            | central_federal_district         | kaluga_oblast                            | sPp     |
| kamchatka_krai                           | far_eastern_federal_district     | kamchatka_krai                           | sPp     |
| kansai                                   | japan                            | kansai                                   | sPp     |
| kanta_hame                               | finland                          | kanta_hame                               | sPp     |
| kanto                                    | japan                            | kanto                                    | sPp     |
| karachay_cherkess_republic               | north_caucasian_federal_district | karachay_cherkess_republic               | sPp     |
| karelia_republic                         | northwestern_federal_district    | karelia_republic                         | sPp     |
| karlovarsky                              | czech_republic                   | karlovarsky                              | sPp     |
| karnataka                                | india                            | karnataka                                | sPp     |
| karnten                                  | austria                          | karnten                                  | sPp     |
| kemerovo_oblast                          | siberian_federal_district        | kemerovo_oblast                          | sPp     |
| kenya                                    | africa                           | kenya                                    | sPp     |
| kerala                                   | india                            | kerala                                   | sPp     |
| kern                                     | california                       | kern                                     | sPp     |
| khabarovsk_krai                          | far_eastern_federal_district     | khabarovsk_krai                          | sPp     |
| khakassia_republic                       | siberian_federal_district        | khakassia_republic                       | sPp     |
| khanty_mansi_autonomous_okrug            | ural_federal_district            | khanty_mansi_autonomous_okrug            | sPp     |
| kharkiv_oblast                           | ukraine                          | kharkiv_oblast                           | sPp     |
| kherson_oblast                           | ukraine                          | kherson_oblast                           | sPp     |
| khmelnytskyi_oblast                      | ukraine                          | khmelnytskyi_oblast                      | sPp     |
| kiev                                     | ukraine                          | kiev                                     | sPp     |
| kiev_oblast                              | ukraine                          | kiev_oblast                              | sPp     |
| kings                                    | california                       | kings                                    | sPp     |
| kiribati                                 | merge                            | kiribati                                 | sPp     |
| kiribati_east                            | oceania                          | kiribati_east                            | sPp     |
| kiribati_west                            | oceania                          | kiribati_west                            | sPp     |
| kirov_oblast                             | volga_federal_district           | kirov_oblast                             | sPp     |
| kirovohrad_oblast                        | ukraine                          | kirovohrad_oblast                        | sPp     |
| koln                                     | nordrhein_westfalen              | koln                                     | sPp     |
| komi_republic                            | northwestern_federal_district    | komi_republic                            | sPp     |
| kosicky                                  | slovakia                         | kosicky                                  | sPp     |
| kostroma_oblast                          | central_federal_district         | kostroma_oblast                          | sPp     |
| kralovehradecky                          | czech_republic                   | kralovehradecky                          | sPp     |
| krasnodar_krai                           | southern_federal_district        | krasnodar_krai                           | sPp     |
| krasnoyarsk_krai                         | siberian_federal_district        | krasnoyarsk_krai                         | sPp     |
| kronoberg                                | sweden                           | kronoberg                                | sPp     |
| kujawsko_pomorskie                       | poland                           | kujawsko_pomorskie                       | sPp     |
| kurgan_oblast                            | ural_federal_district            | kurgan_oblast                            | sPp     |
| kursk_oblast                             | central_federal_district         | kursk_oblast                             | sPp     |
| kuwait                                   | asia                             | kuwait                                   | sPp     |
| kwazulu_natal                            | south_africa                     | kwazulu_natal                            | sPp     |
| kymenlaakso                              | finland                          | kymenlaakso                              | sPp     |
| kyushu                                   | japan                            | kyushu                                   | sPp     |
| la_coruna                                | galicia                          | la_coruna                                | sPp     |
| la_pampa                                 | argentina                        | la_pampa                                 | sPp     |
| la_rioja                                 | spain                            | la_rioja                                 | sPp     |
| ladakh                                   | india                            | ladakh                                   | sPp     |
| lakshadweep                              | india                            | lakshadweep                              | sPp     |
| lampung                                  | indonesia                        | lampung                                  | sPp     |
| lanaudiere                               | quebec                           | lanaudiere                               | sPp     |
| landes                                   | aquitaine                        | landes                                   | sPp     |
| languedoc_roussillon                     | france                           | languedoc_roussillon                     | sPp     |
| laos                                     | asia                             | laos                                     | sPp     |
| lapland                                  | finland                          | lapland                                  | sPp     |
| las_palmas                               | canarias                         | las_palmas                               | sPp     |
| lassen                                   | california                       | lassen                                   | sPp     |
| laurentides                              | quebec                           | laurentides                              | sPp     |
| laval                                    | quebec                           | laval                                    | sPp     |
| lazio                                    | italy                            | lazio                                    | sPp     |
| leiria                                   | portugal                         | leiria                                   | sPp     |
| leningrad_oblast                         | northwestern_federal_district    | leningrad_oblast                         | sPp     |
| leon                                     | castilla_y_leon                  | leon                                     | sPp     |
| lesotho                                  | africa                           | lesotho                                  | sPp     |
| liaoning                                 | china                            | liaoning                                 | sPp     |
| liberecky                                | czech_republic                   | liberecky                                | sPp     |
| liguria                                  | italy                            | liguria                                  | sPp     |
| limburg                                  | netherlands                      | limburg                                  | sPp     |
| limousin                                 | france                           | limousin                                 | sPp     |
| limpopo                                  | south_africa                     | limpopo                                  | sPp     |
| lipetsk_oblast                           | central_federal_district         | lipetsk_oblast                           | sPp     |
| lisbon                                   | portugal                         | lisbon                                   | sPp     |
| lleida                                   | catalunya                        | lleida                                   | sPp     |
| lodzkie                                  | poland                           | lodzkie                                  | sPp     |
| loir_et_cher                             | centre                           | loir_et_cher                             | sPp     |
| loire                                    | rhone_alpes                      | loire                                    | sPp     |
| loire_atlantique                         | pays_de_la_loire                 | loire_atlantique                         | sPp     |
| loiret                                   | centre                           | loiret                                   | sPp     |
| lombardia                                | italy                            | lombardia                                | sPp     |
| long_island                              | new-york                         | long_island                              | sPp     |
| lorraine                                 | france                           | lorraine                                 | sPp     |
| los_angeles                              | california                       | los_angeles                              | sPp     |
| los_lagos                                | chile                            | los_lagos                                | sPp     |
| los_rios                                 | chile                            | los_rios                                 | sPp     |
| lot                                      | midi_pyrenees                    | lot                                      | sPp     |
| lot_et_garonne                           | aquitaine                        | lot_et_garonne                           | sPp     |
| lozere                                   | languedoc_roussillon             | lozere                                   | sPp     |
| lubelskie                                | poland                           | lubelskie                                | sPp     |
| lubuskie                                 | poland                           | lubuskie                                 | sPp     |
| lucerne                                  | switzerland                      | lucerne                                  | sPp     |
| lugo                                     | galicia                          | lugo                                     | sPp     |
| luhansk_oblast                           | ukraine                          | luhansk_oblast                           | sPp     |
| luxembourg                               | europe                           | luxembourg                               | sPp     |
| lviv_oblast                              | ukraine                          | lviv_oblast                              | sPp     |
| macau                                    | china                            | macau                                    | sPp     |
| madeira                                  | portugal                         | madeira                                  | sPp     |
| madera                                   | california                       | madera                                   | sPp     |
| madhya_pradesh                           | india                            | madhya_pradesh                           | sPp     |
| magadan_oblast                           | far_eastern_federal_district     | magadan_oblast                           | sPp     |
| magallanes                               | chile                            | magallanes                               | sPp     |
| maharashtra                              | india                            | maharashtra                              | sPp     |
| maine_et_loire                           | pays_de_la_loire                 | maine_et_loire                           | sPp     |
| malaga                                   | andalucia                        | malaga                                   | sPp     |
| malawi                                   | africa                           | malawi                                   | sPp     |
| malaysia                                 | asia                             | malaysia                                 | sPp     |
| maldives                                 | asia                             | maldives                                 | sPp     |
| mali                                     | africa                           | mali                                     | sPp     |
| malopolskie                              | poland                           | malopolskie                              | sPp     |
| maluku                                   | indonesia                        | maluku                                   | sPp     |
| manche                                   | basse_normandie                  | manche                                   | sPp     |
| manica                                   | mozambique                       | manica                                   | sPp     |
| manipur                                  | india                            | manipur                                  | sPp     |
| manitoba                                 | canada                           | manitoba                                 | sPp     |
| maputo                                   | mozambique                       | maputo                                   | sPp     |
| maputo_city                              | mozambique                       | maputo_city                              | sPp     |
| maranhao                                 | northeast                        | maranhao                                 | sPp     |
| marche                                   | italy                            | marche                                   | sPp     |
| mari_el_republic                         | volga_federal_district           | mari_el_republic                         | sPp     |
| marin                                    | california                       | marin                                    | sPp     |
| mariposa                                 | california                       | mariposa                                 | sPp     |
| marmara                                  | turkey                           | marmara                                  | sPp     |
| marne                                    | champagne_ardenne                | marne                                    | sPp     |
| marshall-islands                         | oceania                          | marshall-islands                         | sPp     |
| marshall_islands                         | oceania                          | marshall_islands                         | sPp     |
| martinique                               | central-america                  | martinique                               | sPp     |
| mato-grosso                              | central-west                     | mato-grosso                              | sPp     |
| mato-grosso-do-sul                       | central-west                     | mato-grosso-do-sul                       | sPp     |
| maule                                    | chile                            | maule                                    | sPp     |
| mauricie                                 | quebec                           | mauricie                                 | sPp     |
| mauritania                               | africa                           | mauritania                               | sPp     |
| mauritius                                | africa                           | mauritius                                | sPp     |
| mayenne                                  | pays_de_la_loire                 | mayenne                                  | sPp     |
| mayotte                                  | africa                           | mayotte                                  | sPp     |
| mazowieckie                              | poland                           | mazowieckie                              | sPp     |
| mediterranean                            | turkey                           | mediterranean                            | sPp     |
| meghalaya                                | india                            | meghalaya                                | sPp     |
| melilla                                  | spain                            | melilla                                  | sPp     |
| mendocino                                | california                       | mendocino                                | sPp     |
| mendoza                                  | argentina                        | mendoza                                  | sPp     |
| merced                                   | california                       | merced                                   | sPp     |
| merge                                    |                                  | merge                                    |         |
| merge_france_taaf                        | merge                            | merge_france_taaf                        | sPp     |
| metro_manila                             | philippines                      | metro_manila                             | sPp     |
| meurthe_et_moselle                       | lorraine                         | meurthe_et_moselle                       | sPp     |
| meuse                                    | lorraine                         | meuse                                    | sPp     |
| mexico                                   | north-america                    | mexico                                   | sPp     |
| mexico_city                              | mexico                           | mexico_city                              | sPp     |
| michigan                                 | us-midwest                       | michigan                                 | sPp     |
| michigan_central                         | michigan                         | michigan_central                         | sPp     |
| michigan_southeast                       | michigan                         | michigan_southeast                       | sPp     |
| michigan_southwest                       | michigan                         | michigan_southwest                       | sPp     |
| michigan_west                            | michigan                         | michigan_west                            | sPp     |
| michoacan                                | mexico                           | michoacan                                | sPp     |
| micronesia                               | oceania                          | micronesia                               | sPp     |
| mid_north_coast                          | new_south_wales                  | mid_north_coast                          | sPp     |
| midi_pyrenees                            | france                           | midi_pyrenees                            | sPp     |
| mimaropa                                 | philippines                      | mimaropa                                 | sPp     |
| minas-gerais                             | southeast                        | minas-gerais                             | sPp     |
| misiones                                 | argentina                        | misiones                                 | sPp     |
| mizoram                                  | india                            | mizoram                                  | sPp     |
| modoc                                    | california                       | modoc                                    | sPp     |
| moere_og_romsdal                         | norway                           | moere_og_romsdal                         | sPp     |
| mohawk_valley                            | new-york                         | mohawk_valley                            | sPp     |
| molise                                   | italy                            | molise                                   | sPp     |
| monaco                                   | europe                           | monaco                                   | sPp     |
| mono                                     | california                       | mono                                     | sPp     |
| monteregie                               | quebec                           | monteregie                               | sPp     |
| monterey                                 | california                       | monterey                                 | sPp     |
| montreal                                 | quebec                           | montreal                                 | sPp     |
| montserrat                               | central-america                  | montserrat                               | sPp     |
| moravskoslezsky                          | czech_republic                   | moravskoslezsky                          | sPp     |
| morbihan                                 | bretagne                         | morbihan                                 | sPp     |
| mordovia_republic                        | volga_federal_district           | mordovia_republic                        | sPp     |
| morelos                                  | mexico                           | morelos                                  | sPp     |
| moscow                                   | central_federal_district         | moscow                                   | sPp     |
| moscow_oblast                            | central_federal_district         | moscow_oblast                            | sPp     |
| moselle                                  | lorraine                         | moselle                                  | sPp     |
| mozambique                               | africa                           | mozambique                               | sPp     |
| mpumalanga                               | south_africa                     | mpumalanga                               | sPp     |
| munster                                  | nordrhein_westfalen              | munster                                  | sPp     |
| murmansk_oblast                          | northwestern_federal_district    | murmansk_oblast                          | sPp     |
| murray                                   | new_south_wales                  | murray                                   | sPp     |
| myanmar                                  | asia                             | myanmar                                  | sPp     |
| mykolaiv_oblast                          | ukraine                          | mykolaiv_oblast                          | sPp     |
| nagaland                                 | india                            | nagaland                                 | sPp     |
| namibia                                  | africa                           | namibia                                  | sPp     |
| nampula                                  | mozambique                       | nampula                                  | sPp     |
| napa                                     | california                       | napa                                     | sPp     |
| national_capital_territory_of_delhi      | india                            | national_capital_territory_of_delhi      | sPp     |
| nauru                                    | oceania                          | nauru                                    | sPp     |
| nayarit                                  | mexico                           | nayarit                                  | sPp     |
| nenets_autonomous_okrug                  | northwestern_federal_district    | nenets_autonomous_okrug                  | sPp     |
| netherlands                              | europe                           | netherlands                              | sPp     |
| neuchatel                                | switzerland                      | neuchatel                                | sPp     |
| neuquen                                  | argentina                        | neuquen                                  | sPp     |
| nevada                                   | california                       | nevada                                   | sPp     |
| new-york                                 | us-northeast                     | new-york                                 | sPp     |
| new_brunswick                            | canada                           | new_brunswick                            | sPp     |
| new_caledonia                            | oceania                          | new_caledonia                            | sPp     |
| new_south_wales                          | australia                        | new_south_wales                          | sPp     |
| new_south_wales_central_west             | new_south_wales                  | new_south_wales_central_west             | sPp     |
| new_south_wales_north_western            | new_south_wales                  | new_south_wales_north_western            | sPp     |
| new_south_wales_northern                 | new_south_wales                  | new_south_wales_northern                 | sPp     |
| new_york_city                            | new-york                         | new_york_city                            | sPp     |
| newfoundland_and_labrador                | canada                           | newfoundland_and_labrador                | sPp     |
| niassa                                   | mozambique                       | niassa                                   | sPp     |
| nicaragua                                | central-america                  | nicaragua                                | sPp     |
| nidwalden                                | switzerland                      | nidwalden                                | sPp     |
| niederosterreich                         | austria                          | niederosterreich                         | sPp     |
| nievre                                   | bourgogne                        | nievre                                   | sPp     |
| niger                                    | africa                           | niger                                    | sPp     |
| nigeria                                  | africa                           | nigeria                                  | sPp     |
| nigeria_north_central                    | nigeria                          | nigeria_north_central                    | sPp     |
| nigeria_north_east                       | nigeria                          | nigeria_north_east                       | sPp     |
| nigeria_north_west                       | nigeria                          | nigeria_north_west                       | sPp     |
| nigeria_south_east                       | nigeria                          | nigeria_south_east                       | sPp     |
| nigeria_south_south                      | nigeria                          | nigeria_south_south                      | sPp     |
| nigeria_south_west                       | nigeria                          | nigeria_south_west                       | sPp     |
| ningxia                                  | china                            | ningxia                                  | sPp     |
| nitriansky                               | slovakia                         | nitriansky                               | sPp     |
| niue                                     | oceania                          | niue                                     | sPp     |
| nizhny_novgorod_oblast                   | volga_federal_district           | nizhny_novgorod_oblast                   | sPp     |
| noord_brabant                            | netherlands                      | noord_brabant                            | sPp     |
| noord_holland                            | netherlands                      | noord_holland                            | sPp     |
| nord                                     | nord_pas_de_calais               | nord                                     | sPp     |
| nord_du_quebec                           | quebec                           | nord_du_quebec                           | sPp     |
| nord_pas_de_calais                       | france                           | nord_pas_de_calais                       | sPp     |
| nordland                                 | norway                           | nordland                                 | sPp     |
| nordrhein_westfalen                      | germany                          | nordrhein_westfalen                      | sPp     |
| norfolk_island                           | australia                        | norfolk_island                           | sPp     |
| norrbotten                               | sweden                           | norrbotten                               | sPp     |
| north                                    | brazil                           | north                                    |         |
| north-america                            |                                  | north-america                            | p       |
| north-carolina                           | us-south                         | north-carolina                           | sPp     |
| north-carolina_north_central             | north-carolina                   | north-carolina_north_central             | sPp     |
| north-carolina_northeast                 | north-carolina                   | north-carolina_northeast                 | sPp     |
| north-carolina_northwest                 | north-carolina                   | north-carolina_northwest                 | sPp     |
| north-carolina_south_central             | north-carolina                   | north-carolina_south_central             | sPp     |
| north-carolina_southeast                 | north-carolina                   | north-carolina_southeast                 | sPp     |
| north-carolina_western                   | north-carolina                   | north-carolina_western                   | sPp     |
| north_caucasian_federal_district         | russia                           | north_caucasian_federal_district         | sPp     |
| north_country                            | new-york                         | north_country                            | sPp     |
| north_kalimantan                         | indonesia                        | north_kalimantan                         | sPp     |
| north_karelia                            | finland                          | north_karelia                            | sPp     |
| north_maluku                             | indonesia                        | north_maluku                             | sPp     |
| north_metro                              | georgia                          | north_metro                              | sPp     |
| north_ossetia_alania_republic            | north_caucasian_federal_district | north_ossetia_alania_republic            | sPp     |
| north_ostrobothnia                       | finland                          | north_ostrobothnia                       | sPp     |
| north_savo                               | finland                          | north_savo                               | sPp     |
| north_sea                                | seas                             | north_sea                                | sPp     |
| north_sulawesi                           | indonesia                        | north_sulawesi                           | sPp     |
| north_sumatra                            | indonesia                        | north_sumatra                            | sPp     |
| northeast                                | brazil                           | northeast                                |         |
| northeastern_ontario                     | ontario                          | northeastern_ontario                     | sPp     |
| northern_cape                            | south_africa                     | northern_cape                            | sPp     |
| northern_ireland                         | united_kingdom                   | northern_ireland                         | sPp     |
| northern_lower                           | michigan                         | northern_lower                           | sPp     |
| northern_mariana_islands                 | oceania                          | northern_mariana_islands                 | sPp     |
| northern_mindanao                        | philippines                      | northern_mindanao                        | sPp     |
| northern_territory                       | australia                        | northern_territory                       | sPp     |
| northwest_territories                    | canada                           | northwest_territories                    | sPp     |
| northwestern_federal_district            | russia                           | northwestern_federal_district            | sPp     |
| northwestern_ontario                     | ontario                          | northwestern_ontario                     | sPp     |
| norway                                   | europe                           | norway                                   | sPp     |
| nova_scotia                              | canada                           | nova_scotia                              | sPp     |
| novgorod_oblast                          | northwestern_federal_district    | novgorod_oblast                          | sPp     |
| novosibirsk_oblast                       | siberian_federal_district        | novosibirsk_oblast                       | sPp     |
| nuble                                    | chile                            | nuble                                    | sPp     |
| nuevo_leon                               | mexico                           | nuevo_leon                               | sPp     |
| nunavut                                  | canada                           | nunavut                                  | sPp     |
| o_higgins                                | chile                            | o_higgins                                | sPp     |
| oaxaca                                   | mexico                           | oaxaca                                   | sPp     |
| oberosterreich                           | austria                          | oberosterreich                           | sPp     |
| obwalden                                 | switzerland                      | obwalden                                 | sPp     |
| oceania                                  |                                  | oceania                                  | sPp     |
| oceania_france_taaf                      | oceania                          | oceania_france_taaf                      | sPp     |
| odessa_oblast                            | ukraine                          | odessa_oblast                            | sPp     |
| odisha                                   | india                            | odisha                                   | sPp     |
| oestfold                                 | norway                           | oestfold                                 | sPp     |
| oise                                     | picardie                         | oise                                     | sPp     |
| olomoucky                                | czech_republic                   | olomoucky                                | sPp     |
| oman                                     | asia                             | oman                                     | sPp     |
| omsk_oblast                              | siberian_federal_district        | omsk_oblast                              | sPp     |
| ontario                                  | canada                           | ontario                                  | sPp     |
| opolskie                                 | poland                           | opolskie                                 | sPp     |
| oppland                                  | norway                           | oppland                                  | sPp     |
| orange                                   | california                       | orange                                   | sPp     |
| orebro                                   | sweden                           | orebro                                   | sPp     |
| orenburg_oblast                          | volga_federal_district           | orenburg_oblast                          | sPp     |
| orne                                     | basse_normandie                  | orne                                     | sPp     |
| oryol_oblast                             | central_federal_district         | oryol_oblast                             | sPp     |
| oslo                                     | norway                           | oslo                                     | sPp     |
| ostergotland                             | sweden                           | ostergotland                             | sPp     |
| ostrobothnia                             | finland                          | ostrobothnia                             | sPp     |
| ourense                                  | galicia                          | ourense                                  | sPp     |
| outaouais                                | quebec                           | outaouais                                | sPp     |
| overijssel                               | netherlands                      | overijssel                               | sPp     |
| paijat_hame                              | finland                          | paijat_hame                              | sPp     |
| palau                                    | oceania                          | palau                                    | sPp     |
| palencia                                 | castilla_y_leon                  | palencia                                 | sPp     |
| palestine                                | asia                             | palestine                                | sPp     |
| panama                                   | central-america                  | panama                                   | sPp     |
| panhandle                                | florida                          | panhandle                                | sPp     |
| papua                                    | indonesia                        | papua                                    | sPp     |
| papua_new_guinea                         | oceania                          | papua_new_guinea                         | sPp     |
| para                                     | north                            | para                                     | sPp     |
| paraguay                                 | south-america                    | paraguay                                 | sPp     |
| paraiba                                  | northeast                        | paraiba                                  | sPp     |
| parana                                   | south                            | parana                                   | sPp     |
| pardubicky                               | czech_republic                   | pardubicky                               | sPp     |
| paris                                    | ile_de_france                    | paris                                    | sPp     |
| pas_de_calais                            | nord_pas_de_calais               | pas_de_calais                            | sPp     |
| pays_de_la_loire                         | france                           | pays_de_la_loire                         | sPp     |
| penza_oblast                             | volga_federal_district           | penza_oblast                             | sPp     |
| perm_krai                                | volga_federal_district           | perm_krai                                | sPp     |
| pernambuco                               | northeast                        | pernambuco                               | sPp     |
| philippines                              | asia                             | philippines                              | sPp     |
| piaui                                    | northeast                        | piaui                                    | sPp     |
| picardie                                 | france                           | picardie                                 | sPp     |
| piedmont_triad                           | north-carolina                   | piedmont_triad                           | sPp     |
| piemonte                                 | italy                            | piemonte                                 | sPp     |
| pirkanmaa                                | finland                          | pirkanmaa                                | sPp     |
| pitcairn                                 | oceania                          | pitcairn                                 | sPp     |
| placer                                   | california                       | placer                                   | sPp     |
| plumas                                   | california                       | plumas                                   | sPp     |
| plzensky                                 | czech_republic                   | plzensky                                 | sPp     |
| podkarpackie                             | poland                           | podkarpackie                             | sPp     |
| podlaskie                                | poland                           | podlaskie                                | sPp     |
| poitou_charentes                         | france                           | poitou_charentes                         | sPp     |
| poland                                   | europe                           | poland                                   | sPp     |
| poltava_oblast                           | ukraine                          | poltava_oblast                           | sPp     |
| polygons-merge                           | merge                            | polygons-merge                           |         |
| polygons-merge_france_taaf               | polygons-merge                   | polygons-merge_france_taaf               | p       |
| polynesie                                | oceania                          | polynesie                                | sPp     |
| pomorskie                                | poland                           | pomorskie                                | sPp     |
| pontevedra                               | galicia                          | pontevedra                               | sPp     |
| portalegre                               | portugal                         | portalegre                               | sPp     |
| porto                                    | portugal                         | porto                                    | sPp     |
| portugal                                 | europe                           | portugal                                 | sPp     |
| praha                                    | czech_republic                   | praha                                    | sPp     |
| presovsky                                | slovakia                         | presovsky                                | sPp     |
| primorsky_krai                           | far_eastern_federal_district     | primorsky_krai                           | sPp     |
| prince_edward_island                     | canada                           | prince_edward_island                     | sPp     |
| provence_alpes_cote_d_azur               | france                           | provence_alpes_cote_d_azur               | sPp     |
| pskov_oblast                             | northwestern_federal_district    | pskov_oblast                             | sPp     |
| puducherry                               | india                            | puducherry                               | sPp     |
| puebla                                   | mexico                           | puebla                                   | sPp     |
| puerto_rico                              | central-america                  | puerto_rico                              | sPp     |
| puglia                                   | italy                            | puglia                                   | sPp     |
| punjab                                   | india                            | punjab                                   | sPp     |
| puy_de_dome                              | auvergne                         | puy_de_dome                              | sPp     |
| pyrenees_atlantiques                     | aquitaine                        | pyrenees_atlantiques                     | sPp     |
| pyrenees_orientales                      | languedoc_roussillon             | pyrenees_orientales                      | sPp     |
| qatar                                    | asia                             | qatar                                    | sPp     |
| qinghai                                  | china                            | qinghai                                  | sPp     |
| quebec                                   | canada                           | quebec                                   | sPp     |
| queensland                               | australia                        | queensland                               | sPp     |
| queretaro                                | mexico                           | queretaro                                | sPp     |
| quintana_roo                             | mexico                           | quintana_roo                             | sPp     |
| rajasthan                                | india                            | rajasthan                                | sPp     |
| region_de_murcia                         | spain                            | region_de_murcia                         | sPp     |
| reunion                                  | africa                           | reunion                                  | sPp     |
| rhone                                    | rhone_alpes                      | rhone                                    | sPp     |
| rhone_alpes                              | france                           | rhone_alpes                              | sPp     |
| riau                                     | indonesia                        | riau                                     | sPp     |
| riau_islands                             | indonesia                        | riau_islands                             | sPp     |
| richmond                                 | virginia                         | richmond                                 | sPp     |
| richmond_tweed                           | new_south_wales                  | richmond_tweed                           | sPp     |
| rio-de-janeiro                           | southeast                        | rio-de-janeiro                           | sPp     |
| rio-grande-do-norte                      | northeast                        | rio-grande-do-norte                      | sPp     |
| rio-grande-do-sul                        | south                            | rio-grande-do-sul                        | sPp     |
| rio_negro                                | argentina                        | rio_negro                                | sPp     |
| riverside                                | california                       | riverside                                | sPp     |
| rivne_oblast                             | ukraine                          | rivne_oblast                             | sPp     |
| robots                                   |                                  | robots                                   |         |
| rogaland                                 | norway                           | rogaland                                 | sPp     |
| rondonia                                 | north                            | rondonia                                 | sPp     |
| roraima                                  | north                            | roraima                                  | sPp     |
| rostov_oblast                            | southern_federal_district        | rostov_oblast                            | sPp     |
| russia                                   |                                  | russia                                   | sPp     |
| rwanda                                   | africa                           | rwanda                                   | sPp     |
| ryazan_oblast                            | central_federal_district         | ryazan_oblast                            | sPp     |
| sacramento                               | california                       | sacramento                               | sPp     |
| saguenay_lac_saint_jean                  | quebec                           | saguenay_lac_saint_jean                  | sPp     |
| saint_barthelemy                         | central-america                  | saint_barthelemy                         | sPp     |
| saint_gallen                             | switzerland                      | saint_gallen                             | sPp     |
| saint_helena_ascension_tristan_da_cunha  | africa                           | saint_helena_ascension_tristan_da_cunha  | sPp     |
| saint_kitts_and_nevis                    | central-america                  | saint_kitts_and_nevis                    | sPp     |
| saint_lucia                              | central-america                  | saint_lucia                              | sPp     |
| saint_martin                             | central-america                  | saint_martin                             | sPp     |
| saint_petersburg                         | northwestern_federal_district    | saint_petersburg                         | sPp     |
| saint_pierre_et_miquelon                 | north-america                    | saint_pierre_et_miquelon                 | sPp     |
| saint_vincent_and_the_grenadines         | central-america                  | saint_vincent_and_the_grenadines         | sPp     |
| sakha_republic                           | far_eastern_federal_district     | sakha_republic                           | sPp     |
| sakhalin_oblast                          | far_eastern_federal_district     | sakhalin_oblast                          | sPp     |
| salamanca                                | castilla_y_leon                  | salamanca                                | sPp     |
| salem                                    | virginia                         | salem                                    | sPp     |
| salta                                    | argentina                        | salta                                    | sPp     |
| salzburg                                 | austria                          | salzburg                                 | sPp     |
| samara_oblast                            | volga_federal_district           | samara_oblast                            | sPp     |
| samoa                                    | oceania                          | samoa                                    | sPp     |
| san_benito                               | california                       | san_benito                               | sPp     |
| san_bernardino                           | california                       | san_bernardino                           | sPp     |
| san_diego                                | california                       | san_diego                                | sPp     |
| san_francisco                            | california                       | san_francisco                            | sPp     |
| san_joaquin                              | california                       | san_joaquin                              | sPp     |
| san_juan                                 | argentina                        | san_juan                                 | sPp     |
| san_luis                                 | argentina                        | san_luis                                 | sPp     |
| san_luis_obispo                          | california                       | san_luis_obispo                          | sPp     |
| san_luis_potosi                          | mexico                           | san_luis_potosi                          | sPp     |
| san_marino                               | europe                           | san_marino                               | sPp     |
| san_mateo                                | california                       | san_mateo                                | sPp     |
| santa-catarina                           | south                            | santa-catarina                           | sPp     |
| santa_barbara                            | california                       | santa_barbara                            | sPp     |
| santa_clara                              | california                       | santa_clara                              | sPp     |
| santa_cruz                               | argentina                        | santa_cruz                               | sPp     |
| santa_cruz_de_tenerife                   | canarias                         | santa_cruz_de_tenerife                   | sPp     |
| santa_fe                                 | argentina                        | santa_fe                                 | sPp     |
| santarem                                 | portugal                         | santarem                                 | sPp     |
| santiago                                 | chile                            | santiago                                 | sPp     |
| santiago_del_estero                      | argentina                        | santiago_del_estero                      | sPp     |
| sao-paulo                                | southeast                        | sao-paulo                                | sPp     |
| sao_tome_and_principe                    | africa                           | sao_tome_and_principe                    | sPp     |
| saone_et_loire                           | bourgogne                        | saone_et_loire                           | sPp     |
| saratov_oblast                           | volga_federal_district           | saratov_oblast                           | sPp     |
| sardegna                                 | italy                            | sardegna                                 | sPp     |
| sarthe                                   | pays_de_la_loire                 | sarthe                                   | sPp     |
| saskatchewan                             | canada                           | saskatchewan                             | sPp     |
| satakunta                                | finland                          | satakunta                                | sPp     |
| saudi_arabia                             | asia                             | saudi_arabia                             | sPp     |
| savoie                                   | rhone_alpes                      | savoie                                   | sPp     |
| schaffhausen                             | switzerland                      | schaffhausen                             | sPp     |
| schwyz                                   | switzerland                      | schwyz                                   | sPp     |
| seas                                     | europe                           | seas                                     |         |
| segovia                                  | castilla_y_leon                  | segovia                                  | sPp     |
| seine_et_marne                           | ile_de_france                    | seine_et_marne                           | sPp     |
| seine_maritime                           | haute_normandie                  | seine_maritime                           | sPp     |
| seine_saint_denis                        | ile_de_france                    | seine_saint_denis                        | sPp     |
| senegal                                  | africa                           | senegal                                  | sPp     |
| sergipe                                  | northeast                        | sergipe                                  | sPp     |
| setubal                                  | portugal                         | setubal                                  | sPp     |
| sevastopol                               | ukraine                          | sevastopol                               | sPp     |
| sevilla                                  | andalucia                        | sevilla                                  | sPp     |
| seychelles                               | africa                           | seychelles                               | sPp     |
| shaanxi                                  | china                            | shaanxi                                  | sPp     |
| shandong                                 | china                            | shandong                                 | sPp     |
| shanghai                                 | china                            | shanghai                                 | sPp     |
| shanxi                                   | china                            | shanxi                                   | sPp     |
| shasta                                   | california                       | shasta                                   | sPp     |
| shikoku                                  | japan                            | shikoku                                  | sPp     |
| siberian_federal_district                | russia                           | siberian_federal_district                | sPp     |
| sichuan                                  | china                            | sichuan                                  | sPp     |
| sicilia                                  | italy                            | sicilia                                  | sPp     |
| sierra                                   | california                       | sierra                                   | sPp     |
| sikkim                                   | india                            | sikkim                                   | sPp     |
| sinaloa                                  | mexico                           | sinaloa                                  | sPp     |
| singapore                                | asia                             | singapore                                | sPp     |
| sint_maarten                             | central-america                  | sint_maarten                             | sPp     |
| siskiyou                                 | california                       | siskiyou                                 | sPp     |
| skane                                    | sweden                           | skane                                    | sPp     |
| slaskie                                  | poland                           | slaskie                                  | sPp     |
| slovakia                                 | europe                           | slovakia                                 | sPp     |
| smolensk_oblast                          | central_federal_district         | smolensk_oblast                          | sPp     |
| soccsksargen                             | philippines                      | soccsksargen                             | sPp     |
| sodermanland                             | sweden                           | sodermanland                             | sPp     |
| sofala                                   | mozambique                       | sofala                                   | sPp     |
| sogn_og_fjordane                         | norway                           | sogn_og_fjordane                         | sPp     |
| solano                                   | california                       | solano                                   | sPp     |
| solomon_islands                          | oceania                          | solomon_islands                          | sPp     |
| solothurn                                | switzerland                      | solothurn                                | sPp     |
| somme                                    | picardie                         | somme                                    | sPp     |
| sonoma                                   | california                       | sonoma                                   | sPp     |
| sonora                                   | mexico                           | sonora                                   | sPp     |
| soria                                    | castilla_y_leon                  | soria                                    | sPp     |
| south                                    | brazil                           | south                                    |         |
| south-america                            |                                  | south-america                            | sPp     |
| south_africa                             | africa                           | south_africa                             | sPp     |
| south_africa_north_west                  | south_africa                     | south_africa_north_west                  | sPp     |
| south_australia                          | australia                        | south_australia                          | sPp     |
| south_east_region                        | new_south_wales                  | south_east_region                        | sPp     |
| south_georgia_and_south_sandwich         | south-america                    | south_georgia_and_south_sandwich         | sPp     |
| south_kalimantan                         | indonesia                        | south_kalimantan                         | sPp     |
| south_karelia                            | finland                          | south_karelia                            | sPp     |
| south_ostrobothnia                       | finland                          | south_ostrobothnia                       | sPp     |
| south_papua                              | indonesia                        | south_papua                              | sPp     |
| south_savo                               | finland                          | south_savo                               | sPp     |
| south_sudan                              | africa                           | south_sudan                              | sPp     |
| south_sulawesi                           | indonesia                        | south_sulawesi                           | sPp     |
| south_sumatra                            | indonesia                        | south_sumatra                            | sPp     |
| southeast                                | brazil                           | southeast                                |         |
| southeast_sulawesi                       | indonesia                        | southeast_sulawesi                       | sPp     |
| southeastern_anatolia                    | turkey                           | southeastern_anatolia                    | sPp     |
| southern_federal_district                | russia                           | southern_federal_district                | sPp     |
| southern_highlands                       | tanzania                         | southern_highlands                       | sPp     |
| southern_tier                            | new-york                         | southern_tier                            | sPp     |
| southwest_finland                        | finland                          | southwest_finland                        | sPp     |
| southwest_papua                          | indonesia                        | southwest_papua                          | sPp     |
| southwestern_ontario                     | ontario                          | southwestern_ontario                     | sPp     |
| spain                                    | europe                           | spain                                    | sPp     |
| stanislaus                               | california                       | stanislaus                               | sPp     |
| state_of_mexico                          | mexico                           | state_of_mexico                          | sPp     |
| stavropol_krai                           | north_caucasian_federal_district | stavropol_krai                           | sPp     |
| steiermark                               | austria                          | steiermark                               | sPp     |
| stockholm                                | sweden                           | stockholm                                | sPp     |
| stredocesky                              | czech_republic                   | stredocesky                              | sPp     |
| sudan                                    | africa                           | sudan                                    | sPp     |
| sumy_oblast                              | ukraine                          | sumy_oblast                              | sPp     |
| suncoast                                 | florida                          | suncoast                                 | sPp     |
| suriname                                 | south-america                    | suriname                                 | sPp     |
| sutter                                   | california                       | sutter                                   | sPp     |
| svalbard                                 | norway                           | svalbard                                 | sPp     |
| sverdlovsk_oblast                        | ural_federal_district            | sverdlovsk_oblast                        | sPp     |
| swaziland                                | africa                           | swaziland                                | sPp     |
| sweden                                   | europe                           | sweden                                   | sPp     |
| swietokrzyskie                           | poland                           | swietokrzyskie                           | sPp     |
| switzerland                              | europe                           | switzerland                              | sPp     |
| sydney_surrounds                         | new_south_wales                  | sydney_surrounds                         | sPp     |
| tabasco                                  | mexico                           | tabasco                                  | sPp     |
| tamaulipas                               | mexico                           | tamaulipas                               | sPp     |
| tambov_oblast                            | central_federal_district         | tambov_oblast                            | sPp     |
| tamil_nadu                               | india                            | tamil_nadu                               | sPp     |
| tanzania                                 | africa                           | tanzania                                 | sPp     |
| tanzania_central                         | tanzania                         | tanzania_central                         | sPp     |
| tanzania_lake                            | tanzania                         | tanzania_lake                            | sPp     |
| tanzania_northern                        | tanzania                         | tanzania_northern                        | sPp     |
| tanzania_western                         | tanzania                         | tanzania_western                         | sPp     |
| tarapaca                                 | chile                            | tarapaca                                 | sPp     |
| tarn                                     | midi_pyrenees                    | tarn                                     | sPp     |
| tarn_et_garonne                          | midi_pyrenees                    | tarn_et_garonne                          | sPp     |
| tarragona                                | catalunya                        | tarragona                                | sPp     |
| tasmania                                 | australia                        | tasmania                                 | sPp     |
| tatarstan_republic                       | volga_federal_district           | tatarstan_republic                       | sPp     |
| tehama                                   | california                       | tehama                                   | sPp     |
| telangana                                | india                            | telangana                                | sPp     |
| telemark                                 | norway                           | telemark                                 | sPp     |
| ternopil_oblast                          | ukraine                          | ternopil_oblast                          | sPp     |
| territoire_de_belfort                    | franche_comte                    | territoire_de_belfort                    | sPp     |
| teruel                                   | aragon                           | teruel                                   | sPp     |
| tete                                     | mozambique                       | tete                                     | sPp     |
| texas                                    | us-south                         | texas                                    | sPp     |
| texas_central                            | texas                            | texas_central                            | sPp     |
| texas_north                              | texas                            | texas_north                              | sPp     |
| texas_northwest                          | texas                            | texas_northwest                          | sPp     |
| texas_south                              | texas                            | texas_south                              | sPp     |
| texas_southeast                          | texas                            | texas_southeast                          | sPp     |
| texas_west                               | texas                            | texas_west                               | sPp     |
| the_riverina                             | new_south_wales                  | the_riverina                             | sPp     |
| thurgau                                  | switzerland                      | thurgau                                  | sPp     |
| tianjin                                  | china                            | tianjin                                  | sPp     |
| tibet                                    | china                            | tibet                                    | sPp     |
| ticino                                   | switzerland                      | ticino                                   | sPp     |
| tierra_del_fuego                         | argentina                        | tierra_del_fuego                         | sPp     |
| tirol                                    | austria                          | tirol                                    | sPp     |
| tlaxcala                                 | mexico                           | tlaxcala                                 | sPp     |
| tocantins                                | north                            | tocantins                                | sPp     |
| togo                                     | africa                           | togo                                     | sPp     |
| tohoku                                   | japan                            | tohoku                                   | sPp     |
| tokelau                                  | oceania                          | tokelau                                  | sPp     |
| toledo                                   | castilla_la_mancha               | toledo                                   | sPp     |
| tomsk_oblast                             | siberian_federal_district        | tomsk_oblast                             | sPp     |
| tonga                                    | oceania                          | tonga                                    | sPp     |
| toscana                                  | italy                            | toscana                                  | sPp     |
| treasure_coast                           | florida                          | treasure_coast                           | sPp     |
| trenciansky                              | slovakia                         | trenciansky                              | sPp     |
| trentino_alto_adige                      | italy                            | trentino_alto_adige                      | sPp     |
| trinidad_and_tobago                      | central-america                  | trinidad_and_tobago                      | sPp     |
| trinity                                  | california                       | trinity                                  | sPp     |
| tripura                                  | india                            | tripura                                  | sPp     |
| trnavsky                                 | slovakia                         | trnavsky                                 | sPp     |
| troendelag                               | norway                           | troendelag                               | sPp     |
| troms                                    | norway                           | troms                                    | sPp     |
| tucuman                                  | argentina                        | tucuman                                  | sPp     |
| tula_oblast                              | central_federal_district         | tula_oblast                              | sPp     |
| tulare                                   | california                       | tulare                                   | sPp     |
| tunisia                                  | africa                           | tunisia                                  | sPp     |
| tuolumne                                 | california                       | tuolumne                                 | sPp     |
| turkey                                   | europe                           | turkey                                   | sPp     |
| turks_and_caicos_islands                 | central-america                  | turks_and_caicos_islands                 | sPp     |
| tuva_republic                            | siberian_federal_district        | tuva_republic                            | sPp     |
| tuvalu                                   | oceania                          | tuvalu                                   | sPp     |
| tver_oblast                              | central_federal_district         | tver_oblast                              | sPp     |
| tyumen_oblast                            | ural_federal_district            | tyumen_oblast                            | sPp     |
| udmurt_republic                          | volga_federal_district           | udmurt_republic                          | sPp     |
| uganda                                   | africa                           | uganda                                   | sPp     |
| uganda_central                           | uganda                           | uganda_central                           | sPp     |
| uganda_eastern                           | uganda                           | uganda_eastern                           | sPp     |
| uganda_northern                          | uganda                           | uganda_northern                          | sPp     |
| uganda_western                           | uganda                           | uganda_western                           | sPp     |
| ukraine                                  | europe                           | ukraine                                  | sPp     |
| ulyanovsk_oblast                         | volga_federal_district           | ulyanovsk_oblast                         | sPp     |
| umbria                                   | italy                            | umbria                                   | sPp     |
| united_arab_emirates                     | asia                             | united_arab_emirates                     | sPp     |
| united_kingdom                           | europe                           | united_kingdom                           | sPp     |
| united_states_virgin_islands             | central-america                  | united_states_virgin_islands             | sPp     |
| upper_peninsula                          | michigan                         | upper_peninsula                          | sPp     |
| uppsala                                  | sweden                           | uppsala                                  | sPp     |
| ural_federal_district                    | russia                           | ural_federal_district                    | sPp     |
| uri                                      | switzerland                      | uri                                      | sPp     |
| us-midwest                               | north-america                    | us-midwest                               | sPp     |
| us-northeast                             | north-america                    | us-northeast                             | sPp     |
| us-south                                 | north-america                    | us-south                                 | sPp     |
| us-west                                  | north-america                    | us-west                                  | sPp     |
| usa_virgin_islands                       | central-america                  | usa_virgin_islands                       | sPp     |
| ustecky                                  | czech_republic                   | ustecky                                  | sPp     |
| utrecht                                  | netherlands                      | utrecht                                  | sPp     |
| uttar_pradesh                            | india                            | uttar_pradesh                            | sPp     |
| uttarakhand                              | india                            | uttarakhand                              | sPp     |
| uusimaa                                  | finland                          | uusimaa                                  | sPp     |
| val_d_oise                               | ile_de_france                    | val_d_oise                               | sPp     |
| val_de_marne                             | ile_de_france                    | val_de_marne                             | sPp     |
| valais                                   | switzerland                      | valais                                   | sPp     |
| valencia                                 | comunitat_valenciana             | valencia                                 | sPp     |
| valladolid                               | castilla_y_leon                  | valladolid                               | sPp     |
| valle_aosta                              | italy                            | valle_aosta                              | sPp     |
| valparaiso                               | chile                            | valparaiso                               | sPp     |
| vanuatu                                  | oceania                          | vanuatu                                  | sPp     |
| var                                      | provence_alpes_cote_d_azur       | var                                      | sPp     |
| varmland                                 | sweden                           | varmland                                 | sPp     |
| vasterbotten                             | sweden                           | vasterbotten                             | sPp     |
| vasternorrland                           | sweden                           | vasternorrland                           | sPp     |
| vastmanland                              | sweden                           | vastmanland                              | sPp     |
| vastra_gotaland                          | sweden                           | vastra_gotaland                          | sPp     |
| vatican_city                             | europe                           | vatican_city                             | sPp     |
| vaucluse                                 | provence_alpes_cote_d_azur       | vaucluse                                 | sPp     |
| vaud                                     | switzerland                      | vaud                                     | sPp     |
| vendee                                   | pays_de_la_loire                 | vendee                                   | sPp     |
| veneto                                   | italy                            | veneto                                   | sPp     |
| venezuela                                | south-america                    | venezuela                                | sPp     |
| ventura                                  | california                       | ventura                                  | sPp     |
| veracruz                                 | mexico                           | veracruz                                 | sPp     |
| vest-agder                               | norway                           | vest-agder                               | sPp     |
| vestfold                                 | norway                           | vestfold                                 | sPp     |
| viana_do_castelo                         | portugal                         | viana_do_castelo                         | sPp     |
| victoria                                 | australia                        | victoria                                 | sPp     |
| vienne                                   | poitou_charentes                 | vienne                                   | sPp     |
| vila_real                                | portugal                         | vila_real                                | sPp     |
| vinnytsia_oblast                         | ukraine                          | vinnytsia_oblast                         | sPp     |
| virginia                                 | us-south                         | virginia                                 | sPp     |
| viseu                                    | portugal                         | viseu                                    | sPp     |
| vizcaya                                  | euskadi                          | vizcaya                                  | sPp     |
| vladimir_oblast                          | central_federal_district         | vladimir_oblast                          | sPp     |
| volga_federal_district                   | russia                           | volga_federal_district                   | sPp     |
| volgograd_oblast                         | southern_federal_district        | volgograd_oblast                         | sPp     |
| vologda_oblast                           | northwestern_federal_district    | vologda_oblast                           | sPp     |
| volyn_oblast                             | ukraine                          | volyn_oblast                             | sPp     |
| vorarlberg                               | austria                          | vorarlberg                               | sPp     |
| voronezh_oblast                          | central_federal_district         | voronezh_oblast                          | sPp     |
| vosges                                   | lorraine                         | vosges                                   | sPp     |
| vysocina                                 | czech_republic                   | vysocina                                 | sPp     |
| wallis_et_futuna                         | oceania                          | wallis_et_futuna                         | sPp     |
| wallonia_french_community                | belgium                          | wallonia_french_community                | sPp     |
| wallonia_german_community                | belgium                          | wallonia_german_community                | sPp     |
| warminsko_mazurskie                      | poland                           | warminsko_mazurskie                      | sPp     |
| west_bengal                              | india                            | west_bengal                              | sPp     |
| west_flanders                            | flanders                         | west_flanders                            | sPp     |
| west_java                                | indonesia                        | west_java                                | sPp     |
| west_kalimantan                          | indonesia                        | west_kalimantan                          | sPp     |
| west_midlands                            | england                          | west_midlands                            | sPp     |
| west_nusa_tenggara                       | indonesia                        | west_nusa_tenggara                       | sPp     |
| west_papua                               | indonesia                        | west_papua                               | sPp     |
| west_sulawesi                            | indonesia                        | west_sulawesi                            | sPp     |
| west_sumatra                             | indonesia                        | west_sumatra                             | sPp     |
| western_australia                        | australia                        | western_australia                        | sPp     |
| western_cape                             | south_africa                     | western_cape                             | sPp     |
| western_new_york                         | new-york                         | western_new_york                         | sPp     |
| western_sahara                           | africa                           | western_sahara                           | sPp     |
| western_visayas                          | philippines                      | western_visayas                          | sPp     |
| wielkopolskie                            | poland                           | wielkopolskie                            | sPp     |
| wien                                     | austria                          | wien                                     | sPp     |
| wytheville                               | virginia                         | wytheville                               | sPp     |
| xinjiang                                 | china                            | xinjiang                                 | sPp     |
| yamalo_nenets_autonomous_okrug           | ural_federal_district            | yamalo_nenets_autonomous_okrug           | sPp     |
| yaroslavl_oblast                         | central_federal_district         | yaroslavl_oblast                         | sPp     |
| yogyakarta                               | indonesia                        | yogyakarta                               | sPp     |
| yolo                                     | california                       | yolo                                     | sPp     |
| yonne                                    | bourgogne                        | yonne                                    | sPp     |
| yorkshire_and_the_humber                 | england                          | yorkshire_and_the_humber                 | sPp     |
| yuba                                     | california                       | yuba                                     | sPp     |
| yucatan                                  | mexico                           | yucatan                                  | sPp     |
| yukon                                    | canada                           | yukon                                    | sPp     |
| yunnan                                   | china                            | yunnan                                   | sPp     |
| yvelines                                 | ile_de_france                    | yvelines                                 | sPp     |
| zabaykalsky_krai                         | siberian_federal_district        | zabaykalsky_krai                         | sPp     |
| zacatecas                                | mexico                           | zacatecas                                | sPp     |
| zachodniopomorskie                       | poland                           | zachodniopomorskie                       | sPp     |
| zakarpattia_oblast                       | ukraine                          | zakarpattia_oblast                       | sPp     |
| zambezia                                 | mozambique                       | zambezia                                 | sPp     |
| zambia                                   | africa                           | zambia                                   | sPp     |
| zamboanga_peninsula                      | philippines                      | zamboanga_peninsula                      | sPp     |
| zamora                                   | castilla_y_leon                  | zamora                                   | sPp     |
| zanzibar                                 | tanzania                         | zanzibar                                 | sPp     |
| zaporizhia_oblast                        | ukraine                          | zaporizhia_oblast                        | sPp     |
| zaragoza                                 | aragon                           | zaragoza                                 | sPp     |
| zeeland                                  | netherlands                      | zeeland                                  | sPp     |
| zhejiang                                 | china                            | zhejiang                                 | sPp     |
| zhytomyr_oblast                          | ukraine                          | zhytomyr_oblast                          | sPp     |
| zilinsky                                 | slovakia                         | zilinsky                                 | sPp     |
| zimbabwe                                 | africa                           | zimbabwe                                 | sPp     |
| zlinsky                                  | czech_republic                   | zlinsky                                  | sPp     |
| zug                                      | switzerland                      | zug                                      | sPp     |
| zuid_holland                             | netherlands                      | zuid_holland                             | sPp     |
| zurich                                   | switzerland                      | zurich                                   | sPp     |
Total elements: 1192

## List of elements from bbbike.org
|    SHORTNAME     | IS IN |    LONG NAME     | FORMATS |
|------------------|-------|------------------|---------|
| Aachen           |       | Aachen           | PGS     |
| Aarhus           |       | Aarhus           | PGS     |
| Adelaide         |       | Adelaide         | PGS     |
| Albuquerque      |       | Albuquerque      | PGS     |
| Alexandria       |       | Alexandria       | PGS     |
| Amsterdam        |       | Amsterdam        | PGS     |
| Antwerpen        |       | Antwerpen        | PGS     |
| Arnhem           |       | Arnhem           | PGS     |
| Auckland         |       | Auckland         | PGS     |
| Augsburg         |       | Augsburg         | PGS     |
| Austin           |       | Austin           | PGS     |
| Baghdad          |       | Baghdad          | PGS     |
| Baku             |       | Baku             | PGS     |
| Balaton          |       | Balaton          | PGS     |
| Bamberg          |       | Bamberg          | PGS     |
| Bangkok          |       | Bangkok          | PGS     |
| Barcelona        |       | Barcelona        | PGS     |
| Basel            |       | Basel            | PGS     |
| Beijing          |       | Beijing          | PGS     |
| Beirut           |       | Beirut           | PGS     |
| Berkeley         |       | Berkeley         | PGS     |
| Berlin           |       | Berlin           | PGS     |
| Bern             |       | Bern             | PGS     |
| Bielefeld        |       | Bielefeld        | PGS     |
| Birmingham       |       | Birmingham       | PGS     |
| Bochum           |       | Bochum           | PGS     |
| Bogota           |       | Bogota           | PGS     |
| Bombay           |       | Bombay           | PGS     |
| Bonn             |       | Bonn             | PGS     |
| Bordeaux         |       | Bordeaux         | PGS     |
| Boulder          |       | Boulder          | PGS     |
| BrandenburgHavel |       | BrandenburgHavel | PGS     |
| Braunschweig     |       | Braunschweig     | PGS     |
| Bremen           |       | Bremen           | PGS     |
| Bremerhaven      |       | Bremerhaven      | PGS     |
| Brisbane         |       | Brisbane         | PGS     |
| Bristol          |       | Bristol          | PGS     |
| Brno             |       | Brno             | PGS     |
| Bruegge          |       | Bruegge          | PGS     |
| Bruessel         |       | Bruessel         | PGS     |
| Budapest         |       | Budapest         | PGS     |
| BuenosAires      |       | BuenosAires      | PGS     |
| Cairo            |       | Cairo            | PGS     |
| Calgary          |       | Calgary          | PGS     |
| Cambridge        |       | Cambridge        | PGS     |
| CambridgeMa      |       | CambridgeMa      | PGS     |
| Canberra         |       | Canberra         | PGS     |
| CapeTown         |       | CapeTown         | PGS     |
| Chemnitz         |       | Chemnitz         | PGS     |
| Chicago          |       | Chicago          | PGS     |
| ClermontFerrand  |       | ClermontFerrand  | PGS     |
| Colmar           |       | Colmar           | PGS     |
| Copenhagen       |       | Copenhagen       | PGS     |
| Cork             |       | Cork             | PGS     |
| Corsica          |       | Corsica          | PGS     |
| Corvallis        |       | Corvallis        | PGS     |
| Cottbus          |       | Cottbus          | PGS     |
| Cracow           |       | Cracow           | PGS     |
| CraterLake       |       | CraterLake       | PGS     |
| Curitiba         |       | Curitiba         | PGS     |
| Cusco            |       | Cusco            | PGS     |
| Dallas           |       | Dallas           | PGS     |
| Darmstadt        |       | Darmstadt        | PGS     |
| Davis            |       | Davis            | PGS     |
| DenHaag          |       | DenHaag          | PGS     |
| Denver           |       | Denver           | PGS     |
| Dessau           |       | Dessau           | PGS     |
| Dortmund         |       | Dortmund         | PGS     |
| Dresden          |       | Dresden          | PGS     |
| Dublin           |       | Dublin           | PGS     |
| Duesseldorf      |       | Duesseldorf      | PGS     |
| Duisburg         |       | Duisburg         | PGS     |
| Edinburgh        |       | Edinburgh        | PGS     |
| Eindhoven        |       | Eindhoven        | PGS     |
| Emden            |       | Emden            | PGS     |
| Erfurt           |       | Erfurt           | PGS     |
| Erlangen         |       | Erlangen         | PGS     |
| Eugene           |       | Eugene           | PGS     |
| Flensburg        |       | Flensburg        | PGS     |
| FortCollins      |       | FortCollins      | PGS     |
| Frankfurt        |       | Frankfurt        | PGS     |
| FrankfurtOder    |       | FrankfurtOder    | PGS     |
| Freiburg         |       | Freiburg         | PGS     |
| Gdansk           |       | Gdansk           | PGS     |
| Genf             |       | Genf             | PGS     |
| Gent             |       | Gent             | PGS     |
| Gera             |       | Gera             | PGS     |
| Glasgow          |       | Glasgow          | PGS     |
| Gliwice          |       | Gliwice          | PGS     |
| Goerlitz         |       | Goerlitz         | PGS     |
| Goeteborg        |       | Goeteborg        | PGS     |
| Goettingen       |       | Goettingen       | PGS     |
| Graz             |       | Graz             | PGS     |
| Groningen        |       | Groningen        | PGS     |
| Halifax          |       | Halifax          | PGS     |
| Halle            |       | Halle            | PGS     |
| Hamburg          |       | Hamburg          | PGS     |
| Hamm             |       | Hamm             | PGS     |
| Hannover         |       | Hannover         | PGS     |
| Heilbronn        |       | Heilbronn        | PGS     |
| Helsinki         |       | Helsinki         | PGS     |
| Hertogenbosch    |       | Hertogenbosch    | PGS     |
| Huntsville       |       | Huntsville       | PGS     |
| Innsbruck        |       | Innsbruck        | PGS     |
| Istanbul         |       | Istanbul         | PGS     |
| Jena             |       | Jena             | PGS     |
| Jerusalem        |       | Jerusalem        | PGS     |
| Johannesburg     |       | Johannesburg     | PGS     |
| Kaiserslautern   |       | Kaiserslautern   | PGS     |
| Karlsruhe        |       | Karlsruhe        | PGS     |
| Kassel           |       | Kassel           | PGS     |
| Katowice         |       | Katowice         | PGS     |
| Kaunas           |       | Kaunas           | PGS     |
| Kiel             |       | Kiel             | PGS     |
| Kiew             |       | Kiew             | PGS     |
| Koblenz          |       | Koblenz          | PGS     |
| Koeln            |       | Koeln            | PGS     |
| Konstanz         |       | Konstanz         | PGS     |
| LaPaz            |       | LaPaz            | PGS     |
| LaPlata          |       | LaPlata          | PGS     |
| LakeGarda        |       | LakeGarda        | PGS     |
| Lausanne         |       | Lausanne         | PGS     |
| Leeds            |       | Leeds            | PGS     |
| Leipzig          |       | Leipzig          | PGS     |
| Lima             |       | Lima             | PGS     |
| Linz             |       | Linz             | PGS     |
| Lisbon           |       | Lisbon           | PGS     |
| Liverpool        |       | Liverpool        | PGS     |
| Ljubljana        |       | Ljubljana        | PGS     |
| Lodz             |       | Lodz             | PGS     |
| London           |       | London           | PGS     |
| LosAngeles       |       | LosAngeles       | PGS     |
| Luebeck          |       | Luebeck          | PGS     |
| Luxemburg        |       | Luxemburg        | PGS     |
| Lyon             |       | Lyon             | PGS     |
| Maastricht       |       | Maastricht       | PGS     |
| Madison          |       | Madison          | PGS     |
| Madrid           |       | Madrid           | PGS     |
| Magdeburg        |       | Magdeburg        | PGS     |
| Mainz            |       | Mainz            | PGS     |
| Malmoe           |       | Malmoe           | PGS     |
| Manchester       |       | Manchester       | PGS     |
| Mannheim         |       | Mannheim         | PGS     |
| Marseille        |       | Marseille        | PGS     |
| Melbourne        |       | Melbourne        | PGS     |
| Memphis          |       | Memphis          | PGS     |
| MexicoCity       |       | MexicoCity       | PGS     |
| Miami            |       | Miami            | PGS     |
| Minsk            |       | Minsk            | PGS     |
| Moenchengladbach |       | Moenchengladbach | PGS     |
| Montevideo       |       | Montevideo       | PGS     |
| Montpellier      |       | Montpellier      | PGS     |
| Montreal         |       | Montreal         | PGS     |
| Moscow           |       | Moscow           | PGS     |
| Muenchen         |       | Muenchen         | PGS     |
| Muenster         |       | Muenster         | PGS     |
| NewDelhi         |       | NewDelhi         | PGS     |
| NewOrleans       |       | NewOrleans       | PGS     |
| NewYork          |       | NewYork          | PGS     |
| Nuernberg        |       | Nuernberg        | PGS     |
| Oldenburg        |       | Oldenburg        | PGS     |
| Oranienburg      |       | Oranienburg      | PGS     |
| Orlando          |       | Orlando          | PGS     |
| Oslo             |       | Oslo             | PGS     |
| Osnabrueck       |       | Osnabrueck       | PGS     |
| Ostrava          |       | Ostrava          | PGS     |
| Ottawa           |       | Ottawa           | PGS     |
| Paderborn        |       | Paderborn        | PGS     |
| Palma            |       | Palma            | PGS     |
| PaloAlto         |       | PaloAlto         | PGS     |
| Paris            |       | Paris            | PGS     |
| Perth            |       | Perth            | PGS     |
| Philadelphia     |       | Philadelphia     | PGS     |
| PhnomPenh        |       | PhnomPenh        | PGS     |
| Portland         |       | Portland         | PGS     |
| PortlandME       |       | PortlandME       | PGS     |
| Porto            |       | Porto            | PGS     |
| PortoAlegre      |       | PortoAlegre      | PGS     |
| Potsdam          |       | Potsdam          | PGS     |
| Poznan           |       | Poznan           | PGS     |
| Prag             |       | Prag             | PGS     |
| Providence       |       | Providence       | PGS     |
| Regensburg       |       | Regensburg       | PGS     |
| Riga             |       | Riga             | PGS     |
| RiodeJaneiro     |       | RiodeJaneiro     | PGS     |
| Rostock          |       | Rostock          | PGS     |
| Rotterdam        |       | Rotterdam        | PGS     |
| Ruegen           |       | Ruegen           | PGS     |
| Saarbruecken     |       | Saarbruecken     | PGS     |
| Sacramento       |       | Sacramento       | PGS     |
| Saigon           |       | Saigon           | PGS     |
| Salzburg         |       | Salzburg         | PGS     |
| SanFrancisco     |       | SanFrancisco     | PGS     |
| SanJose          |       | SanJose          | PGS     |
| SanktPetersburg  |       | SanktPetersburg  | PGS     |
| SantaBarbara     |       | SantaBarbara     | PGS     |
| SantaCruz        |       | SantaCruz        | PGS     |
| Santiago         |       | Santiago         | PGS     |
| Sarajewo         |       | Sarajewo         | PGS     |
| Schwerin         |       | Schwerin         | PGS     |
| Seattle          |       | Seattle          | PGS     |
| Seoul            |       | Seoul            | PGS     |
| Sheffield        |       | Sheffield        | PGS     |
| Singapore        |       | Singapore        | PGS     |
| Sofia            |       | Sofia            | PGS     |
| Stockholm        |       | Stockholm        | PGS     |
| Stockton         |       | Stockton         | PGS     |
| Strassburg       |       | Strassburg       | PGS     |
| Stuttgart        |       | Stuttgart        | PGS     |
| Sucre            |       | Sucre            | PGS     |
| Sydney           |       | Sydney           | PGS     |
| Szczecin         |       | Szczecin         | PGS     |
| Tallinn          |       | Tallinn          | PGS     |
| Tehran           |       | Tehran           | PGS     |
| Tilburg          |       | Tilburg          | PGS     |
| Tokyo            |       | Tokyo            | PGS     |
| Toronto          |       | Toronto          | PGS     |
| Toulouse         |       | Toulouse         | PGS     |
| Trondheim        |       | Trondheim        | PGS     |
| Tucson           |       | Tucson           | PGS     |
| Turin            |       | Turin            | PGS     |
| UlanBator        |       | UlanBator        | PGS     |
| Ulm              |       | Ulm              | PGS     |
| Usedom           |       | Usedom           | PGS     |
| Utrecht          |       | Utrecht          | PGS     |
| Vancouver        |       | Vancouver        | PGS     |
| Victoria         |       | Victoria         | PGS     |
| WarenMueritz     |       | WarenMueritz     | PGS     |
| Warsaw           |       | Warsaw           | PGS     |
| WashingtonDC     |       | WashingtonDC     | PGS     |
| Waterloo         |       | Waterloo         | PGS     |
| Wien             |       | Wien             | PGS     |
| Wroclaw          |       | Wroclaw          | PGS     |
| Wuerzburg        |       | Wuerzburg        | PGS     |
| Wuppertal        |       | Wuppertal        | PGS     |
| Zagreb           |       | Zagreb           | PGS     |
| Zuerich          |       | Zuerich          | PGS     |
Total elements: 237

## List of elements from osmtoday
|                SHORTNAME                 |             IS IN              |           LONG NAME            | FORMATS |
|------------------------------------------|--------------------------------|--------------------------------|---------|
|                                          |                                |                                | Ppg     |
| aargau                                   | Switzerland                    | Aargau                         | Ppg     |
| abitibi_temiscamingue                    | Quebec                         | Abitibi - Temiskaming          | Ppg     |
| abruzzo                                  | Italy                          | Abruzzo                        | Ppg     |
| aceh                                     | Indonesia                      | Aceh                           | Ppg     |
| acre                                     | North                          | Acre                           | Ppg     |
| adygea_republic                          | Southern Federal District      | Republic of Adygea             | Ppg     |
| afghanistan                              | Asia                           | Afghanistan                    | Ppg     |
| africa                                   |                                | Africa                         | Ppg     |
| aguascalientes                           | Mexico                         | Aguascalientes                 | Ppg     |
| akershus                                 | Norway                         | Akershus                       | Ppg     |
| alabama                                  | United States of America       | Alabama                        | Ppg     |
| alagoas                                  | Northeast                      | Alagoas                        | Ppg     |
| aland                                    | Finland                        | Åland Islands                  | Ppg     |
| alaska                                   | United States of America       | Alaska                         | Ppg     |
| alava                                    | Euskadi                        | Alava                          | Ppg     |
| albacete                                 | Castile-La Mancha              | Albacete                       | Ppg     |
| alberta                                  | Canada                         | Alberta                        | Ppg     |
| algeria                                  | Africa                         | Algeria                        | Ppg     |
| alicante                                 | Valencia                       | Alicante                       | Ppg     |
| almeria                                  | Andalusia                      | Almeria                        | Ppg     |
| alsace                                   | France                         | Alsace                         | Ppg     |
| altai_krai                               | Siberian Federal District      | Altai region                   | Ppg     |
| altai_republic                           | Siberian Federal District      | Altai Republic                 | Ppg     |
| amapa                                    | North                          | Amapa                          | Ppg     |
| amazonas                                 | North                          | Amazons                        | Ppg     |
| american-oceania                         | Australia Oceania              | America-Oceania                | Ppg     |
| amur_oblast                              | Far Eastern Federal District   | Amur region                    | Ppg     |
| andalucia                                | Spain                          | Andalusia                      | Ppg     |
| andaman_and_nicobar_islands              | India                          | Andaman and Nicobar Islands    | Ppg     |
| andhra_pradesh                           | India                          | Andhra Pradesh                 | Ppg     |
| angola                                   | Africa                         | Angola                         | Ppg     |
| anguilla                                 | Central America                | Anguilla                       | Ppg     |
| anhui                                    | China                          | anhui                          | Ppg     |
| antarctica                               |                                | Antarctica                     | Ppg     |
| antigua_and_barbuda                      | Central America                | Antigua and Barbuda            | Ppg     |
| antwerp                                  | Flanders                       | Antwerp                        | Ppg     |
| appenzell_ausserrhoden                   | Switzerland                    | Appenzell Ausserrhoden         | Ppg     |
| appenzell_innerrhoden                    | Switzerland                    | Appenzell-Innerrhoden          | Ppg     |
| aquitaine                                | France                         | Aquitaine                      | Ppg     |
| aragon                                   | Spain                          | Aragon                         | Ppg     |
| argentina                                | South America                  | Argentina                      | Ppg     |
| arizona                                  | United States of America       | Arizona                        | Ppg     |
| arkansas                                 | United States of America       | Arkansas                       | Ppg     |
| arkhangelsk_oblast                       | Northwestern Federal District  | Arkhangelsk region             | Ppg     |
| armenia                                  | Asia                           | Armenia                        | Ppg     |
| arnsberg                                 | North Rhine Westphalia         | Arnsberg                       | Ppg     |
| aruba                                    | Central America                | Aruba                          | Ppg     |
| arunachal_pradesh                        | India                          | Arunachal Pradesh              | Ppg     |
| asia                                     |                                | Asia                           | Ppg     |
| assam                                    | India                          | Assam                          | Ppg     |
| astrakhan_oblast                         | Southern Federal District      | Astrakhan Region               | Ppg     |
| asturias                                 | Spain                          | Asturias                       | Ppg     |
| aust-agder                               | Norway                         | Aust-Agder                     | Ppg     |
| australia                                | Australia Oceania              | Australia                      | Ppg     |
| australia_oceania                        |                                | Australia Oceania              | Ppg     |
| austria                                  | Europe                         | Austria                        | Ppg     |
| auvergne                                 | France                         | Auvergne                       | Ppg     |
| avila                                    | Castile and Leon               | Avila                          | Ppg     |
| azerbaijan                               | Asia                           | Azerbaijan                     | Ppg     |
| azores                                   | Portugal                       | Azores                         | Ppg     |
| badajoz                                  | Extremadura                    | Badajoz                        | Ppg     |
| baden_wuerttemberg                       | Germany                        | Baden Württemberg              | Ppg     |
| bahamas                                  | Central America                | Bahamas                        | Ppg     |
| bahia                                    | Northeast                      | Bahia                          | Ppg     |
| bahrain                                  | Asia                           | Bahrain                        | Ppg     |
| baja_california                          | Mexico                         | Baja California                | Ppg     |
| baja_california_sur                      | Mexico                         | Baja California Sur            | Ppg     |
| bali                                     | Indonesia                      | Bali                           | Ppg     |
| bangka_belitung_islands                  | Indonesia                      | Bangka Belitung Islands        | Ppg     |
| bangladesh                               | Asia                           | Bangladesh                     | Ppg     |
| banskobystricky                          | Slovakia                       | Banskofast                     | Ppg     |
| banten                                   | Indonesia                      | banten                         | Ppg     |
| barbados                                 | Central America                | Barbados                       | Ppg     |
| barcelona                                | Catalonia                      | Barcelona                      | Ppg     |
| bas_saint_laurent                        | Quebec                         | Lower Saint Lawrence           | Ppg     |
| basel_landschaft                         | Switzerland                    | Basel Land                     | Ppg     |
| basel_stadt                              | Switzerland                    | Basel-Stadt                    | Ppg     |
| bashkortostan_republic                   | Volga Federal District         | Republic of Bashkortostan      | Ppg     |
| basilicata                               | Italy                          | Basilicata                     | Ppg     |
| basse_normandie                          | France                         | Lower Normandy                 | Ppg     |
| bayern                                   | Germany                        | Bavaria                        | Ppg     |
| bedfordshire                             | England                        | Bedfordshire                   | Ppg     |
| beijing                                  | China                          | Beijing                        | Ppg     |
| belarus                                  | Europe                         | Belarus                        | Ppg     |
| belgium                                  | Europe                         | Belgium                        | Ppg     |
| belgorod_oblast                          | Central Federal District       | Belgorod region                | Ppg     |
| belize                                   | Central America                | Belize                         | Ppg     |
| bengkulu                                 | Indonesia                      | Bengkulu                       | Ppg     |
| benin                                    | Africa                         | Benin                          | Ppg     |
| berkshire                                | England                        | Berkshire                      | Ppg     |
| berlin                                   | Germany                        | Berlin                         | Ppg     |
| bermuda                                  | North America                  | Bermuda                        | Ppg     |
| bern                                     | Switzerland                    | Berne                          | Ppg     |
| bhutan                                   | Asia                           | Butane                         | Ppg     |
| bihar                                    | India                          | Bihar                          | Ppg     |
| blekinge                                 | Sweden                         | Blekinge                       | Ppg     |
| bolivia                                  | South America                  | Bolivia                        | Ppg     |
| botswana                                 | Africa                         | Botswana                       | Ppg     |
| bourgogne                                | France                         | Burgundian                     | Ppg     |
| bouvet_island                            | Africa                         | Bouvet Island                  | Ppg     |
| brandenburg                              | Germany                        | Brandenburg                    | Ppg     |
| bratislavsky                             | Slovakia                       | Bratislava                     | Ppg     |
| brazil                                   | South America                  | Brazil                         | Ppg     |
| bremen                                   | Germany                        | Bremen                         | Ppg     |
| brestakaya                               | Belarus                        | Brest region                   | Ppg     |
| bretagne                                 | France                         | Brittany                       | Ppg     |
| bristol                                  | England                        | Bristol                        | Ppg     |
| british_columbia                         | Canada                         | British Columbia               | Ppg     |
| british_indian_ocean_territory           | Asia                           | British Indian Ocean Territory | Ppg     |
| british_virgin_islands                   | Central America                | British Virgin Islands         | Ppg     |
| brunei                                   | Asia                           | Brunei                         | Ppg     |
| brussels_capital_region                  | Belgium                        | Brussels Capital Region        | Ppg     |
| bryansk_oblast                           | Central Federal District       | Bryansk Region                 | Ppg     |
| buckinghamshire                          | England                        | Buckinghamshire                | Ppg     |
| buenos_aires                             | Argentina                      | Buenos Aires                   | Ppg     |
| buenos_aires_city                        | Argentina                      | city of Buenos Aires           | Ppg     |
| burgenland                               | Austria                        | Burgenland                     | Ppg     |
| burgos                                   | Castile and Leon               | Burgos                         | Ppg     |
| burkina_faso                             | Africa                         | Burkina Faso                   | Ppg     |
| burundi                                  | Africa                         | Burundi                        | Ppg     |
| buryatia_republic                        | Siberian Federal District      | The Republic of Buryatia       | Ppg     |
| buskerud                                 | Norway                         | Buskerud                       | Ppg     |
| caceres                                  | Extremadura                    | Caceres                        | Ppg     |
| cadiz                                    | Andalusia                      | Cadiz                          | Ppg     |
| calabria                                 | Italy                          | Calabria                       | Ppg     |
| california                               | United States of America       | California                     | Ppg     |
| cambodia                                 | Asia                           | Cambodia                       | Ppg     |
| cambridgeshire                           | England                        | Cambridgeshire                 | Ppg     |
| cameroon                                 | Africa                         | Cameroon                       | Ppg     |
| campania                                 | Italy                          | Campaign                       | Ppg     |
| campeche                                 | Mexico                         | campeche                       | Ppg     |
| canada                                   | North America                  | Canada                         | Ppg     |
| canary-islands                           | Africa                         | Canary Islands                 | Ppg     |
| cantabria                                | Spain                          | Cantabria                      | Ppg     |
| cape_verde                               | Africa                         | Cape Verde                     | Ppg     |
| capitale_nationale                       | Quebec                         | Capital Nacional               | Ppg     |
| caribbean                                | Central America                | caribbean                      | Ppg     |
| castellon                                | Valencia                       | Castellón                      | Ppg     |
| castilla_la_mancha                       | Spain                          | Castile-La Mancha              | Ppg     |
| castilla_y_leon                          | Spain                          | Castile and Leon               | Ppg     |
| catalunya                                | Spain                          | Catalonia                      | Ppg     |
| catamarca                                | Argentina                      | Catamarca                      | Ppg     |
| cayman_islands                           | Central America                | Cayman islands                 | Ppg     |
| ceara                                    | Northeast                      | Ceara                          | Ppg     |
| central-west                             | Brazil                         | Central Western                | Ppg     |
| central-zone                             | India                          | Central Zone                   | Ppg     |
| central_african_republic                 | Africa                         | Central African Republic       | Ppg     |
| central_america                          |                                | Central America                | Ppg     |
| central_federal_district                 | Russia                         | Central Federal District       | Ppg     |
| central_finland                          | Finland                        | Central Finland                | Ppg     |
| central_java                             | Indonesia                      | Central Java                   | Ppg     |
| central_kalimantan                       | Indonesia                      | Central Kalimantan             | Ppg     |
| central_ontario                          | Ontario                        | Central Ontario                | Ppg     |
| central_ostrobothnia                     | Finland                        | Central Ostrobothnia           | Ppg     |
| central_sulawesi                         | Indonesia                      | Central Sulawesi               | Ppg     |
| centre                                   | France                         | Center                         | Ppg     |
| centre_du_quebec                         | Quebec                         | Central Quebec                 | Ppg     |
| chaco                                    | Argentina                      | Chaco                          | Ppg     |
| chad                                     | Africa                         | Chad                           | Ppg     |
| champagne_ardenne                        | France                         | Champagne - Ardennes           | Ppg     |
| chandigarh                               | India                          | Chandigarh                     | Ppg     |
| chaudiere_appalaches                     | Quebec                         | Chaudieres - Appalachians      | Ppg     |
| chechen_republic                         | North Caucasian Federal        | Chechen Republic               | Ppg     |
|                                          | District                       |                                |         |
| chelyabinsk_oblast                       | Ural federal district          | Chelyabinsk region             | Ppg     |
| cherkasy_oblast                          | Ukraine                        | Cherkasy region                | Ppg     |
| chernihiv_oblast                         | Ukraine                        | Chernihiv region               | Ppg     |
| chernivtsi_oblast                        | Ukraine                        | Chernivtsi Region              | Ppg     |
| cheshire                                 | England                        | Cheshire                       | Ppg     |
| chhattisgarh                             | India                          | Chhattisgarh                   | Ppg     |
| chiapas                                  | Mexico                         | Chiapas                        | Ppg     |
| chihuahua                                | Mexico                         | Chihuahua                      | Ppg     |
| chile                                    | South America                  | Chile                          | Ppg     |
| china                                    | Asia                           | China                          | Ppg     |
| chongqing                                | China                          | chongqing                      | Ppg     |
| chubu                                    | Japan                          | Chubu                          | Ppg     |
| chubut                                   | Argentina                      | Chubut                         | Ppg     |
| chugoku                                  | Japan                          | Chugoku                        | Ppg     |
| chukotka_autonomous_okrug                | Far Eastern Federal District   | Chukotka Autonomous Okrug      | Ppg     |
| chuvash_republic                         | Volga Federal District         | Chuvash Republic               | Ppg     |
| ciudad_real                              | Castile-La Mancha              | Ciudad Real                    | Ppg     |
| coahuila                                 | Mexico                         | Coahuila                       | Ppg     |
| colima                                   | Mexico                         | Colima                         | Ppg     |
| colombia                                 | South America                  | Colombia                       | Ppg     |
| colorado                                 | United States of America       | Colorado                       | Ppg     |
| comoros                                  | Africa                         | Comoros                        | Ppg     |
| comunidad_de_madrid                      | Spain                          | Madrid                         | Ppg     |
| comunidad_foral_de_navarra               | Spain                          | Navarre                        | Ppg     |
| comunitat_valenciana                     | Spain                          | Valencia                       | Ppg     |
| congo-democratic-republic                | Africa                         | Democratic Republic of the     | Ppg     |
|                                          |                                | Congo                          |         |
| congo_brazzaville                        | Africa                         | Brazzaville                    | Ppg     |
| congo_kinshasa                           | Africa                         | Kinshasa                       | Ppg     |
| connecticut                              | United States of America       | Connecticut                    | Ppg     |
| cook-islands                             | Australia Oceania              | Cook Islands                   | Ppg     |
| cordoba-andalucia                        | Andalusia                      | Cordoba                        | Pp      |
| cordoba-argentina                        | Argentina                      | Cordoba                        | Pp      |
| cornwall                                 | England                        | Cornwall                       | Ppg     |
| corrientes                               | Argentina                      | Corrientes                     | Ppg     |
| corse                                    | France                         | Kors                           | Ppg     |
| costa_rica                               | Central America                | Costa Rica                     | Ppg     |
| cote_nord                                | Quebec                         | Côte Nor                       | Ppg     |
| crimea                                   | Ukraine                        | Crimea                         | Ppg     |
| crimea_republic                          | Southern Federal District      | Republic of Crimea             | Ppg     |
| cuba                                     | Central America                | Cuba                           | Ppg     |
| cuenca                                   | Castile-La Mancha              | Cuenca                         | Ppg     |
| cumbria                                  | England                        | Cumbria                        | Ppg     |
| curacao                                  | Central America                | Curacao                        | Ppg     |
| czech_republic                           | Europe                         | Czech Republic                 | Ppg     |
| dadra_and_nagar_haveli                   | India                          | Dadra and Nagar Haveli         | Ppg     |
| dadra_and_nagar_haveli_and_daman_and_diu | India                          | Dadra and Nagar Haveli and     | Ppg     |
|                                          |                                | Daman and Diu                  |         |
| dagestan_republic                        | North Caucasian Federal        | The Republic of Dagestan       | Ppg     |
|                                          | District                       |                                |         |
| dalarna                                  | Sweden                         | Dalarna                        | Ppg     |
| daman_and_diu                            | India                          | Daman and Diu                  | Ppg     |
| delaware                                 | United States of America       | Delaware                       | Ppg     |
| derbyshire                               | England                        | Derbyshire                     | Ppg     |
| detmold                                  | North Rhine Westphalia         | Detmold                        | Ppg     |
| devon                                    | England                        | Devonian                       | Ppg     |
| district-of-columbia                     | United States of America       | District of Colombia           | Ppg     |
| distrito-federal                         | Central Western                | Distrito Federal               | Ppg     |
| djibouti                                 | Africa                         | Djibouti                       | Ppg     |
| dnipropetrovsk_oblast                    | Ukraine                        | Dnipropetrovsk Region          | Ppg     |
| dolnoslaskie                             | Poland                         | Dolnoslaskie                   | Ppg     |
| dominica                                 | Central America                | Dominica                       | Ppg     |
| dominican_republic                       | Central America                | Dominican Republic             | Ppg     |
| donetsk_oblast                           | Ukraine                        | Donetsk Oblast                 | Ppg     |
| dorset                                   | England                        | Dorset                         | Ppg     |
| drenthe                                  | Netherlands                    | Drenthe                        | Ppg     |
| durango                                  | Mexico                         | Durango                        | Ppg     |
| durham                                   | England                        | Durham                         | Ppg     |
| dusseldorf                               | North Rhine Westphalia         | Dusseldorf                     | Ppg     |
| east                                     | England                        | East                           | Ppg     |
| east_flanders                            | Flanders                       | East Flemish                   | Ppg     |
| east_java                                | Indonesia                      | East Java                      | Ppg     |
| east_kalimantan                          | Indonesia                      | East Kalimantan                | Ppg     |
| east_midlands                            | England                        | East Midlands                  | Ppg     |
| east_nusa_tenggara                       | Indonesia                      | East Nusa Tenggara             | Ppg     |
| east_sussex                              | England                        | East Sussex                    | Ppg     |
| east_timor                               | Asia                           | East Timor                     | Ppg     |
| east_yorkshire_with_hull                 | England                        | East Riding of Yorkshire       | Ppg     |
| eastern-zone                             | India                          | Eastern Zone                   | Ppg     |
| eastern_ontario                          | Ontario                        | Eastern Ontario                | Ppg     |
| ecuador                                  | South America                  | Ecuador                        | Ppg     |
| egypt                                    | Africa                         | Egypt                          | Ppg     |
| el_salvador                              | Central America                | Salvador                       | Ppg     |
| emilia_romagna                           | Italy                          | Emilia-Romagna                 | Ppg     |
| england                                  | United Kingdom                 | England                        | Ppg     |
| entre_rios                               | Argentina                      | Entre Rios                     | Ppg     |
| equatorial_guinea                        | Africa                         | Equatorial Guinea              | Ppg     |
| eritrea                                  | Africa                         | Eritrea                        | Ppg     |
| espirito-santo                           | Southeast                      | Duhito Santo                   | Ppg     |
| essex                                    | England                        | Essex                          | Ppg     |
| estrie                                   | Quebec                         | Estri                          | Ppg     |
| ethiopia                                 | Africa                         | Ethiopia                       | Ppg     |
| europe                                   |                                | Europe                         | Ppg     |
| euskadi                                  | Spain                          | Euskadi                        | Ppg     |
| extremadura                              | Spain                          | Extremadura                    | Ppg     |
| falkland                                 | South America                  | Falkland                       | Ppg     |
| far_eastern_federal_district             | Russia                         | Far Eastern Federal District   | Ppg     |
| fiji                                     | Australia Oceania              | Fiji                           | Ppg     |
| finland                                  | Europe                         | Finland                        | Ppg     |
| finnmark                                 | Norway                         | Finnmark                       | Ppg     |
| flanders                                 | Belgium                        | Flanders                       | Ppg     |
| flemish_brabant                          | Flanders                       | Flemish Brabant                | Ppg     |
| flevoland                                | Netherlands                    | Flevoland                      | Ppg     |
| florida                                  | United States of America       | Florida                        | Ppg     |
| formosa                                  | Argentina                      | Formosa                        | Ppg     |
| france                                   | Europe                         | France                         | Ppg     |
| france_taaf                              | Africa                         | French Southern and Antarctic  | Ppg     |
|                                          |                                | Territories                    |         |
| franche_comte                            | France                         | franche-comté                  | Ppg     |
| fribourg                                 | Switzerland                    | Friborg                        | Ppg     |
| friesland                                | Netherlands                    | friesland                      | Ppg     |
| friuli_venezia_giulia                    | Italy                          | Friuli Venezia Giulia          | Ppg     |
| fujian                                   | China                          | Fujian                         | Ppg     |
| gabon                                    | Africa                         | Gabon                          | Ppg     |
| galicia                                  | Spain                          | Galicia                        | Ppg     |
| gambia                                   | Africa                         | Gambia                         | Ppg     |
| gansu                                    | China                          | Gansu                          | Ppg     |
| gaspesie_iles_de_la_madeleine            | Quebec                         | Gaspesie - Madeleine Islands   | Ppg     |
| gavleborg                                | Sweden                         | Gävleborg                      | Ppg     |
| gcc_states                               | Asia                           | Cooperation Council for the    | Ppg     |
|                                          |                                | Arab States of the Gulf        |         |
| gelderland                               | Netherlands                    | Gelderland                     | Ppg     |
| geneva                                   | Switzerland                    | Geneva                         | Ppg     |
| georgia-asia                             | Asia                           | Georgia                        | Pp      |
| georgia-us                               | United States of America       | Georgia                        | Pp      |
| germany                                  | Europe                         | Germany                        | Ppg     |
| ghana                                    | Africa                         | Ghana                          | Ppg     |
| gibraltar                                | Europe                         | Gibraltar                      | Ppg     |
| girona                                   | Catalonia                      | Girona                         | Ppg     |
| glarus                                   | Switzerland                    | Glarus                         | Ppg     |
| gloucestershire                          | England                        | Gloucestershire                | Ppg     |
| goa                                      | India                          | Goa                            | Ppg     |
| goias                                    | Central Western                | Goiás                          | Ppg     |
| golden_horseshoe                         | Ontario                        | Golden Horseshoe               | Ppg     |
| gorontalo                                | Indonesia                      | Gorontalo                      | Ppg     |
| gotland                                  | Sweden                         | Gotland                        | Ppg     |
| granada                                  | Andalusia                      | Granada                        | Ppg     |
| greater_london                           | England                        | Greater London                 | Ppg     |
| greater_manchester                       | England                        | Greater Manchester             | Ppg     |
| greenland                                | North America                  | Greenland                      | Ppg     |
| grenada                                  | Central America                | Grenada                        | Ppg     |
| grisons                                  | Switzerland                    | Grisons                        | Ppg     |
| grodnenskaya                             | Belarus                        | The Grodno region              | Ppg     |
| groningen                                | Netherlands                    | Groningen                      | Ppg     |
| guadalajara                              | Castile-La Mancha              | Guadalajara                    | Ppg     |
| guadeloupe                               | Central America                | Guadeloupe                     | Ppg     |
| guanajuato                               | Mexico                         | Guanajuato                     | Ppg     |
| guangdong                                | China                          | Guangdong                      | Ppg     |
| guangxi                                  | China                          | Guangxi                        | Ppg     |
| guatemala                                | Central America                | Guatemala                      | Ppg     |
| guernesey                                | Europe                         | guernsey                       | Ppg     |
| guerrero                                 | Mexico                         | Guerrero                       | Ppg     |
| guinea                                   | Africa                         | Guinea                         | Ppg     |
| guipuzcoa                                | Euskadi                        | Gipuzkoa                       | Ppg     |
| guizhou                                  | China                          | Guizhou                        | Ppg     |
| gujarat                                  | India                          | Gujarat                        | Ppg     |
| guyana                                   | South America                  | Guyana                         | Ppg     |
| guyane                                   | South America                  | Guyan                          | Ppg     |
| guyane-france                            | France                         | Guyan                          | Pp      |
| hainan                                   | China                          | Hainan                         | Ppg     |
| haiti_and_domrep                         | Central America                | Republic of Haiti and          | Ppg     |
|                                          |                                | Domenican Republic             |         |
| halland                                  | Sweden                         | Halland                        | Ppg     |
| hamburg                                  | Germany                        | Hamburg                        | Ppg     |
| hampshire                                | England                        | Hampshire                      | Ppg     |
| haryana                                  | India                          | Haryana                        | Ppg     |
| haute_normandie                          | France                         | Upper Normandy                 | Ppg     |
| hawaii                                   | United States of America       | Hawaii                         | Ppg     |
| hebei                                    | China                          | Hebei                          | Ppg     |
| hedmark                                  | Norway                         | headmark                       | Ppg     |
| heilongjiang                             | China                          | heilongjiang                   | Ppg     |
| henan                                    | China                          | Henan                          | Ppg     |
| herefordshire                            | England                        | Herefordshire                  | Ppg     |
| hertfordshire                            | England                        | Hertfordshire                  | Ppg     |
| hessen                                   | Germany                        | Hesse                          | Ppg     |
| hidalgo                                  | Mexico                         | Hidalgo                        | Ppg     |
| himachal_pradesh                         | India                          | Himachal Pradesh               | Ppg     |
| hokkaido                                 | Japan                          | Hokkaido                       | Ppg     |
| homelskay                                | Belarus                        | Gomel region                   | Ppg     |
| honduras                                 | Central America                | Honduras                       | Ppg     |
| hong_kong                                | China                          | Hong Kong                      | Ppg     |
| hordaland                                | Norway                         | Hordaland                      | Ppg     |
| hubei                                    | China                          | Hubei                          | Ppg     |
| huelva                                   | Andalusia                      | Huelva                         | Ppg     |
| huesca                                   | Aragon                         | Huesca                         | Ppg     |
| hunan                                    | China                          | Hunan                          | Ppg     |
| idaho                                    | United States of America       | Idaho                          | Ppg     |
| ile-de-clipperton                        | Australia Oceania              | Ile De Clipperton              | Ppg     |
| ile_de_france                            | France                         | Île de France                  | Ppg     |
| illes_balears                            | Spain                          | Balearic Islands               | Ppg     |
| illinois                                 | United States of America       | Illinois                       | Ppg     |
| india                                    | Asia                           | India                          | Ppg     |
| indiana                                  | United States of America       | Indiana                        | Ppg     |
| indonesia                                | Asia                           | Indonesia                      | Ppg     |
| ingushetia_republic                      | North Caucasian Federal        | Republic of Ingushetia         | Ppg     |
|                                          | District                       |                                |         |
| inner_mongolia                           | China                          | Inner Mongolia                 | Ppg     |
| iowa                                     | United States of America       | Iowa                           | Ppg     |
| iran                                     | Asia                           | Iran                           | Ppg     |
| iraq                                     | Asia                           | Iraq                           | Ppg     |
| ireland                                  | Europe                         | Ireland                        | Ppg     |
| irkutsk_oblast                           | Siberian Federal District      | Irkutsk region                 | Ppg     |
| isle_of_wight                            | England                        | Isle of Wight                  | Ppg     |
| israel                                   | Asia                           | Israel                         | Ppg     |
| israel_and_palestine                     | Asia                           | Israel and Palestine           | Ppg     |
| italy                                    | Europe                         | Italy                          | Ppg     |
| ivano-frankivsk_oblast                   | Ukraine                        | Ivano-Frankivsk region         | Ppg     |
| ivanovo_oblast                           | Central Federal District       | Ivanovo Region                 | Ppg     |
| ivory_coast                              | Africa                         | Ivory Coast                    | Ppg     |
| jaen                                     | Andalusia                      | Haen                           | Ppg     |
| jakarta                                  | Indonesia                      | Jakarta                        | Ppg     |
| jalisco                                  | Mexico                         | Jalisco                        | Ppg     |
| jamaica                                  | Central America                | Jamaica                        | Ppg     |
| jambi                                    | Indonesia                      | Jambi                          | Ppg     |
| jammu_and_kashmir                        | India                          | Jammu and Kashmir              | Ppg     |
| jamtland                                 | Sweden                         | gemtland                       | Ppg     |
| jan_mayen                                | Norway                         | Jan Mayen                      | Ppg     |
| japan                                    | Asia                           | Japan                          | Ppg     |
| java                                     | Indonesia                      | Java                           | Ppg     |
| jersey                                   | Europe                         | channel islands                | Ppg     |
| jewish_autonomous_oblast                 | Far Eastern Federal District   | Jewish Autonomous Region       | Ppg     |
| jharkhand                                | India                          | Jharkhand                      | Ppg     |
| jiangsu                                  | China                          | Jiangsu                        | Ppg     |
| jiangxi                                  | China                          | Jiangxi                        | Ppg     |
| jihocesky                                | Czech Republic                 | Jigocheski                     | Ppg     |
| jihomoravsky                             | Czech Republic                 | Jichomoravian                  | Ppg     |
| jilin                                    | China                          | Jilin                          | Ppg     |
| jonkoping                                | Sweden                         | Yonkoping                      | Ppg     |
| jordan                                   | Asia                           | Jordan                         | Ppg     |
| jujuy                                    | Argentina                      | Jujuy                          | Ppg     |
| jura                                     | Switzerland                    | Yura                           | Ppg     |
| kabardino_balkar_republic                | North Caucasian Federal        | Kabardino Balkar Republic      | Ppg     |
|                                          | District                       |                                |         |
| kainuu                                   | Finland                        | Kainuu                         | Ppg     |
| kalimantan                               | Indonesia                      | kalimantan                     | Ppg     |
| kaliningrad_oblast                       | Northwestern Federal District  | Kaliningrad region             | Ppg     |
| kalmar                                   | Sweden                         | Squid                          | Ppg     |
| kalmykia_republic                        | Southern Federal District      | Republic of Kalmykia           | Ppg     |
| kaluga_oblast                            | Central Federal District       | Kaluga region                  | Ppg     |
| kamchatka_krai                           | Far Eastern Federal District   | Kamchatka                      | Ppg     |
| kansai                                   | Japan                          | Kansai                         | Ppg     |
| kansas                                   | United States of America       | Kansas                         | Ppg     |
| kanta_hame                               | Finland                        | Kanta-Hame                     | Ppg     |
| kanto                                    | Japan                          | Kanto                          | Ppg     |
| karachay_cherkess_republic               | North Caucasian Federal        | Karachaev Circassian Republic  | Ppg     |
|                                          | District                       |                                |         |
| karelia_republic                         | Northwestern Federal District  | Republic of Karelia            | Ppg     |
| karlovarsky                              | Czech Republic                 | Karlovy Vary                   | Ppg     |
| karnataka                                | India                          | Karnataka                      | Ppg     |
| karnten                                  | Austria                        | Carnten                        | Ppg     |
| kazakhstan                               | Asia                           | Kazakhstan                     | Ppg     |
| kemerovo_oblast                          | Siberian Federal District      | Kemerovo Region                | Ppg     |
| kent                                     | England                        | Kent                           | Ppg     |
| kentucky                                 | United States of America       | Kentucky                       | Ppg     |
| kenya                                    | Africa                         | Kenya                          | Ppg     |
| kerala                                   | India                          | Kerala                         | Ppg     |
| khabarovsk_krai                          | Far Eastern Federal District   | Khabarovsk region              | Ppg     |
| khakassia_republic                       | Siberian Federal District      | The Republic of Khakassia      | Ppg     |
| khanty_mansi_autonomous_okrug            | Ural federal district          | Khanty-Mansi Autonomous Okrug  | Ppg     |
| kharkiv_oblast                           | Ukraine                        | Kharkov region                 | Ppg     |
| kherson_oblast                           | Ukraine                        | Kherson region                 | Ppg     |
| khmelnytskyi_oblast                      | Ukraine                        | Khmelnitsky region             | Ppg     |
| kiev                                     | Ukraine                        | Kyiv                           | Ppg     |
| kiev_oblast                              | Ukraine                        | Kyiv region                    | Ppg     |
| kiribati                                 | Australia Oceania              | Kiribati                       | Ppg     |
| kirov_oblast                             | Volga Federal District         | Kirov region                   | Ppg     |
| kirovohrad_oblast                        | Ukraine                        | Kirovograd Region              | Ppg     |
| koln                                     | North Rhine Westphalia         | Column                         | Ppg     |
| komi_republic                            | Northwestern Federal District  | Komi Republic                  | Ppg     |
| kosicky                                  | Slovakia                       | Kositsky                       | Ppg     |
| kostroma_oblast                          | Central Federal District       | Kostroma region                | Ppg     |
| kralovehradecky                          | Czech Republic                 | Kralovogradetsky               | Ppg     |
| krasnodar_krai                           | Southern Federal District      | Krasnodar region               | Ppg     |
| krasnoyarsk_krai                         | Siberian Federal District      | Krasnoyarsk Region             | Ppg     |
| kronoberg                                | Sweden                         | Kronoberg                      | Ppg     |
| kujawsko_pomorskie                       | Poland                         | Kuyavia-Pomeranian             | Ppg     |
| kurgan_oblast                            | Ural federal district          | Kurgan region                  | Ppg     |
| kursk_oblast                             | Central Federal District       | Kursk Region                   | Ppg     |
| kuwait                                   | Asia                           | Kuwait                         | Ppg     |
| kymenlaakso                              | Finland                        | Kymenlaakso                    | Ppg     |
| kyrgyzstan                               | Asia                           | Kyrgyzstan                     | Ppg     |
| kyushu                                   | Japan                          | Kyushu                         | Ppg     |
| la_coruna                                | Galicia                        | La Coruna                      | Ppg     |
| la_pampa                                 | Argentina                      | La Pampa                       | Ppg     |
| la_rioja-argentina                       | Argentina                      | Rioja                          | Pp      |
| la_rioja-spain                           | Spain                          | Rioja                          | Pp      |
| lakshadweep                              | India                          | Lakshadweep                    | Ppg     |
| lampung                                  | Indonesia                      | Lampung                        | Ppg     |
| lanaudiere                               | Quebec                         | Lanaudière                     | Ppg     |
| lancashire                               | England                        | Lancashire                     | Ppg     |
| languedoc_roussillon                     | France                         | Languedoc Roussillon           | Ppg     |
| laos                                     | Asia                           | Laos                           | Ppg     |
| lapland                                  | Finland                        | Lapland                        | Ppg     |
| laurentides                              | Quebec                         | Laurentides                    | Ppg     |
| laval                                    | Quebec                         | Laval                          | Ppg     |
| lazio                                    | Italy                          | Lazio                          | Ppg     |
| lebanon                                  | Asia                           | Lebanon                        | Ppg     |
| leicestershire                           | England                        | Leicestershire                 | Ppg     |
| leningrad_oblast                         | Northwestern Federal District  | Leningrad region               | Ppg     |
| leon                                     | Castile and Leon               | Leon                           | Ppg     |
| lesotho                                  | Africa                         | Lesotho                        | Ppg     |
| liaoning                                 | China                          | Liaoning                       | Ppg     |
| liberecky                                | Czech Republic                 | Liberec                        | Ppg     |
| liberia                                  | Africa                         | Liberia                        | Ppg     |
| libya                                    | Africa                         | Libya                          | Ppg     |
| liguria                                  | Italy                          | Liguria                        | Ppg     |
| limburg-flanders                         | Flanders                       | Limburg                        | Pp      |
| limburg-netherlands                      | Netherlands                    | Limburg                        | Pp      |
| limousin                                 | France                         | Limousine                      | Ppg     |
| lincolnshire                             | England                        | Lincolnshire                   | Ppg     |
| lipetsk_oblast                           | Central Federal District       | Lipetsk Region                 | Ppg     |
| lleida                                   | Catalonia                      | Lleida                         | Ppg     |
| lodzkie                                  | Poland                         | Lodz                           | Ppg     |
| lombardia                                | Italy                          | Lombardy                       | Ppg     |
| lorraine                                 | France                         | Lorraine                       | Ppg     |
| louisiana                                | United States of America       | Louisiana                      | Ppg     |
| lubelskie                                | Poland                         | Lubelskie                      | Ppg     |
| lubuskie                                 | Poland                         | Lubuski                        | Ppg     |
| lucerne                                  | Switzerland                    | Alfalfa                        | Ppg     |
| lugo                                     | Galicia                        | Lugo                           | Ppg     |
| luhansk_oblast                           | Ukraine                        | Lugansk region                 | Ppg     |
| luxembourg                               | Europe                         | Luxembourg                     | Ppg     |
| lviv_oblast                              | Ukraine                        | Lviv region                    | Ppg     |
| macau                                    | China                          | Macau                          | Ppg     |
| madagascar                               | Africa                         | Madagascar                     | Ppg     |
| madeira                                  | Portugal                       | Madeira                        | Ppg     |
| madhya_pradesh                           | India                          | Madhya Pradesh                 | Ppg     |
| magadan_oblast                           | Far Eastern Federal District   | Magadan Region                 | Ppg     |
| maharashtra                              | India                          | Maharashtra                    | Ppg     |
| maine                                    | United States of America       | Maine                          | Ppg     |
| malaga                                   | Andalusia                      | Malaga                         | Ppg     |
| malawi                                   | Africa                         | Malawi                         | Ppg     |
| malaysia                                 | Asia                           | Malaysia                       | Ppg     |
| maldives                                 | Asia                           | Maldives                       | Ppg     |
| mali                                     | Africa                         | Mali                           | Ppg     |
| malopolskie                              | Poland                         | Lesser Poland                  | Ppg     |
| maluku                                   | Indonesia                      | Maluku                         | Ppg     |
| manipur                                  | India                          | Manipur                        | Ppg     |
| manitoba                                 | Canada                         | Manitoba                       | Ppg     |
| maranhao                                 | Northeast                      | Maranhao                       | Ppg     |
| marche                                   | Italy                          | Marche                         | Ppg     |
| mari_el_republic                         | Volga Federal District         | Mari El Republic               | Ppg     |
| marshall-islands                         | Australia Oceania              | Marshall Islands               | Ppg     |
| martinique                               | Central America                | Martinique                     | Ppg     |
| maryland                                 | United States of America       | Maryland                       | Ppg     |
| massachusetts                            | United States of America       | Massachusetts                  | Ppg     |
| mato-grosso                              | Central Western                | Mata Grosso                    | Ppg     |
| mato-grosso-do-sul                       | Central Western                | Mato Grosso Do Sul             | Ppg     |
| mauricie                                 | Quebec                         | Morisi                         | Ppg     |
| mauritania                               | Africa                         | Mauritania                     | Ppg     |
| mauritius                                | Africa                         | Mauritius                      | Ppg     |
| mayotte                                  | Africa                         | Mayotte                        | Ppg     |
| mazowieckie                              | Poland                         | Masovian                       | Ppg     |
| mecklenburg_vorpommern                   | Germany                        | Mecklenburg Western Pomerania  | Ppg     |
| meghalaya                                | India                          | Megalaya                       | Ppg     |
| mendoza                                  | Argentina                      | Mendoza                        | Ppg     |
| merseyside                               | England                        | Merseyside                     | Ppg     |
| mexico                                   | North America                  | Mexico                         | Ppg     |
| mexico_city                              | Mexico                         | mexico city                    | Ppg     |
| michigan                                 | United States of America       | Michigan                       | Ppg     |
| michoacan                                | Mexico                         | Michoacán                      | Ppg     |
| micronesia                               | Australia Oceania              | micronesia                     | Ppg     |
| midi_pyrenees                            | France                         | Midi Pyrenees                  | Ppg     |
| minas-gerais                             | Southeast                      | Minas Gerais                   | Ppg     |
| minnesota                                | United States of America       | Minnesota                      | Ppg     |
| minskaya                                 | Belarus                        | Minsk Region                   | Ppg     |
| misiones                                 | Argentina                      | missiones                      | Ppg     |
| mississippi                              | United States of America       | Mississippi                    | Ppg     |
| missouri                                 | United States of America       | Missouri                       | Ppg     |
| mittelfranken                            | Bavaria                        | Mittelfranken                  | Ppg     |
| mizoram                                  | India                          | Mizoram                        | Ppg     |
| moere_og_romsdal                         | Norway                         | Møre og Romsdal                | Ppg     |
| mogilevskaya                             | Belarus                        | Mogilev region                 | Ppg     |
| molise                                   | Italy                          | plea                           | Ppg     |
| monaco                                   | Europe                         | Monaco                         | Ppg     |
| mongolia                                 | Asia                           | Mongolia                       | Ppg     |
| montana                                  | United States of America       | Montana                        | Ppg     |
| monteregie                               | Quebec                         | Monteregi                      | Ppg     |
| montreal                                 | Quebec                         | Montreal                       | Ppg     |
| montserrat                               | Central America                | Montserrat                     | Ppg     |
| moravskoslezsky                          | Czech Republic                 | Moravian-Slase                 | Ppg     |
| mordovia_republic                        | Volga Federal District         | The Republic of Mordovia       | Ppg     |
| morelos                                  | Mexico                         | Morelos                        | Ppg     |
| morocco                                  | Africa                         | Morocco                        | Ppg     |
| moscow                                   | Central Federal District       | Moscow                         | Ppg     |
| moscow_oblast                            | Central Federal District       | Moscow region                  | Ppg     |
| mozambique                               | Africa                         | Mozambique                     | Ppg     |
| munster                                  | North Rhine Westphalia         | Munster                        | Ppg     |
| murmansk_oblast                          | Northwestern Federal District  | Murmansk region                | Ppg     |
| myanmar                                  | Asia                           | Myanmar                        | Ppg     |
| mykolaiv_oblast                          | Ukraine                        | Mykolaiv Region                | Ppg     |
| nagaland                                 | India                          | Nagaland                       | Ppg     |
| namibia                                  | Africa                         | Namibia                        | Ppg     |
| national_capital_territory_of_delhi      | India                          | National Capital Territory     | Ppg     |
|                                          |                                | Delhi                          |         |
| nauru                                    | Australia Oceania              | Nauru                          | Ppg     |
| nayarit                                  | Mexico                         | Nayarite                       | Ppg     |
| nebraska                                 | United States of America       | Nebraska                       | Ppg     |
| nenets_autonomous_okrug                  | Northwestern Federal District  | Nenets Autonomous Okrug        | Ppg     |
| nepal                                    | Asia                           | Nepal                          | Ppg     |
| netherlands                              | Europe                         | Netherlands                    | Ppg     |
| neuchatel                                | Switzerland                    | Neuchâtel                      | Ppg     |
| neuquen                                  | Argentina                      | Neuquen                        | Ppg     |
| nevada                                   | United States of America       | Nevada                         | Ppg     |
| new-caledonia                            | Australia Oceania              | New Caledonia                  | Ppg     |
| new-hampshire                            | United States of America       | New Hampshire                  | Ppg     |
| new-jersey                               | United States of America       | New Jersey                     | Ppg     |
| new-mexico                               | United States of America       | New Mexico                     | Ppg     |
| new-york                                 | United States of America       | New York                       | Ppg     |
| new-zealand                              | Australia Oceania              | New Zealand                    | Ppg     |
| new_brunswick                            | Canada                         | New Brunswick                  | Ppg     |
| newfoundland_and_labrador                | Canada                         | Newfoundland and Labrador      | Ppg     |
| nicaragua                                | Central America                | Nicaragua                      | Ppg     |
| nidwalden                                | Switzerland                    | Nidwalden                      | Ppg     |
| niederbayern                             | Bavaria                        | Niederbayern                   | Ppg     |
| niederosterreich                         | Austria                        | Niederosterreich               | Ppg     |
| niedersachsen                            | Germany                        | Lower Saxony                   | Ppg     |
| niger                                    | Africa                         | Niger                          | Ppg     |
| nigeria                                  | Africa                         | Nigeria                        | Ppg     |
| ningxia                                  | China                          | Ningxia                        | Ppg     |
| nitriansky                               | Slovakia                       | Nitryansky                     | Ppg     |
| niue                                     | Australia Oceania              | Niue                           | Ppg     |
| nizhny_novgorod_oblast                   | Volga Federal District         | Nizhny Novgorod Region         | Ppg     |
| noord_brabant                            | Netherlands                    | North Brabant                  | Ppg     |
| noord_holland                            | Netherlands                    | North Holland                  | Ppg     |
| norcal                                   | California                     | California North               | Ppg     |
| nord_du_quebec                           | Quebec                         | North of Quebec                | Ppg     |
| nord_pas_de_calais                       | France                         | Nord - Pas de Calais           | Ppg     |
| nordland                                 | Norway                         | Nordland                       | Ppg     |
| nordrhein_westfalen                      | Germany                        | North Rhine Westphalia         | Ppg     |
| norfolk                                  | England                        | Norfolk                        | Ppg     |
| norrbotten                               | Sweden                         | Norrbotten                     | Ppg     |
| north                                    | Brazil                         | North                          | Ppg     |
| north-carolina                           | United States of America       | North Carolina                 | Ppg     |
| north-dakota                             | United States of America       | North Dakota                   | Ppg     |
| north-eastern-zone                       | India                          | North East Zone                | Ppg     |
| north_america                            |                                | North America                  | Ppg     |
| north_caucasian_federal_district         | Russia                         | North Caucasian Federal        | Ppg     |
|                                          |                                | District                       |         |
| north_east                               | England                        | Northeast                      | Ppg     |
| north_kalimantan                         | Indonesia                      | North Kalimantan               | Ppg     |
| north_karelia                            | Finland                        | North Karelia                  | Ppg     |
| north_korea                              | Asia                           | North Korea                    | Ppg     |
| north_maluku                             | Indonesia                      | North Maluku                   | Ppg     |
| north_ossetia_alania_republic            | North Caucasian Federal        | Republic of North Ossetia      | Ppg     |
|                                          | District                       | Alania                         |         |
| north_ostrobothnia                       | Finland                        | North Ostrobothnia             | Ppg     |
| north_savo                               | Finland                        | North Savo                     | Ppg     |
| north_sulawesi                           | Indonesia                      | North Sulawesi                 | Ppg     |
| north_sumatra                            | Indonesia                      | North Sumatra                  | Ppg     |
| north_west                               | England                        | Northwest                      | Ppg     |
| north_yorkshire                          | England                        | North Yorkshire                | Ppg     |
| northamptonshire                         | England                        | Northamptonshire               | Ppg     |
| northeast                                | Brazil                         | Northeast                      | Ppg     |
| northeastern_ontario                     | Ontario                        | Northeastern Ontario           | Ppg     |
| northern-zone                            | India                          | North Zone                     | Ppg     |
| northern_ireland                         | United Kingdom                 | Northern Ireland               | Ppg     |
| northumberland                           | England                        | Northumberland                 | Ppg     |
| northwest_territories                    | Canada                         | Northwest Territories          | Ppg     |
| northwestern_federal_district            | Russia                         | Northwestern Federal District  | Ppg     |
| northwestern_ontario                     | Ontario                        | Northwestern Ontario           | Ppg     |
| norway                                   | Europe                         | Norway                         | Ppg     |
| nottinghamshire                          | England                        | Nottinghamshire                | Ppg     |
| nova_scotia                              | Canada                         | Nova Scotia                    | Ppg     |
| novgorod_oblast                          | Northwestern Federal District  | Novgorod region                | Ppg     |
| novosibirsk_oblast                       | Siberian Federal District      | Novosibirsk region             | Ppg     |
| nuevo_leon                               | Mexico                         | Nuevo Leon                     | Ppg     |
| nunavut                                  | Canada                         | Nunavut                        | Ppg     |
| nusa_tenggara                            | Indonesia                      | West Nusa Tenggara             | Ppg     |
| oaxaca                                   | Mexico                         | Oaxaca                         | Ppg     |
| oberbayern                               | Bavaria                        | Oberbayern                     | Ppg     |
| oberfranken                              | Bavaria                        | Oberfranken                    | Ppg     |
| oberosterreich                           | Austria                        | Oberosterreich                 | Ppg     |
| oberpfalz                                | Bavaria                        | Oberpfalz                      | Ppg     |
| obwalden                                 | Switzerland                    | obwalden                       | Ppg     |
| odessa_oblast                            | Ukraine                        | Odessa region                  | Ppg     |
| odisha                                   | India                          | Odisha                         | Ppg     |
| oestfold                                 | Norway                         | estfold                        | Ppg     |
| ohio                                     | United States of America       | Ohio                           | Ppg     |
| oklahoma                                 | United States of America       | Oklahoma                       | Ppg     |
| olomoucky                                | Czech Republic                 | Olomouc                        | Ppg     |
| oman                                     | Asia                           | Oman                           | Ppg     |
| omsk_oblast                              | Siberian Federal District      | Omsk Region                    | Ppg     |
| ontario                                  | Canada                         | Ontario                        | Ppg     |
| opolskie                                 | Poland                         | Opole                          | Ppg     |
| oppland                                  | Norway                         | Oppland                        | Ppg     |
| orebro                                   | Sweden                         | Orebro                         | Ppg     |
| oregon                                   | United States of America       | Oregon                         | Ppg     |
| orenburg_oblast                          | Volga Federal District         | Orenburg region                | Ppg     |
| oryol_oblast                             | Central Federal District       | Oryol Region                   | Ppg     |
| oslo                                     | Norway                         | Oslo                           | Ppg     |
| ostergotland                             | Sweden                         | Ostergotland                   | Ppg     |
| ostrobothnia                             | Finland                        | Ostrobothnia                   | Ppg     |
| ourense                                  | Galicia                        | Ourense                        | Ppg     |
| outaouais                                | Quebec                         | Outaouais                      | Ppg     |
| overijssel                               | Netherlands                    | Overijssel                     | Ppg     |
| oxfordshire                              | England                        | Oxfordshire                    | Ppg     |
| paijat_hame                              | Finland                        | Payat-Hame                     | Ppg     |
| pakistan                                 | Asia                           | Pakistan                       | Ppg     |
| palau                                    | Australia Oceania              | Palau                          | Ppg     |
| palencia                                 | Castile and Leon               | Palencia                       | Ppg     |
| palestine                                | Asia                           | Palestine                      | Ppg     |
| panama                                   | Central America                | Panama                         | Ppg     |
| papua                                    | Indonesia                      | Papua                          | Ppg     |
| papua-new-guinea                         | Australia Oceania              | Papua New Guinea               | Ppg     |
| para                                     | North                          | Paragraph                      | Ppg     |
| paraguay                                 | South America                  | Paraguay                       | Ppg     |
| paraiba                                  | Northeast                      | Paraiba                        | Ppg     |
| parana                                   | South                          | Paraná                         | Ppg     |
| pardubicky                               | Czech Republic                 | Pardubice                      | Ppg     |
| pays_de_la_loire                         | France                         | Pays de la Loire               | Ppg     |
| pennsylvania                             | United States of America       | Pennsylvania                   | Ppg     |
| penza_oblast                             | Volga Federal District         | Penza region                   | Ppg     |
| perm_krai                                | Volga Federal District         | Perm region                    | Ppg     |
| pernambuco                               | Northeast                      | Pernambuco                     | Ppg     |
| peru                                     | South America                  | Peru                           | Ppg     |
| philippines                              | Asia                           | Philippines                    | Ppg     |
| piaui                                    | Northeast                      | Piaui                          | Ppg     |
| picardie                                 | France                         | Picardy                        | Ppg     |
| piemonte                                 | Italy                          | Piedmont                       | Ppg     |
| pirkanmaa                                | Finland                        | Pirkanmaa                      | Ppg     |
| pitcairn-islands                         | Australia Oceania              | Pitcairn Islands               | Ppg     |
| plzensky                                 | Czech Republic                 | Pilsensky                      | Ppg     |
| podkarpackie                             | Poland                         | Subcarpathia                   | Ppg     |
| podlaskie                                | Poland                         | Podlaskie Voivodeship          | Ppg     |
| poitou_charentes                         | France                         | Poitou Charente                | Ppg     |
| poland                                   | Europe                         | Poland                         | Ppg     |
| poltava_oblast                           | Ukraine                        | Poltava region                 | Ppg     |
| polynesie-francaise                      | Australia Oceania              | Polynesia-France               | Ppg     |
| pomorskie                                | Poland                         | Pomeranian                     | Ppg     |
| pontevedra                               | Galicia                        | Pontevedra                     | Ppg     |
| portugal                                 | Europe                         | Portugal                       | Ppg     |
| praha                                    | Czech Republic                 | Prague                         | Ppg     |
| presovsky                                | Slovakia                       | Presovsky                      | Ppg     |
| primorsky_krai                           | Far Eastern Federal District   | Primorsky Krai                 | Ppg     |
| prince_edward_island                     | Canada                         | Prince Edward Island           | Ppg     |
| provence_alpes_cote_d_azur               | France                         | Provence - Alpes - Cote d'Azur | Ppg     |
| pskov_oblast                             | Northwestern Federal District  | Pskov region                   | Ppg     |
| puducherry                               | India                          | Puducherry                     | Ppg     |
| puebla                                   | Mexico                         | puebla                         | Ppg     |
| puerto-rico                              | United States of America       | Puerto Rico                    | Ppg     |
| puerto_rico                              | Central America                | Puerto Rico                    | Ppg     |
| puglia                                   | Italy                          | Apulia                         | Ppg     |
| punjab                                   | India                          | Punjab                         | Ppg     |
| qatar                                    | Asia                           | Qatar                          | Ppg     |
| qinghai                                  | China                          | Qinghai                        | Ppg     |
| quebec                                   | Canada                         | Quebec                         | Ppg     |
| queretaro                                | Mexico                         | Querétaro                      | Ppg     |
| quintana_roo                             | Mexico                         | Quintana Roo                   | Ppg     |
| rajasthan                                | India                          | Rajasthan                      | Ppg     |
| region_de_murcia                         | Spain                          | Murcia                         | Ppg     |
| reunion                                  | Africa                         | reunion                        | Ppg     |
| rheinland_pfalz                          | Germany                        | Rhineland Palatinate           | Ppg     |
| rhode-island                             | United States of America       | Rhode Island                   | Ppg     |
| rhone_alpes                              | France                         | Rhone Alps                     | Ppg     |
| riau                                     | Indonesia                      | Riau                           | Ppg     |
| riau_islands                             | Indonesia                      | Kepulauan Riau                 | Ppg     |
| rio-de-janeiro                           | Southeast                      | Rio de Janeiro                 | Ppg     |
| rio-grande-do-norte                      | Northeast                      | Rio Grande do Norte            | Ppg     |
| rio-grande-do-sul                        | South                          | Rio Grande do Sul              | Ppg     |
| rio_negro                                | Argentina                      | Rio Negro                      | Ppg     |
| rivne_oblast                             | Ukraine                        | Rivne Region                   | Ppg     |
| rogaland                                 | Norway                         | Rogaland                       | Ppg     |
| rondonia                                 | North                          | Rondonia                       | Ppg     |
| roraima                                  | North                          | Roraima                        | Ppg     |
| rostov_oblast                            | Southern Federal District      | Rostov region                  | Ppg     |
| russia                                   |                                | Russia                         | Ppg     |
| rutland                                  | England                        | rutland                        | Ppg     |
| rwanda                                   | Africa                         | Rwanda                         | Ppg     |
| ryazan_oblast                            | Central Federal District       | Ryazan Oblast                  | Ppg     |
| saarland                                 | Germany                        | Saar                           | Ppg     |
| sachsen                                  | Germany                        | Sachsen                        | Ppg     |
| sachsen_anhalt                           | Germany                        | Saxony-Anhalt                  | Ppg     |
| saguenay_lac_saint_jean                  | Quebec                         | Saguenay - Lake Saint-Jean     | Ppg     |
| saint_barthelemy                         | Central America                | Saint Barthélemy               | Ppg     |
| saint_gallen                             | Switzerland                    | St. Gallen                     | Ppg     |
| saint_helena_ascension_tristan_da_cunha  | Africa                         | Saint Helena                   | Ppg     |
| saint_kitts_and_nevis                    | Central America                | Saint Kitts and Nevis          | Ppg     |
| saint_lucia                              | Central America                | Saint Lucia                    | Ppg     |
| saint_martin                             | Central America                | Saint Martin                   | Ppg     |
| saint_petersburg                         | Northwestern Federal District  | St. Petersburg                 | Ppg     |
| saint_pierre_et_miquelon                 | North America                  | Saint Pierre and Miquelon      | Ppg     |
| saint_vincent_and_the_grenadines         | Central America                | Saint Vincent and the          | Ppg     |
|                                          |                                | Grenadines                     |         |
| sakha_republic                           | Far Eastern Federal District   | Republic of Yakutia            | Ppg     |
| sakhalin_oblast                          | Far Eastern Federal District   | Sakhalin Region                | Ppg     |
| salamanca                                | Castile and Leon               | Salamanca                      | Ppg     |
| salta                                    | Argentina                      | Salta                          | Ppg     |
| salzburg                                 | Austria                        | Salzburg                       | Ppg     |
| samara_oblast                            | Volga Federal District         | Samara Region                  | Ppg     |
| samoa                                    | Australia Oceania              | Samoa                          | Ppg     |
| san_juan                                 | Argentina                      | San Juan                       | Ppg     |
| san_luis                                 | Argentina                      | San Luis                       | Ppg     |
| san_luis_potosi                          | Mexico                         | San Luis Potosi                | Ppg     |
| san_marino                               | Europe                         | San Marino                     | Ppg     |
| santa-catarina                           | South                          | Santa Catarina                 | Ppg     |
| santa_cruz                               | Argentina                      | Santa Cruz                     | Ppg     |
| santa_fe                                 | Argentina                      | Santa Fe                       | Ppg     |
| santiago_del_estero                      | Argentina                      | Santiago Del Estero            | Ppg     |
| sao-paulo                                | Southeast                      | Sao Paulo                      | Ppg     |
| sao_tome_and_principe                    | Africa                         | Sao Tome and Principe          | Ppg     |
| saratov_oblast                           | Volga Federal District         | Saratov region                 | Ppg     |
| sardegna                                 | Italy                          | Sardinia                       | Ppg     |
| saskatchewan                             | Canada                         | Saskatchewan                   | Ppg     |
| satakunta                                | Finland                        | Satakunta                      | Ppg     |
| saudi_arabia                             | Asia                           | Saudi Arabia                   | Ppg     |
| schaffhausen                             | Switzerland                    | Schaffhausen                   | Ppg     |
| schleswig_holstein                       | Germany                        | Schleswig Holstein             | Ppg     |
| schwaben                                 | Bavaria                        | Schwaben                       | Ppg     |
| schwyz                                   | Switzerland                    | Schwyz                         | Ppg     |
| scotland                                 | United Kingdom                 | Scotland                       | Ppg     |
| segovia                                  | Castile and Leon               | Segovia                        | Ppg     |
| senegal                                  | Africa                         | Senegal                        | Ppg     |
| sergipe                                  | Northeast                      | Sergeep                        | Ppg     |
| sevastopol                               | Southern Federal District      | Sevastopol                     | Ppg     |
| sevastopol-ukraine                       | Ukraine                        | Sevastopol                     | Pp      |
| sevilla                                  | Andalusia                      | Seville                        | Ppg     |
| seychelles                               | Africa                         | Seychelles                     | Ppg     |
| shaanxi                                  | China                          | Shaanxi                        | Ppg     |
| shandong                                 | China                          | Shandong                       | Ppg     |
| shanghai                                 | China                          | Shanghai                       | Ppg     |
| shanxi                                   | China                          | Shanxi                         | Ppg     |
| shikoku                                  | Japan                          | Shikoku                        | Ppg     |
| shropshire                               | England                        | Shropshire                     | Ppg     |
| siberian_federal_district                | Russia                         | Siberian Federal District      | Ppg     |
| sichuan                                  | China                          | Sichuan                        | Ppg     |
| sicilia                                  | Italy                          | Sicily                         | Ppg     |
| sierra-leone                             | Africa                         | Sierra Leone                   | Ppg     |
| sikkim                                   | India                          | Sikkim                         | Ppg     |
| sinaloa                                  | Mexico                         | Sinaloa                        | Ppg     |
| singapore                                | Asia                           | Singapore                      | Ppg     |
| sint_maarten                             | Central America                | Sint Maarten                   | Ppg     |
| skane                                    | Sweden                         | Skane                          | Ppg     |
| slaskie                                  | Poland                         | Sweet                          | Ppg     |
| slovakia                                 | Europe                         | Slovakia                       | Ppg     |
| smolensk_oblast                          | Central Federal District       | Smolensk region                | Ppg     |
| socal                                    | California                     | California South               | Ppg     |
| sodermanland                             | Sweden                         | Södermanland                   | Ppg     |
| sogn_og_fjordane                         | Norway                         | Sogn-og-Fyurane                | Ppg     |
| solomon-islands                          | Australia Oceania              | Solomon islands                | Ppg     |
| solothurn                                | Switzerland                    | Solothurn                      | Ppg     |
| somalia                                  | Africa                         | Somalia                        | Ppg     |
| somerset                                 | England                        | Somerset                       | Ppg     |
| sonora                                   | Mexico                         | Sonora                         | Ppg     |
| soria                                    | Castile and Leon               | Soria                          | Ppg     |
| south                                    | Brazil                         | South                          | Ppg     |
| south-africa-and-lesotho                 | Africa                         | South Africa and Lesotho       | Ppg     |
| south-carolina                           | United States of America       | South Carolina                 | Ppg     |
| south-dakota                             | United States of America       | North Dakota                   | Ppg     |
| south_africa                             | Africa                         | South Africa                   | Ppg     |
| south_america                            |                                | South America                  | Ppg     |
| south_east                               | England                        | South East                     | Ppg     |
| south_georgia_and_south_sandwich         | South America                  | South Georgia And South        | Ppg     |
|                                          |                                | Sandwich                       |         |
| south_kalimantan                         | Indonesia                      | South Kalimantan               | Ppg     |
| south_karelia                            | Finland                        | South Karelia                  | Ppg     |
| south_korea                              | Asia                           | South Korea                    | Ppg     |
| south_ostrobothnia                       | Finland                        | Southern Ostrobothnia          | Ppg     |
| south_savo                               | Finland                        | South Savo                     | Ppg     |
| south_sudan                              | Africa                         | South Sudan                    | Ppg     |
| south_sulawesi                           | Indonesia                      | South Sulawesi                 | Ppg     |
| south_sumatra                            | Indonesia                      | South Sumatra                  | Ppg     |
| south_west                               | England                        | South West                     | Ppg     |
| south_yorkshire                          | England                        | South Yorkshire                | Ppg     |
| southeast                                | Brazil                         | Southeast                      | Ppg     |
| southeast_sulawesi                       | Indonesia                      | Southeast Sulawesi             | Ppg     |
| southern-zone                            | India                          | South Zone                     | Ppg     |
| southern_federal_district                | Russia                         | Southern Federal District      | Ppg     |
| southwest_finland                        | Finland                        | Southwest Finland              | Ppg     |
| southwestern_ontario                     | Ontario                        | Southwestern Ontario           | Ppg     |
| spain                                    | Europe                         | Spain                          | Ppg     |
| sri_lanka                                | Asia                           | Sri Lanka                      | Ppg     |
| staffordshire                            | England                        | Staffordshire                  | Ppg     |
| state_of_mexico                          | Mexico                         | mexico city                    | Ppg     |
| stavropol_krai                           | North Caucasian Federal        | Stavropol region               | Ppg     |
|                                          | District                       |                                |         |
| steiermark                               | Austria                        | Stirmarka                      | Ppg     |
| stockholm                                | Sweden                         | Stockholm                      | Ppg     |
| stredocesky                              | Czech Republic                 | Stredochsky                    | Ppg     |
| sudan                                    | Africa                         | Sudan                          | Ppg     |
| suffolk                                  | England                        | Suffolk                        | Ppg     |
| sulawesi                                 | Indonesia                      | Sulawesi                       | Ppg     |
| sumatra                                  | Indonesia                      | Sumatra                        | Ppg     |
| sumy_oblast                              | Ukraine                        | Sumy region                    | Ppg     |
| suriname                                 | South America                  | Suriname                       | Ppg     |
| surrey                                   | England                        | Surrey                         | Ppg     |
| svalbard                                 | Norway                         | Svalbard                       | Ppg     |
| sverdlovsk_oblast                        | Ural federal district          | Sverdlovsk region              | Ppg     |
| swaziland                                | Africa                         | Swaziland                      | Ppg     |
| sweden                                   | Europe                         | Sweden                         | Ppg     |
| swietokrzyskie                           | Poland                         | Świętokrzyskie                 | Ppg     |
| switzerland                              | Europe                         | Switzerland                    | Ppg     |
| syria                                    | Asia                           | Syria                          | Ppg     |
| tabasco                                  | Mexico                         | Tabasco                        | Ppg     |
| taiwan                                   | Asia                           | Taiwan                         | Ppg     |
| tajikistan                               | Asia                           | Tajikistan                     | Ppg     |
| tamaulipas                               | Mexico                         | Tamaulipas                     | Ppg     |
| tambov_oblast                            | Central Federal District       | Tambov Region                  | Ppg     |
| tamil_nadu                               | India                          | Tamilnadu                      | Ppg     |
| tanzania                                 | Africa                         | Tanzania                       | Ppg     |
| tarragona                                | Catalonia                      | Tarragona                      | Ppg     |
| tatarstan_republic                       | Volga Federal District         | Republic of Tatarstan          | Ppg     |
| telangana                                | India                          | Telengana                      | Ppg     |
| telemark                                 | Norway                         | Telemark                       | Ppg     |
| tennessee                                | United States of America       | Tennessee                      | Ppg     |
| ternopil_oblast                          | Ukraine                        | Ternopil Region                | Ppg     |
| teruel                                   | Aragon                         | Teruel                         | Ppg     |
| texas                                    | United States of America       | Texas                          | Ppg     |
| thailand                                 | Asia                           | Thailand                       | Ppg     |
| thueringen                               | Germany                        | Thuringian                     | Ppg     |
| thurgau                                  | Switzerland                    | Thurgau                        | Ppg     |
| tianjin                                  | China                          | Tianjin                        | Ppg     |
| tibet                                    | China                          | Tibet                          | Ppg     |
| ticino                                   | Switzerland                    | Ticino                         | Ppg     |
| tierra_del_fuego                         | Argentina                      | Tierra del Fuego               | Ppg     |
| tirol                                    | Austria                        | Tyrol                          | Ppg     |
| tlaxcala                                 | Mexico                         | Tlaxcala                       | Ppg     |
| tocantins                                | North                          | Tocantines                     | Ppg     |
| togo                                     | Africa                         | Togo                           | Ppg     |
| tohoku                                   | Japan                          | Tohoku                         | Ppg     |
| tokelau                                  | Australia Oceania              | Tokelau                        | Ppg     |
| toledo                                   | Castile-La Mancha              | Toledo                         | Ppg     |
| tomsk_oblast                             | Siberian Federal District      | Tomsk Region                   | Ppg     |
| tonga                                    | Australia Oceania              | Tonga                          | Ppg     |
| toscana                                  | Italy                          | Tuscany                        | Ppg     |
| trenciansky                              | Slovakia                       | Trenchansky                    | Ppg     |
| trentino_alto_adige                      | Italy                          | Trentino Alto Adige            | Ppg     |
| trinidad_and_tobago                      | Central America                | Trinidad and Tobago            | Ppg     |
| tripura                                  | India                          | Tripura                        | Ppg     |
| trnavsky                                 | Slovakia                       | Trnavsky                       | Ppg     |
| troendelag                               | Norway                         | Trendelag                      | Ppg     |
| troms                                    | Norway                         | Troms                          | Ppg     |
| tucuman                                  | Argentina                      | Tucuman                        | Ppg     |
| tula_oblast                              | Central Federal District       | Tula region                    | Ppg     |
| tunisia                                  | Africa                         | Tunisia                        | Ppg     |
| turkmenistan                             | Asia                           | Turkmenistan                   | Ppg     |
| turks_and_caicos_islands                 | Central America                | Turks and Caicos Islands       | Ppg     |
| tuva_republic                            | Siberian Federal District      | Tyva Republic                  | Ppg     |
| tuvalu                                   | Australia Oceania              | Tuvalu                         | Ppg     |
| tver_oblast                              | Central Federal District       | Tver region                    | Ppg     |
| tyne_and_wear                            | England                        | Tyne And Wear                  | Ppg     |
| tyumen_oblast                            | Ural federal district          | Tyumen region                  | Ppg     |
| udmurt_republic                          | Volga Federal District         | Republic of Udmurtia           | Ppg     |
| uganda                                   | Africa                         | Uganda                         | Ppg     |
| ukraine                                  | Europe                         | Ukraine                        | Ppg     |
| ulyanovsk_oblast                         | Volga Federal District         | Ulyanovsk region               | Ppg     |
| umbria                                   | Italy                          | Umbria                         | Ppg     |
| united_arab_emirates                     | Asia                           | United Arab Emirates           | Ppg     |
| united_kingdom                           | Europe                         | United Kingdom                 | Ppg     |
| united_states_virgin_islands             | Central America                | Virgin Islands                 | Ppg     |
| unterfranken                             | Bavaria                        | uncontracted                   | Ppg     |
| uppsala                                  | Sweden                         | Uppsala                        | Ppg     |
| ural_federal_district                    | Russia                         | Ural federal district          | Ppg     |
| uri                                      | Switzerland                    | Uri                            | Ppg     |
| uruguay                                  | South America                  | Uruguay                        | Ppg     |
| us                                       | North America                  | United States of America       | Ppg     |
| us-midwest                               | North America                  | US-Midwest                     | Ppg     |
| us-northeast                             | North America                  | USA-Northeast                  | Ppg     |
| us-pacific                               | North America                  | USA-Pacific                    | Ppg     |
| us-south                                 | North America                  | US South                       | Ppg     |
| us-virgin-islands                        | United States of America       | US Virgin Islands              | Ppg     |
| us-west                                  | North America                  | USA-West                       | Ppg     |
| usa_virgin_islands                       | Central America                | Virgin Islands                 | Ppg     |
| ustecky                                  | Czech Republic                 | Ustetsky                       | Ppg     |
| utah                                     | United States of America       | Utah                           | Ppg     |
| utrecht                                  | Netherlands                    | Utrecht                        | Ppg     |
| uttar_pradesh                            | India                          | Uttar Pradesh                  | Ppg     |
| uttarakhand                              | India                          | Uttarakhand                    | Ppg     |
| uusimaa                                  | Finland                        | Uusimaa                        | Ppg     |
| uzbekistan                               | Asia                           | Uzbekistan                     | Ppg     |
| valais                                   | Switzerland                    | Valais                         | Ppg     |
| valencia                                 | Valencia                       | Valencia                       | Ppg     |
| valladolid                               | Castile and Leon               | Valladolid                     | Ppg     |
| valle_aosta                              | Italy                          | Valle Aosta                    | Ppg     |
| vanuatu                                  | Australia Oceania              | Vanuatu                        | Ppg     |
| varmland                                 | Sweden                         | Värmland                       | Ppg     |
| vasterbotten                             | Sweden                         | Vasterbotten                   | Ppg     |
| vasternorrland                           | Sweden                         | Vast Severnaya Zemlya          | Ppg     |
| vastmanland                              | Sweden                         | Vast Earth                     | Ppg     |
| vastra_gotaland                          | Sweden                         | Västra Götaland                | Ppg     |
| vatican_city                             | Europe                         | Vatican                        | Ppg     |
| vaud                                     | Switzerland                    | In                             | Ppg     |
| veneto                                   | Italy                          | Veneto                         | Ppg     |
| venezuela                                | South America                  | Venezuela                      | Ppg     |
| veracruz                                 | Mexico                         | Veracruz                       | Ppg     |
| vermont                                  | United States of America       | Vermont                        | Ppg     |
| vest-agder                               | Norway                         | Vest-Agder                     | Ppg     |
| vestfold                                 | Norway                         | Vest                           | Ppg     |
| vietnam                                  | Asia                           | Vietnam                        | Ppg     |
| vinnytsia_oblast                         | Ukraine                        | Vinnytsia region               | Ppg     |
| virginia                                 | United States of America       | Virginia                       | Ppg     |
| vitebskaya                               | Belarus                        | Vitebsk region                 | Ppg     |
| vizcaya                                  | Euskadi                        | Biscay                         | Ppg     |
| vladimir_oblast                          | Central Federal District       | Vladimir region                | Ppg     |
| volga_federal_district                   | Russia                         | Volga Federal District         | Ppg     |
| volgograd_oblast                         | Southern Federal District      | Volgograd region               | Ppg     |
| vologda_oblast                           | Northwestern Federal District  | Vologda Region                 | Ppg     |
| volyn_oblast                             | Ukraine                        | Volyn Region                   | Ppg     |
| vorarlberg                               | Austria                        | Vorarlberg                     | Ppg     |
| voronezh_oblast                          | Central Federal District       | Voronezh region                | Ppg     |
| vysocina                                 | Czech Republic                 | Vysochina                      | Ppg     |
| wales                                    | United Kingdom                 | Wales                          | Ppg     |
| wallis-et-futuna                         | Australia Oceania              | Wallis and Futuna              | Ppg     |
| wallonia_french_community                | Belgium                        | French Community of Belgium    | Ppg     |
| wallonia_german_community                | Belgium                        | wallonia                       | Ppg     |
| warminsko_mazurskie                      | Poland                         | Warmian-Masurian               | Ppg     |
| warwickshire                             | England                        | Warwickshire                   | Ppg     |
| washington                               | United States of America       | Washington                     | Ppg     |
| west-virginia                            | United States of America       | West Virginia                  | Ppg     |
| west_bengal                              | India                          | West Bengal                    | Ppg     |
| west_flanders                            | Flanders                       | West Flanders                  | Ppg     |
| west_java                                | Indonesia                      | West Java                      | Ppg     |
| west_kalimantan                          | Indonesia                      | West Kalimantan                | Ppg     |
| west_midlands                            | England                        | Western Middle-earth           | Ppg     |
| west_nusa_tenggara                       | Indonesia                      | West Nusa Tenggara             | Ppg     |
| west_papua                               | Indonesia                      | West Papua                     | Ppg     |
| west_sulawesi                            | Indonesia                      | West Sulawesi                  | Ppg     |
| west_sumatra                             | Indonesia                      | West Sumatra                   | Ppg     |
| west_sussex                              | England                        | West Sussex                    | Ppg     |
| west_yorkshire                           | England                        | West Yorkshire                 | Ppg     |
| western-zone                             | India                          | West Zone                      | Ppg     |
| western_sahara                           | Africa                         | West Sahara                    | Ppg     |
| wielkopolskie                            | Poland                         | Greater Poland                 | Ppg     |
| wien                                     | Austria                        | Vein                           | Ppg     |
| wiltshire                                | England                        | Wiltshire                      | Ppg     |
| wisconsin                                | United States of America       | Wisconsin                      | Ppg     |
| worcestershire                           | England                        | Worcestershire                 | Ppg     |
| wyoming                                  | United States of America       | Wyoming                        | Ppg     |
| xinjiang                                 | China                          | xinjiang                       | Ppg     |
| yamalo_nenets_autonomous_okrug           | Ural federal district          | Yamalo Nenets Autonomous Okrug | Ppg     |
| yaroslavl_oblast                         | Central Federal District       | Yaroslavskaya oblast           | Ppg     |
| yemen                                    | Asia                           | Yemen                          | Ppg     |
| yogyakarta                               | Indonesia                      | Yogyakarta                     | Ppg     |
| yorkshire_and_the_humber                 | England                        | Yorkshire and the Humber       | Ppg     |
| yucatan                                  | Mexico                         | Yucatan                        | Ppg     |
| yukon                                    | Canada                         | Yukon                          | Ppg     |
| yunnan                                   | China                          | Yunnan                         | Ppg     |
| zabaykalsky_krai                         | Siberian Federal District      | Transbaikal region             | Ppg     |
| zacatecas                                | Mexico                         | Zacatecas                      | Ppg     |
| zachodniopomorskie                       | Poland                         | West Pomeranian                | Ppg     |
| zakarpattia_oblast                       | Ukraine                        | Zakarpattia Oblast             | Ppg     |
| zambia                                   | Africa                         | Zambia                         | Ppg     |
| zamora                                   | Castile and Leon               | Zamora                         | Ppg     |
| zaporizhia_oblast                        | Ukraine                        | Zaporozhye region              | Ppg     |
| zaragoza                                 | Aragon                         | Zaragoza                       | Ppg     |
| zeeland                                  | Netherlands                    | Zealand                        | Ppg     |
| zhejiang                                 | China                          | Zhejiang                       | Ppg     |
| zhytomyr_oblast                          | Ukraine                        | Zhytomyr Oblast                | Ppg     |
| zilinsky                                 | Slovakia                       | Zhilinsky                      | Ppg     |
| zimbabwe                                 | Africa                         | Zimbabwe                       | Ppg     |
| zlinsky                                  | Czech Republic                 | Zlinsky                        | Ppg     |
| zug                                      | Switzerland                    | Zug                            | Ppg     |
| zuid_holland                             | Netherlands                    | South Holland                  | Ppg     |
| zurich                                   | Switzerland                    | Zurich                         | Ppg     |
Total elements: 1003

