package _594_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestHarmSubseqTest(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 3, 2, 2, 5, 2, 3, 7},
			res:  5,
		},
		{
			name: "2",
			nums: []int{1, 1, 1, 1},
			res:  0,
		},
		{
			name: "3",
			nums: []int{1, 2, 3, 4},
			res:  2,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findLHS(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func findLHS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
			continue
		}
		m[num] = 1
	}

	var mx int
	for num, q := range m {
		if v, ok := m[num+1]; ok {
			mx = max(mx, v+q)
		}
	}

	return mx
}
