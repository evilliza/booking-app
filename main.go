package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint8 = 50

var conferenceName = "Go Conference"

var remainingTickets uint8 = 50
var bookings = make([]UserData, 0) // make(map[string]string{})

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// for remainingTickets > 0 && len(bookings) < 50 {

	// checking for valid info
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		// bookticket func return
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)                                              // adding a thread
		go sendTicket(userTickets, firstName, lastName, email) // concurrency

		// call func print first names
		firstNames := getFirstNames()
		fmt.Printf("Names of people attending the confirence: %v \n", firstNames)

		// conditional statement, dada type true or faulse
		if remainingTickets == 0 {
			// end programm
			fmt.Println("Sorry, but our conference is fully booked")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("Your first name or last name is too short")
		}
		if !isValidEmail {
			fmt.Println("Your email does not contain @ sign")

		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets is invalid")
		}

	}

	// }

	wg.Wait() // waites until go sendtickets is done

}

func ValidateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint8) {
	panic("unimplemented")
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // version 1.0 before creating map 	// HOW TO SEPERATE STRINGS BY SPACES GOOGLED pkg "strings"
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint8, firstName, lastName, email string) {
	remainingTickets = remainingTickets - userTickets

	// creating map for users
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// version with maps
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberoftickets"] = strconv.FormatUint(uint64(userTickets), 10) // conv fron uint type to string type

	bookings = append(bookings, userData) // using apppend
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint8, firstName, lastName, email string) {
	time.Sleep(20 * time.Second) // to simulate a delay
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done() // removes the thread for the waiting list
}
