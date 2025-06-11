package _2144_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/faulty-keyboard

func Test(t *testing.T) {
	cases := []struct {
		name string
		cost []int
		res  int
	}{
		{
			name: "1",
			cost: []int{1, 2, 3},
			res:  5,
		},
		{
			name: "2",
			cost: []int{6, 5, 7, 9, 2, 2},
			res:  23,
		},
		{
			name: "3",
			cost: []int{5, 5},
			res:  10,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minimumCost(c.cost))
		})
	}
}

func minimumCost(cost []int) int {
	slices.SortFunc(cost, func(a, b int) int {
		return b - a
	})

	var res int
	for i, c := range cost {
		if (i+1)%3 == 0 {
			continue
		}

		res += c
	}

	return res
}
