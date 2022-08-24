package main
import (
	"fmt"
	"math/rand"
	"time"
)

const (
	My_money = 1000
	Add_money = 500
	Remove_money = 100
)

func input_func() (int){
	var n int
	fmt.Scanln(&n)
	return n
}

func main() {
	rand.Seed(time.Now().UnixNano())
	my_money := My_money

	for {
		rand := rand.Intn(5)+1
		n := input_func()
		fmt.Println("input: ",n," random: ",rand)

		if(n == rand) {
			my_money += 500
			fmt.Println("Congratulation ", my_money)
			if(my_money>=5000) {
				fmt.Println("Bye! 5000")
				break;
			}
		} else {
			my_money -= 100
			fmt.Println("Sorry ",my_money)
			if(my_money=<0) {
				fmt.Println("Sorry Bye!")
				break;
			}
		}
	}
}
