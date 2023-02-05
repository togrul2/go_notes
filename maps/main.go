package main

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 39_250_017,
		"Texas":      27_862_596,
		"Florida":    20_61_2439,
		"New York":   19_745_289,
	}
	fmt.Println(statePopulations)

	m := map[[3]int]string{} // Array is a valid key value(hasable) but slice is not
	fmt.Println(m)           // map[]

	val, ok := statePopulations["New York"]
	fmt.Printf("New york population: %v\nIs available: %v\n", val, ok)
	val, ok = statePopulations["Georgia"] // key is not present, we will get default value for its value type
	fmt.Printf("New york population: %v\nIs available: %v\n", val, ok)

	statePopulations["Georgia"] = 10_310_371 // Adding a new value
	fmt.Println(statePopulations)            // Will have Georgia
	delete(statePopulations, "Georgia")      // Delete Georgia
	fmt.Println(statePopulations)            // Will not have Georgia

	_, ok2 := statePopulations["New york"] // Only can check presence from ok2
	fmt.Println(ok2)

	mapCopy := statePopulations
	delete(mapCopy, "Florida")    // Unlike arrays it will also remove from all linked vars
	fmt.Println(statePopulations) // Will not have a Florida key
	fmt.Println(mapCopy)          // Also will not have a Florida key
}
