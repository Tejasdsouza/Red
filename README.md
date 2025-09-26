# Red

Red is a command-line tool designed for network reconnaissance and exploration. It provides utilities for port scanning, DNS resolution, latency monitoring, and server availability checking.

## Features

1. **Port Scanner**  
   Scan common local ports to check if they are open or closed.

2. **DNS Resolver**  
   Resolve a domain to its associated IP addresses, probe HTTP banners, and scan common ports.

3. **Latency Monitor**  
   Measure the latency.

4. **Server Availability Checker**  
   Check if a server is reachable and available.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Tejasdsouza/Red.git
   cd Red
   ```

2. Build the project:
   ```bash
   go build ./cmd/red.go
   ```

3. Run the project:
   ```bash
   ./red
   ```

## Usage

When you run the program, you will see the following menu:

```plaintext
Welcome to RED:
1. Port Scanner
2. DNS Resolver
3. Latency Monitor
4. Server Availability Checker
5. Exit
```

### 1. Port Scanner
Select option `1` to scan common local ports (e.g., 22, 80, 443, etc.) on your machine.

### 2. DNS Resolver
Select option `2` to resolve a domain. Enter a domain (e.g., `example.com`), and the tool will:
- Resolve the domain to its IP addresses.
- Probe the HTTP banner of the domain.
- Scan common ports (e.g., 21, 22, 80, 443, etc.).

### 3. Latency Monitor
Select option `3` to measure the latency.

### 4. Server Availability Checker
Select option `4` to check if a server is reachable and available.

### 5. Exit
Select option `5` to exit the program.
