package main

import (
	"booking-app/helper"
	"fmt"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var booking []map[string]string

	helper.Greet(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		firstName, lastName, email, userTickets = helper.GetInput()

		booking = append(booking, map[string]string{
			"firstName":   firstName,
			"lastName":    lastName,
			"email":       email,
			"userTickets": fmt.Sprintf("%d", userTickets),
		})

		if helper.ValidateInput(firstName, lastName, email, userTickets, remainingTickets) {

			var noTicketsRemaining bool
			remainingTickets, noTicketsRemaining = helper.BookTicket(remainingTickets, conferenceName, booking, firstName, lastName, email, userTickets)

			type RandomStruct struct {
				Value1 int
				Value2 string
				Value3 float64
			}

			randomStruct := RandomStruct{
				Value1: 42,
				Value2: "randomString",
				Value3: 3.14,
			}

			fmt.Printf("RandomStruct: %+v\n", randomStruct)

			if noTicketsRemaining {
				fmt.Printf("Sorry, we are sold out. Please check back later.\n")
				break
			}

		} else {
			helper.MessError(firstName, lastName, email, userTickets, remainingTickets)
		}
	}
}
