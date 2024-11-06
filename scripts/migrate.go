package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"tonx_backend/internal/database"
	"tonx_backend/internal/database/models"
	"tonx_backend/internal/structure"
)

func main() {
	migrate()
}

func migrate() {
	db := database.Connect()
	db.AutoMigrate(&models.Flight{})

	datas := getTestDatas()
	batchSize := 1000
	db.CreateInBatches(datas, batchSize)
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
