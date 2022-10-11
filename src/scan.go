package src

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/drumer2142/openScan/helpers"
)

var (
	tcpProtocol string = "tcp"
	// icmpProtocol string = "ip4:icmp"
	minPort int = 1
	maxPort int = 1024
	timeout     = time.Millisecond * 200
)

func isOpen(protocol string, host net.IP, port int, timeout time.Duration) bool {
	address := helpers.FormatIPandPort(host, port)
	conn, err := net.DialTimeout(protocol, address, timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}

	return false
}

func NetworkScan(ipAddress string) []string {
	// init discoved ip
	discoveredIPs := []string{}

	//find the network's total Hosts
	startHost, finishHost := CalculateTotalHosts(ipAddress)

	// wg := &sync.WaitGroup{}
	for i := startHost; i < finishHost; i++ {

		ip := ConvertIpFromBinary(i)

		echoCheck := isOpen(tcpProtocol, ip, 80, timeout)
		if echoCheck {
			discoveredIPs = append(discoveredIPs, helpers.FormatIP(ip))
		}
	}

	return discoveredIPs

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
