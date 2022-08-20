package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	var seedMoney int = 1000

	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(5) + 1

	for {
		fmt.Println("1 ~ 5 사이의 숫자를 입력해주세요.")
		n, err := InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if n == answer {
				seedMoney += 500
				fmt.Println("Success !")
				if seedMoney >= 5000 {
					fmt.Println("Win !")
					break
				}
			} else {
				seedMoney -= 500
				fmt.Println("Fail !")
				if seedMoney <= 0 {
					fmt.Println("Lose !")
					break
				}
			}
		}
	}
}
