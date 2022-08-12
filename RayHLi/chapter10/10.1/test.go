package main

import "fmt"

func main() {
	day :=1

	switch day {
	case 1:
		fmt.Println("First Day")
	case 2:
		fmt.Println("Second Day")
	default:
		fmt.Println("Another day")
	}
}
