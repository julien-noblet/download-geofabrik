default: clean all
clean:
	go clean
	rm -rf download-geofabrik_* *.zip
gox:
	gox --output="download-geofabrik_{{.OS}}_{{.Arch}}/{{.Dir}}"
	cd generator && gox --output="../download-geofabrik_{{.OS}}_{{.Arch}}/{{.Dir}}"
geofabrik:
	echo "Generating geofabrik.yml"
	go run generator/generator.go
readme:
	cat .README.md1 > README.md
	go run download-geofabrik.go --help-long >> README.md 
	cat .README.md2 >> README.md
	go run download-geofabrik.go list --markdown >> README.md 
package: gox geofabrik
	for i in download-geofabrik_* ;\
	do \
		  echo "Compressing $$i";\
          cp CHANGELOG.md README.md LICENSE geofabrik.yml $$i/ && cd $$i && zip -9 $$i.zip download-geofabrik* generator* geofabrik.yml LICENSE README.md CHANGELOG.md && mv $$i.zip ../ && cd ..;\
        done

all: package
