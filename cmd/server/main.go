package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development purposes
	},
}

func main() {
	port := getPort()

	http.HandleFunc("/ws", handleWebSocket)

	log.Printf("Server listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getPort() int {
	portStr := os.Getenv("SERVER_PORT")
	if portStr == "" {
		return 8080 // Default port
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid SERVER_PORT: %v", err)
	}

	return port
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(30 * time.Second) // Calculate average every 30 seconds
	defer ticker.Stop()

	latencies := make([]time.Duration, 0)

	for {
		select {
		case <-ticker.C:
			avgLatency := calculateAverageLatency(latencies)
			log.Printf("Average latency: %v", avgLatency)

			err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%f", avgLatency.Seconds())))
			if err != nil {
				log.Println("write:", err)
				return
			}
			latencies = make([]time.Duration, 0) // Reset latencies
		default:
			latency, err := pingCloudflare()
			if err != nil {
				log.Println("ping:", err)
				continue
			}
			latencies = append(latencies, latency)
			time.Sleep(1 * time.Second) // Ping every 1 second
		}
	}
}

func pingCloudflare() (time.Duration, error) {
	start := time.Now()
	resp, err := http.Get("http://1.1.1.1")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	end := time.Now()

	return end.Sub(start), nil
}

func calculateAverageLatency(latencies []time.Duration) time.Duration {
	if len(latencies) == 0 {
		return 0
	}

	var total time.Duration
	for _, latency := range latencies {
		total += latency
	}

	return total / time.Duration(len(latencies))
}
