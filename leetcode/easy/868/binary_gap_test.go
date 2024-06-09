package _868_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-gap/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    22,
			res:  2,
		},
		{
			name: "2",
			n:    8,
			res:  0,
		},
		{
			name: "3",
			n:    5,
			res:  2,
		},
		{
			name: "207",
			n:    6,
			res:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, binaryGap(c.n))
		})
	}
}

func binaryGap(n int) int {
	bits := make([]int, 0)
	var ones int
	for n >= 1 {
		b := n % 2
		n /= 2

		if b == 0 && len(bits) == 0 {
			continue
		}

		bits = append(bits, b)
		if b == 1 {
			ones++
		}
	}

	if ones <= 1 {
		return 0
	}

	var mx int
	var c int
	for _, bit := range bits {
		if bit == 0 {
			c++
			continue
		}

		mx = max(mx, c)
		c = 0
	}

	return mx + 1
}
