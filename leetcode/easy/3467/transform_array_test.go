package _3467_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-elements-with-maximum-frequency

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{4, 3, 2, 1},
			res:  []int{0, 0, 1, 1},
		},
		{
			name: "2",
			nums: []int{1, 5, 1, 4, 2},
			res:  []int{0, 0, 1, 1, 1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, transformArray(c.nums))
		})
	}
}

func transformArray(nums []int) []int {
	var eventCounter int
	for _, num := range nums {
		if num%2 == 0 {
			eventCounter++
		}
	}

	for i := range nums {
		if i < eventCounter {
			nums[i] = 0
			continue
		}
		nums[i] = 1
	}

	return nums
}
