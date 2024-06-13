# Project setup
Go programs are structured in: `repository -> modules -> packages -> source files`

1. create `go.mod` file: `go mod init <module path>`
2. build and install a program: `go install <module path>` (binary at `$HOME/go/bin/`) 

Executable source files must always use `package main` with a `main` function.

Fields and methods are only exported if the first letter is uppercase.

