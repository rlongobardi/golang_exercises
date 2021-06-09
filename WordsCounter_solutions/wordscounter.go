package main

import (
	"fmt"
	"strings"
)

func main() {

	mytext := `
	ciao ciao american girl
	my babe babe doll
	jump babe now right now
	till the moon babe ciao
	`

	words := strings.Fields(mytext)
	fmt.Println(words)
	mapwords := map[string]int{}

	for _, w := range words {
		v, okay := mapwords[w]
		if okay {
			mapwords[w] = v + 1
		} else {
			mapwords[w] = 1
		}

	}

	for k, v := range mapwords {

		fmt.Printf("%s -> %d\n", strings.ToLower(k), v)
	}

}
