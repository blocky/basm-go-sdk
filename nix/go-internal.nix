{
  pkgs,
}:
pkgs.buildGoModule rec {
  pname = "go-internal";
  version = "1.14.1";

  src = pkgs.fetchFromGitHub {
    owner = "rogpeppe";
    repo = "go-internal";
    rev = "v${version}";
    hash = "sha256-6NzhXCCD1Qhj05WHbCDxH5hwfNM6psoAk7uIxm7N55E=";
  };

  vendorHash = "sha256-WoRmZbYYpwVVetlxJDjUu9jGgwLXUD3/PnUF6ksUT70=";

  doCheck = false;

  meta = {
    description = "The go-internal package includes various packages used by the Go team, such as testscript, txtar, and more.";
    homepage = "https://github.com/rogpeppe/go-internal";
    license = pkgs.lib.licenses.bsd3;
  };
}
