package _2798_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-employees-who-met-the-target

func Test(t *testing.T) {
	cases := []struct {
		name   string
		hours  []int
		target int
		res    int
	}{
		{
			name:   "1",
			hours:  []int{0, 1, 2, 3, 4},
			target: 2,
			res:    3,
		},
		{
			name:   "2",
			hours:  []int{5, 1, 4, 2, 2},
			target: 6,
			res:    0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numberOfEmployeesWhoMetTarget(c.hours, c.target))
		})
	}
}

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var res int
	for _, h := range hours {
		var inc int
		if h >= target {
			inc = 1
		}

		res += inc
	}

	return res
}
