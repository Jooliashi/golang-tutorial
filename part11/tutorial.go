package main

import "fmt"

func main() {
	ans := -1

	switch ans{
	case 1, -1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not a case")
	}

	switch {
	case ans > 10:
		fmt.Println("greater than zero")
	case ans < 0:
		fmt.Println("less than 0")
	default:
		fmt.Println("zero")
	}
}