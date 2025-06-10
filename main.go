package main

import "fmt"

func main() {
	var conferenceName = "Olyzano Conference"
	const conferenceTickers = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to our %s booking appliacation\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and are still available %d for booking\n",  conferenceTickers, remainingTickets )
	fmt.Println("Get your tickets here to attend the conference")

	var userName  string
	var userTickets int
	fmt.Println("Enter your name:")
	fmt.Scan(&userName)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	if remainingTickets <= 0 {
		fmt.Println("Our conference is sold out. Please check back later for future events.")
		} else {
			fmt.Printf("We still have %d tickets available for booking.\n", remainingTickets)
			fmt.Printf("Thank you %s for booking %d tickets. You will receive a confirmation email shortly.\n", userName, userTickets)
	}
	fmt.Println("We hope to see you at the conference!")
	fmt.Println("Thank you for using our booking application!")
	fmt.Println("Have a great day!")


}