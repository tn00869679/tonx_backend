package structure

import "time"

type Flight struct {
	ID               uint64
	DepartureAirport string
	DepartureTime    time.Time
	ArrivalAirport   string
	ArrivalTime      time.Time
	Flight           string
	Price            uint32
	AvailableSeats   uint16
	Status           int8
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
