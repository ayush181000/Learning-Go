package main

import (
	"fmt"
	"struct/user"
)

func main() {
	userFirstName := getUserInput("Please enter your first name")
	userLastName := getUserInput("Please enter your last name")
	userBirthdate := getUserInput("Please enter your birth date")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		return
	}

	admin := user.NewAdmin("admin@google.com", "test123")

	admin.OutputUser()
	admin.ClearUserName()
	admin.OutputUser()

	appUser.OutputUser()
	appUser.ClearUserName()
	appUser.OutputUser()
}

func getUserInput(infoText string) string {
	var input string
	fmt.Printf("%v: ", infoText)
	fmt.Scanln(&input)
	return input
}
