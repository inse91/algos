package _1200_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-absolute-difference/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  [][]int
	}{
		{
			name: "1",
			nums: []int{4, 2, 1, 3},
			res:  [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name: "2",
			nums: []int{1, 3, 6, 10, 15},
			res:  [][]int{{1, 3}},
		},
		{
			name: "3",
			nums: []int{3, 8, -10, 23, 19, -4, -14, 27},
			res:  [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minimumAbsDifference(c.nums))
			for i, got := range minimumAbsDifference(c.nums) {
				assert.ElementsMatch(t, c.res[i], got)
			}
		})
	}
}

func minimumAbsDifference(arr []int) [][]int {
	slices.Sort(arr)

	minDiff := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		diff := abs(arr[i] - arr[i-1])
		minDiff = min(minDiff, diff)
	}

	var res [][]int
	for i := 1; i < len(arr); i++ {
		diff := abs(arr[i] - arr[i-1])
		if diff != minDiff {
			continue
		}

		res = append(res, []int{arr[i-1], arr[i]})
	}

	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
