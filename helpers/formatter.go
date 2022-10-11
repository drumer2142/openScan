package helpers

import (
	"fmt"
	"net"
)

func FormatIP(ipAddress net.IP) string {
	return fmt.Sprintf("%s", ipAddress)
}

func FormatIPandPort(ipAddress net.IP, port int) string {
	return fmt.Sprintf("%s:%d", ipAddress, port)
}
