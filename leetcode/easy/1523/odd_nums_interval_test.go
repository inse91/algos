package _1523_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		l    int
		h    int
		res  int
	}{
		{
			name: "1",
			l:    3,
			h:    7,
			res:  3,
		},
		{
			name: "2",
			l:    8,
			h:    10,
			res:  1,
		},
		{
			name: "#1",
			l:    3213142,
			h:    45656565,
			res:  21221712,
		},
		{
			name: "#2",
			l:    0,
			h:    17,
			res:  9,
		},
		{
			name: "#3",
			l:    4,
			h:    4,
			res:  0,
		},
		{
			name: "#4",
			l:    5,
			h:    5,
			res:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countOdds(c.l, c.h))
		})
	}
}

func countOdds(low int, high int) int {
	if low == high {
		if low%2 == 0 {
			return 0
		}

		return 1
	}

	if high%2 == 0 {
		high--
	}

	return (high-low)/2 + 1
}
