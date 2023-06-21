package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning Time format")
	presentTime := time.Now()
	fmt.Println("Today Is:", presentTime)

	// Format time
	// Documentation: https://pkg.go.dev/time#pkg-constants
	fmt.Println("Today Is:", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Custom Date
	customDate := time.Date(2022, time.April, 12, 12, 12, 12, 12, time.UTC)
	fmt.Println("Custom Date:", customDate)
	fmt.Println("Custom Date:", customDate.Format("01-02-2006 Monday"))

}
