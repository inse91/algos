package _3354_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/make-array-elements-equal-to-zero

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 0, 2, 0, 3},
			res:  2,
		},
		{
			name: "2",
			nums: []int{2, 3, 4, 0, 4, 1, 0},
			res:  0,
		},
		{
			name: "12",
			nums: []int{16, 13, 10, 0, 0, 0, 10, 6, 7, 8, 7},
			res:  3,
		},
		{
			name: "13",
			nums: []int{0},
			res:  2,
		},
		{
			name: "14",
			nums: []int{46, 53, 45, 49, 53, 45, 47, 44, 46, 37, 50, 39, 53, 52, 48, 42, 40, 48, 43, 41, 51, 45, 41, 43, 49, 44, 45, 45, 43, 51, 51, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 35, 39, 40, 41, 34, 40, 40, 32, 27, 27, 41, 40, 39, 42, 32, 52, 32, 29, 35, 40, 32, 37, 28, 34, 42, 33, 38, 40, 40, 40, 38, 38, 30, 38, 38, 43, 36, 36, 38, 40},
			res:  22,
		},
		{
			name: "16",
			nums: []int{86, 81, 77, 80, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 18, 16, 18, 16, 17, 14, 11, 17, 12, 14, 17, 12, 16, 14, 11, 16, 15, 13, 16, 14, 9, 17, 14, 16, 10, 14},
			res:  68,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countValidSelections(c.nums))
		})
	}
}

func countValidSelections(nums []int) int {
	var (
		summ   int
		zeroes int
	)

	for i := range nums {
		summ += nums[i]
		if nums[i] == 0 {
			zeroes++
		}
	}

	if zeroes == 0 {
		return 0
	}

	var (
		left  int
		right = summ

		res int
	)

	for _, num := range nums {
		if num != 0 {
			left += num
			right -= num

			continue
		}

		switch right - left {
		case 0:
			res += 2
		case 1, -1:
			res += 1
		}
	}

	return res
}
