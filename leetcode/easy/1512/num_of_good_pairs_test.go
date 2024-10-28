package _1512_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/number-of-good-pairs/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 1, 1, 3},
			res:  4,
		},
		{
			name: "2_1",
			nums: []int{1, 1, 1, 1},
			res:  6,
		},
		{
			name: "2_2",
			nums: []int{1, 1, 1, 1, 1},
			res:  10,
		},
		{
			name: "2_3",
			nums: []int{1, 1, 1, 1, 1, 1},
			res:  15,
		},
		{
			name: "3",
			nums: []int{1, 2, 3},
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numIdenticalPairs(c.nums))
		})
	}
}

func numIdenticalPairs(nums []int) int {
	m := make(map[int]int, len(nums)/2)
	for _, num := range nums {
		m[num]++
	}

	var res int
	for _, v := range m {
		ps := preSum(v)
		res += ps
	}

	return res
}

func preSum(n int) int {
	var res int
	n--
	for ; n != 0; n-- {
		res += n
	}

	return res
}
