package _2190_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/neither-minimum-nor-maximum

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		key  int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 100, 200, 1, 100},
			key:  1,
			res:  100,
		},
		{
			name: "2",
			nums: []int{2, 2, 2, 2, 3},
			key:  2,
			res:  2,
		},
		{
			name: "97",
			nums: []int{1, 1, 1, 1, 2, 2, 2},
			key:  1,
			res:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := mostFrequent(c.nums, c.key)
			assert.Equal(t, c.res, res)
		})
	}
}

func mostFrequent(nums []int, key int) int {
	var (
		freq   [1001]int
		mxFreq int
		res    int
	)
	for i, num := range nums {
		if num != key || i == len(nums)-1 {
			continue
		}

		next := nums[i+1]
		freq[next]++
		if freq[next] <= mxFreq {
			continue
		}

		mxFreq = freq[next]
		res = next
	}

	return res
}
