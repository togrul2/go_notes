package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string // These vars are package scoped, it starts with lowercase
	Episodes   []string // This one is global, can be accessed from outside of package
}

type Animal struct {
	Name   string
	Origin string
}

// This way Bird doesn't become child of Animal but it copies its fields
type Bird struct {
	Animal // Embed Animal struct into Bird struct, kinda like inheritance
	// Animal Animal, this will create field Animal for Bird struct
	SpeedKPH float32
	CanFly   bool
}

type User struct {
	Username  string `json:"username"` // Tags are useful for packages such as json
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

/*
Str repr of struct, behaves like toString() from Java
*/
func (user User) String() string {
	return fmt.Sprintf("User(username: %v, email: %v)", user.Username, user.Email)
}

func main() {
	// Struct declaration
	doctor := Doctor{
		number:     3,
		actorName:  "Josh Hawkins",
		companions: []string{"Ann Mathew", "Bill Mason"},
	}
	// Also possible but discouraged, version above is preferrable
	doctor = Doctor{
		3,
		"Josh Hawkins",
		[]string{"Ann Mathew", "Bill Mason"},
		[]string{},
	}
	fmt.Println("Doctor:", doctor)
	fmt.Println("Doctor name:", doctor.actorName)

	// Anonymous declaration
	doctor2 := struct{ name string }{name: "Josh Pertwee"}
	anotherDoctor2 := doctor2     // assignment by value
	anotherDoctorRef2 := &doctor2 // assignment by reference

	doctor2.name = "Tom Baker"
	fmt.Println(doctor2.name)           // Tom Baker
	fmt.Println(anotherDoctorRef2.name) // Tom Baker
	fmt.Println(anotherDoctor2.name)    // Josh Pertwee

	bird := Bird{}
	bird.Name = "Emu"
	bird.Origin = "Australia"
	bird.SpeedKPH = 48
	bird.CanFly = false
	fmt.Printf("bird: %v\n", bird)

	// Same thing as above
	bird = Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	user := &User{
		Username:  "togrulir",
		FirstName: "Togrul",
		LastName:  "Asadov",
		Email:     "adsads00123@gmail.com",
	}

	out, err := json.Marshal(user)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
