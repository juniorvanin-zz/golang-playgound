package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5+32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32)*5/9)
}
func main() {
	//fmt.Println(CToF(15))
	//fmt.Println(FToC(60))

	// assertions

	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	// compilation error
	//fmt.Printf("%g\n", BoilingC-boilingF)

	var c Celsius
	f := Fahrenheit(15)
	newF := new(Fahrenheit)
	*newF = 19
	fmt.Println(c == 0)
	fmt.Println(f >= 0.0)
	fmt.Println(f == *newF)
	fmt.Println(c == Celsius(f))
	fmt.Println(f.String())
	fmt.Println(newF)

}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("Temperature %.2f F", f)
}