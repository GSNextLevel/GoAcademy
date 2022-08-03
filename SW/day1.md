 1. go는 2009년에 공개 / 실행 전 컴파일 필요(go build) -> go run 으로 실행했을 때는 실행파일 생성되지 않으며 테스트 용도로 사용

 2. 

- `hello2.go` 
 ```go
 package main

 import "fmt"

 func main() {
   fmt.Println("Hello, Golang")
}
```

- `go mod init goproject/hello2`
- `go build hello2.go`
- `./hello2`
