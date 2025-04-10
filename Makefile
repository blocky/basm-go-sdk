SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
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

pre-pr: lint generate
	$(MAKE) clean
	$(MAKE) -C ./example run

.PHONY: clean
clean:
	# - ignore errors due to missing files
	-$(MAKE) -C ./example clean
