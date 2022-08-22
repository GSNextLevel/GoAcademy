package main
import "fmt"

func ToUpper(str string) string {
	var builder strings.Builder

	for _, v := range str {
		if v >='a' && v <='z' {
			builder.WriteRune('A'+(c-'a'))
		}else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
func main() {
	str := "hello World!"
	fmt.Println(ToUppper(str))
}
