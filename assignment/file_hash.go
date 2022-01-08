package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func hash_data(file_name string) {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	var s string = ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		hsha2 := sha256.Sum256([]byte(line))
		fmt.Printf("%x \n", hsha2)
		s += fmt.Sprintf("%x\n", hsha2)

	}
	file1, err2 := os.Create(file_name + "_hash.txt")
	if err2 != nil {
		panic(err2)
	}
	_, err3 := io.WriteString(file1, s)
	if err3 != nil {
		panic(err3)
	}

}

func main() {
	hash_data("test.txt")
}
