## 1
const Gravity = 9.80665

## 2
iota 로 순차적으로 증가하는 상수를 관리할 수 있다.

응용 예제

const (
    _ = iota
    KB ByteSize = 1 << (10 * iota)
    MB ByteSize = 1 << (10 * iota)
    GB ByteSize = 1 << (10 * iota)
    TB ByteSize = 1 << (10 * iota)
)