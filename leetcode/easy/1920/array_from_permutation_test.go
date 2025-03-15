package _1920_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/build-array-from-permutation/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{0, 2, 1, 5, 3, 4},
			res:  []int{0, 1, 2, 4, 5, 3},
		},
		{
			name: "2",
			nums: []int{5, 0, 1, 2, 3, 4},
			res:  []int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, buildArray(c.nums))
		})
	}
}

func buildArray(nums []int) []int {
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = nums[num]
	}

	return res
}
