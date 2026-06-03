{
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  outputs = inputs: let
    system = "x86_64-linux";
    pkgs = import inputs.nixpkgs {
      inherit system;
      config = {allowUnfree = true;};
    };
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages = with pkgs; [
        go
        php
        intelephense

        (pkgs.writers.writeBashBin "aoc" ''
          export $(grep -v '^#' .env | xargs)
          ${pkgs.lib.getExe php} runner.php $@
        '')
      ];
    };
  };
}
