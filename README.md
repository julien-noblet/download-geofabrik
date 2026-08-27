# download-geofabrik

[![Join the chat at https://gitter.im/julien-noblet/download-geofabrik](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/julien-noblet/download-geofabrik?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
![Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/julien-noblet/a509e15ea4734ca3e8e98f32ab5369c0/raw/7344619caf8ac5bce291793711071a9636536fce/coverage.json)
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

download-geofabrik is a CLI tool that downloads OpenStreetMap data from Geofabrik.

Usage:
  download-geofabrik [flags]
  download-geofabrik [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  download    Download element
  generate    Generate configuration file
  help        Help about any command
  list        Show elements available

Flags:
  -c, --config string    config file (default is geofabrik.yml)
  -h, --help             help for download-geofabrik
      --quiet            Quiet mode
  -s, --service string   Service to use (geofabrik, geofabrik-parse, openstreetmap.fr, geo2day, bbbike) (default "geofabrik")
      --verbose          Verbose mode
  -v, --version          version for download-geofabrik

Use "download-geofabrik [command] --help" for more information about a command.
```

## List of elements
| SHORT NAME                                  | IS IN                       | LONG NAME                                                             | FORMATS |
|---------------------------------------------|-----------------------------|-----------------------------------------------------------------------|---------|
| act                                         | Australia                   | Australian Capital Territory                                          | kHPpSs  |
| afghanistan                                 | Asia                        | Afghanistan                                                           | kHPpSs  |
| africa                                      |                             | Africa                                                                | kHPps   |
| albania                                     | Europe                      | Albania                                                               | kHPpSs  |
| alberta                                     | Canada                      | Alberta                                                               | kHPpSs  |
| algeria                                     | Africa                      | Algeria                                                               | kHPpSs  |
| alps                                        | Europe                      | Alps                                                                  | kHPps   |
| alsace                                      | France                      | Alsace                                                                | kHPpSs  |
| american-oceania                            | Australia and Oceania       | American Oceania                                                      | kHPpSs  |
| andalucia                                   | Spain                       | Andalucía                                                             | kHPpSs  |
| andorra                                     | Europe                      | Andorra                                                               | kHPpSs  |
| angola                                      | Africa                      | Angola                                                                | kHPpSs  |
| anhui                                       | China                       | Anhui                                                                 | kHPpSs  |
| antarctica                                  |                             | Antarctica                                                            | kHPpSs  |
| aquitaine                                   | France                      | Aquitaine                                                             | kHPpSs  |
| aragon                                      | Spain                       | Aragón                                                                | kHPpSs  |
| argentina                                   | South America               | Argentina                                                             | kHPpSs  |
| armenia                                     | Asia                        | Armenia                                                               | kHPpSs  |
| arnsberg-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Arnsberg                                             | kHPpSs  |
| ashmore-cartier                             | Australia                   | Ashmore and Cartier Islands                                           | kHPpSs  |
| asia                                        |                             | Asia                                                                  | kHPps   |
| asturias                                    | Spain                       | Asturias                                                              | kHPpSs  |
| australia                                   | Australia and Oceania       | Australia                                                             | kHPpSs  |
| australia-oceania                           |                             | Australia and Oceania                                                 | kHPps   |
| austria                                     | Europe                      | Austria                                                               | kHPpSs  |
| auvergne                                    | France                      | Auvergne                                                              | kHPpSs  |
| azerbaijan                                  | Asia                        | Azerbaijan                                                            | kHPpSs  |
| azores                                      | Europe                      | Azores                                                                | kHPpSs  |
| baden-wuerttemberg                          | Germany                     | Baden-Württemberg                                                     | kHPpSs  |
| bahamas                                     | Central America             | Bahamas                                                               | kHPpSs  |
| bangladesh                                  | Asia                        | Bangladesh                                                            | kHPpSs  |
| basse-normandie                             | France                      | Basse-Normandie                                                       | kHPpSs  |
| bayern                                      | Germany                     | Bayern                                                                | kHPpSs  |
| bedfordshire                                | England                     | Bedfordshire                                                          | kHPpSs  |
| beijing                                     | China                       | Beijing                                                               | kHPpSs  |
| belarus                                     | Europe                      | Belarus                                                               | kHPpSs  |
| belgium                                     | Europe                      | Belgium                                                               | kHPpSs  |
| belize                                      | Central America             | Belize                                                                | kHPpSs  |
| benin                                       | Africa                      | Benin                                                                 | kHPpSs  |
| berkshire                                   | England                     | Berkshire                                                             | kHPpSs  |
| berlin                                      | Germany                     | Berlin                                                                | kHPpSs  |
| bermuda                                     | United Kingdom              | Bermuda                                                               | kHPpSs  |
| bhutan                                      | Asia                        | Bhutan                                                                | kHPpSs  |
| bolivia                                     | South America               | Bolivia                                                               | kHPpSs  |
| bosnia-herzegovina                          | Europe                      | Bosnia-Herzegovina                                                    | kHPpSs  |
| botswana                                    | Africa                      | Botswana                                                              | kHPpSs  |
| bourgogne                                   | France                      | Bourgogne                                                             | kHPpSs  |
| brandenburg                                 | Germany                     | Brandenburg (mit Berlin)                                              | kHPpSs  |
| brazil                                      | South America               | Brazil                                                                | kHPps   |
| bremen                                      | Germany                     | Bremen                                                                | kHPpSs  |
| bretagne                                    | France                      | Bretagne                                                              | kHPpSs  |
| bristol                                     | England                     | Bristol                                                               | kHPpSs  |
| britain-and-ireland                         | Europe                      | Britain and Ireland                                                   | kHPps   |
| british-columbia                            | Canada                      | British Columbia                                                      | kHPpSs  |
| buckinghamshire                             | England                     | Buckinghamshire                                                       | kHPpSs  |
| bulgaria                                    | Europe                      | Bulgaria                                                              | kHPpSs  |
| burkina-faso                                | Africa                      | Burkina Faso                                                          | kHPpSs  |
| burundi                                     | Africa                      | Burundi                                                               | kHPpSs  |
| cambodia                                    | Asia                        | Cambodia                                                              | kHPpSs  |
| cambridgeshire                              | England                     | Cambridgeshire                                                        | kHPpSs  |
| cameroon                                    | Africa                      | Cameroon                                                              | kHPpSs  |
| canada                                      | North America               | Canada                                                                | kHPps   |
| canary-islands                              | Africa                      | Canary Islands                                                        | kHPpSs  |
| cantabria                                   | Spain                       | Cantabria                                                             | kHPpSs  |
| cape-verde                                  | Africa                      | Cape Verde                                                            | kHPpSs  |
| castilla-la-mancha                          | Spain                       | Castilla-La Mancha                                                    | kHPpSs  |
| castilla-y-leon                             | Spain                       | Castilla y León                                                       | kHPpSs  |
| cataluna                                    | Spain                       | Cataluña                                                              | kHPpSs  |
| central-african-republic                    | Africa                      | Central African Republic                                              | kHPpSs  |
| central-america                             |                             | Central America                                                       | kHPps   |
| central-fed-district                        | Russian Federation          | Central Federal District                                              | kHPpSs  |
| central-zone                                | India                       | Central Zone                                                          | kHPpSs  |
| centre                                      | France                      | Centre                                                                | kHPpSs  |
| centro                                      | Italy                       | Centro                                                                | kHPpSs  |
| centro-oeste                                | Brazil                      | Centro-Oeste                                                          | kHPpSs  |
| ceuta                                       | Spain                       | Ceuta                                                                 | kHPpSs  |
| chad                                        | Africa                      | Chad                                                                  | kHPpSs  |
| champagne-ardenne                           | France                      | Champagne Ardenne                                                     | kHPpSs  |
| cheshire                                    | England                     | Cheshire                                                              | kHPpSs  |
| chile                                       | South America               | Chile                                                                 | kHPpSs  |
| china                                       | Asia                        | China                                                                 | kHPpSs  |
| chongqing                                   | China                       | Chongqing                                                             | kHPpSs  |
| christmas-island                            | Australia                   | Christmas Island                                                      | kHPpSs  |
| chubu                                       | Japan                       | Chūbu region                                                          | kHPpSs  |
| chugoku                                     | Japan                       | Chūgoku region                                                        | kHPpSs  |
| cocos-islands                               | Australia                   | Cocos (Keeling) Islands                                               | kHPpSs  |
| colombia                                    | South America               | Colombia                                                              | kHPpSs  |
| comores                                     | Africa                      | Comores                                                               | kHPpSs  |
| congo-brazzaville                           | Africa                      | Congo (Republic/Brazzaville)                                          | kHPpSs  |
| congo-democratic-republic                   | Africa                      | Congo (Democratic Republic/Kinshasa)                                  | kHPpSs  |
| cook-islands                                | Australia and Oceania       | Cook Islands                                                          | kHPpSs  |
| coral-sea-islands                           | Australia                   | Coral Sea Islands                                                     | kHPpSs  |
| cornwall                                    | England                     | Cornwall                                                              | kHPpSs  |
| corse                                       | France                      | Corse                                                                 | kHPpSs  |
| costa-rica                                  | Central America             | Costa Rica                                                            | kHPpSs  |
| crimean-fed-district                        | Russian Federation          | Crimean Federal District                                              | kHPpSs  |
| croatia                                     | Europe                      | Croatia                                                               | kHPpSs  |
| cuba                                        | Central America             | Cuba                                                                  | kHPpSs  |
| cumbria                                     | England                     | Cumbria                                                               | kHPpSs  |
| cyprus                                      | Europe                      | Cyprus                                                                | kHPpSs  |
| czech-republic                              | Europe                      | Czech Republic                                                        | kHPpSs  |
| dach                                        | Europe                      | Germany, Austria, Switzerland                                         | kHPps   |
| denmark                                     | Europe                      | Denmark                                                               | kHPpSs  |
| derbyshire                                  | England                     | Derbyshire                                                            | kHPpSs  |
| detmold-regbez                              | Nordrhein-Westfalen         | Regierungsbezirk Detmold                                              | kHPpSs  |
| devon                                       | England                     | Devon                                                                 | kHPpSs  |
| djibouti                                    | Africa                      | Djibouti                                                              | kHPpSs  |
| dolnoslaskie                                | Poland                      | Województwo dolnośląskie<br />(Lower Silesian Voivodeship)            | kHPpSs  |
| dorset                                      | England                     | Dorset                                                                | kHPpSs  |
| drenthe                                     | Netherlands                 | Drenthe                                                               | kHPpSs  |
| duesseldorf-regbez                          | Nordrhein-Westfalen         | Regierungsbezirk Düsseldorf                                           | kHPpSs  |
| durham                                      | England                     | Durham                                                                | kHPpSs  |
| east-sussex                                 | England                     | East Sussex                                                           | kHPpSs  |
| east-timor                                  | Asia                        | East Timor                                                            | kHPpSs  |
| east-yorkshire-with-hull                    | England                     | East Yorkshire with Hull                                              | kHPpSs  |
| eastern-zone                                | India                       | Eastern Zone                                                          | kHPpSs  |
| ecuador                                     | South America               | Ecuador                                                               | kHPpSs  |
| egypt                                       | Africa                      | Egypt                                                                 | kHPpSs  |
| el-salvador                                 | Central America             | El Salvador                                                           | kHPpSs  |
| enfield                                     | Greater London              | Enfield                                                               | kHPpSs  |
| england                                     | United Kingdom              | England                                                               | kHPpSs  |
| equatorial-guinea                           | Africa                      | Equatorial Guinea                                                     | kHPpSs  |
| eritrea                                     | Africa                      | Eritrea                                                               | kHPpSs  |
| essex                                       | England                     | Essex                                                                 | kHPpSs  |
| estonia                                     | Europe                      | Estonia                                                               | kHPpSs  |
| ethiopia                                    | Africa                      | Ethiopia                                                              | kHPpSs  |
| europe                                      |                             | Europe                                                                | kHPps   |
| extremadura                                 | Spain                       | Extremadura                                                           | kHPpSs  |
| falklands                                   | United Kingdom              | Falkland Islands                                                      | kHPpSs  |
| far-eastern-fed-district                    | Russian Federation          | Far Eastern Federal District                                          | kHPpSs  |
| faroe-islands                               | Europe                      | Faroe Islands                                                         | kHPpSs  |
| fiji                                        | Australia and Oceania       | Fiji                                                                  | kHPpSs  |
| finland                                     | Europe                      | Finland                                                               | kHPpSs  |
| flevoland                                   | Netherlands                 | Flevoland                                                             | kHPpSs  |
| france                                      | Europe                      | France                                                                | kHPps   |
| franche-comte                               | France                      | Franche Comte                                                         | kHPpSs  |
| freiburg-regbez                             | Baden-Württemberg           | Regierungsbezirk Freiburg                                             | kHPpSs  |
| friesland                                   | Netherlands                 | Friesland                                                             | kHPpSs  |
| fujian                                      | China                       | Fujian                                                                | kHPpSs  |
| gabon                                       | Africa                      | Gabon                                                                 | kHPpSs  |
| galicia                                     | Spain                       | Galicia                                                               | kHPpSs  |
| gansu                                       | China                       | Gansu                                                                 | kHPpSs  |
| gcc-states                                  | Asia                        | GCC States                                                            | kHPpSs  |
| gelderland                                  | Netherlands                 | Gelderland                                                            | kHPpSs  |
| georgia                                     | Europe                      | Georgia                                                               | kHPpSs  |
| germany                                     | Europe                      | Germany                                                               | kHPps   |
| ghana                                       | Africa                      | Ghana                                                                 | kHPpSs  |
| gloucestershire                             | England                     | Gloucestershire                                                       | kHPpSs  |
| great-britain                               | Europe                      | Great Britain                                                         | kHPps   |
| greater-london                              | England                     | Greater London                                                        | kHPpSs  |
| greater-manchester                          | England                     | Greater Manchester                                                    | kHPpSs  |
| greece                                      | Europe                      | Greece                                                                | kHPpSs  |
| greenland                                   | North America               | Greenland                                                             | kHPpSs  |
| groningen                                   | Netherlands                 | Groningen                                                             | kHPpSs  |
| guadeloupe                                  | France                      | Guadeloupe                                                            | kHPpSs  |
| guangdong                                   | China                       | Guangdong (with Hong Kong and Macau)                                  | kHPpSs  |
| guangxi                                     | China                       | Guangxi                                                               | kHPpSs  |
| guatemala                                   | Central America             | Guatemala                                                             | kHPpSs  |
| guernsey-jersey                             | Europe                      | Guernsey and Jersey                                                   | kHPpSs  |
| guinea                                      | Africa                      | Guinea                                                                | kHPpSs  |
| guinea-bissau                               | Africa                      | Guinea-Bissau                                                         | kHPpSs  |
| guizhou                                     | China                       | Guizhou                                                               | kHPpSs  |
| guyana                                      | South America               | Guyana                                                                | kHPpSs  |
| guyane                                      | France                      | Guyane                                                                | kHPpSs  |
| hainan                                      | China                       | Hainan                                                                | kHPpSs  |
| haiti-and-domrep                            | Central America             | Haiti and Dominican Republic                                          | kHPpSs  |
| hamburg                                     | Germany                     | Hamburg                                                               | kHPpSs  |
| hampshire                                   | England                     | Hampshire                                                             | kHPpSs  |
| haute-normandie                             | France                      | Haute-Normandie                                                       | kHPpSs  |
| heard-mcdonald                              | Australia                   | Heard Island and McDonald Islands                                     | kHPpSs  |
| hebei                                       | China                       | Hebei (with Beijing and Tianjin)                                      | kHPpSs  |
| heilongjiang                                | China                       | Heilongjiang                                                          | kHPpSs  |
| henan                                       | China                       | Henan                                                                 | kHPpSs  |
| herefordshire                               | England                     | Herefordshire                                                         | kHPpSs  |
| hertfordshire                               | England                     | Hertfordshire                                                         | kHPpSs  |
| hessen                                      | Germany                     | Hessen                                                                | kHPpSs  |
| hokkaido                                    | Japan                       | Hokkaidō                                                              | kHPpSs  |
| honduras                                    | Central America             | Honduras                                                              | kHPpSs  |
| hong-kong                                   | China                       | Hong Kong                                                             | kHPpSs  |
| hubei                                       | China                       | Hubei                                                                 | kHPpSs  |
| hunan                                       | China                       | Hunan                                                                 | kHPpSs  |
| hungary                                     | Europe                      | Hungary                                                               | kHPpSs  |
| iceland                                     | Europe                      | Iceland                                                               | kHPpSs  |
| ile-de-clipperton                           | Australia and Oceania       | Île de Clipperton                                                     | kHPpSs  |
| ile-de-france                               | France                      | Ile-de-France                                                         | kHPpSs  |
| india                                       | Asia                        | India                                                                 | kHPpSs  |
| indonesia                                   | Asia                        | Indonesia (with East Timor)                                           | kHPpSs  |
| inner-mongolia                              | China                       | Inner Mongolia                                                        | kHPpSs  |
| interior-admreg                             | British Columbia            | Interior Administrative Region                                        | kHPpSs  |
| iran                                        | Asia                        | Iran                                                                  | kHPpSs  |
| iraq                                        | Asia                        | Iraq                                                                  | kHPpSs  |
| ireland-and-northern-ireland                | Europe                      | Ireland and Northern Ireland                                          | kHPpSs  |
| island-admreg                               | British Columbia            | Island Administrative Region                                          | kHPpSs  |
| islas-baleares                              | Spain                       | Islas Baleares                                                        | kHPpSs  |
| isle-of-man                                 | Europe                      | Isle of Man                                                           | kHPpSs  |
| isle-of-wight                               | England                     | Isle of Wight                                                         | kHPpSs  |
| isole                                       | Italy                       | Isole                                                                 | kHPpSs  |
| israel-and-palestine                        | Asia                        | Israel and Palestine                                                  | kHPpSs  |
| italy                                       | Europe                      | Italy                                                                 | kHPps   |
| ivory-coast                                 | Africa                      | Ivory Coast                                                           | kHPpSs  |
| jamaica                                     | Central America             | Jamaica                                                               | kHPpSs  |
| japan                                       | Asia                        | Japan                                                                 | kHPps   |
| java                                        | Indonesia (with East Timor) | Java                                                                  | kHPpSs  |
| jiangsu                                     | China                       | Jiangsu                                                               | kHPpSs  |
| jiangxi                                     | China                       | Jiangxi                                                               | kHPpSs  |
| jihocesky                                   | Czech Republic              | Jihočeský kraj                                                        | kHPpSs  |
| jihomoravsky                                | Czech Republic              | Jihomoravský kraj                                                     | kHPpSs  |
| jilin                                       | China                       | Jilin                                                                 | kHPpSs  |
| jordan                                      | Asia                        | Jordan                                                                | kHPpSs  |
| kalimantan                                  | Indonesia (with East Timor) | Kalimantan                                                            | kHPpSs  |
| kaliningrad                                 | Russian Federation          | Kaliningrad                                                           | kHPpSs  |
| kansai                                      | Japan                       | Kansai region (a.k.a. Kinki region)                                   | kHPpSs  |
| kanto                                       | Japan                       | Kantō region                                                          | kHPpSs  |
| karlovarsky                                 | Czech Republic              | Karlovarský kraj                                                      | kHPpSs  |
| karlsruhe-regbez                            | Baden-Württemberg           | Regierungsbezirk Karlsruhe                                            | kHPpSs  |
| kazakhstan                                  | Asia                        | Kazakhstan                                                            | kHPpSs  |
| kent                                        | England                     | Kent                                                                  | kHPpSs  |
| kenya                                       | Africa                      | Kenya                                                                 | kHPpSs  |
| kiribati                                    | Australia and Oceania       | Kiribati                                                              | kHPpSs  |
| kitikmeot                                   | Nunavut                     | Kitikmeot Region                                                      | kHPpSs  |
| kivalliq                                    | Nunavut                     | Kivalliq Region                                                       | kHPpSs  |
| koeln-regbez                                | Nordrhein-Westfalen         | Regierungsbezirk Köln                                                 | kHPpSs  |
| kootenay-admreg                             | British Columbia            | Kootenay Administrative Region                                        | kHPpSs  |
| kosovo                                      | Europe                      | Kosovo                                                                | kHPpSs  |
| kralovehradecky                             | Czech Republic              | Královéhradecký kraj                                                  | kHPpSs  |
| kujawsko-pomorskie                          | Poland                      | Województwo kujawsko-pomorskie<br />(Kuyavian-Pomeranian Voivodeship) | kHPpSs  |
| kyrgyzstan                                  | Asia                        | Kyrgyzstan                                                            | kHPpSs  |
| kyushu                                      | Japan                       | Kyūshū                                                                | kHPpSs  |
| la-rioja                                    | Spain                       | La Rioja                                                              | kHPpSs  |
| lancashire                                  | England                     | Lancashire                                                            | kHPpSs  |
| languedoc-roussillon                        | France                      | Languedoc-Roussillon                                                  | kHPpSs  |
| laos                                        | Asia                        | Laos                                                                  | kHPpSs  |
| latvia                                      | Europe                      | Latvia                                                                | kHPpSs  |
| lebanon                                     | Asia                        | Lebanon                                                               | kHPpSs  |
| leicestershire                              | England                     | Leicestershire                                                        | kHPpSs  |
| lesotho                                     | Africa                      | Lesotho                                                               | kHPpSs  |
| liaoning                                    | China                       | Liaoning                                                              | kHPpSs  |
| liberecky                                   | Czech Republic              | Liberecký kraj                                                        | kHPpSs  |
| liberia                                     | Africa                      | Liberia                                                               | kHPpSs  |
| libya                                       | Africa                      | Libya                                                                 | kHPpSs  |
| liechtenstein                               | Europe                      | Liechtenstein                                                         | kHPpSs  |
| limburg                                     | Netherlands                 | Limburg                                                               | kHPpSs  |
| limousin                                    | France                      | Limousin                                                              | kHPpSs  |
| lincolnshire                                | England                     | Lincolnshire                                                          | kHPpSs  |
| lithuania                                   | Europe                      | Lithuania                                                             | kHPpSs  |
| lodzkie                                     | Poland                      | Województwo łódzkie<br />(Łódź Voivodeship)                           | kHPpSs  |
| lorraine                                    | France                      | Lorraine                                                              | kHPpSs  |
| lubelskie                                   | Poland                      | Województwo lubelskie<br />(Lublin Voivodeship)                       | kHPpSs  |
| lubuskie                                    | Poland                      | Województwo lubuskie<br />(Lubusz Voivodeship)                        | kHPpSs  |
| luxembourg                                  | Europe                      | Luxembourg                                                            | kHPpSs  |
| macau                                       | China                       | Macau                                                                 | kHPpSs  |
| macedonia                                   | Europe                      | Macedonia                                                             | kHPpSs  |
| madagascar                                  | Africa                      | Madagascar                                                            | kHPpSs  |
| madrid                                      | Spain                       | Madrid                                                                | kHPpSs  |
| malawi                                      | Africa                      | Malawi                                                                | kHPpSs  |
| malaysia-singapore-brunei                   | Asia                        | Malaysia, Singapore, and Brunei                                       | kHPpSs  |
| maldives                                    | Asia                        | Maldives                                                              | kHPpSs  |
| mali                                        | Africa                      | Mali                                                                  | kHPpSs  |
| malopolskie                                 | Poland                      | Województwo małopolskie<br />(Lesser Poland Voivodeship)              | kHPpSs  |
| malta                                       | Europe                      | Malta                                                                 | kHPpSs  |
| maluku                                      | Indonesia (with East Timor) | Maluku                                                                | kHPpSs  |
| manitoba                                    | Canada                      | Manitoba                                                              | kHPpSs  |
| marshall-islands                            | Australia and Oceania       | Marshall Islands                                                      | kHPpSs  |
| martinique                                  | France                      | Martinique                                                            | kHPpSs  |
| mauritania                                  | Africa                      | Mauritania                                                            | kHPpSs  |
| mauritius                                   | Africa                      | Mauritius                                                             | kHPpSs  |
| mayotte                                     | France                      | Mayotte                                                               | kHPpSs  |
| mazowieckie                                 | Poland                      | Województwo mazowieckie<br />(Mazovian Voivodeship)                   | kHPpSs  |
| mecklenburg-vorpommern                      | Germany                     | Mecklenburg-Vorpommern                                                | kHPpSs  |
| melilla                                     | Spain                       | Melilla                                                               | kHPpSs  |
| merseyside                                  | England                     | Merseyside                                                            | kHPpSs  |
| mexico                                      | North America               | Mexico                                                                | kHPpSs  |
| micronesia                                  | Australia and Oceania       | Micronesia                                                            | kHPpSs  |
| midi-pyrenees                               | France                      | Midi-Pyrenees                                                         | kHPpSs  |
| mittelfranken                               | Bayern                      | Mittelfranken                                                         | kHPpSs  |
| moldova                                     | Europe                      | Moldova                                                               | kHPpSs  |
| monaco                                      | Europe                      | Monaco                                                                | kHPpSs  |
| mongolia                                    | Asia                        | Mongolia                                                              | kHPpSs  |
| montenegro                                  | Europe                      | Montenegro                                                            | kHPpSs  |
| moravskoslezky                              | Czech Republic              | Moravskoslezský kraj                                                  | kHPpSs  |
| morocco                                     | Africa                      | Morocco                                                               | kHPpSs  |
| mozambique                                  | Africa                      | Mozambique                                                            | kHPpSs  |
| muenster-regbez                             | Nordrhein-Westfalen         | Regierungsbezirk Münster                                              | kHPpSs  |
| murcia                                      | Spain                       | Murcia                                                                | kHPpSs  |
| myanmar                                     | Asia                        | Myanmar (a.k.a. Burma)                                                | kHPpSs  |
| namibia                                     | Africa                      | Namibia                                                               | kHPpSs  |
| nauru                                       | Australia and Oceania       | Nauru                                                                 | kHPpSs  |
| navarra                                     | Spain                       | Navarra                                                               | kHPpSs  |
| nepal                                       | Asia                        | Nepal                                                                 | kHPpSs  |
| netherlands                                 | Europe                      | Netherlands                                                           | kHPpSs  |
| new-brunswick                               | Canada                      | New Brunswick                                                         | kHPpSs  |
| new-caledonia                               | Australia and Oceania       | New Caledonia                                                         | kHPpSs  |
| new-south-wales                             | Australia                   | New South Wales (with ACT and JBT)                                    | kHPpSs  |
| new-zealand                                 | Australia and Oceania       | New Zealand                                                           | kHPpSs  |
| newfoundland-and-labrador                   | Canada                      | Newfoundland and Labrador                                             | kHPpSs  |
| nicaragua                                   | Central America             | Nicaragua                                                             | kHPpSs  |
| niederbayern                                | Bayern                      | Niederbayern                                                          | kHPpSs  |
| niedersachsen                               | Germany                     | Niedersachsen (mit Bremen)                                            | kHPpSs  |
| niger                                       | Africa                      | Niger                                                                 | kHPpSs  |
| nigeria                                     | Africa                      | Nigeria                                                               | kHPpSs  |
| ningxia                                     | China                       | Ningxia                                                               | kHPpSs  |
| niue                                        | Australia and Oceania       | Niue                                                                  | kHPpSs  |
| noord-brabant                               | Netherlands                 | Noord-Brabant                                                         | kHPpSs  |
| noord-holland                               | Netherlands                 | Noord-Holland                                                         | kHPpSs  |
| norcal                                      | us/california               | Northern California                                                   | kHPpSs  |
| nord-est                                    | Italy                       | Nord-Est                                                              | kHPpSs  |
| nord-norge                                  | Norway                      | Nord-Norge<br />(Northern Norway)                                     | kHPpSs  |
| nord-ovest                                  | Italy                       | Nord-Ovest                                                            | kHPpSs  |
| nord-pas-de-calais                          | France                      | Nord-Pas-de-Calais                                                    | kHPpSs  |
| nordeste                                    | Brazil                      | Nordeste                                                              | kHPpSs  |
| nordrhein-westfalen                         | Germany                     | Nordrhein-Westfalen                                                   | kHPpSs  |
| norfolk                                     | England                     | Norfolk                                                               | kHPpSs  |
| norfolk-island                              | Australia                   | Norfolk Island                                                        | kHPpSs  |
| norte                                       | Brazil                      | Norte                                                                 | kHPpSs  |
| north-admreg                                | British Columbia            | North Administrative Region                                           | kHPpSs  |
| north-america                               |                             | North America                                                         | kHPps   |
| north-caucasus-fed-district                 | Russian Federation          | North Caucasus Federal District                                       | kHPpSs  |
| north-eastern-zone                          | India                       | North-Eastern Zone                                                    | kHPpSs  |
| north-korea                                 | Asia                        | North Korea                                                           | kHPpSs  |
| north-yorkshire                             | England                     | North Yorkshire                                                       | kHPpSs  |
| northamptonshire                            | England                     | Northamptonshire                                                      | kHPpSs  |
| northern-territory                          | Australia                   | Northern Territory                                                    | kHPpSs  |
| northern-zone                               | India                       | Northern Zone                                                         | kHPpSs  |
| northumberland                              | England                     | Northumberland                                                        | kHPpSs  |
| northwest-territories                       | Canada                      | Northwest Territories                                                 | kHPpSs  |
| northwestern-fed-district                   | Russian Federation          | Northwestern Federal District                                         | kHPpSs  |
| norway                                      | Europe                      | Norway                                                                | kHPpSs  |
| nottinghamshire                             | England                     | Nottinghamshire                                                       | kHPpSs  |
| nova-scotia                                 | Canada                      | Nova Scotia                                                           | kHPpSs  |
| nunavut                                     | Canada                      | Nunavut                                                               | kHPpSs  |
| nusa-tenggara                               | Indonesia (with East Timor) | Nusa-Tenggara                                                         | kHPpSs  |
| oberbayern                                  | Bayern                      | Oberbayern                                                            | kHPpSs  |
| oberfranken                                 | Bayern                      | Oberfranken                                                           | kHPpSs  |
| oberpfalz                                   | Bayern                      | Oberpfalz                                                             | kHPpSs  |
| okanagan-admreg                             | British Columbia            | Okanagan Administrative Region                                        | kHPpSs  |
| olomoucky                                   | Czech Republic              | Olomoucký kraj                                                        | kHPpSs  |
| ontario                                     | Canada                      | Ontario                                                               | kHPpSs  |
| opolskie                                    | Poland                      | Województwo opolskie<br />(Opole Voivodeship)                         | kHPpSs  |
| ostlandet                                   | Norway                      | Østlandet<br />(Eastern Norway)                                       | kHPpSs  |
| overijssel                                  | Netherlands                 | Overijssel                                                            | kHPpSs  |
| oxfordshire                                 | England                     | Oxfordshire                                                           | kHPpSs  |
| pais-vasco                                  | Spain                       | País Vasco                                                            | kHPpSs  |
| pakistan                                    | Asia                        | Pakistan                                                              | kHPpSs  |
| palau                                       | Australia and Oceania       | Palau                                                                 | kHPpSs  |
| panama                                      | Central America             | Panama                                                                | kHPpSs  |
| papua                                       | Indonesia (with East Timor) | Papua                                                                 | kHPpSs  |
| papua-new-guinea                            | Australia and Oceania       | Papua New Guinea                                                      | kHPpSs  |
| paraguay                                    | South America               | Paraguay                                                              | kHPpSs  |
| pardubicky                                  | Czech Republic              | Pardubický kraj                                                       | kHPpSs  |
| pays-de-la-loire                            | France                      | Pays de la Loire                                                      | kHPpSs  |
| peru                                        | South America               | Peru                                                                  | kHPpSs  |
| philippines                                 | Asia                        | Philippines                                                           | kHPpSs  |
| picardie                                    | France                      | Picardie                                                              | kHPpSs  |
| pitcairn-islands                            | Australia and Oceania       | Pitcairn Islands                                                      | kHPpSs  |
| plzensky                                    | Czech Republic              | Plzeňský kraj                                                         | kHPpSs  |
| podkarpackie                                | Poland                      | Województwo podkarpackie<br />(Subcarpathian Voivodeship)             | kHPpSs  |
| podlaskie                                   | Poland                      | Województwo podlaskie<br />(Podlaskie Voivodeship)                    | kHPpSs  |
| poitou-charentes                            | France                      | Poitou-Charentes                                                      | kHPpSs  |
| poland                                      | Europe                      | Poland                                                                | kHPps   |
| polynesie-francaise                         | Australia and Oceania       | Polynésie française (French Polynesia)                                | kHPpSs  |
| pomorskie                                   | Poland                      | Województwo pomorskie<br />(Pomeranian Voivodeship)                   | kHPpSs  |
| portugal                                    | Europe                      | Portugal                                                              | kHPpSs  |
| praha                                       | Czech Republic              | Praha                                                                 | kHPpSs  |
| prince-edward-island                        | Canada                      | Prince Edward Island                                                  | kHPpSs  |
| provence-alpes-cote-d-azur                  | France                      | Provence Alpes-Cote-d'Azur                                            | kHPpSs  |
| qikiqtaaluk                                 | Nunavut                     | Qikiqtaaluk Region                                                    | kHPpSs  |
| qinghai                                     | China                       | Qinghai                                                               | kHPpSs  |
| quebec                                      | Canada                      | Quebec                                                                | kHPpSs  |
| queensland                                  | Australia                   | Queensland                                                            | kHPpSs  |
| reunion                                     | France                      | Reunion                                                               | kHPpSs  |
| rheinland-pfalz                             | Germany                     | Rheinland-Pfalz                                                       | kHPpSs  |
| rhone-alpes                                 | France                      | Rhone-Alpes                                                           | kHPpSs  |
| romania                                     | Europe                      | Romania                                                               | kHPpSs  |
| russia                                      |                             | Russian Federation                                                    | kHPps   |
| rutland                                     | England                     | Rutland                                                               | kHPpSs  |
| rwanda                                      | Africa                      | Rwanda                                                                | kHPpSs  |
| saarland                                    | Germany                     | Saarland                                                              | kHPpSs  |
| sachsen                                     | Germany                     | Sachsen                                                               | kHPpSs  |
| sachsen-anhalt                              | Germany                     | Sachsen-Anhalt                                                        | kHPpSs  |
| saint-helena-ascension-and-tristan-da-cunha | Africa                      | Saint Helena, Ascension, and Tristan da Cunha                         | kHPpSs  |
| samoa                                       | Australia and Oceania       | Samoa                                                                 | kHPpSs  |
| sao-tome-and-principe                       | Africa                      | Sao Tome and Principe                                                 | kHPpSs  |
| saskatchewan                                | Canada                      | Saskatchewan                                                          | kHPpSs  |
| schleswig-holstein                          | Germany                     | Schleswig-Holstein                                                    | kHPpSs  |
| schwaben                                    | Bayern                      | Schwaben                                                              | kHPpSs  |
| scotland                                    | United Kingdom              | Scotland                                                              | kHPpSs  |
| sea                                         | Asia                        | South-East Asia                                                       | kHPps   |
| senegal-and-gambia                          | Africa                      | Senegal and Gambia                                                    | kHPpSs  |
| serbia                                      | Europe                      | Serbia                                                                | kHPpSs  |
| seychelles                                  | Africa                      | Seychelles                                                            | kHPpSs  |
| shaanxi                                     | China                       | Shaanxi                                                               | kHPpSs  |
| shandong                                    | China                       | Shandong                                                              | kHPpSs  |
| shanghai                                    | China                       | Shanghai                                                              | kHPpSs  |
| shanxi                                      | China                       | Shanxi                                                                | kHPpSs  |
| shikoku                                     | Japan                       | Shikoku                                                               | kHPpSs  |
| shropshire                                  | England                     | Shropshire                                                            | kHPpSs  |
| siberian-fed-district                       | Russian Federation          | Siberian Federal District                                             | kHPpSs  |
| sichuan                                     | China                       | Sichuan                                                               | kHPpSs  |
| sierra-leone                                | Africa                      | Sierra Leone                                                          | kHPpSs  |
| slaskie                                     | Poland                      | Województwo śląskie<br />(Silesian Voivodeship)                       | kHPpSs  |
| slovakia                                    | Europe                      | Slovakia                                                              | kHPpSs  |
| slovenia                                    | Europe                      | Slovenia                                                              | kHPpSs  |
| socal                                       | us/california               | Southern California                                                   | kHPpSs  |
| solomon-islands                             | Australia and Oceania       | Solomon Islands                                                       | kHPpSs  |
| somalia                                     | Africa                      | Somalia                                                               | kHPpSs  |
| somerset                                    | England                     | Somerset                                                              | kHPpSs  |
| sorlandet                                   | Norway                      | Sørlandet<br />(Southern Norway)                                      | kHPpSs  |
| south-africa                                | Africa                      | South Africa                                                          | kHPpSs  |
| south-africa-and-lesotho                    | Africa                      | South Africa (includes Lesotho)                                       | kHPps   |
| south-america                               |                             | South America                                                         | kHPps   |
| south-australia                             | Australia                   | South Australia                                                       | kHPpSs  |
| south-fed-district                          | Russian Federation          | South Federal District                                                | kHPpSs  |
| south-korea                                 | Asia                        | South Korea                                                           | kHPpSs  |
| south-sudan                                 | Africa                      | South Sudan                                                           | kHPpSs  |
| south-yorkshire                             | England                     | South Yorkshire                                                       | kHPpSs  |
| southcoast-admreg                           | British Columbia            | South Coast Administrative Region                                     | kHPpSs  |
| southern-zone                               | India                       | Southern Zone                                                         | kHPpSs  |
| spain                                       | Europe                      | Spain                                                                 | kHPpSs  |
| sri-lanka                                   | Asia                        | Sri Lanka                                                             | kHPpSs  |
| staffordshire                               | England                     | Staffordshire                                                         | kHPpSs  |
| stredocesky                                 | Czech Republic              | Středočeský kraj (with Praha)                                         | kHPpSs  |
| stuttgart-regbez                            | Baden-Württemberg           | Regierungsbezirk Stuttgart                                            | kHPpSs  |
| sud                                         | Italy                       | Sud                                                                   | kHPpSs  |
| sudan                                       | Africa                      | Sudan                                                                 | kHPpSs  |
| sudeste                                     | Brazil                      | Sudeste                                                               | kHPpSs  |
| suffolk                                     | England                     | Suffolk                                                               | kHPpSs  |
| sul                                         | Brazil                      | Sul                                                                   | kHPpSs  |
| sulawesi                                    | Indonesia (with East Timor) | Sulawesi                                                              | kHPpSs  |
| sumatra                                     | Indonesia (with East Timor) | Sumatra                                                               | kHPpSs  |
| suriname                                    | South America               | Suriname                                                              | kHPpSs  |
| surrey                                      | England                     | Surrey                                                                | kHPpSs  |
| svalbard-janmayen                           | Norway                      | Svalbard and Jan Mayen                                                | kHPpSs  |
| swaziland                                   | Africa                      | Swaziland                                                             | kHPpSs  |
| sweden                                      | Europe                      | Sweden                                                                | kHPpSs  |
| swietokrzyskie                              | Poland                      | Województwo świętokrzyskie<br />(Świętokrzyskie Voivodeship)          | kHPpSs  |
| switzerland                                 | Europe                      | Switzerland                                                           | kHPpSs  |
| syria                                       | Asia                        | Syria                                                                 | kHPpSs  |
| taiwan                                      | Asia                        | Taiwan                                                                | kHPpSs  |
| tajikistan                                  | Asia                        | Tajikistan                                                            | kHPpSs  |
| tanzania                                    | Africa                      | Tanzania                                                              | kHPpSs  |
| tasmania                                    | Australia                   | Tasmania                                                              | kHPpSs  |
| thailand                                    | Asia                        | Thailand                                                              | kHPpSs  |
| thueringen                                  | Germany                     | Thüringen                                                             | kHPpSs  |
| tianjin                                     | China                       | Tianjin                                                               | kHPpSs  |
| tibet                                       | China                       | Tibet                                                                 | kHPpSs  |
| togo                                        | Africa                      | Togo                                                                  | kHPpSs  |
| tohoku                                      | Japan                       | Tōhoku region                                                         | kHPpSs  |
| tokelau                                     | Australia and Oceania       | Tokelau                                                               | kHPpSs  |
| tonga                                       | Australia and Oceania       | Tonga                                                                 | kHPpSs  |
| trondelag                                   | Norway                      | Trøndelag<br />(Central Norway)                                       | kHPpSs  |
| tuebingen-regbez                            | Baden-Württemberg           | Regierungsbezirk Tübingen                                             | kHPpSs  |
| tunisia                                     | Africa                      | Tunisia                                                               | kHPpSs  |
| turkey                                      | Europe                      | Turkey                                                                | kHPpSs  |
| turkmenistan                                | Asia                        | Turkmenistan                                                          | kHPpSs  |
| tuvalu                                      | Australia and Oceania       | Tuvalu                                                                | kHPpSs  |
| tyne-and-wear                               | England                     | Tyne and Wear                                                         | kHPpSs  |
| uganda                                      | Africa                      | Uganda                                                                | kHPpSs  |
| ukraine                                     | Europe                      | Ukraine (with Crimea)                                                 | kHPpSs  |
| united-kingdom                              | Europe                      | United Kingdom                                                        | kHPps   |
| unterfranken                                | Bayern                      | Unterfranken                                                          | kHPpSs  |
| ural-fed-district                           | Russian Federation          | Ural Federal District                                                 | kHPpSs  |
| uruguay                                     | South America               | Uruguay                                                               | kHPpSs  |
| us                                          | North America               | United States of America                                              | kHPps   |
| us-midwest                                  | North America               | US Midwest                                                            | kHPps   |
| us-northeast                                | North America               | US Northeast                                                          | kHPps   |
| us-pacific                                  | North America               | US Pacific                                                            | kHPps   |
| us-south                                    | North America               | US South                                                              | kHPps   |
| us-west                                     | North America               | US West                                                               | kHPps   |
| us/alabama                                  | North America               | us/alabama                                                            | kHPpSs  |
| us/alaska                                   | North America               | us/alaska                                                             | kHPpSs  |
| us/arizona                                  | North America               | us/arizona                                                            | kHPpSs  |
| us/arkansas                                 | North America               | us/arkansas                                                           | kHPpSs  |
| us/california                               | North America               | us/california                                                         | kHPps   |
| us/colorado                                 | North America               | us/colorado                                                           | kHPpSs  |
| us/connecticut                              | North America               | us/connecticut                                                        | kHPpSs  |
| us/delaware                                 | North America               | us/delaware                                                           | kHPpSs  |
| us/district-of-columbia                     | North America               | us/district-of-columbia                                               | kHPpSs  |
| us/florida                                  | North America               | us/florida                                                            | kHPpSs  |
| us/georgia                                  | North America               | Georgia                                                               | kHPpSs  |
| us/hawaii                                   | North America               | us/hawaii                                                             | kHPpSs  |
| us/idaho                                    | North America               | us/idaho                                                              | kHPpSs  |
| us/illinois                                 | North America               | us/illinois                                                           | kHPpSs  |
| us/indiana                                  | North America               | us/indiana                                                            | kHPpSs  |
| us/iowa                                     | North America               | us/iowa                                                               | kHPpSs  |
| us/kansas                                   | North America               | us/kansas                                                             | kHPpSs  |
| us/kentucky                                 | North America               | us/kentucky                                                           | kHPpSs  |
| us/louisiana                                | North America               | us/louisiana                                                          | kHPpSs  |
| us/maine                                    | North America               | us/maine                                                              | kHPpSs  |
| us/maryland                                 | North America               | us/maryland                                                           | kHPpSs  |
| us/massachusetts                            | North America               | us/massachusetts                                                      | kHPpSs  |
| us/michigan                                 | North America               | us/michigan                                                           | kHPpSs  |
| us/minnesota                                | North America               | us/minnesota                                                          | kHPpSs  |
| us/mississippi                              | North America               | us/mississippi                                                        | kHPpSs  |
| us/missouri                                 | North America               | us/missouri                                                           | kHPpSs  |
| us/montana                                  | North America               | us/montana                                                            | kHPpSs  |
| us/nebraska                                 | North America               | us/nebraska                                                           | kHPpSs  |
| us/nevada                                   | North America               | us/nevada                                                             | kHPpSs  |
| us/new-hampshire                            | North America               | us/new-hampshire                                                      | kHPpSs  |
| us/new-jersey                               | North America               | us/new-jersey                                                         | kHPpSs  |
| us/new-mexico                               | North America               | us/new-mexico                                                         | kHPpSs  |
| us/new-york                                 | North America               | us/new-york                                                           | kHPpSs  |
| us/north-carolina                           | North America               | us/north-carolina                                                     | kHPpSs  |
| us/north-dakota                             | North America               | us/north-dakota                                                       | kHPpSs  |
| us/ohio                                     | North America               | us/ohio                                                               | kHPpSs  |
| us/oklahoma                                 | North America               | us/oklahoma                                                           | kHPpSs  |
| us/oregon                                   | North America               | us/oregon                                                             | kHPpSs  |
| us/pennsylvania                             | North America               | us/pennsylvania                                                       | kHPpSs  |
| us/puerto-rico                              | North America               | us/puerto-rico                                                        | kHPpSs  |
| us/rhode-island                             | North America               | us/rhode-island                                                       | kHPpSs  |
| us/south-carolina                           | North America               | us/south-carolina                                                     | kHPpSs  |
| us/south-dakota                             | North America               | us/south-dakota                                                       | kHPpSs  |
| us/tennessee                                | North America               | us/tennessee                                                          | kHPpSs  |
| us/texas                                    | North America               | us/texas                                                              | kHPpSs  |
| us/us-virgin-islands                        | North America               | us/us-virgin-islands                                                  | kHPpSs  |
| us/utah                                     | North America               | us/utah                                                               | kHPpSs  |
| us/vermont                                  | North America               | us/vermont                                                            | kHPpSs  |
| us/virginia                                 | North America               | us/virginia                                                           | kHPpSs  |
| us/washington                               | North America               | us/washington                                                         | kHPpSs  |
| us/west-virginia                            | North America               | us/west-virginia                                                      | kHPpSs  |
| us/wisconsin                                | North America               | us/wisconsin                                                          | kHPpSs  |
| us/wyoming                                  | North America               | us/wyoming                                                            | kHPpSs  |
| ustecky                                     | Czech Republic              | Ústecký kraj                                                          | kHPpSs  |
| utrecht                                     | Netherlands                 | Utrecht                                                               | kHPpSs  |
| uzbekistan                                  | Asia                        | Uzbekistan                                                            | kHPpSs  |
| valencia                                    | Spain                       | Valencia                                                              | kHPpSs  |
| vanuatu                                     | Australia and Oceania       | Vanuatu                                                               | kHPpSs  |
| venezuela                                   | South America               | Venezuela                                                             | kHPpSs  |
| vestlandet                                  | Norway                      | Vestlandet<br />(Western Norway)                                      | kHPpSs  |
| victoria                                    | Australia                   | Victoria                                                              | kHPpSs  |
| vietnam                                     | Asia                        | Vietnam                                                               | kHPpSs  |
| volga-fed-district                          | Russian Federation          | Volga Federal District                                                | kHPpSs  |
| vysocina                                    | Czech Republic              | Kraj Vysočina                                                         | kHPpSs  |
| wales                                       | United Kingdom              | Wales                                                                 | kHPpSs  |
| wallis-et-futuna                            | Australia and Oceania       | Wallis et Futuna                                                      | kHPpSs  |
| warminsko-mazurskie                         | Poland                      | Województwo warmińsko-mazurskie<br />(Warmian-Masurian Voivodeship)   | kHPpSs  |
| warwickshire                                | England                     | Warwickshire                                                          | kHPpSs  |
| west-midlands                               | England                     | West Midlands                                                         | kHPpSs  |
| west-sussex                                 | England                     | West Sussex                                                           | kHPpSs  |
| west-yorkshire                              | England                     | West Yorkshire                                                        | kHPpSs  |
| western-australia                           | Australia                   | Western Australia                                                     | kHPpSs  |
| western-zone                                | India                       | Western Zone                                                          | kHPpSs  |
| wielkopolskie                               | Poland                      | Województwo wielkopolskie<br />(Greater Poland Voivodeship)           | kHPpSs  |
| wiltshire                                   | England                     | Wiltshire                                                             | kHPpSs  |
| worcestershire                              | England                     | Worcestershire                                                        | kHPpSs  |
| xinjiang                                    | China                       | Xinjiang                                                              | kHPpSs  |
| yemen                                       | Asia                        | Yemen                                                                 | kHPpSs  |
| yukon                                       | Canada                      | Yukon                                                                 | kHPpSs  |
| yunnan                                      | China                       | Yunnan                                                                | kHPpSs  |
| zachodniopomorskie                          | Poland                      | Województwo zachodniopomorskie<br />(West Pomeranian Voivodeship)     | kHPpSs  |
| zambia                                      | Africa                      | Zambia                                                                | kHPpSs  |
| zeeland                                     | Netherlands                 | Zeeland                                                               | kHPpSs  |
| zhejiang                                    | China                       | Zhejiang                                                              | kHPpSs  |
| zimbabwe                                    | Africa                      | Zimbabwe                                                              | kHPpSs  |
| zlinsky                                     | Czech Republic              | Zlínský kraj                                                          | kHPpSs  |
| zuid-holland                                | Netherlands                 | Zuid-Holland                                                          | kHPpSs  |

## List of elements from openstreetmap.fr
| SHORT NAME                               | IS IN                            | LONG NAME                                | FORMATS |
|------------------------------------------|----------------------------------|------------------------------------------|---------|
| aargau                                   | switzerland                      | aargau                                   | Ps      |
| abitibi_temiscamingue                    | quebec                           | abitibi_temiscamingue                    | Ps      |
| abruzzo                                  | italy                            | abruzzo                                  | Ps      |
| aceh                                     | indonesia                        | aceh                                     | Ps      |
| acre                                     | north                            | acre                                     | Ps      |
| adygea_republic                          | southern_federal_district        | adygea_republic                          | Ps      |
| aegean                                   | turkey                           | aegean                                   | Ps      |
| afghanistan                              | asia                             | afghanistan                              | Ps      |
| africa                                   |                                  | africa                                   | Ps      |
| africa_france_taaf                       | africa                           | africa_france_taaf                       | Ps      |
| aguascalientes                           | mexico                           | aguascalientes                           | Ps      |
| ain                                      | rhone_alpes                      | ain                                      | Ps      |
| aisne                                    | picardie                         | aisne                                    | Ps      |
| akershus                                 | norway                           | akershus                                 | Ps      |
| alagoas                                  | northeast                        | alagoas                                  | Ps      |
| alameda                                  | california                       | alameda                                  | Ps      |
| aland                                    | finland                          | aland                                    | Ps      |
| alava                                    | euskadi                          | alava                                    | Ps      |
| albacete                                 | castilla_la_mancha               | albacete                                 | Ps      |
| alberta                                  | canada                           | alberta                                  | Ps      |
| algeria                                  | africa                           | algeria                                  | Ps      |
| alicante                                 | comunitat_valenciana             | alicante                                 | Ps      |
| allier                                   | auvergne                         | allier                                   | Ps      |
| almeria                                  | andalucia                        | almeria                                  | Ps      |
| alpes_de_haute_provence                  | provence_alpes_cote_d_azur       | alpes_de_haute_provence                  | Ps      |
| alpes_maritimes                          | provence_alpes_cote_d_azur       | alpes_maritimes                          | Ps      |
| alpine                                   | california                       | alpine                                   | Ps      |
| alsace                                   | france                           | alsace                                   | Ps      |
| altai_krai                               | siberian_federal_district        | altai_krai                               | Ps      |
| altai_republic                           | siberian_federal_district        | altai_republic                           | Ps      |
| amador                                   | california                       | amador                                   | Ps      |
| amapa                                    | north                            | amapa                                    | Ps      |
| amazonas                                 | north                            | amazonas                                 | Ps      |
| american_samoa                           | oceania                          | american_samoa                           | Ps      |
| amur_oblast                              | far_eastern_federal_district     | amur_oblast                              | Ps      |
| andalucia                                | spain                            | andalucia                                | Ps      |
| andaman_and_nicobar_islands              | india                            | andaman_and_nicobar_islands              | Ps      |
| andhra_pradesh                           | india                            | andhra_pradesh                           | Ps      |
| andorra                                  | europe                           | andorra                                  | Ps      |
| angola                                   | africa                           | angola                                   | Ps      |
| anguilla                                 | central-america                  | anguilla                                 | Ps      |
| anhui                                    | china                            | anhui                                    | Ps      |
| antigua_and_barbuda                      | central-america                  | antigua_and_barbuda                      | Ps      |
| antofagasta                              | chile                            | antofagasta                              | Ps      |
| antwerp                                  | flanders                         | antwerp                                  | Ps      |
| appenzell_ausserrhoden                   | switzerland                      | appenzell_ausserrhoden                   | Ps      |
| appenzell_innerrhoden                    | switzerland                      | appenzell_innerrhoden                    | Ps      |
| appomattox                               | virginia                         | appomattox                               | Ps      |
| aquitaine                                | france                           | aquitaine                                | Ps      |
| aragon                                   | spain                            | aragon                                   | Ps      |
| araucania                                | chile                            | araucania                                | Ps      |
| ardeche                                  | rhone_alpes                      | ardeche                                  | Ps      |
| ardennes                                 | champagne_ardenne                | ardennes                                 | Ps      |
| argentina                                | south-america                    | argentina                                | Ps      |
| argentina_la_rioja                       | argentina                        | argentina_la_rioja                       | Ps      |
| argentina_santa_cruz                     | argentina                        | argentina_santa_cruz                     | Ps      |
| arica_y_parinacota                       | chile                            | arica_y_parinacota                       | Ps      |
| ariege                                   | midi_pyrenees                    | ariege                                   | Ps      |
| arkhangelsk_oblast                       | northwestern_federal_district    | arkhangelsk_oblast                       | Ps      |
| armenia                                  | asia                             | armenia                                  | Ps      |
| arnsberg                                 | nordrhein_westfalen              | arnsberg                                 | Ps      |
| aruba                                    | central-america                  | aruba                                    | Ps      |
| arunachal_pradesh                        | india                            | arunachal_pradesh                        | Ps      |
| ashmore_and_cartier_islands              | australia                        | ashmore_and_cartier_islands              | Ps      |
| asia                                     |                                  | asia                                     | Ps      |
| assam                                    | india                            | assam                                    | Ps      |
| astrakhan_oblast                         | southern_federal_district        | astrakhan_oblast                         | Ps      |
| asturias                                 | spain                            | asturias                                 | Ps      |
| atacama                                  | chile                            | atacama                                  | Ps      |
| aube                                     | champagne_ardenne                | aube                                     | Ps      |
| aude                                     | languedoc_roussillon             | aude                                     | Ps      |
| aust-agder                               | norway                           | aust-agder                               | Ps      |
| australia                                | oceania                          | australia                                | Ps      |
| australian_capital_territory             | australia                        | australian_capital_territory             | Ps      |
| austria                                  | europe                           | austria                                  | Ps      |
| auvergne                                 | france                           | auvergne                                 | Ps      |
| aveiro                                   | portugal                         | aveiro                                   | Ps      |
| aveyron                                  | midi_pyrenees                    | aveyron                                  | Ps      |
| avila                                    | castilla_y_leon                  | avila                                    | Ps      |
| aysen                                    | chile                            | aysen                                    | Ps      |
| azores                                   | portugal                         | azores                                   | Ps      |
| badajoz                                  | extremadura                      | badajoz                                  | Ps      |
| bahamas                                  | central-america                  | bahamas                                  | Ps      |
| bahia                                    | northeast                        | bahia                                    | Ps      |
| bahrain                                  | asia                             | bahrain                                  | Ps      |
| baja_california                          | mexico                           | baja_california                          | Ps      |
| baja_california_sur                      | mexico                           | baja_california_sur                      | Ps      |
| bali                                     | indonesia                        | bali                                     | Ps      |
| bangka_belitung_islands                  | indonesia                        | bangka_belitung_islands                  | Ps      |
| bangsamoro                               | philippines                      | bangsamoro                               | Ps      |
| banskobystricky                          | slovakia                         | banskobystricky                          | Ps      |
| banten                                   | indonesia                        | banten                                   | Ps      |
| barbados                                 | central-america                  | barbados                                 | Ps      |
| barcelona                                | catalunya                        | barcelona                                | Ps      |
| bas_rhin                                 | alsace                           | bas_rhin                                 | Ps      |
| bas_saint_laurent                        | quebec                           | bas_saint_laurent                        | Ps      |
| basel_landschaft                         | switzerland                      | basel_landschaft                         | Ps      |
| basel_stadt                              | switzerland                      | basel_stadt                              | Ps      |
| bashkortostan_republic                   | volga_federal_district           | bashkortostan_republic                   | Ps      |
| basilicata                               | italy                            | basilicata                               | Ps      |
| basse_normandie                          | france                           | basse_normandie                          | Ps      |
| beijing                                  | china                            | beijing                                  | Ps      |
| beja                                     | portugal                         | beja                                     | Ps      |
| belgium                                  | europe                           | belgium                                  | Ps      |
| belgorod_oblast                          | central_federal_district         | belgorod_oblast                          | Ps      |
| bengkulu                                 | indonesia                        | bengkulu                                 | Ps      |
| benin                                    | africa                           | benin                                    | Ps      |
| bermuda                                  | north-america                    | bermuda                                  | Ps      |
| bern                                     | switzerland                      | bern                                     | Ps      |
| bhutan                                   | asia                             | bhutan                                   | Ps      |
| bicol_region                             | philippines                      | bicol_region                             | Ps      |
| bihar                                    | india                            | bihar                                    | Ps      |
| biobio                                   | chile                            | biobio                                   | Ps      |
| bir_tawil                                | africa                           | bir_tawil                                | Ps      |
| black_sea                                | turkey                           | black_sea                                | Ps      |
| blekinge                                 | sweden                           | blekinge                                 | Ps      |
| bouches_du_rhone                         | provence_alpes_cote_d_azur       | bouches_du_rhone                         | Ps      |
| bourgogne                                | france                           | bourgogne                                | Ps      |
| bouvet_island                            | africa                           | bouvet_island                            | Ps      |
| braga                                    | portugal                         | braga                                    | Ps      |
| braganca                                 | portugal                         | braganca                                 | Ps      |
| bratislavsky                             | slovakia                         | bratislavsky                             | Ps      |
| brazil                                   | south-america                    | brazil                                   | Ps      |
| brazil_central-west                      | brazil                           | brazil_central-west                      | Ps      |
| brazil_north                             | brazil                           | brazil_north                             | Ps      |
| brazil_northeast                         | brazil                           | brazil_northeast                         | Ps      |
| brazil_south                             | brazil                           | brazil_south                             | Ps      |
| brazil_southeast                         | brazil                           | brazil_southeast                         | Ps      |
| bretagne                                 | france                           | bretagne                                 | Ps      |
| british_columbia                         | canada                           | british_columbia                         | Ps      |
| british_indian_ocean_territory           | asia                             | british_indian_ocean_territory           | Ps      |
| british_virgin_islands                   | central-america                  | british_virgin_islands                   | Ps      |
| brunei                                   | asia                             | brunei                                   | Ps      |
| brussels_capital_region                  | belgium                          | brussels_capital_region                  | Ps      |
| bryansk_oblast                           | central_federal_district         | bryansk_oblast                           | Ps      |
| buenos_aires                             | argentina                        | buenos_aires                             | Ps      |
| buenos_aires_city                        | argentina                        | buenos_aires_city                        | Ps      |
| burgenland                               | austria                          | burgenland                               | Ps      |
| burgos                                   | castilla_y_leon                  | burgos                                   | Ps      |
| burkina_faso                             | africa                           | burkina_faso                             | Ps      |
| burundi                                  | africa                           | burundi                                  | Ps      |
| buryatia_republic                        | siberian_federal_district        | buryatia_republic                        | Ps      |
| buskerud                                 | norway                           | buskerud                                 | Ps      |
| butte                                    | california                       | butte                                    | Ps      |
| cabo_delgado                             | mozambique                       | cabo_delgado                             | Ps      |
| caceres                                  | extremadura                      | caceres                                  | Ps      |
| cadiz                                    | andalucia                        | cadiz                                    | Ps      |
| cagayan_valley                           | philippines                      | cagayan_valley                           | Ps      |
| calabarzon                               | philippines                      | calabarzon                               | Ps      |
| calabria                                 | italy                            | calabria                                 | Ps      |
| calaveras                                | california                       | calaveras                                | Ps      |
| california                               | us-west                          | california                               | Ps      |
| california_lake                          | california                       | california_lake                          | Ps      |
| california_santa_cruz                    | california                       | california_santa_cruz                    | Ps      |
| calvados                                 | basse_normandie                  | calvados                                 | Ps      |
| cambodia                                 | asia                             | cambodia                                 | Ps      |
| cameroon                                 | africa                           | cameroon                                 | Ps      |
| campania                                 | italy                            | campania                                 | Ps      |
| campeche                                 | mexico                           | campeche                                 | Ps      |
| canada                                   | north-america                    | canada                                   | Ps      |
| canarias                                 | spain                            | canarias                                 | Ps      |
| cantabria                                | spain                            | cantabria                                | Ps      |
| cantal                                   | auvergne                         | cantal                                   | Ps      |
| cape_verde                               | africa                           | cape_verde                               | Ps      |
| capital_district                         | new-york                         | capital_district                         | Ps      |
| capitale_nationale                       | quebec                           | capitale_nationale                       | Ps      |
| caraga                                   | philippines                      | caraga                                   | Ps      |
| caribbean                                | central-america                  | caribbean                                | Ps      |
| castellon                                | comunitat_valenciana             | castellon                                | Ps      |
| castelo_branco                           | portugal                         | castelo_branco                           | Ps      |
| castilla_la_mancha                       | spain                            | castilla_la_mancha                       | Ps      |
| castilla_y_leon                          | spain                            | castilla_y_leon                          | Ps      |
| catalunya                                | spain                            | catalunya                                | Ps      |
| catamarca                                | argentina                        | catamarca                                | Ps      |
| cayman_islands                           | central-america                  | cayman_islands                           | Ps      |
| ceara                                    | northeast                        | ceara                                    | Ps      |
| central-america                          |                                  | central-america                          | Ps      |
| central-west                             | brazil                           | central-west                             |         |
| central_african_republic                 | africa                           | central_african_republic                 | Ps      |
| central_anatolia                         | turkey                           | central_anatolia                         | Ps      |
| central_federal_district                 | russia                           | central_federal_district                 | Ps      |
| central_finland                          | finland                          | central_finland                          | Ps      |
| central_java                             | indonesia                        | central_java                             | Ps      |
| central_kalimantan                       | indonesia                        | central_kalimantan                       | Ps      |
| central_luzon                            | philippines                      | central_luzon                            | Ps      |
| central_new_york                         | new-york                         | central_new_york                         | Ps      |
| central_ontario                          | ontario                          | central_ontario                          | Ps      |
| central_ostrobothnia                     | finland                          | central_ostrobothnia                     | Ps      |
| central_papua                            | indonesia                        | central_papua                            | Ps      |
| central_sulawesi                         | indonesia                        | central_sulawesi                         | Ps      |
| central_visayas                          | philippines                      | central_visayas                          | Ps      |
| centre                                   | france                           | centre                                   | Ps      |
| centre_du_quebec                         | quebec                           | centre_du_quebec                         | Ps      |
| ceuta                                    | spain                            | ceuta                                    | Ps      |
| chaco                                    | argentina                        | chaco                                    | Ps      |
| chad                                     | africa                           | chad                                     | Ps      |
| champagne_ardenne                        | france                           | champagne_ardenne                        | Ps      |
| chandigarh                               | india                            | chandigarh                               | Ps      |
| charente                                 | poitou_charentes                 | charente                                 | Ps      |
| charente_maritime                        | poitou_charentes                 | charente_maritime                        | Ps      |
| chaudiere_appalaches                     | quebec                           | chaudiere_appalaches                     | Ps      |
| chechen_republic                         | north_caucasian_federal_district | chechen_republic                         | Ps      |
| chelyabinsk_oblast                       | ural_federal_district            | chelyabinsk_oblast                       | Ps      |
| cher                                     | centre                           | cher                                     | Ps      |
| cherkasy_oblast                          | ukraine                          | cherkasy_oblast                          | Ps      |
| chernihiv_oblast                         | ukraine                          | chernihiv_oblast                         | Ps      |
| chernivtsi_oblast                        | ukraine                          | chernivtsi_oblast                        | Ps      |
| chesapeake                               | virginia                         | chesapeake                               | Ps      |
| chhattisgarh                             | india                            | chhattisgarh                             | Ps      |
| chiapas                                  | mexico                           | chiapas                                  | Ps      |
| chihuahua                                | mexico                           | chihuahua                                | Ps      |
| chile                                    | south-america                    | chile                                    | Ps      |
| china                                    | asia                             | china                                    | Ps      |
| chongqing                                | china                            | chongqing                                | Ps      |
| christmas_island                         | australia                        | christmas_island                         | Ps      |
| chubu                                    | japan                            | chubu                                    | Ps      |
| chubut                                   | argentina                        | chubut                                   | Ps      |
| chugoku                                  | japan                            | chugoku                                  | Ps      |
| chukotka_autonomous_okrug                | far_eastern_federal_district     | chukotka_autonomous_okrug                | Ps      |
| chuvash_republic                         | volga_federal_district           | chuvash_republic                         | Ps      |
| ciudad_real                              | castilla_la_mancha               | ciudad_real                              | Ps      |
| clipperton                               | oceania                          | clipperton                               | Ps      |
| coahuila                                 | mexico                           | coahuila                                 | Ps      |
| coastal                                  | tanzania                         | coastal                                  | Ps      |
| cocos_islands                            | australia                        | cocos_islands                            | Ps      |
| coimbra                                  | portugal                         | coimbra                                  | Ps      |
| colima                                   | mexico                           | colima                                   | Ps      |
| colorado                                 | us-west                          | colorado                                 | Ps      |
| colorado_northeast                       | colorado                         | colorado_northeast                       | Ps      |
| colorado_northwest                       | colorado                         | colorado_northwest                       | Ps      |
| colorado_southeast                       | colorado                         | colorado_southeast                       | Ps      |
| colorado_southwest                       | colorado                         | colorado_southwest                       | Ps      |
| colusa                                   | california                       | colusa                                   | Ps      |
| comoros                                  | africa                           | comoros                                  | Ps      |
| comunidad_de_madrid                      | spain                            | comunidad_de_madrid                      | Ps      |
| comunidad_foral_de_navarra               | spain                            | comunidad_foral_de_navarra               | Ps      |
| comunitat_valenciana                     | spain                            | comunitat_valenciana                     | Ps      |
| congo_brazzaville                        | africa                           | congo_brazzaville                        | Ps      |
| congo_kinshasa                           | africa                           | congo_kinshasa                           | Ps      |
| contra_costa                             | california                       | contra_costa                             | Ps      |
| cook                                     | illinois                         | cook                                     | Ps      |
| cook_islands                             | oceania                          | cook_islands                             | Ps      |
| coquimbo                                 | chile                            | coquimbo                                 | Ps      |
| coral_sea_islands                        | australia                        | coral_sea_islands                        | Ps      |
| cordillera_administrative_region         | philippines                      | cordillera_administrative_region         | Ps      |
| cordoba                                  | argentina                        | cordoba                                  | Ps      |
| correze                                  | limousin                         | correze                                  | Ps      |
| corrientes                               | argentina                        | corrientes                               | Ps      |
| corse                                    | france                           | corse                                    | Ps      |
| corse_du_sud                             | corse                            | corse_du_sud                             | Ps      |
| costa_rica                               | central-america                  | costa_rica                               | Ps      |
| cote_d_or                                | bourgogne                        | cote_d_or                                | Ps      |
| cote_nord                                | quebec                           | cote_nord                                | Ps      |
| cotes_d_armor                            | bretagne                         | cotes_d_armor                            | Ps      |
| creuse                                   | limousin                         | creuse                                   | Ps      |
| crimea                                   | ukraine                          | crimea                                   | Ps      |
| crimea_republic                          | southern_federal_district        | crimea_republic                          | Ps      |
| cuenca                                   | castilla_la_mancha               | cuenca                                   | Ps      |
| culpeper                                 | virginia                         | culpeper                                 | Ps      |
| curacao                                  | central-america                  | curacao                                  | Ps      |
| czech_republic                           | europe                           | czech_republic                           | Ps      |
| dadra_and_nagar_haveli                   | india                            | dadra_and_nagar_haveli                   | Ps      |
| dadra_and_nagar_haveli_and_daman_and_diu | india                            | dadra_and_nagar_haveli_and_daman_and_diu | Ps      |
| dagestan_republic                        | north_caucasian_federal_district | dagestan_republic                        | Ps      |
| dalarna                                  | sweden                           | dalarna                                  | Ps      |
| daman_and_diu                            | india                            | daman_and_diu                            | Ps      |
| davao_region                             | philippines                      | davao_region                             | Ps      |
| del_norte                                | california                       | del_norte                                | Ps      |
| denmark                                  | europe                           | denmark                                  | Ps      |
| denver                                   | colorado                         | denver                                   | Ps      |
| detmold                                  | nordrhein_westfalen              | detmold                                  | Ps      |
| detroit_metro                            | michigan                         | detroit_metro                            | Ps      |
| deux_sevres                              | poitou_charentes                 | deux_sevres                              | Ps      |
| distrito-federal                         | central-west                     | distrito-federal                         | Ps      |
| djibouti                                 | africa                           | djibouti                                 | Ps      |
| dnipropetrovsk_oblast                    | ukraine                          | dnipropetrovsk_oblast                    | Ps      |
| dolnoslaskie                             | poland                           | dolnoslaskie                             | Ps      |
| dominica                                 | central-america                  | dominica                                 | Ps      |
| dominican_republic                       | central-america                  | dominican_republic                       | Ps      |
| donetsk_oblast                           | ukraine                          | donetsk_oblast                           | Ps      |
| dordogne                                 | aquitaine                        | dordogne                                 | Ps      |
| doubs                                    | franche_comte                    | doubs                                    | Ps      |
| drenthe                                  | netherlands                      | drenthe                                  | Ps      |
| drome                                    | rhone_alpes                      | drome                                    | Ps      |
| durango                                  | mexico                           | durango                                  | Ps      |
| dusseldorf                               | nordrhein_westfalen              | dusseldorf                               | Ps      |
| east_flanders                            | flanders                         | east_flanders                            | Ps      |
| east_java                                | indonesia                        | east_java                                | Ps      |
| east_kalimantan                          | indonesia                        | east_kalimantan                          | Ps      |
| east_midlands                            | england                          | east_midlands                            | Ps      |
| east_nusa_tenggara                       | indonesia                        | east_nusa_tenggara                       | Ps      |
| east_timor                               | asia                             | east_timor                               | Ps      |
| eastern_anatolia                         | turkey                           | eastern_anatolia                         | Ps      |
| eastern_cape                             | south_africa                     | eastern_cape                             | Ps      |
| eastern_ontario                          | ontario                          | eastern_ontario                          | Ps      |
| eastern_visayas                          | philippines                      | eastern_visayas                          | Ps      |
| egypt                                    | africa                           | egypt                                    | Ps      |
| el_dorado                                | california                       | el_dorado                                | Ps      |
| el_salvador                              | central-america                  | el_salvador                              | Ps      |
| emilia_romagna                           | italy                            | emilia_romagna                           | Ps      |
| england                                  | united_kingdom                   | england                                  | Ps      |
| england_east                             | england                          | england_east                             | Ps      |
| england_north_east                       | england                          | england_north_east                       | Ps      |
| england_north_west                       | england                          | england_north_west                       | Ps      |
| england_south_east                       | england                          | england_south_east                       | Ps      |
| england_south_west                       | england                          | england_south_west                       | Ps      |
| entre_rios                               | argentina                        | entre_rios                               | Ps      |
| equatorial_guinea                        | africa                           | equatorial_guinea                        | Ps      |
| eritrea                                  | africa                           | eritrea                                  | Ps      |
| espirito-santo                           | southeast                        | espirito-santo                           | Ps      |
| essonne                                  | ile_de_france                    | essonne                                  | Ps      |
| estrie                                   | quebec                           | estrie                                   | Ps      |
| ethiopia                                 | africa                           | ethiopia                                 | Ps      |
| eure                                     | haute_normandie                  | eure                                     | Ps      |
| eure_et_loir                             | centre                           | eure_et_loir                             | Ps      |
| europe                                   |                                  | europe                                   | Ps      |
| euskadi                                  | spain                            | euskadi                                  | Ps      |
| evora                                    | portugal                         | evora                                    | Ps      |
| extremadura                              | spain                            | extremadura                              | Ps      |
| fairfax                                  | virginia                         | fairfax                                  | Ps      |
| falkland                                 | south-america                    | falkland                                 | Ps      |
| far_eastern_federal_district             | russia                           | far_eastern_federal_district             | Ps      |
| far_west                                 | new_south_wales                  | far_west                                 | Ps      |
| faro                                     | portugal                         | faro                                     | Ps      |
| fiji                                     | merge                            | fiji                                     | Ps      |
| fiji_east                                | oceania                          | fiji_east                                | Ps      |
| fiji_west                                | oceania                          | fiji_west                                | Ps      |
| finger_lakes                             | new-york                         | finger_lakes                             | Ps      |
| finistere                                | bretagne                         | finistere                                | Ps      |
| finland                                  | europe                           | finland                                  | Ps      |
| finnmark                                 | norway                           | finnmark                                 | Ps      |
| flanders                                 | belgium                          | flanders                                 | Ps      |
| flemish_brabant                          | flanders                         | flemish_brabant                          | Ps      |
| flevoland                                | netherlands                      | flevoland                                | Ps      |
| florida                                  | us-south                         | florida                                  | Ps      |
| florida_east_central                     | florida                          | florida_east_central                     | Ps      |
| florida_northeast                        | florida                          | florida_northeast                        | Ps      |
| florida_northwest                        | florida                          | florida_northwest                        | Ps      |
| florida_southwest                        | florida                          | florida_southwest                        | Ps      |
| formosa                                  | argentina                        | formosa                                  | Ps      |
| france                                   | europe                           | france                                   | Ps      |
| france_metro_dom_com_nc                  | merge                            | france_metro_dom_com_nc                  | Ps      |
| franche_comte                            | france                           | franche_comte                            | Ps      |
| franche_comte_jura                       | franche_comte                    | franche_comte_jura                       | Ps      |
| free_state                               | south_africa                     | free_state                               | Ps      |
| fresno                                   | california                       | fresno                                   | Ps      |
| fribourg                                 | switzerland                      | fribourg                                 | Ps      |
| friesland                                | netherlands                      | friesland                                | Ps      |
| friuli_venezia_giulia                    | italy                            | friuli_venezia_giulia                    | Ps      |
| fujian                                   | china                            | fujian                                   | Ps      |
| gabon                                    | africa                           | gabon                                    | Ps      |
| galicia                                  | spain                            | galicia                                  | Ps      |
| gambia                                   | africa                           | gambia                                   | Ps      |
| gansu                                    | china                            | gansu                                    | Ps      |
| gard                                     | languedoc_roussillon             | gard                                     | Ps      |
| gaspesie_iles_de_la_madeleine            | quebec                           | gaspesie_iles_de_la_madeleine            | Ps      |
| gatorland                                | florida                          | gatorland                                | Ps      |
| gauteng                                  | south_africa                     | gauteng                                  | Ps      |
| gavleborg                                | sweden                           | gavleborg                                | Ps      |
| gaza                                     | mozambique                       | gaza                                     | Ps      |
| gelderland                               | netherlands                      | gelderland                               | Ps      |
| geneva                                   | switzerland                      | geneva                                   | Ps      |
| georgia                                  | asia                             | georgia                                  | Ps      |
| georgia_northeast                        | georgia                          | georgia_northeast                        | Ps      |
| georgia_northwest                        | georgia                          | georgia_northwest                        | Ps      |
| georgia_southeast                        | georgia                          | georgia_southeast                        | Ps      |
| georgia_southwest                        | georgia                          | georgia_southwest                        | Ps      |
| germany                                  | europe                           | germany                                  | Ps      |
| gers                                     | midi_pyrenees                    | gers                                     | Ps      |
| ghana                                    | africa                           | ghana                                    | Ps      |
| gibraltar                                | europe                           | gibraltar                                | Ps      |
| girona                                   | catalunya                        | girona                                   | Ps      |
| gironde                                  | aquitaine                        | gironde                                  | Ps      |
| glarus                                   | switzerland                      | glarus                                   | Ps      |
| glenn                                    | california                       | glenn                                    | Ps      |
| goa                                      | india                            | goa                                      | Ps      |
| goias                                    | central-west                     | goias                                    | Ps      |
| gold_coast                               | florida                          | gold_coast                               | Ps      |
| golden_horseshoe                         | ontario                          | golden_horseshoe                         | Ps      |
| gorontalo                                | indonesia                        | gorontalo                                | Ps      |
| gotland                                  | sweden                           | gotland                                  | Ps      |
| granada                                  | andalucia                        | granada                                  | Ps      |
| greater_london                           | england                          | greater_london                           | Ps      |
| greater_metropolitan_newcastle           | new_south_wales                  | greater_metropolitan_newcastle           | Ps      |
| greater_metropolitan_sydney              | new_south_wales                  | greater_metropolitan_sydney              | Ps      |
| grenada                                  | central-america                  | grenada                                  | Ps      |
| grisons                                  | switzerland                      | grisons                                  | Ps      |
| groningen                                | netherlands                      | groningen                                | Ps      |
| guadalajara                              | castilla_la_mancha               | guadalajara                              | Ps      |
| guadeloupe                               | central-america                  | guadeloupe                               | Ps      |
| guam                                     | oceania                          | guam                                     | Ps      |
| guanajuato                               | mexico                           | guanajuato                               | Ps      |
| guangdong                                | china                            | guangdong                                | Ps      |
| guangxi                                  | china                            | guangxi                                  | Ps      |
| guarda                                   | portugal                         | guarda                                   | Ps      |
| guernesey                                | europe                           | guernesey                                | Ps      |
| guerrero                                 | mexico                           | guerrero                                 | Ps      |
| guinea                                   | africa                           | guinea                                   | Ps      |
| guipuzcoa                                | euskadi                          | guipuzcoa                                | Ps      |
| guizhou                                  | china                            | guizhou                                  | Ps      |
| gujarat                                  | india                            | gujarat                                  | Ps      |
| guyana                                   | south-america                    | guyana                                   | Ps      |
| guyane                                   | south-america                    | guyane                                   | Ps      |
| hainan                                   | china                            | hainan                                   | Ps      |
| haiti                                    | central-america                  | haiti                                    | Ps      |
| halland                                  | sweden                           | halland                                  | Ps      |
| haryana                                  | india                            | haryana                                  | Ps      |
| haut_rhin                                | alsace                           | haut_rhin                                | Ps      |
| haute_corse                              | corse                            | haute_corse                              | Ps      |
| haute_garonne                            | midi_pyrenees                    | haute_garonne                            | Ps      |
| haute_loire                              | auvergne                         | haute_loire                              | Ps      |
| haute_marne                              | champagne_ardenne                | haute_marne                              | Ps      |
| haute_normandie                          | france                           | haute_normandie                          | Ps      |
| haute_saone                              | franche_comte                    | haute_saone                              | Ps      |
| haute_savoie                             | rhone_alpes                      | haute_savoie                             | Ps      |
| haute_vienne                             | limousin                         | haute_vienne                             | Ps      |
| hautes_alpes                             | provence_alpes_cote_d_azur       | hautes_alpes                             | Ps      |
| hautes_pyrenees                          | midi_pyrenees                    | hautes_pyrenees                          | Ps      |
| hauts_de_seine                           | ile_de_france                    | hauts_de_seine                           | Ps      |
| heard_island_and_mcdonald_slands         | australia                        | heard_island_and_mcdonald_slands         | Ps      |
| hebei                                    | china                            | hebei                                    | Ps      |
| hedmark                                  | norway                           | hedmark                                  | Ps      |
| heilongjiang                             | china                            | heilongjiang                             | Ps      |
| henan                                    | china                            | henan                                    | Ps      |
| herault                                  | languedoc_roussillon             | herault                                  | Ps      |
| hidalgo                                  | mexico                           | hidalgo                                  | Ps      |
| highland_papua                           | indonesia                        | highland_papua                           | Ps      |
| himachal_pradesh                         | india                            | himachal_pradesh                         | Ps      |
| hokkaido                                 | japan                            | hokkaido                                 | Ps      |
| honduras                                 | central-america                  | honduras                                 | Ps      |
| hong_kong                                | china                            | hong_kong                                | Ps      |
| hordaland                                | norway                           | hordaland                                | Ps      |
| hovedstaden                              | denmark                          | hovedstaden                              | Ps      |
| hubei                                    | china                            | hubei                                    | Ps      |
| hudson_valley                            | new-york                         | hudson_valley                            | Ps      |
| huelva                                   | andalucia                        | huelva                                   | Ps      |
| huesca                                   | aragon                           | huesca                                   | Ps      |
| humboldt                                 | california                       | humboldt                                 | Ps      |
| hunan                                    | china                            | hunan                                    | Ps      |
| ile_de_france                            | france                           | ile_de_france                            | Ps      |
| ilemi                                    | africa                           | ilemi                                    | Ps      |
| illawarra                                | new_south_wales                  | illawarra                                | Ps      |
| ille_et_vilaine                          | bretagne                         | ille_et_vilaine                          | Ps      |
| illes_balears                            | spain                            | illes_balears                            | Ps      |
| illinois                                 | us-midwest                       | illinois                                 | Ps      |
| illinois_central                         | illinois                         | illinois_central                         | Ps      |
| illinois_east_central                    | illinois                         | illinois_east_central                    | Ps      |
| illinois_north                           | illinois                         | illinois_north                           | Ps      |
| illinois_northeast                       | illinois                         | illinois_northeast                       | Ps      |
| illinois_northwest                       | illinois                         | illinois_northwest                       | Ps      |
| illinois_southern                        | illinois                         | illinois_southern                        | Ps      |
| illinois_southwest                       | illinois                         | illinois_southwest                       | Ps      |
| ilocos_region                            | philippines                      | ilocos_region                            | Ps      |
| imperial                                 | california                       | imperial                                 | Ps      |
| india                                    | asia                             | india                                    | Ps      |
| indonesia                                | asia                             | indonesia                                | Ps      |
| indre                                    | centre                           | indre                                    | Ps      |
| indre_et_loire                           | centre                           | indre_et_loire                           | Ps      |
| ingushetia_republic                      | north_caucasian_federal_district | ingushetia_republic                      | Ps      |
| inhambane                                | mozambique                       | inhambane                                | Ps      |
| inner_mongolia                           | china                            | inner_mongolia                           | Ps      |
| inyo                                     | california                       | inyo                                     | Ps      |
| ionian_sea                               | seas                             | ionian_sea                               | Ps      |
| ireland                                  | europe                           | ireland                                  | Ps      |
| irkutsk_oblast                           | siberian_federal_district        | irkutsk_oblast                           | Ps      |
| isere                                    | rhone_alpes                      | isere                                    | Ps      |
| israel                                   | asia                             | israel                                   | Ps      |
| israel_and_palestine                     | asia                             | israel_and_palestine                     | Ps      |
| israel_west_bank                         | asia                             | israel_west_bank                         | Ps      |
| italy                                    | europe                           | italy                                    | Ps      |
| ivano-frankivsk_oblast                   | ukraine                          | ivano-frankivsk_oblast                   | Ps      |
| ivanovo_oblast                           | central_federal_district         | ivanovo_oblast                           | Ps      |
| ivory_coast                              | africa                           | ivory_coast                              | Ps      |
| jaen                                     | andalucia                        | jaen                                     | Ps      |
| jakarta                                  | indonesia                        | jakarta                                  | Ps      |
| jalisco                                  | mexico                           | jalisco                                  | Ps      |
| jamaica                                  | central-america                  | jamaica                                  | Ps      |
| jambi                                    | indonesia                        | jambi                                    | Ps      |
| jammu_and_kashmir                        | india                            | jammu_and_kashmir                        | Ps      |
| jamtland                                 | sweden                           | jamtland                                 | Ps      |
| jan_mayen                                | norway                           | jan_mayen                                | Ps      |
| japan                                    | asia                             | japan                                    | Ps      |
| jersey                                   | europe                           | jersey                                   | Ps      |
| jewish_autonomous_oblast                 | far_eastern_federal_district     | jewish_autonomous_oblast                 | Ps      |
| jharkhand                                | india                            | jharkhand                                | Ps      |
| jiangsu                                  | china                            | jiangsu                                  | Ps      |
| jiangxi                                  | china                            | jiangxi                                  | Ps      |
| jihocesky                                | czech_republic                   | jihocesky                                | Ps      |
| jihomoravsky                             | czech_republic                   | jihomoravsky                             | Ps      |
| jilin                                    | china                            | jilin                                    | Ps      |
| jonkoping                                | sweden                           | jonkoping                                | Ps      |
| jujuy                                    | argentina                        | jujuy                                    | Ps      |
| kabardino_balkar_republic                | north_caucasian_federal_district | kabardino_balkar_republic                | Ps      |
| kainuu                                   | finland                          | kainuu                                   | Ps      |
| kaliningrad_oblast                       | northwestern_federal_district    | kaliningrad_oblast                       | Ps      |
| kalmar                                   | sweden                           | kalmar                                   | Ps      |
| kalmykia_republic                        | southern_federal_district        | kalmykia_republic                        | Ps      |
| kaluga_oblast                            | central_federal_district         | kaluga_oblast                            | Ps      |
| kamchatka_krai                           | far_eastern_federal_district     | kamchatka_krai                           | Ps      |
| kansai                                   | japan                            | kansai                                   | Ps      |
| kanta_hame                               | finland                          | kanta_hame                               | Ps      |
| kanto                                    | japan                            | kanto                                    | Ps      |
| karachay_cherkess_republic               | north_caucasian_federal_district | karachay_cherkess_republic               | Ps      |
| karelia_republic                         | northwestern_federal_district    | karelia_republic                         | Ps      |
| karlovarsky                              | czech_republic                   | karlovarsky                              | Ps      |
| karnataka                                | india                            | karnataka                                | Ps      |
| karnten                                  | austria                          | karnten                                  | Ps      |
| kemerovo_oblast                          | siberian_federal_district        | kemerovo_oblast                          | Ps      |
| kenya                                    | africa                           | kenya                                    | Ps      |
| kerala                                   | india                            | kerala                                   | Ps      |
| kern                                     | california                       | kern                                     | Ps      |
| khabarovsk_krai                          | far_eastern_federal_district     | khabarovsk_krai                          | Ps      |
| khakassia_republic                       | siberian_federal_district        | khakassia_republic                       | Ps      |
| khanty_mansi_autonomous_okrug            | ural_federal_district            | khanty_mansi_autonomous_okrug            | Ps      |
| kharkiv_oblast                           | ukraine                          | kharkiv_oblast                           | Ps      |
| kherson_oblast                           | ukraine                          | kherson_oblast                           | Ps      |
| khmelnytskyi_oblast                      | ukraine                          | khmelnytskyi_oblast                      | Ps      |
| kiev                                     | ukraine                          | kiev                                     | Ps      |
| kiev_oblast                              | ukraine                          | kiev_oblast                              | Ps      |
| kings                                    | california                       | kings                                    | Ps      |
| kiribati                                 | merge                            | kiribati                                 | Ps      |
| kiribati_east                            | oceania                          | kiribati_east                            | Ps      |
| kiribati_west                            | oceania                          | kiribati_west                            | Ps      |
| kirov_oblast                             | volga_federal_district           | kirov_oblast                             | Ps      |
| kirovohrad_oblast                        | ukraine                          | kirovohrad_oblast                        | Ps      |
| koln                                     | nordrhein_westfalen              | koln                                     | Ps      |
| komi_republic                            | northwestern_federal_district    | komi_republic                            | Ps      |
| kosicky                                  | slovakia                         | kosicky                                  | Ps      |
| kostroma_oblast                          | central_federal_district         | kostroma_oblast                          | Ps      |
| kralovehradecky                          | czech_republic                   | kralovehradecky                          | Ps      |
| krasnodar_krai                           | southern_federal_district        | krasnodar_krai                           | Ps      |
| krasnoyarsk_krai                         | siberian_federal_district        | krasnoyarsk_krai                         | Ps      |
| kronoberg                                | sweden                           | kronoberg                                | Ps      |
| kujawsko_pomorskie                       | poland                           | kujawsko_pomorskie                       | Ps      |
| kurgan_oblast                            | ural_federal_district            | kurgan_oblast                            | Ps      |
| kursk_oblast                             | central_federal_district         | kursk_oblast                             | Ps      |
| kuwait                                   | asia                             | kuwait                                   | Ps      |
| kwazulu_natal                            | south_africa                     | kwazulu_natal                            | Ps      |
| kymenlaakso                              | finland                          | kymenlaakso                              | Ps      |
| kyushu                                   | japan                            | kyushu                                   | Ps      |
| la_coruna                                | galicia                          | la_coruna                                | Ps      |
| la_pampa                                 | argentina                        | la_pampa                                 | Ps      |
| ladakh                                   | india                            | ladakh                                   | Ps      |
| lakshadweep                              | india                            | lakshadweep                              | Ps      |
| lampung                                  | indonesia                        | lampung                                  | Ps      |
| lanaudiere                               | quebec                           | lanaudiere                               | Ps      |
| landes                                   | aquitaine                        | landes                                   | Ps      |
| languedoc_roussillon                     | france                           | languedoc_roussillon                     | Ps      |
| laos                                     | asia                             | laos                                     | Ps      |
| lapland                                  | finland                          | lapland                                  | Ps      |
| las_palmas                               | canarias                         | las_palmas                               | Ps      |
| lassen                                   | california                       | lassen                                   | Ps      |
| laurentides                              | quebec                           | laurentides                              | Ps      |
| laval                                    | quebec                           | laval                                    | Ps      |
| lazio                                    | italy                            | lazio                                    | Ps      |
| lebanon                                  | asia                             | lebanon                                  | Ps      |
| leiria                                   | portugal                         | leiria                                   | Ps      |
| leningrad_oblast                         | northwestern_federal_district    | leningrad_oblast                         | Ps      |
| leon                                     | castilla_y_leon                  | leon                                     | Ps      |
| lesotho                                  | africa                           | lesotho                                  | Ps      |
| liaoning                                 | china                            | liaoning                                 | Ps      |
| liberecky                                | czech_republic                   | liberecky                                | Ps      |
| liguria                                  | italy                            | liguria                                  | Ps      |
| limburg                                  | netherlands                      | limburg                                  | Ps      |
| limousin                                 | france                           | limousin                                 | Ps      |
| limpopo                                  | south_africa                     | limpopo                                  | Ps      |
| lipetsk_oblast                           | central_federal_district         | lipetsk_oblast                           | Ps      |
| lisbon                                   | portugal                         | lisbon                                   | Ps      |
| lleida                                   | catalunya                        | lleida                                   | Ps      |
| lodzkie                                  | poland                           | lodzkie                                  | Ps      |
| loir_et_cher                             | centre                           | loir_et_cher                             | Ps      |
| loire                                    | rhone_alpes                      | loire                                    | Ps      |
| loire_atlantique                         | pays_de_la_loire                 | loire_atlantique                         | Ps      |
| loiret                                   | centre                           | loiret                                   | Ps      |
| lombardia                                | italy                            | lombardia                                | Ps      |
| long_island                              | new-york                         | long_island                              | Ps      |
| lorraine                                 | france                           | lorraine                                 | Ps      |
| los_angeles                              | california                       | los_angeles                              | Ps      |
| los_lagos                                | chile                            | los_lagos                                | Ps      |
| los_rios                                 | chile                            | los_rios                                 | Ps      |
| lot                                      | midi_pyrenees                    | lot                                      | Ps      |
| lot_et_garonne                           | aquitaine                        | lot_et_garonne                           | Ps      |
| lozere                                   | languedoc_roussillon             | lozere                                   | Ps      |
| lubelskie                                | poland                           | lubelskie                                | Ps      |
| lubuskie                                 | poland                           | lubuskie                                 | Ps      |
| lucerne                                  | switzerland                      | lucerne                                  | Ps      |
| lugo                                     | galicia                          | lugo                                     | Ps      |
| luhansk_oblast                           | ukraine                          | luhansk_oblast                           | Ps      |
| luxembourg                               | europe                           | luxembourg                               | Ps      |
| lviv_oblast                              | ukraine                          | lviv_oblast                              | Ps      |
| macau                                    | china                            | macau                                    | Ps      |
| madagascar                               | africa                           | madagascar                               | Ps      |
| madeira                                  | portugal                         | madeira                                  | Ps      |
| madera                                   | california                       | madera                                   | Ps      |
| madhya_pradesh                           | india                            | madhya_pradesh                           | Ps      |
| magadan_oblast                           | far_eastern_federal_district     | magadan_oblast                           | Ps      |
| magallanes                               | chile                            | magallanes                               | Ps      |
| maharashtra                              | india                            | maharashtra                              | Ps      |
| maine_et_loire                           | pays_de_la_loire                 | maine_et_loire                           | Ps      |
| malaga                                   | andalucia                        | malaga                                   | Ps      |
| malawi                                   | africa                           | malawi                                   | Ps      |
| malaysia                                 | asia                             | malaysia                                 | Ps      |
| maldives                                 | asia                             | maldives                                 | Ps      |
| mali                                     | africa                           | mali                                     | Ps      |
| malopolskie                              | poland                           | malopolskie                              | Ps      |
| maluku                                   | indonesia                        | maluku                                   | Ps      |
| manche                                   | basse_normandie                  | manche                                   | Ps      |
| manica                                   | mozambique                       | manica                                   | Ps      |
| manipur                                  | india                            | manipur                                  | Ps      |
| manitoba                                 | canada                           | manitoba                                 | Ps      |
| maputo                                   | mozambique                       | maputo                                   | Ps      |
| maputo_city                              | mozambique                       | maputo_city                              | Ps      |
| maranhao                                 | northeast                        | maranhao                                 | Ps      |
| marche                                   | italy                            | marche                                   | Ps      |
| mari_el_republic                         | volga_federal_district           | mari_el_republic                         | Ps      |
| marin                                    | california                       | marin                                    | Ps      |
| mariposa                                 | california                       | mariposa                                 | Ps      |
| marmara                                  | turkey                           | marmara                                  | Ps      |
| marne                                    | champagne_ardenne                | marne                                    | Ps      |
| marshall-islands                         | oceania                          | marshall-islands                         | Ps      |
| marshall_islands                         | oceania                          | marshall_islands                         | Ps      |
| martinique                               | central-america                  | martinique                               | Ps      |
| mato-grosso                              | central-west                     | mato-grosso                              | Ps      |
| mato-grosso-do-sul                       | central-west                     | mato-grosso-do-sul                       | Ps      |
| maule                                    | chile                            | maule                                    | Ps      |
| mauricie                                 | quebec                           | mauricie                                 | Ps      |
| mauritania                               | africa                           | mauritania                               | Ps      |
| mauritius                                | africa                           | mauritius                                | Ps      |
| mayenne                                  | pays_de_la_loire                 | mayenne                                  | Ps      |
| mayotte                                  | africa                           | mayotte                                  | Ps      |
| mazowieckie                              | poland                           | mazowieckie                              | Ps      |
| mediterranean                            | turkey                           | mediterranean                            | Ps      |
| meghalaya                                | india                            | meghalaya                                | Ps      |
| melilla                                  | spain                            | melilla                                  | Ps      |
| mendocino                                | california                       | mendocino                                | Ps      |
| mendoza                                  | argentina                        | mendoza                                  | Ps      |
| merced                                   | california                       | merced                                   | Ps      |
| merge                                    |                                  | merge                                    |         |
| merge_france_taaf                        | merge                            | merge_france_taaf                        | Ps      |
| metro_manila                             | philippines                      | metro_manila                             | Ps      |
| meurthe_et_moselle                       | lorraine                         | meurthe_et_moselle                       | Ps      |
| meuse                                    | lorraine                         | meuse                                    | Ps      |
| mexico                                   | north-america                    | mexico                                   | Ps      |
| mexico_city                              | mexico                           | mexico_city                              | Ps      |
| michigan                                 | us-midwest                       | michigan                                 | Ps      |
| michigan_central                         | michigan                         | michigan_central                         | Ps      |
| michigan_southeast                       | michigan                         | michigan_southeast                       | Ps      |
| michigan_southwest                       | michigan                         | michigan_southwest                       | Ps      |
| michigan_west                            | michigan                         | michigan_west                            | Ps      |
| michoacan                                | mexico                           | michoacan                                | Ps      |
| micronesia                               | oceania                          | micronesia                               | Ps      |
| mid_north_coast                          | new_south_wales                  | mid_north_coast                          | Ps      |
| midi_pyrenees                            | france                           | midi_pyrenees                            | Ps      |
| midtjylland                              | denmark                          | midtjylland                              | Ps      |
| mimaropa                                 | philippines                      | mimaropa                                 | Ps      |
| minas-gerais                             | southeast                        | minas-gerais                             | Ps      |
| misiones                                 | argentina                        | misiones                                 | Ps      |
| mizoram                                  | india                            | mizoram                                  | Ps      |
| modoc                                    | california                       | modoc                                    | Ps      |
| moere_og_romsdal                         | norway                           | moere_og_romsdal                         | Ps      |
| mohawk_valley                            | new-york                         | mohawk_valley                            | Ps      |
| molise                                   | italy                            | molise                                   | Ps      |
| monaco                                   | europe                           | monaco                                   | Ps      |
| mono                                     | california                       | mono                                     | Ps      |
| monteregie                               | quebec                           | monteregie                               | Ps      |
| monterey                                 | california                       | monterey                                 | Ps      |
| montreal                                 | quebec                           | montreal                                 | Ps      |
| montserrat                               | central-america                  | montserrat                               | Ps      |
| moravskoslezsky                          | czech_republic                   | moravskoslezsky                          | Ps      |
| morbihan                                 | bretagne                         | morbihan                                 | Ps      |
| mordovia_republic                        | volga_federal_district           | mordovia_republic                        | Ps      |
| morelos                                  | mexico                           | morelos                                  | Ps      |
| morocco                                  | africa                           | morocco                                  | Ps      |
| moscow                                   | central_federal_district         | moscow                                   | Ps      |
| moscow_oblast                            | central_federal_district         | moscow_oblast                            | Ps      |
| moselle                                  | lorraine                         | moselle                                  | Ps      |
| mozambique                               | africa                           | mozambique                               | Ps      |
| mpumalanga                               | south_africa                     | mpumalanga                               | Ps      |
| munster                                  | nordrhein_westfalen              | munster                                  | Ps      |
| murmansk_oblast                          | northwestern_federal_district    | murmansk_oblast                          | Ps      |
| murray                                   | new_south_wales                  | murray                                   | Ps      |
| myanmar                                  | asia                             | myanmar                                  | Ps      |
| mykolaiv_oblast                          | ukraine                          | mykolaiv_oblast                          | Ps      |
| nagaland                                 | india                            | nagaland                                 | Ps      |
| namibia                                  | africa                           | namibia                                  | Ps      |
| nampula                                  | mozambique                       | nampula                                  | Ps      |
| napa                                     | california                       | napa                                     | Ps      |
| national_capital_territory_of_delhi      | india                            | national_capital_territory_of_delhi      | Ps      |
| nauru                                    | oceania                          | nauru                                    | Ps      |
| nayarit                                  | mexico                           | nayarit                                  | Ps      |
| negros_island_region                     | philippines                      | negros_island_region                     | Ps      |
| nenets_autonomous_okrug                  | northwestern_federal_district    | nenets_autonomous_okrug                  | Ps      |
| netherlands                              | europe                           | netherlands                              | Ps      |
| neuchatel                                | switzerland                      | neuchatel                                | Ps      |
| neuquen                                  | argentina                        | neuquen                                  | Ps      |
| nevada                                   | california                       | nevada                                   | Ps      |
| new-york                                 | us-northeast                     | new-york                                 | Ps      |
| new_brunswick                            | canada                           | new_brunswick                            | Ps      |
| new_caledonia                            | oceania                          | new_caledonia                            | Ps      |
| new_south_wales                          | australia                        | new_south_wales                          | Ps      |
| new_south_wales_central_west             | new_south_wales                  | new_south_wales_central_west             | Ps      |
| new_south_wales_north_western            | new_south_wales                  | new_south_wales_north_western            | Ps      |
| new_south_wales_northern                 | new_south_wales                  | new_south_wales_northern                 | Ps      |
| new_york_city                            | new-york                         | new_york_city                            | Ps      |
| newfoundland_and_labrador                | canada                           | newfoundland_and_labrador                | Ps      |
| niassa                                   | mozambique                       | niassa                                   | Ps      |
| nicaragua                                | central-america                  | nicaragua                                | Ps      |
| nidwalden                                | switzerland                      | nidwalden                                | Ps      |
| niederosterreich                         | austria                          | niederosterreich                         | Ps      |
| nievre                                   | bourgogne                        | nievre                                   | Ps      |
| niger                                    | africa                           | niger                                    | Ps      |
| nigeria                                  | africa                           | nigeria                                  | Ps      |
| nigeria_north_central                    | nigeria                          | nigeria_north_central                    | Ps      |
| nigeria_north_east                       | nigeria                          | nigeria_north_east                       | Ps      |
| nigeria_north_west                       | nigeria                          | nigeria_north_west                       | Ps      |
| nigeria_south_east                       | nigeria                          | nigeria_south_east                       | Ps      |
| nigeria_south_south                      | nigeria                          | nigeria_south_south                      | Ps      |
| nigeria_south_west                       | nigeria                          | nigeria_south_west                       | Ps      |
| ningxia                                  | china                            | ningxia                                  | Ps      |
| nitriansky                               | slovakia                         | nitriansky                               | Ps      |
| niue                                     | oceania                          | niue                                     | Ps      |
| nizhny_novgorod_oblast                   | volga_federal_district           | nizhny_novgorod_oblast                   | Ps      |
| noord_brabant                            | netherlands                      | noord_brabant                            | Ps      |
| noord_holland                            | netherlands                      | noord_holland                            | Ps      |
| nord                                     | nord_pas_de_calais               | nord                                     | Ps      |
| nord_du_quebec                           | quebec                           | nord_du_quebec                           | Ps      |
| nord_pas_de_calais                       | france                           | nord_pas_de_calais                       | Ps      |
| nordjylland                              | denmark                          | nordjylland                              | Ps      |
| nordland                                 | norway                           | nordland                                 | Ps      |
| nordrhein_westfalen                      | germany                          | nordrhein_westfalen                      | Ps      |
| norfolk_island                           | australia                        | norfolk_island                           | Ps      |
| norrbotten                               | sweden                           | norrbotten                               | Ps      |
| north                                    | brazil                           | north                                    |         |
| north-america                            |                                  | north-america                            | Ps      |
| north-carolina                           | us-south                         | north-carolina                           | Ps      |
| north-carolina_north_central             | north-carolina                   | north-carolina_north_central             | Ps      |
| north-carolina_northeast                 | north-carolina                   | north-carolina_northeast                 | Ps      |
| north-carolina_northwest                 | north-carolina                   | north-carolina_northwest                 | Ps      |
| north-carolina_south_central             | north-carolina                   | north-carolina_south_central             | Ps      |
| north-carolina_southeast                 | north-carolina                   | north-carolina_southeast                 | Ps      |
| north-carolina_western                   | north-carolina                   | north-carolina_western                   | Ps      |
| north_caucasian_federal_district         | russia                           | north_caucasian_federal_district         | Ps      |
| north_country                            | new-york                         | north_country                            | Ps      |
| north_kalimantan                         | indonesia                        | north_kalimantan                         | Ps      |
| north_karelia                            | finland                          | north_karelia                            | Ps      |
| north_maluku                             | indonesia                        | north_maluku                             | Ps      |
| north_metro                              | georgia                          | north_metro                              | Ps      |
| north_ossetia_alania_republic            | north_caucasian_federal_district | north_ossetia_alania_republic            | Ps      |
| north_ostrobothnia                       | finland                          | north_ostrobothnia                       | Ps      |
| north_savo                               | finland                          | north_savo                               | Ps      |
| north_sea                                | seas                             | north_sea                                | Ps      |
| north_sulawesi                           | indonesia                        | north_sulawesi                           | Ps      |
| north_sumatra                            | indonesia                        | north_sumatra                            | Ps      |
| northeast                                | brazil                           | northeast                                |         |
| northeastern_ontario                     | ontario                          | northeastern_ontario                     | Ps      |
| northern_cape                            | south_africa                     | northern_cape                            | Ps      |
| northern_ireland                         | united_kingdom                   | northern_ireland                         | Ps      |
| northern_lower                           | michigan                         | northern_lower                           | Ps      |
| northern_mariana_islands                 | oceania                          | northern_mariana_islands                 | Ps      |
| northern_mindanao                        | philippines                      | northern_mindanao                        | Ps      |
| northern_territory                       | australia                        | northern_territory                       | Ps      |
| northwest_territories                    | canada                           | northwest_territories                    | Ps      |
| northwestern_federal_district            | russia                           | northwestern_federal_district            | Ps      |
| northwestern_ontario                     | ontario                          | northwestern_ontario                     | Ps      |
| norway                                   | europe                           | norway                                   | Ps      |
| nova_scotia                              | canada                           | nova_scotia                              | Ps      |
| novgorod_oblast                          | northwestern_federal_district    | novgorod_oblast                          | Ps      |
| novosibirsk_oblast                       | siberian_federal_district        | novosibirsk_oblast                       | Ps      |
| nuble                                    | chile                            | nuble                                    | Ps      |
| nuevo_leon                               | mexico                           | nuevo_leon                               | Ps      |
| nunavut                                  | canada                           | nunavut                                  | Ps      |
| o_higgins                                | chile                            | o_higgins                                | Ps      |
| oaxaca                                   | mexico                           | oaxaca                                   | Ps      |
| oberosterreich                           | austria                          | oberosterreich                           | Ps      |
| obwalden                                 | switzerland                      | obwalden                                 | Ps      |
| oceania                                  |                                  | oceania                                  | Ps      |
| oceania_france_taaf                      | oceania                          | oceania_france_taaf                      | Ps      |
| odessa_oblast                            | ukraine                          | odessa_oblast                            | Ps      |
| odisha                                   | india                            | odisha                                   | Ps      |
| oestfold                                 | norway                           | oestfold                                 | Ps      |
| oise                                     | picardie                         | oise                                     | Ps      |
| olomoucky                                | czech_republic                   | olomoucky                                | Ps      |
| oman                                     | asia                             | oman                                     | Ps      |
| omsk_oblast                              | siberian_federal_district        | omsk_oblast                              | Ps      |
| ontario                                  | canada                           | ontario                                  | Ps      |
| opolskie                                 | poland                           | opolskie                                 | Ps      |
| oppland                                  | norway                           | oppland                                  | Ps      |
| orange                                   | california                       | orange                                   | Ps      |
| orebro                                   | sweden                           | orebro                                   | Ps      |
| orenburg_oblast                          | volga_federal_district           | orenburg_oblast                          | Ps      |
| orne                                     | basse_normandie                  | orne                                     | Ps      |
| oryol_oblast                             | central_federal_district         | oryol_oblast                             | Ps      |
| oslo                                     | norway                           | oslo                                     | Ps      |
| ostergotland                             | sweden                           | ostergotland                             | Ps      |
| ostrobothnia                             | finland                          | ostrobothnia                             | Ps      |
| ourense                                  | galicia                          | ourense                                  | Ps      |
| outaouais                                | quebec                           | outaouais                                | Ps      |
| overijssel                               | netherlands                      | overijssel                               | Ps      |
| paijat_hame                              | finland                          | paijat_hame                              | Ps      |
| palau                                    | oceania                          | palau                                    | Ps      |
| palencia                                 | castilla_y_leon                  | palencia                                 | Ps      |
| palestine                                | asia                             | palestine                                | Ps      |
| panama                                   | central-america                  | panama                                   | Ps      |
| panhandle                                | florida                          | panhandle                                | Ps      |
| papua                                    | indonesia                        | papua                                    | Ps      |
| papua_new_guinea                         | oceania                          | papua_new_guinea                         | Ps      |
| para                                     | north                            | para                                     | Ps      |
| paraguay                                 | south-america                    | paraguay                                 | Ps      |
| paraiba                                  | northeast                        | paraiba                                  | Ps      |
| parana                                   | south                            | parana                                   | Ps      |
| pardubicky                               | czech_republic                   | pardubicky                               | Ps      |
| paris                                    | ile_de_france                    | paris                                    | Ps      |
| pas_de_calais                            | nord_pas_de_calais               | pas_de_calais                            | Ps      |
| pays_de_la_loire                         | france                           | pays_de_la_loire                         | Ps      |
| penza_oblast                             | volga_federal_district           | penza_oblast                             | Ps      |
| perm_krai                                | volga_federal_district           | perm_krai                                | Ps      |
| pernambuco                               | northeast                        | pernambuco                               | Ps      |
| philippines                              | asia                             | philippines                              | Ps      |
| piaui                                    | northeast                        | piaui                                    | Ps      |
| picardie                                 | france                           | picardie                                 | Ps      |
| piedmont_triad                           | north-carolina                   | piedmont_triad                           | Ps      |
| piemonte                                 | italy                            | piemonte                                 | Ps      |
| pirkanmaa                                | finland                          | pirkanmaa                                | Ps      |
| pitcairn                                 | oceania                          | pitcairn                                 | Ps      |
| placer                                   | california                       | placer                                   | Ps      |
| plumas                                   | california                       | plumas                                   | Ps      |
| plzensky                                 | czech_republic                   | plzensky                                 | Ps      |
| podkarpackie                             | poland                           | podkarpackie                             | Ps      |
| podlaskie                                | poland                           | podlaskie                                | Ps      |
| poitou_charentes                         | france                           | poitou_charentes                         | Ps      |
| poland                                   | europe                           | poland                                   | Ps      |
| poltava_oblast                           | ukraine                          | poltava_oblast                           | Ps      |
| polynesie                                | oceania                          | polynesie                                | Ps      |
| pomorskie                                | poland                           | pomorskie                                | Ps      |
| pontevedra                               | galicia                          | pontevedra                               | Ps      |
| portalegre                               | portugal                         | portalegre                               | Ps      |
| porto                                    | portugal                         | porto                                    | Ps      |
| portugal                                 | europe                           | portugal                                 | Ps      |
| praha                                    | czech_republic                   | praha                                    | Ps      |
| presovsky                                | slovakia                         | presovsky                                | Ps      |
| primorsky_krai                           | far_eastern_federal_district     | primorsky_krai                           | Ps      |
| prince_edward_island                     | canada                           | prince_edward_island                     | Ps      |
| provence_alpes_cote_d_azur               | france                           | provence_alpes_cote_d_azur               | Ps      |
| pskov_oblast                             | northwestern_federal_district    | pskov_oblast                             | Ps      |
| puducherry                               | india                            | puducherry                               | Ps      |
| puebla                                   | mexico                           | puebla                                   | Ps      |
| puerto_rico                              | central-america                  | puerto_rico                              | Ps      |
| puglia                                   | italy                            | puglia                                   | Ps      |
| punjab                                   | india                            | punjab                                   | Ps      |
| puy_de_dome                              | auvergne                         | puy_de_dome                              | Ps      |
| pyrenees_atlantiques                     | aquitaine                        | pyrenees_atlantiques                     | Ps      |
| pyrenees_orientales                      | languedoc_roussillon             | pyrenees_orientales                      | Ps      |
| qatar                                    | asia                             | qatar                                    | Ps      |
| qinghai                                  | china                            | qinghai                                  | Ps      |
| quebec                                   | canada                           | quebec                                   | Ps      |
| queensland                               | australia                        | queensland                               | Ps      |
| queretaro                                | mexico                           | queretaro                                | Ps      |
| quintana_roo                             | mexico                           | quintana_roo                             | Ps      |
| rajasthan                                | india                            | rajasthan                                | Ps      |
| region_de_murcia                         | spain                            | region_de_murcia                         | Ps      |
| reunion                                  | africa                           | reunion                                  | Ps      |
| rhone                                    | rhone_alpes                      | rhone                                    | Ps      |
| rhone_alpes                              | france                           | rhone_alpes                              | Ps      |
| riau                                     | indonesia                        | riau                                     | Ps      |
| riau_islands                             | indonesia                        | riau_islands                             | Ps      |
| richmond                                 | virginia                         | richmond                                 | Ps      |
| richmond_tweed                           | new_south_wales                  | richmond_tweed                           | Ps      |
| rio-de-janeiro                           | southeast                        | rio-de-janeiro                           | Ps      |
| rio-grande-do-norte                      | northeast                        | rio-grande-do-norte                      | Ps      |
| rio-grande-do-sul                        | south                            | rio-grande-do-sul                        | Ps      |
| rio_negro                                | argentina                        | rio_negro                                | Ps      |
| riverside                                | california                       | riverside                                | Ps      |
| rivne_oblast                             | ukraine                          | rivne_oblast                             | Ps      |
| rogaland                                 | norway                           | rogaland                                 | Ps      |
| rondonia                                 | north                            | rondonia                                 | Ps      |
| roraima                                  | north                            | roraima                                  | Ps      |
| rostov_oblast                            | southern_federal_district        | rostov_oblast                            | Ps      |
| russia                                   |                                  | russia                                   | Ps      |
| rwanda                                   | africa                           | rwanda                                   | Ps      |
| ryazan_oblast                            | central_federal_district         | ryazan_oblast                            | Ps      |
| sacramento                               | california                       | sacramento                               | Ps      |
| saguenay_lac_saint_jean                  | quebec                           | saguenay_lac_saint_jean                  | Ps      |
| saint_barthelemy                         | central-america                  | saint_barthelemy                         | Ps      |
| saint_gallen                             | switzerland                      | saint_gallen                             | Ps      |
| saint_helena_ascension_tristan_da_cunha  | africa                           | saint_helena_ascension_tristan_da_cunha  | Ps      |
| saint_kitts_and_nevis                    | central-america                  | saint_kitts_and_nevis                    | Ps      |
| saint_lucia                              | central-america                  | saint_lucia                              | Ps      |
| saint_martin                             | central-america                  | saint_martin                             | Ps      |
| saint_petersburg                         | northwestern_federal_district    | saint_petersburg                         | Ps      |
| saint_pierre_et_miquelon                 | north-america                    | saint_pierre_et_miquelon                 | Ps      |
| saint_vincent_and_the_grenadines         | central-america                  | saint_vincent_and_the_grenadines         | Ps      |
| sakha_republic                           | far_eastern_federal_district     | sakha_republic                           | Ps      |
| sakhalin_oblast                          | far_eastern_federal_district     | sakhalin_oblast                          | Ps      |
| salamanca                                | castilla_y_leon                  | salamanca                                | Ps      |
| salem                                    | virginia                         | salem                                    | Ps      |
| salta                                    | argentina                        | salta                                    | Ps      |
| salzburg                                 | austria                          | salzburg                                 | Ps      |
| samara_oblast                            | volga_federal_district           | samara_oblast                            | Ps      |
| samoa                                    | oceania                          | samoa                                    | Ps      |
| san_benito                               | california                       | san_benito                               | Ps      |
| san_bernardino                           | california                       | san_bernardino                           | Ps      |
| san_diego                                | california                       | san_diego                                | Ps      |
| san_francisco                            | california                       | san_francisco                            | Ps      |
| san_joaquin                              | california                       | san_joaquin                              | Ps      |
| san_juan                                 | argentina                        | san_juan                                 | Ps      |
| san_luis                                 | argentina                        | san_luis                                 | Ps      |
| san_luis_obispo                          | california                       | san_luis_obispo                          | Ps      |
| san_luis_potosi                          | mexico                           | san_luis_potosi                          | Ps      |
| san_marino                               | europe                           | san_marino                               | Ps      |
| san_mateo                                | california                       | san_mateo                                | Ps      |
| santa-catarina                           | south                            | santa-catarina                           | Ps      |
| santa_barbara                            | california                       | santa_barbara                            | Ps      |
| santa_clara                              | california                       | santa_clara                              | Ps      |
| santa_cruz_de_tenerife                   | canarias                         | santa_cruz_de_tenerife                   | Ps      |
| santa_fe                                 | argentina                        | santa_fe                                 | Ps      |
| santarem                                 | portugal                         | santarem                                 | Ps      |
| santiago                                 | chile                            | santiago                                 | Ps      |
| santiago_del_estero                      | argentina                        | santiago_del_estero                      | Ps      |
| sao-paulo                                | southeast                        | sao-paulo                                | Ps      |
| sao_tome_and_principe                    | africa                           | sao_tome_and_principe                    | Ps      |
| saone_et_loire                           | bourgogne                        | saone_et_loire                           | Ps      |
| saratov_oblast                           | volga_federal_district           | saratov_oblast                           | Ps      |
| sardegna                                 | italy                            | sardegna                                 | Ps      |
| sarthe                                   | pays_de_la_loire                 | sarthe                                   | Ps      |
| saskatchewan                             | canada                           | saskatchewan                             | Ps      |
| satakunta                                | finland                          | satakunta                                | Ps      |
| saudi_arabia                             | asia                             | saudi_arabia                             | Ps      |
| savoie                                   | rhone_alpes                      | savoie                                   | Ps      |
| schaffhausen                             | switzerland                      | schaffhausen                             | Ps      |
| schwyz                                   | switzerland                      | schwyz                                   | Ps      |
| seas                                     | europe                           | seas                                     |         |
| segovia                                  | castilla_y_leon                  | segovia                                  | Ps      |
| seine_et_marne                           | ile_de_france                    | seine_et_marne                           | Ps      |
| seine_maritime                           | haute_normandie                  | seine_maritime                           | Ps      |
| seine_saint_denis                        | ile_de_france                    | seine_saint_denis                        | Ps      |
| senegal                                  | africa                           | senegal                                  | Ps      |
| sergipe                                  | northeast                        | sergipe                                  | Ps      |
| setubal                                  | portugal                         | setubal                                  | Ps      |
| sevilla                                  | andalucia                        | sevilla                                  | Ps      |
| seychelles                               | africa                           | seychelles                               | Ps      |
| shaanxi                                  | china                            | shaanxi                                  | Ps      |
| shandong                                 | china                            | shandong                                 | Ps      |
| shanghai                                 | china                            | shanghai                                 | Ps      |
| shanxi                                   | china                            | shanxi                                   | Ps      |
| shasta                                   | california                       | shasta                                   | Ps      |
| shikoku                                  | japan                            | shikoku                                  | Ps      |
| siberian_federal_district                | russia                           | siberian_federal_district                | Ps      |
| sichuan                                  | china                            | sichuan                                  | Ps      |
| sicilia                                  | italy                            | sicilia                                  | Ps      |
| sierra                                   | california                       | sierra                                   | Ps      |
| sikkim                                   | india                            | sikkim                                   | Ps      |
| sinaloa                                  | mexico                           | sinaloa                                  | Ps      |
| singapore                                | asia                             | singapore                                | Ps      |
| sint_maarten                             | central-america                  | sint_maarten                             | Ps      |
| siskiyou                                 | california                       | siskiyou                                 | Ps      |
| sjaelland                                | denmark                          | sjaelland                                | Ps      |
| skane                                    | sweden                           | skane                                    | Ps      |
| slaskie                                  | poland                           | slaskie                                  | Ps      |
| slovakia                                 | europe                           | slovakia                                 | Ps      |
| smolensk_oblast                          | central_federal_district         | smolensk_oblast                          | Ps      |
| soccsksargen                             | philippines                      | soccsksargen                             | Ps      |
| sodermanland                             | sweden                           | sodermanland                             | Ps      |
| sofala                                   | mozambique                       | sofala                                   | Ps      |
| sogn_og_fjordane                         | norway                           | sogn_og_fjordane                         | Ps      |
| solano                                   | california                       | solano                                   | Ps      |
| solomon_islands                          | oceania                          | solomon_islands                          | Ps      |
| solothurn                                | switzerland                      | solothurn                                | Ps      |
| somme                                    | picardie                         | somme                                    | Ps      |
| sonoma                                   | california                       | sonoma                                   | Ps      |
| sonora                                   | mexico                           | sonora                                   | Ps      |
| soria                                    | castilla_y_leon                  | soria                                    | Ps      |
| south                                    | brazil                           | south                                    |         |
| south-america                            |                                  | south-america                            | Ps      |
| south_africa                             | africa                           | south_africa                             | Ps      |
| south_africa_north_west                  | south_africa                     | south_africa_north_west                  | Ps      |
| south_australia                          | australia                        | south_australia                          | Ps      |
| south_east_region                        | new_south_wales                  | south_east_region                        | Ps      |
| south_georgia_and_south_sandwich         | south-america                    | south_georgia_and_south_sandwich         | Ps      |
| south_kalimantan                         | indonesia                        | south_kalimantan                         | Ps      |
| south_karelia                            | finland                          | south_karelia                            | Ps      |
| south_ostrobothnia                       | finland                          | south_ostrobothnia                       | Ps      |
| south_papua                              | indonesia                        | south_papua                              | Ps      |
| south_savo                               | finland                          | south_savo                               | Ps      |
| south_sudan                              | africa                           | south_sudan                              | Ps      |
| south_sulawesi                           | indonesia                        | south_sulawesi                           | Ps      |
| south_sumatra                            | indonesia                        | south_sumatra                            | Ps      |
| southeast                                | brazil                           | southeast                                |         |
| southeast_sulawesi                       | indonesia                        | southeast_sulawesi                       | Ps      |
| southeastern_anatolia                    | turkey                           | southeastern_anatolia                    | Ps      |
| southern_federal_district                | russia                           | southern_federal_district                | Ps      |
| southern_federal_district_sevastopol     | southern_federal_district        | southern_federal_district_sevastopol     | Ps      |
| southern_highlands                       | tanzania                         | southern_highlands                       | Ps      |
| southern_tier                            | new-york                         | southern_tier                            | Ps      |
| southwest_finland                        | finland                          | southwest_finland                        | Ps      |
| southwest_papua                          | indonesia                        | southwest_papua                          | Ps      |
| southwestern_ontario                     | ontario                          | southwestern_ontario                     | Ps      |
| spain                                    | europe                           | spain                                    | Ps      |
| spain_la_rioja                           | spain                            | spain_la_rioja                           | Ps      |
| stanislaus                               | california                       | stanislaus                               | Ps      |
| state_of_mexico                          | mexico                           | state_of_mexico                          | Ps      |
| stavropol_krai                           | north_caucasian_federal_district | stavropol_krai                           | Ps      |
| steiermark                               | austria                          | steiermark                               | Ps      |
| stockholm                                | sweden                           | stockholm                                | Ps      |
| stredocesky                              | czech_republic                   | stredocesky                              | Ps      |
| sudan                                    | africa                           | sudan                                    | Ps      |
| sumy_oblast                              | ukraine                          | sumy_oblast                              | Ps      |
| suncoast                                 | florida                          | suncoast                                 | Ps      |
| suriname                                 | south-america                    | suriname                                 | Ps      |
| sutter                                   | california                       | sutter                                   | Ps      |
| svalbard                                 | norway                           | svalbard                                 | Ps      |
| sverdlovsk_oblast                        | ural_federal_district            | sverdlovsk_oblast                        | Ps      |
| swaziland                                | africa                           | swaziland                                | Ps      |
| sweden                                   | europe                           | sweden                                   | Ps      |
| swietokrzyskie                           | poland                           | swietokrzyskie                           | Ps      |
| switzerland                              | europe                           | switzerland                              | Ps      |
| switzerland_jura                         | switzerland                      | switzerland_jura                         | Ps      |
| syddanmark                               | denmark                          | syddanmark                               | Ps      |
| sydney_surrounds                         | new_south_wales                  | sydney_surrounds                         | Ps      |
| tabasco                                  | mexico                           | tabasco                                  | Ps      |
| tamaulipas                               | mexico                           | tamaulipas                               | Ps      |
| tambov_oblast                            | central_federal_district         | tambov_oblast                            | Ps      |
| tamil_nadu                               | india                            | tamil_nadu                               | Ps      |
| tanzania                                 | africa                           | tanzania                                 | Ps      |
| tanzania_central                         | tanzania                         | tanzania_central                         | Ps      |
| tanzania_lake                            | tanzania                         | tanzania_lake                            | Ps      |
| tanzania_northern                        | tanzania                         | tanzania_northern                        | Ps      |
| tanzania_western                         | tanzania                         | tanzania_western                         | Ps      |
| tarapaca                                 | chile                            | tarapaca                                 | Ps      |
| tarn                                     | midi_pyrenees                    | tarn                                     | Ps      |
| tarn_et_garonne                          | midi_pyrenees                    | tarn_et_garonne                          | Ps      |
| tarragona                                | catalunya                        | tarragona                                | Ps      |
| tasmania                                 | australia                        | tasmania                                 | Ps      |
| tatarstan_republic                       | volga_federal_district           | tatarstan_republic                       | Ps      |
| tehama                                   | california                       | tehama                                   | Ps      |
| telangana                                | india                            | telangana                                | Ps      |
| telemark                                 | norway                           | telemark                                 | Ps      |
| ternopil_oblast                          | ukraine                          | ternopil_oblast                          | Ps      |
| territoire_de_belfort                    | franche_comte                    | territoire_de_belfort                    | Ps      |
| teruel                                   | aragon                           | teruel                                   | Ps      |
| tete                                     | mozambique                       | tete                                     | Ps      |
| texas                                    | us-south                         | texas                                    | Ps      |
| texas_central                            | texas                            | texas_central                            | Ps      |
| texas_north                              | texas                            | texas_north                              | Ps      |
| texas_northwest                          | texas                            | texas_northwest                          | Ps      |
| texas_south                              | texas                            | texas_south                              | Ps      |
| texas_southeast                          | texas                            | texas_southeast                          | Ps      |
| texas_west                               | texas                            | texas_west                               | Ps      |
| the_riverina                             | new_south_wales                  | the_riverina                             | Ps      |
| thurgau                                  | switzerland                      | thurgau                                  | Ps      |
| tianjin                                  | china                            | tianjin                                  | Ps      |
| tibet                                    | china                            | tibet                                    | Ps      |
| ticino                                   | switzerland                      | ticino                                   | Ps      |
| tierra_del_fuego                         | argentina                        | tierra_del_fuego                         | Ps      |
| tirol                                    | austria                          | tirol                                    | Ps      |
| tlaxcala                                 | mexico                           | tlaxcala                                 | Ps      |
| tocantins                                | north                            | tocantins                                | Ps      |
| togo                                     | africa                           | togo                                     | Ps      |
| tohoku                                   | japan                            | tohoku                                   | Ps      |
| tokelau                                  | oceania                          | tokelau                                  | Ps      |
| toledo                                   | castilla_la_mancha               | toledo                                   | Ps      |
| tomsk_oblast                             | siberian_federal_district        | tomsk_oblast                             | Ps      |
| tonga                                    | oceania                          | tonga                                    | Ps      |
| toscana                                  | italy                            | toscana                                  | Ps      |
| treasure_coast                           | florida                          | treasure_coast                           | Ps      |
| trenciansky                              | slovakia                         | trenciansky                              | Ps      |
| trentino_alto_adige                      | italy                            | trentino_alto_adige                      | Ps      |
| trinidad_and_tobago                      | central-america                  | trinidad_and_tobago                      | Ps      |
| trinity                                  | california                       | trinity                                  | Ps      |
| tripura                                  | india                            | tripura                                  | Ps      |
| trnavsky                                 | slovakia                         | trnavsky                                 | Ps      |
| troendelag                               | norway                           | troendelag                               | Ps      |
| troms                                    | norway                           | troms                                    | Ps      |
| tucuman                                  | argentina                        | tucuman                                  | Ps      |
| tula_oblast                              | central_federal_district         | tula_oblast                              | Ps      |
| tulare                                   | california                       | tulare                                   | Ps      |
| tunisia                                  | africa                           | tunisia                                  | Ps      |
| tuolumne                                 | california                       | tuolumne                                 | Ps      |
| turkey                                   | europe                           | turkey                                   | Ps      |
| turks_and_caicos_islands                 | central-america                  | turks_and_caicos_islands                 | Ps      |
| tuva_republic                            | siberian_federal_district        | tuva_republic                            | Ps      |
| tuvalu                                   | oceania                          | tuvalu                                   | Ps      |
| tver_oblast                              | central_federal_district         | tver_oblast                              | Ps      |
| tyumen_oblast                            | ural_federal_district            | tyumen_oblast                            | Ps      |
| udmurt_republic                          | volga_federal_district           | udmurt_republic                          | Ps      |
| uganda                                   | africa                           | uganda                                   | Ps      |
| uganda_central                           | uganda                           | uganda_central                           | Ps      |
| uganda_eastern                           | uganda                           | uganda_eastern                           | Ps      |
| uganda_northern                          | uganda                           | uganda_northern                          | Ps      |
| uganda_western                           | uganda                           | uganda_western                           | Ps      |
| ukraine                                  | europe                           | ukraine                                  | Ps      |
| ukraine_sevastopol                       | ukraine                          | ukraine_sevastopol                       | Ps      |
| ulyanovsk_oblast                         | volga_federal_district           | ulyanovsk_oblast                         | Ps      |
| umbria                                   | italy                            | umbria                                   | Ps      |
| united_arab_emirates                     | asia                             | united_arab_emirates                     | Ps      |
| united_kingdom                           | europe                           | united_kingdom                           | Ps      |
| united_states_virgin_islands             | central-america                  | united_states_virgin_islands             | Ps      |
| upper_peninsula                          | michigan                         | upper_peninsula                          | Ps      |
| uppsala                                  | sweden                           | uppsala                                  | Ps      |
| ural_federal_district                    | russia                           | ural_federal_district                    | Ps      |
| uri                                      | switzerland                      | uri                                      | Ps      |
| us-midwest                               | north-america                    | us-midwest                               | Ps      |
| us-northeast                             | north-america                    | us-northeast                             | Ps      |
| us-south                                 | north-america                    | us-south                                 | Ps      |
| us-west                                  | north-america                    | us-west                                  | Ps      |
| usa_virgin_islands                       | central-america                  | usa_virgin_islands                       | Ps      |
| ustecky                                  | czech_republic                   | ustecky                                  | Ps      |
| utrecht                                  | netherlands                      | utrecht                                  | Ps      |
| uttar_pradesh                            | india                            | uttar_pradesh                            | Ps      |
| uttarakhand                              | india                            | uttarakhand                              | Ps      |
| uusimaa                                  | finland                          | uusimaa                                  | Ps      |
| val_d_oise                               | ile_de_france                    | val_d_oise                               | Ps      |
| val_de_marne                             | ile_de_france                    | val_de_marne                             | Ps      |
| valais                                   | switzerland                      | valais                                   | Ps      |
| valencia                                 | comunitat_valenciana             | valencia                                 | Ps      |
| valladolid                               | castilla_y_leon                  | valladolid                               | Ps      |
| valle_aosta                              | italy                            | valle_aosta                              | Ps      |
| valparaiso                               | chile                            | valparaiso                               | Ps      |
| vanuatu                                  | oceania                          | vanuatu                                  | Ps      |
| var                                      | provence_alpes_cote_d_azur       | var                                      | Ps      |
| varmland                                 | sweden                           | varmland                                 | Ps      |
| vasterbotten                             | sweden                           | vasterbotten                             | Ps      |
| vasternorrland                           | sweden                           | vasternorrland                           | Ps      |
| vastmanland                              | sweden                           | vastmanland                              | Ps      |
| vastra_gotaland                          | sweden                           | vastra_gotaland                          | Ps      |
| vatican_city                             | europe                           | vatican_city                             | Ps      |
| vaucluse                                 | provence_alpes_cote_d_azur       | vaucluse                                 | Ps      |
| vaud                                     | switzerland                      | vaud                                     | Ps      |
| vendee                                   | pays_de_la_loire                 | vendee                                   | Ps      |
| veneto                                   | italy                            | veneto                                   | Ps      |
| venezuela                                | south-america                    | venezuela                                | Ps      |
| ventura                                  | california                       | ventura                                  | Ps      |
| veracruz                                 | mexico                           | veracruz                                 | Ps      |
| vest-agder                               | norway                           | vest-agder                               | Ps      |
| vestfold                                 | norway                           | vestfold                                 | Ps      |
| viana_do_castelo                         | portugal                         | viana_do_castelo                         | Ps      |
| victoria                                 | australia                        | victoria                                 | Ps      |
| vienne                                   | poitou_charentes                 | vienne                                   | Ps      |
| vila_real                                | portugal                         | vila_real                                | Ps      |
| vinnytsia_oblast                         | ukraine                          | vinnytsia_oblast                         | Ps      |
| virginia                                 | us-south                         | virginia                                 | Ps      |
| viseu                                    | portugal                         | viseu                                    | Ps      |
| vizcaya                                  | euskadi                          | vizcaya                                  | Ps      |
| vladimir_oblast                          | central_federal_district         | vladimir_oblast                          | Ps      |
| volga_federal_district                   | russia                           | volga_federal_district                   | Ps      |
| volgograd_oblast                         | southern_federal_district        | volgograd_oblast                         | Ps      |
| vologda_oblast                           | northwestern_federal_district    | vologda_oblast                           | Ps      |
| volyn_oblast                             | ukraine                          | volyn_oblast                             | Ps      |
| vorarlberg                               | austria                          | vorarlberg                               | Ps      |
| voronezh_oblast                          | central_federal_district         | voronezh_oblast                          | Ps      |
| vosges                                   | lorraine                         | vosges                                   | Ps      |
| vysocina                                 | czech_republic                   | vysocina                                 | Ps      |
| wallis_et_futuna                         | oceania                          | wallis_et_futuna                         | Ps      |
| wallonia_french_community                | belgium                          | wallonia_french_community                | Ps      |
| wallonia_german_community                | belgium                          | wallonia_german_community                | Ps      |
| warminsko_mazurskie                      | poland                           | warminsko_mazurskie                      | Ps      |
| west_bengal                              | india                            | west_bengal                              | Ps      |
| west_flanders                            | flanders                         | west_flanders                            | Ps      |
| west_java                                | indonesia                        | west_java                                | Ps      |
| west_kalimantan                          | indonesia                        | west_kalimantan                          | Ps      |
| west_midlands                            | england                          | west_midlands                            | Ps      |
| west_nusa_tenggara                       | indonesia                        | west_nusa_tenggara                       | Ps      |
| west_papua                               | indonesia                        | west_papua                               | Ps      |
| west_sulawesi                            | indonesia                        | west_sulawesi                            | Ps      |
| west_sumatra                             | indonesia                        | west_sumatra                             | Ps      |
| western_australia                        | australia                        | western_australia                        | Ps      |
| western_cape                             | south_africa                     | western_cape                             | Ps      |
| western_new_york                         | new-york                         | western_new_york                         | Ps      |
| western_sahara                           | africa                           | western_sahara                           | Ps      |
| western_visayas                          | philippines                      | western_visayas                          | Ps      |
| wielkopolskie                            | poland                           | wielkopolskie                            | Ps      |
| wien                                     | austria                          | wien                                     | Ps      |
| wytheville                               | virginia                         | wytheville                               | Ps      |
| xinjiang                                 | china                            | xinjiang                                 | Ps      |
| yamalo_nenets_autonomous_okrug           | ural_federal_district            | yamalo_nenets_autonomous_okrug           | Ps      |
| yaroslavl_oblast                         | central_federal_district         | yaroslavl_oblast                         | Ps      |
| yogyakarta                               | indonesia                        | yogyakarta                               | Ps      |
| yolo                                     | california                       | yolo                                     | Ps      |
| yonne                                    | bourgogne                        | yonne                                    | Ps      |
| yorkshire_and_the_humber                 | england                          | yorkshire_and_the_humber                 | Ps      |
| yuba                                     | california                       | yuba                                     | Ps      |
| yucatan                                  | mexico                           | yucatan                                  | Ps      |
| yukon                                    | canada                           | yukon                                    | Ps      |
| yunnan                                   | china                            | yunnan                                   | Ps      |
| yvelines                                 | ile_de_france                    | yvelines                                 | Ps      |
| zabaykalsky_krai                         | siberian_federal_district        | zabaykalsky_krai                         | Ps      |
| zacatecas                                | mexico                           | zacatecas                                | Ps      |
| zachodniopomorskie                       | poland                           | zachodniopomorskie                       | Ps      |
| zakarpattia_oblast                       | ukraine                          | zakarpattia_oblast                       | Ps      |
| zambezia                                 | mozambique                       | zambezia                                 | Ps      |
| zambia                                   | africa                           | zambia                                   | Ps      |
| zamboanga_peninsula                      | philippines                      | zamboanga_peninsula                      | Ps      |
| zamora                                   | castilla_y_leon                  | zamora                                   | Ps      |
| zanzibar                                 | tanzania                         | zanzibar                                 | Ps      |
| zaporizhia_oblast                        | ukraine                          | zaporizhia_oblast                        | Ps      |
| zaragoza                                 | aragon                           | zaragoza                                 | Ps      |
| zeeland                                  | netherlands                      | zeeland                                  | Ps      |
| zhejiang                                 | china                            | zhejiang                                 | Ps      |
| zhytomyr_oblast                          | ukraine                          | zhytomyr_oblast                          | Ps      |
| zilinsky                                 | slovakia                         | zilinsky                                 | Ps      |
| zimbabwe                                 | africa                           | zimbabwe                                 | Ps      |
| zlinsky                                  | czech_republic                   | zlinsky                                  | Ps      |
| zug                                      | switzerland                      | zug                                      | Ps      |
| zuid_holland                             | netherlands                      | zuid_holland                             | Ps      |
| zurich                                   | switzerland                      | zurich                                   | Ps      |

## List of elements from bbbike.org
| SHORT NAME       | IS IN | LONG NAME        | FORMATS |
|------------------|-------|------------------|---------|
| Aachen           |       | Aachen           | gGPpS   |
| Aarhus           |       | Aarhus           | gGPpS   |
| Adelaide         |       | Adelaide         | gGPpS   |
| Albuquerque      |       | Albuquerque      | gGPpS   |
| Alexandria       |       | Alexandria       | gGPpS   |
| Amsterdam        |       | Amsterdam        | gGPpS   |
| Antwerpen        |       | Antwerpen        | gGPpS   |
| Arnhem           |       | Arnhem           | gGPpS   |
| Auckland         |       | Auckland         | gGPpS   |
| Augsburg         |       | Augsburg         | gGPpS   |
| Austin           |       | Austin           | gGPpS   |
| Baghdad          |       | Baghdad          | gGPpS   |
| Baku             |       | Baku             | gGPpS   |
| Balaton          |       | Balaton          | gGPpS   |
| Bamberg          |       | Bamberg          | gGPpS   |
| Bangkok          |       | Bangkok          | gGPpS   |
| Barcelona        |       | Barcelona        | gGPpS   |
| Basel            |       | Basel            | gGPpS   |
| Beijing          |       | Beijing          | gGPpS   |
| Beirut           |       | Beirut           | gGPpS   |
| Berkeley         |       | Berkeley         | gGPpS   |
| Berlin           |       | Berlin           | gGPpS   |
| Bern             |       | Bern             | gGPpS   |
| Bielefeld        |       | Bielefeld        | gGPpS   |
| Birmingham       |       | Birmingham       | gGPpS   |
| Bochum           |       | Bochum           | gGPpS   |
| Bogota           |       | Bogota           | gGPpS   |
| Bombay           |       | Bombay           | gGPpS   |
| Bonn             |       | Bonn             | gGPpS   |
| Bordeaux         |       | Bordeaux         | gGPpS   |
| Boulder          |       | Boulder          | gGPpS   |
| BrandenburgHavel |       | BrandenburgHavel | gGPpS   |
| Braunschweig     |       | Braunschweig     | gGPpS   |
| Bremen           |       | Bremen           | gGPpS   |
| Bremerhaven      |       | Bremerhaven      | gGPpS   |
| Brisbane         |       | Brisbane         | gGPpS   |
| Bristol          |       | Bristol          | gGPpS   |
| Brno             |       | Brno             | gGPpS   |
| Bruegge          |       | Bruegge          | gGPpS   |
| Bruessel         |       | Bruessel         | gGPpS   |
| Budapest         |       | Budapest         | gGPpS   |
| BuenosAires      |       | BuenosAires      | gGPpS   |
| Cairo            |       | Cairo            | gGPpS   |
| Calgary          |       | Calgary          | gGPpS   |
| Cambridge        |       | Cambridge        | gGPpS   |
| CambridgeMa      |       | CambridgeMa      | gGPpS   |
| Canberra         |       | Canberra         | gGPpS   |
| CapeTown         |       | CapeTown         | gGPpS   |
| Chemnitz         |       | Chemnitz         | gGPpS   |
| Chicago          |       | Chicago          | gGPpS   |
| ClermontFerrand  |       | ClermontFerrand  | gGPpS   |
| Colmar           |       | Colmar           | gGPpS   |
| Copenhagen       |       | Copenhagen       | gGPpS   |
| Cork             |       | Cork             | gGPpS   |
| Corsica          |       | Corsica          | gGPpS   |
| Corvallis        |       | Corvallis        | gGPpS   |
| Cottbus          |       | Cottbus          | gGPpS   |
| Cracow           |       | Cracow           | gGPpS   |
| CraterLake       |       | CraterLake       | gGPpS   |
| Curitiba         |       | Curitiba         | gGPpS   |
| Cusco            |       | Cusco            | gGPpS   |
| Dallas           |       | Dallas           | gGPpS   |
| Damaskus         |       | Damaskus         | gGPpS   |
| Darmstadt        |       | Darmstadt        | gGPpS   |
| Davis            |       | Davis            | gGPpS   |
| DenHaag          |       | DenHaag          | gGPpS   |
| Denver           |       | Denver           | gGPpS   |
| Dessau           |       | Dessau           | gGPpS   |
| Dortmund         |       | Dortmund         | gGPpS   |
| Dresden          |       | Dresden          | gGPpS   |
| Dublin           |       | Dublin           | gGPpS   |
| Duesseldorf      |       | Duesseldorf      | gGPpS   |
| Duisburg         |       | Duisburg         | gGPpS   |
| Edinburgh        |       | Edinburgh        | gGPpS   |
| Eindhoven        |       | Eindhoven        | gGPpS   |
| Emden            |       | Emden            | gGPpS   |
| Erfurt           |       | Erfurt           | gGPpS   |
| Erlangen         |       | Erlangen         | gGPpS   |
| Eugene           |       | Eugene           | gGPpS   |
| Flensburg        |       | Flensburg        | gGPpS   |
| FortCollins      |       | FortCollins      | gGPpS   |
| Frankfurt        |       | Frankfurt        | gGPpS   |
| FrankfurtOder    |       | FrankfurtOder    | gGPpS   |
| Freiburg         |       | Freiburg         | gGPpS   |
| Gdansk           |       | Gdansk           | gGPpS   |
| Genf             |       | Genf             | gGPpS   |
| Gent             |       | Gent             | gGPpS   |
| Gera             |       | Gera             | gGPpS   |
| Glasgow          |       | Glasgow          | gGPpS   |
| Gliwice          |       | Gliwice          | gGPpS   |
| Goerlitz         |       | Goerlitz         | gGPpS   |
| Goeteborg        |       | Goeteborg        | gGPpS   |
| Goettingen       |       | Goettingen       | gGPpS   |
| Graz             |       | Graz             | gGPpS   |
| Groningen        |       | Groningen        | gGPpS   |
| Halifax          |       | Halifax          | gGPpS   |
| Halle            |       | Halle            | gGPpS   |
| Hamburg          |       | Hamburg          | gGPpS   |
| Hamm             |       | Hamm             | gGPpS   |
| Hannover         |       | Hannover         | gGPpS   |
| Heilbronn        |       | Heilbronn        | gGPpS   |
| Helsinki         |       | Helsinki         | gGPpS   |
| Hertogenbosch    |       | Hertogenbosch    | gGPpS   |
| Huntsville       |       | Huntsville       | gGPpS   |
| Innsbruck        |       | Innsbruck        | gGPpS   |
| Istanbul         |       | Istanbul         | gGPpS   |
| Jena             |       | Jena             | gGPpS   |
| Jerusalem        |       | Jerusalem        | gGPpS   |
| Johannesburg     |       | Johannesburg     | gGPpS   |
| Kaiserslautern   |       | Kaiserslautern   | gGPpS   |
| Karlsruhe        |       | Karlsruhe        | gGPpS   |
| Kassel           |       | Kassel           | gGPpS   |
| Katowice         |       | Katowice         | gGPpS   |
| Kaunas           |       | Kaunas           | gGPpS   |
| Kiel             |       | Kiel             | gGPpS   |
| Kiew             |       | Kiew             | gGPpS   |
| Koblenz          |       | Koblenz          | gGPpS   |
| Koeln            |       | Koeln            | gGPpS   |
| Konstanz         |       | Konstanz         | gGPpS   |
| LaPaz            |       | LaPaz            | gGPpS   |
| LaPlata          |       | LaPlata          | gGPpS   |
| LakeGarda        |       | LakeGarda        | gGPpS   |
| Lausanne         |       | Lausanne         | gGPpS   |
| Leeds            |       | Leeds            | gGPpS   |
| Leipzig          |       | Leipzig          | gGPpS   |
| Lima             |       | Lima             | gGPpS   |
| Linz             |       | Linz             | gGPpS   |
| Lisbon           |       | Lisbon           | gGPpS   |
| Liverpool        |       | Liverpool        | gGPpS   |
| Ljubljana        |       | Ljubljana        | gGPpS   |
| Lodz             |       | Lodz             | gGPpS   |
| London           |       | London           | gGPpS   |
| LosAngeles       |       | LosAngeles       | gGPpS   |
| Luebeck          |       | Luebeck          | gGPpS   |
| Luxemburg        |       | Luxemburg        | gGPpS   |
| Lyon             |       | Lyon             | gGPpS   |
| Maastricht       |       | Maastricht       | gGPpS   |
| Madison          |       | Madison          | gGPpS   |
| Madrid           |       | Madrid           | gGPpS   |
| Magdeburg        |       | Magdeburg        | gGPpS   |
| Mainz            |       | Mainz            | gGPpS   |
| Malmoe           |       | Malmoe           | gGPpS   |
| Manchester       |       | Manchester       | gGPpS   |
| Mannheim         |       | Mannheim         | gGPpS   |
| Marseille        |       | Marseille        | gGPpS   |
| Melbourne        |       | Melbourne        | gGPpS   |
| Memphis          |       | Memphis          | gGPpS   |
| MexicoCity       |       | MexicoCity       | gGPpS   |
| Miami            |       | Miami            | gGPpS   |
| Minsk            |       | Minsk            | gGPpS   |
| Moenchengladbach |       | Moenchengladbach | gGPpS   |
| Montevideo       |       | Montevideo       | gGPpS   |
| Montpellier      |       | Montpellier      | gGPpS   |
| Montreal         |       | Montreal         | gGPpS   |
| Moscow           |       | Moscow           | gGPpS   |
| Muenchen         |       | Muenchen         | gGPpS   |
| Muenster         |       | Muenster         | gGPpS   |
| NewDelhi         |       | NewDelhi         | gGPpS   |
| NewOrleans       |       | NewOrleans       | gGPpS   |
| NewYork          |       | NewYork          | gGPpS   |
| Nuernberg        |       | Nuernberg        | gGPpS   |
| Oldenburg        |       | Oldenburg        | gGPpS   |
| Oranienburg      |       | Oranienburg      | gGPpS   |
| Orlando          |       | Orlando          | gGPpS   |
| Oslo             |       | Oslo             | gGPpS   |
| Osnabrueck       |       | Osnabrueck       | gGPpS   |
| Ostrava          |       | Ostrava          | gGPpS   |
| Ottawa           |       | Ottawa           | gGPpS   |
| Paderborn        |       | Paderborn        | gGPpS   |
| Palma            |       | Palma            | gGPpS   |
| PaloAlto         |       | PaloAlto         | gGPpS   |
| Paris            |       | Paris            | gGPpS   |
| Perth            |       | Perth            | gGPpS   |
| Philadelphia     |       | Philadelphia     | gGPpS   |
| PhnomPenh        |       | PhnomPenh        | gGPpS   |
| Portland         |       | Portland         | gGPpS   |
| PortlandME       |       | PortlandME       | gGPpS   |
| Porto            |       | Porto            | gGPpS   |
| PortoAlegre      |       | PortoAlegre      | gGPpS   |
| Potsdam          |       | Potsdam          | gGPpS   |
| Poznan           |       | Poznan           | gGPpS   |
| Prag             |       | Prag             | gGPpS   |
| Providence       |       | Providence       | gGPpS   |
| Regensburg       |       | Regensburg       | gGPpS   |
| Riga             |       | Riga             | gGPpS   |
| RiodeJaneiro     |       | RiodeJaneiro     | gGPpS   |
| Rostock          |       | Rostock          | gGPpS   |
| Rotterdam        |       | Rotterdam        | gGPpS   |
| Ruegen           |       | Ruegen           | gGPpS   |
| Saarbruecken     |       | Saarbruecken     | gGPpS   |
| Sacramento       |       | Sacramento       | gGPpS   |
| Saigon           |       | Saigon           | gGPpS   |
| Salzburg         |       | Salzburg         | gGPpS   |
| SanFrancisco     |       | SanFrancisco     | gGPpS   |
| SanJose          |       | SanJose          | gGPpS   |
| SanktPetersburg  |       | SanktPetersburg  | gGPpS   |
| SantaBarbara     |       | SantaBarbara     | gGPpS   |
| SantaCruz        |       | SantaCruz        | gGPpS   |
| Santiago         |       | Santiago         | gGPpS   |
| Sarajewo         |       | Sarajewo         | gGPpS   |
| Schwerin         |       | Schwerin         | gGPpS   |
| Seattle          |       | Seattle          | gGPpS   |
| Seoul            |       | Seoul            | gGPpS   |
| Sheffield        |       | Sheffield        | gGPpS   |
| Singapore        |       | Singapore        | gGPpS   |
| Sofia            |       | Sofia            | gGPpS   |
| Stockholm        |       | Stockholm        | gGPpS   |
| Stockton         |       | Stockton         | gGPpS   |
| Strassburg       |       | Strassburg       | gGPpS   |
| Stuttgart        |       | Stuttgart        | gGPpS   |
| Sucre            |       | Sucre            | gGPpS   |
| Sydney           |       | Sydney           | gGPpS   |
| Szczecin         |       | Szczecin         | gGPpS   |
| Tallinn          |       | Tallinn          | gGPpS   |
| Tehran           |       | Tehran           | gGPpS   |
| Tilburg          |       | Tilburg          | gGPpS   |
| Tokyo            |       | Tokyo            | gGPpS   |
| Toronto          |       | Toronto          | gGPpS   |
| Toulouse         |       | Toulouse         | gGPpS   |
| Trondheim        |       | Trondheim        | gGPpS   |
| Tucson           |       | Tucson           | gGPpS   |
| Turin            |       | Turin            | gGPpS   |
| UlanBator        |       | UlanBator        | gGPpS   |
| Ulm              |       | Ulm              | gGPpS   |
| Usedom           |       | Usedom           | gGPpS   |
| Utrecht          |       | Utrecht          | gGPpS   |
| Vancouver        |       | Vancouver        | gGPpS   |
| Victoria         |       | Victoria         | gGPpS   |
| WarenMueritz     |       | WarenMueritz     | gGPpS   |
| Warsaw           |       | Warsaw           | gGPpS   |
| WashingtonDC     |       | WashingtonDC     | gGPpS   |
| Waterloo         |       | Waterloo         | gGPpS   |
| Wien             |       | Wien             | gGPpS   |
| Wroclaw          |       | Wroclaw          | gGPpS   |
| Wuerzburg        |       | Wuerzburg        | gGPpS   |
| Wuppertal        |       | Wuppertal        | gGPpS   |
| Zagreb           |       | Zagreb           | gGPpS   |
| Zuerich          |       | Zuerich          | gGPpS   |

## List of elements from geo2day
| SHORT NAME        | IS IN | LONG NAME                | FORMATS |
|-------------------|-------|--------------------------|---------|
| africa            |       | [africa.poly]            | Pp      |
| antarctica        |       | [antarctica.poly]        | Pp      |
| asia              |       | [asia.poly]              | Pp      |
| australia_oceania |       | [australia_oceania.poly] | Pp      |
| central_america   |       | [central_america.poly]   | Pp      |
| europe            |       | [europe.poly]            | Pp      |
| north_america     |       | [north_america.poly]     | Pp      |
| russia            |       | [russia.poly]            | Pp      |
| south_america     |       | [south_america.poly]     | Pp      |

