package _645_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{1, 2, 2, 4},
			res:  []int{2, 3},
		},
		{
			name: "2",
			nums: []int{1, 1},
			res:  []int{1, 2},
		},
		{
			name: "3",
			nums: []int{2, 2},
			res:  []int{2, 1},
		},
		{
			name: "4",
			nums: []int{3, 2, 2},
			res:  []int{2, 1},
		},
		{
			name: "5",
			nums: []int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9},
			res:  []int{2, 10},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findErrorNums(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func findErrorNums(nums []int) []int {
	mx := nums[0]
	mn := nums[0]
	m := make(map[int]bool, len(nums))

	for _, n := range nums {
		mx = max(mx, n)
		mn = min(mn, n)

		v, ok := m[n]
		if !ok {
			m[n] = false
			continue
		}

		if !v {
			m[n] = true
		}
	}

	if mn == mx {
		if mn == 1 {
			return []int{1, 2}
		}
		return []int{2, 1}
	}

	var mm, nn int

	for i := 1; i <= mx; i++ {
		v, ok := m[i]
		if !ok {
			nn = i
		}
		if v {
			mm = i
		}
	}
	if nn == 0 {
		nn = mx + 1
	}

	return []int{mm, nn}
}
