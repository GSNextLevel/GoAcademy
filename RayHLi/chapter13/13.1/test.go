package main
import "fmt"

type Product struct {
	Name string
	Price int
	ReviewScore float64
}

func main() {
	var test Product = Product{"hello", 120, 12.4}

	fmt.Println(test.Name, test.Price, test.ReviewScore)
}
