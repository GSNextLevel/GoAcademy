## Q1

30 X 2 X 2 + 8 = 128, overflow 발생해 -128

## Q2

```go
package main

import "fmt"

func main() {
	var a uint8

	a |= 2		// 0000 0010
	a |= 4		// 0000 0010 | 0000 0100 => 0000 0110
	a |= 8		// 0000 0110 | 0000 1000 => 0000 1110 (14)

	var b uint8
	b = 4		// b : 0000 0100

	a &^= b		// 0000 1010 (3번째 자리 클리어) => 10
	fmt.Println(a)
}
```

## Q3

```go
package main

import "fmt"

func main() {
	var x int8 = 1	// 0000 0001
	x <<= 7			// 1000 0000
	x >>= 7			// 1111 1111

	fmt.Printf("%d\n", x)	// -1
}
```
