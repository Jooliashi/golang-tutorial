package main

import "fmt"

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	test := func (x int) int {
		return x * -1
	}
	test3 := func(x int) int{
		return x * 7
	}

	test2(test)
	test2(test3)

	func() {
		fmt.Println("test")
	}()

	returnFunc("hello")()
}
