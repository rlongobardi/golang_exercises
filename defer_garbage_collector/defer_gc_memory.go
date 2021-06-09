package main

import (
	"fmt"
)

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func increase(a int) {
	fmt.Println("Value a=", a)
	a++
	fmt.Println("Value a=", a)
}

func worker() {
	defer cleanup("A")
	defer cleanup("B")
	fmt.Println("worker")

}

func worker2() {
	var a int = 23
	defer increase(a)
	fmt.Println("Worker 2")
}

func main() {

	worker()
	worker2()

}
