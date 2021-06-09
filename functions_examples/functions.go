package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divideMod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {

	s := add(3, 2)
	fmt.Println("the sum is", s)
	q, r := divideMod(5, 2)
	fmt.Printf("The div is %d and the mod is %d\n", q, r)

}
