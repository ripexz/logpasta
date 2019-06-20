.PHONY: build clean

SHELL=/bin/bash
MAKE=/usr/bin/make
platforms="windows/amd64" "windows/386" "darwin/amd64" "darwin/386" "linux/amd64" "linux/386"

build: build-windows-amd64 build-windows-386 build-darwin-amd64 build-darwin-386 build-linux-amd64 build-linux-386

clean:
	rm -f logpasta-*

build-windows-amd64:
	GOOS=windows GOARCH=amd64 ./build.sh

build-windows-386:
	GOOS=windows GOARCH=386 ./build.sh

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 ./build.sh

build-darwin-386:
	GOOS=darwin GOARCH=386 ./build.sh

build-linux-amd64:
	GOOS=linux GOARCH=amd64 ./build.sh

build-linux-386:
	GOOS=linux GOARCH=386 ./build.sh