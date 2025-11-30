package _2520_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-the-digits-that-divide-a-number

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  int
	}{
		{
			name: "1",
			num:  7,
			res:  1,
		},
		{
			name: "2",
			num:  121,
			res:  2,
		},
		{
			name: "3",
			num:  1248,
			res:  4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countDigits(c.num))
		})
	}
}

func countDigits(num int) int {
	return 0
}
