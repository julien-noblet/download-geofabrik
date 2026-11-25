{
  pkgs,
  lib,
  ...
}:

{
  env.CGO_ENABLED = "0";

  languages.go = {
    enable = true;
    package = pkgs.go_1_27;
  };

  packages = [
    pkgs.actionlint
    pkgs.deadnix
    pkgs.delve
    pkgs.gitleaks
    pkgs.golangci-lint
    pkgs.gopls
    pkgs.goreleaser
    pkgs.gotools
    pkgs.govulncheck
    pkgs.markdownlint-cli
    pkgs.shellcheck
    pkgs.statix
    pkgs.yamllint
  ];

  scripts = {
    download-geofabrik.exec = ''
      exec go run ./cmd/download-geofabrik "$@"
    '';

    update-vendor-hash.exec = ''
      set -euo pipefail
      if [ ! -f flake.nix ]; then
        echo "flake.nix not found in current directory" >&2
        exit 1
      fi

      if nix --extra-experimental-features "nix-command flakes" build --no-link >/dev/null 2>&1; then
        echo "vendorHash in flake.nix is already up to date"
        exit 0
      fi

      sed -i 's|vendorHash = ".*";|vendorHash = "sha256-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=";|' flake.nix
      NEW_HASH=$(nix --extra-experimental-features "nix-command flakes" build --no-link 2>&1 | sed -n 's/.*got: *\(sha256-[^ ]*\).*/\1/p' | head -n1 || true)
      if [ -n "$NEW_HASH" ]; then
        sed -i "s|vendorHash = \".*\";|vendorHash = \"$NEW_HASH\";|" flake.nix
        echo "Updated vendorHash in flake.nix to $NEW_HASH"
      else
        echo "Failed to determine vendorHash from nix build" >&2
        exit 1
      fi
    '';
  };

  tasks = {
    "go:test" = {
      description = "Run Go unit tests";
      exec = "go test ./...";
    };
    "go:lint" = {
      description = "Run golangci-lint";
      exec = "golangci-lint run";
    };
    "go:vuln" = {
      description = "Run govulncheck";
      exec = "govulncheck ./...";
    };
    "go:build" = {
      description = "Build download-geofabrik binary";
      exec = "go build -o ./download-geofabrik ./cmd/download-geofabrik";
    };
  };

  enterShell = ''
    echo "download-geofabrik dev environment (Go $(go version | awk '{print $3}'))"
    echo "  Scripts : download-geofabrik, update-vendor-hash"
    echo "  Tasks   : devenv tasks run go:test | go:lint | go:vuln | go:build"
  '';

  treefmt = {
    enable = true;
    config = {
      programs = {
        gofmt.enable = true;
        nixfmt.enable = true;
        prettier.enable = true;
      };
      settings.global.excludes = [
        "devenv.lock"
        "flake.lock"
        "*.yml"
        "README.md"
        "CHANGELOG.md"
        "docs/catalogs/*"
      ];
    };
  };

  git-hooks.hooks = {
    treefmt.enable = true;
    golangci-lint.enable = true;
    govet.enable = true;
    deadnix.enable = true;
    statix.enable = true;
    actionlint.enable = true;
    shellcheck.enable = true;
    gitleaks.enable = true;
    yamllint = {
      enable = true;
      excludes = [ "^[^/]+\\.yml$" ];
      settings = {
        preset = "relaxed";
        configuration = ''
          extends: relaxed
          rules:
            line-length: disable
            truthy:
              allowed-values: ["true", "false", "on"]
        '';
      };
    };
    markdownlint = {
      enable = true;
      excludes = [
        "^README\\.md$"
        "^CHANGELOG\\.md$"
        "^\\.chglog/"
        "^docs/catalogs/"
      ];
    };
    go-mod-tidy = {
      enable = true;
      name = "go-mod-tidy";
      entry = "${lib.getExe pkgs.go_1_27} mod tidy";
      files = "^go\\.(mod|sum)$";
      pass_filenames = false;
    };
    update-vendor-hash = {
      enable = true;
      name = "update-vendor-hash";
      entry = "${pkgs.bash}/bin/bash -c 'update-vendor-hash && git add flake.nix'";
      files = "^go\\.(mod|sum)$";
      pass_filenames = false;
    };
    end-of-file-fixer = {
      enable = true;
      excludes = [
        "devenv\\.lock"
        "flake\\.lock"
        "\\.yml$"
        "\\.md5$"
      ];
    };
    trim-trailing-whitespace = {
      enable = true;
      excludes = [
        "devenv\\.lock"
        "flake\\.lock"
        "\\.yml$"
        "\\.md5$"
      ];
    };
  };

  enterTest = ''
    set -euo pipefail
    go test ./...
    golangci-lint run
    govulncheck ./...
    go build -o "$TMPDIR/download-geofabrik" ./cmd/download-geofabrik
    "$TMPDIR/download-geofabrik" --version
    if command -v nix >/dev/null 2>&1; then
      nix --extra-experimental-features "nix-command flakes" build --no-link
    fi
  '';
}
