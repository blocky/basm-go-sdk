lint:
	golangci-lint run --config golangci.yaml

pre-pr: lint

easyjson:
	easyjson --all dto.go
