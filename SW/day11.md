# Structure

1.

```go
type Product struct {
    Name string
    Price int
    RewiewScore float64
}
```

2.

```go
package main

import "fmt"

type Actor struct {
	Name  stirng
	HP    int
	Speed float64
}

type Monster struct {
	Actor
	Attack int
	Speed  int
}

func main() {
	var monster = Monster{
		Actor{"NPCA", 100, 8.7},
		500,
		200,
	}
	fmt.Println(monster.Speed)       // 200
	fmt.Println(monster.Actor.Speed) // 8.7
}
```

3.

```go
type Padding struct {
    A int8 // 1
    B int // 8
    C float64 // 8
    D uint16 // 2
    E int // 4
    F float32 // 4
    G int // 4
}
```

- 총 32비트(8/8/8/8)

```go
type Padding struct {
    A int8 // 1
    G int8 // 1
    D uint16 // 2
    F float32 // 4

    B int // 8
    C float64 // 8
    E int //  8
}
```

- 7/8/8/8 로 패딩 최적화
