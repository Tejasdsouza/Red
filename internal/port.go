package internal

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ScanLocalPorts() {
	fmt.Println("[RED] Scanning common local ports...")
	ports := []int{22, 80, 443, 8080, 3306, 5432, 6379, 25565}
	for _, port := range ports {
		address := "127.0.0.1:" + strconv.Itoa(port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			fmt.Printf("[RED] Port %d closed or unreachable\n", port)
			continue
		}
		fmt.Printf("[RED] Port %d is open!\n", port)
		conn.Close()
	}
}
