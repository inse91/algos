package _2903_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-indices-with-index-and-value-difference-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		id   int
		vd   int
		res  []int
	}{
		{
			name: "1",
			nums: []int{5, 1, 4, 1},
			id:   2,
			vd:   4,
			res:  []int{0, 3},
		},
		{
			name: "2",
			nums: []int{2, 1},
			id:   0,
			vd:   0,
			res:  []int{0, 0},
		},
		{
			name: "3",
			nums: []int{1, 2, 3},
			id:   2,
			vd:   4,
			res:  []int{-1, -1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findIndices(c.nums, c.id, c.vd))
		})
	}
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	for i, num1 := range nums {
		for j := i; j < len(nums); j++ {
			if abs(i-j) < indexDifference {
				continue
			}
			num2 := nums[j]
			if abs(num1-num2) < valueDifference {
				continue
			}

			return []int{i, j}
		}
	}

	return []int{-1, -1}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
