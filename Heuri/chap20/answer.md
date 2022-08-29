## Q1

File 타입은 `Read()` 메서드만 가지고 있으므로, `Read()` & `Write()` 두 가지 메서드를 가진 `ReadWrite` 인터페이스를 사용할 수 없다.

## Q2

```go
type OurDB interface {
    GetData() string
    WriteData(data string)
    Close() error
}
```

## Q3

```go
func CheckAndRun(stringer Stringer) {
    if r, ok := stringer.(Reader); ok {
        r.Read()
    }
}
```
