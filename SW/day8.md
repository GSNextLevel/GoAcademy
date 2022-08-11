# Switch

1. 
```go
package main

import "fmt"

func main() {
    day := 1

    switch day {
    case 1:
        fmt.Println("First day")
    case 2:
        fmt.Println("Second day")
    default:
        fmt.Println("Another day")
    }
}
```

2. 
case count < 20 && score > 80:
    fmt.Println("긍정적인 평가")
에 해당됨

3.
```go
func GetDirection(angle float64) Direction {
    switch angle {
    case angle > 315:
        return "North"
    case angle >= 0 && angle < 45:
        return "North"
    case angle >= 45 && angle < 135:
        return "East"
    case angle > 135 && angle < 225:
        return "South"
    case angle >= 225 && angle < 315:
        return "West"
    default:
        return "None"
    }
}
```
