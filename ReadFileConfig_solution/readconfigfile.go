package main

import (
	"fmt"
	toml "github.com/pelletier/go-toml"
	"log"
	"os"
)

type ApplicationConfig struct {
	Credentials struct {
		Username string
		Password string
	}
	ServerDetails struct {
		Url  string
		Port int
	}
}

func main() {

	file, err := os.Open("appconf.toml")
	if err != nil {
		log.Fatalf("error: can't open this file - %s", err)
	}
	defer file.Close()

	cfg := &ApplicationConfig{}

	dec := toml.NewDecoder(file)
	if err := dec.Decode(cfg); err != nil {
		log.Fatal("error: can't decode configuration file - %s", err)
	}

	fmt.Printf(" Application Configuration -> %+v\n", cfg)
	fmt.Printf("Credentials are username=%s and password=%s", cfg.Credentials.Username, cfg.Credentials.Password)

}
