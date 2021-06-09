package main

import (
	"fmt"
	//"os"
)

func main() {
	//PANIC EXAMPLE 1:
	/*vals := []int{3, 4, 5}
	//this would cause panic
	v := vals[10]
	fmt.Println(v)*/
	//-----------------------
	//PANIC EXAMPLE 2:
	/*file, err := os.Open("no-such-file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("file opened")*/
	//PANIC EXAMPLE 3:
	v := safeValue([]int{1, 2, 3}, 10)
	fmt.Println("value is ->", v)

}

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()
	return vals[index]
}
