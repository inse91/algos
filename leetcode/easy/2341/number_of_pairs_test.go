package _2341_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-number-of-pairs-in-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{1, 3, 2, 1, 3, 2, 2},
			res:  []int{3, 1},
		},
		{
			name: "2",
			nums: []int{1, 1},
			res:  []int{1, 0},
		},
		{
			name: "3",
			nums: []int{0},
			res:  []int{0, 1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, numberOfPairs(c.nums))
		})
	}
}

func numberOfPairs(nums []int) []int {
	var (
		digits     [101]int
		pairsCount int
		leftCount  int
	)
	for _, num := range nums {
		digits[num]++
	}

	for _, dig := range digits {
		if dig == 0 {
			continue
		}

		if dig%2 != 0 {
			leftCount++
		}

		pairsCount += dig / 2
	}

	return []int{pairsCount, leftCount}
}
