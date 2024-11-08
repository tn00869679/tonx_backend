package routes

import (
	"tonx_backend/internal/middleware"
	"tonx_backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(g *gin.Engine, pdb *gorm.DB) {
	service := service.NewService(pdb)
	g.Use(middleware.Cors)
	api := g.Group("/flight")
	{
		api.GET("/list", service.FlightList)
		api.PATCH("/bookTicket", service.BookTicket)
	}
}
