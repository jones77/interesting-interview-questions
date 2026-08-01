package main

import (
	"fmt"
	"log"

	"github.com/jones/interesting-interview-questions/internal/generate"
)

func main() {
	specs, err := generate.EmitAll(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("generated %d questions\n", len(specs))
}
