package _121

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStocks(t *testing.T) {
	testCases := []struct {
		name   string
		in     []int
		expRes int
	}{
		{
			name:   "ok",
			in:     []int{7, 1, 5, 3, 6, 4},
			expRes: 5,
		},
		{
			name:   "!ok",
			in:     []int{7, 6, 4, 3, 1},
			expRes: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := maxProfit(tc.in)
			assert.Equal(t, tc.expRes, res)
		})
	}
}

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	var profit int
	var start int = prices[0]

	var curProfit int

	for i := range prices {
		if prices[i] < start {
			start = prices[i]
			continue
		}
		curProfit = prices[i] - start
		if curProfit > profit {
			profit = curProfit
		}
	}

	return profit
}
