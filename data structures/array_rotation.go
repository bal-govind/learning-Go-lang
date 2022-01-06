package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func rotate(x []int64) {
	var temp int64 = x[0]
	for i := 0; i < len(x)-1; i++ {
		x[i] = x[i+1]
	}
	x[len(x)-1] = temp
	fmt.Println(x)
}

func main() {

	fmt.Println("Hello world", "hi")
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("enter n")
	sc.Scan()
	n, _ := strconv.ParseInt(sc.Text(), 10, 64)
	fmt.Println(n)
	var x = []int64{}
	var i int64 = 0
	for i < n {
		fmt.Println("enter x")
		sc.Scan()
		temp, _ := strconv.ParseInt(sc.Text(), 10, 64)
		x = append(x, temp)
		i++
	}
	fmt.Println(x)
	for i := 0; i < 5; i++ {
		rotate(x)
	}

}
