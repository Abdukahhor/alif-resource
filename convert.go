package main

import (
	"fmt"
)
//Don’t use floating point numbers for money.
func main() {
	var x float32
	x = 80917.65
	fmt.Println(int(x * 100))

	var y float64
	y = 80917.65
	fmt.Println(int(y * 100))
}
