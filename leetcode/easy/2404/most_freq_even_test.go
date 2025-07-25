package _2404_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/most-frequent-even-element

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{0, 1, 2, 2, 4, 4, 1},
			res:  2,
		},
		{
			name: "1_1",
			nums: []int{0, 1, 2, 4, 4, 2, 1},
			res:  2,
		},
		{
			name: "2",
			nums: []int{4, 4, 4, 9, 2, 4},
			res:  4,
		},
		{
			name: "3",
			nums: []int{29, 47, 21, 41, 13, 37, 25, 7},
			res:  -1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, mostFrequentEven(c.nums))
		})
	}
}

func mostFrequentEven(nums []int) int {
	var (
		maxQ  int
		maxEl int
	)
	m := make(map[int]int, len(nums)/2)
	for _, n := range nums {
		if n%2 != 0 {
			continue
		}

		m[n]++
		if m[n] == maxQ {
			maxEl = min(maxEl, n)
		}
		if m[n] > maxQ {
			maxQ = m[n]
			maxEl = n
		}
	}

	if len(m) == 0 {
		return -1
	}

	return maxEl
}
