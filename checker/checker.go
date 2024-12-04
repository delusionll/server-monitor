package checker

import (
	"net/http"
	"time"
)

type ServerStatus struct {
	URL          string
	IsAvailable  bool
	ResponseTime time.Duration
}

func CheckHTTP(url string) ServerStatus {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return ServerStatus{URL: url, IsAvailable: false}
	}
	return ServerStatus{
		URL:          url,
		IsAvailable:  true,
		ResponseTime: time.Since(start),
	}
}

func CheckServers(urls []string) []ServerStatus {
	results := make([]ServerStatus, len(urls))
	done := make(chan ServerStatus)

	for _, url := range urls {
		go func(u string) {
			done <- CheckHTTP(u)
		}(url)
	}

	for i := range urls {
		results[i] = <-done
	}

	return results
}
