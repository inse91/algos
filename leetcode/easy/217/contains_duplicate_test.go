package _217

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicates(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		res  bool
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 1},
			res:  true,
		},
		{
			name: "2",
			nums: []int{1, 2, 3, 4},
			res:  false,
		},
		{
			name: "3",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			res:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := containsDuplicate(tc.nums)
			assert.Equal(t, tc.res, res)
		})
	}
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}

	return false
}
