MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := generate
.DELETE_ON_ERROR:

# add easyjson source file targets here
easyjson_sources := basm/dto.go x/xbasm/parse.go
easyjson_generated := $(easyjson_sources:.go=_easyjson.go)

# Rule to generate *_easyjson.go files
%_easyjson.go: %.go
	@easyjson $<

# Generate all *_easyjson.go files
generate: $(easyjson_generated)

.PHONY: lint
lint:
	golangci-lint run --config golangci.yaml

.PHONY: pre-pr
pre-pr: lint generate
	$(MAKE) clean
	$(MAKE) -C ./example run

.PHONY: clean
clean:
	$(MAKE) -C ./example clean
