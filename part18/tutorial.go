package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 1000
}

func main() {
	//immutable data types
	var x int = 5
	y := x
	y = 7
	fmt.Println(x, y)

	//mutable data types: slice and maps
	var z []int = []int{3, 4, 5}
	k := z
	k[0] = 100
	fmt.Println(z, k)

	var a map[string]int = map[string]int{"hello": 3}
	b := a
	b["y"] = 100
	fmt.Println(a, b)

	var c [2]int = [2]int{2, 3}
	d := c
	d[0] = 100
	fmt.Println(c, d)

	var s []int = []int{3, 4, 5}
	fmt.Println(s)
	changeFirst(s)
	fmt.Println(s)
}
