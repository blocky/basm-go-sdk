{
  # This value controls the default version of the bky-as cli that is setup in
  # the development shell. This value should be updated as the version of
  # the bky-as cli the example test works against is updated.
  #
  # This default value can be overwritten from the command line.
  bkyAsVersion ? "e7a2c061da429c66dcfadaf6007e0a4161fea6dc",
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

  bkyAsVersion = bkyAsVersion;

  # Note that these package versions are determined by the nixpkgs version
  # version used to build the shell.
  devDependencies = [
    pkgs.git # for project management
    pkgs.less # for viewing files
    pkgs.gnumake # for project management
    pkgs.go # for prepping for building wasm, golangci-lint, and easyjson
    pkgs.golangci-lint # for linting go files
    pkgs.easyjson # for generating json marshaling go code
    pkgs.gotools # for tools like goimports
    pkgs.jq # for processing data in examples
    pkgs.nixfmt-rfc-style # for formatting nix files
    pkgs.tinygo # for building wasm
    pkgs.toybox # include common unix commands for convenience
  ];
}
