package main
import (
	"fmt"
	"unsafe"
)

type User struct {
	FirstName string
	LastName string
	Age int
}

func NewUser(firstname string, lastname string, age int) *User {
	var a = User{firstname, lastname, age}
	return &a
}

func main() {
	var a = User{}
	var user = NewUser("AA","BB",12)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(user))
}
