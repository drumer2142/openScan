package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/drumer2142/openScan/src"
)

var (
	defaultSubnetMask string = "/24"
)

func main() {
	ip := flag.String("ip", "", "Must specify ip")
	portScan := flag.Bool("p", false, "Port Scan")
	aliveHost := flag.Bool("a", false, "Is host Alive")
	netScan := flag.Bool("sc", false, "Scan the network")
	flag.Parse()

	if len(*ip) == 0 {
		os.Exit(0)
	}

	ipFormatted := CheckIPforMask(*ip)

	if *portScan {
		fmt.Printf("Scanning %s for open ports....", ipFormatted)
		src.PortScan(ipFormatted)
	}

	if *aliveHost {
		fmt.Printf("Scanning for host %s....", ipFormatted)
		hostAlive := src.IsHostAlive(ipFormatted)
		if hostAlive {
			fmt.Printf("Host %s is alive\n", ipFormatted)
			os.Exit(0)
		}
		fmt.Printf("Host %s is down\n", ipFormatted)
	}

	if *netScan {
		color.Red("Scanning the subnet for possible host alive....It may take a while")
		ipArray := src.NetworkScan(ipFormatted)
		for i := 0; i < len(ipArray); i++ {
			fmt.Println("Host Found Alive", ipArray[i])
		}
	}

	os.Exit(0)
}

func CheckIPforMask(ip string) string {
	maskSplit := strings.Split(ip, "/")
	if len(maskSplit) <= 1 {
		return ip + defaultSubnetMask
	}

	return ip
}
