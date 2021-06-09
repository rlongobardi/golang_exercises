package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value, it's not possible -> (%f)", n)
	}
	return math.Sqrt(n), nil
}

func errorHandler(ex error, value float64) {

	if ex != nil {
		fmt.Printf("ERROR: %s\n", ex)
	} else {
		fmt.Println(value)
	}

}

func main() {

	s1, err := sqrt(2.0)
	errorHandler(err, s1)
	s2, err := sqrt(-2.0)
	errorHandler(err, s2)

}
