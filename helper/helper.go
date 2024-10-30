package helper

import (
	"fmt"
	"strings"
)

func Greet(conferenceName string, conferenceTickets int, remainingTickets int) {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaible.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

}

func PrintFirstName(bookings []string) []string {
	var firstNames []string
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}

func ValidateInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) bool {
	isValideName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValideName && isValidEmail && isValidTickets
}
func GetInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

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

func BookTicket(remainingTickets int, conferenceName string, bookings []map[string]string, firstName string, lastName string, email string, userTickets int) (int, bool) {
	remainingTickets = remainingTickets - userTickets

	// Create a map to store user information
	userInfo := map[string]string{
		"firstName": firstName,
		"lastName":  lastName,
		"email":     email,
		"tickets":   fmt.Sprintf("%d", userTickets),
	}

	bookings = append(bookings, userInfo)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets remaining for %v\n", remainingTickets, conferenceName)

	// Extract first names from the bookings
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	fmt.Printf("These are the first names of all the attendees: %v\n", firstNames)

	noTicketsRemaining := remainingTickets == 0 || remainingTickets < 0

	return remainingTickets, noTicketsRemaining
}

func MessError(firstName string, lastName string, email string, userTickets int, remainingTickets int) {
	isValideName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	if !isValideName {
		fmt.Println("Please enter a valid first and last name.")
	}
	if !isValidEmail {
		fmt.Println("Please enter a valid email address.")
	}
	if !isValidTickets {
		fmt.Println("Please enter a valid number of tickets.")
	}
	if userTickets > remainingTickets {
		fmt.Printf("Sorry, we only have %v tickets remaining. Please try again.\n", remainingTickets)
	}
}
