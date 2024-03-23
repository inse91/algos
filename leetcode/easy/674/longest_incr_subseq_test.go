package _674__test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			res:  3,
		},
		{
			name: "2",
			nums: []int{0, 1, 0, 3, 2, 3},
			res:  2,
		},
		{
			name: "3",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			res:  1,
		},
		{
			name: "4",
			nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6, 2},
			res:  5,
		},
		{
			name: "5",
			nums: []int{1, 3, 5, 4, 7},
			res:  3,
		},
		{
			name: "6",
			nums: []int{1, 3, 5, 7},
			res:  4,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findLengthOfLCIS(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func findLengthOfLCIS(nums []int) (res int) {
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur++
			continue
		}

		res = max(res, cur)
		cur = 1
	}
	return max(res, cur)
}
