package main

import "log"

func main() {
	var isTrue bool

	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100
	isTrue_ := false

	if myNum > 99 && isTrue_ == false {
		log.Println("myNum is greater than 99 and isTrue_ is set to true")
	} else if myNum <= 100 || !isTrue_ {
		log.Println("myNum is greater than 99 and isTrue_ is set to true")
	} else if myNum > 100 || isTrue_ {
		log.Println("myNum is greater than 99 and isTrue_ is set to true")
	} else {
		log.Println("Not found!")
	}

	myName := "Dayo"

	switch myName {
	case "Dayo":
		log.Println("Dayo is not ", myName)
	case "tayo":
		log.Println("Dayo is not ", myName)
	case "tobi":
		log.Println("Dayo is not ", myName)
	case "samuel":
		log.Println("Dayo is not ", myName)
	default:
		log.Println("Nothing to show")

	}

}
