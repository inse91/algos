package _2176_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{
			name: "1",
			nums: []int{3, 1, 2, 2, 2, 1, 3},
			k:    2,
			res:  4,
		},
		{
			name: "2",
			nums: []int{1, 2, 3, 4},
			k:    1,
			res:  0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countPairs(c.nums, c.k))
		})
	}
}

func countPairs(nums []int, k int) int {
	m := make(map[int][]int)
	for i, num := range nums {
		m[num] = append(m[num], i)
	}

	var res int
	for _, idxs := range m {
		if len(idxs) <= 1 {
			continue
		}

		for i, idx1 := range idxs {
			for _, idx2 := range idxs[i+1:] {
				if (idx1*idx2)%k != 0 {
					continue
				}

				res++
			}
		}
	}

	return res
}
