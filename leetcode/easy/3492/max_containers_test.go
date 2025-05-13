package _3492_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-containers-on-a-ship

func Test(t *testing.T) {
	cases := []struct {
		name      string
		n         int
		w         int
		maxWeight int
		res       int
	}{
		{
			name:      "1",
			n:         2,
			w:         3,
			maxWeight: 15,
			res:       4,
		},
		{
			name:      "2",
			n:         3,
			w:         5,
			maxWeight: 20,
			res:       4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxContainers(c.n, c.w, c.maxWeight))
		})
	}
}

func maxContainers(n int, w int, maxWeight int) int {
	cells := n * n
	capacity := cells * w
	if capacity <= maxWeight {
		return cells
	}

	return cells * maxWeight / capacity
}
