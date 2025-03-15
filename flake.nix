{
  inputs = { nixpkgs.url = "github:NixOS/nixpkgs"; };
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in with pkgs; {
        devShells.default = pkgs.mkShell {
          packages = [ protobuf protoc-gen-go go-task nodejs_23 ];
        };
      });
}
