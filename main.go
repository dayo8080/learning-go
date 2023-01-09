package main

import "log"

// defining an interface with a name of Animal, all function inside the interface must be implemented
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// create a Dog struct
type Dog struct {
	Name  string
	Breed string
}

// create a gorilla struct
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	// create an object of Dog struct
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}
	// create an object of Gorilla struct
	gorilla := Gorilla{
		Name:          "Gree",
		Color:         "Black",
		NumberOfTeeth: 34,
	}
	// call the printinfo func with object dog as the parameter
	PrintInfo(dog)
	// call the printinfo func with object gorilla as the parameter
	PrintInfo(gorilla)
}

// create a function Says with a receivable g of Gorilla struct, which returns a string
func (g Gorilla) Says() string {
	return "Killll"
}

// create a function NumberOfLegs with a receivable g of Gorilla struct, which returns a int
func (g Gorilla) NumberOfLegs() int {
	return 6
}

// create a function Says with a receivable d of Dog struct, which returns a string
func (d Dog) Says() string {
	return "woof"
}

// create a function NumberOfLegs with a receivable d of Dog struct, which returns a int
func (d Dog) NumberOfLegs() int {
	return 4
}

// create a function PrintInfo with an argument a as the interface
func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and have a number of ", a.NumberOfLegs(), "legs")
}
