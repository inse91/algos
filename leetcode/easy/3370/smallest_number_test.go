package _3370_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/smallest-number-with-all-set-bits

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    5,
			res:  7,
		},
		{
			name: "2",
			n:    10,
			res:  15,
		},
		{
			name: "3",
			n:    3,
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, smallestNumber(c.n))
		})
	}
}

func smallestNumber(n int) int {
	n += 1
	for n&(n-1) != 0 {
		n += 1
	}

	return n - 1
}
