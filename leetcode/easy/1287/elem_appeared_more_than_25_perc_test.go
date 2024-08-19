package _1287_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// вфывфывфв123123123

// https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		res  int
	}{
		{name: "1", arr: []int{1, 2, 2, 6, 6, 6, 6, 7, 10}, res: 6},
		{name: "2", arr: []int{1, 1}, res: 1},
		{name: "3", arr: []int{1}, res: 1},
		{name: "28", arr: []int{1, 1, 1, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 12, 12, 12}, res: 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findSpecialInteger(c.arr))
		})
	}
}

func findSpecialInteger(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) < 4 {
		return arr[0]
	}

	c := 1
	q := len(arr) / 4
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			c = 1
			continue
		}

		c++
		if c > q {
			return arr[i]
		}
	}

	panic("no_values")
}
