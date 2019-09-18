VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECT_NAME := $(shell basename "$(PWD)")
API_COMMAND := api
OBSERVER_COMMAND := observer

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOPKG := $(cmd)

# Go files
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECT_NAME)-stderr.txt

# PID file will keep the process id of the server
PID_API := /tmp/.$(PROJECT_NAME).$(API_COMMAND).pid
PID_OBSERVER := /tmp/.$(PROJECT_NAME).$(OBSERVER_COMMAND).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## start: Start API and Observer in development mode.
start:
	@bash -c "$(MAKE) clean compile start-api start-observer"

## stop: Stop development mode.
stop: stop-server

## start-api: Start API in development mode.
start-api: stop-server
	@echo "  >  Starting $(PROJECT_NAME) API"
	@-$(GOBIN)/$(PROJECT_NAME) $(API_COMMAND) 2>&1 & echo $$! > $(PID_API)
	@cat $(PID_API) | sed "/^/s/^/  \>  API PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-observer: Start Observer in development mode.
start-observer: stop-server
	@echo "  >  Starting $(PROJECT_NAME) Observer"
	@-$(GOBIN)/$(PROJECT_NAME) $(OBSERVER_COMMAND) 2>&1 & echo $$! > $(PID_OBSERVER)
	@cat $(PID_OBSERVER) | sed "/^/s/^/  \>  Observer PID: /"
	@echo "  >  Error log: $(STDERR)"

stop-server:
	@-touch $(PID_API) $(PID_OBSERVER)
	@-kill `cat $(PID_API)` 2> /dev/null || true
	@-kill `cat $(PID_OBSERVER)` 2> /dev/null || true
	@-rm $(PID_API) $(PID_OBSERVER)

restart-server: stop-server start-server

## compile: Compile the binary.
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

## exec: Run given command. e.g; make exec run="go test ./..."
exec:
	GOBIN=$(GOBIN) $(run)

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-rm $(GOBIN)/$(PROJECT_NAME) 2> /dev/null
	@-$(MAKE) go-clean

## test: Run all unit tests.
test: go-test

## fmt: Run `go fmt` for all go files.
fmt: go-fmt

go-compile: go-get go-build

go-build:
	@echo "  >  Building binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECT_NAME) -v ./cmd/

go-generate:
	@echo "  >  Generating dependency files..."
	GOBIN=$(GOBIN) go generate $(generate)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	GOBIN=$(GOBIN) go get $(get)

go-install:
	GOBIN=$(GOBIN) go install $(GOPKG)

go-clean:
	@echo "  >  Cleaning build cache"
	GOBIN=$(GOBIN) go clean

go-test:
	@echo "  >  Cleaning build cache"
	GOBIN=$(GOBIN) go test -v ./...

go-fmt:
	@echo "  >  Format all go files"
	GOBIN=$(GOBIN) gofmt -w ${GOFMT_FILES}

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo