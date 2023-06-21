package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays on Go")
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Mango"
	fruits[3] = "Orange"

	fmt.Println("Fruit list is", fruits)
	fmt.Println("Fruit list is", len(fruits))

	var vegerables = [3]string{"potato", "tomato", "brinjal"}
	fmt.Println("Vegerables list is", vegerables)
}
