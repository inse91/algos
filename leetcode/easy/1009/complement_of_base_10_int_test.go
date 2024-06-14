package _1009_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/complement-of-base-10-integer/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    5,
			res:  2,
		},
		{
			name: "2",
			n:    7,
			res:  0,
		},
		{
			name: "3",
			n:    10,
			res:  5,
		},
	}

	t.Run("test", func(t *testing.T) {
		assert.Equal(t, 5^7, 2)
		assert.Equal(t, 7^7, 0)
		assert.Equal(t, 10^15, 5)
	})

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, bitwiseComplement(c.n))
		})
	}
}

func bitwiseComplement(n int) int {
	div := 2
	nn := n
	for nn > 1 {
		div *= 2
		nn /= 2
	}

	return n ^ (div - 1)
}
