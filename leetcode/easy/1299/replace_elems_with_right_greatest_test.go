package _1299_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{17, 18, 5, 4, 6, 1},
			res:  []int{18, 6, 6, 6, 1, -1},
		},
		{
			name: "2",
			nums: []int{17},
			res:  []int{-1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, replaceElements(c.nums))
		})
	}

}

func replaceElements(arr []int) []int {
	res := make([]int, len(arr))

	curMax := findMax(arr)
	for i := range arr {
		if arr[i] == curMax {
			curMax = findMax(arr[i+1:])
		}

		res[i] = curMax
	}

	return res
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	return slices.Max(arr)
}
