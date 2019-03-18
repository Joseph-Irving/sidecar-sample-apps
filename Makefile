ARCH = amd64
BIN  = bin/sample
BIN_LINUX  = $(BIN)-linux-$(ARCH)
BIN_DARWIN = $(BIN)-darwin-$(ARCH)

SOURCES := $(shell find . -iname '*.go')

.PHONY: test clean all

all: build-darwin build-linux

build-darwin: $(SOURCES)
	GOARCH=$(ARCH) GOOS=darwin go build -o $(BIN_DARWIN) cmd/*.go

build-linux: $(SOURCES)
	GOARCH=$(ARCH) GOOS=linux CGO_ENABLED=0 go build -o $(BIN_LINUX) cmd/*.go

test: $(SOURCES)
	go test -v -cover $(shell go list ./... | grep -v /vendor)

bench: $(SOURCES)
	go test -run=XX -bench=. $(shell go list ./... | grep -v /vendor)

docker: Dockerfile $(BIN_LINUX)
	docker image build -t quay.io/joseph_irving/sidecar-sample-app:devel .

clean:
	rm -rf bin/
