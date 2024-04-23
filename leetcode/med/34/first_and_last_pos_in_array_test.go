package _34_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFirstAndLastPositionInArray(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		res    []int
	}{
		{
			name:   "1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			res:    []int{3, 4},
		},
		{
			name:   "2",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			res:    []int{-1, -1},
		},
		{
			name:   "3",
			nums:   []int{},
			target: 0,
			res:    []int{-1, -1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res[0], searchRange(c.nums, c.target)[0])
			assert.Equal(t, c.res[1], searchRange(c.nums, c.target)[1])
		})
	}
}

func searchRange(nums []int, target int) []int {
	idx := search(nums, target)
	if idx == -1 {
		return []int{-1, -1}
	}

	mx := idx
	mn := idx
	for {
		mn--
		if mn < 0 {
			mn++
			break
		}
		if nums[mn] != target {
			mn++
			break
		}
	}

	for {
		mx++
		if mx > len(nums)-1 {
			mx--
			break
		}
		if nums[mx] != target {
			mx--
			break
		}
	}

	return []int{mn, mx}
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
