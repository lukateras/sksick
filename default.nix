{ pkgs ? import ./pkgs.nix {} }: with pkgs;

buildGoPackage {
  name = "sks-exploit";
  src = lib.cleanSource ./.;

  goPackagePath = "gitlab.com/yegortimoshenko/sks-exploit";
}
