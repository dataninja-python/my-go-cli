package cli

import (
	"fmt"
	"log"
)

// These variables define the application's defaults
const conferenceTickets uint = 50

type Seat struct {
	ID  string `json:"id"`
	Row string `json:"row"`
	Col uint   `json:"col"`
}

// Event type defines an event. It provides a unique id for each event, tells the type of event,
// its name, and available tickets initially
type Event struct {
	ID               string `json:"id"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	RemainingTickets uint   `json:"remainingTickets"`
	Seats            []Seat `json:"seats"`
}

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}

type Attendee struct {
	ID      string   `json:"id"`
	Tickets []Ticket `json:"tickets"`
	Person
}

type Booking struct {
	ID string `json:"id"`
}

// initApp inserts reasonable defaults to be used to create a new Event.
func initApp() (string, uint, error) {
	var (
		defaultConferenceName   string = "Go Conference"
		defaultRemainingTickets        = conferenceTickets
	)
	return defaultConferenceName, defaultRemainingTickets, nil
}

// RunBookingApp implements the command line version of a booking application.
// It creates an event and books tickets by attaching it to gathered user information.
func RunBookingApp() {
	conferenceName, remainingTickets, err := initApp()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Define the conference's name

	line1 := fmt.Sprintf("Welcome to the %s booking application", conferenceName)
	line2 := fmt.Sprintf("We have a total of %d tickets and %d are still available", conferenceTickets, remainingTickets)
	line3 := fmt.Sprintf("Get your tickets here to attend.")
	introductionLines := []string{line1, line2, line3}

	for _, line := range introductionLines {
		fmt.Println(line)
	}

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
