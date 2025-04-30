package main

import (
	"bank/fileops"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	balance := fileops.ReadFloatFromFile(accountBalanceFile)

	fmt.Println("Welcome to the bank")

	for {
		presentOptions()

		choice := getUserInput("Your choice")

		switch choice {
		case 1:
			fmt.Printf("Your balance: %.2f", balance)

		case 2:
			amount := getUserInput("Amount")

			if amount <= 0 {
				fmt.Println("Balance could not be added")
			} else {
				balance += amount
				fmt.Println("Balance has been added")
			}

		case 3:
			amount := getUserInput("Amount")

			if amount <= 0 || balance-amount < 0 {
				fmt.Println("Balance could not be deducted")
			} else {
				balance -= amount
				fmt.Println("Balance has been deducted")
			}

		case 4:
			fmt.Println(fileops.ReadFloatFromFile(accountBalanceFile))

		default:
			fmt.Println("\nThank you for choosing us!")
			return
		}

		fileops.WriteToFile(accountBalanceFile, balance)
	}
}

func getUserInput(infoText string) float64 {
	var input float64
	fmt.Printf("%v: ", infoText)
	fmt.Scan(&input)
	return input
}
