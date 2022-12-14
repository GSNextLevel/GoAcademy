package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

// Pointer Method
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// Value Type Method
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance) // 70

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance) // 70

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance) // 50

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance) // 20
}
