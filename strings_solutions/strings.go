package main

import (
	"fmt"
)

func main() {

	var s string = "Hello Golanger"

	fmt.Println(s)
	fmt.Printf("first character of this 's' is %v and last char is %v and type is %T\n", s[0], s[len(s)-1], s)
	fmt.Println(s[0:5])
	fmt.Println(s[:len(s)])
	fmt.Println(s[0:])
	fmt.Println("string size is %v", len(s))

	n := 72
	st := fmt.Sprintf("%d", n)
	fmt.Printf("st = %v and type is %T \n", st, st)
	fmt.Printf("st = %q (type is %T)", st, st)
}
