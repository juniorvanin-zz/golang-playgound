package main

import "fmt"

func main() {
	number := *f()

	fmt.Println(number)

	increment(&number)

	fmt.Println(number)

}

func increment(count *int) {
	*count++
}

func f() *int {
	number := 10

	return &number
}
