# Pointer

- 초기화 방법
  - `p1 := $Data{}`
  - `var p2 = new(Data)`

1.

```go
package main

import "fmt"

func add(p1, p2, p3 *int) {
	*p3 = *p1 + *p2
}

func main() {
	a := 3
	b := 5
	c := 0

	add(&a, &b, &c)
	fmt.Println(c) // 8
}
```

2.

```go
package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	return &Actor{name, hp, speed}
}

func main() {
	var actor = NewActor("Gold Rabbit", 99, 100)
	fmt.Println(actor.Speed) // 100
	fmt.Println(actor.Name) // Gold Rabbit
}
```

3.

```go
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	newUser := NewUser("AAA", 23) // 1개 생성
	var p *User = newUser // 해당 인스턴스 가리키는 포인터
	p.Age += 10

	fmt.Println(p.Age)
}
```
