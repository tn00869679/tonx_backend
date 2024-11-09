package scripts

import (
	"encoding/json"
	"io"
	"log"
	"os"
	_ "tonx_backend/init"
	"tonx_backend/internal/database"
	"tonx_backend/internal/database/models"
	"tonx_backend/internal/structure"
)

func Migrate() {
	db := database.Connect()
	db.AutoMigrate(&models.Flight{})

	datas := getTestDatas()
	batchSize := 1000
	db.CreateInBatches(datas, batchSize)

	postgres, err := db.DB()
	if err != nil {
		log.Fatalln("Failed to get postgres from gorm DB :", err)
	}
	postgres.Close()
}

func getTestDatas() *[]structure.Flight {
	jsonFile, err := os.Open("./scripts/flight.json")
	if err != nil {
		log.Fatalln("無法打開檔案 :", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln("無法讀取檔案 :", err)
	}

	var datas []structure.Flight
	if err := json.Unmarshal(byteValue, &datas); err != nil {
		log.Fatalln("無法解析 JSON :", err)
	}

	return &datas
}
