package _2535_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 15, 6, 3},
			res:  9,
		},
		{
			name: "2",
			nums: []int{1, 2, 3, 4},
			res:  0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, differenceOfSum(c.nums))
		})
	}
}

func differenceOfSum(nums []int) int {
	var sum int
	for i, num := range nums {
		if num < 10 {
			continue
		}

		var digSum int
		for num > 0 {
			digSum += num % 10
			num /= 10
		}

		diff := digSum - nums[i]
		if diff < 0 {
			diff = -diff
		}

		sum += diff
	}

	return sum
}
