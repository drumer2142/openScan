package src

import "testing"

func TestCalculateSubnetRange(t *testing.T) {

	type Case struct {
		Range      string
		TotalHosts uint64
	}

	cases := []Case{
		{
			Range:      "10.10.20.0/24",
			TotalHosts: 256,
		},
		{
			Range:      "10.10.20.0/32",
			TotalHosts: 1,
		},
	}

	for _, testCase := range cases {
		totalHosts := CalculateTotalHosts(testCase.Range)

		if testCase.TotalHosts != totalHosts {
			t.Errorf("Last %d does not much with result: %d", testCase.TotalHosts, totalHosts)
		}
	}
}
