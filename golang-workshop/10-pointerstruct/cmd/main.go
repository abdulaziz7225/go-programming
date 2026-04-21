package main

import (
	"fmt"

	"golang.source-fellows.com/training/pointerstruct"
)

func main() {
	a := &pointerstruct.Audi{}
	// a.SpeedInKmH = 10
	(*a).SpeedInKmH = 25 // same as above
	fmt.Printf("Audi: %+v\n", a)
	a.Accelerate()
	fmt.Printf("Audi: %+v\n", a)
}
