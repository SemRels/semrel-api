# SPDX-License-Identifier: Apache-2.0
# SPDX-FileCopyrightText: 2026 The SemRels Authors

.PHONY: generate build test lint tidy

generate:
	buf generate

build:
	go build ./...

test:
	go test -race ./...

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

.DEFAULT_GOAL := build
