REPO := $(shell git rev-parse --show-toplevel)

build:
	GOBIN=$(REPO)/bin go install ./...

all: build

.PHONY: all build

