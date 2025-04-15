{
  pkgs,
  bkyAsVersion,
  devDependencies,
}:
let
  system = import ./system.nix { pkgs = pkgs; };
  goos = system.goos;
  goarch = system.goarch;

  isCommit = x: builtins.match "^[0-9a-f]{40}$" x != null;

  bkyAsStable = pkgs.stdenv.mkDerivation {
    pname = "bky-as";
    version = bkyAsVersion;
    src = builtins.fetchurl {
      url = "https://github.com/blocky/attestation-service-demo/releases/download/${bkyAsVersion}/bky-as_${goos}_${goarch}";
    };
    unpackPhase = ":";
    installPhase = ''
      install -D -m 555 $src $out/bin/bky-as
    '';
  };

  stableShell = pkgs.mkShell {
    packages = devDependencies ++ [ bkyAsStable ];
    shellHook = ''
      echo "Stable bky-as version: ${bkyAsVersion}"
    '';
  };

  bkyAsUnstable = pkgs.stdenv.mkDerivation {
    pname = "bky-as";
    version = bkyAsVersion;
    src = ./fetch-bky-as.sh;
    unpackPhase = ":";
    installPhase = ''
      install -D -m 555 $src $out/bin/fetch-bky-as.sh
    '';
  };

  unstableShell = pkgs.mkShell {
    packages =
      devDependencies
      ++ [ bkyAsUnstable ]
      #dependencies required by the fetch script
      ++ [
        pkgs.gh
        pkgs.awscli2
        pkgs.jq
      ];
    shellHook = ''
      bin=$(pwd)/tmp/bin
      fetch-bky-as.sh $bin ${bkyAsVersion} ${goos} ${goarch}
      export PATH=$bin:$PATH
      echo "Unstable bky-as version: ${bkyAsVersion}"
    '';
  };
in
if isCommit bkyAsVersion || bkyAsVersion == "latest" then
  unstableShell
else
  # If the version is not a commit hash or "latest", we assume it is a stable
  # release version.
  stableShell
