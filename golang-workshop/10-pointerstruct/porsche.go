package pointerstruct

type Porsche struct {
	LicensePlate LicensePlate
	Transmission Transmission
}

func (p *Porsche) StartEngine() error {
	return nil
}

func (p *Porsche) Accelerate() {}

func (p *Porsche) Brake() {}
