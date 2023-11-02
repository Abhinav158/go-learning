package main

import (
	"fmt" 
	"strconv"

	"booking-app/helper"
)

//Package level variables - defined outside all the functions so that it can be accessed by all functions 
var conferenceName = "GoLang sessions"
// Alternate way of declaration - conferenceName := "Go Conference"
// Cannot be used for package level variables like this one
// Does not work for constants, or specify an explicit datatype
const conferenceTickets   = 50
var availableTickets uint = 50
// Arrays to store homogenous data elements as a list 
// Here we use Slices which are basically dynamic arrays
var bookings = make([]map[string]string, 0)


func main(){	

	
	
	greetUsers()	

	// Loop to accept new booking information after one user is done booking
	for {		
		firstName, lastName, email, userTickets := getUserInput()

		isValidEmail, isValidName, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, availableTickets)
		
		if isValidName && isValidEmail && isValidTickets{

			bookTickets(userTickets, firstName, lastName, email)
		
			firstNames := getFirstNames()
			fmt.Printf("The bookings we have so far are: %v\n", firstNames)

			if availableTickets == 0 {
				// End the program 

				fmt.Println("Sorry! Our conference is booked out! Please try again next year!")
				break
			}
			
		} else {
			if !isValidName{
				fmt.Println("Oops! You have entered a very short name!")
			}
			if !isValidEmail{
				fmt.Println("Oops! Your email does not contain the '@' character!")
			}
			if !isValidTickets{
				fmt.Println("Oops! You have entered an invalid number of tickets!")
			}
		}
		

	}
	
}

func greetUsers() {
	fmt.Printf("Welcome to %v Conference Booking App!\n", conferenceName)
	fmt.Printf("We have %v tickets available for this Conference\n", conferenceTickets)
	fmt.Printf("Grab your tickets here to attend!\n")
}

func getFirstNames() []string {
	// Alternate declaration of a slice to store only the first names of people who booked
	firstNames := []string{}

	for _, booking := range(bookings){
		// Now this booking is a map containing all the user Data
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}


func getUserInput() (string, string, string, uint){
	
	// Obtaining values from the user
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	
	fmt.Print("How many tickets would you like to purchase? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string){
	// Update remaining tickets in the booking center 
	availableTickets = availableTickets - userTickets

	// Create a map to contain information about each booked user 
	// All keys have the same datatype. Similarly, the values also have same datatypes
	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"]  = lastName
	userData["email"]     = email
	// We need to convert the number of tickets into a string datatype before storing inside the map
	userData["tickets"]   = strconv.FormatUint(uint64(userTickets), 10)
	
	// Add new elements to this slice
	bookings = append(bookings, userData)

	

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirming the same purchase on %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", availableTickets, conferenceName)
}