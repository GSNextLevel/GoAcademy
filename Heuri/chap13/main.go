package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

type Monster struct {
	Actor
	Attack int
	Speed  int
}

func main() {
	var monster = Monster{
		Actor{"NPCA", 100, 8.7},
		500,
		200,
	}

	fmt.Println(monster.Speed)
	fmt.Println(monster.Actor.Speed)
}
