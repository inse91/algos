package _3074_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/apple-redistribution-into-boxes

func Test(t *testing.T) {
	cases := []struct {
		name  string
		apple []int
		cap   []int
		res   int
	}{
		{
			name:  "1",
			apple: []int{1, 3, 2},
			cap:   []int{4, 3, 1, 5, 2},
			res:   2,
		},
		{
			name:  "2",
			apple: []int{5, 5, 5},
			cap:   []int{2, 4, 2, 7},
			res:   4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minimumBoxes(c.apple, c.cap))
		})
	}
}

func minimumBoxes(apple []int, capacity []int) int {
	var applesTotal int
	for _, apl := range apple {
		applesTotal += apl
	}

	slices.SortFunc(capacity, func(a, b int) int {
		if a > b {
			return -1
		}

		return 1
	})

	var c int
	for _, cp := range capacity {
		applesTotal -= cp
		c++
		if applesTotal <= 0 {
			return c
		}
	}

	return len(capacity)
}
