# Function

1.

```go
package main

import "fmt"

func Multiple(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(Multiple(1, 2))
}
```

2.

```go
package main

import "fmt"

func AAA() {
	fmt.Println("start AAA()")
	BBB()
	fmt.Println("end AAA()")
}

func BBB() {
	fmt.Println("BBB()")
}

func main() {
	AAA()
}

// start AAA()
// BBB()
// end AAA()
```

3.

```go
package main

import "fmt"

func F(n int) int {
	if n <= 1 {
		return n
	}
	return F(n-2) + F(n-1)
}

func main() {
	fmt.Println(F(9))
}

```
