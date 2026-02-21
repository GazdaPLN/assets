#! /usr/bin/make -f

# Zmienne projektu.
VERSION := $(shell git describe --tags 2>/dev/null || git describe --all)
BUILD := $(shell git rev-parse --short HEAD)
PROJECT_NAME := $(shell basename "$(PWD)")
BUILD_TARGETS := $(shell find cmd -name \*main.go | awk -F'/' '{print $$0}')

# Użyj flag linkera, aby podać ustawienia wersji/kompilacji
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Make jest gadatliwy w Linux. Ustaw tryb cichy.
MAKEFLAGS += --silent

# Pliki Go.
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

# Podstawowe polecenia.
all: fmt lint test

build:
	@echo "  >  Budowanie main.go do bin/assets"
	go build $(LDFLAGS) -o bin/assets ./cmd

test:
	@echo "  >  Uruchamianie testów jednostkowych"
	go test -cover -race -coverprofile=coverage.txt -covermode=atomic -v ./...

fmt:
	@echo "  >  Formatowanie wszystkich plików go"
	gofmt -w ${GOFMT_FILES}

lint-install:
ifeq (,$(wildcard test -f bin/golangci-lint))
	@echo "  >  Instalowanie golangci-lint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- v1.50.1
endif

lint: lint-install
	@echo "  >  Uruchamianie golangci-lint"
	bin/golangci-lint run --timeout=2m

# Polecenia Assets.
check: build
	bin/assets check

fix: build
	bin/assets fix

update-auto: build
	bin/assets update-auto

# Polecenia pomocnicze.
add-token: build
	bin/assets add-token $(asset_id)

add-tokenlist: build
	bin/assets add-tokenlist $(asset_id)

add-tokenlist-extended: build
	bin/assets add-tokenlist-extended $(asset_id)
