package src

import (
	"fmt"
	"testing"
)

func TestCalculateSubnetRange(t *testing.T) {

	type Case struct {
		Range  string
		Start  uint32
		Finish uint32
	}

	cases := []Case{
		{
			Range:  "10.10.20.0/24",
			Start:  168432640,
			Finish: 168432895,
		},
		{
			Range:  "10.10.20.0/32",
			Start:  168432640,
			Finish: 168432640,
		},
	}

	for _, testCase := range cases {
		startHost, finishHost := CalculateTotalHosts(testCase.Range)
		// fmt.Println(startHost)
		// fmt.Println(finishHost)

		if testCase.Start != startHost {
			t.Errorf("First Test Host %d does not much with Given First Host: %d", testCase.Start, startHost)
		}
		if testCase.Finish != finishHost {
			t.Errorf("First Test Host %d does not much with Given First Host: %d", testCase.Finish, finishHost)
		}
	}
}

func TestConvertIpFromBinary(t *testing.T) {
	var start uint32 = 168432640
	var finish uint32 = 168432895

	for i := start; i < finish; i++ {
		ip := ConvertIpFromBinary(i)

		fmt.Println(ip)
	}
}

func TestHostAlive(t *testing.T) {

	type Case struct {
		IP    string
		Found bool
	}

	ipList := []Case{
		{
			IP:    "10.10.20.24/24",
			Found: false,
		},
		{
			IP:    "10.10.20.1/32",
			Found: true,
		},
		{
			IP:    "10.10.99.1/24",
			Found: true,
		},
	}

	for _, host := range ipList {
		hostStatus := IsHostAlive(host.IP)
		if hostStatus == host.Found {
			fmt.Printf("Host %s Found Alive %v \n", host.IP, IsHostAlive(host.IP))
		} else {
			t.Errorf("Host %s did not much the expected status. Expected %v but got %v", host.IP, host.Found, hostStatus)
		}

	}

}

func TestNetworkScan(t *testing.T) {

	type Case struct {
		IP            string
		DiscoveredIPs []string
	}

	ipList := []Case{
		{
			IP: "10.10.70.1/24",
			DiscoveredIPs: []string{
				"10.10.70.1",
			},
		},
	}

	for _, discovery := range ipList {
		networkDiscoveries := NetworkScan(discovery.IP)

		if len(networkDiscoveries) != len(discovery.DiscoveredIPs) {
			t.Errorf("Array length does not much")
		}

		fmt.Printf("Host Found Alive %v \n", networkDiscoveries)
	}

}
