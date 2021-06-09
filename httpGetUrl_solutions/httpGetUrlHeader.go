package main

import (
	"fmt"
	"io"
	"net/http"
)

func contentType(url string) (string, error) {
	fmt.Printf("request for url=%s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer closeRespBody(resp)

	ctype := resp.Header.Get("Content-Type")

	if ctype == "" {
		return "", fmt.Errorf("can't find Content-Type Header")
	}
	return ctype, nil
}

func closeRespBody(resp *http.Response) {
	func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic("error closing file")
		}
	}(resp.Body)
}

func respHandler(ex error, value string) {

	if ex != nil {
		fmt.Printf("ERROR: %s\n", ex)
	} else {
		fmt.Println(value)
	}

}

func main() {
	url := "https://golang.com"
	resp, err := contentType(url)
	respHandler(err, resp)
	url2 := "https://blablah.sd"
	resp2, err2 := contentType(url2)
	respHandler(err2, resp2)

}
