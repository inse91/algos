package _1839_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/description/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		ranges [][]int
		l, r   int
		res    bool
	}{
		{
			name:   "1",
			ranges: [][]int{{1, 2}, {3, 4}, {5, 6}},
			l:      2,
			r:      5,
			res:    true,
		},
		{
			name:   "2",
			ranges: [][]int{{1, 10}, {10, 20}},
			l:      21,
			r:      21,
			res:    false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isCovered(c.ranges, c.l, c.r))
		})
	}
}

func isCovered(ranges [][]int, left int, right int) bool {
	m := make(map[int]struct{}, 50)
	for _, r := range ranges {
		if len(r) <= 1 {
			panic("invalid ranges length")
		}

		for i := r[0]; i <= r[1]; i++ {
			m[i] = struct{}{}
		}
	}

	for i := left; i <= right; i++ {
		if _, ok := m[i]; !ok {
			return false
		}
	}

	return true
}
