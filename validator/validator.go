package helper

import (
	"fmt"
	"strings"
)

// Capetalize the variable/function name meaning that the variable/function has been exported and its a global

func Input[T any](feild_ptr *T, msg string) {
	fmt.Println(msg)
	// fmt.Scan( <- address of the variable -> )
	fmt.Scan(feild_ptr)
}

func HandleInValidUserInput(firstName *string, lastName *string, email *string, userTickets *uint) {
	isValisName := len(*firstName) > 0 && len(*lastName) > 0
	isValidEmail := strings.Contains(*email, "@")
	isValidTicketNumber := *userTickets > 0

	if !isValisName {
		fmt.Println("first name or last name you entered is too short")
		Input(firstName, "Enter your first name    :")
		Input(lastName, "Enter your last name     :")
	}
	if !isValidEmail {
		fmt.Println("email address you entered doesn't contain @ sign")
		Input(email, "Enter your email address :")
	}
	if !isValidTicketNumber {
		fmt.Println("number of tickets you entered is invalid")
		Input(userTickets, "Enter number of tickets  :")
	}
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValisName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValisName, isValidEmail, isValidTicketNumber
}

func IsConferenceBookedOut(remainingTickets uint) bool {
	isTicketsAvailable := remainingTickets > 0
	if !isTicketsAvailable {
		fmt.Printf("\nOur conference is booked out. Come back next year.")
		return true
	}
	return false
}
