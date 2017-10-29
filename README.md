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
      --help        Show context-sensitive help (also try --help-long and
                    --help-man).
  -c, --config="./geofabrik.yml"
                    Set Config file.
  -n, --nodownload  Do not download file (test only)
  -v, --verbose     Be verbose

Commands:
  help [<command>...]
    Show help.


  update [<flags>]
    Update geofabrik.yml from github

    --url="https://raw.githubusercontent.com/julien-noblet/download-geofabrik/stable/geofabrik.yml"
      Url for config source

  list
    Show elements available


  download [<flags>] <element>
    Download element

    -B, --osm.bz2  Download osm.bz2 if available
    -S, --shp.zip  Download shp.zip if available
    -P, --osm.pbf  Download osm.pbf (default)
    -H, --osh.pbf  Download osh.pbf (default)
    -s, --state    Download state.txt file
    -p, --poly     Download poly file
```

## List of elements
|                  SHORTNAME                  |         IS IN         |               LONG NAME                | FORMATS |
|---------------------------------------------|-----------------------|----------------------------------------|---------|
| afghanistan                                 | Asia                  | Afghanistan                            | sPBHpS  |
| africa                                      |                       | Africa                                 | sPBHp   |
| alabama                                     | North America         | Alabama                                | sPBHpS  |
| alaska                                      | North America         | Alaska                                 | sPBHpS  |
| albania                                     | Europe                | Albania                                | sPBHpS  |
| alberta                                     | Canada                | Alberta                                | sPBHpS  |
| algeria                                     | Africa                | Algeria                                | sPBHpS  |
| alps                                        | Europe                | Alps                                   | sPBHp   |
| alsace                                      | France                | Alsace                                 | sPBHpS  |
| andorra                                     | Europe                | Andorra                                | sPBHpS  |
| angola                                      | Africa                | Angola                                 | sPBHpS  |
| antarctica                                  |                       | Antarctica                             | sPBHpS  |
| aquitaine                                   | France                | Aquitaine                              | sPBHpS  |
| argentina                                   | South America         | Argentina                              | sPBHpS  |
| arizona                                     | North America         | Arizona                                | sPBHpS  |
| arkansas                                    | North America         | Arkansas                               | sPBHpS  |
| arnsberg-regbez                             | Nordrhein-Westfalen   | Regierungsbezirk Arnsberg              | sPBHpS  |
| asia                                        |                       | Asia                                   | sPBHp   |
| australia                                   | Australia and Oceania | Australia                              | sPBHp   |
| australia-oceania                           |                       | Australia and Oceania                  | sPBHp   |
| austria                                     | Europe                | Austria                                | sPBHpS  |
| auvergne                                    | France                | Auvergne                               | sPBHpS  |
| azerbaijan                                  | Asia                  | Azerbaijan                             | sPBHpS  |
| azores                                      | Europe                | Azores                                 | sPBHpS  |
| baden-wuerttemberg                          | Germany               | Baden-Württemberg                      | sPBHpS  |
| bangladesh                                  | Asia                  | Bangladesh                             | sPBHpS  |
| basse-normandie                             | France                | Basse-Normandie                        | sPBHpS  |
| bayern                                      | Germany               | Bayern                                 | sPBHpS  |
| belarus                                     | Europe                | Belarus                                | sPBHpS  |
| belgium                                     | Europe                | Belgium                                | sPBHpS  |
| belize                                      | Central America       | Belize                                 | sPBHpS  |
| benin                                       | Africa                | Benin                                  | sPBHpS  |
| berkshire                                   | England               | Berkshire                              | sPBHpS  |
| berlin                                      | Germany               | Berlin                                 | sPBHpS  |
| bhutan                                      | Asia                  | Bhutan                                 | sPBHpS  |
| bolivia                                     | South America         | Bolivia                                | sPBHpS  |
| bosnia-herzegovina                          | Europe                | Bosnia-Herzegovina                     | sPBHpS  |
| botswana                                    | Africa                | Botswana                               | sPBHpS  |
| bourgogne                                   | France                | Bourgogne                              | sPBHpS  |
| brandenburg                                 | Germany               | Brandenburg (mit Berlin)               | sPBHpS  |
| brazil                                      | South America         | Brazil                                 | sPBHpS  |
| bremen                                      | Germany               | Bremen                                 | sPBHpS  |
| bretagne                                    | France                | Bretagne                               | sPBHpS  |
| british-columbia                            | Canada                | British Columbia                       | sPBHpS  |
| british-isles                               | Europe                | British Isles                          | sPBHp   |
| buckinghamshire                             | England               | Buckinghamshire                        | sPBHpS  |
| bulgaria                                    | Europe                | Bulgaria                               | sPBHpS  |
| burkina-faso                                | Africa                | Burkina Faso                           | sPBHpS  |
| burundi                                     | Africa                | Burundi                                | sPBHpS  |
| california                                  | North America         | California                             | sPBHpS  |
| cambodia                                    | Asia                  | Cambodia                               | sPBHpS  |
| cambridgeshire                              | England               | Cambridgeshire                         | sPBHpS  |
| cameroon                                    | Africa                | Cameroon                               | sPBHpS  |
| canada                                      | North America         | Canada                                 | sPBHp   |
| canary-islands                              | Africa                | Canary Islands                         | sPBHpS  |
| cape-verde                                  | Africa                | Cape Verde                             | sPBHpS  |
| central-african-republic                    | Africa                | Central African Republic               | sPBHpS  |
| central-america                             |                       | Central America                        | sPBHp   |
| central-fed-district                        | Russian Federation    | Central Federal District               | sPBHpS  |
| centre                                      | France                | Centre                                 | sPBHpS  |
| centro                                      | Italy                 | Centro                                 | sPBHpS  |
| chad                                        | Africa                | Chad                                   | sPBHpS  |
| champagne-ardenne                           | France                | Champagne Ardenne                      | sPBHpS  |
| cheshire                                    | England               | Cheshire                               | sPBHpS  |
| chile                                       | South America         | Chile                                  | sPBHpS  |
| china                                       | Asia                  | China                                  | sPBHpS  |
| chubu                                       | Japan                 | Chūbu region                           | sPBHpS  |
| chugoku                                     | Japan                 | Chūgoku region                         | sPBHpS  |
| colombia                                    | South America         | Colombia                               | sPBHpS  |
| colorado                                    | North America         | Colorado                               | sPBHpS  |
| comores                                     | Africa                | Comores                                | sPBHpS  |
| congo-brazzaville                           | Africa                | Congo (Republic/Brazzaville)           | sPBHpS  |
| congo-democratic-republic                   | Africa                | Congo (Democratic                      | sPBHpS  |
|                                             |                       | Republic/Kinshasa)                     |         |
| connecticut                                 | North America         | Connecticut                            | sPBHpS  |
| cornwall                                    | England               | Cornwall                               | sPBHpS  |
| corse                                       | France                | Corse                                  | sPBHpS  |
| crimean-fed-district                        | Russian Federation    | Crimean Federal District               | sPBHpS  |
| croatia                                     | Europe                | Croatia                                | sPBHpS  |
| cuba                                        | Central America       | Cuba                                   | sPBHpS  |
| cumbria                                     | England               | Cumbria                                | sPBHpS  |
| cyprus                                      | Europe                | Cyprus                                 | sPBHpS  |
| czech-republic                              | Europe                | Czech Republic                         | sPBHpS  |
| dach                                        | Europe                | Germany, Austria, Switzerland          | sPBHp   |
| delaware                                    | North America         | Delaware                               | sPBHpS  |
| denmark                                     | Europe                | Denmark                                | sPBHpS  |
| derbyshire                                  | England               | Derbyshire                             | sPBHpS  |
| detmold-regbez                              | Nordrhein-Westfalen   | Regierungsbezirk Detmold               | sPBHpS  |
| devon                                       | England               | Devon                                  | sPBHpS  |
| district-of-columbia                        | North America         | District of Columbia                   | sPBHpS  |
| djibouti                                    | Africa                | Djibouti                               | sPBHpS  |
| dolnoslaskie                                | Poland                | Województwo dolnośląskie(Lower         | sPBHpS  |
|                                             |                       | Silesian Voivodeship)                  |         |
| dorset                                      | England               | Dorset                                 | sPBHpS  |
| duesseldorf-regbez                          | Nordrhein-Westfalen   | Regierungsbezirk Düsseldorf            | sPBHpS  |
| east-sussex                                 | England               | East Sussex                            | sPBHpS  |
| east-yorkshire-with-hull                    | England               | East Yorkshire with Hull               | sPBHpS  |
| ecuador                                     | South America         | Ecuador                                | sPBHp   |
| egypt                                       | Africa                | Egypt                                  | sPBHpS  |
| england                                     | Great Britain         | England                                | sPBHpS  |
| equatorial-guinea                           | Africa                | Equatorial Guinea                      | sPBHpS  |
| eritrea                                     | Africa                | Eritrea                                | sPBHpS  |
| essex                                       | England               | Essex                                  | sPBHpS  |
| estonia                                     | Europe                | Estonia                                | sPBHpS  |
| ethiopia                                    | Africa                | Ethiopia                               | sPBHpS  |
| europe                                      |                       | Europe                                 | sPBHp   |
| far-eastern-fed-district                    | Russian Federation    | Far Eastern Federal District           | sPBHpS  |
| faroe-islands                               | Europe                | Faroe Islands                          | sPBHpS  |
| fiji                                        | Australia and Oceania | Fiji                                   | sPBHpS  |
| finland                                     | Europe                | Finland                                | sPBHpS  |
| florida                                     | North America         | Florida                                | sPBHpS  |
| france                                      | Europe                | France                                 | sPBHp   |
| franche-comte                               | France                | Franche Comte                          | sPBHpS  |
| freiburg-regbez                             | Baden-Württemberg     | Regierungsbezirk Freiburg              | sPBHpS  |
| gabon                                       | Africa                | Gabon                                  | sPBHpS  |
| gcc-states                                  | Asia                  | GCC States                             | sPBHpS  |
| georgia-eu                                  | Europe                | Georgia (Europe country)               | sPBHpS  |
| georgia-us                                  | North America         | Georgia (US State)                     | sPBHpS  |
| germany                                     | Europe                | Germany                                | sPBHp   |
| ghana                                       | Africa                | Ghana                                  | sPBHpS  |
| gloucestershire                             | England               | Gloucestershire                        | sPBHpS  |
| great-britain                               | Europe                | Great Britain                          | sPBHp   |
| greater-london                              | England               | Greater London                         | sPBHpS  |
| greater-manchester                          | England               | Greater Manchester                     | sPBHpS  |
| greece                                      | Europe                | Greece                                 | sPBHpS  |
| greenland                                   | North America         | Greenland                              | sPBHpS  |
| guadeloupe                                  | France                | Guadeloupe                             | sPBHp   |
| guatemala                                   | Central America       | Guatemala                              | sPBHpS  |
| guinea                                      | Africa                | Guinea                                 | sPBHpS  |
| guinea-bissau                               | Africa                | Guinea-Bissau                          | sPBHpS  |
| guyane                                      | France                | Guyane                                 | sPBHp   |
| haiti-and-domrep                            | Central America       | Haiti and Dominican Republic           | sPBHpS  |
| hamburg                                     | Germany               | Hamburg                                | sPBHpS  |
| hampshire                                   | England               | Hampshire                              | sPBHpS  |
| haute-normandie                             | France                | Haute-Normandie                        | sPBHpS  |
| hawaii                                      | North America         | Hawaii                                 | sPBHpS  |
| herefordshire                               | England               | Herefordshire                          | sPBHpS  |
| hertfordshire                               | England               | Hertfordshire                          | sPBHpS  |
| hessen                                      | Germany               | Hessen                                 | sPBHpS  |
| hokkaido                                    | Japan                 | Hokkaidō                               | sPBHpS  |
| hungary                                     | Europe                | Hungary                                | sPBHpS  |
| iceland                                     | Europe                | Iceland                                | sPBHpS  |
| idaho                                       | North America         | Idaho                                  | sPBHpS  |
| ile-de-france                               | France                | Ile-de-France                          | sPBHpS  |
| illinois                                    | North America         | Illinois                               | sPBHpS  |
| india                                       | Asia                  | India                                  | sPBHpS  |
| indiana                                     | North America         | Indiana                                | sPBHpS  |
| indonesia                                   | Asia                  | Indonesia                              | sPBHpS  |
| iowa                                        | North America         | Iowa                                   | sPBHpS  |
| iran                                        | Asia                  | Iran                                   | sPBHpS  |
| iraq                                        | Asia                  | Iraq                                   | sPBHpS  |
| ireland-and-northern-ireland                | Europe                | Ireland and Northern Ireland           | sPBHpS  |
| isle-of-man                                 | Europe                | Isle of Man                            | sPBHpS  |
| isle-of-wight                               | England               | Isle of Wight                          | sPBHpS  |
| isole                                       | Italy                 | Isole                                  | sPBHpS  |
| israel-and-palestine                        | Asia                  | Israel and Palestine                   | sPBHpS  |
| italy                                       | Europe                | Italy                                  | sPBHp   |
| ivory-coast                                 | Africa                | Ivory Coast                            | sPBHpS  |
| japan                                       | Asia                  | Japan                                  | sPBHp   |
| jordan                                      | Asia                  | Jordan                                 | sPBHpS  |
| kaliningrad                                 | Russian Federation    | Kaliningrad                            | sPBHpS  |
| kansai                                      | Japan                 | Kansai region (a.k.a. Kinki            | sPBHpS  |
|                                             |                       | region)                                |         |
| kansas                                      | North America         | Kansas                                 | sPBHpS  |
| kanto                                       | Japan                 | Kantō region                           | sPBHpS  |
| karlsruhe-regbez                            | Baden-Württemberg     | Regierungsbezirk Karlsruhe             | sPBHpS  |
| kazakhstan                                  | Asia                  | Kazakhstan                             | sPBHpS  |
| kent                                        | England               | Kent                                   | sPBHpS  |
| kentucky                                    | North America         | Kentucky                               | sPBHpS  |
| kenya                                       | Africa                | Kenya                                  | sPBHpS  |
| koeln-regbez                                | Nordrhein-Westfalen   | Regierungsbezirk Köln                  | sPBHpS  |
| kosovo                                      | Europe                | Kosovo                                 | sPBHpS  |
| kujawsko-pomorskie                          | Poland                | Województwo                            | sPBHpS  |
|                                             |                       | kujawsko-pomorskie(Kuyavian-Pomeranian |         |
|                                             |                       | Voivodeship)                           |         |
| kyrgyzstan                                  | Asia                  | Kyrgyzstan                             | sPBHpS  |
| kyushu                                      | Japan                 | Kyūshū                                 | sPBHpS  |
| lancashire                                  | England               | Lancashire                             | sPBHpS  |
| languedoc-roussillon                        | France                | Languedoc-Roussillon                   | sPBHpS  |
| latvia                                      | Europe                | Latvia                                 | sPBHpS  |
| lebanon                                     | Asia                  | Lebanon                                | sPBHpS  |
| leicestershire                              | England               | Leicestershire                         | sPBHpS  |
| lesotho                                     | Africa                | Lesotho                                | sPBHpS  |
| liberia                                     | Africa                | Liberia                                | sPBHpS  |
| libya                                       | Africa                | Libya                                  | sPBHpS  |
| liechtenstein                               | Europe                | Liechtenstein                          | sPBHpS  |
| limousin                                    | France                | Limousin                               | sPBHpS  |
| lithuania                                   | Europe                | Lithuania                              | sPBHpS  |
| lodzkie                                     | Poland                | Województwo łódzkie(Łódź Voivodeship)  | sPBHpS  |
| lorraine                                    | France                | Lorraine                               | sPBHpS  |
| louisiana                                   | North America         | Louisiana                              | sPBHpS  |
| lubelskie                                   | Poland                | Województwo lubelskie(Lublin           | sPBHpS  |
|                                             |                       | Voivodeship)                           |         |
| lubuskie                                    | Poland                | Województwo lubuskie(Lubusz            | sPBHpS  |
|                                             |                       | Voivodeship)                           |         |
| luxembourg                                  | Europe                | Luxembourg                             | sPBHpS  |
| macedonia                                   | Europe                | Macedonia                              | sPBHpS  |
| madagascar                                  | Africa                | Madagascar                             | sPBHpS  |
| maine                                       | North America         | Maine                                  | sPBHpS  |
| malawi                                      | Africa                | Malawi                                 | sPBHpS  |
| malaysia-singapore-brunei                   | Asia                  | Malaysia, Singapore, and Brunei        | sPBHp   |
| maldives                                    | Asia                  | Maldives                               | sPBHpS  |
| mali                                        | Africa                | Mali                                   | sPBHpS  |
| malopolskie                                 | Poland                | Województwo małopolskie(Lesser Poland  | sPBHpS  |
|                                             |                       | Voivodeship)                           |         |
| malta                                       | Europe                | Malta                                  | sPBHpS  |
| manitoba                                    | Canada                | Manitoba                               | sPBHpS  |
| martinique                                  | France                | Martinique                             | sPBHp   |
| maryland                                    | North America         | Maryland                               | sPBHpS  |
| massachusetts                               | North America         | Massachusetts                          | sPBHpS  |
| mauritania                                  | Africa                | Mauritania                             | sPBHpS  |
| mauritius                                   | Africa                | Mauritius                              | sPBHpS  |
| mayotte                                     | France                | Mayotte                                | sPBHp   |
| mazowieckie                                 | Poland                | Województwo mazowieckie(Mazovian       | sPBHpS  |
|                                             |                       | Voivodeship)                           |         |
| mecklenburg-vorpommern                      | Germany               | Mecklenburg-Vorpommern                 | sPBHpS  |
| mexico                                      | North America         | Mexico                                 | sPBHpS  |
| michigan                                    | North America         | Michigan                               | sPBHpS  |
| midi-pyrenees                               | France                | Midi-Pyrenees                          | sPBHpS  |
| minnesota                                   | North America         | Minnesota                              | sPBHpS  |
| mississippi                                 | North America         | Mississippi                            | sPBHpS  |
| missouri                                    | North America         | Missouri                               | sPBHpS  |
| mittelfranken                               | Bayern                | Mittelfranken                          | sPBHpS  |
| moldova                                     | Europe                | Moldova                                | sPBHpS  |
| monaco                                      | Europe                | Monaco                                 | sPBHpS  |
| mongolia                                    | Asia                  | Mongolia                               | sPBHpS  |
| montana                                     | North America         | Montana                                | sPBHpS  |
| montenegro                                  | Europe                | Montenegro                             | sPBHpS  |
| morocco                                     | Africa                | Morocco                                | sPBHpS  |
| mozambique                                  | Africa                | Mozambique                             | sPBHpS  |
| muenster-regbez                             | Nordrhein-Westfalen   | Regierungsbezirk Münster               | sPBHpS  |
| myanmar                                     | Asia                  | Myanmar (a.k.a. Burma)                 | sPBHpS  |
| namibia                                     | Africa                | Namibia                                | sPBHpS  |
| nebraska                                    | North America         | Nebraska                               | sPBHpS  |
| nepal                                       | Asia                  | Nepal                                  | sPBHpS  |
| netherlands                                 | Europe                | Netherlands                            | sPBHpS  |
| nevada                                      | North America         | Nevada                                 | sPBHpS  |
| new-brunswick                               | Canada                | New Brunswick                          | sPBHpS  |
| new-caledonia                               | Australia and Oceania | New Caledonia                          | sPBHpS  |
| new-hampshire                               | North America         | New Hampshire                          | sPBHpS  |
| new-jersey                                  | North America         | New Jersey                             | sPBHpS  |
| new-mexico                                  | North America         | New Mexico                             | sPBHpS  |
| new-york                                    | North America         | New York                               | sPBHpS  |
| new-zealand                                 | Australia and Oceania | New Zealand                            | sPBHpS  |
| newfoundland-and-labrador                   | Canada                | Newfoundland and Labrador              | sPBHpS  |
| nicaragua                                   | Central America       | Nicaragua                              | sPBHp   |
| niederbayern                                | Bayern                | Niederbayern                           | sPBHpS  |
| niedersachsen                               | Germany               | Niedersachsen                          | sPBHpS  |
| niger                                       | Africa                | Niger                                  | sPBHpS  |
| nigeria                                     | Africa                | Nigeria                                | sPBHpS  |
| nord-est                                    | Italy                 | Nord-Est                               | sPBHpS  |
| nord-ovest                                  | Italy                 | Nord-Ovest                             | sPBHpS  |
| nord-pas-de-calais                          | France                | Nord-Pas-de-Calais                     | sPBHpS  |
| nordrhein-westfalen                         | Germany               | Nordrhein-Westfalen                    | sPBHpS  |
| norfolk                                     | England               | Norfolk                                | sPBHpS  |
| north-america                               |                       | North America                          | sPBHp   |
| north-carolina                              | North America         | North Carolina                         | sPBHpS  |
| north-caucasus-fed-district                 | Russian Federation    | North Caucasus Federal District        | sPBHpS  |
| north-dakota                                | North America         | North Dakota                           | sPBHpS  |
| north-korea                                 | Asia                  | North Korea                            | sPBHpS  |
| north-yorkshire                             | England               | North Yorkshire                        | sPBHpS  |
| northumberland                              | England               | Northumberland                         | sPBHpS  |
| northwest-territories                       | Canada                | Northwest Territories                  | sPBHpS  |
| northwestern-fed-district                   | Russian Federation    | Northwestern Federal District          | sPBHpS  |
| norway                                      | Europe                | Norway                                 | sPBHpS  |
| nottinghamshire                             | England               | Nottinghamshire                        | sPBHpS  |
| nova-scotia                                 | Canada                | Nova Scotia                            | sPBHpS  |
| nunavut                                     | Canada                | Nunavut                                | sPBHpS  |
| oberbayern                                  | Bayern                | Oberbayern                             | sPBHpS  |
| oberfranken                                 | Bayern                | Oberfranken                            | sPBHpS  |
| oberpfalz                                   | Bayern                | Oberpfalz                              | sPBHpS  |
| ohio                                        | North America         | Ohio                                   | sPBHpS  |
| oklahoma                                    | North America         | Oklahoma                               | sPBHpS  |
| ontario                                     | Canada                | Ontario                                | sPBHpS  |
| opolskie                                    | Poland                | Województwo opolskie(Opole             | sPBHpS  |
|                                             |                       | Voivodeship)                           |         |
| oregon                                      | North America         | Oregon                                 | sPBHpS  |
| oxfordshire                                 | England               | Oxfordshire                            | sPBHpS  |
| pakistan                                    | Asia                  | Pakistan                               | sPBHpS  |
| papua-new-guinea                            | Australia and Oceania | Papua New Guinea                       | sPBHpS  |
| paraguay                                    | South America         | Paraguay                               | sPBHpS  |
| pays-de-la-loire                            | France                | Pays de la Loire                       | sPBHpS  |
| pennsylvania                                | North America         | Pennsylvania                           | sPBHpS  |
| peru                                        | South America         | Peru                                   | sPBHpS  |
| philippines                                 | Asia                  | Philippines                            | sPBHpS  |
| picardie                                    | France                | Picardie                               | sPBHpS  |
| podkarpackie                                | Poland                | Województwo podkarpackie(Subcarpathian | sPBHpS  |
|                                             |                       | Voivodeship)                           |         |
| podlaskie                                   | Poland                | Województwo podlaskie(Podlaskie        | sPBHpS  |
|                                             |                       | Voivodeship)                           |         |
| poitou-charentes                            | France                | Poitou-Charentes                       | sPBHpS  |
| poland                                      | Europe                | Poland                                 | sPBHp   |
| pomorskie                                   | Poland                | Województwo pomorskie(Pomeranian       | sPBHpS  |
|                                             |                       | Voivodeship)                           |         |
| portugal                                    | Europe                | Portugal                               | sPBHpS  |
| prince-edward-island                        | Canada                | Prince Edward Island                   | sPBHpS  |
| provence-alpes-cote-d-azur                  | France                | Provence Alpes-Cote-d'Azur             | sPBHpS  |
| puerto-rico                                 | North America         | Puerto Rico                            | sPBHpS  |
| quebec                                      | Canada                | Quebec                                 | sPBHpS  |
| reunion                                     | France                | Reunion                                | sPBHp   |
| rheinland-pfalz                             | Germany               | Rheinland-Pfalz                        | sPBHpS  |
| rhode-island                                | North America         | Rhode Island                           | sPBHpS  |
| rhone-alpes                                 | France                | Rhone-Alpes                            | sPBHpS  |
| romania                                     | Europe                | Romania                                | sPBHpS  |
| russia                                      |                       | Russian Federation                     | sPBHp   |
| rwanda                                      | Africa                | Rwanda                                 | sPBHpS  |
| saarland                                    | Germany               | Saarland                               | sPBHpS  |
| sachsen                                     | Germany               | Sachsen                                | sPBHpS  |
| sachsen-anhalt                              | Germany               | Sachsen-Anhalt                         | sPBHpS  |
| saint-helena-ascension-and-tristan-da-cunha | Africa                | Saint Helena, Ascension, and Tristan   | sPBHpS  |
|                                             |                       | da Cunha                               |         |
| sao-tome-and-principe                       | Africa                | Sao Tome and Principe                  | sPBHpS  |
| saskatchewan                                | Canada                | Saskatchewan                           | sPBHpS  |
| schleswig-holstein                          | Germany               | Schleswig-Holstein                     | sPBHpS  |
| schwaben                                    | Bayern                | Schwaben                               | sPBHpS  |
| scotland                                    | Great Britain         | Scotland                               | sPBHpS  |
| senegal-and-gambia                          | Africa                | Senegal and Gambia                     | sPBHpS  |
| serbia                                      | Europe                | Serbia                                 | sPBHpS  |
| seychelles                                  | Africa                | Seychelles                             | sPBHpS  |
| shikoku                                     | Japan                 | Shikoku                                | sPBHpS  |
| shropshire                                  | England               | Shropshire                             | sPBHpS  |
| siberian-fed-district                       | Russian Federation    | Siberian Federal District              | sPBHpS  |
| sierra-leone                                | Africa                | Sierra Leone                           | sPBHpS  |
| slaskie                                     | Poland                | Województwo śląskie(Silesian           | sPBHpS  |
|                                             |                       | Voivodeship)                           |         |
| slovakia                                    | Europe                | Slovakia                               | sPBHpS  |
| slovenia                                    | Europe                | Slovenia                               | sPBHpS  |
| somalia                                     | Africa                | Somalia                                | sPBHpS  |
| somerset                                    | England               | Somerset                               | sPBHpS  |
| south-africa                                | Africa                | South Africa                           | sPBHpS  |
| south-africa-and-lesotho                    | Africa                | South Africa (includes Lesotho)        | sPBHp   |
| south-america                               |                       | South America                          | sPBHp   |
| south-carolina                              | North America         | South Carolina                         | sPBHpS  |
| south-dakota                                | North America         | South Dakota                           | sPBHpS  |
| south-fed-district                          | Russian Federation    | South Federal District                 | sPBHpS  |
| south-korea                                 | Asia                  | South Korea                            | sPBHpS  |
| south-sudan                                 | Africa                | South Sudan                            | sPBHpS  |
| south-yorkshire                             | England               | South Yorkshire                        | sPBHpS  |
| spain                                       | Europe                | Spain                                  | sPBHpS  |
| sri-lanka                                   | Asia                  | Sri Lanka                              | sPBHpS  |
| staffordshire                               | England               | Staffordshire                          | sPBHpS  |
| stuttgart-regbez                            | Baden-Württemberg     | Regierungsbezirk Stuttgart             | sPBHpS  |
| sud                                         | Italy                 | Sud                                    | sPBHpS  |
| sudan                                       | Africa                | Sudan                                  | sPBHpS  |
| suffolk                                     | England               | Suffolk                                | sPBHpS  |
| suriname                                    | South America         | suriname                               | sPBHpS  |
| surrey                                      | England               | Surrey                                 | sPBHpS  |
| swaziland                                   | Africa                | Swaziland                              | sPBHpS  |
| sweden                                      | Europe                | Sweden                                 | sPBHpS  |
| swietokrzyskie                              | Poland                | Województwo                            | sPBHpS  |
|                                             |                       | świętokrzyskie(Świętokrzyskie          |         |
|                                             |                       | Voivodeship)                           |         |
| switzerland                                 | Europe                | Switzerland                            | sPBHpS  |
| syria                                       | Asia                  | Syria                                  | sPBHpS  |
| taiwan                                      | Asia                  | Taiwan                                 | sPBHpS  |
| tajikistan                                  | Asia                  | Tajikistan                             | sPBHp   |
| tanzania                                    | Africa                | Tanzania                               | sPBHpS  |
| tennessee                                   | North America         | Tennessee                              | sPBHpS  |
| texas                                       | North America         | Texas                                  | sPBHpS  |
| thailand                                    | Asia                  | Thailand                               | sPBHpS  |
| thueringen                                  | Germany               | Thüringen                              | sPBHpS  |
| togo                                        | Africa                | Togo                                   | sPBHpS  |
| tohoku                                      | Japan                 | Tōhoku region                          | sPBHpS  |
| tuebingen-regbez                            | Baden-Württemberg     | Regierungsbezirk Tübingen              | sPBHpS  |
| tunisia                                     | Africa                | Tunisia                                | sPBHpS  |
| turkey                                      | Europe                | Turkey                                 | sPBHpS  |
| turkmenistan                                | Asia                  | Turkmenistan                           | sPBHpS  |
| uganda                                      | Africa                | Uganda                                 | sPBHpS  |
| ukraine                                     | Europe                | Ukraine                                | sPBHpS  |
| unterfranken                                | Bayern                | Unterfranken                           | sPBHpS  |
| ural-fed-district                           | Russian Federation    | Ural Federal District                  | sPBHpS  |
| uruguay                                     | South America         | Uruguay                                | sPBHpS  |
| us-midwest                                  | North America         | US Midwest                             | sPBHp   |
| us-northeast                                | North America         | US Northeast                           | sPBHp   |
| us-pacific                                  | North America         | US Pacific                             | sPBHp   |
| us-south                                    | North America         | US South                               | sPBHp   |
| us-west                                     | North America         | US West                                | sPBHp   |
| utah                                        | North America         | Utah                                   | sPBHpS  |
| uzbekistan                                  | Asia                  | Uzbekistan                             | sPBHpS  |
| vermont                                     | North America         | Vermont                                | sPBHpS  |
| vietnam                                     | Asia                  | Vietnam                                | sPBHpS  |
| virginia                                    | North America         | Virginia                               | sPBHpS  |
| volga-fed-district                          | Russian Federation    | Volga Federal District                 | sPBHpS  |
| wales                                       | Great Britain         | Wales                                  | sPBHpS  |
| warminsko-mazurskie                         | Poland                | Województwo                            | sPBHpS  |
|                                             |                       | warmińsko-mazurskie(Warmian-Masurian   |         |
|                                             |                       | Voivodeship)                           |         |
| washington                                  | North America         | Washington                             | sPBHpS  |
| west-midlands                               | England               | West Midlands                          | sPBHpS  |
| west-sussex                                 | England               | West Sussex                            | sPBHpS  |
| west-virginia                               | North America         | West Virginia                          | sPBHpS  |
| west-yorkshire                              | England               | West Yorkshire                         | sPBHpS  |
| wielkopolskie                               | Poland                | Województwo wielkopolskie(Greater      | sPBHpS  |
|                                             |                       | Poland Voivodeship)                    |         |
| wiltshire                                   | England               | Wiltshire                              | sPBHpS  |
| wisconsin                                   | North America         | Wisconsin                              | sPBHpS  |
| worcestershire                              | England               | worcestershire                         | sPBHpS  |
| wyoming                                     | North America         | Wyoming                                | sPBHpS  |
| yemen                                       | Asia                  | Yemen                                  | sPBHpS  |
| yukon                                       | Canada                | Yukon                                  | sPBHpS  |
| zachodniopomorskie                          | Poland                | Województwo zachodniopomorskie(West    | sPBHpS  |
|                                             |                       | Pomeranian Voivodeship)                |         |
| zambia                                      | Africa                | Zambia                                 | sPBHpS  |
| zimbabwe                                    | Africa                | Zimbabwe                               | sPBHpS  |
