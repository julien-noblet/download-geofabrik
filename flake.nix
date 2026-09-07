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
        go1_27_1 = pkgs.symlinkJoin {
          name = "go-1.27.1";
          paths = [ go ];
          postBuild = ''
            ln -s $out/bin/go $out/bin/go1.27.1
          '';
        };
        buildGoModule = pkgs.buildGo127Module;
      in
      {
        packages = {
          default = buildGoModule {
            pname = "download-geofabrik";
            version = "unstable";
            src = ./.;
            vendorHash = "sha256-08yBURftwCHqlZKnu6ek3mbkbTRF007iLgNvP9Xhf1g=";
            subPackages = [ "cmd/download-geofabrik" ];
            env.CGO_ENABLED = 0;

            ldflags = [
              "-s"
              "-w"
            ];
          };
        };

        apps = {
          default = flake-utils.lib.mkApp {
            drv = self.packages.${system}.default;
          };
        };

        devShells.default = pkgs.mkShell {
          packages = [
            go1_27_1
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
            export CGO_ENABLED=0
          '';
        };
      }
    );
}
