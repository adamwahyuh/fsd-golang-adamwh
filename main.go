package main

import (
	"log"
	"time"

	utils "github.com/adamwahyuh/fsd-golang-adamwh/Utils"
)

func main() {
	log.Println("Service Running")

	utils.SendRequestMetrics()

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	// Looping selama service berjalan
	for range ticker.C {
		utils.SendRequestMetrics()
	}
}
