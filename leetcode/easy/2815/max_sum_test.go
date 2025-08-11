package _2815_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/max-pair-sum-in-an-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{112, 131, 411},
			res:  -1,
		},
		{
			name: "2",
			nums: []int{2536, 1613, 3366, 162},
			res:  5902,
		},
		{
			name: "3",
			nums: []int{51, 71, 17, 24, 42},
			res:  88,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxSum(c.nums))
		})
	}
}

func maxSum(nums []int) int {
	m := make(map[int][]int)
	for _, num := range nums {
		d := findMaxDigit(num)
		m[d] = append(m[d], num)
	}

	res := -1
	for i := 9; i >= 0; i-- {
		slc := m[i]
		if len(slc) <= 1 {
			continue
		}

		slices.Sort(slc)
		lastIdx := len(slc) - 1
		res = max(res, slc[lastIdx]+slc[lastIdx-1])
	}

	return res
}

func findMaxDigit(num int) int {
	var res int
	for num > 0 {
		dig := num - num/10*10
		res = max(res, dig)
		num /= 10
	}

	return res
}

func TestMD(t *testing.T) {
	assert.Equal(t, 3, findMaxDigit(231))
	assert.Equal(t, 9, findMaxDigit(931))
}
