package _2769_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-maximum-achievable-number

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		t    int
		res  int
	}{
		{
			name: "1",
			num:  4,
			t:    1,
			res:  6,
		},
		{
			name: "2",
			num:  3,
			t:    2,
			res:  7,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, theMaximumAchievableX(c.num, c.t))
		})
	}
}

func theMaximumAchievableX(num int, t int) int {
	return num + t*2
}
