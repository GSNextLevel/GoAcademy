package main

import (
	"fmt"
)

func main() {
	var a = 345
	var b = 3.1415

	fmt.Printf("%05d\n", a)
	fmt.Printf("%5.2f\n", b)
}

// func main() {
// 	var a int
// 	var b int

// 	fmt.Scanln(&a, &b)
// 	fmt.Println(a, b)
// }
