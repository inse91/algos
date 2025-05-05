package _2073_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/time-needed-to-buy-tickets/

func Test(t *testing.T) {
	cases := []struct {
		name    string
		tickets []int
		k       int
		res     int
	}{
		{
			name:    "1",
			tickets: []int{2, 3, 2},
			k:       2,
			res:     6,
		},
		{
			name:    "2",
			tickets: []int{5, 1, 1, 1},
			k:       0,
			res:     8,
		},
		{
			name:    "9",
			tickets: []int{84, 49, 5, 24, 70, 77, 87, 8},
			k:       3,
			res:     154,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, timeRequiredToBuy(c.tickets, c.k))
		})
	}
}

func timeRequiredToBuy(tickets []int, k int) int {
	goal := tickets[k]
	var res int
	for i, ticket := range tickets {
		if ticket < goal {
			res += ticket
			continue
		}

		res += goal
		if i > k {
			res -= 1
		}
	}

	return res
}
