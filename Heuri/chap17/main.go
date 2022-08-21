package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Balance       = 1000
	EarnPoint     = 500
	LosePoint     = 100
	VictoryPoint  = 5000
	GameoverPoint = 0
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
	rand.Seed(time.Now().UnixNano())
	balance := Balance

	for {
		fmt.Println("1 ~ 5 사이의 숫자를 입력해주세요.")
		n, err := InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else if n < 1 || n > 5 {
			fmt.Println("1~5 사이의 값만 입력하세요.")
		} else {
			r := rand.Intn(5) + 1
			if n == r {
				balance += EarnPoint
				fmt.Println("Success !")
				if balance >= VictoryPoint {
					fmt.Println("Win !")
					break
				}
			} else {
				balance -= LosePoint
				fmt.Println("Fail !")
				if balance <= GameoverPoint {
					fmt.Println("Lose !")
					break
				}
			}
		}
	}
}
