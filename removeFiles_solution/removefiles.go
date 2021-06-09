package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := "data.txt"
	err := os.Remove("data")
	if err != nil {
		log.Println("was it removed", err)

	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("File was not found:", err)
	}

	log.Println(string(file))
}
