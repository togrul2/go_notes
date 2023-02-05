package main

import "fmt"

type Formatted interface {
	String() string
}

type IntCounter int

type Incrementer interface {
	Increment() int
}

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type Car struct {
	brand   string
	model   string
	mileage uint
	year    uint
}

func (c Car) String() string {
	return fmt.Sprintf("Car(brand: %v, model: %v)", c.brand, c.model)
}

func (c Car) Write(data []byte) (int, error) {
	return 1, nil
}

func main() {
	var car Formatted = &Car{}
	fmt.Println(car)

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}
