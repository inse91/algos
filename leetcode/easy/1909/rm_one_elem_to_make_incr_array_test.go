package _1909_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		res  bool
	}{
		{
			name: "1",
			arr:  []int{1, 2, 5, 7},
			res:  true,
		},
		{
			name: "2",
			arr:  []int{2, 3, 1, 2},
			res:  false,
		},
		{
			name: "3",
			arr:  []int{1, 1, 1},
			res:  false,
		},
		{
			name: "13",
			arr:  []int{1, 2},
			res:  true,
		},
		{
			name: "15",
			arr:  []int{1, 1},
			res:  true,
		},
		{
			name: "17",
			arr:  []int{3, 1, 2},
			res:  true,
		},
		{
			name: "19",
			arr:  []int{1, 3, 1, 2},
			res:  false,
		},
		{
			name: "83",
			arr:  []int{1, 3, 2, 4},
			res:  true,
		},
		{
			name: "84",
			arr:  []int{2, 3, 1, 4},
			res:  true,
		},
		{
			name: "84_1",
			arr:  []int{2, 3, 3, 4, 5},
			res:  true,
		},
		{
			name: "86",
			arr:  []int{105, 924, 32, 968},
			res:  true,
		},
		{
			name: "89",
			arr:  []int{2, 1, 2},
			res:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, canBeIncreasing(c.arr))
		})
	}
}

func canBeIncreasing(nums []int) bool {
	var oneCheck bool
	for i := 1; i < len(nums); i++ {
		prv, cur := nums[i-1], nums[i]
		if cur > prv {
			continue
		}

		if oneCheck { // один элемент уже убрали
			return false
		}

		if i == 1 || i == len(nums)-1 {
			oneCheck = true

			continue
		}

		if i == len(nums)-1 {
			return true
		}

		// нужно попробовать убрать либо текущий либо предыдущий

		// 1) пробуем убрать предыдущий
		prvPrv := nums[i-2]
		if cur > prvPrv {
			oneCheck = true

			continue
		}

		// 2) пробуем текущий
		next := nums[i+1]
		if next <= prv {
			return false
		}

		oneCheck = true
	}

	return true
}
