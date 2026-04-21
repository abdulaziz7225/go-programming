package errorhandling

import (
	"errors"
	"fmt"
)

type Audi struct {
	LicensePlate  LicensePlate
	Transmission  *Transmission
	SpeedInKmH    int
	engineStarted bool
}

func (a *Audi) StartEngine() error {
	fmt.Println("Audi starts engine")

	if a.engineStarted {
		return errors.New("could not start engine, because it is already running")
	}

	if  a.Transmission == nil {
		return errors.New("no transmission")
	}
	err := a.Transmission.initialize()
	if err != nil {
		return fmt.Errorf("error while init transmission: %v", err)
	}

	a.engineStarted = true
	return nil
}

func (a *Audi) Accelerate() {
	a.SpeedInKmH += 10
}

func (a *Audi) Brake() {}
