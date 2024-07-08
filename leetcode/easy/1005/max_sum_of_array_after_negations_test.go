package _1005_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

// https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{
			name: "1",
			nums: []int{4, 2, 3},
			k:    1,
			res:  5,
		},
		{
			name: "2",
			nums: []int{3, -1, 0, 2},
			k:    3,
			res:  6,
		},
		{
			name: "2_1",
			nums: []int{3, -1, 1, 2},
			k:    3,
			res:  7,
		},
		{
			name: "3",
			nums: []int{2, -3, -1, 5, -4},
			k:    2,
			res:  13,
		},
		{
			name: "67",
			nums: []int{-2, 5, 0, 2, -2},
			k:    3,
			res:  11,
		},
		{
			name: "78",
			nums: []int{-8, 3, -5, -3, -5, -2},
			k:    6,
			res:  22,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, largestSumAfterKNegations(c.nums, c.k))
		})
	}
}

func largestSumAfterKNegations(nums []int, k int) int {
	negs := make([]int, 0, len(nums)/2)
	poses := make([]int, 0, len(nums)/2)
	var posSum int

	for _, num := range nums {
		if num < 0 {
			negs = append(negs, num)
		}
		if num >= 0 {
			poses = append(poses, num)
			posSum += num
		}
	}

	slices.Sort(negs)
	var negSum int
	for _, num := range negs {
		if k == 0 {
			negSum += num
			continue
		}
		poses = append(poses, -num)
		posSum += -num
		k--
	}

	if k == 0 || k%2 == 0 || len(poses) == 0 {
		return negSum + posSum
	}

	// find min elem
	minPosElem := poses[0]
	for _, posElem := range poses {
		minPosElem = min(minPosElem, posElem)
	}
	posSum -= minPosElem * 2

	return posSum + negSum
}
