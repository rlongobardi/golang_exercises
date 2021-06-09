package main

import (
	"fmt"
)

func doubleVal(a []int, index int) {
	a[index] *= 2
}

func doubleP(a *int) {
	*a *= 2
}

func main() {
	myarray := []int{3, 4, 5, 7, 2}
	doubleVal(myarray, 2)
	fmt.Println(myarray)

	val := 9
	doubleP(&val)
	fmt.Println(val)
}
