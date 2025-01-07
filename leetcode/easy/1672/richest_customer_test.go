package _1672_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/richest-customer-wealth/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		w    [][]int
		res  int
	}{
		{
			name: "1",
			w:    [][]int{{1, 2, 3}, {3, 2, 1}},
			res:  6,
		},
		{
			name: "2",
			w:    [][]int{{1, 5}, {7, 3}, {3, 5}},
			res:  10,
		},
		{
			name: "3",
			w:    [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			res:  17,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maximumWealth(c.w))
		})
	}
}

func maximumWealth(accounts [][]int) int {
	mx := -1
	for _, acc := range accounts {
		var summ int
		for _, n := range acc {
			summ += n
		}

		mx = max(mx, summ)
	}

	return mx
}
