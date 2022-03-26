package main

import "fmt"

func add(x, y int) (z1, z2 int){
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	fmt.Println("before return")
	return 
}

func main() {
	ans1, ans2 := add(6, 7)
	fmt.Println(ans1, ans2)
}
