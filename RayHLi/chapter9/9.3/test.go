package main
import "fmt"

func main() {
  temp := 30
  rain := 40

  if temp > 25 {
	  if rain > 80 {
		  fmt.Println("덥고 비옴.")
	  } else if rain > 20 {
		  fmt.Println("덥고 습함")
	  } else {
		  fmt.Println("좋음")
	  }
  } else if temp < 10 || rain > 80 {
	  fmt.Println("야외 활동 안좋음")
  } else {
	  fmt.Println("좋은 날씨")
  }
}
