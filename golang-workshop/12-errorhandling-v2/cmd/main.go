package main

import (
	// "fmt"
	"errors"
	"log"

	"golang.source-fellows.com/training/errorhandling/v2"
)

func main() {
	a := &errorhandling.Audi{}
	err := a.StartEngine()
	if err != nil {
		if errors.Is(err, errorhandling.ErrF42625) {
			log.Println("Specific error: Transmission initialization failed")
		}
		log.Fatal(err)
	}
	err = a.StartEngine()
	if err != nil {
		log.Fatal(err)
	}
}
