package main

import (
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	yaml "gopkg.in/yaml.v2"
)

//Sample data:
const osmfrValidHTML1 = `
	<h1>Index of /extracts</h1>
	<table><tbody><tr><th><img src="/icons/blank.gif" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr><tr><th colspan="5"><hr></th></tr>
	<tr><td valign="top"><img src="/icons/back.gif" alt="[DIR]"></td><td><a href="/">Parent Directory</a></td><td>&nbsp;</td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="africa-latest.osm.pbf">africa-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:43  </td><td align="right">2.3G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="africa.osm.pbf">africa.osm.pbf</a></td><td align="right">22-Mar-2018 02:43  </td><td align="right">2.3G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="africa.state.txt">africa.state.txt</a></td><td align="right">22-Mar-2018 02:25  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="africa/">africa/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="asia-latest.osm.pbf">asia-latest.osm.pbf</a></td><td align="right">22-Mar-2018 01:39  </td><td align="right">6.0G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="asia.osm.pbf">asia.osm.pbf</a></td><td align="right">22-Mar-2018 01:39  </td><td align="right">6.0G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="asia.state.txt">asia.state.txt</a></td><td align="right">22-Mar-2018 00:50  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="asia/">asia/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="central-america-latest.osm.pbf">central-america-latest.osm.pbf</a></td><td align="right">22-Mar-2018 00:52  </td><td align="right">336M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="central-america.osm.pbf">central-america.osm.pbf</a></td><td align="right">22-Mar-2018 00:52  </td><td align="right">336M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="central-america.state.txt">central-america.state.txt</a></td><td align="right">22-Mar-2018 00:49  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="central-america/">central-america/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="europe-latest.osm.pbf">europe-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:31  </td><td align="right"> 20G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="europe.osm.pbf">europe.osm.pbf</a></td><td align="right">22-Mar-2018 02:31  </td><td align="right"> 20G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="europe.state.txt">europe.state.txt</a></td><td align="right">22-Mar-2018 00:09  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="europe/">europe/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="merge/">merge/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="north-america/">north-america/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="oceania-latest.osm.pbf">oceania-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:42  </td><td align="right">570M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="oceania.osm.pbf">oceania.osm.pbf</a></td><td align="right">22-Mar-2018 02:42  </td><td align="right">570M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="oceania.state.txt">oceania.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="oceania/">oceania/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="south-america-latest.osm.pbf">south-america-latest.osm.pbf</a></td><td align="right">22-Mar-2018 01:00  </td><td align="right">1.3G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="south-america.osm.pbf">south-america.osm.pbf</a></td><td align="right">22-Mar-2018 01:00  </td><td align="right">1.3G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="south-america.state.txt">south-america.state.txt</a></td><td align="right">22-Mar-2018 00:49  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="south-america/">south-america/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><th colspan="5"><hr></th></tr>
	</tbody></table>
	<address>Apache/2.2.22 (Debian) Server at download.openstreetmap.fr Port 80</address>
	`

const osmfrValidHTML2 = `
	<h1>Index of /extracts/merge</h1>
	<table><tbody><tr><th><img src="/icons/blank.gif" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr><tr><th colspan="5"><hr></th></tr>
	<tr><td valign="top"><img src="/icons/back.gif" alt="[DIR]"></td><td><a href="/extracts/">Parent Directory</a></td><td>&nbsp;</td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="fiji-latest.osm.pbf">fiji-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">7.2M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="fiji.osm.pbf">fiji.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">7.2M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="fiji.state.txt">fiji.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_metro_dom_com_nc-latest.osm.pbf">france_metro_dom_com_nc-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:56  </td><td align="right">3.6G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_metro_dom_com_nc.osm.pbf">france_metro_dom_com_nc.osm.pbf</a></td><td align="right">22-Mar-2018 02:56  </td><td align="right">3.6G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="france_metro_dom_com_nc.state.txt">france_metro_dom_com_nc.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_taaf-latest.osm.pbf">france_taaf-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">1.0M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_taaf.osm.pbf">france_taaf.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">1.0M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="france_taaf.state.txt">france_taaf.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="israel_and_palestine-latest.osm.pbf">israel_and_palestine-latest.osm.pbf</a></td><td align="right">21-Mar-2017 02:37  </td><td align="right"> 60M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="israel_and_palestine.osm.pbf">israel_and_palestine.osm.pbf</a></td><td align="right">21-Mar-2017 02:37  </td><td align="right"> 60M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="israel_and_palestine.state.txt">israel_and_palestine.state.txt</a></td><td align="right">11-Sep-2017 13:52  </td><td align="right">190 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="kiribati-latest.osm.pbf">kiribati-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">1.1M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="kiribati.osm.pbf">kiribati.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">1.1M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="kiribati.state.txt">kiribati.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><th colspan="5"><hr></th></tr>
	</tbody></table>
	<address>Apache/2.2.22 (Debian) Server at download.openstreetmap.fr Port 80</address>
	`

const osmfrPolygonJapanValidHTML = `
	<h1>Index of /polygons/asia/japan</h1>
	<table><tbody><tr><th><img src="/icons/blank.gif" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr><tr><th colspan="5"><hr></th></tr>
	<tr><td valign="top"><img src="/icons/back.gif" alt="[DIR]"></td><td><a href="/polygons/asia/">Parent Directory</a></td><td>&nbsp;</td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="chubu.poly">chubu.poly</a></td><td align="right">05-Jun-2016 21:40  </td><td align="right">5.5K</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="chugoku.poly">chugoku.poly</a></td><td align="right">05-Jun-2016 21:40  </td><td align="right">3.5K</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="hokkaido.poly">hokkaido.poly</a></td><td align="right">05-Jun-2016 21:40  </td><td align="right">2.2K</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="kansai.poly">kansai.poly</a></td><td align="right">05-Jun-2016 22:47  </td><td align="right">3.1K</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="kanto.poly">kanto.poly</a></td><td align="right">05-Jun-2016 21:40  </td><td align="right">5.8K</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="kyushu.poly">kyushu.poly</a></td><td align="right">05-Jun-2016 21:40  </td><td align="right">5.9K</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="shikoku.poly">shikoku.poly</a></td><td align="right">05-Jun-2016 21:40  </td><td align="right">2.0K</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="tohoku.poly">tohoku.poly</a></td><td align="right">05-Jun-2016 21:40  </td><td align="right">3.2K</td><td>&nbsp;</td></tr>
	<tr><th colspan="5"><hr></th></tr>
	</tbody></table>
	<address>Apache/2.2.22 (Debian) Server at download.openstreetmap.fr Port 80</address>
	`

