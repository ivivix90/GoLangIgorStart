package main

import "fmt"

func main() {

	//var conferenceName string = "Go conference" it is the same as:
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Conference tickets is: %T, remaining tickets is: %T, conference name is: %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availabe\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email adress: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do You want to buy?: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	bookings[0] = firstName + " " + lastName

	fmt.Printf("THE HOLE ARRAY: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The array type: %T\n", bookings)
	fmt.Printf("The array length: %v\n", len(bookings))

	fmt.Printf("Thank You %v %v for buying %v tickets! Confirmation sent at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v\n", remainingTickets, conferenceName)

}
