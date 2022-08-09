package main

import "fmt"

func main() {
	var x int8 = 1 // 0000 0001
	x <<= 7        // 1000 0000
	x >>= 7        // 1111 1111

	fmt.Printf("%d\n", x) // -1
}
