package _2652_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-multiples

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  int
	}{
		{
			name: "1",
			num:  7,
			res:  21,
		},
		{
			name: "2",
			num:  10,
			res:  40,
		},
		{
			name: "3",
			num:  9,
			res:  30,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sumOfMultiples(c.num))
		})
	}
}

func sumOfMultiples(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			res += i
		}
	}

	return res
}
