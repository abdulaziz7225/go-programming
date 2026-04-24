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
	Plate         LicensePlate  `json:"plate"`
	Transmission  *Transmission `json:"transmission"`
	SpeedInKmH    int           `json:"-"`
	EngineStarted bool          `json:"engined_started"`
}

func (a *Audi) StartEngine() error {
	return nil
}

func (a *Audi) Accelerate() {}

func (a *Audi) Brake() {}
