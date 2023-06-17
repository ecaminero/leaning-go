package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule func(key string, value interface{}) error
type Rules []Rule

func ValidateRange(min, max int) Rule {
	return func(key string, value interface{}) error {
		num, ok := value.(float64)
		if !ok {
			fmt.Printf("Variable is type %T \n", num)
		}
		if num < float64(min) || num > float64(max) {
			return fmt.Errorf("%s must be between %d and %d", key, min, max)
		}
		return nil
	}
}

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("invalid input. It should be a number, got this error instead")
		fmt.Println(err)
		os.Exit(1)
	}

	// Execute validation rules and get errors
	err = ValidateRange(1, 5)("rating", numRating)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Thanks for rating our pizza")
	fmt.Println("Added 1 to your rating: ", numRating+1)

}
