lint:
	golangci-lint run --config golangci.yaml

pre-pr: lint
