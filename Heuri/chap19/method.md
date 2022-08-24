## 메서드

method는 함수의 일종으로 Go 언어에서는 클래스가 없으므로 구조체 밖에 메서드를 지정

메서드 명 : `info()` / 리시버 : `(r Rabbit)`

```go
func (r Rabbit) info() int {
    return r.width * r.height
}
```

리시버 덕분에 `info()` 메서드는 Rabbit 타입이며, 구조체 변수 r은 해당 메서드에서 매개변수 처럼 사용된다.

> 함수와 달리, 메서드는 리시버에 속하므로 데이터와 관련 기능을 묶어 코드의 응집도를 높일 수 있음.

```go
package main

import "fmt"

type account struct {
    balance int
}

// Function
func withdrawFunc(a *account, amount int) {
    a.balance -= amount
}

// Method
func (a *account) withdrawMethod(amount int) {
    a.balance -= amount
}

func main() {
    a := &account{ 100 }    // balance가 100인 account 포인터 변수 생성
    withdrawFunc(a, 30)     // Using Function
    a.withdrawMethod(30)    // Using Method
    fmt.Printf("%d \n", a.balance)  // 40
}
```

## 별칭 리시버 타입

모든 로컬 타입이 리시버 타입으로 가능하기 때문에 별칭 타입도 리시버가 될 수 있고 메서드를 가질 수 있습니다.

```go
package main

import "fmt"

type myInt int      // 사용자 정의 별칭 타입

func (a myInt) add(b int) int {
    return int(a) + b
}

func main() {
    var a myInt = 10
    fmt.Println(a.add(30))      // 40
    var b int = 20
    fmt.Println(myInt(b).add(50))   // 70, int 타입을 타입 변환
}
```

## 포인터 메서드 vs 값 타입 메서드

포인터 메서드는 내부에서 리시버의 값을 변경시킬 수 있다. 반면 값 타입 메서드는 호출쪽과 메서드 내부의 값은 별도 인스턴스로 독립되어 리시버의 값을 변경시킬 수 없다.
