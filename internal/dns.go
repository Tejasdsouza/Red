package internal

import (
    "fmt"
    "net"
    "net/http"
    "strings"
    "time"
)

func ResolveDomain(domain string) {
    fmt.Printf("[RED] Resolving DNS for %s...\n", domain)
    domain = strings.TrimSpace(domain)
    domain = strings.TrimPrefix(domain, "http://")
    domain = strings.TrimPrefix(domain, "https://")
    domain = strings.TrimSuffix(domain, "/")

    ips, err := net.LookupHost(domain)
    if err != nil {
        fmt.Printf("[RED] DNS resolution failed: %v\n", err)
        return
    }
    fmt.Printf("[RED] Found IPs:\n")
    for _, ip := range ips {
        fmt.Printf("  ➤ %s\n", ip)
    }

    fmt.Printf("[RED] Probing HTTP banner for %s...\n", domain)
    client := http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get("http://" + domain)
    if err != nil {
        fmt.Printf("[RED] Failed to get HTTP banner: %v\n", err)
    } else {
        defer resp.Body.Close()
        fmt.Printf("[RED] HTTP Response: %s %s\n", resp.Proto, resp.Status)
        for key, val := range resp.Header {
            fmt.Printf("[RED] %s: %s\n", key, strings.Join(val, ", "))
        }
    }

    if len(ips) > 0 {
        fmt.Printf("[RED] Scanning common ports on %s...\n", ips[0])
        commonPorts := []int{21, 22, 23, 25, 53, 80, 110, 143, 443, 8080}
        for _, port := range commonPorts {
            address := net.JoinHostPort(ips[0], fmt.Sprintf("%d", port))
            conn, err := net.DialTimeout("tcp", address, 2*time.Second)
            if err != nil {
                fmt.Printf("  ➤ Port %d: Closed\n", port)
            } else {
                fmt.Printf("  ➤ Port %d: Open\n", port)
                conn.Close()
            }
        }
    }
	fmt.Printf("[RED] DNS resolution and probing completed.\n")
}