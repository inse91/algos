package _231_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	for {
		if n == 1 {
			break
		}
		if n%2 == 1 {
			return false
		}
		n /= 2
	}
	return true
}

func TestIsPowerOfTwo(t *testing.T) {
	testCases := []struct {
		name string
		in   int
		res  bool
	}{
		{
			name: "1",
			in:   1,
			res:  true,
		},
		{
			name: "2",
			in:   16,
			res:  true,
		},
		{
			name: "3",
			in:   3,
			res:  false,
		},
		{
			name: "4",
			in:   17,
			res:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isPowerOfTwo(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}
}
