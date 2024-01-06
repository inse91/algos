package _448_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		res  []int
	}{
		{
			name: "1",
			in:   []int{4, 3, 2, 7, 8, 2, 3, 1},
			res:  []int{5, 6},
		},
		{
			name: "2",
			in:   []int{1, 1},
			res:  []int{2},
		},
		{
			name: "3",
			in:   []int{1, 1, 2},
			res:  []int{3},
		},
		{
			name: "4",
			in:   []int{1, 1, 2, 3},
			res:  []int{4},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findDisappearedNumbers(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func findDisappearedNumbers(nums []int) []int {
	m := make([]bool, len(nums))
	for _, num := range nums {
		m[num-1] = true
	}

	var res []int
	for i, b := range m {
		if b {
			continue
		}
		res = append(res, i+1)
	}

	return res
}
