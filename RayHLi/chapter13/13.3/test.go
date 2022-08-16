package main
import (
	"fmt"
	"unsafe"
)

type Padding struct {
	A int8
	G int8
	D uint16
	F float32
	B int
	C float64
	E int
}

func main() {
	var test Padding = Padding{1,2,3,1.1,4,1.2,5}
	fmt.Println("A(int8): ",&test.A, "G(int8): ",&test.G, "D(uint16): ",&test.D, "F(float32): ",&test.F, "B(int): ",&test.B, "C(float64): ",&test.C, "E(int): ",&test.E)
	fmt.Println(unsafe.Sizeof(test))
}
