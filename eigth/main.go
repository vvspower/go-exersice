package main

import "fmt"

func main() {
	fmt.Println("HELLO")

	var fruitlist [4]string

	fruitlist[0] = "tomato"
	fruitlist[1] = "apple"
	fruitlist[3] = "potato"

	fmt.Println("Fruit list:", len(fruitlist))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggies:", len(vegList))

}
