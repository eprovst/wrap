{
  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

  outputs =
    { self, nixpkgs }:
    let
      # Dear reader, for now I only have access to a Linux machine
      supportedSystems = [
        "x86_64-linux"
      ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    in
    {
      devShells = forAllSystems (
        system:
        let
          pkgs = import nixpkgs { inherit system; };
        in
        {
          default = pkgs.mkShell {
            packages = with pkgs; [
              go
              gnumake

              nsis
              (buildGoModule rec {
                pname = "goversioninfo";
                version = "1.7.0";
                subPackages = [ "cmd/goversioninfo" ];
                src = pkgs.fetchFromGitHub {
                  owner = "josephspurrier";
                  repo = "goversioninfo";
                  rev = "v${version}";
                  sha256 = "sha256-P6e1hu9DY6gh7SBYIw2o5qLl9LD3u4FGdjxKmqtNGmg=";
                };
                vendorHash = "sha256-cbx+iq7uZWEfjyZ8wtEEso3S/3wTf/0o932YIaUvA0I=";
              })
            ];
          };
        }
      );
    };
}
