package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `{
	"user": "Mr Buffy",
	"type": "deposit",
	"amount": 100000.3
}`

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	rdr := bytes.NewBufferString(data)

	dec := json.NewDecoder(rdr)

	req := &Request{}

	if err := dec.Decode(req); err != nil {
		log.Fatalf("Error: can't decode -%s", err)
	}

	fmt.Printf("Json requeest is -> %+v\n", req)

	prevBalance := 850000.0 // got from a datasource

	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}

	//fmt.Printf("response ->%+v\n", resp)

}
