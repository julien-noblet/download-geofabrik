{
  description = "download-geofabrik development and build environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        go = pkgs.go_1_27;
        buildGoModule = pkgs.buildGo127Module;
      in
      {
        packages.default = buildGoModule {
          pname = "download-geofabrik";
          version = "unstable";
          src = ./.;
          vendorHash = "sha256-KjGBhl8ZtD0TT0n/nrl186NREsiCYgqCcOzWOOJhRok=";
          subPackages = [ "cmd/download-geofabrik" ];

          ldflags = [
            "-s"
            "-w"
          ];
        };

        apps.default = flake-utils.lib.mkApp {
          drv = self.packages.${system}.default;
        };

        devShells.default = pkgs.mkShell {
          packages = [
            go
            pkgs.gopls
            pkgs.golangci-lint
            pkgs.goreleaser
            pkgs.delve
            pkgs.gotools
          ];

          shellHook = ''
            export SHELL="''${SHELL:-${pkgs.bashInteractive}/bin/bash}"
            export GOPATH="''${GOPATH:-$HOME/go}"
            export PATH="$GOPATH/bin:$PATH"
          '';
        };
      }
    );
}
