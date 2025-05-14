package _3452_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/sum-of-good-numbers

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 3, 2, 1, 5, 4},
			k:    2,
			res:  12,
		},
		{
			name: "2",
			nums: []int{2, 1},
			k:    1,
			res:  2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sumOfGoodNumbers(c.nums, c.k))
		})
	}
}

func sumOfGoodNumbers(nums []int, k int) int {
	var res int
	for i, num := range nums {
		leftIdx := i - k
		rightIdx := i + k
		if leftIdx < 0 && rightIdx > len(nums)-1 {
			res += num
			continue
		}

		if leftIdx >= 0 && num <= nums[leftIdx] {
			continue
		}

		if rightIdx <= len(nums)-1 && num <= nums[rightIdx] {
			continue
		}

		res += num
	}

	return res
}
