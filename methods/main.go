package main

import "fmt"

func main() {
	fmt.Println("Methods")
	sus := User{"Mustafa", "sus@sus.com", true, 17}
	sus.GetStatus()
	sus.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:  ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
