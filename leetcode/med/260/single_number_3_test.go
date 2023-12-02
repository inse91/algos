package _260

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		res  []int
	}{
		{
			name: "nil",
			in:   nil,
			res:  []int{},
		},
		{
			name: "1",
			in:   []int{1, 2, 1, 3, 2, 5},
			res:  []int{5, 3},
		},
		{
			name: "2",
			in:   []int{-1, 0},
			res:  []int{-1, 0},
		},
		{
			name: "3",
			in:   []int{0, 1},
			res:  []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := singleNumber(tc.in)
			assert.ElementsMatch(t, tc.res, res)
		})
	}

}

func singleNumber(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	m := map[int]struct{}{}

	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = struct{}{}
			continue
		}
		delete(m, num)
	}

	out := make([]int, 0, 2)
	for num, _ := range m {
		out = append(out, num)
	}

	return out
}
