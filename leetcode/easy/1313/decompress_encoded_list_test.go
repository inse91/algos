package _1313_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/decompress-run-length-encoded-list/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 4},
			res:  []int{2, 4, 4, 4},
		},
		{
			name: "2",
			nums: []int{1, 1, 2, 3},
			res:  []int{1, 3, 3},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, decompressRLElist(c.nums))
		})
	}
}

func decompressRLElist(nums []int) []int {
	var l int
	for i := 0; i < len(nums); i += 2 {
		l += nums[i]
	}

	var num, q int
	res := make([]int, 0, l)

	for i := 1; i < len(nums); i += 2 {
		num, q = nums[i], nums[i-1]
		for j := 0; j < q; j++ {
			res = append(res, num)
		}
	}

	return res
}
