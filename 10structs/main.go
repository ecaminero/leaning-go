package main

import "fmt"

type User struct {
	Name     string
	LastName string
	Email    string
	Age      int
}

func main() {
	fmt.Println("Learning Structs on go")
	// No inheritance in go, but we can use composition, No super or parent

	ana := User{"Ana", "Lekshova", "anal@mail.io", 20}
	fmt.Println(ana)
	fmt.Printf("Ana's details are %+v\n", ana)
	fmt.Printf("Name is %v and email is %v\n", ana.Name, ana.Email)
}
