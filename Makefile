#  _____ ____  __  __
# |  ___/ ___||  \/  |
# | |_  \___ \| |\/| |
# |  _|  ___) | |  | |
# |_|   |____/|_|  |_|
#
# Finite State Machine
# (c) 2018 Adam K Dean

#
# Variables
#
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

SRC ?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKGS = $(shell go list ./... | grep -v /vendor)
BUILD ?= $(shell git rev-parse --short HEAD)
TARGET = bin/$(BUILD)
ENTRYPOINT = cmd/main.go

#
# Rules
#
.DEFAULT_GOAL: $(TARGET)
.PHONY: build clean run hooks fmt test lint

$(TARGET): $(SRC)
	$(GOBUILD) -o $(TARGET) $(ENTRYPOINT)

build: $(TARGET)
	@true

clean:
	@$(GOCLEAN)

run: build
	@$(TARGET)

hooks: .git/hooks/pre-commit

.git/hooks/pre-commit: scripts/hooks/pre-commit.sh
	@cp -f scripts/hooks/pre-commit.sh .git/hooks/pre-commit

test: lint
	$(GOTEST) $(PKGS)

lint: $(GOMETALINTER)
	$(GOMETALINTER) ./... --vendor --fast --disable=maligned

$(GOMETALINTER):
	$(GOGET) -u github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install 1>/dev/null

fmt:
	gofmt -l -w $(SRC)
