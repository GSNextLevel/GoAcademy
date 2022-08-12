package main

import "fmt"

func main() {
  score := 78
  count := 20

  if count < 10 {
	  fmt.Println("평가 수가 모자랍니다.")
  } else if count < 20 {
	  if score > 80 {
		  fmt.Println("긍정적인 평가")
	  } else {
		  fmt.Println("판단할 단계가 아닙니다.")
	  }
  } else {
	  if score > 90 {
		  fmt.Println("Good Score")
	  } else if score > 80 {
		  fmt.Println("살만한 물건입니다.")
	  } else if score > 70 {
		  fmt.Println("신중히 생각하세요.")
	  } else {
		  fmt.Println("좋은 물건이 아닙니다.")
	  }
  }
}
