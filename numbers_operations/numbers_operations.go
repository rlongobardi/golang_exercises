package main

import (
	"fmt"
)

func main() {

	var x int = 3
	var y int = 4
	sum := x + y
	fmt.Printf("x=%v plus y=%v is equal to %v ", x, y, sum)
	fmt.Printf("the type of x=%v is %T\n", x, x)
	fmt.Printf("thee type of y=%v is %T\n", y, y)
	p := x * y
	fmt.Printf("The product between x=%v & y=%v is  p=%v\n", x, y, p)
	var q float64 = 0.0
	q = (float64)((x * y) / 2)
	fmt.Printf("the value of x*y=%v divided by 2 is q=%v\n", (x * y), q)
	fmt.Printf("the type of q is %T and its value is %v\n", q, q)
	var qf float32
	fmt.Printf("the value of 'mean' is %v and type of qf is %T\n", qf, qf)

}
