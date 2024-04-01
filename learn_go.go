package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
	"time"
)

// run with go run learn_go.go
// format this file go fmt learn_go.go

func main() {
	me()
}

// FUNCTIONS
func add1(x int, y int) int { return x + y }

// type omitting
func add2(x, y int) int { return x + y }

// multiple results
func swap(x, y string) (string, string) { return y, x }

// naked returns (example returns x, y)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// functions as function argument
func compute(fn func(int, int) int) int {
	return fn(5, 5)
}

func nested() {
	addNested := func(x, y int) int { return x + y }
	fmt.Println(addNested(2, 2))
	fmt.Println(compute(addNested))
}

// A closure is a function value that references variables from outside its body.
func adder() func(int) int {
	sum := 0 // each closure is bound to sum
	return func(x int) int {
		sum += x
		return sum
	}
}

func closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// METHODS
// Go does not have classes. However, you can define methods on types.
// call with:
// v := Vertex1{3, 4}
// fmt.Println(v.Abs())
type Vertex1 struct {
	X, Y float64
}

func (v Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method pointer receivers
// used to change variable or work with big structs more efficiently
// all methods on a given type should have either value or pointer receivers
// call with v.Scale(10)
func (v *Vertex1) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// type parameters (generic functions)
// func Index[T comparable](s []T, x T) int
// This declaration means that s is a slice of any type T that fulfills the built-in constraint comparable. x is also a value of the same type.

// generic types
type List[T any] struct {
	next *List[T]
	val  T
}

// INTERFACES
// call with:
//
//	var i I = T{"hello"}
//	i.M()
type I interface {
	M()
}

type T struct {
	S string
}

// with this method type T implements interface I
func (t T) M() {
	fmt.Println(t.S)
}

// interface type assertions
// A type assertion provides access to an interface value's underlying concrete value.
func typeAssertions() {
	var i interface{} = "hello"
	// asserts that the interface value i holds a string and assigns the underlying string value to the variable s
	// panics if no string value
	s := i.(string)
	fmt.Println(s)
	// additionally asserts that the interface value holds a specific type and saves it to the variable ok
	// does panic
	s, ok := i.(string)
	fmt.Println(s, ok)
}

// interface type assertions with switch
func typeAssertionsSwitch(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
		return 0, nil
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
		return 0, nil
	default:
		fmt.Printf("I don't know about type %T!\n", v)
		return 0, nil
	}
}

// ERRORS
// A nil error denotes success; a non-nil error denotes failure.
func errorHandling() {
	var i interface{} = "hello"
	if v, err := typeAssertionsSwitch(i); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}

// READERS
// read a string 8 bytes at a time
func reader() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		if err == io.EOF {
			break
		}
	}
}

// VARIABLES
// types
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32, represents a Unicode code point
// float32 float64
// complex64 complex128

// variables declared without an explicit initial value are given their zero value
// 0 for numeric types,
// false for the boolean type,
// "" (the empty string) for strings.

// declare variables
const Truth = true

var c, python, java bool
var c1, python1, java1 = true, false, "No!"

func variables() {
	c2, python2, java2 := true, false, "No!"
	fmt.Println("%s, %s, %s", c2, python2, java2)

}

// type conversion
var i = 1
var f float64 = float64(i)

// LOOPS / CONDITIONALS
func conditionals() {
	// loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// while loop
	whileI := 0
	for whileI < 10 {
		whileI += 1
		fmt.Println(whileI)
	}
	// print "A" before if
	a := 2
	if fmt.Println("A"); a > 1 {
		fmt.Println("hello")
	}
}

// DEFER STATEMENT
// defers the execution of a function until the surrounding function returns
//
// defer fmt.Println("world")
// fmt.Println("hello")
//
// outputs hello world

// POINTERS
func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// STRUCTS
type Vertex struct {
	X int
	Y int
}

func structs() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	// pointer
	p := &v
	p.X = 1e9
	fmt.Println(v)
	v2 := Vertex{X: 1} // Y:0 is implicit
	fmt.Println(v2)
}

// ARRAYS / SLICES
// arrays: fixed size; slices: dynanimc size, zero value is nil, A nil slice has a length and capacity of 0 and has no underlying array.
func arrays() {
	var array1 [2]string
	array1[0] = "hello"
	array1[1] = "world"
	array2 := [5]int{1, 2, 3, 4, 5}
	var slice []int = array2[2:5] // includes elements 2 to 4 of array2
	slice2 := []int{}
	sliceWholeArray := array1[:]
	slice3 := make([]int, 3, 3)
	slice3 = append(slice3, 9, 8, 7)
	fmt.Println(array1, array2, slice, slice2, sliceWholeArray, slice3)
	// loop over slice
	for i, v := range slice3 {
		fmt.Printf("\ni: %d, v: %d", i, v)
	}
	fmt.Println()
}

// MAPS
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
func maps() {
	m1 := make(map[string]int)
	m1["first"] = 1
	m2 := map[string]int{"first": 1, "second": 2}
	delete(m2, "first")
	v, ok := m1["Answer"]
	v1, ok1 := m1["first"]
	fmt.Println(m1, m2, v, ok, v1, ok1)
}

// GOROUTINES
// A goroutine is a lightweight thread managed by the Go runtime.
// go f(x, y, z)  // start new goroutine/ run function in new thread
// The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
// Goroutines run in the same address space, so access to shared memory must be synchronized.

// goroutine channels
// wit channels you can send and receive values to and from goroutines
// By default, sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.
// example uses two goroutines to sum up the integers in a slice
func sumConc(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

func sumUpConcurrent() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sumConc(s[:len(s)/2], c)
	go sumConc(s[len(s)/2:], c)
	x, y := <-c, <-c // receive sum from channel c and save value to x and y
	fmt.Println(x, y, x+y)
}

// goroutine channels buffered
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
func bufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// goroutine channels Range and Close
// A sender can close a channel to indicate that no more values will be sent.
// Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
// Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
// channels must not be closed always
func fibonacci9(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeClose() {
	c := make(chan int, 10)
	go fibonacci9(cap(c), c)
	for i := range c { // receive values until channel closed
		fmt.Println(i)
	}
}

// goroutine select
// lets a goroutine wait for multiple communication operations
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
// following example explained:
// When the program runs, the main function starts executing.
// The fibonacci function is called concurrently with the goroutine started inside the main function.
// The fibonacci function generates Fibonacci numbers and sends them to channel c.
// The goroutine inside main receives these values and prints them.
// After printing 10 Fibonacci numbers, the goroutine sends a value to channel quit, which causes the fibonacci function to terminate.
// The program exits after printing "quit".
func fibonacci834(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func goroutineSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci834(c, quit)
}

// goroutines mutual exclusion
// Makes sure only one goroutine can access a variable at a time
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock() // using defer will unlock also if the function panics
	return c.v[key]
}
func me() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("someKey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("someKey"))
}
