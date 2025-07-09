package network

import (
	"net/http"
	"time"
)

// URLCheckResult holds the result of checking a URL
type URLCheckResult struct {
	URL      string `json:"url"`
	Verified bool   `json:"verified"`
	Status   int    `json:"status"`
}

func CheckURL(url string, timeout int) URLCheckResult {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		return URLCheckResult{
			URL:      url,
			Verified: false,
			Status:   0,
		}
	}
	defer resp.Body.Close()

	verified := resp.StatusCode >= 200 && resp.StatusCode < 400

	return URLCheckResult{
		URL:      url,
		Verified: verified,
		Status:   resp.StatusCode,
	}

}
