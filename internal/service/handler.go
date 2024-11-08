package service

import (
	"net/http"
	"tonx_backend/internal/database/models"
	"tonx_backend/internal/structure/error_code"
	"tonx_backend/internal/structure/request"
	"tonx_backend/internal/structure/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func (s *Service) getFlightList(c *gin.Context, flights *[]models.Flight, req *request.FlightList) bool {
	offset := (req.Page - 1) * req.PerPage
	if result := s.pdb.Where("departure_airport = ?", req.DepartureAirport).
		Where("arrival_airport = ?", req.ArrivalAirport).
		Where("departure_time >= ?", req.DepartureTime).
		Order("id ASC").
		Offset(offset).
		Limit(req.PerPage).
		Find(flights); result.Error != nil {
		response.ServerError(c)
		return false
	}
	return true
}

func (s *Service) updateFlightInfo(c *gin.Context, flight *models.Flight, req *request.BookTicket) bool {
	tx := s.pdb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("flight = ?", req.Flight).
		Where("departure_time >= ?", req.DepartureTime).
		Order("departure_time ASC").
		First(flight); result.Error != nil {
		tx.Rollback()
		response.Error(c, http.StatusServiceUnavailable, result.Error.Error(), error_code.Server)
		return false
	}

	if !flight.Status {
		tx.Rollback()
		response.NotAcceptableError(c)
		return false
	}

	if flight.AvailableSeats != 0 {
		flight.AvailableSeats -= 1
	} else {
		flight.Overbooking -= 1
	}

	if flight.Overbooking == 0 {
		flight.Status = false
	}

	if err := tx.Save(flight).Error; err != nil {
		tx.Rollback()
		response.ServerError(c)
		return false
	}
	return tx.Commit().Error == nil
}
