package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle") // will be run after this func returns
	fmt.Println("End")

	// Three, Two, One
	one()

	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // Closes response after all operations are done
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)

	a := 5
	defer fmt.Println(a) // 5
	a = 10

	// i, j := 1, 0
	// ans := i / j // this is gonna throw a panic
	// fmt.Println(ans)

	// deferAndPanic()
	// Start
	// Middle
	// ...Panic

	fmt.Println("Start")
	deferFuncRecover()
	fmt.Println("End")

	// Start
	// smth bad happened
}

func one() {
	// Last line will run first
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
}

func runPanic() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func deferAndPanic() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	panic("something bad happened")
	fmt.Println("End")
}

func deferFuncRecover() {
	// Recover is only useful in defered functions
	// Function containing recover won't run till the end but higher function in stack will
	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			// rethrow panic if we can't deal with it, so we handle it in higher level
			panic(err)
		}
	}()
	panic("smth bad happened")
	fmt.Println("End")
}
