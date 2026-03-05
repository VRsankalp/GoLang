package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {
	case time.Saturday , time.Sunday:
		fmt.Println("weekday")
	default:
		fmt.Println("it is weoking")
	}
}