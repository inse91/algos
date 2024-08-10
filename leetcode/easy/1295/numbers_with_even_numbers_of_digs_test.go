package _1295_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{12, 345, 2, 6, 7896},
			res:  2,
		},
		{
			name: "2",
			nums: []int{555, 901, 482, 1771},
			res:  1,
		},
		{
			name: "23",
			nums: []int{100000},
			res:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findNumbers(c.nums))
		})
	}
}

func findNumbers(nums []int) int {
	var res int
	for _, num := range nums {
		if check(num) {
			res++
		}
	}

	return res
}

func check(v int) bool {
	switch {
	case v < 10:
		return false
	case v < 100:
		return true
	case v < 1000:
		return false
	case v < 10000:
		return true
	case v < 100000:
		return false
	case v == 100000:
		return true
	default:
		panic("wrong num")
	}
}