const geofabrikSouthAmericaHTML = `<div id="download-top">
<div id="download-top-right">
   <img src="/img/geofabrik-downloads.png" height="20">
</div>
</div>
<div id="download-main">
<div id="download-titlebar">
<h1>Download OpenStreetMap data for this region:</h1>
</div>
<h2>South America</h2><p><a href="index.html">[one level up]</a></p><div id="download-right">
<div id="noscroll">
<div id="map" class="olMap"><div id="OpenLayers.Map_2_OpenLayers_ViewPort" style="position: relative; overflow: hidden; width: 100%; height: 100%;" class="olMapViewport olControlDragPanActive olControlZoomBoxActive olControlPinchZoomActive olControlNavigationActive"><div id="OpenLayers.Map_2_OpenLayers_Container" style="position: absolute; width: 100px; height: 100px; z-index: 749; left: 0px; top: 0px;"><div id="OpenLayers.Layer.OSM_18" style="position: absolute; width: 100%; height: 100%; z-index: 100; left: 0%; top: 0%;" dir="ltr" class="olLayerDiv olLayerGrid"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 89%; top: 119%; width: 256%; height: 256%;" src="http://c.tile.openstreetmap.org/1/0/1.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 89%; top: -137%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/1/0/0.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: -167%; top: 119%; width: 256%; height: 256%;" src="http://c.tile.openstreetmap.org/1/1/1.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: -167%; top: -137%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/1/1/0.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 345%; top: 119%; width: 256%; height: 256%;" src="http://c.tile.openstreetmap.org/1/1/1.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 345%; top: -137%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/1/1/0.png"></div><div id="OpenLayers.Layer.Vector_25" style="position: absolute; width: 100%; height: 100%; z-index: 330; left: 0px; top: 0px;" dir="ltr" class="olLayerDiv"><svg id="OpenLayers.Layer.Vector_25_svgRoot" style="display: block;" width="400" height="300" viewBox="0 0 400 300"><g id="OpenLayers.Layer.Vector_25_root" style="visibility: visible;" transform=""><g id="OpenLayers.Layer.Vector_25_vroot"><path id="OpenLayers.Geometry.Polygon_60" d=" M 226.97429759999682,102.46657608783201 235.22762666666347,104.33386087744958 234.31447466666347,106.89489007355303 234.63228444444127,107.61385521051724 234.2780088888857,108.15568499019342 233.97582933333015,107.97854880928976 233.46994488888572,108.99916614281356 228.24178488888572,112.90719015171317 208.9314631111079,114.3573731316243 172.0356337777746,72.6566195681074 88.5291448888857,109.15460817598903 88.64989155555239,227.34338043189038 311.47085511110794,227.23363288549862 303.6760675555524,116.1457460664315 280.64504177777457,104.5352087514895 257.5038648888857,104.89979861214199 257.0681528888857,104.75563864636494 256.5545031111079,104.73148662414059 256.27094044444124,104.66188590579017 256.5629226666635,103.91162653924778 256.65462755555234,103.90999166523264 256.6543999999968,103.70826258960781 256.78646755555235,103.32938530013536 256.69954133333016,102.32810875532044 255.26294044444126,102.27579635460584 252.96565333333012,101.59480150565409 245.5126257777746,100.45962936427753 242.21998222221902,100.18086167624867 237.81449244444124,102.3203683428004 232.24999964444123,94.62500019179268 228.12500053333014,98.18750007675122 226.97429759999682,102.46657608783201 z" fill-rule="evenodd" fill="#ffcc66" fill-opacity="0.5" stroke="#ff9933" stroke-opacity="1" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path></g><g id="OpenLayers.Layer.Vector_25_troot"></g></g></svg></div></div><div id="OpenLayers.Control.Zoom_5" style="position: absolute; z-index: 1004;" class="olControlZoom olControlNoSelect" unselectable="on"><a href="#zoomIn" class="olControlZoomIn olButton">+</a><a href="#zoomOut" class="olControlZoomOut olButton">-</a></div><div id="OpenLayers.Control.ArgParser_6" style="position: absolute; z-index: 1004;" class="olControlArgParser olControlNoSelect" unselectable="on"></div><div id="OpenLayers.Control.Attribution_7" style="position: absolute; z-index: 1004;" class="olControlAttribution olControlNoSelect" unselectable="on"></div></div></div>

<div class="customized_shapes">
<p>
   <img src="/img/en.png">
Not what you were looking for?
Geofabrik is a consulting and software development firm based in Karlsruhe, Germany
specializing in OpenStreetMap services. We're happy to help you with data
preparation, processing, server setup and the like. <a href="http://www.geofabrik.de/data/shapefiles.html">Check out our web site</a>
and contact us if we can be of service.

   </p>
</div>

<div class="customized_shapes">
   <p>
   <img src="/img/de.png">
   Nicht das Richtige dabei? 
Die Geofabrik ist ein auf OpenStreetMap spezialisiertes Beratungs- und Softwareentwicklungsunternehmen in Karlsruhe.
Gern helfen wir Ihnen bei der Datenaufbereitung, Datenkonvertierung, Serverinstallation und ähnlichen Aufgaben.
<a href="http://www.geofabrik.de/data/shapefiles.html">Besuchen Sie unsere Webseite</a> und sprechen Sie mit uns, wenn wir Ihnen 
helfen können.
</p></div>
</div>
</div>

<div id="download-left">
<h3>Commonly Used Formats</h3>
<ul><li><a href="south-america-latest.osm.pbf">south-america-latest.osm.pbf,</a> suitable for
	Osmium, Osmosis, imposm, osm2pgsql, mkgmap, and others. This file was last 
	modified 8 hours ago and contains all OSM data up to 2018-03-22T21:43:02Z. File size: 1.3&nbsp;GB; MD5 sum: <a href="south-america-latest.osm.pbf.md5">0da8246eafb397f64c1ddaa105e258b4</a>. </li>
<li><s>south-america-latest-free.shp.zip</s> is not available for this region; try one of the sub-regions.</li></ul><h3>Other Formats and Auxiliary Files</h3>
<ul><li><a href="south-america-latest.osm.bz2">south-america-latest.osm.bz2,</a> yields
		 OSM XML when decompressed; use for programs that cannot process the .pbf format.
		 This file was last modified 11 days ago. File size: 2.3&nbsp;GB; MD5 sum: <a href="south-america-latest.osm.bz2.md5">4223edaf50e53687b8fe9d648039895a</a>. </li>
<li><a href="south-america.osh.pbf">south-america.osh.pbf,</a> a file that
		 contains the full OSM history for this region for processing with e.g. osmium.
		 This file was last modified 4 days ago. File size: 2.1&nbsp;GB; MD5 sum: <a href="south-america.osh.pbf.md5">e400f12ca8e39aca9725c8571bf3afc2</a>. </li>
<li><a href="south-america.poly">.poly file</a> that describes the extent
		of this region.</li><li><a href="south-america-updates">.osc.gz files</a> that contain 
		all changes in this region, suitable e.g. for Osmosis updates</li><li><a href="#" onclick="OpenLayers.Util.getElement('details').style.display='inline';">raw directory index</a> allowing you to
	 see and download older files</li></ul><div id="details" style="display:none"><table id="subregions"><tbody><tr><th class="subregion">file</th><th>date</th><th>size</th></tr><tr><td class="subregion"><a href="south-america-140101.osm.pbf">south-america-140101.osm.pbf</a></td><td class="subright">2014-01-01 22:22</td><td class="subright">469796473</td></tr>
<tr><td class="subregion"><a href="south-america-150101.osm.pbf">south-america-150101.osm.pbf</a></td><td class="subright">2015-01-01 23:09</td><td class="subright">644041399</td></tr>
<tr><td class="subregion"><a href="south-america-160101.osm.pbf">south-america-160101.osm.pbf</a></td><td class="subright">2016-01-01 22:19</td><td class="subright">791972821</td></tr>
<tr><td class="subregion"><a href="south-america-170101.osm.pbf">south-america-170101.osm.pbf</a></td><td class="subright">2017-01-01 22:39</td><td class="subright">1048007171</td></tr>
<tr><td class="subregion"><a href="south-america-170101.osm.pbf.md5">south-america-170101.osm.pbf.md5</a></td><td class="subright">2017-01-01 22:41</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-180101.osm.pbf">south-america-180101.osm.pbf</a></td><td class="subright">2018-01-01 22:49</td><td class="subright">1306297360</td></tr>
<tr><td class="subregion"><a href="south-america-180101.osm.pbf.md5">south-america-180101.osm.pbf.md5</a></td><td class="subright">2018-01-01 22:49</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-180201.osm.pbf">south-america-180201.osm.pbf</a></td><td class="subright">2018-02-02 00:24</td><td class="subright">1325657445</td></tr>
<tr><td class="subregion"><a href="south-america-180201.osm.pbf.md5">south-america-180201.osm.pbf.md5</a></td><td class="subright">2018-02-02 00:30</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-180301.osm.pbf">south-america-180301.osm.pbf</a></td><td class="subright">2018-03-01 22:53</td><td class="subright">1353074577</td></tr>
<tr><td class="subregion"><a href="south-america-180301.osm.pbf.md5">south-america-180301.osm.pbf.md5</a></td><td class="subright">2018-03-01 22:54</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-180319.osm.pbf">south-america-180319.osm.pbf</a></td><td class="subright">2018-03-19 23:11</td><td class="subright">1365714353</td></tr>
<tr><td class="subregion"><a href="south-america-180319.osm.pbf.md5">south-america-180319.osm.pbf.md5</a></td><td class="subright">2018-03-19 23:12</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-180320.osm.pbf">south-america-180320.osm.pbf</a></td><td class="subright">2018-03-20 22:52</td><td class="subright">1366973188</td></tr>
<tr><td class="subregion"><a href="south-america-180320.osm.pbf.md5">south-america-180320.osm.pbf.md5</a></td><td class="subright">2018-03-20 22:52</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-180321.osm.pbf">south-america-180321.osm.pbf</a></td><td class="subright">2018-03-21 23:00</td><td class="subright">1367669194</td></tr>
<tr><td class="subregion"><a href="south-america-180321.osm.pbf.md5">south-america-180321.osm.pbf.md5</a></td><td class="subright">2018-03-21 23:03</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-180322.osm.pbf">south-america-180322.osm.pbf</a></td><td class="subright">2018-03-22 22:54</td><td class="subright">1368961348</td></tr>
<tr><td class="subregion"><a href="south-america-180322.osm.pbf.md5">south-america-180322.osm.pbf.md5</a></td><td class="subright">2018-03-22 22:55</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-latest.osm.bz2">south-america-latest.osm.bz2</a></td><td class="subright">2018-03-12 00:52</td><td class="subright">2462112782</td></tr>
<tr><td class="subregion"><a href="south-america-latest.osm.bz2.md5">south-america-latest.osm.bz2.md5</a></td><td class="subright">2018-03-12 00:52</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-latest.osm.pbf">south-america-latest.osm.pbf</a></td><td class="subright">2018-03-22 22:54</td><td class="subright">1368961348</td></tr>
<tr><td class="subregion"><a href="south-america-latest.osm.pbf.md5">south-america-latest.osm.pbf.md5</a></td><td class="subright">2018-03-22 22:55</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="south-america-updates">south-america-updates</a></td><td class="subright">2013-02-26 08:45</td><td class="subright">4096</td></tr>
<tr><td class="subregion"><a href="south-america.kml">south-america.kml</a></td><td class="subright">2018-03-22 22:55</td><td class="subright">1108</td></tr>
<tr><td class="subregion"><a href="south-america.osh.pbf">south-america.osh.pbf</a></td><td class="subright">2018-03-19 06:49</td><td class="subright">2275642028</td></tr>
<tr><td class="subregion"><a href="south-america.osh.pbf.md5">south-america.osh.pbf.md5</a></td><td class="subright">2018-03-19 06:49</td><td class="subright">56</td></tr>
<tr><td class="subregion"><a href="south-america.poly">south-america.poly</a></td><td class="subright">2018-03-22 22:43</td><td class="subright">850</td></tr>
</tbody></table></div><h3>Sub Regions</h3><p>Click on the region name to see the overview page for that region, or select one of the
file extension links for quick access.</p>
<table id="subregions">
<tbody><tr>
<th class="subregion">Sub Region</th>
<th colspan="4">Quick Links</th>
</tr>
<tr>
<th>&nbsp;</th>
<th colspan="2">.osm.pbf</th>
<th>.shp.zip</th>
<th>.osm.bz2</th>
</tr>
<tr onmouseover="loadkml('south-america/argentina.kml')"><td class="subregion"><a href="south-america/argentina.html">Argentina</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/argentina-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(165&nbsp;MB)</td><td> <a href="south-america/argentina-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="south-america/argentina-latest.osm.bz2">[.osm.bz2]</a></td></tr>
<tr onmouseover="loadkml('south-america/bolivia.kml')"><td class="subregion"><a href="south-america/bolivia.html">Bolivia</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/bolivia-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(70&nbsp;MB)</td><td> <a href="south-america/bolivia-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="south-america/bolivia-latest.osm.bz2">[.osm.bz2]</a></td></tr>
<tr onmouseover="loadkml('south-america/brazil.kml')"><td class="subregion"><a href="south-america/brazil.html">Brazil</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/brazil-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(558&nbsp;MB)</td><td> <a href="south-america/brazil-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="south-america/brazil-latest.osm.bz2">[.osm.bz2]</a></td></tr>
<tr onmouseover="loadkml('south-america/chile.kml')"><td class="subregion"><a href="south-america/chile.html">Chile</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/chile-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(133&nbsp;MB)</td><td> <a href="south-america/chile-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="south-america/chile-latest.osm.bz2">[.osm.bz2]</a></td></tr>
<tr onmouseover="loadkml('south-america/colombia.kml')"><td class="subregion"><a href="south-america/colombia.html">Colombia</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/colombia-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(112&nbsp;MB)</td><td> <a href="south-america/colombia-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="south-america/colombia-latest.osm.bz2">[.osm.bz2]</a></td></tr>
<tr onmouseover="loadkml('south-america/ecuador.kml')"><td class="subregion"><a href="south-america/ecuador.html">Ecuador</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/ecuador-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(60&nbsp;MB)</td><td> <img src="/img/cross.png"></td><td> <a href="south-america/ecuador-latest.osm.bz2">[.osm.bz2]</a></td></tr>
<tr onmouseover="loadkml('south-america/paraguay.kml')"><td class="subregion"><a href="south-america/paraguay.html">Paraguay</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/paraguay-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(44.6&nbsp;MB)</td><td> <a href="south-america/paraguay-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="south-america/paraguay-latest.osm.bz2">[.osm.bz2]</a></td></tr>
<tr onmouseover="loadkml('south-america/peru.kml')"><td class="subregion"><a href="south-america/peru.html">Peru</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/peru-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(82&nbsp;MB)</td><td> <a href="south-america/peru-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="south-america/peru-latest.osm.bz2">[.osm.bz2]</a></td></tr>
<tr onmouseover="loadkml('south-america/suriname.kml')"><td class="subregion"><a href="south-america/suriname.html">suriname</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/suriname-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(5.1&nbsp;MB)</td><td> <a href="south-america/suriname-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="south-america/suriname-latest.osm.bz2">[.osm.bz2]</a></td></tr>
<tr onmouseover="loadkml('south-america/uruguay.kml')"><td class="subregion"><a href="south-america/uruguay.html">Uruguay</a></td>
<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/uruguay-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(21.9&nbsp;MB)</td><td> <a href="south-america/uruguay-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="south-america/uruguay-latest.osm.bz2">[.osm.bz2]</a></td></tr>
</tbody></table>
</div>
</div>
<div id="download-copyright">
Data/Maps Copyright 2016
<a href="http://www.geofabrik.de/">Geofabrik GmbH</a>
and
<a href="http://www.openstreetmap.org/">OpenStreetMap Contributors</a>
|
<a href="http://creativecommons.org/licenses/by-sa/2.0/" rel="license">Map tiles: Creative Commons BY-SA 2.0</a>
<a href="http://opendatacommons.org/licenses/odbl/">Data: ODbL 1.0</a>
|
<a href="http://www.geofabrik.de/geofabrik/contact.html">Contact</a>
</div>
`
const geofabrikDistrictOfColumbiaHTML = `<div id="download-top">
<div id="download-top-right">
   <img src="/img/geofabrik-downloads.png" height="20">
</div>
</div>
<div id="download-main">
<div id="download-titlebar">
<h1>Download OpenStreetMap data for this region:</h1>
</div>
<h2>District of Columbia</h2><p><a href="../../north-america.html">[one level up]</a></p><div id="download-right">
<div id="noscroll">
<div id="map" class="olMap"><div id="OpenLayers.Map_2_OpenLayers_ViewPort" style="position: relative; overflow: hidden; width: 100%; height: 100%;" class="olMapViewport olControlDragPanActive olControlZoomBoxActive olControlPinchZoomActive olControlNavigationActive"><div id="OpenLayers.Map_2_OpenLayers_Container" style="position: absolute; width: 100px; height: 100px; z-index: 749; left: 0px; top: 0px;"><div id="OpenLayers.Layer.OSM_18" style="position: absolute; width: 100%; height: 100%; z-index: 100; left: 0%; top: 0%;" dir="ltr" class="olLayerDiv olLayerGrid"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: -40%; top: -39%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/10/292/391.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 216%; top: -39%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/10/293/391.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: -40%; top: 217%; width: 256%; height: 256%;" src="http://a.tile.openstreetmap.org/10/292/392.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 216%; top: 217%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/10/293/392.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 472%; top: -39%; width: 256%; height: 256%;" src="http://c.tile.openstreetmap.org/10/294/391.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 472%; top: 217%; width: 256%; height: 256%;" src="http://a.tile.openstreetmap.org/10/294/392.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: -40%; top: 473%; width: 256%; height: 256%;" src="http://c.tile.openstreetmap.org/10/292/393.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 216%; top: 473%; width: 256%; height: 256%;" src="http://c.tile.openstreetmap.org/10/293/393.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 472%; top: 473%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/10/294/393.png"></div><div id="OpenLayers.Layer.Vector_25" style="position: absolute; width: 100%; height: 100%; z-index: 330; left: 0px; top: 0px;" dir="ltr" class="olLayerDiv"><svg id="OpenLayers.Layer.Vector_25_svgRoot" style="display: block;" width="400" height="300" viewBox="0 0 400 300"><g id="OpenLayers.Layer.Vector_25_root" style="visibility: visible;" transform=""><g id="OpenLayers.Layer.Vector_25_vroot"><path id="OpenLayers.Geometry.Polygon_56" d=" M 182.14508088861476,245.75261818065337 181.99944533305097,236.78330270510196 183.06258488860476,229.04629751964603 183.0771484441575,224.5045930424567 180.57949866638228,216.0932107948356 182.8659768885991,209.63441688832972 179.2396515552755,208.5687958523995 177.87795911082503,205.5120513494221 175.75896177749382,199.03353225926185 184.03834311082755,199.11767291441356 185.8951964441585,196.9206320278172 187.0821262219397,190.60022892575216 186.57968355526828,187.21538742149642 185.15973688860686,183.06357862068762 179.45082311083388,178.2288259379893 182.79315911082085,176.73250629217 182.028572444171,173.30019589286894 179.4653866663939,169.75547136894238 173.4433564441715,172.1124068831341 165.00377599972126,158.74597924339105 160.41625599971303,144.94646175361777 146.1730986663897,141.30665375300669 137.10000355527882,135.0558240792925 125.87878399971669,117.99405395378199 123.1626808886067,112.10611754643833 180.7833884441643,54.247381819783186 276.83731911083305,150.7847090522264 182.14508088861476,245.75261818065337 z" fill-rule="evenodd" fill="#ffcc66" fill-opacity="0.5" stroke="#ff9933" stroke-opacity="1" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path></g><g id="OpenLayers.Layer.Vector_25_troot"></g></g></svg></div></div><div id="OpenLayers.Control.Zoom_5" style="position: absolute; z-index: 1004;" class="olControlZoom olControlNoSelect" unselectable="on"><a href="#zoomIn" class="olControlZoomIn olButton">+</a><a href="#zoomOut" class="olControlZoomOut olButton">-</a></div><div id="OpenLayers.Control.ArgParser_6" style="position: absolute; z-index: 1004;" class="olControlArgParser olControlNoSelect" unselectable="on"></div><div id="OpenLayers.Control.Attribution_7" style="position: absolute; z-index: 1004;" class="olControlAttribution olControlNoSelect" unselectable="on"></div></div></div>

<div class="customized_shapes">
<p>
   <img src="/img/en.png">
Not what you were looking for?
Geofabrik is a consulting and software development firm based in Karlsruhe, Germany
specializing in OpenStreetMap services. We're happy to help you with data
preparation, processing, server setup and the like. <a href="http://www.geofabrik.de/data/shapefiles.html">Check out our web site</a>
and contact us if we can be of service.

   </p>
</div>

<div class="customized_shapes">
   <p>
   <img src="/img/de.png">
   Nicht das Richtige dabei? 
Die Geofabrik ist ein auf OpenStreetMap spezialisiertes Beratungs- und Softwareentwicklungsunternehmen in Karlsruhe.
Gern helfen wir Ihnen bei der Datenaufbereitung, Datenkonvertierung, Serverinstallation und ähnlichen Aufgaben.
<a href="http://www.geofabrik.de/data/shapefiles.html">Besuchen Sie unsere Webseite</a> und sprechen Sie mit uns, wenn wir Ihnen 
helfen können.
</p></div>
</div>
</div>

<div id="download-left">
<h3>Commonly Used Formats</h3>
<ul><li><a href="district-of-columbia-latest.osm.pbf">district-of-columbia-latest.osm.pbf,</a> suitable for
	Osmium, Osmosis, imposm, osm2pgsql, mkgmap, and others. This file was last 
	modified 13 hours ago and contains all OSM data up to 2018-03-22T21:43:02Z. File size: 16.3&nbsp;MB; MD5 sum: <a href="district-of-columbia-latest.osm.pbf.md5">63ba83670813da52f2a749e02b1b9c78</a>. </li>
<li><a href="district-of-columbia-latest-free.shp.zip">district-of-columbia-latest-free.shp.zip,</a> yields
	a number of ESRI compatible shape files when unzipped. <a href="/osm-data-in-gis-formats-free.pdf">(Format description PDF)</a>
	This file was last modified 10 hours ago. File size: 23.9&nbsp;MB. </li>
</ul><h3>Other Formats and Auxiliary Files</h3>
<ul><li><a href="district-of-columbia-latest.osm.bz2">district-of-columbia-latest.osm.bz2,</a> yields
		 OSM XML when decompressed; use for programs that cannot process the .pbf format.
		 This file was last modified 6 days ago. File size: 24.8&nbsp;MB; MD5 sum: <a href="district-of-columbia-latest.osm.bz2.md5">0ff713de1d64cf6c69c5c5391bbae34f</a>. </li>
<li><a href="district-of-columbia.osh.pbf">district-of-columbia.osh.pbf,</a> a file that
		 contains the full OSM history for this region for processing with e.g. osmium.
		 This file was last modified 1 hour ago. File size: 37.8&nbsp;MB; MD5 sum: <a href="district-of-columbia.osh.pbf.md5">86058f3672737fbe4a26de908f348bba</a>. </li>
<li><a href="district-of-columbia.poly">.poly file</a> that describes the extent
		of this region.</li><li><a href="district-of-columbia-updates">.osc.gz files</a> that contain 
		all changes in this region, suitable e.g. for Osmosis updates</li><li><a href="#" onclick="OpenLayers.Util.getElement('details').style.display='inline';">raw directory index</a> allowing you to
	 see and download older files</li></ul><div id="details" style="display:none"><table id="subregions"><tbody><tr><th class="subregion">file</th><th>date</th><th>size</th></tr><tr><td class="subregion"><a href="district-of-columbia-140101.osm.pbf">district-of-columbia-140101.osm.pbf</a></td><td class="subright">2014-01-02 00:11</td><td class="subright">12307550</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-150101.osm.pbf">district-of-columbia-150101.osm.pbf</a></td><td class="subright">2015-01-02 00:50</td><td class="subright">15630802</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-160101.osm.pbf">district-of-columbia-160101.osm.pbf</a></td><td class="subright">2016-01-01 23:40</td><td class="subright">16116372</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-170101-free.shp.zip">district-of-columbia-170101-free.shp.zip</a></td><td class="subright">2017-01-02 06:19</td><td class="subright">24409201</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-170101.osm.pbf">district-of-columbia-170101.osm.pbf</a></td><td class="subright">2017-01-02 00:10</td><td class="subright">16405727</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-170101.osm.pbf.md5">district-of-columbia-170101.osm.pbf.md5</a></td><td class="subright">2017-01-02 00:12</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180101-free.shp.zip">district-of-columbia-180101-free.shp.zip</a></td><td class="subright">2018-01-02 01:31</td><td class="subright">24678996</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180101.osm.pbf">district-of-columbia-180101.osm.pbf</a></td><td class="subright">2018-01-01 23:17</td><td class="subright">16894218</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180101.osm.pbf.md5">district-of-columbia-180101.osm.pbf.md5</a></td><td class="subright">2018-01-01 23:17</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180201-free.shp.zip">district-of-columbia-180201-free.shp.zip</a></td><td class="subright">2018-02-02 06:34</td><td class="subright">24710532</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180201.osm.pbf">district-of-columbia-180201.osm.pbf</a></td><td class="subright">2018-02-02 02:56</td><td class="subright">16933070</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180201.osm.pbf.md5">district-of-columbia-180201.osm.pbf.md5</a></td><td class="subright">2018-02-02 03:00</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180301-free.shp.zip">district-of-columbia-180301-free.shp.zip</a></td><td class="subright">2018-03-02 01:59</td><td class="subright">24992593</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180301.osm.pbf">district-of-columbia-180301.osm.pbf</a></td><td class="subright">2018-03-01 23:20</td><td class="subright">17082337</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180301.osm.pbf.md5">district-of-columbia-180301.osm.pbf.md5</a></td><td class="subright">2018-03-01 23:21</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180319-free.shp.zip">district-of-columbia-180319-free.shp.zip</a></td><td class="subright">2018-03-20 02:41</td><td class="subright">25079801</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180319.osm.pbf">district-of-columbia-180319.osm.pbf</a></td><td class="subright">2018-03-19 23:50</td><td class="subright">17112900</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180319.osm.pbf.md5">district-of-columbia-180319.osm.pbf.md5</a></td><td class="subright">2018-03-19 23:51</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180320-free.shp.zip">district-of-columbia-180320-free.shp.zip</a></td><td class="subright">2018-03-21 02:45</td><td class="subright">25080021</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180320.osm.pbf">district-of-columbia-180320.osm.pbf</a></td><td class="subright">2018-03-20 23:15</td><td class="subright">17113449</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180320.osm.pbf.md5">district-of-columbia-180320.osm.pbf.md5</a></td><td class="subright">2018-03-20 23:16</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180321.osm.pbf">district-of-columbia-180321.osm.pbf</a></td><td class="subright">2018-03-21 23:53</td><td class="subright">17113978</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180321.osm.pbf.md5">district-of-columbia-180321.osm.pbf.md5</a></td><td class="subright">2018-03-21 23:54</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180322-free.shp.zip">district-of-columbia-180322-free.shp.zip</a></td><td class="subright">2018-03-23 02:41</td><td class="subright">25079950</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180322.osm.pbf">district-of-columbia-180322.osm.pbf</a></td><td class="subright">2018-03-22 23:44</td><td class="subright">17114989</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-180322.osm.pbf.md5">district-of-columbia-180322.osm.pbf.md5</a></td><td class="subright">2018-03-22 23:45</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-latest-free.shp.zip">district-of-columbia-latest-free.shp.zip</a></td><td class="subright">2018-03-23 02:41</td><td class="subright">25079950</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-latest.osm.bz2">district-of-columbia-latest.osm.bz2</a></td><td class="subright">2018-03-17 12:13</td><td class="subright">26038378</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-latest.osm.bz2.md5">district-of-columbia-latest.osm.bz2.md5</a></td><td class="subright">2018-03-17 12:13</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-latest.osm.pbf">district-of-columbia-latest.osm.pbf</a></td><td class="subright">2018-03-22 23:44</td><td class="subright">17114989</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-latest.osm.pbf.md5">district-of-columbia-latest.osm.pbf.md5</a></td><td class="subright">2018-03-22 23:45</td><td class="subright">70</td></tr>
<tr><td class="subregion"><a href="district-of-columbia-updates">district-of-columbia-updates</a></td><td class="subright">2013-02-26 08:53</td><td class="subright">4096</td></tr>
<tr><td class="subregion"><a href="district-of-columbia.kml">district-of-columbia.kml</a></td><td class="subright">2018-03-22 23:45</td><td class="subright">1022</td></tr>
<tr><td class="subregion"><a href="district-of-columbia.osh.pbf">district-of-columbia.osh.pbf</a></td><td class="subright">2018-03-23 11:52</td><td class="subright">39605103</td></tr>
<tr><td class="subregion"><a href="district-of-columbia.osh.pbf.md5">district-of-columbia.osh.pbf.md5</a></td><td class="subright">2018-03-23 11:52</td><td class="subright">63</td></tr>
<tr><td class="subregion"><a href="district-of-columbia.poly">district-of-columbia.poly</a></td><td class="subright">2018-03-22 23:40</td><td class="subright">911</td></tr>
</tbody></table></div><h3>Sub Regions</h3><p>No sub regions are defined for this region.</p></div>
</div>
<div id="download-copyright">
Data/Maps Copyright 2016
<a href="http://www.geofabrik.de/">Geofabrik GmbH</a>
and
<a href="http://www.openstreetmap.org/">OpenStreetMap Contributors</a>
|
<a href="http://creativecommons.org/licenses/by-sa/2.0/" rel="license">Map tiles: Creative Commons BY-SA 2.0</a>
<a href="http://opendatacommons.org/licenses/odbl/">Data: ODbL 1.0</a>
|
<a href="http://www.geofabrik.de/geofabrik/contact.html">Contact</a>
</div>
`
const geofabrikShikokuHTML = `<div id="download-top">
<div id="download-top-right">
   <img src="/img/geofabrik-downloads.png" height="20">
</div>
</div>
<div id="download-main">
<div id="download-titlebar">
<h1>Download OpenStreetMap data for this region:</h1>
</div>
<h2>Shikoku</h2><p><a href="../japan.html">[one level up]</a></p><div id="download-right">
<div id="noscroll">
<div id="map" class="olMap"><div id="OpenLayers.Map_2_OpenLayers_ViewPort" style="position: relative; overflow: hidden; width: 100%; height: 100%;" class="olMapViewport olControlDragPanActive olControlZoomBoxActive olControlPinchZoomActive olControlNavigationActive"><div id="OpenLayers.Map_2_OpenLayers_Container" style="position: absolute; width: 100px; height: 100px; z-index: 749; left: 0px; top: 0px;"><div id="OpenLayers.Layer.OSM_18" style="position: absolute; width: 100%; height: 100%; z-index: 100; left: 0%; top: 0%;" dir="ltr" class="olLayerDiv olLayerGrid"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 83%; top: 56%; width: 256%; height: 256%;" src="http://a.tile.openstreetmap.org/7/111/51.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 83%; top: -200%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/7/111/50.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: -173%; top: 56%; width: 256%; height: 256%;" src="http://c.tile.openstreetmap.org/7/110/51.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 339%; top: 56%; width: 256%; height: 256%;" src="http://a.tile.openstreetmap.org/7/112/51.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 83%; top: 312%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/7/111/52.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: -173%; top: -200%; width: 256%; height: 256%;" src="http://a.tile.openstreetmap.org/7/110/50.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 339%; top: -200%; width: 256%; height: 256%;" src="http://b.tile.openstreetmap.org/7/112/50.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: -173%; top: 312%; width: 256%; height: 256%;" src="http://a.tile.openstreetmap.org/7/110/52.png"><img class="olTileImage" style="visibility: inherit; opacity: 1; position: absolute; left: 339%; top: 312%; width: 256%; height: 256%;" src="http://c.tile.openstreetmap.org/7/112/52.png"></div><div id="OpenLayers.Layer.Vector_25" style="position: absolute; width: 100%; height: 100%; z-index: 330; left: 0px; top: 0px;" dir="ltr" class="olLayerDiv"><svg id="OpenLayers.Layer.Vector_25_svgRoot" style="display: block;" width="400" height="300" viewBox="0 0 400 300"><g id="OpenLayers.Layer.Vector_25_root" style="visibility: visible;" transform=""><g id="OpenLayers.Layer.Vector_25_vroot"><path id="OpenLayers.Geometry.Polygon_116" d=" M 83.08195555560451,282.1092612926973 73.24245333338331,180.09412454570474 71.98634666671387,179.8048586922887 44.82531555560672,148.12052176616453 73.26065777782605,131.66086453575736 103.51644444449448,114.67966015193542 106.31992888893728,102.3901072266608 106.49287111116246,102.11160844187225 109.96992000005048,97.84756785916124 111.01667555560562,94.50450691005426 108.70471111115876,93.17550799559831 107.70346666671685,88.62958138512204 132.19754666671724,73.46741776411181 137.74990222227098,78.73856798806901 145.67793777782572,73.53670758644375 144.9770666667173,69.96375850830054 144.01223111115905,68.69397443518164 144.03043555560544,67.58359624170998 145.37756444449406,67.3788935850871 150.13802666671654,66.61066806491954 150.2472533333821,66.19351036293574 149.81034666671803,65.01788322854873 152.7230577778264,56.32374131776851 153.66058666671415,55.90061050659824 156.0726755556043,56.99255707152997 156.9920000000493,56.30611125359201 158.0933688889363,55.06531111380036 160.4326400000482,55.24273589451377 160.99697777782785,56.4691880656178 162.48974222227116,58.548165591620545 161.96181333338427,59.64420584423442 162.00732444449386,60.085889916393626 162.6171733333831,60.82381880473258 162.94485333338343,61.83809974955602 167.02264888893842,61.72907813389884 168.02389333338215,60.05725268496644 170.5725155556047,58.125136461806505 170.79096888893764,57.22062651618853 172.01976888893842,57.20299816048737 172.60231111116082,57.55115196928 173.18485333338322,58.76848560845474 174.44096000004902,59.73342552164377 175.13272888894062,59.27079564515225 175.44220444449275,58.26064008410549 175.78808888893946,57.7505632842699 176.3979377778287,57.406823973102746 177.12611555560215,56.05157345885209 178.36401777782703,55.15126888591385 182.75128888893778,53.38897487967961 204.11420444449323,58.37631233885941 208.00085333338393,55.301142065469776 211.65084444449167,48.24344686412951 209.35708444449483,43.043188117856516 215.92888888893867,37.6434670008598 226.05966222227107,44.576874070819485 226.99719111115883,43.76923813590156 228.12586666671632,43.24180930162993 228.2897066667174,42.5433058808685 228.93596444449213,42.115129136804626 229.56401777782594,42.0323607211285 230.63808000004974,42.3656370156782 230.97486222227053,43.20649918148638 231.7667555556036,42.85780448808282 232.33109333338325,43.81668318475067 232.95914666671524,43.448150113876636 243.17184000005,40.619665226227426 243.17184000005,38.65587966736257 242.50737777782706,37.60923954703958 242.6803200000486,37.046124471991334 244.7556266667143,35.36319897774047 244.55537777782592,33.36184000615003 245.22894222227114,32.14446516675025 245.9298133333832,31.936766264187554 246.3212088889377,32.241684318938496 247.05848888893888,34.04999429263444 247.65013333338175,34.857377802158226 248.733297777826,34.669620513142945 249.16110222227144,34.841915581845115 249.85287111116122,33.80809680313132 252.90211555560563,30.857319341410403 277.7238755556027,17.890738707366836 302.11783111115983,37.953716508533034 302.9006222222706,52.673601503314785 303.1372800000481,53.42204164823261 307.70659555560815,62.865482884074936 335.18620444449516,89.84285024681094 355.1746844444933,207.91293948871908 83.08195555560451,282.1092612926973 z" fill-rule="evenodd" fill="#ffcc66" fill-opacity="0.5" stroke="#ff9933" stroke-opacity="1" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path></g><g id="OpenLayers.Layer.Vector_25_troot"></g></g></svg></div></div><div id="OpenLayers.Control.Zoom_5" style="position: absolute; z-index: 1004;" class="olControlZoom olControlNoSelect" unselectable="on"><a href="#zoomIn" class="olControlZoomIn olButton">+</a><a href="#zoomOut" class="olControlZoomOut olButton">-</a></div><div id="OpenLayers.Control.ArgParser_6" style="position: absolute; z-index: 1004;" class="olControlArgParser olControlNoSelect" unselectable="on"></div><div id="OpenLayers.Control.Attribution_7" style="position: absolute; z-index: 1004;" class="olControlAttribution olControlNoSelect" unselectable="on"></div></div></div>

<div class="customized_shapes">
<p>
   <img src="/img/en.png">
Not what you were looking for?
Geofabrik is a consulting and software development firm based in Karlsruhe, Germany
specializing in OpenStreetMap services. We're happy to help you with data
preparation, processing, server setup and the like. <a href="http://www.geofabrik.de/data/shapefiles.html">Check out our web site</a>
and contact us if we can be of service.

   </p>
</div>

<div class="customized_shapes">
   <p>
   <img src="/img/de.png">
   Nicht das Richtige dabei? 
Die Geofabrik ist ein auf OpenStreetMap spezialisiertes Beratungs- und Softwareentwicklungsunternehmen in Karlsruhe.
Gern helfen wir Ihnen bei der Datenaufbereitung, Datenkonvertierung, Serverinstallation und ähnlichen Aufgaben.
<a href="http://www.geofabrik.de/data/shapefiles.html">Besuchen Sie unsere Webseite</a> und sprechen Sie mit uns, wenn wir Ihnen 
helfen können.
</p></div>
</div>
</div>

<div id="download-left">
<h3>Commonly Used Formats</h3>
<ul><li><a href="shikoku-latest.osm.pbf">shikoku-latest.osm.pbf,</a> suitable for
	Osmium, Osmosis, imposm, osm2pgsql, mkgmap, and others. This file was last 
	modified 13 hours ago and contains all OSM data up to 2018-03-22T21:43:02Z. File size: 54&nbsp;MB; MD5 sum: <a href="shikoku-latest.osm.pbf.md5">844cd127c4e67f313c1fabe81e3c20a1</a>. </li>
<li><a href="shikoku-latest-free.shp.zip">shikoku-latest-free.shp.zip,</a> yields
	a number of ESRI compatible shape files when unzipped. <a href="/osm-data-in-gis-formats-free.pdf">(Format description PDF)</a>
	This file was last modified 10 hours ago. File size: 89&nbsp;MB. </li>
</ul><h3>Other Formats and Auxiliary Files</h3>
<ul><li><a href="shikoku-latest.osm.bz2">shikoku-latest.osm.bz2,</a> yields
		 OSM XML when decompressed; use for programs that cannot process the .pbf format.
		 This file was last modified 5 days ago. File size: 92&nbsp;MB; MD5 sum: <a href="shikoku-latest.osm.bz2.md5">6b89aef0a0ddfdfa388cf72599d6884a</a>. </li>
<li><a href="shikoku.osh.pbf">shikoku.osh.pbf,</a> a file that
		 contains the full OSM history for this region for processing with e.g. osmium.
		 This file was last modified 4 days ago. File size: 118&nbsp;MB; MD5 sum: <a href="shikoku.osh.pbf.md5">4b392433c51502f84d1a1870282b354e</a>. </li>
<li><a href="shikoku.poly">.poly file</a> that describes the extent
		of this region.</li><li><a href="shikoku-updates">.osc.gz files</a> that contain 
		all changes in this region, suitable e.g. for Osmosis updates</li><li><a href="#" onclick="OpenLayers.Util.getElement('details').style.display='inline';">raw directory index</a> allowing you to
	 see and download older files</li></ul><div id="details" style="display:none"><table id="subregions"><tbody><tr><th class="subregion">file</th><th>date</th><th>size</th></tr><tr><td class="subregion"><a href="shikoku-180101-free.shp.zip">shikoku-180101-free.shp.zip</a></td><td class="subright">2018-01-02 01:39</td><td class="subright">90829809</td></tr>
<tr><td class="subregion"><a href="shikoku-180101.osm.pbf">shikoku-180101.osm.pbf</a></td><td class="subright">2018-01-01 23:49</td><td class="subright">55839357</td></tr>
<tr><td class="subregion"><a href="shikoku-180101.osm.pbf.md5">shikoku-180101.osm.pbf.md5</a></td><td class="subright">2018-01-01 23:49</td><td class="subright">57</td></tr>
<tr><td class="subregion"><a href="shikoku-180201-free.shp.zip">shikoku-180201-free.shp.zip</a></td><td class="subright">2018-02-02 06:59</td><td class="subright">91912267</td></tr>
<tr><td class="subregion"><a href="shikoku-180201.osm.pbf">shikoku-180201.osm.pbf</a></td><td class="subright">2018-02-02 03:34</td><td class="subright">56535911</td></tr>
<tr><td class="subregion"><a href="shikoku-180201.osm.pbf.md5">shikoku-180201.osm.pbf.md5</a></td><td class="subright">2018-02-02 03:35</td><td class="subright">57</td></tr>
<tr><td class="subregion"><a href="shikoku-180301-free.shp.zip">shikoku-180301-free.shp.zip</a></td><td class="subright">2018-03-02 01:41</td><td class="subright">92620703</td></tr>
<tr><td class="subregion"><a href="shikoku-180301.osm.pbf">shikoku-180301.osm.pbf</a></td><td class="subright">2018-03-02 00:01</td><td class="subright">57041819</td></tr>
<tr><td class="subregion"><a href="shikoku-180301.osm.pbf.md5">shikoku-180301.osm.pbf.md5</a></td><td class="subright">2018-03-02 00:03</td><td class="subright">57</td></tr>
<tr><td class="subregion"><a href="shikoku-180317.osm.bz2">shikoku-180317.osm.bz2</a></td><td class="subright">2018-03-18 05:47</td><td class="subright">96492886</td></tr>
<tr><td class="subregion"><a href="shikoku-180317.osm.bz2.md5">shikoku-180317.osm.bz2.md5</a></td><td class="subright">2018-03-18 05:47</td><td class="subright">57</td></tr>
<tr><td class="subregion"><a href="shikoku-180319-free.shp.zip">shikoku-180319-free.shp.zip</a></td><td class="subright">2018-03-20 03:27</td><td class="subright">93301997</td></tr>
<tr><td class="subregion"><a href="shikoku-180319.osm.pbf">shikoku-180319.osm.pbf</a></td><td class="subright">2018-03-20 01:23</td><td class="subright">57426748</td></tr>
<tr><td class="subregion"><a href="shikoku-180319.osm.pbf.md5">shikoku-180319.osm.pbf.md5</a></td><td class="subright">2018-03-20 01:24</td><td class="subright">57</td></tr>
<tr><td class="subregion"><a href="shikoku-180320-free.shp.zip">shikoku-180320-free.shp.zip</a></td><td class="subright">2018-03-21 01:59</td><td class="subright">93344785</td></tr>
<tr><td class="subregion"><a href="shikoku-180320.osm.pbf">shikoku-180320.osm.pbf</a></td><td class="subright">2018-03-20 23:55</td><td class="subright">57452017</td></tr>
<tr><td class="subregion"><a href="shikoku-180320.osm.pbf.md5">shikoku-180320.osm.pbf.md5</a></td><td class="subright">2018-03-20 23:56</td><td class="subright">57</td></tr>
<tr><td class="subregion"><a href="shikoku-180322-free.shp.zip">shikoku-180322-free.shp.zip</a></td><td class="subright">2018-03-23 03:43</td><td class="subright">93400545</td></tr>
<tr><td class="subregion"><a href="shikoku-180322.osm.pbf">shikoku-180322.osm.pbf</a></td><td class="subright">2018-03-23 00:05</td><td class="subright">57484436</td></tr>
<tr><td class="subregion"><a href="shikoku-180322.osm.pbf.md5">shikoku-180322.osm.pbf.md5</a></td><td class="subright">2018-03-23 00:06</td><td class="subright">57</td></tr>
<tr><td class="subregion"><a href="shikoku-latest-free.shp.zip">shikoku-latest-free.shp.zip</a></td><td class="subright">2018-03-23 03:43</td><td class="subright">93400545</td></tr>
<tr><td class="subregion"><a href="shikoku-latest.osm.bz2">shikoku-latest.osm.bz2</a></td><td class="subright">2018-03-18 05:47</td><td class="subright">96492886</td></tr>
<tr><td class="subregion"><a href="shikoku-latest.osm.bz2.md5">shikoku-latest.osm.bz2.md5</a></td><td class="subright">2018-03-18 05:47</td><td class="subright">57</td></tr>
<tr><td class="subregion"><a href="shikoku-latest.osm.pbf">shikoku-latest.osm.pbf</a></td><td class="subright">2018-03-23 00:05</td><td class="subright">57484436</td></tr>
<tr><td class="subregion"><a href="shikoku-latest.osm.pbf.md5">shikoku-latest.osm.pbf.md5</a></td><td class="subright">2018-03-23 00:06</td><td class="subright">57</td></tr>
<tr><td class="subregion"><a href="shikoku-updates">shikoku-updates</a></td><td class="subright">2017-04-02 22:42</td><td class="subright">4096</td></tr>
<tr><td class="subregion"><a href="shikoku.kml">shikoku.kml</a></td><td class="subright">2018-03-23 00:06</td><td class="subright">2402</td></tr>
<tr><td class="subregion"><a href="shikoku.osh.pbf">shikoku.osh.pbf</a></td><td class="subright">2018-03-19 10:23</td><td class="subright">124186461</td></tr>
<tr><td class="subregion"><a href="shikoku.osh.pbf.md5">shikoku.osh.pbf.md5</a></td><td class="subright">2018-03-19 10:23</td><td class="subright">50</td></tr>
<tr><td class="subregion"><a href="shikoku.poly">shikoku.poly</a></td><td class="subright">2018-03-23 00:01</td><td class="subright">2743</td></tr>
</tbody></table></div><h3>Sub Regions</h3><p>No sub regions are defined for this region.</p></div>
</div>
<div id="download-copyright">
Data/Maps Copyright 2016
<a href="http://www.geofabrik.de/">Geofabrik GmbH</a>
and
<a href="http://www.openstreetmap.org/">OpenStreetMap Contributors</a>
|
<a href="http://creativecommons.org/licenses/by-sa/2.0/" rel="license">Map tiles: Creative Commons BY-SA 2.0</a>
<a href="http://opendatacommons.org/licenses/odbl/">Data: ODbL 1.0</a>
|
<a href="http://www.geofabrik.de/geofabrik/contact.html">Contact</a>
</div>
`

