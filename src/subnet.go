package src

import (
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
)

func CalculateTotalHosts(ipAddress string) uint64 {
	_, network, _ := net.ParseCIDR(ipAddress)
	return cidr.AddressCount(network)
}
