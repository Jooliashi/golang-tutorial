package main

import "fmt"

func main() {
	var number = 2000.98 //implicit
	fmt.Printf("%T", number)

	number2 := 6
	fmt.Printf("%T", number2)

	number3 := true
	fmt.Printf("%T", number3)

	var number4 uint64
	fmt.Println(number4)
	//all types have default values

	var bl bool
	fmt.Println(bl)
}
