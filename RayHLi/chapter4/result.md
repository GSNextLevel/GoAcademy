# 연습문제
## 1. 다음 a,b,c,d,e 변수 타입 확인
##### a : = 3 (int)
##### var b = 3.1415 (float64)
- 타입 확인 방안 (fmt.Println(reflect.TypeOf(x)) - package(reflect)
##### c := "hello world" (string)
##### d := int32(10) (int32)
##### var e float32 = 3.1415 (float32)

## 2. 출력 확인

##### int8 : 1바이트 부호 있는 정수 (-128~127)
##### int32: 4바이트 부호 있는 정수 (-2147483648 ~ 2147483647)
-> 바이트 범위가 벗어나기에 360은 이진수로 (1)01101000이며 괄호 부분이 오버플로우로 사라지기에 10진수로 104가 출력됨

## 3. f1, f2 다른 값 이유

##### f1의 경우 (42707.406)927942이 나오며 변수형 타입에 의해서 (4바이트 실수) 42707.406이 되며 f2의 경우 123.546789 -> 123.54679에서 345.678을 곱하여 42.707.40727362가 되었는데 역시 동일하게 변수타입에 의해 42707.41이 되었음
-> reflect.TypeOf(m)으로 확인
