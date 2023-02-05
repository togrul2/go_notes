package main

import (
	"fmt"
	"strconv"
)

// Only way to declare variable at package level
// Will be scoped to package since begins with lowercase letter
var i int = 42

// Will be scoped globally since begins with uppercase letter
var S int = 54

var (
	actorFirstName string = "Elisabeth"
	actorLastName  string = "Smith"
	id             int    = 1
)

func main() {
	var n int = 27
	fmt.Printf("%v %T\n", n, n) // 27 int

	var f float32 = 127 // Some types can only be assined explicitly
	fmt.Printf("%v %T\n", f, f)

	var a, b int = 2, 3
	// Another way with separate declaration and assignment
	// var a, b int
	// a, b = 2, 3

	fmt.Println(a, b) // 2 3

	// This way only can be used in functions, only for declaration not for assignment
	// Also this way we can implicitly assign float64 type
	k := 99.
	fmt.Printf("%v %T\n", k, k) // 99 float64

	// Best practice to have acronyms all uppercase
	var theURL string = "https://google.com"
	fmt.Println(theURL)

	var z float32 = float32(i)
	fmt.Printf("%v %T\n", z, z) // 42 float32

	// var s string = string(z) // Will give ascii character with index i, not the result we want
	var s string = strconv.Itoa(i)
	fmt.Printf("%v %T\n", s, s) // 42 string

	// We can't redeclare variables but shadow them
	fmt.Println(i) // i from package scope
	var i int = 750
	fmt.Println(i) // i from func scope
}
