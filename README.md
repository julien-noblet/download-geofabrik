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
|         SHORTNAME          |         IS IN         |           LONG NAME            | FORMATS |
|----------------------------|-----------------------|--------------------------------|---------|
| africa                     |                       | Africa                         | pb      |
| albania                    | Europe                | Albania                        | pbs     |
| alsace                     | France                | Alsace                         | pbs     |
| andorra                    | Europe                | Andorra                        | pbs     |
| antarctica                 |                       | Antarctica                     | pbs     |
| aquitaine                  | France                | Aquitaine                      | pbs     |
| asia                       |                       | Asia                           | pb      |
| australia                  | Australia and Oceania | Australia                      | pb      |
| australia-oceania          |                       | Australia and Oceania          | pb      |
| austria                    | Europe                | Austria                        | pbs     |
| auvergne                   | France                | Auvergne                       | pbs     |
| azerbaijan                 | Asia                  | Azerbaijan                     | pbs     |
| azores                     | Europe                | Azores                         | pbs     |
| bangladesh                 | Asia                  | Bangladesh                     | pbs     |
| basse-normandie            | France                | Basse-Normandie                | pbs     |
| belarus                    | Europe                | Belarus                        | pbs     |
| belgium                    | Europe                | Belgium                        | pbs     |
| belize                     | Central America       | Belize                         | pbs     |
| bosnia-herzegovina         | Europe                | Bosnia-Herzegovina             | pbs     |
| botswana                   | Africa                | Botswana                       | pbs     |
| bourgogne                  | France                | Bourgogne                      | pbs     |
| bretagne                   | France                | Bretagne                       | pbs     |
| bulgaria                   | Europe                | Bulgaria                       | pbs     |
| burkina-faso               | Africa                | Burkina Faso                   | pb      |
| cameroon                   | Africa                | Cameroon                       | pb      |
| canary-islands             | Africa                | Canary Islands                 | pbs     |
| central-america            |                       | Central America                | pb      |
| centre                     | France                | Centre                         | pbs     |
| champagne-ardenne          | France                | Champagne Ardenne              | pbs     |
| china                      | Asia                  | China                          | pbs     |
| congo-democratic-republic  | Africa                | Congo (Democratic Republic)    | pb      |
| corse                      | France                | Corse                          | pbs     |
| croatia                    | Europe                | Croatia                        | pbs     |
| cuba                       | Central America       | Cuba                           | pbs     |
| cyprus                     | Europe                | Cyprus                         | pbs     |
| czech-republic             | Europe                | Czech Republic                 | pbs     |
| denmark                    | Europe                | Denmark                        | pbs     |
| egypt                      | Africa                | Egypt                          | pb      |
| estonia                    | Europe                | Estonia                        | pbs     |
| ethiopia                   | Africa                | Ethiopia                       | pbs     |
| europe                     |                       | Europe                         | pb      |
| faroe-islands              | Europe                | Faroe Islands                  | pbs     |
| fiji                       | Australia and Oceania | Fiji                           | pbs     |
| finland                    | Europe                | Finland                        | pbs     |
| france                     | Europe                | France                         | pb      |
| franche-comte              | France                | Franche Comte                  | pbs     |
| gcc-states                 | Asia                  | GCC States                     | pbs     |
| guadeloupe                 | France                | Guadeloupe                     | pb      |
| guatemala                  | Central America       | Guatemala                      | pbs     |
| guinea                     | Africa                | Guinea                         | pbs     |
| guinea-bissau              | Africa                | Guinea-Bissau                  | pb      |
| guyane                     | France                | Guyane                         | pb      |
| haiti-and-domrep           | Central America       | Haiti and Dominican Republic   | pbs     |
| haute-normandie            | France                | Haute-Normandie                | pbs     |
| ile-de-france              | France                | Ile-de-France                  | pbs     |
| india                      | Asia                  | India                          | pbs     |
| indonesia                  | Asia                  | Indonesia                      | pbs     |
| irak                       | Asia                  | Irak                           | pbs     |
| iran                       | Asia                  | Iran                           | pbs     |
| israel-and-palestine       | Asia                  | Israel and Palestine           | pbs     |
| ivory-coast                | Africa                | Ivory Coast                    | pbs     |
| japan                      | Asia                  | Japan                          | pbs     |
| jordan                     | Asia                  | Jordan                         | pbs     |
| kazakhstan                 | Asia                  | Kazakhstan                     | pbs     |
| kenya                      | Africa                | Kenya                          | pbs     |
| kyrgyzstan                 | Asia                  | Kyrgyzstan                     | pbs     |
| languedoc-roussillon       | France                | Languedoc-Roussillon           | pbs     |
| lebanon                    | Asia                  | Lebanon                        | pbs     |
| lesotho                    | Africa                | Lesotho                        | pbs     |
| liberia                    | Africa                | Liberia                        | pbs     |
| libya                      | Africa                | Libya                          | pbs     |
| limousin                   | France                | Limousin                       | pbs     |
| lorraine                   | France                | Lorraine                       | pbs     |
| madagascar                 | Africa                | Madagascar                     | pbs     |
| malaysia-singapore-brunei  | Asia                  | Malaysia, Singapore, and       | pb      |
|                            |                       | Brunei                         |         |
| martinique                 | France                | Martinique                     | pb      |
| mayotte                    | France                | Mayotte                        | pb      |
| midi-pyrenees              | France                | Midi-Pyrenees                  | pbs     |
| mongolia                   | Asia                  | Mongolia                       | pbs     |
| morocco                    | Africa                | Morocco                        | pbs     |
| nepal                      | Asia                  | Nepal                          | pbs     |
| new-caledonia              | Australia and Oceania | New Caledonia                  | pbs     |
| new-zealand                | Australia and Oceania | New Zealand                    | pbs     |
| nigeria                    | Africa                | Nigeria                        | pb      |
| nord-pas-de-calais         | France                | Nord-Pas-de-Calais             | pbs     |
| north-korea                | Asia                  | North Korea                    | pbs     |
| pakistan                   | Asia                  | Pakistan                       | pbs     |
| pays-de-la-loire           | France                | Pays de la Loire               | pbs     |
| philippines                | Asia                  | Philippines                    | pbs     |
| picardie                   | France                | Picardie                       | pbs     |
| poitou-charentes           | France                | Poitou-Charentes               | pbs     |
| provence-alpes-cote-d-azur | France                | Provence Alpes-Cote-d'Azur     | pbs     |
| reunion                    | France                | Reunion                        | pb      |
| rhone-alpes                | France                | Rhone-Alpes                    | pbs     |
| russia-asian-part          | Asia                  | Russia (Asian part)            | pbs     |
| sierra-leone               | Africa                | Sierra Leone                   | pbs     |
| somalia                    | Africa                | Somalia                        | pb      |
| south-africa-and-lesotho   | Africa                | South Africa (includes Lesotho)| pbs     |
| south-korea                | Asia                  | South Korea                    | pbs     |
| sri-lanka                  | Asia                  | Sri Lanka                      | pbs     |
| syria                      | Asia                  | Syria                          | pbs     |
| taiwan                     | Asia                  | Taiwan                         | pbs     |
| tajikistan                 | Asia                  | Tajikistan                     | pb      |
| tanzania                   | Africa                | Tanzania                       | pbs     |
| thailand                   | Asia                  | Thailand                       | pbs     |
| turkmenistan               | Asia                  | Turkmenistan                   | pbs     |
| uzbekistan                 | Asia                  | Uzbekistan                     | pbs     |
| vietnam                    | Asia                  | Vietnam                        | pbs     |
