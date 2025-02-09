package _1732_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		gain []int
		res  int
	}{
		{
			name: "1",
			gain: []int{-5, 1, 5, 0, -7},
			res:  1,
		},
		{
			name: "2",
			gain: []int{-4, -3, -2, -1, 4, 3, 2},
			res:  0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, largestAltitude(c.gain))
		})
	}
}

func largestAltitude(gain []int) int {
	var res, cur int
	for _, gap := range gain {
		cur += gap
		res = max(res, cur)
	}

	return res
}
