# download-geofabrik

[![Build Status](https://travis-ci.org/julien-noblet/download-geofabrik.svg?branch=master)](https://travis-ci.org/julien-noblet/download-geofabrik)
[![Join the chat at https://gitter.im/julien-noblet/download-geofabrik](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/julien-noblet/download-geofabrik?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/julien-noblet/download-geofabrik)](https://goreportcard.com/report/github.com/julien-noblet/download-geofabrik)

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
                             "geofabrik" or "openstreetmap.fr". It automatically
                             change config file if -c is unused.
  -c, --config="./geofabrik.yml"  
                             Set Config file.
  -n, --nodownload           Do not download file (test only)
  -v, --verbose              Be verbose
  -q, --quiet                Be quiet
      --proxy-http=""        Use http proxy, format: proxy_address:port
      --proxy-sock5=""       Use Sock5 proxy, format: proxy_address:port
      --proxy-user=""        Proxy user
      --proxy-pass=""        Proxy password

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

    -B, --osm.bz2   Download osm.bz2 if available
    -S, --shp.zip   Download shp.zip if available
    -P, --osm.pbf   Download osm.pbf (default)
    -H, --osh.pbf   Download osh.pbf
    -s, --state     Download state.txt file
    -p, --poly      Download poly file
    -k, --kml       Download kml file
        --no-check  dont control downloaded file with checksum

  generate
    Generate a new config file



```

## List of elements
|                  SHORTNAME                  |          IS IN           |               LONG NAME                | FORMATS |
|---------------------------------------------|--------------------------|----------------------------------------|---------|
| afghanistan                                 | Asia                     | Afghanistan                            | sPBHpSk |
| africa                                      |                          | Africa                                 | sPBHpk  |
| alabama                                     | United States of America | Alabama                                | sPBHpSk |
| alaska                                      | United States of America | Alaska                                 | sPBHpSk |
| albania                                     | Europe                   | Albania                                | sPBHpSk |
| alberta                                     | Canada                   | Alberta                                | sPBHpSk |
| algeria                                     | Africa                   | Algeria                                | sPBHpSk |
| alps                                        | Europe                   | Alps                                   | sPBHpk  |
| alsace                                      | France                   | Alsace                                 | sPBHpSk |
| andorra                                     | Europe                   | Andorra                                | sPBHpSk |
| angola                                      | Africa                   | Angola                                 | sPBHpSk |
| antarctica                                  |                          | Antarctica                             | sPBHpSk |
| aquitaine                                   | France                   | Aquitaine                              | sPBHpSk |
| argentina                                   | South America            | Argentina                              | sPBHpSk |
| arizona                                     | United States of America | Arizona                                | sPBHpSk |
| arkansas                                    | United States of America | Arkansas                               | sPBHpSk |
| arnsberg-regbez                             | Nordrhein-Westfalen      | Regierungsbezirk Arnsberg              | sPBHpSk |
| asia                                        |                          | Asia                                   | sPBHpk  |
| australia                                   | Australia and Oceania    | Australia                              | sPBHpk  |
| australia-oceania                           |                          | Australia and Oceania                  | sPBHpk  |
| austria                                     | Europe                   | Austria                                | sPBHpSk |
| auvergne                                    | France                   | Auvergne                               | sPBHpSk |
| azerbaijan                                  | Asia                     | Azerbaijan                             | sPBHpSk |
| azores                                      | Europe                   | Azores                                 | sPBHpSk |
| baden-wuerttemberg                          | Germany                  | Baden-Württemberg                      | sPBHpSk |
| bangladesh                                  | Asia                     | Bangladesh                             | sPBHpSk |
| basse-normandie                             | France                   | Basse-Normandie                        | sPBHpSk |
| bayern                                      | Germany                  | Bayern                                 | sPBHpSk |
| bedfordshire                                | England                  | Bedfordshire                           | sPBHpSk |
| belarus                                     | Europe                   | Belarus                                | sPBHpSk |
| belgium                                     | Europe                   | Belgium                                | sPBHpSk |
| belize                                      | Central America          | Belize                                 | sPBHpSk |
| benin                                       | Africa                   | Benin                                  | sPBHpSk |
| berkshire                                   | England                  | Berkshire                              | sPBHpSk |
| berlin                                      | Germany                  | Berlin                                 | sPBHpSk |
| bhutan                                      | Asia                     | Bhutan                                 | sPBHpSk |
| bolivia                                     | South America            | Bolivia                                | sPBHpSk |
| bosnia-herzegovina                          | Europe                   | Bosnia-Herzegovina                     | sPBHpSk |
| botswana                                    | Africa                   | Botswana                               | sPBHpSk |
| bourgogne                                   | France                   | Bourgogne                              | sPBHpSk |
| brandenburg                                 | Germany                  | Brandenburg (mit Berlin)               | sPBHpSk |
| brazil                                      | South America            | Brazil                                 | sPBHpSk |
| bremen                                      | Germany                  | Bremen                                 | sPBHpSk |
| bretagne                                    | France                   | Bretagne                               | sPBHpSk |
| bristol                                     | England                  | Bristol                                | sPBHpSk |
| british-columbia                            | Canada                   | British Columbia                       | sPBHpSk |
| british-isles                               | Europe                   | British Isles                          | sPBHpk  |
| buckinghamshire                             | England                  | Buckinghamshire                        | sPBHpSk |
| bulgaria                                    | Europe                   | Bulgaria                               | sPBHpSk |
| burkina-faso                                | Africa                   | Burkina Faso                           | sPBHpSk |
| burundi                                     | Africa                   | Burundi                                | sPBHpSk |
| california                                  | United States of America | California                             | sPBHpSk |
| cambodia                                    | Asia                     | Cambodia                               | sPBHpSk |
| cambridgeshire                              | England                  | Cambridgeshire                         | sPBHpSk |
| cameroon                                    | Africa                   | Cameroon                               | sPBHpSk |
| canada                                      | North America            | Canada                                 | sPBHpk  |
| canary-islands                              | Africa                   | Canary Islands                         | sPBHpSk |
| cape-verde                                  | Africa                   | Cape Verde                             | sPBHpSk |
| central-african-republic                    | Africa                   | Central African Republic               | sPBHpSk |
| central-america                             |                          | Central America                        | sPBHpk  |
| central-fed-district                        | Russian Federation       | Central Federal District               | sPBHpSk |
| centre                                      | France                   | Centre                                 | sPBHpSk |
| centro                                      | Italy                    | Centro                                 | sPBHpSk |
| chad                                        | Africa                   | Chad                                   | sPBHpSk |
| champagne-ardenne                           | France                   | Champagne Ardenne                      | sPBHpSk |
| cheshire                                    | England                  | Cheshire                               | sPBHpSk |
| chile                                       | South America            | Chile                                  | sPBHpSk |
| china                                       | Asia                     | China                                  | sPBHpSk |
| chubu                                       | Japan                    | Chūbu region                           | sPBHpSk |
| chugoku                                     | Japan                    | Chūgoku region                         | sPBHpSk |
| colombia                                    | South America            | Colombia                               | sPBHpSk |
| colorado                                    | United States of America | Colorado                               | sPBHpSk |
| comores                                     | Africa                   | Comores                                | sPBHpSk |
| congo-brazzaville                           | Africa                   | Congo (Republic/Brazzaville)           | sPBHpSk |
| congo-democratic-republic                   | Africa                   | Congo (Democratic                      | sPBHpSk |
|                                             |                          | Republic/Kinshasa)                     |         |
| connecticut                                 | United States of America | Connecticut                            | sPBHpSk |
| cornwall                                    | England                  | Cornwall                               | sPBHpSk |
| corse                                       | France                   | Corse                                  | sPBHpSk |
| crimean-fed-district                        | Russian Federation       | Crimean Federal District               | sPBHpSk |
| croatia                                     | Europe                   | Croatia                                | sPBHpSk |
| cuba                                        | Central America          | Cuba                                   | sPBHpSk |
| cumbria                                     | England                  | Cumbria                                | sPBHpSk |
| cyprus                                      | Europe                   | Cyprus                                 | sPBHpSk |
| czech-republic                              | Europe                   | Czech Republic                         | sPBHpSk |
| dach                                        | Europe                   | Germany, Austria, Switzerland          | sPBHpk  |
| delaware                                    | United States of America | Delaware                               | sPBHpSk |
| denmark                                     | Europe                   | Denmark                                | sPBHpSk |
| derbyshire                                  | England                  | Derbyshire                             | sPBHpSk |
| detmold-regbez                              | Nordrhein-Westfalen      | Regierungsbezirk Detmold               | sPBHpSk |
| devon                                       | England                  | Devon                                  | sPBHpSk |
| district-of-columbia                        | United States of America | District of Columbia                   | sPBHpSk |
| djibouti                                    | Africa                   | Djibouti                               | sPBHpSk |
| dolnoslaskie                                | Poland                   | Województwo dolnośląskie(Lower         | sPBHpSk |
|                                             |                          | Silesian Voivodeship)                  |         |
| dorset                                      | England                  | Dorset                                 | sPBHpSk |
| duesseldorf-regbez                          | Nordrhein-Westfalen      | Regierungsbezirk Düsseldorf            | sPBHpSk |
| east-sussex                                 | England                  | East Sussex                            | sPBHpSk |
| east-yorkshire-with-hull                    | England                  | East Yorkshire with Hull               | sPBHpSk |
| ecuador                                     | South America            | Ecuador                                | sPBHpk  |
| egypt                                       | Africa                   | Egypt                                  | sPBHpSk |
| england                                     | Great Britain            | England                                | sPBHpSk |
| equatorial-guinea                           | Africa                   | Equatorial Guinea                      | sPBHpSk |
| eritrea                                     | Africa                   | Eritrea                                | sPBHpSk |
| essex                                       | England                  | Essex                                  | sPBHpSk |
| estonia                                     | Europe                   | Estonia                                | sPBHpSk |
| ethiopia                                    | Africa                   | Ethiopia                               | sPBHpSk |
| europe                                      |                          | Europe                                 | sPBHpk  |
| far-eastern-fed-district                    | Russian Federation       | Far Eastern Federal District           | sPBHpSk |
| faroe-islands                               | Europe                   | Faroe Islands                          | sPBHpSk |
| fiji                                        | Australia and Oceania    | Fiji                                   | sPBHpSk |
| finland                                     | Europe                   | Finland                                | sPBHpSk |
| florida                                     | United States of America | Florida                                | sPBHpSk |
| france                                      | Europe                   | France                                 | sPBHpk  |
| franche-comte                               | France                   | Franche Comte                          | sPBHpSk |
| freiburg-regbez                             | Baden-Württemberg        | Regierungsbezirk Freiburg              | sPBHpSk |
| gabon                                       | Africa                   | Gabon                                  | sPBHpSk |
| gcc-states                                  | Asia                     | GCC States                             | sPBHpSk |
| georgia-eu                                  | Europe                   | Georgia (Europe country)               | sPBHpSk |
| georgia-us                                  | United States of America | Georgia (US State)                     | sPBHpSk |
| germany                                     | Europe                   | Germany                                | sPBHpk  |
| ghana                                       | Africa                   | Ghana                                  | sPBHpSk |
| gloucestershire                             | England                  | Gloucestershire                        | sPBHpSk |
| great-britain                               | Europe                   | Great Britain                          | sPBHpk  |
| greater-london                              | England                  | Greater London                         | sPBHpSk |
| greater-manchester                          | England                  | Greater Manchester                     | sPBHpSk |
| greece                                      | Europe                   | Greece                                 | sPBHpSk |
| greenland                                   | North America            | Greenland                              | sPBHpSk |
| guadeloupe                                  | France                   | Guadeloupe                             | sPBHpk  |
| guatemala                                   | Central America          | Guatemala                              | sPBHpSk |
| guinea                                      | Africa                   | Guinea                                 | sPBHpSk |
| guinea-bissau                               | Africa                   | Guinea-Bissau                          | sPBHpSk |
| guyane                                      | France                   | Guyane                                 | sPBHpk  |
| haiti-and-domrep                            | Central America          | Haiti and Dominican Republic           | sPBHpSk |
| hamburg                                     | Germany                  | Hamburg                                | sPBHpSk |
| hampshire                                   | England                  | Hampshire                              | sPBHpSk |
| haute-normandie                             | France                   | Haute-Normandie                        | sPBHpSk |
| hawaii                                      | United States of America | Hawaii                                 | sPBHpSk |
| herefordshire                               | England                  | Herefordshire                          | sPBHpSk |
| hertfordshire                               | England                  | Hertfordshire                          | sPBHpSk |
| hessen                                      | Germany                  | Hessen                                 | sPBHpSk |
| hokkaido                                    | Japan                    | Hokkaidō                               | sPBHpSk |
| hungary                                     | Europe                   | Hungary                                | sPBHpSk |
| iceland                                     | Europe                   | Iceland                                | sPBHpSk |
| idaho                                       | United States of America | Idaho                                  | sPBHpSk |
| ile-de-france                               | France                   | Ile-de-France                          | sPBHpSk |
| illinois                                    | United States of America | Illinois                               | sPBHpSk |
| india                                       | Asia                     | India                                  | sPBHpSk |
| indiana                                     | United States of America | Indiana                                | sPBHpSk |
| indonesia                                   | Asia                     | Indonesia                              | sPBHpSk |
| iowa                                        | United States of America | Iowa                                   | sPBHpSk |
| iran                                        | Asia                     | Iran                                   | sPBHpSk |
| iraq                                        | Asia                     | Iraq                                   | sPBHpSk |
| ireland-and-northern-ireland                | Europe                   | Ireland and Northern Ireland           | sPBHpSk |
| isle-of-man                                 | Europe                   | Isle of Man                            | sPBHpSk |
| isle-of-wight                               | England                  | Isle of Wight                          | sPBHpSk |
| isole                                       | Italy                    | Isole                                  | sPBHpSk |
| israel-and-palestine                        | Asia                     | Israel and Palestine                   | sPBHpSk |
| italy                                       | Europe                   | Italy                                  | sPBHpk  |
| ivory-coast                                 | Africa                   | Ivory Coast                            | sPBHpSk |
| jamaica                                     | Central America          | jamaica                                | sPBHpSk |
| japan                                       | Asia                     | Japan                                  | sPBHpk  |
| jordan                                      | Asia                     | Jordan                                 | sPBHpSk |
| kaliningrad                                 | Russian Federation       | Kaliningrad                            | sPBHpSk |
| kansai                                      | Japan                    | Kansai region (a.k.a. Kinki            | sPBHpSk |
|                                             |                          | region)                                |         |
| kansas                                      | United States of America | Kansas                                 | sPBHpSk |
| kanto                                       | Japan                    | Kantō region                           | sPBHpSk |
| karlsruhe-regbez                            | Baden-Württemberg        | Regierungsbezirk Karlsruhe             | sPBHpSk |
| kazakhstan                                  | Asia                     | Kazakhstan                             | sPBHpSk |
| kent                                        | England                  | Kent                                   | sPBHpSk |
| kentucky                                    | United States of America | Kentucky                               | sPBHpSk |
| kenya                                       | Africa                   | Kenya                                  | sPBHpSk |
| koeln-regbez                                | Nordrhein-Westfalen      | Regierungsbezirk Köln                  | sPBHpSk |
| kosovo                                      | Europe                   | Kosovo                                 | sPBHpSk |
| kujawsko-pomorskie                          | Poland                   | Województwo                            | sPBHpSk |
|                                             |                          | kujawsko-pomorskie(Kuyavian-Pomeranian |         |
|                                             |                          | Voivodeship)                           |         |
| kyrgyzstan                                  | Asia                     | Kyrgyzstan                             | sPBHpSk |
| kyushu                                      | Japan                    | Kyūshū                                 | sPBHpSk |
| lancashire                                  | England                  | Lancashire                             | sPBHpSk |
| languedoc-roussillon                        | France                   | Languedoc-Roussillon                   | sPBHpSk |
| laos                                        | Asia                     | Laos                                   | sPBHpSk |
| latvia                                      | Europe                   | Latvia                                 | sPBHpSk |
| lebanon                                     | Asia                     | Lebanon                                | sPBHpSk |
| leicestershire                              | England                  | Leicestershire                         | sPBHpSk |
| lesotho                                     | Africa                   | Lesotho                                | sPBHpSk |
| liberia                                     | Africa                   | Liberia                                | sPBHpSk |
| libya                                       | Africa                   | Libya                                  | sPBHpSk |
| liechtenstein                               | Europe                   | Liechtenstein                          | sPBHpSk |
| limousin                                    | France                   | Limousin                               | sPBHpSk |
| lincolnshire                                | England                  | Lincolnshire                           | sPBHpSk |
| lithuania                                   | Europe                   | Lithuania                              | sPBHpSk |
| lodzkie                                     | Poland                   | Województwo łódzkie(Łódź               | sPBHpSk |
|                                             |                          | Voivodeship)                           |         |
| lorraine                                    | France                   | Lorraine                               | sPBHpSk |
| louisiana                                   | United States of America | Louisiana                              | sPBHpSk |
| lubelskie                                   | Poland                   | Województwo lubelskie(Lublin           | sPBHpSk |
|                                             |                          | Voivodeship)                           |         |
| lubuskie                                    | Poland                   | Województwo lubuskie(Lubusz            | sPBHpSk |
|                                             |                          | Voivodeship)                           |         |
| luxembourg                                  | Europe                   | Luxembourg                             | sPBHpSk |
| macedonia                                   | Europe                   | Macedonia                              | sPBHpSk |
| madagascar                                  | Africa                   | Madagascar                             | sPBHpSk |
| maine                                       | United States of America | Maine                                  | sPBHpSk |
| malawi                                      | Africa                   | Malawi                                 | sPBHpSk |
| malaysia-singapore-brunei                   | Asia                     | Malaysia, Singapore, and               | sPBHpk  |
|                                             |                          | Brunei                                 |         |
| maldives                                    | Asia                     | Maldives                               | sPBHpSk |
| mali                                        | Africa                   | Mali                                   | sPBHpSk |
| malopolskie                                 | Poland                   | Województwo małopolskie(Lesser         | sPBHpSk |
|                                             |                          | Poland Voivodeship)                    |         |
| malta                                       | Europe                   | Malta                                  | sPBHpSk |
| manitoba                                    | Canada                   | Manitoba                               | sPBHpSk |
| martinique                                  | France                   | Martinique                             | sPBHpk  |
| maryland                                    | United States of America | Maryland                               | sPBHpSk |
| massachusetts                               | United States of America | Massachusetts                          | sPBHpSk |
| mauritania                                  | Africa                   | Mauritania                             | sPBHpSk |
| mauritius                                   | Africa                   | Mauritius                              | sPBHpSk |
| mayotte                                     | France                   | Mayotte                                | sPBHpk  |
| mazowieckie                                 | Poland                   | Województwo                            | sPBHpSk |
|                                             |                          | mazowieckie(Mazovian                   |         |
|                                             |                          | Voivodeship)                           |         |
| mecklenburg-vorpommern                      | Germany                  | Mecklenburg-Vorpommern                 | sPBHpSk |
| merseyside                                  | England                  | Merseyside                             | sPBHpSk |
| mexico                                      | North America            | Mexico                                 | sPBHpSk |
| michigan                                    | United States of America | Michigan                               | sPBHpSk |
| midi-pyrenees                               | France                   | Midi-Pyrenees                          | sPBHpSk |
| minnesota                                   | United States of America | Minnesota                              | sPBHpSk |
| mississippi                                 | United States of America | Mississippi                            | sPBHpSk |
| missouri                                    | United States of America | Missouri                               | sPBHpSk |
| mittelfranken                               | Bayern                   | Mittelfranken                          | sPBHpSk |
| moldova                                     | Europe                   | Moldova                                | sPBHpSk |
| monaco                                      | Europe                   | Monaco                                 | sPBHpSk |
| mongolia                                    | Asia                     | Mongolia                               | sPBHpSk |
| montana                                     | United States of America | Montana                                | sPBHpSk |
| montenegro                                  | Europe                   | Montenegro                             | sPBHpSk |
| morocco                                     | Africa                   | Morocco                                | sPBHpSk |
| mozambique                                  | Africa                   | Mozambique                             | sPBHpSk |
| muenster-regbez                             | Nordrhein-Westfalen      | Regierungsbezirk Münster               | sPBHpSk |
| myanmar                                     | Asia                     | Myanmar (a.k.a. Burma)                 | sPBHpSk |
| namibia                                     | Africa                   | Namibia                                | sPBHpSk |
| nebraska                                    | United States of America | Nebraska                               | sPBHpSk |
| nepal                                       | Asia                     | Nepal                                  | sPBHpSk |
| netherlands                                 | Europe                   | Netherlands                            | sPBHpSk |
| nevada                                      | United States of America | Nevada                                 | sPBHpSk |
| new-brunswick                               | Canada                   | New Brunswick                          | sPBHpSk |
| new-caledonia                               | Australia and Oceania    | New Caledonia                          | sPBHpSk |
| new-hampshire                               | United States of America | New Hampshire                          | sPBHpSk |
| new-jersey                                  | United States of America | New Jersey                             | sPBHpSk |
| new-mexico                                  | United States of America | New Mexico                             | sPBHpSk |
| new-york                                    | United States of America | New York                               | sPBHpSk |
| new-zealand                                 | Australia and Oceania    | New Zealand                            | sPBHpSk |
| newfoundland-and-labrador                   | Canada                   | Newfoundland and Labrador              | sPBHpSk |
| nicaragua                                   | Central America          | Nicaragua                              | sPBHpk  |
| niederbayern                                | Bayern                   | Niederbayern                           | sPBHpSk |
| niedersachsen                               | Germany                  | Niedersachsen                          | sPBHpSk |
| niger                                       | Africa                   | Niger                                  | sPBHpSk |
| nigeria                                     | Africa                   | Nigeria                                | sPBHpSk |
| nord-est                                    | Italy                    | Nord-Est                               | sPBHpSk |
| nord-ovest                                  | Italy                    | Nord-Ovest                             | sPBHpSk |
| nord-pas-de-calais                          | France                   | Nord-Pas-de-Calais                     | sPBHpSk |
| nordrhein-westfalen                         | Germany                  | Nordrhein-Westfalen                    | sPBHpSk |
| norfolk                                     | England                  | Norfolk                                | sPBHpSk |
| north-america                               |                          | North America                          | sPBHpk  |
| north-carolina                              | United States of America | North Carolina                         | sPBHpSk |
| north-caucasus-fed-district                 | Russian Federation       | North Caucasus Federal                 | sPBHpSk |
|                                             |                          | District                               |         |
| north-dakota                                | United States of America | North Dakota                           | sPBHpSk |
| north-korea                                 | Asia                     | North Korea                            | sPBHpSk |
| north-yorkshire                             | England                  | North Yorkshire                        | sPBHpSk |
| northamptonshire                            | England                  | Northamptonshire                       | sPBHpSk |
| northumberland                              | England                  | Northumberland                         | sPBHpSk |
| northwest-territories                       | Canada                   | Northwest Territories                  | sPBHpSk |
| northwestern-fed-district                   | Russian Federation       | Northwestern Federal District          | sPBHpSk |
| norway                                      | Europe                   | Norway                                 | sPBHpSk |
| nottinghamshire                             | England                  | Nottinghamshire                        | sPBHpSk |
| nova-scotia                                 | Canada                   | Nova Scotia                            | sPBHpSk |
| nunavut                                     | Canada                   | Nunavut                                | sPBHpSk |
| oberbayern                                  | Bayern                   | Oberbayern                             | sPBHpSk |
| oberfranken                                 | Bayern                   | Oberfranken                            | sPBHpSk |
| oberpfalz                                   | Bayern                   | Oberpfalz                              | sPBHpSk |
| ohio                                        | United States of America | Ohio                                   | sPBHpSk |
| oklahoma                                    | United States of America | Oklahoma                               | sPBHpSk |
| ontario                                     | Canada                   | Ontario                                | sPBHpSk |
| opolskie                                    | Poland                   | Województwo opolskie(Opole             | sPBHpSk |
|                                             |                          | Voivodeship)                           |         |
| oregon                                      | United States of America | Oregon                                 | sPBHpSk |
| oxfordshire                                 | England                  | Oxfordshire                            | sPBHpSk |
| pakistan                                    | Asia                     | Pakistan                               | sPBHpSk |
| papua-new-guinea                            | Australia and Oceania    | Papua New Guinea                       | sPBHpSk |
| paraguay                                    | South America            | Paraguay                               | sPBHpSk |
| pays-de-la-loire                            | France                   | Pays de la Loire                       | sPBHpSk |
| pennsylvania                                | United States of America | Pennsylvania                           | sPBHpSk |
| peru                                        | South America            | Peru                                   | sPBHpSk |
| philippines                                 | Asia                     | Philippines                            | sPBHpSk |
| picardie                                    | France                   | Picardie                               | sPBHpSk |
| podkarpackie                                | Poland                   | Województwo                            | sPBHpSk |
|                                             |                          | podkarpackie(Subcarpathian             |         |
|                                             |                          | Voivodeship)                           |         |
| podlaskie                                   | Poland                   | Województwo                            | sPBHpSk |
|                                             |                          | podlaskie(Podlaskie                    |         |
|                                             |                          | Voivodeship)                           |         |
| poitou-charentes                            | France                   | Poitou-Charentes                       | sPBHpSk |
| poland                                      | Europe                   | Poland                                 | sPBHpk  |
| pomorskie                                   | Poland                   | Województwo                            | sPBHpSk |
|                                             |                          | pomorskie(Pomeranian                   |         |
|                                             |                          | Voivodeship)                           |         |
| portugal                                    | Europe                   | Portugal                               | sPBHpSk |
| prince-edward-island                        | Canada                   | Prince Edward Island                   | sPBHpSk |
| provence-alpes-cote-d-azur                  | France                   | Provence Alpes-Cote-d'Azur             | sPBHpSk |
| puerto-rico                                 | United States of America | Puerto Rico                            | sPBHpSk |
| quebec                                      | Canada                   | Quebec                                 | sPBHpSk |
| reunion                                     | France                   | Reunion                                | sPBHpk  |
| rheinland-pfalz                             | Germany                  | Rheinland-Pfalz                        | sPBHpSk |
| rhode-island                                | United States of America | Rhode Island                           | sPBHpSk |
| rhone-alpes                                 | France                   | Rhone-Alpes                            | sPBHpSk |
| romania                                     | Europe                   | Romania                                | sPBHpSk |
| russia                                      |                          | Russian Federation                     | sPBHpk  |
| rutland                                     | England                  | Rutland                                | sPBHpSk |
| rwanda                                      | Africa                   | Rwanda                                 | sPBHpSk |
| saarland                                    | Germany                  | Saarland                               | sPBHpSk |
| sachsen                                     | Germany                  | Sachsen                                | sPBHpSk |
| sachsen-anhalt                              | Germany                  | Sachsen-Anhalt                         | sPBHpSk |
| saint-helena-ascension-and-tristan-da-cunha | Africa                   | Saint Helena, Ascension, and           | sPBHpSk |
|                                             |                          | Tristan da Cunha                       |         |
| sao-tome-and-principe                       | Africa                   | Sao Tome and Principe                  | sPBHpSk |
| saskatchewan                                | Canada                   | Saskatchewan                           | sPBHpSk |
| schleswig-holstein                          | Germany                  | Schleswig-Holstein                     | sPBHpSk |
| schwaben                                    | Bayern                   | Schwaben                               | sPBHpSk |
| scotland                                    | Great Britain            | Scotland                               | sPBHpSk |
| senegal-and-gambia                          | Africa                   | Senegal and Gambia                     | sPBHpSk |
| serbia                                      | Europe                   | Serbia                                 | sPBHpSk |
| seychelles                                  | Africa                   | Seychelles                             | sPBHpSk |
| shikoku                                     | Japan                    | Shikoku                                | sPBHpSk |
| shropshire                                  | England                  | Shropshire                             | sPBHpSk |
| siberian-fed-district                       | Russian Federation       | Siberian Federal District              | sPBHpSk |
| sierra-leone                                | Africa                   | Sierra Leone                           | sPBHpSk |
| slaskie                                     | Poland                   | Województwo śląskie(Silesian           | sPBHpSk |
|                                             |                          | Voivodeship)                           |         |
| slovakia                                    | Europe                   | Slovakia                               | sPBHpSk |
| slovenia                                    | Europe                   | Slovenia                               | sPBHpSk |
| somalia                                     | Africa                   | Somalia                                | sPBHpSk |
| somerset                                    | England                  | Somerset                               | sPBHpSk |
| south-africa                                | Africa                   | South Africa                           | sPBHpSk |
| south-africa-and-lesotho                    | Africa                   | South Africa (includes                 | sPBHpk  |
|                                             |                          | Lesotho)                               |         |
| south-america                               |                          | South America                          | sPBHpk  |
| south-carolina                              | United States of America | South Carolina                         | sPBHpSk |
| south-dakota                                | United States of America | South Dakota                           | sPBHpSk |
| south-fed-district                          | Russian Federation       | South Federal District                 | sPBHpSk |
| south-korea                                 | Asia                     | South Korea                            | sPBHpSk |
| south-sudan                                 | Africa                   | South Sudan                            | sPBHpSk |
| south-yorkshire                             | England                  | South Yorkshire                        | sPBHpSk |
| spain                                       | Europe                   | Spain                                  | sPBHpSk |
| sri-lanka                                   | Asia                     | Sri Lanka                              | sPBHpSk |
| staffordshire                               | England                  | Staffordshire                          | sPBHpSk |
| stuttgart-regbez                            | Baden-Württemberg        | Regierungsbezirk Stuttgart             | sPBHpSk |
| sud                                         | Italy                    | Sud                                    | sPBHpSk |
| sudan                                       | Africa                   | Sudan                                  | sPBHpSk |
| suffolk                                     | England                  | Suffolk                                | sPBHpSk |
| suriname                                    | South America            | suriname                               | sPBHpSk |
| surrey                                      | England                  | Surrey                                 | sPBHpSk |
| swaziland                                   | Africa                   | Swaziland                              | sPBHpSk |
| sweden                                      | Europe                   | Sweden                                 | sPBHpSk |
| swietokrzyskie                              | Poland                   | Województwo                            | sPBHpSk |
|                                             |                          | świętokrzyskie(Świętokrzyskie          |         |
|                                             |                          | Voivodeship)                           |         |
| switzerland                                 | Europe                   | Switzerland                            | sPBHpSk |
| syria                                       | Asia                     | Syria                                  | sPBHpSk |
| taiwan                                      | Asia                     | Taiwan                                 | sPBHpSk |
| tajikistan                                  | Asia                     | Tajikistan                             | sPBHpk  |
| tanzania                                    | Africa                   | Tanzania                               | sPBHpSk |
| tennessee                                   | United States of America | Tennessee                              | sPBHpSk |
| texas                                       | United States of America | Texas                                  | sPBHpSk |
| thailand                                    | Asia                     | Thailand                               | sPBHpSk |
| thueringen                                  | Germany                  | Thüringen                              | sPBHpSk |
| togo                                        | Africa                   | Togo                                   | sPBHpSk |
| tohoku                                      | Japan                    | Tōhoku region                          | sPBHpSk |
| tuebingen-regbez                            | Baden-Württemberg        | Regierungsbezirk Tübingen              | sPBHpSk |
| tunisia                                     | Africa                   | Tunisia                                | sPBHpSk |
| turkey                                      | Europe                   | Turkey                                 | sPBHpSk |
| turkmenistan                                | Asia                     | Turkmenistan                           | sPBHpSk |
| uganda                                      | Africa                   | Uganda                                 | sPBHpSk |
| ukraine                                     | Europe                   | Ukraine                                | sPBHpSk |
| unterfranken                                | Bayern                   | Unterfranken                           | sPBHpSk |
| ural-fed-district                           | Russian Federation       | Ural Federal District                  | sPBHpSk |
| uruguay                                     | South America            | Uruguay                                | sPBHpSk |
| us                                          | North America            | United States of America               |         |
| us-midwest                                  | North America            | US Midwest                             | sPBHpk  |
| us-northeast                                | North America            | US Northeast                           | sPBHpk  |
| us-pacific                                  | North America            | US Pacific                             | sPBHpk  |
| us-south                                    | North America            | US South                               | sPBHpk  |
| us-west                                     | North America            | US West                                | sPBHpk  |
| utah                                        | United States of America | Utah                                   | sPBHpSk |
| uzbekistan                                  | Asia                     | Uzbekistan                             | sPBHpSk |
| vermont                                     | United States of America | Vermont                                | sPBHpSk |
| vietnam                                     | Asia                     | Vietnam                                | sPBHpSk |
| virginia                                    | United States of America | Virginia                               | sPBHpSk |
| volga-fed-district                          | Russian Federation       | Volga Federal District                 | sPBHpSk |
| wales                                       | Great Britain            | Wales                                  | sPBHpSk |
| warminsko-mazurskie                         | Poland                   | Województwo                            | sPBHpSk |
|                                             |                          | warmińsko-mazurskie(Warmian-Masurian   |         |
|                                             |                          | Voivodeship)                           |         |
| warwickshire                                | England                  | Warwickshire                           | sPBHpSk |
| washington                                  | United States of America | Washington                             | sPBHpSk |
| west-midlands                               | England                  | West Midlands                          | sPBHpSk |
| west-sussex                                 | England                  | West Sussex                            | sPBHpSk |
| west-virginia                               | United States of America | West Virginia                          | sPBHpSk |
| west-yorkshire                              | England                  | West Yorkshire                         | sPBHpSk |
| wielkopolskie                               | Poland                   | Województwo                            | sPBHpSk |
|                                             |                          | wielkopolskie(Greater Poland           |         |
|                                             |                          | Voivodeship)                           |         |
| wiltshire                                   | England                  | Wiltshire                              | sPBHpSk |
| wisconsin                                   | United States of America | Wisconsin                              | sPBHpSk |
| worcestershire                              | England                  | Worcestershire                         | sPBHpSk |
| wyoming                                     | United States of America | Wyoming                                | sPBHpSk |
| yemen                                       | Asia                     | Yemen                                  | sPBHpSk |
| yukon                                       | Canada                   | Yukon                                  | sPBHpSk |
| zachodniopomorskie                          | Poland                   | Województwo                            | sPBHpSk |
|                                             |                          | zachodniopomorskie(West                |         |
|                                             |                          | Pomeranian Voivodeship)                |         |
| zambia                                      | Africa                   | Zambia                                 | sPBHpSk |
| zimbabwe                                    | Africa                   | Zimbabwe                               | sPBHpSk |
Total elements: 393

## List of elements from openstreetmap.fr
|                SHORTNAME                |        IS IN        |                LONG NAME                | FORMATS |
|-----------------------------------------|---------------------|-----------------------------------------|---------|
| afghanistan                             | asia                | afghanistan                             | sPp     |
| africa                                  |                     | africa                                  | sPp     |
| algeria                                 | africa              | algeria                                 | sPp     |
| american_samoa                          | south-america       | american_samoa                          | sPp     |
| andalucia                               | spain               | andalucia                               | sPp     |
| andaman_and_nicobar_islands             | india               | andaman_and_nicobar_islands             | sPp     |
| andhra_pradesh                          | india               | andhra_pradesh                          | sPp     |
| angola                                  | africa              | angola                                  | sPp     |
| anguilla                                | central-america     | anguilla                                | sPp     |
| anhui                                   | china               | anhui                                   | sPp     |
| antigua_and_barbuda                     | central-america     | antigua_and_barbuda                     | sPp     |
| aragon                                  | spain               | aragon                                  | sPp     |
| armenia                                 | asia                | armenia                                 | sPp     |
| arnsberg                                | nordrhein_westfalen | arnsberg                                | sPp     |
| aruba                                   | central-america     | aruba                                   | sPp     |
| arunachal_pradesh                       | india               | arunachal_pradesh                       | sPp     |
| asia                                    |                     | asia                                    | sPp     |
| assam                                   | india               | assam                                   | sPp     |
| asturias                                | spain               | asturias                                | sPp     |
| australia                               | oceania             | australia                               | sPp     |
| australian_capital_territory            | australia           | australian_capital_territory            | sPp     |
| austria                                 | europe              | austria                                 | sPp     |
| bahamas                                 | central-america     | bahamas                                 | sPp     |
| bahrain                                 | asia                | bahrain                                 | sPp     |
| banskobystricky                         | slovakia            | banskobystricky                         | sPp     |
| barbados                                | central-america     | barbados                                | sPp     |
| beijing                                 | china               | beijing                                 | sPp     |
| belgium                                 | europe              | belgium                                 | sPp     |
| benin                                   | africa              | benin                                   | sPp     |
| bermuda                                 | north-america       | bermuda                                 | sPp     |
| bhutan                                  | asia                | bhutan                                  | sPp     |
| bihar                                   | india               | bihar                                   | sPp     |
| bouvet_island                           | africa              | bouvet_island                           | sPp     |
| bratislavsky                            | slovakia            | bratislavsky                            | sPp     |
| british_indian_ocean_territory          | asia                | british_indian_ocean_territory          | sPp     |
| british_virgin_islands                  | central-america     | british_virgin_islands                  | sPp     |
| brunei                                  | asia                | brunei                                  | sPp     |
| brussels_capital_region                 | belgium             | brussels_capital_region                 | sPp     |
| burgenland                              | austria             | burgenland                              | sPp     |
| burkina_faso                            | africa              | burkina_faso                            | sPp     |
| burundi                                 | africa              | burundi                                 | sPp     |
| cambodia                                | asia                | cambodia                                | sPp     |
| cameroon                                | africa              | cameroon                                | sPp     |
| canada                                  | north-america       | canada                                  | p       |
| canarias                                | spain               | canarias                                | sPp     |
| cantabria                               | spain               | cantabria                               | sPp     |
| cape_verde                              | africa              | cape_verde                              | sPp     |
| caribbean                               | central-america     | caribbean                               | sPp     |
| castilla_la_mancha                      | spain               | castilla_la_mancha                      | sPp     |
| castilla_y_leon                         | spain               | castilla_y_leon                         | sPp     |
| catalunya                               | spain               | catalunya                               | sPp     |
| cayman_islands                          | central-america     | cayman_islands                          | sPp     |
| central-america                         |                     | central-america                         | sPp     |
| central_african_republic                | africa              | central_african_republic                | sPp     |
| ceuta                                   | spain               | ceuta                                   | sPp     |
| chad                                    | africa              | chad                                    | sPp     |
| chandigarh                              | india               | chandigarh                              | sPp     |
| chhattisgarh                            | india               | chhattisgarh                            | sPp     |
| china                                   | asia                | china                                   | sPp     |
| chongqing                               | china               | chongqing                               | sPp     |
| christmas_island                        | australia           | christmas_island                        | sPp     |
| chubu                                   | japan               | chubu                                   | sPp     |
| chugoku                                 | japan               | chugoku                                 | sPp     |
| cocos_islands                           | australia           | cocos_islands                           | sPp     |
| comoros                                 | africa              | comoros                                 | sPp     |
| comunidad_de_madrid                     | spain               | comunidad_de_madrid                     | sPp     |
| comunidad_foral_de_navarra              | spain               | comunidad_foral_de_navarra              | sPp     |
| comunitat_valenciana                    | spain               | comunitat_valenciana                    | sPp     |
| congo_brazzaville                       | africa              | congo_brazzaville                       | sPp     |
| congo_kinshasa                          | africa              | congo_kinshasa                          | sPp     |
| cook_islands                            | south-america       | cook_islands                            | sPp     |
| coral_sea_islands                       | australia           | coral_sea_islands                       | sPp     |
| costa_rica                              | central-america     | costa_rica                              | sPp     |
| curacao                                 | central-america     | curacao                                 | sPp     |
| czech_republic                          | europe              | czech_republic                          | sPp     |
| dadra_and_nagar_haveli                  | india               | dadra_and_nagar_haveli                  | sPp     |
| daman_and_diu                           | india               | daman_and_diu                           | sPp     |
| detmold                                 | nordrhein_westfalen | detmold                                 | sPp     |
| djibouti                                | africa              | djibouti                                | sPp     |
| dolnoslaskie                            | poland              | dolnoslaskie                            | sPp     |
| dominica                                | central-america     | dominica                                | sPp     |
| dominican_republic                      | central-america     | dominican_republic                      | sPp     |
| drenthe                                 | netherlands         | drenthe                                 | sPp     |
| dusseldorf                              | nordrhein_westfalen | dusseldorf                              | sPp     |
| east                                    | england             | east                                    | sPp     |
| east_midlands                           | england             | east_midlands                           | sPp     |
| east_timor                              | asia                | east_timor                              | sPp     |
| el_salvador                             | central-america     | el_salvador                             | sPp     |
| england                                 | united_kingdom      | england                                 | sPp     |
| equatorial_guinea                       | africa              | equatorial_guinea                       | sPp     |
| eritrea                                 | africa              | eritrea                                 | sPp     |
| europe                                  |                     | europe                                  | sPp     |
| euskadi                                 | spain               | euskadi                                 | sPp     |
| extremadura                             | spain               | extremadura                             | sPp     |
| falkland                                | south-america       | falkland                                | sPp     |
| fiji                                    |                     | fiji                                    | sP      |
| fiji_east                               | south-america       | fiji_east                               | sPp     |
| fiji_west                               | oceania             | fiji_west                               | sPp     |
| flanders                                | belgium             | flanders                                | sPp     |
| flevoland                               | netherlands         | flevoland                               | sPp     |
| france                                  | europe              | france                                  | sPp     |
| france_metro_dom_com_nc                 |                     | france_metro_dom_com_nc                 | sP      |
| france_taaf                             | africa              | france_taaf                             | sPp     |
| friesland                               | netherlands         | friesland                               | sPp     |
| fujian                                  | china               | fujian                                  | sPp     |
| gabon                                   | africa              | gabon                                   | sPp     |
| galicia                                 | spain               | galicia                                 | sPp     |
| gambia                                  | africa              | gambia                                  | sPp     |
| gansu                                   | china               | gansu                                   | sPp     |
| gelderland                              | netherlands         | gelderland                              | sPp     |
| georgia                                 | asia                | georgia                                 | sPp     |
| germany                                 | europe              | germany                                 | sPp     |
| ghana                                   | africa              | ghana                                   | sPp     |
| gibraltar                               | europe              | gibraltar                               | sPp     |
| goa                                     | india               | goa                                     | sPp     |
| greater_london                          | england             | greater_london                          | sPp     |
| grenada                                 | central-america     | grenada                                 | sPp     |
| groningen                               | netherlands         | groningen                               | sPp     |
| guadeloupe                              | central-america     | guadeloupe                              | sPp     |
| guam                                    | oceania             | guam                                    | sPp     |
| guangdong                               | china               | guangdong                               | sPp     |
| guangxi                                 | china               | guangxi                                 | sPp     |
| guernesey                               | europe              | guernesey                               | sPp     |
| guinea                                  | africa              | guinea                                  | sPp     |
| guizhou                                 | china               | guizhou                                 | sPp     |
| gujarat                                 | india               | gujarat                                 | sPp     |
| guyana                                  | south-america       | guyana                                  | sPp     |
| guyane                                  | south-america       | guyane                                  | sPp     |
| hainan                                  | china               | hainan                                  | sPp     |
| haryana                                 | india               | haryana                                 | sPp     |
| hebei                                   | china               | hebei                                   | sPp     |
| heilongjiang                            | china               | heilongjiang                            | sPp     |
| henan                                   | china               | henan                                   | sPp     |
| himachal_pradesh                        | india               | himachal_pradesh                        | sPp     |
| hokkaido                                | japan               | hokkaido                                | sPp     |
| honduras                                | central-america     | honduras                                | sPp     |
| hong_kong                               | china               | hong_kong                               | sPp     |
| hubei                                   | china               | hubei                                   | sPp     |
| hunan                                   | china               | hunan                                   | sPp     |
| illes_balears                           | spain               | illes_balears                           | sPp     |
| india                                   | asia                | india                                   | sPp     |
| inner_mongolia                          | china               | inner_mongolia                          | sPp     |
| ireland                                 | europe              | ireland                                 | sPp     |
| israel                                  | asia                | israel                                  | sPp     |
| israel_and_palestine                    | asia                | israel_and_palestine                    | sPp     |
| ivory_coast                             | africa              | ivory_coast                             | sPp     |
| jamaica                                 | central-america     | jamaica                                 | sPp     |
| jammu_and_kashmir                       | india               | jammu_and_kashmir                       | sPp     |
| japan                                   | asia                | japan                                   | sPp     |
| jersey                                  | europe              | jersey                                  | sPp     |
| jharkhand                               | india               | jharkhand                               | sPp     |
| jiangsu                                 | china               | jiangsu                                 | sPp     |
| jiangxi                                 | china               | jiangxi                                 | sPp     |
| jihocesky                               | czech_republic      | jihocesky                               | sPp     |
| jihomoravsky                            | czech_republic      | jihomoravsky                            | sPp     |
| jilin                                   | china               | jilin                                   | sPp     |
| kansai                                  | japan               | kansai                                  | sPp     |
| kanto                                   | japan               | kanto                                   | sPp     |
| karlovarsky                             | czech_republic      | karlovarsky                             | sPp     |
| karnataka                               | india               | karnataka                               | sPp     |
| karnten                                 | austria             | karnten                                 | sPp     |
| kenya                                   | africa              | kenya                                   | sPp     |
| kerala                                  | india               | kerala                                  | sPp     |
| kiribati                                |                     | kiribati                                | sP      |
| kiribati_east                           | south-america       | kiribati_east                           | sPp     |
| kiribati_west                           | oceania             | kiribati_west                           | sPp     |
| koln                                    | nordrhein_westfalen | koln                                    | sPp     |
| kosicky                                 | slovakia            | kosicky                                 | sPp     |
| kralovehradecky                         | czech_republic      | kralovehradecky                         | sPp     |
| kujawsko_pomorskie                      | poland              | kujawsko_pomorskie                      | sPp     |
| kuwait                                  | asia                | kuwait                                  | sPp     |
| kyushu                                  | japan               | kyushu                                  | sPp     |
| la_rioja                                | spain               | la_rioja                                | sPp     |
| lakshadweep                             | india               | lakshadweep                             | sPp     |
| laos                                    | asia                | laos                                    | sPp     |
| lesotho                                 | africa              | lesotho                                 | sPp     |
| liaoning                                | china               | liaoning                                | sPp     |
| liberecky                               | czech_republic      | liberecky                               | sPp     |
| limburg                                 | netherlands         | limburg                                 | sPp     |
| lodzkie                                 | poland              | lodzkie                                 | sPp     |
| lubelskie                               | poland              | lubelskie                               | sPp     |
| lubuskie                                | poland              | lubuskie                                | sPp     |
| macau                                   | china               | macau                                   | sPp     |
| madhya_pradesh                          | india               | madhya_pradesh                          | sPp     |
| maharashtra                             | india               | maharashtra                             | sPp     |
| malawi                                  | africa              | malawi                                  | sPp     |
| malaysia                                | asia                | malaysia                                | sPp     |
| maldives                                | asia                | maldives                                | sPp     |
| mali                                    | africa              | mali                                    | sPp     |
| malopolskie                             | poland              | malopolskie                             | sPp     |
| manipur                                 | india               | manipur                                 | sPp     |
| marshall-islands                        | south-america       | marshall-islands                        | sPp     |
| marshall_islands                        | oceania             | marshall_islands                        | sPp     |
| martinique                              | central-america     | martinique                              | sPp     |
| mauritania                              | africa              | mauritania                              | sPp     |
| mauritius                               | africa              | mauritius                               | sPp     |
| mayotte                                 | africa              | mayotte                                 | sPp     |
| mazowieckie                             | poland              | mazowieckie                             | sPp     |
| meghalaya                               | india               | meghalaya                               | sPp     |
| melilla                                 | spain               | melilla                                 | sPp     |
| micronesia                              | oceania             | micronesia                              | sPp     |
| mizoram                                 | india               | mizoram                                 | sPp     |
| monaco                                  | europe              | monaco                                  | sPp     |
| montserrat                              | central-america     | montserrat                              | sPp     |
| moravskoslezsky                         | czech_republic      | moravskoslezsky                         | sPp     |
| mozambique                              | africa              | mozambique                              | sPp     |
| munster                                 | nordrhein_westfalen | munster                                 | sPp     |
| myanmar                                 | asia                | myanmar                                 | sPp     |
| nagaland                                | india               | nagaland                                | sPp     |
| namibia                                 | africa              | namibia                                 | sPp     |
| national_capital_territory_of_delhi     | india               | national_capital_territory_of_delhi     | sPp     |
| nauri                                   | south-america       | nauri                                   | sPp     |
| nauru                                   | oceania             | nauru                                   | sPp     |
| netherlands                             | europe              | netherlands                             | sPp     |
| new_caledonia                           | oceania             | new_caledonia                           | sPp     |
| new_south_wales                         | australia           | new_south_wales                         | sPp     |
| nicaragua                               | central-america     | nicaragua                               | sPp     |
| niederosterreich                        | austria             | niederosterreich                        | sPp     |
| niger                                   | africa              | niger                                   | sPp     |
| ningxia                                 | china               | ningxia                                 | sPp     |
| nitriansky                              | slovakia            | nitriansky                              | sPp     |
| niue                                    | south-america       | niue                                    | sPp     |
| noord_brabant                           | netherlands         | noord_brabant                           | sPp     |
| noord_holland                           | netherlands         | noord_holland                           | sPp     |
| nordrhein_westfalen                     | germany             | nordrhein_westfalen                     | sPp     |
| norfolk_island                          | australia           | norfolk_island                          | sPp     |
| north-america                           |                     | north-america                           | p       |
| north_east                              | england             | north_east                              | sPp     |
| north_west                              | england             | north_west                              | sPp     |
| northern_ireland                        | united_kingdom      | northern_ireland                        | sPp     |
| northern_mariana_islands                | oceania             | northern_mariana_islands                | sPp     |
| northern_territory                      | australia           | northern_territory                      | sPp     |
| oberosterreich                          | austria             | oberosterreich                          | sPp     |
| oceania                                 |                     | oceania                                 | sPp     |
| odisha                                  | india               | odisha                                  | sPp     |
| olomoucky                               | czech_republic      | olomoucky                               | sPp     |
| oman                                    | asia                | oman                                    | sPp     |
| opolskie                                | poland              | opolskie                                | sPp     |
| overijssel                              | netherlands         | overijssel                              | sPp     |
| palau                                   | oceania             | palau                                   | sPp     |
| palestine                               | asia                | palestine                               | sPp     |
| panama                                  | central-america     | panama                                  | sPp     |
| papua_new_guinea                        | oceania             | papua_new_guinea                        | sPp     |
| paraguay                                | south-america       | paraguay                                | sPp     |
| pardubicky                              | czech_republic      | pardubicky                              | sPp     |
| pitcairn                                | south-america       | pitcairn                                | sPp     |
| plzensky                                | czech_republic      | plzensky                                | sPp     |
| podkarpackie                            | poland              | podkarpackie                            | sPp     |
| podlaskie                               | poland              | podlaskie                               | sPp     |
| poland                                  | europe              | poland                                  | sPp     |
| polynesie                               | oceania             | polynesie                               | sPp     |
| pomorskie                               | poland              | pomorskie                               | sPp     |
| praha                                   | czech_republic      | praha                                   | sPp     |
| presovsky                               | slovakia            | presovsky                               | sPp     |
| puducherry                              | india               | puducherry                              | sPp     |
| puerto_rico                             | central-america     | puerto_rico                             | sPp     |
| punjab                                  | india               | punjab                                  | sPp     |
| qatar                                   | asia                | qatar                                   | sPp     |
| qinghai                                 | china               | qinghai                                 | sPp     |
| quebec                                  | canada              | quebec                                  | sPp     |
| queensland                              | australia           | queensland                              | sPp     |
| rajasthan                               | india               | rajasthan                               | sPp     |
| region_de_murcia                        | spain               | region_de_murcia                        | sPp     |
| reunion                                 | africa              | reunion                                 | sPp     |
| rwanda                                  | africa              | rwanda                                  | sPp     |
| saint_barthelemy                        | central-america     | saint_barthelemy                        | sPp     |
| saint_helena_ascension_tristan_da_cunha | africa              | saint_helena_ascension_tristan_da_cunha | sPp     |
| saint_kitts_and_nevis                   | central-america     | saint_kitts_and_nevis                   | sPp     |
| saint_lucia                             | central-america     | saint_lucia                             | sPp     |
| saint_martin                            | central-america     | saint_martin                            | sPp     |
| saint_pierre_et_miquelon                | north-america       | saint_pierre_et_miquelon                | sPp     |
| saint_vincent_and_the_grenadines        | central-america     | saint_vincent_and_the_grenadines        | sPp     |
| salzburg                                | austria             | salzburg                                | sPp     |
| samoa                                   | south-america       | samoa                                   | sPp     |
| sao_tome_and_principe                   | africa              | sao_tome_and_principe                   | sPp     |
| saudi_arabia                            | asia                | saudi_arabia                            | sPp     |
| senegal                                 | africa              | senegal                                 | sPp     |
| seychelles                              | africa              | seychelles                              | sPp     |
| shaanxi                                 | china               | shaanxi                                 | sPp     |
| shandong                                | china               | shandong                                | sPp     |
| shanghai                                | china               | shanghai                                | sPp     |
| shanxi                                  | china               | shanxi                                  | sPp     |
| shikoku                                 | japan               | shikoku                                 | sPp     |
| sichuan                                 | china               | sichuan                                 | sPp     |
| sikkim                                  | india               | sikkim                                  | sPp     |
| singapore                               | asia                | singapore                               | sPp     |
| sint_maarten                            | central-america     | sint_maarten                            | sPp     |
| slaskie                                 | poland              | slaskie                                 | sPp     |
| slovakia                                | europe              | slovakia                                | sPp     |
| solomon_islands                         | oceania             | solomon_islands                         | sPp     |
| south-america                           |                     | south-america                           | sPp     |
| south_africa                            | africa              | south_africa                            | sPp     |
| south_australia                         | australia           | south_australia                         | sPp     |
| south_east                              | england             | south_east                              | sPp     |
| south_georgia_and_south_sandwich        | south-america       | south_georgia_and_south_sandwich        | sPp     |
| south_sudan                             | africa              | south_sudan                             | sPp     |
| south_west                              | england             | south_west                              | sPp     |
| spain                                   | europe              | spain                                   | sPp     |
| steiermark                              | austria             | steiermark                              | sPp     |
| stredocesky                             | czech_republic      | stredocesky                             | sPp     |
| sudan                                   | africa              | sudan                                   | sPp     |
| suriname                                | south-america       | suriname                                | sPp     |
| swaziland                               | africa              | swaziland                               | sPp     |
| swietokrzyskie                          | poland              | swietokrzyskie                          | sPp     |
| taaf                                    | south-america       | taaf                                    | sPp     |
| tamil_nadu                              | india               | tamil_nadu                              | sPp     |
| tasmania                                | australia           | tasmania                                | sPp     |
| telangana                               | india               | telangana                               | sPp     |
| tianjin                                 | china               | tianjin                                 | sPp     |
| tibet                                   | china               | tibet                                   | sPp     |
| tirol                                   | austria             | tirol                                   | sPp     |
| togo                                    | africa              | togo                                    | sPp     |
| tohoku                                  | japan               | tohoku                                  | sPp     |
| tokelau                                 | south-america       | tokelau                                 | sPp     |
| tonga                                   | south-america       | tonga                                   | sPp     |
| trenciansky                             | slovakia            | trenciansky                             | sPp     |
| trinidad_and_tobago                     | south-america       | trinidad_and_tobago                     | sPp     |
| tripura                                 | india               | tripura                                 | sPp     |
| trnavsky                                | slovakia            | trnavsky                                | sPp     |
| tunisia                                 | africa              | tunisia                                 | sPp     |
| turks_and_caicos_islands                | central-america     | turks_and_caicos_islands                | sPp     |
| tuvalu                                  | oceania             | tuvalu                                  | sPp     |
| uganda                                  | africa              | uganda                                  | sPp     |
| united_arab_emirates                    | asia                | united_arab_emirates                    | sPp     |
| united_kingdom                          | europe              | united_kingdom                          | sPp     |
| united_states_virgin_islands            | central-america     | united_states_virgin_islands            | sPp     |
| usa_virgin_islands                      | central-america     | usa_virgin_islands                      | sPp     |
| ustecky                                 | czech_republic      | ustecky                                 | sPp     |
| utrecht                                 | netherlands         | utrecht                                 | sPp     |
| uttar_pradesh                           | india               | uttar_pradesh                           | sPp     |
| uttarakhand                             | india               | uttarakhand                             | sPp     |
| vanuatu                                 | oceania             | vanuatu                                 | sPp     |
| venezuela                               | south-america       | venezuela                               | sPp     |
| victoria                                | australia           | victoria                                | sPp     |
| vorarlberg                              | austria             | vorarlberg                              | sPp     |
| vysocina                                | czech_republic      | vysocina                                | sPp     |
| wallis_et_futuna                        | oceania             | wallis_et_futuna                        | sPp     |
| wallonia_french_community               | belgium             | wallonia_french_community               | sPp     |
| wallonia_german_community               | belgium             | wallonia_german_community               | sPp     |
| warminsko_mazurskie                     | poland              | warminsko_mazurskie                     | sPp     |
| west_bengal                             | india               | west_bengal                             | sPp     |
| west_midlands                           | england             | west_midlands                           | sPp     |
| western_australia                       | australia           | western_australia                       | sPp     |
| western_sahara                          | africa              | western_sahara                          | sPp     |
| wielkopolskie                           | poland              | wielkopolskie                           | sPp     |
| wien                                    | austria             | wien                                    | sPp     |
| xinjiang                                | china               | xinjiang                                | sPp     |
| yorkshire_and_the_humber                | england             | yorkshire_and_the_humber                | sPp     |
| yunnan                                  | china               | yunnan                                  | sPp     |
| zachodniopomorskie                      | poland              | zachodniopomorskie                      | sPp     |
| zambia                                  | africa              | zambia                                  | sPp     |
| zeeland                                 | netherlands         | zeeland                                 | sPp     |
| zhejiang                                | china               | zhejiang                                | sPp     |
| zilinsky                                | slovakia            | zilinsky                                | sPp     |
| zimbabwe                                | africa              | zimbabwe                                | sPp     |
| zlinsky                                 | czech_republic      | zlinsky                                 | sPp     |
| zuid_holland                            | netherlands         | zuid_holland                            | sPp     |
Total elements: 357
