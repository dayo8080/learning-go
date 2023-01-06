package main

import "log"

func main() {
	myString := "Green"
	log.Println("Color before the pointer function",myString)
	ChangeColor(&myString)
	log.Println("color after changecolor func", myString)

}

func ChangeColor(s *string) {
	newColor := "Red"
	*s = newColor

}
