package main

import (
	"fmt"
	"math"
)

// format this file go fmt learn_go.go

func main() {
	v := Vertex1{3, 4}
	fmt.Println(v.Abs())
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
func compute (fn func(int, int) int) int {
	return fn(5, 5)
}

func nested() {
	addNested := func(x, y int) int { return x+y }
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
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
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
	var array1[2]string
	array1[0] = "hello"
	array1[1] = "world"
	array2 := [5]int{1,2,3,4,5}
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
	m2 := map[string]int{ "first": 1, "second": 2 }
	delete(m2, "first")
	v, ok := m1["Answer"]
	v1, ok1 := m1["first"]
	fmt.Println(m1, m2, v, ok, v1, ok1)
}
