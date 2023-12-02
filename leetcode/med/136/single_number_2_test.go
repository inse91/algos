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
			in:   []int{2, 2, 3, 2},
			res:  3,
		},
		{
			name: "2",
			in:   []int{0, 1, 0, 1, 0, 1, 99},
			res:  99,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := singleNumber2(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}

}

func singleNumber2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := map[int]bool{}

	for _, num := range nums {
		v, ok := m[num]
		if !ok {
			m[num] = true
			continue
		}
		if v {
			m[num] = false
			continue
		}

		delete(m, num)
	}

	for num, _ := range m {
		return num
	}

	return 0
}
