package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	fmt.Printf("url: %s has content-type: %s\n", url, ctype)
}

func main() {
	fmt.Println("getting url requests")
	urls := []string{"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
		"https://httpbin.org/json",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()

}
