package _1748_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-unique-elements/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 2},
			res:  4,
		},
		{
			name: "2",
			nums: []int{1, 1, 1, 1, 1},
			res:  0,
		},
		{
			name: "3",
			nums: []int{1, 2, 3, 4, 5},
			res:  15,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sumOfUnique(c.nums))
		})
	}
}

func sumOfUnique(nums []int) int {
	counts := make(map[int]int, len(nums))
	for _, num := range nums {
		counts[num]++
	}

	var summ int
	for num, count := range counts {
		if count != 1 {
			continue
		}

		summ += num
	}

	return summ
}
