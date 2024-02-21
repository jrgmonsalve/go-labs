package models

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type customString string

func (s customString) Upper() string {
	return strings.ToUpper(string(s))
}

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) OutPutUserDetails() {
	var firstName customString = customString(u.firstName)
	fmt.Println(firstName.Upper(), u.lastName, u.birthDate)
}

func (u *User) ClearUser() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
	u.createdAt = time.Time{}
}

// construct
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("All fields are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
