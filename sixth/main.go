package main

import "fmt"

func main() {
	// var pointer *int
	// fmt.Println("Value of pointer is :", pointer)

	var myNumber int = 23
	var pointer = &myNumber

	fmt.Println("value of actual ppointer is : ", pointer)
	fmt.Println("value of actual pointer is: ", *pointer)

	*pointer = *pointer * 2

	fmt.Println("New value is ", myNumber)
}
