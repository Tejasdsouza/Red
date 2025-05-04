package main

import (
	"bufio"
	"fmt"
	"os"
	"RED/internal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nWelcome to RED:")
		fmt.Println("1. Port Scanner")
		fmt.Println("2. DNS Resolver")
		fmt.Println("3. Latency Monitor")
		fmt.Println("4. Server Availability Checker")
		fmt.Println("5. Exit")
		fmt.Print("> ")

		input, _, _ := reader.ReadLine()

		switch string(input) {
		case "1":
			internal.ScanLocalPorts()
		case "2":
			fmt.Print("Enter domain: ")
			domain, _ := reader.ReadString('\n')
			internal.ResolveDomain(domain)
		case "3":
			internal.CheckLatency()
		case "4":
			fmt.Print("Enter domain: ")
			domain, _ := reader.ReadString('\n')
			internal.CheckAvailability(domain)
		case "5":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
