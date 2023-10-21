with import <nixpkgs> {};

pkgs.mkShell {
  nativeBuildInputs = with pkgs; [
    gnumake
    gcc
  ];

  hardeningDisable = [ "all" ];
}
