## 인스턴스

메모리에 할당된 데이터의 실체

가리키는 포인터 변수 개수는 인스턴스 개수와 무관

```go
var p1 *Data = &Data{}  // 인스턴스를 생성해 해당 주소를 포인터 변수의 초깃값으로 대입
var p2 *Data = p1       // p2도 Data 인스턴를 가리키도록
var p3 *Data = p1
```

> 쓸모 없어진 인스턴스는 가비지 컬렉터가 자동으로 지워준다.

## new() 내장 함수

`new()` 함수는 인수를 타입으로 받아, 타입을 메모리에 할당하고 기본값으로 채워 그 주소를 반환.
new를 이용해서 내부 필드값을 원하는 값으로 초기화할 수는 없다.

```go
p1 := &Data{}       // &를 사용하는 초기화
var p2 = new(Data)  // new()를 사용하는 초기화
```

> &를 사용하는 초기화는 `p1 := &Data{3, 4}` 와 같은 방식으로 초기화가 가능

## Stack Memory & Heap Memory

Go언어는 **escape analysis**를 통해 어느 메모리에 할당할지 결정.

함수 내부에서 선언된 변수는 함수가 종료되면 사라지지만, **escape analysis**를 통해 해당 변수가 외부로 공개된다면 **힙 메모리**에 할당.