// go.mod is the Go module definition file — like package.json in Node or Cargo.toml in Rust.
// It tells Go the name of this module and which Go version to use.

// "module" declares the module path — a unique name used for imports.
// By convention it matches the repository URL, but it can be any string.
module github.com/iamthebonbon/go-sample

// "go" sets the minimum Go version this module requires.
// The toolchain will warn if you try to build with an older version.
go 1.26.3