const gislabSampleHTML = `
<table>
  <tbody><tr>
    <th>ISO&nbsp;3166</th>
    <th>Страна/регион</th>
    <th>osm.pbf</th>
    <th>osm.bz2</th>
    <th colspan="2">&nbsp;</th>
  </tr>
    <tr><td>AM</td><td>Армения</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/AM.osm.pbf">03 Apr (20 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/AM.osm.bz2">03 Apr (38 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/AM/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/AM/" target="_blank">обновления</a></td></tr>    <tr><td>AZ</td><td>Азербайджан</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/AZ.osm.pbf">03 Apr (20 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/AZ.osm.bz2">03 Apr (37 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/AZ/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/AZ/" target="_blank">обновления</a></td></tr>    <tr><td>BY</td><td>Беларусь</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/BY.osm.pbf">03 Apr (194 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/BY.osm.bz2">03 Apr (319 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/BY/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/BY/" target="_blank">обновления</a></td></tr>    <tr><td>EE</td><td>Эстония</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/EE.osm.pbf">03 Apr (70 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/EE.osm.bz2">03 Apr (127 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/EE/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/EE/" target="_blank">обновления</a></td></tr>   <tr><td>RU</td><td>Российская Федерация</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/RU.osm.pbf">03 Apr (2313 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/RU.osm.bz2">03 Apr (4084 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/RU/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/RU/" target="_blank">обновления</a></td></tr>    <tr><td>TJ</td><td>Таджикистан</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/TJ.osm.pbf">03 Apr (19 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/TJ.osm.bz2">03 Apr (37 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/TJ/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/TJ/" target="_blank">обновления</a></td></tr>    <tr><td>TM</td><td>Туркмения</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/TM.osm.pbf">03 Apr (9 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/TM.osm.bz2">03 Apr (16 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/TM/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/TM/" target="_blank">обновления</a></td></tr>    <tr><td>UA</td><td>Украина</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/UA.osm.pbf">03 Apr (476 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/UA.osm.bz2">03 Apr (844 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/UA/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/UA/" target="_blank">обновления</a></td></tr>    <tr><td>UZ</td><td>Узбекистан</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/UZ.osm.pbf">03 Apr (37 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/UZ.osm.bz2">03 Apr (65 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/UZ/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/UZ/" target="_blank">обновления</a></td></tr>    <tr><td>local</td><td>локальное покрытие</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/local.osm.pbf">03 Apr (3444 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/local.osm.bz2">03 Apr (5999 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/local/" target="_blank">архив</a></td><td><a href="http://data.gis-lab.info/osm_dump/diff/local/" target="_blank">обновления</a></td></tr><tr><th colspan="6">Регионы России</th></tr>    <tr><td>RU-AD</td><td>Адыгея</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/RU-AD.osm.pbf">03 Apr (16 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/RU-AD.osm.bz2">03 Apr (26 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/RU-AD/" target="_blank">архив</a></td><td>&nbsp;</td></tr>    <tr><td>RU-AL</td><td>Алтай</td>
  <td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/RU-AL.osm.pbf">03 Apr (11 MB)</a></td><td style="white-space:nowrap;"><a href="http://data.gis-lab.info/osm_dump/dump/latest/RU-AL.osm.bz2">03 Apr (22 MB)</a></td><td><a href="http://data.gis-lab.info/osm_dump/dump/RU-AL/" target="_blank">архив</a></td><td>&nbsp;</td></tr>   
</tbody></table>
`

