package main

import "golang.source-fellows.com/training/carinterface"

func main() {
	transmission := carinterface.Transmission{EngagedGear: 1}

	var car carinterface.Car

	car = carinterface.Audi{Transmission: transmission}
	car.StartEngine()
	car.Accelerate()

	car = carinterface.Porsche{}
	car.StartEngine()
	car.Accelerate()
}
