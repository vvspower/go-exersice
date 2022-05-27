package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	fmt.Println("sus")
	defer fmt.Println("\ntwo")

	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
