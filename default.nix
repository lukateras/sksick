with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "sks-tools";

  buildInputs = [ gnupg go python3Packages.scrapy ];

  GOPATH = runCommand "gopath" {} ''
    mkdir $out && ln -s ${toString ./.} $out/src
  '';
}
