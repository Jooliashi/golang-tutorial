package main

import (
	"fmt"
)
func main() {
	//if statement
	age := 17

	if age >= 18 {
		fmt.Println("You can ride alone!")
	} else if age >= 14 {
		fmt.Println("You can ride with a parent!")
	} else {
		fmt.Println("You cannot ride!")
		fmt.Printf("wait %d years", 18 - age)
	}
}