package main

import "fmt"

func main() {
	fmt.Println("HELLO")

	// no inheritance in golang: No Super or parents
	sus := User{"Mustafa", "sus@sus.com", true, 17}
	fmt.Println(sus)
	fmt.Printf("sus details are : %+v\n", sus)
	fmt.Printf("Name is : %v\n", sus.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


