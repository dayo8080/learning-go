package main

import "log"

func main() {

	word2, word1 := SaySomething("Hello Dayo. ", "How old are you?")
	log.Println(word1,"\n", word2)

	// saySomethingEsle = SaySomething("Good Bye!", "You are mad")
	// log.Println(saySomethingEsle)

	// log.Println(SaySomething("Finally!!!"))
	// log.Println(i)

}

func SaySomething(val1 string, val2 string) (string, string) {
	return val1, val2
}
