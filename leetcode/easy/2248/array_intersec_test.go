package _2248_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/intersection-of-multiple-arrays

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums [][]int
		res  []int
	}{
		{
			name: "1",
			nums: [][]int{
				{3, 1, 2, 4, 5},
				{1, 2, 3, 4},
				{3, 4, 5, 6},
			},
			res: []int{3, 4},
		},
		{
			name: "2",
			nums: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			res: []int{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, intersection(c.nums))
		})
	}
}

func intersection(nums [][]int) []int {
	if len(nums) == 0 {
		return nil
	}

	set := make(map[int]struct{}, len(nums[0]))
	for _, num := range nums[0] {
		set[num] = struct{}{}
	}

	for _, numz := range nums[1:] {
		compSet := make(map[int]struct{}, len(numz))
		for _, num := range numz {
			compSet[num] = struct{}{}
		}

		for num := range set {
			_, ok := compSet[num]
			if ok {
				continue
			}

			delete(set, num)
		}
	}

	res := make([]int, 0, len(set))
	for num := range set {
		res = append(res, num)
	}

	return res
}
