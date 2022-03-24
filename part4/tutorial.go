package main

import "fmt"

func main() {
	fmt.Printf("Hello %T %v", 10, 10)
	var x string = fmt.Sprintf("Hello %T %v", 10, 10)
	fmt.Println(x)

	fmt.Printf("Number: %b", 1025)

	fmt.Printf("Number: %e", 2.932094820)

	fmt.Printf("Number: %q", "tim")

	fmt.Printf("Number: %9q", "tim")

	fmt.Printf("Number: %.2f is cool", 3.244324)

	fmt.Printf("Number: %07d is cool", 45)
}

//Boolean
	//%t
//%
	//%%
//Integer
	//%b (base 2)
	//%d (base 10)
//Floating Points
	//%e (scientific notation)
	//%f or %F (decimals no exponents)
	//%g (large exponents)
//String
	//%s 
	//%q (with quotation marks)
//Width & Precision
	//%f (default width and precision)
	//%9f (width 9)
	//%.2f (default width, precision 2)
	//%9.2f (9 width, precision 2)
//Padding
	//%09d (pads digit to length 9 with preceeding 0's)