# Go SDK for the Blocky Attestation Service Wasm Runtime (BASM)

This SDK provides functions for interacting with Blocky Attestation Service
WASM runtime.

## Contributing

### Dependencies

- Go 1.23.8
- Tinygo v0.34.0
- golangci-lint
- jq
- [easyjson](https://github.com/mailru/easyjson) v0.9.0
    - Used for generating JSON serialization code

### Development

#### Integration Testing

SDK examples are compiled to wasm and tested against the Blocky Attestation
Service using the `bky-as` CLI. The tests are run using the [`testscript`
library](https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript).

Run the integration tests with:

```bash
make test-integration
```
