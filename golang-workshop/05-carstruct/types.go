package carstruct

type Car struct {
	Color        string
	Transmission Transmission
}

type Transmission struct {
	Rotations   int
	EngagedGear int
}