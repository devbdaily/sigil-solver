.PHONY: all fmt dep test build clean

default: all

all: dep test build

fmt:
	go fmt ./...

dep:
	go mod download
	go mod verify

test:
	go test ./...

build:
	go build .

clean:
	rm sigil-solver
