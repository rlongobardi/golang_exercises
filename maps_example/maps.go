package main

import (
	"fmt"
)

func main() {

	library := map[string]string{
		"123": "THE BIBLE",
		"234": "MAHABHARAT",
		"456": "THE PILLARS OF EARTH",
	}

	library["321"] = "RAMAYANA"

	fmt.Println("Map size is: ", len(library))
	fmt.Printf(library["123"] + "\n")
	fmt.Println(library)

	fmt.Println(library["444"])

	value, okay := library["234"]

	if okay {
		fmt.Println("found book with title", value)
	} else {
		fmt.Println("Value not found!")
	}

	fmt.Println(library["321"])
	delete(library, "321")
	fmt.Println(library)

	for k := range library {
		fmt.Println(k)
	}

	for k, v := range library {
		fmt.Printf("%s->%s\n", k, v)
	}

}
