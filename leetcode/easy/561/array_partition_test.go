package _561_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPartition(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 4, 3, 2},
			res:  4,
		},
		{
			name: "2",
			nums: []int{6, 2, 6, 5, 1, 2},
			res:  9,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := arrayPairSum(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func arrayPairSum(nums []int) int {
	slices.Sort(nums)

	var sum int
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum
}
