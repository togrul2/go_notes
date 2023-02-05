package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	a := 42
	b := a
	fmt.Println(a, b) // 42 42
	b = 74
	fmt.Println(a, b) // 42 74

	// & - gives the location, * - gives the value in the location
	var a2 int = 42
	var b2 *int = &a2
	fmt.Println(a2, *b2) // 42 42
	a2 = 68
	fmt.Println(&a2, b2) // &a2 == b2
	fmt.Println(a2, *b2) // 68 68
	*b2 = 14
	fmt.Println(a2, *b2) // 14 14

	l := [3]int{1, 2, 3}
	l1 := &l[0]
	l2 := &l[1]
	// l3 := &l[1] - 4                     // Will return error
	fmt.Printf("%v %p %p\n", l, l1, l2) // [1 2 3], x, x + 4 where x is a random hex number

	var ms *myStruct // at this moment its value is nil
	ms = &myStruct{foo: 42}
	fmt.Println(*ms) // {42}

	var ms2 *myStruct = new(myStruct) // new function initializes type without passing empty values
	(*ms2).foo = 42
	ms2.foo = 42     // Absolutely the same as above
	fmt.Println(ms2) // &{42}

}
