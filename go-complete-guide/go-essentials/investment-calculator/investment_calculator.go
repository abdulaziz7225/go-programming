package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation_rate float64 = 2.5
	var investment_amount float64
	var years float64
	var expected_return_rate float64

	fmt.Print("Investment Amount : ")
	fmt.Scan(&investment_amount)
	
	fmt.Print("Years : ")
	fmt.Scan(&years)
	
	fmt.Print("Expected Return Rate : ")
	fmt.Scan(&expected_return_rate)

	future_value := investment_amount * math.Pow(1+expected_return_rate/100, years)
	future_real_value := future_value / math.Pow(1+inflation_rate/100, years)

	formatted_fv := fmt.Sprintf("Future Value : %.2f\n", future_value)
	formatted_frv := fmt.Sprintf("Future Real Value : %.2f\n", future_real_value)
	fmt.Print(formatted_fv, formatted_frv)
}
