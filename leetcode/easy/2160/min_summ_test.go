package _2160_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  int
	}{
		{
			name: "1",
			num:  2932,
			res:  52,
		},
		{
			name: "2",
			num:  4009,
			res:  13,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minimumSum(c.num))
		})
	}
}

func minimumSum(num int) int {
	digs := make([]int, 0, 4)
	for ; num > 0; num /= 10 {
		digs = append(digs, num%10)
	}

	slices.Sort(digs)

	return (digs[0]+digs[1])*10 + digs[2] + digs[3]
}
