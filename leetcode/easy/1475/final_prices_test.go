package _1475_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop

func Test(t *testing.T) {
	cases := []struct {
		name   string
		prices []int
		res    []int
	}{
		{
			name:   "1",
			prices: []int{8, 4, 6, 2, 3},
			res:    []int{4, 2, 4, 2, 3},
		},
		{
			name:   "1_1",
			prices: []int{8, 4, 6, 5, 2, 3},
			res:    []int{4, 2, 1, 3, 2, 3},
		},
		{
			name:   "2",
			prices: []int{1, 2, 3, 4, 5},
			res:    []int{1, 2, 3, 4, 5},
		},
		{
			name:   "3",
			prices: []int{10, 1, 1, 6},
			res:    []int{9, 0, 1, 6},
		},
		{
			name:   "58",
			prices: []int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1},
			res:    []int{1, 3, 2, 1, 7, 0, 0, 6, 9, 1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, finalPrices(c.prices))
		})
	}
}

func finalPrices(prices []int) []int {
	for i := 0; i <= len(prices)-2; i++ {
		current := prices[i]

		var c int
		for j := i + 1; j <= len(prices)-1; j++ {
			next := prices[j]
			if next > current {
				continue
			}

			c = next
			break
		}

		prices[i] = current - c
	}

	return prices
}
