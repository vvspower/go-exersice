package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())

	var diceNumber int = rand.Intn(6) + 1
	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice val is 1 and you can open")
		fallthrough
	case 2:
		fmt.Println("You can move 2 spot")

	}

}
