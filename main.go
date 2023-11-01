package main

import (
	"fmt" 
	"strings"
)

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
		

		if userTickets > availableTickets{
			// Invalid input 
			fmt.Printf("Oops! We only have %v tickets remaining. Try again!", availableTickets)
			continue
		}
		// Update remaining tickets in the booking center 
		availableTickets = availableTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirming the same purchase on %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", availableTickets, conferenceName)
		
		// Alternate declaration of a slice to store only the first names of people who booked
		firstNames := []string{}

		for _, booking := range(bookings){
			// Now we have access to each element 

			// Strings package to split based on the space in the string 
			var  names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are all our bookings so far: %v\n", firstNames)
		
		if availableTickets == 0{
			// End the booking process
			fmt.Printf("Our conference is booked out. Hope to see you next year!")
			break 
		}

	}
	
}