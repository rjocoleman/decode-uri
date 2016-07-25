all: build

build:
	go build

linux: *.go
	GOOS=linux GOARCH=amd64 go build

clean:
	rm -f decode-uri
