package main

import (
	"fmt"
)

func main() {
	slices := []int{14, 49, 78, 99, 1, 32, 89}
	fmt.Printf("slices is of size %v and type of elements is %T \n", len(slices), slices)
	fmt.Println(slices)
	fmt.Println("let's find the max..")
	max := slices[0]
	for i := 1; i < len(slices); i++ {
		if max < slices[i] {
			max = slices[i]
		}
	}
	fmt.Println("Max element in the array/slices is ", max)
}
