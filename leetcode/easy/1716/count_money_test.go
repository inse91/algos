package _1716_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/calculate-money-in-leetcode-bank/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    4,
			res:  10,
		},
		{
			name: "2",
			n:    10,
			res:  37,
		},
		{
			name: "3",
			n:    20,
			res:  96,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, totalMoney(c.n))
		})
	}
}

func totalMoney(n int) int {
	var res int
	fullWeeks := n / 7
	lastWeekDays := n % 7
	start := 1 + fullWeeks
	for fullWeeks != 0 {
		res += 28 + (fullWeeks-1)*7
		fullWeeks--
	}

	for i := 0; i < lastWeekDays; i++ {
		res += start + i
	}

	return res
}
