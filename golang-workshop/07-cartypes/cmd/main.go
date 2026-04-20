package main

import (
	"fmt"

	"golang.source-fellows.com/training/cartypes"
)

func checkType(obj any) {
	switch obj.(type) {
	case cartypes.Audi:
		fmt.Println("It's an Audi")
	case cartypes.Porsche:
		fmt.Println("It's a Porsche")
	}
}
func main() {
	audi := cartypes.Audi{
		LicensePlate: "ABC123",
		Transmission: cartypes.Transmission{
			Rotations:   1000,
			EngagedGear: 3,
		},
	}

	porsche := cartypes.Porsche{
		LicensePlate: "XYZ789",
		Transmission: cartypes.Transmission{
			Rotations:   2000,
			EngagedGear: 4,
		},
	}

	checkType(audi)
	checkType(porsche)
}
