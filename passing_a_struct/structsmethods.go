package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {

	t := Trade{
		Symbol: "MIL",
		Volume: 10,
		Buy:    true,
		Price:  54.33,
	}
	fmt.Printf("t=%v\n", t)
	fmt.Printf("%0.4f\n", t.Value())
}
