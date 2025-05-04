package internal

import (
	"fmt"
	"time"
	"net"
)

func CheckLatency() {
	host := "google.com"
	const count = 5
	var total time.Duration

	fmt.Printf("[RED] Pinging %s...\n", host)

	for i := 0; i < count; i++ {
		start := time.Now()
		_, err := net.LookupIP(host)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		elapsed := time.Since(start)
		total += elapsed
		time.Sleep(200 * time.Millisecond)
	}

	avg := total / count
	fmt.Printf("[RED] Average latency to %s: %v\n", host, avg)
}