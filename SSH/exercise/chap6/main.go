package main

import "fmt"

// func main() {
// 	var a int8 = 30 // 00011110

// 	a <<= 2 // 01111000 = 8+16+32+64 = 120
// 	a += 8  // 128, 2^8-1 를 넘어가므로 반전

// 	fmt.Println(a)

// }

// func main() {
// 	var a uint8    // go의 특징 var로 변수 지정하고 init 값을 안주면 auto value. 숫자는 0, 문자열은 ''
// 	fmt.Println(a) // 0

// 	a |= 2
// 	fmt.Println(a) // 2

// 	a |= 4
// 	fmt.Println(a) // 6

// 	a |= 8
// 	fmt.Println(a) // 14

// 	var b uint8
// 	b = 4

// 	a &^= b
// 	fmt.Println(a) // exclusive 한 값을 and 연산 10

// }

func main() {
	var a uint8    // go의 특징 var로 변수 지정하고 init 값을 안주면 auto value. 숫자는 0, 문자열은 ''
	fmt.Println(a) // 0

	a |= 2
	fmt.Println(a) // 2

	a |= 4
	fmt.Println(a) // 6

	a |= 8
	fmt.Println(a) // 14

	var b uint8
	b = 4

	a &^= b
	fmt.Println(a) // exclusive 한 값을 and 연산 10

}
