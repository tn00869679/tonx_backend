package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	apiURL      = "http://localhost:8888/flight/bookTicket"
	totalCalls  = 500 // 總請求數
	concurrency = 250 // 同時併發的 Goroutine 數
)

func main() {
	var wg sync.WaitGroup
	var successCount, failureCount int
	var mu sync.Mutex

	requestBody := []byte(`{
        "flight": "航班 ForTestBookSuccess",
        "departureTime": "2024-11-01T00:00:00Z"
    }`)

	semaphore := make(chan struct{}, concurrency)
	start := time.Now()

	for i := 0; i < totalCalls; i++ {
		wg.Add(1)
		semaphore <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()

			req, err := http.NewRequest("PATCH", apiURL, bytes.NewBuffer(requestBody))
			if err != nil {
				fmt.Println("Failed to create request :", err)
				mu.Lock()
				failureCount++
				mu.Unlock()
				return
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Request failed :", err)
				mu.Lock()
				failureCount++
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			mu.Lock()
			if resp.StatusCode == http.StatusOK {
				successCount++
			} else {
				failureCount++
			}
			mu.Unlock()
		}()
	}

	wg.Wait()
	close(semaphore)

	duration := time.Since(start)
	fmt.Printf("Total Calls: %d\n", totalCalls)
	fmt.Printf("Success: %d\n", successCount)
	fmt.Printf("Failure: %d\n", failureCount)
	fmt.Printf("Duration: %v\n", duration)
}
