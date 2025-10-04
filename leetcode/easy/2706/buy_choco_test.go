package _2706_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/buy-two-chocolates

func Test(t *testing.T) {
	cases := []struct {
		name   string
		prices []int
		money  int
		res    int
	}{
		{
			name:   "1",
			prices: []int{1, 2, 2},
			money:  3,
			res:    0,
		},
		{
			name:   "2",
			prices: []int{3, 2, 3},
			money:  3,
			res:    3,
		},
		{
			name:   "3",
			prices: []int{1, 1, 1},
			money:  3,
			res:    1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, buyChoco(c.prices, c.money))
		})
	}
}

func buyChoco(prices []int, money int) int {
	slices.Sort(prices)

	summ := prices[0] + prices[1]
	if summ > money {
		return money
	}

	return money - summ
}
