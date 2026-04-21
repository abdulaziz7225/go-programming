package errorhandling

import (
	"errors"
	"fmt"
)

type Audi struct {
	LicensePlate  LicensePlate
	Transmission  Transmission
	SpeedInKmH    int
	engineStarted bool
}

func (a *Audi) StartEngine() error {
	if a.engineStarted {
		return errors.New("could not start engine, because it is already running")
	}

	fmt.Println("Audi starts engine")
	a.engineStarted = true
	return nil
}

func (a *Audi) Accelerate() {
	a.SpeedInKmH += 10
}

func (a *Audi) Brake() {}
