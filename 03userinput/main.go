package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to my playground!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || error ok
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", rating)
}
