# Idiomatics
## Project setup
Go programs are structured in: `repository -> modules -> packages -> source files`

Run single go file with `go run <filename>`

1. create `go.mod` file: `go mod init <module path>`
2. build and install a program: `go install <module path>` (binary at `$HOME/go/bin/`) 

Executable source files must always use `package main` with a `main` function.

Fields and methods are only exported if the first letter is uppercase.

Use always full package name to import, for remote ones the module has to be downloaded. (`go mod tidy`)

Module dependencies are downloaded to pkg/mod indicated by the GOPATH environment variable. Remove all downloaded modules: `go clean -modcache` 

## Writing style
Format code with `go fmt`

Local variable declaration inside if or for loop definition not before.

It's idiomatic—to write an if-else-if-else chain as a switch.

With multiple return values use return value names and an unadorned `return`

Idiomatic slice allocation: `v := make([]int, 100)`

Use MixedCaps or mixedCaps

Use the package structure to help you choose good names, use single-word names. A helpful doc comment can be more valuable than an extra long name.

Getter have no "get".

One-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader

Test files ending in `_test.go`, functions are named TestXXX with signature `func (t *testing.T).

Run tests with `go test`

## Data
New allocation, `New(T)`, returns a pointer to a newly allocated zero value of type T

Composite literals create a new instance: T{...}

Constructores should be in the form `func NewT(...) \*T`

Make allocation, `make(T, args)`, creates slices, maps and channels only. Returns an initialized value of type T

Use "comma ok" idiom to distinguish missing entries from zero values

Use slices over arrays


# Data types
## Array
Arrays are useful when planning the detailed layout of memory, but primarily they are a building block for slices.

## Slice
Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.

### Two-dimensional slice
Diffferent ways to create:
```go
type Text [][]byte
t := Text{
    []byte("Hello"),
    []byte("World),
}

p := make([][]uint8, YSize)
for i := range p {
    p[i] = make([]uint8, XSize)
}
```
### Map
```go
// initialize
t := make(map[string]int)
t := map[string]int{}

// initialize with data
t := map[string]int{
    "UTC": 0,
    "EST": 1,
}

// add value
t["aaa"] = 2

// retrieve value; key not exists => zero value
x := t["aaa"]

// delete value
delete(t, "UTC")

// length
len(t)

// test key exists
value, ok := t["aaa"]

// iterate not ordered
for key, value := range t {
}

// iterate ordered
m := map[int]string{2: "2", 1: "1"}
var keys []int
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
}

// concurrent use
var c = struct{
    sync.RWMutex
    m map[string]int
}{m: make(map[string]int)}
// concurrent read
c.RLock()
n := counter.m["someKey"]
c.RUnlock()
//concurrent write
c.Lock()
c.m["someKey"]++
c.Unlock()
```

## Set
Use bool maps.
```go
// set of strings
s := make(map[string]bool)
```























