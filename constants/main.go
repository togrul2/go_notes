package main

import (
	"fmt"
)

const rectangleSide = 5 // will not be exported
const RectangleSide = 5 // will be exported

// iota, enumeration constant
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

// Same as above, just shorter
// const (
// 	a = iota
// 	b
// 	c
// )

const (
	// it will start from 0 since it is in different scope
	a2 = iota // 0
	b2        // 1
	c2        // 2
)

// This way our constants will start with 1
const (
	_    = iota
	dog  // 1
	cat  // 2
	bird // 3
)

// Same thing
// const (
// 	dog  = iota + 1 // 1
// 	cat             // 2
// 	bird            // 3
// )

const (
	_  = iota             // ignore 0 value
	KB = 1 << (10 * iota) // 1024
	MB                    // 1048576
	GB                    // 1073741824
	TB                    // 1099511627776
	PB                    // 1125899906842624
	EB                    // 1152921504606846976
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

// typed constants can only work with the same types, while untyped ones work with any matching ones
const f int16 = 46

func main() {
	const s = 42
	// Function return values can't be used for constants
	// const sin = math.Sin(1.57) // Error

	fmt.Printf("%v %T\n", f, f) // 46 int16
	const f = 128               // constants can be shadowed
	fmt.Printf("%v %T\n", f, f) // 128 int

	const per = 42
	var side int16 = 16
	// Will still since const values will be replaced with its values
	fmt.Printf("%v %T\n", per+side, per+side) // 58 int16

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB) // 3.73GB

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin: %v\n", roles&isAdmin == isAdmin)                 // true
	fmt.Printf("Can see africa: %v\n", roles&canSeeAfrica == canSeeAfrica) // false
}
