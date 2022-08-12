## Q1

```go
package main

import "fmt"

func main() {
	n := 10
	for i := n; i > 0; i-- {
		fmt.Printf("%d ", i)
	}
}

```

## Q2

```go
if i > 2 && i < 7 {
    continue
}
```

## Q3

```go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i, "*", i, "=", i*i)
		}
	}
}

```

## Q4

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

```
