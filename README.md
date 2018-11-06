# download-geofabrik

[![Build Status](https://travis-ci.org/julien-noblet/download-geofabrik.svg?branch=master)](https://travis-ci.org/julien-noblet/download-geofabrik)
[![Join the chat at https://gitter.im/julien-noblet/download-geofabrik](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/julien-noblet/download-geofabrik?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/julien-noblet/download-geofabrik)](https://goreportcard.com/report/github.com/julien-noblet/download-geofabrik)
[![Coverage Status](https://coveralls.io/repos/github/julien-noblet/download-geofabrik/badge.svg?branch=master)](https://coveralls.io/github/julien-noblet/download-geofabrik?branch=master)

## Version 2
Warning! command line have changed from V1
see [Usage](#usage)

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
      --help                 Show context-sensitive help (also try --help-long
                             and --help-man).
      --service="geofabrik"  Can switch to another service. You can use
                             "geofabrik", "openstreetmap.fr" or "gislab". It
                             automatically change config file if -c is unused.
  -c, --config="./geofabrik.yml"  
                             Set Config file.
  -n, --nodownload           Do not download file (test only)
  -v, --verbose              Be verbose
  -q, --quiet                Be quiet
      --progress             Add a progress bar
      --proxy-http=""        Use http proxy, format: proxy_address:port
      --proxy-sock5=""       Use Sock5 proxy, format: proxy_address:port
      --proxy-user=""        Proxy user
      --proxy-pass=""        Proxy password
      --version              Show application version.

Commands:
  help [<command>...]
    Show help.


  update [<flags>]
    Update geofabrik.yml from github *** DEPRECATED you should prefer use
    generate ***

    --url="https://raw.githubusercontent.com/julien-noblet/download-geofabrik/master/geofabrik.yml"  
      Url for config source

  list [<flags>]
    Show elements available

    --markdown  generate list in Markdown format

  download [<flags>] <element>
    Download element

    -B, --osm.bz2  Download osm.bz2 if available
    -S, --shp.zip  Download shp.zip if available
    -P, --osm.pbf  Download osm.pbf (default)
    -H, --osh.pbf  Download osh.pbf
    -s, --state    Download state.txt file
    -p, --poly     Download poly file
    -k, --kml      Download kml file
        --check    Control with checksum (default) Use --no-check to discard
                   control

  generate
    Generate a new config file



```

## List of elements
|                  SHORTNAME                  |          IS IN           |               LONG NAME                | FORMATS |
|---------------------------------------------|--------------------------|----------------------------------------|---------|
| afghanistan                                 | Asia                     | Afghanistan                            | sPBpSk  |
| africa                                      |                          | Africa                                 | sPBpk   |
| alabama                                     | United States of America | Alabama                                | sPBpSk  |
| alaska                                      | United States of America | Alaska                                 | sPBpSk  |
| albania                                     | Europe                   | Albania                                | sPBpSk  |
| alberta                                     | Canada                   | Alberta                                | sPBpSk  |
| algeria                                     | Africa                   | Algeria                                | sPBpSk  |
| alps                                        | Europe                   | Alps                                   | sPBpk   |
| alsace                                      | France                   | Alsace                                 | sPBpSk  |
| andorra                                     | Europe                   | Andorra                                | sPBpSk  |
| angola                                      | Africa                   | Angola                                 | sPBpSk  |
| antarctica                                  |                          | Antarctica                             | sPBpSk  |
| aquitaine                                   | France                   | Aquitaine                              | sPBpSk  |
| argentina                                   | South America            | Argentina                              | sPBpSk  |
| arizona                                     | United States of America | Arizona                                | sPBpSk  |
| arkansas                                    | United States of America | Arkansas                               | sPBpSk  |
| armenia                                     | Asia                     | Armenia                                | sPBpSk  |
| arnsberg-regbez                             | Nordrhein-Westfalen      | Regierungsbezirk Arnsberg              | sPBpSk  |
| asia                                        |                          | Asia                                   | sPBpk   |
| australia                                   | Australia and Oceania    | Australia                              | sPBpSk  |
| australia-oceania                           |                          | Australia and Oceania                  | sPBpk   |
| austria                                     | Europe                   | Austria                                | sPBpSk  |
| auvergne                                    | France                   | Auvergne                               | sPBpSk  |
| azerbaijan                                  | Asia                     | Azerbaijan                             | sPBpSk  |
| azores                                      | Europe                   | Azores                                 | sPBpSk  |
| baden-wuerttemberg                          | Germany                  | Baden-Württemberg                      | sPBpSk  |
| bangladesh                                  | Asia                     | Bangladesh                             | sPBpSk  |
| basse-normandie                             | France                   | Basse-Normandie                        | sPBpSk  |
| bayern                                      | Germany                  | Bayern                                 | sPBpSk  |
| bedfordshire                                | England                  | Bedfordshire                           | sPBpSk  |
| belarus                                     | Europe                   | Belarus                                | sPBpSk  |
| belgium                                     | Europe                   | Belgium                                | sPBpSk  |
| belize                                      | Central America          | Belize                                 | sPBpSk  |
| benin                                       | Africa                   | Benin                                  | sPBpSk  |
| berkshire                                   | England                  | Berkshire                              | sPBpSk  |
| berlin                                      | Germany                  | Berlin                                 | sPBpSk  |
| bhutan                                      | Asia                     | Bhutan                                 | sPBpSk  |
| bolivia                                     | South America            | Bolivia                                | sPBpSk  |
| bosnia-herzegovina                          | Europe                   | Bosnia-Herzegovina                     | sPBpSk  |
| botswana                                    | Africa                   | Botswana                               | sPBpSk  |
| bourgogne                                   | France                   | Bourgogne                              | sPBpSk  |
| brandenburg                                 | Germany                  | Brandenburg (mit Berlin)               | sPBpSk  |
| brazil                                      | South America            | Brazil                                 | sPBpSk  |
| bremen                                      | Germany                  | Bremen                                 | sPBpSk  |
| bretagne                                    | France                   | Bretagne                               | sPBpSk  |
| bristol                                     | England                  | Bristol                                | sPBpSk  |
| british-columbia                            | Canada                   | British Columbia                       | sPBpSk  |
| british-isles                               | Europe                   | British Isles                          | sPBpk   |
| buckinghamshire                             | England                  | Buckinghamshire                        | sPBpSk  |
| bulgaria                                    | Europe                   | Bulgaria                               | sPBpSk  |
| burkina-faso                                | Africa                   | Burkina Faso                           | sPBpSk  |
| burundi                                     | Africa                   | Burundi                                | sPBpSk  |
| california                                  | United States of America | California                             | sPBpSk  |
| cambodia                                    | Asia                     | Cambodia                               | sPBpSk  |
| cambridgeshire                              | England                  | Cambridgeshire                         | sPBpSk  |
| cameroon                                    | Africa                   | Cameroon                               | sPBpSk  |
| canada                                      | North America            | Canada                                 | sPBpk   |
| canary-islands                              | Africa                   | Canary Islands                         | sPBpSk  |
| cape-verde                                  | Africa                   | Cape Verde                             | sPBpSk  |
| central-african-republic                    | Africa                   | Central African Republic               | sPBpSk  |
| central-america                             |                          | Central America                        | sPBpk   |
| central-fed-district                        | Russian Federation       | Central Federal District               | sPBpSk  |
| centre                                      | France                   | Centre                                 | sPBpSk  |
| centro                                      | Italy                    | Centro                                 | sPBpSk  |
| chad                                        | Africa                   | Chad                                   | sPBpSk  |
| champagne-ardenne                           | France                   | Champagne Ardenne                      | sPBpSk  |
| cheshire                                    | England                  | Cheshire                               | sPBpSk  |
| chile                                       | South America            | Chile                                  | sPBpSk  |
| china                                       | Asia                     | China                                  | sPBpSk  |
| chubu                                       | Japan                    | Chūbu region                           | sPBpSk  |
| chugoku                                     | Japan                    | Chūgoku region                         | sPBpSk  |
| colombia                                    | South America            | Colombia                               | sPBpSk  |
| colorado                                    | United States of America | Colorado                               | sPBpSk  |
| comores                                     | Africa                   | Comores                                | sPBpSk  |
| congo-brazzaville                           | Africa                   | Congo (Republic/Brazzaville)           | sPBpSk  |
| congo-democratic-republic                   | Africa                   | Congo (Democratic                      | sPBpSk  |
|                                             |                          | Republic/Kinshasa)                     |         |
| connecticut                                 | United States of America | Connecticut                            | sPBpSk  |
| cornwall                                    | England                  | Cornwall                               | sPBpSk  |
| corse                                       | France                   | Corse                                  | sPBpSk  |
| crimean-fed-district                        | Russian Federation       | Crimean Federal District               | sPBpSk  |
| croatia                                     | Europe                   | Croatia                                | sPBpSk  |
| cuba                                        | Central America          | Cuba                                   | sPBpSk  |
| cumbria                                     | England                  | Cumbria                                | sPBpSk  |
| cyprus                                      | Europe                   | Cyprus                                 | sPBpSk  |
| czech-republic                              | Europe                   | Czech Republic                         | sPBpSk  |
| dach                                        | Europe                   | Germany, Austria, Switzerland          | sPBpk   |
| delaware                                    | United States of America | Delaware                               | sPBpSk  |
| denmark                                     | Europe                   | Denmark                                | sPBpSk  |
| derbyshire                                  | England                  | Derbyshire                             | sPBpSk  |
| detmold-regbez                              | Nordrhein-Westfalen      | Regierungsbezirk Detmold               | sPBpSk  |
| devon                                       | England                  | Devon                                  | sPBpSk  |
| district-of-columbia                        | United States of America | District of Columbia                   | sPBpSk  |
| djibouti                                    | Africa                   | Djibouti                               | sPBpSk  |
| dolnoslaskie                                | Poland                   | Województwo dolnośląskie(Lower         | sPBpSk  |
|                                             |                          | Silesian Voivodeship)                  |         |
| dorset                                      | England                  | Dorset                                 | sPBpSk  |
| duesseldorf-regbez                          | Nordrhein-Westfalen      | Regierungsbezirk Düsseldorf            | sPBpSk  |
| durham                                      | England                  | Durham                                 | sPBpSk  |
| east-sussex                                 | England                  | East Sussex                            | sPBpSk  |
| east-yorkshire-with-hull                    | England                  | East Yorkshire with Hull               | sPBpSk  |
| ecuador                                     | South America            | Ecuador                                | sPBpk   |
| egypt                                       | Africa                   | Egypt                                  | sPBpSk  |
| england                                     | Great Britain            | England                                | sPBpSk  |
| equatorial-guinea                           | Africa                   | Equatorial Guinea                      | sPBpSk  |
| eritrea                                     | Africa                   | Eritrea                                | sPBpSk  |
| essex                                       | England                  | Essex                                  | sPBpSk  |
| estonia                                     | Europe                   | Estonia                                | sPBpSk  |
| ethiopia                                    | Africa                   | Ethiopia                               | sPBpSk  |
| europe                                      |                          | Europe                                 | sPBpk   |
| far-eastern-fed-district                    | Russian Federation       | Far Eastern Federal District           | sPBpSk  |
| faroe-islands                               | Europe                   | Faroe Islands                          | sPBpSk  |
| fiji                                        | Australia and Oceania    | Fiji                                   | sPBpSk  |
| finland                                     | Europe                   | Finland                                | sPBpSk  |
| florida                                     | United States of America | Florida                                | sPBpSk  |
| france                                      | Europe                   | France                                 | sPBpk   |
| franche-comte                               | France                   | Franche Comte                          | sPBpSk  |
| freiburg-regbez                             | Baden-Württemberg        | Regierungsbezirk Freiburg              | sPBpSk  |
| gabon                                       | Africa                   | Gabon                                  | sPBpSk  |
| gcc-states                                  | Asia                     | GCC States                             | sPBpSk  |
| georgia-eu                                  | Europe                   | Georgia (Europe country)               | sPBpSk  |
| georgia-us                                  | United States of America | Georgia (US State)                     | sPBpSk  |
| germany                                     | Europe                   | Germany                                | sPBpk   |
| ghana                                       | Africa                   | Ghana                                  | sPBpSk  |
| gloucestershire                             | England                  | Gloucestershire                        | sPBpSk  |
| great-britain                               | Europe                   | Great Britain                          | sPBpk   |
| greater-london                              | England                  | Greater London                         | sPBpSk  |
| greater-manchester                          | England                  | Greater Manchester                     | sPBpSk  |
| greece                                      | Europe                   | Greece                                 | sPBpSk  |
| greenland                                   | North America            | Greenland                              | sPBpSk  |
| guadeloupe                                  | France                   | Guadeloupe                             | sPBpk   |
| guatemala                                   | Central America          | Guatemala                              | sPBpSk  |
| guinea                                      | Africa                   | Guinea                                 | sPBpSk  |
| guinea-bissau                               | Africa                   | Guinea-Bissau                          | sPBpSk  |
| guyane                                      | France                   | Guyane                                 | sPBpk   |
| haiti-and-domrep                            | Central America          | Haiti and Dominican Republic           | sPBpSk  |
| hamburg                                     | Germany                  | Hamburg                                | sPBpSk  |
| hampshire                                   | England                  | Hampshire                              | sPBpSk  |
| haute-normandie                             | France                   | Haute-Normandie                        | sPBpSk  |
| hawaii                                      | United States of America | Hawaii                                 | sPBpSk  |
| herefordshire                               | England                  | Herefordshire                          | sPBpSk  |
| hertfordshire                               | England                  | Hertfordshire                          | sPBpSk  |
| hessen                                      | Germany                  | Hessen                                 | sPBpSk  |
| hokkaido                                    | Japan                    | Hokkaidō                               | sPBpSk  |
| hungary                                     | Europe                   | Hungary                                | sPBpSk  |
| iceland                                     | Europe                   | Iceland                                | sPBpSk  |
| idaho                                       | United States of America | Idaho                                  | sPBpSk  |
| ile-de-france                               | France                   | Ile-de-France                          | sPBpSk  |
| illinois                                    | United States of America | Illinois                               | sPBpSk  |
| india                                       | Asia                     | India                                  | sPBpSk  |
| indiana                                     | United States of America | Indiana                                | sPBpSk  |
| indonesia                                   | Asia                     | Indonesia                              | sPBpSk  |
| iowa                                        | United States of America | Iowa                                   | sPBpSk  |
| iran                                        | Asia                     | Iran                                   | sPBpSk  |
| iraq                                        | Asia                     | Iraq                                   | sPBpSk  |
| ireland-and-northern-ireland                | Europe                   | Ireland and Northern Ireland           | sPBpSk  |
| isle-of-man                                 | Europe                   | Isle of Man                            | sPBpSk  |
| isle-of-wight                               | England                  | Isle of Wight                          | sPBpSk  |
| isole                                       | Italy                    | Isole                                  | sPBpSk  |
| israel-and-palestine                        | Asia                     | Israel and Palestine                   | sPBpSk  |
| italy                                       | Europe                   | Italy                                  | sPBpk   |
| ivory-coast                                 | Africa                   | Ivory Coast                            | sPBpSk  |
| jamaica                                     | Central America          | jamaica                                | sPBpSk  |
| japan                                       | Asia                     | Japan                                  | sPBpk   |
| jordan                                      | Asia                     | Jordan                                 | sPBpSk  |
| kaliningrad                                 | Russian Federation       | Kaliningrad                            | sPBpSk  |
| kansai                                      | Japan                    | Kansai region (a.k.a. Kinki            | sPBpSk  |
|                                             |                          | region)                                |         |
| kansas                                      | United States of America | Kansas                                 | sPBpSk  |
| kanto                                       | Japan                    | Kantō region                           | sPBpSk  |
| karlsruhe-regbez                            | Baden-Württemberg        | Regierungsbezirk Karlsruhe             | sPBpSk  |
| kazakhstan                                  | Asia                     | Kazakhstan                             | sPBpSk  |
| kent                                        | England                  | Kent                                   | sPBpSk  |
| kentucky                                    | United States of America | Kentucky                               | sPBpSk  |
| kenya                                       | Africa                   | Kenya                                  | sPBpSk  |
| koeln-regbez                                | Nordrhein-Westfalen      | Regierungsbezirk Köln                  | sPBpSk  |
| kosovo                                      | Europe                   | Kosovo                                 | sPBpSk  |
| kujawsko-pomorskie                          | Poland                   | Województwo                            | sPBpSk  |
|                                             |                          | kujawsko-pomorskie(Kuyavian-Pomeranian |         |
|                                             |                          | Voivodeship)                           |         |
| kyrgyzstan                                  | Asia                     | Kyrgyzstan                             | sPBpSk  |
| kyushu                                      | Japan                    | Kyūshū                                 | sPBpSk  |
| lancashire                                  | England                  | Lancashire                             | sPBpSk  |
| languedoc-roussillon                        | France                   | Languedoc-Roussillon                   | sPBpSk  |
| laos                                        | Asia                     | Laos                                   | sPBpSk  |
| latvia                                      | Europe                   | Latvia                                 | sPBpSk  |
| lebanon                                     | Asia                     | Lebanon                                | sPBpSk  |
| leicestershire                              | England                  | Leicestershire                         | sPBpSk  |
| lesotho                                     | Africa                   | Lesotho                                | sPBpSk  |
| liberia                                     | Africa                   | Liberia                                | sPBpSk  |
| libya                                       | Africa                   | Libya                                  | sPBpSk  |
| liechtenstein                               | Europe                   | Liechtenstein                          | sPBpSk  |
| limousin                                    | France                   | Limousin                               | sPBpSk  |
| lincolnshire                                | England                  | Lincolnshire                           | sPBpSk  |
| lithuania                                   | Europe                   | Lithuania                              | sPBpSk  |
| lodzkie                                     | Poland                   | Województwo łódzkie(Łódź               | sPBpSk  |
|                                             |                          | Voivodeship)                           |         |
| lorraine                                    | France                   | Lorraine                               | sPBpSk  |
| louisiana                                   | United States of America | Louisiana                              | sPBpSk  |
| lubelskie                                   | Poland                   | Województwo lubelskie(Lublin           | sPBpSk  |
|                                             |                          | Voivodeship)                           |         |
| lubuskie                                    | Poland                   | Województwo lubuskie(Lubusz            | sPBpSk  |
|                                             |                          | Voivodeship)                           |         |
| luxembourg                                  | Europe                   | Luxembourg                             | sPBpSk  |
| macedonia                                   | Europe                   | Macedonia                              | sPBpSk  |
| madagascar                                  | Africa                   | Madagascar                             | sPBpSk  |
| maine                                       | United States of America | Maine                                  | sPBpSk  |
| malawi                                      | Africa                   | Malawi                                 | sPBpSk  |
| malaysia-singapore-brunei                   | Asia                     | Malaysia, Singapore, and               | sPBpk   |
|                                             |                          | Brunei                                 |         |
| maldives                                    | Asia                     | Maldives                               | sPBpSk  |
| mali                                        | Africa                   | Mali                                   | sPBpSk  |
| malopolskie                                 | Poland                   | Województwo małopolskie(Lesser         | sPBpSk  |
|                                             |                          | Poland Voivodeship)                    |         |
| malta                                       | Europe                   | Malta                                  | sPBpSk  |
| manitoba                                    | Canada                   | Manitoba                               | sPBpSk  |
| martinique                                  | France                   | Martinique                             | sPBpk   |
| maryland                                    | United States of America | Maryland                               | sPBpSk  |
| massachusetts                               | United States of America | Massachusetts                          | sPBpSk  |
| mauritania                                  | Africa                   | Mauritania                             | sPBpSk  |
| mauritius                                   | Africa                   | Mauritius                              | sPBpSk  |
| mayotte                                     | France                   | Mayotte                                | sPBpk   |
| mazowieckie                                 | Poland                   | Województwo                            | sPBpSk  |
|                                             |                          | mazowieckie(Mazovian                   |         |
|                                             |                          | Voivodeship)                           |         |
| mecklenburg-vorpommern                      | Germany                  | Mecklenburg-Vorpommern                 | sPBpSk  |
| merseyside                                  | England                  | Merseyside                             | sPBpSk  |
| mexico                                      | North America            | Mexico                                 | sPBpSk  |
| michigan                                    | United States of America | Michigan                               | sPBpSk  |
| midi-pyrenees                               | France                   | Midi-Pyrenees                          | sPBpSk  |
| minnesota                                   | United States of America | Minnesota                              | sPBpSk  |
| mississippi                                 | United States of America | Mississippi                            | sPBpSk  |
| missouri                                    | United States of America | Missouri                               | sPBpSk  |
| mittelfranken                               | Bayern                   | Mittelfranken                          | sPBpSk  |
| moldova                                     | Europe                   | Moldova                                | sPBpSk  |
| monaco                                      | Europe                   | Monaco                                 | sPBpSk  |
| mongolia                                    | Asia                     | Mongolia                               | sPBpSk  |
| montana                                     | United States of America | Montana                                | sPBpSk  |
| montenegro                                  | Europe                   | Montenegro                             | sPBpSk  |
| morocco                                     | Africa                   | Morocco                                | sPBpSk  |
| mozambique                                  | Africa                   | Mozambique                             | sPBpSk  |
| muenster-regbez                             | Nordrhein-Westfalen      | Regierungsbezirk Münster               | sPBpSk  |
| myanmar                                     | Asia                     | Myanmar (a.k.a. Burma)                 | sPBpSk  |
| namibia                                     | Africa                   | Namibia                                | sPBpSk  |
| nebraska                                    | United States of America | Nebraska                               | sPBpSk  |
| nepal                                       | Asia                     | Nepal                                  | sPBpSk  |
| netherlands                                 | Europe                   | Netherlands                            | sPBpSk  |
| nevada                                      | United States of America | Nevada                                 | sPBpSk  |
| new-brunswick                               | Canada                   | New Brunswick                          | sPBpSk  |
| new-caledonia                               | Australia and Oceania    | New Caledonia                          | sPBpSk  |
| new-hampshire                               | United States of America | New Hampshire                          | sPBpSk  |
| new-jersey                                  | United States of America | New Jersey                             | sPBpSk  |
| new-mexico                                  | United States of America | New Mexico                             | sPBpSk  |
| new-york                                    | United States of America | New York                               | sPBpSk  |
| new-zealand                                 | Australia and Oceania    | New Zealand                            | sPBpSk  |
| newfoundland-and-labrador                   | Canada                   | Newfoundland and Labrador              | sPBpSk  |
| nicaragua                                   | Central America          | Nicaragua                              | sPBpk   |
| niederbayern                                | Bayern                   | Niederbayern                           | sPBpSk  |
| niedersachsen                               | Germany                  | Niedersachsen                          | sPBpSk  |
| niger                                       | Africa                   | Niger                                  | sPBpSk  |
| nigeria                                     | Africa                   | Nigeria                                | sPBpSk  |
| nord-est                                    | Italy                    | Nord-Est                               | sPBpSk  |
| nord-ovest                                  | Italy                    | Nord-Ovest                             | sPBpSk  |
| nord-pas-de-calais                          | France                   | Nord-Pas-de-Calais                     | sPBpSk  |
| nordrhein-westfalen                         | Germany                  | Nordrhein-Westfalen                    | sPBpSk  |
| norfolk                                     | England                  | Norfolk                                | sPBpSk  |
| north-america                               |                          | North America                          | sPBpk   |
| north-carolina                              | United States of America | North Carolina                         | sPBpSk  |
| north-caucasus-fed-district                 | Russian Federation       | North Caucasus Federal                 | sPBpSk  |
|                                             |                          | District                               |         |
| north-dakota                                | United States of America | North Dakota                           | sPBpSk  |
| north-korea                                 | Asia                     | North Korea                            | sPBpSk  |
| north-yorkshire                             | England                  | North Yorkshire                        | sPBpSk  |
| northamptonshire                            | England                  | Northamptonshire                       | sPBpSk  |
| northumberland                              | England                  | Northumberland                         | sPBpSk  |
| northwest-territories                       | Canada                   | Northwest Territories                  | sPBpSk  |
| northwestern-fed-district                   | Russian Federation       | Northwestern Federal District          | sPBpSk  |
| norway                                      | Europe                   | Norway                                 | sPBpSk  |
| nottinghamshire                             | England                  | Nottinghamshire                        | sPBpSk  |
| nova-scotia                                 | Canada                   | Nova Scotia                            | sPBpSk  |
| nunavut                                     | Canada                   | Nunavut                                | sPBpSk  |
| oberbayern                                  | Bayern                   | Oberbayern                             | sPBpSk  |
| oberfranken                                 | Bayern                   | Oberfranken                            | sPBpSk  |
| oberpfalz                                   | Bayern                   | Oberpfalz                              | sPBpSk  |
| ohio                                        | United States of America | Ohio                                   | sPBpSk  |
| oklahoma                                    | United States of America | Oklahoma                               | sPBpSk  |
| ontario                                     | Canada                   | Ontario                                | sPBpSk  |
| opolskie                                    | Poland                   | Województwo opolskie(Opole             | sPBpSk  |
|                                             |                          | Voivodeship)                           |         |
| oregon                                      | United States of America | Oregon                                 | sPBpSk  |
| oxfordshire                                 | England                  | Oxfordshire                            | sPBpSk  |
| pakistan                                    | Asia                     | Pakistan                               | sPBpSk  |
| papua-new-guinea                            | Australia and Oceania    | Papua New Guinea                       | sPBpSk  |
| paraguay                                    | South America            | Paraguay                               | sPBpSk  |
| pays-de-la-loire                            | France                   | Pays de la Loire                       | sPBpSk  |
| pennsylvania                                | United States of America | Pennsylvania                           | sPBpSk  |
| peru                                        | South America            | Peru                                   | sPBpSk  |
| philippines                                 | Asia                     | Philippines                            | sPBpSk  |
| picardie                                    | France                   | Picardie                               | sPBpSk  |
| podkarpackie                                | Poland                   | Województwo                            | sPBpSk  |
|                                             |                          | podkarpackie(Subcarpathian             |         |
|                                             |                          | Voivodeship)                           |         |
| podlaskie                                   | Poland                   | Województwo                            | sPBpSk  |
|                                             |                          | podlaskie(Podlaskie                    |         |
|                                             |                          | Voivodeship)                           |         |
| poitou-charentes                            | France                   | Poitou-Charentes                       | sPBpSk  |
| poland                                      | Europe                   | Poland                                 | sPBpk   |
| pomorskie                                   | Poland                   | Województwo                            | sPBpSk  |
|                                             |                          | pomorskie(Pomeranian                   |         |
|                                             |                          | Voivodeship)                           |         |
| portugal                                    | Europe                   | Portugal                               | sPBpSk  |
| prince-edward-island                        | Canada                   | Prince Edward Island                   | sPBpSk  |
| provence-alpes-cote-d-azur                  | France                   | Provence Alpes-Cote-d'Azur             | sPBpSk  |
| puerto-rico                                 | United States of America | Puerto Rico                            | sPBpSk  |
| quebec                                      | Canada                   | Quebec                                 | sPBpSk  |
| reunion                                     | France                   | Reunion                                | sPBpk   |
| rheinland-pfalz                             | Germany                  | Rheinland-Pfalz                        | sPBpSk  |
| rhode-island                                | United States of America | Rhode Island                           | sPBpSk  |
| rhone-alpes                                 | France                   | Rhone-Alpes                            | sPBpSk  |
| romania                                     | Europe                   | Romania                                | sPBpSk  |
| russia                                      |                          | Russian Federation                     | sPBpk   |
| rutland                                     | England                  | Rutland                                | sPBpSk  |
| rwanda                                      | Africa                   | Rwanda                                 | sPBpSk  |
| saarland                                    | Germany                  | Saarland                               | sPBpSk  |
| sachsen                                     | Germany                  | Sachsen                                | sPBpSk  |
| sachsen-anhalt                              | Germany                  | Sachsen-Anhalt                         | sPBpSk  |
| saint-helena-ascension-and-tristan-da-cunha | Africa                   | Saint Helena, Ascension, and           | sPBpSk  |
|                                             |                          | Tristan da Cunha                       |         |
| sao-tome-and-principe                       | Africa                   | Sao Tome and Principe                  | sPBpSk  |
| saskatchewan                                | Canada                   | Saskatchewan                           | sPBpSk  |
| schleswig-holstein                          | Germany                  | Schleswig-Holstein                     | sPBpSk  |
| schwaben                                    | Bayern                   | Schwaben                               | sPBpSk  |
| scotland                                    | Great Britain            | Scotland                               | sPBpSk  |
| senegal-and-gambia                          | Africa                   | Senegal and Gambia                     | sPBpSk  |
| serbia                                      | Europe                   | Serbia                                 | sPBpSk  |
| seychelles                                  | Africa                   | Seychelles                             | sPBpSk  |
| shikoku                                     | Japan                    | Shikoku                                | sPBpSk  |
| shropshire                                  | England                  | Shropshire                             | sPBpSk  |
| siberian-fed-district                       | Russian Federation       | Siberian Federal District              | sPBpSk  |
| sierra-leone                                | Africa                   | Sierra Leone                           | sPBpSk  |
| slaskie                                     | Poland                   | Województwo śląskie(Silesian           | sPBpSk  |
|                                             |                          | Voivodeship)                           |         |
| slovakia                                    | Europe                   | Slovakia                               | sPBpSk  |
| slovenia                                    | Europe                   | Slovenia                               | sPBpSk  |
| somalia                                     | Africa                   | Somalia                                | sPBpSk  |
| somerset                                    | England                  | Somerset                               | sPBpSk  |
| south-africa                                | Africa                   | South Africa                           | sPBpSk  |
| south-africa-and-lesotho                    | Africa                   | South Africa (includes                 | sPBpk   |
|                                             |                          | Lesotho)                               |         |
| south-america                               |                          | South America                          | sPBpk   |
| south-carolina                              | United States of America | South Carolina                         | sPBpSk  |
| south-dakota                                | United States of America | South Dakota                           | sPBpSk  |
| south-fed-district                          | Russian Federation       | South Federal District                 | sPBpSk  |
| south-korea                                 | Asia                     | South Korea                            | sPBpSk  |
| south-sudan                                 | Africa                   | South Sudan                            | sPBpSk  |
| south-yorkshire                             | England                  | South Yorkshire                        | sPBpSk  |
| spain                                       | Europe                   | Spain                                  | sPBpSk  |
| sri-lanka                                   | Asia                     | Sri Lanka                              | sPBpSk  |
| staffordshire                               | England                  | Staffordshire                          | sPBpSk  |
| stuttgart-regbez                            | Baden-Württemberg        | Regierungsbezirk Stuttgart             | sPBpSk  |
| sud                                         | Italy                    | Sud                                    | sPBpSk  |
| sudan                                       | Africa                   | Sudan                                  | sPBpSk  |
| suffolk                                     | England                  | Suffolk                                | sPBpSk  |
| suriname                                    | South America            | suriname                               | sPBpSk  |
| surrey                                      | England                  | Surrey                                 | sPBpSk  |
| swaziland                                   | Africa                   | Swaziland                              | sPBpSk  |
| sweden                                      | Europe                   | Sweden                                 | sPBpSk  |
| swietokrzyskie                              | Poland                   | Województwo                            | sPBpSk  |
|                                             |                          | świętokrzyskie(Świętokrzyskie          |         |
|                                             |                          | Voivodeship)                           |         |
| switzerland                                 | Europe                   | Switzerland                            | sPBpSk  |
| syria                                       | Asia                     | Syria                                  | sPBpSk  |
| taiwan                                      | Asia                     | Taiwan                                 | sPBpSk  |
| tajikistan                                  | Asia                     | Tajikistan                             | sPBpk   |
| tanzania                                    | Africa                   | Tanzania                               | sPBpSk  |
| tennessee                                   | United States of America | Tennessee                              | sPBpSk  |
| texas                                       | United States of America | Texas                                  | sPBpSk  |
| thailand                                    | Asia                     | Thailand                               | sPBpSk  |
| thueringen                                  | Germany                  | Thüringen                              | sPBpSk  |
| togo                                        | Africa                   | Togo                                   | sPBpSk  |
| tohoku                                      | Japan                    | Tōhoku region                          | sPBpSk  |
| tuebingen-regbez                            | Baden-Württemberg        | Regierungsbezirk Tübingen              | sPBpSk  |
| tunisia                                     | Africa                   | Tunisia                                | sPBpSk  |
| turkey                                      | Europe                   | Turkey                                 | sPBpSk  |
| turkmenistan                                | Asia                     | Turkmenistan                           | sPBpSk  |
| tyne-and-wear                               | England                  | Tyne and Wear                          | sPBpSk  |
| uganda                                      | Africa                   | Uganda                                 | sPBpSk  |
| ukraine                                     | Europe                   | Ukraine                                | sPBpSk  |
| unterfranken                                | Bayern                   | Unterfranken                           | sPBpSk  |
| ural-fed-district                           | Russian Federation       | Ural Federal District                  | sPBpSk  |
| uruguay                                     | South America            | Uruguay                                | sPBpSk  |
| us                                          | North America            | United States of America               |         |
| us-midwest                                  | North America            | US Midwest                             | sPBpk   |
| us-northeast                                | North America            | US Northeast                           | sPBpk   |
| us-pacific                                  | North America            | US Pacific                             | sPBpk   |
| us-south                                    | North America            | US South                               | sPBpk   |
| us-west                                     | North America            | US West                                | sPBpk   |
| utah                                        | United States of America | Utah                                   | sPBpSk  |
| uzbekistan                                  | Asia                     | Uzbekistan                             | sPBpSk  |
| venezuela                                   | South America            | Venezuela                              | sPBpSk  |
| vermont                                     | United States of America | Vermont                                | sPBpSk  |
| vietnam                                     | Asia                     | Vietnam                                | sPBpSk  |
| virginia                                    | United States of America | Virginia                               | sPBpSk  |
| volga-fed-district                          | Russian Federation       | Volga Federal District                 | sPBpSk  |
| wales                                       | Great Britain            | Wales                                  | sPBpSk  |
| warminsko-mazurskie                         | Poland                   | Województwo                            | sPBpSk  |
|                                             |                          | warmińsko-mazurskie(Warmian-Masurian   |         |
|                                             |                          | Voivodeship)                           |         |
| warwickshire                                | England                  | Warwickshire                           | sPBpSk  |
| washington                                  | United States of America | Washington                             | sPBpSk  |
| west-midlands                               | England                  | West Midlands                          | sPBpSk  |
| west-sussex                                 | England                  | West Sussex                            | sPBpSk  |
| west-virginia                               | United States of America | West Virginia                          | sPBpSk  |
| west-yorkshire                              | England                  | West Yorkshire                         | sPBpSk  |
| wielkopolskie                               | Poland                   | Województwo                            | sPBpSk  |
|                                             |                          | wielkopolskie(Greater Poland           |         |
|                                             |                          | Voivodeship)                           |         |
| wiltshire                                   | England                  | Wiltshire                              | sPBpSk  |
| wisconsin                                   | United States of America | Wisconsin                              | sPBpSk  |
| worcestershire                              | England                  | Worcestershire                         | sPBpSk  |
| wyoming                                     | United States of America | Wyoming                                | sPBpSk  |
| yemen                                       | Asia                     | Yemen                                  | sPBpSk  |
| yukon                                       | Canada                   | Yukon                                  | sPBpSk  |
| zachodniopomorskie                          | Poland                   | Województwo                            | sPBpSk  |
|                                             |                          | zachodniopomorskie(West                |         |
|                                             |                          | Pomeranian Voivodeship)                |         |
| zambia                                      | Africa                   | Zambia                                 | sPBpSk  |
| zimbabwe                                    | Africa                   | Zimbabwe                               | sPBpSk  |
Total elements: 397

## List of elements from openstreetmap.fr
|                SHORTNAME                |              IS IN               |                LONG NAME                | FORMATS |
|-----------------------------------------|----------------------------------|-----------------------------------------|---------|
| abruzzo                                 | italy                            | abruzzo                                 | sPp     |
| aceh                                    | indonesia                        | aceh                                    | sPp     |
| adygea_republic                         | southern_federal_district        | adygea_republic                         | sPp     |
| afghanistan                             | asia                             | afghanistan                             | sPp     |
| africa                                  |                                  | africa                                  | sPp     |
| algeria                                 | africa                           | algeria                                 | sPp     |
| altai_krai                              | siberian_federal_district        | altai_krai                              | sPp     |
| altai_republic                          | siberian_federal_district        | altai_republic                          | sPp     |
| american_samoa                          | south-america                    | american_samoa                          | sPp     |
| amur_oblast                             | far_eastern_federal_district     | amur_oblast                             | sPp     |
| andalucia                               | spain                            | andalucia                               | sPp     |
| andaman_and_nicobar_islands             | india                            | andaman_and_nicobar_islands             | sPp     |
| andhra_pradesh                          | india                            | andhra_pradesh                          | sPp     |
| angola                                  | africa                           | angola                                  | sPp     |
| anguilla                                | central-america                  | anguilla                                | sPp     |
| anhui                                   | china                            | anhui                                   | sPp     |
| antigua_and_barbuda                     | central-america                  | antigua_and_barbuda                     | sPp     |
| aragon                                  | spain                            | aragon                                  | sPp     |
| arkhangelsk_oblast                      | northwestern_federal_district    | arkhangelsk_oblast                      | sPp     |
| armenia                                 | asia                             | armenia                                 | sPp     |
| arnsberg                                | nordrhein_westfalen              | arnsberg                                | sPp     |
| aruba                                   | central-america                  | aruba                                   | sPp     |
| arunachal_pradesh                       | india                            | arunachal_pradesh                       | sPp     |
| asia                                    |                                  | asia                                    | sPp     |
| assam                                   | india                            | assam                                   | sPp     |
| astrakhan_oblast                        | southern_federal_district        | astrakhan_oblast                        | sPp     |
| asturias                                | spain                            | asturias                                | sPp     |
| australia                               | oceania                          | australia                               | sPp     |
| australian_capital_territory            | australia                        | australian_capital_territory            | sPp     |
| austria                                 | europe                           | austria                                 | sPp     |
| bahamas                                 | central-america                  | bahamas                                 | sPp     |
| bahrain                                 | asia                             | bahrain                                 | sPp     |
| bali                                    | indonesia                        | bali                                    | sPp     |
| bangka_belitung_islands                 | indonesia                        | bangka_belitung_islands                 | sPp     |
| banskobystricky                         | slovakia                         | banskobystricky                         | sPp     |
| banten                                  | indonesia                        | banten                                  | sPp     |
| barbados                                | central-america                  | barbados                                | sPp     |
| bashkortostan_republic                  | volga_federal_district           | bashkortostan_republic                  | sPp     |
| basilicata                              | italy                            | basilicata                              | sPp     |
| beijing                                 | china                            | beijing                                 | sPp     |
| belgium                                 | europe                           | belgium                                 | sPp     |
| belgorod_oblast                         | central_federal_district         | belgorod_oblast                         | sPp     |
| bengkulu                                | indonesia                        | bengkulu                                | sPp     |
| benin                                   | africa                           | benin                                   | sPp     |
| bermuda                                 | north-america                    | bermuda                                 | sPp     |
| bhutan                                  | asia                             | bhutan                                  | sPp     |
| bihar                                   | india                            | bihar                                   | sPp     |
| bouvet_island                           | africa                           | bouvet_island                           | sPp     |
| bratislavsky                            | slovakia                         | bratislavsky                            | sPp     |
| british_indian_ocean_territory          | asia                             | british_indian_ocean_territory          | sPp     |
| british_virgin_islands                  | central-america                  | british_virgin_islands                  | sPp     |
| brunei                                  | asia                             | brunei                                  | sPp     |
| brussels_capital_region                 | belgium                          | brussels_capital_region                 | sPp     |
| bryansk_oblast                          | central_federal_district         | bryansk_oblast                          | sPp     |
| burgenland                              | austria                          | burgenland                              | sPp     |
| burkina_faso                            | africa                           | burkina_faso                            | sPp     |
| burundi                                 | africa                           | burundi                                 | sPp     |
| buryatia_republic                       | siberian_federal_district        | buryatia_republic                       | sPp     |
| calabria                                | italy                            | calabria                                | sPp     |
| cambodia                                | asia                             | cambodia                                | sPp     |
| cameroon                                | africa                           | cameroon                                | sPp     |
| campania                                | italy                            | campania                                | sPp     |
| canada                                  | north-america                    | canada                                  | p       |
| canarias                                | spain                            | canarias                                | sPp     |
| cantabria                               | spain                            | cantabria                               | sPp     |
| cape_verde                              | africa                           | cape_verde                              | sPp     |
| caribbean                               | central-america                  | caribbean                               | sPp     |
| castilla_la_mancha                      | spain                            | castilla_la_mancha                      | sPp     |
| castilla_y_leon                         | spain                            | castilla_y_leon                         | sPp     |
| catalunya                               | spain                            | catalunya                               | sPp     |
| cayman_islands                          | central-america                  | cayman_islands                          | sPp     |
| central-america                         |                                  | central-america                         | sPp     |
| central_african_republic                | africa                           | central_african_republic                | sPp     |
| central_federal_district                | russia                           | central_federal_district                | sPp     |
| central_java                            | indonesia                        | central_java                            | sPp     |
| central_kalimantan                      | indonesia                        | central_kalimantan                      | sPp     |
| central_sulawesi                        | indonesia                        | central_sulawesi                        | sPp     |
| ceuta                                   | spain                            | ceuta                                   | sPp     |
| chad                                    | africa                           | chad                                    | sPp     |
| chandigarh                              | india                            | chandigarh                              | sPp     |
| chechen_republic                        | north_caucasian_federal_district | chechen_republic                        | sPp     |
| chelyabinsk_oblast                      | ural_federal_district            | chelyabinsk_oblast                      | sPp     |
| chhattisgarh                            | india                            | chhattisgarh                            | sPp     |
| china                                   | asia                             | china                                   | sPp     |
| chongqing                               | china                            | chongqing                               | sPp     |
| christmas_island                        | australia                        | christmas_island                        | sPp     |
| chubu                                   | japan                            | chubu                                   | sPp     |
| chugoku                                 | japan                            | chugoku                                 | sPp     |
| chukotka_autonomous_okrug               | far_eastern_federal_district     | chukotka_autonomous_okrug               | sPp     |
| chuvash_republic                        | volga_federal_district           | chuvash_republic                        | sPp     |
| cocos_islands                           | australia                        | cocos_islands                           | sPp     |
| comoros                                 | africa                           | comoros                                 | sPp     |
| comunidad_de_madrid                     | spain                            | comunidad_de_madrid                     | sPp     |
| comunidad_foral_de_navarra              | spain                            | comunidad_foral_de_navarra              | sPp     |
| comunitat_valenciana                    | spain                            | comunitat_valenciana                    | sPp     |
| congo_brazzaville                       | africa                           | congo_brazzaville                       | sPp     |
| congo_kinshasa                          | africa                           | congo_kinshasa                          | sPp     |
| cook_islands                            | south-america                    | cook_islands                            | sPp     |
| coral_sea_islands                       | australia                        | coral_sea_islands                       | sPp     |
| costa_rica                              | central-america                  | costa_rica                              | sPp     |
| curacao                                 | central-america                  | curacao                                 | sPp     |
| czech_republic                          | europe                           | czech_republic                          | sPp     |
| dadra_and_nagar_haveli                  | india                            | dadra_and_nagar_haveli                  | sPp     |
| dagestan_republic                       | north_caucasian_federal_district | dagestan_republic                       | sPp     |
| daman_and_diu                           | india                            | daman_and_diu                           | sPp     |
| detmold                                 | nordrhein_westfalen              | detmold                                 | sPp     |
| djibouti                                | africa                           | djibouti                                | sPp     |
| dolnoslaskie                            | poland                           | dolnoslaskie                            | sPp     |
| dominica                                | central-america                  | dominica                                | sPp     |
| dominican_republic                      | central-america                  | dominican_republic                      | sPp     |
| drenthe                                 | netherlands                      | drenthe                                 | sPp     |
| dusseldorf                              | nordrhein_westfalen              | dusseldorf                              | sPp     |
| east                                    | england                          | east                                    | sPp     |
| east_java                               | indonesia                        | east_java                               | sPp     |
| east_kalimantan                         | indonesia                        | east_kalimantan                         | sPp     |
| east_midlands                           | england                          | east_midlands                           | sPp     |
| east_nusa_tenggara                      | indonesia                        | east_nusa_tenggara                      | sPp     |
| east_timor                              | asia                             | east_timor                              | sPp     |
| el_salvador                             | central-america                  | el_salvador                             | sPp     |
| emilia_romagna                          | italy                            | emilia_romagna                          | sPp     |
| england                                 | united_kingdom                   | england                                 | sPp     |
| equatorial_guinea                       | africa                           | equatorial_guinea                       | sPp     |
| eritrea                                 | africa                           | eritrea                                 | sPp     |
| europe                                  |                                  | europe                                  | sPp     |
| euskadi                                 | spain                            | euskadi                                 | sPp     |
| extremadura                             | spain                            | extremadura                             | sPp     |
| falkland                                | south-america                    | falkland                                | sPp     |
| far_eastern_federal_district            | russia                           | far_eastern_federal_district            | sPp     |
| fiji                                    |                                  | fiji                                    | sP      |
| fiji_east                               | south-america                    | fiji_east                               | sPp     |
| fiji_west                               | oceania                          | fiji_west                               | sPp     |
| flanders                                | belgium                          | flanders                                | sPp     |
| flevoland                               | netherlands                      | flevoland                               | sPp     |
| france                                  | europe                           | france                                  | sPp     |
| france_metro_dom_com_nc                 |                                  | france_metro_dom_com_nc                 | sP      |
| france_taaf                             | africa                           | france_taaf                             | sPp     |
| friesland                               | netherlands                      | friesland                               | sPp     |
| friuli_venezia_giulia                   | italy                            | friuli_venezia_giulia                   | sPp     |
| fujian                                  | china                            | fujian                                  | sPp     |
| gabon                                   | africa                           | gabon                                   | sPp     |
| galicia                                 | spain                            | galicia                                 | sPp     |
| gambia                                  | africa                           | gambia                                  | sPp     |
| gansu                                   | china                            | gansu                                   | sPp     |
| gelderland                              | netherlands                      | gelderland                              | sPp     |
| georgia                                 | asia                             | georgia                                 | sPp     |
| germany                                 | europe                           | germany                                 | sPp     |
| ghana                                   | africa                           | ghana                                   | sPp     |
| gibraltar                               | europe                           | gibraltar                               | sPp     |
| goa                                     | india                            | goa                                     | sPp     |
| gorontalo                               | indonesia                        | gorontalo                               | sPp     |
| greater_london                          | england                          | greater_london                          | sPp     |
| grenada                                 | central-america                  | grenada                                 | sPp     |
| groningen                               | netherlands                      | groningen                               | sPp     |
| guadeloupe                              | central-america                  | guadeloupe                              | sPp     |
| guam                                    | oceania                          | guam                                    | sPp     |
| guangdong                               | china                            | guangdong                               | sPp     |
| guangxi                                 | china                            | guangxi                                 | sPp     |
| guernesey                               | europe                           | guernesey                               | sPp     |
| guinea                                  | africa                           | guinea                                  | sPp     |
| guizhou                                 | china                            | guizhou                                 | sPp     |
| gujarat                                 | india                            | gujarat                                 | sPp     |
| guyana                                  | south-america                    | guyana                                  | sPp     |
| guyane                                  | south-america                    | guyane                                  | sPp     |
| hainan                                  | china                            | hainan                                  | sPp     |
| haryana                                 | india                            | haryana                                 | sPp     |
| hebei                                   | china                            | hebei                                   | sPp     |
| heilongjiang                            | china                            | heilongjiang                            | sPp     |
| henan                                   | china                            | henan                                   | sPp     |
| himachal_pradesh                        | india                            | himachal_pradesh                        | sPp     |
| hokkaido                                | japan                            | hokkaido                                | sPp     |
| honduras                                | central-america                  | honduras                                | sPp     |
| hong_kong                               | china                            | hong_kong                               | sPp     |
| hubei                                   | china                            | hubei                                   | sPp     |
| hunan                                   | china                            | hunan                                   | sPp     |
| illes_balears                           | spain                            | illes_balears                           | sPp     |
| india                                   | asia                             | india                                   | sPp     |
| indonesia                               | asia                             | indonesia                               | sPp     |
| ingushetia_republic                     | north_caucasian_federal_district | ingushetia_republic                     | sPp     |
| inner_mongolia                          | china                            | inner_mongolia                          | sPp     |
| ireland                                 | europe                           | ireland                                 | sPp     |
| irkutsk_oblast                          | siberian_federal_district        | irkutsk_oblast                          | sPp     |
| israel                                  | asia                             | israel                                  | sPp     |
| israel_and_palestine                    | asia                             | israel_and_palestine                    | sPp     |
| italy                                   | europe                           | italy                                   | sPp     |
| ivanovo_oblast                          | central_federal_district         | ivanovo_oblast                          | sPp     |
| ivory_coast                             | africa                           | ivory_coast                             | sPp     |
| jakarta                                 | indonesia                        | jakarta                                 | sPp     |
| jamaica                                 | central-america                  | jamaica                                 | sPp     |
| jambi                                   | indonesia                        | jambi                                   | sPp     |
| jammu_and_kashmir                       | india                            | jammu_and_kashmir                       | sPp     |
| japan                                   | asia                             | japan                                   | sPp     |
| jersey                                  | europe                           | jersey                                  | sPp     |
| jewish_autonomous_oblast                | far_eastern_federal_district     | jewish_autonomous_oblast                | sPp     |
| jharkhand                               | india                            | jharkhand                               | sPp     |
| jiangsu                                 | china                            | jiangsu                                 | sPp     |
| jiangxi                                 | china                            | jiangxi                                 | sPp     |
| jihocesky                               | czech_republic                   | jihocesky                               | sPp     |
| jihomoravsky                            | czech_republic                   | jihomoravsky                            | sPp     |
| jilin                                   | china                            | jilin                                   | sPp     |
| kabardino_balkar_republic               | north_caucasian_federal_district | kabardino_balkar_republic               | sPp     |
| kaliningrad_oblast                      | northwestern_federal_district    | kaliningrad_oblast                      | sPp     |
| kalmykia_republic                       | southern_federal_district        | kalmykia_republic                       | sPp     |
| kaluga_oblast                           | central_federal_district         | kaluga_oblast                           | sPp     |
| kamchatka_krai                          | far_eastern_federal_district     | kamchatka_krai                          | sPp     |
| kansai                                  | japan                            | kansai                                  | sPp     |
| kanto                                   | japan                            | kanto                                   | sPp     |
| karachay_cherkess_republic              | north_caucasian_federal_district | karachay_cherkess_republic              | sPp     |
| karelia_republic                        | northwestern_federal_district    | karelia_republic                        | sPp     |
| karlovarsky                             | czech_republic                   | karlovarsky                             | sPp     |
| karnataka                               | india                            | karnataka                               | sPp     |
| karnten                                 | austria                          | karnten                                 | sPp     |
| kemerovo_oblast                         | siberian_federal_district        | kemerovo_oblast                         | sPp     |
| kenya                                   | africa                           | kenya                                   | sPp     |
| kerala                                  | india                            | kerala                                  | sPp     |
| khabarovsk_krai                         | far_eastern_federal_district     | khabarovsk_krai                         | sPp     |
| khakassia_republic                      | siberian_federal_district        | khakassia_republic                      | sPp     |
| khanty_mansi_autonomous_okrug           | ural_federal_district            | khanty_mansi_autonomous_okrug           | sPp     |
| kiribati                                |                                  | kiribati                                | sP      |
| kiribati_east                           | south-america                    | kiribati_east                           | sPp     |
| kiribati_west                           | oceania                          | kiribati_west                           | sPp     |
| kirov_oblast                            | volga_federal_district           | kirov_oblast                            | sPp     |
| koln                                    | nordrhein_westfalen              | koln                                    | sPp     |
| komi_republic                           | northwestern_federal_district    | komi_republic                           | sPp     |
| kosicky                                 | slovakia                         | kosicky                                 | sPp     |
| kostroma_oblast                         | central_federal_district         | kostroma_oblast                         | sPp     |
| kralovehradecky                         | czech_republic                   | kralovehradecky                         | sPp     |
| krasnodar_krai                          | southern_federal_district        | krasnodar_krai                          | sPp     |
| krasnoyarsk_krai                        | siberian_federal_district        | krasnoyarsk_krai                        | sPp     |
| kujawsko_pomorskie                      | poland                           | kujawsko_pomorskie                      | sPp     |
| kurgan_oblast                           | ural_federal_district            | kurgan_oblast                           | sPp     |
| kursk_oblast                            | central_federal_district         | kursk_oblast                            | sPp     |
| kuwait                                  | asia                             | kuwait                                  | sPp     |
| kyushu                                  | japan                            | kyushu                                  | sPp     |
| la_rioja                                | spain                            | la_rioja                                | sPp     |
| lakshadweep                             | india                            | lakshadweep                             | sPp     |
| lampung                                 | indonesia                        | lampung                                 | sPp     |
| laos                                    | asia                             | laos                                    | sPp     |
| lazio                                   | italy                            | lazio                                   | sPp     |
| leningrad_oblast                        | northwestern_federal_district    | leningrad_oblast                        | sPp     |
| lesotho                                 | africa                           | lesotho                                 | sPp     |
| liaoning                                | china                            | liaoning                                | sPp     |
| liberecky                               | czech_republic                   | liberecky                               | sPp     |
| liguria                                 | italy                            | liguria                                 | sPp     |
| limburg                                 | netherlands                      | limburg                                 | sPp     |
| lipetsk_oblast                          | central_federal_district         | lipetsk_oblast                          | sPp     |
| lodzkie                                 | poland                           | lodzkie                                 | sPp     |
| lombardia                               | italy                            | lombardia                               | sPp     |
| lubelskie                               | poland                           | lubelskie                               | sPp     |
| lubuskie                                | poland                           | lubuskie                                | sPp     |
| luxembourg                              | europe                           | luxembourg                              | sPp     |
| macau                                   | china                            | macau                                   | sPp     |
| madhya_pradesh                          | india                            | madhya_pradesh                          | sPp     |
| magadan_oblast                          | far_eastern_federal_district     | magadan_oblast                          | sPp     |
| maharashtra                             | india                            | maharashtra                             | sPp     |
| malawi                                  | africa                           | malawi                                  | sPp     |
| malaysia                                | asia                             | malaysia                                | sPp     |
| maldives                                | asia                             | maldives                                | sPp     |
| mali                                    | africa                           | mali                                    | sPp     |
| malopolskie                             | poland                           | malopolskie                             | sPp     |
| maluku                                  | indonesia                        | maluku                                  | sPp     |
| manipur                                 | india                            | manipur                                 | sPp     |
| marche                                  | italy                            | marche                                  | sPp     |
| mari_el_republic                        | volga_federal_district           | mari_el_republic                        | sPp     |
| marshall-islands                        | south-america                    | marshall-islands                        | sPp     |
| marshall_islands                        | oceania                          | marshall_islands                        | sPp     |
| martinique                              | central-america                  | martinique                              | sPp     |
| mauritania                              | africa                           | mauritania                              | sPp     |
| mauritius                               | africa                           | mauritius                               | sPp     |
| mayotte                                 | africa                           | mayotte                                 | sPp     |
| mazowieckie                             | poland                           | mazowieckie                             | sPp     |
| meghalaya                               | india                            | meghalaya                               | sPp     |
| melilla                                 | spain                            | melilla                                 | sPp     |
| micronesia                              | oceania                          | micronesia                              | sPp     |
| mizoram                                 | india                            | mizoram                                 | sPp     |
| molise                                  | italy                            | molise                                  | sPp     |
| monaco                                  | europe                           | monaco                                  | sPp     |
| montserrat                              | central-america                  | montserrat                              | sPp     |
| moravskoslezsky                         | czech_republic                   | moravskoslezsky                         | sPp     |
| mordovia_republic                       | volga_federal_district           | mordovia_republic                       | sPp     |
| moscow                                  | central_federal_district         | moscow                                  | sPp     |
| moscow_oblast                           | central_federal_district         | moscow_oblast                           | sPp     |
| mozambique                              | africa                           | mozambique                              | sPp     |
| munster                                 | nordrhein_westfalen              | munster                                 | sPp     |
| murmansk_oblast                         | northwestern_federal_district    | murmansk_oblast                         | sPp     |
| myanmar                                 | asia                             | myanmar                                 | sPp     |
| nagaland                                | india                            | nagaland                                | sPp     |
| namibia                                 | africa                           | namibia                                 | sPp     |
| national_capital_territory_of_delhi     | india                            | national_capital_territory_of_delhi     | sPp     |
| nauri                                   | south-america                    | nauri                                   | sPp     |
| nauru                                   | oceania                          | nauru                                   | sPp     |
| nenets_autonomous_okrug                 | northwestern_federal_district    | nenets_autonomous_okrug                 | sPp     |
| netherlands                             | europe                           | netherlands                             | sPp     |
| new_caledonia                           | oceania                          | new_caledonia                           | sPp     |
| new_south_wales                         | australia                        | new_south_wales                         | sPp     |
| nicaragua                               | central-america                  | nicaragua                               | sPp     |
| niederosterreich                        | austria                          | niederosterreich                        | sPp     |
| niger                                   | africa                           | niger                                   | sPp     |
| ningxia                                 | china                            | ningxia                                 | sPp     |
| nitriansky                              | slovakia                         | nitriansky                              | sPp     |
| niue                                    | south-america                    | niue                                    | sPp     |
| nizhny_novgorod_oblast                  | volga_federal_district           | nizhny_novgorod_oblast                  | sPp     |
| noord_brabant                           | netherlands                      | noord_brabant                           | sPp     |
| noord_holland                           | netherlands                      | noord_holland                           | sPp     |
| nordrhein_westfalen                     | germany                          | nordrhein_westfalen                     | sPp     |
| norfolk_island                          | australia                        | norfolk_island                          | sPp     |
| north-america                           |                                  | north-america                           | p       |
| north_caucasian_federal_district        | russia                           | north_caucasian_federal_district        | sPp     |
| north_east                              | england                          | north_east                              | sPp     |
| north_kalimantan                        | indonesia                        | north_kalimantan                        | sPp     |
| north_maluku                            | indonesia                        | north_maluku                            | sPp     |
| north_ossetia_alania_republic           | north_caucasian_federal_district | north_ossetia_alania_republic           | sPp     |
| north_sulawesi                          | indonesia                        | north_sulawesi                          | sPp     |
| north_sumatra                           | indonesia                        | north_sumatra                           | sPp     |
| north_west                              | england                          | north_west                              | sPp     |
| northern_ireland                        | united_kingdom                   | northern_ireland                        | sPp     |
| northern_mariana_islands                | oceania                          | northern_mariana_islands                | sPp     |
| northern_territory                      | australia                        | northern_territory                      | sPp     |
| northwestern_federal_district           | russia                           | northwestern_federal_district           | sPp     |
| novgorod_oblast                         | northwestern_federal_district    | novgorod_oblast                         | sPp     |
| novosibirsk_oblast                      | siberian_federal_district        | novosibirsk_oblast                      | sPp     |
| oberosterreich                          | austria                          | oberosterreich                          | sPp     |
| oceania                                 |                                  | oceania                                 | sPp     |
| odisha                                  | india                            | odisha                                  | sPp     |
| olomoucky                               | czech_republic                   | olomoucky                               | sPp     |
| oman                                    | asia                             | oman                                    | sPp     |
| omsk_oblast                             | siberian_federal_district        | omsk_oblast                             | sPp     |
| opolskie                                | poland                           | opolskie                                | sPp     |
| orenburg_oblast                         | volga_federal_district           | orenburg_oblast                         | sPp     |
| oryol_oblast                            | central_federal_district         | oryol_oblast                            | sPp     |
| overijssel                              | netherlands                      | overijssel                              | sPp     |
| palau                                   | oceania                          | palau                                   | sPp     |
| palestine                               | asia                             | palestine                               | sPp     |
| panama                                  | central-america                  | panama                                  | sPp     |
| papua                                   | indonesia                        | papua                                   | sPp     |
| papua_new_guinea                        | oceania                          | papua_new_guinea                        | sPp     |
| paraguay                                | south-america                    | paraguay                                | sPp     |
| pardubicky                              | czech_republic                   | pardubicky                              | sPp     |
| penza_oblast                            | volga_federal_district           | penza_oblast                            | sPp     |
| perm_krai                               | volga_federal_district           | perm_krai                               | sPp     |
| piemonte                                | italy                            | piemonte                                | sPp     |
| pitcairn                                | south-america                    | pitcairn                                | sPp     |
| plzensky                                | czech_republic                   | plzensky                                | sPp     |
| podkarpackie                            | poland                           | podkarpackie                            | sPp     |
| podlaskie                               | poland                           | podlaskie                               | sPp     |
| poland                                  | europe                           | poland                                  | sPp     |
| polynesie                               | oceania                          | polynesie                               | sPp     |
| pomorskie                               | poland                           | pomorskie                               | sPp     |
| praha                                   | czech_republic                   | praha                                   | sPp     |
| presovsky                               | slovakia                         | presovsky                               | sPp     |
| primorsky_krai                          | far_eastern_federal_district     | primorsky_krai                          | sPp     |
| pskov_oblast                            | northwestern_federal_district    | pskov_oblast                            | sPp     |
| puducherry                              | india                            | puducherry                              | sPp     |
| puerto_rico                             | central-america                  | puerto_rico                             | sPp     |
| puglia                                  | italy                            | puglia                                  | sPp     |
| punjab                                  | india                            | punjab                                  | sPp     |
| qatar                                   | asia                             | qatar                                   | sPp     |
| qinghai                                 | china                            | qinghai                                 | sPp     |
| quebec                                  | canada                           | quebec                                  | sPp     |
| queensland                              | australia                        | queensland                              | sPp     |
| rajasthan                               | india                            | rajasthan                               | sPp     |
| region_de_murcia                        | spain                            | region_de_murcia                        | sPp     |
| reunion                                 | africa                           | reunion                                 | sPp     |
| riau                                    | indonesia                        | riau                                    | sPp     |
| riau_islands                            | indonesia                        | riau_islands                            | sPp     |
| rostov_oblast                           | southern_federal_district        | rostov_oblast                           | sPp     |
| russia                                  |                                  | russia                                  | sPp     |
| rwanda                                  | africa                           | rwanda                                  | sPp     |
| ryazan_oblast                           | central_federal_district         | ryazan_oblast                           | sPp     |
| saint_barthelemy                        | central-america                  | saint_barthelemy                        | sPp     |
| saint_helena_ascension_tristan_da_cunha | africa                           | saint_helena_ascension_tristan_da_cunha | sPp     |
| saint_kitts_and_nevis                   | central-america                  | saint_kitts_and_nevis                   | sPp     |
| saint_lucia                             | central-america                  | saint_lucia                             | sPp     |
| saint_martin                            | central-america                  | saint_martin                            | sPp     |
| saint_petersburg                        | northwestern_federal_district    | saint_petersburg                        | sPp     |
| saint_pierre_et_miquelon                | north-america                    | saint_pierre_et_miquelon                | sPp     |
| saint_vincent_and_the_grenadines        | central-america                  | saint_vincent_and_the_grenadines        | sPp     |
| sakha_republic                          | far_eastern_federal_district     | sakha_republic                          | sPp     |
| sakhalin_oblast                         | far_eastern_federal_district     | sakhalin_oblast                         | sPp     |
| salzburg                                | austria                          | salzburg                                | sPp     |
| samara_oblast                           | volga_federal_district           | samara_oblast                           | sPp     |
| samoa                                   | south-america                    | samoa                                   | sPp     |
| sao_tome_and_principe                   | africa                           | sao_tome_and_principe                   | sPp     |
| saratov_oblast                          | volga_federal_district           | saratov_oblast                          | sPp     |
| sardegna                                | italy                            | sardegna                                | sPp     |
| saudi_arabia                            | asia                             | saudi_arabia                            | sPp     |
| senegal                                 | africa                           | senegal                                 | sPp     |
| seychelles                              | africa                           | seychelles                              | sPp     |
| shaanxi                                 | china                            | shaanxi                                 | sPp     |
| shandong                                | china                            | shandong                                | sPp     |
| shanghai                                | china                            | shanghai                                | sPp     |
| shanxi                                  | china                            | shanxi                                  | sPp     |
| shikoku                                 | japan                            | shikoku                                 | sPp     |
| siberian_federal_district               | russia                           | siberian_federal_district               | sPp     |
| sichuan                                 | china                            | sichuan                                 | sPp     |
| sicilia                                 | italy                            | sicilia                                 | sPp     |
| sikkim                                  | india                            | sikkim                                  | sPp     |
| singapore                               | asia                             | singapore                               | sPp     |
| sint_maarten                            | central-america                  | sint_maarten                            | sPp     |
| slaskie                                 | poland                           | slaskie                                 | sPp     |
| slovakia                                | europe                           | slovakia                                | sPp     |
| smolensk_oblast                         | central_federal_district         | smolensk_oblast                         | sPp     |
| solomon_islands                         | oceania                          | solomon_islands                         | sPp     |
| south-america                           |                                  | south-america                           | sPp     |
| south_africa                            | africa                           | south_africa                            | sPp     |
| south_australia                         | australia                        | south_australia                         | sPp     |
| south_east                              | england                          | south_east                              | sPp     |
| south_georgia_and_south_sandwich        | south-america                    | south_georgia_and_south_sandwich        | sPp     |
| south_kalimantan                        | indonesia                        | south_kalimantan                        | sPp     |
| south_sudan                             | africa                           | south_sudan                             | sPp     |
| south_sulawesi                          | indonesia                        | south_sulawesi                          | sPp     |
| south_sumatra                           | indonesia                        | south_sumatra                           | sPp     |
| south_west                              | england                          | south_west                              | sPp     |
| southeast_sulawesi                      | indonesia                        | southeast_sulawesi                      | sPp     |
| southern_federal_district               | russia                           | southern_federal_district               | sPp     |
| spain                                   | europe                           | spain                                   | sPp     |
| stavropol_krai                          | north_caucasian_federal_district | stavropol_krai                          | sPp     |
| steiermark                              | austria                          | steiermark                              | sPp     |
| stredocesky                             | czech_republic                   | stredocesky                             | sPp     |
| sudan                                   | africa                           | sudan                                   | sPp     |
| suriname                                | south-america                    | suriname                                | sPp     |
| sverdlovsk_oblast                       | ural_federal_district            | sverdlovsk_oblast                       | sPp     |
| swaziland                               | africa                           | swaziland                               | sPp     |
| swietokrzyskie                          | poland                           | swietokrzyskie                          | sPp     |
| taaf                                    | south-america                    | taaf                                    | sPp     |
| tambov_oblast                           | central_federal_district         | tambov_oblast                           | sPp     |
| tamil_nadu                              | india                            | tamil_nadu                              | sPp     |
| tasmania                                | australia                        | tasmania                                | sPp     |
| tatarstan_republic                      | volga_federal_district           | tatarstan_republic                      | sPp     |
| telangana                               | india                            | telangana                               | sPp     |
| tianjin                                 | china                            | tianjin                                 | sPp     |
| tibet                                   | china                            | tibet                                   | sPp     |
| tirol                                   | austria                          | tirol                                   | sPp     |
| togo                                    | africa                           | togo                                    | sPp     |
| tohoku                                  | japan                            | tohoku                                  | sPp     |
| tokelau                                 | south-america                    | tokelau                                 | sPp     |
| tomsk_oblast                            | siberian_federal_district        | tomsk_oblast                            | sPp     |
| tonga                                   | south-america                    | tonga                                   | sPp     |
| toscana                                 | italy                            | toscana                                 | sPp     |
| trenciansky                             | slovakia                         | trenciansky                             | sPp     |
| trentino_alto_adige                     | italy                            | trentino_alto_adige                     | sPp     |
| trinidad_and_tobago                     | central-america                  | trinidad_and_tobago                     | sPp     |
| tripura                                 | india                            | tripura                                 | sPp     |
| trnavsky                                | slovakia                         | trnavsky                                | sPp     |
| tula_oblast                             | central_federal_district         | tula_oblast                             | sPp     |
| tunisia                                 | africa                           | tunisia                                 | sPp     |
| turks_and_caicos_islands                | central-america                  | turks_and_caicos_islands                | sPp     |
| tuva_republic                           | siberian_federal_district        | tuva_republic                           | sPp     |
| tuvalu                                  | oceania                          | tuvalu                                  | sPp     |
| tver_oblast                             | central_federal_district         | tver_oblast                             | sPp     |
| tyumen_oblast                           | ural_federal_district            | tyumen_oblast                           | sPp     |
| udmurt_republic                         | volga_federal_district           | udmurt_republic                         | sPp     |
| uganda                                  | africa                           | uganda                                  | sPp     |
| ulyanovsk_oblast                        | volga_federal_district           | ulyanovsk_oblast                        | sPp     |
| umbria                                  | italy                            | umbria                                  | sPp     |
| united_arab_emirates                    | asia                             | united_arab_emirates                    | sPp     |
| united_kingdom                          | europe                           | united_kingdom                          | sPp     |
| united_states_virgin_islands            | central-america                  | united_states_virgin_islands            | sPp     |
| ural_federal_district                   | russia                           | ural_federal_district                   | sPp     |
| usa_virgin_islands                      | central-america                  | usa_virgin_islands                      | sPp     |
| ustecky                                 | czech_republic                   | ustecky                                 | sPp     |
| utrecht                                 | netherlands                      | utrecht                                 | sPp     |
| uttar_pradesh                           | india                            | uttar_pradesh                           | sPp     |
| uttarakhand                             | india                            | uttarakhand                             | sPp     |
| valle_aosta                             | italy                            | valle_aosta                             | sPp     |
| vanuatu                                 | oceania                          | vanuatu                                 | sPp     |
| veneto                                  | italy                            | veneto                                  | sPp     |
| venezuela                               | south-america                    | venezuela                               | sPp     |
| victoria                                | australia                        | victoria                                | sPp     |
| vladimir_oblast                         | central_federal_district         | vladimir_oblast                         | sPp     |
| volga_federal_district                  | russia                           | volga_federal_district                  | sPp     |
| volgograd_oblast                        | southern_federal_district        | volgograd_oblast                        | sPp     |
| vologda_oblast                          | northwestern_federal_district    | vologda_oblast                          | sPp     |
| vorarlberg                              | austria                          | vorarlberg                              | sPp     |
| voronezh_oblast                         | central_federal_district         | voronezh_oblast                         | sPp     |
| vysocina                                | czech_republic                   | vysocina                                | sPp     |
| wallis_et_futuna                        | oceania                          | wallis_et_futuna                        | sPp     |
| wallonia_french_community               | belgium                          | wallonia_french_community               | sPp     |
| wallonia_german_community               | belgium                          | wallonia_german_community               | sPp     |
| warminsko_mazurskie                     | poland                           | warminsko_mazurskie                     | sPp     |
| west_bengal                             | india                            | west_bengal                             | sPp     |
| west_java                               | indonesia                        | west_java                               | sPp     |
| west_kalimantan                         | indonesia                        | west_kalimantan                         | sPp     |
| west_midlands                           | england                          | west_midlands                           | sPp     |
| west_nusa_tenggara                      | indonesia                        | west_nusa_tenggara                      | sPp     |
| west_papua                              | indonesia                        | west_papua                              | sPp     |
| west_sulawesi                           | indonesia                        | west_sulawesi                           | sPp     |
| west_sumatra                            | indonesia                        | west_sumatra                            | sPp     |
| western_australia                       | australia                        | western_australia                       | sPp     |
| western_sahara                          | africa                           | western_sahara                          | sPp     |
| wielkopolskie                           | poland                           | wielkopolskie                           | sPp     |
| wien                                    | austria                          | wien                                    | sPp     |
| xinjiang                                | china                            | xinjiang                                | sPp     |
| yamalo_nenets_autonomous_okrug          | ural_federal_district            | yamalo_nenets_autonomous_okrug          | sPp     |
| yaroslavl_oblast                        | central_federal_district         | yaroslavl_oblast                        | sPp     |
| yogyakarta                              | indonesia                        | yogyakarta                              | sPp     |
| yorkshire_and_the_humber                | england                          | yorkshire_and_the_humber                | sPp     |
| yunnan                                  | china                            | yunnan                                  | sPp     |
| zabaykalsky_krai                        | siberian_federal_district        | zabaykalsky_krai                        | sPp     |
| zachodniopomorskie                      | poland                           | zachodniopomorskie                      | sPp     |
| zambia                                  | africa                           | zambia                                  | sPp     |
| zeeland                                 | netherlands                      | zeeland                                 | sPp     |
| zhejiang                                | china                            | zhejiang                                | sPp     |
| zilinsky                                | slovakia                         | zilinsky                                | sPp     |
| zimbabwe                                | africa                           | zimbabwe                                | sPp     |
| zlinsky                                 | czech_republic                   | zlinsky                                 | sPp     |
| zuid_holland                            | netherlands                      | zuid_holland                            | sPp     |
Total elements: 506

## List of elements from glis-lab.info
| SHORTNAME | IS IN |           LONG NAME            | FORMATS |
|-----------|-------|--------------------------------|---------|
| AM        |       | Армения                        | PBp     |
| AZ        |       | Азербайджан                    | PBp     |
| BY        |       | Беларусь                       | PBp     |
| EE        |       | Эстония                        | PBp     |
| GE        |       | Грузия                         | PBp     |
| KG        |       | Киргизия                       | PBp     |
| KZ        |       | Казахстан                      | PBp     |
| LT        |       | Литва                          | PBp     |
| LV        |       | Латвия                         | PBp     |
| MD        |       | Молдова                        | PBp     |
| RU        |       | Российская Федерация           | PBp     |
| RU-AD     |       | Адыгея                         | PBp     |
| RU-AL     |       | Алтай                          | PBp     |
| RU-ALT    |       | Алтайский край                 | PBp     |
| RU-AMU    |       | Амурская область               | PBp     |
| RU-ARK    |       | Архангельская область          | PBp     |
| RU-AST    |       | Астраханская область           | PBp     |
| RU-BA     |       | Башкирия                       | PBp     |
| RU-BEL    |       | Белгородская область           | PBp     |
| RU-BRY    |       | Брянская область               | PBp     |
| RU-BU     |       | Бурятия                        | PBp     |
| RU-CE     |       | Чечня                          | PBp     |
| RU-CHE    |       | Челябинская область            | PBp     |
| RU-CHU    |       | Чукотский автономный округ     | PBp     |
| RU-CR     |       | Крым                           | PBp     |
| RU-CU     |       | Чувашия                        | PBp     |
| RU-DA     |       | Дагестан                       | PBp     |
| RU-IN     |       | Ингушетия                      | PBp     |
| RU-IRK    |       | Иркутская область              | PBp     |
| RU-IVA    |       | Ивановская область             | PBp     |
| RU-KAM    |       | Камчатский край                | PBp     |
| RU-KB     |       | Кабардино-Балкария             | PBp     |
| RU-KC     |       | Карачаево-Черкессия            | PBp     |
| RU-KDA    |       | Краснодарский край             | PBp     |
| RU-KEM    |       | Кемеровская область            | PBp     |
| RU-KGD    |       | Калининградская область        | PBp     |
| RU-KGN    |       | Курганская область             | PBp     |
| RU-KHA    |       | Хабаровский край               | PBp     |
| RU-KHM    |       | Ханты-Мансийский автономный    | PBp     |
|           |       | округ                          |         |
| RU-KIR    |       | Кировская область              | PBp     |
| RU-KK     |       | Хакасия                        | PBp     |
| RU-KL     |       | Калмыкия                       | PBp     |
| RU-KLU    |       | Калужская область              | PBp     |
| RU-KO     |       | Коми                           | PBp     |
| RU-KOS    |       | Костромская область            | PBp     |
| RU-KR     |       | Карелия                        | PBp     |
| RU-KRS    |       | Курская область                | PBp     |
| RU-KYA    |       | Красноярский край              | PBp     |
| RU-LEN    |       | Ленинградская область          | PBp     |
| RU-LIP    |       | Липецкая область               | PBp     |
| RU-MAG    |       | Магаданская область            | PBp     |
| RU-ME     |       | Марий Эл                       | PBp     |
| RU-MO     |       | Мордовия                       | PBp     |
| RU-MOS    |       | Московская область             | PBp     |
| RU-MOW    |       | Москва                         | PBp     |
| RU-MUR    |       | Мурманская область             | PBp     |
| RU-NEN    |       | Ненецкий автономный округ      | PBp     |
| RU-NGR    |       | Новгородская область           | PBp     |
| RU-NIZ    |       | Нижегородская область          | PBp     |
| RU-NVS    |       | Новосибирская область          | PBp     |
| RU-OMS    |       | Омская область                 | PBp     |
| RU-ORE    |       | Оренбургская область           | PBp     |
| RU-ORL    |       | Орловская область              | PBp     |
| RU-PER    |       | Пермский край                  | PBp     |
| RU-PNZ    |       | Пензенская область             | PBp     |
| RU-PRI    |       | Приморский край                | PBp     |
| RU-PSK    |       | Псковская область              | PBp     |
| RU-ROS    |       | Ростовская область             | PBp     |
| RU-RYA    |       | Рязанская область              | PBp     |
| RU-SA     |       | Якутия                         | PBp     |
| RU-SAK    |       | Сахалинская область            | PBp     |
| RU-SAM    |       | Самарская область              | PBp     |
| RU-SAR    |       | Саратовская область            | PBp     |
| RU-SE     |       | Северная Осетия                | PBp     |
| RU-SEV    |       | Севастополь                    | PBp     |
| RU-SMO    |       | Смоленская область             | PBp     |
| RU-SPE    |       | Санкт-Петербург                | PBp     |
| RU-STA    |       | Ставропольский край            | PBp     |
| RU-SVE    |       | Свердловская область           | PBp     |
| RU-TA     |       | Татарстан                      | PBp     |
| RU-TAM    |       | Тамбовская область             | PBp     |
| RU-TOM    |       | Томская область                | PBp     |
| RU-TUL    |       | Тульская область               | PBp     |
| RU-TVE    |       | Тверская область               | PBp     |
| RU-TY     |       | Тува                           | PBp     |
| RU-TYU    |       | Тюменская область              | PBp     |
| RU-UD     |       | Удмуртия                       | PBp     |
| RU-ULY    |       | Ульяновская область            | PBp     |
| RU-VGG    |       | Волгоградская область          | PBp     |
| RU-VLA    |       | Владимирская область           | PBp     |
| RU-VLG    |       | Вологодская область            | PBp     |
| RU-VOR    |       | Воронежская область            | PBp     |
| RU-YAN    |       | Ямало-Ненецкий автономный      | PBp     |
|           |       | округ                          |         |
| RU-YAR    |       | Ярославская область            | PBp     |
| RU-YEV    |       | Еврейская автономная область   | PBp     |
| RU-ZAB    |       | Забайкальский край             | PBp     |
| TJ        |       | Таджикистан                    | PBp     |
| TM        |       | Туркмения                      | PBp     |
| UA        |       | Украина                        | PBp     |
| UZ        |       | Узбекистан                     | PBp     |
| local     |       | локальное покрытие             | PBp     |
Total elements: 101
