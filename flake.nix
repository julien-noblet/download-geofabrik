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

        download-geofabrik-dev = pkgs.runCommand "download-geofabrik" {
          nativeBuildInputs = [ pkgs.installShellFiles ];
        } ''
          mkdir -p $out/bin
          cat << 'SCRIPT' > $out/bin/download-geofabrik
          #!${pkgs.bash}/bin/bash
          if [ -f "./cmd/download-geofabrik/main.go" ]; then
            exec ${go}/bin/go run ./cmd/download-geofabrik "$@"
          elif command -v git >/dev/null 2>&1 && git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
            REPO_ROOT="$(git rev-parse --show-toplevel 2>/dev/null)"
            if [ -n "$REPO_ROOT" ] && [ -f "$REPO_ROOT/cmd/download-geofabrik/main.go" ]; then
              exec ${go}/bin/go run "$REPO_ROOT/cmd/download-geofabrik" "$@"
            fi
          fi
          exec ${self.packages.${system}.default}/bin/download-geofabrik "$@"
          SCRIPT
          chmod +x $out/bin/download-geofabrik

          ${self.packages.${system}.default}/bin/download-geofabrik completion bash > download-geofabrik.bash
          ${self.packages.${system}.default}/bin/download-geofabrik completion fish > download-geofabrik.fish
          ${self.packages.${system}.default}/bin/download-geofabrik completion zsh > _download-geofabrik

          installShellCompletion --cmd download-geofabrik \
            --bash download-geofabrik.bash \
            --fish download-geofabrik.fish \
            --zsh _download-geofabrik
        '';
      in
      {
        packages = {
          default = buildGoModule {
            pname = "download-geofabrik";
            version = "unstable";
            src = ./.;
            vendorHash = "sha256-sWU1ayWMz9XIyv5qaWHZVOBDfWWdq5PeYLooRNlND4M=";
            subPackages = [ "cmd/download-geofabrik" ];
            env.CGO_ENABLED = 0;

            nativeBuildInputs = [ pkgs.installShellFiles ];

            postInstall = ''
              $out/bin/download-geofabrik completion bash > download-geofabrik.bash
              $out/bin/download-geofabrik completion fish > download-geofabrik.fish
              $out/bin/download-geofabrik completion zsh > _download-geofabrik

              installShellCompletion --cmd download-geofabrik \
                --bash download-geofabrik.bash \
                --fish download-geofabrik.fish \
                --zsh _download-geofabrik
            '';

            flags = [
              "-trimpath"
            ];

            ldflags = [
              "-s"
              "-w"
            ];
          };
        };

        apps = {
          default = flake-utils.lib.mkApp {
            drv = download-geofabrik-dev;
          };
          download-geofabrik = flake-utils.lib.mkApp {
            drv = download-geofabrik-dev;
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
            download-geofabrik-dev
          ];

          shellHook = ''
            export SHELL="''${SHELL:-${pkgs.bashInteractive}/bin/bash}"
            export GOPATH="''${GOPATH:-$HOME/go}"
            export PATH="$GOPATH/bin:$PATH"
            export CGO_ENABLED=0

            if [ -n "''${BASH_VERSION:-}" ]; then
              source <(${download-geofabrik-dev}/bin/download-geofabrik completion bash 2>/dev/null) || true
            elif [ -n "''${ZSH_VERSION:-}" ]; then
              source <(${download-geofabrik-dev}/bin/download-geofabrik completion zsh 2>/dev/null) || true
            fi
          '';
        };
      }
    );
}
