package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices on Go")

	// needs to initialize the list
	fruits := []string{"Apple", "Banana", "Mango", "Orange"}
	fmt.Println("Fruit list is", fruits)
	fmt.Printf("Type of fruits is %T\n", fruits)

	// Add  new value to the list
	fruits = append(fruits, "Grape")
	fmt.Println("Fruit list  now is", fruits)

	// Replace a value from the list
	fruits = append(fruits[1:], "")
	fmt.Println("Fruit list  now is", fruits)

	// Replace a value from the list
	fruits = append(fruits[1:3], "")
	fmt.Println("Fruit list  now is", fruits)

	// Replace a value from the list
	fruits = append(fruits[:3], "")
	fmt.Println("Fruit list  now is", fruits)

	fmt.Println("-----------------------------")

	higScores := make([]int, 4)
	higScores[0] = 234
	higScores[1] = 100
	fmt.Println("HigScores list  now is", higScores)

	// Realocate memory
	higScores = append(higScores, 234, 566, 777)
	fmt.Println("HigScores list  now is", higScores)
	fmt.Println("HigScores list is sorted?", sort.IntsAreSorted(higScores))

	// Sorting in ascending order
	sort.Ints(higScores)
	fmt.Println("HigScores list is sorted", higScores)

	// Sorting in descending order
	sort.IntsAreSorted(higScores)
	fmt.Println("HigScores list is sorted?", sort.IntsAreSorted(higScores))

	fmt.Println("-----------------------------")

	// How to remove a value from slices based on index
	courses := []string{"reactjs", "javascript", "swift", "python", "ruby", "go"}
	fmt.Println("Courses list is", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses list is", courses)

}
