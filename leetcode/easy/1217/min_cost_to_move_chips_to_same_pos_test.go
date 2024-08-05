package _1217_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		pos  []int
		res  int
	}{
		{
			name: "1",
			pos:  []int{1, 2, 3},
			res:  1,
		},
		{
			name: "2",
			pos:  []int{2, 2, 2, 3, 3},
			res:  2,
		},
		{
			name: "3",
			pos:  []int{1, 1000000000},
			res:  1,
		},
		{
			name: "39",
			pos:  []int{6, 4, 7, 8, 2, 10, 2, 7, 9, 7},
			res:  4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minCostToMoveChips(c.pos))
		})
	}
}

func minCostToMoveChips(position []int) int {
	var even, odd int
	for _, p := range position {
		if p%2 == 0 {
			odd++
		} else {
			even++
		}
	}

	return min(odd, even)
}
