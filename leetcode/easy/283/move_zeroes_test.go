package _283_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{0, 1, 0, 3, 12},
			res:  []int{1, 3, 12, 0, 0},
		},
		{
			name: "2",
			nums: []int{0},
			res:  []int{0},
		},
		{
			name: "3",
			nums: []int{1, 2, 4},
			res:  []int{1, 2, 4},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			moveZeroes(c.nums)
			assert.Equal(t, c.res, c.nums)
		})
	}
}

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	var cur int
	for _, num := range nums {
		if num == 0 {
			continue
		}
		nums[cur] = num
		cur++
	}
	for cur < len(nums) {
		nums[cur] = 0
		cur++
	}
}
