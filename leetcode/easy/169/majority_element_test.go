package _169

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		res  int
	}{
		{
			name: "1",
			in:   []int{3, 2, 3},
			res:  3,
		},
		{
			name: "2",
			in:   []int{2, 2, 1, 1, 1, 2, 2},
			res:  2,
		},
		{
			name: "4",
			in:   []int{1},
			res:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := majorityElement(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}
}

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 1
			continue
		}

		m[num]++
		if m[num] > len(nums)/2 {
			return num
		}
	}

	return 0
}
