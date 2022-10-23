package main

import "testing"

func TestCheckIPforMask(t *testing.T) {
	type Case struct {
		IP            string
		ExpectedValue string
	}

	cases := []Case{
		{
			IP:            "10.10.20.21/24",
			ExpectedValue: "10.10.20.21/24",
		},
		{
			IP:            "10.10.20.21",
			ExpectedValue: "10.10.20.21/24",
		},
		{
			IP:            "192.168.2.1/32",
			ExpectedValue: "192.168.2.1/32",
		},
	}

	for _, ipList := range cases {
		formattedIP := CheckIPforMask(ipList.IP)
		if formattedIP != ipList.ExpectedValue {
			t.Errorf("Formatted IP %s does not much the expected value %s", formattedIP, ipList.ExpectedValue)
		}
	}

}
