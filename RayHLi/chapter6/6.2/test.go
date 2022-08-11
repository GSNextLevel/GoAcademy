package main
import "fmt"

func main() {
  var a uint8
  a |= 2
  fmt.Println(a)
  a |= 4
  fmt.Println(a)
  a |= 8
  fmt.Println(a)

  var b uint8
  b = 4

  a&^=b
  fmt.Println(a)
}
