package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

var conferenceName = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketsNumber := helper.IsValidInputs(firstName, lastName, email, uint(userTickets), remainingTickets)

	if isValidEmail && isValidName && isValidTicketsNumber {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(uint(userTickets), firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}

		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}

		if !isValidTicketsNumber {
			fmt.Println("umber of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking applaication\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets her to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email adress:")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets int, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - uint(userTickets)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: uint(userTickets),
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v", bookings)

	fmt.Printf("Thnak you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %vto email adress %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done()
}
