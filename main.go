package main

import (
	"server-monitor/checker"
	"server-monitor/logger"
	"time"
)

func main() {
	servers := []string{"https://example.com", "https://google.com"}
	log := logger.SetupLogger()

	for {
		results := checker.CheckServers(servers)
		for _, result := range results {
			logger.LogResult(log, result)
		}
		time.Sleep(10 * time.Second)
	}
}
