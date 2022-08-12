# Array

1. 
```go
package main

import "fmt"

func main() {
    a := [5]int{1,2,3,4,5}

    for i, v := range a {
        a[i] = v * 2
    }

    fmt.Println(a[2])
}
```
- 3*2 = 6

2. 
- float64 -> 8 byte
- 3 * 2 * 5 * 8 = 240 byte

3. 
```go
package main

import "fmt"

func ChangeArray(arr [5]int) {
    arr[3] = 3000
}

func main() {
    a := [5]int{1,2,3,4,5}
    ChangeArray(a)
    fmt.Println(a[3])
}
```
- a[3] 값은 그대로 4 (복사만 되고 return되지 않음)
- 3000으로 넣어주려면 아래처럼
```go
func ChangeArray(arr [5]int) [5]int {
    arr[3] = 3000
    return arr
}

func main() {
    a := [5]int{1,2,3,4,5}
    a = ChangeArray(a)
    fmt.Println(a[3])
}
```
