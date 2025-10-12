package _3502_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-cost-to-reach-every-position

func Test(t *testing.T) {
	cases := []struct {
		name string
		cost []int
		res  []int
	}{
		{
			name: "1",
			cost: []int{5, 3, 4, 1, 3, 2},
			res:  []int{5, 3, 3, 1, 1, 1},
		},
		{
			name: "2",
			cost: []int{1, 2, 4, 6, 7},
			res:  []int{1, 1, 1, 1, 1},
		},
		{
			name: "3",
			cost: []int{5, 3, 5, 2, 2, 7, 10, 11, 4, 2, 4, 1, 3, 2},
			res:  []int{5, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minCosts(c.cost))
		})
	}
}

func minCosts(cost []int) []int {
	prevMax := cost[0]
	for i, v := range cost {
		if v < prevMax {
			prevMax = v
			continue
		}

		cost[i] = prevMax
	}

	return cost
}
