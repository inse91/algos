package _2119_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/a-number-after-a-double-reversal/

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  bool
	}{
		{
			name: "1",
			num:  526,
			res:  true,
		},
		{
			name: "2",
			num:  1800,
			res:  false,
		},
		{
			name: "3",
			num:  0,
			res:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isSameAfterReversals(c.num))
		})
	}
}

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}

	return num%10 != 0
}
