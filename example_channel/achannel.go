package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func returnType(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	log.Println(ct)
	return ct
}

func main() {

	//simple example of a channel int usage
	ch := make(chan int)

	go func() {
		ch <- 42
	}()
	v := <-ch

	fmt.Printf("value of v = %d\n", v)

	chString := make(chan string)

	fmt.Println("*********************************************************")
	//example of channel usage with wanted time delay
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://httpbin.org/xml",
		"https://httpbin.org/json",
	}
	go func() {
		for _, url := range urls {
			ct := returnType(url)
			fmt.Printf("sending content-type of the request to the channel: %s\n", ct)
			chString <- ct
			time.Sleep(3 * time.Second) //delay of three seconds
		}
	}()

	for _, url := range urls {
		ctype := <-chString
		fmt.Printf("url: %s has content-type: %s\n", url, ctype)
	}

}
