package main

import (
	"bufio" //input from keyboard
	"fmt" //standard output functions
	"os" 
	"strconv" //conversion of console input from string to desired data type
)

func main() {

	fmt.Println("Hello world", "hi")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name, _ := strconv.ParseInt(sc.Text(), 10, 64)
	fmt.Println("You are ", 2022-name, " years old")

}
