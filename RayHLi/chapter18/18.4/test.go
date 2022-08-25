package main
import (
	"fmt"
	"sort"
)

type Player struct {
	name string
	age int
	score int
	pass float32
}

type Players []Player

func(s Players) Len()  int {return len(s)}
func(s Players) Less(i, j int) bool {return s[i].score < s[j].score}
func(s Players) Swap(i, j int) {s[i], s[j] = s[j], s[i]}

func main() {
	s:=[]Player{
		{"A", 13, 45, 78.4},{"B", 16, 24, 67.4},{"C", 18, 54, 50.8},{"D", 16, 46, 89.7},
	}
	sort.Sort(sort.Reverse(Players(s)))
	fmt.Println(s)
}
