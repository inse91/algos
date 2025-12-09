package _2748_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-beautiful-pairs

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{2, 5, 1, 4},
			res:  5,
		},
		{
			name: "11",
			nums: []int{2, 5, 1, 4, 2},
			res:  7,
		},
		{
			name: "2",
			nums: []int{11, 21, 12},
			res:  2,
		},
		{
			name: "17",
			nums: []int{11, 21, 12, 12},
			res:  4,
		},
		{
			name: "22",
			nums: []int{31, 25, 72, 79, 74},
			res:  7,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countBeautifulPairs(c.nums))
		})
	}
}

func countBeautifulPairs(nums []int) int {
	var c int
	for i, num1 := range nums {
		for _, num2 := range nums[i+1:] {
			fir := first(num1)
			las := last(num2)
			if check(fir, las) {
				c += 1
			}
		}
	}

	return c
}

func check(x, y int) bool {
	mx := min(x, y)
	for i := 2; i <= mx; i++ {
		if x%i != 0 {
			continue
		}

		if y%i != 0 {
			continue
		}

		return false
	}

	return true
}

func first(x int) int {
	if x < 10 {
		return x
	}

	if x < 100 {
		return x / 10
	}

	if x < 1000 {
		return x / 100
	}

	return x / 1000
}

func last(x int) int {
	return x % 10
}
