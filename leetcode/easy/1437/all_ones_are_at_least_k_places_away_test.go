package _1437_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  bool
	}{
		{
			name: "1",
			nums: []int{1, 0, 0, 0, 1, 0, 0, 1},
			k:    2,
			res:  true,
		},
		{
			name: "2",
			nums: []int{1, 0, 0, 1, 0, 1},
			k:    2,
			res:  false,
		},
		{
			name: "3",
			nums: []int{0, 0, 0, 0, 1},
			k:    2,
			res:  true,
		},
		{
			name: "4_1",
			nums: []int{0, 1},
			k:    2,
			res:  true,
		},
		{
			name: "4_2",
			nums: []int{1, 1},
			k:    2,
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, kLengthApart(c.nums, c.k))
		})
	}
}

func kLengthApart(nums []int, k int) bool {
	var gotOne bool
	var dist int
	for _, num := range nums {
		switch num {
		case 0:
			if gotOne {
				dist++
			}
			continue
		case 1:
			if !gotOne {
				gotOne = true
				dist = 0
				continue
			}

			if dist < k {
				return false
			}

			dist = 0
		}
	}

	return true
}
