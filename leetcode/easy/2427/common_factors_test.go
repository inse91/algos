package _2427_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-common-factors

func Test(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		res  int
	}{
		{
			name: "1",
			a:    12,
			b:    6,
			res:  4,
		},
		{
			name: "2",
			a:    25,
			b:    30,
			res:  2,
		},
		{
			name: "3",
			a:    885,
			b:    885,
			res:  8,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, commonFactors(c.a, c.b))
		})
	}
}

func commonFactors(a int, b int) int {
	top := max(a, b)
	if a != b {
		top /= 2
	}

	var res int
	for i := 1; i <= top; i++ {
		if a%i == 0 && b%i == 0 {
			res++
		}
	}

	return res
}
