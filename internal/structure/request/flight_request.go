package request

import "time"

type FlightList struct {
	DepartureAirport string    `form:"departureAirport"     binding:"required,min=1,max=100"`
	DepartureTime    time.Time `form:"departureTime"     binding:"required"`
	ArrivalAirport   string    `form:"arrivalAirport"     binding:"required,min=1,max=100"`
	Page             int       `form:"page"    binding:"required,gte=1"`
	PerPage          int       `form:"perPage"    binding:"required,gte=1"`
}

type BookTicket struct {
	DepartureTime time.Time `json:"departureTime"     binding:"required"`
	Flight        string    `json:"flight"     binding:"required,min=1,max=50"`
}
