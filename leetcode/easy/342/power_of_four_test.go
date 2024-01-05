package _342_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowerOfFour(t *testing.T) {
	cases := []struct {
		name string
		in   int
		res  bool
	}{
		{
			name: "1",
			in:   16,
			res:  true,
		},
		{
			name: "2",
			in:   17,
			res:  false,
		},
		{
			name: "3",
			in:   -1,
			res:  false,
		},
		{
			name: "4",
			in:   1,
			res:  true,
		},
		{
			name: "5",
			in:   4,
			res:  true,
		},
		{
			name: "6",
			in:   5,
			res:  false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := isPowerOfFour(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	for n >= 4 {
		if n%4 != 0 {
			return false
		}
		n /= 4
	}
	return n == 1
}
