package structure

import "time"

type Flight struct {
	ID               uint64
	DepartureAirport string
	DepartureTime    time.Time
	ArrivalAirport   string
	ArrivalTime      time.Time
	Flight           string
	Price            int32
	AvailableSeats   int
	Overbooking      int
	Status           bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
