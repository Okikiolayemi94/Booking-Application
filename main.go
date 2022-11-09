package main

import "fmt"

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets = 50
var userName string
var userTicket int
var userChoice int

func welcome() {
	fmt.Printf("Welcome to %s booking application", conferenceName)
	fmt.Println("\n-----------------------------------")
	fmt.Printf("\nWe have a total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("\nGet your tickets here to attend")
	fmt.Println("-----------------------------------")
}
func getUsername() {
	fmt.Print("Please enter your name to book a ticket: ")
	_, err := fmt.Scan(&userName)
	if err != nil {
		return
	}
}
func getUserChoice() {
	fmt.Println("Press 1 to get a ticket or Press 2 to exit the app")
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
}
func main() {

	welcome()
	getUsername()
	getUserChoice()

	fmt.Printf("User %s has booked %v ticket(s)", userName, userTicket)
	fmt.Printf("\nWe have %v tickets available for grabs...", remainingTickets)
	fmt.Println("\n-----------------------------------")
	fmt.Println("\nWant to make another booking?\t 1. Yes\t 2. No.")
}
