# download-geofabrik

[![Build Status](https://travis-ci.org/julien-noblet/download-geofabrik.svg?branch=master)](https://travis-ci.org/julien-noblet/download-geofabrik)
[![Join the chat at https://gitter.im/julien-noblet/download-geofabrik](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/julien-noblet/download-geofabrik?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

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
    -s, --state    Download state.txt file
    -p, --poly     Download poly file
```

## List of elements
|          SHORTNAME           |         IS IN         |           LONG NAME                    | FORMATS |
|------------------------------|-----------------------|----------------------------------------|---------|
| afghanistan                  | Asia                  | Afghanistan                            | sPBpS   |
| africa                       |                       | Africa                                 | sPBp    |
| alabama                      | North America         | Alabama                                | sPBpS   |
| alaska                       | North America         | Alaska                                 | sPBpS   |
| albania                      | Europe                | Albania                                | sPBpS   |
| alberta                      | Canada                | Alberta                                | sPBpS   |
| alps                         | Europe                | Alps                                   | sPBp    |
| alsace                       | France                | Alsace                                 | sPBpS   |
| andorra                      | Europe                | Andorra                                | sPBpS   |
| antarctica                   |                       | Antarctica                             | sPBpS   |
| aquitaine                    | France                | Aquitaine                              | sPBpS   |
| argentina                    | South America         | Argentina                              | sPBpS   |
| arizona                      | North America         | Arizona                                | sPBpS   |
| arkansas                     | North America         | Arkansas                               | sPBpS   |
| arnsberg-regbez              | Nordrhein-Westfalen   | Regierungsbezirk Arnsberg              | sPBpS   |
| asia                         |                       | Asia                                   | sPBp    |
| australia                    | Australia and Oceania | Australia                              | sPBp    |
| australia-oceania            |                       | Australia and Oceania                  | sPBp    |
| austria                      | Europe                | Austria                                | sPBpS   |
| auvergne                     | France                | Auvergne                               | sPBpS   |
| azerbaijan                   | Asia                  | Azerbaijan                             | sPBpS   |
| azores                       | Europe                | Azores                                 | sPBpS   |
| baden-wuerttemberg           | Germany               | Baden-Württemberg                      | sPBpS   |
| bangladesh                   | Asia                  | Bangladesh                             | sPBpS   |
| basse-normandie              | France                | Basse-Normandie                        | sPBpS   |
| bayern                       | Germany               | Bayern                                 | sPBpS   |
| belarus                      | Europe                | Belarus                                | sPBpS   |
| belgium                      | Europe                | Belgium                                | sPBpS   |
| belize                       | Central America       | Belize                                 | sPBpS   |
| benin                        | Africa                | Benin                                  | sPBp    |
| berkshire                    | England               | Berkshire                              | sPBpS   |
| berlin                       | Germany               | Berlin                                 | sPBpS   |
| bolivia                      | South America         | Bolivia                                | sPBpS   |
| bosnia-herzegovina           | Europe                | Bosnia-Herzegovina                     | sPBpS   |
| botswana                     | Africa                | Botswana                               | sPBpS   |
| bourgogne                    | France                | Bourgogne                              | sPBpS   |
| brandenburg                  | Germany               | Brandenburg                            | sPBpS   |
| brazil                       | South America         | Brazil                                 | sPBpS   |
| bremen                       | Germany               | Bremen                                 | sPBpS   |
| bretagne                     | France                | Bretagne                               | sPBpS   |
| british-columbia             | Canada                | British Columbia                       | sPBpS   |
| british-isles                | Europe                | British Isles                          | sPBp    |
| buckinghamshire              | England               | Buckinghamshire                        | sPBpS   |
| bulgaria                     | Europe                | Bulgaria                               | sPBpS   |
| burkina-faso                 | Africa                | Burkina Faso                           | sPBp    |
| california                   | North America         | California                             | sPBpS   |
| cambodia                     | Asia                  | Cambodia                               | sPBpS   |
| cambridgeshire               | England               | Cambridgeshire                         | sPBpS   |
| cameroon                     | Africa                | Cameroon                               | sPBp    |
| canada                       | North America         | Canada                                 | sPBp    |
| canary-islands               | Africa                | Canary Islands                         | sPBpS   |
| central-america              |                       | Central America                        | sPBp    |
| central-fed-district         | Russian Federation    | Central Federal District               | sPBpS   |
| centre                       | France                | Centre                                 | sPBpS   |
| centro                       | Italy                 | Centro                                 | sPBpS   |
| chad                         | Africa                | Chad                                   | sPBpS   |
| champagne-ardenne            | France                | Champagne Ardenne                      | sPBpS   |
| cheshire                     | England               | Cheshire                               | sPBpS   |
| chile                        | South America         | Chile                                  | sPBpS   |
| china                        | Asia                  | China                                  | sPBpS   |
| colombia                     | South America         | Colombia                               | sPBpS   |
| colorado                     | North America         | Colorado                               | sPBpS   |
| congo-democratic-republic    | Africa                | Congo (Democratic Republic)            | sPBp    |
| connecticut                  | North America         | Connecticut                            | sPBpS   |
| cornwall                     | England               | Cornwall                               | sPBpS   |
| corse                        | France                | Corse                                  | sPBpS   |
| crimean-fed-district         | Russian Federation    | Crimean Federal District               | sPBpS   |
| croatia                      | Europe                | Croatia                                | sPBpS   |
| cuba                         | Central America       | Cuba                                   | sPBpS   |
| cumbria                      | England               | Cumbria                                | sPBpS   |
| cyprus                       | Europe                | Cyprus                                 | sPBpS   |
| czech-republic               | Europe                | Czech Republic                         | sPBpS   |
| dach                         | Europe                | Germany, Austria, Switzerland          | sPBp    |
| delaware                     | North America         | Delaware                               | sPBpS   |
| denmark                      | Europe                | Denmark                                | sPBpS   |
| derbyshire                   | England               | Derbyshire                             | sPBpS   |
| detmold-regbez               | Nordrhein-Westfalen   | Regierungsbezirk Detmold               | sPBpS   |
| devon                        | England               | Devon                                  | sPBpS   |
| district-of-columbia         | North America         | District of Columbia                   | sPBpS   |
| dolnoslaskie                 | Poland                | Województwo dolnośląskie(Lower Silesian Voivodeship) | sPBpS   |
| dorset                       | England               | Dorset                                 | sPBpS   |
| duesseldorf-regbez           | Nordrhein-Westfalen   | Regierungsbezirk Düsseldorf            | sPBpS   |
| east-sussex                  | England               | East Sussex                            | sPBpS   |
| east-yorkshire-with-hull     | England               | East Yorkshire with Hull               | sPBpS   |
| ecuador                      | South America         | Ecuador                                | sPBp    |
| egypt                        | Africa                | Egypt                                  | sPBp    |
| england                      | Great Britain         | England                                | sPBpS   |
| essex                        | England               | Essex                                  | sPBpS   |
| estonia                      | Europe                | Estonia                                | sPBpS   |
| ethiopia                     | Africa                | Ethiopia                               | sPBpS   |
| europe                       |                       | Europe                                 | sPBp    |
| far-eastern-fed-district     | Russian Federation    | Far Eastern Federal District           | sPBpS   |
| faroe-islands                | Europe                | Faroe Islands                          | sPBpS   |
| fiji                         | Australia and Oceania | Fiji                                   | sPBpS   |
| finland                      | Europe                | Finland                                | sPBpS   |
| florida                      | North America         | Florida                                | sPBpS   |
| france                       | Europe                | France                                 | sPBp    |
| franche-comte                | France                | Franche Comte                          | sPBpS   |
| freiburg-regbez              | Baden-Württemberg     | Regierungsbezirk Freiburg              | sPBpS   |
| gcc-states                   | Asia                  | GCC States                             | sPBpS   |
| georgia-eu                   | Europe                | Georgia (Europe country)               | sPBpS   |
| georgia-us                   | North America         | Georgia (US State)                     | sPBpS   |
| germany                      | Europe                | Germany                                | sPBp    |
| gloucestershire              | England               | Gloucestershire                        | sPBpS   |
| great-britain                | Europe                | Great Britain                          | sPBp    |
| greater-london               | England               | Greater London                         | sPBpS   |
| greater-manchester           | England               | Greater Manchester                     | sPBpS   |
| greece                       | Europe                | Greece                                 | sPBpS   |
| greenland                    | North America         | Greenland                              | sPBpS   |
| guadeloupe                   | France                | Guadeloupe                             | sPBp    |
| guatemala                    | Central America       | Guatemala                              | sPBpS   |
| guinea                       | Africa                | Guinea                                 | sPBpS   |
| guinea-bissau                | Africa                | Guinea-Bissau                          | sPBp    |
| guyane                       | France                | Guyane                                 | sPBp    |
| haiti-and-domrep             | Central America       | Haiti and Dominican Republic           | sPBpS   |
| hamburg                      | Germany               | Hamburg                                | sPBpS   |
| hampshire                    | England               | Hampshire                              | sPBpS   |
| haute-normandie              | France                | Haute-Normandie                        | sPBpS   |
| hawaii                       | North America         | Hawaii                                 | sPBpS   |
| herefordshire                | England               | Herefordshire                          | sPBpS   |
| hertfordshire                | England               | Hertfordshire                          | sPBpS   |
| hessen                       | Germany               | Hessen                                 | sPBpS   |
| hungary                      | Europe                | Hungary                                | sPBpS   |
| iceland                      | Europe                | Iceland                                | sPBpS   |
| idaho                        | North America         | Idaho                                  | sPBpS   |
| ile-de-france                | France                | Ile-de-France                          | sPBpS   |
| illinois                     | North America         | Illinois                               | sPBpS   |
| india                        | Asia                  | India                                  | sPBpS   |
| indiana                      | North America         | Indiana                                | sPBpS   |
| indonesia                    | Asia                  | Indonesia                              | sPBpS   |
| iowa                         | North America         | Iowa                                   | sPBpS   |
| iran                         | Asia                  | Iran                                   | sPBpS   |
| iraq                         | Asia                  | Iraq                                   | sPBpS   |
| ireland-and-northern-ireland | Europe                | Ireland and Northern Ireland           | sPBpS   |
| isle-of-man                  | Europe                | Isle of Man                            | sPBpS   |
| isle-of-wight                | England               | Isle of Wight                          | sPBpS   |
| isole                        | Italy                 | Isole                                  | sPBpS   |
| israel-and-palestine         | Asia                  | Israel and Palestine                   | sPBpS   |
| italy                        | Europe                | Italy                                  | sPBp    |
| ivory-coast                  | Africa                | Ivory Coast                            | sPBpS   |
| japan                        | Asia                  | Japan                                  | sPBpS   |
| jordan                       | Asia                  | Jordan                                 | sPBpS   |
| kansas                       | North America         | Kansas                                 | sPBpS   |
| karlsruhe-regbez             | Baden-Württemberg     | Regierungsbezirk Karlsruhe             | sPBpS   |
| kazakhstan                   | Asia                  | Kazakhstan                             | sPBpS   |
| kent                         | England               | Kent                                   | sPBpS   |
| kentucky                     | North America         | Kentucky                               | sPBpS   |
| kenya                        | Africa                | Kenya                                  | sPBpS   |
| koeln-regbez                 | Nordrhein-Westfalen   | Regierungsbezirk Köln                  | sPBpS   |
| kosovo                       | Europe                | Kosovo                                 | sPBpS   |
| kujawsko-pomorskie           | Poland                | Województwo kujawsko-pomorskie(Kuyavian-Pomeranian Voivodeship) | sPBpS   |
| kyrgyzstan                   | Asia                  | Kyrgyzstan                             | sPBpS   |
| lancashire                   | England               | Lancashire                             | sPBpS   |
| languedoc-roussillon         | France                | Languedoc-Roussillon                   | sPBpS   |
| latvia                       | Europe                | Latvia                                 | sPBpS   |
| lebanon                      | Asia                  | Lebanon                                | sPBpS   |
| leicestershire               | England               | Leicestershire                         | sPBpS   |
| lesotho                      | Africa                | Lesotho                                | sPBpS   |
| liberia                      | Africa                | Liberia                                | sPBpS   |
| libya                        | Africa                | Libya                                  | sPBpS   |
| liechtenstein                | Europe                | Liechtenstein                          | sPBpS   |
| limousin                     | France                | Limousin                               | sPBpS   |
| lithuania                    | Europe                | Lithuania                              | sPBpS   |
| lodzkie                      | Poland                | Województwo łódzkie(Łódź Voivodeship)  | sPBpS   |
| lorraine                     | France                | Lorraine                               | sPBpS   |
| louisiana                    | North America         | Louisiana                              | sPBpS   |
| lubelskie                    | Poland                | Województwo lubelskie(Lublin Voivodeship) | sPBpS   |
| lubuskie                     | Poland                | Województwo lubuskie(Lubusz Voivodeship) | sPBpS   |
| luxembourg                   | Europe                | Luxembourg                             | sPBpS   |
| macedonia                    | Europe                | Macedonia                              | sPBpS   |
| madagascar                   | Africa                | Madagascar                             | sPBpS   |
| maine                        | North America         | Maine                                  | sPBpS   |
| malaysia-singapore-brunei    | Asia                  | Malaysia, Singapore, and Brunei        | sPBp    |
| maldives                     | Asia                  | Maldives                               | sPBpS   |
| mali                         | Africa                | Mali                                   | sPBpS   |
| malopolskie                  | Poland                | Województwo małopolskie(Lesser Poland Voivodeship) | sPBpS   |
| malta                        | Europe                | Malta                                  | sPBpS   |
| manitoba                     | Canada                | Manitoba                               | sPBpS   |
| martinique                   | France                | Martinique                             | sPBp    |
| maryland                     | North America         | Maryland                               | sPBpS   |
| massachusetts                | North America         | Massachusetts                          | sPBpS   |
| mayotte                      | France                | Mayotte                                | sPBp    |
| mazowieckie                  | Poland                | Województwo mazowieckie(Mazovian Voivodeship) | sPBpS   |
| mecklenburg-vorpommern       | Germany               | Mecklenburg-Vorpommern                 | sPBpS   |
| mexico                       | North America         | Mexico                                 | sPBpS   |
| michigan                     | North America         | Michigan                               | sPBpS   |
| midi-pyrenees                | France                | Midi-Pyrenees                          | sPBpS   |
| minnesota                    | North America         | Minnesota                              | sPBpS   |
| mississippi                  | North America         | Mississippi                            | sPBpS   |
| missouri                     | North America         | Missouri                               | sPBpS   |
| mittelfranken                | Bayern                | Mittelfranken                          | sPBpS   |
| moldova                      | Europe                | Moldova                                | sPBpS   |
| monaco                       | Europe                | Monaco                                 | sPBpS   |
| mongolia                     | Asia                  | Mongolia                               | sPBpS   |
| montana                      | North America         | Montana                                | sPBpS   |
| montenegro                   | Europe                | Montenegro                             | sPBpS   |
| morocco                      | Africa                | Morocco                                | sPBpS   |
| muenster-regbez              | Nordrhein-Westfalen   | Regierungsbezirk Münster               | sPBpS   |
| namibia                      | Africa                | Namibia                                | sPBpS   |
| nebraska                     | North America         | Nebraska                               | sPBpS   |
| nepal                        | Asia                  | Nepal                                  | sPBpS   |
| netherlands                  | Europe                | Netherlands                            | sPBpS   |
| nevada                       | North America         | Nevada                                 | sPBpS   |
| new-brunswick                | Canada                | New Brunswick                          | sPBpS   |
| new-caledonia                | Australia and Oceania | New Caledonia                          | sPBpS   |
| new-hampshire                | North America         | New Hampshire                          | sPBpS   |
| new-jersey                   | North America         | New Jersey                             | sPBpS   |
| new-mexico                   | North America         | New Mexico                             | sPBpS   |
| new-york                     | North America         | New York                               | sPBpS   |
| new-zealand                  | Australia and Oceania | New Zealand                            | sPBpS   |
| newfoundland-and-labrador    | Canada                | Newfoundland and Labrador              | sPBpS   |
| nicaragua                    | Central America       | Nicaragua                              | sPBp    |
| niederbayern                 | Bayern                | Niederbayern                           | sPBpS   |
| niedersachsen                | Germany               | Niedersachsen                          | sPBpS   |
| niger                        | Africa                | Niger                                  | sPBpS   |
| nigeria                      | Africa                | Nigeria                                | sPBpS   |
| nord-est                     | Italy                 | Nord-Est                               | sPBpS   |
| nord-ovest                   | Italy                 | Nord-Ovest                             | sPBpS   |
| nord-pas-de-calais           | France                | Nord-Pas-de-Calais                     | sPBpS   |
| nordrhein-westfalen          | Germany               | Nordrhein-Westfalen                    | sPBpS   |
| norfolk                      | England               | Norfolk                                | sPBpS   |
| north-america                |                       | North America                          | sPBp    |
| north-carolina               | North America         | North Carolina                         | sPBpS   |
| north-caucasus-fed-district  | Russian Federation    | North Caucasus Federal District        | sPBpS   |
| north-dakota                 | North America         | North Dakota                           | sPBpS   |
| north-korea                  | Asia                  | North Korea                            | sPBpS   |
| north-yorkshire              | England               | North Yorkshire                        | sPBpS   |
| northwest-territories        | Canada                | Northwest Territories                  | sPBpS   |
| northwestern-fed-district    | Russian Federation    | Northwestern Federal District          | sPBpS   |
| norway                       | Europe                | Norway                                 | sPBpS   |
| nottinghamshire              | England               | Nottinghamshire                        | sPBpS   |
| nova-scotia                  | Canada                | Nova Scotia                            | sPBpS   |
| nunavut                      | Canada                | Nunavut                                | sPBpS   |
| oberbayern                   | Bayern                | Oberbayern                             | sPBpS   |
| oberfranken                  | Bayern                | Oberfranken                            | sPBpS   |
| oberpfalz                    | Bayern                | Oberpfalz                              | sPBpS   |
| ohio                         | North America         | Ohio                                   | sPBpS   |
| oklahoma                     | North America         | Oklahoma                               | sPBpS   |
| ontario                      | Canada                | Ontario                                | sPBpS   |
| opolskie                     | Poland                | Województwo opolskie(Opole Voivodeship) | sPBpS   |
| oregon                       | North America         | Oregon                                 | sPBpS   |
| oxfordshire                  | England               | Oxfordshire                            | sPBpS   |
| pakistan                     | Asia                  | Pakistan                               | sPBpS   |
| paraguay                     | South America         | Paraguay                               | sPBpS   |
| pays-de-la-loire             | France                | Pays de la Loire                       | sPBpS   |
| pennsylvania                 | North America         | Pennsylvania                           | sPBpS   |
| peru                         | South America         | Peru                                   | sPBpS   |
| philippines                  | Asia                  | Philippines                            | sPBpS   |
| picardie                     | France                | Picardie                               | sPBpS   |
| podkarpackie                 | Poland                | Województwo podkarpackie(Subcarpathian Voivodeship) | sPBpS   |
| podlaskie                    | Poland                | Województwo podlaskie(Podlaskie Voivodeship) | sPBpS   |
| poitou-charentes             | France                | Poitou-Charentes                       | sPBpS   |
| poland                       | Europe                | Poland                                 | sPBp    |
| pomorskie                    | Poland                | Województwo pomorskie(Pomeranian Voivodeship) | sPBpS   |
| portugal                     | Europe                | Portugal                               | sPBpS   |
| prince-edward-island         | Canada                | Prince Edward Island                   | sPBpS   |
| provence-alpes-cote-d-azur   | France                | Provence Alpes-Cote-d'Azur             | sPBpS   |
| quebec                       | Canada                | Quebec                                 | sPBpS   |
| reunion                      | France                | Reunion                                | sPBp    |
| rheinland-pfalz              | Germany               | Rheinland-Pfalz                        | sPBpS   |
| rhode-island                 | North America         | Rhode Island                           | sPBpS   |
| rhone-alpes                  | France                | Rhone-Alpes                            | sPBpS   |
| romania                      | Europe                | Romania                                | sPBpS   |
| russia                       |                       | Russian Federation                     | sPBp    |
| saarland                     | Germany               | Saarland                               | sPBpS   |
| sachsen                      | Germany               | Sachsen                                | sPBpS   |
| sachsen-anhalt               | Germany               | Sachsen-Anhalt                         | sPBpS   |
| saskatchewan                 | Canada                | Saskatchewan                           | sPBpS   |
| schleswig-holstein           | Germany               | Schleswig-Holstein                     | sPBpS   |
| schwaben                     | Bayern                | Schwaben                               | sPBpS   |
| scotland                     | Great Britain         | Scotland                               | sPBpS   |
| serbia                       | Europe                | Serbia                                 | sPBpS   |
| shropshire                   | England               | Shropshire                             | sPBpS   |
| siberian-fed-district        | Russian Federation    | Siberian Federal District              | sPBpS   |
| sierra-leone                 | Africa                | Sierra Leone                           | sPBpS   |
| slaskie                      | Poland                | Województwo śląskie(Silesian Voivodeship) | sPBpS   |
| slovakia                     | Europe                | Slovakia                               | sPBpS   |
| slovenia                     | Europe                | Slovenia                               | sPBpS   |
| somalia                      | Africa                | Somalia                                | sPBp    |
| somerset                     | England               | Somerset                               | sPBpS   |
| south-africa-and-lesotho     | Africa                | South Africa (includes Lesotho)        | sPBpS   |
| south-america                |                       | South America                          | sPBp    |
| south-carolina               | North America         | South Carolina                         | sPBpS   |
| south-dakota                 | North America         | South Dakota                           | sPBpS   |
| south-fed-district           | Russian Federation    | South Federal District                 | sPBpS   |
| south-korea                  | Asia                  | South Korea                            | sPBpS   |
| south-yorkshire              | England               | South Yorkshire                        | sPBpS   |
| spain                        | Europe                | Spain                                  | sPBpS   |
| sri-lanka                    | Asia                  | Sri Lanka                              | sPBpS   |
| staffordshire                | England               | Staffordshire                          | sPBpS   |
| stuttgart-regbez             | Baden-Württemberg     | Regierungsbezirk Stuttgart             | sPBpS   |
| sud                          | Italy                 | Sud                                    | sPBpS   |
| suffolk                      | England               | Suffolk                                | sPBpS   |
| suriname                     | South America         | suriname                               | sPBpS   |
| surrey                       | England               | Surrey                                 | sPBpS   |
| sweden                       | Europe                | Sweden                                 | sPBpS   |
| swietokrzyskie               | Poland                | Województwo                            | sPBpS   |
|                              |                       | świętokrzyskie(Świętokrzyskie Voivodeship) |         |
| switzerland                  | Europe                | Switzerland                            | sPBpS   |
| syria                        | Asia                  | Syria                                  | sPBpS   |
| taiwan                       | Asia                  | Taiwan                                 | sPBpS   |
| tajikistan                   | Asia                  | Tajikistan                             | sPBp    |
| tanzania                     | Africa                | Tanzania                               | sPBpS   |
| tennessee                    | North America         | Tennessee                              | sPBpS   |
| texas                        | North America         | Texas                                  | sPBpS   |
| thailand                     | Asia                  | Thailand                               | sPBpS   |
| thueringen                   | Germany               | Thüringen                              | sPBpS   |
| tuebingen-regbez             | Baden-Württemberg     | Regierungsbezirk Tübingen              | sPBpS   |
| turkey                       | Europe                | Turkey                                 | sPBpS   |
| turkmenistan                 | Asia                  | Turkmenistan                           | sPBpS   |
| ukraine                      | Europe                | Ukraine                                | sPBpS   |
| unterfranken                 | Bayern                | Unterfranken                           | sPBpS   |
| ural-fed-district            | Russian Federation    | Ural Federal District                  | sPBpS   |
| uruguay                      | South America         | Uruguay                                | sPBpS   |
| us-midwest                   | North America         | US Midwest                             | sPBp    |
| us-northeast                 | North America         | US Northeast                           | sPBp    |
| us-pacific                   | North America         | US Pacific                             | sPBp    |
| us-south                     | North America         | US South                               | sPBp    |
| us-west                      | North America         | US West                                | sPBp    |
| utah                         | North America         | Utah                                   | sPBpS   |
| uzbekistan                   | Asia                  | Uzbekistan                             | sPBpS   |
| vermont                      | North America         | Vermont                                | sPBpS   |
| vietnam                      | Asia                  | Vietnam                                | sPBpS   |
| virginia                     | North America         | Virginia                               | sPBpS   |
| volga-fed-district           | Russian Federation    | Volga Federal District                 | sPBpS   |
| wales                        | Great Britain         | Wales                                  | sPBpS   |
| warminsko-mazurskie          | Poland                | Województwo warmińsko-mazurskie(Warmian-Masurian Voivodeship) | sPBpS   |
| washington                   | North America         | Washington                             | sPBpS   |
| west-midlands                | England               | West Midlands                          | sPBpS   |
| west-sussex                  | England               | West Sussex                            | sPBpS   |
| west-virginia                | North America         | West Virginia                          | sPBpS   |
| west-yorkshire               | England               | West Yorkshire                         | sPBpS   |
| wielkopolskie                | Poland                | Województwo wielkopolskie(Greater Poland Voivodeship) | sPBpS   |
| wiltshire                    | England               | Wiltshire                              | sPBpS   |
| wisconsin                    | North America         | Wisconsin                              | sPBpS   |
| wyoming                      | North America         | Wyoming                                | sPBpS   |
| yukon                        | Canada                | Yukon                                  | sPBpS   |
| zachodniopomorskie           | Poland                | Województwo zachodniopomorskie(West Pomeranian Voivodeship) | sPBpS   |
