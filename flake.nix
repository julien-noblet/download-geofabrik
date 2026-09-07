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
        upx = pkgs.upx.overrideAttrs (oldAttrs: rec {
          version = "5.2.1";
          src = pkgs.fetchFromGitHub {
            owner = "upx";
            repo = "upx";
            tag = "v${version}";
            fetchSubmodules = true;
            hash = "sha256-Fy+BqntQcHLo5FGhWTP6gqe5OIGC6w7SMc+tlV+VqCM=";
          };
        });
      in
      {
        packages = {
          default = buildGoModule {
            pname = "download-geofabrik";
            version = "unstable";
            src = ./.;
            vendorHash = "sha256-KjGBhl8ZtD0TT0n/nrl186NREsiCYgqCcOzWOOJhRok=";
            subPackages = [ "cmd/download-geofabrik" ];
            env.CGO_ENABLED = 0;

            ldflags = [
              "-s"
              "-w"
            ];
          };
          upx = upx;
        };

        apps = {
          default = flake-utils.lib.mkApp {
            drv = self.packages.${system}.default;
          };
          upx = flake-utils.lib.mkApp {
            drv = upx;
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
            upx
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
