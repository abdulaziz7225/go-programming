package errorhandling

type LicensePlate string

type Car interface {
	StartEngine() error
	Accelerate()
	Brake()
}

type Transmission struct {
	Rotations   int
	EngagedGear int
}
