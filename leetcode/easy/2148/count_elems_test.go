package _2148_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-elements-with-strictly-smaller-and-greater-elements

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{11, 7, 2, 15},
			res:  2,
		},
		{
			name: "11",
			nums: []int{7, 7, 7, 15},
			res:  0,
		},
		{
			name: "2",
			nums: []int{-3, 3, 3, 90},
			res:  2,
		},
		{
			name: "27",
			nums: []int{-71, -71, 93, -71, 40},
			res:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := countElements(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func countElements(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}

	slices.Sort(nums)
	var (
		ln                          = len(nums)
		first                       = nums[0]
		last                        = nums[ln-1]
		foundLeft, foundRight       bool
		foundLeftIdx, foundRightIdx int
	)

	if first == last {
		return 0
	}

	for i := 1; i < ln-1; i++ {
		if foundLeft && foundRight {
			break
		}

		curLeft := nums[i]
		if !foundLeft && curLeft != first {
			foundLeftIdx = i
			foundLeft = true
		}

		curRight := nums[ln-i-1]
		if !foundRight && curRight != last {
			foundRightIdx = ln - i
			foundRight = true
		}
	}

	if !foundLeft || !foundRight {
		return 0
	}

	return foundRightIdx - foundLeftIdx
}
