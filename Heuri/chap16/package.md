## Package

Go 언어에서 코드를 묶는 가장 큰 단위

- main : 프로그램 시작점(`main()` 함수)을 포함한 패키지.
- [표준 패키지](https://goloang.org/pkg/)
- [유용한 패키지](https://github.com/avelino/awesome-go/)

## Import된 패키지 찾는 방법

1. Go 설치 경로의 기본 패키지
2. 외부 저장소의 패키지의 경우 `GOPATH\pkg` 폴더에 설치
3. 현재 모듈 아래 위치한 패키지

## 겹치는 Package Alias 사용하기

마지막 폴더명이 같은 패키지

```go
import (
    "text/template"
    "html/template"
)
```

Alias 부여 후 호출 방법

```go
import (
    "text/template"
    htemplate "html/template"
)

template.New("foo").Parse(`{{define "T"}}Hello`)
htemplate.New("foo").Parse(`{{define "T"}}Hello`)
```

## 사용하지 않는 패키지 포함하기

패키지를 직접 사용하지 않지만 부가효과를 얻고자 임포트하는 경우

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)
```

> 위 코드는 `init()` 함수 기능만 사용하기 원할 경우 밑줄 \_을 이용해서 임포트

### 패키지 초기화

패키지를 임포트하면 컴파일러는 패키지 내 전역 변수를 초기화한다. 그다음 패키지에 `init()` 함수가 있다면 호출해 패키지를 초기화

> `init()` : 입력 매개변수가 없고 반환값도 없는 함수

## Go 모듈

`go mod tidy` Go 모듈에 필요한 패키지를 찾아서 다운로드해주고 필요한 패키지 정보를 `go.mod` 파일과 `go.sum` 파일에 기재
