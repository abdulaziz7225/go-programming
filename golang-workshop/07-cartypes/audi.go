package cartypes

import "fmt"

type Audi struct {
	LicensePlate LicensePlate
	Transmission Transmission
}

func (a Audi) StartEngine() error {
	fmt.Println("Audi starts engine")
	return nil
}

func (a Audi) Accelerate() {
	fmt.Printf("Audi accelerates in gear: %d\n", a.Transmission.EngagedGear)
}

func (a Audi) Brake() {
	fmt.Println("Audi brakes")
}
