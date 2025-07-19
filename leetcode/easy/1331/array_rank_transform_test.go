package _1331_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/rank-transform-of-an-array

func Test(t *testing.T) {
	cases := []struct {
		name  string
		array []int
		res   []int
	}{
		{
			name:  "1",
			array: []int{40, 10, 20, 30},
			res:   []int{4, 1, 2, 3},
		},
		{
			name:  "2",
			array: []int{10, 10, 10},
			res:   []int{1, 1, 1},
		},
		{
			name:  "3",
			array: []int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			res:   []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
		{
			name:  "4",
			array: []int{10, 20, 10},
			res:   []int{1, 2, 1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, arrayRankTransform(c.array))
		})
	}
}

func arrayRankTransform(nums []int) []int {
	cpy := make([]int, len(nums))
	copy(cpy, nums)

	slices.Sort(cpy)
	m := make(map[int]int, len(nums))
	c := 1
	for i, num := range cpy {
		if i == 0 {
			m[num] = 1
			continue
		}

		prv := cpy[i-1]
		if prv == num {
			continue
		}

		c++
		m[num] = c
	}

	for i, num := range nums {
		cpy[i] = m[num]
	}

	return cpy
}
