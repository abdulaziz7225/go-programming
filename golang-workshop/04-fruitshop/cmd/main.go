package main

import (
	"fmt"
	"math"
)

const applesCount int = 5
const orangesCount int = 3
const applePriceCents int = 75  // 0.75 Euro
const orangePriceCents int = 90 // 0.90 Euro

var totalCents int = 0

func main() {
	fmt.Println("Price calculation FRUITSHOP")

	fmt.Println()
	calculateApplePrices()

	fmt.Println()
	calculateOrangePrices()

	fmt.Println()
	totalEuros := float64(totalCents) / 100.0
	fmt.Printf("Total price for all fruits is: %.2f Euros\n", totalEuros)
}

func calculateApplePrices() {
	fmt.Println("Calculating the price of apples")
	for i := 0; i < applesCount; i++ {
		totalCents += applePriceCents
		// Use if statement for a special apple discount
		if i%3 == 2 {
			totalCents -= 10 // 10 Cents
			fmt.Printf("Apple number %d costs %d cents (discount applied)\n", i, applePriceCents-10)
		} else {
			fmt.Printf("Apple number %d costs %d cents\n", i, applePriceCents)
		}
	}
}

func calculateOrangePrices() {
	for i := 1; i <= orangesCount; i++ {
		switch {
		case i%4 == 0:
			price := int(math.Trunc(float64(orangePriceCents) * 0.85))
			totalCents += price
			fmt.Printf("Orange number %d costs %d cents (discount applied)\n", i, price)
		default:
			totalCents += orangePriceCents
			fmt.Printf("Orange number %d costs %d cents\n", i, orangePriceCents)
		}
	}
}
