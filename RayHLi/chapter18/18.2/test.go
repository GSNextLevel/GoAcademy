package main
import "fmt"

func main() {
	slice := []int{1,2,3,4,5,6,7,8,9,10}
	slice1 := slice[8:]

	fmt.Println(slice1)
}
