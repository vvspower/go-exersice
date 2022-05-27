package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["Rb"] = "Ruby"
	languages["Go"] = "GoLang"

	fmt.Println(languages)
	fmt.Println("JS short for:", languages["JS"])

	delete(languages, "Go")
	fmt.Println(languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
