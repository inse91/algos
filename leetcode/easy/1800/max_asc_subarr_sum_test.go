package _1800_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-ascending-subarray-sum/description/?envType=problem-list-v2&envId=24zducyh

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{10, 20, 30, 5, 10, 50},
			res:  65,
		},
		{
			name: "2",
			nums: []int{10, 20, 30, 40, 50},
			res:  150,
		},
		{
			name: "3",
			nums: []int{12, 17, 15, 13, 10, 11, 12},
			res:  33,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxAscendingSum(c.nums))
		})
	}
}

func maxAscendingSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		cur = nums[0]
		res = 0
	)
	for i, num := range nums {
		if i == 0 {
			continue
		}

		prev := nums[i-1]
		if prev < num {
			cur += num
			continue
		}

		res = max(res, cur)
		cur = num
	}

	res = max(res, cur)

	return res
}
