package models

import (
	"time"
)

type Flight struct {
	ID               uint64    `gorm:"primaryKey;auto_increment"`
	DepartureAirport string    `gorm:"type:varchar(100)"`
	DepartureTime    time.Time `gorm:"index:idx_departure_time"`
	ArrivalAirport   string    `gorm:"type:varchar(100)"`
	ArrivalTime      time.Time
	Flight           string `gorm:"type:varchar(50);index:idx_flight"`
	Price            int32
	AvailableSeats   int
	Overbooking      int
	Status           bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
