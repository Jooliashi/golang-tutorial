package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple": 5,
		"pear": 6,
		"orange": 9,
	}

	mp2 := make(map[string]int)
	fmt.Println(mp2)
	fmt.Println(mp["apple"])
	mp["tim"] = 900
	fmt.Println(mp)
	
	delete(mp, "apple")
	fmt.Println(mp)

	val, ok := mp["apple"]
	fmt.Println(val, ok)
}
