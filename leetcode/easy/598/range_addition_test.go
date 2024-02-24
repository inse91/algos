package _598_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeAddition(t *testing.T) {
	cases := []struct {
		name string
		m    int
		n    int
		ops  [][]int
		res  int
	}{
		{
			name: "1",
			m:    3,
			n:    3,
			ops:  [][]int{{2, 2}, {3, 3}},
			res:  4,
		},
		{
			name: "2",
			m:    2,
			n:    3,
			ops:  [][]int{{0, 0, 0}, {0, 0, 0}},
			res:  3,
		},
		{
			name: "3",
			m:    3,
			n:    3,
			ops:  [][]int{},
			res:  9,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := maxCount(c.m, c.n, c.ops)
			assert.Equal(t, c.res, res)
		})
	}
}

func maxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		m = min(m, op[0])
		n = min(n, op[1])
	}

	return m * n
}
