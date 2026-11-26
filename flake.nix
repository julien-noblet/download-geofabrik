{
  description = "download-geofabrik development and build environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
        buildGoModule = pkgs.buildGo127Module;
      in
      {
        packages = {
          default = buildGoModule {
            pname = "download-geofabrik";
            version = "unstable";
            src = ./.;
            vendorHash = "sha256-qDdRRRUWV5dJYeBwlM9mCHZQ/rwW5c2tb+eeRfoLqtc=";
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
            drv = self.packages.${system}.default;
          };
          download-geofabrik = flake-utils.lib.mkApp {
            drv = self.packages.${system}.default;
          };
        };
      }
    );
}
