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
|          SHORTNAME           |         IS IN         |           LONG NAME            | FORMATS |
|------------------------------|-----------------------|--------------------------------|---------|
| africa                       |                          | Africa                         | pb      |
| alabama                      | United States of America | Alabama                        | pbs     |
| alaska                       | United States of America | Alaska                         | pbs     |
| albania                      | Europe                   | Albania                        | pbs     |
| alberta                      | Canada                   | Alberta                        | pbs     |
| alps                         | Europe                   | Alps                           | pb      |
| alsace                       | France                   | Alsace                         | pbs     |
| andorra                      | Europe                   | Andorra                        | pbs     |
| antarctica                   |                          | Antarctica                     | pbs     |
| aquitaine                    | France                   | Aquitaine                      | pbs     |
| argentina                    | South America            | Argentina                      | pbs     |
| arizona                      | United States of America | Arizona                        | pbs     |
| arkansas                     | United States of America | Arkansas                       | pbs     |
| arnsberg-regbez              | Nordrhein-Westfalen      | Regierungsbezirk Arnsberg      | pbs     |
| asia                         |                          | Asia                           | pb      |
| australia                    | Australia and Oceania    | Australia                      | pb      |
| australia-oceania            |                          | Australia and Oceania          | pb      |
| austria                      | Europe                   | Austria                        | pbs     |
| auvergne                     | France                   | Auvergne                       | pbs     |
| azerbaijan                   | Asia                     | Azerbaijan                     | pbs     |
| azores                       | Europe                   | Azores                         | pbs     |
| baden-wuerttemberg           | Germany                  | Baden-Württemberg              | pbs     |
| bangladesh                   | Asia                     | Bangladesh                     | pbs     |
| basse-normandie              | France                   | Basse-Normandie                | pbs     |
| bayern                       | Germany                  | Bayern                         | pbs     |
| belarus                      | Europe                   | Belarus                        | pbs     |
| belgium                      | Europe                   | Belgium                        | pbs     |
| belize                       | Central America          | Belize                         | pbs     |
| berlin                       | Germany                  | Berlin                         | pbs     |
| bolivia                      | South America            | Bolivia                        | pbs     |
| bosnia-herzegovina           | Europe                   | Bosnia-Herzegovina             | pbs     |
| botswana                     | Africa                   | Botswana                       | pbs     |
| bourgogne                    | France                   | Bourgogne                      | pbs     |
| brandenburg                  | Germany                  | Brandenburg                    | pbs     |
| brazil                       | South America            | Brazil                         | pbs     |
| bremen                       | Germany                  | Bremen                         | pbs     |
| bretagne                     | France                   | Bretagne                       | pbs     |
| british-columbia             | Canada                   | British Columbia               | pbs     |
| british-isles                | Europe                   | British Isles                  | pbs     |
| buckinghamshire              | England                  | Buckinghamshire                | pbs     |
| bulgaria                     | Europe                   | Bulgaria                       | pbs     |
| burkina-faso                 | Africa                   | Burkina Faso                   | pb      |
| california                   | United States of America | California                     | pbs     |
| cambridgeshire               | England                  | Cambridgeshire                 | pbs     |
| cameroon                     | Africa                   | Cameroon                       | pb      |
| canada                       | North America            | Canada                         | pb      |
| canary-islands               | Africa                   | Canary Islands                 | pbs     |
| central-america              |                          | Central America                | pb      |
| centre                       | France                   | Centre                         | pbs     |
| champagne-ardenne            | France                   | Champagne Ardenne              | pbs     |
| cheshire                     | England                  | Cheshire                       | pbs     |
| chile                        | South America            | Chile                          | pbs     |
| china                        | Asia                     | China                          | pbs     |
| colombia                     | South America            | Colombia                       | pbs     |
| colorado                     | United States of America | Colorado                       | pbs     |
| congo-democratic-republic    | Africa                   | Congo (Democratic Republic)    | pb      |
| connecticut                  | United States of America | Connecticut                    | pbs     |
| cornwall                     | England                  | Cornwall                       | pbs     |
| corse                        | France                   | Corse                          | pbs     |
| croatia                      | Europe                   | Croatia                        | pbs     |
| cuba                         | Central America          | Cuba                           | pbs     |
| cumbria                      | England                  | Cumbria                        | pbs     |
| cyprus                       | Europe                   | Cyprus                         | pbs     |
| czech-republic               | Europe                   | Czech Republic                 | pbs     |
| dach                         | Europe                   | Germany, Austria, Switzerland  | pb      |
| delaware                     | United States of America | Delaware                       | pbs     |
| denmark                      | Europe                   | Denmark                        | pbs     |
| derbyshire                   | England                  | Derbyshire                     | pbs     |
| detmold-regbez               | Nordrhein-Westfalen      | Regierungsbezirk Detmold       | pbs     |
| devon                        | England                  | Devon                          | pbs     |
| district-of-columbia         | United States of America | District of Columbia           | pbs     |
| dorset                       | England                  | Dorset                         | pbs     |
| duesseldorf-regbez           | Nordrhein-Westfalen      | Regierungsbezirk Düsseldorf    | pbs     |
| east-sussex                  | England                  | East Sussex                    | pbs     |
| east-yorkshire-with-hull     | England                  | East Yorkshire with Hull       | pbs     |
| ecuador                      | South America            | Ecuador                        | pb      |
| egypt                        | Africa                   | Egypt                          | pb      |
| england                      | Great Britain            | England                        | pbs     |
| essex                        | England                  | Essex                          | pbs     |
| estonia                      | Europe                   | Estonia                        | pbs     |
| ethiopia                     | Africa                   | Ethiopia                       | pbs     |
| europe                       |                          | Europe                         | pb      |
| faroe-islands                | Europe                   | Faroe Islands                  | pbs     |
| fiji                         | Australia and Oceania    | Fiji                           | pbs     |
| finland                      | Europe                   | Finland                        | pbs     |
| florida                      | United States of America | Florida                        | pbs     |
| france                       | Europe                   | France                         | pb      |
| franche-comte                | France                   | Franche Comte                  | pbs     |
| freiburg-regbez              | Baden-Württemberg        | Regierungsbezirk Freiburg      | pbs     |
| gcc-states                   | Asia                     | GCC States                     | pbs     |
| georgia-eu                   | Europe                   | Georgia (Eastern Europe)       | pbs     |
| georgia-us                   | United States of America | Georgia (US State)             | pbs     |
| germany                      | Europe                   | Germany                        | pb      |
| gloucestershire              | England                  | Gloucestershire                | pbs     |
| great-britain                | Europe                   | Great Britain                  | pbs     |
| greater-london               | England                  | Greater London                 | pbs     |
| greater-manchester           | England                  | Greater Manchester             | pbs     |
| greece                       | Europe                   | Greece                         | pbs     |
| greenland                    | North America            | Greenland                      | pbs     |
| guadeloupe                   | France                   | Guadeloupe                     | pb      |
| guatemala                    | Central America          | Guatemala                      | pbs     |
| guinea                       | Africa                   | Guinea                         | pbs     |
| guinea-bissau                | Africa                   | Guinea-Bissau                  | pb      |
| guyane                       | France                   | Guyane                         | pb      |
| haiti-and-domrep             | Central America          | Haiti and Dominican Republic   | pbs     |
| hamburg                      | Germany                  | Hamburg                        | pbs     |
| hampshire                    | England                  | Hampshire                      | pbs     |
| haute-normandie              | France                   | Haute-Normandie                | pbs     |
| hawaii                       | United States of America | Hawaii                         | pbs     |
| herefordshire                | England                  | Herefordshire                  | pbs     |
| hertfordshire                | England                  | Hertfordshire                  | pbs     |
| hessen                       | Germany                  | Hessen                         | pbs     |
| hungary                      | Europe                   | Hungary                        | pbs     |
| iceland                      | Europe                   | Iceland                        | pbs     |
| idaho                        | United States of America | Idaho                          | pbs     |
| ile-de-france                | France                   | Ile-de-France                  | pbs     |
| illinois                     | United States of America | Illinois                       | pbs     |
| india                        | Asia                     | India                          | pbs     |
| indiana                      | United States of America | Indiana                        | pbs     |
| indonesia                    | Asia                     | Indonesia                      | pbs     |
| iowa                         | United States of America | Iowa                           | pbs     |
| irak                         | Asia                     | Irak                           | pbs     |
| iran                         | Asia                     | Iran                           | pbs     |
| ireland-and-northern-ireland | Europe                   | Ireland and Northern Ireland   | pbs     |
| isle-of-man                  | Europe                   | Isle of Man                    | pbs     |
| isle-of-wight                | England                  | Isle of Wight                  | pbs     |
| israel-and-palestine         | Asia                     | Israel and Palestine           | pbs     |
| italy                        | Europe                   | Italy                          | pbs     |
| ivory-coast                  | Africa                   | Ivory Coast                    | pbs     |
| japan                        | Asia                     | Japan                          | pbs     |
| jordan                       | Asia                     | Jordan                         | pbs     |
| kansas                       | United States of America | Kansas                         | pbs     |
| karlsruhe-regbez             | Baden-Württemberg        | Regierungsbezirk Karlsruhe     | pbs     |
| kazakhstan                   | Asia                     | Kazakhstan                     | pbs     |
| kent                         | England                  | Kent                           | pbs     |
| kentucky                     | United States of America | Kentucky                       | pbs     |
| kenya                        | Africa                   | Kenya                          | pbs     |
| koeln-regbez                 | Nordrhein-Westfalen      | Regierungsbezirk Köln          | pbs     |
| kosovo                       | Europe                   | Kosovo                         | pbs     |
| kyrgyzstan                   | Asia                     | Kyrgyzstan                     | pbs     |
| lancashire                   | England                  | Lancashire                     | pbs     |
| languedoc-roussillon         | France                   | Languedoc-Roussillon           | pbs     |
| latvia                       | Europe                   | Latvia                         | pbs     |
| lebanon                      | Asia                     | Lebanon                        | pbs     |
| leicestershire               | England                  | Leicestershire                 | pbs     |
| lesotho                      | Africa                   | Lesotho                        | pbs     |
| liberia                      | Africa                   | Liberia                        | pbs     |
| libya                        | Africa                   | Libya                          | pbs     |
| liechtenstein                | Europe                   | Liechtenstein                  | pbs     |
| limousin                     | France                   | Limousin                       | pbs     |
| lithuania                    | Europe                   | Lithuania                      | pbs     |
| lorraine                     | France                   | Lorraine                       | pbs     |
| louisiana                    | United States of America | Louisiana                      | pbs     |
| luxembourg                   | Europe                   | Luxembourg                     | pbs     |
| macedonia                    | Europe                   | Macedonia                      | pbs     |
| madagascar                   | Africa                   | Madagascar                     | pbs     |
| maine                        | United States of America | Maine                          | pbs     |
| malaysia-singapore-brunei    | Asia                     | Malaysia, Singapore, and Brunei| pb      |
| malta                        | Europe                   | Malta                          | pbs     |
| manitoba                     | Canada                   | Manitoba                       | pbs     |
| martinique                   | France                   | Martinique                     | pb      |
| maryland                     | United States of America | Maryland                       | pbs     |
| massachusetts                | United States of America | Massachusetts                  | pbs     |
| mayotte                      | France                   | Mayotte                        | pb      |
| mecklenburg-vorpommern       | Germany                  | Mecklenburg-Vorpommern         | pbs     |
| mexico                       | North America            | Mexico                         | pbs     |
| michigan                     | United States of America | Michigan                       | pbs     |
| midi-pyrenees                | France                   | Midi-Pyrenees                  | pbs     |
| minnesota                    | United States of America | Minnesota                      | pbs     |
| mississippi                  | United States of America | Mississippi                    | pbs     |
| missouri                     | United States of America | Missouri                       | pbs     |
| mittelfranken                | Bayern                   | Mittelfranken                  | pbs     |
| moldova                      | Europe                   | Moldova                        | pbs     |
| monaco                       | Europe                   | Monaco                         | pbs     |
| mongolia                     | Asia                     | Mongolia                       | pbs     |
| montana                      | United States of America | Montana                        | pbs     |
| montenegro                   | Europe                   | Montenegro                     | pbs     |
| morocco                      | Africa                   | Morocco                        | pbs     |
| muenster-regbez              | Nordrhein-Westfalen      | Regierungsbezirk Münster       | pbs     |
| nebraska                     | United States of America | Nebraska                       | pbs     |
| nepal                        | Asia                     | Nepal                          | pbs     |
| netherlands                  | Europe                   | Netherlands                    | pbs     |
| nevada                       | United States of America | Nevada                         | pbs     |
| new-brunswick                | Canada                   | New Brunswick                  | pbs     |
| new-caledonia                | Australia and Oceania    | New Caledonia                  | pbs     |
| new-hampshire                | United States of America | New Hampshire                  | pbs     |
| new-jersey                   | United States of America | New Jersey                     | pbs     |
| new-mexico                   | United States of America | New Mexico                     | pbs     |
| new-york                     | United States of America | New York                       | pbs     |
| new-zealand                  | Australia and Oceania    | New Zealand                    | pbs     |
| newfoundland-and-labrador    | Canada                   | Newfoundland and Labrador      | pbs     |
| niederbayern                 | Bayern                   | Niederbayern                   | pbs     |
| niedersachsen                | Germany                  | Niedersachsen                  | pbs     |
| nigeria                      | Africa                   | Nigeria                        | pb      |
| nord-pas-de-calais           | France                   | Nord-Pas-de-Calais             | pbs     |
| nordrhein-westfalen          | Germany                  | Nordrhein-Westfalen            | pbs     |
| norfolk                      | England                  | Norfolk                        | pbs     |
| north-america                |                          | North America                  | pb      |
| north-carolina               | United States of America | North Carolina                 | pbs     |
| north-dakota                 | United States of America | North Dakota                   | pbs     |
| north-korea                  | Asia                     | North Korea                    | pbs     |
| north-yorkshire              | England                  | North Yorkshire                | pbs     |
| northwest-territories        | Canada                   | Northwest Territories          | pbs     |
| norway                       | Europe                   | Norway                         | pbs     |
| nottinghamshire              | England                  | Nottinghamshire                | pbs     |
| nova-scotia                  | Canada                   | Nova Scotia                    | pbs     |
| nunavut                      | Canada                   | Nunavut                        | pbs     |
| oberbayern                   | Bayern                   | Oberbayern                     | pbs     |
| oberfranken                  | Bayern                   | Oberfranken                    | pbs     |
| oberpfalz                    | Bayern                   | Oberpfalz                      | pbs     |
| ohio                         | United States of America | Ohio                           | pbs     |
| oklahoma                     | United States of America | Oklahoma                       | pbs     |
| ontario                      | Canada                   | Ontario                        | pbs     |
| oregon                       | United States of America | Oregon                         | pbs     |
| oxfordshire                  | England                  | Oxfordshire                    | pbs     |
| pakistan                     | Asia                     | Pakistan                       | pbs     |
| paraguay                     | South America            | Paraguay                       | pbs     |
| pays-de-la-loire             | France                   | Pays de la Loire               | pbs     |
| pennsylvania                 | United States of America | Pennsylvania                   | pbs     |
| peru                         | South America            | Peru                           | pbs     |
| philippines                  | Asia                     | Philippines                    | pbs     |
| picardie                     | France                   | Picardie                       | pbs     |
| poitou-charentes             | France                   | Poitou-Charentes               | pbs     |
| poland                       | Europe                   | Poland                         | pbs     |
| portugal                     | Europe                   | Portugal                       | pbs     |
| prince-edward-island         | Canada                   | Prince Edward Island           | pbs     |
| provence-alpes-cote-d-azur   | France                   | Provence Alpes-Cote-d'Azur     | pbs     |
| quebec                       | Canada                   | Quebec                         | pbs     |
| reunion                      | France                   | Reunion                        | pb      |
| rheinland-pfalz              | Germany                  | Rheinland-Pfalz                | pbs     |
| rhode-island                 | United States of America | Rhode Island                   | pbs     |
| rhone-alpes                  | France                   | Rhone-Alpes                    | pbs     |
| romania                      | Europe                   | Romania                        | pbs     |
| russia-asian-part            | Asia                     | Russia (Asian part)            | pbs     |
| russia-european-part         | Europe                   | Russia (European part)         | pbs     |
| saarland                     | Germany                  | Saarland                       | pbs     |
| sachsen                      | Germany                  | Sachsen                        | pbs     |
| sachsen-anhalt               | Germany                  | Sachsen-Anhalt                 | pbs     |
| saskatchewan                 | Canada                   | Saskatchewan                   | pbs     |
| schleswig-holstein           | Germany                  | Schleswig-Holstein             | pbs     |
| schwaben                     | Bayern                   | Schwaben                       | pbs     |
| scotland                     | Great Britain            | Scotland                       | pbs     |
| serbia                       | Europe                   | Serbia                         | pbs     |
| shropshire                   | England                  | Shropshire                     | pbs     |
| sierra-leone                 | Africa                   | Sierra Leone                   | pbs     |
| slovakia                     | Europe                   | Slovakia                       | pbs     |
| slovenia                     | Europe                   | Slovenia                       | pbs     |
| somalia                      | Africa                   | Somalia                        | pb      |
| somerset                     | England                  | Somerset                       | pbs     |
| south-africa-and-lesotho     | Africa                   | South Africa (includes Lesotho)| pbs     |
| south-america                |                          | South America                  | pb      |
| south-carolina               | United States of America | South Carolina                 | pbs     |
| south-dakota                 | United States of America | South Dakota                   | pbs     |
| south-korea                  | Asia                     | South Korea                    | pbs     |
| south-yorkshire              | England                  | South Yorkshire                | pbs     |
| spain                        | Europe                   | Spain                          | pbs     |
| sri-lanka                    | Asia                     | Sri Lanka                      | pbs     |
| staffordshire                | England                  | Staffordshire                  | pbs     |
| stuttgart-regbez             | Baden-Württemberg        | Regierungsbezirk Stuttgart     | pbs     |
| suffolk                      | England                  | Suffolk                        | pbs     |
| surrey                       | England                  | Surrey                         | pbs     |
| sweden                       | Europe                   | Sweden                         | pbs     |
| switzerland                  | Europe                   | Switzerland                    | pbs     |
| syria                        | Asia                     | Syria                          | pbs     |
| taiwan                       | Asia                     | Taiwan                         | pbs     |
| tajikistan                   | Asia                     | Tajikistan                     | pb      |
| tanzania                     | Africa                   | Tanzania                       | pbs     |
| tennessee                    | United States of America | Tennessee                      | pbs     |
| texas                        | United States of America | Texas                          | pbs     |
| thailand                     | Asia                     | Thailand                       | pbs     |
| thueringen                   | Germany                  | Thüringen                      | pbs     |
| tuebingen-regbez             | Baden-Württemberg        | Regierungsbezirk Tübingen      | pbs     |
| turkey                       | Europe                   | Turkey                         | pbs     |
| turkmenistan                 | Asia                     | Turkmenistan                   | pbs     |
| ukraine                      | Europe                   | Ukraine                        | pbs     |
| unterfranken                 | Bayern                   | Unterfranken                   | pbs     |
| uruguay                      | South America            | Uruguay                        | pbs     |
| us                           | North America            | United States of America       |         |
| us-midwest                   | North America            | US Midwest                     | pb      |
| us-northeast                 | North America            | US Northeast                   | pb      |
| us-pacific                   | North America            | US Pacific                     | pb      |
| us-south                     | North America            | US South                       | pb      |
| us-west                      | North America            | US West                        | pb      |
| utah                         | United States of America | Utah                           | pbs     |
| uzbekistan                   | Asia                     | Uzbekistan                     | pbs     |
| vermont                      | United States of America | Vermont                        | pbs     |
| vietnam                      | Asia                     | Vietnam                        | pbs     |
| virginia                     | United States of America | Virginia                       | pbs     |
| wales                        | Great Britain            | Wales                          | pbs     |
| washington                   | United States of America | Washington                     | pbs     |
| west-midlands                | England                  | West Midlands                  | pbs     |
| west-sussex                  | England                  | West Sussex                    | pbs     |
| west-virginia                | United States of America | West Virginia                  | pbs     |
| west-yorkshire               | England                  | West Yorkshire                 | pbs     |
| wiltshire                    | England                  | Wiltshire                      | pbs     |
| wisconsin                    | United States of America | Wisconsin                      | pbs     |
| wyoming                      | United States of America | Wyoming                        | pbs     |
| yukon                        | Canada                   | Yukon                          | pbs     |
