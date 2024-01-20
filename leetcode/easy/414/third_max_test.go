package _414_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestThirdMax(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		res  int
	}{
		{
			name: "1",
			in:   []int{3, 2, 1},
			res:  1,
		},
		{
			name: "2",
			in:   []int{1, 2},
			res:  2,
		},
		{
			name: "3",
			in:   []int{2, 2, 3, 1},
			res:  1,
		},
		{
			name: "4",
			in:   []int{1, 2, 2, 5, 3, 5},
			res:  2,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := thirdMax(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func thirdMax(nums []int) int {
	slices.Sort(nums)
	mx := nums[len(nums)-1]
	var count int
	for i := len(nums) - 2; i >= 0; i-- {
		if count == 2 {
			return mx
		}
		if nums[i] != mx {
			mx = nums[i]
			count++
		}
	}
	if count == 2 {
		return mx
	}
	return nums[len(nums)-1]
}
