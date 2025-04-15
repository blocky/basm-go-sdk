# Go SDK for the Blocky Attestation Service Wasm Runtime (BASM)

This SDK provides functions for interacting with Blocky Attestation Service
WASM runtime.

## Contributing

### Dependencies

- Go (see `go.mod` for version)
- Tinygo v0.34.0
- jq
- golangci-lint
- [easyjson](https://github.com/mailru/easyjson) v0.9.0
    - Used for generating JSON serialization code
- `bky-as` - Blocky Attestation Service
    - The SDK is designed to work with the Blocky Attestation Service
    - The version compatible with this SDK is pinned in the `shell.nix` file.

Additional project dependencies are specified in tbe `shell.nix` file.

### Development

### Nix Shell
To enter a development shell with all dependencies, run:

```bash
nix-shell --pure
```

The development shell can be started with a specific version of `bky-as` by
specifying the version via the `--argstr` flag:

```bash
nix-shell --pure --argstr bkyAsVersion v0.1.0-beta.5 # stable version
nix-shell --pure --argstr bkyAsVersion <full git commit sha> # specific unstable version
nix-shell --pure --argstr bkyAsVersion latest # latest unstable version
```

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
