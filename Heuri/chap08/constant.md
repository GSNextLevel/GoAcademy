## Constants

Constants are also called literals.

## iota

Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants.

```go
const (
    Red     int = iota      // 0
    Blue    int = iota      // 1
    Green   int = iota      // 2
)

**iota**는 소괄호를 벗어나면 다시 초기화

const (
    C1 uint = iota + 1      // 1 = 0 + 1
    C2                      // 2 = 1 + 1
    C3                      // 3 = 2 + 1
)
```
