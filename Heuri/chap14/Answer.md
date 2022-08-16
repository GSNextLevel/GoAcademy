## Q1

8

## Q2

```go
// My Answer
func NewActor(name string, hp int, speed float64) *Actor {
	var u = Actor{name, hp, speed}
	return &u
}

// Another Answer
func NewActor(name string, hp int, speed float64) *Actor {
	return &Actor{name, hp, speed}
}

func NewActor(name string, hp int, speed float64) *Actor {
	a := new(Actor)
    a.Name = name
    a.Hp = hp
    a.Speed = speed
	return a
}
```

## Q3

1개
