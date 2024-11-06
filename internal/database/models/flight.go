package models

import (
	"time"
)

type Flight struct {
	ID               int64  `gorm:"primaryKey;auto_increment"`
	DepartureAirport string `gorm:"type:varchar(100)"`
	DepartureTime    time.Time
	ArrivalAirport   string `gorm:"type:varchar(100)"`
	ArrivalTime      time.Time
	Flight           string `gorm:"type:varchar(50)"`
	Price            int32  `gorm:"type:INTEGER"`
	AvailableSeats   int    `gorm:"type:INTEGER"`
	Status           int    `gorm:"type:smallint"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
