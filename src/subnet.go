package src

import (
	"encoding/binary"
	"log"
	"net"
)

func CalculateTotalHosts(ipAddress string) (uint32, uint32) {
	_, network, err := net.ParseCIDR(ipAddress)
	if err != nil {
		log.Fatal(err)
	}

	mask := binary.BigEndian.Uint32(network.Mask)
	start := binary.BigEndian.Uint32(network.IP)

	finish := (start & mask) | (mask ^ 0xffffffff)

	return start, finish
}

func ConvertIpFromBinary(i uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, i)

	return ip
}
