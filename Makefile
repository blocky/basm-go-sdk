lint:
	golangci-lint run --config golangci.yaml

pre-pr: lint

generate:
	easyjson --all dto.go
	easyjson x/xbasm/parse.go
