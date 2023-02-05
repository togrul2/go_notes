package main

import "fmt"

func main() {
	// var n bool = true
	n := 1 == 1
	m := 1 == 2
	var l bool                  // will assign default 0 value (false in this case)
	fmt.Println(l)              // false
	fmt.Printf("%v %T\n", n, n) // true bool
	fmt.Printf("%v %T\n", m, m) // false bool

	num := 42
	fmt.Printf("%v %T\n", num, num) // true bool
	// int8   from 2^7 to 2^7 - 1
	// int16  from 2^15 to 2^15 - 1
	// int32  from 2^31 to 2^31 - 1
	// int64  from 2^63 to 2^63 - 1

	// uint8 from 0 to 2^8 - 1, byte is an alias for this
	// uint16 from 0 to 2^16 - 1
	// uint32 from 0 to 2^32 - 1
	// uint64 from 0 to 2^64 - 1
	i, j := 10, 3 // Only same types can be used in arithmetic operations
	fmt.Println(i + j)
	fmt.Println(i - j)
	fmt.Println(i * j)
	fmt.Println(i / j)
	fmt.Println(i % j)

	// 10 - 1010
	//  3 - 0011
	fmt.Println("OR", i|j)   // 1011
	fmt.Println("AND", i&j)  // 0010
	fmt.Println("XOR:", i^j) // 1001
	fmt.Println("NOT", ^j)   // -100

	fmt.Println("Rshift", i>>j) // Rshift
	fmt.Println("Lshift", i<<j) // Lshift

	// float32 from +/- 1.18E-38 to +/- 3.4E-38
	// float64 from +/- 2.23E-308 to +/- 1.8E-308

	floatNum := 3.14   // float64 is default
	floatNum = 13.7e72 //13.7E72 also possible
	fmt.Printf("%v %T\n", floatNum, floatNum)

	var float64num float64 = 3.14
	fmt.Printf("%v %T\n", float64num, float64num)

	floatA, floatB := 10.2, 3.7
	fmt.Println(floatA + floatB) // 13.899999999999999
	fmt.Println(floatA - floatB) // 6.499999999999999
	fmt.Println(floatA * floatB) // 37.74
	fmt.Println(floatA / floatB) // 2.7567567567567566

	cnum := 1 + 2i                    // default type is complex 128
	fmt.Printf("%v %T\n", cnum, cnum) // (1 + 2i) complex128

	complexA, complexB := 1+2i, 2+5.2i
	fmt.Println(complexA + complexB) // (3+7.2i)
	fmt.Println(complexA - complexB) // (-1-3.2i)
	fmt.Println(complexA * complexB) // (-8.4+9.2i)
	fmt.Println(complexA / complexB) // (0.3994845360824742-0.038659793814433i)

	fmt.Printf("%v %T\n", real(cnum), real(cnum)) // 1 float32, real part
	fmt.Printf("%v %T\n", imag(cnum), imag(cnum)) // 2 float32, imaginary part

	var cnum2 complex128 = complex(5, 2) // 5 + 2i
	fmt.Printf("%v %T\n", cnum2, cnum2)

	s := "This is a string"
	// s[2] = "u" // can't do this
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", s[2], s[2])         // 105 uint8
	fmt.Printf("%v %T\n", string(s[2]), s[2]) // i uint8

	s1, s2 := "Hello,", " world!"
	fmt.Println(s1 + s2) // Hello, world!

	b := []byte(s)
	fmt.Printf("%v %T\n", b, b)

	r := 'a' // rune, alias for int32
	fmt.Printf("%v %T\n", r, r)
}
