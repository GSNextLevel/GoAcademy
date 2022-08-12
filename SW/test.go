package main

import "fmt"

func main() {
    temp := 11
    rain := 25

    if temp >= 25 {
        if rain >= 80 {
            fmt.Println("Hot and Rainy")
        } else if rain >= 20 {
            fmt.Println("Hot and Humid")
        } else {
            fmt.Println("Good to go out")
        }
    } else if temp < 10 || rain >= 80 {
        fmt.Println("Bad to go out")
    } else {
        fmt.Println("Good weather")
    }
}
