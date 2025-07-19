package _2485_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-pivot-integer

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    8,
			res:  6,
		},
		{
			name: "2",
			n:    1,
			res:  1,
		},
		{
			name: "3",
			n:    4,
			res:  -1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, pivotInteger(c.n))
		})
	}
}

func pivotInteger(n int) int {
	var fact int
	for i := 0; i <= n; i++ {
		fact += i
	}

	var acc int
	for i := n; i > 0; i-- {
		acc += i
		if fact == acc {
			return i
		}

		if fact < acc {
			return -1
		}

		fact -= i
	}

	return -1
}
