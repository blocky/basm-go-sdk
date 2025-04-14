MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.DEFAULT_GOAL := run
.DELETE_ON_ERROR:

srcs := $(wildcard ../**/*.go)

x.wasm: main.go $(srcs)
	@echo "Building WASM module..."
	@tinygo build -o $@ -target=wasi $<

.PHONY: build
build: x.wasm

result.json: x.wasm
	@echo "Running WASM module..."
	@cat fn-call.json \
		| BKY_AS_HOST="local-server" bky-as attest-fn-call > $@

.PHONY: run
run: result.json
	@echo "## Output"
	@jq -r '.transitive_attested_function_call.claims.output | @base64d | fromjson ' $<
	@echo ""
	@echo "## Log"
	@jq -r '.transitive_attested_function_call.logs | @base64d' $<

.PHONY: clean
clean:
	# ignore errors due to missing files
	-rm x.wasm result.json
