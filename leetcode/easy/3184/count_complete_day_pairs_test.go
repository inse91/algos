package _3184_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		hrs  []int
		res  int
	}{
		{
			name: "1",
			hrs:  []int{12, 12, 30, 24, 24},
			res:  2,
		},
		{
			name: "2",
			hrs:  []int{72, 48, 24, 3},
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countCompleteDayPairs(c.hrs))
		})
	}
}

func countCompleteDayPairs(hours []int) int {
	var res int
	for i, h1 := range hours {
		for j := i + 1; j < len(hours); j++ {
			if (h1+hours[j])%24 == 0 {
				res++
			}
		}
	}

	return res
}
