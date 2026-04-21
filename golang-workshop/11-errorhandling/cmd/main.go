package main

import (
	// "fmt"
	"log"

	"golang.source-fellows.com/training/errorhandling"
)

func main() {
	a := &errorhandling.Audi{}
	err := a.StartEngine()
	if err != nil {
		log.Fatal(err)
	}
	err = a.StartEngine()
	if err != nil {
		log.Fatal(err)
	}
}
