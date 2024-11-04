package _1470_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/shuffle-the-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		n    int
		res  []int
	}{
		{
			name: "1",
			nums: []int{2, 5, 1, 3, 4, 7},
			n:    3,
			res:  []int{2, 3, 5, 4, 1, 7},
		},
		{
			name: "2",
			nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
			n:    4,
			res:  []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name: "3",
			nums: []int{1, 1, 2, 2},
			n:    2,
			res:  []int{1, 2, 1, 2},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, shuffle(c.nums, c.n))
		})
	}

}

func shuffle(nums []int, n int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		res = append(res, nums[i], nums[i+n])
	}

	return res
}
