package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	var diceNumber int = rand.Intn(6) + 1 // random number

	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
		case

	}

}
