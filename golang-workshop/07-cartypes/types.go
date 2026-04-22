package cartypes

//go:generate goplantuml -aggregate-private-members -show-aggregations -show-compositions -output cartypes.puml .

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
