package _905_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{3, 1, 2, 4},
			res:  []int{2, 4, 3, 1},
		},
		{
			name: "2",
			nums: []int{0},
			res:  []int{0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, sortArrayByParity(c.nums))
		})
	}
}

func sortArrayByParity(nums []int) []int {
	res := make([]int, 0, len(nums))
	odd := make([]int, 0, len(nums)/2)

	for _, n := range nums {
		if n%2 == 0 {
			res = append(res, n)
			continue
		}

		odd = append(odd, n)
	}

	res = append(res, odd...)

	return res
}
