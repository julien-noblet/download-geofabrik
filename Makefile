gofiles  = download-geofabrik.go config.go download.go element.go formats.go generator.go bbbike.go geofabrik.go openstreetmap.fr.go scrapper.go
geofabrik:
	echo "Generating geofabrik.yml"
	go run $(gofiles) generate --progress
osmfr:
	echo "Generating openstreetmap.fr.yml"
	go run $(gofiles) --service="openstreetmap.fr" generate --progress
bbbike:
	echo "Generating bbbike.yml"
	go run $(gofiles) --service="bbbike" generate -v
readme: 
	cat .README.md1 > README.md
	go run $(gofiles) --help-long >> README.md 
	cat .README.md2 >> README.md
	go run $(gofiles) list --markdown >> README.md 
	echo "" >> README.md
	echo "## List of elements from openstreetmap.fr" >> README.md
	go run $(gofiles) --service "openstreetmap.fr" list --markdown >> README.md
	echo "" >> README.md
	echo "## List of elements from bbbike.org" >> README.md
	go run $(gofiles) --service "bbbike" list --markdown >> README.md
	echo "" >> README.md

