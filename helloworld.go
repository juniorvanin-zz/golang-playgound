package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result string
	var whitespaceChar = " "

	for index, arg := range os.Args {
		result += strconv.Itoa(index) + whitespaceChar + arg
		result += whitespaceChar
	}

	fmt.Println(result)
}
