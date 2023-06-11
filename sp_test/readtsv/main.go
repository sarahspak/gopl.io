package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fname := os.Args[1]
	file, err := os.Open(fname)
	if err == nil {
		input := bufio.NewScanner(file)
		for input.Scan() {
			fmt.Println(input.Text())
		}
	}
}
