package publicpkg

import "fmt"

const (
	PI = 3.1415
	pi = 3.141516
)

var ScreenSize int = 1080 // Public Var
var screenHeight int      // Private Var

func PublicFunc() {
	const MyConst = 100 // 함수 내부에서 선언되었기 때문에 패키지 외부로 공개 X
	fmt.Println("This is a public function", MyConst)
}

func privateFunc() {
	fmt.Println("This is a private function")
}

type MyInt int       // Public Alias Type
type myString string // Non-public Alias Type

type MyStruct struct {
	Age  int    // Public Struct Field
	name string // Non-public struct Field
}

func (m MyStruct) PublicMethod() {
	fmt.Println("This is a public method")
}

func (m MyStruct) privateMethod() {
	fmt.Println("This is a private method")
}

type myPrivateStruct struct {
	Age  int
	name string
}

func (m myPrivateStruct) PrivateMethod() {
	fmt.Println("This is a private method")
}
