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

func main() {

	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Printf("t1=%v\n", t1)
	fmt.Printf("%+v\n", t1)
	fmt.Printf("Price for %s is %0.2f\n", t1.Symbol, t1.Price)
	t2 := Trade{
		Symbol: "NASDAQ",
		Volume: 20,
		Price:  123.45,
		Buy:    true,
	}

	fmt.Printf("t2=%v\n", t2)
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v", t3)

}
