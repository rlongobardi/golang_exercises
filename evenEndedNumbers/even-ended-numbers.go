package main

import (
	"fmt"
)

func main() {

	var count int = 0

	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			m := a * b

			res := fmt.Sprintf("%d", m)
			if res[0] == res[len(res)-1] {
				count++
			}

			if b == 1000 {
				fmt.Println("Japan")
			}
		}
	}

	fmt.Println("total even-ended-numbers is:", count)

}