func TestElementSlice_Generate(t *testing.T) {
	myConfig := &SampleConfigValidPtr
	myYaml, _ := yaml.Marshal(*myConfig)
	myConfig.Elements = map[string]Element{} // void Elements

	type args struct {
		myConfig *Config
	}
	tests := []struct {
		name    string
		e       ElementSlice
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Marshalling OK, no error",
			e:       sampleElementValidPtr,
			args:    args{myConfig: myConfig},
			want:    myYaml,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.Generate(tt.args.myConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElementSlice.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementSlice.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		configfile string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Generate(tt.args.configfile)
		})
	}
}

func TestExt_mergeElement(t *testing.T) {
	myFakeGeorgia := sampleFakeGeorgia2Ptr
	myFakeGeorgia.ID = "georgia-us"
	type args struct {
		element *Element
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "adding a new element",
			args: args{
				element: &sampleFakeGeorgiaPtr,
			},
			wantErr: false,
		},
		{
			name: "adding a same element",
			args: args{
				element: &sampleAfricaElementPtr,
			},
			wantErr: false,
		},
		{
			name: "adding a meta element",
			args: args{
				element: &sampleUsElementPtr,
			},
			wantErr: false,
		},
		{
			name: "adding a same element parent differ",
			args: args{
				element: &myFakeGeorgia,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Ext{
				DefaultExtender: new(gocrawl.DefaultExtender),
				Elements:        sampleElementValidPtr,
			}
			if err := e.mergeElement(tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("Ext.mergeElement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Benchmark_Element_addHash_noHash(b *testing.B) {
	sampleElement := Element{
		ID:      "test",
		Formats: []string{"osm.pbf"},
	}
	sampleHTML := `<p><a href="example.osm.pbf" >example.osm.pbf</a></p>`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTML))
	if err != nil {
		b.Error(err)
	}
	sampleSelection := doc.Find("p")
	for n := 0; n < b.N; n++ {
		sampleElement.addHash(sampleSelection)
	}
}
func Benchmark_Element_addHash_Hash(b *testing.B) {
	sampleElement := Element{
		ID:      "test",
		Formats: []string{"osm.pbf"},
	}
	sampleHTML := `<p><a href="example.osm.pbf" >example.osm.pbf</a><a href="example.osm.pbf.md5" >example.osm.pbf.md5</a></p>`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTML))
	if err != nil {
		b.Error(err)
	}
	sampleSelection := doc.Find("p")
	for n := 0; n < b.N; n++ {
		sampleElement.addHash(sampleSelection)
	}
}

