package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //we ensure will be ended

	fmt.Printf("Worker %d has started..\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d has completed!\n", id)

}

func main() {
	//execution of a task in sequential order
	var wg sync.WaitGroup
	for i := 1; i < 6; i++ {
		wg.Add(1)
		go worker(i, &wg)
		wg.Wait()
	}

	fmt.Println("***************")
	//unorder execution of tasks
	var wg2 sync.WaitGroup
	for i := 6; i < 11; i++ {
		wg2.Add(1)
		worker(i, &wg2)
	}
	wg2.Wait()

}
