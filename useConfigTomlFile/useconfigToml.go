package main

import (
	"fmt"
	toml "github.com/pelletier/go-toml"
	"log"
	"os"
)

type Config struct {
	Login struct {
		User     string
		Password string
	}
}

func main() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("error: can't open config file - %s", err)
	}

	defer file.Close()

	cfg := &Config{}
	dec := toml.NewDecoder(file)
	if err := dec.Decode(cfg); err != nil {
		log.Fatalf("error: can't decode configuration file - %s", err)
	}
	fmt.Printf("%+v\n", cfg)
}