func TestElement_addHash(t *testing.T) {
	sampleElement := Element{
		ID:      "test",
		Formats: []string{"osm.pbf"},
	}
	sampleElementWithHash := sampleElement
	sampleElementWithHash.Formats = append(sampleElementWithHash.Formats, "osm.pbf.md5")
	sampleHTMLWithHash := "<p><a href=\"example.osm.pbf\">example.osm.pbf</a> and <a href=\"example.osm.pbf.md5\">example.osm.pbf.md5</a></p>"
	docWithHash, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTMLWithHash))
	if err != nil {
		t.Error(err)
	}
	sampleSelectionWithHash := docWithHash.Find("p")
	sampleHTMLWithoutHash := `<p><a href="example.osm.pbf" >example.osm.pbf</a></p>`
	docWithoutHash, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTMLWithoutHash))
	if err != nil {
		t.Error(err)
	}
	sampleSelectionWithoutHash := docWithoutHash.Find("p")
	sampleHTMLWithWrongHash := "<p><a href=\"example.osm.pbf\">example.osm.pbf</a> and <a href=\"example.osm.pbf.sha\">example.osm.pbf.sha</a></p>"
	docWithWrongHash, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTMLWithWrongHash))
	if err != nil {
		t.Error(err)
	}
	sampleSelectionWithWrongHash := docWithWrongHash.Find("p")

	type args struct {
		myel *goquery.Selection
	}
	tests := []struct {
		name        string
		fields      Element
		args        args
		wantElement Element
	}{
		// TODO: Add test cases.
		{
			name:        "Hash found",
			fields:      sampleElement,
			args:        args{myel: sampleSelectionWithHash},
			wantElement: sampleElementWithHash,
		}, {
			name:        "NoHash found",
			fields:      sampleElement,
			args:        args{myel: sampleSelectionWithoutHash},
			wantElement: sampleElement,
		}, {
			name:        "Wrong Hash found",
			fields:      sampleElement,
			args:        args{myel: sampleSelectionWithWrongHash},
			wantElement: sampleElement,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Element{
				ID:      tt.fields.ID,
				File:    tt.fields.File,
				Meta:    tt.fields.Meta,
				Name:    tt.fields.Name,
				Formats: tt.fields.Formats,
				Parent:  tt.fields.Parent,
			}

			e.addHash(tt.args.myel)
			if !reflect.DeepEqual(*e, tt.wantElement) {
				t.Errorf("addHash(%v) e=%v, want %v", tt.args.myel, *e, tt.wantElement)
			}
		})
	}
}

