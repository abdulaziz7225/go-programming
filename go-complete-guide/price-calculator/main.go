package main

import (
	"fmt"

	// "example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.08, 0.12, 0.17}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmd := cmdmanager.New()
		var priceJob = prices.New(taxRate, fm)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process the job")
			fmt.Println(err)
		}
	}
}
