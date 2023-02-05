package main

import "fmt"

func main() {
	// grades := [3]int{97, 78, 85}
	grades := [...]int{97, 78, 85} // same thing as above, here go assigns the length itself
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ann"
	students[2] = "Josh"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #2: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))

	var identityMatrix [3][3]int

	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{1, 0, 0}
	identityMatrix[2] = [3]int{1, 0, 0}

	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a // Will copy array

	sliceA := []int{1, 2, 3} // This is a slice
	sliceB := sliceA         // slice is reference so it will behave like python list

	// This will assign b as alias for a, this way we will have similar behavior of list from Python
	// b := &a
	// fmt.Println(*b) //
	b[1] = 5
	fmt.Println(a) // [1, 2, 3]
	fmt.Println(b) // [1, 5, 3]

	sliceB[1] = 5
	fmt.Println(sliceA) // [1, 5, 3]
	fmt.Println(sliceB) // [1, 5, 3]

	fmt.Printf("Length: %v\n", len(sliceA))
	fmt.Printf("Capacity: %v\n", cap(sliceA))

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Slicing also work with array, they create reference though
	allElements := slice[:] // slice of all elements
	fromNtoEnd := slice[3:]
	beginTillTheN := slice[:6]
	fromNtoM := slice[3:6]
	fmt.Println("[:]:", allElements)    // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("[3:]:", fromNtoEnd)    // [4 5 6 7 8 9 10]
	fmt.Println("[:6]:", beginTillTheN) // [1 2 3 4 5 6]
	fmt.Println("[3:6]:", fromNtoM)     // [4 5 6]

	m := make([]int, 3, 100)             // arguments: type, length, capacity(default=length)
	fmt.Println(m)                       // [0, 0, 0]
	fmt.Printf("Length: %v\n", len(m))   // 3
	fmt.Printf("Capacity: %v\n", cap(m)) // 100

	rm := append(m, 1)
	fmt.Println(rm)                       // [0, 0, 0, 1]
	fmt.Printf("Length: %v\n", len(rm))   // 4
	fmt.Printf("Capacity: %v\n", cap(rm)) // 100

	e := []int{}
	re := append(e, 1)
	fmt.Println(re)                       // [1]
	fmt.Printf("Length: %v\n", len(re))   // 1
	fmt.Printf("Capacity: %v\n", cap(re)) // 1

	re2 := append(e, 2, 3, 4, 5, 6)
	fmt.Println(re2)                       // [1, 2, 3, 4, 5, 6]
	fmt.Printf("Length: %v\n", len(re2))   // 5
	fmt.Printf("Capacity: %v\n", cap(re2)) // 6

	re3 := append(e, []int{7, 8, 9}...) // spread operator from java, js
	fmt.Println(re3)                    // [7,8,9]

	stack := []int{1, 2, 3, 4}
	unshifted := stack[1:]                                          // Remove from the bottom
	popped := stack[:len(stack)-1]                                  // Remove from the top
	fmt.Println(popped)                                             // [1,2,3]
	target := 2                                                     // target index
	removeFromMiddle := append(stack[:target], stack[target+1:]...) // removes element at target index
	fmt.Println("Unshifted:", unshifted)
	fmt.Println("Removed from middle", removeFromMiddle)
}
