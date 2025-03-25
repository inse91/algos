package _1913_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{5, 6, 2, 7, 4},
			res:  34,
		},
		{
			name: "2",
			nums: []int{4, 2, 5, 9, 7, 4, 8},
			res:  64,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxProductDifference(c.nums))
		})
	}
}

func maxProductDifference(nums []int) int {
	slices.Sort(nums)
	// todo можно не сортировать
	// 	а найти два максимальных и два минимальных элемента за O(N)

	return nums[len(nums)-1]*nums[len(nums)-2] - nums[0]*nums[1]
}
