package main

import (
	"fmt"
	"math"
)

func main() {
	flag := true

	if flag { // Curly braces always are required even with one-liners
		fmt.Println("Flag is true")
	}

	statePopulations := map[string]int{
		"California": 39_250_017,
		"Texas":      27_862_596,
		"Florida":    20_61_2439,
		"New York":   19_745_289,
	}

	// Go way of declaring value, like walrus in python
	if val, ok := statePopulations["Texas"]; ok {
		fmt.Println("Texas:", val)
	}

	number := 50
	guess := 30

	if guess < 1 || guess > 100 {
		fmt.Println("Number out of range")
	} else {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("guessed")
		}
	}

	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else if myNum/math.Pow(math.Sqrt(myNum), 2)-1 < 0.001 {
		fmt.Printf("These are nearly the same")
	} else {
		fmt.Println("These are the different")
	}

	// This declaration also works with switch statememnts
	// Go's switch doesn't require break keyword
	switch op := 2; op {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 5, 6, 7:
		fmt.Println("5 6 or 7")
	default:
		fmt.Println("Neither one nor two")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less or equal to 10")
		fallthrough // This is like break in other languages but works opposite
	case i <= 20:
		fmt.Println("Less or equal to 20")
		fallthrough
	default:
		fmt.Println("Greater than 20")
	}

	// Type switch
	var j interface{}
	switch j = [3]int{1, 2, 3}; j.(type) {
	case int:
		fmt.Println("I is a type of int")
	case float32:
		fmt.Println("I is a type of float32")
	case string:
		fmt.Println("I is a type of string")
	case [3]int:
		fmt.Println("I is a type of int array of len 3")
	default:
		fmt.Println("I is a some different type")
	}

	switch k := 1; k {
	case 1:
		fmt.Println("1")
		if 1 == 1 {
			break // This will quit the switch early
		}
		fmt.Print("11") // Will not be run
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}
}
