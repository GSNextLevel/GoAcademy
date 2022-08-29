package main
import "fmt"

func main() {
	array := [5]int{1,2,3,4,5}
	slice:=array[1:2]

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100

	fmt.Println("array:", array)
	slice = append(slice, 500)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}