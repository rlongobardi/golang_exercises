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
	//fmt.Println(words)
	mapwords := map[string]int{}

	for _, w := range words {
		mapwords[strings.ToLower(w)]++

	}

	fmt.Println(mapwords)

}
