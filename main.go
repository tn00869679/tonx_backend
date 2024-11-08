package main

import (
	"fmt"
	_ "tonx_backend/init"
	"tonx_backend/internal/database"
	"tonx_backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server Start!")

	pdb := database.Connect()
	gin.SetMode(gin.DebugMode)
	g := gin.Default()
	routes.Setup(g, pdb)

	g.Run(":8888")
}
