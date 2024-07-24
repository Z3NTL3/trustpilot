package main

import (
	"fmt"
	"log"

	"github.com/z3ntl3/trustpilot/client"
)

var URL = "https://trustpilot.com/review/www.ikea.com"

func main() {
	client := client.New()

	result, err := client.Execute(URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Resulting date model: %+v \n", result)
}