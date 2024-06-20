# Project setup
Go programs are structured in: `repository -> modules -> packages -> source files`

1. create `go.mod` file: `go mod init <module path>`
2. build and install a program: `go install <module path>` (binary at `$HOME/go/bin/`) 

Executable source files must always use `package main` with a `main` function.

Fields and methods are only exported if the first letter is uppercase.

Use always full package name to import, for remote ones the module has to be downloaded. (`go mod tidy`)

Module dependencies are downloaded to pkg/mod indicated by the GOPATH environment variable. Remove all downloaded modules: `go clean -modcache` 

# Idiomatics
Format code with `go fmt`

Local variable declaration inside if or for loop definition not before.

It's idiomatic—to write an if-else-if-else chain as a switch.

# Testing
Test files ending in `_test.go`, functions are named TestXXX with signature `func (t *testing.T).

Run tests with `go test`

# Names
Use MixedCaps or mixedCaps

Use the package structure to help you choose good names, use single-word names. A helpful doc comment can be more valuable than an extra long name.

Getter have no "get".

One-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader






