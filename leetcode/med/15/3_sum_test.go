package _15_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test3Sum(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  [][]int
	}{
		{
			name: "1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			res:  [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "2",
			nums: []int{0, 0, 0},
			res:  [][]int{{0, 0, 0}},
		},
		{
			name: "3",
			nums: []int{0, 1, 1},
			res:  [][]int{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := threeSum(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	slices.Sort(nums)

	m := make(map[[3]int]struct{})

	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		for j, k := i+1, len(nums)-1; j < k; {
			a1 := nums[j]
			a2 := nums[k]
			if a1+a2 == target {
				j++
				k--
				if _, ok := m[[3]int{-target, a1, a2}]; ok {
					continue
				}
				res = append(res, []int{-target, a1, a2})
				m[[3]int{-target, a1, a2}] = struct{}{}
				continue
			}
			if a1+a2 < target {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
