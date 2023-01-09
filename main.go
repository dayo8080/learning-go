package main

import "log"

func main() {

	// create a map with a variable myMap. Maps are key value pairs
	myMap := make(map[string]string)

	// append data to the myMap variable
	myMap["key 0"] = "value 0"
	myMap["key 1"] = "value 1"
	myMap["key 2"] = "value 2"
	myMap["key 3"] = "value 3"

	// loop through a map
	for mapKey, mapValue := range myMap{
		log.Println(mapKey, mapValue)
	}

	// create a User struct
	type User struct {
		FirstName string
		LastName string
		Age int
	}

	// create objects of type User
	user1 := User {
		FirstName: "Dayo",
		LastName: "Oluseye",
	}

		user2 := User {
		FirstName: "Folake",
		LastName: "Oluseye",
	}

	// create a variable mySliceNew with a type User Slice
	var mySliceNew []User
	
	// append values to a slice
	mySliceNew = append(mySliceNew, user1, user2)

	// loop through the slice
	for y, x:=range mySliceNew{
		log.Println(y, x.FirstName, x.LastName)
	}


	

	// a slice also called an array
	mySlice := []string{"dog", "cat", "horse", "fish", "banana"}

	// _ is the index which is ignored,
	for _, x := range mySlice {
		log.Println(x)
	}

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
}
