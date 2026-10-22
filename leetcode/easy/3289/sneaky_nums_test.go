package _3289_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{0, 1, 1, 0},
			res:  []int{0, 1},
		},
		{
			name: "2",
			nums: []int{0, 3, 2, 1, 3, 2},
			res:  []int{2, 3},
		},
		{
			name: "3",
			nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2},
			res:  []int{4, 5},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, getSneakyNumbers(c.nums))
		})
	}
}

func getSneakyNumbers(nums []int) []int {
	m := make(map[int]struct{}, len(nums))
	var res []int

	for _, num := range nums {
		_, ok := m[num]
		if !ok {
			m[num] = struct{}{}
			continue
		}

		res = append(res, num)
		if len(res) == 2 {
			break
		}
	}

	return res
}
