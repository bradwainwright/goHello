package main

import (
	"fmt"
	"math"
)

func newton(z, x float64) float64 {
	return (z - ((z*z)-x)/(2*z))
}
func Sqrt(x float64) float64 {
	z := 1.0
	p := 0.0
	for math.Abs(p-z) > 0.000001 {
		z = newton(z, x)
		p = newton(z, x)
	}
	return z

}
func main() {
	fmt.Println("HELLO WORLD")
	fmt.Println("test")
	fmt.Println(math.Abs(3.142))
	a := "test string"
	fmt.Println(a)
	fmt.Println(Sqrt(2))
}
