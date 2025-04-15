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
- [`testscript`](https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript)
  - Used for running executable and filesystem tests.

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

Note that any command can be run using the shell via the `--run` flag:

```bash
nix-shell --pure --run "make pre-pr"
```bash

#### Testing

SDK examples are compiled to wasm and tested against the Blocky Attestation
Service using the `bky-as` CLI. The tests are run using the [`testscript`
CLI tool](https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript).

Run the test scripts with:

```bash
make test
```
