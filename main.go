package main

import "log"

func main() {

	word2, word1 := SaySomething("Hello Dayo. ", "How old are you?")
	log.Println(word1,"\n", word2)
// Testing branches

}

// testing second commit

func SaySomething(val1 string, val2 string) (string, string) {
	return val1, val2
}
