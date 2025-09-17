package _3079_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-sum-of-encrypted-integers

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3},
			res:  6,
		},
		{
			name: "2",
			nums: []int{10, 21, 31},
			res:  66,
		},
		{
			name: "3",
			nums: []int{143,227,39},
			res:  1320,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sumOfEncryptedInt(c.nums))
		})
	}
}

func sumOfEncryptedInt(nums []int) int {
	var res int
	for _, num := range nums {
		res += decr(num)
	}

	return res
}

func decr(x int) int {
	if x < 10 {
		return x
	}

	var (
		c, mx int
	)
	for x > 0 {
		dig := x - x/10*10
		mx = max(mx, dig)

		x /= 10
		c++
	}

	mul := 1
	for range c {
		x += mx*mul
		mul *= 10
	}

	return x
}
