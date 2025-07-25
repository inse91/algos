package _2180_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-integers-with-even-digit-sum

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  int
	}{
		{
			name: "1",
			num:  4,
			res:  2,
		},
		{
			name: "2",
			num:  30,
			res:  14,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countEven(c.num))
		})
	}
}

func countEven(num int) int {
	var res int
	for n := 2; n <= num; n++ {
		nc := n
		var sum int
		for nc > 0 {
			sum += nc % 10
			nc /= 10
		}

		if sum%2 == 0 {
			res++
		}
	}

	return res
}
