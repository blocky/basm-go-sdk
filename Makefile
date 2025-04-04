lint:
	golangci-lint run --config golangci.yaml

pre-pr: lint
	$(MAKE) -C ./example clean run

generate:
	easyjson --all basm/dto.go
	easyjson x/xbasm/parse.go
