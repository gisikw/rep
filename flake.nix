{
  description = "crane — headless agent dispatcher";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "crane";
          version = "0.1.0";
          src = ./.;
          vendorHash = null; # Update after first build
        };

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [ go gopls just ];

          shellHook = ''
            echo "crane dev shell"
            echo "  just test    — run tests"
            echo "  just build   — build binary"
          '';
        };
      }
    );
}
