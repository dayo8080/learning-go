package main

import (
	"log"

	"github.com/dayo8080/learning-golang/helpers"
)

const numPool = 9999

func Calculatevalue(intChannel chan int) {
	randomNumber := helpers.RandomNumber(numPool)

	intChannel <- randomNumber
}

func main() {

	intChannel := make(chan int)
	defer close(intChannel)

	go Calculatevalue(intChannel)

	num := <- intChannel
	log.Println(num)
}
