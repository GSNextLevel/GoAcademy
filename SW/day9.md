# for 

1. 
```go
package main

import "fmt"

func main() {
    for i := 10; i > 0; i-- {
        fmt.Print(i, " ")
    }
}
```

2. 
```go
package main

import "fmt"

func main() {
    for i := 2; i < 10; i++ {
        if i >= 3 && i <= 6 {
            continue
        } 
        for j := 1; j < 10; j++ {
            fmt.Println(i, "*", j, "=", i*j)
        }
        fmt.Println()
    }    
}
```

3. 
```go
package main

import "fmt"

func main() {
    for i := 1; i < 10; i++ {
        if i % 2 == 0 {
            continue
        }
        for j := 1; j < 10; j++ {
            if i == j {
                fmt.Println(i, "*", j, "=", i*j)
            }
        } 
    }
}
```

4.
```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        for j := 0; j < 5-i ; j++ {
            fmt.Print("*")
        } 
        fmt.Println()
    }
}
```
