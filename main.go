package main

import (
	"flag"
	"fmt"
	_ "tonx_backend/init"
	"tonx_backend/internal/database"
	"tonx_backend/internal/routes"
	"tonx_backend/scripts"

	"github.com/gin-gonic/gin"
)

func main() {
	migrate := flag.Bool("migrate", false, "migrate & build datas")
	flag.Parse()
	if *migrate {
		scripts.Migrate()
	}

	fmt.Println("Server Start!")

	pdb := database.Connect()
	gin.SetMode(gin.DebugMode)
	g := gin.Default()
	routes.Setup(g, pdb)

	g.Run(":8888")
}
