package cartypes

import "fmt"

type Porsche struct {
	LicensePlate LicensePlate
	Transmission Transmission
}

func (p Porsche) StartEngine() error {
	fmt.Println("Prosche starts engine")
	return nil
}

func (p Porsche) Accelerate() {
	fmt.Printf("Porsche accelerates in gear: %d\n", p.Transmission.EngagedGear)
}

func (p Porsche) Brake() {
	fmt.Println("Porsche brakes")
}
