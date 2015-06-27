clean:
	go clean
	rm -rf darwin32 darwin64 linux32 linux64 win32 win64 *.zip

darwin32:
	mkdir darwin32
	cp geofabrik.yml LICENSE README.md darwin32/
	GOOS=darwin GOARCH=386 go build -o darwin32/download-geofabrik

darwin64:
	mkdir darwin64
	cp geofabrik.yml LICENSE README.md darwin64/
	GOOS=darwin GOARCH=amd64 go build -o darwin64/download-geofabrik

linux32:
	mkdir linux32
	cp geofabrik.yml LICENSE README.md linux32/
	GOOS=linux GOARCH=386 go build -o linux32/download-geofabrik

linux64:
	mkdir linux64
	cp geofabrik.yml LICENSE README.md linux64/
	GOOS=linux GOARCH=amd64 go build -o linux64/download-geofabrik

win32:
	mkdir win32
	cp geofabrik.yml LICENSE README.md win32/
	GOOS=windows GOARCH=386 go build -o win32/download-geofabrik.exe

win64:
	mkdir win64
	cp geofabrik.yml LICENSE README.md win64/
	GOOS=windows GOARCH=amd64 go build -o win64/download-geofabrik.exe

package-darwin32: darwin32
	cd darwin32 && zip -9 download-geofabrik-darwin32.zip download-geofabrik geofabrik.yml LICENSE README.md
	mv darwin32/download-geofabrik-darwin32.zip ./

package-darwin64: darwin64
	cd darwin64 && zip -9 download-geofabrik-darwin64.zip download-geofabrik geofabrik.yml LICENSE README.md
	mv darwin64/download-geofabrik-darwin64.zip ./

package-linux32: linux32
	cd linux32 && zip -9 download-geofabrik-linux32.zip download-geofabrik geofabrik.yml LICENSE README.md
	mv linux32/download-geofabrik-linux32.zip ./

package-linux64: linux64
	cd linux64 && zip -9 download-geofabrik-linux64.zip download-geofabrik geofabrik.yml LICENSE README.md
	mv linux64/download-geofabrik-linux64.zip ./

package-win32: win32
	cd win32 && zip -9 download-geofabrik-win32.zip download-geofabrik.exe geofabrik.yml LICENSE README.md
	mv win32/download-geofabrik-win32.zip ./

package-win64: win64
	cd win64 && zip -9 download-geofabrik-win64.zip download-geofabrik.exe geofabrik.yml LICENSE README.md
	mv win64/download-geofabrik-win64.zip ./

all: package-darwin32 package-darwin64 package-linux32 package-linux64 package-win32 package-win64
