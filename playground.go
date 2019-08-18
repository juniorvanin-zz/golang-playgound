package main

import "fmt"

func main() {
	number:= new(int)
	var number2 *int
	number =  number2
	number = new(int)
	*number = 15
	fmt.Println(number)
	fmt.Println(number2)

	fmt.Println(gcd(15,4))
	fmt.Println(fib(15))

}

func gcd( x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i:= 0 ;i < n ;i++  {
		x, y = y, x+y
	}
	return x
}