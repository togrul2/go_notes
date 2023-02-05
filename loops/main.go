package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println() // newline

	// Working with two vars
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Printf("%v %v\n", i, j)
	}

	i := 0
	for ; i < 10; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println() // newline

	j := 5
	for j < 10 {
		fmt.Printf("j: %v\n", j)
		j++
	}

	for { // infinite loop
		i++
		if i == 20 {
			break
		}
		if i%3 == 0 {
			continue
		}
		fmt.Printf("%v ", i)
	}
	fmt.Println() // newline

Loop: // Loop label
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Println(i, j)
			if i*j > 20 {
				break Loop // breaks outer loop
			}
		}
	}

	s := []int{1, 2, 3}
	for k, v := range s { // this works with arrays, slices, maps, strings and channels
		fmt.Printf("%v: %v\n", k, v)
	}
}
