package main

import "fmt"

func main() {
	fmt.Println("if else statements")
	var result string
	var LoginCount int = 0

	if LoginCount > 10 {
		result = "Regular user"
	} else {
		result = "New User"
	}

	var number int = 1423445

	if number%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	}

	fmt.Println(result)

}
