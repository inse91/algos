package _3396_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 4, 2, 3, 3, 5, 7},
			res:  2,
		},
		{
			name: "2",
			nums: []int{4, 5, 6, 4, 4},
			res:  2,
		},
		{
			name: "3",
			nums: []int{6, 7, 8, 9},
			res:  0,
		},
		{
			name: "23",
			nums: []int{1, 2, 7, 15, 23, 3, 4, 2, 13, 13, 15, 7, 23},
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minimumOperations(c.nums))
		})
	}
}

func minimumOperations(nums []int) int {
	set := make(map[int]int)
	var nonDistincts int
	for _, num := range nums {
		if c, ok := set[num]; ok && c == 1 {
			nonDistincts++
		}
		set[num]++
	}

	if nonDistincts == 0 {
		return 0
	}

	var (
		res       = 1
		firstPart []int
	)
	for {
		if len(nums) <= 3 {
			return res
		}

		firstPart = nums[:3]
		for _, num := range firstPart {
			if set[num] <= 1 {
				continue
			}

			set[num]--
			if set[num] == 1 {
				nonDistincts--
			}

			if nonDistincts == 0 {
				return res
			}
		}

		nums = nums[3:]
		res++
	}
}
