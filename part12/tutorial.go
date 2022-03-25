package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 100
	arr[4] = 900
	fmt.Println(arr[0])

	arr2 := [3]int{4, 6, 7}
	fmt.Println(len(arr2))

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)

	arr2D := [2][2]int{{1,2},{3,4}}
	fmt.Println(arr2D)
}