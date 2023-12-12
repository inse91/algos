package _219

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicates2new(t *testing.T) {

	testCases := []struct {
		name string
		nums []int
		k    int
		res  bool
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 1},
			k:    3,
			res:  true,
		},
		{
			name: "2",
			nums: []int{1, 0, 1, 1},
			k:    1,
			res:  true,
		},
		{
			name: "3",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			res:  false,
		},
		{
			name: "4",
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    3,
			res:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := containsNearbyDuplicate(tc.nums, tc.k)
			assert.Equal(t, tc.res, res)
		})
	}
}

func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int]int)
	for i, num := range nums {
		if v, ok := m[num]; ok && abs(i, v) <= k {
			return true
		}
		m[num] = i
	}

	return false
}

func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}
