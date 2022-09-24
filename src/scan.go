package src

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	// icmpProtocol string = "ip4:icmp"
	tcpProtocol string = "tcp"
	minPort     int    = 1
	maxPort     int    = 1024
	timeout            = time.Microsecond * 200
)

func isOpen(protocol string, host string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout(protocol, fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}

	return false
}

func IsHostAlive(ipAddress string) bool {
	return isOpen(tcpProtocol, ipAddress, 80, timeout)
}

func PortScan(ipAddress string) {
	ports := []int{}

	wg := &sync.WaitGroup{}
	for port := minPort; port < maxPort; port++ {
		wg.Add(1)
		go func(port int) {
			opened := isOpen(tcpProtocol, ipAddress, port, timeout)
			if opened {
				ports = append(ports, port)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()

	if len(ports) == 0 {
		fmt.Printf("No Ports found open for this IP\n")
	} else {
		fmt.Printf("Opened ports: %v\n", ports)
	}
}
