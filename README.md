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
example in the `example` directory. The go.mod file in the example directory
has been set up to use the local version of the sdk.

When iterating on changes to the SDK, make sure to run the example to verify
the changes work.

```bash
make run
```

Verify that the output is as expected.
