# The Go SDK for the Blocky Attestation Service Wasm Runtime

(working name)

This SDK provides functions for interacting with Blocky Attestation Service
WASM runtime.

## Contributing

### Dependencies

- Go 1.22.6
- Tinygo v0.32.0
- golangci-lint
- [easyjson](https://github.com/mailru/easyjson) v0.9.0
    - Used for generating JSON serialization code

### Development

#### Testing

Until a test harness is created, the best way to test the SDK is to use the
example in the `example` directory. If you are working off of a development
branch, in the example directory use `go get` to fetch the version you are
working on:

```bash
go get github.com/blocky/attestation-sdk-go@<branch>
```

Then update the example in `example/main.go` and run:

```bash
make run
```

Verify that the output is as expected.
