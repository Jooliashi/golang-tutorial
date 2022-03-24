package main

import "fmt"

func main() {
	var number uint16 = 260
	number = number + 5
	fmt.Println("Hello World!", number)
}

//Data Types:
	//Unsigned Integers (no negatives)
		//uint8/byte (0-255)
		//uint16 (<= 65535)
		//uint32
		//uint64
	//Signed Integers
		//int8(-128 to 127)
		//int16(-32769 to 32767)
		//int32
		//int64
	//Machine Dependent Types
		//uint
		//int
		//unitptr
	//Float
		//float32
		//float64
	//Complex
		//complex64
		//complex128
	//Strings
	//Boleans
		//true
		//false
