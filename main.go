package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %s booking application", conferenceName)
	fmt.Printf("\nWe have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("\nGet your tickets here to attend")

	var userName string
	var userTicket int

	fmt.Print("Please enter your name to book a ticket: ")
	fmt.Scan(&userName)

	fmt.Println("Press 1 to get a ticket or Press 2 to exit the app")
	var userChoice int
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		fmt.Println("Your booking was Successful")
	case 2:
		fmt.Println("Oops, Goodbye")
	default:
		fmt.Println("Invalid request, Try Again")

	}
	if userChoice == 1 {
		remainingTickets = conferenceTickets - 1
		userTicket++
	} else {
		remainingTickets = conferenceTickets
	}
	fmt.Printf("User %s has booked %v ticket(s)", userName, userTicket)
	fmt.Printf("\nWe have %v tickets available for grabs...", remainingTickets)
}
