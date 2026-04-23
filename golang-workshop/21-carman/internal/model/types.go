package model

type LicensePlate string

type Transmission struct {
	Rotations   int
	EngagedGear int
}

type Car interface {
	StartEngine() error
	Accelerate()
	Brake()
}

type Audi struct {
	Plate         LicensePlate
	Transmission  *Transmission
	SpeedInKmH    int
	engineStarted bool
}

func (a *Audi) StartEngine() error {
	return nil
}

func (a *Audi) Accelerate() {}

func (a *Audi) Brake() {}
