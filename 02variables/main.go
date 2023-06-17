package main

import "fmt"

const LoginToken = "A-LOGIN-TOKEN"

func main() {
	var username string = "John Doe"
	fmt.Println("Hello: ", username)
	fmt.Printf("Variable is type %T \n", username)
	fmt.Println("-------------------------")

	// bool
	var isLooged bool = true
	fmt.Println("Logged in?:", isLooged)
	fmt.Printf("Variable is type %T \n", isLooged)
	fmt.Println("-------------------------")

	// int
	var smallValue int = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is type %T \n", smallValue)
	fmt.Println("-------------------------")

	// float32
	var smallFloat float32 = 255.6066440002
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type %T \n", smallValue)
	fmt.Println("-------------------------")

	// default Int value
	var defaultIntValue int
	fmt.Println(defaultIntValue)
	fmt.Printf("Variable is type %T \n", defaultIntValue)
	fmt.Println("-------------------------")

	// implicit type
	var implicit = "Implicit"
	fmt.Printf("Variable is type %T \n", implicit)
	fmt.Println("-------------------------")

	// variable no style
	numberOfUsers := 300000
	fmt.Println("Number of user:", numberOfUsers)
	fmt.Println("-------------------------")

	// const
	fmt.Println(LoginToken)

}
