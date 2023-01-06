package main

import (
	"log"
	"time"
)

var s = "package"

func main() {
	var s = "main"

	type User struct {
		FirstName   string
		LastName    string
		PhoneNumber string
		Age         int
		BirthDate   time.Time
	}

	user := User{
		FirstName:   "Dayo",
		LastName:    "Oluseye",
		Age:         40,
		PhoneNumber: "07085405644",
	}

	log.Println(user.FirstName, user.LastName, user.BirthDate)

	log.Println("s is", s)
	log.Println("s2 is", s)

	i, j := saySomething("SAY")
	log.Println(i, j)

	log.Println(saySomething("SUP"))

}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	log.Println("s from the saySomething func is", s3)
	return s3, "Hello"
}
