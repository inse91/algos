package _1550_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/three-consecutive-odds/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		res  bool
	}{
		{
			name: "1",
			arr:  []int{2, 6, 4, 1},
			res:  false,
		},
		{
			name: "2",
			arr:  []int{1, 2, 34, 3, 4, 5, 7, 23, 12},
			res:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, threeConsecutiveOdds(c.arr))
		})
	}
}

func threeConsecutiveOdds(arr []int) bool {
	var c int
	for _, num := range arr {
		if num%2 == 0 {
			c = 0
			continue
		}

		c++
		if c == 3 {
			return true
		}
	}

	return false
}
