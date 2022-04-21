# Variables

## Which Go binary to use

GO?=go


# Development setup

.PHONY: bindl
bindl:
	curl --location https://bindl.dev/bootstrap.sh | OUTDIR=bin bash


# Test

.PHONY: test
test:
	${GO} test -race -v ./...


# Linters

.PHONY: lint
lint: bin/golangci-lint
	bin/golangci-lint run

.PHONY: lint/fix
lint/fix: bin/golangci-lint
	bin/golangci-lint run --fix


# GitHub Actions

.PHONY: gh/lint
gh/lint: bin/golangci-lint
	bin/golangci-lint run --out-format github-actions


# Bindl programs
#
include Makefile.bindl
