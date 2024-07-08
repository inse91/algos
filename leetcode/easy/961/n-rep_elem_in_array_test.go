package _961_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 3},
			res:  3,
		},
		{
			name: "2",
			nums: []int{2, 1, 2, 5, 3, 2},
			res:  2,
		},
		{
			name: "3",
			nums: []int{5, 1, 5, 2, 5, 3, 5, 4},
			res:  5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, repeatedNTimes(c.nums))
		})
	}
}

func repeatedNTimes(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		_, ok := set[num]
		if ok {
			return num
		}

		set[num] = struct{}{}
	}

	return -1
}
