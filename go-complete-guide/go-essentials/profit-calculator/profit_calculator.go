package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := get_user_info("Revenue")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := get_user_info("Expenses")
	if err != nil {
		fmt.Println(err)
		return
	}

	tax_rate, err := get_user_info("Tax Rate")
	if err != nil {
		fmt.Println(err)
		return
	}

	earning_before_tax, earning_after_tax, ratio := calculate_financial(revenue, expenses, tax_rate)

	fmt.Printf("Earning before tax : %.1f\n", earning_before_tax)
	fmt.Printf("Earning after tax : %.1f\n", earning_after_tax)
	fmt.Printf("Ratio : %.3f\n", ratio)
	store_results(earning_before_tax, earning_after_tax, ratio)
}

func get_user_info(info_text string) (float64, error) {
	var user_input float64
	fmt.Print(info_text, " : ")
	fmt.Scan(&user_input)

	if user_input <= 0 {
		return 0, errors.New("value should be greater than 0")
	}

	return user_input, nil
}

func store_results(before_tax, after_tax, ratio float64) {
	results := fmt.Sprintf("Earning before tax : %.1f\nEarning after tax : %.1f\nRatio : %.3f\n", before_tax, after_tax, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculate_financial(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	var earning_before_tax float64 = revenue - expenses
	var earning_after_tax float64 = earning_before_tax * (1 - tax_rate/100)
	var ratio float64 = earning_before_tax / earning_after_tax

	return earning_before_tax, earning_after_tax, ratio
}
