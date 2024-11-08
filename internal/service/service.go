package service

import (
	"tonx_backend/internal/database/models"
	"tonx_backend/internal/structure/request"
	"tonx_backend/internal/structure/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	pdb *gorm.DB
}

func NewService(pdb *gorm.DB) *Service {
	return &Service{pdb}
}

func checkInput[T any](c *gin.Context, bind *T) bool {
	if err := c.ShouldBind(bind); err != nil {
		response.InputError(c)
		return false
	}
	return true
}

func (s *Service) FlightList(c *gin.Context) {
	var req request.FlightList
	if !checkInput(c, &req) {
		return
	}

	var flights []models.Flight
	if !s.getFlightList(c, &flights, &req) {
		return
	}

	response.OK(c, &flights, "")
}

func (s *Service) BookTicket(c *gin.Context) {
	var req request.BookTicket
	if !checkInput(c, &req) {
		return
	}

	var flights models.Flight
	if !s.updateFlightInfo(c, &flights, &req) {
		return
	}

	response.OK(c, nil, "")
}
