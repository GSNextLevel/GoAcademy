# fmt

1.

```go
package main

import "fmt"

func main() {
	var a = 345
	var b = 3.1415

	fmt.Printf("%05d\n", a)  //00345
	fmt.Printf("%5.2f\n", b) // 3.14
}
```

2.

```go
package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scanln(&a, &b) //&
	fmt.Println(a, b)
}
```

3.

```go
package main

import "fmt"

func main() {
	var a = 123
	var b int = 4567
	f := 3.14159269

	fmt.Printf("%6d\n", a)
	fmt.Printf("%06d\n", b)
	fmt.Printf("%6.2f\n", f)
}
```
