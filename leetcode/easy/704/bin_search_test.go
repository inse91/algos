package _704_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		t    int
		res  int
	}{
		{
			name: "1",
			nums: []int{-1, 0, 3, 5, 9, 12},
			t:    9,
			res:  4,
		},
		{
			name: "2",
			nums: []int{-1, 0, 3, 5, 9, 12},
			t:    -1,
			res:  0,
		},
		{
			name: "3",
			nums: []int{-1, 0, 3, 5, 9, 12},
			t:    2,
			res:  -1,
		},
		{
			name: "4",
			nums: []int{5},
			t:    5,
			res:  0,
		},
		{
			name: "5",
			nums: []int{1, 5},
			t:    5,
			res:  1,
		},
		{
			name: "6",
			nums: []int{1, 5},
			t:    1,
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := search(c.nums, c.t)
			assert.Equal(t, c.res, res)
		})
	}
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if nums[0] == target {
		return 0
	}

	iters := math.Log2(float64(len(nums)))
	iters = math.Ceil(iters)

	var gap int
	for i := 0; i < int(iters); i++ {
		pividx := len(nums) / 2
		if nums[pividx] == target {
			return pividx + gap
		}
		if nums[pividx] > target {
			// left
			nums = nums[:pividx]
		} else {
			gap += len(nums[:pividx])
			nums = nums[pividx:]
		}
	}
	return -1
}
