package main

import "fmt"

func main() {
	var number int
	number = 265
	number += 5
	var out = fmt.Sprint(10)
	fmt.Println("Hello world", number+number, "hi", out)
	fmt.Printf("%T", out)
}
