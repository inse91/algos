package _136

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		res  int
	}{
		{
			name: "nil",
			in:   nil,
			res:  0,
		},
		{
			name: "1",
			in:   []int{2, 2, 1},
			res:  1,
		},
		{
			name: "2",
			in:   []int{4, 1, 2, 1, 2},
			res:  4,
		},
		{
			name: "3",
			in:   []int{1},
			res:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := singleNumber(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}

}

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := map[int]struct{}{}

	for i := range nums {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = struct{}{}
			continue
		}
		delete(m, nums[i])
	}

	for num, _ := range m {
		return num
	}

	return 0
}
