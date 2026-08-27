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
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "download-geofabrik";
          version = "unstable";
          src = ./.;
          vendorHash = "sha256-CSMeEBwyQb7broVgnm939baps+3IgR8UqnwXVtVw/lQ=";
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
          packages = with pkgs; [
            go
            gopls
            golangci-lint
            goreleaser
            delve
            gotools
          ];

          shellHook = ''
            export GOPATH="''${GOPATH:-$HOME/go}"
            export PATH="$GOPATH/bin:$PATH"
          '';
        };
      }
    );
}
