package main

import (
	"fmt"
)
func main() {
	x := 5
	y := 6.5
	val := float64(x) + 1.5 == y
	fmt.Printf("%t\n", val)

	str1 := "tim"
	str2 := "Tim"
	val2 := str1 < str2
	fmt.Printf("%t\n", val2)

	cha1 := 'c'
	cha2 := 'c'
	val3 := cha1 == cha2
	fmt.Printf("%t\n", val3)
}