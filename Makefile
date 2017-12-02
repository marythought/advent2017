VERSION ?= $(shell git rev-parse --abbrev-ref HEAD)-$(shell git describe --long --always)
BUILD_TIME ?=`date +%FT%T%z`

all: build

compile:
	mkdir -p bin && go build  -ldflags="-X main.BuildTime=${BUILD_TIME} -X main.Version=${VERSION}" -o bin/advent main.go

build: compile

run: build
	./bin/advent

test:
	go test `glide novendor` -cover