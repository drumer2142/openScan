package src

import (
	"fmt"
	"log"
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

func isOpen(protocol string, host net.IP, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout(protocol, fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}

	return false
}

func NetworkScan(ipAddress string) {
	// init discoved ip
	discoveredIPs := []string{}

	//find the network's total Hosts
	startHost, finishHost := CalculateTotalHosts(ipAddress)

	wg := &sync.WaitGroup{}
	for i := startHost; i < finishHost; i++ {

		ip := ConvertIpFromBinary(i)

		wg.Add(1)
		go func() {
			opened := isOpen(tcpProtocol, ip, 80, timeout)
			if opened {
				discoveredIPs = append(discoveredIPs, string(ip))
			}
			wg.Done()
		}()
	}

}

func IsHostAlive(ipAddress string) bool {
	ip, _, err := net.ParseCIDR(ipAddress)
	if err != nil {
		log.Fatal(err)
	}
	return isOpen(tcpProtocol, ip, 80, timeout)
}

func PortScan(ipAddress string) {
	ports := []int{}

	ip, _, err := net.ParseCIDR(ipAddress)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	for port := minPort; port < maxPort; port++ {
		wg.Add(1)
		go func(port int) {
			opened := isOpen(tcpProtocol, ip, port, timeout)
			if opened {
				ports = append(ports, port)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()

	if len(ports) == 0 {
		fmt.Printf("No open porst found\n")
	} else {
		fmt.Printf("Open ports: %v\n", ports)
	}
}
