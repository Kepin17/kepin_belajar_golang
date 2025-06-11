package main

import "fmt"

func main() {
	var conferenceName = "Olyzano Conference"
	const conferenceTickers = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to our %s booking appliacation\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and are still available %d for booking\n",  conferenceTickers, remainingTickets )
	fmt.Println("Get your tickets here to attend the conference")


	var bookings = [50]string{}

	var userFirstName  string
	var userLastName  string
	var userEmail string
	var userTickets int
	fmt.Println("Enter your first name:")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter your Last name:")
	fmt.Scan(&userLastName)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets

	bookings[0] = userName


	if userTickets <= remainingTickets {	
		fmt.Printf("The whole booking list is: %s\n", bookings)
		fmt.Printf("The whole firstbooking list is: %v\n", bookings[0])
		fmt.Printf("%s booked %d tickets. You have %d tickets remaining.\n", userName, userTickets, remainingTickets)
	}else {
		fmt.Printf("Sorry %s, we only have %d tickets remaining. You cannot book %d tickets.\n", userName, remainingTickets, userTickets)
	}
}