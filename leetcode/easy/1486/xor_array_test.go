package _1486_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/xor-operation-in-an-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		start int
		n     int
		res   int
	}{
		{
			name:  "1",
			start: 0,
			n:     5,
			res:   8,
		},
		{
			name:  "2",
			start: 3,
			n:     4,
			res:   8,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, xorOperation(c.n, c.start))
		})
	}
}

func xorOperation(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res ^= start + 2*i
	}

	return res
}
