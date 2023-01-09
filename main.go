package main

import (
	"log"

	"github.com/dayo8080/learning-golang/helpers"
)

func main() {
	var myVar helpers.SomeType

	myVar.FirstName = "Dayo"
	myVar.LastName = "Oluseye"

	log.Println(myVar.FirstName, myVar.LastName)

}
