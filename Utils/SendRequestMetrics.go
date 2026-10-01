package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adamwahyuh/fsd-golang-adamwh/config"
)

func SendRequestMetrics() {
	config.LoadEnv()

	ApiUrl := config.GetEnv("API_URL") + "/metrics"

	// Ambil data Metrics
	metrics, err := GetSystemMetrics()
	if err != nil {
		log.Printf("Error mengambil data sistem: %v\n", err)
		return
	}

	// Ubah data menjadi JSON format
	payload, err := json.Marshal(metrics)
	if err != nil {
		log.Printf("Error mengubah data ke JSON: %v\n", err)
		return
	}

	// Hit Endpoint
	req, err := http.NewRequest("POST", ApiUrl, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Error membuat request: %v\n", err)
		return
	}

	// Set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Gagal mengirim data ke API: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Cek status response dari Laravel (204 No Content = berhasil)
	if resp.StatusCode == http.StatusNoContent || resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		fmt.Printf("[%s] Berhasil mengirim metrics ke API.\n", time.Now().Format("15:04:05"))
	} else {
		log.Printf("Gagal mengirim data. API merespon dengan status code: %d\n", resp.StatusCode)
	}
}
