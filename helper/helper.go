package helper

import "strings"

func IsValidInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsNumber := userTickets <= remainingTickets && userTickets > 0

	return isValidName, isValidEmail, isValidTicketsNumber
}