func TestExt_parseOSMfr(t *testing.T) {
	var f func(s string, myUrl string) *goquery.Document
	f = func(s string, myUrl string) *goquery.Document {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
		doc.Url, err = url.Parse(myUrl)
		if err != nil {
			t.Error(err)
		}
		return doc
	}
	type args struct {
		ctx *gocrawl.URLContext // not used
		res *http.Response      // not used
		doc *goquery.Document
	}
	tests := []struct {
		name   string
		fields Ext
		args   args
		want   ElementSlice
		want1  bool
	}{
		// TODO: Add test cases.
		{
			name:   "osm valid1",
			fields: Ext{Elements: ElementSlice{}},
			args:   args{doc: f(osmfrValidHTML1, "http://download.openstreetmap.fr/extracts/")},
			want1:  true,
			want: ElementSlice{
				"africa":          {ID: "africa", Meta: false, Name: "africa", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"asia":            {ID: "asia", Meta: false, Name: "asia", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"central-america": {ID: "central-america", Meta: false, Name: "central-america", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"europe":          {ID: "europe", Meta: false, Name: "europe", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"oceania":         {ID: "oceania", Meta: false, Name: "oceania", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"south-america":   {ID: "south-america", Meta: false, Name: "south-america", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
			},
		},
		{
			name:   "osm valid2",
			fields: Ext{Elements: ElementSlice{}},
			args:   args{doc: f(osmfrValidHTML2, "http://download.openstreetmap.fr/extracts/merge/")},
			want: ElementSlice{
				"france_metro_dom_com_nc": {ID: "france_metro_dom_com_nc", File: "", Meta: false, Name: "france_metro_dom_com_nc", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"france_taaf":             {ID: "france_taaf", File: "", Meta: false, Name: "france_taaf", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"israel_and_palestine":    {ID: "israel_and_palestine", File: "", Meta: false, Name: "israel_and_palestine", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"kiribati":                {ID: "kiribati", File: "", Meta: false, Name: "kiribati", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"fiji":                    {ID: "fiji", File: "", Meta: false, Name: "fiji", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
			},
			want1: true,
		}, {
			name:   "osm Poly valid",
			fields: Ext{Elements: ElementSlice{}},
			args:   args{doc: f(osmfrPolygonJapanValidHTML, "http://download.openstreetmap.fr/polygons/asia/japan/")},
			want1:  true,
			want: ElementSlice{
				"chubu":    {ID: "chubu", Meta: false, Name: "chubu", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"chugoku":  {ID: "chugoku", Meta: false, Name: "chugoku", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"hokkaido": {ID: "hokkaido", Meta: false, Name: "hokkaido", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"kansai":   {ID: "kansai", Meta: false, Name: "kansai", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"kanto":    {ID: "kanto", Meta: false, Name: "kanto", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"kyushu":   {ID: "kyushu", Meta: false, Name: "kyushu", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"shikoku":  {ID: "shikoku", Meta: false, Name: "shikoku", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"tohoku":   {ID: "tohoku", Meta: false, Name: "tohoku", Formats: []string{"poly"}, Parent: "japan", File: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: tt.fields.DefaultExtender,
				Elements:        tt.fields.Elements,
			}
			_, got1 := e.parseOSMfr(tt.args.ctx, tt.args.res, tt.args.doc)
			if !reflect.DeepEqual(e.Elements, tt.want) {
				t.Errorf("Ext.parseOSMfr() \nExt.Elements= %+v, want %+v", e.Elements, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Ext.parseOSMfr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExt_parseGisLab(t *testing.T) {
	var f func(s string, myUrl string) *goquery.Document
	f = func(s string, myUrl string) *goquery.Document {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
		doc.Url, err = url.Parse(myUrl)
		if err != nil {
			t.Error(err)
		}
		return doc
	}
	type args struct {
		ctx *gocrawl.URLContext // not used
		res *http.Response      // not used
		doc *goquery.Document
	}
	tests := []struct {
		name   string
		fields Ext
		args   args
		want   ElementSlice
		want1  bool
	}{
		// TODO: Add test cases.
		{
			name:   "osm valid1",
			fields: Ext{Elements: ElementSlice{}},
			args:   args{doc: f(gislabSampleHTML, "http://be.gis-lab.info/project/osm_dump/iframe.php")},
			want1:  true,
			want: ElementSlice{
				"local": {ID: "local", Name: "локальное покрытие", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"AM":    {ID: "AM", Name: "Армения", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"AZ":    {ID: "AZ", Name: "Азербайджан", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"BY":    {ID: "BY", Name: "Беларусь", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"EE":    {ID: "EE", Name: "Эстония", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"TJ":    {ID: "TJ", Name: "Таджикистан", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"TM":    {ID: "TM", Name: "Туркмения", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"UA":    {ID: "UA", Name: "Украина", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"UZ":    {ID: "UZ", Name: "Узбекистан", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"RU-AD": {ID: "RU-AD", Name: "Адыгея", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"RU-AL": {ID: "RU-AL", Name: "Алтай", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"RU":    {ID: "RU", Name: "Российская Федерация", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: tt.fields.DefaultExtender,
				Elements:        tt.fields.Elements,
			}
			_, got1 := e.parseGisLab(tt.args.ctx, tt.args.res, tt.args.doc)
			if !reflect.DeepEqual(e.Elements, tt.want) {
				t.Errorf("Ext.parseGisLab() \nExt.Elements= %+v, want %+v", e.Elements, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Ext.parseGisLab() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Benchmark_Parser_osmfrValidHTML1(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(osmfrValidHTML1))
	doc.Url, err = url.Parse("http://download.openstreetmap.fr/extracts/")
	if err != nil {
		b.Error(err)
	}

	ext := Ext{Elements: ElementSlice{}}

	for n := 0; n < b.N; n++ {
		ext.parseOSMfr(nil, nil, doc)
	}
}
func Benchmark_Parser_osmfrValidHTML2(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(osmfrValidHTML2))
	doc.Url, err = url.Parse("http://download.openstreetmap.fr/extracts/merge/")
	if err != nil {
		b.Error(err)
	}

	ext := Ext{Elements: ElementSlice{}}

	for n := 0; n < b.N; n++ {
		ext.parseOSMfr(nil, nil, doc)
	}
}

func Benchmark_Parser_osmfrPolygonJapanValidHTML(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(osmfrPolygonJapanValidHTML))
	doc.Url, err = url.Parse("http://download.openstreetmap.fr/polygons/asia/japan/")
	if err != nil {
		b.Error(err)
	}

	ext := Ext{Elements: ElementSlice{}}

	for n := 0; n < b.N; n++ {
		ext.parseOSMfr(nil, nil, doc)
	}
}

func TestExt_parseGeofabrik(t *testing.T) {
	var f func(s string) *goquery.Document //, myUrl string) *goquery.Document
	f = func(s string) *goquery.Document { //, myUrl string) *goquery.Document {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
		//	doc.Url, err = url.Parse(myUrl) // not needed
		if err != nil {
			t.Error(err)
		}
		return doc
	}
	type args struct {
		ctx *gocrawl.URLContext // not used
		res *http.Response      // not used
		doc *goquery.Document
	}
	tests := []struct {
		name  string
		args  args         // Contain tests
		want  ElementSlice // to compare with e.Elements
		want1 bool         // always true
	}{
		// TODO: Add test cases.
		{
			name:  "Parse Geofabrik SouthAmerica",
			args:  args{doc: f(geofabrikSouthAmericaHTML)},
			want1: true,
			want: ElementSlice{
				"us":            {ID: "us", File: "", Meta: true, Name: "United States of America", Formats: []string{}, Parent: "north-america"},
				"south-america": {ID: "south-america", File: "", Meta: false, Name: "South America", Formats: []string{"osm.pbf", "osm.pbf.md5", "osm.bz2", "osm.bz2.md5", "osh.pbf", "osh.pbf.md5", "poly", "kml", "state"}, Parent: ""},
			},
		},
		{
			name:  "Parse Geofabrik District of Columbia",
			args:  args{doc: f(geofabrikDistrictOfColumbiaHTML)},
			want1: true,
			want: ElementSlice{
				"us": {ID: "us", File: "", Meta: true, Name: "United States of America", Formats: []string{}, Parent: "north-america"},
				"district-of-columbia": {ID: "district-of-columbia", File: "", Meta: false, Name: "District of Columbia", Formats: []string{"osm.pbf", "osm.pbf.md5", "shp.zip", "osm.bz2", "osm.bz2.md5", "osh.pbf", "osh.pbf.md5", "poly", "kml", "state"}, Parent: "us"},
			},
		},
		{
			name:  "Parse Geofabrik Shikoku",
			args:  args{doc: f(geofabrikShikokuHTML)},
			want1: true,
			want: ElementSlice{
				"us":      {ID: "us", File: "", Meta: true, Name: "United States of America", Formats: []string{}, Parent: "north-america"},
				"shikoku": {ID: "shikoku", File: "", Meta: false, Name: "Shikoku", Formats: []string{"osm.pbf", "osm.pbf.md5", "shp.zip", "osm.bz2", "osm.bz2.md5", "osh.pbf", "osh.pbf.md5", "poly", "kml", "state"}, Parent: "japan"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: new(gocrawl.DefaultExtender),
				Elements:        ElementSlice{},
			}
			e.parseGeofabrik(tt.args.ctx, tt.args.res, tt.args.doc)

			if !reflect.DeepEqual(e.Elements, tt.want) {
				t.Errorf("Ext.parseGeofabrik() got = %+v, want %+v", e.Elements, tt.want)
			}
			/* // Since it's always true, not need to check it!
			if got1 != tt.want1 {
				t.Errorf("Ext.parseGeofabrik() got1 = %v, want %v", got1, tt.want1)
			}
			*/
		})
	}
}

func Benchmark_Parser_geofabrikSouthAmericaHTML(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(geofabrikSouthAmericaHTML))
	if err != nil {
		b.Error(err)
	}
	ext := Ext{Elements: ElementSlice{}}
	for n := 0; n < b.N; n++ {
		ext.parseGeofabrik(nil, nil, doc)
	}
}

func Benchmark_Parser_geofabrikDistrictOfColumbiaHTML(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(geofabrikDistrictOfColumbiaHTML))
	if err != nil {
		b.Error(err)
	}
	ext := Ext{Elements: ElementSlice{}}
	for n := 0; n < b.N; n++ {
		ext.parseGeofabrik(nil, nil, doc)
	}
}

func Benchmark_Parser_geofabrikShikokuHTML(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(geofabrikShikokuHTML))
	if err != nil {
		b.Error(err)
	}
	ext := Ext{Elements: ElementSlice{}}
	for n := 0; n < b.N; n++ {
		ext.parseGeofabrik(nil, nil, doc)
	}
}
