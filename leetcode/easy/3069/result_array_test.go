package _3069_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/distribute-elements-into-two-arrays-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{2, 1, 3},
			res:  []int{2, 3, 1},
		},
		{
			name: "2",
			nums: []int{5, 4, 3, 8},
			res:  []int{5, 3, 4, 8},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, resultArray(c.nums))
		})
	}
}

func resultArray(nums []int) []int {
	res1 := make([]int, 0, len(nums))
	res2 := make([]int, 0, len(nums)/2)
	res1 = append(res1, nums[0])
	res2 = append(res2, nums[1])

	for _, n := range nums[2:] {
		last1 := res1[len(res1)-1]
		last2 := res2[len(res2)-1]
		if last1 > last2 {
			res1 = append(res1, n)
			continue
		}

		res2 = append(res2, n)
	}

	return append(res1, res2...)
}
