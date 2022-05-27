package main

import "fmt"

const LoginToken string = "adggdewg" // Public

func main() {

	var username string = "mustafa"
	fmt.Println("username")
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 223.42424424
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)

	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var website = "hello"
	fmt.Println(website)

	word := 3000
	fmt.Println(word)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
