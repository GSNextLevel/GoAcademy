package main

import "fmt"

// func Multiple(a int, b int) int {
// 	return a * b
// }

// func main() {
// 	var result int

// 	result = Multiple(3, 4)

// 	fmt.Println(result)

// }

func F(n int) int {
	if n < 2 {
		return n
	}

	return F(n-2) + F(n-1)
}

func main() {
	fmt.Println(F(9))
}
