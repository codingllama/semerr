# Makefile is used to capture non-intuitive commands or often forgotten tasks.
#
# Partly follows
# https://www.gnu.org/software/make/manual/html_node/Standard-Targets.html.

SRCS ?= $(shell find . -name '*.go')

GO ?= go

# go install golang.org/x/tools/cmd/godoc@latest
GODOC ?= godoc

# go install golang.org/x/tools/cmd/goimports@latest
GOIMPORTS ?= goimports

# https://golangci-lint.run/usage/install/#local-installation
#
# or simply:
# go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
GOLANGCI-LINT ?= golangci-lint

# go install github.com/google/addlicense@latest
ADDLICENSE ?= addlicense
ADDLICENSE_FLAGS ?= -c 'Alan Parra' -l mit

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -f cover.out

.PHONY: check
check: test

.PHONY: build
build:
	$(GO) build ./...

.PHONY:
test:
	$(GO) test ./...

.PHONY: gen
gen:
	$(GO) generate ./...

.PHONY: lint
lint: lint/go lint/license

.PHONY: lint/go
lint/go:
	$(GOLANGCI-LINT) run ./...
	cd ./internal/cmd/generate && $(GOLANGCI-LINT) run ./...

.PHONY: lint/license
lint/license:
	$(ADDLICENSE) $(ADDLICENSE_FLAGS) -check .

.PHONY: fix
fix: fix/go fix/license fix/mod

.PHONY: fix/go
fix/go:
	$(GOIMPORTS) -w $(SRCS)

.PHONY: fix/license
fix/license:
	$(ADDLICENSE) $(ADDLICENSE_FLAGS) .

.PHONY: fix/mod
fix/mod:
	$(GO) mod tidy
	cd ./internal/cmd/generate && $(GO) mod tidy

.PHONY: cover
cover: cover/html

.PHONY: cover/html
cover/html:
	$(GO) test ./... -coverprofile=cover.out
	$(GO) tool cover -html=cover.out

.PHONY: cover/text
cover/text:
	$(GO) test ./... -coverprofile=cover.out
	$(GO) tool cover -func=cover.out | grep -Pv '\t0.0%$$'

.PHONY: docs
docs: docs/html

.PHONY: docs/html
docs/html:
	$(GODOC) -http=':6060'

.PHONY: docs/text
docs/text:
	$(GO) doc -all .

# Disable builtin rules.
.SUFFIXES:
