{
  # This value controls the default version of the bky-as cli that is setup in
  # the development shell. This value should be updated as the version of
  # the bky-as cli the example test works against is updated.
  #
  # This default value can be overwritten from the command line by using a valid
  # semver tag to grab a stable version, a git commit to grab a specific unstable
  # version, or "latest" to grab the latest unstable version, e.g.
  #   `nix-shell --argstr bkyas_version v0.1.0-beta.5`
  #   `nix-shell --argstr bkyas_version <full git commit sha>`
  #   `nix-shell --argstr bkyas_version latest`
  # or use the default value by omitting the argument, e.g.
  #   `nix-shell`
  bkyas_version ? "e7a2c061da429c66dcfadaf6007e0a4161fea6dc",
}:
let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-24.11";
  pkgs = import nixpkgs {
    config = { };
    overlays = [ ];
  };

  mkDevShell = import ./nix/mkDevShell.nix;
in
mkDevShell {
  pkgs = pkgs;

  version = bkyas_version;

  devDependencies = [
    pkgs.git # for project management
    pkgs.gnumake # for project management
    pkgs.go_1_22 # for prepping for building wasm
    pkgs.golangci-lint # for linting go files
    pkgs.easyjson # for generating json marshaling go code
    pkgs.gotools # for tools like goimports
    pkgs.jq # for processing data in examples
    pkgs.nixfmt-rfc-style # for formatting nix files
    pkgs.tinygo # for building wasm
  ];
}
