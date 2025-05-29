package _3099_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/harshad-number

func Test(t *testing.T) {
	cases := []struct {
		name string
		x    int
		res  int
	}{
		{
			name: "1",
			x:    18,
			res:  9,
		},
		{
			name: "2",
			x:    23,
			res:  -1,
		},
		{
			name: "13",
			x:    50,
			res:  5,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sumOfTheDigitsOfHarshadNumber(c.x))
		})
	}
}

func sumOfTheDigitsOfHarshadNumber(x int) int {
	xInit := x
	var sum int
	for x > 0 {
		sum += x % 10
		x /= 10
	}

	if xInit%sum != 0 {
		return -1
	}

	return sum
}
