package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch1, ch2 := make(chan int), make(chan string)

	go func() {
		flag := rand.Intn(6)
		fmt.Println("flag =", flag)
		if flag < 3 {
			ch2 <- "Test"
		} else {
			ch1 <- 42
		}
	}()

	select {
	case val := <-ch1:
		fmt.Printf("Got value value %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("Got value = %s from ch2\n", val)
	}

	fmt.Println("_____________*____*________________")
	out := make(chan float64)
	go func() {
		time.Sleep(time.Second)
		out <- 3.15
	}()
	select {
	case val := <-out:
		fmt.Printf("got value %f from 'out' channel \n", val)
	case <-time.After(2000 * time.Millisecond):
		fmt.Println("timeout")
	}

}
