package main

import "fmt"

func main() {
	var age = 32
	var agePointer *int = &age
	fmt.Println("Age", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult years", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
