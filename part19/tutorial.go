package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string){
	str = "not changed!"
}

func main() {
	x := 7
	fmt.Println(&x) //computer memory where x sits which is the pointer
	y := &x
	fmt.Println(x, y)
	*y = 8 //dereference the pointer and set it to a different value
	fmt.Println(x, y)

	toChange := "hello"
	changeValue(&toChange)
	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
}
