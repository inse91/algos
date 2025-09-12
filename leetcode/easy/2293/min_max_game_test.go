package _2293_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/min-max-game

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 3, 5, 2, 4, 8, 2, 2},
			res:  1,
		},
		{
			name: "2",
			nums: []int{3},
			res:  3,
		},
		{
			name: "3",
			nums: []int{12132, 33, 15, 442, 412, 138, 2312, 1422},
			res:  33,
		},
		{
			name: "79",
			nums: []int{70, 38, 21, 22},
			res:  22,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minMaxGame(c.nums))
		})
	}
}

func minMaxGame(nums []int) int {
	var mx bool
	for {
		if len(nums) == 1 {
			return nums[0]
		}

		for i := 0; i < len(nums); i += 2 {
			n1 := nums[i]
			n2 := nums[i+1]

			if mx {
				nums[i/2] = max(n1, n2)
			} else {
				nums[i/2] = min(n1, n2)
			}

			mx = !mx
		}
		nums = nums[:len(nums)/2]
	}
}
