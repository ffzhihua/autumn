GOPATH:=$(shell pwd)/vendor:$(shell go env GOPATH)

build:
	go build -o autumn

clean:
	rm -rf autumn

