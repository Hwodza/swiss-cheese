{
  description = "Go development environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {inherit system;};
      in rec {
        packages = {
          default = pkgs.buildGoModule {
            pname = "swiss-cheese";
            version = "0.0.0";
            src = ./.;
            vendorHash = null;
          };
        };
        apps.default = flake-utils.lib.mkApp {
          drv = packages.default;
        };
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gotools
            gopls
            golangci-lint
            delve
          ];
        };
      }
    );
}
