package _1365_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{8, 1, 2, 2, 3},
			res:  []int{4, 0, 1, 1, 3},
		},
		{
			name: "2",
			nums: []int{6, 5, 4, 8},
			res:  []int{2, 1, 0, 3},
		},
		{
			name: "3",
			nums: []int{7, 7, 7, 7},
			res:  []int{0, 0, 0, 0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, smallerNumbersThanCurrent(c.nums))
		})
	}
}

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, 0, len(nums))
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for _, num := range nums {
		var c int
		for k, v := range m {
			if k >= num {
				continue
			}

			c += v
		}

		res = append(res, c)
	}

	return res
}
