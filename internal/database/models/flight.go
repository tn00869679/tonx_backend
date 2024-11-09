package models

import (
	"time"
)

type Flight struct {
	ID               uint64    `gorm:"primaryKey;auto_increment"`
	DepartureAirport string    `gorm:"type:varchar(100);index:idx_dep_arr_airport"`
	DepartureTime    time.Time `gorm:"index:idx_departuretime_flight"`
	ArrivalAirport   string    `gorm:"type:varchar(100);index:idx_dep_arr_airport"`
	ArrivalTime      time.Time
	Flight           string `gorm:"type:varchar(50);index:idx_departuretime_flight"`
	Price            int32
	AvailableSeats   int
	Overbooking      int
	Status           bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
