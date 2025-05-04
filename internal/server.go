package internal

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func CheckAvailability(domain string) {
	domain = strings.TrimSpace(domain)
	if !strings.HasPrefix(domain, "http") {
		domain = "https://" + domain
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(domain)
	if err != nil {
		fmt.Println("[RED] Server unreachable:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("[RED] Server %s is reachable, Status Code: %d\n", domain, resp.StatusCode)
}
