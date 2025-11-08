package _1979_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-greatest-common-divisor-of-array

func Test1(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{2, 5, 6, 9, 10},
			res:  2,
		},
		{
			name: "2",
			nums: []int{7, 5, 6, 8, 3},
			res:  1,
		},
		{
			name: "3",
			nums: []int{3, 3},
			res:  3,
		},
		{
			name: "13",
			nums: []int{14, 10, 17, 25, 19},
			res:  5,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findGCD(c.nums))
		})
	}
}

func findGCD(nums []int) int {
	mx := nums[0]
	mn := mx
	for _, num := range nums {
		mx = max(mx, num)
		mn = min(mn, num)
	}

	if mx%mn == 0 {
		return mn
	}

	for div := 2; div <= mn; div++ {
		if mn%div != 0 {
			continue
		}

		pot := mn / div
		if mx%pot == 0 {
			return pot
		}

	}

	return 1
}
