package logger

import (
	"log"
	"os"
	"server-monitor/checker"
)

func SetupLogger() *log.Logger {
	file, err := os.OpenFile("monitor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Could not create log file: %v", err)
	}
	return log.New(file, "MONITOR ", log.LstdFlags)
}

func LogResult(logger *log.Logger, result checker.ServerStatus) {
	status := "UP"
	if !result.IsAvailable {
		status = "DOWN"
	}
	logger.Printf(
		"Server: %s | Status: %s | Response Time: %v\n",
		result.URL,
		status,
		result.ResponseTime,
	)
}
