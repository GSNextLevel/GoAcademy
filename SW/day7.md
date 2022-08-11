# If
- if 초기문, 조건문 
    - ex> `if a,b := test(), a` : test()의 return 값이 a,b에 할당되며, a가 true이면 true

1.
```go
package main

import "fmt"

func main() {
    age := 22

    if age < 10 {
        fmt.Println("Child")
    } else if age >= 20 && age < 30 {
        fmt.Println("Good Age") {}
    } else {
        fmt.Println("Grown")
    }
}
```

2. 
else if score > 70 {
    fmt.Println("신중히 생각해보세요")
}
에 해당

3. 
```go
package main

import "fmt"

func main() {
    temp := 30
    rain := 40

    if temp >= 25 {
        if rain >= 80 {
            fmt.Println("Hot and Rainy")
        } else if rain >= 20 {
            fmt.Println("Hot and Humid")
        } else {
            fmt.Println("Good to go out")
        }
    } else if temp < 10 || rain >= 80 {
        fmt.Println("Bad to go out")
    } else {
        fmt.Println("Good weather")
    }
}

```
