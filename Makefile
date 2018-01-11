clean:
	go clean
	rm -rf download-geofabrik_* *.zip
gox:
	~/go/bin/gox --output="download-geofabrik_{{.OS}}_{{.Arch}}/{{.Dir}}"
	cd generator && ~/go/bin/gox --output="../download-geofabrik_{{.OS}}_{{.Arch}}/{{.Dir}}"

package: gox
	for i in download-geofabrik_* ;\
	do \
          cp README.md LICENSE geofabrik.yml $$i/ && cd $$i && zip -9 $$i.zip download-geofabrik* generator* geofabrik.yml LICENSE README.md && mv $$i.zip ../ && cd ..;\
        done

all: package
