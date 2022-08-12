package main

import "fmt"

func main() {
  score := 85
  count := 15
  switch {
  case count < 10:
      fmt.Println("평가 수가 모자릅니다.")
  case count < 20 && score > 80:
      fmt.Println("긍정적인 평가")
  case count < 20:
      fmt.Println("판단할 단계 아님")
  case score > 90:
      fmt.Println("좋은 평가")
  case score > 80:
      fmt.Println("살만한 물건")
  case score > 70:
	  fmt.Println("신중히 생각")
  default:
	  fmt.Println("좋은 물건이 아님")
  }
}
