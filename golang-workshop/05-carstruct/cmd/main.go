package main

import "fmt"

import "golang.source-fellows.com/training/carstruct"

func main() {
	// a := Car {}
	// fmt.Printf("%v\n", a)

	// var c Car
	// fmt.Printf("%+v\n", c)
	// c.color = "blue"
	// fmt.Printf("Car color: %s\n", c.color)

	t := carstruct.Transmission{}
	t.Rotations = 200
	t.EngagedGear = 1
	c := carstruct.Car{Transmission: t, Color: "blue"}
	fmt.Printf("%+v\n", c)
}
