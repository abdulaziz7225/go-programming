package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/brianvoe/gofakeit/v7"
)

const ACCOUNT_BALANCE_FILE = "balance.txt"
const DEFAULT_ACCOUNT_BALANCE = 1000

func main() {
	var accountBalance, err = fileops.ReadFloatFromFile(ACCOUNT_BALANCE_FILE, DEFAULT_ACCOUNT_BALANCE)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err, "\n-----------------------------")
	}

	fmt.Println("Welcome to GO Bank!")
	fmt.Println("Reach us 24/7 at :", gofakeit.PhoneFormatted())
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice : ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance :", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Your deposit : ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount :", accountBalance)
			fileops.WriteFloatToFile(ACCOUNT_BALANCE_FILE, accountBalance)
		} else if choice == 3 {
			var withdrawalAmount float64
			fmt.Print("Withdrawal amount : ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Withdraw failed! You account doesn't have such amount of money")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Money withdrew! New balance :", accountBalance)
			fileops.WriteFloatToFile(ACCOUNT_BALANCE_FILE, accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thank you for choosing our bank")
}
