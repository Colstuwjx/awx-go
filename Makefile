GO           := GO15VENDOREXPERIMENT=1 go
FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
pkgs          = $(shell $(GO) list ./... | grep -v /vendor/ )

BUILD_ENTRY             ?= ./*.go
BINARY                  ?= awx-go

all: format build test

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

install-test-reqs:
	@$(GO) get -u github.com/kylelemons/godebug/pretty
	@$(GO) get -u github.com/kylelemons/godebug/diff

test-short: install-test-reqs
	@echo ">> running short tests"
	@$(GO) test -short $(pkgs)

test: install-test-reqs
	@echo ">> running all tests"
	@$(GO) test $(pkgs)

lint: install-test-reqs
	@gometalinter --vendor --disable-all --enable=gosimple --enable=golint --enable=vet --enable=ineffassign --enable=unconvert \
    --exclude="by other packages, and that stutters; consider calling this" \
    --skip=proto \
    --skip=vendor \
    --skip=.git \
    --tests ./...

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
