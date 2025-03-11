lint:
	golangci-lint run --config golangci.yaml

pre-pr: lint
	$(MAKE) -C ./example run

generate:
	easyjson --all dto.go
	easyjson x/xbasm/parse.go
