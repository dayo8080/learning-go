package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myFirstMap := make(map[string]User)

	me := User{
		FirstName: "Dayo",
		LastName:  "Oluseye",
	}

	myFirstMap["me"] = me
	myMap := make(map[string]string)

	myMap["dog"] = "samson"

	myMap["other-dog"] = "cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])
	log.Println(myFirstMap["me"])

	var mySlice []string
	mySlice = append(mySlice, "This")
	mySlice = append(mySlice, "is")
	mySlice = append(mySlice, "a")
	mySlice = append(mySlice, "slice")

	var noSlice []int
	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	noSlice = append(noSlice, 40)
	noSlice = append(noSlice, 5)
	noSlice = append(noSlice, 10)

	log.Println(noSlice)

	sort.Ints(noSlice)

	log.Println(mySlice)
	log.Println(noSlice)
	log.Println(numSlice)
	log.Println(numSlice[:4])
	log.Println(numSlice[2:6])
	log.Println(numSlice[3:])

}
