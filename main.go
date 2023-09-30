package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
	
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter no. of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber := userTickets > 0 && userTickets > remainingTickets

		// isInvalidCity := city != "Singapore" || city != "London"

		if !isValidName || !isValidEmail || !isValidTicketNumber {
			if !isValidName {
				fmt.Println("Your first name or last name is too short")
			} else if !isValidEmail {
				fmt.Println("Enter correct email address")
			} else if !isValidTicketNumber {
				fmt.Println("No. of tickets you entered is invalid")
			}
			continue
		}
	
		remainingTickets -= userTickets
		bookings = append(bookings, firstName + " " + lastName)
	
		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		fmt.Printf("Array length: %v\n", len(bookings))
	
		fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation email at %v soon.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining here to attend.", remainingTickets)
	
		firstNames := []string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These first names of our bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Comeback next year")
			break
		}
	}


	city := "Delhi"

	switch city {
	case "Lucknow":

	case "Agra", "Varanasi":
	
	case "Allahabad", "Mirzapur":

	case "Hapur":

	default:
		fmt.Println("No valid city selected")
	}

}
