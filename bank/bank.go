package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	balance := readBalanceFromFile()

	fmt.Println("Welcome to the bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Print passbook")
		fmt.Println("5. Exit")

		choice := getUserInput("Your choice")

		switch choice {
		case 1:
			fmt.Printf("Your balance: %.2f", balance)
			continue

		case 2:
			amount := getUserInput("Amount")

			if amount <= 0 {
				fmt.Println("Balance could not be added")
			} else {
				balance += amount
				fmt.Println("Balance has been added")
			}

			continue

		case 3:
			amount := getUserInput("Amount")

			if amount <= 0 || balance-amount < 0 {
				fmt.Println("Balance could not be deducted")
			} else {
				balance -= amount
				fmt.Println("Balance has been deducted")
			}

			continue

		case 4:
			writeBalanceToFile(balance)
			continue

		default:
			fmt.Print("Goodbye")
			return
		}

	}
	fmt.Println()
	fmt.Println("Thank you for choosing us!")
}

func getUserInput(infoText string) float64 {
	var input float64
	fmt.Printf("%v: ", infoText)
	fmt.Scan(&input)
	return input
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}
