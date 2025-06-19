package _2164_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sort-even-and-odd-indices-independently/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{4, 1, 2, 3},
			res:  []int{2, 3, 4, 1},
		},
		{
			name: "2",
			nums: []int{2, 1},
			res:  []int{2, 1},
		},
		{
			name: "3",
			nums: []int{5, 12, 23, 46, 99, 33, 34},
			res:  []int{5, 46, 23, 33, 34, 12, 99},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sortEvenOdd(c.nums))
		})
	}
}

func sortEvenOdd(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}

	odds := make([]int, 0, len(nums)/2)
	evens := make([]int, 0, len(nums)/2)
	for i := range nums {
		if i%2 == 0 {
			evens = append(evens, nums[i])
			continue
		}

		odds = append(odds, nums[i])
	}

	slices.Sort(evens)
	slices.SortFunc(odds, func(a, b int) int {
		return b - a
	})

	var o, e int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = evens[e]
			e++
			continue
		}

		nums[i] = odds[o]
		o++
	}

	return nums
}
