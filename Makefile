GO           := GO15VENDOREXPERIMENT=1 go
FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
pkgs          = $(shell $(GO) list ./... | grep -v /vendor/ )

BUILD_ENTRY             ?= ./*
BINARY                  ?= awx-go

all: format build test

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

test-short:
	@echo ">> running short tests"
	@$(GO) test -short $(pkgs)

test:
	@echo ">> running all tests"
	@$(GO) test $(pkgs)

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

build:
	@echo ">> building binaries"
	@$(GO) build -o $(BINARY) $(BUILD_ENTRY)

.PHONY: all style format build test test-short vet