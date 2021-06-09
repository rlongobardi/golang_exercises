package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call the url -%s", err)
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
