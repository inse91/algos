package _268_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{3, 0, 1},
			res:  2,
		},
		{
			name: "2",
			nums: []int{0, 1},
			res:  2,
		},
		{
			name: "3",
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			res:  8,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := missingNumber(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)
	prev := nums[0]
	if prev != 0 {
		return 0
	}
	for _, num := range nums[1:] {
		if num != prev+1 {
			return prev + 1
		}
		prev = num
	}
	return prev + 1
}
