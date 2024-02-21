package main

import (
	"fmt"

	"github.com/jrgmonsalve/go-labs/udemy/5s_structs/models"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	var appUser1 *models.User

	appUser1, err := models.NewUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser1.OutPutUserDetails()
}

func getUserData(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}
