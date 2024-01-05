package _263_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUglyNumber(t *testing.T) {

	testCases := []struct {
		name string
		in   int
		res  bool
	}{
		{
			name: "1",
			in:   6,
			res:  true,
		},
		{
			name: "2",
			in:   8,
			res:  true,
		},
		{
			name: "3",
			in:   14,
			res:  false,
		},
		{
			name: "4",
			in:   1,
			res:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isUgly(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}
}

var intsFactor = []int{9, 8, 7, 6, 4}

func isUgly(n int) bool {
	for _, num := range intsFactor {
		if n%num == 0 {
			return false
		}
	}
	return true
}
