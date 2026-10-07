package _1822_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sign-of-the-product-of-an-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{-1, -2, -3, -4, 3, 2, 1},
			res:  1,
		},
		{
			name: "2",
			nums: []int{1, 5, 0, 2, -3},
			res:  0,
		},
		{
			name: "3",
			nums: []int{-1, 1, -1, 1, -1},
			res:  -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, arraySign(c.nums))
		})
	}
}

func arraySign(nums []int) int {
	var negCount int
	for _, num := range nums {
		if num == 0 {
			return 0
		}

		if num < 0 {
			negCount++
		}
	}

	if negCount%2 == 1 {
		return -1
	}

	return 1
}
