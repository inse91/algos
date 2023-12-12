package _102

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestHeppyNumber(t *testing.T) {
	testCases := []struct {
		name string
		in   int
		res  bool
	}{
		{
			name: "1",
			in:   19,
			res:  true,
		},
		//Explanation:
		//1^2 + 9^2 = 82
		//8^2 + 2^2 = 68
		//6^2 + 8^2 = 100
		//1^2 + 0^2 + 0^2 = 1
		{
			name: "2",
			in:   2,
			res:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isHappy(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}
}

func isHappy(n int) bool {

	s := strconv.FormatInt(int64(n), 2)
	for _, r := range s {
		ss := string(r)
		_ = ss
	}
	return false
}

func sq(n int) int {
	return n * n
}
