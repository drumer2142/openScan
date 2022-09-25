package src

import "testing"

func TestCalculateSubnetRange(t *testing.T) {

	type Case struct {
		Range  string
		LastIp uint64
	}

	cases := []Case{
		{
			Range:  "10.10.20.0/24",
			LastIp: 256,
		},
		{
			Range:  "10.10.20.0/32",
			LastIp: 1,
		},
	}

	for _, testCase := range cases {
		lastIp := CalculateSubnetRange(testCase.Range)

		if testCase.LastIp != lastIp {
			t.Errorf("Last %d does not much with result: %d", testCase.LastIp, lastIp)
		}
	}
}
