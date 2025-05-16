package _3005_test

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-elements-with-maximum-frequency

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 2, 3, 1, 4},
			res:  4,
		},
		{
			name: "2",
			nums: []int{1, 2, 3, 4, 5},
			res:  5,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxFrequencyElements(c.nums))
		})
	}
}

func maxFrequencyElements(nums []int) int {
	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}

	var maxQ, res int
	for v := range maps.Values(freq) {
		if v < maxQ {
			continue
		}

		if v == maxQ {
			res += v
			continue
		}

		maxQ = v
		res = v
	}

	return res
}
