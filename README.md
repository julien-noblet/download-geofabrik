# download-geofabrik

## Usage
```shell
./download-geofabrik element
```
where ```element``` is one of geofabrik's files.
```shell
./download-geofabrik -h
Usage of ./download-geofabrik:
  -config="./geofabrik.yml": Config for downloading OSMFiles
  -n=false: Don't download (testing only)
  -list=false: list all elements
  -osm.bz2=false: Download osm.bz2 if available
  -osm.pbf=false: Download osm.pbf (default)
  -shp.zip=false: Download shp.zip if available
  -update=false: Update geofabrik.yml from github
```

## List of elements
|          SHORTNAME           |         IS IN         |           LONG NAME            | FORMATS |
|------------------------------|-----------------------|--------------------------------|---------|
| africa                       |                       | Africa                         | pb      |
| albania                      | Europe                | Albania                        | pbs     |
| alberta                      | Canada                | Alberta                        | pbs     |
| alps                         | Europe                | Alps                           | pb      |
| alsace                       | France                | Alsace                         | pbs     |
| andorra                      | Europe                | Andorra                        | pbs     |
| antarctica                   |                       | Antarctica                     | pbs     |
| aquitaine                    | France                | Aquitaine                      | pbs     |
| arnsberg-regbez              | Nordrhein-Westfalen   | Regierungsbezirk Arnsberg      | pbs     |
| asia                         |                       | Asia                           | pb      |
| australia                    | Australia and Oceania | Australia                      | pb      |
| australia-oceania            |                       | Australia and Oceania          | pb      |
| austria                      | Europe                | Austria                        | pbs     |
| auvergne                     | France                | Auvergne                       | pbs     |
| azerbaijan                   | Asia                  | Azerbaijan                     | pbs     |
| azores                       | Europe                | Azores                         | pbs     |
| baden-wuerttemberg           | Germany               | Baden-Württemberg              | pbs     |
| bangladesh                   | Asia                  | Bangladesh                     | pbs     |
| basse-normandie              | France                | Basse-Normandie                | pbs     |
| bayern                       | Germany               | Bayern                         | pbs     |
| belarus                      | Europe                | Belarus                        | pbs     |
| belgium                      | Europe                | Belgium                        | pbs     |
| belize                       | Central America       | Belize                         | pbs     |
| berlin                       | Germany               | Berlin                         | pbs     |
| bosnia-herzegovina           | Europe                | Bosnia-Herzegovina             | pbs     |
| botswana                     | Africa                | Botswana                       | pbs     |
| bourgogne                    | France                | Bourgogne                      | pbs     |
| brandenburg                  | Germany               | Brandenburg                    | pbs     |
| bremen                       | Germany               | Bremen                         | pbs     |
| bretagne                     | France                | Bretagne                       | pbs     |
| british-columbia             | Canada                | British Columbia               | pbs     |
| british-isles                | Europe                | British Isles                  | pbs     |
| buckinghamshire              | England               | Buckinghamshire                | pbs     |
| bulgaria                     | Europe                | Bulgaria                       | pbs     |
| burkina-faso                 | Africa                | Burkina Faso                   | pb      |
| cambridgeshire               | England               | Cambridgeshire                 | pbs     |
| cameroon                     | Africa                | Cameroon                       | pb      |
| canada                       | North America         | Canada                         | pb      |
| canary-islands               | Africa                | Canary Islands                 | pbs     |
| central-america              |                       | Central America                | pb      |
| centre                       | France                | Centre                         | pbs     |
| champagne-ardenne            | France                | Champagne Ardenne              | pbs     |
| cheshire                     | England               | Cheshire                       | pbs     |
| china                        | Asia                  | China                          | pbs     |
| congo-democratic-republic    | Africa                | Congo (Democratic Republic)    | pb      |
| cornwall                     | England               | Cornwall                       | pbs     |
| corse                        | France                | Corse                          | pbs     |
| croatia                      | Europe                | Croatia                        | pbs     |
| cuba                         | Central America       | Cuba                           | pbs     |
| cumbria                      | England               | Cumbria                        | pbs     |
| cyprus                       | Europe                | Cyprus                         | pbs     |
| czech-republic               | Europe                | Czech Republic                 | pbs     |
| dach                         | Europe                | Germany, Austria, Switzerland  | pb      |
| denmark                      | Europe                | Denmark                        | pbs     |
| derbyshire                   | England               | Derbyshire                     | pbs     |
| detmold-regbez               | Nordrhein-Westfalen   | Regierungsbezirk Detmold       | pbs     |
| devon                        | England               | Devon                          | pbs     |
| dorset                       | England               | Dorset                         | pbs     |
| duesseldorf-regbez           | Nordrhein-Westfalen   | Regierungsbezirk Düsseldorf    | pbs     |
| east-sussex                  | England               | East Sussex                    | pbs     |
| east-yorkshire-with-hull     | England               | East Yorkshire with Hull       | pbs     |
| egypt                        | Africa                | Egypt                          | pb      |
| england                      | Great Britain         | England                        | pbs     |
| essex                        | England               | Essex                          | pbs     |
| estonia                      | Europe                | Estonia                        | pbs     |
| ethiopia                     | Africa                | Ethiopia                       | pbs     |
| europe                       |                       | Europe                         | pb      |
| faroe-islands                | Europe                | Faroe Islands                  | pbs     |
| fiji                         | Australia and Oceania | Fiji                           | pbs     |
| finland                      | Europe                | Finland                        | pbs     |
| france                       | Europe                | France                         | pb      |
| franche-comte                | France                | Franche Comte                  | pbs     |
| freiburg-regbez              | Baden-Württemberg     | Regierungsbezirk Freiburg      | pbs     |
| gcc-states                   | Asia                  | GCC States                     | pbs     |
| georgia                      | Europe                | Georgia (Eastern Europe)       | pbs     |
| germany                      | Europe                | Germany                        | pb      |
| gloucestershire              | England               | Gloucestershire                | pbs     |
| great-britain                | Europe                | Great Britain                  | pbs     |
| greater-london               | England               | Greater London                 | pbs     |
| greater-manchester           | England               | Greater Manchester             | pbs     |
| greece                       | Europe                | Greece                         | pbs     |
| guadeloupe                   | France                | Guadeloupe                     | pb      |
| guatemala                    | Central America       | Guatemala                      | pbs     |
| guinea                       | Africa                | Guinea                         | pbs     |
| guinea-bissau                | Africa                | Guinea-Bissau                  | pb      |
| guyane                       | France                | Guyane                         | pb      |
| haiti-and-domrep             | Central America       | Haiti and Dominican Republic   | pbs     |
| hamburg                      | Germany               | Hamburg                        | pbs     |
| hampshire                    | England               | Hampshire                      | pbs     |
| haute-normandie              | France                | Haute-Normandie                | pbs     |
| herefordshire                | England               | Herefordshire                  | pbs     |
| hertfordshire                | England               | Hertfordshire                  | pbs     |
| hessen                       | Germany               | Hessen                         | pbs     |
| hungary                      | Europe                | Hungary                        | pbs     |
| iceland                      | Europe                | Iceland                        | pbs     |
| ile-de-france                | France                | Ile-de-France                  | pbs     |
| india                        | Asia                  | India                          | pbs     |
| indonesia                    | Asia                  | Indonesia                      | pbs     |
| irak                         | Asia                  | Irak                           | pbs     |
| iran                         | Asia                  | Iran                           | pbs     |
| ireland-and-northern-ireland | Europe                | Ireland and Northern Ireland   | pbs     |
| isle-of-man                  | Europe                | Isle of Man                    | pbs     |
| isle-of-wight                | England               | Isle of Wight                  | pbs     |
| israel-and-palestine         | Asia                  | Israel and Palestine           | pbs     |
| italy                        | Europe                | Italy                          | pbs     |
| ivory-coast                  | Africa                | Ivory Coast                    | pbs     |
| japan                        | Asia                  | Japan                          | pbs     |
| jordan                       | Asia                  | Jordan                         | pbs     |
| karlsruhe-regbez             | Baden-Württemberg     | Regierungsbezirk Karlsruhe     | pbs     |
| kazakhstan                   | Asia                  | Kazakhstan                     | pbs     |
| kent                         | England               | Kent                           | pbs     |
| kenya                        | Africa                | Kenya                          | pbs     |
| koeln-regbez                 | Nordrhein-Westfalen   | Regierungsbezirk Köln          | pbs     |
| kosovo                       | Europe                | Kosovo                         | pbs     |
| kyrgyzstan                   | Asia                  | Kyrgyzstan                     | pbs     |
| lancashire                   | England               | Lancashire                     | pbs     |
| languedoc-roussillon         | France                | Languedoc-Roussillon           | pbs     |
| latvia                       | Europe                | Latvia                         | pbs     |
| lebanon                      | Asia                  | Lebanon                        | pbs     |
| leicestershire               | England               | Leicestershire                 | pbs     |
| lesotho                      | Africa                | Lesotho                        | pbs     |
| liberia                      | Africa                | Liberia                        | pbs     |
| libya                        | Africa                | Libya                          | pbs     |
| liechtenstein                | Europe                | Liechtenstein                  | pbs     |
| limousin                     | France                | Limousin                       | pbs     |
| lithuania                    | Europe                | Lithuania                      | pbs     |
| lorraine                     | France                | Lorraine                       | pbs     |
| luxembourg                   | Europe                | Luxembourg                     | pbs     |
| macedonia                    | Europe                | Macedonia                      | pbs     |
| madagascar                   | Africa                | Madagascar                     | pbs     |
| malaysia-singapore-brunei    | Asia                  | Malaysia, Singapore, and Brunei| pb      |
| malta                        | Europe                | Malta                          | pbs     |
| manitoba                     | Canada                | Manitoba                       | pbs     |
| martinique                   | France                | Martinique                     | pb      |
| mayotte                      | France                | Mayotte                        | pb      |
| mecklenburg-vorpommern       | Germany               | Mecklenburg-Vorpommern         | pbs     |
| midi-pyrenees                | France                | Midi-Pyrenees                  | pbs     |
| mittelfranken                | Bayern                | Mittelfranken                  | pbs     |
| moldova                      | Europe                | Moldova                        | pbs     |
| monaco                       | Europe                | Monaco                         | pbs     |
| mongolia                     | Asia                  | Mongolia                       | pbs     |
| montenegro                   | Europe                | Montenegro                     | pbs     |
| morocco                      | Africa                | Morocco                        | pbs     |
| muenster-regbez              | Nordrhein-Westfalen   | Regierungsbezirk Münster       | pbs     |
| nepal                        | Asia                  | Nepal                          | pbs     |
| netherlands                  | Europe                | Netherlands                    | pbs     |
| new-brunswick                | Canada                | New Brunswick                  | pbs     |
| new-caledonia                | Australia and Oceania | New Caledonia                  | pbs     |
| new-zealand                  | Australia and Oceania | New Zealand                    | pbs     |
| newfoundland-and-labrador    | Canada                | Newfoundland and Labrador      | pbs     |
| niederbayern                 | Bayern                | Niederbayern                   | pbs     |
| niedersachsen                | Germany               | Niedersachsen                  | pbs     |
| nigeria                      | Africa                | Nigeria                        | pb      |
| nord-pas-de-calais           | France                | Nord-Pas-de-Calais             | pbs     |
| nordrhein-westfalen          | Germany               | Nordrhein-Westfalen            | pbs     |
| norfolk                      | England               | Norfolk                        | pbs     |
| north-america                |                       | North America                  | pb      |
| north-korea                  | Asia                  | North Korea                    | pbs     |
| north-yorkshire              | England               | North Yorkshire                | pbs     |
| northwest-territories        | Canada                | Northwest Territories          | pbs     |
| norway                       | Europe                | Norway                         | pbs     |
| nottinghamshire              | England               | Nottinghamshire                | pbs     |
| nova-scotia                  | Canada                | Nova Scotia                    | pbs     |
| nunavut                      | Canada                | Nunavut                        | pbs     |
| oberbayern                   | Bayern                | Oberbayern                     | pbs     |
| oberfranken                  | Bayern                | Oberfranken                    | pbs     |
| oberpfalz                    | Bayern                | Oberpfalz                      | pbs     |
| ontario                      | Canada                | Ontario                        | pbs     |
| oxfordshire                  | England               | Oxfordshire                    | pbs     |
| pakistan                     | Asia                  | Pakistan                       | pbs     |
| pays-de-la-loire             | France                | Pays de la Loire               | pbs     |
| philippines                  | Asia                  | Philippines                    | pbs     |
| picardie                     | France                | Picardie                       | pbs     |
| poitou-charentes             | France                | Poitou-Charentes               | pbs     |
| poland                       | Europe                | Poland                         | pbs     |
| portugal                     | Europe                | Portugal                       | pbs     |
| prince-edward-island         | Canada                | Prince Edward Island           | pbs     |
| provence-alpes-cote-d-azur   | France                | Provence Alpes-Cote-d'Azur     | pbs     |
| quebec                       | Canada                | Quebec                         | pbs     |
| reunion                      | France                | Reunion                        | pb      |
| rheinland-pfalz              | Germany               | Rheinland-Pfalz                | pbs     |
| rhone-alpes                  | France                | Rhone-Alpes                    | pbs     |
| romania                      | Europe                | Romania                        | pbs     |
| russia-asian-part            | Asia                  | Russia (Asian part)            | pbs     |
| russia-european-part         | Europe                | Russia (European part)         | pbs     |
| saarland                     | Germany               | Saarland                       | pbs     |
| sachsen                      | Germany               | Sachsen                        | pbs     |
| sachsen-anhalt               | Germany               | Sachsen-Anhalt                 | pbs     |
| saskatchewan                 | Canada                | Saskatchewan                   | pbs     |
| schleswig-holstein           | Germany               | Schleswig-Holstein             | pbs     |
| schwaben                     | Bayern                | Schwaben                       | pbs     |
| scotland                     | Great Britain         | Scotland                       | pbs     |
| serbia                       | Europe                | Serbia                         | pbs     |
| shropshire                   | England               | Shropshire                     | pbs     |
| sierra-leone                 | Africa                | Sierra Leone                   | pbs     |
| slovakia                     | Europe                | Slovakia                       | pbs     |
| slovenia                     | Europe                | Slovenia                       | pbs     |
| somalia                      | Africa                | Somalia                        | pb      |
| somerset                     | England               | Somerset                       | pbs     |
| south-africa-and-lesotho     | Africa                | South Africa (includes         | pbs     |
|                              |                       | Lesotho)                       |         |
| south-korea                  | Asia                  | South Korea                    | pbs     |
| south-yorkshire              | England               | South Yorkshire                | pbs     |
| spain                        | Europe                | Spain                          | pbs     |
| sri-lanka                    | Asia                  | Sri Lanka                      | pbs     |
| staffordshire                | England               | Staffordshire                  | pbs     |
| stuttgart-regbez             | Baden-Württemberg     | Regierungsbezirk Stuttgart     | pbs     |
| suffolk                      | England               | Suffolk                        | pbs     |
| surrey                       | England               | Surrey                         | pbs     |
| sweden                       | Europe                | Sweden                         | pbs     |
| switzerland                  | Europe                | Switzerland                    | pbs     |
| syria                        | Asia                  | Syria                          | pbs     |
| taiwan                       | Asia                  | Taiwan                         | pbs     |
| tajikistan                   | Asia                  | Tajikistan                     | pb      |
| tanzania                     | Africa                | Tanzania                       | pbs     |
| thailand                     | Asia                  | Thailand                       | pbs     |
| thueringen                   | Germany               | Thüringen                      | pbs     |
| tuebingen-regbez             | Baden-Württemberg     | Regierungsbezirk Tübingen      | pbs     |
| turkey                       | Europe                | Turkey                         | pbs     |
| turkmenistan                 | Asia                  | Turkmenistan                   | pbs     |
| ukraine                      | Europe                | Ukraine                        | pbs     |
| unterfranken                 | Bayern                | Unterfranken                   | pbs     |
| uzbekistan                   | Asia                  | Uzbekistan                     | pbs     |
| vietnam                      | Asia                  | Vietnam                        | pbs     |
| wales                        | Great Britain         | Wales                          | pbs     |
| west-midlands                | England               | West Midlands                  | pbs     |
| west-sussex                  | England               | West Sussex                    | pbs     |
| west-yorkshire               | England               | West Yorkshire                 | pbs     |
| wiltshire                    | England               | Wiltshire                      | pbs     |
| yukon                        | Canada                | Yukon                          | pbs     |
