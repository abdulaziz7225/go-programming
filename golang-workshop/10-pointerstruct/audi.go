package pointerstruct

type Audi struct {
	LicensePlate LicensePlate
	Transmission Transmission
	SpeedInKmH   int
}

func (a *Audi) StartEngine() error {
	return nil
}

func (a *Audi) Accelerate() {
	a.SpeedInKmH += 10
}

func (a *Audi) Brake() {}
