package _16_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test3SumClosest(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		res    int
	}{
		{
			name:   "1",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			res:    2,
		},
		{
			name:   "2",
			nums:   []int{1, 1, -1},
			target: 0,
			res:    1,
		},
		{
			name:   "3",
			nums:   []int{1, 1, 1, 0},
			target: -100,
			res:    2,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := threeSumClosest(c.nums, c.target)
			assert.Equal(t, c.res, res)
		})
	}
}

func threeSumClosest(nums []int, target int) int {
	st := nums[0] + nums[1] + nums[2]
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	gap := target - st
	closest := st

	r0 := nums
	for i := range r0 {
		n1 := r0[i]
		r1 := r0[i+1:]

		for j := range r1 {
			n2 := r1[j]
			r2 := r1[j+1:]

			for k := range r2 {
				n3 := r2[k]

				sum := n1 + n2 + n3
				if sum == target {
					return target
				}

				if abs(sum-target) < abs(gap) {
					gap = sum - target
				}
				if abs(sum-target) < abs(closest-target) {
					closest = sum
				}
			}
		}

	}

	//return target - gap
	return closest
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
