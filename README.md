# download-geofabrik

[![Build Status](https://travis-ci.org/julien-noblet/download-geofabrik.svg?branch=master)](https://travis-ci.org/julien-noblet/download-geofabrik)
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
                             "geofabrik", "openstreetmap.fr" or "bbbike". It
                             automatically change config file if -c is unused.
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
| centro-oeste                                | Brazil                      | centro-oeste                   | sPBHpSk |
| ceuta                                       | Spain                       | Ceuta                          | sPBHpSk |
| chad                                        | Africa                      | Chad                           | sPBHpSk |
| champagne-ardenne                           | France                      | Champagne Ardenne              | sPBHpSk |
| cheshire                                    | England                     | Cheshire                       | sPBHpSk |
| chile                                       | South America               | Chile                          | sPBHpSk |
| china                                       | Asia                        | China                          | sPBHpSk |
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
| england                                     | Great Britain               | England                        | sPBHpSk |
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
| gabon                                       | Africa                      | Gabon                          | sPBHpSk |
| galicia                                     | Spain                       | Galicia                        | sPBHpSk |
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
| guatemala                                   | Central America             | Guatemala                      | sPBHpSk |
| guernsey-jersey                             | Europe                      | guernsey-jersey                | sPBHpSk |
| guinea                                      | Africa                      | Guinea                         | sPBHpSk |
| guinea-bissau                               | Africa                      | Guinea-Bissau                  | sPBHpSk |
| guyana                                      | South America               | guyana                         | sPBHpSk |
| guyane                                      | France                      | Guyane                         | sPBHpk  |
| haiti-and-domrep                            | Central America             | Haiti and Dominican Republic   | sPBHpSk |
| hamburg                                     | Germany                     | Hamburg                        | sPBHpSk |
| hampshire                                   | England                     | Hampshire                      | sPBHpSk |
| haute-normandie                             | France                      | Haute-Normandie                | sPBHpSk |
| herefordshire                               | England                     | Herefordshire                  | sPBHpSk |
| hertfordshire                               | England                     | Hertfordshire                  | sPBHpSk |
| hessen                                      | Germany                     | Hessen                         | sPBHpSk |
| hokkaido                                    | Japan                       | Hokkaidō                       | sPBHpSk |
| honduras                                    | Central America             | Honduras                       | sPBHpSk |
| hungary                                     | Europe                      | Hungary                        | sPBHpSk |
| iceland                                     | Europe                      | Iceland                        | sPBHpSk |
| ile-de-clipperton                           | Australia and Oceania       | Île de Clipperton              | sPBHpSk |
| ile-de-france                               | France                      | Ile-de-France                  | sPBHpSk |
| india                                       | Asia                        | India                          | sPBHpk  |
| indonesia                                   | Asia                        | Indonesia (with East Timor)    | sPBHpk  |
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
| niue                                        | Australia and Oceania       | Niue                           | sPBHpSk |
| noord-brabant                               | Netherlands                 | Noord-Brabant                  | sPBHpSk |
| noord-holland                               | Netherlands                 | Noord-Holland                  | sPBHpSk |
| norcal                                      | us/california               | Northern California            | sPBHpSk |
| nord-est                                    | Italy                       | Nord-Est                       | sPBHpSk |
| nord-ovest                                  | Italy                       | Nord-Ovest                     | sPBHpSk |
| nord-pas-de-calais                          | France                      | Nord-Pas-de-Calais             | sPBHpSk |
| nordeste                                    | Brazil                      | nordeste                       | sPBHpSk |
| nordrhein-westfalen                         | Germany                     | Nordrhein-Westfalen            | sPBHpk  |
| norfolk                                     | England                     | Norfolk                        | sPBHpSk |
| norte                                       | Brazil                      | norte                          | sPBHpSk |
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
| polynesie-francaise                         | Australia and Oceania       | polynesie-francaise            | sPBHpSk |
| pomorskie                                   | Poland                      | Województwo pomorskie<br       | sPBHpSk |
|                                             |                             | />(Pomeranian Voivodeship)     |         |
| portugal                                    | Europe                      | Portugal                       | sPBHpSk |
| prince-edward-island                        | Canada                      | Prince Edward Island           | sPBHpSk |
| provence-alpes-cote-d-azur                  | France                      | Provence Alpes-Cote-d'Azur     | sPBHpSk |
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
| scotland                                    | Great Britain               | Scotland                       | sPBHpSk |
| senegal-and-gambia                          | Africa                      | Senegal and Gambia             | sPBHpSk |
| serbia                                      | Europe                      | Serbia                         | sPBHpSk |
| seychelles                                  | Africa                      | Seychelles                     | sPBHpSk |
| shikoku                                     | Japan                       | Shikoku                        | sPBHpSk |
| shropshire                                  | England                     | Shropshire                     | sPBHpSk |
| siberian-fed-district                       | Russian Federation          | Siberian Federal District      | sPBHpSk |
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
| sudeste                                     | Brazil                      | sudeste                        | sPBHpSk |
| suffolk                                     | England                     | Suffolk                        | sPBHpSk |
| sul                                         | Brazil                      | sul                            | sPBHpSk |
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
| wales                                       | Great Britain               | Wales                          | sPBHpSk |
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
| yemen                                       | Asia                        | Yemen                          | sPBHpSk |
| yukon                                       | Canada                      | Yukon                          | sPBHpSk |
| zachodniopomorskie                          | Poland                      | Województwo                    | sPBHpSk |
|                                             |                             | zachodniopomorskie<br />(West  |         |
|                                             |                             | Pomeranian Voivodeship)        |         |
| zambia                                      | Africa                      | Zambia                         | sPBHpSk |
| zeeland                                     | Netherlands                 | Zeeland                        | sPBHpSk |
| zimbabwe                                    | Africa                      | Zimbabwe                       | sPBHpSk |
| zuid-holland                                | Netherlands                 | Zuid-Holland                   | sPBHpSk |
Total elements: 475

## List of elements from openstreetmap.fr
| SHORTNAME | IS IN | LONG NAME | FORMATS |
|-----------|-------|-----------|---------|
|           |       |           | sPp     |
Total elements: 1

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
Total elements: 236

