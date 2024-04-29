package _33_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"slices"
	"testing"
)

func TestBinSearchRotatedSortedArray(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		res    int
	}{
		{
			name:   "1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			res:    4,
		},
		{
			name:   "2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			res:    -1,
		},
		{
			name:   "3",
			nums:   []int{1},
			target: 0,
			res:    -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, search(c.nums, c.target))
		})
	}
}

func search(nums []int, target int) int {
	iters := math.Log2(float64(len(nums)))
	iters = math.Ceil(iters)

	var gap int
	var piv int

	last := nums[len(nums)-1]

	numsClone := slices.Clone(nums)

	for i := 0; i < int(iters); i++ {
		pividx := len(numsClone) / 2
		if pividx == 0 {
			break
		}
		if numsClone[pividx] < numsClone[pividx-1] {
			piv = pividx + gap
			break
		}
		if numsClone[pividx] < last {
			// left
			numsClone = numsClone[:pividx]
		} else {
			// right
			gap += len(numsClone[:pividx])
			numsClone = numsClone[pividx:]
		}
	}

	nums = append(nums[piv:], nums[:piv]...)

	idx := binSearch(nums, target)
	if idx == -1 {
		return -1
	}

	return (idx + piv) % len(nums)
}

func binSearch(nums []int, target int) int {
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
