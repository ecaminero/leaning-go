package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")
	myNumber := 2
	var emptyPointer *int
	var ptr = &myNumber

	fmt.Println("Empty Pointer", emptyPointer)

	fmt.Println("Memory Access (myNumber)", ptr)
	fmt.Println("Value of MyNumber", *ptr)
	*ptr = *ptr + 2
	fmt.Println("new value ", *ptr)

}
