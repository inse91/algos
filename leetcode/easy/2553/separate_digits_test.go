package _2553_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/separate-the-digits-in-an-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{13, 25, 83, 77},
			res:  []int{1, 3, 2, 5, 8, 3, 7, 7},
		},
		{
			name: "2",
			nums: []int{1, 2, 3},
			res:  []int{1, 2, 3},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, separateDigits(c.nums))
		})
	}
}

func separateDigits(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, num := range nums {
		res = append(res, digs(num)...)
	}

	return res
}

func digs(num int) []int {
	var res []int
	for num > 0 {
		res = append(res, num%10)
		num /= 10
	}
	slices.Reverse(res)

	return res
}
