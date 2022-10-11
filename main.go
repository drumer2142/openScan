package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/drumer2142/openScan/src"
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

	if *portScan {
		fmt.Printf("Scanning %s for open ports....", *ip)
		src.PortScan(*ip)
	}

	if *aliveHost {
		fmt.Printf("Scanning for host %s....", *ip)
		hostAlive := src.IsHostAlive(*ip)
		if hostAlive {
			fmt.Printf("Host %s is alive\n", *ip)
			os.Exit(0)
		}
		fmt.Printf("Host %s is down\n", *ip)
	}

	if *netScan {
		fmt.Println("Scanning the subnet for possible host alive....")
		fmt.Printf("Host Found Alive %v", src.NetworkScan(*ip))
	}

	os.Exit(0)
}
