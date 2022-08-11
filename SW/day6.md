# constant

- const 로 선언
- iota 증가하는  상수값일 때 (0부터 시작하여 +1씩 증가됨)

1. 
```go
const Gravity = 9.80665
```

2. 
```go
package main

import "fmt"

const (
    C1 = iota // 0
    C2 // 1
    C3 // 2
)

const (
    D1 = iota + 1 // 1
    D2 // 2
    D3 // 3
)

func main() {
    fmt.Println(C3, D3) // 2 3
}
```
