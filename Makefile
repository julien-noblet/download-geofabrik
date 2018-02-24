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
package: gox geofabrik
	for i in download-geofabrik_* ;\
	do \
		  echo "Compressing $$i";\
          cp README.md LICENSE geofabrik.yml $$i/ && cd $$i && zip -9 $$i.zip download-geofabrik* generator* geofabrik.yml LICENSE README.md && mv $$i.zip ../ && cd ..;\
        done

all: package
