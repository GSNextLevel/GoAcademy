# Operator

1.

```go
package main

import "fmt"

func main() {
	var a int8 = 30

	a <<= 2        // 30 * 2^2 = 120
	a += 8         // 120 + 8 = 128
	fmt.Println(a) // int8 -> -128 ~ 127
}
```

2.

```go
package main

import "fmt"

func main() {
	var a uint8 // default a=0
    a |= 2 // 0 | 2 = 2
	a |= 4 // 2 | 4 = 6
	a |= 8 // 6 | 8 = 14 -> 0000 1110

	var b uint8
	b = 4 // 0000 0100

	a &^= b // 00001110 & 11111011 = 00001010 -> 10
	fmt.Println(a)
}
```

3.

```go
package main

import "fmt"

func main() {
	var x int8 = 1 // 0000 0001
	x <<= 7        // 1000 0000 -> -128 (int8 : -128 ~ 127)
	x >>= 7        // 1111 1111 -> -1
	fmt.Printf("%d\n", x)
}
```
