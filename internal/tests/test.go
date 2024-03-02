package tests

/*

// ImmutableEventData
type ImmutableEventData struct {
	maxEventTickets uint `json:"max_event_tickets"`
	minEventTickets uint `json:"min_event_tickets"`
}

// Event
type Event struct {
	EventID          string `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventDate        string `json:"event_date"`
	EventTime        string `json:"event_time"`
	EventTimezone    string `json:"event_timezone"`
	ImmutableEventData
}

// Create a constructor to make creating events easier
func NewEvent(name string, ticketsold uint, remainingTickets uint, maxTickets ...uint, minTickets ...uint) *Event {
	// make the minimum number of tickets 0 unless otherwise provided
	var minT, maxT uint = 0, 50
	fmt.Println("Defaults are Minimum Tickets: %v and Maximum Tickets: %v", minT, maxT)
	if len(minTickets) > 0 {
		minT = minTickets[0]
	}
	// make the default max tickets 50 unless otherwise provided
	if len(maxTickets) > 0 {
		maxT = maxTickets[0]
	}

	return &Event{
		Name:             name,
		TicketsSold:      ticketsold,
		RemainingTickets: remainingTickets,
		ImmutableEventData: ImmutableEventData{
			minTickets: minT,
			maxTickets: maxT,
		},
	}
}

// Functions to work with mutable data
func (e *Event) SetName(name string) {
	e.Name = name
}

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

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

// Functions to work with immutable data
func (e *Event) GetMaxTickets() uint {
	return e.maxTickets
}

func (e *Event) GetMinTickets() uint {
	return e.minTickets
}

*/
