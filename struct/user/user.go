package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u User) OutputUser() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email string, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "admin",
			lastName:  "ss",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(firstName string, lastName string, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" {
		return nil, errors.New("Firstname or lastname cannot be empty")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil

}
