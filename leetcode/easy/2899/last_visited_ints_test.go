package _2899_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/last-visited-integers

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{1, 2, -1, -1, -1},
			res:  []int{2, 1, -1},
		},
		{
			name: "2",
			nums: []int{1, -1, 2, -1, -1},
			res:  []int{1, 2, 1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, lastVisitedIntegers(c.nums))
		})
	}
}

func lastVisitedIntegers(nums []int) []int {
	var (
		k    int
		seen []int
		res  []int
	)

	for _, num := range nums {
		if num > -1 {
			seen = append(seen, num)
			k = 0

			continue
		}

		// num == -1
		if k < len(seen) {
			invertedIdx := len(seen) - k - 1
			num = seen[invertedIdx]

			k++
		}

		res = append(res, num)
	}

	return res
}
