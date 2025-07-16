package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ACCOUNT_BALANCE_FILE = "balance.txt"

func read_balance_from_file() (float64, error) {
	data, err := os.ReadFile(ACCOUNT_BALANCE_FILE)

	if err != nil {
		return 1000, errors.New("failed to find the file")
	}

	balance_text := string(data)
	balance, err := strconv.ParseFloat(balance_text, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func write_balance_to_file(balance float64) {
	var balance_text string = fmt.Sprint(balance)
	os.WriteFile(ACCOUNT_BALANCE_FILE, []byte(balance_text), 0644)
}

func main() {
	var account_balance, err = read_balance_from_file()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err, "\n-----------------------------")
	}

	fmt.Println("Welcome to GO Bank!")
	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice : ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance :", account_balance)
		} else if choice == 2 {
			var deposit_amount float64
			fmt.Print("Your deposit : ")
			fmt.Scan(&deposit_amount)

			if deposit_amount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			account_balance += deposit_amount
			fmt.Println("Balance updated! New amount :", account_balance)
			write_balance_to_file(account_balance)
		} else if choice == 3 {
			var withdrawal_amount float64
			fmt.Print("Withdrawal amount : ")
			fmt.Scan(&withdrawal_amount)

			if withdrawal_amount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawal_amount > account_balance {
				fmt.Println("Withdraw failed! You account doesn't have such amount of money")
				continue
			}
			account_balance -= withdrawal_amount
			fmt.Println("Money withdrew! New balance :", account_balance)
			write_balance_to_file(account_balance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thank you for choosing our bank")
}
