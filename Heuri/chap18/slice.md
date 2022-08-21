## Slice

동적 배열인 슬라이스는 초기화하지 않으면 길이가 0으로 생성

```go
var slice1 = []int{1, 2, 3}
var slice2 = []int{1, 5:2, 10:3}   // 5번째 인덱스는 2, 10번째 인덱스는 3

// Using make()
var slice = make([]int, 3)  // 길이 3개 짜리 int 슬라이스, 각 요솟값은 0
```

> `var array = [...]int{1, 2, 3}` : 길이가 3인 고정 배열 선언

### 여러 값 추가

```go
slice = append(slice, 3, 4, 5, 6, 7)    // append()를 추가하여 값을 하나 이상 추가
```

## 슬라이스 동작 원리

```go
type SliceHeader struct {
    Data uintptr    // 실제 배열을 가리키는 포인터
    Len int         // 요소 개수
    Cap int         // 실제 배열의 길이
}
```

> `var slice = make([]int, 3, 5)` 로 선언할 경우, 배열 길이는 5, 요소 길이는 3이므로 앞 인덱스 2까지는 0 3개, 나머지 2개의 공간은 빈 공간이 생성
