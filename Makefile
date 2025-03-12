lint:
	golangci-lint run --config golangci.yaml

pre-pr: lint
	$(MAKE) -C ./example clean run

generate:
	easyjson --all dto.go
	easyjson x/xbasm/parse.go
