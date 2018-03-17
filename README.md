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
      --help            Show context-sensitive help (also try --help-long and
                        --help-man).
  -c, --config="./geofabrik.yml"  
                        Set Config file.
  -n, --nodownload      Do not download file (test only)
  -v, --verbose         Be verbose
  -q, --quiet           Be quiet
      --proxy-http=""   Use http proxy, format: proxy_address:port
      --proxy-sock5=""  Use Sock5 proxy, format: proxy_address:port
      --proxy-user=""   Proxy user
      --proxy-pass=""   Proxy password

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
    -H, --osh.pbf  Download osh.pbf (default)
    -s, --state    Download state.txt file
    -p, --poly     Download poly file
    -k, --kml      Download kml file

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
| lodzkie                                     | Poland                   | Województwo łódzkie(Łódź Voivodeship)  | sPBHpSk |
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
| malaysia-singapore-brunei                   | Asia                     | Malaysia, Singapore, and Brunei        | sPBHpk  |
| maldives                                    | Asia                     | Maldives                               | sPBHpSk |
| mali                                        | Africa                   | Mali                                   | sPBHpSk |
| malopolskie                                 | Poland                   | Województwo małopolskie(Lesser Poland  | sPBHpSk |
|                                             |                          | Voivodeship)                           |         |
| malta                                       | Europe                   | Malta                                  | sPBHpSk |
| manitoba                                    | Canada                   | Manitoba                               | sPBHpSk |
| martinique                                  | France                   | Martinique                             | sPBHpk  |
| maryland                                    | United States of America | Maryland                               | sPBHpSk |
| massachusetts                               | United States of America | Massachusetts                          | sPBHpSk |
| mauritania                                  | Africa                   | Mauritania                             | sPBHpSk |
| mauritius                                   | Africa                   | Mauritius                              | sPBHpSk |
| mayotte                                     | France                   | Mayotte                                | sPBHpk  |
| mazowieckie                                 | Poland                   | Województwo mazowieckie(Mazovian       | sPBHpSk |
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
| north-caucasus-fed-district                 | Russian Federation       | North Caucasus Federal District        | sPBHpSk |
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
| podkarpackie                                | Poland                   | Województwo podkarpackie(Subcarpathian | sPBHpSk |
|                                             |                          | Voivodeship)                           |         |
| podlaskie                                   | Poland                   | Województwo podlaskie(Podlaskie        | sPBHpSk |
|                                             |                          | Voivodeship)                           |         |
| poitou-charentes                            | France                   | Poitou-Charentes                       | sPBHpSk |
| poland                                      | Europe                   | Poland                                 | sPBHpk  |
| pomorskie                                   | Poland                   | Województwo pomorskie(Pomeranian       | sPBHpSk |
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
| saint-helena-ascension-and-tristan-da-cunha | Africa                   | Saint Helena, Ascension, and Tristan   | sPBHpSk |
|                                             |                          | da Cunha                               |         |
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
| south-africa-and-lesotho                    | Africa                   | South Africa (includes Lesotho)        | sPBHpk  |
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
| wielkopolskie                               | Poland                   | Województwo wielkopolskie(Greater      | sPBHpSk |
|                                             |                          | Poland Voivodeship)                    |         |
| wiltshire                                   | England                  | Wiltshire                              | sPBHpSk |
| wisconsin                                   | United States of America | Wisconsin                              | sPBHpSk |
| worcestershire                              | England                  | Worcestershire                         | sPBHpSk |
| wyoming                                     | United States of America | Wyoming                                | sPBHpSk |
| yemen                                       | Asia                     | Yemen                                  | sPBHpSk |
| yukon                                       | Canada                   | Yukon                                  | sPBHpSk |
| zachodniopomorskie                          | Poland                   | Województwo zachodniopomorskie(West    | sPBHpSk |
|                                             |                          | Pomeranian Voivodeship)                |         |
| zambia                                      | Africa                   | Zambia                                 | sPBHpSk |
| zimbabwe                                    | Africa                   | Zimbabwe                               | sPBHpSk |
Total elements: 393
