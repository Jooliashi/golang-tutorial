package main

import (
	"fmt"
	//"math"
)
func main() {
	var num1 float64 = 8
	var num2 int = 4
	answer := num1 / float64(num2)
	fmt.Printf("%f\n", answer)

	var num3 int = 9
	var num4 int = 4
	answer2 := num3 / num4
	fmt.Printf("%d\n", answer2)

	var num5 int = 9
	var num6 int = 4
	answer3 := num5 % num6
	fmt.Printf("%d\n", answer3)
}