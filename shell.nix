{pkgs ? import <nixpkgs> {}}:
pkgs.stdenv.mkDerivation {
  name = "wrap-shell";

  buildInputs = with pkgs; [
    alejandra
    go
    dep
  ];

  shellHook = ''
    unset GOPATH
 '';
}
