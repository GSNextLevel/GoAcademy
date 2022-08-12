package main

import "fmt"

func Mutiple(a, b int) int {
  return a*b
}

func main() {
	var rel int = Mutiple(12, 3)
	fmt.Println(rel)
}
