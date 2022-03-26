package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func changeX(pt *Point){
	pt.x = 100
}

type Circle struct {
	radius float64
	*Point
}

func main() {
	var p1 Point = Point{1,  2}
	var p2 Point = Point{-5, 7}
	fmt.Println(p1.x, p2.x)

	p1.x = 7
	fmt.Println(p1.x, p2.x)

	p3 := &Point{y: 3}
	fmt.Println(p3)
	changeX(p3)
	fmt.Println(p3)
	p3.x = 8
	fmt.Println(p3)

	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1.x)
}
