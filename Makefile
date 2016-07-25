all: build

build:
	go build

linux: *.go
	GOOS=linux GOARCH=amd64 go build

clean:
	rm -f decode-uri

dist: linux
	tar -cvzf linux-amd64-decode-uri.tar.gz ./decode-uri 
