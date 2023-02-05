package main

import "fmt"

// Generic interface
type Writer[T any] interface {
	Write([]T) error
	Read() error
}

type Number interface {
	int32 | float32
}

func show[T Number]() int {
	return 1
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// Also can be used with structs
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	s_int := []int{10, 20, 15, -10}
	fmt.Println(Index(s_int, 10))

	s_string := []string{"Hi", "world", "!"}
	fmt.Println(Index(s_string, "Hi"))
}
