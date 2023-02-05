package main

import "fmt"

type Animal struct {
	name string
	age  int
}

// If we want to be able to modify attrs of struct, we have to pass pointer to it like below
// func (animal *Animal) eat() {}

func (animal Animal) GetInfo() {
	fmt.Printf("Hi, my name is %s. I'am %d years old", animal.name, animal.age)
}

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hi!", i)
	}
	greeting("Hi", "Togrul")
	greet, name := "Hi", "Togrul"
	greeting2(&greet, &name) // values of greet and name can be changed since we pass reference to them
	fmt.Println("Name", name)
	fmt.Println("Average", avg(1, 2, 3, 4, 5))

	res, err := divide(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// Anonymous function
	func() {
		fmt.Println("From anonymous!")
	}()

	for i := 0; i < 5; i++ {
		// Better practice to pass values to functions instead of taking from outer scope
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	f := func() {
		fmt.Println("Declarated function")
	}
	f()

	var divide func(a, b float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0., fmt.Errorf("Cannot pass 0 as a divider")
		}
		return a / b, nil
	}
	fmt.Println(divide(5.0, 3.0))
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Printf("The value of index is %v\n", idx)
}

func greeting(greeting, name string) { // If types are the same, it can be written at the end
	fmt.Println(greeting, name)
	name = "Ted" // Is not gonna change passed value from outside
}

func greeting2(greting, name *string) {
	fmt.Println(*greting, *name)
	*name = "Ted" // Gonna change variable, outside one also
}

func avg(values ...int) int {
	result := 0
	for _, n := range values {
		result += n
	}
	return result / len(values)
}

func sum(values ...int) *int {
	result := 0
	for _, n := range values {
		result += n
	}
	return &result
}

// same as above
func sum2(values ...int) (result int) { // default val is given (0)
	for _, n := range values {
		result += n
	}
	return // will return result, named return type
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0, fmt.Errorf("Cannot provide 0 as a divider")
	}
	return a / b, nil
}
