package _2562_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-array-concatenation-value

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int64
	}{
		{
			name: "1",
			nums: []int{7, 52, 2, 4},
			res:  596,
		},
		{
			name: "2",
			nums: []int{5, 14, 13, 8, 12},
			res:  673,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findTheArrayConcVal(c.nums))
		})
	}
}

func findTheArrayConcVal(nums []int) int64 {
	var res int64
	br := len(nums) / 2
	if len(nums)%2 == 1 {
		br++
	}

	for i := range nums {
		if i == br {
			break
		}
		left := nums[i]
		right := nums[len(nums)-1-i]

		switch {
		case i == len(nums)/2: // center
			res += int64(left)
			continue
		case right >= 1e4:
			left *= 100000
		case right >= 1e3:
			left *= 10000
		case right >= 1e2:
			left *= 1000
		case right >= 1e1:
			left *= 100
		default:
			left *= 10
		}

		res += int64(left + right)
	}

	return res
}
