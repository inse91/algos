package _2475_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-unequal-triplets-in-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{4, 4, 2, 4, 3},
			res:  3,
		},
		{
			name: "2",
			nums: []int{1, 1, 1, 1, 1},
			res:  0,
		},
		{
			name: "3",
			nums: []int{4, 4, 2, 4, 3, 2, 4},
			res:  8,
		},
		{
			name: "39",
			nums: []int{1, 3, 1, 2, 4},
			res:  7,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, unequalTriplets(c.nums))
		})
	}
}

func unequalTriplets(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	if len(m) < 3 {
		return 0
	}

	var res int
	leftCounter, rightCounter := 0, len(nums)
	for _, cnt := range m {
		rightCounter -= cnt
		res += leftCounter * cnt * rightCounter
		leftCounter += cnt
	}

	return res
}
