// Helper functions of the main file 

package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, availableTickets uint) (bool, bool, bool){

	// Valid name Criteria: Username must be atleast 2 characters long
	var isValidName = len(firstName) >=2 && len(lastName) >= 2

	// Email Validation Criteria: Must contain certain characters in the input 
	var isValidEmail = strings.Contains(email, "@")

	// Number of tickets entered by user should be greater than 0
	var isValidTickets = userTickets >= 0 && userTickets <= availableTickets

	// In Go, we can return any number of values from a function instead of the conventional single value
	return isValidEmail, isValidName, isValidTickets
}


