package main

import "fmt"

func main(){	

	var conferenceName = "GoLang sessions"
	// Alternate way of declaration - conferenceName := "Go Conference"
	// Does not work for constants, or specify an explicit datatype


	const conferenceTickets = 50
	var availableTickets uint = 50

	fmt.Println("Welcome to", conferenceName, "Booking Application!")

	// We can also format our strings to help make display of variables easier
	fmt.Printf("We have a total of %v tickets \n", conferenceTickets)
	fmt.Println("Only", availableTickets, "tickets available!")
	fmt.Println("Get your tickets here to attend!")

	// Arrays to store homogenous data elements as a list 
	// Maximum size of the array is 50

	// Here we use slices which are basically dynamic arrays
	var bookings []string

	// Loop to accept new booking information after one user is done booking
	for {
		// Obtaining values from the user
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// Users can only book a non-negative number of tickets 

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		
		fmt.Print("How many tickets would you like to purchase? ")
		fmt.Scan(&userTickets)

		

		// Add new elements to this slice
		bookings = append(bookings, firstName + " " + lastName)
		


		// Update remaining tickets in the booking center 
		availableTickets = availableTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirming the same purchase on %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", availableTickets, conferenceName)
		
		fmt.Printf("These are all our bookings: %v\n", bookings)

	}
	
}